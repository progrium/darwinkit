// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLFormElement] class.
var DOMHTMLFormElementClass = _DOMHTMLFormElementClass{objc.GetClass("DOMHTMLFormElement")}

type _DOMHTMLFormElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLFormElement] class.
type IDOMHTMLFormElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlformelement?language=objc
type DOMHTMLFormElement struct {
	DOMHTMLElement
}

func DOMHTMLFormElementFrom(ptr unsafe.Pointer) DOMHTMLFormElement {
	return DOMHTMLFormElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLFormElementClass) Alloc() DOMHTMLFormElement {
	rv := objc.Call[DOMHTMLFormElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLFormElement_Alloc() DOMHTMLFormElement {
	return DOMHTMLFormElementClass.Alloc()
}

func (dc _DOMHTMLFormElementClass) New() DOMHTMLFormElement {
	rv := objc.Call[DOMHTMLFormElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLFormElement() DOMHTMLFormElement {
	return DOMHTMLFormElementClass.New()
}

func (d_ DOMHTMLFormElement) Init() DOMHTMLFormElement {
	rv := objc.Call[DOMHTMLFormElement](d_, objc.Sel("init"))
	return rv
}
