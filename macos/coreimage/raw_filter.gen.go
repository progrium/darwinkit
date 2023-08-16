// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/imageio"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RAWFilter] class.
var RAWFilterClass = _RAWFilterClass{objc.GetClass("CIRAWFilter")}

type _RAWFilterClass struct {
	objc.Class
}

// An interface definition for the [RAWFilter] class.
type IRAWFilter interface {
	IFilter
	SemanticSegmentationSkyMatte() Image
	IsSharpnessSupported() bool
	SemanticSegmentationTeethMatte() Image
	ScaleFactor() float64
	SetScaleFactor(value float64)
	IsLocalToneMapSupported() bool
	NeutralChromaticity() coregraphics.Point
	SetNeutralChromaticity(value coregraphics.Point)
	DecoderVersion() RAWDecoderVersion
	SetDecoderVersion(value RAWDecoderVersion)
	IsLensCorrectionSupported() bool
	LocalToneMapAmount() float64
	SetLocalToneMapAmount(value float64)
	IsLensCorrectionEnabled() bool
	SetLensCorrectionEnabled(value bool)
	IsGamutMappingEnabled() bool
	SetGamutMappingEnabled(value bool)
	IsDetailSupported() bool
	IsMoireReductionSupported() bool
	SharpnessAmount() float64
	SetSharpnessAmount(value float64)
	ContrastAmount() float64
	SetContrastAmount(value float64)
	IsLuminanceNoiseReductionSupported() bool
	Properties() foundation.Dictionary
	IsContrastSupported() bool
	DetailAmount() float64
	SetDetailAmount(value float64)
	ExtendedDynamicRangeAmount() float64
	SetExtendedDynamicRangeAmount(value float64)
	IsColorNoiseReductionSupported() bool
	IsDraftModeEnabled() bool
	SetDraftModeEnabled(value bool)
	Exposure() float64
	SetExposure(value float64)
	BoostAmount() float64
	SetBoostAmount(value float64)
	SemanticSegmentationHairMatte() Image
	SemanticSegmentationGlassesMatte() Image
	NeutralTint() float64
	SetNeutralTint(value float64)
	NativeSize() coregraphics.Size
	BoostShadowAmount() float64
	SetBoostShadowAmount(value float64)
	ColorNoiseReductionAmount() float64
	SetColorNoiseReductionAmount(value float64)
	SemanticSegmentationSkinMatte() Image
	NeutralTemperature() float64
	SetNeutralTemperature(value float64)
	PreviewImage() Image
	LuminanceNoiseReductionAmount() float64
	SetLuminanceNoiseReductionAmount(value float64)
	ShadowBias() float64
	SetShadowBias(value float64)
	SupportedDecoderVersions() []RAWDecoderVersion
	NeutralLocation() coregraphics.Point
	SetNeutralLocation(value coregraphics.Point)
	BaselineExposure() float64
	SetBaselineExposure(value float64)
	PortraitEffectsMatte() Image
	LinearSpaceFilter() Filter
	SetLinearSpaceFilter(value IFilter)
	MoireReductionAmount() float64
	SetMoireReductionAmount(value float64)
	Orientation() imageio.ImagePropertyOrientation
	SetOrientation(value imageio.ImagePropertyOrientation)
}

// A filter subclass that produces an image by manipulating RAW image sensor data from a digital camera or scanner. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter?language=objc
type RAWFilter struct {
	Filter
}

func RAWFilterFrom(ptr unsafe.Pointer) RAWFilter {
	return RAWFilter{
		Filter: FilterFrom(ptr),
	}
}

func (rc _RAWFilterClass) FilterWithCVPixelBufferProperties(buffer corevideo.PixelBufferRef, properties foundation.Dictionary) RAWFilter {
	rv := objc.Call[RAWFilter](rc, objc.Sel("filterWithCVPixelBuffer:properties:"), buffer, properties)
	return rv
}

// Creates a RAW filter from the pixel buffer and its properties that you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801628-filterwithcvpixelbuffer?language=objc
func RAWFilter_FilterWithCVPixelBufferProperties(buffer corevideo.PixelBufferRef, properties foundation.Dictionary) RAWFilter {
	return RAWFilterClass.FilterWithCVPixelBufferProperties(buffer, properties)
}

func (rc _RAWFilterClass) FilterWithImageURL(url foundation.IURL) RAWFilter {
	rv := objc.Call[RAWFilter](rc, objc.Sel("filterWithImageURL:"), objc.Ptr(url))
	return rv
}

// Creates a RAW filter from the image at the URL location that you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801630-filterwithimageurl?language=objc
func RAWFilter_FilterWithImageURL(url foundation.IURL) RAWFilter {
	return RAWFilterClass.FilterWithImageURL(url)
}

func (rc _RAWFilterClass) FilterWithImageDataIdentifierHint(data []byte, identifierHint string) RAWFilter {
	rv := objc.Call[RAWFilter](rc, objc.Sel("filterWithImageData:identifierHint:"), data, identifierHint)
	return rv
}

// Creates a RAW filter from the image data and type hint that you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801629-filterwithimagedata?language=objc
func RAWFilter_FilterWithImageDataIdentifierHint(data []byte, identifierHint string) RAWFilter {
	return RAWFilterClass.FilterWithImageDataIdentifierHint(data, identifierHint)
}

func (rc _RAWFilterClass) Alloc() RAWFilter {
	rv := objc.Call[RAWFilter](rc, objc.Sel("alloc"))
	return rv
}

func RAWFilter_Alloc() RAWFilter {
	return RAWFilterClass.Alloc()
}

func (rc _RAWFilterClass) New() RAWFilter {
	rv := objc.Call[RAWFilter](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRAWFilter() RAWFilter {
	return RAWFilterClass.New()
}

func (r_ RAWFilter) Init() RAWFilter {
	rv := objc.Call[RAWFilter](r_, objc.Sel("init"))
	return rv
}

// An optional auxiliary image that represents the semantic segmentation sky matte of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801654-semanticsegmentationskymatte?language=objc
func (r_ RAWFilter) SemanticSegmentationSkyMatte() Image {
	rv := objc.Call[Image](r_, objc.Sel("semanticSegmentationSkyMatte"))
	return rv
}

// A Boolean that indicates if the current image supports sharpness adjustments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801658-sharpnesssupported?language=objc
func (r_ RAWFilter) IsSharpnessSupported() bool {
	rv := objc.Call[bool](r_, objc.Sel("isSharpnessSupported"))
	return rv
}

// An optional auxiliary image that represents the semantic segmentation teeth matte of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801655-semanticsegmentationteethmatte?language=objc
func (r_ RAWFilter) SemanticSegmentationTeethMatte() Image {
	rv := objc.Call[Image](r_, objc.Sel("semanticSegmentationTeethMatte"))
	return rv
}

// An array containing the names of all supported camera models. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801659-supportedcameramodels?language=objc
func (rc _RAWFilterClass) SupportedCameraModels() []string {
	rv := objc.Call[[]string](rc, objc.Sel("supportedCameraModels"))
	return rv
}

// An array containing the names of all supported camera models. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801659-supportedcameramodels?language=objc
func RAWFilter_SupportedCameraModels() []string {
	return RAWFilterClass.SupportedCameraModels()
}

// A value that indicates the desired scale factor to draw the output image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801650-scalefactor?language=objc
func (r_ RAWFilter) ScaleFactor() float64 {
	rv := objc.Call[float64](r_, objc.Sel("scaleFactor"))
	return rv
}

// A value that indicates the desired scale factor to draw the output image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801650-scalefactor?language=objc
func (r_ RAWFilter) SetScaleFactor(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setScaleFactor:"), value)
}

// A Boolean that indicates if the current image supports local tone curve adjustments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801636-localtonemapsupported?language=objc
func (r_ RAWFilter) IsLocalToneMapSupported() bool {
	rv := objc.Call[bool](r_, objc.Sel("isLocalToneMapSupported"))
	return rv
}

// A value that indicates the amount of white balance based on chromaticity values to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801642-neutralchromaticity?language=objc
func (r_ RAWFilter) NeutralChromaticity() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](r_, objc.Sel("neutralChromaticity"))
	return rv
}

// A value that indicates the amount of white balance based on chromaticity values to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801642-neutralchromaticity?language=objc
func (r_ RAWFilter) SetNeutralChromaticity(value coregraphics.Point) {
	objc.Call[objc.Void](r_, objc.Sel("setNeutralChromaticity:"), value)
}

// A value that indicates the decoder version to use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801622-decoderversion?language=objc
func (r_ RAWFilter) DecoderVersion() RAWDecoderVersion {
	rv := objc.Call[RAWDecoderVersion](r_, objc.Sel("decoderVersion"))
	return rv
}

// A value that indicates the decoder version to use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801622-decoderversion?language=objc
func (r_ RAWFilter) SetDecoderVersion(value RAWDecoderVersion) {
	objc.Call[objc.Void](r_, objc.Sel("setDecoderVersion:"), value)
}

// A Boolean that indicates if you can enable lens correction for the current image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801633-lenscorrectionsupported?language=objc
func (r_ RAWFilter) IsLensCorrectionSupported() bool {
	rv := objc.Call[bool](r_, objc.Sel("isLensCorrectionSupported"))
	return rv
}

// A value that indicates the amount of local tone curve to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801635-localtonemapamount?language=objc
func (r_ RAWFilter) LocalToneMapAmount() float64 {
	rv := objc.Call[float64](r_, objc.Sel("localToneMapAmount"))
	return rv
}

// A value that indicates the amount of local tone curve to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801635-localtonemapamount?language=objc
func (r_ RAWFilter) SetLocalToneMapAmount(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setLocalToneMapAmount:"), value)
}

// A Boolean that indicates whether to enable lens correction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801632-lenscorrectionenabled?language=objc
func (r_ RAWFilter) IsLensCorrectionEnabled() bool {
	rv := objc.Call[bool](r_, objc.Sel("isLensCorrectionEnabled"))
	return rv
}

// A Boolean that indicates whether to enable lens correction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801632-lenscorrectionenabled?language=objc
func (r_ RAWFilter) SetLensCorrectionEnabled(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setLensCorrectionEnabled:"), value)
}

// A Boolean that indicates whether to enable gamut mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801631-gamutmappingenabled?language=objc
func (r_ RAWFilter) IsGamutMappingEnabled() bool {
	rv := objc.Call[bool](r_, objc.Sel("isGamutMappingEnabled"))
	return rv
}

// A Boolean that indicates whether to enable gamut mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801631-gamutmappingenabled?language=objc
func (r_ RAWFilter) SetGamutMappingEnabled(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setGamutMappingEnabled:"), value)
}

// A Boolean that indicates if the current image supports detail enhancement adjustments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801624-detailsupported?language=objc
func (r_ RAWFilter) IsDetailSupported() bool {
	rv := objc.Call[bool](r_, objc.Sel("isDetailSupported"))
	return rv
}

// A Boolean that indicates if the current image supports moire artifact reduction adjustments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801640-moirereductionsupported?language=objc
func (r_ RAWFilter) IsMoireReductionSupported() bool {
	rv := objc.Call[bool](r_, objc.Sel("isMoireReductionSupported"))
	return rv
}

// A value that indicates the amount of sharpness to apply to the edges of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801657-sharpnessamount?language=objc
func (r_ RAWFilter) SharpnessAmount() float64 {
	rv := objc.Call[float64](r_, objc.Sel("sharpnessAmount"))
	return rv
}

// A value that indicates the amount of sharpness to apply to the edges of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801657-sharpnessamount?language=objc
func (r_ RAWFilter) SetSharpnessAmount(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setSharpnessAmount:"), value)
}

// A value that indicates the amount of local contrast to apply to the edges of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801620-contrastamount?language=objc
func (r_ RAWFilter) ContrastAmount() float64 {
	rv := objc.Call[float64](r_, objc.Sel("contrastAmount"))
	return rv
}

// A value that indicates the amount of local contrast to apply to the edges of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801620-contrastamount?language=objc
func (r_ RAWFilter) SetContrastAmount(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setContrastAmount:"), value)
}

// A Boolean that indicates if the current image supports luminance noise reduction adjustments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801638-luminancenoisereductionsupported?language=objc
func (r_ RAWFilter) IsLuminanceNoiseReductionSupported() bool {
	rv := objc.Call[bool](r_, objc.Sel("isLuminanceNoiseReductionSupported"))
	return rv
}

// A dictionary that contains properties of the image source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801649-properties?language=objc
func (r_ RAWFilter) Properties() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](r_, objc.Sel("properties"))
	return rv
}

// A Boolean that indicates if the current image supports contrast adjustments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801621-contrastsupported?language=objc
func (r_ RAWFilter) IsContrastSupported() bool {
	rv := objc.Call[bool](r_, objc.Sel("isContrastSupported"))
	return rv
}

// A value that indicates the amount of detail enhancement to apply to the edges of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801623-detailamount?language=objc
func (r_ RAWFilter) DetailAmount() float64 {
	rv := objc.Call[float64](r_, objc.Sel("detailAmount"))
	return rv
}

// A value that indicates the amount of detail enhancement to apply to the edges of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801623-detailamount?language=objc
func (r_ RAWFilter) SetDetailAmount(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setDetailAmount:"), value)
}

// A value that indicates the amount of extended dynamic range (EDR) to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3820998-extendeddynamicrangeamount?language=objc
func (r_ RAWFilter) ExtendedDynamicRangeAmount() float64 {
	rv := objc.Call[float64](r_, objc.Sel("extendedDynamicRangeAmount"))
	return rv
}

// A value that indicates the amount of extended dynamic range (EDR) to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3820998-extendeddynamicrangeamount?language=objc
func (r_ RAWFilter) SetExtendedDynamicRangeAmount(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setExtendedDynamicRangeAmount:"), value)
}

// A Boolean that indicates if the current image supports color noise reduction adjustments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801619-colornoisereductionsupported?language=objc
func (r_ RAWFilter) IsColorNoiseReductionSupported() bool {
	rv := objc.Call[bool](r_, objc.Sel("isColorNoiseReductionSupported"))
	return rv
}

// A Boolean that indicates whether to enable draft mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801625-draftmodeenabled?language=objc
func (r_ RAWFilter) IsDraftModeEnabled() bool {
	rv := objc.Call[bool](r_, objc.Sel("isDraftModeEnabled"))
	return rv
}

// A Boolean that indicates whether to enable draft mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801625-draftmodeenabled?language=objc
func (r_ RAWFilter) SetDraftModeEnabled(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setDraftModeEnabled:"), value)
}

// A value that indicates the amount of exposure to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801626-exposure?language=objc
func (r_ RAWFilter) Exposure() float64 {
	rv := objc.Call[float64](r_, objc.Sel("exposure"))
	return rv
}

// A value that indicates the amount of exposure to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801626-exposure?language=objc
func (r_ RAWFilter) SetExposure(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setExposure:"), value)
}

// A value that indicates the amount of global tone curve to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801616-boostamount?language=objc
func (r_ RAWFilter) BoostAmount() float64 {
	rv := objc.Call[float64](r_, objc.Sel("boostAmount"))
	return rv
}

// A value that indicates the amount of global tone curve to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801616-boostamount?language=objc
func (r_ RAWFilter) SetBoostAmount(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setBoostAmount:"), value)
}

// An optional auxiliary image that represents the semantic segmentation hair matte of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801652-semanticsegmentationhairmatte?language=objc
func (r_ RAWFilter) SemanticSegmentationHairMatte() Image {
	rv := objc.Call[Image](r_, objc.Sel("semanticSegmentationHairMatte"))
	return rv
}

// An optional auxiliary image that represents the semantic segmentation glasses matte of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801651-semanticsegmentationglassesmatte?language=objc
func (r_ RAWFilter) SemanticSegmentationGlassesMatte() Image {
	rv := objc.Call[Image](r_, objc.Sel("semanticSegmentationGlassesMatte"))
	return rv
}

// A value that indicates the amount of white balance based on tint values to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801645-neutraltint?language=objc
func (r_ RAWFilter) NeutralTint() float64 {
	rv := objc.Call[float64](r_, objc.Sel("neutralTint"))
	return rv
}

// A value that indicates the amount of white balance based on tint values to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801645-neutraltint?language=objc
func (r_ RAWFilter) SetNeutralTint(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setNeutralTint:"), value)
}

// The full native size of the unscaled image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801641-nativesize?language=objc
func (r_ RAWFilter) NativeSize() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](r_, objc.Sel("nativeSize"))
	return rv
}

// A value that indicates the amount to boost the shadow areas of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801617-boostshadowamount?language=objc
func (r_ RAWFilter) BoostShadowAmount() float64 {
	rv := objc.Call[float64](r_, objc.Sel("boostShadowAmount"))
	return rv
}

// A value that indicates the amount to boost the shadow areas of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801617-boostshadowamount?language=objc
func (r_ RAWFilter) SetBoostShadowAmount(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setBoostShadowAmount:"), value)
}

// A value that indicates the amount of chroma noise reduction to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801618-colornoisereductionamount?language=objc
func (r_ RAWFilter) ColorNoiseReductionAmount() float64 {
	rv := objc.Call[float64](r_, objc.Sel("colorNoiseReductionAmount"))
	return rv
}

// A value that indicates the amount of chroma noise reduction to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801618-colornoisereductionamount?language=objc
func (r_ RAWFilter) SetColorNoiseReductionAmount(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setColorNoiseReductionAmount:"), value)
}

// An optional auxiliary image that represents the semantic segmentation skin matte of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801653-semanticsegmentationskinmatte?language=objc
func (r_ RAWFilter) SemanticSegmentationSkinMatte() Image {
	rv := objc.Call[Image](r_, objc.Sel("semanticSegmentationSkinMatte"))
	return rv
}

// A value that indicates the amount of white balance based on temperature values to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801644-neutraltemperature?language=objc
func (r_ RAWFilter) NeutralTemperature() float64 {
	rv := objc.Call[float64](r_, objc.Sel("neutralTemperature"))
	return rv
}

// A value that indicates the amount of white balance based on temperature values to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801644-neutraltemperature?language=objc
func (r_ RAWFilter) SetNeutralTemperature(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setNeutralTemperature:"), value)
}

// An optional auxiliary image that represents a preview of the original image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801648-previewimage?language=objc
func (r_ RAWFilter) PreviewImage() Image {
	rv := objc.Call[Image](r_, objc.Sel("previewImage"))
	return rv
}

// A value that indicates the amount of luminance noise reduction to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801637-luminancenoisereductionamount?language=objc
func (r_ RAWFilter) LuminanceNoiseReductionAmount() float64 {
	rv := objc.Call[float64](r_, objc.Sel("luminanceNoiseReductionAmount"))
	return rv
}

// A value that indicates the amount of luminance noise reduction to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801637-luminancenoisereductionamount?language=objc
func (r_ RAWFilter) SetLuminanceNoiseReductionAmount(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setLuminanceNoiseReductionAmount:"), value)
}

// A value that indicates the amount to subtract from the shadows in the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801656-shadowbias?language=objc
func (r_ RAWFilter) ShadowBias() float64 {
	rv := objc.Call[float64](r_, objc.Sel("shadowBias"))
	return rv
}

// A value that indicates the amount to subtract from the shadows in the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801656-shadowbias?language=objc
func (r_ RAWFilter) SetShadowBias(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setShadowBias:"), value)
}

// An array of all supported decoder versions for the given image type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801660-supporteddecoderversions?language=objc
func (r_ RAWFilter) SupportedDecoderVersions() []RAWDecoderVersion {
	rv := objc.Call[[]RAWDecoderVersion](r_, objc.Sel("supportedDecoderVersions"))
	return rv
}

// A value that indicates the amount of white balance based on pixel coordinates to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801643-neutrallocation?language=objc
func (r_ RAWFilter) NeutralLocation() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](r_, objc.Sel("neutralLocation"))
	return rv
}

// A value that indicates the amount of white balance based on pixel coordinates to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801643-neutrallocation?language=objc
func (r_ RAWFilter) SetNeutralLocation(value coregraphics.Point) {
	objc.Call[objc.Void](r_, objc.Sel("setNeutralLocation:"), value)
}

// A value that indicates the baseline exposure to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801615-baselineexposure?language=objc
func (r_ RAWFilter) BaselineExposure() float64 {
	rv := objc.Call[float64](r_, objc.Sel("baselineExposure"))
	return rv
}

// A value that indicates the baseline exposure to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801615-baselineexposure?language=objc
func (r_ RAWFilter) SetBaselineExposure(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setBaselineExposure:"), value)
}

// An optional auxiliary image that represents the portrait effects matte of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801647-portraiteffectsmatte?language=objc
func (r_ RAWFilter) PortraitEffectsMatte() Image {
	rv := objc.Call[Image](r_, objc.Sel("portraitEffectsMatte"))
	return rv
}

// An optional filter you can apply to the RAW image while it’s in linear space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801634-linearspacefilter?language=objc
func (r_ RAWFilter) LinearSpaceFilter() Filter {
	rv := objc.Call[Filter](r_, objc.Sel("linearSpaceFilter"))
	return rv
}

// An optional filter you can apply to the RAW image while it’s in linear space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801634-linearspacefilter?language=objc
func (r_ RAWFilter) SetLinearSpaceFilter(value IFilter) {
	objc.Call[objc.Void](r_, objc.Sel("setLinearSpaceFilter:"), objc.Ptr(value))
}

// A value that indicates the amount of moire artifact reduction to apply to high frequency areas of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801639-moirereductionamount?language=objc
func (r_ RAWFilter) MoireReductionAmount() float64 {
	rv := objc.Call[float64](r_, objc.Sel("moireReductionAmount"))
	return rv
}

// A value that indicates the amount of moire artifact reduction to apply to high frequency areas of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801639-moirereductionamount?language=objc
func (r_ RAWFilter) SetMoireReductionAmount(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setMoireReductionAmount:"), value)
}

// A value that indicates the orientation of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801646-orientation?language=objc
func (r_ RAWFilter) Orientation() imageio.ImagePropertyOrientation {
	rv := objc.Call[imageio.ImagePropertyOrientation](r_, objc.Sel("orientation"))
	return rv
}

// A value that indicates the orientation of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/3801646-orientation?language=objc
func (r_ RAWFilter) SetOrientation(value imageio.ImagePropertyOrientation) {
	objc.Call[objc.Void](r_, objc.Sel("setOrientation:"), value)
}
