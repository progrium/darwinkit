package typing

import (
	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/internal/set"
)

var _ Type = (*VoidType)(nil)

// VoidType for void return
type VoidType struct {
}

func (d *VoidType) GoImports() set.Set[string] {
	return nil
}

func (d *VoidType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	return ""
}

func (d *VoidType) ObjcName() string {
	return "void"
}

func (d *VoidType) DeclareModule() *modules.Module {
	return nil
}
