package typing

import (
	"strings"

	"github.com/progrium/macdriver/internal/set"
)

var _ Type = (*BlockType)(nil)

type BlockParam struct {
	Type Type
	Name string // the param name
}

// BlockType
type BlockType struct {
	ReturnType Type
	Params     []BlockParam
}

func (a *BlockType) GoImports() set.Set[string] {
	imports := set.New("github.com/progrium/macdriver/objc")
	imports.AddSet(a.ReturnType.GoImports())
	for _, p := range a.Params {
		imports.AddSet(p.Type.GoImports())
	}
	return imports
}

func (a *BlockType) GoName(currentModule *Module, receiveFromObjc bool) string {
	var sb strings.Builder
	sb.WriteString("func (")
	for i, p := range a.Params {
		if i > 0 {
			sb.WriteString(", ")
		}
		switch p.Name {
		case "range":
			sb.WriteString(p.Name + "_")
		default:
			sb.WriteString(p.Name)
		}
		sb.WriteString(" ")
		sb.WriteString(p.Type.GoName(currentModule, true))
	}
	sb.WriteString(")")
	if _, ok := a.ReturnType.(*VoidType); !ok {
		sb.WriteByte(' ')
		sb.WriteString(a.ReturnType.GoName(currentModule, true))
	}
	return sb.String()
}

func (a *BlockType) ObjcName() string {
	var sb strings.Builder
	sb.WriteString(a.ReturnType.ObjcName())
	sb.WriteString(" (^)(")
	for i, p := range a.Params {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(p.Type.ObjcName())
		sb.WriteString(" ")
		sb.WriteString(p.Name)
	}
	sb.WriteString(")")
	return sb.String()
}

func (a *BlockType) DeclareModule() *Module {
	return nil
}
