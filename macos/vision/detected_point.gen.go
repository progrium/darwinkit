// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DetectedPoint] class.
var DetectedPointClass = _DetectedPointClass{objc.GetClass("VNDetectedPoint")}

type _DetectedPointClass struct {
	objc.Class
}

// An interface definition for the [DetectedPoint] class.
type IDetectedPoint interface {
	IPoint
	Confidence() Confidence
}

// An object that represents a normalized point in an image, along with a confidence value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedpoint?language=objc
type DetectedPoint struct {
	Point
}

func DetectedPointFrom(ptr unsafe.Pointer) DetectedPoint {
	return DetectedPoint{
		Point: PointFrom(ptr),
	}
}

func (dc _DetectedPointClass) Alloc() DetectedPoint {
	rv := objc.Call[DetectedPoint](dc, objc.Sel("alloc"))
	return rv
}

func DetectedPoint_Alloc() DetectedPoint {
	return DetectedPointClass.Alloc()
}

func (dc _DetectedPointClass) New() DetectedPoint {
	rv := objc.Call[DetectedPoint](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDetectedPoint() DetectedPoint {
	return DetectedPointClass.New()
}

func (d_ DetectedPoint) Init() DetectedPoint {
	rv := objc.Call[DetectedPoint](d_, objc.Sel("init"))
	return rv
}

func (d_ DetectedPoint) InitWithXY(x float64, y float64) DetectedPoint {
	rv := objc.Call[DetectedPoint](d_, objc.Sel("initWithX:y:"), x, y)
	return rv
}

// Creates a point object with the specified coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/3548331-initwithx?language=objc
func NewDetectedPointWithXY(x float64, y float64) DetectedPoint {
	instance := DetectedPointClass.Alloc().InitWithXY(x, y)
	instance.Autorelease()
	return instance
}

func (d_ DetectedPoint) InitWithLocation(location coregraphics.Point) DetectedPoint {
	rv := objc.Call[DetectedPoint](d_, objc.Sel("initWithLocation:"), location)
	return rv
}

// Creates a point object from the specified Core Graphics point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/3548330-initwithlocation?language=objc
func NewDetectedPointWithLocation(location coregraphics.Point) DetectedPoint {
	instance := DetectedPointClass.Alloc().InitWithLocation(location)
	instance.Autorelease()
	return instance
}

// A confidence score that indicates the detected point’s accuracy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedpoint/3548297-confidence?language=objc
func (d_ DetectedPoint) Confidence() Confidence {
	rv := objc.Call[Confidence](d_, objc.Sel("confidence"))
	return rv
}
