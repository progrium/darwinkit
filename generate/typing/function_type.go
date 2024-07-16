package typing

import (
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/internal/set"
)

var _ Type = (*FunctionType)(nil)

type Parameter struct {
	Type Type
	Name string // the param name
}

// FunctionType Objective-c function type
type FunctionType struct {
	Name   string          // objc type name
	GName  string          // Go name, usually is objc type name without prefix 'NS'
	Module *modules.Module // object-c module

	Parameters []Parameter // function parameters
	ReturnType Type        // function return type
}

var Function = &FunctionType{
	Name:   "Function",
	GName:  "Function",
	Module: modules.Get("objc"),
}

func (c *FunctionType) GoImports() set.Set[string] {
	return set.New("github.com/progrium/darwinkit/macos/" + c.Module.Package)
}

func (c *FunctionType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	if receiveFromObjc {
		return FullGoName(*c.Module, c.GoStructName(), *currentModule)
	} else {
		return FullGoName(*c.Module, c.GoInterfaceName(), *currentModule)
	}
}

func (c *FunctionType) ObjcName() string {
	return c.Name + "*"
}

func (c *FunctionType) CName() string {
	return c.Name + "*"
}

func (c *FunctionType) DeclareModule() *modules.Module {
	return c.Module
}

// GoInterfaceName return the go wrapper interface name
func (c *FunctionType) GoInterfaceName() string {
	return "I" + c.GName
}

// GoStructName return the go wrapper struct name
func (c *FunctionType) GoStructName() string {
	return c.GName
}
