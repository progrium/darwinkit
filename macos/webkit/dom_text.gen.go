// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMText] class.
var DOMTextClass = _DOMTextClass{objc.GetClass("DOMText")}

type _DOMTextClass struct {
	objc.Class
}

// An interface definition for the [DOMText] class.
type IDOMText interface {
	IDOMCharacterData
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domtext?language=objc
type DOMText struct {
	DOMCharacterData
}

func DOMTextFrom(ptr unsafe.Pointer) DOMText {
	return DOMText{
		DOMCharacterData: DOMCharacterDataFrom(ptr),
	}
}

func (dc _DOMTextClass) Alloc() DOMText {
	rv := objc.Call[DOMText](dc, objc.Sel("alloc"))
	return rv
}

func DOMText_Alloc() DOMText {
	return DOMTextClass.Alloc()
}

func (dc _DOMTextClass) New() DOMText {
	rv := objc.Call[DOMText](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMText() DOMText {
	return DOMTextClass.New()
}

func (d_ DOMText) Init() DOMText {
	rv := objc.Call[DOMText](d_, objc.Sel("init"))
	return rv
}
