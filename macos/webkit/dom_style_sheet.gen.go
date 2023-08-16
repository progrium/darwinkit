// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMStyleSheet] class.
var DOMStyleSheetClass = _DOMStyleSheetClass{objc.GetClass("DOMStyleSheet")}

type _DOMStyleSheetClass struct {
	objc.Class
}

// An interface definition for the [DOMStyleSheet] class.
type IDOMStyleSheet interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domstylesheet?language=objc
type DOMStyleSheet struct {
	DOMObject
}

func DOMStyleSheetFrom(ptr unsafe.Pointer) DOMStyleSheet {
	return DOMStyleSheet{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMStyleSheetClass) Alloc() DOMStyleSheet {
	rv := objc.Call[DOMStyleSheet](dc, objc.Sel("alloc"))
	return rv
}

func DOMStyleSheet_Alloc() DOMStyleSheet {
	return DOMStyleSheetClass.Alloc()
}

func (dc _DOMStyleSheetClass) New() DOMStyleSheet {
	rv := objc.Call[DOMStyleSheet](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMStyleSheet() DOMStyleSheet {
	return DOMStyleSheetClass.New()
}

func (d_ DOMStyleSheet) Init() DOMStyleSheet {
	rv := objc.Call[DOMStyleSheet](d_, objc.Sel("init"))
	return rv
}
