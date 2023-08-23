// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ClassificationObservation] class.
var ClassificationObservationClass = _ClassificationObservationClass{objc.GetClass("VNClassificationObservation")}

type _ClassificationObservationClass struct {
	objc.Class
}

// An interface definition for the [ClassificationObservation] class.
type IClassificationObservation interface {
	IObservation
	HasMinimumRecallForPrecision(minimumRecall float64, precision float64) bool
	HasMinimumPrecisionForRecall(minimumPrecision float64, recall float64) bool
	HasPrecisionRecallCurve() bool
	Identifier() string
}

// An object that represents classification information that an image analysis request produces. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnclassificationobservation?language=objc
type ClassificationObservation struct {
	Observation
}

func ClassificationObservationFrom(ptr unsafe.Pointer) ClassificationObservation {
	return ClassificationObservation{
		Observation: ObservationFrom(ptr),
	}
}

func (cc _ClassificationObservationClass) Alloc() ClassificationObservation {
	rv := objc.Call[ClassificationObservation](cc, objc.Sel("alloc"))
	return rv
}

func ClassificationObservation_Alloc() ClassificationObservation {
	return ClassificationObservationClass.Alloc()
}

func (cc _ClassificationObservationClass) New() ClassificationObservation {
	rv := objc.Call[ClassificationObservation](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewClassificationObservation() ClassificationObservation {
	return ClassificationObservationClass.New()
}

func (c_ ClassificationObservation) Init() ClassificationObservation {
	rv := objc.Call[ClassificationObservation](c_, objc.Sel("init"))
	return rv
}

// Determines whether the observation for a specific precision has a minimum recall value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnclassificationobservation/3152625-hasminimumrecall?language=objc
func (c_ ClassificationObservation) HasMinimumRecallForPrecision(minimumRecall float64, precision float64) bool {
	rv := objc.Call[bool](c_, objc.Sel("hasMinimumRecall:forPrecision:"), minimumRecall, precision)
	return rv
}

// Determines whether the observation for a specific recall has a minimum precision value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnclassificationobservation/3152624-hasminimumprecision?language=objc
func (c_ ClassificationObservation) HasMinimumPrecisionForRecall(minimumPrecision float64, recall float64) bool {
	rv := objc.Call[bool](c_, objc.Sel("hasMinimumPrecision:forRecall:"), minimumPrecision, recall)
	return rv
}

// A Boolean variable indicating whether the observation contains precision and recall curves. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnclassificationobservation/3152626-hasprecisionrecallcurve?language=objc
func (c_ ClassificationObservation) HasPrecisionRecallCurve() bool {
	rv := objc.Call[bool](c_, objc.Sel("hasPrecisionRecallCurve"))
	return rv
}

// Classification label identifying the type of observation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnclassificationobservation/2867259-identifier?language=objc
func (c_ ClassificationObservation) Identifier() string {
	rv := objc.Call[string](c_, objc.Sel("identifier"))
	return rv
}
