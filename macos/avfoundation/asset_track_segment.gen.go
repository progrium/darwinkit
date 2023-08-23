// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetTrackSegment] class.
var AssetTrackSegmentClass = _AssetTrackSegmentClass{objc.GetClass("AVAssetTrackSegment")}

type _AssetTrackSegmentClass struct {
	objc.Class
}

// An interface definition for the [AssetTrackSegment] class.
type IAssetTrackSegment interface {
	objc.IObject
	IsEmpty() bool
	TimeMapping() coremedia.TimeMapping
}

// An object that represents a time range segment of an asset track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettracksegment?language=objc
type AssetTrackSegment struct {
	objc.Object
}

func AssetTrackSegmentFrom(ptr unsafe.Pointer) AssetTrackSegment {
	return AssetTrackSegment{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetTrackSegmentClass) Alloc() AssetTrackSegment {
	rv := objc.Call[AssetTrackSegment](ac, objc.Sel("alloc"))
	return rv
}

func AssetTrackSegment_Alloc() AssetTrackSegment {
	return AssetTrackSegmentClass.Alloc()
}

func (ac _AssetTrackSegmentClass) New() AssetTrackSegment {
	rv := objc.Call[AssetTrackSegment](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetTrackSegment() AssetTrackSegment {
	return AssetTrackSegmentClass.New()
}

func (a_ AssetTrackSegment) Init() AssetTrackSegment {
	rv := objc.Call[AssetTrackSegment](a_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the segment is empty. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettracksegment/1385714-empty?language=objc
func (a_ AssetTrackSegment) IsEmpty() bool {
	rv := objc.Call[bool](a_, objc.Sel("isEmpty"))
	return rv
}

// The time range of the track that this segment presents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettracksegment/1390557-timemapping?language=objc
func (a_ AssetTrackSegment) TimeMapping() coremedia.TimeMapping {
	rv := objc.Call[coremedia.TimeMapping](a_, objc.Sel("timeMapping"))
	return rv
}
