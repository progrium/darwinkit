// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMRGBColor] class.
var DOMRGBColorClass = _DOMRGBColorClass{objc.GetClass("DOMRGBColor")}

type _DOMRGBColorClass struct {
	objc.Class
}

// An interface definition for the [DOMRGBColor] class.
type IDOMRGBColor interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domrgbcolor?language=objc
type DOMRGBColor struct {
	DOMObject
}

func DOMRGBColorFrom(ptr unsafe.Pointer) DOMRGBColor {
	return DOMRGBColor{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMRGBColorClass) Alloc() DOMRGBColor {
	rv := objc.Call[DOMRGBColor](dc, objc.Sel("alloc"))
	return rv
}

func DOMRGBColor_Alloc() DOMRGBColor {
	return DOMRGBColorClass.Alloc()
}

func (dc _DOMRGBColorClass) New() DOMRGBColor {
	rv := objc.Call[DOMRGBColor](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMRGBColor() DOMRGBColor {
	return DOMRGBColorClass.New()
}

func (d_ DOMRGBColor) Init() DOMRGBColor {
	rv := objc.Call[DOMRGBColor](d_, objc.Sel("init"))
	return rv
}
