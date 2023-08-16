// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coreimage"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BitmapImageRep] class.
var BitmapImageRepClass = _BitmapImageRepClass{objc.GetClass("NSBitmapImageRep")}

type _BitmapImageRepClass struct {
	objc.Class
}

// An interface definition for the [BitmapImageRep] class.
type IBitmapImageRep interface {
	IImageRep
	BitmapImageRepByConvertingToColorSpaceRenderingIntent(targetSpace IColorSpace, renderingIntent ColorRenderingIntent) BitmapImageRep
	SetCompressionFactor(compression TIFFCompression, factor float64)
	TIFFRepresentationUsingCompressionFactor(comp TIFFCompression, factor float64) []byte
	CanBeCompressedUsing(compression TIFFCompression) bool
	GetPixelAtXY(p *uint, x int, y int)
	IncrementalLoadFromDataComplete(data []byte, complete bool) int
	ColorAtXY(x int, y int) Color
	GetCompressionFactor(compression *TIFFCompression, factor *float64)
	SetPropertyWithValue(property BitmapImageRepPropertyKey, value objc.IObject)
	SetPixelAtXY(p *uint, x int, y int)
	GetBitmapDataPlanes(data *uint8)
	ColorizeByMappingGrayToColorBlackMappingWhiteMapping(midPoint float64, midPointColor IColor, shadowColor IColor, lightColor IColor)
	ValueForProperty(property BitmapImageRepPropertyKey) objc.Object
	SetColorAtXY(color IColor, x int, y int)
	RepresentationUsingTypeProperties(storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.IObject) []byte
	BitmapImageRepByRetaggingWithColorSpace(newSpace IColorSpace) BitmapImageRep
	TIFFRepresentation() []byte
	CGImage() coregraphics.ImageRef
	SamplesPerPixel() int
	BytesPerPlane() int
	BytesPerRow() int
	ColorSpace() ColorSpace
	BitmapFormat() BitmapFormat
	BitmapData() *uint8
	IsPlanar() bool
	BitsPerPixel() int
	NumberOfPlanes() int
}

// An object that renders an image from bitmap data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep?language=objc
type BitmapImageRep struct {
	ImageRep
}

func BitmapImageRepFrom(ptr unsafe.Pointer) BitmapImageRep {
	return BitmapImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}

func (b_ BitmapImageRep) InitWithCIImage(ciImage coreimage.IImage) BitmapImageRep {
	rv := objc.Call[BitmapImageRep](b_, objc.Sel("initWithCIImage:"), objc.Ptr(ciImage))
	return rv
}

// Returns a bitmap image representation from a Core Image object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395587-initwithciimage?language=objc
func BitmapImageRep_InitWithCIImage(ciImage coreimage.IImage) BitmapImageRep {
	return BitmapImageRepClass.Alloc().InitWithCIImage(ciImage)
}

func (b_ BitmapImageRep) InitForIncrementalLoad() BitmapImageRep {
	rv := objc.Call[BitmapImageRep](b_, objc.Sel("initForIncrementalLoad"))
	return rv
}

// Initializes a newly allocated bitmap image representation for incremental loading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395522-initforincrementalload?language=objc
func BitmapImageRep_InitForIncrementalLoad() BitmapImageRep {
	return BitmapImageRepClass.Alloc().InitForIncrementalLoad()
}

func (b_ BitmapImageRep) InitWithData(data []byte) BitmapImageRep {
	rv := objc.Call[BitmapImageRep](b_, objc.Sel("initWithData:"), data)
	return rv
}

// Initializes a newly allocated bitmap image representation from the specified data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395569-initwithdata?language=objc
func BitmapImageRep_InitWithData(data []byte) BitmapImageRep {
	return BitmapImageRepClass.Alloc().InitWithData(data)
}

func (b_ BitmapImageRep) InitWithBitmapDataPlanesPixelsWidePixelsHighBitsPerSampleSamplesPerPixelHasAlphaIsPlanarColorSpaceNameBytesPerRowBitsPerPixel(planes *uint8, width int, height int, bps int, spp int, alpha bool, isPlanar bool, colorSpaceName ColorSpaceName, rBytes int, pBits int) BitmapImageRep {
	rv := objc.Call[BitmapImageRep](b_, objc.Sel("initWithBitmapDataPlanes:pixelsWide:pixelsHigh:bitsPerSample:samplesPerPixel:hasAlpha:isPlanar:colorSpaceName:bytesPerRow:bitsPerPixel:"), planes, width, height, bps, spp, alpha, isPlanar, colorSpaceName, rBytes, pBits)
	return rv
}

// Initializes a newly allocated bitmap image representation so it can render the specified image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395540-initwithbitmapdataplanes?language=objc
func BitmapImageRep_InitWithBitmapDataPlanesPixelsWidePixelsHighBitsPerSampleSamplesPerPixelHasAlphaIsPlanarColorSpaceNameBytesPerRowBitsPerPixel(planes *uint8, width int, height int, bps int, spp int, alpha bool, isPlanar bool, colorSpaceName ColorSpaceName, rBytes int, pBits int) BitmapImageRep {
	return BitmapImageRepClass.Alloc().InitWithBitmapDataPlanesPixelsWidePixelsHighBitsPerSampleSamplesPerPixelHasAlphaIsPlanarColorSpaceNameBytesPerRowBitsPerPixel(planes, width, height, bps, spp, alpha, isPlanar, colorSpaceName, rBytes, pBits)
}

func (bc _BitmapImageRepClass) ImageRepWithData(data []byte) BitmapImageRep {
	rv := objc.Call[BitmapImageRep](bc, objc.Sel("imageRepWithData:"), data)
	return rv
}

// Creates and returns a bitmap image representation with the first image in the specified data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395502-imagerepwithdata?language=objc
func BitmapImageRep_ImageRepWithData(data []byte) BitmapImageRep {
	return BitmapImageRepClass.ImageRepWithData(data)
}

func (b_ BitmapImageRep) InitWithCGImage(cgImage coregraphics.ImageRef) BitmapImageRep {
	rv := objc.Call[BitmapImageRep](b_, objc.Sel("initWithCGImage:"), cgImage)
	return rv
}

// Returns a bitmap image representation from a Core Graphics image object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395423-initwithcgimage?language=objc
func BitmapImageRep_InitWithCGImage(cgImage coregraphics.ImageRef) BitmapImageRep {
	return BitmapImageRepClass.Alloc().InitWithCGImage(cgImage)
}

func (bc _BitmapImageRepClass) Alloc() BitmapImageRep {
	rv := objc.Call[BitmapImageRep](bc, objc.Sel("alloc"))
	return rv
}

func BitmapImageRep_Alloc() BitmapImageRep {
	return BitmapImageRepClass.Alloc()
}

func (bc _BitmapImageRepClass) New() BitmapImageRep {
	rv := objc.Call[BitmapImageRep](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBitmapImageRep() BitmapImageRep {
	return BitmapImageRepClass.New()
}

func (b_ BitmapImageRep) Init() BitmapImageRep {
	rv := objc.Call[BitmapImageRep](b_, objc.Sel("init"))
	return rv
}

// Formats the specified bitmap images using the specified storage type and properties and returns them in a data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395520-representationofimagerepsinarray?language=objc
func (bc _BitmapImageRepClass) RepresentationOfImageRepsInArrayUsingTypeProperties(imageReps []IImageRep, storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.IObject) []byte {
	rv := objc.Call[[]byte](bc, objc.Sel("representationOfImageRepsInArray:usingType:properties:"), imageReps, storageType, properties)
	return rv
}

// Formats the specified bitmap images using the specified storage type and properties and returns them in a data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395520-representationofimagerepsinarray?language=objc
func BitmapImageRep_RepresentationOfImageRepsInArrayUsingTypeProperties(imageReps []IImageRep, storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.IObject) []byte {
	return BitmapImageRepClass.RepresentationOfImageRepsInArrayUsingTypeProperties(imageReps, storageType, properties)
}

// Converts the bitmap image representation to the specified color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395470-bitmapimagerepbyconvertingtocolo?language=objc
func (b_ BitmapImageRep) BitmapImageRepByConvertingToColorSpaceRenderingIntent(targetSpace IColorSpace, renderingIntent ColorRenderingIntent) BitmapImageRep {
	rv := objc.Call[BitmapImageRep](b_, objc.Sel("bitmapImageRepByConvertingToColorSpace:renderingIntent:"), objc.Ptr(targetSpace), renderingIntent)
	return rv
}

// Sets the bitmap image representation’s compression type and compression factor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395478-setcompression?language=objc
func (b_ BitmapImageRep) SetCompressionFactor(compression TIFFCompression, factor float64) {
	objc.Call[objc.Void](b_, objc.Sel("setCompression:factor:"), compression, factor)
}

// Returns a TIFF representation of the image using the specified compression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395524-tiffrepresentationusingcompressi?language=objc
func (b_ BitmapImageRep) TIFFRepresentationUsingCompressionFactor(comp TIFFCompression, factor float64) []byte {
	rv := objc.Call[[]byte](b_, objc.Sel("TIFFRepresentationUsingCompression:factor:"), comp, factor)
	return rv
}

// Tests whether the bitmap image representation can be compressed by the specified compression scheme. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395427-canbecompressedusing?language=objc
func (b_ BitmapImageRep) CanBeCompressedUsing(compression TIFFCompression) bool {
	rv := objc.Call[bool](b_, objc.Sel("canBeCompressedUsing:"), compression)
	return rv
}

// Returns by indirection the pixel data for the specified location in the bitmap image representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395567-getpixel?language=objc
func (b_ BitmapImageRep) GetPixelAtXY(p *uint, x int, y int) {
	objc.Call[objc.Void](b_, objc.Sel("getPixel:atX:y:"), p, x, y)
}

// Returns by indirection an array of all available compression types that can be used when writing a TIFF image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395513-gettiffcompressiontypes?language=objc
func (bc _BitmapImageRepClass) GetTIFFCompressionTypesCount(list *TIFFCompression, numTypes *int) {
	objc.Call[objc.Void](bc, objc.Sel("getTIFFCompressionTypes:count:"), list, numTypes)
}

// Returns by indirection an array of all available compression types that can be used when writing a TIFF image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395513-gettiffcompressiontypes?language=objc
func BitmapImageRep_GetTIFFCompressionTypesCount(list *TIFFCompression, numTypes *int) {
	BitmapImageRepClass.GetTIFFCompressionTypesCount(list, numTypes)
}

// Loads the current image data into an incrementally-loaded image representation and returns the current status of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395464-incrementalloadfromdata?language=objc
func (b_ BitmapImageRep) IncrementalLoadFromDataComplete(data []byte, complete bool) int {
	rv := objc.Call[int](b_, objc.Sel("incrementalLoadFromData:complete:"), data, complete)
	return rv
}

// Returns the color of the pixel at the specified coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395583-coloratx?language=objc
func (b_ BitmapImageRep) ColorAtXY(x int, y int) Color {
	rv := objc.Call[Color](b_, objc.Sel("colorAtX:y:"), x, y)
	return rv
}

// Returns by indirection the bitmap image representation’s compression type and compression factor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395515-getcompression?language=objc
func (b_ BitmapImageRep) GetCompressionFactor(compression *TIFFCompression, factor *float64) {
	objc.Call[objc.Void](b_, objc.Sel("getCompression:factor:"), compression, factor)
}

// Sets the specified property of the bitmap image representation to the specified value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395486-setproperty?language=objc
func (b_ BitmapImageRep) SetPropertyWithValue(property BitmapImageRepPropertyKey, value objc.IObject) {
	objc.Call[objc.Void](b_, objc.Sel("setProperty:withValue:"), property, value)
}

// Sets the bitmap image representation’s pixel at the specified coordinates to the specified raw pixel values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395446-setpixel?language=objc
func (b_ BitmapImageRep) SetPixelAtXY(p *uint, x int, y int) {
	objc.Call[objc.Void](b_, objc.Sel("setPixel:atX:y:"), p, x, y)
}

// Returns by indirection bitmap data of the bitmap image representation separated into planes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395490-getbitmapdataplanes?language=objc
func (b_ BitmapImageRep) GetBitmapDataPlanes(data *uint8) {
	objc.Call[objc.Void](b_, objc.Sel("getBitmapDataPlanes:"), data)
}

// Colorizes a grayscale image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395530-colorizebymappinggray?language=objc
func (b_ BitmapImageRep) ColorizeByMappingGrayToColorBlackMappingWhiteMapping(midPoint float64, midPointColor IColor, shadowColor IColor, lightColor IColor) {
	objc.Call[objc.Void](b_, objc.Sel("colorizeByMappingGray:toColor:blackMapping:whiteMapping:"), midPoint, objc.Ptr(midPointColor), objc.Ptr(shadowColor), objc.Ptr(lightColor))
}

// Creates and returns an array of bitmap image representation objects that correspond to the images in the specified data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395419-imagerepswithdata?language=objc
func (bc _BitmapImageRepClass) ImageRepsWithData(data []byte) []ImageRep {
	rv := objc.Call[[]ImageRep](bc, objc.Sel("imageRepsWithData:"), data)
	return rv
}

// Creates and returns an array of bitmap image representation objects that correspond to the images in the specified data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395419-imagerepswithdata?language=objc
func BitmapImageRep_ImageRepsWithData(data []byte) []ImageRep {
	return BitmapImageRepClass.ImageRepsWithData(data)
}

// Returns a TIFF representation of the specified images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395561-tiffrepresentationofimagerepsina?language=objc
func (bc _BitmapImageRepClass) TIFFRepresentationOfImageRepsInArray(array []IImageRep) []byte {
	rv := objc.Call[[]byte](bc, objc.Sel("TIFFRepresentationOfImageRepsInArray:"), array)
	return rv
}

// Returns a TIFF representation of the specified images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395561-tiffrepresentationofimagerepsina?language=objc
func BitmapImageRep_TIFFRepresentationOfImageRepsInArray(array []IImageRep) []byte {
	return BitmapImageRepClass.TIFFRepresentationOfImageRepsInArray(array)
}

// Returns the value for the specified property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395492-valueforproperty?language=objc
func (b_ BitmapImageRep) ValueForProperty(property BitmapImageRepPropertyKey) objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("valueForProperty:"), property)
	return rv
}

// Changes the color of the pixel at the specified coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395472-setcolor?language=objc
func (b_ BitmapImageRep) SetColorAtXY(color IColor, x int, y int) {
	objc.Call[objc.Void](b_, objc.Sel("setColor:atX:y:"), objc.Ptr(color), x, y)
}

// Formats the bitmap representation’s image data using the specified storage type and properties and returns it in a data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395458-representationusingtype?language=objc
func (b_ BitmapImageRep) RepresentationUsingTypeProperties(storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.IObject) []byte {
	rv := objc.Call[[]byte](b_, objc.Sel("representationUsingType:properties:"), storageType, properties)
	return rv
}

// Returns an autoreleased string containing the localized name for the specified compression type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395548-localizednamefortiffcompressiont?language=objc
func (bc _BitmapImageRepClass) LocalizedNameForTIFFCompressionType(compression TIFFCompression) string {
	rv := objc.Call[string](bc, objc.Sel("localizedNameForTIFFCompressionType:"), compression)
	return rv
}

// Returns an autoreleased string containing the localized name for the specified compression type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395548-localizednamefortiffcompressiont?language=objc
func BitmapImageRep_LocalizedNameForTIFFCompressionType(compression TIFFCompression) string {
	return BitmapImageRepClass.LocalizedNameForTIFFCompressionType(compression)
}

// Changes the color space tag of the bitmap image representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395512-bitmapimagerepbyretaggingwithcol?language=objc
func (b_ BitmapImageRep) BitmapImageRepByRetaggingWithColorSpace(newSpace IColorSpace) BitmapImageRep {
	rv := objc.Call[BitmapImageRep](b_, objc.Sel("bitmapImageRepByRetaggingWithColorSpace:"), objc.Ptr(newSpace))
	return rv
}

// A TIFF representation of the bitmap image data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395557-tiffrepresentation?language=objc
func (b_ BitmapImageRep) TIFFRepresentation() []byte {
	rv := objc.Call[[]byte](b_, objc.Sel("TIFFRepresentation"))
	return rv
}

// A Core Graphics image object based on the bitmap image representation’s data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395506-cgimage?language=objc
func (b_ BitmapImageRep) CGImage() coregraphics.ImageRef {
	rv := objc.Call[coregraphics.ImageRef](b_, objc.Sel("CGImage"))
	return rv
}

// The number of components for each pixel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395577-samplesperpixel?language=objc
func (b_ BitmapImageRep) SamplesPerPixel() int {
	rv := objc.Call[int](b_, objc.Sel("samplesPerPixel"))
	return rv
}

// The number of bytes in each plane or channel of data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395559-bytesperplane?language=objc
func (b_ BitmapImageRep) BytesPerPlane() int {
	rv := objc.Call[int](b_, objc.Sel("bytesPerPlane"))
	return rv
}

// The minimum number of bytes required to specify a scan line in each data plane. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395454-bytesperrow?language=objc
func (b_ BitmapImageRep) BytesPerRow() int {
	rv := objc.Call[int](b_, objc.Sel("bytesPerRow"))
	return rv
}

// The color space of the bitmap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395468-colorspace?language=objc
func (b_ BitmapImageRep) ColorSpace() ColorSpace {
	rv := objc.Call[ColorSpace](b_, objc.Sel("colorSpace"))
	return rv
}

// The format of the bitmap image representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395508-bitmapformat?language=objc
func (b_ BitmapImageRep) BitmapFormat() BitmapFormat {
	rv := objc.Call[BitmapFormat](b_, objc.Sel("bitmapFormat"))
	return rv
}

// A pointer to the bitmap data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395421-bitmapdata?language=objc
func (b_ BitmapImageRep) BitmapData() *uint8 {
	rv := objc.Call[*uint8](b_, objc.Sel("bitmapData"))
	return rv
}

// A Boolean value that indicates whether the image data is in a planar configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395482-planar?language=objc
func (b_ BitmapImageRep) IsPlanar() bool {
	rv := objc.Call[bool](b_, objc.Sel("isPlanar"))
	return rv
}

// The number of bits allocated for each pixel in each plane of data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395488-bitsperpixel?language=objc
func (b_ BitmapImageRep) BitsPerPixel() int {
	rv := objc.Call[int](b_, objc.Sel("bitsPerPixel"))
	return rv
}

// The number of separate planes into which the image data is organized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/1395474-numberofplanes?language=objc
func (b_ BitmapImageRep) NumberOfPlanes() int {
	rv := objc.Call[int](b_, objc.Sel("numberOfPlanes"))
	return rv
}
