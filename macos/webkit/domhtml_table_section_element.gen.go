// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLTableSectionElement] class.
var DOMHTMLTableSectionElementClass = _DOMHTMLTableSectionElementClass{objc.GetClass("DOMHTMLTableSectionElement")}

type _DOMHTMLTableSectionElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLTableSectionElement] class.
type IDOMHTMLTableSectionElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmltablesectionelement?language=objc
type DOMHTMLTableSectionElement struct {
	DOMHTMLElement
}

func DOMHTMLTableSectionElementFrom(ptr unsafe.Pointer) DOMHTMLTableSectionElement {
	return DOMHTMLTableSectionElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLTableSectionElementClass) Alloc() DOMHTMLTableSectionElement {
	rv := objc.Call[DOMHTMLTableSectionElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLTableSectionElement_Alloc() DOMHTMLTableSectionElement {
	return DOMHTMLTableSectionElementClass.Alloc()
}

func (dc _DOMHTMLTableSectionElementClass) New() DOMHTMLTableSectionElement {
	rv := objc.Call[DOMHTMLTableSectionElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLTableSectionElement() DOMHTMLTableSectionElement {
	return DOMHTMLTableSectionElementClass.New()
}

func (d_ DOMHTMLTableSectionElement) Init() DOMHTMLTableSectionElement {
	rv := objc.Call[DOMHTMLTableSectionElement](d_, objc.Sel("init"))
	return rv
}
