// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMRange] class.
var DOMRangeClass = _DOMRangeClass{objc.GetClass("DOMRange")}

type _DOMRangeClass struct {
	objc.Class
}

// An interface definition for the [DOMRange] class.
type IDOMRange interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domrange?language=objc
type DOMRange struct {
	DOMObject
}

func DOMRangeFrom(ptr unsafe.Pointer) DOMRange {
	return DOMRange{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMRangeClass) Alloc() DOMRange {
	rv := objc.Call[DOMRange](dc, objc.Sel("alloc"))
	return rv
}

func DOMRange_Alloc() DOMRange {
	return DOMRangeClass.Alloc()
}

func (dc _DOMRangeClass) New() DOMRange {
	rv := objc.Call[DOMRange](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMRange() DOMRange {
	return DOMRangeClass.New()
}

func (d_ DOMRange) Init() DOMRange {
	rv := objc.Call[DOMRange](d_, objc.Sel("init"))
	return rv
}
