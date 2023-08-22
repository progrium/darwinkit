// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DetectHumanHandPoseRequest] class.
var DetectHumanHandPoseRequestClass = _DetectHumanHandPoseRequestClass{objc.GetClass("VNDetectHumanHandPoseRequest")}

type _DetectHumanHandPoseRequestClass struct {
	objc.Class
}

// An interface definition for the [DetectHumanHandPoseRequest] class.
type IDetectHumanHandPoseRequest interface {
	IImageBasedRequest
	MaximumHandCount() uint
	SetMaximumHandCount(value uint)
}

// A request that detects a human hand pose. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest?language=objc
type DetectHumanHandPoseRequest struct {
	ImageBasedRequest
}

func DetectHumanHandPoseRequestFrom(ptr unsafe.Pointer) DetectHumanHandPoseRequest {
	return DetectHumanHandPoseRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (dc _DetectHumanHandPoseRequestClass) Alloc() DetectHumanHandPoseRequest {
	rv := objc.Call[DetectHumanHandPoseRequest](dc, objc.Sel("alloc"))
	return rv
}

func DetectHumanHandPoseRequest_Alloc() DetectHumanHandPoseRequest {
	return DetectHumanHandPoseRequestClass.Alloc()
}

func (dc _DetectHumanHandPoseRequestClass) New() DetectHumanHandPoseRequest {
	rv := objc.Call[DetectHumanHandPoseRequest](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDetectHumanHandPoseRequest() DetectHumanHandPoseRequest {
	return DetectHumanHandPoseRequestClass.New()
}

func (d_ DetectHumanHandPoseRequest) Init() DetectHumanHandPoseRequest {
	rv := objc.Call[DetectHumanHandPoseRequest](d_, objc.Sel("init"))
	return rv
}

func (d_ DetectHumanHandPoseRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) DetectHumanHandPoseRequest {
	rv := objc.Call[DetectHumanHandPoseRequest](d_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewDetectHumanHandPoseRequestWithCompletionHandler(completionHandler RequestCompletionHandler) DetectHumanHandPoseRequest {
	instance := DetectHumanHandPoseRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}

// The maximum number of hands to detect in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/3571271-maximumhandcount?language=objc
func (d_ DetectHumanHandPoseRequest) MaximumHandCount() uint {
	rv := objc.Call[uint](d_, objc.Sel("maximumHandCount"))
	return rv
}

// The maximum number of hands to detect in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/3571271-maximumhandcount?language=objc
func (d_ DetectHumanHandPoseRequest) SetMaximumHandCount(value uint) {
	objc.Call[objc.Void](d_, objc.Sel("setMaximumHandCount:"), value)
}
