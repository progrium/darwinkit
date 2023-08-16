// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMRect] class.
var DOMRectClass = _DOMRectClass{objc.GetClass("DOMRect")}

type _DOMRectClass struct {
	objc.Class
}

// An interface definition for the [DOMRect] class.
type IDOMRect interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domrect?language=objc
type DOMRect struct {
	DOMObject
}

func DOMRectFrom(ptr unsafe.Pointer) DOMRect {
	return DOMRect{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMRectClass) Alloc() DOMRect {
	rv := objc.Call[DOMRect](dc, objc.Sel("alloc"))
	return rv
}

func DOMRect_Alloc() DOMRect {
	return DOMRectClass.Alloc()
}

func (dc _DOMRectClass) New() DOMRect {
	rv := objc.Call[DOMRect](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMRect() DOMRect {
	return DOMRectClass.New()
}

func (d_ DOMRect) Init() DOMRect {
	rv := objc.Call[DOMRect](d_, objc.Sel("init"))
	return rv
}
