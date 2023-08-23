// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FaceObservation] class.
var FaceObservationClass = _FaceObservationClass{objc.GetClass("VNFaceObservation")}

type _FaceObservationClass struct {
	objc.Class
}

// An interface definition for the [FaceObservation] class.
type IFaceObservation interface {
	IDetectedObjectObservation
	Roll() foundation.Number
	Yaw() foundation.Number
	FaceCaptureQuality() foundation.Number
	Pitch() foundation.Number
	Landmarks() FaceLandmarks2D
}

// Face or facial-feature information that an image analysis request detects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservation?language=objc
type FaceObservation struct {
	DetectedObjectObservation
}

func FaceObservationFrom(ptr unsafe.Pointer) FaceObservation {
	return FaceObservation{
		DetectedObjectObservation: DetectedObjectObservationFrom(ptr),
	}
}

func (fc _FaceObservationClass) FaceObservationWithRequestRevisionBoundingBoxRollYawPitch(requestRevision uint, boundingBox coregraphics.Rect, roll foundation.INumber, yaw foundation.INumber, pitch foundation.INumber) FaceObservation {
	rv := objc.Call[FaceObservation](fc, objc.Sel("faceObservationWithRequestRevision:boundingBox:roll:yaw:pitch:"), requestRevision, boundingBox, objc.Ptr(roll), objc.Ptr(yaw), objc.Ptr(pitch))
	return rv
}

// Creates an observation that contains the roll, yaw, and pitch of the face. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservation/3763078-faceobservationwithrequestrevisi?language=objc
func FaceObservation_FaceObservationWithRequestRevisionBoundingBoxRollYawPitch(requestRevision uint, boundingBox coregraphics.Rect, roll foundation.INumber, yaw foundation.INumber, pitch foundation.INumber) FaceObservation {
	return FaceObservationClass.FaceObservationWithRequestRevisionBoundingBoxRollYawPitch(requestRevision, boundingBox, roll, yaw, pitch)
}

func (fc _FaceObservationClass) Alloc() FaceObservation {
	rv := objc.Call[FaceObservation](fc, objc.Sel("alloc"))
	return rv
}

func FaceObservation_Alloc() FaceObservation {
	return FaceObservationClass.Alloc()
}

func (fc _FaceObservationClass) New() FaceObservation {
	rv := objc.Call[FaceObservation](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFaceObservation() FaceObservation {
	return FaceObservationClass.New()
}

func (f_ FaceObservation) Init() FaceObservation {
	rv := objc.Call[FaceObservation](f_, objc.Sel("init"))
	return rv
}

func (fc _FaceObservationClass) ObservationWithBoundingBox(boundingBox coregraphics.Rect) FaceObservation {
	rv := objc.Call[FaceObservation](fc, objc.Sel("observationWithBoundingBox:"), boundingBox)
	return rv
}

// Creates an observation with a bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/2879297-observationwithboundingbox?language=objc
func FaceObservation_ObservationWithBoundingBox(boundingBox coregraphics.Rect) FaceObservation {
	return FaceObservationClass.ObservationWithBoundingBox(boundingBox)
}

func (fc _FaceObservationClass) ObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox coregraphics.Rect) FaceObservation {
	rv := objc.Call[FaceObservation](fc, objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return rv
}

// Creates an observation with a revision number and bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/2967105-observationwithrequestrevision?language=objc
func FaceObservation_ObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox coregraphics.Rect) FaceObservation {
	return FaceObservationClass.ObservationWithRequestRevisionBoundingBox(requestRevision, boundingBox)
}

// The roll angle of a face in radians. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservation/2980939-roll?language=objc
func (f_ FaceObservation) Roll() foundation.Number {
	rv := objc.Call[foundation.Number](f_, objc.Sel("roll"))
	return rv
}

// The yaw angle of a face in radians. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservation/2980940-yaw?language=objc
func (f_ FaceObservation) Yaw() foundation.Number {
	rv := objc.Call[foundation.Number](f_, objc.Sel("yaw"))
	return rv
}

// A value that indicates the quality of the face capture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservation/3152627-facecapturequality?language=objc
func (f_ FaceObservation) FaceCaptureQuality() foundation.Number {
	rv := objc.Call[foundation.Number](f_, objc.Sel("faceCaptureQuality"))
	return rv
}

// The pitch angle of a face in radians. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservation/3750998-pitch?language=objc
func (f_ FaceObservation) Pitch() foundation.Number {
	rv := objc.Call[foundation.Number](f_, objc.Sel("pitch"))
	return rv
}

// The facial features of the detected face. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservation/2867250-landmarks?language=objc
func (f_ FaceObservation) Landmarks() FaceLandmarks2D {
	rv := objc.Call[FaceLandmarks2D](f_, objc.Sel("landmarks"))
	return rv
}
