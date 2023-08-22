// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corefoundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VideoProcessorTimeIntervalCadence] class.
var VideoProcessorTimeIntervalCadenceClass = _VideoProcessorTimeIntervalCadenceClass{objc.GetClass("VNVideoProcessorTimeIntervalCadence")}

type _VideoProcessorTimeIntervalCadenceClass struct {
	objc.Class
}

// An interface definition for the [VideoProcessorTimeIntervalCadence] class.
type IVideoProcessorTimeIntervalCadence interface {
	IVideoProcessorCadence
	TimeInterval() corefoundation.TimeInterval
}

// An object that defines a time-based cadence for processing a video stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessortimeintervalcadence?language=objc
type VideoProcessorTimeIntervalCadence struct {
	VideoProcessorCadence
}

func VideoProcessorTimeIntervalCadenceFrom(ptr unsafe.Pointer) VideoProcessorTimeIntervalCadence {
	return VideoProcessorTimeIntervalCadence{
		VideoProcessorCadence: VideoProcessorCadenceFrom(ptr),
	}
}

func (v_ VideoProcessorTimeIntervalCadence) InitWithTimeInterval(timeInterval corefoundation.TimeInterval) VideoProcessorTimeIntervalCadence {
	rv := objc.Call[VideoProcessorTimeIntervalCadence](v_, objc.Sel("initWithTimeInterval:"), timeInterval)
	return rv
}

// Creates a new time-based cadence with a time interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessortimeintervalcadence/3675686-initwithtimeinterval?language=objc
func NewVideoProcessorTimeIntervalCadenceWithTimeInterval(timeInterval corefoundation.TimeInterval) VideoProcessorTimeIntervalCadence {
	instance := VideoProcessorTimeIntervalCadenceClass.Alloc().InitWithTimeInterval(timeInterval)
	instance.Autorelease()
	return instance
}

func (vc _VideoProcessorTimeIntervalCadenceClass) Alloc() VideoProcessorTimeIntervalCadence {
	rv := objc.Call[VideoProcessorTimeIntervalCadence](vc, objc.Sel("alloc"))
	return rv
}

func VideoProcessorTimeIntervalCadence_Alloc() VideoProcessorTimeIntervalCadence {
	return VideoProcessorTimeIntervalCadenceClass.Alloc()
}

func (vc _VideoProcessorTimeIntervalCadenceClass) New() VideoProcessorTimeIntervalCadence {
	rv := objc.Call[VideoProcessorTimeIntervalCadence](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVideoProcessorTimeIntervalCadence() VideoProcessorTimeIntervalCadence {
	return VideoProcessorTimeIntervalCadenceClass.New()
}

func (v_ VideoProcessorTimeIntervalCadence) Init() VideoProcessorTimeIntervalCadence {
	rv := objc.Call[VideoProcessorTimeIntervalCadence](v_, objc.Sel("init"))
	return rv
}

// The time interval of the cadence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessortimeintervalcadence/3675687-timeinterval?language=objc
func (v_ VideoProcessorTimeIntervalCadence) TimeInterval() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](v_, objc.Sel("timeInterval"))
	return rv
}
