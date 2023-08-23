// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DetectHumanBodyPoseRequest] class.
var DetectHumanBodyPoseRequestClass = _DetectHumanBodyPoseRequestClass{objc.GetClass("VNDetectHumanBodyPoseRequest")}

type _DetectHumanBodyPoseRequestClass struct {
	objc.Class
}

// An interface definition for the [DetectHumanBodyPoseRequest] class.
type IDetectHumanBodyPoseRequest interface {
	IImageBasedRequest
}

// A request that detects a human body pose. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodyposerequest?language=objc
type DetectHumanBodyPoseRequest struct {
	ImageBasedRequest
}

func DetectHumanBodyPoseRequestFrom(ptr unsafe.Pointer) DetectHumanBodyPoseRequest {
	return DetectHumanBodyPoseRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (dc _DetectHumanBodyPoseRequestClass) Alloc() DetectHumanBodyPoseRequest {
	rv := objc.Call[DetectHumanBodyPoseRequest](dc, objc.Sel("alloc"))
	return rv
}

func DetectHumanBodyPoseRequest_Alloc() DetectHumanBodyPoseRequest {
	return DetectHumanBodyPoseRequestClass.Alloc()
}

func (dc _DetectHumanBodyPoseRequestClass) New() DetectHumanBodyPoseRequest {
	rv := objc.Call[DetectHumanBodyPoseRequest](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDetectHumanBodyPoseRequest() DetectHumanBodyPoseRequest {
	return DetectHumanBodyPoseRequestClass.New()
}

func (d_ DetectHumanBodyPoseRequest) Init() DetectHumanBodyPoseRequest {
	rv := objc.Call[DetectHumanBodyPoseRequest](d_, objc.Sel("init"))
	return rv
}

func (d_ DetectHumanBodyPoseRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) DetectHumanBodyPoseRequest {
	rv := objc.Call[DetectHumanBodyPoseRequest](d_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewDetectHumanBodyPoseRequestWithCompletionHandler(completionHandler RequestCompletionHandler) DetectHumanBodyPoseRequest {
	instance := DetectHumanBodyPoseRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}
