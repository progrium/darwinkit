// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLSelectElement] class.
var DOMHTMLSelectElementClass = _DOMHTMLSelectElementClass{objc.GetClass("DOMHTMLSelectElement")}

type _DOMHTMLSelectElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLSelectElement] class.
type IDOMHTMLSelectElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlselectelement?language=objc
type DOMHTMLSelectElement struct {
	DOMHTMLElement
}

func DOMHTMLSelectElementFrom(ptr unsafe.Pointer) DOMHTMLSelectElement {
	return DOMHTMLSelectElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLSelectElementClass) Alloc() DOMHTMLSelectElement {
	rv := objc.Call[DOMHTMLSelectElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLSelectElement_Alloc() DOMHTMLSelectElement {
	return DOMHTMLSelectElementClass.Alloc()
}

func (dc _DOMHTMLSelectElementClass) New() DOMHTMLSelectElement {
	rv := objc.Call[DOMHTMLSelectElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLSelectElement() DOMHTMLSelectElement {
	return DOMHTMLSelectElementClass.New()
}

func (d_ DOMHTMLSelectElement) Init() DOMHTMLSelectElement {
	rv := objc.Call[DOMHTMLSelectElement](d_, objc.Sel("init"))
	return rv
}
