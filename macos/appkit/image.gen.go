// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Image] class.
var ImageClass = _ImageClass{objc.GetClass("NSImage")}

type _ImageClass struct {
	objc.Class
}

// An interface definition for the [Image] class.
type IImage interface {
	objc.IObject
	ImageWithSymbolConfiguration(configuration IImageSymbolConfiguration) Image
	DrawAtPointFromRectOperationFraction(point foundation.Point, fromRect foundation.Rect, op CompositingOperation, delta float64)
	TIFFRepresentationUsingCompressionFactor(comp TIFFCompression, factor float64) []byte
	RecommendedLayerContentsScale(preferredContentsScale float64) float64
	DrawInRectFromRectOperationFraction(rect foundation.Rect, fromRect foundation.Rect, op CompositingOperation, delta float64)
	Name() ImageName
	AddRepresentations(imageReps []IImageRep)
	LayerContentsForContentsScale(layerContentsScale float64) objc.Object
	SetName(string_ ImageName) bool
	RemoveRepresentation(imageRep IImageRep)
	HitTestRectWithImageDestinationRectContextHintsFlipped(testRectDestSpace foundation.Rect, imageRectDestSpace foundation.Rect, context IGraphicsContext, hints map[ImageHintKey]objc.IObject, flipped bool) bool
	Recache()
	CGImageForProposedRectContextHints(proposedDestRect *foundation.Rect, referenceContext IGraphicsContext, hints map[ImageHintKey]objc.IObject) coregraphics.ImageRef
	BestRepresentationForRectContextHints(rect foundation.Rect, referenceContext IGraphicsContext, hints map[ImageHintKey]objc.IObject) ImageRep
	DrawRepresentationInRect(imageRep IImageRep, rect foundation.Rect) bool
	AddRepresentation(imageRep IImageRep)
	TIFFRepresentation() []byte
	AccessibilityDescription() string
	SetAccessibilityDescription(value string)
	IsTemplate() bool
	SetTemplate(value bool)
	Representations() []ImageRep
	Delegate() ImageDelegateWrapper
	SetDelegate(value PImageDelegate)
	SetDelegateObject(valueObject objc.IObject)
	IsValid() bool
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	SymbolConfiguration() ImageSymbolConfiguration
	UsesEPSOnResolutionMismatch() bool
	SetUsesEPSOnResolutionMismatch(value bool)
	PrefersColorMatch() bool
	SetPrefersColorMatch(value bool)
	CapInsets() foundation.EdgeInsets
	SetCapInsets(value foundation.EdgeInsets)
	CacheMode() ImageCacheMode
	SetCacheMode(value ImageCacheMode)
	AlignmentRect() foundation.Rect
	SetAlignmentRect(value foundation.Rect)
	MatchesOnMultipleResolution() bool
	SetMatchesOnMultipleResolution(value bool)
	Size() foundation.Size
	SetSize(value foundation.Size)
	ResizingMode() ImageResizingMode
	SetResizingMode(value ImageResizingMode)
	MatchesOnlyOnBestFittingAxis() bool
	SetMatchesOnlyOnBestFittingAxis(value bool)
}

// A high-level interface for manipulating image data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage?language=objc
type Image struct {
	objc.Object
}

func ImageFrom(ptr unsafe.Pointer) Image {
	return Image{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ic _ImageClass) ImageWithSizeFlippedDrawingHandler(size foundation.Size, drawingHandlerShouldBeCalledWithFlippedContext bool, drawingHandler func(dstRect foundation.Rect) bool) Image {
	rv := objc.Call[Image](ic, objc.Sel("imageWithSize:flipped:drawingHandler:"), size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
	return rv
}

// Creates and returns an image object whose contents are drawn using the specified block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519860-imagewithsize?language=objc
func Image_ImageWithSizeFlippedDrawingHandler(size foundation.Size, drawingHandlerShouldBeCalledWithFlippedContext bool, drawingHandler func(dstRect foundation.Rect) bool) Image {
	return ImageClass.ImageWithSizeFlippedDrawingHandler(size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
}

func (i_ Image) InitWithPasteboard(pasteboard IPasteboard) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithPasteboard:"), objc.Ptr(pasteboard))
	return rv
}

// Initializes and returns an image object with data from the specified pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519952-initwithpasteboard?language=objc
func NewImageWithPasteboard(pasteboard IPasteboard) Image {
	instance := ImageClass.Alloc().InitWithPasteboard(pasteboard)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitWithContentsOfFile(fileName string) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithContentsOfFile:"), fileName)
	return rv
}

// Initializes and returns an image object with the contents of the specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519918-initwithcontentsoffile?language=objc
func NewImageWithContentsOfFile(fileName string) Image {
	instance := ImageClass.Alloc().InitWithContentsOfFile(fileName)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitWithSize(size foundation.Size) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithSize:"), size)
	return rv
}

// Initializes and returns an image object with the specified dimensions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1520033-initwithsize?language=objc
func NewImageWithSize(size foundation.Size) Image {
	instance := ImageClass.Alloc().InitWithSize(size)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitWithData(data []byte) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithData:"), data)
	return rv
}

// Initializes and returns an image object using the provided image data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519941-initwithdata?language=objc
func NewImageWithData(data []byte) Image {
	instance := ImageClass.Alloc().InitWithData(data)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitWithDataIgnoringOrientation(data []byte) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithDataIgnoringOrientation:"), data)
	return rv
}

// Initializes and returns an image object using the provided image data and ignoring the EXIF orientation tags. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519915-initwithdataignoringorientation?language=objc
func NewImageWithDataIgnoringOrientation(data []byte) Image {
	instance := ImageClass.Alloc().InitWithDataIgnoringOrientation(data)
	instance.Autorelease()
	return instance
}

func (ic _ImageClass) ImageWithSystemSymbolNameAccessibilityDescription(name string, description string) Image {
	rv := objc.Call[Image](ic, objc.Sel("imageWithSystemSymbolName:accessibilityDescription:"), name, description)
	return rv
}

// Creates a symbol image with the system symbol name and accessibility description you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/3622472-imagewithsystemsymbolname?language=objc
func Image_ImageWithSystemSymbolNameAccessibilityDescription(name string, description string) Image {
	return ImageClass.ImageWithSystemSymbolNameAccessibilityDescription(name, description)
}

func (i_ Image) InitWithContentsOfURL(url foundation.IURL) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithContentsOfURL:"), objc.Ptr(url))
	return rv
}

// Initializes and returns an image object with the contents of the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519907-initwithcontentsofurl?language=objc
func NewImageWithContentsOfURL(url foundation.IURL) Image {
	instance := ImageClass.Alloc().InitWithContentsOfURL(url)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitByReferencingURL(url foundation.IURL) Image {
	rv := objc.Call[Image](i_, objc.Sel("initByReferencingURL:"), objc.Ptr(url))
	return rv
}

// Initializes and returns an image object using the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519990-initbyreferencingurl?language=objc
func NewImageByReferencingURL(url foundation.IURL) Image {
	instance := ImageClass.Alloc().InitByReferencingURL(url)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitWithCGImageSize(cgImage coregraphics.ImageRef, size foundation.Size) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithCGImage:size:"), cgImage, size)
	return rv
}

// Creates a new image using the contents of the provided image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519939-initwithcgimage?language=objc
func NewImageWithCGImageSize(cgImage coregraphics.ImageRef, size foundation.Size) Image {
	instance := ImageClass.Alloc().InitWithCGImageSize(cgImage, size)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitByReferencingFile(fileName string) Image {
	rv := objc.Call[Image](i_, objc.Sel("initByReferencingFile:"), fileName)
	return rv
}

// Initializes and returns an image object using the specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519955-initbyreferencingfile?language=objc
func NewImageByReferencingFile(fileName string) Image {
	instance := ImageClass.Alloc().InitByReferencingFile(fileName)
	instance.Autorelease()
	return instance
}

func (ic _ImageClass) Alloc() Image {
	rv := objc.Call[Image](ic, objc.Sel("alloc"))
	return rv
}

func Image_Alloc() Image {
	return ImageClass.Alloc()
}

func (ic _ImageClass) New() Image {
	rv := objc.Call[Image](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImage() Image {
	return ImageClass.New()
}

func (i_ Image) Init() Image {
	rv := objc.Call[Image](i_, objc.Sel("init"))
	return rv
}

// Creates a new symbol image with the specified configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/3656508-imagewithsymbolconfiguration?language=objc
func (i_ Image) ImageWithSymbolConfiguration(configuration IImageSymbolConfiguration) Image {
	rv := objc.Call[Image](i_, objc.Sel("imageWithSymbolConfiguration:"), objc.Ptr(configuration))
	return rv
}

// Draws all or part of the image at the specified point in the current coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519981-drawatpoint?language=objc
func (i_ Image) DrawAtPointFromRectOperationFraction(point foundation.Point, fromRect foundation.Rect, op CompositingOperation, delta float64) {
	objc.Call[objc.Void](i_, objc.Sel("drawAtPoint:fromRect:operation:fraction:"), point, fromRect, op, delta)
}

// Returns a data object that contains TIFF data with the specified compression settings for all of the image representations in the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519949-tiffrepresentationusingcompressi?language=objc
func (i_ Image) TIFFRepresentationUsingCompressionFactor(comp TIFFCompression, factor float64) []byte {
	rv := objc.Call[[]byte](i_, objc.Sel("TIFFRepresentationUsingCompression:factor:"), comp, factor)
	return rv
}

// Returns the recommended layer contents scale for this image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519878-recommendedlayercontentsscale?language=objc
func (i_ Image) RecommendedLayerContentsScale(preferredContentsScale float64) float64 {
	rv := objc.Call[float64](i_, objc.Sel("recommendedLayerContentsScale:"), preferredContentsScale)
	return rv
}

// Draws all or part of the image in the specified rectangle in the current coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1520067-drawinrect?language=objc
func (i_ Image) DrawInRectFromRectOperationFraction(rect foundation.Rect, fromRect foundation.Rect, op CompositingOperation, delta float64) {
	objc.Call[objc.Void](i_, objc.Sel("drawInRect:fromRect:operation:fraction:"), rect, fromRect, op, delta)
}

// Returns the name associated with the image, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519884-name?language=objc
func (i_ Image) Name() ImageName {
	rv := objc.Call[ImageName](i_, objc.Sel("name"))
	return rv
}

// Adds an array of image representation objects to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519964-addrepresentations?language=objc
func (i_ Image) AddRepresentations(imageReps []IImageRep) {
	objc.Call[objc.Void](i_, objc.Sel("addRepresentations:"), imageReps)
}

// Returns an object that may be used as the contents of a layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519851-layercontentsforcontentsscale?language=objc
func (i_ Image) LayerContentsForContentsScale(layerContentsScale float64) objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("layerContentsForContentsScale:"), layerContentsScale)
	return rv
}

// Registers the image object under the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1520025-setname?language=objc
func (i_ Image) SetName(string_ ImageName) bool {
	rv := objc.Call[bool](i_, objc.Sel("setName:"), string_)
	return rv
}

// Returns the image object associated with the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1520015-imagenamed?language=objc
func (ic _ImageClass) ImageNamed(name ImageName) Image {
	rv := objc.Call[Image](ic, objc.Sel("imageNamed:"), name)
	return rv
}

// Returns the image object associated with the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1520015-imagenamed?language=objc
func Image_ImageNamed(name ImageName) Image {
	return ImageClass.ImageNamed(name)
}

// Tests whether the image can create an instance of itself using pasteboard data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1520039-caninitwithpasteboard?language=objc
func (ic _ImageClass) CanInitWithPasteboard(pasteboard IPasteboard) bool {
	rv := objc.Call[bool](ic, objc.Sel("canInitWithPasteboard:"), objc.Ptr(pasteboard))
	return rv
}

// Tests whether the image can create an instance of itself using pasteboard data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1520039-caninitwithpasteboard?language=objc
func Image_CanInitWithPasteboard(pasteboard IPasteboard) bool {
	return ImageClass.CanInitWithPasteboard(pasteboard)
}

// Removes and releases the specified image representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519995-removerepresentation?language=objc
func (i_ Image) RemoveRepresentation(imageRep IImageRep) {
	objc.Call[objc.Void](i_, objc.Sel("removeRepresentation:"), objc.Ptr(imageRep))
}

// Returns whether the destination rectangle would intersect a non-transparent portion of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519922-hittestrect?language=objc
func (i_ Image) HitTestRectWithImageDestinationRectContextHintsFlipped(testRectDestSpace foundation.Rect, imageRectDestSpace foundation.Rect, context IGraphicsContext, hints map[ImageHintKey]objc.IObject, flipped bool) bool {
	rv := objc.Call[bool](i_, objc.Sel("hitTestRect:withImageDestinationRect:context:hints:flipped:"), testRectDestSpace, imageRectDestSpace, objc.Ptr(context), hints, flipped)
	return rv
}

// Invalidates and frees offscreen caches of all image representations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519890-recache?language=objc
func (i_ Image) Recache() {
	objc.Call[objc.Void](i_, objc.Sel("recache"))
}

// Returns a Core Graphics image based on the contents of the current image object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519861-cgimageforproposedrect?language=objc
func (i_ Image) CGImageForProposedRectContextHints(proposedDestRect *foundation.Rect, referenceContext IGraphicsContext, hints map[ImageHintKey]objc.IObject) coregraphics.ImageRef {
	rv := objc.Call[coregraphics.ImageRef](i_, objc.Sel("CGImageForProposedRect:context:hints:"), proposedDestRect, objc.Ptr(referenceContext), hints)
	return rv
}

// Returns the best representation of the image for the specified rectangle using the provided hints. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519961-bestrepresentationforrect?language=objc
func (i_ Image) BestRepresentationForRectContextHints(rect foundation.Rect, referenceContext IGraphicsContext, hints map[ImageHintKey]objc.IObject) ImageRep {
	rv := objc.Call[ImageRep](i_, objc.Sel("bestRepresentationForRect:context:hints:"), rect, objc.Ptr(referenceContext), hints)
	return rv
}

// Draws the image using the specified image representation object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519904-drawrepresentation?language=objc
func (i_ Image) DrawRepresentationInRect(imageRep IImageRep, rect foundation.Rect) bool {
	rv := objc.Call[bool](i_, objc.Sel("drawRepresentation:inRect:"), objc.Ptr(imageRep), rect)
	return rv
}

// Adds the specified image representation object to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519911-addrepresentation?language=objc
func (i_ Image) AddRepresentation(imageRep IImageRep) {
	objc.Call[objc.Void](i_, objc.Sel("addRepresentation:"), objc.Ptr(imageRep))
}

// A data object containing TIFF data for all of the image representations in the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519841-tiffrepresentation?language=objc
func (i_ Image) TIFFRepresentation() []byte {
	rv := objc.Call[[]byte](i_, objc.Sel("TIFFRepresentation"))
	return rv
}

// The image’s accessibility description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519943-accessibilitydescription?language=objc
func (i_ Image) AccessibilityDescription() string {
	rv := objc.Call[string](i_, objc.Sel("accessibilityDescription"))
	return rv
}

// The image’s accessibility description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519943-accessibilitydescription?language=objc
func (i_ Image) SetAccessibilityDescription(value string) {
	objc.Call[objc.Void](i_, objc.Sel("setAccessibilityDescription:"), value)
}

// A Boolean value that determines whether the image represents a template image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1520017-template?language=objc
func (i_ Image) IsTemplate() bool {
	rv := objc.Call[bool](i_, objc.Sel("isTemplate"))
	return rv
}

// A Boolean value that determines whether the image represents a template image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1520017-template?language=objc
func (i_ Image) SetTemplate(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setTemplate:"), value)
}

// An array containing all of the image object’s image representations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519858-representations?language=objc
func (i_ Image) Representations() []ImageRep {
	rv := objc.Call[[]ImageRep](i_, objc.Sel("representations"))
	return rv
}

// The image’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519926-delegate?language=objc
func (i_ Image) Delegate() ImageDelegateWrapper {
	rv := objc.Call[ImageDelegateWrapper](i_, objc.Sel("delegate"))
	return rv
}

// The image’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519926-delegate?language=objc
func (i_ Image) SetDelegate(value PImageDelegate) {
	po0 := objc.WrapAsProtocol("NSImageDelegate", value)
	objc.SetAssociatedObject(i_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](i_, objc.Sel("setDelegate:"), po0)
}

// The image’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519926-delegate?language=objc
func (i_ Image) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](i_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// A Boolean value that indicates whether it is possible to draw an image representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519991-valid?language=objc
func (i_ Image) IsValid() bool {
	rv := objc.Call[bool](i_, objc.Sel("isValid"))
	return rv
}

// The background color for the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1520059-backgroundcolor?language=objc
func (i_ Image) BackgroundColor() Color {
	rv := objc.Call[Color](i_, objc.Sel("backgroundColor"))
	return rv
}

// The background color for the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1520059-backgroundcolor?language=objc
func (i_ Image) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](i_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// Returns an array of UTI strings identifying the image types supported by the registered image representation objects, either directly or through a user-installed filter service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519988-imagetypes?language=objc
func (ic _ImageClass) ImageTypes() []string {
	rv := objc.Call[[]string](ic, objc.Sel("imageTypes"))
	return rv
}

// Returns an array of UTI strings identifying the image types supported by the registered image representation objects, either directly or through a user-installed filter service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519988-imagetypes?language=objc
func Image_ImageTypes() []string {
	return ImageClass.ImageTypes()
}

// The configuration details for a symbol image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/3852559-symbolconfiguration?language=objc
func (i_ Image) SymbolConfiguration() ImageSymbolConfiguration {
	rv := objc.Call[ImageSymbolConfiguration](i_, objc.Sel("symbolConfiguration"))
	return rv
}

// A Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519868-usesepsonresolutionmismatch?language=objc
func (i_ Image) UsesEPSOnResolutionMismatch() bool {
	rv := objc.Call[bool](i_, objc.Sel("usesEPSOnResolutionMismatch"))
	return rv
}

// A Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519868-usesepsonresolutionmismatch?language=objc
func (i_ Image) SetUsesEPSOnResolutionMismatch(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setUsesEPSOnResolutionMismatch:"), value)
}

// Returns an array of UTI strings identifying the image types supported directly by the registered image representation objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519899-imageunfilteredtypes?language=objc
func (ic _ImageClass) ImageUnfilteredTypes() []string {
	rv := objc.Call[[]string](ic, objc.Sel("imageUnfilteredTypes"))
	return rv
}

// Returns an array of UTI strings identifying the image types supported directly by the registered image representation objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519899-imageunfilteredtypes?language=objc
func Image_ImageUnfilteredTypes() []string {
	return ImageClass.ImageUnfilteredTypes()
}

// A Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1520010-preferscolormatch?language=objc
func (i_ Image) PrefersColorMatch() bool {
	rv := objc.Call[bool](i_, objc.Sel("prefersColorMatch"))
	return rv
}

// A Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1520010-preferscolormatch?language=objc
func (i_ Image) SetPrefersColorMatch(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setPrefersColorMatch:"), value)
}

// The cap insets for the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1520012-capinsets?language=objc
func (i_ Image) CapInsets() foundation.EdgeInsets {
	rv := objc.Call[foundation.EdgeInsets](i_, objc.Sel("capInsets"))
	return rv
}

// The cap insets for the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1520012-capinsets?language=objc
func (i_ Image) SetCapInsets(value foundation.EdgeInsets) {
	objc.Call[objc.Void](i_, objc.Sel("setCapInsets:"), value)
}

// The image’s caching mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519850-cachemode?language=objc
func (i_ Image) CacheMode() ImageCacheMode {
	rv := objc.Call[ImageCacheMode](i_, objc.Sel("cacheMode"))
	return rv
}

// The image’s caching mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519850-cachemode?language=objc
func (i_ Image) SetCacheMode(value ImageCacheMode) {
	objc.Call[objc.Void](i_, objc.Sel("setCacheMode:"), value)
}

// A rectangle that you can use to position the image during layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519905-alignmentrect?language=objc
func (i_ Image) AlignmentRect() foundation.Rect {
	rv := objc.Call[foundation.Rect](i_, objc.Sel("alignmentRect"))
	return rv
}

// A rectangle that you can use to position the image during layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519905-alignmentrect?language=objc
func (i_ Image) SetAlignmentRect(value foundation.Rect) {
	objc.Call[objc.Void](i_, objc.Sel("setAlignmentRect:"), value)
}

// A Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519963-matchesonmultipleresolution?language=objc
func (i_ Image) MatchesOnMultipleResolution() bool {
	rv := objc.Call[bool](i_, objc.Sel("matchesOnMultipleResolution"))
	return rv
}

// A Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519963-matchesonmultipleresolution?language=objc
func (i_ Image) SetMatchesOnMultipleResolution(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setMatchesOnMultipleResolution:"), value)
}

// The size of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519987-size?language=objc
func (i_ Image) Size() foundation.Size {
	rv := objc.Call[foundation.Size](i_, objc.Sel("size"))
	return rv
}

// The size of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519987-size?language=objc
func (i_ Image) SetSize(value foundation.Size) {
	objc.Call[objc.Void](i_, objc.Sel("setSize:"), value)
}

// The resizing mode for the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1520060-resizingmode?language=objc
func (i_ Image) ResizingMode() ImageResizingMode {
	rv := objc.Call[ImageResizingMode](i_, objc.Sel("resizingMode"))
	return rv
}

// The resizing mode for the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1520060-resizingmode?language=objc
func (i_ Image) SetResizingMode(value ImageResizingMode) {
	objc.Call[objc.Void](i_, objc.Sel("setResizingMode:"), value)
}

// A Boolean value that indicates whether the image matches only on the best fitting axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519848-matchesonlyonbestfittingaxis?language=objc
func (i_ Image) MatchesOnlyOnBestFittingAxis() bool {
	rv := objc.Call[bool](i_, objc.Sel("matchesOnlyOnBestFittingAxis"))
	return rv
}

// A Boolean value that indicates whether the image matches only on the best fitting axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/1519848-matchesonlyonbestfittingaxis?language=objc
func (i_ Image) SetMatchesOnlyOnBestFittingAxis(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setMatchesOnlyOnBestFittingAxis:"), value)
}
