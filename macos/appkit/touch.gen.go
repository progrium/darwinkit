// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Touch] class.
var TouchClass = _TouchClass{objc.GetClass("NSTouch")}

type _TouchClass struct {
	objc.Class
}

// An interface definition for the [Touch] class.
type ITouch interface {
	objc.IObject
	PreviousLocationInView(view IView) foundation.Point
	LocationInView(view IView) foundation.Point
	Device() objc.Object
	NormalizedPosition() foundation.Point
	Identity() objc.Object
	DeviceSize() foundation.Size
	Type() TouchType
	Phase() TouchPhase
	IsResting() bool
}

// A snapshot of a particular touch at an instant in time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouch?language=objc
type Touch struct {
	objc.Object
}

func TouchFrom(ptr unsafe.Pointer) Touch {
	return Touch{
		Object: objc.ObjectFrom(ptr),
	}
}

func (tc _TouchClass) Alloc() Touch {
	rv := objc.Call[Touch](tc, objc.Sel("alloc"))
	return rv
}

func Touch_Alloc() Touch {
	return TouchClass.Alloc()
}

func (tc _TouchClass) New() Touch {
	rv := objc.Call[Touch](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTouch() Touch {
	return TouchClass.New()
}

func (t_ Touch) Init() Touch {
	rv := objc.Call[Touch](t_, objc.Sel("init"))
	return rv
}

// Indicates the previous location of the touch in the view's coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouch/2588251-previouslocationinview?language=objc
func (t_ Touch) PreviousLocationInView(view IView) foundation.Point {
	rv := objc.Call[foundation.Point](t_, objc.Sel("previousLocationInView:"), objc.Ptr(view))
	return rv
}

// Indicates the location of the touch in the view's coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouch/2588250-locationinview?language=objc
func (t_ Touch) LocationInView(view IView) foundation.Point {
	rv := objc.Call[foundation.Point](t_, objc.Sel("locationInView:"), objc.Ptr(view))
	return rv
}

// The digitizer that generates the touch. Useful to distinguish touches emanating from multiple-device scenarios. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouch/1533562-device?language=objc
func (t_ Touch) Device() objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("device"))
	return rv
}

// The normalized position of the touch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouch/1534031-normalizedposition?language=objc
func (t_ Touch) NormalizedPosition() foundation.Point {
	rv := objc.Call[foundation.Point](t_, objc.Sel("normalizedPosition"))
	return rv
}

// The changes to a particular touch during its lifetime. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouch/1535399-identity?language=objc
func (t_ Touch) Identity() objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("identity"))
	return rv
}

// The range of the touch device in points, such as 72 ppi. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouch/1528476-devicesize?language=objc
func (t_ Touch) DeviceSize() foundation.Size {
	rv := objc.Call[foundation.Size](t_, objc.Sel("deviceSize"))
	return rv
}

// A type of touch from a Touch Bar interaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouch/2544708-type?language=objc
func (t_ Touch) Type() TouchType {
	rv := objc.Call[TouchType](t_, objc.Sel("type"))
	return rv
}

// The current phase of the touch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouch/1531520-phase?language=objc
func (t_ Touch) Phase() TouchPhase {
	rv := objc.Call[TouchPhase](t_, objc.Sel("phase"))
	return rv
}

// The indicator for a resting touch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouch/1525663-resting?language=objc
func (t_ Touch) IsResting() bool {
	rv := objc.Call[bool](t_, objc.Sel("isResting"))
	return rv
}
