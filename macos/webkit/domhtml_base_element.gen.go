// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLBaseElement] class.
var DOMHTMLBaseElementClass = _DOMHTMLBaseElementClass{objc.GetClass("DOMHTMLBaseElement")}

type _DOMHTMLBaseElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLBaseElement] class.
type IDOMHTMLBaseElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlbaseelement?language=objc
type DOMHTMLBaseElement struct {
	DOMHTMLElement
}

func DOMHTMLBaseElementFrom(ptr unsafe.Pointer) DOMHTMLBaseElement {
	return DOMHTMLBaseElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLBaseElementClass) Alloc() DOMHTMLBaseElement {
	rv := objc.Call[DOMHTMLBaseElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLBaseElement_Alloc() DOMHTMLBaseElement {
	return DOMHTMLBaseElementClass.Alloc()
}

func (dc _DOMHTMLBaseElementClass) New() DOMHTMLBaseElement {
	rv := objc.Call[DOMHTMLBaseElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLBaseElement() DOMHTMLBaseElement {
	return DOMHTMLBaseElementClass.New()
}

func (d_ DOMHTMLBaseElement) Init() DOMHTMLBaseElement {
	rv := objc.Call[DOMHTMLBaseElement](d_, objc.Sel("init"))
	return rv
}
