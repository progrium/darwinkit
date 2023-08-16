// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLElement] class.
var DOMHTMLElementClass = _DOMHTMLElementClass{objc.GetClass("DOMHTMLElement")}

type _DOMHTMLElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLElement] class.
type IDOMHTMLElement interface {
	IDOMElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlelement?language=objc
type DOMHTMLElement struct {
	DOMElement
}

func DOMHTMLElementFrom(ptr unsafe.Pointer) DOMHTMLElement {
	return DOMHTMLElement{
		DOMElement: DOMElementFrom(ptr),
	}
}

func (dc _DOMHTMLElementClass) Alloc() DOMHTMLElement {
	rv := objc.Call[DOMHTMLElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLElement_Alloc() DOMHTMLElement {
	return DOMHTMLElementClass.Alloc()
}

func (dc _DOMHTMLElementClass) New() DOMHTMLElement {
	rv := objc.Call[DOMHTMLElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLElement() DOMHTMLElement {
	return DOMHTMLElementClass.New()
}

func (d_ DOMHTMLElement) Init() DOMHTMLElement {
	rv := objc.Call[DOMHTMLElement](d_, objc.Sel("init"))
	return rv
}
