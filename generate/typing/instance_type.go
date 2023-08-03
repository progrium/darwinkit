package typing

import (
	"github.com/progrium/macdriver/internal/set"
)

var _ Type = (*InstanceType)(nil)

// InstanceType the objc instancetype.
// class method start with alloc or new, instance method start with autorelease，init，retain or self, return instancetype.
type InstanceType struct {
}

func (i *InstanceType) GoImports() set.Set[string] {
	panic("implement me")
}

func (i *InstanceType) GoName(currentModule *Module, receiveFromObjc bool) string {
	panic("implement me")
}

func (i *InstanceType) ObjcName() string {
	panic("implement me")
}

func (i *InstanceType) DeclareModule() *Module {
	return nil
}
