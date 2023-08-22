// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableMovieTrack] class.
var MutableMovieTrackClass = _MutableMovieTrackClass{objc.GetClass("AVMutableMovieTrack")}

type _MutableMovieTrackClass struct {
	objc.Class
}

// An interface definition for the [MutableMovieTrack] class.
type IMutableMovieTrack interface {
	IMovieTrack
	RemoveTrackAssociationToTrackType(movieTrack IMovieTrack, trackAssociationType TrackAssociationType)
	AddTrackAssociationToTrackType(movieTrack IMovieTrack, trackAssociationType TrackAssociationType)
	InsertMediaTimeRangeIntoTimeRange(mediaTimeRange coremedia.TimeRange, trackTimeRange coremedia.TimeRange) bool
	RemoveTimeRange(timeRange coremedia.TimeRange)
	AppendSampleBufferDecodeTimePresentationTimeError(sampleBuffer coremedia.SampleBufferRef, outDecodeTime *coremedia.Time, outPresentationTime *coremedia.Time, outError foundation.IError) bool
	InsertTimeRangeOfTrackAtTimeCopySampleDataError(timeRange coremedia.TimeRange, track IAssetTrack, startTime coremedia.Time, copySampleData bool, outError foundation.IError) bool
	ReplaceFormatDescriptionWithFormatDescription(formatDescription coremedia.FormatDescriptionRef, newFormatDescription coremedia.FormatDescriptionRef)
	ScaleTimeRangeToDuration(timeRange coremedia.TimeRange, duration coremedia.Time)
	InsertEmptyTimeRange(timeRange coremedia.TimeRange)
	Layer() int
	SetLayer(value int)
	PreferredMediaChunkSize() int
	SetPreferredMediaChunkSize(value int)
	SetExtendedLanguageTag(value string)
	ProductionApertureDimensions() coregraphics.Size
	SetProductionApertureDimensions(value coregraphics.Size)
	IsModified() bool
	SetModified(value bool)
	PreferredMediaChunkAlignment() int
	SetPreferredMediaChunkAlignment(value int)
	SetPreferredVolume(value float64)
	SetAlternateGroupID(value int)
	SetMetadata(value []IMetadataItem)
	SampleReferenceBaseURL() foundation.URL
	SetSampleReferenceBaseURL(value foundation.IURL)
	SetMediaDataStorage(value IMediaDataStorage)
	CleanApertureDimensions() coregraphics.Size
	SetCleanApertureDimensions(value coregraphics.Size)
	SetNaturalSize(value coregraphics.Size)
	EncodedPixelsDimensions() coregraphics.Size
	SetEncodedPixelsDimensions(value coregraphics.Size)
	PreferredMediaChunkDuration() coremedia.Time
	SetPreferredMediaChunkDuration(value coremedia.Time)
	HasProtectedContent() bool
	SetLanguageCode(value string)
	SetEnabled(value bool)
	Timescale() coremedia.TimeScale
	SetTimescale(value coremedia.TimeScale)
	SetPreferredTransform(value coregraphics.AffineTransform)
}

// A mutable track that conforms to the QuickTime or ISO base media file format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack?language=objc
type MutableMovieTrack struct {
	MovieTrack
}

func MutableMovieTrackFrom(ptr unsafe.Pointer) MutableMovieTrack {
	return MutableMovieTrack{
		MovieTrack: MovieTrackFrom(ptr),
	}
}

func (mc _MutableMovieTrackClass) Alloc() MutableMovieTrack {
	rv := objc.Call[MutableMovieTrack](mc, objc.Sel("alloc"))
	return rv
}

func MutableMovieTrack_Alloc() MutableMovieTrack {
	return MutableMovieTrackClass.Alloc()
}

func (mc _MutableMovieTrackClass) New() MutableMovieTrack {
	rv := objc.Call[MutableMovieTrack](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableMovieTrack() MutableMovieTrack {
	return MutableMovieTrackClass.New()
}

func (m_ MutableMovieTrack) Init() MutableMovieTrack {
	rv := objc.Call[MutableMovieTrack](m_, objc.Sel("init"))
	return rv
}

// Removes a specific type of track association between two tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1389620-removetrackassociationtotrack?language=objc
func (m_ MutableMovieTrack) RemoveTrackAssociationToTrackType(movieTrack IMovieTrack, trackAssociationType TrackAssociationType) {
	objc.Call[objc.Void](m_, objc.Sel("removeTrackAssociationToTrack:type:"), objc.Ptr(movieTrack), trackAssociationType)
}

// Creates a specific type of track association between two tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1390163-addtrackassociationtotrack?language=objc
func (m_ MutableMovieTrack) AddTrackAssociationToTrackType(movieTrack IMovieTrack, trackAssociationType TrackAssociationType) {
	objc.Call[objc.Void](m_, objc.Sel("addTrackAssociationToTrack:type:"), objc.Ptr(movieTrack), trackAssociationType)
}

// Inserts a reference to a media time range into a track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1638038-insertmediatimerange?language=objc
func (m_ MutableMovieTrack) InsertMediaTimeRangeIntoTimeRange(mediaTimeRange coremedia.TimeRange, trackTimeRange coremedia.TimeRange) bool {
	rv := objc.Call[bool](m_, objc.Sel("insertMediaTimeRange:intoTimeRange:"), mediaTimeRange, trackTimeRange)
	return rv
}

// Removes the specified time range from a track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1385962-removetimerange?language=objc
func (m_ MutableMovieTrack) RemoveTimeRange(timeRange coremedia.TimeRange) {
	objc.Call[objc.Void](m_, objc.Sel("removeTimeRange:"), timeRange)
}

// Appends sample data to a media file and adds sample references for the added data to a track’s media sample tables. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1638041-appendsamplebuffer?language=objc
func (m_ MutableMovieTrack) AppendSampleBufferDecodeTimePresentationTimeError(sampleBuffer coremedia.SampleBufferRef, outDecodeTime *coremedia.Time, outPresentationTime *coremedia.Time, outError foundation.IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("appendSampleBuffer:decodeTime:presentationTime:error:"), sampleBuffer, outDecodeTime, outPresentationTime, objc.Ptr(outError))
	return rv
}

// Inserts a portion of an asset track into the target movie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1387665-inserttimerange?language=objc
func (m_ MutableMovieTrack) InsertTimeRangeOfTrackAtTimeCopySampleDataError(timeRange coremedia.TimeRange, track IAssetTrack, startTime coremedia.Time, copySampleData bool, outError foundation.IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("insertTimeRange:ofTrack:atTime:copySampleData:error:"), timeRange, objc.Ptr(track), startTime, copySampleData, objc.Ptr(outError))
	return rv
}

// Replaces the track’s format description with a new format description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/2876160-replaceformatdescription?language=objc
func (m_ MutableMovieTrack) ReplaceFormatDescriptionWithFormatDescription(formatDescription coremedia.FormatDescriptionRef, newFormatDescription coremedia.FormatDescriptionRef) {
	objc.Call[objc.Void](m_, objc.Sel("replaceFormatDescription:withFormatDescription:"), formatDescription, newFormatDescription)
}

// Changes the duration of a time range in a track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1388618-scaletimerange?language=objc
func (m_ MutableMovieTrack) ScaleTimeRangeToDuration(timeRange coremedia.TimeRange, duration coremedia.Time) {
	objc.Call[objc.Void](m_, objc.Sel("scaleTimeRange:toDuration:"), timeRange, duration)
}

// Adds an empty time range to a track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1389441-insertemptytimerange?language=objc
func (m_ MutableMovieTrack) InsertEmptyTimeRange(timeRange coremedia.TimeRange) {
	objc.Call[objc.Void](m_, objc.Sel("insertEmptyTimeRange:"), timeRange)
}

// The layer level for the visual media of the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1387655-layer?language=objc
func (m_ MutableMovieTrack) Layer() int {
	rv := objc.Call[int](m_, objc.Sel("layer"))
	return rv
}

// The layer level for the visual media of the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1387655-layer?language=objc
func (m_ MutableMovieTrack) SetLayer(value int) {
	objc.Call[objc.Void](m_, objc.Sel("setLayer:"), value)
}

// The maximum size to use for each chunk of sample data written to the file for file types that support media chunk duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1390149-preferredmediachunksize?language=objc
func (m_ MutableMovieTrack) PreferredMediaChunkSize() int {
	rv := objc.Call[int](m_, objc.Sel("preferredMediaChunkSize"))
	return rv
}

// The maximum size to use for each chunk of sample data written to the file for file types that support media chunk duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1390149-preferredmediachunksize?language=objc
func (m_ MutableMovieTrack) SetPreferredMediaChunkSize(value int) {
	objc.Call[objc.Void](m_, objc.Sel("setPreferredMediaChunkSize:"), value)
}

// An IETF BCP 47 language identifier that identifies the language tag associated with the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1389056-extendedlanguagetag?language=objc
func (m_ MutableMovieTrack) SetExtendedLanguageTag(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setExtendedLanguageTag:"), value)
}

// The production aperture dimensions of the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1390108-productionaperturedimensions?language=objc
func (m_ MutableMovieTrack) ProductionApertureDimensions() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](m_, objc.Sel("productionApertureDimensions"))
	return rv
}

// The production aperture dimensions of the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1390108-productionaperturedimensions?language=objc
func (m_ MutableMovieTrack) SetProductionApertureDimensions(value coregraphics.Size) {
	objc.Call[objc.Void](m_, objc.Sel("setProductionApertureDimensions:"), value)
}

// A Boolean value that indicates whether a track is in a modified state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1390201-modified?language=objc
func (m_ MutableMovieTrack) IsModified() bool {
	rv := objc.Call[bool](m_, objc.Sel("isModified"))
	return rv
}

// A Boolean value that indicates whether a track is in a modified state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1390201-modified?language=objc
func (m_ MutableMovieTrack) SetModified(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setModified:"), value)
}

// The boundary for media chunk alignment for file types that support media chunk alignment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1390504-preferredmediachunkalignment?language=objc
func (m_ MutableMovieTrack) PreferredMediaChunkAlignment() int {
	rv := objc.Call[int](m_, objc.Sel("preferredMediaChunkAlignment"))
	return rv
}

// The boundary for media chunk alignment for file types that support media chunk alignment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1390504-preferredmediachunkalignment?language=objc
func (m_ MutableMovieTrack) SetPreferredMediaChunkAlignment(value int) {
	objc.Call[objc.Void](m_, objc.Sel("setPreferredMediaChunkAlignment:"), value)
}

// The preferred volume for the audible medata data of the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1390391-preferredvolume?language=objc
func (m_ MutableMovieTrack) SetPreferredVolume(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setPreferredVolume:"), value)
}

// A number that identifies the track as a member of a particular alternate group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1387206-alternategroupid?language=objc
func (m_ MutableMovieTrack) SetAlternateGroupID(value int) {
	objc.Call[objc.Void](m_, objc.Sel("setAlternateGroupID:"), value)
}

// An array of metadata stored by the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1390023-metadata?language=objc
func (m_ MutableMovieTrack) SetMetadata(value []IMetadataItem) {
	objc.Call[objc.Void](m_, objc.Sel("setMetadata:"), value)
}

// The base URL for sample references. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1385583-samplereferencebaseurl?language=objc
func (m_ MutableMovieTrack) SampleReferenceBaseURL() foundation.URL {
	rv := objc.Call[foundation.URL](m_, objc.Sel("sampleReferenceBaseURL"))
	return rv
}

// The base URL for sample references. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1385583-samplereferencebaseurl?language=objc
func (m_ MutableMovieTrack) SetSampleReferenceBaseURL(value foundation.IURL) {
	objc.Call[objc.Void](m_, objc.Sel("setSampleReferenceBaseURL:"), objc.Ptr(value))
}

// A storage container for the media data to be added to a track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1386532-mediadatastorage?language=objc
func (m_ MutableMovieTrack) SetMediaDataStorage(value IMediaDataStorage) {
	objc.Call[objc.Void](m_, objc.Sel("setMediaDataStorage:"), objc.Ptr(value))
}

// The clean aperture dimension of the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1386454-cleanaperturedimensions?language=objc
func (m_ MutableMovieTrack) CleanApertureDimensions() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](m_, objc.Sel("cleanApertureDimensions"))
	return rv
}

// The clean aperture dimension of the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1386454-cleanaperturedimensions?language=objc
func (m_ MutableMovieTrack) SetCleanApertureDimensions(value coregraphics.Size) {
	objc.Call[objc.Void](m_, objc.Sel("setCleanApertureDimensions:"), value)
}

// The dimensions used to display the visual media data for the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1385900-naturalsize?language=objc
func (m_ MutableMovieTrack) SetNaturalSize(value coregraphics.Size) {
	objc.Call[objc.Void](m_, objc.Sel("setNaturalSize:"), value)
}

// The encoded pixels dimensions of the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1389417-encodedpixelsdimensions?language=objc
func (m_ MutableMovieTrack) EncodedPixelsDimensions() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](m_, objc.Sel("encodedPixelsDimensions"))
	return rv
}

// The encoded pixels dimensions of the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1389417-encodedpixelsdimensions?language=objc
func (m_ MutableMovieTrack) SetEncodedPixelsDimensions(value coregraphics.Size) {
	objc.Call[objc.Void](m_, objc.Sel("setEncodedPixelsDimensions:"), value)
}

// The maximum duration to use for each chunk of sample data written to the file for file types that support media chunk duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1390292-preferredmediachunkduration?language=objc
func (m_ MutableMovieTrack) PreferredMediaChunkDuration() coremedia.Time {
	rv := objc.Call[coremedia.Time](m_, objc.Sel("preferredMediaChunkDuration"))
	return rv
}

// The maximum duration to use for each chunk of sample data written to the file for file types that support media chunk duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1390292-preferredmediachunkduration?language=objc
func (m_ MutableMovieTrack) SetPreferredMediaChunkDuration(value coremedia.Time) {
	objc.Call[objc.Void](m_, objc.Sel("setPreferredMediaChunkDuration:"), value)
}

// A Boolean value that indicates whether a track contains protected content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1389542-hasprotectedcontent?language=objc
func (m_ MutableMovieTrack) HasProtectedContent() bool {
	rv := objc.Call[bool](m_, objc.Sel("hasProtectedContent"))
	return rv
}

// A ISO 639-2/T language code that indicates the language associated with the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1389736-languagecode?language=objc
func (m_ MutableMovieTrack) SetLanguageCode(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setLanguageCode:"), value)
}

// A Boolean value that indicates whether the track is enabled by default for presentation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1386340-enabled?language=objc
func (m_ MutableMovieTrack) SetEnabled(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setEnabled:"), value)
}

// The time scale for tracks that contain the moov atom. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1388055-timescale?language=objc
func (m_ MutableMovieTrack) Timescale() coremedia.TimeScale {
	rv := objc.Call[coremedia.TimeScale](m_, objc.Sel("timescale"))
	return rv
}

// The time scale for tracks that contain the moov atom. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1388055-timescale?language=objc
func (m_ MutableMovieTrack) SetTimescale(value coremedia.TimeScale) {
	objc.Call[objc.Void](m_, objc.Sel("setTimescale:"), value)
}

// The transform performed on the visual media data of the track for display purposes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablemovietrack/1386593-preferredtransform?language=objc
func (m_ MutableMovieTrack) SetPreferredTransform(value coregraphics.AffineTransform) {
	objc.Call[objc.Void](m_, objc.Sel("setPreferredTransform:"), value)
}
