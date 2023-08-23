// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coreimage"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GenerateOpticalFlowRequest] class.
var GenerateOpticalFlowRequestClass = _GenerateOpticalFlowRequestClass{objc.GetClass("VNGenerateOpticalFlowRequest")}

type _GenerateOpticalFlowRequestClass struct {
	objc.Class
}

// An interface definition for the [GenerateOpticalFlowRequest] class.
type IGenerateOpticalFlowRequest interface {
	ITargetedImageRequest
	ComputationAccuracy() GenerateOpticalFlowRequestComputationAccuracy
	SetComputationAccuracy(value GenerateOpticalFlowRequestComputationAccuracy)
	OutputPixelFormat() uint
	SetOutputPixelFormat(value uint)
}

// An object that generates directional change vectors for each pixel in the targeted image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequest?language=objc
type GenerateOpticalFlowRequest struct {
	TargetedImageRequest
}

func GenerateOpticalFlowRequestFrom(ptr unsafe.Pointer) GenerateOpticalFlowRequest {
	return GenerateOpticalFlowRequest{
		TargetedImageRequest: TargetedImageRequestFrom(ptr),
	}
}

func (gc _GenerateOpticalFlowRequestClass) Alloc() GenerateOpticalFlowRequest {
	rv := objc.Call[GenerateOpticalFlowRequest](gc, objc.Sel("alloc"))
	return rv
}

func GenerateOpticalFlowRequest_Alloc() GenerateOpticalFlowRequest {
	return GenerateOpticalFlowRequestClass.Alloc()
}

func (gc _GenerateOpticalFlowRequestClass) New() GenerateOpticalFlowRequest {
	rv := objc.Call[GenerateOpticalFlowRequest](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGenerateOpticalFlowRequest() GenerateOpticalFlowRequest {
	return GenerateOpticalFlowRequestClass.New()
}

func (g_ GenerateOpticalFlowRequest) Init() GenerateOpticalFlowRequest {
	rv := objc.Call[GenerateOpticalFlowRequest](g_, objc.Sel("init"))
	return rv
}

func (g_ GenerateOpticalFlowRequest) InitWithTargetedImageURLOptions(imageURL foundation.IURL, options map[ImageOption]objc.IObject) GenerateOpticalFlowRequest {
	rv := objc.Call[GenerateOpticalFlowRequest](g_, objc.Sel("initWithTargetedImageURL:options:"), objc.Ptr(imageURL), options)
	return rv
}

// Creates a new request targeting an image at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923458-initwithtargetedimageurl?language=objc
func NewGenerateOpticalFlowRequestWithTargetedImageURLOptions(imageURL foundation.IURL, options map[ImageOption]objc.IObject) GenerateOpticalFlowRequest {
	instance := GenerateOpticalFlowRequestClass.Alloc().InitWithTargetedImageURLOptions(imageURL, options)
	instance.Autorelease()
	return instance
}

func (g_ GenerateOpticalFlowRequest) InitWithTargetedCMSampleBufferOptions(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject) GenerateOpticalFlowRequest {
	rv := objc.Call[GenerateOpticalFlowRequest](g_, objc.Sel("initWithTargetedCMSampleBuffer:options:"), sampleBuffer, options)
	return rv
}

// Creates a new request that targets an image in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/3571274-initwithtargetedcmsamplebuffer?language=objc
func NewGenerateOpticalFlowRequestWithTargetedCMSampleBufferOptions(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject) GenerateOpticalFlowRequest {
	instance := GenerateOpticalFlowRequestClass.Alloc().InitWithTargetedCMSampleBufferOptions(sampleBuffer, options)
	instance.Autorelease()
	return instance
}

func (g_ GenerateOpticalFlowRequest) InitWithTargetedCIImageOptions(ciImage coreimage.IImage, options map[ImageOption]objc.IObject) GenerateOpticalFlowRequest {
	rv := objc.Call[GenerateOpticalFlowRequest](g_, objc.Sel("initWithTargetedCIImage:options:"), objc.Ptr(ciImage), options)
	return rv
}

// Creates a new request targeting a CIImage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923447-initwithtargetedciimage?language=objc
func NewGenerateOpticalFlowRequestWithTargetedCIImageOptions(ciImage coreimage.IImage, options map[ImageOption]objc.IObject) GenerateOpticalFlowRequest {
	instance := GenerateOpticalFlowRequestClass.Alloc().InitWithTargetedCIImageOptions(ciImage, options)
	instance.Autorelease()
	return instance
}

func (g_ GenerateOpticalFlowRequest) InitWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject) GenerateOpticalFlowRequest {
	rv := objc.Call[GenerateOpticalFlowRequest](g_, objc.Sel("initWithTargetedCVPixelBuffer:options:"), pixelBuffer, options)
	return rv
}

// Creates a new request targeting an image in a CVPixelBufferRef. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923445-initwithtargetedcvpixelbuffer?language=objc
func NewGenerateOpticalFlowRequestWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject) GenerateOpticalFlowRequest {
	instance := GenerateOpticalFlowRequestClass.Alloc().InitWithTargetedCVPixelBufferOptions(pixelBuffer, options)
	instance.Autorelease()
	return instance
}

func (g_ GenerateOpticalFlowRequest) InitWithTargetedCGImageOptions(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject) GenerateOpticalFlowRequest {
	rv := objc.Call[GenerateOpticalFlowRequest](g_, objc.Sel("initWithTargetedCGImage:options:"), cgImage, options)
	return rv
}

// Creates a new request targeting a Core Graphics image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923452-initwithtargetedcgimage?language=objc
func NewGenerateOpticalFlowRequestWithTargetedCGImageOptions(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject) GenerateOpticalFlowRequest {
	instance := GenerateOpticalFlowRequestClass.Alloc().InitWithTargetedCGImageOptions(cgImage, options)
	instance.Autorelease()
	return instance
}

func (g_ GenerateOpticalFlowRequest) InitWithTargetedImageDataOptions(imageData []byte, options map[ImageOption]objc.IObject) GenerateOpticalFlowRequest {
	rv := objc.Call[GenerateOpticalFlowRequest](g_, objc.Sel("initWithTargetedImageData:options:"), imageData, options)
	return rv
}

// Creates a new request targeting an image as raw data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923460-initwithtargetedimagedata?language=objc
func NewGenerateOpticalFlowRequestWithTargetedImageDataOptions(imageData []byte, options map[ImageOption]objc.IObject) GenerateOpticalFlowRequest {
	instance := GenerateOpticalFlowRequestClass.Alloc().InitWithTargetedImageDataOptions(imageData, options)
	instance.Autorelease()
	return instance
}

func (g_ GenerateOpticalFlowRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) GenerateOpticalFlowRequest {
	rv := objc.Call[GenerateOpticalFlowRequest](g_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewGenerateOpticalFlowRequestWithCompletionHandler(completionHandler RequestCompletionHandler) GenerateOpticalFlowRequest {
	instance := GenerateOpticalFlowRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}

// The accuracy level for computing optical flow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequest/3672181-computationaccuracy?language=objc
func (g_ GenerateOpticalFlowRequest) ComputationAccuracy() GenerateOpticalFlowRequestComputationAccuracy {
	rv := objc.Call[GenerateOpticalFlowRequestComputationAccuracy](g_, objc.Sel("computationAccuracy"))
	return rv
}

// The accuracy level for computing optical flow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequest/3672181-computationaccuracy?language=objc
func (g_ GenerateOpticalFlowRequest) SetComputationAccuracy(value GenerateOpticalFlowRequestComputationAccuracy) {
	objc.Call[objc.Void](g_, objc.Sel("setComputationAccuracy:"), value)
}

// The output buffer’s pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequest/3548304-outputpixelformat?language=objc
func (g_ GenerateOpticalFlowRequest) OutputPixelFormat() uint {
	rv := objc.Call[uint](g_, objc.Sel("outputPixelFormat"))
	return rv
}

// The output buffer’s pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequest/3548304-outputpixelformat?language=objc
func (g_ GenerateOpticalFlowRequest) SetOutputPixelFormat(value uint) {
	objc.Call[objc.Void](g_, objc.Sel("setOutputPixelFormat:"), value)
}
