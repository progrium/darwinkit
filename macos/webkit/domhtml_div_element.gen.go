// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLDivElement] class.
var DOMHTMLDivElementClass = _DOMHTMLDivElementClass{objc.GetClass("DOMHTMLDivElement")}

type _DOMHTMLDivElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLDivElement] class.
type IDOMHTMLDivElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmldivelement?language=objc
type DOMHTMLDivElement struct {
	DOMHTMLElement
}

func DOMHTMLDivElementFrom(ptr unsafe.Pointer) DOMHTMLDivElement {
	return DOMHTMLDivElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLDivElementClass) Alloc() DOMHTMLDivElement {
	rv := objc.Call[DOMHTMLDivElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLDivElement_Alloc() DOMHTMLDivElement {
	return DOMHTMLDivElementClass.Alloc()
}

func (dc _DOMHTMLDivElementClass) New() DOMHTMLDivElement {
	rv := objc.Call[DOMHTMLDivElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLDivElement() DOMHTMLDivElement {
	return DOMHTMLDivElementClass.New()
}

func (d_ DOMHTMLDivElement) Init() DOMHTMLDivElement {
	rv := objc.Call[DOMHTMLDivElement](d_, objc.Sel("init"))
	return rv
}
