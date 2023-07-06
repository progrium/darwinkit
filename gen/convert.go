package gen

import (
	"fmt"
	"strings"

	"github.com/progrium/macschema/schema"
)

func Convert(desc PackageDescription, imports []PackageContents, schemas ...*schema.Schema) (*GoPackage, error) {
	pkg := &GoPackage{
		PackageDescription: desc,
	}
	consumedImports := map[Import]bool{
		{Path: "unsafe"}: true,
		{Path: "github.com/progrium/macdriver/objc"}: true,
	}
	for _, s := range schemas {
		if s.Class != nil {
			classDef, err := processClassSchema(pkg, s, imports, consumedImports)
			if err != nil {
				return pkg, fmt.Errorf("issue with class %s failed to parse declaration %+v: %w", s.Class.Name, s.Class, err)
			}
			pkg.Classes = append(pkg.Classes, classDef)
		} else if s.Struct != nil {
			return pkg, fmt.Errorf("%v: structs are not yet supported", s.Struct.Name)
		} else {
			return pkg, fmt.Errorf("invalid schema kind %v", s.Kind)
		}
	}
	for imp := range consumedImports {
		pkg.Imports = append(pkg.Imports, imp)
	}
	return pkg, nil
}

// processClassSchema converts a schema.Class into a ClassDef
func processClassSchema(pkg *GoPackage, s *schema.Schema, imports []PackageContents, consumedImports map[Import]bool) (ClassDef, error) {
	cb := classBuilder{
		Class:           *s.Class,
		Imports:         imports,
		consumedImports: consumedImports,
		generatedNames:  map[string]string{},
	}
	classDef := ClassDef{
		Name: cb.Class.Name,
		Base: "objc.Object",
	}
	decl, err := parseClassDeclaration(cb.Class.Declaration)
	if err != nil {
		return classDef, err
	}

	if decl.Base != "NSObject" {
		if cls := cb.mapClass(decl.Base); cls != nil {
			classDef.Base = cls.GoType
		}
	}

	cb.EachTypeMethod(func(m schema.Method) {
		defer ignoreIfUnimplemented(fmt.Sprintf("%s.%s", s.Class.Name, m.Name))

		msg := cb.msgSend(m, true)
		ident := selectorNameToGoIdent(cb.generatedNames, m.Name)
		name := fmt.Sprintf("%s_%s", cb.Class.Name, ident)
		wrapper := MethodDef{
			Description: formatComment(m, name),
			Name:        name,
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

	return classDef, nil
}

func formatComment(m schema.Method, ident string) string {
	ld := strings.ToLower(m.Description)
	firstWord := strings.Split(ld, " ")[0]

	if firstWord == "a" || firstWord == "the" {
		ld = "returns " + ld
	}
	return fmt.Sprintf("// %s %s\n//\n// See %s for details.", ident, ld, m.TopicURL)

}
