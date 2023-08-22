// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetImageGenerator] class.
var AssetImageGeneratorClass = _AssetImageGeneratorClass{objc.GetClass("AVAssetImageGenerator")}

type _AssetImageGeneratorClass struct {
	objc.Class
}

// An interface definition for the [AssetImageGenerator] class.
type IAssetImageGenerator interface {
	objc.IObject
	GenerateCGImagesAsynchronouslyForTimesCompletionHandler(requestedTimes []foundation.IValue, handler AssetImageGeneratorCompletionHandler)
	CancelAllCGImageGeneration()
	MaximumSize() coregraphics.Size
	SetMaximumSize(value coregraphics.Size)
	VideoComposition() VideoComposition
	SetVideoComposition(value IVideoComposition)
	ApertureMode() AssetImageGeneratorApertureMode
	SetApertureMode(value AssetImageGeneratorApertureMode)
	AppliesPreferredTrackTransform() bool
	SetAppliesPreferredTrackTransform(value bool)
	RequestedTimeToleranceAfter() coremedia.Time
	SetRequestedTimeToleranceAfter(value coremedia.Time)
	CustomVideoCompositor() VideoCompositingWrapper
	RequestedTimeToleranceBefore() coremedia.Time
	SetRequestedTimeToleranceBefore(value coremedia.Time)
	Asset() Asset
}

// An object that generates images from a video asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator?language=objc
type AssetImageGenerator struct {
	objc.Object
}

func AssetImageGeneratorFrom(ptr unsafe.Pointer) AssetImageGenerator {
	return AssetImageGenerator{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetImageGeneratorClass) AssetImageGeneratorWithAsset(asset IAsset) AssetImageGenerator {
	rv := objc.Call[AssetImageGenerator](ac, objc.Sel("assetImageGeneratorWithAsset:"), objc.Ptr(asset))
	return rv
}

// Returns a new object that generates images for times within a video asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator/1426634-assetimagegeneratorwithasset?language=objc
func AssetImageGenerator_AssetImageGeneratorWithAsset(asset IAsset) AssetImageGenerator {
	return AssetImageGeneratorClass.AssetImageGeneratorWithAsset(asset)
}

func (a_ AssetImageGenerator) InitWithAsset(asset IAsset) AssetImageGenerator {
	rv := objc.Call[AssetImageGenerator](a_, objc.Sel("initWithAsset:"), objc.Ptr(asset))
	return rv
}

// Creates an object that generates images for times within a video asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator/1387855-initwithasset?language=objc
func NewAssetImageGeneratorWithAsset(asset IAsset) AssetImageGenerator {
	instance := AssetImageGeneratorClass.Alloc().InitWithAsset(asset)
	instance.Autorelease()
	return instance
}

func (ac _AssetImageGeneratorClass) Alloc() AssetImageGenerator {
	rv := objc.Call[AssetImageGenerator](ac, objc.Sel("alloc"))
	return rv
}

func AssetImageGenerator_Alloc() AssetImageGenerator {
	return AssetImageGeneratorClass.Alloc()
}

func (ac _AssetImageGeneratorClass) New() AssetImageGenerator {
	rv := objc.Call[AssetImageGenerator](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetImageGenerator() AssetImageGenerator {
	return AssetImageGeneratorClass.New()
}

func (a_ AssetImageGenerator) Init() AssetImageGenerator {
	rv := objc.Call[AssetImageGenerator](a_, objc.Sel("init"))
	return rv
}

// Generates images asynchronously for an array of requested times, and returns the results in a callback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator/1388100-generatecgimagesasynchronouslyfo?language=objc
func (a_ AssetImageGenerator) GenerateCGImagesAsynchronouslyForTimesCompletionHandler(requestedTimes []foundation.IValue, handler AssetImageGeneratorCompletionHandler) {
	objc.Call[objc.Void](a_, objc.Sel("generateCGImagesAsynchronouslyForTimes:completionHandler:"), requestedTimes, handler)
}

// Cancels all pending image generation requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator/1385859-cancelallcgimagegeneration?language=objc
func (a_ AssetImageGenerator) CancelAllCGImageGeneration() {
	objc.Call[objc.Void](a_, objc.Sel("cancelAllCGImageGeneration"))
}

// The maximum size of images to generate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator/1387560-maximumsize?language=objc
func (a_ AssetImageGenerator) MaximumSize() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](a_, objc.Sel("maximumSize"))
	return rv
}

// The maximum size of images to generate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator/1387560-maximumsize?language=objc
func (a_ AssetImageGenerator) SetMaximumSize(value coregraphics.Size) {
	objc.Call[objc.Void](a_, objc.Sel("setMaximumSize:"), value)
}

// A video composition to use when extracting images from assets with multiple video tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator/1390189-videocomposition?language=objc
func (a_ AssetImageGenerator) VideoComposition() VideoComposition {
	rv := objc.Call[VideoComposition](a_, objc.Sel("videoComposition"))
	return rv
}

// A video composition to use when extracting images from assets with multiple video tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator/1390189-videocomposition?language=objc
func (a_ AssetImageGenerator) SetVideoComposition(value IVideoComposition) {
	objc.Call[objc.Void](a_, objc.Sel("setVideoComposition:"), objc.Ptr(value))
}

// Specifies the aperture mode for the generated image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator/1389314-aperturemode?language=objc
func (a_ AssetImageGenerator) ApertureMode() AssetImageGeneratorApertureMode {
	rv := objc.Call[AssetImageGeneratorApertureMode](a_, objc.Sel("apertureMode"))
	return rv
}

// Specifies the aperture mode for the generated image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator/1389314-aperturemode?language=objc
func (a_ AssetImageGenerator) SetApertureMode(value AssetImageGeneratorApertureMode) {
	objc.Call[objc.Void](a_, objc.Sel("setApertureMode:"), value)
}

// A Boolean value that specifies whether to apply the track matrix or matrices when generating an image from the asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator/1390616-appliespreferredtracktransform?language=objc
func (a_ AssetImageGenerator) AppliesPreferredTrackTransform() bool {
	rv := objc.Call[bool](a_, objc.Sel("appliesPreferredTrackTransform"))
	return rv
}

// A Boolean value that specifies whether to apply the track matrix or matrices when generating an image from the asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator/1390616-appliespreferredtracktransform?language=objc
func (a_ AssetImageGenerator) SetAppliesPreferredTrackTransform(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAppliesPreferredTrackTransform:"), value)
}

// A maximum length of time after the requested time to allow image generation to occur. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator/1387751-requestedtimetoleranceafter?language=objc
func (a_ AssetImageGenerator) RequestedTimeToleranceAfter() coremedia.Time {
	rv := objc.Call[coremedia.Time](a_, objc.Sel("requestedTimeToleranceAfter"))
	return rv
}

// A maximum length of time after the requested time to allow image generation to occur. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator/1387751-requestedtimetoleranceafter?language=objc
func (a_ AssetImageGenerator) SetRequestedTimeToleranceAfter(value coremedia.Time) {
	objc.Call[objc.Void](a_, objc.Sel("setRequestedTimeToleranceAfter:"), value)
}

// A custom video compositor to use when extracting images from assets with multiple video tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator/1386469-customvideocompositor?language=objc
func (a_ AssetImageGenerator) CustomVideoCompositor() VideoCompositingWrapper {
	rv := objc.Call[VideoCompositingWrapper](a_, objc.Sel("customVideoCompositor"))
	return rv
}

// A maximum length of time before the requested time to allow image generation to occur. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator/1390571-requestedtimetolerancebefore?language=objc
func (a_ AssetImageGenerator) RequestedTimeToleranceBefore() coremedia.Time {
	rv := objc.Call[coremedia.Time](a_, objc.Sel("requestedTimeToleranceBefore"))
	return rv
}

// A maximum length of time before the requested time to allow image generation to occur. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator/1390571-requestedtimetolerancebefore?language=objc
func (a_ AssetImageGenerator) SetRequestedTimeToleranceBefore(value coremedia.Time) {
	objc.Call[objc.Void](a_, objc.Sel("setRequestedTimeToleranceBefore:"), value)
}

// The asset that initialized the image generator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegenerator/1390689-asset?language=objc
func (a_ AssetImageGenerator) Asset() Asset {
	rv := objc.Call[Asset](a_, objc.Sel("asset"))
	return rv
}
