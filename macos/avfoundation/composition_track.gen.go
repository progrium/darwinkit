// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CompositionTrack] class.
var CompositionTrackClass = _CompositionTrackClass{objc.GetClass("AVCompositionTrack")}

type _CompositionTrackClass struct {
	objc.Class
}

// An interface definition for the [CompositionTrack] class.
type ICompositionTrack interface {
	IAssetTrack
	SegmentForTrackTime(trackTime coremedia.Time) CompositionTrackSegment
	FormatDescriptionReplacements() []CompositionTrackFormatDescriptionReplacement
}

// A track in a composition that presents media of a uniform type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack?language=objc
type CompositionTrack struct {
	AssetTrack
}

func CompositionTrackFrom(ptr unsafe.Pointer) CompositionTrack {
	return CompositionTrack{
		AssetTrack: AssetTrackFrom(ptr),
	}
}

func (cc _CompositionTrackClass) Alloc() CompositionTrack {
	rv := objc.Call[CompositionTrack](cc, objc.Sel("alloc"))
	return rv
}

func CompositionTrack_Alloc() CompositionTrack {
	return CompositionTrackClass.Alloc()
}

func (cc _CompositionTrackClass) New() CompositionTrack {
	rv := objc.Call[CompositionTrack](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCompositionTrack() CompositionTrack {
	return CompositionTrackClass.New()
}

func (c_ CompositionTrack) Init() CompositionTrack {
	rv := objc.Call[CompositionTrack](c_, objc.Sel("init"))
	return rv
}

// Returns a segment whose target time range contains, or is closest to, the specified track time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/2865555-segmentfortracktime?language=objc
func (c_ CompositionTrack) SegmentForTrackTime(trackTime coremedia.Time) CompositionTrackSegment {
	rv := objc.Call[CompositionTrackSegment](c_, objc.Sel("segmentForTrackTime:"), trackTime)
	return rv
}

// The replacement format descriptions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/3194194-formatdescriptionreplacements?language=objc
func (c_ CompositionTrack) FormatDescriptionReplacements() []CompositionTrackFormatDescriptionReplacement {
	rv := objc.Call[[]CompositionTrackFormatDescriptionReplacement](c_, objc.Sel("formatDescriptionReplacements"))
	return rv
}
