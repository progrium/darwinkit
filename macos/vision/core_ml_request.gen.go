// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CoreMLRequest] class.
var CoreMLRequestClass = _CoreMLRequestClass{objc.GetClass("VNCoreMLRequest")}

type _CoreMLRequestClass struct {
	objc.Class
}

// An interface definition for the [CoreMLRequest] class.
type ICoreMLRequest interface {
	IImageBasedRequest
	Model() CoreMLModel
	ImageCropAndScaleOption() ImageCropAndScaleOption
	SetImageCropAndScaleOption(value ImageCropAndScaleOption)
}

// An image analysis request that uses a Core ML model to process images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlrequest?language=objc
type CoreMLRequest struct {
	ImageBasedRequest
}

func CoreMLRequestFrom(ptr unsafe.Pointer) CoreMLRequest {
	return CoreMLRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (c_ CoreMLRequest) InitWithModel(model ICoreMLModel) CoreMLRequest {
	rv := objc.Call[CoreMLRequest](c_, objc.Sel("initWithModel:"), objc.Ptr(model))
	return rv
}

// Creates a model container to use with an image analysis request based on the model you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlrequest/2890146-initwithmodel?language=objc
func NewCoreMLRequestWithModel(model ICoreMLModel) CoreMLRequest {
	instance := CoreMLRequestClass.Alloc().InitWithModel(model)
	instance.Autorelease()
	return instance
}

func (cc _CoreMLRequestClass) Alloc() CoreMLRequest {
	rv := objc.Call[CoreMLRequest](cc, objc.Sel("alloc"))
	return rv
}

func CoreMLRequest_Alloc() CoreMLRequest {
	return CoreMLRequestClass.Alloc()
}

func (cc _CoreMLRequestClass) New() CoreMLRequest {
	rv := objc.Call[CoreMLRequest](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCoreMLRequest() CoreMLRequest {
	return CoreMLRequestClass.New()
}

func (c_ CoreMLRequest) Init() CoreMLRequest {
	rv := objc.Call[CoreMLRequest](c_, objc.Sel("init"))
	return rv
}

func (c_ CoreMLRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) CoreMLRequest {
	rv := objc.Call[CoreMLRequest](c_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewCoreMLRequestWithCompletionHandler(completionHandler RequestCompletionHandler) CoreMLRequest {
	instance := CoreMLRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}

// The model to base the image analysis request on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlrequest/2890150-model?language=objc
func (c_ CoreMLRequest) Model() CoreMLModel {
	rv := objc.Call[CoreMLModel](c_, objc.Sel("model"))
	return rv
}

// An optional setting that tells the Vision algorithm how to scale an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlrequest/2890144-imagecropandscaleoption?language=objc
func (c_ CoreMLRequest) ImageCropAndScaleOption() ImageCropAndScaleOption {
	rv := objc.Call[ImageCropAndScaleOption](c_, objc.Sel("imageCropAndScaleOption"))
	return rv
}

// An optional setting that tells the Vision algorithm how to scale an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlrequest/2890144-imagecropandscaleoption?language=objc
func (c_ CoreMLRequest) SetImageCropAndScaleOption(value ImageCropAndScaleOption) {
	objc.Call[objc.Void](c_, objc.Sel("setImageCropAndScaleOption:"), value)
}
