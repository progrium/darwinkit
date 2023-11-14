package generate

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/progrium/macdriver/generate/declparse"
	"github.com/progrium/macdriver/generate/modules"
)

var blacklist = []string{
	"CMPersistentTrackID",                      // type alias, enum overlap. causes infinite loop
	"CMSubtitleFormatType",                     // type alias, enum overlap. causes infinite loop
	"CMAttachmentMode",                         // type alias, enum overlap. causes infinite loop
	"CMMetadataFormatType",                     // type alias, enum overlap. causes infinite loop
	"CMMediaType",                              // type alias, enum overlap. causes infinite loop
	"CMAudioCodecType",                         // type alias, enum overlap. causes infinite loop
	"CMTimeCodeFormatType",                     // type alias, enum overlap. causes infinite loop
	"CMMuxedStreamType",                        // type alias, enum overlap. causes infinite loop
	"AUEventSampleTime",                        // type alias, enum overlap. causes infinite loop
	"os_workgroup_t",                           // os module struct ref thing
	"NSFileProviderItem",                       // type alias, protocol overlap
	"NSItemProviderReading",                    // protocol with class methods
	"NSSecureCoding",                           // protocol with class methods
	"NSProxy",                                  // complicated
	"NSDistantObject",                          // based on proxy
	"NSProtocolChecker",                        // based on proxy
	"AUInternalRenderBlock",                    // "not supported pointer to: AURenderEvent"
	"CGDataProviderGetBytesAtPositionCallback", // uses a weird kernel type not sure how to handle
	"off_t",
	"NSAccessibilityColor", // "cannot find protocol declaration" when added to protocols.gen.m
	"WebView",              // gets picked up instead of WKWebView
}

var pathBlacklist = []string{
	"foundation/nshashtable/legacy_hash_table_implementation/nshashtable", // found instead of NSHashTable class
	"foundation/nsmaptable/legacy_map_table_implementation/nsmaptable",    // same
}

type Symbol struct {
	Name string
	Path string
	Kind string // Class, Constant, Enum, Framework, Function, Macro, Method, Property, Protocol, Struct, Type

	Description   string
	Type          string
	Parent        string
	Modules       []string
	Platforms     []Platform
	Declaration   string
	Declarations  map[string]string
	Parameters    []Parameter
	Return        string
	Deprecated    bool
	InheritedFrom string
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
	name = strings.ReplaceAll(name, " ", "")
	for _, m := range s.Modules {
		m = strings.ReplaceAll(m, " ", "")
		if strings.EqualFold(m, name) {
			return true
		}
	}
	return false
}

func (s Symbol) MainModule() *modules.Module {
	if s.Name == "IOSurface" {
		return modules.Get("iosurface")
	}
	if len(s.Modules) == 0 {
		return nil
	}
	sort.Strings(s.Modules)
	defer func() {
		if r := recover(); r != nil {
			if strings.Contains(r.(string), "module not found") {
				if !modules.CanIgnoreNotFound(r) {
					panic(r)
				}
			}
		}
	}()
	mod := modules.Get(s.Modules[0])
	if strings.HasPrefix(s.Name, "CG") && mod.Package == "corefoundation" {
		// lets just normalize CG symbols under corefoundation to coregraphics
		mod = modules.Get("coregraphics")
	}
	return mod
}

func (s Symbol) HasPlatform(name string, version int, deprecated bool) bool {
	for _, p := range s.Platforms {
		if strings.EqualFold(p.Name, name) {
			if !deprecated && p.Deprecated {
				return false
			}
			ver := strings.Split(p.IntroducedAt, ".")
			major, _ := strconv.Atoi(ver[0])
			//minor, _ := strconv.Atoi(ver[1])
			if version >= major {
				return true
			}
			return false
		}
	}
	return false
}

func (s Symbol) Parse(platform string) (*declparse.Statement, error) {
	decl := s.Declaration
	if decl == "" && len(s.Declarations) > 0 {
		decl = s.Declarations[platform]
	}
	p := declparse.NewStringParser(decl)
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
	cache   map[string]Symbol
	all     []Symbol
	modSyms map[string][]Symbol
}

func OpenSymbols(filename string) (*SymbolCache, error) {
	db, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	return &SymbolCache{
		ReadCloser: db,
		cache:      make(map[string]Symbol),
		modSyms:    make(map[string][]Symbol),
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
	if strIn(blacklist, v.Name) {
		return v, fmt.Errorf("blacklisted symbol: %s", v.Name)
	}
	if strIn(pathBlacklist, v.Path) {
		return v, fmt.Errorf("blacklisted path: %s", v.Path)
	}
	if v.Kind != "Property" && v.Kind != "Method" && v.Kind != "Framework" {
		db.cache[v.Name] = v
	}
	return
}

func (db *SymbolCache) ModuleSymbol(m modules.Module) *Symbol {
	for _, s := range db.AllSymbols() {
		if s.Kind != "Framework" {
			continue
		}
		if strings.EqualFold(s.Path, m.Name) {
			fs := s
			return &fs
		}
	}
	return nil
}

func (db *SymbolCache) FindTypeSymbol(name string) *Symbol {
	if s, ok := db.cache[name]; ok {
		return &s
	}
	var found *Symbol
	for _, s := range db.AllSymbols() {
		if s.Kind == "Property" || s.Kind == "Method" || s.Kind == "Framework" {
			continue
		}
		if strings.EqualFold(s.Name, name) {
			found = &s
			break
		}
	}
	if found != nil {
		db.cache[found.Name] = *found
	}
	return found
}

func (db *SymbolCache) AllSymbols() (symbols []Symbol) {
	if len(db.all) > 0 {
		return db.all
	}
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
	db.all = symbols
	return
}

func (db *SymbolCache) ModuleSymbols(module string) (symbols []Symbol) {
	m := modules.Get(module)
	if m == nil {
		return
	}
	if s, ok := db.modSyms[m.Name]; ok {
		return s
	}
	for _, file := range db.File {
		if !file.FileInfo().IsDir() &&
			(strings.HasPrefix(file.Name, fmt.Sprintf("symbols/%s/", strings.ToLower(m.Name))) ||
				file.Name == fmt.Sprintf("symbols/%s.json", strings.ToLower(m.Name))) {
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
	db.modSyms[m.Name] = symbols
	return
}
