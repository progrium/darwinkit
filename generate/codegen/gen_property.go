package codegen

import (
	"strings"

	"github.com/progrium/darwinkit/generate/typing"
	"github.com/progrium/darwinkit/internal/stringx"
)

// Property is code generator for objective-c property
type Property struct {
	Name          string
	GoName        string
	ReadOnly      bool
	GetterName    string
	SetterName    string
	Type          typing.Type
	ClassProperty bool
	Weak          bool
	Deprecated    bool
	Required      bool //for protocol
	Description   string
	DocURL        string
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
		Description:  p.Description,
		DocURL:       p.DocURL,
	}
}

func (p *Property) setter() *Method {
	name := "set" + stringx.Capitalize(p.Name)
	if p.SetterName != "" {
		name = strings.TrimSuffix(p.SetterName, ":")
	}
	goName := "set" + stringx.Capitalize(p.GoName)
	if p.SetterName != "" {
		goName = strings.TrimSuffix(p.SetterName, ":")
	}
	return &Method{
		Name:         name,
		GoName:       goName,
		Params:       []*Param{{Name: "value", Type: p.Type, FieldName: goName}},
		ReturnType:   &typing.VoidType{},
		WeakProperty: p.Weak,
		ClassMethod:  p.ClassProperty,
		Deprecated:   p.Deprecated,
		Description:  p.Description,
		DocURL:       p.DocURL,
	}
}
