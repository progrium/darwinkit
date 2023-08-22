// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FeaturePrintObservation] class.
var FeaturePrintObservationClass = _FeaturePrintObservationClass{objc.GetClass("VNFeaturePrintObservation")}

type _FeaturePrintObservationClass struct {
	objc.Class
}

// An interface definition for the [FeaturePrintObservation] class.
type IFeaturePrintObservation interface {
	IObservation
	ComputeDistanceToFeaturePrintObservationError(outDistance *float64, featurePrint IFeaturePrintObservation, error foundation.IError) bool
	Data() []byte
	ElementType() ElementType
	ElementCount() uint
}

// An observation that provides the recognized feature print. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfeatureprintobservation?language=objc
type FeaturePrintObservation struct {
	Observation
}

func FeaturePrintObservationFrom(ptr unsafe.Pointer) FeaturePrintObservation {
	return FeaturePrintObservation{
		Observation: ObservationFrom(ptr),
	}
}

func (fc _FeaturePrintObservationClass) Alloc() FeaturePrintObservation {
	rv := objc.Call[FeaturePrintObservation](fc, objc.Sel("alloc"))
	return rv
}

func FeaturePrintObservation_Alloc() FeaturePrintObservation {
	return FeaturePrintObservationClass.Alloc()
}

func (fc _FeaturePrintObservationClass) New() FeaturePrintObservation {
	rv := objc.Call[FeaturePrintObservation](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFeaturePrintObservation() FeaturePrintObservation {
	return FeaturePrintObservationClass.New()
}

func (f_ FeaturePrintObservation) Init() FeaturePrintObservation {
	rv := objc.Call[FeaturePrintObservation](f_, objc.Sel("init"))
	return rv
}

// Computes the distance between two feature print observations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfeatureprintobservation/3182823-computedistance?language=objc
func (f_ FeaturePrintObservation) ComputeDistanceToFeaturePrintObservationError(outDistance *float64, featurePrint IFeaturePrintObservation, error foundation.IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("computeDistance:toFeaturePrintObservation:error:"), outDistance, objc.Ptr(featurePrint), objc.Ptr(error))
	return rv
}

// The feature print data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfeatureprintobservation/3152630-data?language=objc
func (f_ FeaturePrintObservation) Data() []byte {
	rv := objc.Call[[]byte](f_, objc.Sel("data"))
	return rv
}

// The type of each element in the data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfeatureprintobservation/3152632-elementtype?language=objc
func (f_ FeaturePrintObservation) ElementType() ElementType {
	rv := objc.Call[ElementType](f_, objc.Sel("elementType"))
	return rv
}

// The total number of elements in the data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfeatureprintobservation/3152631-elementcount?language=objc
func (f_ FeaturePrintObservation) ElementCount() uint {
	rv := objc.Call[uint](f_, objc.Sel("elementCount"))
	return rv
}
