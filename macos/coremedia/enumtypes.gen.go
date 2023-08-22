// AUTO-GENERATED CODE, DO NOT MODIFY

package coremedia

import "github.com/progrium/macdriver/macos/corefoundation"

// A type for mask bits that represent parts of an audio format description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmaudioformatdescriptionmask?language=objc
type AudioFormatDescriptionMask uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmbaseclassversion?language=objc
type BaseClassVersion uint

// A type for parameters that contain block buffer feature and control flags. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmblockbufferflags?language=objc
type BlockBufferFlags uint32

// A type to specify conditions to associate with a buffer queue trigger. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmbufferqueuetriggercondition?language=objc
type BufferQueueTriggerCondition int32

// A closed caption format type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmclosedcaptionformattype?language=objc
type ClosedCaptionFormatType uint

// A datatype that represents an item count. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmitemcount?language=objc
type ItemCount corefoundation.Index

// A datatype that represents an item index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmitemindex?language=objc
type ItemIndex corefoundation.Index

// A pixel format type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmpixelformattype?language=objc
type PixelFormatType uint

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmstructversion?language=objc
type StructVersion uint

// An integer value that describes the display mode flags for text media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmtextdisplayflags?language=objc
type TextDisplayFlags uint32

// A text format type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmtextformattype?language=objc
type TextFormatType uint

// An integer value that describes the justification modes for text media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmtextjustificationvalue?language=objc
type TextJustificationValue int8

// An epoch for a time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmtimeepoch?language=objc
type TimeEpoch int64

// A structure that defines the flags for a time value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmtimeflags?language=objc
type TimeFlags uint32

const (
	kCMTimeFlags_HasBeenRounded        TimeFlags = 2
	kCMTimeFlags_ImpliedValueFlagsMask TimeFlags = 28
	kCMTimeFlags_Indefinite            TimeFlags = 16
	kCMTimeFlags_NegativeInfinity      TimeFlags = 8
	kCMTimeFlags_PositiveInfinity      TimeFlags = 4
	kCMTimeFlags_Valid                 TimeFlags = 1
)

// An enumeration of rounding methods to use when performing time calculations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmtimeroundingmethod?language=objc
type TimeRoundingMethod uint32

const (
	kCMTimeRoundingMethod_Default                     TimeRoundingMethod = 1
	kCMTimeRoundingMethod_QuickTime                   TimeRoundingMethod = 4
	kCMTimeRoundingMethod_RoundAwayFromZero           TimeRoundingMethod = 3
	kCMTimeRoundingMethod_RoundHalfAwayFromZero       TimeRoundingMethod = 1
	kCMTimeRoundingMethod_RoundTowardNegativeInfinity TimeRoundingMethod = 6
	kCMTimeRoundingMethod_RoundTowardPositiveInfinity TimeRoundingMethod = 5
	kCMTimeRoundingMethod_RoundTowardZero             TimeRoundingMethod = 2
)

// An integer timescale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmtimescale?language=objc
type TimeScale int32

// An integer time value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmtimevalue?language=objc
type TimeValue int64

// A video codec type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmvideocodectype?language=objc
type VideoCodecType uint
