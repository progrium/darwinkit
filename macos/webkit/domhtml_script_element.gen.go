// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLScriptElement] class.
var DOMHTMLScriptElementClass = _DOMHTMLScriptElementClass{objc.GetClass("DOMHTMLScriptElement")}

type _DOMHTMLScriptElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLScriptElement] class.
type IDOMHTMLScriptElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlscriptelement?language=objc
type DOMHTMLScriptElement struct {
	DOMHTMLElement
}

func DOMHTMLScriptElementFrom(ptr unsafe.Pointer) DOMHTMLScriptElement {
	return DOMHTMLScriptElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLScriptElementClass) Alloc() DOMHTMLScriptElement {
	rv := objc.Call[DOMHTMLScriptElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLScriptElement_Alloc() DOMHTMLScriptElement {
	return DOMHTMLScriptElementClass.Alloc()
}

func (dc _DOMHTMLScriptElementClass) New() DOMHTMLScriptElement {
	rv := objc.Call[DOMHTMLScriptElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLScriptElement() DOMHTMLScriptElement {
	return DOMHTMLScriptElementClass.New()
}

func (d_ DOMHTMLScriptElement) Init() DOMHTMLScriptElement {
	rv := objc.Call[DOMHTMLScriptElement](d_, objc.Sel("init"))
	return rv
}
