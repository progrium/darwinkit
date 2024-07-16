package typing

import (
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/internal/set"
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

func (c *VoidPointerType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	return "unsafe.Pointer"
}

func (c *VoidPointerType) ObjcName() string {
	return "void*"
}

func (c *VoidPointerType) CName() string {
	return "void*"
}

func (c *VoidPointerType) CSignature() string {
	return "void *"
}

func (c *VoidPointerType) DeclareModule() *modules.Module {
	return nil
}
