// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DetectHorizonRequest] class.
var DetectHorizonRequestClass = _DetectHorizonRequestClass{objc.GetClass("VNDetectHorizonRequest")}

type _DetectHorizonRequestClass struct {
	objc.Class
}

// An interface definition for the [DetectHorizonRequest] class.
type IDetectHorizonRequest interface {
	IImageBasedRequest
}

// An image analysis request that determines the horizon angle in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthorizonrequest?language=objc
type DetectHorizonRequest struct {
	ImageBasedRequest
}

func DetectHorizonRequestFrom(ptr unsafe.Pointer) DetectHorizonRequest {
	return DetectHorizonRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (dc _DetectHorizonRequestClass) Alloc() DetectHorizonRequest {
	rv := objc.Call[DetectHorizonRequest](dc, objc.Sel("alloc"))
	return rv
}

func DetectHorizonRequest_Alloc() DetectHorizonRequest {
	return DetectHorizonRequestClass.Alloc()
}

func (dc _DetectHorizonRequestClass) New() DetectHorizonRequest {
	rv := objc.Call[DetectHorizonRequest](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDetectHorizonRequest() DetectHorizonRequest {
	return DetectHorizonRequestClass.New()
}

func (d_ DetectHorizonRequest) Init() DetectHorizonRequest {
	rv := objc.Call[DetectHorizonRequest](d_, objc.Sel("init"))
	return rv
}

func (d_ DetectHorizonRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) DetectHorizonRequest {
	rv := objc.Call[DetectHorizonRequest](d_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewDetectHorizonRequestWithCompletionHandler(completionHandler RequestCompletionHandler) DetectHorizonRequest {
	instance := DetectHorizonRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}
