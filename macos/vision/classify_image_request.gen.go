// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ClassifyImageRequest] class.
var ClassifyImageRequestClass = _ClassifyImageRequestClass{objc.GetClass("VNClassifyImageRequest")}

type _ClassifyImageRequestClass struct {
	objc.Class
}

// An interface definition for the [ClassifyImageRequest] class.
type IClassifyImageRequest interface {
	IImageBasedRequest
	SupportedIdentifiersAndReturnError(error foundation.IError) []string
}

// A request to classify an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnclassifyimagerequest?language=objc
type ClassifyImageRequest struct {
	ImageBasedRequest
}

func ClassifyImageRequestFrom(ptr unsafe.Pointer) ClassifyImageRequest {
	return ClassifyImageRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (cc _ClassifyImageRequestClass) Alloc() ClassifyImageRequest {
	rv := objc.Call[ClassifyImageRequest](cc, objc.Sel("alloc"))
	return rv
}

func ClassifyImageRequest_Alloc() ClassifyImageRequest {
	return ClassifyImageRequestClass.Alloc()
}

func (cc _ClassifyImageRequestClass) New() ClassifyImageRequest {
	rv := objc.Call[ClassifyImageRequest](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewClassifyImageRequest() ClassifyImageRequest {
	return ClassifyImageRequestClass.New()
}

func (c_ ClassifyImageRequest) Init() ClassifyImageRequest {
	rv := objc.Call[ClassifyImageRequest](c_, objc.Sel("init"))
	return rv
}

func (c_ ClassifyImageRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) ClassifyImageRequest {
	rv := objc.Call[ClassifyImageRequest](c_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewClassifyImageRequestWithCompletionHandler(completionHandler RequestCompletionHandler) ClassifyImageRequest {
	instance := ClassifyImageRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}

// Returns the classification identifiers that the request supports in its current configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnclassifyimagerequest/3750957-supportedidentifiersandreturnerr?language=objc
func (c_ ClassifyImageRequest) SupportedIdentifiersAndReturnError(error foundation.IError) []string {
	rv := objc.Call[[]string](c_, objc.Sel("supportedIdentifiersAndReturnError:"), objc.Ptr(error))
	return rv
}
