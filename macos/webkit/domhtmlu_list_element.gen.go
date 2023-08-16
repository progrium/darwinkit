// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLUListElement] class.
var DOMHTMLUListElementClass = _DOMHTMLUListElementClass{objc.GetClass("DOMHTMLUListElement")}

type _DOMHTMLUListElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLUListElement] class.
type IDOMHTMLUListElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlulistelement?language=objc
type DOMHTMLUListElement struct {
	DOMHTMLElement
}

func DOMHTMLUListElementFrom(ptr unsafe.Pointer) DOMHTMLUListElement {
	return DOMHTMLUListElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLUListElementClass) Alloc() DOMHTMLUListElement {
	rv := objc.Call[DOMHTMLUListElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLUListElement_Alloc() DOMHTMLUListElement {
	return DOMHTMLUListElementClass.Alloc()
}

func (dc _DOMHTMLUListElementClass) New() DOMHTMLUListElement {
	rv := objc.Call[DOMHTMLUListElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLUListElement() DOMHTMLUListElement {
	return DOMHTMLUListElementClass.New()
}

func (d_ DOMHTMLUListElement) Init() DOMHTMLUListElement {
	rv := objc.Call[DOMHTMLUListElement](d_, objc.Sel("init"))
	return rv
}
