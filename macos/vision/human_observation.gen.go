// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [HumanObservation] class.
var HumanObservationClass = _HumanObservationClass{objc.GetClass("VNHumanObservation")}

type _HumanObservationClass struct {
	objc.Class
}

// An interface definition for the [HumanObservation] class.
type IHumanObservation interface {
	IDetectedObjectObservation
	UpperBodyOnly() bool
}

// An object that represents a person that the request detects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanobservation?language=objc
type HumanObservation struct {
	DetectedObjectObservation
}

func HumanObservationFrom(ptr unsafe.Pointer) HumanObservation {
	return HumanObservation{
		DetectedObjectObservation: DetectedObjectObservationFrom(ptr),
	}
}

func (hc _HumanObservationClass) Alloc() HumanObservation {
	rv := objc.Call[HumanObservation](hc, objc.Sel("alloc"))
	return rv
}

func HumanObservation_Alloc() HumanObservation {
	return HumanObservationClass.Alloc()
}

func (hc _HumanObservationClass) New() HumanObservation {
	rv := objc.Call[HumanObservation](hc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewHumanObservation() HumanObservation {
	return HumanObservationClass.New()
}

func (h_ HumanObservation) Init() HumanObservation {
	rv := objc.Call[HumanObservation](h_, objc.Sel("init"))
	return rv
}

func (hc _HumanObservationClass) ObservationWithBoundingBox(boundingBox coregraphics.Rect) HumanObservation {
	rv := objc.Call[HumanObservation](hc, objc.Sel("observationWithBoundingBox:"), boundingBox)
	return rv
}

// Creates an observation with a bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/2879297-observationwithboundingbox?language=objc
func HumanObservation_ObservationWithBoundingBox(boundingBox coregraphics.Rect) HumanObservation {
	return HumanObservationClass.ObservationWithBoundingBox(boundingBox)
}

func (hc _HumanObservationClass) ObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox coregraphics.Rect) HumanObservation {
	rv := objc.Call[HumanObservation](hc, objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return rv
}

// Creates an observation with a revision number and bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/2967105-observationwithrequestrevision?language=objc
func HumanObservation_ObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox coregraphics.Rect) HumanObservation {
	return HumanObservationClass.ObservationWithRequestRevisionBoundingBox(requestRevision, boundingBox)
}

// A Boolean value that indicates whether the observation represents an upper-body or full-body rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanobservation/3751000-upperbodyonly?language=objc
func (h_ HumanObservation) UpperBodyOnly() bool {
	rv := objc.Call[bool](h_, objc.Sel("upperBodyOnly"))
	return rv
}
