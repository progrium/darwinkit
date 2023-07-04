package gen

import (
	"github.com/progrium/macschema/schema"
)

type structBuilder struct {
	Struct          schema.Struct
	Imports         []PackageContents
	consumedImports map[Import]bool
}

// func (sb *structBuilder) EachTypeMethod(f func(schema.Method)) {
// 	f(schema.Method{
// 		Name:   "alloc",
// 		Return: schema.DataType{Name: "instancetype"},
// 	})
// 	for _, m := range sb.Struct.TypeMethods {
// 		f(m)
// 	}
// 	for _, p := range sb.Struct.TypeProperties {
// 		for _, m := range propertyMethods(p) {
// 			f(m)
// 		}
// 	}
// }

// func (sb *structBuilder) EachInstanceMethod(f func(schema.Method)) {
// 	seen := make(map[string]bool)
// 	for _, m := range sb.Struct.InstanceMethods {
// 		seen[m.Name] = true
// 		f(m)
// 	}
// 	if !seen["init"] {
// 		f(schema.Method{
// 			Name:   "init",
// 			Return: schema.DataType{Name: "instancetype"},
// 		})
// 	}
// 	for _, p := range sb.Struct.InstanceProperties {
// 		func() {
// 			defer ignoreIfUnimplemented(fmt.Sprintf("%s.%s", sb.Struct.Name, p.Name))
// 			for _, m := range propertyMethods(p) {
// 				// properties sometimes specify a getter or setter method that is also
// 				// in `InstanceMethods`, so skip those if they've already been handled
// 				if !seen[m.Name] {
// 					f(m)
// 				}
// 			}
// 		}()
// 	}
// }

// func (sb *structBuilder) instanceMethod(method schema.Method) MethodDef {
// 	r := MethodDef{
// 		Name:        toExportedName(selectorNameToGoIdent(method.Name)),
// 		WrappedFunc: sb.cgoWrapperFunc(method, false),
// 	}
// 	if isInstanceType(method.Return) {
// 		r.Name += "_as" + sb.Struct.Name
// 	}
// 	return r
// }

// func (sb *structBuilder) msgSend(method schema.Method, isTypeMethod bool) CGoMsgSend {
// 	msg := CGoMsgSend{
// 		Name:   msgSendFuncName(sb.Struct, method.Name, isTypeMethod),
// 		Class:  sb.Struct.Name,
// 		Return: sb.toMsgSendReturn(method.Return),
// 	}
// 	for i, key := range strings.SplitAfter(method.Name, ":") {
// 		if key == "" {
// 			continue
// 		}
// 		part := SelectorPart{
// 			Key: key,
// 		}
// 		if i < len(method.Args) {
// 			arg := method.Args[i]
// 			typ := sb.mapType(arg.Type)
// 			part.Arg = &Arg{
// 				Name: arg.Name,
// 				Type: typ.CType,
// 			}
// 		}
// 		msg.Selector = append(msg.Selector, part)
// 	}
// 	return msg
// }

func (sb *structBuilder) toMsgSendReturn(dt schema.DataType) string {
	if isVoid(dt) {
		return "void"
	}
	typ := sb.mapType(dt)
	return typ.CType
}

// func (sb *structBuilder) cgoWrapperFunc(method schema.Method, isTypeMethod bool) CGoWrapperFunc {
// 	r := CGoWrapperFunc{
// 		Name:    msgSendFuncName(sb.Struct, method.Name, isTypeMethod),
// 		Args:    []CGoWrapperArg{},
// 		Returns: sb.toCGoWrapperReturn(method.Return),
// 	}
// 	for _, arg := range method.Args {
// 		typ := sb.mapType(arg.Type)
// 		goType := typ.GoSimpleRefType
// 		if goType == "" {
// 			goType = typ.GoType
// 		}
// 		r.Args = append(r.Args, CGoWrapperArg{
// 			Name:     arg.Name,
// 			Type:     goType,
// 			ToCGoFmt: typ.ToCGoFmt,
// 		})
// 	}
// 	return r
// }

func (sb *structBuilder) toCGoWrapperReturn(dt schema.DataType) []CGoWrapperReturn {
	if isVoid(dt) {
		return nil
	}
	typ := sb.mapType(dt)
	return []CGoWrapperReturn{
		{Type: typ.GoType, FromCGoFmt: typ.FromCGoFmt},
	}
}

func (sb *structBuilder) pkgPrefixForStruct(name string) (_ string, _found bool) {
	if name == sb.Struct.Name {
		return "", true
	}
	for _, imp := range sb.Imports {
		if !imp.Classes[name] {
			continue
		}
		if imp.Import == nil {
			return "", true
		}
		sb.consumedImports[*imp.Import] = true
		return imp.Import.Alias + ".", true
	}
	return "", false
}

func (sb *structBuilder) mapStruct(name string) *typeMapping {
	pkgPrefix, found := sb.pkgPrefixForStruct(name)
	if !found {
		return nil
	}
	return &typeMapping{
		GoType:          pkgPrefix + name,
		GoSimpleRefType: pkgPrefix + name + "Ref",
		CType:           "void*",
		FromCGoFmt:      pkgPrefix + name + "_fromPointer(%s)",
		ToCGoFmt:        "objc.RefPointer(%s)",
	}
}

func (sb *structBuilder) mapType(dt schema.DataType) typeMapping {
	if dt.IsPtr {
		if structType := sb.mapStruct(dt.Name); structType != nil {
			return *structType
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
		return *sb.mapStruct(sb.Struct.Name)
	}
	// FIXME(mgood): look these up based on the schema, but for now just use
	// "NSString" as a known struct expected to be present to resolve to the
	// "core" package.
	corePkg, found := sb.pkgPrefixForStruct("NSString")
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
			FromCGoFmt:      "objc.Object_fromPointer(%s)",
			ToCGoFmt:        "objc.RefPointer(%s)",
		}
	default:
		panic(unimplemented("mapType: %s", dt.Name))
	}
}
