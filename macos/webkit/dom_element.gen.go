// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMElement] class.
var DOMElementClass = _DOMElementClass{objc.GetClass("DOMElement")}

type _DOMElementClass struct {
	objc.Class
}

// An interface definition for the [DOMElement] class.
type IDOMElement interface {
	IDOMNode
	Image() appkit.Image
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domelement?language=objc
type DOMElement struct {
	DOMNode
}

func DOMElementFrom(ptr unsafe.Pointer) DOMElement {
	return DOMElement{
		DOMNode: DOMNodeFrom(ptr),
	}
}

func (dc _DOMElementClass) Alloc() DOMElement {
	rv := objc.Call[DOMElement](dc, objc.Sel("alloc"))
	return rv
}

func DOMElement_Alloc() DOMElement {
	return DOMElementClass.Alloc()
}

func (dc _DOMElementClass) New() DOMElement {
	rv := objc.Call[DOMElement](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMElement() DOMElement {
	return DOMElementClass.New()
}

func (d_ DOMElement) Init() DOMElement {
	rv := objc.Call[DOMElement](d_, objc.Sel("init"))
	return rv
}

// Returns an image associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domelement/1536421-image?language=objc
func (d_ DOMElement) Image() appkit.Image {
	rv := objc.Call[appkit.Image](d_, objc.Sel("image"))
	return rv
}
