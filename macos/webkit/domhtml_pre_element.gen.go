// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLPreElement] class.
var DOMHTMLPreElementClass = _DOMHTMLPreElementClass{objc.GetClass("DOMHTMLPreElement")}

type _DOMHTMLPreElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLPreElement] class.
type IDOMHTMLPreElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlpreelement?language=objc
type DOMHTMLPreElement struct {
	DOMHTMLElement
}

func DOMHTMLPreElementFrom(ptr unsafe.Pointer) DOMHTMLPreElement {
	return DOMHTMLPreElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLPreElementClass) Alloc() DOMHTMLPreElement {
	rv := objc.Call[DOMHTMLPreElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLPreElement_Alloc() DOMHTMLPreElement {
	return DOMHTMLPreElementClass.Alloc()
}

func (dc _DOMHTMLPreElementClass) New() DOMHTMLPreElement {
	rv := objc.Call[DOMHTMLPreElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLPreElement() DOMHTMLPreElement {
	return DOMHTMLPreElementClass.New()
}

func (d_ DOMHTMLPreElement) Init() DOMHTMLPreElement {
	rv := objc.Call[DOMHTMLPreElement](d_, objc.Sel("init"))
	return rv
}
