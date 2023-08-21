// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DateInterval] class.
var DateIntervalClass = _DateIntervalClass{objc.GetClass("NSDateInterval")}

type _DateIntervalClass struct {
	objc.Class
}

// An interface definition for the [DateInterval] class.
type IDateInterval interface {
	objc.IObject
	IsEqualToDateInterval(dateInterval IDateInterval) bool
	IntersectionWithDateInterval(dateInterval IDateInterval) DateInterval
	ContainsDate(date IDate) bool
	IntersectsDateInterval(dateInterval IDateInterval) bool
	Compare(dateInterval IDateInterval) ComparisonResult
	StartDate() Date
	Duration() TimeInterval
	EndDate() Date
}

// An object representing the span of time between a specific start date and end date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateinterval?language=objc
type DateInterval struct {
	objc.Object
}

func DateIntervalFrom(ptr unsafe.Pointer) DateInterval {
	return DateInterval{
		Object: objc.ObjectFrom(ptr),
	}
}

func (d_ DateInterval) InitWithStartDateEndDate(startDate IDate, endDate IDate) DateInterval {
	rv := objc.Call[DateInterval](d_, objc.Sel("initWithStartDate:endDate:"), objc.Ptr(startDate), objc.Ptr(endDate))
	return rv
}

// Initializes a date interval from a given start date and end date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateinterval/1641639-initwithstartdate?language=objc
func NewDateIntervalWithStartDateEndDate(startDate IDate, endDate IDate) DateInterval {
	instance := DateIntervalClass.Alloc().InitWithStartDateEndDate(startDate, endDate)
	instance.Autorelease()
	return instance
}

func (d_ DateInterval) Init() DateInterval {
	rv := objc.Call[DateInterval](d_, objc.Sel("init"))
	return rv
}

func (dc _DateIntervalClass) Alloc() DateInterval {
	rv := objc.Call[DateInterval](dc, objc.Sel("alloc"))
	return rv
}

func DateInterval_Alloc() DateInterval {
	return DateIntervalClass.Alloc()
}

func (dc _DateIntervalClass) New() DateInterval {
	rv := objc.Call[DateInterval](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDateInterval() DateInterval {
	return DateIntervalClass.New()
}

// Indicates whether the receiver is equal to the specified date interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateinterval/1641650-isequaltodateinterval?language=objc
func (d_ DateInterval) IsEqualToDateInterval(dateInterval IDateInterval) bool {
	rv := objc.Call[bool](d_, objc.Sel("isEqualToDateInterval:"), objc.Ptr(dateInterval))
	return rv
}

// Returns the intersection between the receiver and the specified date interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateinterval/1641645-intersectionwithdateinterval?language=objc
func (d_ DateInterval) IntersectionWithDateInterval(dateInterval IDateInterval) DateInterval {
	rv := objc.Call[DateInterval](d_, objc.Sel("intersectionWithDateInterval:"), objc.Ptr(dateInterval))
	return rv
}

// Indicates whether the receiver contains the specified date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateinterval/1641647-containsdate?language=objc
func (d_ DateInterval) ContainsDate(date IDate) bool {
	rv := objc.Call[bool](d_, objc.Sel("containsDate:"), objc.Ptr(date))
	return rv
}

// Indicates whether the receiver intersects with the specified date interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateinterval/1641654-intersectsdateinterval?language=objc
func (d_ DateInterval) IntersectsDateInterval(dateInterval IDateInterval) bool {
	rv := objc.Call[bool](d_, objc.Sel("intersectsDateInterval:"), objc.Ptr(dateInterval))
	return rv
}

// Compares the receiver with the specified date interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateinterval/1641636-compare?language=objc
func (d_ DateInterval) Compare(dateInterval IDateInterval) ComparisonResult {
	rv := objc.Call[ComparisonResult](d_, objc.Sel("compare:"), objc.Ptr(dateInterval))
	return rv
}

// The start date of the date interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateinterval/1641656-startdate?language=objc
func (d_ DateInterval) StartDate() Date {
	rv := objc.Call[Date](d_, objc.Sel("startDate"))
	return rv
}

// The duration of the date interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateinterval/1641643-duration?language=objc
func (d_ DateInterval) Duration() TimeInterval {
	rv := objc.Call[TimeInterval](d_, objc.Sel("duration"))
	return rv
}

// The end date of the date interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateinterval/1641651-enddate?language=objc
func (d_ DateInterval) EndDate() Date {
	rv := objc.Call[Date](d_, objc.Sel("endDate"))
	return rv
}
