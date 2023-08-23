// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VideoProcessorFrameRateCadence] class.
var VideoProcessorFrameRateCadenceClass = _VideoProcessorFrameRateCadenceClass{objc.GetClass("VNVideoProcessorFrameRateCadence")}

type _VideoProcessorFrameRateCadenceClass struct {
	objc.Class
}

// An interface definition for the [VideoProcessorFrameRateCadence] class.
type IVideoProcessorFrameRateCadence interface {
	IVideoProcessorCadence
	FrameRate() int
}

// An object that defines a frame-based cadence for processing a video stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessorframeratecadence?language=objc
type VideoProcessorFrameRateCadence struct {
	VideoProcessorCadence
}

func VideoProcessorFrameRateCadenceFrom(ptr unsafe.Pointer) VideoProcessorFrameRateCadence {
	return VideoProcessorFrameRateCadence{
		VideoProcessorCadence: VideoProcessorCadenceFrom(ptr),
	}
}

func (v_ VideoProcessorFrameRateCadence) InitWithFrameRate(frameRate int) VideoProcessorFrameRateCadence {
	rv := objc.Call[VideoProcessorFrameRateCadence](v_, objc.Sel("initWithFrameRate:"), frameRate)
	return rv
}

// Creates a new frame-based cadence with a frame rate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessorframeratecadence/3675682-initwithframerate?language=objc
func NewVideoProcessorFrameRateCadenceWithFrameRate(frameRate int) VideoProcessorFrameRateCadence {
	instance := VideoProcessorFrameRateCadenceClass.Alloc().InitWithFrameRate(frameRate)
	instance.Autorelease()
	return instance
}

func (vc _VideoProcessorFrameRateCadenceClass) Alloc() VideoProcessorFrameRateCadence {
	rv := objc.Call[VideoProcessorFrameRateCadence](vc, objc.Sel("alloc"))
	return rv
}

func VideoProcessorFrameRateCadence_Alloc() VideoProcessorFrameRateCadence {
	return VideoProcessorFrameRateCadenceClass.Alloc()
}

func (vc _VideoProcessorFrameRateCadenceClass) New() VideoProcessorFrameRateCadence {
	rv := objc.Call[VideoProcessorFrameRateCadence](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVideoProcessorFrameRateCadence() VideoProcessorFrameRateCadence {
	return VideoProcessorFrameRateCadenceClass.New()
}

func (v_ VideoProcessorFrameRateCadence) Init() VideoProcessorFrameRateCadence {
	rv := objc.Call[VideoProcessorFrameRateCadence](v_, objc.Sel("init"))
	return rv
}

// The frame rate at which to process video. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessorframeratecadence/3675681-framerate?language=objc
func (v_ VideoProcessorFrameRateCadence) FrameRate() int {
	rv := objc.Call[int](v_, objc.Sel("frameRate"))
	return rv
}
