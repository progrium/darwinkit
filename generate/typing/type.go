package typing

import (
	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/internal/set"
)

// Type interface for all type
type Type interface {
	// GoName Go type name
	GoName(currentModule *modules.Module, receiveFromObjc bool) string
	// ObjcName Objective-c type name
	ObjcName() string

	// GoImports go imports for this type
	GoImports() set.Set[string]

	// DeclareModule the module of this type. return nil if is a built in type
	DeclareModule() *modules.Module
}

func FullGoName(module modules.Module, name string, currentModule modules.Module) string {
	if module == currentModule {
		return name
	}
	return module.Package + "." + name
}

func PrependPackage(module modules.Module, s string, currentModule modules.Module) string {
	if module == currentModule {
		return s
	}
	return module.Package + "." + s
}
