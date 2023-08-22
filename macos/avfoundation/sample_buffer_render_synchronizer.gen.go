// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SampleBufferRenderSynchronizer] class.
var SampleBufferRenderSynchronizerClass = _SampleBufferRenderSynchronizerClass{objc.GetClass("AVSampleBufferRenderSynchronizer")}

type _SampleBufferRenderSynchronizerClass struct {
	objc.Class
}

// An interface definition for the [SampleBufferRenderSynchronizer] class.
type ISampleBufferRenderSynchronizer interface {
	objc.IObject
	RemoveTimeObserver(observer objc.IObject)
	SetRateTimeAtHostTime(rate float64, time coremedia.Time, hostTime coremedia.Time)
	AddBoundaryTimeObserverForTimesQueueUsingBlock(times []foundation.IValue, queue dispatch.Queue, block func()) objc.Object
	AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval coremedia.Time, queue dispatch.Queue, block func(time coremedia.Time)) objc.Object
	RemoveRendererAtTimeCompletionHandler(renderer PQueuedSampleBufferRendering, time coremedia.Time, completionHandler func(didRemoveRenderer bool))
	RemoveRendererObjectAtTimeCompletionHandler(rendererObject objc.IObject, time coremedia.Time, completionHandler func(didRemoveRenderer bool))
	AddRenderer(renderer PQueuedSampleBufferRendering)
	AddRendererObject(rendererObject objc.IObject)
	CurrentTime() coremedia.Time
	DelaysRateChangeUntilHasSufficientMediaData() bool
	SetDelaysRateChangeUntilHasSufficientMediaData(value bool)
	Rate() float64
	SetRate(value float64)
	Timebase() coremedia.TimebaseRef
	Renderers() []QueuedSampleBufferRenderingWrapper
}

// An object used to synchronize multiple queued sample buffers to a single timeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer?language=objc
type SampleBufferRenderSynchronizer struct {
	objc.Object
}

func SampleBufferRenderSynchronizerFrom(ptr unsafe.Pointer) SampleBufferRenderSynchronizer {
	return SampleBufferRenderSynchronizer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SampleBufferRenderSynchronizerClass) Alloc() SampleBufferRenderSynchronizer {
	rv := objc.Call[SampleBufferRenderSynchronizer](sc, objc.Sel("alloc"))
	return rv
}

func SampleBufferRenderSynchronizer_Alloc() SampleBufferRenderSynchronizer {
	return SampleBufferRenderSynchronizerClass.Alloc()
}

func (sc _SampleBufferRenderSynchronizerClass) New() SampleBufferRenderSynchronizer {
	rv := objc.Call[SampleBufferRenderSynchronizer](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSampleBufferRenderSynchronizer() SampleBufferRenderSynchronizer {
	return SampleBufferRenderSynchronizerClass.New()
}

func (s_ SampleBufferRenderSynchronizer) Init() SampleBufferRenderSynchronizer {
	rv := objc.Call[SampleBufferRenderSynchronizer](s_, objc.Sel("init"))
	return rv
}

// Cancels the specified time observer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867829-removetimeobserver?language=objc
func (s_ SampleBufferRenderSynchronizer) RemoveTimeObserver(observer objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("removeTimeObserver:"), observer)
}

// Sets the playback rate and the relationship between the current time and host time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/3726157-setrate?language=objc
func (s_ SampleBufferRenderSynchronizer) SetRateTimeAtHostTime(rate float64, time coremedia.Time, hostTime coremedia.Time) {
	objc.Call[objc.Void](s_, objc.Sel("setRate:time:atHostTime:"), rate, time, hostTime)
}

// Requests invocation of a block when specified times are traversed during normal rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867824-addboundarytimeobserverfortimes?language=objc
func (s_ SampleBufferRenderSynchronizer) AddBoundaryTimeObserverForTimesQueueUsingBlock(times []foundation.IValue, queue dispatch.Queue, block func()) objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("addBoundaryTimeObserverForTimes:queue:usingBlock:"), times, queue, block)
	return rv
}

// Requests invocation of a block during rendering at specified time intervals. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867825-addperiodictimeobserverforinterv?language=objc
func (s_ SampleBufferRenderSynchronizer) AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval coremedia.Time, queue dispatch.Queue, block func(time coremedia.Time)) objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("addPeriodicTimeObserverForInterval:queue:usingBlock:"), interval, queue, block)
	return rv
}

// Removes a renderer from the synchronizer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867826-removerenderer?language=objc
func (s_ SampleBufferRenderSynchronizer) RemoveRendererAtTimeCompletionHandler(renderer PQueuedSampleBufferRendering, time coremedia.Time, completionHandler func(didRemoveRenderer bool)) {
	po0 := objc.WrapAsProtocol("AVQueuedSampleBufferRendering", renderer)
	objc.Call[objc.Void](s_, objc.Sel("removeRenderer:atTime:completionHandler:"), po0, time, completionHandler)
}

// Removes a renderer from the synchronizer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867826-removerenderer?language=objc
func (s_ SampleBufferRenderSynchronizer) RemoveRendererObjectAtTimeCompletionHandler(rendererObject objc.IObject, time coremedia.Time, completionHandler func(didRemoveRenderer bool)) {
	objc.Call[objc.Void](s_, objc.Sel("removeRenderer:atTime:completionHandler:"), objc.Ptr(rendererObject), time, completionHandler)
}

// Adds a renderer to the list of renderers under the synchronizer's control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867828-addrenderer?language=objc
func (s_ SampleBufferRenderSynchronizer) AddRenderer(renderer PQueuedSampleBufferRendering) {
	po0 := objc.WrapAsProtocol("AVQueuedSampleBufferRendering", renderer)
	objc.Call[objc.Void](s_, objc.Sel("addRenderer:"), po0)
}

// Adds a renderer to the list of renderers under the synchronizer's control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867828-addrenderer?language=objc
func (s_ SampleBufferRenderSynchronizer) AddRendererObject(rendererObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("addRenderer:"), objc.Ptr(rendererObject))
}

// Returns the current time of the synchronizer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/3022467-currenttime?language=objc
func (s_ SampleBufferRenderSynchronizer) CurrentTime() coremedia.Time {
	rv := objc.Call[coremedia.Time](s_, objc.Sel("currentTime"))
	return rv
}

// A Boolean value that Indicates whether the playback should start immediately on rate change requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/3726156-delaysratechangeuntilhassufficie?language=objc
func (s_ SampleBufferRenderSynchronizer) DelaysRateChangeUntilHasSufficientMediaData() bool {
	rv := objc.Call[bool](s_, objc.Sel("delaysRateChangeUntilHasSufficientMediaData"))
	return rv
}

// A Boolean value that Indicates whether the playback should start immediately on rate change requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/3726156-delaysratechangeuntilhassufficie?language=objc
func (s_ SampleBufferRenderSynchronizer) SetDelaysRateChangeUntilHasSufficientMediaData(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setDelaysRateChangeUntilHasSufficientMediaData:"), value)
}

// The current playback rate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867823-rate?language=objc
func (s_ SampleBufferRenderSynchronizer) Rate() float64 {
	rv := objc.Call[float64](s_, objc.Sel("rate"))
	return rv
}

// The current playback rate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867823-rate?language=objc
func (s_ SampleBufferRenderSynchronizer) SetRate(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setRate:"), value)
}

// The synchronizer’s rendering timebase which determines how it interprets timestamps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867822-timebase?language=objc
func (s_ SampleBufferRenderSynchronizer) Timebase() coremedia.TimebaseRef {
	rv := objc.Call[coremedia.TimebaseRef](s_, objc.Sel("timebase"))
	return rv
}

// An array of queued sample buffer renderers currently attached to the synchronizer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867827-renderers?language=objc
func (s_ SampleBufferRenderSynchronizer) Renderers() []QueuedSampleBufferRenderingWrapper {
	rv := objc.Call[[]QueuedSampleBufferRenderingWrapper](s_, objc.Sel("renderers"))
	return rv
}
