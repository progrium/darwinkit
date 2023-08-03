package codegen

import (
	"github.com/progrium/macdriver/generate/typing"
)

// Param is code generator for objective-c method param
type Param struct {
	Name      string
	Type      typing.Type
	FieldName string // objc param field name(part of function name)
}

func (p *Param) String() string {
	return p.FieldName + ":" + p.Name
}

// GoDeclare return go param declare code
func (p *Param) GoDeclare(currentModule *typing.Module, receiveFromObjc bool) string {
	return p.GoName() + " " + p.Type.GoName(currentModule, receiveFromObjc)
}

func (p *Param) ObjcDeclare() string {
	return p.Type.ObjcName() + " " + p.GoName()
}

func (p *Param) GoName() string {
	switch p.Name {
	case "type", "range", "map", "string", "select":
		return p.Name + "_"
	}
	return p.Name
}
