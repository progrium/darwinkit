// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var DateComponentsClass = _DateComponentsClass{objc.GetClass("NSDateComponents")}

type _DateComponentsClass struct {
	objc.Class
}

type IDateComponents interface {
	objc.IObject
	IsValidDateInCalendar(calendar ICalendar) bool
	ValueForComponent(unit CalendarUnit) int
	SetValueForComponent(value int, unit CalendarUnit)
	Calendar() Calendar
	SetCalendar(value ICalendar)
	TimeZone() TimeZone
	SetTimeZone(value ITimeZone)
	IsValidDate() bool
	Date() Date
	Era() int
	SetEra(value int)
	Year() int
	SetYear(value int)
	YearForWeekOfYear() int
	SetYearForWeekOfYear(value int)
	Quarter() int
	SetQuarter(value int)
	Month() int
	SetMonth(value int)
	IsLeapMonth() bool
	SetLeapMonth(value bool)
	Weekday() int
	SetWeekday(value int)
	WeekdayOrdinal() int
	SetWeekdayOrdinal(value int)
	WeekOfMonth() int
	SetWeekOfMonth(value int)
	WeekOfYear() int
	SetWeekOfYear(value int)
	Day() int
	SetDay(value int)
	Hour() int
	SetHour(value int)
	Minute() int
	SetMinute(value int)
	Second() int
	SetSecond(value int)
	Nanosecond() int
	SetNanosecond(value int)
}

type DateComponents struct {
	objc.Object
}

func MakeDateComponents(ptr unsafe.Pointer) DateComponents {
	return DateComponents{
		Object: objc.MakeObject(ptr),
	}
}

func (dc _DateComponentsClass) Alloc() DateComponents {
	rv := objc.CallMethod[DateComponents](dc, objc.GetSelector("alloc"))
	return rv
}

func DateComponents_Alloc() DateComponents {
	return DateComponentsClass.Alloc()
}

func (dc _DateComponentsClass) New() DateComponents {
	rv := objc.CallMethod[DateComponents](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDateComponents() DateComponents {
	return DateComponentsClass.New()
}

func DateComponents_New() DateComponents {
	return DateComponentsClass.New()
}

func (d_ DateComponents) Init() DateComponents {
	rv := objc.CallMethod[DateComponents](d_, objc.GetSelector("init"))
	return rv
}

func DateComponents_Init() DateComponents {
	return DateComponentsClass.Alloc().Init()
}

func (d_ DateComponents) IsValidDateInCalendar(calendar ICalendar) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isValidDateInCalendar:"), objc.ExtractPtr(calendar))
	return rv
}

func (d_ DateComponents) ValueForComponent(unit CalendarUnit) int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("valueForComponent:"), unit)
	return rv
}

func (d_ DateComponents) SetValueForComponent(value int, unit CalendarUnit) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setValue:forComponent:"), value, unit)
}

func (d_ DateComponents) Calendar() Calendar {
	rv := objc.CallMethod[Calendar](d_, objc.GetSelector("calendar"))
	return rv
}

func (d_ DateComponents) SetCalendar(value ICalendar) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setCalendar:"), objc.ExtractPtr(value))
}

func (d_ DateComponents) TimeZone() TimeZone {
	rv := objc.CallMethod[TimeZone](d_, objc.GetSelector("timeZone"))
	return rv
}

func (d_ DateComponents) SetTimeZone(value ITimeZone) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setTimeZone:"), objc.ExtractPtr(value))
}

func (d_ DateComponents) IsValidDate() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isValidDate"))
	return rv
}

func (d_ DateComponents) Date() Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("date"))
	return rv
}

func (d_ DateComponents) Era() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("era"))
	return rv
}

func (d_ DateComponents) SetEra(value int) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setEra:"), value)
}

func (d_ DateComponents) Year() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("year"))
	return rv
}

func (d_ DateComponents) SetYear(value int) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setYear:"), value)
}

func (d_ DateComponents) YearForWeekOfYear() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("yearForWeekOfYear"))
	return rv
}

func (d_ DateComponents) SetYearForWeekOfYear(value int) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setYearForWeekOfYear:"), value)
}

func (d_ DateComponents) Quarter() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("quarter"))
	return rv
}

func (d_ DateComponents) SetQuarter(value int) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setQuarter:"), value)
}

func (d_ DateComponents) Month() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("month"))
	return rv
}

func (d_ DateComponents) SetMonth(value int) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setMonth:"), value)
}

func (d_ DateComponents) IsLeapMonth() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isLeapMonth"))
	return rv
}

func (d_ DateComponents) SetLeapMonth(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setLeapMonth:"), value)
}

func (d_ DateComponents) Weekday() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("weekday"))
	return rv
}

func (d_ DateComponents) SetWeekday(value int) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setWeekday:"), value)
}

func (d_ DateComponents) WeekdayOrdinal() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("weekdayOrdinal"))
	return rv
}

func (d_ DateComponents) SetWeekdayOrdinal(value int) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setWeekdayOrdinal:"), value)
}

func (d_ DateComponents) WeekOfMonth() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("weekOfMonth"))
	return rv
}

func (d_ DateComponents) SetWeekOfMonth(value int) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setWeekOfMonth:"), value)
}

func (d_ DateComponents) WeekOfYear() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("weekOfYear"))
	return rv
}

func (d_ DateComponents) SetWeekOfYear(value int) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setWeekOfYear:"), value)
}

func (d_ DateComponents) Day() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("day"))
	return rv
}

func (d_ DateComponents) SetDay(value int) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDay:"), value)
}

func (d_ DateComponents) Hour() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("hour"))
	return rv
}

func (d_ DateComponents) SetHour(value int) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setHour:"), value)
}

func (d_ DateComponents) Minute() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("minute"))
	return rv
}

func (d_ DateComponents) SetMinute(value int) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setMinute:"), value)
}

func (d_ DateComponents) Second() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("second"))
	return rv
}

func (d_ DateComponents) SetSecond(value int) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setSecond:"), value)
}

func (d_ DateComponents) Nanosecond() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("nanosecond"))
	return rv
}

func (d_ DateComponents) SetNanosecond(value int) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setNanosecond:"), value)
}
