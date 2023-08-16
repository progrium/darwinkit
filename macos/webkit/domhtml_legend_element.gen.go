// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLLegendElement] class.
var DOMHTMLLegendElementClass = _DOMHTMLLegendElementClass{objc.GetClass("DOMHTMLLegendElement")}

type _DOMHTMLLegendElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLLegendElement] class.
type IDOMHTMLLegendElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmllegendelement?language=objc
type DOMHTMLLegendElement struct {
	DOMHTMLElement
}

func DOMHTMLLegendElementFrom(ptr unsafe.Pointer) DOMHTMLLegendElement {
	return DOMHTMLLegendElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLLegendElementClass) Alloc() DOMHTMLLegendElement {
	rv := objc.Call[DOMHTMLLegendElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLLegendElement_Alloc() DOMHTMLLegendElement {
	return DOMHTMLLegendElementClass.Alloc()
}

func (dc _DOMHTMLLegendElementClass) New() DOMHTMLLegendElement {
	rv := objc.Call[DOMHTMLLegendElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLLegendElement() DOMHTMLLegendElement {
	return DOMHTMLLegendElementClass.New()
}

func (d_ DOMHTMLLegendElement) Init() DOMHTMLLegendElement {
	rv := objc.Call[DOMHTMLLegendElement](d_, objc.Sel("init"))
	return rv
}
