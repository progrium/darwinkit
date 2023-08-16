// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMCSSStyleSheet] class.
var DOMCSSStyleSheetClass = _DOMCSSStyleSheetClass{objc.GetClass("DOMCSSStyleSheet")}

type _DOMCSSStyleSheetClass struct {
	objc.Class
}

// An interface definition for the [DOMCSSStyleSheet] class.
type IDOMCSSStyleSheet interface {
	IDOMStyleSheet
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstylesheet?language=objc
type DOMCSSStyleSheet struct {
	DOMStyleSheet
}

func DOMCSSStyleSheetFrom(ptr unsafe.Pointer) DOMCSSStyleSheet {
	return DOMCSSStyleSheet{
		DOMStyleSheet: DOMStyleSheetFrom(ptr),
	}
}

func (dc _DOMCSSStyleSheetClass) Alloc() DOMCSSStyleSheet {
	rv := objc.Call[DOMCSSStyleSheet](dc, objc.Sel("alloc"))
	return rv
}

func DOMCSSStyleSheet_Alloc() DOMCSSStyleSheet {
	return DOMCSSStyleSheetClass.Alloc()
}

func (dc _DOMCSSStyleSheetClass) New() DOMCSSStyleSheet {
	rv := objc.Call[DOMCSSStyleSheet](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMCSSStyleSheet() DOMCSSStyleSheet {
	return DOMCSSStyleSheetClass.New()
}

func (d_ DOMCSSStyleSheet) Init() DOMCSSStyleSheet {
	rv := objc.Call[DOMCSSStyleSheet](d_, objc.Sel("init"))
	return rv
}
