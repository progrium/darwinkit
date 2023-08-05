// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var CoderClass = _CoderClass{objc.GetClass("NSCoder")}

type _CoderClass struct {
	objc.Class
}

type ICoder interface {
	objc.IObject
}

type Coder struct {
	objc.Object
}

func MakeCoder(ptr unsafe.Pointer) Coder {
	return Coder{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CoderClass) Alloc() Coder {
	rv := objc.CallMethod[Coder](cc, objc.GetSelector("alloc"))
	return rv
}

func Coder_Alloc() Coder {
	return CoderClass.Alloc()
}

func (cc _CoderClass) New() Coder {
	rv := objc.CallMethod[Coder](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCoder() Coder {
	return CoderClass.New()
}

func Coder_New() Coder {
	return CoderClass.New()
}

func (c_ Coder) Init() Coder {
	rv := objc.CallMethod[Coder](c_, objc.GetSelector("init"))
	return rv
}

func Coder_Init() Coder {
	return CoderClass.Alloc().Init()
}
