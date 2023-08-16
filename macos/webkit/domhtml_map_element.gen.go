// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLMapElement] class.
var DOMHTMLMapElementClass = _DOMHTMLMapElementClass{objc.GetClass("DOMHTMLMapElement")}

type _DOMHTMLMapElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLMapElement] class.
type IDOMHTMLMapElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlmapelement?language=objc
type DOMHTMLMapElement struct {
	DOMHTMLElement
}

func DOMHTMLMapElementFrom(ptr unsafe.Pointer) DOMHTMLMapElement {
	return DOMHTMLMapElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLMapElementClass) Alloc() DOMHTMLMapElement {
	rv := objc.Call[DOMHTMLMapElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLMapElement_Alloc() DOMHTMLMapElement {
	return DOMHTMLMapElementClass.Alloc()
}

func (dc _DOMHTMLMapElementClass) New() DOMHTMLMapElement {
	rv := objc.Call[DOMHTMLMapElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLMapElement() DOMHTMLMapElement {
	return DOMHTMLMapElementClass.New()
}

func (d_ DOMHTMLMapElement) Init() DOMHTMLMapElement {
	rv := objc.Call[DOMHTMLMapElement](d_, objc.Sel("init"))
	return rv
}
