// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TrackingRequest] class.
var TrackingRequestClass = _TrackingRequestClass{objc.GetClass("VNTrackingRequest")}

type _TrackingRequestClass struct {
	objc.Class
}

// An interface definition for the [TrackingRequest] class.
type ITrackingRequest interface {
	IImageBasedRequest
	IsLastFrame() bool
	SetLastFrame(value bool)
	InputObservation() DetectedObjectObservation
	SetInputObservation(value IDetectedObjectObservation)
	TrackingLevel() RequestTrackingLevel
	SetTrackingLevel(value RequestTrackingLevel)
}

// The abstract superclass for image analysis requests that track unique features across multiple images or video frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackingrequest?language=objc
type TrackingRequest struct {
	ImageBasedRequest
}

func TrackingRequestFrom(ptr unsafe.Pointer) TrackingRequest {
	return TrackingRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (tc _TrackingRequestClass) Alloc() TrackingRequest {
	rv := objc.Call[TrackingRequest](tc, objc.Sel("alloc"))
	return rv
}

func TrackingRequest_Alloc() TrackingRequest {
	return TrackingRequestClass.Alloc()
}

func (tc _TrackingRequestClass) New() TrackingRequest {
	rv := objc.Call[TrackingRequest](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTrackingRequest() TrackingRequest {
	return TrackingRequestClass.New()
}

func (t_ TrackingRequest) Init() TrackingRequest {
	rv := objc.Call[TrackingRequest](t_, objc.Sel("init"))
	return rv
}

func (t_ TrackingRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) TrackingRequest {
	rv := objc.Call[TrackingRequest](t_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewTrackingRequestWithCompletionHandler(completionHandler RequestCompletionHandler) TrackingRequest {
	instance := TrackingRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}

// A Boolean that indicates the last frame in a tracking sequence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackingrequest/2887356-lastframe?language=objc
func (t_ TrackingRequest) IsLastFrame() bool {
	rv := objc.Call[bool](t_, objc.Sel("isLastFrame"))
	return rv
}

// A Boolean that indicates the last frame in a tracking sequence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackingrequest/2887356-lastframe?language=objc
func (t_ TrackingRequest) SetLastFrame(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setLastFrame:"), value)
}

// The observation object defining a region to track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackingrequest/2887350-inputobservation?language=objc
func (t_ TrackingRequest) InputObservation() DetectedObjectObservation {
	rv := objc.Call[DetectedObjectObservation](t_, objc.Sel("inputObservation"))
	return rv
}

// The observation object defining a region to track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackingrequest/2887350-inputobservation?language=objc
func (t_ TrackingRequest) SetInputObservation(value IDetectedObjectObservation) {
	objc.Call[objc.Void](t_, objc.Sel("setInputObservation:"), objc.Ptr(value))
}

// A value for specifying whether to prioritize speed or location accuracy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackingrequest/2887353-trackinglevel?language=objc
func (t_ TrackingRequest) TrackingLevel() RequestTrackingLevel {
	rv := objc.Call[RequestTrackingLevel](t_, objc.Sel("trackingLevel"))
	return rv
}

// A value for specifying whether to prioritize speed or location accuracy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackingrequest/2887353-trackinglevel?language=objc
func (t_ TrackingRequest) SetTrackingLevel(value RequestTrackingLevel) {
	objc.Call[objc.Void](t_, objc.Sel("setTrackingLevel:"), value)
}
