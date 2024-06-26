// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/coremedia"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [MutableTimedMetadataGroup] class.
var MutableTimedMetadataGroupClass = _MutableTimedMetadataGroupClass{objc.GetClass("AVMutableTimedMetadataGroup")}

type _MutableTimedMetadataGroupClass struct {
	objc.Class
}

// An interface definition for the [MutableTimedMetadataGroup] class.
type IMutableTimedMetadataGroup interface {
	ITimedMetadataGroup
	SetItems(value []IMetadataItem)
	SetTimeRange(value coremedia.TimeRange)
}

// A mutable collection of metadata items that are valid for use during a specific time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutabletimedmetadatagroup?language=objc
type MutableTimedMetadataGroup struct {
	TimedMetadataGroup
}

func MutableTimedMetadataGroupFrom(ptr unsafe.Pointer) MutableTimedMetadataGroup {
	return MutableTimedMetadataGroup{
		TimedMetadataGroup: TimedMetadataGroupFrom(ptr),
	}
}

func (mc _MutableTimedMetadataGroupClass) Alloc() MutableTimedMetadataGroup {
	rv := objc.Call[MutableTimedMetadataGroup](mc, objc.Sel("alloc"))
	return rv
}

func (mc _MutableTimedMetadataGroupClass) New() MutableTimedMetadataGroup {
	rv := objc.Call[MutableTimedMetadataGroup](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableTimedMetadataGroup() MutableTimedMetadataGroup {
	return MutableTimedMetadataGroupClass.New()
}

func (m_ MutableTimedMetadataGroup) Init() MutableTimedMetadataGroup {
	rv := objc.Call[MutableTimedMetadataGroup](m_, objc.Sel("init"))
	return rv
}

func (m_ MutableTimedMetadataGroup) InitWithItemsTimeRange(items []IMetadataItem, timeRange coremedia.TimeRange) MutableTimedMetadataGroup {
	rv := objc.Call[MutableTimedMetadataGroup](m_, objc.Sel("initWithItems:timeRange:"), items, timeRange)
	return rv
}

// Creates a timed metadata group initialized with the given metadata items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtimedmetadatagroup/1389632-initwithitems?language=objc
func NewMutableTimedMetadataGroupWithItemsTimeRange(items []IMetadataItem, timeRange coremedia.TimeRange) MutableTimedMetadataGroup {
	instance := MutableTimedMetadataGroupClass.Alloc().InitWithItemsTimeRange(items, timeRange)
	instance.Autorelease()
	return instance
}

func (m_ MutableTimedMetadataGroup) InitWithSampleBuffer(sampleBuffer coremedia.SampleBufferRef) MutableTimedMetadataGroup {
	rv := objc.Call[MutableTimedMetadataGroup](m_, objc.Sel("initWithSampleBuffer:"), sampleBuffer)
	return rv
}

// Creates a timed metadata group with a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtimedmetadatagroup/1387128-initwithsamplebuffer?language=objc
func NewMutableTimedMetadataGroupWithSampleBuffer(sampleBuffer coremedia.SampleBufferRef) MutableTimedMetadataGroup {
	instance := MutableTimedMetadataGroupClass.Alloc().InitWithSampleBuffer(sampleBuffer)
	instance.Autorelease()
	return instance
}

// An array of metadata items in the timed metadata group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutabletimedmetadatagroup/1386481-items?language=objc
func (m_ MutableTimedMetadataGroup) SetItems(value []IMetadataItem) {
	objc.Call[objc.Void](m_, objc.Sel("setItems:"), value)
}

// The time range of the timed metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutabletimedmetadatagroup/1387595-timerange?language=objc
func (m_ MutableTimedMetadataGroup) SetTimeRange(value coremedia.TimeRange) {
	objc.Call[objc.Void](m_, objc.Sel("setTimeRange:"), value)
}
