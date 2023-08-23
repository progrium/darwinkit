// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// Methods you can implement to enqueue sample buffers for presentation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avqueuedsamplebufferrendering?language=objc
type PQueuedSampleBufferRendering interface {
	// optional
	StopRequestingMediaData()
	HasStopRequestingMediaData() bool

	// optional
	RequestMediaDataWhenReadyOnQueueUsingBlock(queue dispatch.Queue, block func())
	HasRequestMediaDataWhenReadyOnQueueUsingBlock() bool

	// optional
	Flush()
	HasFlush() bool

	// optional
	EnqueueSampleBuffer(sampleBuffer coremedia.SampleBufferRef)
	HasEnqueueSampleBuffer() bool

	// optional
	IsReadyForMoreMediaData() bool
	HasIsReadyForMoreMediaData() bool

	// optional
	Timebase() coremedia.TimebaseRef
	HasTimebase() bool

	// optional
	HasSufficientMediaDataForReliablePlaybackStart() bool
	HasHasSufficientMediaDataForReliablePlaybackStart() bool
}

// A concrete type wrapper for the [PQueuedSampleBufferRendering] protocol.
type QueuedSampleBufferRenderingWrapper struct {
	objc.Object
}

func (q_ QueuedSampleBufferRenderingWrapper) HasStopRequestingMediaData() bool {
	return q_.RespondsToSelector(objc.Sel("stopRequestingMediaData"))
}

// Cancels any current requestMediaDataWhenReadyOnQueue:usingBlock: call. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avqueuedsamplebufferrendering/2867642-stoprequestingmediadata?language=objc
func (q_ QueuedSampleBufferRenderingWrapper) StopRequestingMediaData() {
	objc.Call[objc.Void](q_, objc.Sel("stopRequestingMediaData"))
}

func (q_ QueuedSampleBufferRenderingWrapper) HasRequestMediaDataWhenReadyOnQueueUsingBlock() bool {
	return q_.RespondsToSelector(objc.Sel("requestMediaDataWhenReadyOnQueue:usingBlock:"))
}

// Tells the target to invoke a client-supplied block in order to gather sample buffers for playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avqueuedsamplebufferrendering/2867643-requestmediadatawhenreadyonqueue?language=objc
func (q_ QueuedSampleBufferRenderingWrapper) RequestMediaDataWhenReadyOnQueueUsingBlock(queue dispatch.Queue, block func()) {
	objc.Call[objc.Void](q_, objc.Sel("requestMediaDataWhenReadyOnQueue:usingBlock:"), queue, block)
}

func (q_ QueuedSampleBufferRenderingWrapper) HasFlush() bool {
	return q_.RespondsToSelector(objc.Sel("flush"))
}

// Discards all pending enqueued sample buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avqueuedsamplebufferrendering/2867639-flush?language=objc
func (q_ QueuedSampleBufferRenderingWrapper) Flush() {
	objc.Call[objc.Void](q_, objc.Sel("flush"))
}

func (q_ QueuedSampleBufferRenderingWrapper) HasEnqueueSampleBuffer() bool {
	return q_.RespondsToSelector(objc.Sel("enqueueSampleBuffer:"))
}

// Sends a sample buffer to the queue for rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avqueuedsamplebufferrendering/2867641-enqueuesamplebuffer?language=objc
func (q_ QueuedSampleBufferRenderingWrapper) EnqueueSampleBuffer(sampleBuffer coremedia.SampleBufferRef) {
	objc.Call[objc.Void](q_, objc.Sel("enqueueSampleBuffer:"), sampleBuffer)
}

func (q_ QueuedSampleBufferRenderingWrapper) HasIsReadyForMoreMediaData() bool {
	return q_.RespondsToSelector(objc.Sel("isReadyForMoreMediaData"))
}

// A Boolean value that indicates whether the receiver is able to accept more sample buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avqueuedsamplebufferrendering/2867638-readyformoremediadata?language=objc
func (q_ QueuedSampleBufferRenderingWrapper) IsReadyForMoreMediaData() bool {
	rv := objc.Call[bool](q_, objc.Sel("isReadyForMoreMediaData"))
	return rv
}

func (q_ QueuedSampleBufferRenderingWrapper) HasTimebase() bool {
	return q_.RespondsToSelector(objc.Sel("timebase"))
}

// The timebase for a renderer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avqueuedsamplebufferrendering/2867640-timebase?language=objc
func (q_ QueuedSampleBufferRenderingWrapper) Timebase() coremedia.TimebaseRef {
	rv := objc.Call[coremedia.TimebaseRef](q_, objc.Sel("timebase"))
	return rv
}

func (q_ QueuedSampleBufferRenderingWrapper) HasHasSufficientMediaDataForReliablePlaybackStart() bool {
	return q_.RespondsToSelector(objc.Sel("hasSufficientMediaDataForReliablePlaybackStart"))
}

// A Boolean value that indicates whether the enqued media meets the required preroll level for reliable playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avqueuedsamplebufferrendering/3726153-hassufficientmediadataforreliabl?language=objc
func (q_ QueuedSampleBufferRenderingWrapper) HasSufficientMediaDataForReliablePlaybackStart() bool {
	rv := objc.Call[bool](q_, objc.Sel("hasSufficientMediaDataForReliablePlaybackStart"))
	return rv
}
