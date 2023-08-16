package typing

import "C"
import (
	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/internal/set"
)

// for weird struct refs like those ending in "Ref"
type RefType struct {
	Name string // c and objc type name
	// GName  string          // the go struct name
	// Module *modules.Module // the module
}

func (s *RefType) GoImports() set.Set[string] {
	return set.New("unsafe")
}

func (s *RefType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	return "unsafe.Pointer"
}

func (s *RefType) ObjcName() string {
	return s.Name
}

func (s *RefType) DeclareModule() *modules.Module {
	return nil
}
