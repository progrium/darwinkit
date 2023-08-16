// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicontextoption?language=objc
type ContextOption string

const (
	kCIContextAllowLowPower         ContextOption = "kCIContextAllowLowPower"
	kCIContextCacheIntermediates    ContextOption = "kCIContextCacheIntermediates"
	kCIContextHighQualityDownsample ContextOption = "high_quality_downsample"
	kCIContextName                  ContextOption = "kCIContextName"
	kCIContextOutputColorSpace      ContextOption = "output_color_space"
	kCIContextOutputPremultiplied   ContextOption = "output_premultiplied"
	kCIContextPriorityRequestLow    ContextOption = "priority_request_low"
	kCIContextUseSoftwareRenderer   ContextOption = "software_renderer"
	kCIContextWorkingColorSpace     ContextOption = "working_color_space"
	kCIContextWorkingFormat         ContextOption = "working_format"
)

// Constants concerning Data Matrix code ECC version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidatamatrixcodeeccversion?language=objc
type DataMatrixCodeECCVersion int

const (
	DataMatrixCodeECCVersion000 DataMatrixCodeECCVersion = 0
	DataMatrixCodeECCVersion050 DataMatrixCodeECCVersion = 50
	DataMatrixCodeECCVersion080 DataMatrixCodeECCVersion = 80
	DataMatrixCodeECCVersion100 DataMatrixCodeECCVersion = 100
	DataMatrixCodeECCVersion140 DataMatrixCodeECCVersion = 140
	DataMatrixCodeECCVersion200 DataMatrixCodeECCVersion = 200
)

// Pixel data formats for image input, output, and processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciformat?language=objc
type Format int

const (
	kCIFormatA16    Format = 1793
	kCIFormatA8     Format = 257
	kCIFormatABGR8  Format = 267
	kCIFormatARGB8  Format = 265
	kCIFormatAf     Format = 2305
	kCIFormatAh     Format = 2049
	kCIFormatBGRA8  Format = 266
	kCIFormatL16    Format = 1795
	kCIFormatL8     Format = 259
	kCIFormatLA16   Format = 1796
	kCIFormatLA8    Format = 260
	kCIFormatLAf    Format = 2308
	kCIFormatLAh    Format = 2052
	kCIFormatLf     Format = 2307
	kCIFormatLh     Format = 2051
	kCIFormatR16    Format = 1797
	kCIFormatR8     Format = 261
	kCIFormatRG16   Format = 1798
	kCIFormatRG8    Format = 262
	kCIFormatRGBA16 Format = 1800
	kCIFormatRGBA8  Format = 264
	kCIFormatRGBAf  Format = 2312
	kCIFormatRGBAh  Format = 2056
	kCIFormatRGf    Format = 2310
	kCIFormatRGh    Format = 2054
	kCIFormatRf     Format = 2309
	kCIFormatRh     Format = 2053
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageautoadjustmentoption?language=objc
type ImageAutoAdjustmentOption string

const (
	kCIImageAutoAdjustCrop     ImageAutoAdjustmentOption = "kCIImageAutoAdjustCrop"
	kCIImageAutoAdjustEnhance  ImageAutoAdjustmentOption = "kCIImageAutoAdjustEnhance"
	kCIImageAutoAdjustFeatures ImageAutoAdjustmentOption = "kCIImageAutoAdjustFeatures"
	kCIImageAutoAdjustLevel    ImageAutoAdjustmentOption = "kCIImageAutoAdjustLevel"
	kCIImageAutoAdjustRedEye   ImageAutoAdjustmentOption = "kCIImageAutoAdjustRedEye"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageoption?language=objc
type ImageOption string

const (
	kCIImageApplyOrientationProperty                  ImageOption = "kCIImageApplyOrientationProperty"
	kCIImageAuxiliaryDepth                            ImageOption = "kCIImageAuxiliaryDepth"
	kCIImageAuxiliaryDisparity                        ImageOption = "kCIImageAuxiliaryDisparity"
	kCIImageAuxiliaryPortraitEffectsMatte             ImageOption = "kCIImageAuxiliaryPortraitEffectsMatte"
	kCIImageAuxiliarySemanticSegmentationGlassesMatte ImageOption = "kCIImageAuxiliarySemanticSegmentationGlassesMatte"
	kCIImageAuxiliarySemanticSegmentationHairMatte    ImageOption = "kCIImageAuxiliarySemanticSegmentationHairMatte"
	kCIImageAuxiliarySemanticSegmentationSkinMatte    ImageOption = "kCIImageAuxiliarySemanticSegmentationSkinMatte"
	kCIImageAuxiliarySemanticSegmentationSkyMatte     ImageOption = "kCIImageAuxiliarySemanticSegmentationSkyMatte"
	kCIImageAuxiliarySemanticSegmentationTeethMatte   ImageOption = "kCIImageAuxiliarySemanticSegmentationTeethMatte"
	kCIImageColorSpace                                ImageOption = "CIImageColorSpace"
	kCIImageNearestSampling                           ImageOption = "CIImageNearestSampling"
	kCIImageProperties                                ImageOption = "CIImageProperties"
	kCIImageProviderTileSize                          ImageOption = "tile_size"
	kCIImageProviderUserInfo                          ImageOption = "user_info"
	kCIImageTextureFormat                             ImageOption = "kCIImageTextureFormat"
	kCIImageTextureTarget                             ImageOption = "kCIImageTextureTarget"
	kCIImageToneMapHDRtoSDR                           ImageOption = "kCIImageToneMapHDRtoSDR"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimagerepresentationoption?language=objc
type ImageRepresentationOption string

const (
	kCIImageRepresentationAVDepthData                           ImageRepresentationOption = "kCIImageRepresentationAVDepthData"
	kCIImageRepresentationAVPortraitEffectsMatte                ImageRepresentationOption = "kCIImageRepresentationAVPortraitEffectsMatte"
	kCIImageRepresentationAVSemanticSegmentationMattes          ImageRepresentationOption = "kCIImageRepresentationAVSemanticSegmentationMattes"
	kCIImageRepresentationDepthImage                            ImageRepresentationOption = "kCIImageRepresentationDepthImage"
	kCIImageRepresentationDisparityImage                        ImageRepresentationOption = "kCIImageRepresentationDisparityImage"
	kCIImageRepresentationPortraitEffectsMatteImage             ImageRepresentationOption = "kCIImageRepresentationPortraitEffectsMatteImage"
	kCIImageRepresentationSemanticSegmentationGlassesMatteImage ImageRepresentationOption = "kCIImageRepresentationSemanticSegmentationGlassesMatteImage"
	kCIImageRepresentationSemanticSegmentationHairMatteImage    ImageRepresentationOption = "kCIImageRepresentationSemanticSegmentationHairMatteImage"
	kCIImageRepresentationSemanticSegmentationSkinMatteImage    ImageRepresentationOption = "kCIImageRepresentationSemanticSegmentationSkinMatteImage"
	kCIImageRepresentationSemanticSegmentationSkyMatteImage     ImageRepresentationOption = "kCIImageRepresentationSemanticSegmentationSkyMatteImage"
	kCIImageRepresentationSemanticSegmentationTeethMatteImage   ImageRepresentationOption = "kCIImageRepresentationSemanticSegmentationTeethMatteImage"
)

// Constants that indicate the percentage of the symbol dedicated to error correction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodeerrorcorrectionlevel?language=objc
type QRCodeErrorCorrectionLevel int

const (
	QRCodeErrorCorrectionLevelH QRCodeErrorCorrectionLevel = 72
	QRCodeErrorCorrectionLevelL QRCodeErrorCorrectionLevel = 76
	QRCodeErrorCorrectionLevelM QRCodeErrorCorrectionLevel = 77
	QRCodeErrorCorrectionLevelQ QRCodeErrorCorrectionLevel = 81
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawdecoderversion?language=objc
type RAWDecoderVersion string

const (
	RAWDecoderVersion6    RAWDecoderVersion = "8"
	RAWDecoderVersion6DNG RAWDecoderVersion = "8.dng"
	RAWDecoderVersion7    RAWDecoderVersion = "8"
	RAWDecoderVersion7DNG RAWDecoderVersion = "8.dng"
	RAWDecoderVersion8    RAWDecoderVersion = "8"
	RAWDecoderVersion8DNG RAWDecoderVersion = "8.dng"
	RAWDecoderVersionNone RAWDecoderVersion = "None"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilteroption?language=objc
type RAWFilterOption string

const (
	kCIActiveKeys                            RAWFilterOption = "activeKeys"
	kCIInputAllowDraftModeKey                RAWFilterOption = "inputDraftMode"
	kCIInputBaselineExposureKey              RAWFilterOption = "inputBaselineExposure"
	kCIInputBoostKey                         RAWFilterOption = "inputBoost"
	kCIInputBoostShadowAmountKey             RAWFilterOption = "inputBoostShadowAmount"
	kCIInputColorNoiseReductionAmountKey     RAWFilterOption = "inputColorNoiseReductionAmount"
	kCIInputDecoderVersionKey                RAWFilterOption = "inputDecoderVersion"
	kCIInputDisableGamutMapKey               RAWFilterOption = "inputDisableGamutMap"
	kCIInputEnableChromaticNoiseTrackingKey  RAWFilterOption = "inputEnableNoiseTracking"
	kCIInputEnableEDRModeKey                 RAWFilterOption = "inputEnableEDRMode"
	kCIInputEnableSharpeningKey              RAWFilterOption = "inputEnableSharpening"
	kCIInputEnableVendorLensCorrectionKey    RAWFilterOption = "inputEnableVendorLensCorrection"
	kCIInputIgnoreImageOrientationKey        RAWFilterOption = "inputIgnoreOrientation"
	kCIInputImageOrientationKey              RAWFilterOption = "inputImageOrientation"
	kCIInputLinearSpaceFilter                RAWFilterOption = "inputLinearSpaceFilter"
	kCIInputLocalToneMapAmountKey            RAWFilterOption = "inputLocalToneMapAmount"
	kCIInputLuminanceNoiseReductionAmountKey RAWFilterOption = "inputLuminanceNoiseReductionAmount"
	kCIInputMoireAmountKey                   RAWFilterOption = "inputMoireAmount"
	kCIInputNeutralChromaticityXKey          RAWFilterOption = "inputNeutralChromaticityX"
	kCIInputNeutralChromaticityYKey          RAWFilterOption = "inputNeutralChromaticityY"
	kCIInputNeutralLocationKey               RAWFilterOption = "inputNeutralLocation"
	kCIInputNeutralTemperatureKey            RAWFilterOption = "inputNeutralTemperature"
	kCIInputNeutralTintKey                   RAWFilterOption = "inputNeutralTint"
	kCIInputNoiseReductionAmountKey          RAWFilterOption = "inputNoiseReductionAmount"
	kCIInputNoiseReductionContrastAmountKey  RAWFilterOption = "inputNoiseReductionContrastAmount"
	kCIInputNoiseReductionDetailAmountKey    RAWFilterOption = "inputNoiseReductionDetailAmount"
	kCIInputNoiseReductionSharpnessAmountKey RAWFilterOption = "inputNoiseReductionSharpnessAmount"
	kCIInputScaleFactorKey                   RAWFilterOption = "inputScaleFactor"
	kCIOutputNativeSizeKey                   RAWFilterOption = "outputNativeSize"
	kCIPropertiesKey                         RAWFilterOption = "properties"
	kCISupportedDecoderVersionsKey           RAWFilterOption = "supportedDecoderVersions"
)

// Different ways of representing alpha. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestinationalphamode?language=objc
type RenderDestinationAlphaMode uint

const (
	RenderDestinationAlphaNone            RenderDestinationAlphaMode = 0
	RenderDestinationAlphaPremultiplied   RenderDestinationAlphaMode = 1
	RenderDestinationAlphaUnpremultiplied RenderDestinationAlphaMode = 2
)
