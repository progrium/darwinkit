// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableCompositionTrack] class.
var MutableCompositionTrackClass = _MutableCompositionTrackClass{objc.GetClass("AVMutableCompositionTrack")}

type _MutableCompositionTrackClass struct {
	objc.Class
}

// An interface definition for the [MutableCompositionTrack] class.
type IMutableCompositionTrack interface {
	ICompositionTrack
	RemoveTrackAssociationToTrackType(compositionTrack ICompositionTrack, trackAssociationType TrackAssociationType)
	AddTrackAssociationToTrackType(compositionTrack ICompositionTrack, trackAssociationType TrackAssociationType)
	RemoveTimeRange(timeRange coremedia.TimeRange)
	ValidateTrackSegmentsError(trackSegments []ICompositionTrackSegment, outError foundation.IError) bool
	InsertTimeRangesOfTracksAtTimeError(timeRanges []foundation.IValue, tracks []IAssetTrack, startTime coremedia.Time, outError foundation.IError) bool
	InsertTimeRangeOfTrackAtTimeError(timeRange coremedia.TimeRange, track IAssetTrack, startTime coremedia.Time, outError foundation.IError) bool
	ReplaceFormatDescriptionWithFormatDescription(originalFormatDescription coremedia.FormatDescriptionRef, replacementFormatDescription coremedia.FormatDescriptionRef)
	ScaleTimeRangeToDuration(timeRange coremedia.TimeRange, duration coremedia.Time)
	InsertEmptyTimeRange(timeRange coremedia.TimeRange)
	SetExtendedLanguageTag(value string)
	SetPreferredVolume(value float64)
	SetSegments(value []ICompositionTrackSegment)
	SetLanguageCode(value string)
	SetEnabled(value bool)
	SetNaturalTimeScale(value coremedia.TimeScale)
	SetPreferredTransform(value coregraphics.AffineTransform)
}

// A mutable track in a composition that you use to insert, remove, and scale track segments without affecting their low-level representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack?language=objc
type MutableCompositionTrack struct {
	CompositionTrack
}

func MutableCompositionTrackFrom(ptr unsafe.Pointer) MutableCompositionTrack {
	return MutableCompositionTrack{
		CompositionTrack: CompositionTrackFrom(ptr),
	}
}

func (mc _MutableCompositionTrackClass) Alloc() MutableCompositionTrack {
	rv := objc.Call[MutableCompositionTrack](mc, objc.Sel("alloc"))
	return rv
}

func MutableCompositionTrack_Alloc() MutableCompositionTrack {
	return MutableCompositionTrackClass.Alloc()
}

func (mc _MutableCompositionTrackClass) New() MutableCompositionTrack {
	rv := objc.Call[MutableCompositionTrack](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableCompositionTrack() MutableCompositionTrack {
	return MutableCompositionTrackClass.New()
}

func (m_ MutableCompositionTrack) Init() MutableCompositionTrack {
	rv := objc.Call[MutableCompositionTrack](m_, objc.Sel("init"))
	return rv
}

// Removes an association from a composition track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack/3013765-removetrackassociationtotrack?language=objc
func (m_ MutableCompositionTrack) RemoveTrackAssociationToTrackType(compositionTrack ICompositionTrack, trackAssociationType TrackAssociationType) {
	objc.Call[objc.Void](m_, objc.Sel("removeTrackAssociationToTrack:type:"), objc.Ptr(compositionTrack), trackAssociationType)
}

// Establishes a track association of a specific type between two tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack/3013764-addtrackassociationtotrack?language=objc
func (m_ MutableCompositionTrack) AddTrackAssociationToTrackType(compositionTrack ICompositionTrack, trackAssociationType TrackAssociationType) {
	objc.Call[objc.Void](m_, objc.Sel("addTrackAssociationToTrack:type:"), objc.Ptr(compositionTrack), trackAssociationType)
}

// Removes a time range of media from a composition track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack/1386048-removetimerange?language=objc
func (m_ MutableCompositionTrack) RemoveTimeRange(timeRange coremedia.TimeRange) {
	objc.Call[objc.Void](m_, objc.Sel("removeTimeRange:"), timeRange)
}

// Returns a Boolean value that indicates whether a given array of track segments conform to the timing rules for a composition track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack/1388746-validatetracksegments?language=objc
func (m_ MutableCompositionTrack) ValidateTrackSegmentsError(trackSegments []ICompositionTrackSegment, outError foundation.IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("validateTrackSegments:error:"), trackSegments, objc.Ptr(outError))
	return rv
}

// Inserts the time ranges of multiple source tracks into a track of a composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack/1388629-inserttimeranges?language=objc
func (m_ MutableCompositionTrack) InsertTimeRangesOfTracksAtTimeError(timeRanges []foundation.IValue, tracks []IAssetTrack, startTime coremedia.Time, outError foundation.IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("insertTimeRanges:ofTracks:atTime:error:"), timeRanges, tracks, startTime, objc.Ptr(outError))
	return rv
}

// Inserts a time range of media from a source track into a composition track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack/1390691-inserttimerange?language=objc
func (m_ MutableCompositionTrack) InsertTimeRangeOfTrackAtTimeError(timeRange coremedia.TimeRange, track IAssetTrack, startTime coremedia.Time, outError foundation.IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("insertTimeRange:ofTrack:atTime:error:"), timeRange, objc.Ptr(track), startTime, objc.Ptr(outError))
	return rv
}

// Replaces a format description with another or cancels a previous replacement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack/3180005-replaceformatdescription?language=objc
func (m_ MutableCompositionTrack) ReplaceFormatDescriptionWithFormatDescription(originalFormatDescription coremedia.FormatDescriptionRef, replacementFormatDescription coremedia.FormatDescriptionRef) {
	objc.Call[objc.Void](m_, objc.Sel("replaceFormatDescription:withFormatDescription:"), originalFormatDescription, replacementFormatDescription)
}

// Changes the duration of a time range of the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack/1388212-scaletimerange?language=objc
func (m_ MutableCompositionTrack) ScaleTimeRangeToDuration(timeRange coremedia.TimeRange, duration coremedia.Time) {
	objc.Call[objc.Void](m_, objc.Sel("scaleTimeRange:toDuration:"), timeRange, duration)
}

// Adds or extends an empty time range within the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack/1389483-insertemptytimerange?language=objc
func (m_ MutableCompositionTrack) InsertEmptyTimeRange(timeRange coremedia.TimeRange) {
	objc.Call[objc.Void](m_, objc.Sel("insertEmptyTimeRange:"), timeRange)
}

// The language tag associated with the track, as an RFC 4646 language tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack/1388866-extendedlanguagetag?language=objc
func (m_ MutableCompositionTrack) SetExtendedLanguageTag(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setExtendedLanguageTag:"), value)
}

// The volume the track prefers for its audible media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack/1388649-preferredvolume?language=objc
func (m_ MutableCompositionTrack) SetPreferredVolume(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setPreferredVolume:"), value)
}

// The track segments that a composition track contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack/1390321-segments?language=objc
func (m_ MutableCompositionTrack) SetSegments(value []ICompositionTrackSegment) {
	objc.Call[objc.Void](m_, objc.Sel("setSegments:"), value)
}

// The language associated with the track, as an ISO 639-2/T language code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack/1387192-languagecode?language=objc
func (m_ MutableCompositionTrack) SetLanguageCode(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setLanguageCode:"), value)
}

// A Boolean value that indicates whether the tracks is in an enabled state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack/3334926-enabled?language=objc
func (m_ MutableCompositionTrack) SetEnabled(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setEnabled:"), value)
}

// The time scale in which you can perform time-based operations without extra numerical conversion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack/1389030-naturaltimescale?language=objc
func (m_ MutableCompositionTrack) SetNaturalTimeScale(value coremedia.TimeScale) {
	objc.Call[objc.Void](m_, objc.Sel("setNaturalTimeScale:"), value)
}

// The preferred transformation of the visual media data for display purposes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecompositiontrack/1388304-preferredtransform?language=objc
func (m_ MutableCompositionTrack) SetPreferredTransform(value coregraphics.AffineTransform) {
	objc.Call[objc.Void](m_, objc.Sel("setPreferredTransform:"), value)
}
