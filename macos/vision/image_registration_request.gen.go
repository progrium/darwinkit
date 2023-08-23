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

// The class instance for the [ImageRegistrationRequest] class.
var ImageRegistrationRequestClass = _ImageRegistrationRequestClass{objc.GetClass("VNImageRegistrationRequest")}

type _ImageRegistrationRequestClass struct {
	objc.Class
}

// An interface definition for the [ImageRegistrationRequest] class.
type IImageRegistrationRequest interface {
	ITargetedImageRequest
}

// The abstract superclass for image analysis requests that align images according to their content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimageregistrationrequest?language=objc
type ImageRegistrationRequest struct {
	TargetedImageRequest
}

func ImageRegistrationRequestFrom(ptr unsafe.Pointer) ImageRegistrationRequest {
	return ImageRegistrationRequest{
		TargetedImageRequest: TargetedImageRequestFrom(ptr),
	}
}

func (ic _ImageRegistrationRequestClass) Alloc() ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](ic, objc.Sel("alloc"))
	return rv
}

func ImageRegistrationRequest_Alloc() ImageRegistrationRequest {
	return ImageRegistrationRequestClass.Alloc()
}

func (ic _ImageRegistrationRequestClass) New() ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageRegistrationRequest() ImageRegistrationRequest {
	return ImageRegistrationRequestClass.New()
}

func (i_ ImageRegistrationRequest) Init() ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageRegistrationRequest) InitWithTargetedImageURLOptions(imageURL foundation.IURL, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedImageURL:options:"), objc.Ptr(imageURL), options)
	return rv
}

// Creates a new request targeting an image at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923458-initwithtargetedimageurl?language=objc
func NewImageRegistrationRequestWithTargetedImageURLOptions(imageURL foundation.IURL, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedImageURLOptions(imageURL, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCMSampleBufferOptions(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCMSampleBuffer:options:"), sampleBuffer, options)
	return rv
}

// Creates a new request that targets an image in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/3571274-initwithtargetedcmsamplebuffer?language=objc
func NewImageRegistrationRequestWithTargetedCMSampleBufferOptions(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCMSampleBufferOptions(sampleBuffer, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCIImageOptions(ciImage coreimage.IImage, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCIImage:options:"), objc.Ptr(ciImage), options)
	return rv
}

// Creates a new request targeting a CIImage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923447-initwithtargetedciimage?language=objc
func NewImageRegistrationRequestWithTargetedCIImageOptions(ciImage coreimage.IImage, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCIImageOptions(ciImage, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCVPixelBuffer:options:"), pixelBuffer, options)
	return rv
}

// Creates a new request targeting an image in a CVPixelBufferRef. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923445-initwithtargetedcvpixelbuffer?language=objc
func NewImageRegistrationRequestWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCVPixelBufferOptions(pixelBuffer, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCGImageOptions(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCGImage:options:"), cgImage, options)
	return rv
}

// Creates a new request targeting a Core Graphics image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923452-initwithtargetedcgimage?language=objc
func NewImageRegistrationRequestWithTargetedCGImageOptions(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCGImageOptions(cgImage, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedImageDataOptions(imageData []byte, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedImageData:options:"), imageData, options)
	return rv
}

// Creates a new request targeting an image as raw data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923460-initwithtargetedimagedata?language=objc
func NewImageRegistrationRequestWithTargetedImageDataOptions(imageData []byte, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedImageDataOptions(imageData, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewImageRegistrationRequestWithCompletionHandler(completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}
