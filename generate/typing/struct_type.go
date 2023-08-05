package typing

import "C"
import (
	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/internal/set"
)

var _ Type = (*StructType)(nil)

// StructType struct type
type StructType struct {
	Name   string          // c and objc type name
	GName  string          // the go struct name
	Module *modules.Module // the module
}

func (s *StructType) GoImports() set.Set[string] {
	return set.New("github.com/progrium/macdriver/macos/" + s.Module.Package)
}

func (s *StructType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	return FullGoName(*s.Module, s.GName, *currentModule)
}

func (s *StructType) ObjcName() string {
	return s.Name
}

func (s *StructType) DeclareModule() *modules.Module {
	return s.Module
}
