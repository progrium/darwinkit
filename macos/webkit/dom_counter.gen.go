// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMCounter] class.
var DOMCounterClass = _DOMCounterClass{objc.GetClass("DOMCounter")}

type _DOMCounterClass struct {
	objc.Class
}

// An interface definition for the [DOMCounter] class.
type IDOMCounter interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcounter?language=objc
type DOMCounter struct {
	DOMObject
}

func DOMCounterFrom(ptr unsafe.Pointer) DOMCounter {
	return DOMCounter{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMCounterClass) Alloc() DOMCounter {
	rv := objc.Call[DOMCounter](dc, objc.Sel("alloc"))
	return rv
}

func DOMCounter_Alloc() DOMCounter {
	return DOMCounterClass.Alloc()
}

func (dc _DOMCounterClass) New() DOMCounter {
	rv := objc.Call[DOMCounter](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMCounter() DOMCounter {
	return DOMCounterClass.New()
}

func (d_ DOMCounter) Init() DOMCounter {
	rv := objc.Call[DOMCounter](d_, objc.Sel("init"))
	return rv
}
