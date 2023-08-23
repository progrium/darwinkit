// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetVariant] class.
var AssetVariantClass = _AssetVariantClass{objc.GetClass("AVAssetVariant")}

type _AssetVariantClass struct {
	objc.Class
}

// An interface definition for the [AssetVariant] class.
type IAssetVariant interface {
	objc.IObject
	PeakBitRate() float64
	AudioAttributes() AssetVariantAudioAttributes
	AverageBitRate() float64
	VideoAttributes() AssetVariantVideoAttributes
}

// An object that represents a bit rate variant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant?language=objc
type AssetVariant struct {
	objc.Object
}

func AssetVariantFrom(ptr unsafe.Pointer) AssetVariant {
	return AssetVariant{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetVariantClass) Alloc() AssetVariant {
	rv := objc.Call[AssetVariant](ac, objc.Sel("alloc"))
	return rv
}

func AssetVariant_Alloc() AssetVariant {
	return AssetVariantClass.Alloc()
}

func (ac _AssetVariantClass) New() AssetVariant {
	rv := objc.Call[AssetVariant](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetVariant() AssetVariant {
	return AssetVariantClass.New()
}

func (a_ AssetVariant) Init() AssetVariant {
	rv := objc.Call[AssetVariant](a_, objc.Sel("init"))
	return rv
}

// The peak bit rate for the variant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/3746545-peakbitrate?language=objc
func (a_ AssetVariant) PeakBitRate() float64 {
	rv := objc.Call[float64](a_, objc.Sel("peakBitRate"))
	return rv
}

// The audio rendition attributes for the variant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/3746543-audioattributes?language=objc
func (a_ AssetVariant) AudioAttributes() AssetVariantAudioAttributes {
	rv := objc.Call[AssetVariantAudioAttributes](a_, objc.Sel("audioAttributes"))
	return rv
}

// The average bit rate for the variant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/3746544-averagebitrate?language=objc
func (a_ AssetVariant) AverageBitRate() float64 {
	rv := objc.Call[float64](a_, objc.Sel("averageBitRate"))
	return rv
}

// The audio rendition attributes for the variant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/3746546-videoattributes?language=objc
func (a_ AssetVariant) VideoAttributes() AssetVariantVideoAttributes {
	rv := objc.Call[AssetVariantVideoAttributes](a_, objc.Sel("videoAttributes"))
	return rv
}
