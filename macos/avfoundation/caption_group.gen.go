// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptionGroup] class.
var CaptionGroupClass = _CaptionGroupClass{objc.GetClass("AVCaptionGroup")}

type _CaptionGroupClass struct {
	objc.Class
}

// An interface definition for the [CaptionGroup] class.
type ICaptionGroup interface {
	objc.IObject
	Captions() []Caption
	TimeRange() coremedia.TimeRange
}

// An object that represents zero or more captions that intersect in time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptiongroup?language=objc
type CaptionGroup struct {
	objc.Object
}

func CaptionGroupFrom(ptr unsafe.Pointer) CaptionGroup {
	return CaptionGroup{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ CaptionGroup) InitWithTimeRange(timeRange coremedia.TimeRange) CaptionGroup {
	rv := objc.Call[CaptionGroup](c_, objc.Sel("initWithTimeRange:"), timeRange)
	return rv
}

// Creates a caption group with a time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptiongroup/3752962-initwithtimerange?language=objc
func NewCaptionGroupWithTimeRange(timeRange coremedia.TimeRange) CaptionGroup {
	instance := CaptionGroupClass.Alloc().InitWithTimeRange(timeRange)
	instance.Autorelease()
	return instance
}

func (c_ CaptionGroup) InitWithCaptionsTimeRange(captions []ICaption, timeRange coremedia.TimeRange) CaptionGroup {
	rv := objc.Call[CaptionGroup](c_, objc.Sel("initWithCaptions:timeRange:"), captions, timeRange)
	return rv
}

// Creates a caption group with captions and a time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptiongroup/3752961-initwithcaptions?language=objc
func NewCaptionGroupWithCaptionsTimeRange(captions []ICaption, timeRange coremedia.TimeRange) CaptionGroup {
	instance := CaptionGroupClass.Alloc().InitWithCaptionsTimeRange(captions, timeRange)
	instance.Autorelease()
	return instance
}

func (cc _CaptionGroupClass) Alloc() CaptionGroup {
	rv := objc.Call[CaptionGroup](cc, objc.Sel("alloc"))
	return rv
}

func CaptionGroup_Alloc() CaptionGroup {
	return CaptionGroupClass.Alloc()
}

func (cc _CaptionGroupClass) New() CaptionGroup {
	rv := objc.Call[CaptionGroup](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptionGroup() CaptionGroup {
	return CaptionGroupClass.New()
}

func (c_ CaptionGroup) Init() CaptionGroup {
	rv := objc.Call[CaptionGroup](c_, objc.Sel("init"))
	return rv
}

// The captions associated with the caption group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptiongroup/3752960-captions?language=objc
func (c_ CaptionGroup) Captions() []Caption {
	rv := objc.Call[[]Caption](c_, objc.Sel("captions"))
	return rv
}

// The time range of the caption group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptiongroup/3752963-timerange?language=objc
func (c_ CaptionGroup) TimeRange() coremedia.TimeRange {
	rv := objc.Call[coremedia.TimeRange](c_, objc.Sel("timeRange"))
	return rv
}
