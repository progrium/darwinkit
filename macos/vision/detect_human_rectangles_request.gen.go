// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DetectHumanRectanglesRequest] class.
var DetectHumanRectanglesRequestClass = _DetectHumanRectanglesRequestClass{objc.GetClass("VNDetectHumanRectanglesRequest")}

type _DetectHumanRectanglesRequestClass struct {
	objc.Class
}

// An interface definition for the [DetectHumanRectanglesRequest] class.
type IDetectHumanRectanglesRequest interface {
	IImageBasedRequest
	UpperBodyOnly() bool
	SetUpperBodyOnly(value bool)
}

// A request that finds rectangular regions that contain people in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanrectanglesrequest?language=objc
type DetectHumanRectanglesRequest struct {
	ImageBasedRequest
}

func DetectHumanRectanglesRequestFrom(ptr unsafe.Pointer) DetectHumanRectanglesRequest {
	return DetectHumanRectanglesRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (dc _DetectHumanRectanglesRequestClass) Alloc() DetectHumanRectanglesRequest {
	rv := objc.Call[DetectHumanRectanglesRequest](dc, objc.Sel("alloc"))
	return rv
}

func DetectHumanRectanglesRequest_Alloc() DetectHumanRectanglesRequest {
	return DetectHumanRectanglesRequestClass.Alloc()
}

func (dc _DetectHumanRectanglesRequestClass) New() DetectHumanRectanglesRequest {
	rv := objc.Call[DetectHumanRectanglesRequest](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDetectHumanRectanglesRequest() DetectHumanRectanglesRequest {
	return DetectHumanRectanglesRequestClass.New()
}

func (d_ DetectHumanRectanglesRequest) Init() DetectHumanRectanglesRequest {
	rv := objc.Call[DetectHumanRectanglesRequest](d_, objc.Sel("init"))
	return rv
}

func (d_ DetectHumanRectanglesRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) DetectHumanRectanglesRequest {
	rv := objc.Call[DetectHumanRectanglesRequest](d_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewDetectHumanRectanglesRequestWithCompletionHandler(completionHandler RequestCompletionHandler) DetectHumanRectanglesRequest {
	instance := DetectHumanRectanglesRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}

// A Boolean value that indicates whether the request requires detecting a full body or upper body only to produce a result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanrectanglesrequest/3750977-upperbodyonly?language=objc
func (d_ DetectHumanRectanglesRequest) UpperBodyOnly() bool {
	rv := objc.Call[bool](d_, objc.Sel("upperBodyOnly"))
	return rv
}

// A Boolean value that indicates whether the request requires detecting a full body or upper body only to produce a result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanrectanglesrequest/3750977-upperbodyonly?language=objc
func (d_ DetectHumanRectanglesRequest) SetUpperBodyOnly(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setUpperBodyOnly:"), value)
}
