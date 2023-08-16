// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLFrameSetElement] class.
var DOMHTMLFrameSetElementClass = _DOMHTMLFrameSetElementClass{objc.GetClass("DOMHTMLFrameSetElement")}

type _DOMHTMLFrameSetElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLFrameSetElement] class.
type IDOMHTMLFrameSetElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlframesetelement?language=objc
type DOMHTMLFrameSetElement struct {
	DOMHTMLElement
}

func DOMHTMLFrameSetElementFrom(ptr unsafe.Pointer) DOMHTMLFrameSetElement {
	return DOMHTMLFrameSetElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLFrameSetElementClass) Alloc() DOMHTMLFrameSetElement {
	rv := objc.Call[DOMHTMLFrameSetElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLFrameSetElement_Alloc() DOMHTMLFrameSetElement {
	return DOMHTMLFrameSetElementClass.Alloc()
}

func (dc _DOMHTMLFrameSetElementClass) New() DOMHTMLFrameSetElement {
	rv := objc.Call[DOMHTMLFrameSetElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLFrameSetElement() DOMHTMLFrameSetElement {
	return DOMHTMLFrameSetElementClass.New()
}

func (d_ DOMHTMLFrameSetElement) Init() DOMHTMLFrameSetElement {
	rv := objc.Call[DOMHTMLFrameSetElement](d_, objc.Sel("init"))
	return rv
}
