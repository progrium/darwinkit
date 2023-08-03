package typing

import (
	"github.com/progrium/macdriver/internal/set"
)

var _ Type = (*StringType)(nil)

// StringType string
type StringType struct {
	NeedNil bool // string type need nil value.If set to true, will use foundation.String instread string for generated go code.
}

func (s *StringType) GoImports() set.Set[string] {
	return set.New("github.com/progrium/macdriver/macos/foundation")
}

func (s *StringType) GoName(currentModule *Module, receiveFromObjc bool) string {
	if s.NeedNil {
		return "foundation.String"
	}
	return "string"
}

func (s *StringType) ObjcName() string {
	return "NSString*"
}

func (s *StringType) DeclareModule() *Module {
	return Foundation
}
