// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageSymbolConfiguration] class.
var ImageSymbolConfigurationClass = _ImageSymbolConfigurationClass{objc.GetClass("NSImageSymbolConfiguration")}

type _ImageSymbolConfigurationClass struct {
	objc.Class
}

// An interface definition for the [ImageSymbolConfiguration] class.
type IImageSymbolConfiguration interface {
	objc.IObject
}

// An object that contains the specific font, style, and weight attributes to apply to a symbol image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagesymbolconfiguration?language=objc
type ImageSymbolConfiguration struct {
	objc.Object
}

func ImageSymbolConfigurationFrom(ptr unsafe.Pointer) ImageSymbolConfiguration {
	return ImageSymbolConfiguration{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithPaletteColors(paletteColors []IColor) ImageSymbolConfiguration {
	rv := objc.Call[ImageSymbolConfiguration](ic, objc.Sel("configurationWithPaletteColors:"), paletteColors)
	return rv
}

// Creates a color configuration by specifying a palette of colors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagesymbolconfiguration/3852563-configurationwithpalettecolors?language=objc
func ImageSymbolConfiguration_ConfigurationWithPaletteColors(paletteColors []IColor) ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.ConfigurationWithPaletteColors(paletteColors)
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithPointSizeWeight(pointSize float64, weight FontWeight) ImageSymbolConfiguration {
	rv := objc.Call[ImageSymbolConfiguration](ic, objc.Sel("configurationWithPointSize:weight:"), pointSize, weight)
	return rv
}

// Creates a symbol configuration with the specified point size and font weight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagesymbolconfiguration/3656510-configurationwithpointsize?language=objc
func ImageSymbolConfiguration_ConfigurationWithPointSizeWeight(pointSize float64, weight FontWeight) ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.ConfigurationWithPointSizeWeight(pointSize, weight)
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithScale(scale ImageSymbolScale) ImageSymbolConfiguration {
	rv := objc.Call[ImageSymbolConfiguration](ic, objc.Sel("configurationWithScale:"), scale)
	return rv
}

// Creates a symbol configuration using the scale you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagesymbolconfiguration/3667455-configurationwithscale?language=objc
func ImageSymbolConfiguration_ConfigurationWithScale(scale ImageSymbolScale) ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.ConfigurationWithScale(scale)
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithHierarchicalColor(hierarchicalColor IColor) ImageSymbolConfiguration {
	rv := objc.Call[ImageSymbolConfiguration](ic, objc.Sel("configurationWithHierarchicalColor:"), objc.Ptr(hierarchicalColor))
	return rv
}

// Creates a hierarchical color configuration using the color you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagesymbolconfiguration/3852562-configurationwithhierarchicalcol?language=objc
func ImageSymbolConfiguration_ConfigurationWithHierarchicalColor(hierarchicalColor IColor) ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.ConfigurationWithHierarchicalColor(hierarchicalColor)
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithTextStyle(style FontTextStyle) ImageSymbolConfiguration {
	rv := objc.Call[ImageSymbolConfiguration](ic, objc.Sel("configurationWithTextStyle:"), style)
	return rv
}

// Creates a symbol configuration with the specified text style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagesymbolconfiguration/3656512-configurationwithtextstyle?language=objc
func ImageSymbolConfiguration_ConfigurationWithTextStyle(style FontTextStyle) ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.ConfigurationWithTextStyle(style)
}

func (ic _ImageSymbolConfigurationClass) ConfigurationPreferringMulticolor() ImageSymbolConfiguration {
	rv := objc.Call[ImageSymbolConfiguration](ic, objc.Sel("configurationPreferringMulticolor"))
	return rv
}

// Creates a configuration that specifies that the symbol should prefer its multicolor variant if one exists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagesymbolconfiguration/3852561-configurationpreferringmulticolo?language=objc
func ImageSymbolConfiguration_ConfigurationPreferringMulticolor() ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.ConfigurationPreferringMulticolor()
}

func (i_ ImageSymbolConfiguration) ConfigurationByApplyingConfiguration(configuration IImageSymbolConfiguration) ImageSymbolConfiguration {
	rv := objc.Call[ImageSymbolConfiguration](i_, objc.Sel("configurationByApplyingConfiguration:"), objc.Ptr(configuration))
	return rv
}

// Creates a configuration object by applying the values from the configuration you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagesymbolconfiguration/3852560-configurationbyapplyingconfigura?language=objc
func ImageSymbolConfiguration_ConfigurationByApplyingConfiguration(configuration IImageSymbolConfiguration) ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.Alloc().ConfigurationByApplyingConfiguration(configuration)
}

func (ic _ImageSymbolConfigurationClass) Alloc() ImageSymbolConfiguration {
	rv := objc.Call[ImageSymbolConfiguration](ic, objc.Sel("alloc"))
	return rv
}

func ImageSymbolConfiguration_Alloc() ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.Alloc()
}

func (ic _ImageSymbolConfigurationClass) New() ImageSymbolConfiguration {
	rv := objc.Call[ImageSymbolConfiguration](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageSymbolConfiguration() ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.New()
}

func (i_ ImageSymbolConfiguration) Init() ImageSymbolConfiguration {
	rv := objc.Call[ImageSymbolConfiguration](i_, objc.Sel("init"))
	return rv
}