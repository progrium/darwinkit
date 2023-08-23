// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VideoProcessor] class.
var VideoProcessorClass = _VideoProcessorClass{objc.GetClass("VNVideoProcessor")}

type _VideoProcessorClass struct {
	objc.Class
}

// An interface definition for the [VideoProcessor] class.
type IVideoProcessor interface {
	objc.IObject
	AddRequestProcessingOptionsError(request IRequest, processingOptions IVideoProcessorRequestProcessingOptions, error foundation.IError) bool
	AnalyzeTimeRangeError(timeRange coremedia.TimeRange, error foundation.IError) bool
	RemoveRequestError(request IRequest, error foundation.IError) bool
	Cancel()
}

// An object that performs offline analysis of video content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessor?language=objc
type VideoProcessor struct {
	objc.Object
}

func VideoProcessorFrom(ptr unsafe.Pointer) VideoProcessor {
	return VideoProcessor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (v_ VideoProcessor) InitWithURL(videoURL foundation.IURL) VideoProcessor {
	rv := objc.Call[VideoProcessor](v_, objc.Sel("initWithURL:"), objc.Ptr(videoURL))
	return rv
}

// Creates a video processor to perform Vision requests against the specified video asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessor/3548386-initwithurl?language=objc
func NewVideoProcessorWithURL(videoURL foundation.IURL) VideoProcessor {
	instance := VideoProcessorClass.Alloc().InitWithURL(videoURL)
	instance.Autorelease()
	return instance
}

func (vc _VideoProcessorClass) Alloc() VideoProcessor {
	rv := objc.Call[VideoProcessor](vc, objc.Sel("alloc"))
	return rv
}

func VideoProcessor_Alloc() VideoProcessor {
	return VideoProcessorClass.Alloc()
}

func (vc _VideoProcessorClass) New() VideoProcessor {
	rv := objc.Call[VideoProcessor](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVideoProcessor() VideoProcessor {
	return VideoProcessorClass.New()
}

func (v_ VideoProcessor) Init() VideoProcessor {
	rv := objc.Call[VideoProcessor](v_, objc.Sel("init"))
	return rv
}

// Adds a request with processing options to the video processor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessor/3675677-addrequest?language=objc
func (v_ VideoProcessor) AddRequestProcessingOptionsError(request IRequest, processingOptions IVideoProcessorRequestProcessingOptions, error foundation.IError) bool {
	rv := objc.Call[bool](v_, objc.Sel("addRequest:processingOptions:error:"), objc.Ptr(request), objc.Ptr(processingOptions), objc.Ptr(error))
	return rv
}

// Analyzes a time range of video content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessor/3675678-analyzetimerange?language=objc
func (v_ VideoProcessor) AnalyzeTimeRangeError(timeRange coremedia.TimeRange, error foundation.IError) bool {
	rv := objc.Call[bool](v_, objc.Sel("analyzeTimeRange:error:"), timeRange, objc.Ptr(error))
	return rv
}

// Removes a Vision request from the video processor’s request queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessor/3548387-removerequest?language=objc
func (v_ VideoProcessor) RemoveRequestError(request IRequest, error foundation.IError) bool {
	rv := objc.Call[bool](v_, objc.Sel("removeRequest:error:"), objc.Ptr(request), objc.Ptr(error))
	return rv
}

// Cancels the video processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessor/3548385-cancel?language=objc
func (v_ VideoProcessor) Cancel() {
	objc.Call[objc.Void](v_, objc.Sel("cancel"))
}
