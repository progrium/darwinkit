// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLTableElement] class.
var DOMHTMLTableElementClass = _DOMHTMLTableElementClass{objc.GetClass("DOMHTMLTableElement")}

type _DOMHTMLTableElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLTableElement] class.
type IDOMHTMLTableElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmltableelement?language=objc
type DOMHTMLTableElement struct {
	DOMHTMLElement
}

func DOMHTMLTableElementFrom(ptr unsafe.Pointer) DOMHTMLTableElement {
	return DOMHTMLTableElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLTableElementClass) Alloc() DOMHTMLTableElement {
	rv := objc.Call[DOMHTMLTableElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLTableElement_Alloc() DOMHTMLTableElement {
	return DOMHTMLTableElementClass.Alloc()
}

func (dc _DOMHTMLTableElementClass) New() DOMHTMLTableElement {
	rv := objc.Call[DOMHTMLTableElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLTableElement() DOMHTMLTableElement {
	return DOMHTMLTableElementClass.New()
}

func (d_ DOMHTMLTableElement) Init() DOMHTMLTableElement {
	rv := objc.Call[DOMHTMLTableElement](d_, objc.Sel("init"))
	return rv
}
