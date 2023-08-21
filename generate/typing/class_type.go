package typing

import (
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/internal/set"
)

var _ Type = (*ClassType)(nil)

// ClassType Objective-c interface type
type ClassType struct {
	Name   string          // objc type name
	GName  string          // Go name, usually is objc type name without prefix 'NS'
	Module *modules.Module // object-c module
}

// Object is the objc object root class type
var Object = &ClassType{
	Name:   "NSObject",
	GName:  "Object",
	Module: modules.Get("objc"),
}

var Class = &ClassType{
	Name:   "Class",
	GName:  "Class",
	Module: modules.Get("objc"),
}

func (c *ClassType) GoImports() set.Set[string] {
	return set.New("github.com/progrium/darwinkit/macos/" + c.Module.Package)
}

func (c *ClassType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	if receiveFromObjc {
		return FullGoName(*c.Module, c.GoStructName(), *currentModule)
	} else {
		return FullGoName(*c.Module, c.GoInterfaceName(), *currentModule)
	}
}

func (c *ClassType) ObjcName() string {
	return c.Name + "*"
}

func (c *ClassType) CName() string {
	return c.Name + "*"
}

func (c *ClassType) DeclareModule() *modules.Module {
	return c.Module
}

// GoInterfaceName return the go wrapper interface name
func (c *ClassType) GoInterfaceName() string {
	return "I" + c.GName
}

// GoStructName return the go wrapper struct name
func (c *ClassType) GoStructName() string {
	return c.GName
}
