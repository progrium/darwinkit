package typing

import (
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/internal/set"
	"github.com/progrium/darwinkit/internal/stringx"
)

// Type interface for all type
type Type interface {
	// GoName Go type name
	GoName(currentModule *modules.Module, receiveFromObjc bool) string
	// ObjcName Objective-c type name
	ObjcName() string

	// CName C type name
	CName() string

	// GoImports go imports for this type
	GoImports() set.Set[string]

	// DeclareModule the module of this type. return nil if is a built in type
	DeclareModule() *modules.Module
}

func FullGoName(module modules.Module, name string, currentModule modules.Module) string {
	if module.Name == currentModule.Name {
		return stringx.Capitalize(name)
	}
	return module.Package + "." + stringx.Capitalize(name)
}
