// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLTableColElement] class.
var DOMHTMLTableColElementClass = _DOMHTMLTableColElementClass{objc.GetClass("DOMHTMLTableColElement")}

type _DOMHTMLTableColElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLTableColElement] class.
type IDOMHTMLTableColElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmltablecolelement?language=objc
type DOMHTMLTableColElement struct {
	DOMHTMLElement
}

func DOMHTMLTableColElementFrom(ptr unsafe.Pointer) DOMHTMLTableColElement {
	return DOMHTMLTableColElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLTableColElementClass) Alloc() DOMHTMLTableColElement {
	rv := objc.Call[DOMHTMLTableColElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLTableColElement_Alloc() DOMHTMLTableColElement {
	return DOMHTMLTableColElementClass.Alloc()
}

func (dc _DOMHTMLTableColElementClass) New() DOMHTMLTableColElement {
	rv := objc.Call[DOMHTMLTableColElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLTableColElement() DOMHTMLTableColElement {
	return DOMHTMLTableColElementClass.New()
}

func (d_ DOMHTMLTableColElement) Init() DOMHTMLTableColElement {
	rv := objc.Call[DOMHTMLTableColElement](d_, objc.Sel("init"))
	return rv
}
