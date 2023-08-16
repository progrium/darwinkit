// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMCSSValue] class.
var DOMCSSValueClass = _DOMCSSValueClass{objc.GetClass("DOMCSSValue")}

type _DOMCSSValueClass struct {
	objc.Class
}

// An interface definition for the [DOMCSSValue] class.
type IDOMCSSValue interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssvalue?language=objc
type DOMCSSValue struct {
	DOMObject
}

func DOMCSSValueFrom(ptr unsafe.Pointer) DOMCSSValue {
	return DOMCSSValue{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMCSSValueClass) Alloc() DOMCSSValue {
	rv := objc.Call[DOMCSSValue](dc, objc.Sel("alloc"))
	return rv
}

func DOMCSSValue_Alloc() DOMCSSValue {
	return DOMCSSValueClass.Alloc()
}

func (dc _DOMCSSValueClass) New() DOMCSSValue {
	rv := objc.Call[DOMCSSValue](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMCSSValue() DOMCSSValue {
	return DOMCSSValueClass.New()
}

func (d_ DOMCSSValue) Init() DOMCSSValue {
	rv := objc.Call[DOMCSSValue](d_, objc.Sel("init"))
	return rv
}
