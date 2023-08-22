// AUTO-GENERATED CODE, DO NOT MODIFY

package coremediaio

import (
	"unsafe"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmiostreamscheduledoutputnotificationproc?language=objc
type StreamScheduledOutputNotificationProc = func(sequenceNumberOfBufferThatWasOutput uint64, outputHostTime uint64, scheduledOutputNotificationRefCon unsafe.Pointer)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmiodevicegetsmptetimeproc?language=objc
type DeviceGetSMPTETimeProc = func(refCon unsafe.Pointer, frameNumber *uint64, isDropFrame *bool, tolerance *uint32) uint

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmiodevicestreamqueuealteredproc?language=objc
type DeviceStreamQueueAlteredProc = func(streamID StreamID, token unsafe.Pointer, refCon unsafe.Pointer)
