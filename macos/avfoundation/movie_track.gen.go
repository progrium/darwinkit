// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MovieTrack] class.
var MovieTrackClass = _MovieTrackClass{objc.GetClass("AVMovieTrack")}

type _MovieTrackClass struct {
	objc.Class
}

// An interface definition for the [MovieTrack] class.
type IMovieTrack interface {
	IAssetTrack
	MediaPresentationTimeRange() coremedia.TimeRange
	AlternateGroupID() int
	MediaDataStorage() MediaDataStorage
	MediaDecodeTimeRange() coremedia.TimeRange
}

// A track in a movie that conforms to the QuickTime or ISO base media file format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovietrack?language=objc
type MovieTrack struct {
	AssetTrack
}

func MovieTrackFrom(ptr unsafe.Pointer) MovieTrack {
	return MovieTrack{
		AssetTrack: AssetTrackFrom(ptr),
	}
}

func (mc _MovieTrackClass) Alloc() MovieTrack {
	rv := objc.Call[MovieTrack](mc, objc.Sel("alloc"))
	return rv
}

func MovieTrack_Alloc() MovieTrack {
	return MovieTrackClass.Alloc()
}

func (mc _MovieTrackClass) New() MovieTrack {
	rv := objc.Call[MovieTrack](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMovieTrack() MovieTrack {
	return MovieTrackClass.New()
}

func (m_ MovieTrack) Init() MovieTrack {
	rv := objc.Call[MovieTrack](m_, objc.Sel("init"))
	return rv
}

// A range of presentation times for the track’s media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovietrack/1389982-mediapresentationtimerange?language=objc
func (m_ MovieTrack) MediaPresentationTimeRange() coremedia.TimeRange {
	rv := objc.Call[coremedia.TimeRange](m_, objc.Sel("mediaPresentationTimeRange"))
	return rv
}

// A value that identifies the track as a member of a particular alternate group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovietrack/1387020-alternategroupid?language=objc
func (m_ MovieTrack) AlternateGroupID() int {
	rv := objc.Call[int](m_, objc.Sel("alternateGroupID"))
	return rv
}

// The storage container for media data added to a track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovietrack/1386868-mediadatastorage?language=objc
func (m_ MovieTrack) MediaDataStorage() MediaDataStorage {
	rv := objc.Call[MediaDataStorage](m_, objc.Sel("mediaDataStorage"))
	return rv
}

// A range of decode times for the track’s media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovietrack/1388187-mediadecodetimerange?language=objc
func (m_ MovieTrack) MediaDecodeTimeRange() coremedia.TimeRange {
	rv := objc.Call[coremedia.TimeRange](m_, objc.Sel("mediaDecodeTimeRange"))
	return rv
}
