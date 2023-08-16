// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLTableCaptionElement] class.
var DOMHTMLTableCaptionElementClass = _DOMHTMLTableCaptionElementClass{objc.GetClass("DOMHTMLTableCaptionElement")}

type _DOMHTMLTableCaptionElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLTableCaptionElement] class.
type IDOMHTMLTableCaptionElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmltablecaptionelement?language=objc
type DOMHTMLTableCaptionElement struct {
	DOMHTMLElement
}

func DOMHTMLTableCaptionElementFrom(ptr unsafe.Pointer) DOMHTMLTableCaptionElement {
	return DOMHTMLTableCaptionElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLTableCaptionElementClass) Alloc() DOMHTMLTableCaptionElement {
	rv := objc.Call[DOMHTMLTableCaptionElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLTableCaptionElement_Alloc() DOMHTMLTableCaptionElement {
	return DOMHTMLTableCaptionElementClass.Alloc()
}

func (dc _DOMHTMLTableCaptionElementClass) New() DOMHTMLTableCaptionElement {
	rv := objc.Call[DOMHTMLTableCaptionElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLTableCaptionElement() DOMHTMLTableCaptionElement {
	return DOMHTMLTableCaptionElementClass.New()
}

func (d_ DOMHTMLTableCaptionElement) Init() DOMHTMLTableCaptionElement {
	rv := objc.Call[DOMHTMLTableCaptionElement](d_, objc.Sel("init"))
	return rv
}
