// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLDirectoryElement] class.
var DOMHTMLDirectoryElementClass = _DOMHTMLDirectoryElementClass{objc.GetClass("DOMHTMLDirectoryElement")}

type _DOMHTMLDirectoryElementClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLDirectoryElement] class.
type IDOMHTMLDirectoryElement interface {
	IDOMHTMLElement
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmldirectoryelement?language=objc
type DOMHTMLDirectoryElement struct {
	DOMHTMLElement
}

func DOMHTMLDirectoryElementFrom(ptr unsafe.Pointer) DOMHTMLDirectoryElement {
	return DOMHTMLDirectoryElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

func (dc _DOMHTMLDirectoryElementClass) Alloc() DOMHTMLDirectoryElement {
	rv := objc.Call[DOMHTMLDirectoryElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLDirectoryElement_Alloc() DOMHTMLDirectoryElement {
	return DOMHTMLDirectoryElementClass.Alloc()
}

func (dc _DOMHTMLDirectoryElementClass) New() DOMHTMLDirectoryElement {
	rv := objc.Call[DOMHTMLDirectoryElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLDirectoryElement() DOMHTMLDirectoryElement {
	return DOMHTMLDirectoryElementClass.New()
}

func (d_ DOMHTMLDirectoryElement) Init() DOMHTMLDirectoryElement {
	rv := objc.Call[DOMHTMLDirectoryElement](d_, objc.Sel("init"))
	return rv
}
