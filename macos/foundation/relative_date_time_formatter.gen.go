// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RelativeDateTimeFormatter] class.
var RelativeDateTimeFormatterClass = _RelativeDateTimeFormatterClass{objc.GetClass("NSRelativeDateTimeFormatter")}

type _RelativeDateTimeFormatterClass struct {
	objc.Class
}

// An interface definition for the [RelativeDateTimeFormatter] class.
type IRelativeDateTimeFormatter interface {
	IFormatter
	LocalizedStringFromTimeInterval(timeInterval TimeInterval) string
	LocalizedStringFromDateComponents(dateComponents IDateComponents) string
	LocalizedStringForDateRelativeToDate(date IDate, referenceDate IDate) string
	UnitsStyle() RelativeDateTimeFormatterUnitsStyle
	SetUnitsStyle(value RelativeDateTimeFormatterUnitsStyle)
	Locale() Locale
	SetLocale(value ILocale)
	DateTimeStyle() RelativeDateTimeFormatterStyle
	SetDateTimeStyle(value RelativeDateTimeFormatterStyle)
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	Calendar() Calendar
	SetCalendar(value ICalendar)
}

// A formatter that creates locale-aware string representations of a relative date or time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativedatetimeformatter?language=objc
type RelativeDateTimeFormatter struct {
	Formatter
}

func RelativeDateTimeFormatterFrom(ptr unsafe.Pointer) RelativeDateTimeFormatter {
	return RelativeDateTimeFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

func (rc _RelativeDateTimeFormatterClass) Alloc() RelativeDateTimeFormatter {
	rv := objc.Call[RelativeDateTimeFormatter](rc, objc.Sel("alloc"))
	return rv
}

func RelativeDateTimeFormatter_Alloc() RelativeDateTimeFormatter {
	return RelativeDateTimeFormatterClass.Alloc()
}

func (rc _RelativeDateTimeFormatterClass) New() RelativeDateTimeFormatter {
	rv := objc.Call[RelativeDateTimeFormatter](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRelativeDateTimeFormatter() RelativeDateTimeFormatter {
	return RelativeDateTimeFormatterClass.New()
}

func (r_ RelativeDateTimeFormatter) Init() RelativeDateTimeFormatter {
	rv := objc.Call[RelativeDateTimeFormatter](r_, objc.Sel("init"))
	return rv
}

// Formats the specified time interval using the formatter’s calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativedatetimeformatter/3131000-localizedstringfromtimeinterval?language=objc
func (r_ RelativeDateTimeFormatter) LocalizedStringFromTimeInterval(timeInterval TimeInterval) string {
	rv := objc.Call[string](r_, objc.Sel("localizedStringFromTimeInterval:"), timeInterval)
	return rv
}

// Formats a relative time represented by the specified date components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativedatetimeformatter/3130999-localizedstringfromdatecomponent?language=objc
func (r_ RelativeDateTimeFormatter) LocalizedStringFromDateComponents(dateComponents IDateComponents) string {
	rv := objc.Call[string](r_, objc.Sel("localizedStringFromDateComponents:"), objc.Ptr(dateComponents))
	return rv
}

// Formats the date interval from the reference date to the specified date using the formatter’s calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativedatetimeformatter/3130998-localizedstringfordate?language=objc
func (r_ RelativeDateTimeFormatter) LocalizedStringForDateRelativeToDate(date IDate, referenceDate IDate) string {
	rv := objc.Call[string](r_, objc.Sel("localizedStringForDate:relativeToDate:"), objc.Ptr(date), objc.Ptr(referenceDate))
	return rv
}

// The style to use when formatting the quantity or the name of the unit, such as “1 day ago” or “one day ago”. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativedatetimeformatter/3131002-unitsstyle?language=objc
func (r_ RelativeDateTimeFormatter) UnitsStyle() RelativeDateTimeFormatterUnitsStyle {
	rv := objc.Call[RelativeDateTimeFormatterUnitsStyle](r_, objc.Sel("unitsStyle"))
	return rv
}

// The style to use when formatting the quantity or the name of the unit, such as “1 day ago” or “one day ago”. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativedatetimeformatter/3131002-unitsstyle?language=objc
func (r_ RelativeDateTimeFormatter) SetUnitsStyle(value RelativeDateTimeFormatterUnitsStyle) {
	objc.Call[objc.Void](r_, objc.Sel("setUnitsStyle:"), value)
}

// The locale to use when formatting the date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativedatetimeformatter/3130997-locale?language=objc
func (r_ RelativeDateTimeFormatter) Locale() Locale {
	rv := objc.Call[Locale](r_, objc.Sel("locale"))
	return rv
}

// The locale to use when formatting the date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativedatetimeformatter/3130997-locale?language=objc
func (r_ RelativeDateTimeFormatter) SetLocale(value ILocale) {
	objc.Call[objc.Void](r_, objc.Sel("setLocale:"), objc.Ptr(value))
}

// The style to use when describing a relative date, for example “yesterday” or “1 day ago”. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativedatetimeformatter/3130995-datetimestyle?language=objc
func (r_ RelativeDateTimeFormatter) DateTimeStyle() RelativeDateTimeFormatterStyle {
	rv := objc.Call[RelativeDateTimeFormatterStyle](r_, objc.Sel("dateTimeStyle"))
	return rv
}

// The style to use when describing a relative date, for example “yesterday” or “1 day ago”. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativedatetimeformatter/3130995-datetimestyle?language=objc
func (r_ RelativeDateTimeFormatter) SetDateTimeStyle(value RelativeDateTimeFormatterStyle) {
	objc.Call[objc.Void](r_, objc.Sel("setDateTimeStyle:"), value)
}

// A description of where the formatted string will appear, allowing the formatter to capitalize the output appropriately. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativedatetimeformatter/3130996-formattingcontext?language=objc
func (r_ RelativeDateTimeFormatter) FormattingContext() FormattingContext {
	rv := objc.Call[FormattingContext](r_, objc.Sel("formattingContext"))
	return rv
}

// A description of where the formatted string will appear, allowing the formatter to capitalize the output appropriately. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativedatetimeformatter/3130996-formattingcontext?language=objc
func (r_ RelativeDateTimeFormatter) SetFormattingContext(value FormattingContext) {
	objc.Call[objc.Void](r_, objc.Sel("setFormattingContext:"), value)
}

// The calendar to use for formatting values that don’t have an inherent calendar of their own. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativedatetimeformatter/3130994-calendar?language=objc
func (r_ RelativeDateTimeFormatter) Calendar() Calendar {
	rv := objc.Call[Calendar](r_, objc.Sel("calendar"))
	return rv
}

// The calendar to use for formatting values that don’t have an inherent calendar of their own. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrelativedatetimeformatter/3130994-calendar?language=objc
func (r_ RelativeDateTimeFormatter) SetCalendar(value ICalendar) {
	objc.Call[objc.Void](r_, objc.Sel("setCalendar:"), objc.Ptr(value))
}
