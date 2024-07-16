package typing

import (
	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/internal/set"
)

var Bool = &PrimitiveType{GoName_: "bool", ObjcName_: "BOOL", CName_: "bool"}
var Boolean = &PrimitiveType{GoName_: "bool", ObjcName_: "Boolean", CName_: "bool"}

var Int = &PrimitiveType{GoName_: "int", ObjcName_: "NSInteger", CName_: "long"}
var UInt = &PrimitiveType{GoName_: "uint", ObjcName_: "NSUInteger", CName_: "uint"}

var Float = &PrimitiveType{GoName_: "float32", ObjcName_: "float"}
var Double = &PrimitiveType{GoName_: "float64", ObjcName_: "double"}

var Int8 = &PrimitiveType{GoName_: "int8", ObjcName_: "int8_t"}
var UInt8 = &PrimitiveType{GoName_: "uint8", ObjcName_: "uint8_t"}
var Byte = &PrimitiveType{GoName_: "byte", ObjcName_: "char"}
var OffT = &PrimitiveType{GoName_: "float64", ObjcName_: "off_t"}

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
	primitiveMap["char"] = UInt8
	primitiveMap["bool"] = Bool
	primitiveMap["int"] = Int
	primitiveMap["short"] = Int
	primitiveMap["long"] = Int32
	primitiveMap["long long"] = Int64
	primitiveMap["size_t"] = UInt
	primitiveMap["uintptr_t"] = UInt

	primitiveMap["Boolean"] = Boolean
	primitiveMap["SInt8"] = Int8
	primitiveMap["SInt16"] = Int16
	primitiveMap["SInt32"] = Int32
	primitiveMap["SInt64"] = Int64
	primitiveMap["UInt8"] = UInt8
	primitiveMap["UInt16"] = UInt16
	primitiveMap["UInt32"] = UInt32
	primitiveMap["UInt64"] = UInt64

	primitiveMap["Float32"] = Float
	primitiveMap["Float64"] = Double
}

func GetPrimitiveType(typeName string) (*PrimitiveType, bool) {
	pt, ok := primitiveMap[typeName]
	return pt, ok
}

// PrimitiveType primitive types
type PrimitiveType struct {
	GoName_   string // go type name
	ObjcName_ string // objc type name
	CName_    string
}

func (p *PrimitiveType) GoImports() set.Set[string] {
	return nil
}

func (p *PrimitiveType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	return p.GoName_
}

func (p *PrimitiveType) ObjcName() string {
	return p.ObjcName_
}

func (p *PrimitiveType) CName() string {
	n := p.CName_
	if n == "" {
		return p.ObjcName_
	}
	return n
}

func (p *PrimitiveType) DeclareModule() *modules.Module {
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
