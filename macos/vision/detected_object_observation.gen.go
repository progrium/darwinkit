// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DetectedObjectObservation] class.
var DetectedObjectObservationClass = _DetectedObjectObservationClass{objc.GetClass("VNDetectedObjectObservation")}

type _DetectedObjectObservationClass struct {
	objc.Class
}

// An interface definition for the [DetectedObjectObservation] class.
type IDetectedObjectObservation interface {
	IObservation
	GlobalSegmentationMask() PixelBufferObservation
	BoundingBox() coregraphics.Rect
}

// An observation that provides the position and extent of an image feature that an image analysis request detects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation?language=objc
type DetectedObjectObservation struct {
	Observation
}

func DetectedObjectObservationFrom(ptr unsafe.Pointer) DetectedObjectObservation {
	return DetectedObjectObservation{
		Observation: ObservationFrom(ptr),
	}
}

func (dc _DetectedObjectObservationClass) ObservationWithBoundingBox(boundingBox coregraphics.Rect) DetectedObjectObservation {
	rv := objc.Call[DetectedObjectObservation](dc, objc.Sel("observationWithBoundingBox:"), boundingBox)
	return rv
}

// Creates an observation with a bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/2879297-observationwithboundingbox?language=objc
func DetectedObjectObservation_ObservationWithBoundingBox(boundingBox coregraphics.Rect) DetectedObjectObservation {
	return DetectedObjectObservationClass.ObservationWithBoundingBox(boundingBox)
}

func (dc _DetectedObjectObservationClass) ObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox coregraphics.Rect) DetectedObjectObservation {
	rv := objc.Call[DetectedObjectObservation](dc, objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return rv
}

// Creates an observation with a revision number and bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/2967105-observationwithrequestrevision?language=objc
func DetectedObjectObservation_ObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox coregraphics.Rect) DetectedObjectObservation {
	return DetectedObjectObservationClass.ObservationWithRequestRevisionBoundingBox(requestRevision, boundingBox)
}

func (dc _DetectedObjectObservationClass) Alloc() DetectedObjectObservation {
	rv := objc.Call[DetectedObjectObservation](dc, objc.Sel("alloc"))
	return rv
}

func DetectedObjectObservation_Alloc() DetectedObjectObservation {
	return DetectedObjectObservationClass.Alloc()
}

func (dc _DetectedObjectObservationClass) New() DetectedObjectObservation {
	rv := objc.Call[DetectedObjectObservation](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDetectedObjectObservation() DetectedObjectObservation {
	return DetectedObjectObservationClass.New()
}

func (d_ DetectedObjectObservation) Init() DetectedObjectObservation {
	rv := objc.Call[DetectedObjectObservation](d_, objc.Sel("init"))
	return rv
}

// A resulting pixel buffer from a request to generate a segmentation mask for an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/3798796-globalsegmentationmask?language=objc
func (d_ DetectedObjectObservation) GlobalSegmentationMask() PixelBufferObservation {
	rv := objc.Call[PixelBufferObservation](d_, objc.Sel("globalSegmentationMask"))
	return rv
}

// The bounding box of the object that the request detects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/2867227-boundingbox?language=objc
func (d_ DetectedObjectObservation) BoundingBox() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](d_, objc.Sel("boundingBox"))
	return rv
}
