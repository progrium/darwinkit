package typing

import (
	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/internal/set"
)

var _ Type = (*SelectorType)(nil)

// SelectorType objc selector type
type SelectorType struct {
}

func (s *SelectorType) GoImports() set.Set[string] {
	return set.New("github.com/progrium/macdriver/objc")
}

func (s *SelectorType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	if currentModule.Package == "objc" {
		return "Selector"
	} else {
		return "objc.Selector"
	}
}

func (s *SelectorType) ObjcName() string {
	return "SEL"
}

func (s *SelectorType) DeclareModule() *modules.Module {
	return modules.Get("objc")
}
