// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GenerateAttentionBasedSaliencyImageRequest] class.
var GenerateAttentionBasedSaliencyImageRequestClass = _GenerateAttentionBasedSaliencyImageRequestClass{objc.GetClass("VNGenerateAttentionBasedSaliencyImageRequest")}

type _GenerateAttentionBasedSaliencyImageRequestClass struct {
	objc.Class
}

// An interface definition for the [GenerateAttentionBasedSaliencyImageRequest] class.
type IGenerateAttentionBasedSaliencyImageRequest interface {
	IImageBasedRequest
}

// An object that produces a heat map that identifies the parts of an image most likely to draw attention. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateattentionbasedsaliencyimagerequest?language=objc
type GenerateAttentionBasedSaliencyImageRequest struct {
	ImageBasedRequest
}

func GenerateAttentionBasedSaliencyImageRequestFrom(ptr unsafe.Pointer) GenerateAttentionBasedSaliencyImageRequest {
	return GenerateAttentionBasedSaliencyImageRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (gc _GenerateAttentionBasedSaliencyImageRequestClass) Alloc() GenerateAttentionBasedSaliencyImageRequest {
	rv := objc.Call[GenerateAttentionBasedSaliencyImageRequest](gc, objc.Sel("alloc"))
	return rv
}

func GenerateAttentionBasedSaliencyImageRequest_Alloc() GenerateAttentionBasedSaliencyImageRequest {
	return GenerateAttentionBasedSaliencyImageRequestClass.Alloc()
}

func (gc _GenerateAttentionBasedSaliencyImageRequestClass) New() GenerateAttentionBasedSaliencyImageRequest {
	rv := objc.Call[GenerateAttentionBasedSaliencyImageRequest](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGenerateAttentionBasedSaliencyImageRequest() GenerateAttentionBasedSaliencyImageRequest {
	return GenerateAttentionBasedSaliencyImageRequestClass.New()
}

func (g_ GenerateAttentionBasedSaliencyImageRequest) Init() GenerateAttentionBasedSaliencyImageRequest {
	rv := objc.Call[GenerateAttentionBasedSaliencyImageRequest](g_, objc.Sel("init"))
	return rv
}

func (g_ GenerateAttentionBasedSaliencyImageRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) GenerateAttentionBasedSaliencyImageRequest {
	rv := objc.Call[GenerateAttentionBasedSaliencyImageRequest](g_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewGenerateAttentionBasedSaliencyImageRequestWithCompletionHandler(completionHandler RequestCompletionHandler) GenerateAttentionBasedSaliencyImageRequest {
	instance := GenerateAttentionBasedSaliencyImageRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}
