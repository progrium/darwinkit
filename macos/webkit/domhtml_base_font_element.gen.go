// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLBaseFontElement] class.
var DOMHTMLBaseFontElementClass = _DOMHTMLBaseFontElementClass{objc.GetClass("DOMHTMLBaseFontElement")}

type _DOMHTMLBaseFontElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLBaseFontElement] class.
type IDOMHTMLBaseFontElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlbasefontelement?language=objc
type DOMHTMLBaseFontElement struct {
	DOMHTMLElement
}

func DOMHTMLBaseFontElementFrom(ptr unsafe.Pointer) DOMHTMLBaseFontElement {
	return DOMHTMLBaseFontElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLBaseFontElementClass) Alloc() DOMHTMLBaseFontElement {
	rv := objc.Call[DOMHTMLBaseFontElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLBaseFontElement_Alloc() DOMHTMLBaseFontElement {
	return DOMHTMLBaseFontElementClass.Alloc()
}

func (dc _DOMHTMLBaseFontElementClass) New() DOMHTMLBaseFontElement {
	rv := objc.Call[DOMHTMLBaseFontElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLBaseFontElement() DOMHTMLBaseFontElement {
	return DOMHTMLBaseFontElementClass.New()
}

func (d_ DOMHTMLBaseFontElement) Init() DOMHTMLBaseFontElement {
	rv := objc.Call[DOMHTMLBaseFontElement](d_, objc.Sel("init"))
	return rv
}
