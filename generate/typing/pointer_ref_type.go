package typing

import (
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/internal/set"
)

var _ Type = (*PointerRefType)(nil)

// PointerRefType c pointer type def
type PointerRefType struct {
	Name   string          // objc type name
	GName  string          // Go name, usually is objc type name without prefix 'NS'
	Module *modules.Module // object-c module
}

func (c *PointerRefType) GoImports() set.Set[string] {
	return set.New("github.com/progrium/darwinkit/macos/" + c.Module.Package)
}

func (c *PointerRefType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	return FullGoName(*c.Module, c.GName, *currentModule)
}

func (c *PointerRefType) ObjcName() string {
	return c.Name
}

func (c *PointerRefType) CName() string {
	return c.Name
}

func (c *PointerRefType) DeclareModule() *modules.Module {
	return c.Module
}
