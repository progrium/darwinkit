// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLAppletElement] class.
var DOMHTMLAppletElementClass = _DOMHTMLAppletElementClass{objc.GetClass("DOMHTMLAppletElement")}

type _DOMHTMLAppletElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLAppletElement] class.
type IDOMHTMLAppletElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlappletelement?language=objc
type DOMHTMLAppletElement struct {
	DOMHTMLElement
}

func DOMHTMLAppletElementFrom(ptr unsafe.Pointer) DOMHTMLAppletElement {
	return DOMHTMLAppletElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLAppletElementClass) Alloc() DOMHTMLAppletElement {
	rv := objc.Call[DOMHTMLAppletElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLAppletElement_Alloc() DOMHTMLAppletElement {
	return DOMHTMLAppletElementClass.Alloc()
}

func (dc _DOMHTMLAppletElementClass) New() DOMHTMLAppletElement {
	rv := objc.Call[DOMHTMLAppletElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLAppletElement() DOMHTMLAppletElement {
	return DOMHTMLAppletElementClass.New()
}

func (d_ DOMHTMLAppletElement) Init() DOMHTMLAppletElement {
	rv := objc.Call[DOMHTMLAppletElement](d_, objc.Sel("init"))
	return rv
}
