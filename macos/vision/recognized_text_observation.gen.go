// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RecognizedTextObservation] class.
var RecognizedTextObservationClass = _RecognizedTextObservationClass{objc.GetClass("VNRecognizedTextObservation")}

type _RecognizedTextObservationClass struct {
	objc.Class
}

// An interface definition for the [RecognizedTextObservation] class.
type IRecognizedTextObservation interface {
	IRectangleObservation
	TopCandidates(maxCandidateCount uint) []RecognizedText
}

// A request that detects and recognizes regions of text in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedtextobservation?language=objc
type RecognizedTextObservation struct {
	RectangleObservation
}

func RecognizedTextObservationFrom(ptr unsafe.Pointer) RecognizedTextObservation {
	return RecognizedTextObservation{
		RectangleObservation: RectangleObservationFrom(ptr),
	}
}

func (rc _RecognizedTextObservationClass) Alloc() RecognizedTextObservation {
	rv := objc.Call[RecognizedTextObservation](rc, objc.Sel("alloc"))
	return rv
}

func RecognizedTextObservation_Alloc() RecognizedTextObservation {
	return RecognizedTextObservationClass.Alloc()
}

func (rc _RecognizedTextObservationClass) New() RecognizedTextObservation {
	rv := objc.Call[RecognizedTextObservation](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRecognizedTextObservation() RecognizedTextObservation {
	return RecognizedTextObservationClass.New()
}

func (r_ RecognizedTextObservation) Init() RecognizedTextObservation {
	rv := objc.Call[RecognizedTextObservation](r_, objc.Sel("init"))
	return rv
}

func (rc _RecognizedTextObservationClass) ObservationWithBoundingBox(boundingBox coregraphics.Rect) RecognizedTextObservation {
	rv := objc.Call[RecognizedTextObservation](rc, objc.Sel("observationWithBoundingBox:"), boundingBox)
	return rv
}

// Creates an observation with a bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/2879297-observationwithboundingbox?language=objc
func RecognizedTextObservation_ObservationWithBoundingBox(boundingBox coregraphics.Rect) RecognizedTextObservation {
	return RecognizedTextObservationClass.ObservationWithBoundingBox(boundingBox)
}

func (rc _RecognizedTextObservationClass) ObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox coregraphics.Rect) RecognizedTextObservation {
	rv := objc.Call[RecognizedTextObservation](rc, objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return rv
}

// Creates an observation with a revision number and bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/2967105-observationwithrequestrevision?language=objc
func RecognizedTextObservation_ObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox coregraphics.Rect) RecognizedTextObservation {
	return RecognizedTextObservationClass.ObservationWithRequestRevisionBoundingBox(requestRevision, boundingBox)
}

// Requests the n top candidates for a recognized text string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedtextobservation/3152637-topcandidates?language=objc
func (r_ RecognizedTextObservation) TopCandidates(maxCandidateCount uint) []RecognizedText {
	rv := objc.Call[[]RecognizedText](r_, objc.Sel("topCandidates:"), maxCandidateCount)
	return rv
}
