// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLModElement] class.
var DOMHTMLModElementClass = _DOMHTMLModElementClass{objc.GetClass("DOMHTMLModElement")}

type _DOMHTMLModElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLModElement] class.
type IDOMHTMLModElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlmodelement?language=objc
type DOMHTMLModElement struct {
	DOMHTMLElement
}

func DOMHTMLModElementFrom(ptr unsafe.Pointer) DOMHTMLModElement {
	return DOMHTMLModElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLModElementClass) Alloc() DOMHTMLModElement {
	rv := objc.Call[DOMHTMLModElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLModElement_Alloc() DOMHTMLModElement {
	return DOMHTMLModElementClass.Alloc()
}

func (dc _DOMHTMLModElementClass) New() DOMHTMLModElement {
	rv := objc.Call[DOMHTMLModElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLModElement() DOMHTMLModElement {
	return DOMHTMLModElementClass.New()
}

func (d_ DOMHTMLModElement) Init() DOMHTMLModElement {
	rv := objc.Call[DOMHTMLModElement](d_, objc.Sel("init"))
	return rv
}
