// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLTextAreaElement] class.
var DOMHTMLTextAreaElementClass = _DOMHTMLTextAreaElementClass{objc.GetClass("DOMHTMLTextAreaElement")}

type _DOMHTMLTextAreaElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLTextAreaElement] class.
type IDOMHTMLTextAreaElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmltextareaelement?language=objc
type DOMHTMLTextAreaElement struct {
	DOMHTMLElement
}

func DOMHTMLTextAreaElementFrom(ptr unsafe.Pointer) DOMHTMLTextAreaElement {
	return DOMHTMLTextAreaElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLTextAreaElementClass) Alloc() DOMHTMLTextAreaElement {
	rv := objc.Call[DOMHTMLTextAreaElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLTextAreaElement_Alloc() DOMHTMLTextAreaElement {
	return DOMHTMLTextAreaElementClass.Alloc()
}

func (dc _DOMHTMLTextAreaElementClass) New() DOMHTMLTextAreaElement {
	rv := objc.Call[DOMHTMLTextAreaElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLTextAreaElement() DOMHTMLTextAreaElement {
	return DOMHTMLTextAreaElementClass.New()
}

func (d_ DOMHTMLTextAreaElement) Init() DOMHTMLTextAreaElement {
	rv := objc.Call[DOMHTMLTextAreaElement](d_, objc.Sel("init"))
	return rv
}
