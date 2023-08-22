// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetVariantAudioRenditionSpecificAttributes] class.
var AssetVariantAudioRenditionSpecificAttributesClass = _AssetVariantAudioRenditionSpecificAttributesClass{objc.GetClass("AVAssetVariantAudioRenditionSpecificAttributes")}

type _AssetVariantAudioRenditionSpecificAttributesClass struct {
	objc.Class
}

// An interface definition for the [AssetVariantAudioRenditionSpecificAttributes] class.
type IAssetVariantAudioRenditionSpecificAttributes interface {
	objc.IObject
	ChannelCount() int
}

// An object that represents attributes specific to a particular rendition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantaudiorenditionspecificattributes?language=objc
type AssetVariantAudioRenditionSpecificAttributes struct {
	objc.Object
}

func AssetVariantAudioRenditionSpecificAttributesFrom(ptr unsafe.Pointer) AssetVariantAudioRenditionSpecificAttributes {
	return AssetVariantAudioRenditionSpecificAttributes{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetVariantAudioRenditionSpecificAttributesClass) Alloc() AssetVariantAudioRenditionSpecificAttributes {
	rv := objc.Call[AssetVariantAudioRenditionSpecificAttributes](ac, objc.Sel("alloc"))
	return rv
}

func AssetVariantAudioRenditionSpecificAttributes_Alloc() AssetVariantAudioRenditionSpecificAttributes {
	return AssetVariantAudioRenditionSpecificAttributesClass.Alloc()
}

func (ac _AssetVariantAudioRenditionSpecificAttributesClass) New() AssetVariantAudioRenditionSpecificAttributes {
	rv := objc.Call[AssetVariantAudioRenditionSpecificAttributes](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetVariantAudioRenditionSpecificAttributes() AssetVariantAudioRenditionSpecificAttributes {
	return AssetVariantAudioRenditionSpecificAttributesClass.New()
}

func (a_ AssetVariantAudioRenditionSpecificAttributes) Init() AssetVariantAudioRenditionSpecificAttributes {
	rv := objc.Call[AssetVariantAudioRenditionSpecificAttributes](a_, objc.Sel("init"))
	return rv
}

// The count of audio channels in the rendition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantaudiorenditionspecificattributes/3746551-channelcount?language=objc
func (a_ AssetVariantAudioRenditionSpecificAttributes) ChannelCount() int {
	rv := objc.Call[int](a_, objc.Sel("channelCount"))
	return rv
}
