// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageTranslationAlignmentObservation] class.
var ImageTranslationAlignmentObservationClass = _ImageTranslationAlignmentObservationClass{objc.GetClass("VNImageTranslationAlignmentObservation")}

type _ImageTranslationAlignmentObservationClass struct {
	objc.Class
}

// An interface definition for the [ImageTranslationAlignmentObservation] class.
type IImageTranslationAlignmentObservation interface {
	IImageAlignmentObservation
	AlignmentTransform() coregraphics.AffineTransform
}

// Affine transform information that an image alignment request produces. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimagetranslationalignmentobservation?language=objc
type ImageTranslationAlignmentObservation struct {
	ImageAlignmentObservation
}

func ImageTranslationAlignmentObservationFrom(ptr unsafe.Pointer) ImageTranslationAlignmentObservation {
	return ImageTranslationAlignmentObservation{
		ImageAlignmentObservation: ImageAlignmentObservationFrom(ptr),
	}
}

func (ic _ImageTranslationAlignmentObservationClass) Alloc() ImageTranslationAlignmentObservation {
	rv := objc.Call[ImageTranslationAlignmentObservation](ic, objc.Sel("alloc"))
	return rv
}

func ImageTranslationAlignmentObservation_Alloc() ImageTranslationAlignmentObservation {
	return ImageTranslationAlignmentObservationClass.Alloc()
}

func (ic _ImageTranslationAlignmentObservationClass) New() ImageTranslationAlignmentObservation {
	rv := objc.Call[ImageTranslationAlignmentObservation](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageTranslationAlignmentObservation() ImageTranslationAlignmentObservation {
	return ImageTranslationAlignmentObservationClass.New()
}

func (i_ ImageTranslationAlignmentObservation) Init() ImageTranslationAlignmentObservation {
	rv := objc.Call[ImageTranslationAlignmentObservation](i_, objc.Sel("init"))
	return rv
}

// The alignment transform to align the floating image with the reference image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimagetranslationalignmentobservation/2887130-alignmenttransform?language=objc
func (i_ ImageTranslationAlignmentObservation) AlignmentTransform() coregraphics.AffineTransform {
	rv := objc.Call[coregraphics.AffineTransform](i_, objc.Sel("alignmentTransform"))
	return rv
}
