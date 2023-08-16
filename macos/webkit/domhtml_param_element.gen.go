// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLParamElement] class.
var DOMHTMLParamElementClass = _DOMHTMLParamElementClass{objc.GetClass("DOMHTMLParamElement")}

type _DOMHTMLParamElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLParamElement] class.
type IDOMHTMLParamElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlparamelement?language=objc
type DOMHTMLParamElement struct {
	DOMHTMLElement
}

func DOMHTMLParamElementFrom(ptr unsafe.Pointer) DOMHTMLParamElement {
	return DOMHTMLParamElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLParamElementClass) Alloc() DOMHTMLParamElement {
	rv := objc.Call[DOMHTMLParamElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLParamElement_Alloc() DOMHTMLParamElement {
	return DOMHTMLParamElementClass.Alloc()
}

func (dc _DOMHTMLParamElementClass) New() DOMHTMLParamElement {
	rv := objc.Call[DOMHTMLParamElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLParamElement() DOMHTMLParamElement {
	return DOMHTMLParamElementClass.New()
}

func (d_ DOMHTMLParamElement) Init() DOMHTMLParamElement {
	rv := objc.Call[DOMHTMLParamElement](d_, objc.Sel("init"))
	return rv
}
