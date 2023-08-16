// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLFontElement] class.
var DOMHTMLFontElementClass = _DOMHTMLFontElementClass{objc.GetClass("DOMHTMLFontElement")}

type _DOMHTMLFontElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLFontElement] class.
type IDOMHTMLFontElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlfontelement?language=objc
type DOMHTMLFontElement struct {
	DOMHTMLElement
}

func DOMHTMLFontElementFrom(ptr unsafe.Pointer) DOMHTMLFontElement {
	return DOMHTMLFontElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLFontElementClass) Alloc() DOMHTMLFontElement {
	rv := objc.Call[DOMHTMLFontElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLFontElement_Alloc() DOMHTMLFontElement {
	return DOMHTMLFontElementClass.Alloc()
}

func (dc _DOMHTMLFontElementClass) New() DOMHTMLFontElement {
	rv := objc.Call[DOMHTMLFontElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLFontElement() DOMHTMLFontElement {
	return DOMHTMLFontElementClass.New()
}

func (d_ DOMHTMLFontElement) Init() DOMHTMLFontElement {
	rv := objc.Call[DOMHTMLFontElement](d_, objc.Sel("init"))
	return rv
}
