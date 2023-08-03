// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var DateIntervalFormatterClass = _DateIntervalFormatterClass{objc.GetClass("NSDateIntervalFormatter")}

type _DateIntervalFormatterClass struct {
	objc.Class
}

type IDateIntervalFormatter interface {
	IFormatter
	StringFromDateToDate(fromDate IDate, toDate IDate) string
	StringFromDateInterval(dateInterval IDateInterval) string
	DateStyle() DateIntervalFormatterStyle
	SetDateStyle(value DateIntervalFormatterStyle)
	TimeStyle() DateIntervalFormatterStyle
	SetTimeStyle(value DateIntervalFormatterStyle)
	DateTemplate() string
	SetDateTemplate(value string)
	Calendar() Calendar
	SetCalendar(value ICalendar)
	Locale() Locale
	SetLocale(value ILocale)
	TimeZone() TimeZone
	SetTimeZone(value ITimeZone)
}

type DateIntervalFormatter struct {
	Formatter
}

func MakeDateIntervalFormatter(ptr unsafe.Pointer) DateIntervalFormatter {
	return DateIntervalFormatter{
		Formatter: MakeFormatter(ptr),
	}
}

func (dc _DateIntervalFormatterClass) Alloc() DateIntervalFormatter {
	rv := objc.CallMethod[DateIntervalFormatter](dc, objc.GetSelector("alloc"))
	return rv
}

func DateIntervalFormatter_Alloc() DateIntervalFormatter {
	return DateIntervalFormatterClass.Alloc()
}

func (dc _DateIntervalFormatterClass) New() DateIntervalFormatter {
	rv := objc.CallMethod[DateIntervalFormatter](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDateIntervalFormatter() DateIntervalFormatter {
	return DateIntervalFormatterClass.New()
}

func DateIntervalFormatter_New() DateIntervalFormatter {
	return DateIntervalFormatterClass.New()
}

func (d_ DateIntervalFormatter) Init() DateIntervalFormatter {
	rv := objc.CallMethod[DateIntervalFormatter](d_, objc.GetSelector("init"))
	return rv
}

func DateIntervalFormatter_Init() DateIntervalFormatter {
	return DateIntervalFormatterClass.Alloc().Init()
}

func (d_ DateIntervalFormatter) StringFromDateToDate(fromDate IDate, toDate IDate) string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("stringFromDate:toDate:"), objc.ExtractPtr(fromDate), objc.ExtractPtr(toDate))
	return rv
}

func (d_ DateIntervalFormatter) StringFromDateInterval(dateInterval IDateInterval) string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("stringFromDateInterval:"), objc.ExtractPtr(dateInterval))
	return rv
}

func (d_ DateIntervalFormatter) DateStyle() DateIntervalFormatterStyle {
	rv := objc.CallMethod[DateIntervalFormatterStyle](d_, objc.GetSelector("dateStyle"))
	return rv
}

func (d_ DateIntervalFormatter) SetDateStyle(value DateIntervalFormatterStyle) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDateStyle:"), value)
}

func (d_ DateIntervalFormatter) TimeStyle() DateIntervalFormatterStyle {
	rv := objc.CallMethod[DateIntervalFormatterStyle](d_, objc.GetSelector("timeStyle"))
	return rv
}

func (d_ DateIntervalFormatter) SetTimeStyle(value DateIntervalFormatterStyle) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setTimeStyle:"), value)
}

func (d_ DateIntervalFormatter) DateTemplate() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("dateTemplate"))
	return rv
}

func (d_ DateIntervalFormatter) SetDateTemplate(value string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDateTemplate:"), value)
}

func (d_ DateIntervalFormatter) Calendar() Calendar {
	rv := objc.CallMethod[Calendar](d_, objc.GetSelector("calendar"))
	return rv
}

func (d_ DateIntervalFormatter) SetCalendar(value ICalendar) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setCalendar:"), objc.ExtractPtr(value))
}

func (d_ DateIntervalFormatter) Locale() Locale {
	rv := objc.CallMethod[Locale](d_, objc.GetSelector("locale"))
	return rv
}

func (d_ DateIntervalFormatter) SetLocale(value ILocale) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setLocale:"), objc.ExtractPtr(value))
}

func (d_ DateIntervalFormatter) TimeZone() TimeZone {
	rv := objc.CallMethod[TimeZone](d_, objc.GetSelector("timeZone"))
	return rv
}

func (d_ DateIntervalFormatter) SetTimeZone(value ITimeZone) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setTimeZone:"), objc.ExtractPtr(value))
}
