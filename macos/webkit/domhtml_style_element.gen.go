// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLStyleElement] class.
var DOMHTMLStyleElementClass = _DOMHTMLStyleElementClass{objc.GetClass("DOMHTMLStyleElement")}

type _DOMHTMLStyleElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLStyleElement] class.
type IDOMHTMLStyleElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlstyleelement?language=objc
type DOMHTMLStyleElement struct {
	DOMHTMLElement
}

func DOMHTMLStyleElementFrom(ptr unsafe.Pointer) DOMHTMLStyleElement {
	return DOMHTMLStyleElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLStyleElementClass) Alloc() DOMHTMLStyleElement {
	rv := objc.Call[DOMHTMLStyleElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLStyleElement_Alloc() DOMHTMLStyleElement {
	return DOMHTMLStyleElementClass.Alloc()
}

func (dc _DOMHTMLStyleElementClass) New() DOMHTMLStyleElement {
	rv := objc.Call[DOMHTMLStyleElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLStyleElement() DOMHTMLStyleElement {
	return DOMHTMLStyleElementClass.New()
}

func (d_ DOMHTMLStyleElement) Init() DOMHTMLStyleElement {
	rv := objc.Call[DOMHTMLStyleElement](d_, objc.Sel("init"))
	return rv
}
