// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ImageRepClass = _ImageRepClass{objc.GetClass("NSImageRep")}

type _ImageRepClass struct {
	objc.Class
}

type IImageRep interface {
	objc.IObject
	CGImageForProposedRectContextHints(proposedDestRect *foundation.Rect, context IGraphicsContext, hints map[ImageHintKey]objc.IObject) coregraphics.ImageRef
	Draw() bool
	DrawAtPoint(point foundation.Point) bool
	DrawInRect(rect foundation.Rect) bool
	DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect foundation.Rect, srcSpacePortionRect foundation.Rect, op CompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints map[ImageHintKey]objc.IObject) bool
	Size() foundation.Size
	SetSize(value foundation.Size)
	BitsPerSample() int
	SetBitsPerSample(value int)
	ColorSpaceName() ColorSpaceName
	SetColorSpaceName(value ColorSpaceName)
	HasAlpha() bool
	SetAlpha(value bool)
	IsOpaque() bool
	SetOpaque(value bool)
	PixelsHigh() int
	SetPixelsHigh(value int)
	PixelsWide() int
	SetPixelsWide(value int)
	LayoutDirection() ImageLayoutDirection
	SetLayoutDirection(value ImageLayoutDirection)
}

type ImageRep struct {
	objc.Object
}

func MakeImageRep(ptr unsafe.Pointer) ImageRep {
	return ImageRep{
		Object: objc.MakeObject(ptr),
	}
}

func (i_ ImageRep) Init() ImageRep {
	rv := objc.CallMethod[ImageRep](i_, objc.GetSelector("init"))
	return rv
}

func ImageRep_Init() ImageRep {
	return ImageRepClass.Alloc().Init()
}

func (ic _ImageRepClass) Alloc() ImageRep {
	rv := objc.CallMethod[ImageRep](ic, objc.GetSelector("alloc"))
	return rv
}

func ImageRep_Alloc() ImageRep {
	return ImageRepClass.Alloc()
}

func (ic _ImageRepClass) New() ImageRep {
	rv := objc.CallMethod[ImageRep](ic, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewImageRep() ImageRep {
	return ImageRepClass.New()
}

func ImageRep_New() ImageRep {
	return ImageRepClass.New()
}

func (ic _ImageRepClass) ImageRepsWithContentsOfFile(filename string) []ImageRep {
	rv := objc.CallMethod[[]ImageRep](ic, objc.GetSelector("imageRepsWithContentsOfFile:"), filename)
	return rv
}

func ImageRep_ImageRepsWithContentsOfFile(filename string) []ImageRep {
	return ImageRepClass.ImageRepsWithContentsOfFile(filename)
}

func (ic _ImageRepClass) ImageRepsWithPasteboard(pasteboard IPasteboard) []ImageRep {
	rv := objc.CallMethod[[]ImageRep](ic, objc.GetSelector("imageRepsWithPasteboard:"), objc.ExtractPtr(pasteboard))
	return rv
}

func ImageRep_ImageRepsWithPasteboard(pasteboard IPasteboard) []ImageRep {
	return ImageRepClass.ImageRepsWithPasteboard(pasteboard)
}

func (ic _ImageRepClass) ImageRepsWithContentsOfURL(url foundation.IURL) []ImageRep {
	rv := objc.CallMethod[[]ImageRep](ic, objc.GetSelector("imageRepsWithContentsOfURL:"), objc.ExtractPtr(url))
	return rv
}

func ImageRep_ImageRepsWithContentsOfURL(url foundation.IURL) []ImageRep {
	return ImageRepClass.ImageRepsWithContentsOfURL(url)
}

func (ic _ImageRepClass) ImageRepWithContentsOfFile(filename string) ImageRep {
	rv := objc.CallMethod[ImageRep](ic, objc.GetSelector("imageRepWithContentsOfFile:"), filename)
	return rv
}

func ImageRep_ImageRepWithContentsOfFile(filename string) ImageRep {
	return ImageRepClass.ImageRepWithContentsOfFile(filename)
}

func (ic _ImageRepClass) ImageRepWithPasteboard(pasteboard IPasteboard) ImageRep {
	rv := objc.CallMethod[ImageRep](ic, objc.GetSelector("imageRepWithPasteboard:"), objc.ExtractPtr(pasteboard))
	return rv
}

func ImageRep_ImageRepWithPasteboard(pasteboard IPasteboard) ImageRep {
	return ImageRepClass.ImageRepWithPasteboard(pasteboard)
}

func (ic _ImageRepClass) ImageRepWithContentsOfURL(url foundation.IURL) ImageRep {
	rv := objc.CallMethod[ImageRep](ic, objc.GetSelector("imageRepWithContentsOfURL:"), objc.ExtractPtr(url))
	return rv
}

func ImageRep_ImageRepWithContentsOfURL(url foundation.IURL) ImageRep {
	return ImageRepClass.ImageRepWithContentsOfURL(url)
}

func (ic _ImageRepClass) CanInitWithData(data []byte) bool {
	rv := objc.CallMethod[bool](ic, objc.GetSelector("canInitWithData:"), data)
	return rv
}

func ImageRep_CanInitWithData(data []byte) bool {
	return ImageRepClass.CanInitWithData(data)
}

func (ic _ImageRepClass) CanInitWithPasteboard(pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](ic, objc.GetSelector("canInitWithPasteboard:"), objc.ExtractPtr(pasteboard))
	return rv
}

func ImageRep_CanInitWithPasteboard(pasteboard IPasteboard) bool {
	return ImageRepClass.CanInitWithPasteboard(pasteboard)
}

func (i_ ImageRep) CGImageForProposedRectContextHints(proposedDestRect *foundation.Rect, context IGraphicsContext, hints map[ImageHintKey]objc.IObject) coregraphics.ImageRef {
	rv := objc.CallMethod[coregraphics.ImageRef](i_, objc.GetSelector("CGImageForProposedRect:context:hints:"), proposedDestRect, objc.ExtractPtr(context), hints)
	return rv
}

func (i_ ImageRep) Draw() bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("draw"))
	return rv
}

func (i_ ImageRep) DrawAtPoint(point foundation.Point) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("drawAtPoint:"), point)
	return rv
}

func (i_ ImageRep) DrawInRect(rect foundation.Rect) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("drawInRect:"), rect)
	return rv
}

func (i_ ImageRep) DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect foundation.Rect, srcSpacePortionRect foundation.Rect, op CompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints map[ImageHintKey]objc.IObject) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("drawInRect:fromRect:operation:fraction:respectFlipped:hints:"), dstSpacePortionRect, srcSpacePortionRect, op, requestedAlpha, respectContextIsFlipped, hints)
	return rv
}

func (ic _ImageRepClass) ImageRepClassForType(type_ string) objc.Class {
	rv := objc.CallMethod[objc.Class](ic, objc.GetSelector("imageRepClassForType:"), type_)
	return rv
}

func ImageRep_ImageRepClassForType(type_ string) objc.Class {
	return ImageRepClass.ImageRepClassForType(type_)
}

func (ic _ImageRepClass) ImageRepClassForData(data []byte) objc.Class {
	rv := objc.CallMethod[objc.Class](ic, objc.GetSelector("imageRepClassForData:"), data)
	return rv
}

func ImageRep_ImageRepClassForData(data []byte) objc.Class {
	return ImageRepClass.ImageRepClassForData(data)
}

func (ic _ImageRepClass) RegisterImageRepClass(imageRepClass objc.IClass) {
	objc.CallMethod[objc.Void](ic, objc.GetSelector("registerImageRepClass:"), objc.ExtractPtr(imageRepClass))
}

func ImageRep_RegisterImageRepClass(imageRepClass objc.IClass) {
	ImageRepClass.RegisterImageRepClass(imageRepClass)
}

func (ic _ImageRepClass) UnregisterImageRepClass(imageRepClass objc.IClass) {
	objc.CallMethod[objc.Void](ic, objc.GetSelector("unregisterImageRepClass:"), objc.ExtractPtr(imageRepClass))
}

func ImageRep_UnregisterImageRepClass(imageRepClass objc.IClass) {
	ImageRepClass.UnregisterImageRepClass(imageRepClass)
}

func (ic _ImageRepClass) ImageTypes() []string {
	rv := objc.CallMethod[[]string](ic, objc.GetSelector("imageTypes"))
	return rv
}

func ImageRep_ImageTypes() []string {
	return ImageRepClass.ImageTypes()
}

func (ic _ImageRepClass) ImageUnfilteredTypes() []string {
	rv := objc.CallMethod[[]string](ic, objc.GetSelector("imageUnfilteredTypes"))
	return rv
}

func ImageRep_ImageUnfilteredTypes() []string {
	return ImageRepClass.ImageUnfilteredTypes()
}

func (i_ ImageRep) Size() foundation.Size {
	rv := objc.CallMethod[foundation.Size](i_, objc.GetSelector("size"))
	return rv
}

func (i_ ImageRep) SetSize(value foundation.Size) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setSize:"), value)
}

func (i_ ImageRep) BitsPerSample() int {
	rv := objc.CallMethod[int](i_, objc.GetSelector("bitsPerSample"))
	return rv
}

func (i_ ImageRep) SetBitsPerSample(value int) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setBitsPerSample:"), value)
}

func (i_ ImageRep) ColorSpaceName() ColorSpaceName {
	rv := objc.CallMethod[ColorSpaceName](i_, objc.GetSelector("colorSpaceName"))
	return rv
}

func (i_ ImageRep) SetColorSpaceName(value ColorSpaceName) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setColorSpaceName:"), value)
}

func (i_ ImageRep) HasAlpha() bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("hasAlpha"))
	return rv
}

func (i_ ImageRep) SetAlpha(value bool) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setAlpha:"), value)
}

func (i_ ImageRep) IsOpaque() bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("isOpaque"))
	return rv
}

func (i_ ImageRep) SetOpaque(value bool) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setOpaque:"), value)
}

func (i_ ImageRep) PixelsHigh() int {
	rv := objc.CallMethod[int](i_, objc.GetSelector("pixelsHigh"))
	return rv
}

func (i_ ImageRep) SetPixelsHigh(value int) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setPixelsHigh:"), value)
}

func (i_ ImageRep) PixelsWide() int {
	rv := objc.CallMethod[int](i_, objc.GetSelector("pixelsWide"))
	return rv
}

func (i_ ImageRep) SetPixelsWide(value int) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setPixelsWide:"), value)
}

func (i_ ImageRep) LayoutDirection() ImageLayoutDirection {
	rv := objc.CallMethod[ImageLayoutDirection](i_, objc.GetSelector("layoutDirection"))
	return rv
}

func (i_ ImageRep) SetLayoutDirection(value ImageLayoutDirection) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setLayoutDirection:"), value)
}

func (ic _ImageRepClass) RegisteredImageRepClasses() []objc.Class {
	rv := objc.CallMethod[[]objc.Class](ic, objc.GetSelector("registeredImageRepClasses"))
	return rv
}

func ImageRep_RegisteredImageRepClasses() []objc.Class {
	return ImageRepClass.RegisteredImageRepClasses()
}
