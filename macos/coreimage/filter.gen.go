// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Filter] class.
var FilterClass = _FilterClass{objc.GetClass("CIFilter")}

type _FilterClass struct {
	objc.Class
}

// An interface definition for the [Filter] class.
type IFilter interface {
	objc.IObject
	SetDefaults()
	Name() string
	Apply(k IKernel, args ...any) Image
	OutputKeys() []string
	InputKeys() []string
	IsEnabled() bool
	SetEnabled(value bool)
	Attributes() map[string]objc.Object
	OutputImage() Image
}

// An image processor that produces an image by manipulating one or more input images or by generating new image data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter?language=objc
type Filter struct {
	objc.Object
}

func FilterFrom(ptr unsafe.Pointer) Filter {
	return Filter{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FilterClass) Alloc() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("alloc"))
	return rv
}

func Filter_Alloc() Filter {
	return FilterClass.Alloc()
}

func (fc _FilterClass) New() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFilter() Filter {
	return FilterClass.New()
}

func (f_ Filter) Init() Filter {
	rv := objc.Call[Filter](f_, objc.Sel("init"))
	return rv
}

// Returns a star-shine generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228415-starshinegeneratorfilter?language=objc
func (fc _FilterClass) StarShineGeneratorFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("starShineGeneratorFilter"))
	return rv
}

// Returns a star-shine generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228415-starshinegeneratorfilter?language=objc
func Filter_StarShineGeneratorFilter() Filter {
	return FilterClass.StarShineGeneratorFilter()
}

// Returns a saturation blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228400-saturationblendmodefilter?language=objc
func (fc _FilterClass) SaturationBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("saturationBlendModeFilter"))
	return rv
}

// Returns a saturation blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228400-saturationblendmodefilter?language=objc
func Filter_SaturationBlendModeFilter() Filter {
	return FilterClass.SaturationBlendModeFilter()
}

// Returns a sunbeams generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228419-sunbeamsgeneratorfilter?language=objc
func (fc _FilterClass) SunbeamsGeneratorFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("sunbeamsGeneratorFilter"))
	return rv
}

// Returns a sunbeams generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228419-sunbeamsgeneratorfilter?language=objc
func Filter_SunbeamsGeneratorFilter() Filter {
	return FilterClass.SunbeamsGeneratorFilter()
}

// Returns a document enhancer filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228317-documentenhancerfilter?language=objc
func (fc _FilterClass) DocumentEnhancerFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("documentEnhancerFilter"))
	return rv
}

// Returns a document enhancer filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228317-documentenhancerfilter?language=objc
func Filter_DocumentEnhancerFilter() Filter {
	return FilterClass.DocumentEnhancerFilter()
}

// Returns a ripple transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228397-rippletransitionfilter?language=objc
func (fc _FilterClass) RippleTransitionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("rippleTransitionFilter"))
	return rv
}

// Returns a ripple transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228397-rippletransitionfilter?language=objc
func Filter_RippleTransitionFilter() Filter {
	return FilterClass.RippleTransitionFilter()
}

// Publishes a custom filter that is not packaged as an image unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1437889-registerfiltername?language=objc
func (fc _FilterClass) RegisterFilterNameConstructorClassAttributes(name string, anObject PFilterConstructor, attributes map[string]objc.IObject) {
	po1 := objc.WrapAsProtocol("CIFilterConstructor", anObject)
	objc.Call[objc.Void](fc, objc.Sel("registerFilterName:constructor:classAttributes:"), name, po1, attributes)
}

// Publishes a custom filter that is not packaged as an image unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1437889-registerfiltername?language=objc
func Filter_RegisterFilterNameConstructorClassAttributes(name string, anObject PFilterConstructor, attributes map[string]objc.IObject) {
	FilterClass.RegisterFilterNameConstructorClassAttributes(name, anObject, attributes)
}

// Publishes a custom filter that is not packaged as an image unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1437889-registerfiltername?language=objc
func (fc _FilterClass) RegisterFilterNameConstructorObjectClassAttributes(name string, anObjectObject objc.IObject, attributes map[string]objc.IObject) {
	objc.Call[objc.Void](fc, objc.Sel("registerFilterName:constructor:classAttributes:"), name, objc.Ptr(anObjectObject), attributes)
}

// Publishes a custom filter that is not packaged as an image unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1437889-registerfiltername?language=objc
func Filter_RegisterFilterNameConstructorObjectClassAttributes(name string, anObjectObject objc.IObject, attributes map[string]objc.IObject) {
	FilterClass.RegisterFilterNameConstructorObjectClassAttributes(name, anObjectObject, attributes)
}

// Returns the localized name for the specified filter name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1437697-localizednameforfiltername?language=objc
func (fc _FilterClass) LocalizedNameForFilterName(filterName string) string {
	rv := objc.Call[string](fc, objc.Sel("localizedNameForFilterName:"), filterName)
	return rv
}

// Returns the localized name for the specified filter name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1437697-localizednameforfiltername?language=objc
func Filter_LocalizedNameForFilterName(filterName string) string {
	return FilterClass.LocalizedNameForFilterName(filterName)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600150-pinchdistortionfilter?language=objc
func (fc _FilterClass) PinchDistortionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("pinchDistortionFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600150-pinchdistortionfilter?language=objc
func Filter_PinchDistortionFilter() Filter {
	return FilterClass.PinchDistortionFilter()
}

// Returns a parallelogram tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228379-parallelogramtilefilter?language=objc
func (fc _FilterClass) ParallelogramTileFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("parallelogramTileFilter"))
	return rv
}

// Returns a parallelogram tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228379-parallelogramtilefilter?language=objc
func Filter_ParallelogramTileFilter() Filter {
	return FilterClass.ParallelogramTileFilter()
}

// Returns a color cube mixed with mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228289-colorcubesmixedwithmaskfilter?language=objc
func (fc _FilterClass) ColorCubesMixedWithMaskFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorCubesMixedWithMaskFilter"))
	return rv
}

// Returns a color cube mixed with mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228289-colorcubesmixedwithmaskfilter?language=objc
func Filter_ColorCubesMixedWithMaskFilter() Filter {
	return FilterClass.ColorCubesMixedWithMaskFilter()
}

// Returns a kaleidoscope filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228343-kaleidoscopefilter?language=objc
func (fc _FilterClass) KaleidoscopeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("kaleidoscopeFilter"))
	return rv
}

// Returns a kaleidoscope filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228343-kaleidoscopefilter?language=objc
func Filter_KaleidoscopeFilter() Filter {
	return FilterClass.KaleidoscopeFilter()
}

// Returns a lighten blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228346-lightenblendmodefilter?language=objc
func (fc _FilterClass) LightenBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("lightenBlendModeFilter"))
	return rv
}

// Returns a lighten blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228346-lightenblendmodefilter?language=objc
func Filter_LightenBlendModeFilter() Filter {
	return FilterClass.LightenBlendModeFilter()
}

// Returns a fourfold rotated tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228328-fourfoldrotatedtilefilter?language=objc
func (fc _FilterClass) FourfoldRotatedTileFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("fourfoldRotatedTileFilter"))
	return rv
}

// Returns a fourfold rotated tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228328-fourfoldrotatedtilefilter?language=objc
func Filter_FourfoldRotatedTileFilter() Filter {
	return FilterClass.FourfoldRotatedTileFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547111-areaaveragefilter?language=objc
func (fc _FilterClass) AreaAverageFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("areaAverageFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547111-areaaveragefilter?language=objc
func Filter_AreaAverageFilter() Filter {
	return FilterClass.AreaAverageFilter()
}

// Returns an edge-work filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228320-edgeworkfilter?language=objc
func (fc _FilterClass) EdgeWorkFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("edgeWorkFilter"))
	return rv
}

// Returns an edge-work filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228320-edgeworkfilter?language=objc
func Filter_EdgeWorkFilter() Filter {
	return FilterClass.EdgeWorkFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3801605-vividlightblendmodefilter?language=objc
func (fc _FilterClass) VividLightBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("vividLightBlendModeFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3801605-vividlightblendmodefilter?language=objc
func Filter_VividLightBlendModeFilter() Filter {
	return FilterClass.VividLightBlendModeFilter()
}

// Returns a dissolve transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228314-dissolvetransitionfilter?language=objc
func (fc _FilterClass) DissolveTransitionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("dissolveTransitionFilter"))
	return rv
}

// Returns a dissolve transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228314-dissolvetransitionfilter?language=objc
func Filter_DissolveTransitionFilter() Filter {
	return FilterClass.DissolveTransitionFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3750385-convolutionrgb3x3filter?language=objc
func (fc _FilterClass) ConvolutionRGB3X3Filter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("convolutionRGB3X3Filter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3750385-convolutionrgb3x3filter?language=objc
func Filter_ConvolutionRGB3X3Filter() Filter {
	return FilterClass.ConvolutionRGB3X3Filter()
}

// Returns a sixfold reflected tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228405-sixfoldreflectedtilefilter?language=objc
func (fc _FilterClass) SixfoldReflectedTileFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("sixfoldReflectedTileFilter"))
	return rv
}

// Returns a sixfold reflected tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228405-sixfoldreflectedtilefilter?language=objc
func Filter_SixfoldReflectedTileFilter() Filter {
	return FilterClass.SixfoldReflectedTileFilter()
}

// Returns a keystone correction horizontal filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3325510-keystonecorrectionhorizontalfilt?language=objc
func (fc _FilterClass) KeystoneCorrectionHorizontalFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("keystoneCorrectionHorizontalFilter"))
	return rv
}

// Returns a keystone correction horizontal filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3325510-keystonecorrectionhorizontalfilt?language=objc
func Filter_KeystoneCorrectionHorizontalFilter() Filter {
	return FilterClass.KeystoneCorrectionHorizontalFilter()
}

// Returns a convolution 5 x 5 filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228300-convolution5x5filter?language=objc
func (fc _FilterClass) Convolution5X5Filter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("convolution5X5Filter"))
	return rv
}

// Returns a convolution 5 x 5 filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228300-convolution5x5filter?language=objc
func Filter_Convolution5X5Filter() Filter {
	return FilterClass.Convolution5X5Filter()
}

// Blurs a rectangular area by enlarging contrasting pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228367-morphologyrectanglemaximumfilter?language=objc
func (fc _FilterClass) MorphologyRectangleMaximumFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("morphologyRectangleMaximumFilter"))
	return rv
}

// Blurs a rectangular area by enlarging contrasting pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228367-morphologyrectanglemaximumfilter?language=objc
func Filter_MorphologyRectangleMaximumFilter() Filter {
	return FilterClass.MorphologyRectangleMaximumFilter()
}

// Returns a screen blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228401-screenblendmodefilter?language=objc
func (fc _FilterClass) ScreenBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("screenBlendModeFilter"))
	return rv
}

// Returns a screen blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228401-screenblendmodefilter?language=objc
func Filter_ScreenBlendModeFilter() Filter {
	return FilterClass.ScreenBlendModeFilter()
}

// Returns a blend with red mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228275-blendwithredmaskfilter?language=objc
func (fc _FilterClass) BlendWithRedMaskFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("blendWithRedMaskFilter"))
	return rv
}

// Returns a blend with red mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228275-blendwithredmaskfilter?language=objc
func Filter_BlendWithRedMaskFilter() Filter {
	return FilterClass.BlendWithRedMaskFilter()
}

// Calculates the median of an image to refine detail. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228358-medianfilter?language=objc
func (fc _FilterClass) MedianFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("medianFilter"))
	return rv
}

// Calculates the median of an image to refine detail. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228358-medianfilter?language=objc
func Filter_MedianFilter() Filter {
	return FilterClass.MedianFilter()
}

// Returns a color blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228282-colorblendmodefilter?language=objc
func (fc _FilterClass) ColorBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorBlendModeFilter"))
	return rv
}

// Returns a color blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228282-colorblendmodefilter?language=objc
func Filter_ColorBlendModeFilter() Filter {
	return FilterClass.ColorBlendModeFilter()
}

// Applies a bokeh effect to an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228277-bokehblurfilter?language=objc
func (fc _FilterClass) BokehBlurFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("bokehBlurFilter"))
	return rv
}

// Applies a bokeh effect to an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228277-bokehblurfilter?language=objc
func Filter_BokehBlurFilter() Filter {
	return FilterClass.BokehBlurFilter()
}

// Returns a line screen filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228348-linescreenfilter?language=objc
func (fc _FilterClass) LineScreenFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("lineScreenFilter"))
	return rv
}

// Returns a line screen filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228348-linescreenfilter?language=objc
func Filter_LineScreenFilter() Filter {
	return FilterClass.LineScreenFilter()
}

// Returns a Core ML model filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228305-coremlmodelfilter?language=objc
func (fc _FilterClass) CoreMLModelFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("coreMLModelFilter"))
	return rv
}

// Returns a Core ML model filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228305-coremlmodelfilter?language=objc
func Filter_CoreMLModelFilter() Filter {
	return FilterClass.CoreMLModelFilter()
}

// Returns a keystone correction combined filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3325509-keystonecorrectioncombinedfilter?language=objc
func (fc _FilterClass) KeystoneCorrectionCombinedFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("keystoneCorrectionCombinedFilter"))
	return rv
}

// Returns a keystone correction combined filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3325509-keystonecorrectioncombinedfilter?language=objc
func Filter_KeystoneCorrectionCombinedFilter() Filter {
	return FilterClass.KeystoneCorrectionCombinedFilter()
}

// Returns a hue adjust filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228340-hueadjustfilter?language=objc
func (fc _FilterClass) HueAdjustFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("hueAdjustFilter"))
	return rv
}

// Returns a hue adjust filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228340-hueadjustfilter?language=objc
func Filter_HueAdjustFilter() Filter {
	return FilterClass.HueAdjustFilter()
}

// Returns a hue-saturation-value gradient filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228342-huesaturationvaluegradientfilter?language=objc
func (fc _FilterClass) HueSaturationValueGradientFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("hueSaturationValueGradientFilter"))
	return rv
}

// Returns a hue-saturation-value gradient filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228342-huesaturationvaluegradientfilter?language=objc
func Filter_HueSaturationValueGradientFilter() Filter {
	return FilterClass.HueSaturationValueGradientFilter()
}

// Returns a photo-effect mono filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228387-photoeffectmonofilter?language=objc
func (fc _FilterClass) PhotoEffectMonoFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("photoEffectMonoFilter"))
	return rv
}

// Returns a photo-effect mono filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228387-photoeffectmonofilter?language=objc
func Filter_PhotoEffectMonoFilter() Filter {
	return FilterClass.PhotoEffectMonoFilter()
}

// Returns a convolution 9 vertical filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228303-convolution9verticalfilter?language=objc
func (fc _FilterClass) Convolution9VerticalFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("convolution9VerticalFilter"))
	return rv
}

// Returns a convolution 9 vertical filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228303-convolution9verticalfilter?language=objc
func Filter_Convolution9VerticalFilter() Filter {
	return FilterClass.Convolution9VerticalFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547117-areaminimumalphafilter?language=objc
func (fc _FilterClass) AreaMinimumAlphaFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("areaMinimumAlphaFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547117-areaminimumalphafilter?language=objc
func Filter_AreaMinimumAlphaFilter() Filter {
	return FilterClass.AreaMinimumAlphaFilter()
}

// Returns a dither filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228315-ditherfilter?language=objc
func (fc _FilterClass) DitherFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("ditherFilter"))
	return rv
}

// Returns a dither filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228315-ditherfilter?language=objc
func Filter_DitherFilter() Filter {
	return FilterClass.DitherFilter()
}

// Returns a photo-effect transfer filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228391-photoeffecttransferfilter?language=objc
func (fc _FilterClass) PhotoEffectTransferFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("photoEffectTransferFilter"))
	return rv
}

// Returns a photo-effect transfer filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228391-photoeffecttransferfilter?language=objc
func Filter_PhotoEffectTransferFilter() Filter {
	return FilterClass.PhotoEffectTransferFilter()
}

// Returns a stripes generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228417-stripesgeneratorfilter?language=objc
func (fc _FilterClass) StripesGeneratorFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("stripesGeneratorFilter"))
	return rv
}

// Returns a stripes generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228417-stripesgeneratorfilter?language=objc
func Filter_StripesGeneratorFilter() Filter {
	return FilterClass.StripesGeneratorFilter()
}

// Returns a white-point adjust filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228432-whitepointadjustfilter?language=objc
func (fc _FilterClass) WhitePointAdjustFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("whitePointAdjustFilter"))
	return rv
}

// Returns a white-point adjust filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228432-whitepointadjustfilter?language=objc
func Filter_WhitePointAdjustFilter() Filter {
	return FilterClass.WhitePointAdjustFilter()
}

// Returns a Code 128 barcode generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228281-code128barcodegeneratorfilter?language=objc
func (fc _FilterClass) Code128BarcodeGeneratorFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("code128BarcodeGeneratorFilter"))
	return rv
}

// Returns a Code 128 barcode generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228281-code128barcodegeneratorfilter?language=objc
func Filter_Code128BarcodeGeneratorFilter() Filter {
	return FilterClass.Code128BarcodeGeneratorFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600143-drostefilter?language=objc
func (fc _FilterClass) DrosteFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("drosteFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600143-drostefilter?language=objc
func Filter_DrosteFilter() Filter {
	return FilterClass.DrosteFilter()
}

// Returns a circular screen filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228280-circularscreenfilter?language=objc
func (fc _FilterClass) CircularScreenFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("circularScreenFilter"))
	return rv
}

// Returns a circular screen filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228280-circularscreenfilter?language=objc
func Filter_CircularScreenFilter() Filter {
	return FilterClass.CircularScreenFilter()
}

// Returns the location of the localized reference documentation that describes the filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1437642-localizedreferencedocumentationf?language=objc
func (fc _FilterClass) LocalizedReferenceDocumentationForFilterName(filterName string) foundation.URL {
	rv := objc.Call[foundation.URL](fc, objc.Sel("localizedReferenceDocumentationForFilterName:"), filterName)
	return rv
}

// Returns the location of the localized reference documentation that describes the filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1437642-localizedreferencedocumentationf?language=objc
func Filter_LocalizedReferenceDocumentationForFilterName(filterName string) foundation.URL {
	return FilterClass.LocalizedReferenceDocumentationForFilterName(filterName)
}

// Returns a gamma adjust filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228330-gammaadjustfilter?language=objc
func (fc _FilterClass) GammaAdjustFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("gammaAdjustFilter"))
	return rv
}

// Returns a gamma adjust filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228330-gammaadjustfilter?language=objc
func Filter_GammaAdjustFilter() Filter {
	return FilterClass.GammaAdjustFilter()
}

// Returns a perspective tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228381-perspectivetilefilter?language=objc
func (fc _FilterClass) PerspectiveTileFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("perspectiveTileFilter"))
	return rv
}

// Returns a perspective tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228381-perspectivetilefilter?language=objc
func Filter_PerspectiveTileFilter() Filter {
	return FilterClass.PerspectiveTileFilter()
}

// Returns the localized description of a filter for display in the user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1437591-localizeddescriptionforfilternam?language=objc
func (fc _FilterClass) LocalizedDescriptionForFilterName(filterName string) string {
	rv := objc.Call[string](fc, objc.Sel("localizedDescriptionForFilterName:"), filterName)
	return rv
}

// Returns the localized description of a filter for display in the user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1437591-localizeddescriptionforfilternam?language=objc
func Filter_LocalizedDescriptionForFilterName(filterName string) string {
	return FilterClass.LocalizedDescriptionForFilterName(filterName)
}

// Returns a comic effect filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228298-comiceffectfilter?language=objc
func (fc _FilterClass) ComicEffectFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("comicEffectFilter"))
	return rv
}

// Returns a comic effect filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228298-comiceffectfilter?language=objc
func Filter_ComicEffectFilter() Filter {
	return FilterClass.ComicEffectFilter()
}

// Returns a vibrance filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228429-vibrancefilter?language=objc
func (fc _FilterClass) VibranceFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("vibranceFilter"))
	return rv
}

// Returns a vibrance filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228429-vibrancefilter?language=objc
func Filter_VibranceFilter() Filter {
	return FilterClass.VibranceFilter()
}

// Returns a photo-effect instant filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228386-photoeffectinstantfilter?language=objc
func (fc _FilterClass) PhotoEffectInstantFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("photoEffectInstantFilter"))
	return rv
}

// Returns a photo-effect instant filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228386-photoeffectinstantfilter?language=objc
func Filter_PhotoEffectInstantFilter() Filter {
	return FilterClass.PhotoEffectInstantFilter()
}

// Returns a subtract blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228418-subtractblendmodefilter?language=objc
func (fc _FilterClass) SubtractBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("subtractBlendModeFilter"))
	return rv
}

// Returns a subtract blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228418-subtractblendmodefilter?language=objc
func Filter_SubtractBlendModeFilter() Filter {
	return FilterClass.SubtractBlendModeFilter()
}

// Returns a hard-light blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228335-hardlightblendmodefilter?language=objc
func (fc _FilterClass) HardLightBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("hardLightBlendModeFilter"))
	return rv
}

// Returns a hard-light blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228335-hardlightblendmodefilter?language=objc
func Filter_HardLightBlendModeFilter() Filter {
	return FilterClass.HardLightBlendModeFilter()
}

// Returns a mesh generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228359-meshgeneratorfilter?language=objc
func (fc _FilterClass) MeshGeneratorFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("meshGeneratorFilter"))
	return rv
}

// Returns a mesh generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228359-meshgeneratorfilter?language=objc
func Filter_MeshGeneratorFilter() Filter {
	return FilterClass.MeshGeneratorFilter()
}

// Returns a darken blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228307-darkenblendmodefilter?language=objc
func (fc _FilterClass) DarkenBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("darkenBlendModeFilter"))
	return rv
}

// Returns a darken blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228307-darkenblendmodefilter?language=objc
func Filter_DarkenBlendModeFilter() Filter {
	return FilterClass.DarkenBlendModeFilter()
}

// Returns an accordion fold transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228263-accordionfoldtransitionfilter?language=objc
func (fc _FilterClass) AccordionFoldTransitionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("accordionFoldTransitionFilter"))
	return rv
}

// Returns an accordion fold transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228263-accordionfoldtransitionfilter?language=objc
func Filter_AccordionFoldTransitionFilter() Filter {
	return FilterClass.AccordionFoldTransitionFilter()
}

// Returns a lenticular halo generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228345-lenticularhalogeneratorfilter?language=objc
func (fc _FilterClass) LenticularHaloGeneratorFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("lenticularHaloGeneratorFilter"))
	return rv
}

// Returns a lenticular halo generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228345-lenticularhalogeneratorfilter?language=objc
func Filter_LenticularHaloGeneratorFilter() Filter {
	return FilterClass.LenticularHaloGeneratorFilter()
}

// Returns a color cross-polynomial filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228286-colorcrosspolynomialfilter?language=objc
func (fc _FilterClass) ColorCrossPolynomialFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorCrossPolynomialFilter"))
	return rv
}

// Returns a color cross-polynomial filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228286-colorcrosspolynomialfilter?language=objc
func Filter_ColorCrossPolynomialFilter() Filter {
	return FilterClass.ColorCrossPolynomialFilter()
}

// Returns a soft-light blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228408-softlightblendmodefilter?language=objc
func (fc _FilterClass) SoftLightBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("softLightBlendModeFilter"))
	return rv
}

// Returns a soft-light blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228408-softlightblendmodefilter?language=objc
func Filter_SoftLightBlendModeFilter() Filter {
	return FilterClass.SoftLightBlendModeFilter()
}

// Returns a bloom filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228276-bloomfilter?language=objc
func (fc _FilterClass) BloomFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("bloomFilter"))
	return rv
}

// Returns a bloom filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228276-bloomfilter?language=objc
func Filter_BloomFilter() Filter {
	return FilterClass.BloomFilter()
}

// Returns a barcode generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228269-barcodegeneratorfilter?language=objc
func (fc _FilterClass) BarcodeGeneratorFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("barcodeGeneratorFilter"))
	return rv
}

// Returns a barcode generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228269-barcodegeneratorfilter?language=objc
func Filter_BarcodeGeneratorFilter() Filter {
	return FilterClass.BarcodeGeneratorFilter()
}

// Returns a PDF417 barcode generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228261-pdf417barcodegenerator?language=objc
func (fc _FilterClass) PDF417BarcodeGenerator() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("PDF417BarcodeGenerator"))
	return rv
}

// Returns a PDF417 barcode generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228261-pdf417barcodegenerator?language=objc
func Filter_PDF417BarcodeGenerator() Filter {
	return FilterClass.PDF417BarcodeGenerator()
}

// Returns a photo-effect noir filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228388-photoeffectnoirfilter?language=objc
func (fc _FilterClass) PhotoEffectNoirFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("photoEffectNoirFilter"))
	return rv
}

// Returns a photo-effect noir filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228388-photoeffectnoirfilter?language=objc
func Filter_PhotoEffectNoirFilter() Filter {
	return FilterClass.PhotoEffectNoirFilter()
}

// Returns a copy machine transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228304-copymachinetransitionfilter?language=objc
func (fc _FilterClass) CopyMachineTransitionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("copyMachineTransitionFilter"))
	rv.Autorelease()
	return rv
}

// Returns a copy machine transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228304-copymachinetransitionfilter?language=objc
func Filter_CopyMachineTransitionFilter() Filter {
	return FilterClass.CopyMachineTransitionFilter()
}

// Returns a hatched screen filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228336-hatchedscreenfilter?language=objc
func (fc _FilterClass) HatchedScreenFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("hatchedScreenFilter"))
	return rv
}

// Returns a hatched screen filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228336-hatchedscreenfilter?language=objc
func Filter_HatchedScreenFilter() Filter {
	return FilterClass.HatchedScreenFilter()
}

// Returns a Aztec code generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228268-azteccodegeneratorfilter?language=objc
func (fc _FilterClass) AztecCodeGeneratorFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("aztecCodeGeneratorFilter"))
	return rv
}

// Returns a Aztec code generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228268-azteccodegeneratorfilter?language=objc
func Filter_AztecCodeGeneratorFilter() Filter {
	return FilterClass.AztecCodeGeneratorFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600154-vortexdistortionfilter?language=objc
func (fc _FilterClass) VortexDistortionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("vortexDistortionFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600154-vortexdistortionfilter?language=objc
func Filter_VortexDistortionFilter() Filter {
	return FilterClass.VortexDistortionFilter()
}

// Returns a saliency map filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228399-saliencymapfilter?language=objc
func (fc _FilterClass) SaliencyMapFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("saliencyMapFilter"))
	return rv
}

// Returns a saliency map filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228399-saliencymapfilter?language=objc
func Filter_SaliencyMapFilter() Filter {
	return FilterClass.SaliencyMapFilter()
}

// Returns a photo-effect fade filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228385-photoeffectfadefilter?language=objc
func (fc _FilterClass) PhotoEffectFadeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("photoEffectFadeFilter"))
	return rv
}

// Returns a photo-effect fade filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228385-photoeffectfadefilter?language=objc
func Filter_PhotoEffectFadeFilter() Filter {
	return FilterClass.PhotoEffectFadeFilter()
}

// Returns a sharpen luminance filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228404-sharpenluminancefilter?language=objc
func (fc _FilterClass) SharpenLuminanceFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("sharpenLuminanceFilter"))
	return rv
}

// Returns a sharpen luminance filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228404-sharpenluminancefilter?language=objc
func Filter_SharpenLuminanceFilter() Filter {
	return FilterClass.SharpenLuminanceFilter()
}

// Returns a triangle kaleidoscope filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228425-trianglekaleidoscopefilter?language=objc
func (fc _FilterClass) TriangleKaleidoscopeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("triangleKaleidoscopeFilter"))
	return rv
}

// Returns a triangle kaleidoscope filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228425-trianglekaleidoscopefilter?language=objc
func Filter_TriangleKaleidoscopeFilter() Filter {
	return FilterClass.TriangleKaleidoscopeFilter()
}

// Returns an addition compositing filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228264-additioncompositingfilter?language=objc
func (fc _FilterClass) AdditionCompositingFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("additionCompositingFilter"))
	return rv
}

// Returns an addition compositing filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228264-additioncompositingfilter?language=objc
func Filter_AdditionCompositingFilter() Filter {
	return FilterClass.AdditionCompositingFilter()
}

// Returns a keystone correction vertical filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3325511-keystonecorrectionverticalfilter?language=objc
func (fc _FilterClass) KeystoneCorrectionVerticalFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("keystoneCorrectionVerticalFilter"))
	return rv
}

// Returns a keystone correction vertical filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3325511-keystonecorrectionverticalfilter?language=objc
func Filter_KeystoneCorrectionVerticalFilter() Filter {
	return FilterClass.KeystoneCorrectionVerticalFilter()
}

// Returns a Lab Delta E filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228260-labdeltae?language=objc
func (fc _FilterClass) LabDeltaE() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("LabDeltaE"))
	return rv
}

// Returns a Lab Delta E filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228260-labdeltae?language=objc
func Filter_LabDeltaE() Filter {
	return FilterClass.LabDeltaE()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547121-columnaveragefilter?language=objc
func (fc _FilterClass) ColumnAverageFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("columnAverageFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547121-columnaveragefilter?language=objc
func Filter_ColumnAverageFilter() Filter {
	return FilterClass.ColumnAverageFilter()
}

// Creates a zoom blur centered around a single point on the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228434-zoomblurfilter?language=objc
func (fc _FilterClass) ZoomBlurFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("zoomBlurFilter"))
	return rv
}

// Creates a zoom blur centered around a single point on the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228434-zoomblurfilter?language=objc
func Filter_ZoomBlurFilter() Filter {
	return FilterClass.ZoomBlurFilter()
}

// Returns a maximum compositing filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228357-maximumcompositingfilter?language=objc
func (fc _FilterClass) MaximumCompositingFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("maximumCompositingFilter"))
	return rv
}

// Returns a maximum compositing filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228357-maximumcompositingfilter?language=objc
func Filter_MaximumCompositingFilter() Filter {
	return FilterClass.MaximumCompositingFilter()
}

// Returns a false color filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228325-falsecolorfilter?language=objc
func (fc _FilterClass) FalseColorFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("falseColorFilter"))
	return rv
}

// Returns a false color filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228325-falsecolorfilter?language=objc
func Filter_FalseColorFilter() Filter {
	return FilterClass.FalseColorFilter()
}

// Returns an unsharp mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228428-unsharpmaskfilter?language=objc
func (fc _FilterClass) UnsharpMaskFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("unsharpMaskFilter"))
	return rv
}

// Returns an unsharp mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228428-unsharpmaskfilter?language=objc
func Filter_UnsharpMaskFilter() Filter {
	return FilterClass.UnsharpMaskFilter()
}

// Returns a gloom filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228334-gloomfilter?language=objc
func (fc _FilterClass) GloomFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("gloomFilter"))
	return rv
}

// Returns a gloom filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228334-gloomfilter?language=objc
func Filter_GloomFilter() Filter {
	return FilterClass.GloomFilter()
}

// Blurs an image with a Gaussian distribution pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228331-gaussianblurfilter?language=objc
func (fc _FilterClass) GaussianBlurFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("gaussianBlurFilter"))
	return rv
}

// Blurs an image with a Gaussian distribution pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228331-gaussianblurfilter?language=objc
func Filter_GaussianBlurFilter() Filter {
	return FilterClass.GaussianBlurFilter()
}

// Returns a perspective rotate filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3325512-perspectiverotatefilter?language=objc
func (fc _FilterClass) PerspectiveRotateFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("perspectiveRotateFilter"))
	return rv
}

// Returns a perspective rotate filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3325512-perspectiverotatefilter?language=objc
func Filter_PerspectiveRotateFilter() Filter {
	return FilterClass.PerspectiveRotateFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3750387-convolutionrgb7x7filter?language=objc
func (fc _FilterClass) ConvolutionRGB7X7Filter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("convolutionRGB7X7Filter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3750387-convolutionrgb7x7filter?language=objc
func Filter_ConvolutionRGB7X7Filter() Filter {
	return FilterClass.ConvolutionRGB7X7Filter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600146-holedistortionfilter?language=objc
func (fc _FilterClass) HoleDistortionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("holeDistortionFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600146-holedistortionfilter?language=objc
func Filter_HoleDistortionFilter() Filter {
	return FilterClass.HoleDistortionFilter()
}

// Returns a checkerboard generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228279-checkerboardgeneratorfilter?language=objc
func (fc _FilterClass) CheckerboardGeneratorFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("checkerboardGeneratorFilter"))
	return rv
}

// Returns a checkerboard generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228279-checkerboardgeneratorfilter?language=objc
func Filter_CheckerboardGeneratorFilter() Filter {
	return FilterClass.CheckerboardGeneratorFilter()
}

// Detects and highlights edges of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228364-morphologygradientfilter?language=objc
func (fc _FilterClass) MorphologyGradientFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("morphologyGradientFilter"))
	return rv
}

// Detects and highlights edges of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228364-morphologygradientfilter?language=objc
func Filter_MorphologyGradientFilter() Filter {
	return FilterClass.MorphologyGradientFilter()
}

// Returns a mix filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228362-mixfilter?language=objc
func (fc _FilterClass) MixFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("mixFilter"))
	return rv
}

// Returns a mix filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228362-mixfilter?language=objc
func Filter_MixFilter() Filter {
	return FilterClass.MixFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600153-twirldistortionfilter?language=objc
func (fc _FilterClass) TwirlDistortionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("twirlDistortionFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600153-twirldistortionfilter?language=objc
func Filter_TwirlDistortionFilter() Filter {
	return FilterClass.TwirlDistortionFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600149-nineparttiledfilter?language=objc
func (fc _FilterClass) NinePartTiledFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("ninePartTiledFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600149-nineparttiledfilter?language=objc
func Filter_NinePartTiledFilter() Filter {
	return FilterClass.NinePartTiledFilter()
}

// Returns a disparity-to-depth filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228313-disparitytodepthfilter?language=objc
func (fc _FilterClass) DisparityToDepthFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("disparityToDepthFilter"))
	return rv
}

// Returns a disparity-to-depth filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228313-disparitytodepthfilter?language=objc
func Filter_DisparityToDepthFilter() Filter {
	return FilterClass.DisparityToDepthFilter()
}

// Returns a palette centroid filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228377-palettecentroidfilter?language=objc
func (fc _FilterClass) PaletteCentroidFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("paletteCentroidFilter"))
	return rv
}

// Returns a palette centroid filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228377-palettecentroidfilter?language=objc
func Filter_PaletteCentroidFilter() Filter {
	return FilterClass.PaletteCentroidFilter()
}

// Returns a random generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228396-randomgeneratorfilter?language=objc
func (fc _FilterClass) RandomGeneratorFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("randomGeneratorFilter"))
	return rv
}

// Returns a random generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228396-randomgeneratorfilter?language=objc
func Filter_RandomGeneratorFilter() Filter {
	return FilterClass.RandomGeneratorFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600141-circularwrapfilter?language=objc
func (fc _FilterClass) CircularWrapFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("circularWrapFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600141-circularwrapfilter?language=objc
func Filter_CircularWrapFilter() Filter {
	return FilterClass.CircularWrapFilter()
}

// Returns a bicubic scale transform filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228271-bicubicscaletransformfilter?language=objc
func (fc _FilterClass) BicubicScaleTransformFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("bicubicScaleTransformFilter"))
	return rv
}

// Returns a bicubic scale transform filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228271-bicubicscaletransformfilter?language=objc
func Filter_BicubicScaleTransformFilter() Filter {
	return FilterClass.BicubicScaleTransformFilter()
}

// Returns a crystalize filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228306-crystallizefilter?language=objc
func (fc _FilterClass) CrystallizeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("crystallizeFilter"))
	return rv
}

// Returns a crystalize filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228306-crystallizefilter?language=objc
func Filter_CrystallizeFilter() Filter {
	return FilterClass.CrystallizeFilter()
}

// Returns a fourfold reflected tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228327-fourfoldreflectedtilefilter?language=objc
func (fc _FilterClass) FourfoldReflectedTileFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("fourfoldReflectedTileFilter"))
	return rv
}

// Returns a fourfold reflected tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228327-fourfoldreflectedtilefilter?language=objc
func Filter_FourfoldReflectedTileFilter() Filter {
	return FilterClass.FourfoldReflectedTileFilter()
}

// Returns a triangle tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228426-triangletilefilter?language=objc
func (fc _FilterClass) TriangleTileFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("triangleTileFilter"))
	return rv
}

// Returns a triangle tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228426-triangletilefilter?language=objc
func Filter_TriangleTileFilter() Filter {
	return FilterClass.TriangleTileFilter()
}

// Returns a fourfold translated tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228329-fourfoldtranslatedtilefilter?language=objc
func (fc _FilterClass) FourfoldTranslatedTileFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("fourfoldTranslatedTileFilter"))
	return rv
}

// Returns a fourfold translated tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228329-fourfoldtranslatedtilefilter?language=objc
func Filter_FourfoldTranslatedTileFilter() Filter {
	return FilterClass.FourfoldTranslatedTileFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3801604-linearlightblendmodefilter?language=objc
func (fc _FilterClass) LinearLightBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("linearLightBlendModeFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3801604-linearlightblendmodefilter?language=objc
func Filter_LinearLightBlendModeFilter() Filter {
	return FilterClass.LinearLightBlendModeFilter()
}

// Returns a flash transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228326-flashtransitionfilter?language=objc
func (fc _FilterClass) FlashTransitionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("flashTransitionFilter"))
	return rv
}

// Returns a flash transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228326-flashtransitionfilter?language=objc
func Filter_FlashTransitionFilter() Filter {
	return FilterClass.FlashTransitionFilter()
}

// Returns a color invert filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228292-colorinvertfilter?language=objc
func (fc _FilterClass) ColorInvertFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorInvertFilter"))
	return rv
}

// Returns a color invert filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228292-colorinvertfilter?language=objc
func Filter_ColorInvertFilter() Filter {
	return FilterClass.ColorInvertFilter()
}

// Returns a color cube with color space filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228288-colorcubewithcolorspacefilter?language=objc
func (fc _FilterClass) ColorCubeWithColorSpaceFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorCubeWithColorSpaceFilter"))
	return rv
}

// Returns a color cube with color space filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228288-colorcubewithcolorspacefilter?language=objc
func Filter_ColorCubeWithColorSpaceFilter() Filter {
	return FilterClass.ColorCubeWithColorSpaceFilter()
}

// Sets all input values for a filter to default values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1437902-setdefaults?language=objc
func (f_ Filter) SetDefaults() {
	objc.Call[objc.Void](f_, objc.Sel("setDefaults"))
}

// Returns a color polynomial filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228296-colorpolynomialfilter?language=objc
func (fc _FilterClass) ColorPolynomialFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorPolynomialFilter"))
	return rv
}

// Returns a color polynomial filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228296-colorpolynomialfilter?language=objc
func Filter_ColorPolynomialFilter() Filter {
	return FilterClass.ColorPolynomialFilter()
}

// Returns a divide blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228316-divideblendmodefilter?language=objc
func (fc _FilterClass) DivideBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("divideBlendModeFilter"))
	return rv
}

// Returns a divide blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228316-divideblendmodefilter?language=objc
func Filter_DivideBlendModeFilter() Filter {
	return FilterClass.DivideBlendModeFilter()
}

// Returns a twelvefold reflected tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228427-twelvefoldreflectedtilefilter?language=objc
func (fc _FilterClass) TwelvefoldReflectedTileFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("twelvefoldReflectedTileFilter"))
	return rv
}

// Returns a twelvefold reflected tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228427-twelvefoldreflectedtilefilter?language=objc
func Filter_TwelvefoldReflectedTileFilter() Filter {
	return FilterClass.TwelvefoldReflectedTileFilter()
}

// Returns a spotlight filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228414-spotlightfilter?language=objc
func (fc _FilterClass) SpotLightFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("spotLightFilter"))
	return rv
}

// Returns a spotlight filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228414-spotlightfilter?language=objc
func Filter_SpotLightFilter() Filter {
	return FilterClass.SpotLightFilter()
}

// Returns a rounded rectangle generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3335007-roundedrectanglegeneratorfilter?language=objc
func (fc _FilterClass) RoundedRectangleGeneratorFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("roundedRectangleGeneratorFilter"))
	return rv
}

// Returns a rounded rectangle generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3335007-roundedrectanglegeneratorfilter?language=objc
func Filter_RoundedRectangleGeneratorFilter() Filter {
	return FilterClass.RoundedRectangleGeneratorFilter()
}

// Returns a perspective transform filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228382-perspectivetransformfilter?language=objc
func (fc _FilterClass) PerspectiveTransformFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("perspectiveTransformFilter"))
	return rv
}

// Returns a perspective transform filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228382-perspectivetransformfilter?language=objc
func Filter_PerspectiveTransformFilter() Filter {
	return FilterClass.PerspectiveTransformFilter()
}

// Returns a perspective correction filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228380-perspectivecorrectionfilter?language=objc
func (fc _FilterClass) PerspectiveCorrectionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("perspectiveCorrectionFilter"))
	return rv
}

// Returns a perspective correction filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228380-perspectivecorrectionfilter?language=objc
func Filter_PerspectiveCorrectionFilter() Filter {
	return FilterClass.PerspectiveCorrectionFilter()
}

// A name associated with a filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1437997-setname?language=objc
func (f_ Filter) Name() string {
	rv := objc.Call[string](f_, objc.Sel("name"))
	return rv
}

// Returns a multiply compositing filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228371-multiplycompositingfilter?language=objc
func (fc _FilterClass) MultiplyCompositingFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("multiplyCompositingFilter"))
	return rv
}

// Returns a multiply compositing filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228371-multiplycompositingfilter?language=objc
func Filter_MultiplyCompositingFilter() Filter {
	return FilterClass.MultiplyCompositingFilter()
}

// Returns a hue blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228341-hueblendmodefilter?language=objc
func (fc _FilterClass) HueBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("hueBlendModeFilter"))
	return rv
}

// Returns a hue blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228341-hueblendmodefilter?language=objc
func Filter_HueBlendModeFilter() Filter {
	return FilterClass.HueBlendModeFilter()
}

// Reduces noise by sharpening the edges of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228372-noisereductionfilter?language=objc
func (fc _FilterClass) NoiseReductionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("noiseReductionFilter"))
	return rv
}

// Reduces noise by sharpening the edges of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228372-noisereductionfilter?language=objc
func Filter_NoiseReductionFilter() Filter {
	return FilterClass.NoiseReductionFilter()
}

// Returns a page-curl-with-shadow transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228376-pagecurlwithshadowtransitionfilt?language=objc
func (fc _FilterClass) PageCurlWithShadowTransitionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("pageCurlWithShadowTransitionFilter"))
	return rv
}

// Returns a page-curl-with-shadow transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228376-pagecurlwithshadowtransitionfilt?language=objc
func Filter_PageCurlWithShadowTransitionFilter() Filter {
	return FilterClass.PageCurlWithShadowTransitionFilter()
}

// Returns a color posterize filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228297-colorposterizefilter?language=objc
func (fc _FilterClass) ColorPosterizeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorPosterizeFilter"))
	return rv
}

// Returns a color posterize filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228297-colorposterizefilter?language=objc
func Filter_ColorPosterizeFilter() Filter {
	return FilterClass.ColorPosterizeFilter()
}

// Returns an exclusion blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228323-exclusionblendmodefilter?language=objc
func (fc _FilterClass) ExclusionBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("exclusionBlendModeFilter"))
	return rv
}

// Returns an exclusion blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228323-exclusionblendmodefilter?language=objc
func Filter_ExclusionBlendModeFilter() Filter {
	return FilterClass.ExclusionBlendModeFilter()
}

// Returns a temperature and tint filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228421-temperatureandtintfilter?language=objc
func (fc _FilterClass) TemperatureAndTintFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("temperatureAndTintFilter"))
	return rv
}

// Returns a temperature and tint filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228421-temperatureandtintfilter?language=objc
func Filter_TemperatureAndTintFilter() Filter {
	return FilterClass.TemperatureAndTintFilter()
}

// Blurs a circular area by enlarging contrasting pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228365-morphologymaximumfilter?language=objc
func (fc _FilterClass) MorphologyMaximumFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("morphologyMaximumFilter"))
	return rv
}

// Blurs a circular area by enlarging contrasting pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228365-morphologymaximumfilter?language=objc
func Filter_MorphologyMaximumFilter() Filter {
	return FilterClass.MorphologyMaximumFilter()
}

// Returns a pointillize filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228394-pointillizefilter?language=objc
func (fc _FilterClass) PointillizeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("pointillizeFilter"))
	return rv
}

// Returns a pointillize filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228394-pointillizefilter?language=objc
func Filter_PointillizeFilter() Filter {
	return FilterClass.PointillizeFilter()
}

// Returns a vignette filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228431-vignettefilter?language=objc
func (fc _FilterClass) VignetteFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("vignetteFilter"))
	return rv
}

// Returns a vignette filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228431-vignettefilter?language=objc
func Filter_VignetteFilter() Filter {
	return FilterClass.VignetteFilter()
}

// Returns a Gabor gradients filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3325508-gaborgradientsfilter?language=objc
func (fc _FilterClass) GaborGradientsFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("gaborGradientsFilter"))
	return rv
}

// Returns a Gabor gradients filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3325508-gaborgradientsfilter?language=objc
func Filter_GaborGradientsFilter() Filter {
	return FilterClass.GaborGradientsFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547114-areamaximumfilter?language=objc
func (fc _FilterClass) AreaMaximumFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("areaMaximumFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547114-areamaximumfilter?language=objc
func Filter_AreaMaximumFilter() Filter {
	return FilterClass.AreaMaximumFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547123-rowaveragefilter?language=objc
func (fc _FilterClass) RowAverageFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("rowAverageFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547123-rowaveragefilter?language=objc
func Filter_RowAverageFilter() Filter {
	return FilterClass.RowAverageFilter()
}

// Returns a source-out compositing filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228411-sourceoutcompositingfilter?language=objc
func (fc _FilterClass) SourceOutCompositingFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("sourceOutCompositingFilter"))
	return rv
}

// Returns a source-out compositing filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228411-sourceoutcompositingfilter?language=objc
func Filter_SourceOutCompositingFilter() Filter {
	return FilterClass.SourceOutCompositingFilter()
}

// Returns  the localized name for the specified filter category. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1438057-localizednameforcategory?language=objc
func (fc _FilterClass) LocalizedNameForCategory(category string) string {
	rv := objc.Call[string](fc, objc.Sel("localizedNameForCategory:"), category)
	return rv
}

// Returns  the localized name for the specified filter category. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1438057-localizednameforcategory?language=objc
func Filter_LocalizedNameForCategory(category string) string {
	return FilterClass.LocalizedNameForCategory(category)
}

// Returns a photo-effect chrome filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228384-photoeffectchromefilter?language=objc
func (fc _FilterClass) PhotoEffectChromeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("photoEffectChromeFilter"))
	return rv
}

// Returns a photo-effect chrome filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228384-photoeffectchromefilter?language=objc
func Filter_PhotoEffectChromeFilter() Filter {
	return FilterClass.PhotoEffectChromeFilter()
}

// Returns an edges filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228321-edgesfilter?language=objc
func (fc _FilterClass) EdgesFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("edgesFilter"))
	return rv
}

// Returns an edges filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228321-edgesfilter?language=objc
func Filter_EdgesFilter() Filter {
	return FilterClass.EdgesFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600147-lighttunnelfilter?language=objc
func (fc _FilterClass) LightTunnelFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("lightTunnelFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600147-lighttunnelfilter?language=objc
func Filter_LightTunnelFilter() Filter {
	return FilterClass.LightTunnelFilter()
}

// Returns a glide reflected tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228333-glidereflectedtilefilter?language=objc
func (fc _FilterClass) GlideReflectedTileFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("glideReflectedTileFilter"))
	return rv
}

// Returns a glide reflected tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228333-glidereflectedtilefilter?language=objc
func Filter_GlideReflectedTileFilter() Filter {
	return FilterClass.GlideReflectedTileFilter()
}

// Returns a photo-effect tonal filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228390-photoeffecttonalfilter?language=objc
func (fc _FilterClass) PhotoEffectTonalFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("photoEffectTonalFilter"))
	return rv
}

// Returns a photo-effect tonal filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228390-photoeffecttonalfilter?language=objc
func Filter_PhotoEffectTonalFilter() Filter {
	return FilterClass.PhotoEffectTonalFilter()
}

// Returns an overlay blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228374-overlayblendmodefilter?language=objc
func (fc _FilterClass) OverlayBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("overlayBlendModeFilter"))
	return rv
}

// Returns an overlay blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228374-overlayblendmodefilter?language=objc
func Filter_OverlayBlendModeFilter() Filter {
	return FilterClass.OverlayBlendModeFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600152-toruslensdistortionfilter?language=objc
func (fc _FilterClass) TorusLensDistortionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("torusLensDistortionFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600152-toruslensdistortionfilter?language=objc
func Filter_TorusLensDistortionFilter() Filter {
	return FilterClass.TorusLensDistortionFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547116-areaminmaxredfilter?language=objc
func (fc _FilterClass) AreaMinMaxRedFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("areaMinMaxRedFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547116-areaminmaxredfilter?language=objc
func Filter_AreaMinMaxRedFilter() Filter {
	return FilterClass.AreaMinMaxRedFilter()
}

// Returns an exposure adjust filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228324-exposureadjustfilter?language=objc
func (fc _FilterClass) ExposureAdjustFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("exposureAdjustFilter"))
	return rv
}

// Returns an exposure adjust filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228324-exposureadjustfilter?language=objc
func Filter_ExposureAdjustFilter() Filter {
	return FilterClass.ExposureAdjustFilter()
}

// Returns a linear-to-sRGB filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228352-lineartosrgbtonecurvefilter?language=objc
func (fc _FilterClass) LinearToSRGBToneCurveFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("linearToSRGBToneCurveFilter"))
	return rv
}

// Returns a linear-to-sRGB filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228352-lineartosrgbtonecurvefilter?language=objc
func Filter_LinearToSRGBToneCurveFilter() Filter {
	return FilterClass.LinearToSRGBToneCurveFilter()
}

// Returns a multiply blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228370-multiplyblendmodefilter?language=objc
func (fc _FilterClass) MultiplyBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("multiplyBlendModeFilter"))
	return rv
}

// Returns a multiply blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228370-multiplyblendmodefilter?language=objc
func Filter_MultiplyBlendModeFilter() Filter {
	return FilterClass.MultiplyBlendModeFilter()
}

// Returns a color map filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228293-colormapfilter?language=objc
func (fc _FilterClass) ColorMapFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorMapFilter"))
	return rv
}

// Returns a color map filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228293-colormapfilter?language=objc
func Filter_ColorMapFilter() Filter {
	return FilterClass.ColorMapFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3750388-convolutionrgb9horizontalfilter?language=objc
func (fc _FilterClass) ConvolutionRGB9HorizontalFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("convolutionRGB9HorizontalFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3750388-convolutionrgb9horizontalfilter?language=objc
func Filter_ConvolutionRGB9HorizontalFilter() Filter {
	return FilterClass.ConvolutionRGB9HorizontalFilter()
}

// Returns a color matrix filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228294-colormatrixfilter?language=objc
func (fc _FilterClass) ColorMatrixFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorMatrixFilter"))
	return rv
}

// Returns a color matrix filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228294-colormatrixfilter?language=objc
func Filter_ColorMatrixFilter() Filter {
	return FilterClass.ColorMatrixFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547110-kmeansfilter?language=objc
func (fc _FilterClass) KMeansFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("KMeansFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547110-kmeansfilter?language=objc
func Filter_KMeansFilter() Filter {
	return FilterClass.KMeansFilter()
}

// Applies a circle-shaped blur to an area of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228311-discblurfilter?language=objc
func (fc _FilterClass) DiscBlurFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("discBlurFilter"))
	return rv
}

// Applies a circle-shaped blur to an area of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228311-discblurfilter?language=objc
func Filter_DiscBlurFilter() Filter {
	return FilterClass.DiscBlurFilter()
}

// Returns a color clamp filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228284-colorclampfilter?language=objc
func (fc _FilterClass) ColorClampFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorClampFilter"))
	return rv
}

// Returns a color clamp filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228284-colorclampfilter?language=objc
func Filter_ColorClampFilter() Filter {
	return FilterClass.ColorClampFilter()
}

// Returns a pin-light blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228392-pinlightblendmodefilter?language=objc
func (fc _FilterClass) PinLightBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("pinLightBlendModeFilter"))
	return rv
}

// Returns a pin-light blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228392-pinlightblendmodefilter?language=objc
func Filter_PinLightBlendModeFilter() Filter {
	return FilterClass.PinLightBlendModeFilter()
}

// Returns a photo-effect process filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228389-photoeffectprocessfilter?language=objc
func (fc _FilterClass) PhotoEffectProcessFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("photoEffectProcessFilter"))
	return rv
}

// Returns a photo-effect process filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228389-photoeffectprocessfilter?language=objc
func Filter_PhotoEffectProcessFilter() Filter {
	return FilterClass.PhotoEffectProcessFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547113-areamaximumalphafilter?language=objc
func (fc _FilterClass) AreaMaximumAlphaFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("areaMaximumAlphaFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547113-areamaximumalphafilter?language=objc
func Filter_AreaMaximumAlphaFilter() Filter {
	return FilterClass.AreaMaximumAlphaFilter()
}

// Returns an X-ray filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228433-xrayfilter?language=objc
func (fc _FilterClass) XRayFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("xRayFilter"))
	return rv
}

// Returns an X-ray filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228433-xrayfilter?language=objc
func Filter_XRayFilter() Filter {
	return FilterClass.XRayFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547115-areaminmaxfilter?language=objc
func (fc _FilterClass) AreaMinMaxFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("areaMinMaxFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547115-areaminmaxfilter?language=objc
func Filter_AreaMinMaxFilter() Filter {
	return FilterClass.AreaMinMaxFilter()
}

// Returns a Gaussian gradient filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228332-gaussiangradientfilter?language=objc
func (fc _FilterClass) GaussianGradientFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("gaussianGradientFilter"))
	return rv
}

// Returns a Gaussian gradient filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228332-gaussiangradientfilter?language=objc
func Filter_GaussianGradientFilter() Filter {
	return FilterClass.GaussianGradientFilter()
}

// Returns a highlight-shadow adjust filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228339-highlightshadowadjustfilter?language=objc
func (fc _FilterClass) HighlightShadowAdjustFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("highlightShadowAdjustFilter"))
	return rv
}

// Returns a highlight-shadow adjust filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228339-highlightshadowadjustfilter?language=objc
func Filter_HighlightShadowAdjustFilter() Filter {
	return FilterClass.HighlightShadowAdjustFilter()
}

// Returns a convolution 7 x 7 filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228301-convolution7x7filter?language=objc
func (fc _FilterClass) Convolution7X7Filter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("convolution7X7Filter"))
	return rv
}

// Returns a convolution 7 x 7 filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228301-convolution7x7filter?language=objc
func Filter_Convolution7X7Filter() Filter {
	return FilterClass.Convolution7X7Filter()
}

// Returns a blend with blue mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228273-blendwithbluemaskfilter?language=objc
func (fc _FilterClass) BlendWithBlueMaskFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("blendWithBlueMaskFilter"))
	return rv
}

// Returns a blend with blue mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228273-blendwithbluemaskfilter?language=objc
func Filter_BlendWithBlueMaskFilter() Filter {
	return FilterClass.BlendWithBlueMaskFilter()
}

// Returns a pixellate filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228393-pixellatefilter?language=objc
func (fc _FilterClass) PixellateFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("pixellateFilter"))
	return rv
}

// Returns a pixellate filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228393-pixellatefilter?language=objc
func Filter_PixellateFilter() Filter {
	return FilterClass.PixellateFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547119-colorabsolutedifferencefilter?language=objc
func (fc _FilterClass) ColorAbsoluteDifferenceFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorAbsoluteDifferenceFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547119-colorabsolutedifferencefilter?language=objc
func Filter_ColorAbsoluteDifferenceFilter() Filter {
	return FilterClass.ColorAbsoluteDifferenceFilter()
}

// Returns a depth-to-disparity filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228309-depthtodisparityfilter?language=objc
func (fc _FilterClass) DepthToDisparityFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("depthToDisparityFilter"))
	return rv
}

// Returns a depth-to-disparity filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228309-depthtodisparityfilter?language=objc
func Filter_DepthToDisparityFilter() Filter {
	return FilterClass.DepthToDisparityFilter()
}

// Returns a thermal filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228423-thermalfilter?language=objc
func (fc _FilterClass) ThermalFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("thermalFilter"))
	return rv
}

// Returns a thermal filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228423-thermalfilter?language=objc
func Filter_ThermalFilter() Filter {
	return FilterClass.ThermalFilter()
}

// Produces a CIImage object by applying a kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1562058-apply?language=objc
func (f_ Filter) Apply(k IKernel, args ...any) Image {
	rv := objc.Call[Image](f_, objc.Sel("apply:"), append([]any{objc.Ptr(k)}, args...)...)
	return rv
}

// Returns a sixfold rotated tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228406-sixfoldrotatedtilefilter?language=objc
func (fc _FilterClass) SixfoldRotatedTileFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("sixfoldRotatedTileFilter"))
	return rv
}

// Returns a sixfold rotated tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228406-sixfoldrotatedtilefilter?language=objc
func Filter_SixfoldRotatedTileFilter() Filter {
	return FilterClass.SixfoldRotatedTileFilter()
}

// Returns a line overlay filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228347-lineoverlayfilter?language=objc
func (fc _FilterClass) LineOverlayFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("lineOverlayFilter"))
	return rv
}

// Returns a line overlay filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228347-lineoverlayfilter?language=objc
func Filter_LineOverlayFilter() Filter {
	return FilterClass.LineOverlayFilter()
}

// Returns a color-burn blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228283-colorburnblendmodefilter?language=objc
func (fc _FilterClass) ColorBurnBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorBurnBlendModeFilter"))
	return rv
}

// Returns a color-burn blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228283-colorburnblendmodefilter?language=objc
func Filter_ColorBurnBlendModeFilter() Filter {
	return FilterClass.ColorBurnBlendModeFilter()
}

// Returns a smooth linear gradient filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228407-smoothlineargradientfilter?language=objc
func (fc _FilterClass) SmoothLinearGradientFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("smoothLinearGradientFilter"))
	return rv
}

// Returns a smooth linear gradient filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228407-smoothlineargradientfilter?language=objc
func Filter_SmoothLinearGradientFilter() Filter {
	return FilterClass.SmoothLinearGradientFilter()
}

// Returns a vignette-effect filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228430-vignetteeffectfilter?language=objc
func (fc _FilterClass) VignetteEffectFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("vignetteEffectFilter"))
	return rv
}

// Returns a vignette-effect filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228430-vignetteeffectfilter?language=objc
func Filter_VignetteEffectFilter() Filter {
	return FilterClass.VignetteEffectFilter()
}

// Returns an eightfold reflected tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228322-eightfoldreflectedtilefilter?language=objc
func (fc _FilterClass) EightfoldReflectedTileFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("eightfoldReflectedTileFilter"))
	return rv
}

// Returns an eightfold reflected tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228322-eightfoldreflectedtilefilter?language=objc
func Filter_EightfoldReflectedTileFilter() Filter {
	return FilterClass.EightfoldReflectedTileFilter()
}

// Returns a color-dodge blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228291-colordodgeblendmodefilter?language=objc
func (fc _FilterClass) ColorDodgeBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorDodgeBlendModeFilter"))
	return rv
}

// Returns a color-dodge blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228291-colordodgeblendmodefilter?language=objc
func Filter_ColorDodgeBlendModeFilter() Filter {
	return FilterClass.ColorDodgeBlendModeFilter()
}

// Returns a shaded material filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228403-shadedmaterialfilter?language=objc
func (fc _FilterClass) ShadedMaterialFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("shadedMaterialFilter"))
	return rv
}

// Returns a shaded material filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228403-shadedmaterialfilter?language=objc
func Filter_ShadedMaterialFilter() Filter {
	return FilterClass.ShadedMaterialFilter()
}

// Blurs a rectangular area by reducing contrasting pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228368-morphologyrectangleminimumfilter?language=objc
func (fc _FilterClass) MorphologyRectangleMinimumFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("morphologyRectangleMinimumFilter"))
	return rv
}

// Blurs a rectangular area by reducing contrasting pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228368-morphologyrectangleminimumfilter?language=objc
func Filter_MorphologyRectangleMinimumFilter() Filter {
	return FilterClass.MorphologyRectangleMinimumFilter()
}

// Returns a sepia-tone filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228402-sepiatonefilter?language=objc
func (fc _FilterClass) SepiaToneFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("sepiaToneFilter"))
	return rv
}

// Returns a sepia-tone filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228402-sepiatonefilter?language=objc
func Filter_SepiaToneFilter() Filter {
	return FilterClass.SepiaToneFilter()
}

// Returns a hexagonal pixellate filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228338-hexagonalpixellatefilter?language=objc
func (fc _FilterClass) HexagonalPixellateFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("hexagonalPixellateFilter"))
	return rv
}

// Returns a hexagonal pixellate filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228338-hexagonalpixellatefilter?language=objc
func Filter_HexagonalPixellateFilter() Filter {
	return FilterClass.HexagonalPixellateFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547112-areahistogramfilter?language=objc
func (fc _FilterClass) AreaHistogramFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("areaHistogramFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547112-areahistogramfilter?language=objc
func Filter_AreaHistogramFilter() Filter {
	return FilterClass.AreaHistogramFilter()
}

// Returns a mod transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228363-modtransitionfilter?language=objc
func (fc _FilterClass) ModTransitionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("modTransitionFilter"))
	return rv
}

// Returns a mod transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228363-modtransitionfilter?language=objc
func Filter_ModTransitionFilter() Filter {
	return FilterClass.ModTransitionFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600148-ninepartstretchedfilter?language=objc
func (fc _FilterClass) NinePartStretchedFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("ninePartStretchedFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600148-ninepartstretchedfilter?language=objc
func Filter_NinePartStretchedFilter() Filter {
	return FilterClass.NinePartStretchedFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600138-bumpdistortionfilter?language=objc
func (fc _FilterClass) BumpDistortionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("bumpDistortionFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600138-bumpdistortionfilter?language=objc
func Filter_BumpDistortionFilter() Filter {
	return FilterClass.BumpDistortionFilter()
}

// Returns a convolution 9 horizontal filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228302-convolution9horizontalfilter?language=objc
func (fc _FilterClass) Convolution9HorizontalFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("convolution9HorizontalFilter"))
	return rv
}

// Returns a convolution 9 horizontal filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228302-convolution9horizontalfilter?language=objc
func Filter_Convolution9HorizontalFilter() Filter {
	return FilterClass.Convolution9HorizontalFilter()
}

// Returns a luminosity blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228353-luminosityblendmodefilter?language=objc
func (fc _FilterClass) LuminosityBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("luminosityBlendModeFilter"))
	return rv
}

// Returns a luminosity blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228353-luminosityblendmodefilter?language=objc
func Filter_LuminosityBlendModeFilter() Filter {
	return FilterClass.LuminosityBlendModeFilter()
}

// Returns an optical tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228373-optilefilter?language=objc
func (fc _FilterClass) OpTileFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("opTileFilter"))
	return rv
}

// Returns an optical tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228373-optilefilter?language=objc
func Filter_OpTileFilter() Filter {
	return FilterClass.OpTileFilter()
}

// Returns a minimum compositing filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228361-minimumcompositingfilter?language=objc
func (fc _FilterClass) MinimumCompositingFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("minimumCompositingFilter"))
	return rv
}

// Returns a minimum compositing filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228361-minimumcompositingfilter?language=objc
func Filter_MinimumCompositingFilter() Filter {
	return FilterClass.MinimumCompositingFilter()
}

// Returns a maximum component filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228356-maximumcomponentfilter?language=objc
func (fc _FilterClass) MaximumComponentFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("maximumComponentFilter"))
	return rv
}

// Returns a maximum component filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228356-maximumcomponentfilter?language=objc
func Filter_MaximumComponentFilter() Filter {
	return FilterClass.MaximumComponentFilter()
}

// Returns an affine clamp filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228265-affineclampfilter?language=objc
func (fc _FilterClass) AffineClampFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("affineClampFilter"))
	return rv
}

// Returns an affine clamp filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228265-affineclampfilter?language=objc
func Filter_AffineClampFilter() Filter {
	return FilterClass.AffineClampFilter()
}

// Applies a square-shaped blur to an area of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228278-boxblurfilter?language=objc
func (fc _FilterClass) BoxBlurFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("boxBlurFilter"))
	return rv
}

// Applies a square-shaped blur to an area of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228278-boxblurfilter?language=objc
func Filter_BoxBlurFilter() Filter {
	return FilterClass.BoxBlurFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600140-circlesplashdistortionfilter?language=objc
func (fc _FilterClass) CircleSplashDistortionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("circleSplashDistortionFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600140-circlesplashdistortionfilter?language=objc
func Filter_CircleSplashDistortionFilter() Filter {
	return FilterClass.CircleSplashDistortionFilter()
}

// Returns a Lanczos scale transform filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228344-lanczosscaletransformfilter?language=objc
func (fc _FilterClass) LanczosScaleTransformFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("lanczosScaleTransformFilter"))
	return rv
}

// Returns a Lanczos scale transform filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228344-lanczosscaletransformfilter?language=objc
func Filter_LanczosScaleTransformFilter() Filter {
	return FilterClass.LanczosScaleTransformFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3750390-personsegmentationfilter?language=objc
func (fc _FilterClass) PersonSegmentationFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("personSegmentationFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3750390-personsegmentationfilter?language=objc
func Filter_PersonSegmentationFilter() Filter {
	return FilterClass.PersonSegmentationFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547120-colorthresholdfilter?language=objc
func (fc _FilterClass) ColorThresholdFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorThresholdFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547120-colorthresholdfilter?language=objc
func Filter_ColorThresholdFilter() Filter {
	return FilterClass.ColorThresholdFilter()
}

// Returns a tone curve filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228424-tonecurvefilter?language=objc
func (fc _FilterClass) ToneCurveFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("toneCurveFilter"))
	return rv
}

// Returns a tone curve filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228424-tonecurvefilter?language=objc
func Filter_ToneCurveFilter() Filter {
	return FilterClass.ToneCurveFilter()
}

// Returns a linear-burn blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228349-linearburnblendmodefilter?language=objc
func (fc _FilterClass) LinearBurnBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("linearBurnBlendModeFilter"))
	return rv
}

// Returns a linear-burn blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228349-linearburnblendmodefilter?language=objc
func Filter_LinearBurnBlendModeFilter() Filter {
	return FilterClass.LinearBurnBlendModeFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600139-bumpdistortionlinearfilter?language=objc
func (fc _FilterClass) BumpDistortionLinearFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("bumpDistortionLinearFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600139-bumpdistortionlinearfilter?language=objc
func Filter_BumpDistortionLinearFilter() Filter {
	return FilterClass.BumpDistortionLinearFilter()
}

// Returns a mask-to-alpha filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228354-masktoalphafilter?language=objc
func (fc _FilterClass) MaskToAlphaFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("maskToAlphaFilter"))
	return rv
}

// Returns a mask-to-alpha filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228354-masktoalphafilter?language=objc
func Filter_MaskToAlphaFilter() Filter {
	return FilterClass.MaskToAlphaFilter()
}

// Returns a linear-dodge blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228350-lineardodgeblendmodefilter?language=objc
func (fc _FilterClass) LinearDodgeBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("linearDodgeBlendModeFilter"))
	return rv
}

// Returns a linear-dodge blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228350-lineardodgeblendmodefilter?language=objc
func Filter_LinearDodgeBlendModeFilter() Filter {
	return FilterClass.LinearDodgeBlendModeFilter()
}

// Returns a page curl transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228375-pagecurltransitionfilter?language=objc
func (fc _FilterClass) PageCurlTransitionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("pageCurlTransitionFilter"))
	return rv
}

// Returns a page curl transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228375-pagecurltransitionfilter?language=objc
func Filter_PageCurlTransitionFilter() Filter {
	return FilterClass.PageCurlTransitionFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600151-stretchcropfilter?language=objc
func (fc _FilterClass) StretchCropFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("stretchCropFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600151-stretchcropfilter?language=objc
func Filter_StretchCropFilter() Filter {
	return FilterClass.StretchCropFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547122-histogramdisplayfilter?language=objc
func (fc _FilterClass) HistogramDisplayFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("histogramDisplayFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547122-histogramdisplayfilter?language=objc
func Filter_HistogramDisplayFilter() Filter {
	return FilterClass.HistogramDisplayFilter()
}

// Returns a dot screen filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228318-dotscreenfilter?language=objc
func (fc _FilterClass) DotScreenFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("dotScreenFilter"))
	return rv
}

// Returns a dot screen filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228318-dotscreenfilter?language=objc
func Filter_DotScreenFilter() Filter {
	return FilterClass.DotScreenFilter()
}

// Returns a affine tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228266-affinetilefilter?language=objc
func (fc _FilterClass) AffineTileFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("affineTileFilter"))
	return rv
}

// Returns a affine tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228266-affinetilefilter?language=objc
func Filter_AffineTileFilter() Filter {
	return FilterClass.AffineTileFilter()
}

// Returns a depth-of-field filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228308-depthoffieldfilter?language=objc
func (fc _FilterClass) DepthOfFieldFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("depthOfFieldFilter"))
	return rv
}

// Returns a depth-of-field filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228308-depthoffieldfilter?language=objc
func Filter_DepthOfFieldFilter() Filter {
	return FilterClass.DepthOfFieldFilter()
}

// Returns a difference blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228310-differenceblendmodefilter?language=objc
func (fc _FilterClass) DifferenceBlendModeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("differenceBlendModeFilter"))
	return rv
}

// Returns a difference blend mode filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228310-differenceblendmodefilter?language=objc
func Filter_DifferenceBlendModeFilter() Filter {
	return FilterClass.DifferenceBlendModeFilter()
}

// Returns a bars swipe transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228270-barsswipetransitionfilter?language=objc
func (fc _FilterClass) BarsSwipeTransitionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("barsSwipeTransitionFilter"))
	return rv
}

// Returns a bars swipe transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228270-barsswipetransitionfilter?language=objc
func Filter_BarsSwipeTransitionFilter() Filter {
	return FilterClass.BarsSwipeTransitionFilter()
}

// Returns a radial gradient filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228395-radialgradientfilter?language=objc
func (fc _FilterClass) RadialGradientFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("radialGradientFilter"))
	return rv
}

// Returns a radial gradient filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228395-radialgradientfilter?language=objc
func Filter_RadialGradientFilter() Filter {
	return FilterClass.RadialGradientFilter()
}

// Returns an edge preserve upsample filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228319-edgepreserveupsamplefilter?language=objc
func (fc _FilterClass) EdgePreserveUpsampleFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("edgePreserveUpsampleFilter"))
	return rv
}

// Returns an edge preserve upsample filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228319-edgepreserveupsamplefilter?language=objc
func Filter_EdgePreserveUpsampleFilter() Filter {
	return FilterClass.EdgePreserveUpsampleFilter()
}

// Returns a color cube filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228287-colorcubefilter?language=objc
func (fc _FilterClass) ColorCubeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorCubeFilter"))
	return rv
}

// Returns a color cube filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228287-colorcubefilter?language=objc
func Filter_ColorCubeFilter() Filter {
	return FilterClass.ColorCubeFilter()
}

// Creates motion blur on an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228369-motionblurfilter?language=objc
func (fc _FilterClass) MotionBlurFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("motionBlurFilter"))
	return rv
}

// Creates motion blur on an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228369-motionblurfilter?language=objc
func Filter_MotionBlurFilter() Filter {
	return FilterClass.MotionBlurFilter()
}

// Returns an array of all published filter names that match all the specified categories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1437595-filternamesincategories?language=objc
func (fc _FilterClass) FilterNamesInCategories(categories []string) []string {
	rv := objc.Call[[]string](fc, objc.Sel("filterNamesInCategories:"), categories)
	return rv
}

// Returns an array of all published filter names that match all the specified categories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1437595-filternamesincategories?language=objc
func Filter_FilterNamesInCategories(categories []string) []string {
	return FilterClass.FilterNamesInCategories(categories)
}

// Blurs a specified portion of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228355-maskedvariableblurfilter?language=objc
func (fc _FilterClass) MaskedVariableBlurFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("maskedVariableBlurFilter"))
	return rv
}

// Blurs a specified portion of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228355-maskedvariableblurfilter?language=objc
func Filter_MaskedVariableBlurFilter() Filter {
	return FilterClass.MaskedVariableBlurFilter()
}

// Returns a spot color filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228413-spotcolorfilter?language=objc
func (fc _FilterClass) SpotColorFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("spotColorFilter"))
	return rv
}

// Returns a spot color filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228413-spotcolorfilter?language=objc
func Filter_SpotColorFilter() Filter {
	return FilterClass.SpotColorFilter()
}

// Returns an array of all published filter names in the specified category. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1438145-filternamesincategory?language=objc
func (fc _FilterClass) FilterNamesInCategory(category string) []string {
	rv := objc.Call[[]string](fc, objc.Sel("filterNamesInCategory:"), category)
	return rv
}

// Returns an array of all published filter names in the specified category. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1438145-filternamesincategory?language=objc
func Filter_FilterNamesInCategory(category string) []string {
	return FilterClass.FilterNamesInCategory(category)
}

// Returns a source-atop compositing filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228409-sourceatopcompositingfilter?language=objc
func (fc _FilterClass) SourceAtopCompositingFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("sourceAtopCompositingFilter"))
	return rv
}

// Returns a source-atop compositing filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228409-sourceatopcompositingfilter?language=objc
func Filter_SourceAtopCompositingFilter() Filter {
	return FilterClass.SourceAtopCompositingFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3584770-colorthresholdotsufilter?language=objc
func (fc _FilterClass) ColorThresholdOtsuFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorThresholdOtsuFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3584770-colorthresholdotsufilter?language=objc
func Filter_ColorThresholdOtsuFilter() Filter {
	return FilterClass.ColorThresholdOtsuFilter()
}

// Blurs a circular area by reducing contrasting pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228366-morphologyminimumfilter?language=objc
func (fc _FilterClass) MorphologyMinimumFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("morphologyMinimumFilter"))
	return rv
}

// Blurs a circular area by reducing contrasting pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228366-morphologyminimumfilter?language=objc
func Filter_MorphologyMinimumFilter() Filter {
	return FilterClass.MorphologyMinimumFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547118-areaminimumfilter?language=objc
func (fc _FilterClass) AreaMinimumFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("areaMinimumFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3547118-areaminimumfilter?language=objc
func Filter_AreaMinimumFilter() Filter {
	return FilterClass.AreaMinimumFilter()
}

// Returns a dictionary that contains key-value pairs that describe the filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228047-customattributes?language=objc
func (fc _FilterClass) CustomAttributes() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](fc, objc.Sel("customAttributes"))
	return rv
}

// Returns a dictionary that contains key-value pairs that describe the filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228047-customattributes?language=objc
func Filter_CustomAttributes() map[string]objc.Object {
	return FilterClass.CustomAttributes()
}

// Returns a text image generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228422-textimagegeneratorfilter?language=objc
func (fc _FilterClass) TextImageGeneratorFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("textImageGeneratorFilter"))
	return rv
}

// Returns a text image generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228422-textimagegeneratorfilter?language=objc
func Filter_TextImageGeneratorFilter() Filter {
	return FilterClass.TextImageGeneratorFilter()
}

// Returns an sRGB-to-linear filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228398-srgbtonecurvetolinearfilter?language=objc
func (fc _FilterClass) SRGBToneCurveToLinearFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("sRGBToneCurveToLinearFilter"))
	return rv
}

// Returns an sRGB-to-linear filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228398-srgbtonecurvetolinearfilter?language=objc
func Filter_SRGBToneCurveToLinearFilter() Filter {
	return FilterClass.SRGBToneCurveToLinearFilter()
}

// Returns a source-in compositing filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228410-sourceincompositingfilter?language=objc
func (fc _FilterClass) SourceInCompositingFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("sourceInCompositingFilter"))
	return rv
}

// Returns a source-in compositing filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228410-sourceincompositingfilter?language=objc
func Filter_SourceInCompositingFilter() Filter {
	return FilterClass.SourceInCompositingFilter()
}

// Returns a minimum component filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228360-minimumcomponentfilter?language=objc
func (fc _FilterClass) MinimumComponentFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("minimumComponentFilter"))
	return rv
}

// Returns a minimum component filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228360-minimumcomponentfilter?language=objc
func Filter_MinimumComponentFilter() Filter {
	return FilterClass.MinimumComponentFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3750389-convolutionrgb9verticalfilter?language=objc
func (fc _FilterClass) ConvolutionRGB9VerticalFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("convolutionRGB9VerticalFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3750389-convolutionrgb9verticalfilter?language=objc
func Filter_ConvolutionRGB9VerticalFilter() Filter {
	return FilterClass.ConvolutionRGB9VerticalFilter()
}

// Returns a perspective transform with extent filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228383-perspectivetransformwithextentfi?language=objc
func (fc _FilterClass) PerspectiveTransformWithExtentFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("perspectiveTransformWithExtentFilter"))
	return rv
}

// Returns a perspective transform with extent filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228383-perspectivetransformwithextentfi?language=objc
func Filter_PerspectiveTransformWithExtentFilter() Filter {
	return FilterClass.PerspectiveTransformWithExtentFilter()
}

// Returns a swipe transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228420-swipetransitionfilter?language=objc
func (fc _FilterClass) SwipeTransitionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("swipeTransitionFilter"))
	return rv
}

// Returns a swipe transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228420-swipetransitionfilter?language=objc
func Filter_SwipeTransitionFilter() Filter {
	return FilterClass.SwipeTransitionFilter()
}

// Returns a height-field-from-mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228337-heightfieldfrommaskfilter?language=objc
func (fc _FilterClass) HeightFieldFromMaskFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("heightFieldFromMaskFilter"))
	return rv
}

// Returns a height-field-from-mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228337-heightfieldfrommaskfilter?language=objc
func Filter_HeightFieldFromMaskFilter() Filter {
	return FilterClass.HeightFieldFromMaskFilter()
}

// Returns a color controls filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228285-colorcontrolsfilter?language=objc
func (fc _FilterClass) ColorControlsFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorControlsFilter"))
	return rv
}

// Returns a color controls filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228285-colorcontrolsfilter?language=objc
func Filter_ColorControlsFilter() Filter {
	return FilterClass.ColorControlsFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600142-displacementdistortionfilter?language=objc
func (fc _FilterClass) DisplacementDistortionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("displacementDistortionFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600142-displacementdistortionfilter?language=objc
func Filter_DisplacementDistortionFilter() Filter {
	return FilterClass.DisplacementDistortionFilter()
}

// Returns a color monochrome filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228295-colormonochromefilter?language=objc
func (fc _FilterClass) ColorMonochromeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorMonochromeFilter"))
	return rv
}

// Returns a color monochrome filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228295-colormonochromefilter?language=objc
func Filter_ColorMonochromeFilter() Filter {
	return FilterClass.ColorMonochromeFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600145-glasslozengefilter?language=objc
func (fc _FilterClass) GlassLozengeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("glassLozengeFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600145-glasslozengefilter?language=objc
func Filter_GlassLozengeFilter() Filter {
	return FilterClass.GlassLozengeFilter()
}

// Returns a blend with alpha mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228272-blendwithalphamaskfilter?language=objc
func (fc _FilterClass) BlendWithAlphaMaskFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("blendWithAlphaMaskFilter"))
	return rv
}

// Returns a blend with alpha mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228272-blendwithalphamaskfilter?language=objc
func Filter_BlendWithAlphaMaskFilter() Filter {
	return FilterClass.BlendWithAlphaMaskFilter()
}

// Returns a color curves filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228290-colorcurvesfilter?language=objc
func (fc _FilterClass) ColorCurvesFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("colorCurvesFilter"))
	return rv
}

// Returns a color curves filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228290-colorcurvesfilter?language=objc
func Filter_ColorCurvesFilter() Filter {
	return FilterClass.ColorCurvesFilter()
}

// Returns a source-over compositing filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228412-sourceovercompositingfilter?language=objc
func (fc _FilterClass) SourceOverCompositingFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("sourceOverCompositingFilter"))
	return rv
}

// Returns a source-over compositing filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228412-sourceovercompositingfilter?language=objc
func Filter_SourceOverCompositingFilter() Filter {
	return FilterClass.SourceOverCompositingFilter()
}

// Returns an attributed-text image generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228267-attributedtextimagegeneratorfilt?language=objc
func (fc _FilterClass) AttributedTextImageGeneratorFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("attributedTextImageGeneratorFilter"))
	return rv
}

// Returns an attributed-text image generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228267-attributedtextimagegeneratorfilt?language=objc
func Filter_AttributedTextImageGeneratorFilter() Filter {
	return FilterClass.AttributedTextImageGeneratorFilter()
}

// Returns a disintegrate-with-mask transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228312-disintegratewithmasktransitionfi?language=objc
func (fc _FilterClass) DisintegrateWithMaskTransitionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("disintegrateWithMaskTransitionFilter"))
	return rv
}

// Returns a disintegrate-with-mask transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228312-disintegratewithmasktransitionfi?language=objc
func Filter_DisintegrateWithMaskTransitionFilter() Filter {
	return FilterClass.DisintegrateWithMaskTransitionFilter()
}

// Returns a palettize filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228378-palettizefilter?language=objc
func (fc _FilterClass) PalettizeFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("palettizeFilter"))
	return rv
}

// Returns a palettize filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228378-palettizefilter?language=objc
func Filter_PalettizeFilter() Filter {
	return FilterClass.PalettizeFilter()
}

// Returns a blend with mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228274-blendwithmaskfilter?language=objc
func (fc _FilterClass) BlendWithMaskFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("blendWithMaskFilter"))
	return rv
}

// Returns a blend with mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228274-blendwithmaskfilter?language=objc
func Filter_BlendWithMaskFilter() Filter {
	return FilterClass.BlendWithMaskFilter()
}

// Returns a straighten filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228416-straightenfilter?language=objc
func (fc _FilterClass) StraightenFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("straightenFilter"))
	return rv
}

// Returns a straighten filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228416-straightenfilter?language=objc
func Filter_StraightenFilter() Filter {
	return FilterClass.StraightenFilter()
}

// Returns a CMYK halftone filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228259-cmykhalftone?language=objc
func (fc _FilterClass) CMYKHalftone() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("CMYKHalftone"))
	return rv
}

// Returns a CMYK halftone filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228259-cmykhalftone?language=objc
func Filter_CMYKHalftone() Filter {
	return FilterClass.CMYKHalftone()
}

// Returns a linear gradient filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228351-lineargradientfilter?language=objc
func (fc _FilterClass) LinearGradientFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("linearGradientFilter"))
	return rv
}

// Returns a linear gradient filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228351-lineargradientfilter?language=objc
func Filter_LinearGradientFilter() Filter {
	return FilterClass.LinearGradientFilter()
}

// Returns a QR code generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228262-qrcodegenerator?language=objc
func (fc _FilterClass) QRCodeGenerator() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("QRCodeGenerator"))
	return rv
}

// Returns a QR code generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228262-qrcodegenerator?language=objc
func Filter_QRCodeGenerator() Filter {
	return FilterClass.QRCodeGenerator()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600144-glassdistortionfilter?language=objc
func (fc _FilterClass) GlassDistortionFilter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("glassDistortionFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3600144-glassdistortionfilter?language=objc
func Filter_GlassDistortionFilter() Filter {
	return FilterClass.GlassDistortionFilter()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3750386-convolutionrgb5x5filter?language=objc
func (fc _FilterClass) ConvolutionRGB5X5Filter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("convolutionRGB5X5Filter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3750386-convolutionrgb5x5filter?language=objc
func Filter_ConvolutionRGB5X5Filter() Filter {
	return FilterClass.ConvolutionRGB5X5Filter()
}

// Creates a CIFilter object for a specific kind of filter and initializes the input values with a nil-terminated list of arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1562057-filterwithname?language=objc
func (fc _FilterClass) FilterWithNameKeysAndValues(name string, key0 objc.IObject, args ...any) Filter {
	rv := objc.Call[Filter](fc, objc.Sel("filterWithName:keysAndValues:"), append([]any{name, key0}, args...)...)
	return rv
}

// Creates a CIFilter object for a specific kind of filter and initializes the input values with a nil-terminated list of arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1562057-filterwithname?language=objc
func Filter_FilterWithNameKeysAndValues(name string, key0 objc.IObject, args ...any) Filter {
	return FilterClass.FilterWithNameKeysAndValues(name, key0, args...)
}

// Returns a convolution 3 x 3 filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228299-convolution3x3filter?language=objc
func (fc _FilterClass) Convolution3X3Filter() Filter {
	rv := objc.Call[Filter](fc, objc.Sel("convolution3X3Filter"))
	return rv
}

// Returns a convolution 3 x 3 filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228299-convolution3x3filter?language=objc
func Filter_Convolution3X3Filter() Filter {
	return FilterClass.Convolution3X3Filter()
}

// The names of all output parameters from the filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1438122-outputkeys?language=objc
func (f_ Filter) OutputKeys() []string {
	rv := objc.Call[[]string](f_, objc.Sel("outputKeys"))
	return rv
}

// The names of all input parameters to the filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1438013-inputkeys?language=objc
func (f_ Filter) InputKeys() []string {
	rv := objc.Call[[]string](f_, objc.Sel("inputKeys"))
	return rv
}

// A Boolean value that determines whether the filter is enabled. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1438276-enabled?language=objc
func (f_ Filter) IsEnabled() bool {
	rv := objc.Call[bool](f_, objc.Sel("isEnabled"))
	return rv
}

// A Boolean value that determines whether the filter is enabled. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1438276-enabled?language=objc
func (f_ Filter) SetEnabled(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setEnabled:"), value)
}

// A dictionary of key-value pairs that describe the filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/1437661-attributes?language=objc
func (f_ Filter) Attributes() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](f_, objc.Sel("attributes"))
	return rv
}

// A CIImage object that encapsulates the operations configured in the filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilter/3228048-outputimage?language=objc
func (f_ Filter) OutputImage() Image {
	rv := objc.Call[Image](f_, objc.Sel("outputImage"))
	return rv
}
