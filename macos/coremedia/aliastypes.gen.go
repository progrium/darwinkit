// AUTO-GENERATED CODE, DO NOT MODIFY

package coremedia

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corefoundation"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmbuffergetbooleanhandler?language=objc
type BufferGetBooleanHandler = func(buf BufferRef) bool

// A block the system calls to make the sample buffer ready for use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmsamplebuffermakedatareadyhandler?language=objc
type SampleBufferMakeDataReadyHandler = func(sbuf SampleBufferRef) uint

// Callback that returns a Boolean value from a CMBuffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmbuffergetbooleancallback?language=objc
type BufferGetBooleanCallback = func(buf BufferRef, refcon unsafe.Pointer) bool

// A callback for the system to invoke when a trigger condition becomes true. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmbufferqueuetriggercallback?language=objc
type BufferQueueTriggerCallback = func(triggerRefcon unsafe.Pointer, triggerToken unsafe.Pointer)

// Client callback called by CMSampleBufferInvalidate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmsamplebufferinvalidatehandler?language=objc
type SampleBufferInvalidateHandler = func(sbuf SampleBufferRef)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmbuffergetsizehandler?language=objc
type BufferGetSizeHandler = func(buf BufferRef) uint

// Client callback called by CMSampleBufferInvalidate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmsamplebufferinvalidatecallback?language=objc
type SampleBufferInvalidateCallback = func(sbuf SampleBufferRef, invalidateRefCon uint64)

// A type alias for a callback that tests whether a buffer is in a valid state to add to a queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmbuffervalidationcallback?language=objc
type BufferValidationCallback = func(queue BufferQueueRef, buf BufferRef, validationRefCon unsafe.Pointer) uint

// A client callback that returns a size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmbuffergetsizecallback?language=objc
type BufferGetSizeCallback = func(buf BufferRef, refcon unsafe.Pointer) uint

// Client callback called by CMSampleBufferMakeDataReady. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmsamplebuffermakedatareadycallback?language=objc
type SampleBufferMakeDataReadyCallback = func(sbuf SampleBufferRef, makeDataReadyRefcon unsafe.Pointer) uint

// A type alias for a trigger handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmbufferqueuetriggerhandler?language=objc
type BufferQueueTriggerHandler = func(triggerToken unsafe.Pointer)

// Callback that returns a CMTime from a CMBuffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmbuffergettimecallback?language=objc
type BufferGetTimeCallback = func(buf BufferRef, refcon unsafe.Pointer) Time

// A type alias for a handler that tests whether a buffer is in a valid state to add to a queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmbuffervalidationhandler?language=objc
type BufferValidationHandler = func(queue BufferQueueRef, buf BufferRef) uint

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmbuffergettimehandler?language=objc
type BufferGetTimeHandler = func(buf BufferRef) Time

// Callback that compares one CMBuffer with another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmbuffercomparecallback?language=objc
type BufferCompareCallback = func(buf1 BufferRef, buf2 BufferRef, refcon unsafe.Pointer) corefoundation.ComparisonResult

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremedia/cmbuffercomparehandler?language=objc
type BufferCompareHandler = func(buf1 BufferRef, buf2 BufferRef) corefoundation.ComparisonResult
