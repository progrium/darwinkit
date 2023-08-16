// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLMenuElement] class.
var DOMHTMLMenuElementClass = _DOMHTMLMenuElementClass{objc.GetClass("DOMHTMLMenuElement")}

type _DOMHTMLMenuElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLMenuElement] class.
type IDOMHTMLMenuElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlmenuelement?language=objc
type DOMHTMLMenuElement struct {
	DOMHTMLElement
}

func DOMHTMLMenuElementFrom(ptr unsafe.Pointer) DOMHTMLMenuElement {
	return DOMHTMLMenuElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLMenuElementClass) Alloc() DOMHTMLMenuElement {
	rv := objc.Call[DOMHTMLMenuElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLMenuElement_Alloc() DOMHTMLMenuElement {
	return DOMHTMLMenuElementClass.Alloc()
}

func (dc _DOMHTMLMenuElementClass) New() DOMHTMLMenuElement {
	rv := objc.Call[DOMHTMLMenuElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLMenuElement() DOMHTMLMenuElement {
	return DOMHTMLMenuElementClass.New()
}

func (d_ DOMHTMLMenuElement) Init() DOMHTMLMenuElement {
	rv := objc.Call[DOMHTMLMenuElement](d_, objc.Sel("init"))
	return rv
}
