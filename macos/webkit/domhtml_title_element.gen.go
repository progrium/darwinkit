// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLTitleElement] class.
var DOMHTMLTitleElementClass = _DOMHTMLTitleElementClass{objc.GetClass("DOMHTMLTitleElement")}

type _DOMHTMLTitleElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLTitleElement] class.
type IDOMHTMLTitleElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmltitleelement?language=objc
type DOMHTMLTitleElement struct {
	DOMHTMLElement
}

func DOMHTMLTitleElementFrom(ptr unsafe.Pointer) DOMHTMLTitleElement {
	return DOMHTMLTitleElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLTitleElementClass) Alloc() DOMHTMLTitleElement {
	rv := objc.Call[DOMHTMLTitleElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLTitleElement_Alloc() DOMHTMLTitleElement {
	return DOMHTMLTitleElementClass.Alloc()
}

func (dc _DOMHTMLTitleElementClass) New() DOMHTMLTitleElement {
	rv := objc.Call[DOMHTMLTitleElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLTitleElement() DOMHTMLTitleElement {
	return DOMHTMLTitleElementClass.New()
}

func (d_ DOMHTMLTitleElement) Init() DOMHTMLTitleElement {
	rv := objc.Call[DOMHTMLTitleElement](d_, objc.Sel("init"))
	return rv
}
