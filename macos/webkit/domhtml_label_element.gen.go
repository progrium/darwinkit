// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLLabelElement] class.
var DOMHTMLLabelElementClass = _DOMHTMLLabelElementClass{objc.GetClass("DOMHTMLLabelElement")}

type _DOMHTMLLabelElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLLabelElement] class.
type IDOMHTMLLabelElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmllabelelement?language=objc
type DOMHTMLLabelElement struct {
	DOMHTMLElement
}

func DOMHTMLLabelElementFrom(ptr unsafe.Pointer) DOMHTMLLabelElement {
	return DOMHTMLLabelElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLLabelElementClass) Alloc() DOMHTMLLabelElement {
	rv := objc.Call[DOMHTMLLabelElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLLabelElement_Alloc() DOMHTMLLabelElement {
	return DOMHTMLLabelElementClass.Alloc()
}

func (dc _DOMHTMLLabelElementClass) New() DOMHTMLLabelElement {
	rv := objc.Call[DOMHTMLLabelElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLLabelElement() DOMHTMLLabelElement {
	return DOMHTMLLabelElementClass.New()
}

func (d_ DOMHTMLLabelElement) Init() DOMHTMLLabelElement {
	rv := objc.Call[DOMHTMLLabelElement](d_, objc.Sel("init"))
	return rv
}
