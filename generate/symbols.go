package generate

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/progrium/darwinkit/generate/declparse"
	"github.com/progrium/darwinkit/generate/modules"
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
	// Core fields (used by both old and new formats)
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

	// New fields from Apple's JSON format
	Abstract []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"abstract,omitempty"`
	Identifier struct {
		URL               string `json:"url"`
		InterfaceLanguage string `json:"interfaceLanguage"`
	} `json:"identifier,omitempty"`
	Metadata struct {
		Title       string `json:"title"`
		Role        string `json:"role"`
		RoleHeading string `json:"roleHeading"`
		Modules     []struct {
			Name string `json:"name"`
		} `json:"modules"`
		Platforms  []Platform `json:"platforms"`
		ExternalID string     `json:"externalID"`
	} `json:"metadata,omitempty"`
	References map[string]struct {
		Title      string `json:"title"`
		Type       string `json:"type"`
		Identifier string `json:"identifier"`
		Kind       string `json:"kind"`
		Role       string `json:"role"`
		URL        string `json:"url"`
	} `json:"references,omitempty"`
	PrimaryContentSections []struct {
		Kind         string `json:"kind"`
		Declarations []struct {
			Languages []string `json:"languages"`
			Tokens    []struct {
				Kind string `json:"kind"`
				Text string `json:"text"`
			} `json:"tokens"`
		} `json:"declarations"`
	} `json:"primaryContentSections,omitempty"`
	TopicSections []struct {
		Kind        string   `json:"kind"`
		Title       string   `json:"title"`
		Identifiers []string `json:"identifiers"`
	} `json:"topicSections,omitempty"`
}

type ContentFragment struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
	Code string `json:"code,omitempty"`
}

type Reference struct {
	Title      string            `json:"title,omitempty"`
	Type       string            `json:"type,omitempty"`
	Identifier string            `json:"identifier,omitempty"`
	URL        string            `json:"url,omitempty"`
	Kind       string            `json:"kind,omitempty"`
	Role       string            `json:"role,omitempty"`
	Abstract   []ContentFragment `json:"abstract,omitempty"`
}

type Platform struct {
	Name         string `json:"name"`
	IntroducedAt string `json:"introducedAt"`
	Current      string `json:"current"`
	Beta         bool   `json:"beta,omitempty"`
	Deprecated   bool   `json:"deprecated,omitempty"`
	DeprecatedAt string `json:"deprecatedAt,omitempty"`
}

type Parameter struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type,omitempty"`
	Optional    bool   `json:"optional,omitempty"`
}

func (s Symbol) DocURL() string {
	return fmt.Sprintf("https://developer.apple.com/documentation/%s?language=objc", s.Path)
}

func (s Symbol) HasFramework(name string) bool {
	name = strings.ReplaceAll(name, " ", "")
	for _, m := range s.Metadata.Modules {
		m.Name = strings.ReplaceAll(m.Name, " ", "")
		if strings.EqualFold(m.Name, name) {
			return true
		}
	}
	return false
}

func (s Symbol) MainModule() *modules.Module {
	if s.Name == "IOSurface" {
		return modules.Get("iosurface")
	}
	var moduleNames []string
	for _, m := range s.Metadata.Modules {
		moduleNames = append(moduleNames, m.Name)
	}
	if len(moduleNames) == 0 {
		return nil
	}
	sort.Strings(moduleNames)
	defer func() {
		if r := recover(); r != nil {
			if strings.Contains(r.(string), "module not found") {
				if !modules.CanIgnoreNotFound(r) {
					panic(r)
				}
			}
		}
	}()
	mod := modules.Get(moduleNames[0])
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

// SymbolCache represents a cache of symbols loaded from Apple API docs
type SymbolCache struct {
	basePath string
	cache    map[string]Symbol
	all      []Symbol
	modSyms  map[string][]Symbol
}

// OpenSymbols creates a new SymbolCache for the given API docs directory
func OpenSymbols(basePath string) (*SymbolCache, error) {
	// Verify the directory exists
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("API docs directory not found: %s", basePath)
	}

	return &SymbolCache{
		basePath: basePath,
		cache:    make(map[string]Symbol),
		modSyms:  make(map[string][]Symbol),
	}, nil
}

// loadFromFile loads a symbol from a JSON file
func (db *SymbolCache) loadFromFile(filePath string) (v Symbol, err error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return v, err
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return v, err
	}

	// Process the symbol
	if v.Metadata.Title != "" {
		// Extract core fields from new format
		v.Name = v.Metadata.Title
		v.Path = strings.TrimPrefix(v.Identifier.URL, "doc://com.apple.documentation/documentation/")
		v.Kind = v.Metadata.RoleHeading

		// Extract modules
		var modules []string
		for _, m := range v.Metadata.Modules {
			modules = append(modules, m.Name)
		}
		v.Modules = modules

		// Extract description from abstract
		var description strings.Builder
		for _, ab := range v.Abstract {
			if ab.Type == "text" {
				description.WriteString(ab.Text)
			} else if ab.Type == "codeVoice" {
				description.WriteString("`" + ab.Text + "`")
			}
		}
		v.Description = description.String()

		// Extract declaration from primary content sections
		for _, section := range v.PrimaryContentSections {
			if section.Kind == "declarations" {
				for _, decl := range section.Declarations {
					if contains(decl.Languages, "occ") {
						var declaration strings.Builder
						
						// First build the full declaration string
						for _, token := range decl.Tokens {
							declaration.WriteString(token.Text)
						}
						
						// Convert Apple's token format to our token format
						var tokens []Token
						for _, token := range decl.Tokens {
							tokens = append(tokens, Token{
								Kind: mapTokenKind(token.Kind),
								Text: token.Text,
							})
						}

						v.Declaration = strings.TrimSpace(declaration.String())
						
						// Use the token parser to extract structured information
						if parsed, err := ParseTokens(tokens); err == nil {
							if parsed.Kind != "" && parsed.Kind != "Unknown" {
								v.Kind = parsed.Kind
							}
							if parsed.ReturnType != "" {
								v.Return = parsed.ReturnType
							}
							
							if len(parsed.Parameters) > 0 {
								var parameters []Parameter
								for _, param := range parsed.Parameters {
									p := Parameter{
										Name:     param.Name,
										Type:     param.Type,
										Optional: param.Nullable,
									}
									parameters = append(parameters, p)
								}
								v.Parameters = parameters
							}
						} else {
							// Fallback to the old manual parsing approach if token parser fails
							var returnType strings.Builder
							var inReturnType bool
							var parameters []Parameter
							var currentParam Parameter
							
							for _, token := range decl.Tokens {
								switch token.Kind {
								case "typeIdentifier", "identifier", "text", "number", "keyword":
									if inReturnType {
										returnType.WriteString(token.Text)
									}
								case "parameter":
									currentParam = Parameter{Name: token.Text}
									parameters = append(parameters, currentParam)
								case "typeParameter":
									if len(parameters) > 0 {
										parameters[len(parameters)-1].Type = token.Text
									}
								case "attribute":
									if token.Text == "nullable" {
										if len(parameters) > 0 {
											parameters[len(parameters)-1].Optional = true
										}
									}
								case "returnArrow":
									inReturnType = true
								}
							}
							
							if returnType.Len() > 0 {
								v.Return = strings.TrimSpace(returnType.String())
							}
							if len(parameters) > 0 {
								v.Parameters = parameters
							}
						}
						break
					}
				}
			}
		}
	}

	if strInList(blacklist, v.Name) {
		return v, fmt.Errorf("blacklisted symbol: %s", v.Name)
	}
	if strInList(pathBlacklist, v.Path) {
		return v, fmt.Errorf("blacklisted path: %s", v.Path)
	}
	if v.Kind != "Property" && v.Kind != "Method" && v.Kind != "Framework" {
		db.cache[v.Name] = v
	}
	return
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// mapTokenKind maps Apple's token kinds to our TokenKind constants
func mapTokenKind(kind string) string {
	switch kind {
	case "identifier", "typeIdentifier":
		return TokenIdentifier
	case "keyword":
		return TokenKeyword
	case "operator":
		return TokenOperator
	case "text", "punctuation", "attribute", "parameter", "typeParameter", "returnArrow":
		return TokenPunctuation
	case "number":
		return TokenNumber
	case "string":
		return TokenString
	default:
		return TokenIdentifier // Default to identifier
	}
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

	// Walk through all files in the base directory
	modulesDirs, err := os.ReadDir(db.basePath)
	if err != nil {
		return
	}

	for _, moduleDir := range modulesDirs {
		if !moduleDir.IsDir() {
			continue
		}

		modulePath := filepath.Join(db.basePath, moduleDir.Name())
		files, err := os.ReadDir(modulePath)
		if err != nil {
			continue
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			// Skip overview.json as it doesn't contain specific symbols
			if file.Name() == "overview.json" {
				continue
			}

			filePath := filepath.Join(modulePath, file.Name())
			s, err := db.loadFromFile(filePath)
			if err != nil {
				continue
			}
			symbols = append(symbols, s)
		}
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

	// Check if the module directory exists
	moduleDir := filepath.Join(db.basePath, strings.ToLower(m.Name))
	
	// If directory exists, read all JSON files in it
	if fi, err := os.Stat(moduleDir); err == nil && fi.IsDir() {
		files, err := os.ReadDir(moduleDir)
		if err == nil {
			for _, file := range files {
				if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
					continue
				}
				
				// Skip overview.json as it doesn't contain specific symbols
				if file.Name() == "overview.json" {
					continue
				}
				
				filePath := filepath.Join(moduleDir, file.Name())
				s, err := db.loadFromFile(filePath)
				if err != nil {
					continue
				}
				symbols = append(symbols, s)
			}
		}
	}

	// Check for overview.json to process references
	overviewPath := filepath.Join(moduleDir, "overview.json")
	if _, err := os.Stat(overviewPath); err == nil {
		overview, err := db.loadFromFile(overviewPath)
		if err == nil {
			// Process references from overview
			for _, ref := range overview.References {
				if ref.Kind == "symbol" && ref.Role == "symbol" {
					symbolPath := strings.TrimPrefix(ref.URL, "/documentation/"+strings.ToLower(m.Name)+"/")
					symbolFilePath := filepath.Join(moduleDir, symbolPath+".json")
					
					if _, err := os.Stat(symbolFilePath); err == nil {
						s, err := db.loadFromFile(symbolFilePath)
						if err != nil {
							continue
						}
						// Check if we already have this symbol
						duplicate := false
						for _, existing := range symbols {
							if existing.Name == s.Name {
								duplicate = true
								break
							}
						}
						if !duplicate {
							symbols = append(symbols, s)
						}
					}
				}
			}
		}
	}

	// Handle special case for AppKit
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

// strInList checks if a string is in a list of strings
func strInList(list []string, s string) bool {
	for _, i := range list {
		if i == s {
			return true
		}
	}
	return false
}