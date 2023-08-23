// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageAlignmentObservation] class.
var ImageAlignmentObservationClass = _ImageAlignmentObservationClass{objc.GetClass("VNImageAlignmentObservation")}

type _ImageAlignmentObservationClass struct {
	objc.Class
}

// An interface definition for the [ImageAlignmentObservation] class.
type IImageAlignmentObservation interface {
	IObservation
}

// The abstract superclass for image analysis results that describe the relative alignment of two images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimagealignmentobservation?language=objc
type ImageAlignmentObservation struct {
	Observation
}

func ImageAlignmentObservationFrom(ptr unsafe.Pointer) ImageAlignmentObservation {
	return ImageAlignmentObservation{
		Observation: ObservationFrom(ptr),
	}
}

func (ic _ImageAlignmentObservationClass) Alloc() ImageAlignmentObservation {
	rv := objc.Call[ImageAlignmentObservation](ic, objc.Sel("alloc"))
	return rv
}

func ImageAlignmentObservation_Alloc() ImageAlignmentObservation {
	return ImageAlignmentObservationClass.Alloc()
}

func (ic _ImageAlignmentObservationClass) New() ImageAlignmentObservation {
	rv := objc.Call[ImageAlignmentObservation](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageAlignmentObservation() ImageAlignmentObservation {
	return ImageAlignmentObservationClass.New()
}

func (i_ ImageAlignmentObservation) Init() ImageAlignmentObservation {
	rv := objc.Call[ImageAlignmentObservation](i_, objc.Sel("init"))
	return rv
}
