// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GenerateObjectnessBasedSaliencyImageRequest] class.
var GenerateObjectnessBasedSaliencyImageRequestClass = _GenerateObjectnessBasedSaliencyImageRequestClass{objc.GetClass("VNGenerateObjectnessBasedSaliencyImageRequest")}

type _GenerateObjectnessBasedSaliencyImageRequestClass struct {
	objc.Class
}

// An interface definition for the [GenerateObjectnessBasedSaliencyImageRequest] class.
type IGenerateObjectnessBasedSaliencyImageRequest interface {
	IImageBasedRequest
}

// A request that generates a heat map that identifies the parts of an image most likely to represent objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateobjectnessbasedsaliencyimagerequest?language=objc
type GenerateObjectnessBasedSaliencyImageRequest struct {
	ImageBasedRequest
}

func GenerateObjectnessBasedSaliencyImageRequestFrom(ptr unsafe.Pointer) GenerateObjectnessBasedSaliencyImageRequest {
	return GenerateObjectnessBasedSaliencyImageRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (gc _GenerateObjectnessBasedSaliencyImageRequestClass) Alloc() GenerateObjectnessBasedSaliencyImageRequest {
	rv := objc.Call[GenerateObjectnessBasedSaliencyImageRequest](gc, objc.Sel("alloc"))
	return rv
}

func GenerateObjectnessBasedSaliencyImageRequest_Alloc() GenerateObjectnessBasedSaliencyImageRequest {
	return GenerateObjectnessBasedSaliencyImageRequestClass.Alloc()
}

func (gc _GenerateObjectnessBasedSaliencyImageRequestClass) New() GenerateObjectnessBasedSaliencyImageRequest {
	rv := objc.Call[GenerateObjectnessBasedSaliencyImageRequest](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGenerateObjectnessBasedSaliencyImageRequest() GenerateObjectnessBasedSaliencyImageRequest {
	return GenerateObjectnessBasedSaliencyImageRequestClass.New()
}

func (g_ GenerateObjectnessBasedSaliencyImageRequest) Init() GenerateObjectnessBasedSaliencyImageRequest {
	rv := objc.Call[GenerateObjectnessBasedSaliencyImageRequest](g_, objc.Sel("init"))
	return rv
}

func (g_ GenerateObjectnessBasedSaliencyImageRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) GenerateObjectnessBasedSaliencyImageRequest {
	rv := objc.Call[GenerateObjectnessBasedSaliencyImageRequest](g_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewGenerateObjectnessBasedSaliencyImageRequestWithCompletionHandler(completionHandler RequestCompletionHandler) GenerateObjectnessBasedSaliencyImageRequest {
	instance := GenerateObjectnessBasedSaliencyImageRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}
