// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RecognizedObjectObservation] class.
var RecognizedObjectObservationClass = _RecognizedObjectObservationClass{objc.GetClass("VNRecognizedObjectObservation")}

type _RecognizedObjectObservationClass struct {
	objc.Class
}

// An interface definition for the [RecognizedObjectObservation] class.
type IRecognizedObjectObservation interface {
	IDetectedObjectObservation
	Labels() []ClassificationObservation
}

// A detected object observation with an array of classification labels that classify the recognized object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedobjectobservation?language=objc
type RecognizedObjectObservation struct {
	DetectedObjectObservation
}

func RecognizedObjectObservationFrom(ptr unsafe.Pointer) RecognizedObjectObservation {
	return RecognizedObjectObservation{
		DetectedObjectObservation: DetectedObjectObservationFrom(ptr),
	}
}

func (rc _RecognizedObjectObservationClass) Alloc() RecognizedObjectObservation {
	rv := objc.Call[RecognizedObjectObservation](rc, objc.Sel("alloc"))
	return rv
}

func RecognizedObjectObservation_Alloc() RecognizedObjectObservation {
	return RecognizedObjectObservationClass.Alloc()
}

func (rc _RecognizedObjectObservationClass) New() RecognizedObjectObservation {
	rv := objc.Call[RecognizedObjectObservation](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRecognizedObjectObservation() RecognizedObjectObservation {
	return RecognizedObjectObservationClass.New()
}

func (r_ RecognizedObjectObservation) Init() RecognizedObjectObservation {
	rv := objc.Call[RecognizedObjectObservation](r_, objc.Sel("init"))
	return rv
}

func (rc _RecognizedObjectObservationClass) ObservationWithBoundingBox(boundingBox coregraphics.Rect) RecognizedObjectObservation {
	rv := objc.Call[RecognizedObjectObservation](rc, objc.Sel("observationWithBoundingBox:"), boundingBox)
	return rv
}

// Creates an observation with a bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/2879297-observationwithboundingbox?language=objc
func RecognizedObjectObservation_ObservationWithBoundingBox(boundingBox coregraphics.Rect) RecognizedObjectObservation {
	return RecognizedObjectObservationClass.ObservationWithBoundingBox(boundingBox)
}

func (rc _RecognizedObjectObservationClass) ObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox coregraphics.Rect) RecognizedObjectObservation {
	rv := objc.Call[RecognizedObjectObservation](rc, objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return rv
}

// Creates an observation with a revision number and bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/2967105-observationwithrequestrevision?language=objc
func RecognizedObjectObservation_ObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox coregraphics.Rect) RecognizedObjectObservation {
	return RecognizedObjectObservationClass.ObservationWithRequestRevisionBoundingBox(requestRevision, boundingBox)
}

// An array of observations that classify the recognized object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedobjectobservation/2980942-labels?language=objc
func (r_ RecognizedObjectObservation) Labels() []ClassificationObservation {
	rv := objc.Call[[]ClassificationObservation](r_, objc.Sel("labels"))
	return rv
}
