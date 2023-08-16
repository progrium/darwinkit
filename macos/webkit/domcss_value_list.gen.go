// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMCSSValueList] class.
var DOMCSSValueListClass = _DOMCSSValueListClass{objc.GetClass("DOMCSSValueList")}

type _DOMCSSValueListClass struct {
	objc.Class
}

// An interface definition for the [DOMCSSValueList] class.
type IDOMCSSValueList interface {
	IDOMCSSValue
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssvaluelist?language=objc
type DOMCSSValueList struct {
	DOMCSSValue
}

func DOMCSSValueListFrom(ptr unsafe.Pointer) DOMCSSValueList {
	return DOMCSSValueList{
		DOMCSSValue: DOMCSSValueFrom(ptr),
	}
}

func (dc _DOMCSSValueListClass) Alloc() DOMCSSValueList {
	rv := objc.Call[DOMCSSValueList](dc, objc.Sel("alloc"))
	return rv
}

func DOMCSSValueList_Alloc() DOMCSSValueList {
	return DOMCSSValueListClass.Alloc()
}

func (dc _DOMCSSValueListClass) New() DOMCSSValueList {
	rv := objc.Call[DOMCSSValueList](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMCSSValueList() DOMCSSValueList {
	return DOMCSSValueListClass.New()
}

func (d_ DOMCSSValueList) Init() DOMCSSValueList {
	rv := objc.Call[DOMCSSValueList](d_, objc.Sel("init"))
	return rv
}
