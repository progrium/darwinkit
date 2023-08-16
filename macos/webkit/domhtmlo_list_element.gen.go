// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLOListElement] class.
var DOMHTMLOListElementClass = _DOMHTMLOListElementClass{objc.GetClass("DOMHTMLOListElement")}

type _DOMHTMLOListElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLOListElement] class.
type IDOMHTMLOListElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlolistelement?language=objc
type DOMHTMLOListElement struct {
	DOMHTMLElement
}

func DOMHTMLOListElementFrom(ptr unsafe.Pointer) DOMHTMLOListElement {
	return DOMHTMLOListElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLOListElementClass) Alloc() DOMHTMLOListElement {
	rv := objc.Call[DOMHTMLOListElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLOListElement_Alloc() DOMHTMLOListElement {
	return DOMHTMLOListElementClass.Alloc()
}

func (dc _DOMHTMLOListElementClass) New() DOMHTMLOListElement {
	rv := objc.Call[DOMHTMLOListElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLOListElement() DOMHTMLOListElement {
	return DOMHTMLOListElementClass.New()
}

func (d_ DOMHTMLOListElement) Init() DOMHTMLOListElement {
	rv := objc.Call[DOMHTMLOListElement](d_, objc.Sel("init"))
	return rv
}
