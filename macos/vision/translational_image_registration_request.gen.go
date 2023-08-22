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

// The class instance for the [TranslationalImageRegistrationRequest] class.
var TranslationalImageRegistrationRequestClass = _TranslationalImageRegistrationRequestClass{objc.GetClass("VNTranslationalImageRegistrationRequest")}

type _TranslationalImageRegistrationRequestClass struct {
	objc.Class
}

// An interface definition for the [TranslationalImageRegistrationRequest] class.
type ITranslationalImageRegistrationRequest interface {
	IImageRegistrationRequest
}

// An image analysis request that determines the affine transform necessary to align the content of two images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntranslationalimageregistrationrequest?language=objc
type TranslationalImageRegistrationRequest struct {
	ImageRegistrationRequest
}

func TranslationalImageRegistrationRequestFrom(ptr unsafe.Pointer) TranslationalImageRegistrationRequest {
	return TranslationalImageRegistrationRequest{
		ImageRegistrationRequest: ImageRegistrationRequestFrom(ptr),
	}
}

func (tc _TranslationalImageRegistrationRequestClass) Alloc() TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](tc, objc.Sel("alloc"))
	return rv
}

func TranslationalImageRegistrationRequest_Alloc() TranslationalImageRegistrationRequest {
	return TranslationalImageRegistrationRequestClass.Alloc()
}

func (tc _TranslationalImageRegistrationRequestClass) New() TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTranslationalImageRegistrationRequest() TranslationalImageRegistrationRequest {
	return TranslationalImageRegistrationRequestClass.New()
}

func (t_ TranslationalImageRegistrationRequest) Init() TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("init"))
	return rv
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedImageURLOptions(imageURL foundation.IURL, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedImageURL:options:"), objc.Ptr(imageURL), options)
	return rv
}

// Creates a new request targeting an image at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923458-initwithtargetedimageurl?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedImageURLOptions(imageURL foundation.IURL, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedImageURLOptions(imageURL, options)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCMSampleBufferOptions(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCMSampleBuffer:options:"), sampleBuffer, options)
	return rv
}

// Creates a new request that targets an image in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/3571274-initwithtargetedcmsamplebuffer?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCMSampleBufferOptions(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCMSampleBufferOptions(sampleBuffer, options)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCIImageOptions(ciImage coreimage.IImage, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCIImage:options:"), objc.Ptr(ciImage), options)
	return rv
}

// Creates a new request targeting a CIImage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923447-initwithtargetedciimage?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCIImageOptions(ciImage coreimage.IImage, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCIImageOptions(ciImage, options)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCVPixelBuffer:options:"), pixelBuffer, options)
	return rv
}

// Creates a new request targeting an image in a CVPixelBufferRef. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923445-initwithtargetedcvpixelbuffer?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCVPixelBufferOptions(pixelBuffer, options)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCGImageOptions(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCGImage:options:"), cgImage, options)
	return rv
}

// Creates a new request targeting a Core Graphics image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923452-initwithtargetedcgimage?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCGImageOptions(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCGImageOptions(cgImage, options)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedImageDataOptions(imageData []byte, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedImageData:options:"), imageData, options)
	return rv
}

// Creates a new request targeting an image as raw data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923460-initwithtargetedimagedata?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedImageDataOptions(imageData []byte, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedImageDataOptions(imageData, options)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewTranslationalImageRegistrationRequestWithCompletionHandler(completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}
