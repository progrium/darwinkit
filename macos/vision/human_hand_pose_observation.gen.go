// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [HumanHandPoseObservation] class.
var HumanHandPoseObservationClass = _HumanHandPoseObservationClass{objc.GetClass("VNHumanHandPoseObservation")}

type _HumanHandPoseObservationClass struct {
	objc.Class
}

// An interface definition for the [HumanHandPoseObservation] class.
type IHumanHandPoseObservation interface {
	IRecognizedPointsObservation
	RecognizedPointsForJointsGroupNameError(jointsGroupName HumanHandPoseObservationJointsGroupName, error foundation.IError) map[HumanHandPoseObservationJointName]RecognizedPoint
	RecognizedPointForJointNameError(jointName HumanHandPoseObservationJointName, error foundation.IError) RecognizedPoint
	Chirality() Chirality
	AvailableJointNames() []HumanHandPoseObservationJointName
	AvailableJointsGroupNames() []HumanHandPoseObservationJointsGroupName
}

// An observation that provides the hand points the analysis recognized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanhandposeobservation?language=objc
type HumanHandPoseObservation struct {
	RecognizedPointsObservation
}

func HumanHandPoseObservationFrom(ptr unsafe.Pointer) HumanHandPoseObservation {
	return HumanHandPoseObservation{
		RecognizedPointsObservation: RecognizedPointsObservationFrom(ptr),
	}
}

func (hc _HumanHandPoseObservationClass) Alloc() HumanHandPoseObservation {
	rv := objc.Call[HumanHandPoseObservation](hc, objc.Sel("alloc"))
	return rv
}

func HumanHandPoseObservation_Alloc() HumanHandPoseObservation {
	return HumanHandPoseObservationClass.Alloc()
}

func (hc _HumanHandPoseObservationClass) New() HumanHandPoseObservation {
	rv := objc.Call[HumanHandPoseObservation](hc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewHumanHandPoseObservation() HumanHandPoseObservation {
	return HumanHandPoseObservationClass.New()
}

func (h_ HumanHandPoseObservation) Init() HumanHandPoseObservation {
	rv := objc.Call[HumanHandPoseObservation](h_, objc.Sel("init"))
	return rv
}

// Retrieves the recognized points associated with the joint group name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanhandposeobservation/3675640-recognizedpointsforjointsgroupna?language=objc
func (h_ HumanHandPoseObservation) RecognizedPointsForJointsGroupNameError(jointsGroupName HumanHandPoseObservationJointsGroupName, error foundation.IError) map[HumanHandPoseObservationJointName]RecognizedPoint {
	rv := objc.Call[map[HumanHandPoseObservationJointName]RecognizedPoint](h_, objc.Sel("recognizedPointsForJointsGroupName:error:"), jointsGroupName, objc.Ptr(error))
	return rv
}

// Retrieves the recognized point for a joint name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanhandposeobservation/3675639-recognizedpointforjointname?language=objc
func (h_ HumanHandPoseObservation) RecognizedPointForJointNameError(jointName HumanHandPoseObservationJointName, error foundation.IError) RecognizedPoint {
	rv := objc.Call[RecognizedPoint](h_, objc.Sel("recognizedPointForJointName:error:"), jointName, objc.Ptr(error))
	return rv
}

// The chirality, or handedness, of a pose. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanhandposeobservation/3750971-chirality?language=objc
func (h_ HumanHandPoseObservation) Chirality() Chirality {
	rv := objc.Call[Chirality](h_, objc.Sel("chirality"))
	return rv
}

// The names of the available joints in the observation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanhandposeobservation/3675637-availablejointnames?language=objc
func (h_ HumanHandPoseObservation) AvailableJointNames() []HumanHandPoseObservationJointName {
	rv := objc.Call[[]HumanHandPoseObservationJointName](h_, objc.Sel("availableJointNames"))
	return rv
}

// The joint group names available in the observation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanhandposeobservation/3675638-availablejointsgroupnames?language=objc
func (h_ HumanHandPoseObservation) AvailableJointsGroupNames() []HumanHandPoseObservationJointsGroupName {
	rv := objc.Call[[]HumanHandPoseObservationJointsGroupName](h_, objc.Sel("availableJointsGroupNames"))
	return rv
}
