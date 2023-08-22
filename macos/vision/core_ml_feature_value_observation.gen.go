// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coreml"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CoreMLFeatureValueObservation] class.
var CoreMLFeatureValueObservationClass = _CoreMLFeatureValueObservationClass{objc.GetClass("VNCoreMLFeatureValueObservation")}

type _CoreMLFeatureValueObservationClass struct {
	objc.Class
}

// An interface definition for the [CoreMLFeatureValueObservation] class.
type ICoreMLFeatureValueObservation interface {
	IObservation
	FeatureName() string
	FeatureValue() coreml.FeatureValue
}

// An object that represents a collection of key-value information that a Core ML image analysis request produces. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlfeaturevalueobservation?language=objc
type CoreMLFeatureValueObservation struct {
	Observation
}

func CoreMLFeatureValueObservationFrom(ptr unsafe.Pointer) CoreMLFeatureValueObservation {
	return CoreMLFeatureValueObservation{
		Observation: ObservationFrom(ptr),
	}
}

func (cc _CoreMLFeatureValueObservationClass) Alloc() CoreMLFeatureValueObservation {
	rv := objc.Call[CoreMLFeatureValueObservation](cc, objc.Sel("alloc"))
	return rv
}

func CoreMLFeatureValueObservation_Alloc() CoreMLFeatureValueObservation {
	return CoreMLFeatureValueObservationClass.Alloc()
}

func (cc _CoreMLFeatureValueObservationClass) New() CoreMLFeatureValueObservation {
	rv := objc.Call[CoreMLFeatureValueObservation](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCoreMLFeatureValueObservation() CoreMLFeatureValueObservation {
	return CoreMLFeatureValueObservationClass.New()
}

func (c_ CoreMLFeatureValueObservation) Init() CoreMLFeatureValueObservation {
	rv := objc.Call[CoreMLFeatureValueObservation](c_, objc.Sel("init"))
	return rv
}

// The name used in the model description of the CoreML model that produced this observation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlfeaturevalueobservation/3131944-featurename?language=objc
func (c_ CoreMLFeatureValueObservation) FeatureName() string {
	rv := objc.Call[string](c_, objc.Sel("featureName"))
	return rv
}

// The feature result of a VNCoreMLRequest that outputs neither a classification nor an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlfeaturevalueobservation/2890135-featurevalue?language=objc
func (c_ CoreMLFeatureValueObservation) FeatureValue() coreml.FeatureValue {
	rv := objc.Call[coreml.FeatureValue](c_, objc.Sel("featureValue"))
	return rv
}
