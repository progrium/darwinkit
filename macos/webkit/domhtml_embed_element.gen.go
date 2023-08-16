// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLEmbedElement] class.
var DOMHTMLEmbedElementClass = _DOMHTMLEmbedElementClass{objc.GetClass("DOMHTMLEmbedElement")}

type _DOMHTMLEmbedElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLEmbedElement] class.
type IDOMHTMLEmbedElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlembedelement?language=objc
type DOMHTMLEmbedElement struct {
	DOMHTMLElement
}

func DOMHTMLEmbedElementFrom(ptr unsafe.Pointer) DOMHTMLEmbedElement {
	return DOMHTMLEmbedElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLEmbedElementClass) Alloc() DOMHTMLEmbedElement {
	rv := objc.Call[DOMHTMLEmbedElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLEmbedElement_Alloc() DOMHTMLEmbedElement {
	return DOMHTMLEmbedElementClass.Alloc()
}

func (dc _DOMHTMLEmbedElementClass) New() DOMHTMLEmbedElement {
	rv := objc.Call[DOMHTMLEmbedElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLEmbedElement() DOMHTMLEmbedElement {
	return DOMHTMLEmbedElementClass.New()
}

func (d_ DOMHTMLEmbedElement) Init() DOMHTMLEmbedElement {
	rv := objc.Call[DOMHTMLEmbedElement](d_, objc.Sel("init"))
	return rv
}
