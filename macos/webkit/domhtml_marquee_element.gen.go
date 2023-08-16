// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLMarqueeElement] class.
var DOMHTMLMarqueeElementClass = _DOMHTMLMarqueeElementClass{objc.GetClass("DOMHTMLMarqueeElement")}

type _DOMHTMLMarqueeElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLMarqueeElement] class.
type IDOMHTMLMarqueeElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlmarqueeelement?language=objc
type DOMHTMLMarqueeElement struct {
	DOMHTMLElement
}

func DOMHTMLMarqueeElementFrom(ptr unsafe.Pointer) DOMHTMLMarqueeElement {
	return DOMHTMLMarqueeElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLMarqueeElementClass) Alloc() DOMHTMLMarqueeElement {
	rv := objc.Call[DOMHTMLMarqueeElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLMarqueeElement_Alloc() DOMHTMLMarqueeElement {
	return DOMHTMLMarqueeElementClass.Alloc()
}

func (dc _DOMHTMLMarqueeElementClass) New() DOMHTMLMarqueeElement {
	rv := objc.Call[DOMHTMLMarqueeElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLMarqueeElement() DOMHTMLMarqueeElement {
	return DOMHTMLMarqueeElementClass.New()
}

func (d_ DOMHTMLMarqueeElement) Init() DOMHTMLMarqueeElement {
	rv := objc.Call[DOMHTMLMarqueeElement](d_, objc.Sel("init"))
	return rv
}
