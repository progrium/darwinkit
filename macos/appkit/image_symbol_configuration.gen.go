// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ImageSymbolConfigurationClass = _ImageSymbolConfigurationClass{objc.GetClass("NSImageSymbolConfiguration")}

type _ImageSymbolConfigurationClass struct {
	objc.Class
}

type IImageSymbolConfiguration interface {
	objc.IObject
}

type ImageSymbolConfiguration struct {
	objc.Object
}

func MakeImageSymbolConfiguration(ptr unsafe.Pointer) ImageSymbolConfiguration {
	return ImageSymbolConfiguration{
		Object: objc.MakeObject(ptr),
	}
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithPointSizeWeight(pointSize float64, weight FontWeight) ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationWithPointSize:weight:"), pointSize, weight)
	return rv
}

func ImageSymbolConfiguration_ConfigurationWithPointSizeWeight(pointSize float64, weight FontWeight) ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.ConfigurationWithPointSizeWeight(pointSize, weight)
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithPointSizeWeightScale(pointSize float64, weight FontWeight, scale ImageSymbolScale) ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationWithPointSize:weight:scale:"), pointSize, weight, scale)
	return rv
}

func ImageSymbolConfiguration_ConfigurationWithPointSizeWeightScale(pointSize float64, weight FontWeight, scale ImageSymbolScale) ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.ConfigurationWithPointSizeWeightScale(pointSize, weight, scale)
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithTextStyle(style FontTextStyle) ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationWithTextStyle:"), style)
	return rv
}

func ImageSymbolConfiguration_ConfigurationWithTextStyle(style FontTextStyle) ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.ConfigurationWithTextStyle(style)
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithTextStyleScale(style FontTextStyle, scale ImageSymbolScale) ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationWithTextStyle:scale:"), style, scale)
	return rv
}

func ImageSymbolConfiguration_ConfigurationWithTextStyleScale(style FontTextStyle, scale ImageSymbolScale) ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.ConfigurationWithTextStyleScale(style, scale)
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithScale(scale ImageSymbolScale) ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationWithScale:"), scale)
	return rv
}

func ImageSymbolConfiguration_ConfigurationWithScale(scale ImageSymbolScale) ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.ConfigurationWithScale(scale)
}

func (i_ ImageSymbolConfiguration) ConfigurationByApplyingConfiguration(configuration IImageSymbolConfiguration) ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](i_, objc.GetSelector("configurationByApplyingConfiguration:"), objc.ExtractPtr(configuration))
	return rv
}

func ImageSymbolConfiguration_ConfigurationByApplyingConfiguration(configuration IImageSymbolConfiguration) ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.Alloc().ConfigurationByApplyingConfiguration(configuration)
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithPaletteColors(paletteColors []IColor) ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationWithPaletteColors:"), paletteColors)
	return rv
}

func ImageSymbolConfiguration_ConfigurationWithPaletteColors(paletteColors []IColor) ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.ConfigurationWithPaletteColors(paletteColors)
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithHierarchicalColor(hierarchicalColor IColor) ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationWithHierarchicalColor:"), objc.ExtractPtr(hierarchicalColor))
	return rv
}

func ImageSymbolConfiguration_ConfigurationWithHierarchicalColor(hierarchicalColor IColor) ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.ConfigurationWithHierarchicalColor(hierarchicalColor)
}

func (ic _ImageSymbolConfigurationClass) ConfigurationPreferringMulticolor() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationPreferringMulticolor"))
	return rv
}

func ImageSymbolConfiguration_ConfigurationPreferringMulticolor() ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.ConfigurationPreferringMulticolor()
}

func (ic _ImageSymbolConfigurationClass) ConfigurationPreferringHierarchical() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationPreferringHierarchical"))
	return rv
}

func ImageSymbolConfiguration_ConfigurationPreferringHierarchical() ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.ConfigurationPreferringHierarchical()
}

func (ic _ImageSymbolConfigurationClass) ConfigurationPreferringMonochrome() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationPreferringMonochrome"))
	return rv
}

func ImageSymbolConfiguration_ConfigurationPreferringMonochrome() ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.ConfigurationPreferringMonochrome()
}

func (ic _ImageSymbolConfigurationClass) Alloc() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("alloc"))
	return rv
}

func ImageSymbolConfiguration_Alloc() ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.Alloc()
}

func (ic _ImageSymbolConfigurationClass) New() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewImageSymbolConfiguration() ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.New()
}

func ImageSymbolConfiguration_New() ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.New()
}

func (i_ ImageSymbolConfiguration) Init() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](i_, objc.GetSelector("init"))
	return rv
}

func ImageSymbolConfiguration_Init() ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.Alloc().Init()
}
