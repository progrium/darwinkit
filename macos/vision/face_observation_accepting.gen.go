// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"github.com/progrium/macdriver/objc"
)

// An image analysis request that operates on face observations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservationaccepting?language=objc
type PFaceObservationAccepting interface {
	// optional
	SetInputFaceObservations(value []FaceObservation)
	HasSetInputFaceObservations() bool

	// optional
	InputFaceObservations() []IFaceObservation
	HasInputFaceObservations() bool
}

// A concrete type wrapper for the [PFaceObservationAccepting] protocol.
type FaceObservationAcceptingWrapper struct {
	objc.Object
}

func (f_ FaceObservationAcceptingWrapper) HasSetInputFaceObservations() bool {
	return f_.RespondsToSelector(objc.Sel("setInputFaceObservations:"))
}

// An array of VNFaceObservation objects to process as part of the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservationaccepting/2877424-inputfaceobservations?language=objc
func (f_ FaceObservationAcceptingWrapper) SetInputFaceObservations(value []IFaceObservation) {
	objc.Call[objc.Void](f_, objc.Sel("setInputFaceObservations:"), value)
}

func (f_ FaceObservationAcceptingWrapper) HasInputFaceObservations() bool {
	return f_.RespondsToSelector(objc.Sel("inputFaceObservations"))
}

// An array of VNFaceObservation objects to process as part of the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservationaccepting/2877424-inputfaceobservations?language=objc
func (f_ FaceObservationAcceptingWrapper) InputFaceObservations() []FaceObservation {
	rv := objc.Call[[]FaceObservation](f_, objc.Sel("inputFaceObservations"))
	return rv
}
