package codegen

import (
	"github.com/progrium/macdriver/generate/typing"
	"github.com/progrium/macdriver/internal/stringx"
)

// Property is code generator for objective-c property
type Property struct {
	Name          string
	GoName        string
	ReadOnly      bool
	GetterName    string
	Type          typing.Type
	ClassProperty bool
	Weak          bool
	Deprecated    bool
	Required      bool //for protocol
}

func (p *Property) String() string {
	return p.Name
}

func (p *Property) getter() *Method {
	methodName := p.GetterName
	if len(methodName) == 0 {
		methodName = p.Name
	}
	goMethodName := p.GetterName
	if len(goMethodName) == 0 {
		goMethodName = p.GoName
	}
	return &Method{
		Name:         methodName,
		GoName:       goMethodName,
		ReturnType:   p.Type,
		WeakProperty: p.Weak,
		ClassMethod:  p.ClassProperty,
		Deprecated:   p.Deprecated,
	}
}

func (p *Property) setter() *Method {
	return &Method{
		Name:         "set" + stringx.Capitalize(p.Name),
		GoName:       "set" + stringx.Capitalize(p.GoName),
		Params:       []*Param{{Name: "value", Type: p.Type}},
		ReturnType:   &typing.VoidType{},
		WeakProperty: p.Weak,
		ClassMethod:  p.ClassProperty,
		Deprecated:   p.Deprecated,
	}
}
