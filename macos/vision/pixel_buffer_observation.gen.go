// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PixelBufferObservation] class.
var PixelBufferObservationClass = _PixelBufferObservationClass{objc.GetClass("VNPixelBufferObservation")}

type _PixelBufferObservationClass struct {
	objc.Class
}

// An interface definition for the [PixelBufferObservation] class.
type IPixelBufferObservation interface {
	IObservation
	FeatureName() string
	PixelBuffer() corevideo.PixelBufferRef
}

// An object that represents an image that an image analysis request produces. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpixelbufferobservation?language=objc
type PixelBufferObservation struct {
	Observation
}

func PixelBufferObservationFrom(ptr unsafe.Pointer) PixelBufferObservation {
	return PixelBufferObservation{
		Observation: ObservationFrom(ptr),
	}
}

func (pc _PixelBufferObservationClass) Alloc() PixelBufferObservation {
	rv := objc.Call[PixelBufferObservation](pc, objc.Sel("alloc"))
	return rv
}

func PixelBufferObservation_Alloc() PixelBufferObservation {
	return PixelBufferObservationClass.Alloc()
}

func (pc _PixelBufferObservationClass) New() PixelBufferObservation {
	rv := objc.Call[PixelBufferObservation](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPixelBufferObservation() PixelBufferObservation {
	return PixelBufferObservationClass.New()
}

func (p_ PixelBufferObservation) Init() PixelBufferObservation {
	rv := objc.Call[PixelBufferObservation](p_, objc.Sel("init"))
	return rv
}

// A feature name that the CoreML model defines. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpixelbufferobservation/3131945-featurename?language=objc
func (p_ PixelBufferObservation) FeatureName() string {
	rv := objc.Call[string](p_, objc.Sel("featureName"))
	return rv
}

// The image that results from a request with image output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpixelbufferobservation/2890132-pixelbuffer?language=objc
func (p_ PixelBufferObservation) PixelBuffer() corevideo.PixelBufferRef {
	rv := objc.Call[corevideo.PixelBufferRef](p_, objc.Sel("pixelBuffer"))
	return rv
}
