// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/imageio"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Context] class.
var ContextClass = _ContextClass{objc.GetClass("CIContext")}

type _ContextClass struct {
	objc.Class
}

// An interface definition for the [Context] class.
type IContext interface {
	objc.IObject
	WriteJPEGRepresentationOfImageToURLColorSpaceOptionsError(image IImage, url foundation.IURL, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject, errorPtr foundation.IError) bool
	DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteHairSemanticSegmentationGlassesMatteGainMapOrientationOptions(image IImage, disparityImage IImage, portraitEffectsMatte IImage, hairSemanticSegmentation IImage, glassesMatte IImage, gainMap IImage, orientation imageio.ImagePropertyOrientation, options foundation.Dictionary) Filter
	RenderToCVPixelBuffer(image IImage, buffer corevideo.PixelBufferRef)
	JPEGRepresentationOfImageColorSpaceOptions(image IImage, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject) []byte
	ReclaimResources()
	StartTaskToClearError(destination IRenderDestination, error foundation.IError) RenderTask
	WritePNGRepresentationOfImageToURLFormatColorSpaceOptionsError(image IImage, url foundation.IURL, format Format, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject, errorPtr foundation.IError) bool
	DepthBlurEffectFilterForImageDataOptions(data []byte, options foundation.Dictionary) Filter
	DepthBlurEffectFilterForImageURLOptions(url foundation.IURL, options foundation.Dictionary) Filter
	PNGRepresentationOfImageFormatColorSpaceOptions(image IImage, format Format, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject) []byte
	CreateCGImageFromRectFormatColorSpaceDeferred(image IImage, fromRect coregraphics.Rect, format Format, colorSpace coregraphics.ColorSpaceRef, deferred bool) coregraphics.ImageRef
	WriteHEIF10RepresentationOfImageToURLColorSpaceOptionsError(image IImage, url foundation.IURL, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject, errorPtr foundation.IError) bool
	PrepareRenderFromRectToDestinationAtPointError(image IImage, fromRect coregraphics.Rect, destination IRenderDestination, atPoint coregraphics.Point, error foundation.IError) bool
	ClearCaches()
	HEIF10RepresentationOfImageColorSpaceOptionsError(image IImage, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject, errorPtr foundation.IError) []byte
	HEIFRepresentationOfImageFormatColorSpaceOptions(image IImage, format Format, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject) []byte
	StartTaskToRenderToDestinationError(image IImage, destination IRenderDestination, error foundation.IError) RenderTask
	WriteHEIFRepresentationOfImageToURLFormatColorSpaceOptionsError(image IImage, url foundation.IURL, format Format, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject, errorPtr foundation.IError) bool
	TIFFRepresentationOfImageFormatColorSpaceOptions(image IImage, format Format, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject) []byte
	WriteTIFFRepresentationOfImageToURLFormatColorSpaceOptionsError(image IImage, url foundation.IURL, format Format, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject, errorPtr foundation.IError) bool
	WorkingColorSpace() coregraphics.ColorSpaceRef
	WorkingFormat() Format
}

// An evaluation context for rendering image processing results and performing image analysis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext?language=objc
type Context struct {
	objc.Object
}

func ContextFrom(ptr unsafe.Pointer) Context {
	return Context{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ Context) Init() Context {
	rv := objc.Call[Context](c_, objc.Sel("init"))
	return rv
}

func (cc _ContextClass) Alloc() Context {
	rv := objc.Call[Context](cc, objc.Sel("alloc"))
	return rv
}

func Context_Alloc() Context {
	return ContextClass.Alloc()
}

func (cc _ContextClass) New() Context {
	rv := objc.Call[Context](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContext() Context {
	return ContextClass.New()
}

// Renders the image and exports the resulting image data as a file in JPEG format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1642218-writejpegrepresentationofimage?language=objc
func (c_ Context) WriteJPEGRepresentationOfImageToURLColorSpaceOptionsError(image IImage, url foundation.IURL, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject, errorPtr foundation.IError) bool {
	rv := objc.Call[bool](c_, objc.Sel("writeJPEGRepresentationOfImage:toURL:colorSpace:options:error:"), objc.Ptr(image), objc.Ptr(url), colorSpace, options, objc.Ptr(errorPtr))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/3600105-depthblureffectfilterforimage?language=objc
func (c_ Context) DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteHairSemanticSegmentationGlassesMatteGainMapOrientationOptions(image IImage, disparityImage IImage, portraitEffectsMatte IImage, hairSemanticSegmentation IImage, glassesMatte IImage, gainMap IImage, orientation imageio.ImagePropertyOrientation, options foundation.Dictionary) Filter {
	rv := objc.Call[Filter](c_, objc.Sel("depthBlurEffectFilterForImage:disparityImage:portraitEffectsMatte:hairSemanticSegmentation:glassesMatte:gainMap:orientation:options:"), objc.Ptr(image), objc.Ptr(disparityImage), objc.Ptr(portraitEffectsMatte), objc.Ptr(hairSemanticSegmentation), objc.Ptr(glassesMatte), objc.Ptr(gainMap), orientation, options)
	return rv
}

// Renders an image into a pixel buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1437853-render?language=objc
func (c_ Context) RenderToCVPixelBuffer(image IImage, buffer corevideo.PixelBufferRef) {
	objc.Call[objc.Void](c_, objc.Sel("render:toCVPixelBuffer:"), objc.Ptr(image), buffer)
}

// Renders the image and exports the resulting image data in JPEG format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1642214-jpegrepresentationofimage?language=objc
func (c_ Context) JPEGRepresentationOfImageColorSpaceOptions(image IImage, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject) []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("JPEGRepresentationOfImage:colorSpace:options:"), objc.Ptr(image), colorSpace, options)
	return rv
}

// Runs the garbage collector to reclaim any resources that the context no longer requires. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1437967-reclaimresources?language=objc
func (c_ Context) ReclaimResources() {
	objc.Call[objc.Void](c_, objc.Sel("reclaimResources"))
}

// Fills the entire destination with black or clear depending on its alphaMode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/2875450-starttasktoclear?language=objc
func (c_ Context) StartTaskToClearError(destination IRenderDestination, error foundation.IError) RenderTask {
	rv := objc.Call[RenderTask](c_, objc.Sel("startTaskToClear:error:"), objc.Ptr(destination), objc.Ptr(error))
	return rv
}

// Renders the image and exports the resulting image data as a file in PNG format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/2866197-writepngrepresentationofimage?language=objc
func (c_ Context) WritePNGRepresentationOfImageToURLFormatColorSpaceOptionsError(image IImage, url foundation.IURL, format Format, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject, errorPtr foundation.IError) bool {
	rv := objc.Call[bool](c_, objc.Sel("writePNGRepresentationOfImage:toURL:format:colorSpace:options:error:"), objc.Ptr(image), objc.Ptr(url), format, colorSpace, options, objc.Ptr(errorPtr))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/3020629-depthblureffectfilterforimagedat?language=objc
func (c_ Context) DepthBlurEffectFilterForImageDataOptions(data []byte, options foundation.Dictionary) Filter {
	rv := objc.Call[Filter](c_, objc.Sel("depthBlurEffectFilterForImageData:options:"), data, options)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/3019316-depthblureffectfilterforimageurl?language=objc
func (c_ Context) DepthBlurEffectFilterForImageURLOptions(url foundation.IURL, options foundation.Dictionary) Filter {
	rv := objc.Call[Filter](c_, objc.Sel("depthBlurEffectFilterForImageURL:options:"), objc.Ptr(url), options)
	return rv
}

// Renders the image and exports the resulting image data in PNG format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/2866196-pngrepresentationofimage?language=objc
func (c_ Context) PNGRepresentationOfImageFormatColorSpaceOptions(image IImage, format Format, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject) []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("PNGRepresentationOfImage:format:colorSpace:options:"), objc.Ptr(image), format, colorSpace, options)
	return rv
}

// Creates a Quartz 2D image from a region of a Core Image image object with deferred rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1642211-createcgimage?language=objc
func (c_ Context) CreateCGImageFromRectFormatColorSpaceDeferred(image IImage, fromRect coregraphics.Rect, format Format, colorSpace coregraphics.ColorSpaceRef, deferred bool) coregraphics.ImageRef {
	rv := objc.Call[coregraphics.ImageRef](c_, objc.Sel("createCGImage:fromRect:format:colorSpace:deferred:"), objc.Ptr(image), fromRect, format, colorSpace, deferred)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/3762900-writeheif10representationofimage?language=objc
func (c_ Context) WriteHEIF10RepresentationOfImageToURLColorSpaceOptionsError(image IImage, url foundation.IURL, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject, errorPtr foundation.IError) bool {
	rv := objc.Call[bool](c_, objc.Sel("writeHEIF10RepresentationOfImage:toURL:colorSpace:options:error:"), objc.Ptr(image), objc.Ptr(url), colorSpace, options, objc.Ptr(errorPtr))
	return rv
}

// An optional call to warm up a CIContext so that subsequent calls to render with the same arguments run more efficiently. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/2875428-preparerender?language=objc
func (c_ Context) PrepareRenderFromRectToDestinationAtPointError(image IImage, fromRect coregraphics.Rect, destination IRenderDestination, atPoint coregraphics.Point, error foundation.IError) bool {
	rv := objc.Call[bool](c_, objc.Sel("prepareRender:fromRect:toDestination:atPoint:error:"), objc.Ptr(image), fromRect, objc.Ptr(destination), atPoint, objc.Ptr(error))
	return rv
}

// Frees any cached data, such as temporary images, associated with the context and runs the garbage collector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1437790-clearcaches?language=objc
func (c_ Context) ClearCaches() {
	objc.Call[objc.Void](c_, objc.Sel("clearCaches"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/3762899-heif10representationofimage?language=objc
func (c_ Context) HEIF10RepresentationOfImageColorSpaceOptionsError(image IImage, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject, errorPtr foundation.IError) []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("HEIF10RepresentationOfImage:colorSpace:options:error:"), objc.Ptr(image), colorSpace, options, objc.Ptr(errorPtr))
	return rv
}

// Renders the image and exports the resulting image data in HEIF format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/2902269-heifrepresentationofimage?language=objc
func (c_ Context) HEIFRepresentationOfImageFormatColorSpaceOptions(image IImage, format Format, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject) []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("HEIFRepresentationOfImage:format:colorSpace:options:"), objc.Ptr(image), format, colorSpace, options)
	return rv
}

// Creates a context without a specific rendering destination, using default options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1642219-context?language=objc
func (cc _ContextClass) Context() Context {
	rv := objc.Call[Context](cc, objc.Sel("context"))
	return rv
}

// Creates a context without a specific rendering destination, using default options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1642219-context?language=objc
func Context_Context() Context {
	return ContextClass.Context()
}

// Renders an image to a destination so that point (0, 0) of the image maps to point (0, 0) of the destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/2875429-starttasktorender?language=objc
func (c_ Context) StartTaskToRenderToDestinationError(image IImage, destination IRenderDestination, error foundation.IError) RenderTask {
	rv := objc.Call[RenderTask](c_, objc.Sel("startTaskToRender:toDestination:error:"), objc.Ptr(image), objc.Ptr(destination), objc.Ptr(error))
	return rv
}

// Renders the image and exports the resulting image data as a file in HEIF format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/2902266-writeheifrepresentationofimage?language=objc
func (c_ Context) WriteHEIFRepresentationOfImageToURLFormatColorSpaceOptionsError(image IImage, url foundation.IURL, format Format, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject, errorPtr foundation.IError) bool {
	rv := objc.Call[bool](c_, objc.Sel("writeHEIFRepresentationOfImage:toURL:format:colorSpace:options:error:"), objc.Ptr(image), objc.Ptr(url), format, colorSpace, options, objc.Ptr(errorPtr))
	return rv
}

// Renders the image and exports the resulting image data in TIFF format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1642220-tiffrepresentationofimage?language=objc
func (c_ Context) TIFFRepresentationOfImageFormatColorSpaceOptions(image IImage, format Format, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject) []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("TIFFRepresentationOfImage:format:colorSpace:options:"), objc.Ptr(image), format, colorSpace, options)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/3365984-contextwithmtlcommandqueue?language=objc
func (cc _ContextClass) ContextWithMTLCommandQueue(commandQueue metal.PCommandQueue) Context {
	po0 := objc.WrapAsProtocol("MTLCommandQueue", commandQueue)
	rv := objc.Call[Context](cc, objc.Sel("contextWithMTLCommandQueue:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/3365984-contextwithmtlcommandqueue?language=objc
func Context_ContextWithMTLCommandQueue(commandQueue metal.PCommandQueue) Context {
	return ContextClass.ContextWithMTLCommandQueue(commandQueue)
}

// Returns the number of GPUs not currently driving a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1437817-offlinegpucount?language=objc
func (cc _ContextClass) OfflineGPUCount() int {
	rv := objc.Call[int](cc, objc.Sel("offlineGPUCount"))
	return rv
}

// Returns the number of GPUs not currently driving a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1437817-offlinegpucount?language=objc
func Context_OfflineGPUCount() int {
	return ContextClass.OfflineGPUCount()
}

// Renders the image and exports the resulting image data as a file in TIFF format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1642213-writetiffrepresentationofimage?language=objc
func (c_ Context) WriteTIFFRepresentationOfImageToURLFormatColorSpaceOptionsError(image IImage, url foundation.IURL, format Format, colorSpace coregraphics.ColorSpaceRef, options map[ImageRepresentationOption]objc.IObject, errorPtr foundation.IError) bool {
	rv := objc.Call[bool](c_, objc.Sel("writeTIFFRepresentationOfImage:toURL:format:colorSpace:options:error:"), objc.Ptr(image), objc.Ptr(url), format, colorSpace, options, objc.Ptr(errorPtr))
	return rv
}

// Creates a Core Image context from a Quartz context, using the specified options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1437864-contextwithcgcontext?language=objc
func (cc _ContextClass) ContextWithCGContextOptions(cgctx coregraphics.ContextRef, options map[ContextOption]objc.IObject) Context {
	rv := objc.Call[Context](cc, objc.Sel("contextWithCGContext:options:"), cgctx, options)
	return rv
}

// Creates a Core Image context from a Quartz context, using the specified options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1437864-contextwithcgcontext?language=objc
func Context_ContextWithCGContextOptions(cgctx coregraphics.ContextRef, options map[ContextOption]objc.IObject) Context {
	return ContextClass.ContextWithCGContextOptions(cgctx, options)
}

// Initializes a context without a specific rendering destination, using the specified options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1438261-initwithoptions?language=objc
func (cc _ContextClass) ContextWithOptions(options map[ContextOption]objc.IObject) Context {
	rv := objc.Call[Context](cc, objc.Sel("contextWithOptions:"), options)
	return rv
}

// Initializes a context without a specific rendering destination, using the specified options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1438261-initwithoptions?language=objc
func Context_ContextWithOptions(options map[ContextOption]objc.IObject) Context {
	return ContextClass.ContextWithOptions(options)
}

// Creates a Core Image context using the specified Metal device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1437609-contextwithmtldevice?language=objc
func (cc _ContextClass) ContextWithMTLDevice(device metal.PDevice) Context {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[Context](cc, objc.Sel("contextWithMTLDevice:"), po0)
	return rv
}

// Creates a Core Image context using the specified Metal device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1437609-contextwithmtldevice?language=objc
func Context_ContextWithMTLDevice(device metal.PDevice) Context {
	return ContextClass.ContextWithMTLDevice(device)
}

// The working color space of the Core Image context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1438061-workingcolorspace?language=objc
func (c_ Context) WorkingColorSpace() coregraphics.ColorSpaceRef {
	rv := objc.Call[coregraphics.ColorSpaceRef](c_, objc.Sel("workingColorSpace"))
	return rv
}

// The working pixel format of the Core Image context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontext/1642215-workingformat?language=objc
func (c_ Context) WorkingFormat() Format {
	rv := objc.Call[Format](c_, objc.Sel("workingFormat"))
	return rv
}
