// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var DateIntervalClass = _DateIntervalClass{objc.GetClass("NSDateInterval")}

type _DateIntervalClass struct {
	objc.Class
}

type IDateInterval interface {
	objc.IObject
	Compare(dateInterval IDateInterval) ComparisonResult
	IsEqualToDateInterval(dateInterval IDateInterval) bool
	IntersectsDateInterval(dateInterval IDateInterval) bool
	IntersectionWithDateInterval(dateInterval IDateInterval) DateInterval
	ContainsDate(date IDate) bool
	StartDate() Date
	EndDate() Date
	Duration() TimeInterval
}

type DateInterval struct {
	objc.Object
}

func MakeDateInterval(ptr unsafe.Pointer) DateInterval {
	return DateInterval{
		Object: objc.MakeObject(ptr),
	}
}

func (d_ DateInterval) Init() DateInterval {
	rv := objc.CallMethod[DateInterval](d_, objc.GetSelector("init"))
	return rv
}

func DateInterval_Init() DateInterval {
	return DateIntervalClass.Alloc().Init()
}

func (d_ DateInterval) InitWithStartDateDuration(startDate IDate, duration TimeInterval) DateInterval {
	rv := objc.CallMethod[DateInterval](d_, objc.GetSelector("initWithStartDate:duration:"), objc.ExtractPtr(startDate), duration)
	return rv
}

func DateInterval_InitWithStartDateDuration(startDate IDate, duration TimeInterval) DateInterval {
	return DateIntervalClass.Alloc().InitWithStartDateDuration(startDate, duration)
}

func (d_ DateInterval) InitWithStartDateEndDate(startDate IDate, endDate IDate) DateInterval {
	rv := objc.CallMethod[DateInterval](d_, objc.GetSelector("initWithStartDate:endDate:"), objc.ExtractPtr(startDate), objc.ExtractPtr(endDate))
	return rv
}

func DateInterval_InitWithStartDateEndDate(startDate IDate, endDate IDate) DateInterval {
	return DateIntervalClass.Alloc().InitWithStartDateEndDate(startDate, endDate)
}

func (dc _DateIntervalClass) Alloc() DateInterval {
	rv := objc.CallMethod[DateInterval](dc, objc.GetSelector("alloc"))
	return rv
}

func DateInterval_Alloc() DateInterval {
	return DateIntervalClass.Alloc()
}

func (dc _DateIntervalClass) New() DateInterval {
	rv := objc.CallMethod[DateInterval](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDateInterval() DateInterval {
	return DateIntervalClass.New()
}

func DateInterval_New() DateInterval {
	return DateIntervalClass.New()
}

func (d_ DateInterval) Compare(dateInterval IDateInterval) ComparisonResult {
	rv := objc.CallMethod[ComparisonResult](d_, objc.GetSelector("compare:"), objc.ExtractPtr(dateInterval))
	return rv
}

func (d_ DateInterval) IsEqualToDateInterval(dateInterval IDateInterval) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isEqualToDateInterval:"), objc.ExtractPtr(dateInterval))
	return rv
}

func (d_ DateInterval) IntersectsDateInterval(dateInterval IDateInterval) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("intersectsDateInterval:"), objc.ExtractPtr(dateInterval))
	return rv
}

func (d_ DateInterval) IntersectionWithDateInterval(dateInterval IDateInterval) DateInterval {
	rv := objc.CallMethod[DateInterval](d_, objc.GetSelector("intersectionWithDateInterval:"), objc.ExtractPtr(dateInterval))
	return rv
}

func (d_ DateInterval) ContainsDate(date IDate) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("containsDate:"), objc.ExtractPtr(date))
	return rv
}

func (d_ DateInterval) StartDate() Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("startDate"))
	return rv
}

func (d_ DateInterval) EndDate() Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("endDate"))
	return rv
}

func (d_ DateInterval) Duration() TimeInterval {
	rv := objc.CallMethod[TimeInterval](d_, objc.GetSelector("duration"))
	return rv
}
