// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DetectFaceRectanglesRequest] class.
var DetectFaceRectanglesRequestClass = _DetectFaceRectanglesRequestClass{objc.GetClass("VNDetectFaceRectanglesRequest")}

type _DetectFaceRectanglesRequestClass struct {
	objc.Class
}

// An interface definition for the [DetectFaceRectanglesRequest] class.
type IDetectFaceRectanglesRequest interface {
	IImageBasedRequest
}

// A request that finds faces within an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacerectanglesrequest?language=objc
type DetectFaceRectanglesRequest struct {
	ImageBasedRequest
}

func DetectFaceRectanglesRequestFrom(ptr unsafe.Pointer) DetectFaceRectanglesRequest {
	return DetectFaceRectanglesRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (dc _DetectFaceRectanglesRequestClass) Alloc() DetectFaceRectanglesRequest {
	rv := objc.Call[DetectFaceRectanglesRequest](dc, objc.Sel("alloc"))
	return rv
}

func DetectFaceRectanglesRequest_Alloc() DetectFaceRectanglesRequest {
	return DetectFaceRectanglesRequestClass.Alloc()
}

func (dc _DetectFaceRectanglesRequestClass) New() DetectFaceRectanglesRequest {
	rv := objc.Call[DetectFaceRectanglesRequest](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDetectFaceRectanglesRequest() DetectFaceRectanglesRequest {
	return DetectFaceRectanglesRequestClass.New()
}

func (d_ DetectFaceRectanglesRequest) Init() DetectFaceRectanglesRequest {
	rv := objc.Call[DetectFaceRectanglesRequest](d_, objc.Sel("init"))
	return rv
}

func (d_ DetectFaceRectanglesRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) DetectFaceRectanglesRequest {
	rv := objc.Call[DetectFaceRectanglesRequest](d_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewDetectFaceRectanglesRequestWithCompletionHandler(completionHandler RequestCompletionHandler) DetectFaceRectanglesRequest {
	instance := DetectFaceRectanglesRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}
