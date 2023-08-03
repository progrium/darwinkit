package typing

import "C"
import (
	"github.com/progrium/macdriver/internal/set"
)

var _ Type = (*DictType)(nil)

// DictType for objc Dictionary types, the element should be StringType or ClassType
type DictType struct {
	KeyType   Type
	ValueType Type
}

func (d *DictType) GoImports() set.Set[string] {
	var imports = set.New("github.com/progrium/macdriver/macos/foundation")
	imports.AddSet(d.KeyType.GoImports())
	imports.AddSet(d.ValueType.GoImports())
	return imports
}

func (d *DictType) GoName(currentModule *Module, receiveFromObjc bool) string {
	if _, ok := UnwrapAlias(d.KeyType).(*StringType); ok {
		if _, ok = d.ValueType.(*ProtocolType); !ok {
			return "map[" + d.KeyType.GoName(currentModule, receiveFromObjc) + "]" + d.ValueType.GoName(currentModule, receiveFromObjc)
		}
	}
	name := "Dictionary"
	if !receiveFromObjc {
		name = "Dictionary"
	}
	return FullGoName(*Foundation, name, *currentModule)
}

func (d *DictType) ObjcName() string {
	return "NSDictionary*"
}

func (d *DictType) DeclareModule() *Module {
	return d.ValueType.DeclareModule()
}
