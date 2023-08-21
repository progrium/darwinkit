package typing

import (
	"strings"
	"unicode"

	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/internal/set"
)

func GetDispatchType(typeName string) (Type, bool) {
	if typeName == "dispatch_block_t" {
		return &BlockType{
			ReturnType: &VoidType{},
		}, true
	}
	for _, name := range []string{
		"dispatch_queue_t",
		"dispatch_data_t",
	} {
		if typeName == name {
			return &DispatchType{ObjcName_: typeName}, true
		}
	}
	return nil, false
}

type DispatchType struct {
	ObjcName_ string // objc type name
}

func (d *DispatchType) GoImports() set.Set[string] {
	return set.New("github.com/progrium/darwinkit/dispatch")
}

func (d *DispatchType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	name := strings.TrimPrefix(d.ObjcName_, "dispatch_")
	name = strings.TrimSuffix(name, "_t")
	r := []rune(name)
	r[0] = unicode.ToUpper(r[0])
	return FullGoName(*d.DeclareModule(), string(r), *currentModule)
}

func (d *DispatchType) ObjcName() string {
	return d.ObjcName_
}

func (d *DispatchType) CName() string {
	return d.ObjcName_
}

func (d *DispatchType) DeclareModule() *modules.Module {
	return modules.Get("dispatch")
}
