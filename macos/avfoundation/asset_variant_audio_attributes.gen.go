// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetVariantAudioAttributes] class.
var AssetVariantAudioAttributesClass = _AssetVariantAudioAttributesClass{objc.GetClass("AVAssetVariantAudioAttributes")}

type _AssetVariantAudioAttributesClass struct {
	objc.Class
}

// An interface definition for the [AssetVariantAudioAttributes] class.
type IAssetVariantAudioAttributes interface {
	objc.IObject
	RenditionSpecificAttributesForMediaOption(mediaSelectionOption IMediaSelectionOption) AssetVariantAudioRenditionSpecificAttributes
	FormatIDs() []foundation.Number
}

// An object that defines the audio attributes for an asset variant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantaudioattributes?language=objc
type AssetVariantAudioAttributes struct {
	objc.Object
}

func AssetVariantAudioAttributesFrom(ptr unsafe.Pointer) AssetVariantAudioAttributes {
	return AssetVariantAudioAttributes{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetVariantAudioAttributesClass) Alloc() AssetVariantAudioAttributes {
	rv := objc.Call[AssetVariantAudioAttributes](ac, objc.Sel("alloc"))
	return rv
}

func AssetVariantAudioAttributes_Alloc() AssetVariantAudioAttributes {
	return AssetVariantAudioAttributesClass.Alloc()
}

func (ac _AssetVariantAudioAttributesClass) New() AssetVariantAudioAttributes {
	rv := objc.Call[AssetVariantAudioAttributes](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetVariantAudioAttributes() AssetVariantAudioAttributes {
	return AssetVariantAudioAttributesClass.New()
}

func (a_ AssetVariantAudioAttributes) Init() AssetVariantAudioAttributes {
	rv := objc.Call[AssetVariantAudioAttributes](a_, objc.Sel("init"))
	return rv
}

// Returns specific attributes for the media option. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantaudioattributes/3746549-renditionspecificattributesforme?language=objc
func (a_ AssetVariantAudioAttributes) RenditionSpecificAttributesForMediaOption(mediaSelectionOption IMediaSelectionOption) AssetVariantAudioRenditionSpecificAttributes {
	rv := objc.Call[AssetVariantAudioRenditionSpecificAttributes](a_, objc.Sel("renditionSpecificAttributesForMediaOption:"), objc.Ptr(mediaSelectionOption))
	return rv
}

// The audio formats of the renditions present in the variant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantaudioattributes/3746548-formatids?language=objc
func (a_ AssetVariantAudioAttributes) FormatIDs() []foundation.Number {
	rv := objc.Call[[]foundation.Number](a_, objc.Sel("formatIDs"))
	return rv
}
