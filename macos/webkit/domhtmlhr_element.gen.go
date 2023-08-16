// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLHRElement] class.
var DOMHTMLHRElementClass = _DOMHTMLHRElementClass{objc.GetClass("DOMHTMLHRElement")}

type _DOMHTMLHRElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLHRElement] class.
type IDOMHTMLHRElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlhrelement?language=objc
type DOMHTMLHRElement struct {
	DOMHTMLElement
}

func DOMHTMLHRElementFrom(ptr unsafe.Pointer) DOMHTMLHRElement {
	return DOMHTMLHRElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLHRElementClass) Alloc() DOMHTMLHRElement {
	rv := objc.Call[DOMHTMLHRElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLHRElement_Alloc() DOMHTMLHRElement {
	return DOMHTMLHRElementClass.Alloc()
}

func (dc _DOMHTMLHRElementClass) New() DOMHTMLHRElement {
	rv := objc.Call[DOMHTMLHRElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLHRElement() DOMHTMLHRElement {
	return DOMHTMLHRElementClass.New()
}

func (d_ DOMHTMLHRElement) Init() DOMHTMLHRElement {
	rv := objc.Call[DOMHTMLHRElement](d_, objc.Sel("init"))
	return rv
}
