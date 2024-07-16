package generate

import (
	"fmt"
	"log"
	"strings"

	"github.com/progrium/darwinkit/generate/declparse"
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/generate/typing"
	"github.com/progrium/darwinkit/internal/stringx"
)

func (db *Generator) TypeFromSymbol(sym Symbol) typing.Type {
	m := sym.MainModule()
	if m == nil {
		return nil
	}
	module := m.Name
	if db.Platform == "macos" && module == "UIKit" {
		module = "AppKit"
	}
	// cases where symbol lives in two places,
	// we want it local if it belongs to the
	// framework we're generating
	if sym.HasFramework(db.Framework) {
		module = db.Framework
	}
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
		st, err := sym.Parse(db.Platform)
		if err != nil || st.Enum == nil {
			fmt.Printf("TypeFromSymbol: name=%s declaration=%s\n", sym.Name, sym.Declaration)
			log.Panicf("invalid enum decl. %s", err)
		}

		return &typing.AliasType{
			Name:   sym.Name,
			GName:  stringx.Capitalize(modules.TrimPrefix(sym.Name)), // reset in ToEnumInfo
			Module: modules.Get(module),
			Type:   db.ParseType(st.Enum.Type),
		}
	case "Union":
		return &typing.RefType{
			Name:  sym.Name,
			GName: modules.TrimPrefix(sym.Name),
		}
	case "Type":
		if sym.Type != "Type Alias" {
			fmt.Printf("TypeFromSymbol: name=%s declaration=%s path=%s\n", sym.Name, sym.Declaration, sym.Path)
			panic("unknown type")
		}

		// special handling of Ref structs
		if (strings.HasSuffix(sym.Name, "Ref") && strings.Contains(sym.Declaration, "struct")) ||
			sym.Name == "AudioComponent" ||
			// sym.Name == "NSZone" ||
			sym.Name == "MusicSequence" {
			return &typing.RefType{
				Name:  sym.Name,
				GName: modules.TrimPrefix(sym.Name),
			}
		}
		st, err := sym.Parse(db.Platform)
		if err != nil {
			fmt.Printf("TypeFromSymbol: name=%s declaration=%s path=%s\n", sym.Name, sym.Declaration, sym.Path)
			panic("bad declaration")
		}
		if st.Struct != nil {
			return &typing.RefType{
				Name:  st.Struct.Name,
				GName: modules.TrimPrefix(sym.Name),
			}
		}
		if st.TypeAlias == nil {
			fmt.Printf("TypeFromSymbol: name=%s declaration=%s path=%s\n", sym.Name, sym.Declaration, sym.Path)
			panic("bad type alias")
		}
		typ := db.ParseType(*st.TypeAlias)
		if typ == nil {
			fmt.Printf("TypeFromSymbol: name=%s declaration=%s path=%s\n", sym.Name, sym.Declaration, sym.Path)
			panic("unable to parse type")
		}
		typ = &typing.AliasType{
			Name:   sym.Name,
			GName:  modules.TrimPrefix(sym.Name),
			Module: modules.Get(module),
			Type:   typ,
		}
		return typ
	case "Struct":
		if strings.HasSuffix(sym.Name, "Ref") {
			return &typing.RefType{
				Name:   sym.Name,
				GName:  modules.TrimPrefix(sym.Name),
				Module: modules.Get(module),
			}
		}
		return &typing.StructType{
			Name:   sym.Name,
			GName:  modules.TrimPrefix(sym.Name),
			Module: modules.Get(module),
		}
	case "Function":
		if sym.Name != "CGDisplayCreateImage" &&
			sym.Name != "CGMainDisplayID" {
			return nil
		}
		typ, err := sym.Parse(db.Platform)
		if err != nil {
			fmt.Printf("TypeFromSymbol: failed to parse %s: %s\n", sym.Declaration, err)
			return nil
		}
		fn := typ.Function
		if fn == nil {
			fmt.Printf("TypeFromSymbol: name=%s declaration=%s\n", sym.Name, sym.Declaration)
			return nil
		}
		ft := &typing.FunctionType{
			Name:   sym.Name,
			GName:  modules.TrimPrefix(sym.Name),
			Module: modules.Get(module),
		}
		for _, arg := range fn.Args {
			ft.Parameters = append(ft.Parameters, typing.Parameter{
				Name: arg.Name,
				Type: db.ParseType(arg.Type),
			})
		}
		ft.ReturnType = db.ParseType(fn.ReturnType)
		return ft
	default:
		fmt.Printf("TypeFromSymbol: kind=%s name=%s path=%s\n", sym.Kind, sym.Name, sym.Path)
		panic("bad type")
	}

}

// ParseType parses a type from a declparse.TypeInfo.
func (db *Generator) ParseType(ti declparse.TypeInfo) (typ typing.Type) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("ParseType: %s panic=%s\n", ti.Name, r)
			if strings.Contains(r.(string), "module not found") {
				if !modules.CanIgnoreNotFound(r) {
					panic(r)
				}
			}
		}
	}()
	if ti.Func != nil {
		var blockParams []typing.BlockParam
		for _, arg := range ti.Func.Args {
			blockParams = append(blockParams, typing.BlockParam{
				Name: arg.Name,
				Type: db.ParseType(arg.Type),
			})
		}
		return &typing.BlockType{
			ReturnType: db.ParseType(ti.Func.ReturnType),
			Params:     blockParams,
		}
	}

	ref := false
	switch ti.Name {
	case "SEL":
		typ = &typing.SelectorType{}
	case "void":
		typ = &typing.VoidType{}
	case "instancetype":
		typ = &typing.InstanceType{}
	case "id":
		typ = &typing.IDType{}
		if len(ti.Params) > 0 && ti.Params[0].Name != "NSObject" {
			idtype := db.FindTypeSymbol(ti.Params[0].Name)
			if idtype != nil {
				ptyp, ok := db.TypeFromSymbol(*idtype).(*typing.ProtocolType)
				if ok {
					typ = ptyp
				}
			}
		}
	case "OSType", "OSStatus", "FourCharCode":
		// to kernel types?
		typ = typing.UInt
	case "kern_return_t":
		// to kernel type
		typ = typing.Int
	case "Class":
		// objc type
		typ = typing.Class
	case "CGFloat", "Float64":
		typ = typing.Double
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
	case "NSZone", "ipc_port_t", "io_object_t":
		typ = &typing.RefType{Name: ti.Name}
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
	case "NSObject":
		// force possible NSObject protocol to be NSObject class
		typ = typing.Object
		ref = true
	default:

		var ok bool
		typ, ok = typing.GetPrimitiveType(ti.Name)
		// log.Println("primitive", ti.Name, ok)
		if !ok {
			typ, ok = typing.GetDispatchType(ti.Name)
		}
		// log.Println("dispatch", ti.Name, ok)
		if !ok {
			typ, ok = typing.GetKernelType(ti.Name)
		}
		// log.Println("kernel", ti.Name, ok)
		if !ok {
			typ = db.TypeFromSymbolName(ti.Name)
			// log.Printf("symbol %v %T %v - %v", ti.Name, typ, ok, string(j))
			switch typ.(type) {
			case *typing.ClassType:
				ref = true
			case *typing.StructType:
				//ref = true
			case *typing.ProtocolType:
				panic("standalone proto type")
			}
		}
	}

	if _, ok := typ.(*typing.CStringType); ok {
		return typ
	}

	//fmt.Printf("ParseType: %s %s %s\n", ti.Name, typ, ti)
	if (ti.IsPtr || ti.IsPtrPtr) && !ref {
		if _, ok := typ.(*typing.VoidType); ok {
			typ = &typing.VoidPointerType{}
		} else {
			typ = &typing.PointerType{
				Type:    typ,
				IsConst: ti.Annots[declparse.TypeAnnotConst],
			}
		}
	}

	return
}
