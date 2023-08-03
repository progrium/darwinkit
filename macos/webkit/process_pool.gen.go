// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ProcessPoolClass = _ProcessPoolClass{objc.GetClass("WKProcessPool")}

type _ProcessPoolClass struct {
	objc.Class
}

type IProcessPool interface {
	objc.IObject
}

type ProcessPool struct {
	objc.Object
}

func MakeProcessPool(ptr unsafe.Pointer) ProcessPool {
	return ProcessPool{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _ProcessPoolClass) Alloc() ProcessPool {
	rv := objc.CallMethod[ProcessPool](pc, objc.GetSelector("alloc"))
	return rv
}

func ProcessPool_Alloc() ProcessPool {
	return ProcessPoolClass.Alloc()
}

func (pc _ProcessPoolClass) New() ProcessPool {
	rv := objc.CallMethod[ProcessPool](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewProcessPool() ProcessPool {
	return ProcessPoolClass.New()
}

func ProcessPool_New() ProcessPool {
	return ProcessPoolClass.New()
}

func (p_ ProcessPool) Init() ProcessPool {
	rv := objc.CallMethod[ProcessPool](p_, objc.GetSelector("init"))
	return rv
}

func ProcessPool_Init() ProcessPool {
	return ProcessPoolClass.Alloc().Init()
}
