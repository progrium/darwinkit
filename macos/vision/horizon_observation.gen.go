// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [HorizonObservation] class.
var HorizonObservationClass = _HorizonObservationClass{objc.GetClass("VNHorizonObservation")}

type _HorizonObservationClass struct {
	objc.Class
}

// An interface definition for the [HorizonObservation] class.
type IHorizonObservation interface {
	IObservation
	Angle() float64
	Transform() coregraphics.AffineTransform
}

// The horizon angle information that an image analysis request detects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhorizonobservation?language=objc
type HorizonObservation struct {
	Observation
}

func HorizonObservationFrom(ptr unsafe.Pointer) HorizonObservation {
	return HorizonObservation{
		Observation: ObservationFrom(ptr),
	}
}

func (hc _HorizonObservationClass) Alloc() HorizonObservation {
	rv := objc.Call[HorizonObservation](hc, objc.Sel("alloc"))
	return rv
}

func HorizonObservation_Alloc() HorizonObservation {
	return HorizonObservationClass.Alloc()
}

func (hc _HorizonObservationClass) New() HorizonObservation {
	rv := objc.Call[HorizonObservation](hc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewHorizonObservation() HorizonObservation {
	return HorizonObservationClass.New()
}

func (h_ HorizonObservation) Init() HorizonObservation {
	rv := objc.Call[HorizonObservation](h_, objc.Sel("init"))
	return rv
}

// The angle of the observed horizon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhorizonobservation/2867230-angle?language=objc
func (h_ HorizonObservation) Angle() float64 {
	rv := objc.Call[float64](h_, objc.Sel("angle"))
	return rv
}

// The transform to apply to the detected horizon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhorizonobservation/2867231-transform?language=objc
func (h_ HorizonObservation) Transform() coregraphics.AffineTransform {
	rv := objc.Call[coregraphics.AffineTransform](h_, objc.Sel("transform"))
	return rv
}
