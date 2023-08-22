// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptionConversionTimeRangeAdjustment] class.
var CaptionConversionTimeRangeAdjustmentClass = _CaptionConversionTimeRangeAdjustmentClass{objc.GetClass("AVCaptionConversionTimeRangeAdjustment")}

type _CaptionConversionTimeRangeAdjustmentClass struct {
	objc.Class
}

// An interface definition for the [CaptionConversionTimeRangeAdjustment] class.
type ICaptionConversionTimeRangeAdjustment interface {
	ICaptionConversionAdjustment
	StartTimeOffset() coremedia.Time
	DurationOffset() coremedia.Time
}

// An object that describes an adjustment to the time range of one or more captions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversiontimerangeadjustment?language=objc
type CaptionConversionTimeRangeAdjustment struct {
	CaptionConversionAdjustment
}

func CaptionConversionTimeRangeAdjustmentFrom(ptr unsafe.Pointer) CaptionConversionTimeRangeAdjustment {
	return CaptionConversionTimeRangeAdjustment{
		CaptionConversionAdjustment: CaptionConversionAdjustmentFrom(ptr),
	}
}

func (cc _CaptionConversionTimeRangeAdjustmentClass) Alloc() CaptionConversionTimeRangeAdjustment {
	rv := objc.Call[CaptionConversionTimeRangeAdjustment](cc, objc.Sel("alloc"))
	return rv
}

func CaptionConversionTimeRangeAdjustment_Alloc() CaptionConversionTimeRangeAdjustment {
	return CaptionConversionTimeRangeAdjustmentClass.Alloc()
}

func (cc _CaptionConversionTimeRangeAdjustmentClass) New() CaptionConversionTimeRangeAdjustment {
	rv := objc.Call[CaptionConversionTimeRangeAdjustment](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptionConversionTimeRangeAdjustment() CaptionConversionTimeRangeAdjustment {
	return CaptionConversionTimeRangeAdjustmentClass.New()
}

func (c_ CaptionConversionTimeRangeAdjustment) Init() CaptionConversionTimeRangeAdjustment {
	rv := objc.Call[CaptionConversionTimeRangeAdjustment](c_, objc.Sel("init"))
	return rv
}

// The time value by which the system offsets the start times of captions to correct a problem. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversiontimerangeadjustment/3857641-starttimeoffset?language=objc
func (c_ CaptionConversionTimeRangeAdjustment) StartTimeOffset() coremedia.Time {
	rv := objc.Call[coremedia.Time](c_, objc.Sel("startTimeOffset"))
	return rv
}

// The time value by which the system offsets the durations of captions to correct a problem. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversiontimerangeadjustment/3857640-durationoffset?language=objc
func (c_ CaptionConversionTimeRangeAdjustment) DurationOffset() coremedia.Time {
	rv := objc.Call[coremedia.Time](c_, objc.Sel("durationOffset"))
	return rv
}
