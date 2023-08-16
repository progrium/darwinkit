// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLAreaElement] class.
var DOMHTMLAreaElementClass = _DOMHTMLAreaElementClass{objc.GetClass("DOMHTMLAreaElement")}

type _DOMHTMLAreaElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLAreaElement] class.
type IDOMHTMLAreaElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlareaelement?language=objc
type DOMHTMLAreaElement struct {
	DOMHTMLElement
}

func DOMHTMLAreaElementFrom(ptr unsafe.Pointer) DOMHTMLAreaElement {
	return DOMHTMLAreaElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLAreaElementClass) Alloc() DOMHTMLAreaElement {
	rv := objc.Call[DOMHTMLAreaElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLAreaElement_Alloc() DOMHTMLAreaElement {
	return DOMHTMLAreaElementClass.Alloc()
}

func (dc _DOMHTMLAreaElementClass) New() DOMHTMLAreaElement {
	rv := objc.Call[DOMHTMLAreaElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLAreaElement() DOMHTMLAreaElement {
	return DOMHTMLAreaElementClass.New()
}

func (d_ DOMHTMLAreaElement) Init() DOMHTMLAreaElement {
	rv := objc.Call[DOMHTMLAreaElement](d_, objc.Sel("init"))
	return rv
}
