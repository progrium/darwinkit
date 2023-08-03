// auto-generated code, do not modify
package uniformtypeidentifiers

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var TypeClass = _TypeClass{objc.GetClass("UTType")}

type _TypeClass struct {
	objc.Class
}

type IType interface {
	objc.IObject
}

type Type struct {
	objc.Object
}

func MakeType(ptr unsafe.Pointer) Type {
	return Type{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TypeClass) Alloc() Type {
	rv := objc.CallMethod[Type](tc, objc.GetSelector("alloc"))
	return rv
}

func Type_Alloc() Type {
	return TypeClass.Alloc()
}

func (tc _TypeClass) New() Type {
	rv := objc.CallMethod[Type](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewType() Type {
	return TypeClass.New()
}

func Type_New() Type {
	return TypeClass.New()
}

func (t_ Type) Init() Type {
	rv := objc.CallMethod[Type](t_, objc.GetSelector("init"))
	return rv
}

func Type_Init() Type {
	return TypeClass.Alloc().Init()
}
