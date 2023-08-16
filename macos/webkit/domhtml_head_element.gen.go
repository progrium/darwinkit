// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLHeadElement] class.
var DOMHTMLHeadElementClass = _DOMHTMLHeadElementClass{objc.GetClass("DOMHTMLHeadElement")}

type _DOMHTMLHeadElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLHeadElement] class.
type IDOMHTMLHeadElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlheadelement?language=objc
type DOMHTMLHeadElement struct {
	DOMHTMLElement
}

func DOMHTMLHeadElementFrom(ptr unsafe.Pointer) DOMHTMLHeadElement {
	return DOMHTMLHeadElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLHeadElementClass) Alloc() DOMHTMLHeadElement {
	rv := objc.Call[DOMHTMLHeadElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLHeadElement_Alloc() DOMHTMLHeadElement {
	return DOMHTMLHeadElementClass.Alloc()
}

func (dc _DOMHTMLHeadElementClass) New() DOMHTMLHeadElement {
	rv := objc.Call[DOMHTMLHeadElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLHeadElement() DOMHTMLHeadElement {
	return DOMHTMLHeadElementClass.New()
}

func (d_ DOMHTMLHeadElement) Init() DOMHTMLHeadElement {
	rv := objc.Call[DOMHTMLHeadElement](d_, objc.Sel("init"))
	return rv
}
