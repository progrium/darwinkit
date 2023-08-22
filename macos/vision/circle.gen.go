// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Circle] class.
var CircleClass = _CircleClass{objc.GetClass("VNCircle")}

type _CircleClass struct {
	objc.Class
}

// An interface definition for the [Circle] class.
type ICircle interface {
	objc.IObject
	ContainsPoint(point IPoint) bool
	Radius() float64
	Diameter() float64
	Center() Point
}

// An immutable, two-dimensional circle represented by its center point and radius. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncircle?language=objc
type Circle struct {
	objc.Object
}

func CircleFrom(ptr unsafe.Pointer) Circle {
	return Circle{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ Circle) InitWithCenterRadius(center IPoint, radius float64) Circle {
	rv := objc.Call[Circle](c_, objc.Sel("initWithCenter:radius:"), objc.Ptr(center), radius)
	return rv
}

// Creates a circle with the specified center and radius. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncircle/3548317-initwithcenter?language=objc
func NewCircleWithCenterRadius(center IPoint, radius float64) Circle {
	instance := CircleClass.Alloc().InitWithCenterRadius(center, radius)
	instance.Autorelease()
	return instance
}

func (cc _CircleClass) Alloc() Circle {
	rv := objc.Call[Circle](cc, objc.Sel("alloc"))
	return rv
}

func Circle_Alloc() Circle {
	return CircleClass.Alloc()
}

func (cc _CircleClass) New() Circle {
	rv := objc.Call[Circle](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCircle() Circle {
	return CircleClass.New()
}

func (c_ Circle) Init() Circle {
	rv := objc.Call[Circle](c_, objc.Sel("init"))
	return rv
}

// Determines if this circle, including its boundary, contains the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncircle/3548313-containspoint?language=objc
func (c_ Circle) ContainsPoint(point IPoint) bool {
	rv := objc.Call[bool](c_, objc.Sel("containsPoint:"), objc.Ptr(point))
	return rv
}

// The circle’s radius. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncircle/3548318-radius?language=objc
func (c_ Circle) Radius() float64 {
	rv := objc.Call[float64](c_, objc.Sel("radius"))
	return rv
}

// The circle’s diameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncircle/3548315-diameter?language=objc
func (c_ Circle) Diameter() float64 {
	rv := objc.Call[float64](c_, objc.Sel("diameter"))
	return rv
}

// A circle object centered at the origin, with a radius of zero. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncircle/3548319-zerocircle?language=objc
func (cc _CircleClass) ZeroCircle() Circle {
	rv := objc.Call[Circle](cc, objc.Sel("zeroCircle"))
	return rv
}

// A circle object centered at the origin, with a radius of zero. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncircle/3548319-zerocircle?language=objc
func Circle_ZeroCircle() Circle {
	return CircleClass.ZeroCircle()
}

// The circle’s center point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncircle/3548312-center?language=objc
func (c_ Circle) Center() Point {
	rv := objc.Call[Point](c_, objc.Sel("center"))
	return rv
}
