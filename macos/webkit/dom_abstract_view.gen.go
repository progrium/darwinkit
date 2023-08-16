// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMAbstractView] class.
var DOMAbstractViewClass = _DOMAbstractViewClass{objc.GetClass("DOMAbstractView")}

type _DOMAbstractViewClass struct {
	objc.Class
}

// An interface definition for the [DOMAbstractView] class.
type IDOMAbstractView interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domabstractview?language=objc
type DOMAbstractView struct {
	DOMObject
}

func DOMAbstractViewFrom(ptr unsafe.Pointer) DOMAbstractView {
	return DOMAbstractView{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMAbstractViewClass) Alloc() DOMAbstractView {
	rv := objc.Call[DOMAbstractView](dc, objc.Sel("alloc"))
	return rv
}

func DOMAbstractView_Alloc() DOMAbstractView {
	return DOMAbstractViewClass.Alloc()
}

func (dc _DOMAbstractViewClass) New() DOMAbstractView {
	rv := objc.Call[DOMAbstractView](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMAbstractView() DOMAbstractView {
	return DOMAbstractViewClass.New()
}

func (d_ DOMAbstractView) Init() DOMAbstractView {
	rv := objc.Call[DOMAbstractView](d_, objc.Sel("init"))
	return rv
}
