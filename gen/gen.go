package gen

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/progrium/macschema/schema"
)

type classBuilder struct {
	Class           schema.Class
	Imports         []PackageContents
	consumedImports map[Import]bool
}

func (cb *classBuilder) EachTypeMethod(f func(schema.Method)) {
	f(schema.Method{
		Name:   "alloc",
		Return: schema.DataType{Name: "instancetype"},
	})
	for _, m := range cb.Class.TypeMethods {
		f(m)
	}
	for _, p := range cb.Class.TypeProperties {
		for _, m := range propertyMethods(p) {
			f(m)
		}
	}
}

func (cb *classBuilder) EachInstanceMethod(f func(schema.Method)) {
	seen := make(map[string]bool)
	for _, m := range cb.Class.InstanceMethods {
		seen[m.Name] = true
		f(m)
	}
	if !seen["init"] {
		f(schema.Method{
			Name:   "init",
			Return: schema.DataType{Name: "instancetype"},
		})
	}
	for _, p := range cb.Class.InstanceProperties {
		func() {
			defer ignoreIfUnimplemented(fmt.Sprintf("%s.%s", cb.Class.Name, p.Name))
			for _, m := range propertyMethods(p) {
				// properties sometimes specify a getter or setter method that is also
				// in `InstanceMethods`, so skip those if they've already been handled
				if !seen[m.Name] {
					f(m)
				}
			}
		}()
	}
}

func (cb *classBuilder) instanceMethod(method schema.Method) MethodDef {
	r := MethodDef{
		Name:        toExportedName(selectorNameToGoIdent(method.Name)),
		WrappedFunc: cb.cgoWrapperFunc(method, false),
	}
	if isInstanceType(method.Return) {
		r.Name += "_as" + cb.Class.Name
	}
	return r
}

func (cb *classBuilder) msgSend(method schema.Method, isTypeMethod bool) CGoMsgSend {
	msg := CGoMsgSend{
		Name:   msgSendFuncName(cb.Class, method.Name, isTypeMethod),
		Class:  cb.Class.Name,
		Return: cb.toMsgSendReturn(method.Return),
	}
	for i, key := range strings.SplitAfter(method.Name, ":") {
		if key == "" {
			continue
		}
		part := SelectorPart{
			Key: key,
		}
		if i < len(method.Args) {
			arg := method.Args[i]
			typ := cb.mapType(arg.Type)
			part.Arg = &Arg{
				Name: arg.Name,
				Type: typ.CType,
			}
		}
		msg.Selector = append(msg.Selector, part)
	}
	return msg
}

func (cb *classBuilder) toMsgSendReturn(dt schema.DataType) string {
	if isVoid(dt) {
		return "void"
	}
	typ := cb.mapType(dt)
	return typ.CType
}

func (cb *classBuilder) cgoWrapperFunc(method schema.Method, isTypeMethod bool) CGoWrapperFunc {
	r := CGoWrapperFunc{
		Name:    msgSendFuncName(cb.Class, method.Name, isTypeMethod),
		Args:    []CGoWrapperArg{},
		Returns: cb.toCGoWrapperReturn(method.Return),
	}
	for _, arg := range method.Args {
		typ := cb.mapType(arg.Type)
		goType := typ.GoSimpleRefType
		if goType == "" {
			goType = typ.GoType
		}
		r.Args = append(r.Args, CGoWrapperArg{
			Name:     arg.Name,
			Type:     goType,
			ToCGoFmt: typ.ToCGoFmt,
		})
	}
	return r
}

func (cb *classBuilder) toCGoWrapperReturn(dt schema.DataType) []CGoWrapperReturn {
	if isVoid(dt) {
		return nil
	}
	typ := cb.mapType(dt)
	return []CGoWrapperReturn{
		{Type: typ.GoType, FromCGoFmt: typ.FromCGoFmt},
	}
}

func isVoid(dt schema.DataType) bool {
	// NOTE: IBAction is a special hint type for Interface Builder, but for our
	// purposes should be `void`:
	// #define IBAction void
	// See https://nshipster.com/ibaction-iboutlet-iboutletcollection/
	if dt.Name == "IBAction" {
		return true
	}
	return dt.Name == "void" && !dt.IsPtr && !dt.IsPtrPtr
}

func msgSendFuncName(cls schema.Class, selector string, isTypeMethod bool) string {
	target := "inst"
	if isTypeMethod {
		target = "type"
	}
	return fmt.Sprintf("%s_%s_%s", cls.Name, target, selectorNameToGoIdent(selector))
}

func selectorNameToGoIdent(sel string) string {
	return strings.ReplaceAll(sel, ":", "_")
}

// Objective-C properties are syntactic sugar for getter/setter methods, this
// translates the property definition into those methods.
func propertyMethods(p schema.Property) []schema.Method {
	readonly := false
	for k, v := range p.Attrs {
		switch k {
		case "readonly":
			readonly = v.(bool)
		case "copy", "strong", "nullable", "class", "getter", "nonatomic", "assign",
			"retain", "weak":
			// no-op
		default:
			panic(unimplemented("unsupported property attr: %s", k))
		}
	}
	r := []schema.Method{
		propertyReadMethod(p),
	}
	if !readonly {
		r = append(r, propertyWriteMethod(p))
	}
	return r
}

func propertyReadMethod(p schema.Property) schema.Method {
	m := schema.Method{
		Name:        p.Name,
		Description: p.Description,
		Declaration: p.Declaration,
		Return:      p.Type,
		Args:        nil,
		Deprecated:  p.Deprecated,
		TopicURL:    p.TopicURL,
	}
	if getter, ok := p.Attrs["getter"]; ok {
		m.Name = getter.(string)
	}
	return m
}

func propertyWriteMethod(p schema.Property) schema.Method {
	return schema.Method{
		Name:        "set" + toExportedName(p.Name) + ":", // TODO rename the function?
		Description: p.Description,
		Declaration: p.Declaration,
		Return:      schema.DataType{Name: "void"},
		Args: []schema.Arg{
			{Name: "value", Type: p.Type},
		},
		Deprecated: p.Deprecated,
		TopicURL:   p.TopicURL,
	}
}

// grabbed from go/token
func toExportedName(name string) string {
	ch, n := utf8.DecodeRuneInString(name)
	if unicode.IsUpper(ch) {
		return name
	}
	ch = unicode.ToUpper(ch)
	// TODO confirm that there is actually an upper-case mapping? or we would need
	// to skip this method, or prepend a letter here to make it a valid exported
	// Go method name.
	return string(ch) + name[n:]
}

func Convert(desc PackageDescription, imports []PackageContents, schemas ...*schema.Schema) GoPackage {
	pkg := GoPackage{
		PackageDescription: desc,
	}
	consumedImports := map[Import]bool{
		{Path: "unsafe"}: true,
		{Path: "github.com/progrium/macdriver/objc"}: true,
	}
	for _, s := range schemas {
		cb := classBuilder{
			Class:           *s.Class,
			Imports:         imports,
			consumedImports: consumedImports,
		}
		classDef := ClassDef{
			Name: cb.Class.Name,
			Base: "objc.Object",
		}
		if decl := parseClassDeclaration(cb.Class.Declaration); decl.Base != "NSObject" {
			// TODO require resolving every base class but for now just fall back on
			// objc.Object if we don't have the base type in the schema
			if cls := cb.mapClass(decl.Base); cls != nil {
				classDef.Base = cls.GoType
			}
		}
		cb.EachTypeMethod(func(m schema.Method) {
			defer ignoreIfUnimplemented(fmt.Sprintf("%s.%s", s.Class.Name, m.Name))

			msg := cb.msgSend(m, true)
			wrapper := MethodDef{
				Name:        fmt.Sprintf("%s_%s", cb.Class.Name, selectorNameToGoIdent(m.Name)),
				WrappedFunc: cb.cgoWrapperFunc(m, true),
			}

			pkg.ClassMsgSendWrappers = append(pkg.ClassMsgSendWrappers, msg)
			pkg.CGoWrapperFuncs = append(pkg.CGoWrapperFuncs, wrapper)
		})
		cb.EachInstanceMethod(func(m schema.Method) {
			defer ignoreIfUnimplemented(fmt.Sprintf("%s.%s", s.Class.Name, m.Name))

			method := cb.instanceMethod(m)
			msg := cb.msgSend(m, false)

			classDef.InstanceMethods = append(classDef.InstanceMethods, method)
			pkg.MsgSendWrappers = append(pkg.MsgSendWrappers, msg)
		})
		pkg.Classes = append(pkg.Classes, classDef)
	}
	for imp := range consumedImports {
		pkg.Imports = append(pkg.Imports, imp)
	}
	return pkg
}

type declaration struct {
	Name string
	Base string
}

func parseClassDeclaration(decl string) declaration {
	decl = strings.TrimPrefix(decl, "@interface ")
	parts := strings.Split(decl, ":")
	if len(parts) != 2 {
		panic(fmt.Errorf("unable to parse: %q", decl))
	}
	return declaration{
		Name: strings.TrimSpace(parts[0]),
		Base: strings.TrimSpace(parts[1]),
	}
}
