// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FragmentedAssetTrack] class.
var FragmentedAssetTrackClass = _FragmentedAssetTrackClass{objc.GetClass("AVFragmentedAssetTrack")}

type _FragmentedAssetTrackClass struct {
	objc.Class
}

// An interface definition for the [FragmentedAssetTrack] class.
type IFragmentedAssetTrack interface {
	IAssetTrack
}

// An object that provides the track-level interface to inspect a fragmented asset’s media tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedassettrack?language=objc
type FragmentedAssetTrack struct {
	AssetTrack
}

func FragmentedAssetTrackFrom(ptr unsafe.Pointer) FragmentedAssetTrack {
	return FragmentedAssetTrack{
		AssetTrack: AssetTrackFrom(ptr),
	}
}

func (fc _FragmentedAssetTrackClass) Alloc() FragmentedAssetTrack {
	rv := objc.Call[FragmentedAssetTrack](fc, objc.Sel("alloc"))
	return rv
}

func FragmentedAssetTrack_Alloc() FragmentedAssetTrack {
	return FragmentedAssetTrackClass.Alloc()
}

func (fc _FragmentedAssetTrackClass) New() FragmentedAssetTrack {
	rv := objc.Call[FragmentedAssetTrack](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFragmentedAssetTrack() FragmentedAssetTrack {
	return FragmentedAssetTrackClass.New()
}

func (f_ FragmentedAssetTrack) Init() FragmentedAssetTrack {
	rv := objc.Call[FragmentedAssetTrack](f_, objc.Sel("init"))
	return rv
}
