// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLFieldSetElement] class.
var DOMHTMLFieldSetElementClass = _DOMHTMLFieldSetElementClass{objc.GetClass("DOMHTMLFieldSetElement")}

type _DOMHTMLFieldSetElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLFieldSetElement] class.
type IDOMHTMLFieldSetElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlfieldsetelement?language=objc
type DOMHTMLFieldSetElement struct {
	DOMHTMLElement
}

func DOMHTMLFieldSetElementFrom(ptr unsafe.Pointer) DOMHTMLFieldSetElement {
	return DOMHTMLFieldSetElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLFieldSetElementClass) Alloc() DOMHTMLFieldSetElement {
	rv := objc.Call[DOMHTMLFieldSetElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLFieldSetElement_Alloc() DOMHTMLFieldSetElement {
	return DOMHTMLFieldSetElementClass.Alloc()
}

func (dc _DOMHTMLFieldSetElementClass) New() DOMHTMLFieldSetElement {
	rv := objc.Call[DOMHTMLFieldSetElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLFieldSetElement() DOMHTMLFieldSetElement {
	return DOMHTMLFieldSetElementClass.New()
}

func (d_ DOMHTMLFieldSetElement) Init() DOMHTMLFieldSetElement {
	rv := objc.Call[DOMHTMLFieldSetElement](d_, objc.Sel("init"))
	return rv
}
