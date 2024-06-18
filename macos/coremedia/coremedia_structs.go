package coremedia

import "github.com/progrium/darwinkit/macos/corefoundation"

// A structure that maps a segment of a source time range to a target time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmtimemapping?language=objc
type TimeMapping struct {
	Source TimeRange
	Target TimeRange
}

// An object that describes a media format descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmformatdescriptionref?language=objc
type FormatDescriptionRef uintptr

// A structure that represents time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmtime?language=objc
type Time struct {
	Value     int64
	Timescale int32
	Flags     uint32
	Epoch     int64
}

// A structure to support custom memory allocation and deallocation for a block used in a block buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmblockbuffercustomblocksource?language=objc
type BlockBufferCustomBlockSource struct {
	Version   uint32
	Pad_cgo_0 [24]byte
}

// A structure that stores the callbacks that perform buffer operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmbuffercallbacks?language=objc
type BufferCallbacks struct {
	Version   uint32
	Pad_cgo_0 [64]byte
}

// A reference to a block buffer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmblockbufferref?language=objc
type BlockBufferRef uintptr

// A structure that stores the handlers that perform buffer operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmbufferhandlers?language=objc
type BufferHandlers struct {
	Version                     uint64
	GetDecodeTimeStamp          uintptr                  // *_Ctype_struct___3
	GetPresentationTimeStamp    uintptr                  // *_Ctype_struct___3
	GetDuration                 uintptr                  // *_Ctype_struct___3
	IsDataReady                 uintptr                  // *_Ctype_struct___4
	Compare                     uintptr                  // *_Ctype_struct___5
	DataBecameReadyNotification corefoundation.StringRef //*_Ctype_struct___CFString
	GetSize                     uintptr                  // *_Ctype_struct___6
}

// A structure that represents a time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmtimerange?language=objc
type TimeRange struct {
	Start    Time
	Duration Time
}

// A reference to an object that provides a simple lockless queue of elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmsimplequeueref?language=objc
type SimpleQueueRef uintptr

// A collection of timing information for a sample in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmsampletiminginfo?language=objc
type SampleTimingInfo struct {
	Duration              Time
	PresentationTimeStamp Time
	DecodeTimeStamp       Time
}

// A reference to an immutable sample buffer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmsamplebufferref?language=objc
type SampleBufferRef uintptr

// An object that represents a source of time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmclockref?language=objc
type ClockRef uintptr

// A structure that represents video dimensions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmvideodimensions?language=objc
type VideoDimensions struct {
	Width  int32
	Height int32
}

// A reference to a memory pool object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmmemorypoolref?language=objc
type MemoryPoolRef uintptr

// A model of a timeline under application control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmtimebaseref?language=objc
type TimebaseRef uintptr

// A reference to a buffer queue object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmbufferqueueref?language=objc
type BufferQueueRef uintptr
