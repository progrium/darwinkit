package typing

import "C"
import (
	"github.com/progrium/macdriver/internal/set"
)

var _ Type = (*StructType)(nil)

// StructType struct type
type StructType struct {
	Name   string  // c and objc type name
	GName  string  // the go struct name
	Module *Module // the module
}

func (s *StructType) GoImports() set.Set[string] {
	return set.New("github.com/progrium/macdriver/macos/" + s.Module.Package)
}

func (s *StructType) GoName(currentModule *Module, receiveFromObjc bool) string {
	return FullGoName(*s.Module, s.GName, *currentModule)
}

func (s *StructType) ObjcName() string {
	return s.Name
}

func (s *StructType) DeclareModule() *Module {
	return s.Module
}
