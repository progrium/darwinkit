package typing

import (
	"github.com/progrium/macdriver/internal/set"
)

var _ Type = (*PrimitiveType)(nil)

var Bool = &PrimitiveType{GoName_: "bool", ObjcName_: "BOOL"}

var Int = &PrimitiveType{GoName_: "int", ObjcName_: "NSInteger"}
var UInt = &PrimitiveType{GoName_: "uint", ObjcName_: "NSUInteger"}

var Float = &PrimitiveType{GoName_: "float32", ObjcName_: "float"}
var Double = &PrimitiveType{GoName_: "float64", ObjcName_: "double"}

var Int8 = &PrimitiveType{GoName_: "int8", ObjcName_: "int8_t"}
var UInt8 = &PrimitiveType{GoName_: "uint8", ObjcName_: "uint8_t"}
var Byte = &PrimitiveType{GoName_: "byte", ObjcName_: "char"}

var Int16 = &PrimitiveType{GoName_: "int16", ObjcName_: "int16_t"}
var UInt16 = &PrimitiveType{GoName_: "uint16", ObjcName_: "uint16_t"}

var Int32 = &PrimitiveType{GoName_: "int32", ObjcName_: "int32_t"}
var UInt32 = &PrimitiveType{GoName_: "uint32", ObjcName_: "uint32_t"}

var Int64 = &PrimitiveType{GoName_: "int64", ObjcName_: "int64_t"}
var UInt64 = &PrimitiveType{GoName_: "uint64", ObjcName_: "uint64_t"}

var primitiveMap map[string]*PrimitiveType

func init() {
	primitiveMap = map[string]*PrimitiveType{}
	for _, pt := range []*PrimitiveType{Bool, Int, UInt, Float, Double, Int8, UInt8, Byte, Int16, UInt16, Int32, UInt32, Int64, UInt64} {
		primitiveMap[pt.ObjcName_] = pt
	}
}

func GetPrimitiveType(typeName string) (*PrimitiveType, bool) {
	pt, ok := primitiveMap[typeName]
	return pt, ok
}

// PrimitiveType primitive types
type PrimitiveType struct {
	GoName_   string // go type name
	ObjcName_ string // objc type name
}

func (p *PrimitiveType) GoImports() set.Set[string] {
	return nil
}

func (p *PrimitiveType) GoName(currentModule *Module, receiveFromObjc bool) string {
	return p.GoName_
}

func (p *PrimitiveType) ObjcName() string {
	return p.ObjcName_
}

func (p *PrimitiveType) DeclareModule() *Module {
	return nil
}

func (p *PrimitiveType) ToUnsigned() *PrimitiveType {
	switch p.GoName_ {
	case "int":
		return UInt
	case "int8":
		return UInt8
	case "int16":
		return UInt16
	case "int32":
		return UInt32
	case "int64":
		return UInt64
	default:
		return p
	}
}
