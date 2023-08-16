// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLTableCellElement] class.
var DOMHTMLTableCellElementClass = _DOMHTMLTableCellElementClass{objc.GetClass("DOMHTMLTableCellElement")}

type _DOMHTMLTableCellElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLTableCellElement] class.
type IDOMHTMLTableCellElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmltablecellelement?language=objc
type DOMHTMLTableCellElement struct {
	DOMHTMLElement
}

func DOMHTMLTableCellElementFrom(ptr unsafe.Pointer) DOMHTMLTableCellElement {
	return DOMHTMLTableCellElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLTableCellElementClass) Alloc() DOMHTMLTableCellElement {
	rv := objc.Call[DOMHTMLTableCellElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLTableCellElement_Alloc() DOMHTMLTableCellElement {
	return DOMHTMLTableCellElementClass.Alloc()
}

func (dc _DOMHTMLTableCellElementClass) New() DOMHTMLTableCellElement {
	rv := objc.Call[DOMHTMLTableCellElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLTableCellElement() DOMHTMLTableCellElement {
	return DOMHTMLTableCellElementClass.New()
}

func (d_ DOMHTMLTableCellElement) Init() DOMHTMLTableCellElement {
	rv := objc.Call[DOMHTMLTableCellElement](d_, objc.Sel("init"))
	return rv
}
