// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLLIElement] class.
var DOMHTMLLIElementClass = _DOMHTMLLIElementClass{objc.GetClass("DOMHTMLLIElement")}

type _DOMHTMLLIElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLLIElement] class.
type IDOMHTMLLIElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmllielement?language=objc
type DOMHTMLLIElement struct {
	DOMHTMLElement
}

func DOMHTMLLIElementFrom(ptr unsafe.Pointer) DOMHTMLLIElement {
	return DOMHTMLLIElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLLIElementClass) Alloc() DOMHTMLLIElement {
	rv := objc.Call[DOMHTMLLIElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLLIElement_Alloc() DOMHTMLLIElement {
	return DOMHTMLLIElementClass.Alloc()
}

func (dc _DOMHTMLLIElementClass) New() DOMHTMLLIElement {
	rv := objc.Call[DOMHTMLLIElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLLIElement() DOMHTMLLIElement {
	return DOMHTMLLIElementClass.New()
}

func (d_ DOMHTMLLIElement) Init() DOMHTMLLIElement {
	rv := objc.Call[DOMHTMLLIElement](d_, objc.Sel("init"))
	return rv
}
