// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/kernel"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageHomographicAlignmentObservation] class.
var ImageHomographicAlignmentObservationClass = _ImageHomographicAlignmentObservationClass{objc.GetClass("VNImageHomographicAlignmentObservation")}

type _ImageHomographicAlignmentObservationClass struct {
	objc.Class
}

// An interface definition for the [ImageHomographicAlignmentObservation] class.
type IImageHomographicAlignmentObservation interface {
	IImageAlignmentObservation
	WarpTransform() kernel.Matrix_float3x3
}

// An object that represents a perspective warp transformation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimagehomographicalignmentobservation?language=objc
type ImageHomographicAlignmentObservation struct {
	ImageAlignmentObservation
}

func ImageHomographicAlignmentObservationFrom(ptr unsafe.Pointer) ImageHomographicAlignmentObservation {
	return ImageHomographicAlignmentObservation{
		ImageAlignmentObservation: ImageAlignmentObservationFrom(ptr),
	}
}

func (ic _ImageHomographicAlignmentObservationClass) Alloc() ImageHomographicAlignmentObservation {
	rv := objc.Call[ImageHomographicAlignmentObservation](ic, objc.Sel("alloc"))
	return rv
}

func ImageHomographicAlignmentObservation_Alloc() ImageHomographicAlignmentObservation {
	return ImageHomographicAlignmentObservationClass.Alloc()
}

func (ic _ImageHomographicAlignmentObservationClass) New() ImageHomographicAlignmentObservation {
	rv := objc.Call[ImageHomographicAlignmentObservation](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageHomographicAlignmentObservation() ImageHomographicAlignmentObservation {
	return ImageHomographicAlignmentObservationClass.New()
}

func (i_ ImageHomographicAlignmentObservation) Init() ImageHomographicAlignmentObservation {
	rv := objc.Call[ImageHomographicAlignmentObservation](i_, objc.Sel("init"))
	return rv
}

// The warp transform matrix to morph the floating image into the reference image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimagehomographicalignmentobservation/2887129-warptransform?language=objc
func (i_ ImageHomographicAlignmentObservation) WarpTransform() kernel.Matrix_float3x3 {
	rv := objc.Call[kernel.Matrix_float3x3](i_, objc.Sel("warpTransform"))
	return rv
}
