// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ProcessPool] class.
var ProcessPoolClass = _ProcessPoolClass{objc.GetClass("WKProcessPool")}

type _ProcessPoolClass struct {
	objc.Class
}

// An interface definition for the [ProcessPool] class.
type IProcessPool interface {
	objc.IObject
}

// An opaque token that you use to run multiple web views in a single process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkprocesspool?language=objc
type ProcessPool struct {
	objc.Object
}

func ProcessPoolFrom(ptr unsafe.Pointer) ProcessPool {
	return ProcessPool{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _ProcessPoolClass) Alloc() ProcessPool {
	rv := objc.Call[ProcessPool](pc, objc.Sel("alloc"))
	return rv
}

func ProcessPool_Alloc() ProcessPool {
	return ProcessPoolClass.Alloc()
}

func (pc _ProcessPoolClass) New() ProcessPool {
	rv := objc.Call[ProcessPool](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewProcessPool() ProcessPool {
	return ProcessPoolClass.New()
}

func (p_ ProcessPool) Init() ProcessPool {
	rv := objc.Call[ProcessPool](p_, objc.Sel("init"))
	return rv
}
