// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GenerateImageFeaturePrintRequest] class.
var GenerateImageFeaturePrintRequestClass = _GenerateImageFeaturePrintRequestClass{objc.GetClass("VNGenerateImageFeaturePrintRequest")}

type _GenerateImageFeaturePrintRequestClass struct {
	objc.Class
}

// An interface definition for the [GenerateImageFeaturePrintRequest] class.
type IGenerateImageFeaturePrintRequest interface {
	IImageBasedRequest
	ImageCropAndScaleOption() ImageCropAndScaleOption
	SetImageCropAndScaleOption(value ImageCropAndScaleOption)
}

// An image-based request to generate feature prints from an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateimagefeatureprintrequest?language=objc
type GenerateImageFeaturePrintRequest struct {
	ImageBasedRequest
}

func GenerateImageFeaturePrintRequestFrom(ptr unsafe.Pointer) GenerateImageFeaturePrintRequest {
	return GenerateImageFeaturePrintRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (gc _GenerateImageFeaturePrintRequestClass) Alloc() GenerateImageFeaturePrintRequest {
	rv := objc.Call[GenerateImageFeaturePrintRequest](gc, objc.Sel("alloc"))
	return rv
}

func GenerateImageFeaturePrintRequest_Alloc() GenerateImageFeaturePrintRequest {
	return GenerateImageFeaturePrintRequestClass.Alloc()
}

func (gc _GenerateImageFeaturePrintRequestClass) New() GenerateImageFeaturePrintRequest {
	rv := objc.Call[GenerateImageFeaturePrintRequest](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGenerateImageFeaturePrintRequest() GenerateImageFeaturePrintRequest {
	return GenerateImageFeaturePrintRequestClass.New()
}

func (g_ GenerateImageFeaturePrintRequest) Init() GenerateImageFeaturePrintRequest {
	rv := objc.Call[GenerateImageFeaturePrintRequest](g_, objc.Sel("init"))
	return rv
}

func (g_ GenerateImageFeaturePrintRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) GenerateImageFeaturePrintRequest {
	rv := objc.Call[GenerateImageFeaturePrintRequest](g_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewGenerateImageFeaturePrintRequestWithCompletionHandler(completionHandler RequestCompletionHandler) GenerateImageFeaturePrintRequest {
	instance := GenerateImageFeaturePrintRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}

// An optional setting that tells the algorithm how to scale an input image before generating the feature print. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateimagefeatureprintrequest/3152620-imagecropandscaleoption?language=objc
func (g_ GenerateImageFeaturePrintRequest) ImageCropAndScaleOption() ImageCropAndScaleOption {
	rv := objc.Call[ImageCropAndScaleOption](g_, objc.Sel("imageCropAndScaleOption"))
	return rv
}

// An optional setting that tells the algorithm how to scale an input image before generating the feature print. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateimagefeatureprintrequest/3152620-imagecropandscaleoption?language=objc
func (g_ GenerateImageFeaturePrintRequest) SetImageCropAndScaleOption(value ImageCropAndScaleOption) {
	objc.Call[objc.Void](g_, objc.Sel("setImageCropAndScaleOption:"), value)
}
