package typing

import (
	"fmt"
	"strings"

	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/internal/set"
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
	imports := set.New("github.com/progrium/darwinkit/objc")
	imports.AddSet(a.ReturnType.GoImports())
	for _, p := range a.Params {
		imports.AddSet(p.Type.GoImports())
	}
	return imports
}

func (a *BlockType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	var sb strings.Builder
	sb.WriteString("func (")
	for i, p := range a.Params {
		if i > 0 {
			sb.WriteString(", ")
		}
		if _, ok := p.Type.(*VoidType); ok {
			continue
		}
		switch p.Name {
		case "range", "type":
			sb.WriteString(p.Name + "_")
		case "":
			sb.WriteString(fmt.Sprintf("arg%d", i))
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

func (a *BlockType) CName() string {
	return "implement me"
}

func (a *BlockType) DeclareModule() *modules.Module {
	return nil
}
