// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLIFrameElement] class.
var DOMHTMLIFrameElementClass = _DOMHTMLIFrameElementClass{objc.GetClass("DOMHTMLIFrameElement")}

type _DOMHTMLIFrameElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLIFrameElement] class.
type IDOMHTMLIFrameElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmliframeelement?language=objc
type DOMHTMLIFrameElement struct {
	DOMHTMLElement
}

func DOMHTMLIFrameElementFrom(ptr unsafe.Pointer) DOMHTMLIFrameElement {
	return DOMHTMLIFrameElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLIFrameElementClass) Alloc() DOMHTMLIFrameElement {
	rv := objc.Call[DOMHTMLIFrameElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLIFrameElement_Alloc() DOMHTMLIFrameElement {
	return DOMHTMLIFrameElementClass.Alloc()
}

func (dc _DOMHTMLIFrameElementClass) New() DOMHTMLIFrameElement {
	rv := objc.Call[DOMHTMLIFrameElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLIFrameElement() DOMHTMLIFrameElement {
	return DOMHTMLIFrameElementClass.New()
}

func (d_ DOMHTMLIFrameElement) Init() DOMHTMLIFrameElement {
	rv := objc.Call[DOMHTMLIFrameElement](d_, objc.Sel("init"))
	return rv
}
