// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLParagraphElement] class.
var DOMHTMLParagraphElementClass = _DOMHTMLParagraphElementClass{objc.GetClass("DOMHTMLParagraphElement")}

type _DOMHTMLParagraphElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLParagraphElement] class.
type IDOMHTMLParagraphElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlparagraphelement?language=objc
type DOMHTMLParagraphElement struct {
	DOMHTMLElement
}

func DOMHTMLParagraphElementFrom(ptr unsafe.Pointer) DOMHTMLParagraphElement {
	return DOMHTMLParagraphElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLParagraphElementClass) Alloc() DOMHTMLParagraphElement {
	rv := objc.Call[DOMHTMLParagraphElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLParagraphElement_Alloc() DOMHTMLParagraphElement {
	return DOMHTMLParagraphElementClass.Alloc()
}

func (dc _DOMHTMLParagraphElementClass) New() DOMHTMLParagraphElement {
	rv := objc.Call[DOMHTMLParagraphElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLParagraphElement() DOMHTMLParagraphElement {
	return DOMHTMLParagraphElementClass.New()
}

func (d_ DOMHTMLParagraphElement) Init() DOMHTMLParagraphElement {
	rv := objc.Call[DOMHTMLParagraphElement](d_, objc.Sel("init"))
	return rv
}
