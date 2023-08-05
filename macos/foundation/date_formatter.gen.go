// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var DateFormatterClass = _DateFormatterClass{objc.GetClass("NSDateFormatter")}

type _DateFormatterClass struct {
	objc.Class
}

type IDateFormatter interface {
	IFormatter
	DateFromString(string_ string) Date
	StringFromDate(date IDate) string
	SetLocalizedDateFormatFromTemplate(dateFormatTemplate string)
	DateStyle() DateFormatterStyle
	SetDateStyle(value DateFormatterStyle)
	TimeStyle() DateFormatterStyle
	SetTimeStyle(value DateFormatterStyle)
	DateFormat() string
	SetDateFormat(value string)
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	Calendar() Calendar
	SetCalendar(value ICalendar)
	DefaultDate() Date
	SetDefaultDate(value IDate)
	Locale() Locale
	SetLocale(value ILocale)
	TimeZone() TimeZone
	SetTimeZone(value ITimeZone)
	TwoDigitStartDate() Date
	SetTwoDigitStartDate(value IDate)
	GregorianStartDate() Date
	SetGregorianStartDate(value IDate)
	FormatterBehavior() DateFormatterBehavior
	SetFormatterBehavior(value DateFormatterBehavior)
	IsLenient() bool
	SetLenient(value bool)
	DoesRelativeDateFormatting() bool
	SetDoesRelativeDateFormatting(value bool)
	AMSymbol() string
	SetAMSymbol(value string)
	PMSymbol() string
	SetPMSymbol(value string)
	WeekdaySymbols() []string
	SetWeekdaySymbols(value []string)
	ShortWeekdaySymbols() []string
	SetShortWeekdaySymbols(value []string)
	VeryShortWeekdaySymbols() []string
	SetVeryShortWeekdaySymbols(value []string)
	StandaloneWeekdaySymbols() []string
	SetStandaloneWeekdaySymbols(value []string)
	ShortStandaloneWeekdaySymbols() []string
	SetShortStandaloneWeekdaySymbols(value []string)
	VeryShortStandaloneWeekdaySymbols() []string
	SetVeryShortStandaloneWeekdaySymbols(value []string)
	MonthSymbols() []string
	SetMonthSymbols(value []string)
	ShortMonthSymbols() []string
	SetShortMonthSymbols(value []string)
	VeryShortMonthSymbols() []string
	SetVeryShortMonthSymbols(value []string)
	StandaloneMonthSymbols() []string
	SetStandaloneMonthSymbols(value []string)
	ShortStandaloneMonthSymbols() []string
	SetShortStandaloneMonthSymbols(value []string)
	VeryShortStandaloneMonthSymbols() []string
	SetVeryShortStandaloneMonthSymbols(value []string)
	QuarterSymbols() []string
	SetQuarterSymbols(value []string)
	ShortQuarterSymbols() []string
	SetShortQuarterSymbols(value []string)
	StandaloneQuarterSymbols() []string
	SetStandaloneQuarterSymbols(value []string)
	ShortStandaloneQuarterSymbols() []string
	SetShortStandaloneQuarterSymbols(value []string)
	EraSymbols() []string
	SetEraSymbols(value []string)
	LongEraSymbols() []string
	SetLongEraSymbols(value []string)
	GeneratesCalendarDates() bool
	SetGeneratesCalendarDates(value bool)
}

type DateFormatter struct {
	Formatter
}

func MakeDateFormatter(ptr unsafe.Pointer) DateFormatter {
	return DateFormatter{
		Formatter: MakeFormatter(ptr),
	}
}

func (dc _DateFormatterClass) Alloc() DateFormatter {
	rv := objc.CallMethod[DateFormatter](dc, objc.GetSelector("alloc"))
	return rv
}

func DateFormatter_Alloc() DateFormatter {
	return DateFormatterClass.Alloc()
}

func (dc _DateFormatterClass) New() DateFormatter {
	rv := objc.CallMethod[DateFormatter](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDateFormatter() DateFormatter {
	return DateFormatterClass.New()
}

func DateFormatter_New() DateFormatter {
	return DateFormatterClass.New()
}

func (d_ DateFormatter) Init() DateFormatter {
	rv := objc.CallMethod[DateFormatter](d_, objc.GetSelector("init"))
	return rv
}

func DateFormatter_Init() DateFormatter {
	return DateFormatterClass.Alloc().Init()
}

func (d_ DateFormatter) DateFromString(string_ string) Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("dateFromString:"), string_)
	return rv
}

func (d_ DateFormatter) StringFromDate(date IDate) string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("stringFromDate:"), objc.ExtractPtr(date))
	return rv
}

func (dc _DateFormatterClass) LocalizedStringFromDateDateStyleTimeStyle(date IDate, dstyle DateFormatterStyle, tstyle DateFormatterStyle) string {
	rv := objc.CallMethod[string](dc, objc.GetSelector("localizedStringFromDate:dateStyle:timeStyle:"), objc.ExtractPtr(date), dstyle, tstyle)
	return rv
}

func DateFormatter_LocalizedStringFromDateDateStyleTimeStyle(date IDate, dstyle DateFormatterStyle, tstyle DateFormatterStyle) string {
	return DateFormatterClass.LocalizedStringFromDateDateStyleTimeStyle(date, dstyle, tstyle)
}

func (d_ DateFormatter) SetLocalizedDateFormatFromTemplate(dateFormatTemplate string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setLocalizedDateFormatFromTemplate:"), dateFormatTemplate)
}

func (dc _DateFormatterClass) DateFormatFromTemplateOptionsLocale(tmplate string, opts uint, locale ILocale) string {
	rv := objc.CallMethod[string](dc, objc.GetSelector("dateFormatFromTemplate:options:locale:"), tmplate, opts, objc.ExtractPtr(locale))
	return rv
}

func DateFormatter_DateFormatFromTemplateOptionsLocale(tmplate string, opts uint, locale ILocale) string {
	return DateFormatterClass.DateFormatFromTemplateOptionsLocale(tmplate, opts, locale)
}

func (d_ DateFormatter) DateStyle() DateFormatterStyle {
	rv := objc.CallMethod[DateFormatterStyle](d_, objc.GetSelector("dateStyle"))
	return rv
}

func (d_ DateFormatter) SetDateStyle(value DateFormatterStyle) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDateStyle:"), value)
}

func (d_ DateFormatter) TimeStyle() DateFormatterStyle {
	rv := objc.CallMethod[DateFormatterStyle](d_, objc.GetSelector("timeStyle"))
	return rv
}

func (d_ DateFormatter) SetTimeStyle(value DateFormatterStyle) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setTimeStyle:"), value)
}

func (d_ DateFormatter) DateFormat() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("dateFormat"))
	return rv
}

func (d_ DateFormatter) SetDateFormat(value string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDateFormat:"), value)
}

func (d_ DateFormatter) FormattingContext() FormattingContext {
	rv := objc.CallMethod[FormattingContext](d_, objc.GetSelector("formattingContext"))
	return rv
}

func (d_ DateFormatter) SetFormattingContext(value FormattingContext) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setFormattingContext:"), value)
}

func (d_ DateFormatter) Calendar() Calendar {
	rv := objc.CallMethod[Calendar](d_, objc.GetSelector("calendar"))
	return rv
}

func (d_ DateFormatter) SetCalendar(value ICalendar) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setCalendar:"), objc.ExtractPtr(value))
}

func (d_ DateFormatter) DefaultDate() Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("defaultDate"))
	return rv
}

func (d_ DateFormatter) SetDefaultDate(value IDate) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDefaultDate:"), objc.ExtractPtr(value))
}

func (d_ DateFormatter) Locale() Locale {
	rv := objc.CallMethod[Locale](d_, objc.GetSelector("locale"))
	return rv
}

func (d_ DateFormatter) SetLocale(value ILocale) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setLocale:"), objc.ExtractPtr(value))
}

func (d_ DateFormatter) TimeZone() TimeZone {
	rv := objc.CallMethod[TimeZone](d_, objc.GetSelector("timeZone"))
	return rv
}

func (d_ DateFormatter) SetTimeZone(value ITimeZone) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setTimeZone:"), objc.ExtractPtr(value))
}

func (d_ DateFormatter) TwoDigitStartDate() Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("twoDigitStartDate"))
	return rv
}

func (d_ DateFormatter) SetTwoDigitStartDate(value IDate) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setTwoDigitStartDate:"), objc.ExtractPtr(value))
}

func (d_ DateFormatter) GregorianStartDate() Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("gregorianStartDate"))
	return rv
}

func (d_ DateFormatter) SetGregorianStartDate(value IDate) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setGregorianStartDate:"), objc.ExtractPtr(value))
}

func (d_ DateFormatter) FormatterBehavior() DateFormatterBehavior {
	rv := objc.CallMethod[DateFormatterBehavior](d_, objc.GetSelector("formatterBehavior"))
	return rv
}

func (d_ DateFormatter) SetFormatterBehavior(value DateFormatterBehavior) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setFormatterBehavior:"), value)
}

func (dc _DateFormatterClass) DefaultFormatterBehavior() DateFormatterBehavior {
	rv := objc.CallMethod[DateFormatterBehavior](dc, objc.GetSelector("defaultFormatterBehavior"))
	return rv
}

func DateFormatter_DefaultFormatterBehavior() DateFormatterBehavior {
	return DateFormatterClass.DefaultFormatterBehavior()
}

func (dc _DateFormatterClass) SetDefaultFormatterBehavior(value DateFormatterBehavior) {
	objc.CallMethod[objc.Void](dc, objc.GetSelector("setDefaultFormatterBehavior:"), value)
}

func DateFormatter_SetDefaultFormatterBehavior(value DateFormatterBehavior) {
	DateFormatterClass.SetDefaultFormatterBehavior(value)
}

func (d_ DateFormatter) IsLenient() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isLenient"))
	return rv
}

func (d_ DateFormatter) SetLenient(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setLenient:"), value)
}

func (d_ DateFormatter) DoesRelativeDateFormatting() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("doesRelativeDateFormatting"))
	return rv
}

func (d_ DateFormatter) SetDoesRelativeDateFormatting(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDoesRelativeDateFormatting:"), value)
}

func (d_ DateFormatter) AMSymbol() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("AMSymbol"))
	return rv
}

func (d_ DateFormatter) SetAMSymbol(value string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setAMSymbol:"), value)
}

func (d_ DateFormatter) PMSymbol() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("PMSymbol"))
	return rv
}

func (d_ DateFormatter) SetPMSymbol(value string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setPMSymbol:"), value)
}

func (d_ DateFormatter) WeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("weekdaySymbols"))
	return rv
}

func (d_ DateFormatter) SetWeekdaySymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setWeekdaySymbols:"), value)
}

func (d_ DateFormatter) ShortWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("shortWeekdaySymbols"))
	return rv
}

func (d_ DateFormatter) SetShortWeekdaySymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setShortWeekdaySymbols:"), value)
}

func (d_ DateFormatter) VeryShortWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("veryShortWeekdaySymbols"))
	return rv
}

func (d_ DateFormatter) SetVeryShortWeekdaySymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setVeryShortWeekdaySymbols:"), value)
}

func (d_ DateFormatter) StandaloneWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("standaloneWeekdaySymbols"))
	return rv
}

func (d_ DateFormatter) SetStandaloneWeekdaySymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setStandaloneWeekdaySymbols:"), value)
}

func (d_ DateFormatter) ShortStandaloneWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("shortStandaloneWeekdaySymbols"))
	return rv
}

func (d_ DateFormatter) SetShortStandaloneWeekdaySymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setShortStandaloneWeekdaySymbols:"), value)
}

func (d_ DateFormatter) VeryShortStandaloneWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("veryShortStandaloneWeekdaySymbols"))
	return rv
}

func (d_ DateFormatter) SetVeryShortStandaloneWeekdaySymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setVeryShortStandaloneWeekdaySymbols:"), value)
}

func (d_ DateFormatter) MonthSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("monthSymbols"))
	return rv
}

func (d_ DateFormatter) SetMonthSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setMonthSymbols:"), value)
}

func (d_ DateFormatter) ShortMonthSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("shortMonthSymbols"))
	return rv
}

func (d_ DateFormatter) SetShortMonthSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setShortMonthSymbols:"), value)
}

func (d_ DateFormatter) VeryShortMonthSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("veryShortMonthSymbols"))
	return rv
}

func (d_ DateFormatter) SetVeryShortMonthSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setVeryShortMonthSymbols:"), value)
}

func (d_ DateFormatter) StandaloneMonthSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("standaloneMonthSymbols"))
	return rv
}

func (d_ DateFormatter) SetStandaloneMonthSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setStandaloneMonthSymbols:"), value)
}

func (d_ DateFormatter) ShortStandaloneMonthSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("shortStandaloneMonthSymbols"))
	return rv
}

func (d_ DateFormatter) SetShortStandaloneMonthSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setShortStandaloneMonthSymbols:"), value)
}

func (d_ DateFormatter) VeryShortStandaloneMonthSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("veryShortStandaloneMonthSymbols"))
	return rv
}

func (d_ DateFormatter) SetVeryShortStandaloneMonthSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setVeryShortStandaloneMonthSymbols:"), value)
}

func (d_ DateFormatter) QuarterSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("quarterSymbols"))
	return rv
}

func (d_ DateFormatter) SetQuarterSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setQuarterSymbols:"), value)
}

func (d_ DateFormatter) ShortQuarterSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("shortQuarterSymbols"))
	return rv
}

func (d_ DateFormatter) SetShortQuarterSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setShortQuarterSymbols:"), value)
}

func (d_ DateFormatter) StandaloneQuarterSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("standaloneQuarterSymbols"))
	return rv
}

func (d_ DateFormatter) SetStandaloneQuarterSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setStandaloneQuarterSymbols:"), value)
}

func (d_ DateFormatter) ShortStandaloneQuarterSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("shortStandaloneQuarterSymbols"))
	return rv
}

func (d_ DateFormatter) SetShortStandaloneQuarterSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setShortStandaloneQuarterSymbols:"), value)
}

func (d_ DateFormatter) EraSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("eraSymbols"))
	return rv
}

func (d_ DateFormatter) SetEraSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setEraSymbols:"), value)
}

func (d_ DateFormatter) LongEraSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("longEraSymbols"))
	return rv
}

func (d_ DateFormatter) SetLongEraSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setLongEraSymbols:"), value)
}

func (d_ DateFormatter) GeneratesCalendarDates() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("generatesCalendarDates"))
	return rv
}

func (d_ DateFormatter) SetGeneratesCalendarDates(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setGeneratesCalendarDates:"), value)
}
