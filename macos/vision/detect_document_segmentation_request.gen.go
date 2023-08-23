// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DetectDocumentSegmentationRequest] class.
var DetectDocumentSegmentationRequestClass = _DetectDocumentSegmentationRequestClass{objc.GetClass("VNDetectDocumentSegmentationRequest")}

type _DetectDocumentSegmentationRequestClass struct {
	objc.Class
}

// An interface definition for the [DetectDocumentSegmentationRequest] class.
type IDetectDocumentSegmentationRequest interface {
	IImageBasedRequest
}

// An object that detects rectangular regions that contain text in the input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectdocumentsegmentationrequest?language=objc
type DetectDocumentSegmentationRequest struct {
	ImageBasedRequest
}

func DetectDocumentSegmentationRequestFrom(ptr unsafe.Pointer) DetectDocumentSegmentationRequest {
	return DetectDocumentSegmentationRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (dc _DetectDocumentSegmentationRequestClass) Alloc() DetectDocumentSegmentationRequest {
	rv := objc.Call[DetectDocumentSegmentationRequest](dc, objc.Sel("alloc"))
	return rv
}

func DetectDocumentSegmentationRequest_Alloc() DetectDocumentSegmentationRequest {
	return DetectDocumentSegmentationRequestClass.Alloc()
}

func (dc _DetectDocumentSegmentationRequestClass) New() DetectDocumentSegmentationRequest {
	rv := objc.Call[DetectDocumentSegmentationRequest](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDetectDocumentSegmentationRequest() DetectDocumentSegmentationRequest {
	return DetectDocumentSegmentationRequestClass.New()
}

func (d_ DetectDocumentSegmentationRequest) Init() DetectDocumentSegmentationRequest {
	rv := objc.Call[DetectDocumentSegmentationRequest](d_, objc.Sel("init"))
	return rv
}

func (d_ DetectDocumentSegmentationRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) DetectDocumentSegmentationRequest {
	rv := objc.Call[DetectDocumentSegmentationRequest](d_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewDetectDocumentSegmentationRequestWithCompletionHandler(completionHandler RequestCompletionHandler) DetectDocumentSegmentationRequest {
	instance := DetectDocumentSegmentationRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}
