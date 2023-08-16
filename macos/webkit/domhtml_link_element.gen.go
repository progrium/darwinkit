// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLLinkElement] class.
var DOMHTMLLinkElementClass = _DOMHTMLLinkElementClass{objc.GetClass("DOMHTMLLinkElement")}

type _DOMHTMLLinkElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLLinkElement] class.
type IDOMHTMLLinkElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmllinkelement?language=objc
type DOMHTMLLinkElement struct {
	DOMHTMLElement
}

func DOMHTMLLinkElementFrom(ptr unsafe.Pointer) DOMHTMLLinkElement {
	return DOMHTMLLinkElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLLinkElementClass) Alloc() DOMHTMLLinkElement {
	rv := objc.Call[DOMHTMLLinkElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLLinkElement_Alloc() DOMHTMLLinkElement {
	return DOMHTMLLinkElementClass.Alloc()
}

func (dc _DOMHTMLLinkElementClass) New() DOMHTMLLinkElement {
	rv := objc.Call[DOMHTMLLinkElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLLinkElement() DOMHTMLLinkElement {
	return DOMHTMLLinkElementClass.New()
}

func (d_ DOMHTMLLinkElement) Init() DOMHTMLLinkElement {
	rv := objc.Call[DOMHTMLLinkElement](d_, objc.Sel("init"))
	return rv
}
