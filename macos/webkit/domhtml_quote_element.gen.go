// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLQuoteElement] class.
var DOMHTMLQuoteElementClass = _DOMHTMLQuoteElementClass{objc.GetClass("DOMHTMLQuoteElement")}

type _DOMHTMLQuoteElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLQuoteElement] class.
type IDOMHTMLQuoteElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlquoteelement?language=objc
type DOMHTMLQuoteElement struct {
	DOMHTMLElement
}

func DOMHTMLQuoteElementFrom(ptr unsafe.Pointer) DOMHTMLQuoteElement {
	return DOMHTMLQuoteElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLQuoteElementClass) Alloc() DOMHTMLQuoteElement {
	rv := objc.Call[DOMHTMLQuoteElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLQuoteElement_Alloc() DOMHTMLQuoteElement {
	return DOMHTMLQuoteElementClass.Alloc()
}

func (dc _DOMHTMLQuoteElementClass) New() DOMHTMLQuoteElement {
	rv := objc.Call[DOMHTMLQuoteElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLQuoteElement() DOMHTMLQuoteElement {
	return DOMHTMLQuoteElementClass.New()
}

func (d_ DOMHTMLQuoteElement) Init() DOMHTMLQuoteElement {
	rv := objc.Call[DOMHTMLQuoteElement](d_, objc.Sel("init"))
	return rv
}
