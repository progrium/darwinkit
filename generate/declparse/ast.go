package declparse

import (
	"fmt"
	"strings"
)

type Node interface {
	String() string
}

type Statement struct {
	Method    *MethodDecl
	Property  *PropertyDecl
	Interface *InterfaceDecl
	Protocol  *ProtocolDecl
	Function  *FunctionDecl
	Variable  *VariableDecl
	Enum      *EnumDecl
	Struct    *StructDecl
	TypeAlias *TypeInfo
	Typedef   string
}

type ProtocolDecl struct {
	Name      string
	SuperName string
}

type InterfaceDecl struct {
	Name      string
	SuperName string
}

type PropertyDecl struct {
	Name     string
	Type     TypeInfo
	Attrs    map[PropAttr]string
	IsOutlet bool
}

type FunctionDecl struct {
	Name       string
	ReturnType TypeInfo
	Args       FuncArgs
	IsBlock    bool
	IsPtr      bool
	Variadic   bool
}

func (f *FunctionDecl) Ident() string {
	if f.IsBlock {
		return fmt.Sprintf("(^%s)", f.Name)
	}
	if f.IsPtr {
		return fmt.Sprintf("(*%s)", f.Name)
	}
	return f.Name
}

type FuncArgs []ArgInfo

type MethodDecl struct {
	TypeMethod  bool // instance method otherwise
	ReturnType  TypeInfo
	NameParts   []string
	Args        []ArgInfo
	Variadic    bool
	Unavailable bool
}

func (m *MethodDecl) Name() string {
	if len(m.NameParts) == 0 {
		return ""
	}
	if len(m.NameParts) == 1 && len(m.Args) == 0 {
		return m.NameParts[0]
	}
	return strings.Join(append(m.NameParts, ""), ":")
}

type TypeInfo struct {
	Name     string
	IsPtr    bool
	IsPtrPtr bool
	Annots   map[TypeAnnotation]bool
	Func     *FunctionDecl
	Params   []TypeInfo
}

type ArgInfo struct {
	Name string
	Type TypeInfo
}

type VariableDecl struct {
	Name  string
	Type  TypeInfo
	Value string
}

type EnumDecl struct {
	Name  string
	Type  TypeInfo
	Cases []VariableDecl
}

type StructDecl struct {
	Name   string
	Fields []VariableDecl
}
