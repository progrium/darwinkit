package typing

import "C"
import (
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/internal/set"
)

// for weird struct refs like those ending in "Ref"
type RefType struct {
	Name   string          // c and objc type name
	GName  string          // the go struct name
	Module *modules.Module // the module
}

func (s *RefType) GoImports() set.Set[string] {
	if s.Module == nil {
		return set.New("unsafe")
	}
	return set.New("github.com/progrium/darwinkit/macos/" + s.Module.Package)
}

func (s *RefType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	if s.Module == nil {
		return "unsafe.Pointer"
	}
	return FullGoName(*s.Module, s.GName, *currentModule)
}

func (s *RefType) ObjcName() string {
	return s.Name
}

func (s *RefType) CName() string {
	return "void *"
}

func (s *RefType) DeclareModule() *modules.Module {
	return s.Module
}
