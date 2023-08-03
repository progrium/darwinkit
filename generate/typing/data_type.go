package typing

import (
	"github.com/progrium/macdriver/internal/set"
)

var _ Type = (*DataType)(nil)

// DataType objc binary data type
type DataType struct {
}

func (d *DataType) GoImports() set.Set[string] {
	return nil
}

func (d *DataType) GoName(currentModule *Module, receiveFromObjc bool) string {
	return "[]byte"
}

func (d *DataType) ObjcName() string {
	return "NSData*"
}

func (d *DataType) DeclareModule() *Module {
	return nil
}
