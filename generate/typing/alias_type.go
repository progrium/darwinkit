package typing

import (
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/internal/set"
)

var _ Type = (*AliasType)(nil)

// AliasType type def types
type AliasType struct {
	Type   Type            // the real type
	GName  string          // alias name for new go type
	Name   string          // the objc name
	Module *modules.Module // used when Alias is not empty
}

func (a *AliasType) GoImports() set.Set[string] {
	var imports = set.New[string]()
	if a.Module != nil {
		imports.Add("github.com/progrium/darwinkit/macos/" + a.Module.Package)
	}
	imports.AddSet(a.Type.GoImports())
	return imports
}

func (a *AliasType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	mod := *a.Module
	if currentModule.Package == "appkit" && mod.Package == "uikit" {
		mod = *currentModule
	}
	return FullGoName(mod, a.GName, *currentModule)
}

func (a *AliasType) ObjcName() string {
	return a.Name
}

func (a *AliasType) CName() string {
	return a.Name
}

func (a *AliasType) CSignature() string {
	return a.Name
}

func (a *AliasType) DeclareModule() *modules.Module {
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
