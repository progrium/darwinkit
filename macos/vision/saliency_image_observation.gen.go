// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SaliencyImageObservation] class.
var SaliencyImageObservationClass = _SaliencyImageObservationClass{objc.GetClass("VNSaliencyImageObservation")}

type _SaliencyImageObservationClass struct {
	objc.Class
}

// An interface definition for the [SaliencyImageObservation] class.
type ISaliencyImageObservation interface {
	IPixelBufferObservation
	SalientObjects() []RectangleObservation
}

// An observation that contains a grayscale heat map of important areas across an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnsaliencyimageobservation?language=objc
type SaliencyImageObservation struct {
	PixelBufferObservation
}

func SaliencyImageObservationFrom(ptr unsafe.Pointer) SaliencyImageObservation {
	return SaliencyImageObservation{
		PixelBufferObservation: PixelBufferObservationFrom(ptr),
	}
}

func (sc _SaliencyImageObservationClass) Alloc() SaliencyImageObservation {
	rv := objc.Call[SaliencyImageObservation](sc, objc.Sel("alloc"))
	return rv
}

func SaliencyImageObservation_Alloc() SaliencyImageObservation {
	return SaliencyImageObservationClass.Alloc()
}

func (sc _SaliencyImageObservationClass) New() SaliencyImageObservation {
	rv := objc.Call[SaliencyImageObservation](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSaliencyImageObservation() SaliencyImageObservation {
	return SaliencyImageObservationClass.New()
}

func (s_ SaliencyImageObservation) Init() SaliencyImageObservation {
	rv := objc.Call[SaliencyImageObservation](s_, objc.Sel("init"))
	return rv
}

// A collection of objects describing the distinct areas of the saliency heat map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnsaliencyimageobservation/3114160-salientobjects?language=objc
func (s_ SaliencyImageObservation) SalientObjects() []RectangleObservation {
	rv := objc.Call[[]RectangleObservation](s_, objc.Sel("salientObjects"))
	return rv
}
