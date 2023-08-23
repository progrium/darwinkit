// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [HumanBodyPoseObservation] class.
var HumanBodyPoseObservationClass = _HumanBodyPoseObservationClass{objc.GetClass("VNHumanBodyPoseObservation")}

type _HumanBodyPoseObservationClass struct {
	objc.Class
}

// An interface definition for the [HumanBodyPoseObservation] class.
type IHumanBodyPoseObservation interface {
	IRecognizedPointsObservation
	RecognizedPointsForJointsGroupNameError(jointsGroupName HumanBodyPoseObservationJointsGroupName, error foundation.IError) map[HumanBodyPoseObservationJointName]RecognizedPoint
	RecognizedPointForJointNameError(jointName HumanBodyPoseObservationJointName, error foundation.IError) RecognizedPoint
	AvailableJointNames() []HumanBodyPoseObservationJointName
	AvailableJointsGroupNames() []HumanBodyPoseObservationJointsGroupName
}

// An observation that provides the body points the analysis recognized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyposeobservation?language=objc
type HumanBodyPoseObservation struct {
	RecognizedPointsObservation
}

func HumanBodyPoseObservationFrom(ptr unsafe.Pointer) HumanBodyPoseObservation {
	return HumanBodyPoseObservation{
		RecognizedPointsObservation: RecognizedPointsObservationFrom(ptr),
	}
}

func (hc _HumanBodyPoseObservationClass) Alloc() HumanBodyPoseObservation {
	rv := objc.Call[HumanBodyPoseObservation](hc, objc.Sel("alloc"))
	return rv
}

func HumanBodyPoseObservation_Alloc() HumanBodyPoseObservation {
	return HumanBodyPoseObservationClass.Alloc()
}

func (hc _HumanBodyPoseObservationClass) New() HumanBodyPoseObservation {
	rv := objc.Call[HumanBodyPoseObservation](hc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewHumanBodyPoseObservation() HumanBodyPoseObservation {
	return HumanBodyPoseObservationClass.New()
}

func (h_ HumanBodyPoseObservation) Init() HumanBodyPoseObservation {
	rv := objc.Call[HumanBodyPoseObservation](h_, objc.Sel("init"))
	return rv
}

// Retrieves the recognized points associated with the joint group name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyposeobservation/3675604-recognizedpointsforjointsgroupna?language=objc
func (h_ HumanBodyPoseObservation) RecognizedPointsForJointsGroupNameError(jointsGroupName HumanBodyPoseObservationJointsGroupName, error foundation.IError) map[HumanBodyPoseObservationJointName]RecognizedPoint {
	rv := objc.Call[map[HumanBodyPoseObservationJointName]RecognizedPoint](h_, objc.Sel("recognizedPointsForJointsGroupName:error:"), jointsGroupName, objc.Ptr(error))
	return rv
}

// Retrieves the recognized point for a joint name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyposeobservation/3675603-recognizedpointforjointname?language=objc
func (h_ HumanBodyPoseObservation) RecognizedPointForJointNameError(jointName HumanBodyPoseObservationJointName, error foundation.IError) RecognizedPoint {
	rv := objc.Call[RecognizedPoint](h_, objc.Sel("recognizedPointForJointName:error:"), jointName, objc.Ptr(error))
	return rv
}

// The names of the available joints in the observation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyposeobservation/3675601-availablejointnames?language=objc
func (h_ HumanBodyPoseObservation) AvailableJointNames() []HumanBodyPoseObservationJointName {
	rv := objc.Call[[]HumanBodyPoseObservationJointName](h_, objc.Sel("availableJointNames"))
	return rv
}

// The available joint group names in the observation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyposeobservation/3675602-availablejointsgroupnames?language=objc
func (h_ HumanBodyPoseObservation) AvailableJointsGroupNames() []HumanBodyPoseObservationJointsGroupName {
	rv := objc.Call[[]HumanBodyPoseObservationJointsGroupName](h_, objc.Sel("availableJointsGroupNames"))
	return rv
}
