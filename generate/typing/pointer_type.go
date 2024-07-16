package typing

import (
	"fmt"

	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/internal/set"
)

var _ Type = (*PointerType)(nil)

var objCToCMap = map[string]string{
	"NSInteger":  "int",
	"NSUInteger": "uint",
}

// PointerType is c pointer type
type PointerType struct {
	Type    Type
	IsConst bool
}

func (c *PointerType) GoImports() set.Set[string] {
	return c.Type.GoImports()
}

func (c *PointerType) GoName(currentModule *modules.Module, receiveFromObjc bool) string {
	switch UnwrapAlias(c.Type).(type) {
	case *StructType, *PrimitiveType, *KernelType:
		return "*" + c.Type.GoName(currentModule, receiveFromObjc)
	case *ClassType:
		return "*" + c.Type.GoName(currentModule, true)
	case *StringType:
		return "*" + (&ClassType{Name: "NSString", GName: "String", Module: modules.Get("foundation")}).GoName(currentModule, true)
	case *DataType:
		return "*" + (&ClassType{Name: "NSData", GName: "Data", Module: modules.Get("foundation")}).GoName(currentModule, true)
	case *ArrayType:
		return "*" + (&ClassType{Name: "NSArray", GName: "Array", Module: modules.Get("foundation")}).GoName(currentModule, true)
	case *DictType:
		return "*" + (&ClassType{Name: "NSDictionary", GName: "Dictionary", Module: modules.Get("foundation")}).GoName(currentModule, true)
	case *PointerType:
		return "*" + c.Type.GoName(currentModule, receiveFromObjc)
	case *RefType:
		return "*" + c.Type.GoName(currentModule, receiveFromObjc)
	case *ProtocolType:
		return "*" + c.Type.GoName(currentModule, receiveFromObjc)
	default:
		panic(fmt.Sprintf("not supported pointer to: %T %v", c.Type, c.Type.ObjcName()))
	}

}

func (c *PointerType) ObjcName() string {
	return c.Type.ObjcName() + "*"
}

func (c *PointerType) CName() string {
	switch tt := c.Type.(type) {
	case *VoidPointerType:
		return ""
	case *RefType:
		return tt.CName() + "*"
	case *PointerType:
		return tt.Type.CName()
	default:
		return c.Type.CName()
	}
}

type hasCSignature interface {
	CSignature() string
}

func (c *PointerType) CSignature() string {
	t := c.Type.CName() + "*"
	if sig, ok := c.Type.(hasCSignature); ok {
		t = sig.CSignature() + "*"
	}
	if v, ok := objCToCMap[t]; ok {
		t = v
	}
	if c.IsConst {
		t = "const " + t
	}
	return t
}

func (c *PointerType) DeclareModule() *modules.Module {
	return c.Type.DeclareModule()
}
