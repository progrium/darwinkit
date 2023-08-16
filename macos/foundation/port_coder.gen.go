// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PortCoder] class.
var PortCoderClass = _PortCoderClass{objc.GetClass("NSPortCoder")}

type _PortCoderClass struct {
	objc.Class
}

// An interface definition for the [PortCoder] class.
type IPortCoder interface {
	ICoder
}

// A coder used to transmit object proxies (and sometimes objects themselves) between connections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsportcoder?language=objc
type PortCoder struct {
	Coder
}

func PortCoderFrom(ptr unsafe.Pointer) PortCoder {
	return PortCoder{
		Coder: CoderFrom(ptr),
	}
}

func (pc _PortCoderClass) Alloc() PortCoder {
	rv := objc.Call[PortCoder](pc, objc.Sel("alloc"))
	return rv
}

func PortCoder_Alloc() PortCoder {
	return PortCoderClass.Alloc()
}

func (pc _PortCoderClass) New() PortCoder {
	rv := objc.Call[PortCoder](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPortCoder() PortCoder {
	return PortCoderClass.New()
}

func (p_ PortCoder) Init() PortCoder {
	rv := objc.Call[PortCoder](p_, objc.Sel("init"))
	return rv
}
