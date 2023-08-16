// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLMetaElement] class.
var DOMHTMLMetaElementClass = _DOMHTMLMetaElementClass{objc.GetClass("DOMHTMLMetaElement")}

type _DOMHTMLMetaElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLMetaElement] class.
type IDOMHTMLMetaElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlmetaelement?language=objc
type DOMHTMLMetaElement struct {
	DOMHTMLElement
}

func DOMHTMLMetaElementFrom(ptr unsafe.Pointer) DOMHTMLMetaElement {
	return DOMHTMLMetaElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLMetaElementClass) Alloc() DOMHTMLMetaElement {
	rv := objc.Call[DOMHTMLMetaElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLMetaElement_Alloc() DOMHTMLMetaElement {
	return DOMHTMLMetaElementClass.Alloc()
}

func (dc _DOMHTMLMetaElementClass) New() DOMHTMLMetaElement {
	rv := objc.Call[DOMHTMLMetaElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLMetaElement() DOMHTMLMetaElement {
	return DOMHTMLMetaElementClass.New()
}

func (d_ DOMHTMLMetaElement) Init() DOMHTMLMetaElement {
	rv := objc.Call[DOMHTMLMetaElement](d_, objc.Sel("init"))
	return rv
}
