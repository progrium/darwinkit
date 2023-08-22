// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetDownloadContentConfiguration] class.
var AssetDownloadContentConfigurationClass = _AssetDownloadContentConfigurationClass{objc.GetClass("AVAssetDownloadContentConfiguration")}

type _AssetDownloadContentConfigurationClass struct {
	objc.Class
}

// An interface definition for the [AssetDownloadContentConfiguration] class.
type IAssetDownloadContentConfiguration interface {
	objc.IObject
	MediaSelections() []MediaSelection
	SetMediaSelections(value []IMediaSelection)
	VariantQualifiers() []AssetVariantQualifier
	SetVariantQualifiers(value []IAssetVariantQualifier)
}

// A configuration object that contains variant qualifiers and media options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadcontentconfiguration?language=objc
type AssetDownloadContentConfiguration struct {
	objc.Object
}

func AssetDownloadContentConfigurationFrom(ptr unsafe.Pointer) AssetDownloadContentConfiguration {
	return AssetDownloadContentConfiguration{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetDownloadContentConfigurationClass) Alloc() AssetDownloadContentConfiguration {
	rv := objc.Call[AssetDownloadContentConfiguration](ac, objc.Sel("alloc"))
	return rv
}

func AssetDownloadContentConfiguration_Alloc() AssetDownloadContentConfiguration {
	return AssetDownloadContentConfigurationClass.Alloc()
}

func (ac _AssetDownloadContentConfigurationClass) New() AssetDownloadContentConfiguration {
	rv := objc.Call[AssetDownloadContentConfiguration](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetDownloadContentConfiguration() AssetDownloadContentConfiguration {
	return AssetDownloadContentConfigurationClass.New()
}

func (a_ AssetDownloadContentConfiguration) Init() AssetDownloadContentConfiguration {
	rv := objc.Call[AssetDownloadContentConfiguration](a_, objc.Sel("init"))
	return rv
}

// The media selections of an asset that a task downloads. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadcontentconfiguration/3750224-mediaselections?language=objc
func (a_ AssetDownloadContentConfiguration) MediaSelections() []MediaSelection {
	rv := objc.Call[[]MediaSelection](a_, objc.Sel("mediaSelections"))
	return rv
}

// The media selections of an asset that a task downloads. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadcontentconfiguration/3750224-mediaselections?language=objc
func (a_ AssetDownloadContentConfiguration) SetMediaSelections(value []IMediaSelection) {
	objc.Call[objc.Void](a_, objc.Sel("setMediaSelections:"), value)
}

// The variant qualifiers for this configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadcontentconfiguration/3750225-variantqualifiers?language=objc
func (a_ AssetDownloadContentConfiguration) VariantQualifiers() []AssetVariantQualifier {
	rv := objc.Call[[]AssetVariantQualifier](a_, objc.Sel("variantQualifiers"))
	return rv
}

// The variant qualifiers for this configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadcontentconfiguration/3750225-variantqualifiers?language=objc
func (a_ AssetDownloadContentConfiguration) SetVariantQualifiers(value []IAssetVariantQualifier) {
	objc.Call[objc.Void](a_, objc.Sel("setVariantQualifiers:"), value)
}
