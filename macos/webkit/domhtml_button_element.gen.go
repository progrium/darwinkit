// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLButtonElement] class.
var DOMHTMLButtonElementClass = _DOMHTMLButtonElementClass{objc.GetClass("DOMHTMLButtonElement")}

type _DOMHTMLButtonElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLButtonElement] class.
type IDOMHTMLButtonElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlbuttonelement?language=objc
type DOMHTMLButtonElement struct {
	DOMHTMLElement
}

func DOMHTMLButtonElementFrom(ptr unsafe.Pointer) DOMHTMLButtonElement {
	return DOMHTMLButtonElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLButtonElementClass) Alloc() DOMHTMLButtonElement {
	rv := objc.Call[DOMHTMLButtonElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLButtonElement_Alloc() DOMHTMLButtonElement {
	return DOMHTMLButtonElementClass.Alloc()
}

func (dc _DOMHTMLButtonElementClass) New() DOMHTMLButtonElement {
	rv := objc.Call[DOMHTMLButtonElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLButtonElement() DOMHTMLButtonElement {
	return DOMHTMLButtonElementClass.New()
}

func (d_ DOMHTMLButtonElement) Init() DOMHTMLButtonElement {
	rv := objc.Call[DOMHTMLButtonElement](d_, objc.Sel("init"))
	return rv
}
