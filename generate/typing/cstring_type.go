package typing

import (
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/internal/set"
)

var _ Type = (*CStringType)(nil)

// CStringType string
type CStringType struct {
	IsConst bool
}

func (s *CStringType) GoImports() set.Set[string] {
	return nil
}

func (s *CStringType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	return "string"
}

func (s *CStringType) ObjcName() string {
	return "char*"
}

func (s *CStringType) CName() string {
	return "char*"
}

func (c *CStringType) CSignature() string {
	t := c.CName()
	if c.IsConst {
		t = "const " + t
	}
	return t
}

func (s *CStringType) DeclareModule() *modules.Module {
	return nil
}
