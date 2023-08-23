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

// The class instance for the [HomographicImageRegistrationRequest] class.
var HomographicImageRegistrationRequestClass = _HomographicImageRegistrationRequestClass{objc.GetClass("VNHomographicImageRegistrationRequest")}

type _HomographicImageRegistrationRequestClass struct {
	objc.Class
}

// An interface definition for the [HomographicImageRegistrationRequest] class.
type IHomographicImageRegistrationRequest interface {
	IImageRegistrationRequest
}

// An image analysis request that determines the perspective warp matrix necessary to align the content of two images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhomographicimageregistrationrequest?language=objc
type HomographicImageRegistrationRequest struct {
	ImageRegistrationRequest
}

func HomographicImageRegistrationRequestFrom(ptr unsafe.Pointer) HomographicImageRegistrationRequest {
	return HomographicImageRegistrationRequest{
		ImageRegistrationRequest: ImageRegistrationRequestFrom(ptr),
	}
}

func (hc _HomographicImageRegistrationRequestClass) Alloc() HomographicImageRegistrationRequest {
	rv := objc.Call[HomographicImageRegistrationRequest](hc, objc.Sel("alloc"))
	return rv
}

func HomographicImageRegistrationRequest_Alloc() HomographicImageRegistrationRequest {
	return HomographicImageRegistrationRequestClass.Alloc()
}

func (hc _HomographicImageRegistrationRequestClass) New() HomographicImageRegistrationRequest {
	rv := objc.Call[HomographicImageRegistrationRequest](hc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewHomographicImageRegistrationRequest() HomographicImageRegistrationRequest {
	return HomographicImageRegistrationRequestClass.New()
}

func (h_ HomographicImageRegistrationRequest) Init() HomographicImageRegistrationRequest {
	rv := objc.Call[HomographicImageRegistrationRequest](h_, objc.Sel("init"))
	return rv
}

func (h_ HomographicImageRegistrationRequest) InitWithTargetedImageURLOptions(imageURL foundation.IURL, options map[ImageOption]objc.IObject) HomographicImageRegistrationRequest {
	rv := objc.Call[HomographicImageRegistrationRequest](h_, objc.Sel("initWithTargetedImageURL:options:"), objc.Ptr(imageURL), options)
	return rv
}

// Creates a new request targeting an image at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923458-initwithtargetedimageurl?language=objc
func NewHomographicImageRegistrationRequestWithTargetedImageURLOptions(imageURL foundation.IURL, options map[ImageOption]objc.IObject) HomographicImageRegistrationRequest {
	instance := HomographicImageRegistrationRequestClass.Alloc().InitWithTargetedImageURLOptions(imageURL, options)
	instance.Autorelease()
	return instance
}

func (h_ HomographicImageRegistrationRequest) InitWithTargetedCMSampleBufferOptions(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject) HomographicImageRegistrationRequest {
	rv := objc.Call[HomographicImageRegistrationRequest](h_, objc.Sel("initWithTargetedCMSampleBuffer:options:"), sampleBuffer, options)
	return rv
}

// Creates a new request that targets an image in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/3571274-initwithtargetedcmsamplebuffer?language=objc
func NewHomographicImageRegistrationRequestWithTargetedCMSampleBufferOptions(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject) HomographicImageRegistrationRequest {
	instance := HomographicImageRegistrationRequestClass.Alloc().InitWithTargetedCMSampleBufferOptions(sampleBuffer, options)
	instance.Autorelease()
	return instance
}

func (h_ HomographicImageRegistrationRequest) InitWithTargetedCIImageOptions(ciImage coreimage.IImage, options map[ImageOption]objc.IObject) HomographicImageRegistrationRequest {
	rv := objc.Call[HomographicImageRegistrationRequest](h_, objc.Sel("initWithTargetedCIImage:options:"), objc.Ptr(ciImage), options)
	return rv
}

// Creates a new request targeting a CIImage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923447-initwithtargetedciimage?language=objc
func NewHomographicImageRegistrationRequestWithTargetedCIImageOptions(ciImage coreimage.IImage, options map[ImageOption]objc.IObject) HomographicImageRegistrationRequest {
	instance := HomographicImageRegistrationRequestClass.Alloc().InitWithTargetedCIImageOptions(ciImage, options)
	instance.Autorelease()
	return instance
}

func (h_ HomographicImageRegistrationRequest) InitWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject) HomographicImageRegistrationRequest {
	rv := objc.Call[HomographicImageRegistrationRequest](h_, objc.Sel("initWithTargetedCVPixelBuffer:options:"), pixelBuffer, options)
	return rv
}

// Creates a new request targeting an image in a CVPixelBufferRef. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923445-initwithtargetedcvpixelbuffer?language=objc
func NewHomographicImageRegistrationRequestWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject) HomographicImageRegistrationRequest {
	instance := HomographicImageRegistrationRequestClass.Alloc().InitWithTargetedCVPixelBufferOptions(pixelBuffer, options)
	instance.Autorelease()
	return instance
}

func (h_ HomographicImageRegistrationRequest) InitWithTargetedCGImageOptions(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject) HomographicImageRegistrationRequest {
	rv := objc.Call[HomographicImageRegistrationRequest](h_, objc.Sel("initWithTargetedCGImage:options:"), cgImage, options)
	return rv
}

// Creates a new request targeting a Core Graphics image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923452-initwithtargetedcgimage?language=objc
func NewHomographicImageRegistrationRequestWithTargetedCGImageOptions(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject) HomographicImageRegistrationRequest {
	instance := HomographicImageRegistrationRequestClass.Alloc().InitWithTargetedCGImageOptions(cgImage, options)
	instance.Autorelease()
	return instance
}

func (h_ HomographicImageRegistrationRequest) InitWithTargetedImageDataOptions(imageData []byte, options map[ImageOption]objc.IObject) HomographicImageRegistrationRequest {
	rv := objc.Call[HomographicImageRegistrationRequest](h_, objc.Sel("initWithTargetedImageData:options:"), imageData, options)
	return rv
}

// Creates a new request targeting an image as raw data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923460-initwithtargetedimagedata?language=objc
func NewHomographicImageRegistrationRequestWithTargetedImageDataOptions(imageData []byte, options map[ImageOption]objc.IObject) HomographicImageRegistrationRequest {
	instance := HomographicImageRegistrationRequestClass.Alloc().InitWithTargetedImageDataOptions(imageData, options)
	instance.Autorelease()
	return instance
}

func (h_ HomographicImageRegistrationRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) HomographicImageRegistrationRequest {
	rv := objc.Call[HomographicImageRegistrationRequest](h_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewHomographicImageRegistrationRequestWithCompletionHandler(completionHandler RequestCompletionHandler) HomographicImageRegistrationRequest {
	instance := HomographicImageRegistrationRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}
