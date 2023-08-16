// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLOptGroupElement] class.
var DOMHTMLOptGroupElementClass = _DOMHTMLOptGroupElementClass{objc.GetClass("DOMHTMLOptGroupElement")}

type _DOMHTMLOptGroupElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLOptGroupElement] class.
type IDOMHTMLOptGroupElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmloptgroupelement?language=objc
type DOMHTMLOptGroupElement struct {
	DOMHTMLElement
}

func DOMHTMLOptGroupElementFrom(ptr unsafe.Pointer) DOMHTMLOptGroupElement {
	return DOMHTMLOptGroupElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLOptGroupElementClass) Alloc() DOMHTMLOptGroupElement {
	rv := objc.Call[DOMHTMLOptGroupElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLOptGroupElement_Alloc() DOMHTMLOptGroupElement {
	return DOMHTMLOptGroupElementClass.Alloc()
}

func (dc _DOMHTMLOptGroupElementClass) New() DOMHTMLOptGroupElement {
	rv := objc.Call[DOMHTMLOptGroupElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLOptGroupElement() DOMHTMLOptGroupElement {
	return DOMHTMLOptGroupElementClass.New()
}

func (d_ DOMHTMLOptGroupElement) Init() DOMHTMLOptGroupElement {
	rv := objc.Call[DOMHTMLOptGroupElement](d_, objc.Sel("init"))
	return rv
}
