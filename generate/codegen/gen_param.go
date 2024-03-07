package codegen

import (
	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/generate/typing"
)

// Param is code generator for objective-c method param
type Param struct {
	Name      string
	Type      typing.Type
	FieldName string // objc param field name(part of function name)
	Object    bool   // version of param generalized to IObject for protocols
	IsPtrPtr  bool
}

func (p *Param) String() string {
	return p.FieldName + ":" + p.Name
}

// Return Go parameter code
func (p *Param) GoDeclare(currentModule *modules.Module, receiveFromObjc bool) string {
	returnValue := p.GoName() + " "
	if p.IsPtrPtr == true { // Example: NSError **error
		returnValue = returnValue + "unsafe.Pointer"
	} else {
		returnValue = returnValue + p.Type.GoName(currentModule, receiveFromObjc)
	}
	return returnValue
}

func (p *Param) ObjcDeclare() string {
	return p.Type.ObjcName() + " " + p.GoName()
}

func (p *Param) GoName() (name string) {
	switch p.Name {
	case "type", "range", "map", "string", "select", "interface":
		name = p.Name + "_"
	default:
		name = p.Name
	}
	if p.Object {
		name = name + "Object"
	}
	return
}
