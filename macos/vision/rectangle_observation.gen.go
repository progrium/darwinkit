// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RectangleObservation] class.
var RectangleObservationClass = _RectangleObservationClass{objc.GetClass("VNRectangleObservation")}

type _RectangleObservationClass struct {
	objc.Class
}

// An interface definition for the [RectangleObservation] class.
type IRectangleObservation interface {
	IDetectedObjectObservation
	BottomRight() coregraphics.Point
	BottomLeft() coregraphics.Point
	TopRight() coregraphics.Point
	TopLeft() coregraphics.Point
}

// An object that represents the four vertices of a detected rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrectangleobservation?language=objc
type RectangleObservation struct {
	DetectedObjectObservation
}

func RectangleObservationFrom(ptr unsafe.Pointer) RectangleObservation {
	return RectangleObservation{
		DetectedObjectObservation: DetectedObjectObservationFrom(ptr),
	}
}

func (rc _RectangleObservationClass) Alloc() RectangleObservation {
	rv := objc.Call[RectangleObservation](rc, objc.Sel("alloc"))
	return rv
}

func RectangleObservation_Alloc() RectangleObservation {
	return RectangleObservationClass.Alloc()
}

func (rc _RectangleObservationClass) New() RectangleObservation {
	rv := objc.Call[RectangleObservation](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRectangleObservation() RectangleObservation {
	return RectangleObservationClass.New()
}

func (r_ RectangleObservation) Init() RectangleObservation {
	rv := objc.Call[RectangleObservation](r_, objc.Sel("init"))
	return rv
}

func (rc _RectangleObservationClass) ObservationWithBoundingBox(boundingBox coregraphics.Rect) RectangleObservation {
	rv := objc.Call[RectangleObservation](rc, objc.Sel("observationWithBoundingBox:"), boundingBox)
	return rv
}

// Creates an observation with a bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/2879297-observationwithboundingbox?language=objc
func RectangleObservation_ObservationWithBoundingBox(boundingBox coregraphics.Rect) RectangleObservation {
	return RectangleObservationClass.ObservationWithBoundingBox(boundingBox)
}

func (rc _RectangleObservationClass) ObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox coregraphics.Rect) RectangleObservation {
	rv := objc.Call[RectangleObservation](rc, objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return rv
}

// Creates an observation with a revision number and bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/2967105-observationwithrequestrevision?language=objc
func RectangleObservation_ObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox coregraphics.Rect) RectangleObservation {
	return RectangleObservationClass.ObservationWithRequestRevisionBoundingBox(requestRevision, boundingBox)
}

// The coordinates of the lower-right corner of the observation bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrectangleobservation/2867226-bottomright?language=objc
func (r_ RectangleObservation) BottomRight() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](r_, objc.Sel("bottomRight"))
	return rv
}

// The coordinates of the lower-left corner of the observation bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrectangleobservation/2867201-bottomleft?language=objc
func (r_ RectangleObservation) BottomLeft() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](r_, objc.Sel("bottomLeft"))
	return rv
}

// The coordinates of the upper-right corner of the observation bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrectangleobservation/2867233-topright?language=objc
func (r_ RectangleObservation) TopRight() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](r_, objc.Sel("topRight"))
	return rv
}

// The coordinates of the upper-left corner of the observation bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrectangleobservation/2867210-topleft?language=objc
func (r_ RectangleObservation) TopLeft() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](r_, objc.Sel("topLeft"))
	return rv
}
