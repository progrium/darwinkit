// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLTableRowElement] class.
var DOMHTMLTableRowElementClass = _DOMHTMLTableRowElementClass{objc.GetClass("DOMHTMLTableRowElement")}

type _DOMHTMLTableRowElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLTableRowElement] class.
type IDOMHTMLTableRowElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmltablerowelement?language=objc
type DOMHTMLTableRowElement struct {
	DOMHTMLElement
}

func DOMHTMLTableRowElementFrom(ptr unsafe.Pointer) DOMHTMLTableRowElement {
	return DOMHTMLTableRowElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLTableRowElementClass) Alloc() DOMHTMLTableRowElement {
	rv := objc.Call[DOMHTMLTableRowElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLTableRowElement_Alloc() DOMHTMLTableRowElement {
	return DOMHTMLTableRowElementClass.Alloc()
}

func (dc _DOMHTMLTableRowElementClass) New() DOMHTMLTableRowElement {
	rv := objc.Call[DOMHTMLTableRowElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLTableRowElement() DOMHTMLTableRowElement {
	return DOMHTMLTableRowElementClass.New()
}

func (d_ DOMHTMLTableRowElement) Init() DOMHTMLTableRowElement {
	rv := objc.Call[DOMHTMLTableRowElement](d_, objc.Sel("init"))
	return rv
}
