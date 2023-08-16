// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLFrameElement] class.
var DOMHTMLFrameElementClass = _DOMHTMLFrameElementClass{objc.GetClass("DOMHTMLFrameElement")}

type _DOMHTMLFrameElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLFrameElement] class.
type IDOMHTMLFrameElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlframeelement?language=objc
type DOMHTMLFrameElement struct {
	DOMHTMLElement
}

func DOMHTMLFrameElementFrom(ptr unsafe.Pointer) DOMHTMLFrameElement {
	return DOMHTMLFrameElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLFrameElementClass) Alloc() DOMHTMLFrameElement {
	rv := objc.Call[DOMHTMLFrameElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLFrameElement_Alloc() DOMHTMLFrameElement {
	return DOMHTMLFrameElementClass.Alloc()
}

func (dc _DOMHTMLFrameElementClass) New() DOMHTMLFrameElement {
	rv := objc.Call[DOMHTMLFrameElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLFrameElement() DOMHTMLFrameElement {
	return DOMHTMLFrameElementClass.New()
}

func (d_ DOMHTMLFrameElement) Init() DOMHTMLFrameElement {
	rv := objc.Call[DOMHTMLFrameElement](d_, objc.Sel("init"))
	return rv
}
