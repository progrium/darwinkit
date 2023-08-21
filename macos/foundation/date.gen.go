// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Date] class.
var DateClass = _DateClass{objc.GetClass("NSDate")}

type _DateClass struct {
	objc.Class
}

// An interface definition for the [Date] class.
type IDate interface {
	objc.IObject
	TimeIntervalSinceDate(anotherDate IDate) TimeInterval
	DescriptionWithLocale(locale objc.IObject) string
	IsEqualToDate(otherDate IDate) bool
	LaterDate(anotherDate IDate) Date
	Compare(other IDate) ComparisonResult
	EarlierDate(anotherDate IDate) Date
	TimeIntervalSince1970() TimeInterval
	Description() string
	TimeIntervalSinceNow() TimeInterval
	TimeIntervalSinceReferenceDate() TimeInterval
}

// A representation of a specific point in time, independent of any calendar or time zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate?language=objc
type Date struct {
	objc.Object
}

func DateFrom(ptr unsafe.Pointer) Date {
	return Date{
		Object: objc.ObjectFrom(ptr),
	}
}

func (dc _DateClass) Date() Date {
	rv := objc.Call[Date](dc, objc.Sel("date"))
	return rv
}

// Creates and returns a new date object set to the current date and time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1591574-date?language=objc
func Date_Date() Date {
	return DateClass.Date()
}

func (d_ Date) DateByAddingTimeInterval(ti TimeInterval) Date {
	rv := objc.Call[Date](d_, objc.Sel("dateByAddingTimeInterval:"), ti)
	return rv
}

// Returns a new date object that is set to a given number of seconds relative to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1408823-datebyaddingtimeinterval?language=objc
func Date_DateByAddingTimeInterval(ti TimeInterval) Date {
	instance := DateClass.Alloc().DateByAddingTimeInterval(ti)
	instance.Autorelease()
	return instance
}

func (dc _DateClass) DateWithTimeIntervalSinceDate(secsToBeAdded TimeInterval, date IDate) Date {
	rv := objc.Call[Date](dc, objc.Sel("dateWithTimeInterval:sinceDate:"), secsToBeAdded, objc.Ptr(date))
	return rv
}

// Creates and returns a date object set to a given number of seconds from the specified date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1591578-datewithtimeinterval?language=objc
func Date_DateWithTimeIntervalSinceDate(secsToBeAdded TimeInterval, date IDate) Date {
	return DateClass.DateWithTimeIntervalSinceDate(secsToBeAdded, date)
}

func (dc _DateClass) DateWithTimeIntervalSinceNow(secs TimeInterval) Date {
	rv := objc.Call[Date](dc, objc.Sel("dateWithTimeIntervalSinceNow:"), secs)
	return rv
}

// Creates and returns a date object set to a given number of seconds from the current date and time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1591579-datewithtimeintervalsincenow?language=objc
func Date_DateWithTimeIntervalSinceNow(secs TimeInterval) Date {
	return DateClass.DateWithTimeIntervalSinceNow(secs)
}

func (dc _DateClass) DateWithTimeIntervalSince1970(secs TimeInterval) Date {
	rv := objc.Call[Date](dc, objc.Sel("dateWithTimeIntervalSince1970:"), secs)
	return rv
}

// Creates and returns a date object set to the given number of seconds from 00:00:00 UTC on 1 January 1970. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1591576-datewithtimeintervalsince1970?language=objc
func Date_DateWithTimeIntervalSince1970(secs TimeInterval) Date {
	return DateClass.DateWithTimeIntervalSince1970(secs)
}

func (d_ Date) InitWithTimeIntervalSinceReferenceDate(ti TimeInterval) Date {
	rv := objc.Call[Date](d_, objc.Sel("initWithTimeIntervalSinceReferenceDate:"), ti)
	return rv
}

// Returns a date object initialized relative to 00:00:00 UTC on 1 January 2001 by a given number of seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1409769-initwithtimeintervalsincereferen?language=objc
func NewDateWithTimeIntervalSinceReferenceDate(ti TimeInterval) Date {
	instance := DateClass.Alloc().InitWithTimeIntervalSinceReferenceDate(ti)
	instance.Autorelease()
	return instance
}

func (d_ Date) InitWithTimeIntervalSinceNow(secs TimeInterval) Date {
	rv := objc.Call[Date](d_, objc.Sel("initWithTimeIntervalSinceNow:"), secs)
	return rv
}

// Returns a date object initialized relative to the current date and time by a given number of seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1411701-initwithtimeintervalsincenow?language=objc
func NewDateWithTimeIntervalSinceNow(secs TimeInterval) Date {
	instance := DateClass.Alloc().InitWithTimeIntervalSinceNow(secs)
	instance.Autorelease()
	return instance
}

func (dc _DateClass) DateWithTimeIntervalSinceReferenceDate(ti TimeInterval) Date {
	rv := objc.Call[Date](dc, objc.Sel("dateWithTimeIntervalSinceReferenceDate:"), ti)
	return rv
}

// Creates and returns a date object set to a given number of seconds from 00:00:00 UTC on 1 January 2001. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1591577-datewithtimeintervalsincereferen?language=objc
func Date_DateWithTimeIntervalSinceReferenceDate(ti TimeInterval) Date {
	return DateClass.DateWithTimeIntervalSinceReferenceDate(ti)
}

func (d_ Date) InitWithTimeIntervalSinceDate(secsToBeAdded TimeInterval, date IDate) Date {
	rv := objc.Call[Date](d_, objc.Sel("initWithTimeInterval:sinceDate:"), secsToBeAdded, objc.Ptr(date))
	return rv
}

// Returns a date object initialized relative to another given date by a given number of seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1414201-initwithtimeinterval?language=objc
func NewDateWithTimeIntervalSinceDate(secsToBeAdded TimeInterval, date IDate) Date {
	instance := DateClass.Alloc().InitWithTimeIntervalSinceDate(secsToBeAdded, date)
	instance.Autorelease()
	return instance
}

func (d_ Date) Init() Date {
	rv := objc.Call[Date](d_, objc.Sel("init"))
	return rv
}

func (d_ Date) InitWithTimeIntervalSince1970(secs TimeInterval) Date {
	rv := objc.Call[Date](d_, objc.Sel("initWithTimeIntervalSince1970:"), secs)
	return rv
}

// Returns a date object initialized relative to 00:00:00 UTC on 1 January 1970 by a given number of seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1416453-initwithtimeintervalsince1970?language=objc
func NewDateWithTimeIntervalSince1970(secs TimeInterval) Date {
	instance := DateClass.Alloc().InitWithTimeIntervalSince1970(secs)
	instance.Autorelease()
	return instance
}

func (dc _DateClass) Alloc() Date {
	rv := objc.Call[Date](dc, objc.Sel("alloc"))
	return rv
}

func Date_Alloc() Date {
	return DateClass.Alloc()
}

func (dc _DateClass) New() Date {
	rv := objc.Call[Date](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDate() Date {
	return DateClass.New()
}

// Returns the interval between the receiver and another given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1410206-timeintervalsincedate?language=objc
func (d_ Date) TimeIntervalSinceDate(anotherDate IDate) TimeInterval {
	rv := objc.Call[TimeInterval](d_, objc.Sel("timeIntervalSinceDate:"), objc.Ptr(anotherDate))
	return rv
}

// Returns a string representation of the date using the given locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1414108-descriptionwithlocale?language=objc
func (d_ Date) DescriptionWithLocale(locale objc.IObject) string {
	rv := objc.Call[string](d_, objc.Sel("descriptionWithLocale:"), locale)
	return rv
}

// Returns a Boolean value that indicates whether a given object is a date that is exactly equal the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1416400-isequaltodate?language=objc
func (d_ Date) IsEqualToDate(otherDate IDate) bool {
	rv := objc.Call[bool](d_, objc.Sel("isEqualToDate:"), objc.Ptr(otherDate))
	return rv
}

// Returns the later of the receiver and another given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1409058-laterdate?language=objc
func (d_ Date) LaterDate(anotherDate IDate) Date {
	rv := objc.Call[Date](d_, objc.Sel("laterDate:"), objc.Ptr(anotherDate))
	return rv
}

// Indicates the temporal ordering of the receiver and another given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1409748-compare?language=objc
func (d_ Date) Compare(other IDate) ComparisonResult {
	rv := objc.Call[ComparisonResult](d_, objc.Sel("compare:"), objc.Ptr(other))
	return rv
}

// Returns the earlier of the receiver and another given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1412791-earlierdate?language=objc
func (d_ Date) EarlierDate(anotherDate IDate) Date {
	rv := objc.Call[Date](d_, objc.Sel("earlierDate:"), objc.Ptr(anotherDate))
	return rv
}

// A date object representing a date in the distant past. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1418197-distantpast?language=objc
func (dc _DateClass) DistantPast() Date {
	rv := objc.Call[Date](dc, objc.Sel("distantPast"))
	return rv
}

// A date object representing a date in the distant past. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1418197-distantpast?language=objc
func Date_DistantPast() Date {
	return DateClass.DistantPast()
}

// The interval between the date object and 00:00:00 UTC on 1 January 1970. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1407504-timeintervalsince1970?language=objc
func (d_ Date) TimeIntervalSince1970() TimeInterval {
	rv := objc.Call[TimeInterval](d_, objc.Sel("timeIntervalSince1970"))
	return rv
}

// A date object representing a date in the distant future. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1415385-distantfuture?language=objc
func (dc _DateClass) DistantFuture() Date {
	rv := objc.Call[Date](dc, objc.Sel("distantFuture"))
	return rv
}

// A date object representing a date in the distant future. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1415385-distantfuture?language=objc
func Date_DistantFuture() Date {
	return DateClass.DistantFuture()
}

// A string representation of the date object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1409767-description?language=objc
func (d_ Date) Description() string {
	rv := objc.Call[string](d_, objc.Sel("description"))
	return rv
}

// The interval between the date object and the current date and time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1407937-timeintervalsincenow?language=objc
func (d_ Date) TimeIntervalSinceNow() TimeInterval {
	rv := objc.Call[TimeInterval](d_, objc.Sel("timeIntervalSinceNow"))
	return rv
}

// The interval between the date object and 00:00:00 UTC on 1 January 2001. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/1417376-timeintervalsincereferencedate?language=objc
func (d_ Date) TimeIntervalSinceReferenceDate() TimeInterval {
	rv := objc.Call[TimeInterval](d_, objc.Sel("timeIntervalSinceReferenceDate"))
	return rv
}

// The current date and time, as of the time of access. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/3180113-now?language=objc
func (dc _DateClass) Now() Date {
	rv := objc.Call[Date](dc, objc.Sel("now"))
	return rv
}

// The current date and time, as of the time of access. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdate/3180113-now?language=objc
func Date_Now() Date {
	return DateClass.Now()
}
