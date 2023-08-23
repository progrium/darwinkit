// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageBasedRequest] class.
var ImageBasedRequestClass = _ImageBasedRequestClass{objc.GetClass("VNImageBasedRequest")}

type _ImageBasedRequestClass struct {
	objc.Class
}

// An interface definition for the [ImageBasedRequest] class.
type IImageBasedRequest interface {
	IRequest
	RegionOfInterest() coregraphics.Rect
	SetRegionOfInterest(value coregraphics.Rect)
}

// The abstract superclass for image analysis requests that focus on a specific part of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimagebasedrequest?language=objc
type ImageBasedRequest struct {
	Request
}

func ImageBasedRequestFrom(ptr unsafe.Pointer) ImageBasedRequest {
	return ImageBasedRequest{
		Request: RequestFrom(ptr),
	}
}

func (ic _ImageBasedRequestClass) Alloc() ImageBasedRequest {
	rv := objc.Call[ImageBasedRequest](ic, objc.Sel("alloc"))
	return rv
}

func ImageBasedRequest_Alloc() ImageBasedRequest {
	return ImageBasedRequestClass.Alloc()
}

func (ic _ImageBasedRequestClass) New() ImageBasedRequest {
	rv := objc.Call[ImageBasedRequest](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageBasedRequest() ImageBasedRequest {
	return ImageBasedRequestClass.New()
}

func (i_ ImageBasedRequest) Init() ImageBasedRequest {
	rv := objc.Call[ImageBasedRequest](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageBasedRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) ImageBasedRequest {
	rv := objc.Call[ImageBasedRequest](i_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewImageBasedRequestWithCompletionHandler(completionHandler RequestCompletionHandler) ImageBasedRequest {
	instance := ImageBasedRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}

// The region of the image in which Vision will perform the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimagebasedrequest/2877482-regionofinterest?language=objc
func (i_ ImageBasedRequest) RegionOfInterest() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](i_, objc.Sel("regionOfInterest"))
	return rv
}

// The region of the image in which Vision will perform the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimagebasedrequest/2877482-regionofinterest?language=objc
func (i_ ImageBasedRequest) SetRegionOfInterest(value coregraphics.Rect) {
	objc.Call[objc.Void](i_, objc.Sel("setRegionOfInterest:"), value)
}
