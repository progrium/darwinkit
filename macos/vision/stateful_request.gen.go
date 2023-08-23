// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [StatefulRequest] class.
var StatefulRequestClass = _StatefulRequestClass{objc.GetClass("VNStatefulRequest")}

type _StatefulRequestClass struct {
	objc.Class
}

// An interface definition for the [StatefulRequest] class.
type IStatefulRequest interface {
	IImageBasedRequest
	MinimumLatencyFrameCount() int
	FrameAnalysisSpacing() coremedia.Time
}

// An abstract request type that builds evidence of a condition over time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnstatefulrequest?language=objc
type StatefulRequest struct {
	ImageBasedRequest
}

func StatefulRequestFrom(ptr unsafe.Pointer) StatefulRequest {
	return StatefulRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (s_ StatefulRequest) InitWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing coremedia.Time, completionHandler RequestCompletionHandler) StatefulRequest {
	rv := objc.Call[StatefulRequest](s_, objc.Sel("initWithFrameAnalysisSpacing:completionHandler:"), frameAnalysisSpacing, completionHandler)
	return rv
}

// Initializes a video-based request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnstatefulrequest/3564828-initwithframeanalysisspacing?language=objc
func NewStatefulRequestWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing coremedia.Time, completionHandler RequestCompletionHandler) StatefulRequest {
	instance := StatefulRequestClass.Alloc().InitWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing, completionHandler)
	instance.Autorelease()
	return instance
}

func (sc _StatefulRequestClass) Alloc() StatefulRequest {
	rv := objc.Call[StatefulRequest](sc, objc.Sel("alloc"))
	return rv
}

func StatefulRequest_Alloc() StatefulRequest {
	return StatefulRequestClass.Alloc()
}

func (sc _StatefulRequestClass) New() StatefulRequest {
	rv := objc.Call[StatefulRequest](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStatefulRequest() StatefulRequest {
	return StatefulRequestClass.New()
}

func (s_ StatefulRequest) Init() StatefulRequest {
	rv := objc.Call[StatefulRequest](s_, objc.Sel("init"))
	return rv
}

func (s_ StatefulRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) StatefulRequest {
	rv := objc.Call[StatefulRequest](s_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewStatefulRequestWithCompletionHandler(completionHandler RequestCompletionHandler) StatefulRequest {
	instance := StatefulRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}

// The minimum number of frames a request processes before reporting an observation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnstatefulrequest/3564829-minimumlatencyframecount?language=objc
func (s_ StatefulRequest) MinimumLatencyFrameCount() int {
	rv := objc.Call[int](s_, objc.Sel("minimumLatencyFrameCount"))
	return rv
}

// A time value that indicates the interval between analysis operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnstatefulrequest/3675676-frameanalysisspacing?language=objc
func (s_ StatefulRequest) FrameAnalysisSpacing() coremedia.Time {
	rv := objc.Call[coremedia.Time](s_, objc.Sel("frameAnalysisSpacing"))
	return rv
}
