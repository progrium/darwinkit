// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLBodyElement] class.
var DOMHTMLBodyElementClass = _DOMHTMLBodyElementClass{objc.GetClass("DOMHTMLBodyElement")}

type _DOMHTMLBodyElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLBodyElement] class.
type IDOMHTMLBodyElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlbodyelement?language=objc
type DOMHTMLBodyElement struct {
	DOMHTMLElement
}

func DOMHTMLBodyElementFrom(ptr unsafe.Pointer) DOMHTMLBodyElement {
	return DOMHTMLBodyElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLBodyElementClass) Alloc() DOMHTMLBodyElement {
	rv := objc.Call[DOMHTMLBodyElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLBodyElement_Alloc() DOMHTMLBodyElement {
	return DOMHTMLBodyElementClass.Alloc()
}

func (dc _DOMHTMLBodyElementClass) New() DOMHTMLBodyElement {
	rv := objc.Call[DOMHTMLBodyElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLBodyElement() DOMHTMLBodyElement {
	return DOMHTMLBodyElementClass.New()
}

func (d_ DOMHTMLBodyElement) Init() DOMHTMLBodyElement {
	rv := objc.Call[DOMHTMLBodyElement](d_, objc.Sel("init"))
	return rv
}
