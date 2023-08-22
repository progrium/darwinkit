// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Point] class.
var PointClass = _PointClass{objc.GetClass("VNPoint")}

type _PointClass struct {
	objc.Class
}

// An interface definition for the [Point] class.
type IPoint interface {
	objc.IObject
	DistanceToPoint(point IPoint) float64
	X() float64
	Y() float64
	Location() coregraphics.Point
}

// An immutable object that represents a single, two-dimensional point in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint?language=objc
type Point struct {
	objc.Object
}

func PointFrom(ptr unsafe.Pointer) Point {
	return Point{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ Point) InitWithXY(x float64, y float64) Point {
	rv := objc.Call[Point](p_, objc.Sel("initWithX:y:"), x, y)
	return rv
}

// Creates a point object with the specified coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/3548331-initwithx?language=objc
func NewPointWithXY(x float64, y float64) Point {
	instance := PointClass.Alloc().InitWithXY(x, y)
	instance.Autorelease()
	return instance
}

func (p_ Point) InitWithLocation(location coregraphics.Point) Point {
	rv := objc.Call[Point](p_, objc.Sel("initWithLocation:"), location)
	return rv
}

// Creates a point object from the specified Core Graphics point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/3548330-initwithlocation?language=objc
func NewPointWithLocation(location coregraphics.Point) Point {
	instance := PointClass.Alloc().InitWithLocation(location)
	instance.Autorelease()
	return instance
}

func (pc _PointClass) Alloc() Point {
	rv := objc.Call[Point](pc, objc.Sel("alloc"))
	return rv
}

func Point_Alloc() Point {
	return PointClass.Alloc()
}

func (pc _PointClass) New() Point {
	rv := objc.Call[Point](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPoint() Point {
	return PointClass.New()
}

func (p_ Point) Init() Point {
	rv := objc.Call[Point](p_, objc.Sel("init"))
	return rv
}

// Creates a point object that’s shifted by the X and Y offsets of the specified vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/3548333-pointbyapplyingvector?language=objc
func (pc _PointClass) PointByApplyingVectorToPoint(vector IVector, point IPoint) Point {
	rv := objc.Call[Point](pc, objc.Sel("pointByApplyingVector:toPoint:"), objc.Ptr(vector), objc.Ptr(point))
	return rv
}

// Creates a point object that’s shifted by the X and Y offsets of the specified vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/3548333-pointbyapplyingvector?language=objc
func Point_PointByApplyingVectorToPoint(vector IVector, point IPoint) Point {
	return PointClass.PointByApplyingVectorToPoint(vector, point)
}

// Returns the distance to another point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/3675674-distancetopoint?language=objc
func (p_ Point) DistanceToPoint(point IPoint) float64 {
	rv := objc.Call[float64](p_, objc.Sel("distanceToPoint:"), objc.Ptr(point))
	return rv
}

// A point object that represents the origin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/3548336-zeropoint?language=objc
func (pc _PointClass) ZeroPoint() Point {
	rv := objc.Call[Point](pc, objc.Sel("zeroPoint"))
	return rv
}

// A point object that represents the origin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/3548336-zeropoint?language=objc
func Point_ZeroPoint() Point {
	return PointClass.ZeroPoint()
}

// The x-coordinate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/3548334-x?language=objc
func (p_ Point) X() float64 {
	rv := objc.Call[float64](p_, objc.Sel("x"))
	return rv
}

// The y-coordinate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/3548335-y?language=objc
func (p_ Point) Y() float64 {
	rv := objc.Call[float64](p_, objc.Sel("y"))
	return rv
}

// The Core Graphics point for this point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/3548332-location?language=objc
func (p_ Point) Location() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](p_, objc.Sel("location"))
	return rv
}
