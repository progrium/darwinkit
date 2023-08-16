// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DateComponents] class.
var DateComponentsClass = _DateComponentsClass{objc.GetClass("NSDateComponents")}

type _DateComponentsClass struct {
	objc.Class
}

// An interface definition for the [DateComponents] class.
type IDateComponents interface {
	objc.IObject
	IsValidDateInCalendar(calendar ICalendar) bool
	ValueForComponent(unit CalendarUnit) int
	SetValueForComponent(value int, unit CalendarUnit)
	IsValidDate() bool
	WeekOfYear() int
	SetWeekOfYear(value int)
	Month() int
	SetMonth(value int)
	Second() int
	SetSecond(value int)
	Date() Date
	Weekday() int
	SetWeekday(value int)
	Hour() int
	SetHour(value int)
	IsLeapMonth() bool
	SetLeapMonth(value bool)
	Quarter() int
	SetQuarter(value int)
	Year() int
	SetYear(value int)
	WeekOfMonth() int
	SetWeekOfMonth(value int)
	TimeZone() TimeZone
	SetTimeZone(value ITimeZone)
	Nanosecond() int
	SetNanosecond(value int)
	Day() int
	SetDay(value int)
	WeekdayOrdinal() int
	SetWeekdayOrdinal(value int)
	Calendar() Calendar
	SetCalendar(value ICalendar)
	Era() int
	SetEra(value int)
	Minute() int
	SetMinute(value int)
	YearForWeekOfYear() int
	SetYearForWeekOfYear(value int)
}

// An object that specifies a date or time in terms of units (such as year, month, day, hour, and minute) to be evaluated in a calendar system and time zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents?language=objc
type DateComponents struct {
	objc.Object
}

func DateComponentsFrom(ptr unsafe.Pointer) DateComponents {
	return DateComponents{
		Object: objc.ObjectFrom(ptr),
	}
}

func (dc _DateComponentsClass) Alloc() DateComponents {
	rv := objc.Call[DateComponents](dc, objc.Sel("alloc"))
	return rv
}

func DateComponents_Alloc() DateComponents {
	return DateComponentsClass.Alloc()
}

func (dc _DateComponentsClass) New() DateComponents {
	rv := objc.Call[DateComponents](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDateComponents() DateComponents {
	return DateComponentsClass.New()
}

func (d_ DateComponents) Init() DateComponents {
	rv := objc.Call[DateComponents](d_, objc.Sel("init"))
	return rv
}

// Returns a Boolean value that indicates whether the current combination of properties represents a date which exists in the specified calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1412707-isvaliddateincalendar?language=objc
func (d_ DateComponents) IsValidDateInCalendar(calendar ICalendar) bool {
	rv := objc.Call[bool](d_, objc.Sel("isValidDateInCalendar:"), objc.Ptr(calendar))
	return rv
}

// Returns the value for a given calendar unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1416763-valueforcomponent?language=objc
func (d_ DateComponents) ValueForComponent(unit CalendarUnit) int {
	rv := objc.Call[int](d_, objc.Sel("valueForComponent:"), unit)
	return rv
}

// Sets a value for a given calendar unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1415961-setvalue?language=objc
func (d_ DateComponents) SetValueForComponent(value int, unit CalendarUnit) {
	objc.Call[objc.Void](d_, objc.Sel("setValue:forComponent:"), value, unit)
}

// A Boolean value that indicates whether the current combination of properties represents a date which exists in the current calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1408788-validdate?language=objc
func (d_ DateComponents) IsValidDate() bool {
	rv := objc.Call[bool](d_, objc.Sel("isValidDate"))
	return rv
}

// The ISO 8601 week date of the year. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1416908-weekofyear?language=objc
func (d_ DateComponents) WeekOfYear() int {
	rv := objc.Call[int](d_, objc.Sel("weekOfYear"))
	return rv
}

// The ISO 8601 week date of the year. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1416908-weekofyear?language=objc
func (d_ DateComponents) SetWeekOfYear(value int) {
	objc.Call[objc.Void](d_, objc.Sel("setWeekOfYear:"), value)
}

// The number of months. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1408724-month?language=objc
func (d_ DateComponents) Month() int {
	rv := objc.Call[int](d_, objc.Sel("month"))
	return rv
}

// The number of months. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1408724-month?language=objc
func (d_ DateComponents) SetMonth(value int) {
	objc.Call[objc.Void](d_, objc.Sel("setMonth:"), value)
}

// The number of second units for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1414045-second?language=objc
func (d_ DateComponents) Second() int {
	rv := objc.Call[int](d_, objc.Sel("second"))
	return rv
}

// The number of second units for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1414045-second?language=objc
func (d_ DateComponents) SetSecond(value int) {
	objc.Call[objc.Void](d_, objc.Sel("setSecond:"), value)
}

// The date calculated from the current components using the stored calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1412861-date?language=objc
func (d_ DateComponents) Date() Date {
	rv := objc.Call[Date](d_, objc.Sel("date"))
	return rv
}

// The number of the weekdays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1410442-weekday?language=objc
func (d_ DateComponents) Weekday() int {
	rv := objc.Call[int](d_, objc.Sel("weekday"))
	return rv
}

// The number of the weekdays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1410442-weekday?language=objc
func (d_ DateComponents) SetWeekday(value int) {
	objc.Call[objc.Void](d_, objc.Sel("setWeekday:"), value)
}

// The number of hour units for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1411355-hour?language=objc
func (d_ DateComponents) Hour() int {
	rv := objc.Call[int](d_, objc.Sel("hour"))
	return rv
}

// The number of hour units for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1411355-hour?language=objc
func (d_ DateComponents) SetHour(value int) {
	objc.Call[objc.Void](d_, objc.Sel("setHour:"), value)
}

// A Boolean value that indicates whether the month is a leap month. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1408836-leapmonth?language=objc
func (d_ DateComponents) IsLeapMonth() bool {
	rv := objc.Call[bool](d_, objc.Sel("isLeapMonth"))
	return rv
}

// A Boolean value that indicates whether the month is a leap month. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1408836-leapmonth?language=objc
func (d_ DateComponents) SetLeapMonth(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setLeapMonth:"), value)
}

// The number of quarters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1416503-quarter?language=objc
func (d_ DateComponents) Quarter() int {
	rv := objc.Call[int](d_, objc.Sel("quarter"))
	return rv
}

// The number of quarters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1416503-quarter?language=objc
func (d_ DateComponents) SetQuarter(value int) {
	objc.Call[objc.Void](d_, objc.Sel("setQuarter:"), value)
}

// The number of years. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1412462-year?language=objc
func (d_ DateComponents) Year() int {
	rv := objc.Call[int](d_, objc.Sel("year"))
	return rv
}

// The number of years. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1412462-year?language=objc
func (d_ DateComponents) SetYear(value int) {
	objc.Call[objc.Void](d_, objc.Sel("setYear:"), value)
}

// The week number of the months. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1413168-weekofmonth?language=objc
func (d_ DateComponents) WeekOfMonth() int {
	rv := objc.Call[int](d_, objc.Sel("weekOfMonth"))
	return rv
}

// The week number of the months. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1413168-weekofmonth?language=objc
func (d_ DateComponents) SetWeekOfMonth(value int) {
	objc.Call[objc.Void](d_, objc.Sel("setWeekOfMonth:"), value)
}

// The time zone used to interpret the date components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1408233-timezone?language=objc
func (d_ DateComponents) TimeZone() TimeZone {
	rv := objc.Call[TimeZone](d_, objc.Sel("timeZone"))
	return rv
}

// The time zone used to interpret the date components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1408233-timezone?language=objc
func (d_ DateComponents) SetTimeZone(value ITimeZone) {
	objc.Call[objc.Void](d_, objc.Sel("setTimeZone:"), objc.Ptr(value))
}

// The number of nanosecond units for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1415730-nanosecond?language=objc
func (d_ DateComponents) Nanosecond() int {
	rv := objc.Call[int](d_, objc.Sel("nanosecond"))
	return rv
}

// The number of nanosecond units for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1415730-nanosecond?language=objc
func (d_ DateComponents) SetNanosecond(value int) {
	objc.Call[objc.Void](d_, objc.Sel("setNanosecond:"), value)
}

// The number of days. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1415267-day?language=objc
func (d_ DateComponents) Day() int {
	rv := objc.Call[int](d_, objc.Sel("day"))
	return rv
}

// The number of days. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1415267-day?language=objc
func (d_ DateComponents) SetDay(value int) {
	objc.Call[objc.Void](d_, objc.Sel("setDay:"), value)
}

// The ordinal number of weekdays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1409476-weekdayordinal?language=objc
func (d_ DateComponents) WeekdayOrdinal() int {
	rv := objc.Call[int](d_, objc.Sel("weekdayOrdinal"))
	return rv
}

// The ordinal number of weekdays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1409476-weekdayordinal?language=objc
func (d_ DateComponents) SetWeekdayOrdinal(value int) {
	objc.Call[objc.Void](d_, objc.Sel("setWeekdayOrdinal:"), value)
}

// The calendar used to interpret the date components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1415799-calendar?language=objc
func (d_ DateComponents) Calendar() Calendar {
	rv := objc.Call[Calendar](d_, objc.Sel("calendar"))
	return rv
}

// The calendar used to interpret the date components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1415799-calendar?language=objc
func (d_ DateComponents) SetCalendar(value ICalendar) {
	objc.Call[objc.Void](d_, objc.Sel("setCalendar:"), objc.Ptr(value))
}

// The number of eras. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1416364-era?language=objc
func (d_ DateComponents) Era() int {
	rv := objc.Call[int](d_, objc.Sel("era"))
	return rv
}

// The number of eras. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1416364-era?language=objc
func (d_ DateComponents) SetEra(value int) {
	objc.Call[objc.Void](d_, objc.Sel("setEra:"), value)
}

// The number of minute units for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1409443-minute?language=objc
func (d_ DateComponents) Minute() int {
	rv := objc.Call[int](d_, objc.Sel("minute"))
	return rv
}

// The number of minute units for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1409443-minute?language=objc
func (d_ DateComponents) SetMinute(value int) {
	objc.Call[objc.Void](d_, objc.Sel("setMinute:"), value)
}

// The ISO 8601 week-numbering year. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1413809-yearforweekofyear?language=objc
func (d_ DateComponents) YearForWeekOfYear() int {
	rv := objc.Call[int](d_, objc.Sel("yearForWeekOfYear"))
	return rv
}

// The ISO 8601 week-numbering year. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/1413809-yearforweekofyear?language=objc
func (d_ DateComponents) SetYearForWeekOfYear(value int) {
	objc.Call[objc.Void](d_, objc.Sel("setYearForWeekOfYear:"), value)
}
