// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coreml"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RecognizedPointsObservation] class.
var RecognizedPointsObservationClass = _RecognizedPointsObservationClass{objc.GetClass("VNRecognizedPointsObservation")}

type _RecognizedPointsObservationClass struct {
	objc.Class
}

// An interface definition for the [RecognizedPointsObservation] class.
type IRecognizedPointsObservation interface {
	IObservation
	RecognizedPointsForGroupKeyError(groupKey RecognizedPointGroupKey, error foundation.IError) map[RecognizedPointKey]RecognizedPoint
	RecognizedPointForKeyError(pointKey RecognizedPointKey, error foundation.IError) RecognizedPoint
	KeypointsMultiArrayAndReturnError(error foundation.IError) coreml.MultiArray
	AvailableKeys() []RecognizedPointKey
	AvailableGroupKeys() []RecognizedPointGroupKey
}

// An observation that provides the points the analysis recognized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedpointsobservation?language=objc
type RecognizedPointsObservation struct {
	Observation
}

func RecognizedPointsObservationFrom(ptr unsafe.Pointer) RecognizedPointsObservation {
	return RecognizedPointsObservation{
		Observation: ObservationFrom(ptr),
	}
}

func (rc _RecognizedPointsObservationClass) Alloc() RecognizedPointsObservation {
	rv := objc.Call[RecognizedPointsObservation](rc, objc.Sel("alloc"))
	return rv
}

func RecognizedPointsObservation_Alloc() RecognizedPointsObservation {
	return RecognizedPointsObservationClass.Alloc()
}

func (rc _RecognizedPointsObservationClass) New() RecognizedPointsObservation {
	rv := objc.Call[RecognizedPointsObservation](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRecognizedPointsObservation() RecognizedPointsObservation {
	return RecognizedPointsObservationClass.New()
}

func (r_ RecognizedPointsObservation) Init() RecognizedPointsObservation {
	rv := objc.Call[RecognizedPointsObservation](r_, objc.Sel("init"))
	return rv
}

// Retrieves the recognized points for a key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedpointsobservation/3618962-recognizedpointsforgroupkey?language=objc
func (r_ RecognizedPointsObservation) RecognizedPointsForGroupKeyError(groupKey RecognizedPointGroupKey, error foundation.IError) map[RecognizedPointKey]RecognizedPoint {
	rv := objc.Call[map[RecognizedPointKey]RecognizedPoint](r_, objc.Sel("recognizedPointsForGroupKey:error:"), groupKey, objc.Ptr(error))
	return rv
}

// Retrieves a recognized point for a key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedpointsobservation/3656175-recognizedpointforkey?language=objc
func (r_ RecognizedPointsObservation) RecognizedPointForKeyError(pointKey RecognizedPointKey, error foundation.IError) RecognizedPoint {
	rv := objc.Call[RecognizedPoint](r_, objc.Sel("recognizedPointForKey:error:"), pointKey, objc.Ptr(error))
	return rv
}

// Retrieves the grouping of normalized point coordinates and confidence scores in a format compatible with Core ML. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedpointsobservation/3618961-keypointsmultiarrayandreturnerro?language=objc
func (r_ RecognizedPointsObservation) KeypointsMultiArrayAndReturnError(error foundation.IError) coreml.MultiArray {
	rv := objc.Call[coreml.MultiArray](r_, objc.Sel("keypointsMultiArrayAndReturnError:"), objc.Ptr(error))
	return rv
}

// The available point keys in the observation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedpointsobservation/3656174-availablekeys?language=objc
func (r_ RecognizedPointsObservation) AvailableKeys() []RecognizedPointKey {
	rv := objc.Call[[]RecognizedPointKey](r_, objc.Sel("availableKeys"))
	return rv
}

// The available point group keys in the observation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedpointsobservation/3618960-availablegroupkeys?language=objc
func (r_ RecognizedPointsObservation) AvailableGroupKeys() []RecognizedPointGroupKey {
	rv := objc.Call[[]RecognizedPointGroupKey](r_, objc.Sel("availableGroupKeys"))
	return rv
}
