// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetTrackGroup] class.
var AssetTrackGroupClass = _AssetTrackGroupClass{objc.GetClass("AVAssetTrackGroup")}

type _AssetTrackGroupClass struct {
	objc.Class
}

// An interface definition for the [AssetTrackGroup] class.
type IAssetTrackGroup interface {
	objc.IObject
	TrackIDs() []foundation.Number
}

// A group of related tracks in an asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrackgroup?language=objc
type AssetTrackGroup struct {
	objc.Object
}

func AssetTrackGroupFrom(ptr unsafe.Pointer) AssetTrackGroup {
	return AssetTrackGroup{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetTrackGroupClass) Alloc() AssetTrackGroup {
	rv := objc.Call[AssetTrackGroup](ac, objc.Sel("alloc"))
	return rv
}

func AssetTrackGroup_Alloc() AssetTrackGroup {
	return AssetTrackGroupClass.Alloc()
}

func (ac _AssetTrackGroupClass) New() AssetTrackGroup {
	rv := objc.Call[AssetTrackGroup](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetTrackGroup() AssetTrackGroup {
	return AssetTrackGroupClass.New()
}

func (a_ AssetTrackGroup) Init() AssetTrackGroup {
	rv := objc.Call[AssetTrackGroup](a_, objc.Sel("init"))
	return rv
}

// The IDs of the tracks in the group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrackgroup/1389024-trackids?language=objc
func (a_ AssetTrackGroup) TrackIDs() []foundation.Number {
	rv := objc.Call[[]foundation.Number](a_, objc.Sel("trackIDs"))
	return rv
}
