// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetDownloadConfiguration] class.
var AssetDownloadConfigurationClass = _AssetDownloadConfigurationClass{objc.GetClass("AVAssetDownloadConfiguration")}

type _AssetDownloadConfigurationClass struct {
	objc.Class
}

// An interface definition for the [AssetDownloadConfiguration] class.
type IAssetDownloadConfiguration interface {
	objc.IObject
	OptimizesAuxiliaryContentConfigurations() bool
	SetOptimizesAuxiliaryContentConfigurations(value bool)
	AuxiliaryContentConfigurations() []AssetDownloadContentConfiguration
	SetAuxiliaryContentConfigurations(value []IAssetDownloadContentConfiguration)
	ArtworkData() []byte
	SetArtworkData(value []byte)
	PrimaryContentConfiguration() AssetDownloadContentConfiguration
}

// An object that provides the configuration for a download task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration?language=objc
type AssetDownloadConfiguration struct {
	objc.Object
}

func AssetDownloadConfigurationFrom(ptr unsafe.Pointer) AssetDownloadConfiguration {
	return AssetDownloadConfiguration{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetDownloadConfigurationClass) DownloadConfigurationWithAssetTitle(asset IURLAsset, title string) AssetDownloadConfiguration {
	rv := objc.Call[AssetDownloadConfiguration](ac, objc.Sel("downloadConfigurationWithAsset:title:"), objc.Ptr(asset), title)
	return rv
}

// Creates a download configuration for a media asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/3750220-downloadconfigurationwithasset?language=objc
func AssetDownloadConfiguration_DownloadConfigurationWithAssetTitle(asset IURLAsset, title string) AssetDownloadConfiguration {
	return AssetDownloadConfigurationClass.DownloadConfigurationWithAssetTitle(asset, title)
}

func (ac _AssetDownloadConfigurationClass) Alloc() AssetDownloadConfiguration {
	rv := objc.Call[AssetDownloadConfiguration](ac, objc.Sel("alloc"))
	return rv
}

func AssetDownloadConfiguration_Alloc() AssetDownloadConfiguration {
	return AssetDownloadConfigurationClass.Alloc()
}

func (ac _AssetDownloadConfigurationClass) New() AssetDownloadConfiguration {
	rv := objc.Call[AssetDownloadConfiguration](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetDownloadConfiguration() AssetDownloadConfiguration {
	return AssetDownloadConfigurationClass.New()
}

func (a_ AssetDownloadConfiguration) Init() AssetDownloadConfiguration {
	rv := objc.Call[AssetDownloadConfiguration](a_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the task optimizes auxiliary content selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/3750221-optimizesauxiliarycontentconfigu?language=objc
func (a_ AssetDownloadConfiguration) OptimizesAuxiliaryContentConfigurations() bool {
	rv := objc.Call[bool](a_, objc.Sel("optimizesAuxiliaryContentConfigurations"))
	return rv
}

// A Boolean value that indicates whether the task optimizes auxiliary content selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/3750221-optimizesauxiliarycontentconfigu?language=objc
func (a_ AssetDownloadConfiguration) SetOptimizesAuxiliaryContentConfigurations(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setOptimizesAuxiliaryContentConfigurations:"), value)
}

// The configuration for the auxiliary content that the task downloads. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/3750219-auxiliarycontentconfigurations?language=objc
func (a_ AssetDownloadConfiguration) AuxiliaryContentConfigurations() []AssetDownloadContentConfiguration {
	rv := objc.Call[[]AssetDownloadContentConfiguration](a_, objc.Sel("auxiliaryContentConfigurations"))
	return rv
}

// The configuration for the auxiliary content that the task downloads. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/3750219-auxiliarycontentconfigurations?language=objc
func (a_ AssetDownloadConfiguration) SetAuxiliaryContentConfigurations(value []IAssetDownloadContentConfiguration) {
	objc.Call[objc.Void](a_, objc.Sel("setAuxiliaryContentConfigurations:"), value)
}

// A data value that represents the asset’s artwork. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/3750218-artworkdata?language=objc
func (a_ AssetDownloadConfiguration) ArtworkData() []byte {
	rv := objc.Call[[]byte](a_, objc.Sel("artworkData"))
	return rv
}

// A data value that represents the asset’s artwork. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/3750218-artworkdata?language=objc
func (a_ AssetDownloadConfiguration) SetArtworkData(value []byte) {
	objc.Call[objc.Void](a_, objc.Sel("setArtworkData:"), value)
}

// The configuration for the primary content that the task downloads. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/3750222-primarycontentconfiguration?language=objc
func (a_ AssetDownloadConfiguration) PrimaryContentConfiguration() AssetDownloadContentConfiguration {
	rv := objc.Call[AssetDownloadContentConfiguration](a_, objc.Sel("primaryContentConfiguration"))
	return rv
}
