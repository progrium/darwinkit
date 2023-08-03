package typing

import (
	"github.com/progrium/macdriver/internal/set"
)

var _ Type = (*AliasType)(nil)

// AliasType type def types
type AliasType struct {
	Type   Type    // the real type
	GName  string  // alias name for new go type
	Name   string  // the objc name
	Module *Module // used when Alias is not empty
}

func (a *AliasType) GoImports() set.Set[string] {
	var imports = set.New[string]()
	if a.Module != nil {
		imports.Add("github.com/progrium/macdriver/macos/" + a.Module.Package)
	}
	imports.AddSet(a.Type.GoImports())
	return imports
}

func (a *AliasType) GoName(currentModule *Module, receiveFromObjc bool) string {
	return FullGoName(*a.Module, a.GName, *currentModule)
}

func (a *AliasType) ObjcName() string {
	return a.Name
}

func (a *AliasType) DeclareModule() *Module {
	return a.Module
}

// struct alias use go type alias, do not need conversion
func (a *AliasType) isStructAlias() bool {
	_, ok := a.Type.(*StructType)
	return ok
}

// UnwrapAlias unwrap alias type to it's real underlying type
func UnwrapAlias(t Type) Type {
	for true {
		at, ok := t.(*AliasType)
		if !ok {
			return t
		}
		t = at.Type
	}
	panic("should not happen")
}
