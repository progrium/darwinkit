// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DetectTrajectoriesRequest] class.
var DetectTrajectoriesRequestClass = _DetectTrajectoriesRequestClass{objc.GetClass("VNDetectTrajectoriesRequest")}

type _DetectTrajectoriesRequestClass struct {
	objc.Class
}

// An interface definition for the [DetectTrajectoriesRequest] class.
type IDetectTrajectoriesRequest interface {
	IStatefulRequest
	ObjectMaximumNormalizedRadius() float64
	SetObjectMaximumNormalizedRadius(value float64)
	TargetFrameTime() coremedia.Time
	SetTargetFrameTime(value coremedia.Time)
	TrajectoryLength() int
	ObjectMinimumNormalizedRadius() float64
	SetObjectMinimumNormalizedRadius(value float64)
}

// A request that detects the trajectories of shapes moving along a parabolic path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest?language=objc
type DetectTrajectoriesRequest struct {
	StatefulRequest
}

func DetectTrajectoriesRequestFrom(ptr unsafe.Pointer) DetectTrajectoriesRequest {
	return DetectTrajectoriesRequest{
		StatefulRequest: StatefulRequestFrom(ptr),
	}
}

func (d_ DetectTrajectoriesRequest) InitWithFrameAnalysisSpacingTrajectoryLengthCompletionHandler(frameAnalysisSpacing coremedia.Time, trajectoryLength int, completionHandler RequestCompletionHandler) DetectTrajectoriesRequest {
	rv := objc.Call[DetectTrajectoriesRequest](d_, objc.Sel("initWithFrameAnalysisSpacing:trajectoryLength:completionHandler:"), frameAnalysisSpacing, trajectoryLength, completionHandler)
	return rv
}

// Creates a new request to detect trajectories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/3600612-initwithframeanalysisspacing?language=objc
func NewDetectTrajectoriesRequestWithFrameAnalysisSpacingTrajectoryLengthCompletionHandler(frameAnalysisSpacing coremedia.Time, trajectoryLength int, completionHandler RequestCompletionHandler) DetectTrajectoriesRequest {
	instance := DetectTrajectoriesRequestClass.Alloc().InitWithFrameAnalysisSpacingTrajectoryLengthCompletionHandler(frameAnalysisSpacing, trajectoryLength, completionHandler)
	instance.Autorelease()
	return instance
}

func (dc _DetectTrajectoriesRequestClass) Alloc() DetectTrajectoriesRequest {
	rv := objc.Call[DetectTrajectoriesRequest](dc, objc.Sel("alloc"))
	return rv
}

func DetectTrajectoriesRequest_Alloc() DetectTrajectoriesRequest {
	return DetectTrajectoriesRequestClass.Alloc()
}

func (dc _DetectTrajectoriesRequestClass) New() DetectTrajectoriesRequest {
	rv := objc.Call[DetectTrajectoriesRequest](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDetectTrajectoriesRequest() DetectTrajectoriesRequest {
	return DetectTrajectoriesRequestClass.New()
}

func (d_ DetectTrajectoriesRequest) Init() DetectTrajectoriesRequest {
	rv := objc.Call[DetectTrajectoriesRequest](d_, objc.Sel("init"))
	return rv
}

func (d_ DetectTrajectoriesRequest) InitWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing coremedia.Time, completionHandler RequestCompletionHandler) DetectTrajectoriesRequest {
	rv := objc.Call[DetectTrajectoriesRequest](d_, objc.Sel("initWithFrameAnalysisSpacing:completionHandler:"), frameAnalysisSpacing, completionHandler)
	return rv
}

// Initializes a video-based request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnstatefulrequest/3564828-initwithframeanalysisspacing?language=objc
func NewDetectTrajectoriesRequestWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing coremedia.Time, completionHandler RequestCompletionHandler) DetectTrajectoriesRequest {
	instance := DetectTrajectoriesRequestClass.Alloc().InitWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing, completionHandler)
	instance.Autorelease()
	return instance
}

func (d_ DetectTrajectoriesRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) DetectTrajectoriesRequest {
	rv := objc.Call[DetectTrajectoriesRequest](d_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewDetectTrajectoriesRequestWithCompletionHandler(completionHandler RequestCompletionHandler) DetectTrajectoriesRequest {
	instance := DetectTrajectoriesRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}

// The maximum radius of the bounding circle of the object to track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/3675670-objectmaximumnormalizedradius?language=objc
func (d_ DetectTrajectoriesRequest) ObjectMaximumNormalizedRadius() float64 {
	rv := objc.Call[float64](d_, objc.Sel("objectMaximumNormalizedRadius"))
	return rv
}

// The maximum radius of the bounding circle of the object to track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/3675670-objectmaximumnormalizedradius?language=objc
func (d_ DetectTrajectoriesRequest) SetObjectMaximumNormalizedRadius(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setObjectMaximumNormalizedRadius:"), value)
}

// The requested target frame time for processing trajectory detection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/3750981-targetframetime?language=objc
func (d_ DetectTrajectoriesRequest) TargetFrameTime() coremedia.Time {
	rv := objc.Call[coremedia.Time](d_, objc.Sel("targetFrameTime"))
	return rv
}

// The requested target frame time for processing trajectory detection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/3750981-targetframetime?language=objc
func (d_ DetectTrajectoriesRequest) SetTargetFrameTime(value coremedia.Time) {
	objc.Call[objc.Void](d_, objc.Sel("setTargetFrameTime:"), value)
}

// The number of points to detect before calculating a trajectory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/3675673-trajectorylength?language=objc
func (d_ DetectTrajectoriesRequest) TrajectoryLength() int {
	rv := objc.Call[int](d_, objc.Sel("trajectoryLength"))
	return rv
}

// The minimum radius of the bounding circle of the object to track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/3675671-objectminimumnormalizedradius?language=objc
func (d_ DetectTrajectoriesRequest) ObjectMinimumNormalizedRadius() float64 {
	rv := objc.Call[float64](d_, objc.Sel("objectMinimumNormalizedRadius"))
	return rv
}

// The minimum radius of the bounding circle of the object to track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/3675671-objectminimumnormalizedradius?language=objc
func (d_ DetectTrajectoriesRequest) SetObjectMinimumNormalizedRadius(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setObjectMinimumNormalizedRadius:"), value)
}
