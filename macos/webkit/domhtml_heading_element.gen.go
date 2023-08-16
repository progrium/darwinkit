// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLHeadingElement] class.
var DOMHTMLHeadingElementClass = _DOMHTMLHeadingElementClass{objc.GetClass("DOMHTMLHeadingElement")}

type _DOMHTMLHeadingElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLHeadingElement] class.
type IDOMHTMLHeadingElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlheadingelement?language=objc
type DOMHTMLHeadingElement struct {
	DOMHTMLElement
}

func DOMHTMLHeadingElementFrom(ptr unsafe.Pointer) DOMHTMLHeadingElement {
	return DOMHTMLHeadingElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLHeadingElementClass) Alloc() DOMHTMLHeadingElement {
	rv := objc.Call[DOMHTMLHeadingElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLHeadingElement_Alloc() DOMHTMLHeadingElement {
	return DOMHTMLHeadingElementClass.Alloc()
}

func (dc _DOMHTMLHeadingElementClass) New() DOMHTMLHeadingElement {
	rv := objc.Call[DOMHTMLHeadingElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLHeadingElement() DOMHTMLHeadingElement {
	return DOMHTMLHeadingElementClass.New()
}

func (d_ DOMHTMLHeadingElement) Init() DOMHTMLHeadingElement {
	rv := objc.Call[DOMHTMLHeadingElement](d_, objc.Sel("init"))
	return rv
}
