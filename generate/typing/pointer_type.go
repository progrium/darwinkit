package typing

import (
	"github.com/progrium/macdriver/internal/set"
)

var _ Type = (*PointerType)(nil)

// PointerType is c pointer type
type PointerType struct {
	Type Type
}

func (c *PointerType) GoImports() set.Set[string] {
	return c.Type.GoImports()
}

func (c *PointerType) GoName(currentModule *Module, receiveFromObjc bool) string {
	switch UnwrapAlias(c.Type).(type) {
	case *StructType:
		return "*" + c.Type.GoName(currentModule, receiveFromObjc)
	case *PrimitiveType:
		return "*" + c.Type.GoName(currentModule, receiveFromObjc)
	case *ClassType:
		return "*" + c.Type.GoName(currentModule, true)
	case *StringType:
		return "*" + (&ClassType{Name: "NSString", GName: "String", Module: Foundation}).GoName(currentModule, true)
	case *DataType:
		return "*" + (&ClassType{Name: "NSData", GName: "Data", Module: Foundation}).GoName(currentModule, true)
	case *ArrayType:
		return "*" + (&ClassType{Name: "NSArray", GName: "Array", Module: Foundation}).GoName(currentModule, true)
	case *DictType:
		return "*" + (&ClassType{Name: "NSDictionary", GName: "Dictionary", Module: Foundation}).GoName(currentModule, true)
	default:
		panic("not supported pointer to: " + c.Type.ObjcName())
	}

}

func (c *PointerType) ObjcName() string {
	return c.Type.ObjcName() + "*"
}

func (c *PointerType) DeclareModule() *Module {
	return c.Type.DeclareModule()
}
