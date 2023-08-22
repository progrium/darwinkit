// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DetectFaceCaptureQualityRequest] class.
var DetectFaceCaptureQualityRequestClass = _DetectFaceCaptureQualityRequestClass{objc.GetClass("VNDetectFaceCaptureQualityRequest")}

type _DetectFaceCaptureQualityRequestClass struct {
	objc.Class
}

// An interface definition for the [DetectFaceCaptureQualityRequest] class.
type IDetectFaceCaptureQualityRequest interface {
	IImageBasedRequest
}

// A request that produces a floating-point number that represents the capture quality of a face in a photo. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacecapturequalityrequest?language=objc
type DetectFaceCaptureQualityRequest struct {
	ImageBasedRequest
}

func DetectFaceCaptureQualityRequestFrom(ptr unsafe.Pointer) DetectFaceCaptureQualityRequest {
	return DetectFaceCaptureQualityRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (dc _DetectFaceCaptureQualityRequestClass) Alloc() DetectFaceCaptureQualityRequest {
	rv := objc.Call[DetectFaceCaptureQualityRequest](dc, objc.Sel("alloc"))
	return rv
}

func DetectFaceCaptureQualityRequest_Alloc() DetectFaceCaptureQualityRequest {
	return DetectFaceCaptureQualityRequestClass.Alloc()
}

func (dc _DetectFaceCaptureQualityRequestClass) New() DetectFaceCaptureQualityRequest {
	rv := objc.Call[DetectFaceCaptureQualityRequest](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDetectFaceCaptureQualityRequest() DetectFaceCaptureQualityRequest {
	return DetectFaceCaptureQualityRequestClass.New()
}

func (d_ DetectFaceCaptureQualityRequest) Init() DetectFaceCaptureQualityRequest {
	rv := objc.Call[DetectFaceCaptureQualityRequest](d_, objc.Sel("init"))
	return rv
}

func (d_ DetectFaceCaptureQualityRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) DetectFaceCaptureQualityRequest {
	rv := objc.Call[DetectFaceCaptureQualityRequest](d_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewDetectFaceCaptureQualityRequestWithCompletionHandler(completionHandler RequestCompletionHandler) DetectFaceCaptureQualityRequest {
	instance := DetectFaceCaptureQualityRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}
