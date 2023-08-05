package typing

import (
	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/internal/set"
)

var _ Type = (*ArrayType)(nil)

// ArrayType the element should be StringType or ClassType
type ArrayType struct {
	Type Type
}

func (a *ArrayType) GoImports() set.Set[string] {
	imports := set.New("github.com/progrium/macdriver/macos/foundation")
	imports.AddSet(a.Type.GoImports())
	return imports
}

func (a *ArrayType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	return "[]" + a.Type.GoName(currentModule, receiveFromObjc)
}

func (a *ArrayType) ObjcName() string {
	return "NSArray*"
}

func (a *ArrayType) DeclareModule() *modules.Module {
	return a.Type.DeclareModule()
}
