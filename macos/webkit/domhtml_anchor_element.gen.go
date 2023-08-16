// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLAnchorElement] class.
var DOMHTMLAnchorElementClass = _DOMHTMLAnchorElementClass{objc.GetClass("DOMHTMLAnchorElement")}

type _DOMHTMLAnchorElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLAnchorElement] class.
type IDOMHTMLAnchorElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlanchorelement?language=objc
type DOMHTMLAnchorElement struct {
	DOMHTMLElement
}

func DOMHTMLAnchorElementFrom(ptr unsafe.Pointer) DOMHTMLAnchorElement {
	return DOMHTMLAnchorElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLAnchorElementClass) Alloc() DOMHTMLAnchorElement {
	rv := objc.Call[DOMHTMLAnchorElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLAnchorElement_Alloc() DOMHTMLAnchorElement {
	return DOMHTMLAnchorElementClass.Alloc()
}

func (dc _DOMHTMLAnchorElementClass) New() DOMHTMLAnchorElement {
	rv := objc.Call[DOMHTMLAnchorElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLAnchorElement() DOMHTMLAnchorElement {
	return DOMHTMLAnchorElementClass.New()
}

func (d_ DOMHTMLAnchorElement) Init() DOMHTMLAnchorElement {
	rv := objc.Call[DOMHTMLAnchorElement](d_, objc.Sel("init"))
	return rv
}
