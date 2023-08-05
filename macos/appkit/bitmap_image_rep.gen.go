// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var BitmapImageRepClass = _BitmapImageRepClass{objc.GetClass("NSBitmapImageRep")}

type _BitmapImageRepClass struct {
	objc.Class
}

type IBitmapImageRep interface {
	IImageRep
	ColorizeByMappingGrayToColorBlackMappingWhiteMapping(midPoint float64, midPointColor IColor, shadowColor IColor, lightColor IColor)
	TIFFRepresentationUsingCompressionFactor(comp TIFFCompression, factor float32) []byte
	RepresentationUsingTypeProperties(storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.IObject) []byte
	CanBeCompressedUsing(compression TIFFCompression) bool
	SetCompressionFactor(compression TIFFCompression, factor float32)
	GetCompressionFactor(compression *TIFFCompression, factor *float32)
	SetPropertyWithValue(property BitmapImageRepPropertyKey, value objc.IObject)
	ValueForProperty(property BitmapImageRepPropertyKey) objc.Object
	IncrementalLoadFromDataComplete(data []byte, complete bool) int
	SetColorAtXY(color IColor, x int, y int)
	ColorAtXY(x int, y int) Color
	SetPixelAtXY(p *uint, x int, y int)
	GetPixelAtXY(p *uint, x int, y int)
	BitmapImageRepByConvertingToColorSpaceRenderingIntent(targetSpace IColorSpace, renderingIntent ColorRenderingIntent) BitmapImageRep
	BitmapImageRepByRetaggingWithColorSpace(newSpace IColorSpace) BitmapImageRep
	BitmapFormat() BitmapFormat
	BitsPerPixel() int
	BytesPerPlane() int
	BytesPerRow() int
	IsPlanar() bool
	NumberOfPlanes() int
	SamplesPerPixel() int
	BitmapData() *byte
	TIFFRepresentation() []byte
	CGImage() coregraphics.ImageRef
	ColorSpace() ColorSpace
}

type BitmapImageRep struct {
	ImageRep
}

func MakeBitmapImageRep(ptr unsafe.Pointer) BitmapImageRep {
	return BitmapImageRep{
		ImageRep: MakeImageRep(ptr),
	}
}

func (bc _BitmapImageRepClass) ImageRepWithData(data []byte) BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](bc, objc.GetSelector("imageRepWithData:"), data)
	return rv
}

func BitmapImageRep_ImageRepWithData(data []byte) BitmapImageRep {
	return BitmapImageRepClass.ImageRepWithData(data)
}

func (b_ BitmapImageRep) InitWithCGImage(cgImage coregraphics.ImageRef) BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](b_, objc.GetSelector("initWithCGImage:"), cgImage)
	return rv
}

func BitmapImageRep_InitWithCGImage(cgImage coregraphics.ImageRef) BitmapImageRep {
	return BitmapImageRepClass.Alloc().InitWithCGImage(cgImage)
}

func (b_ BitmapImageRep) InitWithData(data []byte) BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](b_, objc.GetSelector("initWithData:"), data)
	return rv
}

func BitmapImageRep_InitWithData(data []byte) BitmapImageRep {
	return BitmapImageRepClass.Alloc().InitWithData(data)
}

func (b_ BitmapImageRep) InitForIncrementalLoad() BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](b_, objc.GetSelector("initForIncrementalLoad"))
	return rv
}

func BitmapImageRep_InitForIncrementalLoad() BitmapImageRep {
	return BitmapImageRepClass.Alloc().InitForIncrementalLoad()
}

func (b_ BitmapImageRep) InitWithFocusedViewRect(rect foundation.Rect) BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](b_, objc.GetSelector("initWithFocusedViewRect:"), rect)
	return rv
}

func BitmapImageRep_InitWithFocusedViewRect(rect foundation.Rect) BitmapImageRep {
	return BitmapImageRepClass.Alloc().InitWithFocusedViewRect(rect)
}

func (b_ BitmapImageRep) Init() BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](b_, objc.GetSelector("init"))
	return rv
}

func BitmapImageRep_Init() BitmapImageRep {
	return BitmapImageRepClass.Alloc().Init()
}

func (bc _BitmapImageRepClass) Alloc() BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](bc, objc.GetSelector("alloc"))
	return rv
}

func BitmapImageRep_Alloc() BitmapImageRep {
	return BitmapImageRepClass.Alloc()
}

func (bc _BitmapImageRepClass) New() BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](bc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewBitmapImageRep() BitmapImageRep {
	return BitmapImageRepClass.New()
}

func BitmapImageRep_New() BitmapImageRep {
	return BitmapImageRepClass.New()
}

func (bc _BitmapImageRepClass) ImageRepsWithData(data []byte) []ImageRep {
	rv := objc.CallMethod[[]ImageRep](bc, objc.GetSelector("imageRepsWithData:"), data)
	return rv
}

func BitmapImageRep_ImageRepsWithData(data []byte) []ImageRep {
	return BitmapImageRepClass.ImageRepsWithData(data)
}

func (b_ BitmapImageRep) ColorizeByMappingGrayToColorBlackMappingWhiteMapping(midPoint float64, midPointColor IColor, shadowColor IColor, lightColor IColor) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("colorizeByMappingGray:toColor:blackMapping:whiteMapping:"), midPoint, objc.ExtractPtr(midPointColor), objc.ExtractPtr(shadowColor), objc.ExtractPtr(lightColor))
}

func (bc _BitmapImageRepClass) TIFFRepresentationOfImageRepsInArray(array []IImageRep) []byte {
	rv := objc.CallMethod[[]byte](bc, objc.GetSelector("TIFFRepresentationOfImageRepsInArray:"), array)
	return rv
}

func BitmapImageRep_TIFFRepresentationOfImageRepsInArray(array []IImageRep) []byte {
	return BitmapImageRepClass.TIFFRepresentationOfImageRepsInArray(array)
}

func (bc _BitmapImageRepClass) TIFFRepresentationOfImageRepsInArrayUsingCompressionFactor(array []IImageRep, comp TIFFCompression, factor float32) []byte {
	rv := objc.CallMethod[[]byte](bc, objc.GetSelector("TIFFRepresentationOfImageRepsInArray:usingCompression:factor:"), array, comp, factor)
	return rv
}

func BitmapImageRep_TIFFRepresentationOfImageRepsInArrayUsingCompressionFactor(array []IImageRep, comp TIFFCompression, factor float32) []byte {
	return BitmapImageRepClass.TIFFRepresentationOfImageRepsInArrayUsingCompressionFactor(array, comp, factor)
}

func (b_ BitmapImageRep) TIFFRepresentationUsingCompressionFactor(comp TIFFCompression, factor float32) []byte {
	rv := objc.CallMethod[[]byte](b_, objc.GetSelector("TIFFRepresentationUsingCompression:factor:"), comp, factor)
	return rv
}

func (bc _BitmapImageRepClass) RepresentationOfImageRepsInArrayUsingTypeProperties(imageReps []IImageRep, storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.IObject) []byte {
	rv := objc.CallMethod[[]byte](bc, objc.GetSelector("representationOfImageRepsInArray:usingType:properties:"), imageReps, storageType, properties)
	return rv
}

func BitmapImageRep_RepresentationOfImageRepsInArrayUsingTypeProperties(imageReps []IImageRep, storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.IObject) []byte {
	return BitmapImageRepClass.RepresentationOfImageRepsInArrayUsingTypeProperties(imageReps, storageType, properties)
}

func (b_ BitmapImageRep) RepresentationUsingTypeProperties(storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.IObject) []byte {
	rv := objc.CallMethod[[]byte](b_, objc.GetSelector("representationUsingType:properties:"), storageType, properties)
	return rv
}

func (bc _BitmapImageRepClass) LocalizedNameForTIFFCompressionType(compression TIFFCompression) string {
	rv := objc.CallMethod[string](bc, objc.GetSelector("localizedNameForTIFFCompressionType:"), compression)
	return rv
}

func BitmapImageRep_LocalizedNameForTIFFCompressionType(compression TIFFCompression) string {
	return BitmapImageRepClass.LocalizedNameForTIFFCompressionType(compression)
}

func (b_ BitmapImageRep) CanBeCompressedUsing(compression TIFFCompression) bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("canBeCompressedUsing:"), compression)
	return rv
}

func (b_ BitmapImageRep) SetCompressionFactor(compression TIFFCompression, factor float32) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setCompression:factor:"), compression, factor)
}

func (b_ BitmapImageRep) GetCompressionFactor(compression *TIFFCompression, factor *float32) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("getCompression:factor:"), compression, factor)
}

func (b_ BitmapImageRep) SetPropertyWithValue(property BitmapImageRepPropertyKey, value objc.IObject) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setProperty:withValue:"), property, objc.ExtractPtr(value))
}

func (b_ BitmapImageRep) ValueForProperty(property BitmapImageRepPropertyKey) objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("valueForProperty:"), property)
	return rv
}

func (b_ BitmapImageRep) IncrementalLoadFromDataComplete(data []byte, complete bool) int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("incrementalLoadFromData:complete:"), data, complete)
	return rv
}

func (b_ BitmapImageRep) SetColorAtXY(color IColor, x int, y int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setColor:atX:y:"), objc.ExtractPtr(color), x, y)
}

func (b_ BitmapImageRep) ColorAtXY(x int, y int) Color {
	rv := objc.CallMethod[Color](b_, objc.GetSelector("colorAtX:y:"), x, y)
	return rv
}

func (b_ BitmapImageRep) SetPixelAtXY(p *uint, x int, y int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setPixel:atX:y:"), p, x, y)
}

func (b_ BitmapImageRep) GetPixelAtXY(p *uint, x int, y int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("getPixel:atX:y:"), p, x, y)
}

func (b_ BitmapImageRep) BitmapImageRepByConvertingToColorSpaceRenderingIntent(targetSpace IColorSpace, renderingIntent ColorRenderingIntent) BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](b_, objc.GetSelector("bitmapImageRepByConvertingToColorSpace:renderingIntent:"), objc.ExtractPtr(targetSpace), renderingIntent)
	return rv
}

func (b_ BitmapImageRep) BitmapImageRepByRetaggingWithColorSpace(newSpace IColorSpace) BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](b_, objc.GetSelector("bitmapImageRepByRetaggingWithColorSpace:"), objc.ExtractPtr(newSpace))
	return rv
}

func (b_ BitmapImageRep) BitmapFormat() BitmapFormat {
	rv := objc.CallMethod[BitmapFormat](b_, objc.GetSelector("bitmapFormat"))
	return rv
}

func (b_ BitmapImageRep) BitsPerPixel() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("bitsPerPixel"))
	return rv
}

func (b_ BitmapImageRep) BytesPerPlane() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("bytesPerPlane"))
	return rv
}

func (b_ BitmapImageRep) BytesPerRow() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("bytesPerRow"))
	return rv
}

func (b_ BitmapImageRep) IsPlanar() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("isPlanar"))
	return rv
}

func (b_ BitmapImageRep) NumberOfPlanes() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("numberOfPlanes"))
	return rv
}

func (b_ BitmapImageRep) SamplesPerPixel() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("samplesPerPixel"))
	return rv
}

func (b_ BitmapImageRep) BitmapData() *byte {
	rv := objc.CallMethod[*byte](b_, objc.GetSelector("bitmapData"))
	return rv
}

func (b_ BitmapImageRep) TIFFRepresentation() []byte {
	rv := objc.CallMethod[[]byte](b_, objc.GetSelector("TIFFRepresentation"))
	return rv
}

func (b_ BitmapImageRep) CGImage() coregraphics.ImageRef {
	rv := objc.CallMethod[coregraphics.ImageRef](b_, objc.GetSelector("CGImage"))
	return rv
}

func (b_ BitmapImageRep) ColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](b_, objc.GetSelector("colorSpace"))
	return rv
}
