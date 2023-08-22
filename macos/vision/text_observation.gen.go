// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextObservation] class.
var TextObservationClass = _TextObservationClass{objc.GetClass("VNTextObservation")}

type _TextObservationClass struct {
	objc.Class
}

// An interface definition for the [TextObservation] class.
type ITextObservation interface {
	IRectangleObservation
	CharacterBoxes() []RectangleObservation
}

// Information about regions of text that an image analysis request detects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntextobservation?language=objc
type TextObservation struct {
	RectangleObservation
}

func TextObservationFrom(ptr unsafe.Pointer) TextObservation {
	return TextObservation{
		RectangleObservation: RectangleObservationFrom(ptr),
	}
}

func (tc _TextObservationClass) Alloc() TextObservation {
	rv := objc.Call[TextObservation](tc, objc.Sel("alloc"))
	return rv
}

func TextObservation_Alloc() TextObservation {
	return TextObservationClass.Alloc()
}

func (tc _TextObservationClass) New() TextObservation {
	rv := objc.Call[TextObservation](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextObservation() TextObservation {
	return TextObservationClass.New()
}

func (t_ TextObservation) Init() TextObservation {
	rv := objc.Call[TextObservation](t_, objc.Sel("init"))
	return rv
}

func (tc _TextObservationClass) ObservationWithBoundingBox(boundingBox coregraphics.Rect) TextObservation {
	rv := objc.Call[TextObservation](tc, objc.Sel("observationWithBoundingBox:"), boundingBox)
	return rv
}

// Creates an observation with a bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/2879297-observationwithboundingbox?language=objc
func TextObservation_ObservationWithBoundingBox(boundingBox coregraphics.Rect) TextObservation {
	return TextObservationClass.ObservationWithBoundingBox(boundingBox)
}

func (tc _TextObservationClass) ObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox coregraphics.Rect) TextObservation {
	rv := objc.Call[TextObservation](tc, objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return rv
}

// Creates an observation with a revision number and bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/2967105-observationwithrequestrevision?language=objc
func TextObservation_ObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox coregraphics.Rect) TextObservation {
	return TextObservationClass.ObservationWithRequestRevisionBoundingBox(requestRevision, boundingBox)
}

// An array of detected individual character bounding boxes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntextobservation/2867213-characterboxes?language=objc
func (t_ TextObservation) CharacterBoxes() []RectangleObservation {
	rv := objc.Call[[]RectangleObservation](t_, objc.Sel("characterBoxes"))
	return rv
}
