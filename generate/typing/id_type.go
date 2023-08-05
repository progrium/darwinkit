package typing

import (
	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/internal/set"
)

var _ Type = (*IDType)(nil)

// InstanceType the objc instancetype.
// class method start with alloc or new, instance method start with autorelease，init，retain or self, return instancetype.
type IDType struct {
}

func (i *IDType) GoImports() set.Set[string] {
	return Object.GoImports()
}

func (i *IDType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	return Object.GoName(currentModule, receiveFromObjc)
}

func (i *IDType) ObjcName() string {
	return "id"
}

func (i *IDType) DeclareModule() *modules.Module {
	return Object.DeclareModule()
}
