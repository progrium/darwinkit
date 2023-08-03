package typing

import (
	"github.com/progrium/macdriver/internal/set"
)

var _ Type = (*VoidPointerType)(nil)

// VoidPointerType void*/unsafe.Pointer
type VoidPointerType struct {
	Name  string // objc type name
	GName string // Go name, usually is objc type name without prefix 'NS'
}

func (c *VoidPointerType) GoImports() set.Set[string] {
	return set.New("unsafe")
}

func (c *VoidPointerType) GoName(currentModule *Module, receiveFromObjc bool) string {
	return "unsafe.Pointer"
}

func (c *VoidPointerType) ObjcName() string {
	return "void*"
}

func (c *VoidPointerType) DeclareModule() *Module {
	return nil
}
