// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TouchClass = _TouchClass{objc.GetClass("NSTouch")}

type _TouchClass struct {
	objc.Class
}

type ITouch interface {
	objc.IObject
	LocationInView(view IView) foundation.Point
	PreviousLocationInView(view IView) foundation.Point
	Type() TouchType
	Phase() TouchPhase
	NormalizedPosition() foundation.Point
	IsResting() bool
	Device() objc.Object
	DeviceSize() foundation.Size
}

type Touch struct {
	objc.Object
}

func MakeTouch(ptr unsafe.Pointer) Touch {
	return Touch{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TouchClass) Alloc() Touch {
	rv := objc.CallMethod[Touch](tc, objc.GetSelector("alloc"))
	return rv
}

func Touch_Alloc() Touch {
	return TouchClass.Alloc()
}

func (tc _TouchClass) New() Touch {
	rv := objc.CallMethod[Touch](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTouch() Touch {
	return TouchClass.New()
}

func Touch_New() Touch {
	return TouchClass.New()
}

func (t_ Touch) Init() Touch {
	rv := objc.CallMethod[Touch](t_, objc.GetSelector("init"))
	return rv
}

func Touch_Init() Touch {
	return TouchClass.Alloc().Init()
}

func (t_ Touch) LocationInView(view IView) foundation.Point {
	rv := objc.CallMethod[foundation.Point](t_, objc.GetSelector("locationInView:"), objc.ExtractPtr(view))
	return rv
}

func (t_ Touch) PreviousLocationInView(view IView) foundation.Point {
	rv := objc.CallMethod[foundation.Point](t_, objc.GetSelector("previousLocationInView:"), objc.ExtractPtr(view))
	return rv
}

func (t_ Touch) Type() TouchType {
	rv := objc.CallMethod[TouchType](t_, objc.GetSelector("type"))
	return rv
}

func (t_ Touch) Phase() TouchPhase {
	rv := objc.CallMethod[TouchPhase](t_, objc.GetSelector("phase"))
	return rv
}

func (t_ Touch) NormalizedPosition() foundation.Point {
	rv := objc.CallMethod[foundation.Point](t_, objc.GetSelector("normalizedPosition"))
	return rv
}

func (t_ Touch) IsResting() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isResting"))
	return rv
}

func (t_ Touch) Device() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.GetSelector("device"))
	return rv
}

func (t_ Touch) DeviceSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.GetSelector("deviceSize"))
	return rv
}
