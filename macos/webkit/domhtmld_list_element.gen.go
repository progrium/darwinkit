// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLDListElement] class.
var DOMHTMLDListElementClass = _DOMHTMLDListElementClass{objc.GetClass("DOMHTMLDListElement")}

type _DOMHTMLDListElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLDListElement] class.
type IDOMHTMLDListElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmldlistelement?language=objc
type DOMHTMLDListElement struct {
	DOMHTMLElement
}

func DOMHTMLDListElementFrom(ptr unsafe.Pointer) DOMHTMLDListElement {
	return DOMHTMLDListElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLDListElementClass) Alloc() DOMHTMLDListElement {
	rv := objc.Call[DOMHTMLDListElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLDListElement_Alloc() DOMHTMLDListElement {
	return DOMHTMLDListElementClass.Alloc()
}

func (dc _DOMHTMLDListElementClass) New() DOMHTMLDListElement {
	rv := objc.Call[DOMHTMLDListElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLDListElement() DOMHTMLDListElement {
	return DOMHTMLDListElementClass.New()
}

func (d_ DOMHTMLDListElement) Init() DOMHTMLDListElement {
	rv := objc.Call[DOMHTMLDListElement](d_, objc.Sel("init"))
	return rv
}
