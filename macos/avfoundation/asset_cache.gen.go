// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetCache] class.
var AssetCacheClass = _AssetCacheClass{objc.GetClass("AVAssetCache")}

type _AssetCacheClass struct {
	objc.Class
}

// An interface definition for the [AssetCache] class.
type IAssetCache interface {
	objc.IObject
	MediaSelectionOptionsInMediaSelectionGroup(mediaSelectionGroup IMediaSelectionGroup) []MediaSelectionOption
	IsPlayableOffline() bool
}

// An object that you use to inspect locally cached media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetcache?language=objc
type AssetCache struct {
	objc.Object
}

func AssetCacheFrom(ptr unsafe.Pointer) AssetCache {
	return AssetCache{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetCacheClass) Alloc() AssetCache {
	rv := objc.Call[AssetCache](ac, objc.Sel("alloc"))
	return rv
}

func AssetCache_Alloc() AssetCache {
	return AssetCacheClass.Alloc()
}

func (ac _AssetCacheClass) New() AssetCache {
	rv := objc.Call[AssetCache](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetCache() AssetCache {
	return AssetCacheClass.New()
}

func (a_ AssetCache) Init() AssetCache {
	rv := objc.Call[AssetCache](a_, objc.Sel("init"))
	return rv
}

// Returns an array of locally cached media selection options that are available for offline use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetcache/1823715-mediaselectionoptionsinmediasele?language=objc
func (a_ AssetCache) MediaSelectionOptionsInMediaSelectionGroup(mediaSelectionGroup IMediaSelectionGroup) []MediaSelectionOption {
	rv := objc.Call[[]MediaSelectionOption](a_, objc.Sel("mediaSelectionOptionsInMediaSelectionGroup:"), objc.Ptr(mediaSelectionGroup))
	return rv
}

// A Boolean value that indicates whether the asset is playable without an internet connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetcache/1823708-playableoffline?language=objc
func (a_ AssetCache) IsPlayableOffline() bool {
	rv := objc.Call[bool](a_, objc.Sel("isPlayableOffline"))
	return rv
}
