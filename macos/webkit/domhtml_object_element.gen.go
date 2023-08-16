// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLObjectElement] class.
var DOMHTMLObjectElementClass = _DOMHTMLObjectElementClass{objc.GetClass("DOMHTMLObjectElement")}

type _DOMHTMLObjectElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLObjectElement] class.
type IDOMHTMLObjectElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlobjectelement?language=objc
type DOMHTMLObjectElement struct {
	DOMHTMLElement
}

func DOMHTMLObjectElementFrom(ptr unsafe.Pointer) DOMHTMLObjectElement {
	return DOMHTMLObjectElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLObjectElementClass) Alloc() DOMHTMLObjectElement {
	rv := objc.Call[DOMHTMLObjectElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLObjectElement_Alloc() DOMHTMLObjectElement {
	return DOMHTMLObjectElementClass.Alloc()
}

func (dc _DOMHTMLObjectElementClass) New() DOMHTMLObjectElement {
	rv := objc.Call[DOMHTMLObjectElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLObjectElement() DOMHTMLObjectElement {
	return DOMHTMLObjectElementClass.New()
}

func (d_ DOMHTMLObjectElement) Init() DOMHTMLObjectElement {
	rv := objc.Call[DOMHTMLObjectElement](d_, objc.Sel("init"))
	return rv
}
