// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLInputElement] class.
var DOMHTMLInputElementClass = _DOMHTMLInputElementClass{objc.GetClass("DOMHTMLInputElement")}

type _DOMHTMLInputElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLInputElement] class.
type IDOMHTMLInputElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlinputelement?language=objc
type DOMHTMLInputElement struct {
	DOMHTMLElement
}

func DOMHTMLInputElementFrom(ptr unsafe.Pointer) DOMHTMLInputElement {
	return DOMHTMLInputElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLInputElementClass) Alloc() DOMHTMLInputElement {
	rv := objc.Call[DOMHTMLInputElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLInputElement_Alloc() DOMHTMLInputElement {
	return DOMHTMLInputElementClass.Alloc()
}

func (dc _DOMHTMLInputElementClass) New() DOMHTMLInputElement {
	rv := objc.Call[DOMHTMLInputElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLInputElement() DOMHTMLInputElement {
	return DOMHTMLInputElementClass.New()
}

func (d_ DOMHTMLInputElement) Init() DOMHTMLInputElement {
	rv := objc.Call[DOMHTMLInputElement](d_, objc.Sel("init"))
	return rv
}
