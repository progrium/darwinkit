// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var NullClass = _NullClass{objc.GetClass("NSNull")}

type _NullClass struct {
	objc.Class
}

type INull interface {
	objc.IObject
}

type Null struct {
	objc.Object
}

func MakeNull(ptr unsafe.Pointer) Null {
	return Null{
		Object: objc.MakeObject(ptr),
	}
}

func (nc _NullClass) Alloc() Null {
	rv := objc.CallMethod[Null](nc, objc.GetSelector("alloc"))
	return rv
}

func Null_Alloc() Null {
	return NullClass.Alloc()
}

func (nc _NullClass) New() Null {
	rv := objc.CallMethod[Null](nc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewNull() Null {
	return NullClass.New()
}

func Null_New() Null {
	return NullClass.New()
}

func (n_ Null) Init() Null {
	rv := objc.CallMethod[Null](n_, objc.GetSelector("init"))
	return rv
}

func Null_Init() Null {
	return NullClass.Alloc().Init()
}

func (nc _NullClass) Null() Null {
	rv := objc.CallMethod[Null](nc, objc.GetSelector("null"))
	return rv
}

func Null_Null() Null {
	return NullClass.Null()
}
