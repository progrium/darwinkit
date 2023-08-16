// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLImageElement] class.
var DOMHTMLImageElementClass = _DOMHTMLImageElementClass{objc.GetClass("DOMHTMLImageElement")}

type _DOMHTMLImageElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLImageElement] class.
type IDOMHTMLImageElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlimageelement?language=objc
type DOMHTMLImageElement struct {
	DOMHTMLElement
}

func DOMHTMLImageElementFrom(ptr unsafe.Pointer) DOMHTMLImageElement {
	return DOMHTMLImageElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLImageElementClass) Alloc() DOMHTMLImageElement {
	rv := objc.Call[DOMHTMLImageElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLImageElement_Alloc() DOMHTMLImageElement {
	return DOMHTMLImageElementClass.Alloc()
}

func (dc _DOMHTMLImageElementClass) New() DOMHTMLImageElement {
	rv := objc.Call[DOMHTMLImageElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLImageElement() DOMHTMLImageElement {
	return DOMHTMLImageElementClass.New()
}

func (d_ DOMHTMLImageElement) Init() DOMHTMLImageElement {
	rv := objc.Call[DOMHTMLImageElement](d_, objc.Sel("init"))
	return rv
}
