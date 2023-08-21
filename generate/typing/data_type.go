package typing

import (
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/internal/set"
)

var _ Type = (*DataType)(nil)

// DataType objc binary data type
type DataType struct {
}

func (d *DataType) GoImports() set.Set[string] {
	return nil
}

func (d *DataType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	return "[]byte"
}

func (d *DataType) ObjcName() string {
	return "NSData*"
}

func (d *DataType) CName() string {
	return "NSData*"
}

func (d *DataType) DeclareModule() *modules.Module {
	return nil
}
