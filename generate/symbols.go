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

func (s Symbol) HasPlatform(name string, version int) bool {
	for _, p := range s.Platforms {
		if strings.EqualFold(p.Name, name) {
			if p.Deprecated {
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

var symbolCache = map[string]Symbol{}

func FindSymbolByName(db *zip.ReadCloser, name string) *Symbol {
	if s, ok := symbolCache[name]; ok {
		return &s
	}
	var found *Symbol
	for _, file := range db.File {
		if strings.Contains(file.Name, strings.ToLower(name)) {
			s, err := LoadSymbolFrom(file)
			if err != nil {
				continue
			}
			if s.Name == name {
				found = &s
			}
		}
	}
	if found != nil {
		symbolCache[name] = *found
	}
	return found
}

func LoadSymbolFrom(file *zip.File) (v Symbol, err error) {
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
	symbolCache[v.Name] = v
	return
}
