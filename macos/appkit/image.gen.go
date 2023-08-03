// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ImageClass = _ImageClass{objc.GetClass("NSImage")}

type _ImageClass struct {
	objc.Class
}

type IImage interface {
	objc.IObject
	SetName(string_ ImageName) bool
	Name() ImageName
	ImageWithSymbolConfiguration(configuration IImageSymbolConfiguration) Image
	AddRepresentation(imageRep IImageRep)
	AddRepresentations(imageReps []IImageRep)
	RemoveRepresentation(imageRep IImageRep)
	BestRepresentationForRectContextHints(rect foundation.Rect, referenceContext IGraphicsContext, hints map[ImageHintKey]objc.IObject) ImageRep
	DrawInRect(rect foundation.Rect)
	DrawAtPointFromRectOperationFraction(point foundation.Point, fromRect foundation.Rect, op CompositingOperation, delta float64)
	DrawInRectFromRectOperationFraction(rect foundation.Rect, fromRect foundation.Rect, op CompositingOperation, delta float64)
	DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect foundation.Rect, srcSpacePortionRect foundation.Rect, op CompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints map[ImageHintKey]objc.IObject)
	DrawRepresentationInRect(imageRep IImageRep, rect foundation.Rect) bool
	Recache()
	TIFFRepresentationUsingCompressionFactor(comp TIFFCompression, factor float32) []byte
	CGImageForProposedRectContextHints(proposedDestRect *foundation.Rect, referenceContext IGraphicsContext, hints map[ImageHintKey]objc.IObject) coregraphics.ImageRef
	CancelIncrementalLoad()
	HitTestRectWithImageDestinationRectContextHintsFlipped(testRectDestSpace foundation.Rect, imageRectDestSpace foundation.Rect, context IGraphicsContext, hints map[ImageHintKey]objc.IObject, flipped bool) bool
	LayerContentsForContentsScale(layerContentsScale float64) objc.Object
	RecommendedLayerContentsScale(preferredContentsScale float64) float64
	SymbolConfiguration() ImageSymbolConfiguration
	Delegate() ImageDelegateWrapper
	SetDelegate(value IImageDelegate)
	SetDelegate0(value objc.IObject)
	Size() foundation.Size
	SetSize(value foundation.Size)
	IsTemplate() bool
	SetTemplate(value bool)
	Representations() []ImageRep
	PrefersColorMatch() bool
	SetPrefersColorMatch(value bool)
	UsesEPSOnResolutionMismatch() bool
	SetUsesEPSOnResolutionMismatch(value bool)
	MatchesOnMultipleResolution() bool
	SetMatchesOnMultipleResolution(value bool)
	IsValid() bool
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	CapInsets() foundation.EdgeInsets
	SetCapInsets(value foundation.EdgeInsets)
	ResizingMode() ImageResizingMode
	SetResizingMode(value ImageResizingMode)
	AlignmentRect() foundation.Rect
	SetAlignmentRect(value foundation.Rect)
	CacheMode() ImageCacheMode
	SetCacheMode(value ImageCacheMode)
	TIFFRepresentation() []byte
	AccessibilityDescription() string
	SetAccessibilityDescription(value string)
	MatchesOnlyOnBestFittingAxis() bool
	SetMatchesOnlyOnBestFittingAxis(value bool)
}

type Image struct {
	objc.Object
}

func MakeImage(ptr unsafe.Pointer) Image {
	return Image{
		Object: objc.MakeObject(ptr),
	}
}

func (ic _ImageClass) ImageWithSystemSymbolNameAccessibilityDescription(symbolName string, description string) Image {
	rv := objc.CallMethod[Image](ic, objc.GetSelector("imageWithSystemSymbolName:accessibilityDescription:"), symbolName, description)
	return rv
}

func Image_ImageWithSystemSymbolNameAccessibilityDescription(symbolName string, description string) Image {
	return ImageClass.ImageWithSystemSymbolNameAccessibilityDescription(symbolName, description)
}

func (ic _ImageClass) ImageWithSystemSymbolNameVariableValueAccessibilityDescription(name string, value float64, description string) Image {
	rv := objc.CallMethod[Image](ic, objc.GetSelector("imageWithSystemSymbolName:variableValue:accessibilityDescription:"), name, value, description)
	return rv
}

func Image_ImageWithSystemSymbolNameVariableValueAccessibilityDescription(name string, value float64, description string) Image {
	return ImageClass.ImageWithSystemSymbolNameVariableValueAccessibilityDescription(name, value, description)
}

func (ic _ImageClass) ImageWithSymbolNameVariableValue(name string, value float64) Image {
	rv := objc.CallMethod[Image](ic, objc.GetSelector("imageWithSymbolName:variableValue:"), name, value)
	return rv
}

func Image_ImageWithSymbolNameVariableValue(name string, value float64) Image {
	return ImageClass.ImageWithSymbolNameVariableValue(name, value)
}

func (ic _ImageClass) ImageWithSizeFlippedDrawingHandler(size foundation.Size, drawingHandlerShouldBeCalledWithFlippedContext bool, drawingHandler func(dstRect foundation.Rect) bool) Image {
	rv := objc.CallMethod[Image](ic, objc.GetSelector("imageWithSize:flipped:drawingHandler:"), size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
	return rv
}

func Image_ImageWithSizeFlippedDrawingHandler(size foundation.Size, drawingHandlerShouldBeCalledWithFlippedContext bool, drawingHandler func(dstRect foundation.Rect) bool) Image {
	return ImageClass.ImageWithSizeFlippedDrawingHandler(size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
}

func (i_ Image) InitByReferencingFile(fileName string) Image {
	rv := objc.CallMethod[Image](i_, objc.GetSelector("initByReferencingFile:"), fileName)
	return rv
}

func Image_InitByReferencingFile(fileName string) Image {
	return ImageClass.Alloc().InitByReferencingFile(fileName)
}

func (i_ Image) InitByReferencingURL(url foundation.IURL) Image {
	rv := objc.CallMethod[Image](i_, objc.GetSelector("initByReferencingURL:"), objc.ExtractPtr(url))
	return rv
}

func Image_InitByReferencingURL(url foundation.IURL) Image {
	return ImageClass.Alloc().InitByReferencingURL(url)
}

func (i_ Image) InitWithContentsOfFile(fileName string) Image {
	rv := objc.CallMethod[Image](i_, objc.GetSelector("initWithContentsOfFile:"), fileName)
	return rv
}

func Image_InitWithContentsOfFile(fileName string) Image {
	return ImageClass.Alloc().InitWithContentsOfFile(fileName)
}

func (i_ Image) InitWithContentsOfURL(url foundation.IURL) Image {
	rv := objc.CallMethod[Image](i_, objc.GetSelector("initWithContentsOfURL:"), objc.ExtractPtr(url))
	return rv
}

func Image_InitWithContentsOfURL(url foundation.IURL) Image {
	return ImageClass.Alloc().InitWithContentsOfURL(url)
}

func (i_ Image) InitWithData(data []byte) Image {
	rv := objc.CallMethod[Image](i_, objc.GetSelector("initWithData:"), data)
	return rv
}

func Image_InitWithData(data []byte) Image {
	return ImageClass.Alloc().InitWithData(data)
}

func (i_ Image) InitWithDataIgnoringOrientation(data []byte) Image {
	rv := objc.CallMethod[Image](i_, objc.GetSelector("initWithDataIgnoringOrientation:"), data)
	return rv
}

func Image_InitWithDataIgnoringOrientation(data []byte) Image {
	return ImageClass.Alloc().InitWithDataIgnoringOrientation(data)
}

func (i_ Image) InitWithCGImageSize(cgImage coregraphics.ImageRef, size foundation.Size) Image {
	rv := objc.CallMethod[Image](i_, objc.GetSelector("initWithCGImage:size:"), cgImage, size)
	return rv
}

func Image_InitWithCGImageSize(cgImage coregraphics.ImageRef, size foundation.Size) Image {
	return ImageClass.Alloc().InitWithCGImageSize(cgImage, size)
}

func (i_ Image) InitWithPasteboard(pasteboard IPasteboard) Image {
	rv := objc.CallMethod[Image](i_, objc.GetSelector("initWithPasteboard:"), objc.ExtractPtr(pasteboard))
	return rv
}

func Image_InitWithPasteboard(pasteboard IPasteboard) Image {
	return ImageClass.Alloc().InitWithPasteboard(pasteboard)
}

func (i_ Image) InitWithSize(size foundation.Size) Image {
	rv := objc.CallMethod[Image](i_, objc.GetSelector("initWithSize:"), size)
	return rv
}

func Image_InitWithSize(size foundation.Size) Image {
	return ImageClass.Alloc().InitWithSize(size)
}

func (ic _ImageClass) Alloc() Image {
	rv := objc.CallMethod[Image](ic, objc.GetSelector("alloc"))
	return rv
}

func Image_Alloc() Image {
	return ImageClass.Alloc()
}

func (ic _ImageClass) New() Image {
	rv := objc.CallMethod[Image](ic, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewImage() Image {
	return ImageClass.New()
}

func Image_New() Image {
	return ImageClass.New()
}

func (i_ Image) Init() Image {
	rv := objc.CallMethod[Image](i_, objc.GetSelector("init"))
	return rv
}

func Image_Init() Image {
	return ImageClass.Alloc().Init()
}

func (ic _ImageClass) ImageNamed(name ImageName) Image {
	rv := objc.CallMethod[Image](ic, objc.GetSelector("imageNamed:"), name)
	return rv
}

func Image_ImageNamed(name ImageName) Image {
	return ImageClass.ImageNamed(name)
}

func (i_ Image) SetName(string_ ImageName) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("setName:"), string_)
	return rv
}

func (i_ Image) Name() ImageName {
	rv := objc.CallMethod[ImageName](i_, objc.GetSelector("name"))
	return rv
}

func (i_ Image) ImageWithSymbolConfiguration(configuration IImageSymbolConfiguration) Image {
	rv := objc.CallMethod[Image](i_, objc.GetSelector("imageWithSymbolConfiguration:"), objc.ExtractPtr(configuration))
	return rv
}

func (ic _ImageClass) CanInitWithPasteboard(pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](ic, objc.GetSelector("canInitWithPasteboard:"), objc.ExtractPtr(pasteboard))
	return rv
}

func Image_CanInitWithPasteboard(pasteboard IPasteboard) bool {
	return ImageClass.CanInitWithPasteboard(pasteboard)
}

func (i_ Image) AddRepresentation(imageRep IImageRep) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("addRepresentation:"), objc.ExtractPtr(imageRep))
}

func (i_ Image) AddRepresentations(imageReps []IImageRep) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("addRepresentations:"), imageReps)
}

func (i_ Image) RemoveRepresentation(imageRep IImageRep) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("removeRepresentation:"), objc.ExtractPtr(imageRep))
}

func (i_ Image) BestRepresentationForRectContextHints(rect foundation.Rect, referenceContext IGraphicsContext, hints map[ImageHintKey]objc.IObject) ImageRep {
	rv := objc.CallMethod[ImageRep](i_, objc.GetSelector("bestRepresentationForRect:context:hints:"), rect, objc.ExtractPtr(referenceContext), hints)
	return rv
}

func (i_ Image) DrawInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("drawInRect:"), rect)
}

func (i_ Image) DrawAtPointFromRectOperationFraction(point foundation.Point, fromRect foundation.Rect, op CompositingOperation, delta float64) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("drawAtPoint:fromRect:operation:fraction:"), point, fromRect, op, delta)
}

func (i_ Image) DrawInRectFromRectOperationFraction(rect foundation.Rect, fromRect foundation.Rect, op CompositingOperation, delta float64) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("drawInRect:fromRect:operation:fraction:"), rect, fromRect, op, delta)
}

func (i_ Image) DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect foundation.Rect, srcSpacePortionRect foundation.Rect, op CompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints map[ImageHintKey]objc.IObject) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("drawInRect:fromRect:operation:fraction:respectFlipped:hints:"), dstSpacePortionRect, srcSpacePortionRect, op, requestedAlpha, respectContextIsFlipped, hints)
}

func (i_ Image) DrawRepresentationInRect(imageRep IImageRep, rect foundation.Rect) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("drawRepresentation:inRect:"), objc.ExtractPtr(imageRep), rect)
	return rv
}

func (i_ Image) Recache() {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("recache"))
}

func (i_ Image) TIFFRepresentationUsingCompressionFactor(comp TIFFCompression, factor float32) []byte {
	rv := objc.CallMethod[[]byte](i_, objc.GetSelector("TIFFRepresentationUsingCompression:factor:"), comp, factor)
	return rv
}

func (i_ Image) CGImageForProposedRectContextHints(proposedDestRect *foundation.Rect, referenceContext IGraphicsContext, hints map[ImageHintKey]objc.IObject) coregraphics.ImageRef {
	rv := objc.CallMethod[coregraphics.ImageRef](i_, objc.GetSelector("CGImageForProposedRect:context:hints:"), proposedDestRect, objc.ExtractPtr(referenceContext), hints)
	return rv
}

func (i_ Image) CancelIncrementalLoad() {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("cancelIncrementalLoad"))
}

func (i_ Image) HitTestRectWithImageDestinationRectContextHintsFlipped(testRectDestSpace foundation.Rect, imageRectDestSpace foundation.Rect, context IGraphicsContext, hints map[ImageHintKey]objc.IObject, flipped bool) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("hitTestRect:withImageDestinationRect:context:hints:flipped:"), testRectDestSpace, imageRectDestSpace, objc.ExtractPtr(context), hints, flipped)
	return rv
}

func (i_ Image) LayerContentsForContentsScale(layerContentsScale float64) objc.Object {
	rv := objc.CallMethod[objc.Object](i_, objc.GetSelector("layerContentsForContentsScale:"), layerContentsScale)
	return rv
}

func (i_ Image) RecommendedLayerContentsScale(preferredContentsScale float64) float64 {
	rv := objc.CallMethod[float64](i_, objc.GetSelector("recommendedLayerContentsScale:"), preferredContentsScale)
	return rv
}

func (i_ Image) SymbolConfiguration() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](i_, objc.GetSelector("symbolConfiguration"))
	return rv
}

func (i_ Image) Delegate() ImageDelegateWrapper {
	rv := objc.CallMethod[ImageDelegateWrapper](i_, objc.GetSelector("delegate"))
	return rv
}

func (i_ Image) SetDelegate(value IImageDelegate) {
	po := objc.WrapAsProtocol("NSImageDelegate", value)
	objc.SetAssociatedObject(i_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setDelegate:"), po)
}

func (i_ Image) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (i_ Image) Size() foundation.Size {
	rv := objc.CallMethod[foundation.Size](i_, objc.GetSelector("size"))
	return rv
}

func (i_ Image) SetSize(value foundation.Size) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setSize:"), value)
}

func (i_ Image) IsTemplate() bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("isTemplate"))
	return rv
}

func (i_ Image) SetTemplate(value bool) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setTemplate:"), value)
}

func (ic _ImageClass) ImageTypes() []string {
	rv := objc.CallMethod[[]string](ic, objc.GetSelector("imageTypes"))
	return rv
}

func Image_ImageTypes() []string {
	return ImageClass.ImageTypes()
}

func (ic _ImageClass) ImageUnfilteredTypes() []string {
	rv := objc.CallMethod[[]string](ic, objc.GetSelector("imageUnfilteredTypes"))
	return rv
}

func Image_ImageUnfilteredTypes() []string {
	return ImageClass.ImageUnfilteredTypes()
}

func (i_ Image) Representations() []ImageRep {
	rv := objc.CallMethod[[]ImageRep](i_, objc.GetSelector("representations"))
	// TODO: convert slice items...
	return rv
}

func (i_ Image) PrefersColorMatch() bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("prefersColorMatch"))
	return rv
}

func (i_ Image) SetPrefersColorMatch(value bool) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setPrefersColorMatch:"), value)
}

func (i_ Image) UsesEPSOnResolutionMismatch() bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("usesEPSOnResolutionMismatch"))
	return rv
}

func (i_ Image) SetUsesEPSOnResolutionMismatch(value bool) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setUsesEPSOnResolutionMismatch:"), value)
}

func (i_ Image) MatchesOnMultipleResolution() bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("matchesOnMultipleResolution"))
	return rv
}

func (i_ Image) SetMatchesOnMultipleResolution(value bool) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setMatchesOnMultipleResolution:"), value)
}

func (i_ Image) IsValid() bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("isValid"))
	return rv
}

func (i_ Image) BackgroundColor() Color {
	rv := objc.CallMethod[Color](i_, objc.GetSelector("backgroundColor"))
	return rv
}

func (i_ Image) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (i_ Image) CapInsets() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](i_, objc.GetSelector("capInsets"))
	return rv
}

func (i_ Image) SetCapInsets(value foundation.EdgeInsets) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setCapInsets:"), value)
}

func (i_ Image) ResizingMode() ImageResizingMode {
	rv := objc.CallMethod[ImageResizingMode](i_, objc.GetSelector("resizingMode"))
	return rv
}

func (i_ Image) SetResizingMode(value ImageResizingMode) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setResizingMode:"), value)
}

func (i_ Image) AlignmentRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](i_, objc.GetSelector("alignmentRect"))
	return rv
}

func (i_ Image) SetAlignmentRect(value foundation.Rect) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setAlignmentRect:"), value)
}

func (i_ Image) CacheMode() ImageCacheMode {
	rv := objc.CallMethod[ImageCacheMode](i_, objc.GetSelector("cacheMode"))
	return rv
}

func (i_ Image) SetCacheMode(value ImageCacheMode) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setCacheMode:"), value)
}

func (i_ Image) TIFFRepresentation() []byte {
	rv := objc.CallMethod[[]byte](i_, objc.GetSelector("TIFFRepresentation"))
	return rv
}

func (i_ Image) AccessibilityDescription() string {
	rv := objc.CallMethod[string](i_, objc.GetSelector("accessibilityDescription"))
	return rv
}

func (i_ Image) SetAccessibilityDescription(value string) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setAccessibilityDescription:"), value)
}

func (i_ Image) MatchesOnlyOnBestFittingAxis() bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("matchesOnlyOnBestFittingAxis"))
	return rv
}

func (i_ Image) SetMatchesOnlyOnBestFittingAxis(value bool) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setMatchesOnlyOnBestFittingAxis:"), value)
}
