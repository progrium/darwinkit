package generate

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/progrium/macdriver/generate/declparse"
)

type Symbol struct {
	Name string
	Path string
	Kind string

	Description string
	Type        string
	Parent      string
	Modules     []string
	Platforms   []Platform
	Declaration string
	Parameters  []Parameter
	Return      string
	Deprecated  bool
}

type Platform struct {
	Name         string
	IntroducedAt string
	Current      string
	Beta         bool
	Deprecated   bool
	DeprecatedAt string
}

type Parameter struct {
	Name        string
	Description string
}

func (s Symbol) DocURL() string {
	return fmt.Sprintf("https://developer.apple.com/documentation/%s?language=objc", s.Path)
}

func (s Symbol) HasFramework(name string) bool {
	for _, m := range s.Modules {
		if strings.EqualFold(m, name) {
			return true
		}
	}
	return false
}

func (s Symbol) HasPlatform(name string, version int, deprecated bool) bool {
	for _, p := range s.Platforms {
		if strings.EqualFold(p.Name, name) {
			if !deprecated && p.Deprecated {
				return false
			}
			ver := strings.Split(p.IntroducedAt, ".")
			major, _ := strconv.Atoi(ver[0])
			minor, _ := strconv.Atoi(ver[1])
			if version > major || (version == major && minor == 0) {
				return true
			}
			return false
		}
	}
	return false
}

func (s Symbol) Parse() (*declparse.Statement, error) {
	p := declparse.NewStringParser(s.Declaration)
	switch s.Kind {
	case "Constant", "Property":
		p.Hint = declparse.HintVariable
	case "Function":
		p.Hint = declparse.HintFunction
	default:
	}
	if s.Type == "Enumeration Case" {
		p.Hint = declparse.HintEnumCase
	}
	return p.Parse()
}

type SymbolCache struct {
	*zip.ReadCloser
	cache map[string]Symbol
}

func OpenSymbols(filename string) (*SymbolCache, error) {
	db, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	return &SymbolCache{
		ReadCloser: db,
		cache:      make(map[string]Symbol),
	}, nil
}

func (db *SymbolCache) loadFrom(file *zip.File) (v Symbol, err error) {
	var reader io.ReadCloser
	reader, err = file.Open()
	if err != nil {
		return v, err
	}
	defer reader.Close()

	b, err := io.ReadAll(reader)
	if err != nil {
		return
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return v, err
	}
	db.cache[v.Name] = v
	return
}

func (db *SymbolCache) FindByName(name string) *Symbol {
	if s, ok := db.cache[name]; ok {
		return &s
	}
	var found *Symbol
	for _, file := range db.File {
		if strings.Contains(file.Name, strings.ToLower(name)) {
			s, err := db.loadFrom(file)
			if err != nil {
				continue
			}
			if strings.EqualFold(s.Name, name) {
				found = &s
			}
		}
	}
	if found != nil {
		db.cache[name] = *found
	}
	return found
}

func (db *SymbolCache) AllSymbols() (symbols []Symbol) {
	for _, file := range db.File {
		if file.FileInfo().IsDir() {
			continue
		}
		s, err := db.loadFrom(file)
		if err != nil {
			continue
		}
		symbols = append(symbols, s)
	}
	return
}

func (db *SymbolCache) ModuleSymbols(module string) (symbols []Symbol) {
	for _, file := range db.File {
		if !file.FileInfo().IsDir() && strings.HasPrefix(file.Name, "symbols/"+strings.ToLower(module)) {
			s, err := db.loadFrom(file)
			if err != nil {
				continue
			}
			symbols = append(symbols, s)
		}
	}
	if module == "appkit" {
		for _, s := range db.ModuleSymbols("uikit") {
			if s.HasFramework("appkit") {
				symbols = append(symbols, s)
			}
		}
	}
	return
}
