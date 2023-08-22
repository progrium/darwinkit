// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RecognizedPoint] class.
var RecognizedPointClass = _RecognizedPointClass{objc.GetClass("VNRecognizedPoint")}

type _RecognizedPointClass struct {
	objc.Class
}

// An interface definition for the [RecognizedPoint] class.
type IRecognizedPoint interface {
	IDetectedPoint
	Identifier() RecognizedPointKey
}

// An object that represents a normalized point in an image, along with an identifier label and a confidence value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedpoint?language=objc
type RecognizedPoint struct {
	DetectedPoint
}

func RecognizedPointFrom(ptr unsafe.Pointer) RecognizedPoint {
	return RecognizedPoint{
		DetectedPoint: DetectedPointFrom(ptr),
	}
}

func (rc _RecognizedPointClass) Alloc() RecognizedPoint {
	rv := objc.Call[RecognizedPoint](rc, objc.Sel("alloc"))
	return rv
}

func RecognizedPoint_Alloc() RecognizedPoint {
	return RecognizedPointClass.Alloc()
}

func (rc _RecognizedPointClass) New() RecognizedPoint {
	rv := objc.Call[RecognizedPoint](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRecognizedPoint() RecognizedPoint {
	return RecognizedPointClass.New()
}

func (r_ RecognizedPoint) Init() RecognizedPoint {
	rv := objc.Call[RecognizedPoint](r_, objc.Sel("init"))
	return rv
}

func (r_ RecognizedPoint) InitWithXY(x float64, y float64) RecognizedPoint {
	rv := objc.Call[RecognizedPoint](r_, objc.Sel("initWithX:y:"), x, y)
	return rv
}

// Creates a point object with the specified coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/3548331-initwithx?language=objc
func NewRecognizedPointWithXY(x float64, y float64) RecognizedPoint {
	instance := RecognizedPointClass.Alloc().InitWithXY(x, y)
	instance.Autorelease()
	return instance
}

func (r_ RecognizedPoint) InitWithLocation(location coregraphics.Point) RecognizedPoint {
	rv := objc.Call[RecognizedPoint](r_, objc.Sel("initWithLocation:"), location)
	return rv
}

// Creates a point object from the specified Core Graphics point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/3548330-initwithlocation?language=objc
func NewRecognizedPointWithLocation(location coregraphics.Point) RecognizedPoint {
	instance := RecognizedPointClass.Alloc().InitWithLocation(location)
	instance.Autorelease()
	return instance
}

// The point’s identifier label. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedpoint/3548299-identifier?language=objc
func (r_ RecognizedPoint) Identifier() RecognizedPointKey {
	rv := objc.Call[RecognizedPointKey](r_, objc.Sel("identifier"))
	return rv
}
