// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLOptionElement] class.
var DOMHTMLOptionElementClass = _DOMHTMLOptionElementClass{objc.GetClass("DOMHTMLOptionElement")}

type _DOMHTMLOptionElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLOptionElement] class.
type IDOMHTMLOptionElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmloptionelement?language=objc
type DOMHTMLOptionElement struct {
	DOMHTMLElement
}

func DOMHTMLOptionElementFrom(ptr unsafe.Pointer) DOMHTMLOptionElement {
	return DOMHTMLOptionElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLOptionElementClass) Alloc() DOMHTMLOptionElement {
	rv := objc.Call[DOMHTMLOptionElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLOptionElement_Alloc() DOMHTMLOptionElement {
	return DOMHTMLOptionElementClass.Alloc()
}

func (dc _DOMHTMLOptionElementClass) New() DOMHTMLOptionElement {
	rv := objc.Call[DOMHTMLOptionElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLOptionElement() DOMHTMLOptionElement {
	return DOMHTMLOptionElementClass.New()
}

func (d_ DOMHTMLOptionElement) Init() DOMHTMLOptionElement {
	rv := objc.Call[DOMHTMLOptionElement](d_, objc.Sel("init"))
	return rv
}
