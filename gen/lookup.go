package gen

import (
	"github.com/progrium/macschema/schema"
)

func isInstanceType(dt schema.DataType) bool {
	// TODO verify that other fields are zeroed?
	return dt.Name == "instancetype"
}

type typeMapping struct {
	GoType string
	// For objects that can be represented as a simpler `Ref` interface type this
	// is preferred in some cases (e.g. parameter methods) to make handling
	// passing sub-classes simpler since they satisfy the interface.
	// If this is empty, then fall back on `GoType`.
	GoSimpleRefType string
	CType           string
	FromCGoFmt      string
	ToCGoFmt        string
}

func (cb *classBuilder) pkgPrefixForClass(name string) (_ string, _found bool) {
	if name == cb.Class.Name {
		return "", true
	}
	for _, imp := range cb.Imports {
		if !imp.Classes[name] {
			continue
		}
		if imp.Import == nil {
			return "", true
		}
		cb.consumedImports[*imp.Import] = true
		return imp.Import.Alias + ".", true
	}
	return "", false
}

func (cb *classBuilder) mapClass(name string) *typeMapping {
	pkgPrefix, found := cb.pkgPrefixForClass(name)
	if !found {
		return nil
	}
	return &typeMapping{
		GoType:          pkgPrefix + name,
		GoSimpleRefType: pkgPrefix + name + "Ref",
		CType:           "void*",
		FromCGoFmt:      pkgPrefix + name + "_FromPointer(%s)",
		ToCGoFmt:        "objc.RefPointer(%s)",
	}
}

func (cb *classBuilder) mapType(dt schema.DataType) typeMapping {
	if dt.IsPtr {
		if classType := cb.mapClass(dt.Name); classType != nil {
			return *classType
		}
		if dt.Name == "void" {
			return typeMapping{
				GoType:     "unsafe.Pointer",
				CType:      "void*",
				FromCGoFmt: "%s",
				ToCGoFmt:   "%s",
			}
		}
	}
	if dt.IsPtr || dt.IsPtrPtr {
		panic(unimplemented("pointers %#v", dt))
	}
	if isInstanceType(dt) {
		return *cb.mapClass(cb.Class.Name)
	}
	// FIXME(mgood): look these up based on the schema, but for now just use
	// "NSString" as a known class expected to be present to resolve to the
	// "core" package.
	corePkg, found := cb.pkgPrefixForClass("NSString")
	if !found {
		panic("could not locate the `core` package to resolve primitive types")
	}
	switch dt.Name {
	// FIXME split enums into their own types
	case "NSUInteger", "NSWindowStyleMask", "NSBackingStoreType", "NSWindowOrderingMode", "NSWindowCollectionBehavior":
		return typeMapping{
			GoType:     corePkg + "NSUInteger",
			CType:      "unsigned long",
			FromCGoFmt: corePkg + "NSUInteger(%s)",
			ToCGoFmt:   "C.ulong(%s)",
		}
	case "NSInteger", "NSWindowTitleVisibility", "NSWindowLevel", "NSApplicationActivationPolicy", "NSControlStateValue", "NSPopoverBehavior":
		return typeMapping{
			GoType:     corePkg + "NSInteger",
			CType:      "long",
			FromCGoFmt: corePkg + "NSInteger(%s)",
			ToCGoFmt:   "C.long(%s)",
		}
	case "CGFloat":
		return typeMapping{
			GoType:     corePkg + "CGFloat",
			CType:      "double",
			FromCGoFmt: corePkg + "CGFloat(%s)",
			ToCGoFmt:   "C.double(%s)",
		}
	case "NSStringEncoding":
		return typeMapping{
			GoType:     corePkg + "NSStringEncoding",
			CType:      "unsigned long",
			FromCGoFmt: corePkg + "NSStringEncoding(%s)",
			ToCGoFmt:   "C.ulong(%s)",
		}
	case "unichar":
		return typeMapping{
			GoType:     corePkg + "Unichar",
			CType:      "unsigned short",
			FromCGoFmt: corePkg + "Unichar(%s)",
			ToCGoFmt:   "C.ushort(%s)",
		}
	case "BOOL":
		return typeMapping{
			GoType:     "bool",
			CType:      "BOOL",
			FromCGoFmt: "convertObjCBoolToGo(%s)",
			ToCGoFmt:   "convertToObjCBool(%s)",
		}
	case "int":
		return typeMapping{
			GoType:     "int32",
			CType:      "int",
			FromCGoFmt: "int32(%s)",
			ToCGoFmt:   "C.int(%s)",
		}
	case "SEL":
		return typeMapping{
			GoType:     "objc.Selector",
			CType:      "void*",
			FromCGoFmt: "objc.SelectorAt(%s)",
			ToCGoFmt:   "%s.SelectorAddress()",
		}
	case "NSRect", "CGRect":
		return typeMapping{
			GoType:     corePkg + "NSRect",
			CType:      "NSRect",
			FromCGoFmt: "*(*" + corePkg + "NSRect)(unsafe.Pointer(&%s))",
			ToCGoFmt:   "*(*C.NSRect)(unsafe.Pointer(&%s))",
		}
	case "NSPoint":
		return typeMapping{
			GoType:     corePkg + "NSPoint",
			CType:      "NSPoint",
			FromCGoFmt: "*(*" + corePkg + "NSPoint)(unsafe.Pointer(&%s))",
			ToCGoFmt:   "*(*C.NSPoint)(unsafe.Pointer(&%s))",
		}
	case "NSSize", "CGSize":
		return typeMapping{
			GoType:     corePkg + "NSSize",
			CType:      "NSSize",
			FromCGoFmt: "*(*" + corePkg + "NSSize)(unsafe.Pointer(&%s))",
			ToCGoFmt:   "*(*C.NSSize)(unsafe.Pointer(&%s))",
		}
	case "id":
		return typeMapping{
			GoType:          "objc.Object",
			GoSimpleRefType: "objc.Ref",
			CType:           "void*",
			FromCGoFmt:      "objc.Object_FromPointer(%s)",
			ToCGoFmt:        "objc.RefPointer(%s)",
		}
	default:
		panic(unimplemented("mapType: %s", dt.Name))
	}
}
