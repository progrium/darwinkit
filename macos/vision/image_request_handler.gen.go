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

// The class instance for the [ImageRequestHandler] class.
var ImageRequestHandlerClass = _ImageRequestHandlerClass{objc.GetClass("VNImageRequestHandler")}

type _ImageRequestHandlerClass struct {
	objc.Class
}

// An interface definition for the [ImageRequestHandler] class.
type IImageRequestHandler interface {
	objc.IObject
	PerformRequestsError(requests []IRequest, error foundation.IError) bool
}

// An object that processes one or more image analysis requests pertaining to a single image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimagerequesthandler?language=objc
type ImageRequestHandler struct {
	objc.Object
}

func ImageRequestHandlerFrom(ptr unsafe.Pointer) ImageRequestHandler {
	return ImageRequestHandler{
		Object: objc.ObjectFrom(ptr),
	}
}

func (i_ ImageRequestHandler) InitWithCIImageOptions(image coreimage.IImage, options map[ImageOption]objc.IObject) ImageRequestHandler {
	rv := objc.Call[ImageRequestHandler](i_, objc.Sel("initWithCIImage:options:"), objc.Ptr(image), options)
	return rv
}

// Creates a handler to be used for performing requests on CIImage data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimagerequesthandler/2866549-initwithciimage?language=objc
func NewImageRequestHandlerWithCIImageOptions(image coreimage.IImage, options map[ImageOption]objc.IObject) ImageRequestHandler {
	instance := ImageRequestHandlerClass.Alloc().InitWithCIImageOptions(image, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRequestHandler) InitWithDataOptions(imageData []byte, options map[ImageOption]objc.IObject) ImageRequestHandler {
	rv := objc.Call[ImageRequestHandler](i_, objc.Sel("initWithData:options:"), imageData, options)
	return rv
}

// Creates a handler to be used for performing requests on an image contained in an NSData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimagerequesthandler/2866551-initwithdata?language=objc
func NewImageRequestHandlerWithDataOptions(imageData []byte, options map[ImageOption]objc.IObject) ImageRequestHandler {
	instance := ImageRequestHandlerClass.Alloc().InitWithDataOptions(imageData, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRequestHandler) InitWithURLOptions(imageURL foundation.IURL, options map[ImageOption]objc.IObject) ImageRequestHandler {
	rv := objc.Call[ImageRequestHandler](i_, objc.Sel("initWithURL:options:"), objc.Ptr(imageURL), options)
	return rv
}

// Creates a handler to be used for performing requests on an image at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimagerequesthandler/2866553-initwithurl?language=objc
func NewImageRequestHandlerWithURLOptions(imageURL foundation.IURL, options map[ImageOption]objc.IObject) ImageRequestHandler {
	instance := ImageRequestHandlerClass.Alloc().InitWithURLOptions(imageURL, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRequestHandler) InitWithCMSampleBufferOptions(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject) ImageRequestHandler {
	rv := objc.Call[ImageRequestHandler](i_, objc.Sel("initWithCMSampleBuffer:options:"), sampleBuffer, options)
	return rv
}

// Creates a request handler that performs requests on an image contained within a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimagerequesthandler/3548373-initwithcmsamplebuffer?language=objc
func NewImageRequestHandlerWithCMSampleBufferOptions(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject) ImageRequestHandler {
	instance := ImageRequestHandlerClass.Alloc().InitWithCMSampleBufferOptions(sampleBuffer, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRequestHandler) InitWithCGImageOptions(image coregraphics.ImageRef, options map[ImageOption]objc.IObject) ImageRequestHandler {
	rv := objc.Call[ImageRequestHandler](i_, objc.Sel("initWithCGImage:options:"), image, options)
	return rv
}

// Creates a handler to be used for performing requests on Core Graphics images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimagerequesthandler/2866541-initwithcgimage?language=objc
func NewImageRequestHandlerWithCGImageOptions(image coregraphics.ImageRef, options map[ImageOption]objc.IObject) ImageRequestHandler {
	instance := ImageRequestHandlerClass.Alloc().InitWithCGImageOptions(image, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRequestHandler) InitWithCVPixelBufferOptions(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject) ImageRequestHandler {
	rv := objc.Call[ImageRequestHandler](i_, objc.Sel("initWithCVPixelBuffer:options:"), pixelBuffer, options)
	return rv
}

// Creates a handler for performing requests on a Core Video pixel buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimagerequesthandler/2880309-initwithcvpixelbuffer?language=objc
func NewImageRequestHandlerWithCVPixelBufferOptions(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject) ImageRequestHandler {
	instance := ImageRequestHandlerClass.Alloc().InitWithCVPixelBufferOptions(pixelBuffer, options)
	instance.Autorelease()
	return instance
}

func (ic _ImageRequestHandlerClass) Alloc() ImageRequestHandler {
	rv := objc.Call[ImageRequestHandler](ic, objc.Sel("alloc"))
	return rv
}

func ImageRequestHandler_Alloc() ImageRequestHandler {
	return ImageRequestHandlerClass.Alloc()
}

func (ic _ImageRequestHandlerClass) New() ImageRequestHandler {
	rv := objc.Call[ImageRequestHandler](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageRequestHandler() ImageRequestHandler {
	return ImageRequestHandlerClass.New()
}

func (i_ ImageRequestHandler) Init() ImageRequestHandler {
	rv := objc.Call[ImageRequestHandler](i_, objc.Sel("init"))
	return rv
}

// Schedules Vision requests to perform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimagerequesthandler/2880297-performrequests?language=objc
func (i_ ImageRequestHandler) PerformRequestsError(requests []IRequest, error foundation.IError) bool {
	rv := objc.Call[bool](i_, objc.Sel("performRequests:error:"), requests, objc.Ptr(error))
	return rv
}
