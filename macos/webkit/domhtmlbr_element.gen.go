// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLBRElement] class.
var DOMHTMLBRElementClass = _DOMHTMLBRElementClass{objc.GetClass("DOMHTMLBRElement")}

type _DOMHTMLBRElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLBRElement] class.
type IDOMHTMLBRElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlbrelement?language=objc
type DOMHTMLBRElement struct {
	DOMHTMLElement
}

func DOMHTMLBRElementFrom(ptr unsafe.Pointer) DOMHTMLBRElement {
	return DOMHTMLBRElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLBRElementClass) Alloc() DOMHTMLBRElement {
	rv := objc.Call[DOMHTMLBRElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLBRElement_Alloc() DOMHTMLBRElement {
	return DOMHTMLBRElementClass.Alloc()
}

func (dc _DOMHTMLBRElementClass) New() DOMHTMLBRElement {
	rv := objc.Call[DOMHTMLBRElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLBRElement() DOMHTMLBRElement {
	return DOMHTMLBRElementClass.New()
}

func (d_ DOMHTMLBRElement) Init() DOMHTMLBRElement {
	rv := objc.Call[DOMHTMLBRElement](d_, objc.Sel("init"))
	return rv
}
