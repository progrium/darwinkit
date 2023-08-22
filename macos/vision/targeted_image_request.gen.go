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

// The class instance for the [TargetedImageRequest] class.
var TargetedImageRequestClass = _TargetedImageRequestClass{objc.GetClass("VNTargetedImageRequest")}

type _TargetedImageRequestClass struct {
	objc.Class
}

// An interface definition for the [TargetedImageRequest] class.
type ITargetedImageRequest interface {
	IImageBasedRequest
}

// The abstract superclass for image analysis requests that operate on both the processed image and a secondary image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest?language=objc
type TargetedImageRequest struct {
	ImageBasedRequest
}

func TargetedImageRequestFrom(ptr unsafe.Pointer) TargetedImageRequest {
	return TargetedImageRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (t_ TargetedImageRequest) InitWithTargetedImageURLOptions(imageURL foundation.IURL, options map[ImageOption]objc.IObject) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedImageURL:options:"), objc.Ptr(imageURL), options)
	return rv
}

// Creates a new request targeting an image at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923458-initwithtargetedimageurl?language=objc
func NewTargetedImageRequestWithTargetedImageURLOptions(imageURL foundation.IURL, options map[ImageOption]objc.IObject) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedImageURLOptions(imageURL, options)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCMSampleBufferOptions(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCMSampleBuffer:options:"), sampleBuffer, options)
	return rv
}

// Creates a new request that targets an image in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/3571274-initwithtargetedcmsamplebuffer?language=objc
func NewTargetedImageRequestWithTargetedCMSampleBufferOptions(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCMSampleBufferOptions(sampleBuffer, options)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCIImageOptions(ciImage coreimage.IImage, options map[ImageOption]objc.IObject) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCIImage:options:"), objc.Ptr(ciImage), options)
	return rv
}

// Creates a new request targeting a CIImage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923447-initwithtargetedciimage?language=objc
func NewTargetedImageRequestWithTargetedCIImageOptions(ciImage coreimage.IImage, options map[ImageOption]objc.IObject) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCIImageOptions(ciImage, options)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCVPixelBuffer:options:"), pixelBuffer, options)
	return rv
}

// Creates a new request targeting an image in a CVPixelBufferRef. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923445-initwithtargetedcvpixelbuffer?language=objc
func NewTargetedImageRequestWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCVPixelBufferOptions(pixelBuffer, options)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCGImageOptions(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCGImage:options:"), cgImage, options)
	return rv
}

// Creates a new request targeting a Core Graphics image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923452-initwithtargetedcgimage?language=objc
func NewTargetedImageRequestWithTargetedCGImageOptions(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCGImageOptions(cgImage, options)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedImageDataOptions(imageData []byte, options map[ImageOption]objc.IObject) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedImageData:options:"), imageData, options)
	return rv
}

// Creates a new request targeting an image as raw data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923460-initwithtargetedimagedata?language=objc
func NewTargetedImageRequestWithTargetedImageDataOptions(imageData []byte, options map[ImageOption]objc.IObject) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedImageDataOptions(imageData, options)
	instance.Autorelease()
	return instance
}

func (tc _TargetedImageRequestClass) Alloc() TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](tc, objc.Sel("alloc"))
	return rv
}

func TargetedImageRequest_Alloc() TargetedImageRequest {
	return TargetedImageRequestClass.Alloc()
}

func (tc _TargetedImageRequestClass) New() TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTargetedImageRequest() TargetedImageRequest {
	return TargetedImageRequestClass.New()
}

func (t_ TargetedImageRequest) Init() TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("init"))
	return rv
}

func (t_ TargetedImageRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewTargetedImageRequestWithCompletionHandler(completionHandler RequestCompletionHandler) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}
