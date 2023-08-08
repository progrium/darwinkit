package generate

import (
	"log"
	"strings"

	"github.com/progrium/macdriver/generate/codegen"
	"github.com/progrium/macdriver/generate/declparse"
	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/generate/oldgen"
	"github.com/progrium/macdriver/generate/typing"
	"github.com/progrium/macdriver/internal/set"
)

type Generator struct {
	*SymbolCache
	Platform  string
	Version   int
	enumCache map[string][]Symbol
}

func (db *Generator) ResolveTypeAlias(typeName string) (declparse.TypeInfo, bool) {
	s := db.FindByName(typeName)
	if s == nil {
		return declparse.TypeInfo{}, false
	}
	st, err := s.Parse()
	if err != nil || st.TypeAlias == nil {
		return declparse.TypeInfo{}, false
	}
	return *st.TypeAlias, true
}

func (db *Generator) TypeFromSymbolName(name string) typing.Type {
	s := db.FindByName(name)
	if s == nil {
		log.Panicf("symbol name not found: %s", name)
	}
	return db.TypeFromSymbol(*s)
}

func (db *Generator) TypeFromSymbol(sym Symbol) typing.Type {
	module := sym.Modules[0]
	switch sym.Kind {
	case "Class":
		return &typing.ClassType{
			Name:   sym.Name,
			GName:  modules.TrimPrefix(sym.Name),
			Module: modules.Get(module),
		}
	case "Protocol":
		return &typing.ProtocolType{
			Name:   sym.Name,
			GName:  modules.TrimPrefix(sym.Name),
			Module: modules.Get(module),
		}
	case "Enum":
		st, err := sym.Parse()
		if err != nil || st.Enum == nil {
			panic("invalid enum")
		}
		return &typing.AliasType{
			Name:   sym.Name,
			GName:  modules.TrimPrefix(sym.Name),
			Module: modules.Get(module),
			Type:   db.ParseType(st.Enum.Type),
		}
	case "Type":
		if sym.Type != "Type Alias" {
			panic("unknown type")
		}
		st, err := sym.Parse()
		if err != nil || st.TypeAlias == nil {
			panic("invalid type alias")
		}
		return &typing.AliasType{
			Name:   sym.Name,
			GName:  modules.TrimPrefix(sym.Name),
			Module: modules.Get(module),
			Type:   db.ParseType(*st.TypeAlias),
		}
	case "Struct":
		return &typing.StructType{
			Name:   sym.Name,
			GName:  modules.TrimPrefix(sym.Name),
			Module: modules.Get(module),
		}
	default:
		panic("unknown type")
	}

}

func (db *Generator) ParseType(ti declparse.TypeInfo) (typ typing.Type) {
	ref := false
	switch ti.Name {
	case "SEL":
		typ = &typing.SelectorType{}
	case "void":
		typ = &typing.VoidType{}
	case "instance":
		typ = &typing.InstanceType{}
	case "id":
		typ = &typing.IDType{}
		if len(ti.Params) > 0 {
			typ = db.TypeFromSymbolName(ti.Params[0].Name).(*typing.ProtocolType)
		}
	case "CGFloat":
		typ = typing.Float
	case "NSString":
		typ = &typing.StringType{}
		ref = true
	case "NSStringNeedNil":
		typ = &typing.StringType{NeedNil: true}
		ref = true
	case "NSData":
		typ = &typing.DataType{}
		ref = true
	case "NSArray":
		at := &typing.ArrayType{}
		typ = at
		if len(ti.Params) == 1 {
			at.Type = db.ParseType(ti.Params[0])
		} else {
			at.Type = typing.Object
		}
		ref = true
	case "NSDictionary":
		dt := &typing.DictType{}
		typ = dt
		if len(ti.Params) == 2 {
			dt.KeyType = db.ParseType(ti.Params[0])
			dt.ValueType = db.ParseType(ti.Params[1])
		} else {
			dt.KeyType = typing.Object
			dt.ValueType = typing.Object
		}
		ref = true
	default:
		pt, ok := typing.GetPrimitiveType(ti.Name)
		if ok {
			typ = pt
		} else {
			typ = db.TypeFromSymbolName(ti.Name)
			switch typ.(type) {
			case *typing.ClassType:
				ref = true
			case *typing.ProtocolType:
				panic("standalone proto type")
			}
		}
	}

	if ti.IsPtr && !ref {
		if _, ok := typ.(*typing.VoidType); ok {
			typ = &typing.VoidPointerType{}
		} else {
			typ = &typing.PointerType{
				Type: typ,
			}
		}
	}

	// TODO: handle block
	// currentType = &typing.BlockType{
	// 	ReturnType: currentType,
	// 	Params:     blockParams,
	// }

	return
}

func (db *Generator) ToEnumInfo(fw string, sym Symbol) *codegen.EnumInfo {
	at := db.TypeFromSymbol(sym).(*typing.AliasType)
	if db.Platform == "macos" && fw == "UIKit" {
		fw = "AppKit"
	}
	var enumValues []*codegen.EnumValue
	for _, ev := range db.EnumValues(fw) {
		st, _ := ev.Parse()
		if st.Variable.Type.Name != sym.Name && ev.Parent != sym.Name {
			continue
		}
		if strIn([]string{
			"NS16BitBigEndianBitmapFormat",
			"NS16BitLittleEndianBitmapFormat",
			"NS32BitBigEndianBitmapFormat",
			"NS32BitLittleEndianBitmapFormat",
		}, ev.Name) {
			continue
		}
		c := modules.LookupConstant(db.Platform, fw, ev.Name)
		if c == nil {
			continue
		}
		enumValues = append(enumValues, &codegen.EnumValue{
			Name:       ev.Name,
			GoName:     modules.TrimPrefix(ev.Name),
			Value:      c.Value,
			Arm64Value: c.ArmValue,
			Module:     modules.Get(fw),
		})
	}
	ei := &codegen.EnumInfo{
		AliasType: *at,
		Values:    enumValues,
	}
	ei.Module = modules.Get(fw) // main to convert uikit to appkit
	ei.GName = modules.TrimPrefix(sym.Name)
	return ei
}

func (db *Generator) EnumValues(framework string) (symbols []Symbol) {
	if db.enumCache == nil {
		db.enumCache = make(map[string][]Symbol)
	}
	if cached, ok := db.enumCache[framework]; ok {
		return cached
	}
	seen := map[string]bool{}
	if db.Platform == "macos" && framework == "appkit" {
		for _, s := range db.EnumValues("uikit") {
			if _, exists := seen[s.Name]; !exists {
				seen[s.Name] = true
				symbols = append(symbols, s)
			}
		}
	}
	for _, s := range db.ModuleSymbols(framework) {
		if db.Platform == "macos" && framework == "uikit" {
			if !s.HasFramework("appkit") {
				continue
			}
		}
		if !s.HasPlatform(db.Platform, db.Version, true) {
			continue
		}
		if s.Kind != "Constant" {
			continue
		}
		if s.Type != "Enumeration Case" && s.Type != "Global Variable" {
			continue
		}
		if _, exists := seen[s.Name]; !exists {
			seen[s.Name] = true
			symbols = append(symbols, s)
		}
	}

	db.enumCache[framework] = symbols
	return
}

func (db *Generator) Generate(platform string, version int, rootDir string, framework string, ignoreTypes set.Set[string]) {
	db.Platform = platform
	db.Version = version
	module := modules.Get(framework)
	mw := &codegen.ModuleWriter{
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
			// classGen := getClassGen(tti)
			// classGen.Init()
			// fw := &gen.FileWriter{
			// 	Name:        tti.Name,
			// 	Module:      *classGen.Type.Module,
			// 	PlatformDir: rootDir,
			// }
			// fw.Add(classGen)
			// fw.WriteCode()
		case "Protocol":
			// protocolGen := getProtocolGen(tti)
			// protocolGen.Init()
			// fw := &gen.FileWriter{
			// 	Name:        tti.Name,
			// 	Module:      *protocolGen.Type.Module,
			// 	PlatformDir: rootDir,
			// }
			// fw.Add(protocolGen)
			// fw.WriteCode()
			// mw.Protocols = append(mw.Protocols, protocolGen.Type)
		case "Enum":
			mw.Aliases = append(mw.Aliases, db.ToEnumInfo(framework, s))
		case "Type":
			if db.AllowedEnumAlias(s) {
				mw.Aliases = append(mw.Aliases, db.ToEnumInfo(framework, s))
			}
		}
	}
	// mw.WriteCode()
	mw.WriteAliasesCode()
	oldgen.FormatCode(rootDir + "/" + module.Package)
}

func (db *Generator) AllowedEnumAlias(s Symbol) bool {
	if s.Type != "Type Alias" {
		return false
	}
	allowed := []string{
		"size_t",
		"uint8_t",
		"uint16_t",
		"uint32_t",
		"uint64_t",
		"int8_t",
		"int16_t",
		"int32_t",
		"int64_t",
		"char",
		"long",
		"int",
		"short",
		"float",
		"double",
		"NSUInteger",
		"NSString",
		"NSInteger",
		"CGFloat",
		"SInt8",
		"SInt16",
		"SInt32",
		"SInt64",
		"UInt8",
		"UInt16",
		"UInt32",
		"UInt64",
	}
	st, err := s.Parse()
	if err != nil {
		return false
	}
	if st.TypeAlias != nil && st.TypeAlias.Name != "" {
		if strIn(allowed, st.TypeAlias.Name) {
			return true
		} else {
			ti, ok := db.ResolveTypeAlias(st.TypeAlias.Name)
			if ok && strIn(allowed, ti.Name) {
				return true
			}
		}
	}
	return false
}

func strIn(slice []string, str string) bool {
	for _, s := range slice {
		if strings.EqualFold(str, s) {
			return true
		}
	}
	return false
}
