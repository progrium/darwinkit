package generate

import (
	"strings"

	"github.com/progrium/macdriver/generate/codegen"
	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/generate/typing"
	"github.com/progrium/macdriver/internal/stringx"
)

func (db *Generator) ToEnumInfo(fw string, sym Symbol) *codegen.AliasInfo {
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
	ei := &codegen.AliasInfo{
		AliasType:   *at,
		Description: sym.Description,
		DocURL:      sym.DocURL(),
		Values:      enumValues,
	}
	ei.Module = modules.Get(fw) // main to convert uikit to appkit
	// todo: generalize/improve special case for audiotoolkit enum
	if sym.Name == "AU3DMixerAttenuationCurve" || sym.Name == "AU3DMixerRenderingFlags" {
		sym.Name = strings.ReplaceAll(sym.Name, "3", "Three")
	}
	ei.GName = stringx.Capitalize(modules.TrimPrefix(sym.Name))
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
