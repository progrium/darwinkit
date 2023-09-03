package typing

import (
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/internal/set"
)

var _ Type = (*StringType)(nil)

// StringType string
type StringType struct {
	NeedNil bool // string type need nil value.If set to true, will use foundation.String instead string for generated go code.
}

func (s *StringType) GoImports() set.Set[string] {
	return set.New("github.com/progrium/darwinkit/macos/foundation")
}

func (s *StringType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	if s.NeedNil {
		return "foundation.String"
	}
	return "string"
}

func (s *StringType) ObjcName() string {
	return "NSString*"
}

func (s *StringType) CName() string {
	return "NSString*"
}

func (s *StringType) DeclareModule() *modules.Module {
	return modules.Get("foundation")
}
