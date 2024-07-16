package generate

import (
	"fmt"
	"log"
	"strings"

	"github.com/progrium/darwinkit/generate/codegen"
	"github.com/progrium/darwinkit/generate/declparse"
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/generate/typing"
	"github.com/progrium/darwinkit/internal/set"
)

type Generator struct {
	*SymbolCache
	Platform      string
	Version       int
	Framework     string
	customTypes   set.Set[string]
	customMethods set.Set[string]
	enumCache     map[string][]Symbol
	genCache      map[string]codegen.CodeGen
}

// Generate generates the code for the given platform, version, and framework
func (db *Generator) Generate(platform string, version int, rootDir string, framework string, ignoreTypes set.Set[string]) {
	db.Platform = platform
	db.Version = version
	module := modules.Get(framework)
	db.Framework = module.Package
	modsym := db.ModuleSymbol(*module)

	types, methods := exportedSourceSymbols(fmt.Sprintf("%s/%s/%s_custom.go", rootDir, db.Framework, db.Framework))
	db.customTypes = set.New(types...)
	db.customMethods = set.New(methods...)

	RemoveGeneratedCode(fmt.Sprintf("%s/%s", rootDir, db.Framework))

	mw := &codegen.ModuleWriter{
		Description: modsym.Description,
		DocURL:      modsym.DocURL(),
		Module:      *module,
		PlatformDir: rootDir,
	}
	for _, s := range db.ModuleSymbols(framework) {
		if ignoreTypes.Contains(s.Name) {
			continue
		}
		if !s.HasPlatform(platform, version, true) {
			continue
		}
		switch s.Kind {
		case "Class":
			if db.customTypes.Contains(modules.TrimPrefix(s.Name)) {
				log.Println("skipping class with custom definition", s.Name)
				continue
			}
			classGen := db.ToClassGen(s)
			if classGen == nil {
				log.Println("skipping class", s.Name)
				continue
			}
			classGen.Init()
			fw := &codegen.FileWriter{
				Name:        s.Name,
				Module:      *classGen.Type.Module,
				PlatformDir: rootDir,
			}
			fw.Add(classGen)
			fw.WriteCode()
		case "Protocol":
			protocolGen := db.ToProtocolGen(framework, s)
			if protocolGen == nil {
				log.Println("skipping protocol", s.Name)
				continue
			}
			if db.customTypes.Contains("P" + modules.TrimPrefix(s.Name)) {
				log.Println("skipping protocol interface with custom definition", s.Name)
				protocolGen.SkipInterface = true
			}
			if db.customTypes.Contains(modules.TrimPrefix(s.Name) + "Wrapper") {
				log.Println("skipping protocol wrapper with custom definition", s.Name)
				protocolGen.SkipWrapper = true
			}
			if db.customTypes.Contains(modules.TrimPrefix(s.Name)) {
				log.Println("skipping protocol delegate with custom definition", s.Name)
				protocolGen.SkipDelegate = true
			}
			protocolGen.Init()
			fw := &codegen.FileWriter{
				Name:        s.Name + "Protocol",
				Module:      *protocolGen.Type.Module,
				PlatformDir: rootDir,
			}
			fw.Add(protocolGen)
			fw.WriteCode()
			mw.Protocols = append(mw.Protocols, protocolGen.Type)
		case "Enum":
			if s.Name == "MTLArgumentAccess" {
				// todo: regen metal without this skip to see why
				continue
			}
			if db.customTypes.Contains(modules.TrimPrefix(s.Name)) {
				log.Println("skipping enum with custom definition", s.Name)
				continue
			}
			mw.EnumAliases = append(mw.EnumAliases, db.ToEnumInfo(framework, s))
		case "Type":
			if db.customTypes.Contains(modules.TrimPrefix(s.Name)) {
				log.Println("skipping type with custom definition", s.Name)
				continue
			}
			if db.AllowedEnumAlias(s) {
				mw.EnumAliases = append(mw.EnumAliases, db.ToEnumInfo(framework, s))
				continue
			}
			st, err := s.Parse(db.Platform)
			if err != nil || st.TypeAlias == nil {
				log.Printf("skipping '%s', bad decl: %s", s.Name, s.Declaration)
				continue
			}
			if st.TypeAlias.Func != nil {
				ftyp := db.ParseType(*st.TypeAlias)
				if ftyp == nil {
					fmt.Printf("skipping: name=%s declaration=%s\n", s.Name, s.Declaration)
					continue
				}
				mw.FuncAliases = append(mw.FuncAliases, &codegen.AliasInfo{
					AliasType: typing.AliasType{
						Name:  s.Name,
						GName: modules.TrimPrefix(s.Name),
						Type:  db.ParseType(*st.TypeAlias),
					},
					Description: s.Description,
					DocURL:      s.DocURL(),
				})
				continue
			}

			if st.TypeAlias.Annots[declparse.TypeAnnotStruct] {
				mw.StructAliases = append(mw.StructAliases, &codegen.AliasInfo{
					AliasType: typing.AliasType{
						Name:  s.Name,
						GName: modules.TrimPrefix(s.Name),
						Type:  db.ParseType(*st.TypeAlias),
					},
					Description: s.Description,
					DocURL:      s.DocURL(),
				})
				continue
			}
		case "Function":
			fn := db.ToFunction(framework, s)
			if fn == nil {
				continue
			}
			mw.Functions = append(mw.Functions, fn)
		case "Struct":
			s := db.ToStruct(framework, s)
			if s == nil {
				continue
			}
			mw.Structs = append(mw.Structs, s)
		}
	}
	mw.WriteCode()
	FormatCode(rootDir + "/" + module.Package)
}

func (db *Generator) ResolveTypeAlias(typeName string) (declparse.TypeInfo, bool) {
	s := db.FindTypeSymbol(typeName)
	if s == nil {
		return declparse.TypeInfo{}, false
	}
	st, err := s.Parse(db.Platform)
	if err != nil || st.TypeAlias == nil {
		return declparse.TypeInfo{}, false
	}
	return *st.TypeAlias, true
}

func (db *Generator) TypeFromSymbolName(name string) typing.Type {
	// hardcoded for now
	if db.Platform == "macos" && strings.HasPrefix(name, "UI") && !strings.HasPrefix(name, "UInt") {
		name = "NS" + strings.TrimPrefix(name, "UI")
	}
	s := db.FindTypeSymbol(name)
	if s == nil {

		// __ prefixed symbols are generally structs
		if strings.HasPrefix(name, "__") {
			s := &typing.StructType{
				Name:   name,
				GName:  modules.TrimPrefix(name),
				Module: modules.Get(db.Framework),
			}
			log.Printf("using Struct for unknown symbol: %s\n", name)
			return s
		}
		log.Printf("using NSObject for unknown symbol: %s\n", name)
		return typing.Object
	}
	if m := s.MainModule(); m != nil && modules.CanAbstractModuleCoupling(db.Framework, m.Package) {
		log.Printf("using NSObject for '%s' to decouple '%s'\n", name, m.Package)
		return typing.Object
	}

	return db.TypeFromSymbol(*s)
}

func strIn(slice []string, str string) bool {
	for _, s := range slice {
		if strings.EqualFold(str, s) {
			return true
		}
	}
	return false
}
