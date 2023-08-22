// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SampleCursor] class.
var SampleCursorClass = _SampleCursorClass{objc.GetClass("AVSampleCursor")}

type _SampleCursorClass struct {
	objc.Class
}

// An interface definition for the [SampleCursor] class.
type ISampleCursor interface {
	objc.IObject
	StepInPresentationOrderByCount(stepCount int64) int64
	StepInDecodeOrderByCount(stepCount int64) int64
	SamplesWithLaterDecodeTimeStampsMayHaveEarlierPresentationTimeStampsThanCursor(cursor ISampleCursor) bool
	StepByDecodeTimeWasPinned(deltaDecodeTime coremedia.Time, outWasPinned *bool) coremedia.Time
	ComparePositionInDecodeOrderWithPositionOfCursor(cursor ISampleCursor) foundation.ComparisonResult
	CopyCurrentSampleFormatDescription() coremedia.FormatDescriptionRef
	StepByPresentationTimeWasPinned(deltaPresentationTime coremedia.Time, outWasPinned *bool) coremedia.Time
	SamplesWithEarlierDecodeTimeStampsMayHaveLaterPresentationTimeStampsThanCursor(cursor ISampleCursor) bool
	CurrentChunkStorageRange() SampleCursorStorageRange
	DecodeTimeStamp() coremedia.Time
	CurrentSampleSyncInfo() SampleCursorSyncInfo
	CurrentChunkInfo() SampleCursorChunkInfo
	CurrentSampleDuration() coremedia.Time
	SamplesRequiredForDecoderRefresh() int
	CurrentSampleStorageRange() SampleCursorStorageRange
	PresentationTimeStamp() coremedia.Time
	CurrentSampleDependencyInfo() SampleCursorDependencyInfo
	CurrentSampleDependencyAttachments() foundation.Dictionary
	CurrentSampleIndexInChunk() int64
	CurrentSampleAudioDependencyInfo() SampleCursorAudioDependencyInfo
	CurrentChunkStorageURL() foundation.URL
}

// An object that provides information about the media sample at the cursor’s current position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor?language=objc
type SampleCursor struct {
	objc.Object
}

func SampleCursorFrom(ptr unsafe.Pointer) SampleCursor {
	return SampleCursor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SampleCursorClass) Alloc() SampleCursor {
	rv := objc.Call[SampleCursor](sc, objc.Sel("alloc"))
	return rv
}

func SampleCursor_Alloc() SampleCursor {
	return SampleCursorClass.Alloc()
}

func (sc _SampleCursorClass) New() SampleCursor {
	rv := objc.Call[SampleCursor](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSampleCursor() SampleCursor {
	return SampleCursorClass.New()
}

func (s_ SampleCursor) Init() SampleCursor {
	rv := objc.Call[SampleCursor](s_, objc.Sel("init"))
	return rv
}

// Moves the cursor a given number of samples in presentation order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1388834-stepinpresentationorderbycount?language=objc
func (s_ SampleCursor) StepInPresentationOrderByCount(stepCount int64) int64 {
	rv := objc.Call[int64](s_, objc.Sel("stepInPresentationOrderByCount:"), stepCount)
	return rv
}

// Moves the cursor a given number of samples in decode order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1389606-stepindecodeorderbycount?language=objc
func (s_ SampleCursor) StepInDecodeOrderByCount(stepCount int64) int64 {
	rv := objc.Call[int64](s_, objc.Sel("stepInDecodeOrderByCount:"), stepCount)
	return rv
}

// Determines whether a sample later in decode order can have a presentation timestamp earlier than that of the specified sample cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1390029-sampleswithlaterdecodetimestamps?language=objc
func (s_ SampleCursor) SamplesWithLaterDecodeTimeStampsMayHaveEarlierPresentationTimeStampsThanCursor(cursor ISampleCursor) bool {
	rv := objc.Call[bool](s_, objc.Sel("samplesWithLaterDecodeTimeStampsMayHaveEarlierPresentationTimeStampsThanCursor:"), objc.Ptr(cursor))
	return rv
}

// Moves the cursor by a given delta time on the decode timeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1389152-stepbydecodetime?language=objc
func (s_ SampleCursor) StepByDecodeTimeWasPinned(deltaDecodeTime coremedia.Time, outWasPinned *bool) coremedia.Time {
	rv := objc.Call[coremedia.Time](s_, objc.Sel("stepByDecodeTime:wasPinned:"), deltaDecodeTime, outWasPinned)
	return rv
}

// Compares the relative positions of two sample cursors and returns their relative positions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1390608-comparepositionindecodeorderwith?language=objc
func (s_ SampleCursor) ComparePositionInDecodeOrderWithPositionOfCursor(cursor ISampleCursor) foundation.ComparisonResult {
	rv := objc.Call[foundation.ComparisonResult](s_, objc.Sel("comparePositionInDecodeOrderWithPositionOfCursor:"), objc.Ptr(cursor))
	return rv
}

// Returns the format description of the sample at the cursor’s current position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1390703-copycurrentsampleformatdescripti?language=objc
func (s_ SampleCursor) CopyCurrentSampleFormatDescription() coremedia.FormatDescriptionRef {
	rv := objc.Call[coremedia.FormatDescriptionRef](s_, objc.Sel("copyCurrentSampleFormatDescription"))
	return rv
}

// Moves the cursor by a given delta time on the presentation timeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1387680-stepbypresentationtime?language=objc
func (s_ SampleCursor) StepByPresentationTimeWasPinned(deltaPresentationTime coremedia.Time, outWasPinned *bool) coremedia.Time {
	rv := objc.Call[coremedia.Time](s_, objc.Sel("stepByPresentationTime:wasPinned:"), deltaPresentationTime, outWasPinned)
	return rv
}

// Determines whether a sample earlier in decode order can have a presentation timestamp later than that of the specified sample cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1386558-sampleswithearlierdecodetimestam?language=objc
func (s_ SampleCursor) SamplesWithEarlierDecodeTimeStampsMayHaveLaterPresentationTimeStampsThanCursor(cursor ISampleCursor) bool {
	rv := objc.Call[bool](s_, objc.Sel("samplesWithEarlierDecodeTimeStampsMayHaveLaterPresentationTimeStampsThanCursor:"), objc.Ptr(cursor))
	return rv
}

// The sample range in the storage container to load together with the current sample as a chunk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1390385-currentchunkstoragerange?language=objc
func (s_ SampleCursor) CurrentChunkStorageRange() SampleCursorStorageRange {
	rv := objc.Call[SampleCursorStorageRange](s_, objc.Sel("currentChunkStorageRange"))
	return rv
}

// The decode timestamp of the sample at the current position of the cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1388412-decodetimestamp?language=objc
func (s_ SampleCursor) DecodeTimeStamp() coremedia.Time {
	rv := objc.Call[coremedia.Time](s_, objc.Sel("decodeTimeStamp"))
	return rv
}

// The synchronization information for the current sample for consideration when resynchronizing a decoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1390776-currentsamplesyncinfo?language=objc
func (s_ SampleCursor) CurrentSampleSyncInfo() SampleCursorSyncInfo {
	rv := objc.Call[SampleCursorSyncInfo](s_, objc.Sel("currentSampleSyncInfo"))
	return rv
}

// A value that provides information about the chunk of samples to which the current sample belongs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1387481-currentchunkinfo?language=objc
func (s_ SampleCursor) CurrentChunkInfo() SampleCursorChunkInfo {
	rv := objc.Call[SampleCursorChunkInfo](s_, objc.Sel("currentChunkInfo"))
	return rv
}

// The decode duration of the sample at the cursor’s current position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1389833-currentsampleduration?language=objc
func (s_ SampleCursor) CurrentSampleDuration() coremedia.Time {
	rv := objc.Call[coremedia.Time](s_, objc.Sel("currentSampleDuration"))
	return rv
}

// The number of samples prior to the current sample, in decode order, the decoder requires to achieve a coherent output at the current decode time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1386446-samplesrequiredfordecoderrefresh?language=objc
func (s_ SampleCursor) SamplesRequiredForDecoderRefresh() int {
	rv := objc.Call[int](s_, objc.Sel("samplesRequiredForDecoderRefresh"))
	return rv
}

// The offset and length of the current sample in the current chunk storage URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1386359-currentsamplestoragerange?language=objc
func (s_ SampleCursor) CurrentSampleStorageRange() SampleCursorStorageRange {
	rv := objc.Call[SampleCursorStorageRange](s_, objc.Sel("currentSampleStorageRange"))
	return rv
}

// The presentation timestamp of the sample at the current position of the cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1389740-presentationtimestamp?language=objc
func (s_ SampleCursor) PresentationTimeStamp() coremedia.Time {
	rv := objc.Call[coremedia.Time](s_, objc.Sel("presentationTimeStamp"))
	return rv
}

// The dependency information that describes relationships between a media sample and other media samples in the same sample sequence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1390766-currentsampledependencyinfo?language=objc
func (s_ SampleCursor) CurrentSampleDependencyInfo() SampleCursorDependencyInfo {
	rv := objc.Call[SampleCursorDependencyInfo](s_, objc.Sel("currentSampleDependencyInfo"))
	return rv
}

// A dictionary of dependency-related sample buffer attachments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/3752983-currentsampledependencyattachmen?language=objc
func (s_ SampleCursor) CurrentSampleDependencyAttachments() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](s_, objc.Sel("currentSampleDependencyAttachments"))
	return rv
}

// The index of the current sample within the chunk to which it belongs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1387806-currentsampleindexinchunk?language=objc
func (s_ SampleCursor) CurrentSampleIndexInChunk() int64 {
	rv := objc.Call[int64](s_, objc.Sel("currentSampleIndexInChunk"))
	return rv
}

// The independent decodability information for the audio sample. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/3131266-currentsampleaudiodependencyinfo?language=objc
func (s_ SampleCursor) CurrentSampleAudioDependencyInfo() SampleCursorAudioDependencyInfo {
	rv := objc.Call[SampleCursorAudioDependencyInfo](s_, objc.Sel("currentSampleAudioDependencyInfo"))
	return rv
}

// The URL of the storage container of the current sample and other samples to load in the same operation as a chunk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/1388328-currentchunkstorageurl?language=objc
func (s_ SampleCursor) CurrentChunkStorageURL() foundation.URL {
	rv := objc.Call[foundation.URL](s_, objc.Sel("currentChunkStorageURL"))
	return rv
}
