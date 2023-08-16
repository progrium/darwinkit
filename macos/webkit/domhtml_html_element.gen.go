// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLHtmlElement] class.
var DOMHTMLHtmlElementClass = _DOMHTMLHtmlElementClass{objc.GetClass("DOMHTMLHtmlElement")}

type _DOMHTMLHtmlElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLHtmlElement] class.
type IDOMHTMLHtmlElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlhtmlelement?language=objc
type DOMHTMLHtmlElement struct {
	DOMHTMLElement
}

func DOMHTMLHtmlElementFrom(ptr unsafe.Pointer) DOMHTMLHtmlElement {
	return DOMHTMLHtmlElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLHtmlElementClass) Alloc() DOMHTMLHtmlElement {
	rv := objc.Call[DOMHTMLHtmlElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLHtmlElement_Alloc() DOMHTMLHtmlElement {
	return DOMHTMLHtmlElementClass.Alloc()
}

func (dc _DOMHTMLHtmlElementClass) New() DOMHTMLHtmlElement {
	rv := objc.Call[DOMHTMLHtmlElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLHtmlElement() DOMHTMLHtmlElement {
	return DOMHTMLHtmlElementClass.New()
}

func (d_ DOMHTMLHtmlElement) Init() DOMHTMLHtmlElement {
	rv := objc.Call[DOMHTMLHtmlElement](d_, objc.Sel("init"))
	return rv
}
