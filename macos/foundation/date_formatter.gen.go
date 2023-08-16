// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DateFormatter] class.
var DateFormatterClass = _DateFormatterClass{objc.GetClass("NSDateFormatter")}

type _DateFormatterClass struct {
	objc.Class
}

// An interface definition for the [DateFormatter] class.
type IDateFormatter interface {
	IFormatter
	DateFromString(string_ string) Date
	GetObjectValueForStringRangeError(obj objc.IObject, string_ string, rangep *Range, error IError) bool
	StringFromDate(date IDate) string
	SetLocalizedDateFormatFromTemplate(dateFormatTemplate string)
	IsLenient() bool
	SetLenient(value bool)
	VeryShortStandaloneMonthSymbols() []string
	SetVeryShortStandaloneMonthSymbols(value []string)
	ShortStandaloneWeekdaySymbols() []string
	SetShortStandaloneWeekdaySymbols(value []string)
	DoesRelativeDateFormatting() bool
	SetDoesRelativeDateFormatting(value bool)
	WeekdaySymbols() []string
	SetWeekdaySymbols(value []string)
	PMSymbol() string
	SetPMSymbol(value string)
	StandaloneQuarterSymbols() []string
	SetStandaloneQuarterSymbols(value []string)
	MonthSymbols() []string
	SetMonthSymbols(value []string)
	FormatterBehavior() DateFormatterBehavior
	SetFormatterBehavior(value DateFormatterBehavior)
	DateFormat() string
	SetDateFormat(value string)
	ShortWeekdaySymbols() []string
	SetShortWeekdaySymbols(value []string)
	Locale() Locale
	SetLocale(value ILocale)
	EraSymbols() []string
	SetEraSymbols(value []string)
	StandaloneWeekdaySymbols() []string
	SetStandaloneWeekdaySymbols(value []string)
	DefaultDate() Date
	SetDefaultDate(value IDate)
	TimeZone() TimeZone
	SetTimeZone(value ITimeZone)
	TimeStyle() DateFormatterStyle
	SetTimeStyle(value DateFormatterStyle)
	VeryShortMonthSymbols() []string
	SetVeryShortMonthSymbols(value []string)
	QuarterSymbols() []string
	SetQuarterSymbols(value []string)
	ShortStandaloneQuarterSymbols() []string
	SetShortStandaloneQuarterSymbols(value []string)
	TwoDigitStartDate() Date
	SetTwoDigitStartDate(value IDate)
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	AMSymbol() string
	SetAMSymbol(value string)
	DateStyle() DateFormatterStyle
	SetDateStyle(value DateFormatterStyle)
	VeryShortStandaloneWeekdaySymbols() []string
	SetVeryShortStandaloneWeekdaySymbols(value []string)
	ShortStandaloneMonthSymbols() []string
	SetShortStandaloneMonthSymbols(value []string)
	GregorianStartDate() Date
	SetGregorianStartDate(value IDate)
	Calendar() Calendar
	SetCalendar(value ICalendar)
	LongEraSymbols() []string
	SetLongEraSymbols(value []string)
	ShortMonthSymbols() []string
	SetShortMonthSymbols(value []string)
	VeryShortWeekdaySymbols() []string
	SetVeryShortWeekdaySymbols(value []string)
	ShortQuarterSymbols() []string
	SetShortQuarterSymbols(value []string)
	StandaloneMonthSymbols() []string
	SetStandaloneMonthSymbols(value []string)
	GeneratesCalendarDates() bool
	SetGeneratesCalendarDates(value bool)
}

// A formatter that converts between dates and their textual representations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter?language=objc
type DateFormatter struct {
	Formatter
}

func DateFormatterFrom(ptr unsafe.Pointer) DateFormatter {
	return DateFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

func (dc _DateFormatterClass) Alloc() DateFormatter {
	rv := objc.Call[DateFormatter](dc, objc.Sel("alloc"))
	return rv
}

func DateFormatter_Alloc() DateFormatter {
	return DateFormatterClass.Alloc()
}

func (dc _DateFormatterClass) New() DateFormatter {
	rv := objc.Call[DateFormatter](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDateFormatter() DateFormatter {
	return DateFormatterClass.New()
}

func (d_ DateFormatter) Init() DateFormatter {
	rv := objc.Call[DateFormatter](d_, objc.Sel("init"))
	return rv
}

// Returns a string representation of a specified date, that the system formats for the current locale using the specified date and time styles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1415241-localizedstringfromdate?language=objc
func (dc _DateFormatterClass) LocalizedStringFromDateDateStyleTimeStyle(date IDate, dstyle DateFormatterStyle, tstyle DateFormatterStyle) string {
	rv := objc.Call[string](dc, objc.Sel("localizedStringFromDate:dateStyle:timeStyle:"), objc.Ptr(date), dstyle, tstyle)
	return rv
}

// Returns a string representation of a specified date, that the system formats for the current locale using the specified date and time styles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1415241-localizedstringfromdate?language=objc
func DateFormatter_LocalizedStringFromDateDateStyleTimeStyle(date IDate, dstyle DateFormatterStyle, tstyle DateFormatterStyle) string {
	return DateFormatterClass.LocalizedStringFromDateDateStyleTimeStyle(date, dstyle, tstyle)
}

// Returns a localized date format string representing the given date format components arranged appropriately for the specified locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1408112-dateformatfromtemplate?language=objc
func (dc _DateFormatterClass) DateFormatFromTemplateOptionsLocale(tmplate string, opts uint, locale ILocale) string {
	rv := objc.Call[string](dc, objc.Sel("dateFormatFromTemplate:options:locale:"), tmplate, opts, objc.Ptr(locale))
	return rv
}

// Returns a localized date format string representing the given date format components arranged appropriately for the specified locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1408112-dateformatfromtemplate?language=objc
func DateFormatter_DateFormatFromTemplateOptionsLocale(tmplate string, opts uint, locale ILocale) string {
	return DateFormatterClass.DateFormatFromTemplateOptionsLocale(tmplate, opts, locale)
}

// Returns a date representation of a specified string that the system interprets using the receiver’s current settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1409994-datefromstring?language=objc
func (d_ DateFormatter) DateFromString(string_ string) Date {
	rv := objc.Call[Date](d_, objc.Sel("dateFromString:"), string_)
	return rv
}

// Returns by reference a date representation of a specified string and its date range, as well as a Boolean value that indicates whether the system can parse the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1409248-getobjectvalue?language=objc
func (d_ DateFormatter) GetObjectValueForStringRangeError(obj objc.IObject, string_ string, rangep *Range, error IError) bool {
	rv := objc.Call[bool](d_, objc.Sel("getObjectValue:forString:range:error:"), obj, string_, rangep, objc.Ptr(error))
	return rv
}

// Returns a string representation of a specified date that the system formats using the receiver’s current settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1415810-stringfromdate?language=objc
func (d_ DateFormatter) StringFromDate(date IDate) string {
	rv := objc.Call[string](d_, objc.Sel("stringFromDate:"), objc.Ptr(date))
	return rv
}

// Sets the date format from a template using the specified locale for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1417087-setlocalizeddateformatfromtempla?language=objc
func (d_ DateFormatter) SetLocalizedDateFormatFromTemplate(dateFormatTemplate string) {
	objc.Call[objc.Void](d_, objc.Sel("setLocalizedDateFormatFromTemplate:"), dateFormatTemplate)
}

// A Boolean value that indicates whether the receiver uses heuristics when parsing a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1411441-lenient?language=objc
func (d_ DateFormatter) IsLenient() bool {
	rv := objc.Call[bool](d_, objc.Sel("isLenient"))
	return rv
}

// A Boolean value that indicates whether the receiver uses heuristics when parsing a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1411441-lenient?language=objc
func (d_ DateFormatter) SetLenient(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setLenient:"), value)
}

// The very short month symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1413322-veryshortstandalonemonthsymbols?language=objc
func (d_ DateFormatter) VeryShortStandaloneMonthSymbols() []string {
	rv := objc.Call[[]string](d_, objc.Sel("veryShortStandaloneMonthSymbols"))
	return rv
}

// The very short month symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1413322-veryshortstandalonemonthsymbols?language=objc
func (d_ DateFormatter) SetVeryShortStandaloneMonthSymbols(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setVeryShortStandaloneMonthSymbols:"), value)
}

// The array of short standalone weekday symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1409119-shortstandaloneweekdaysymbols?language=objc
func (d_ DateFormatter) ShortStandaloneWeekdaySymbols() []string {
	rv := objc.Call[[]string](d_, objc.Sel("shortStandaloneWeekdaySymbols"))
	return rv
}

// The array of short standalone weekday symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1409119-shortstandaloneweekdaysymbols?language=objc
func (d_ DateFormatter) SetShortStandaloneWeekdaySymbols(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setShortStandaloneWeekdaySymbols:"), value)
}

// A Boolean value that indicates whether the receiver uses phrases such as “today” and “tomorrow” for the date component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1415848-doesrelativedateformatting?language=objc
func (d_ DateFormatter) DoesRelativeDateFormatting() bool {
	rv := objc.Call[bool](d_, objc.Sel("doesRelativeDateFormatting"))
	return rv
}

// A Boolean value that indicates whether the receiver uses phrases such as “today” and “tomorrow” for the date component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1415848-doesrelativedateformatting?language=objc
func (d_ DateFormatter) SetDoesRelativeDateFormatting(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setDoesRelativeDateFormatting:"), value)
}

// The array of weekday symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1412405-weekdaysymbols?language=objc
func (d_ DateFormatter) WeekdaySymbols() []string {
	rv := objc.Call[[]string](d_, objc.Sel("weekdaySymbols"))
	return rv
}

// The array of weekday symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1412405-weekdaysymbols?language=objc
func (d_ DateFormatter) SetWeekdaySymbols(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setWeekdaySymbols:"), value)
}

// The PM symbol for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1412367-pmsymbol?language=objc
func (d_ DateFormatter) PMSymbol() string {
	rv := objc.Call[string](d_, objc.Sel("PMSymbol"))
	return rv
}

// The PM symbol for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1412367-pmsymbol?language=objc
func (d_ DateFormatter) SetPMSymbol(value string) {
	objc.Call[objc.Void](d_, objc.Sel("setPMSymbol:"), value)
}

// The standalone quarter symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1411487-standalonequartersymbols?language=objc
func (d_ DateFormatter) StandaloneQuarterSymbols() []string {
	rv := objc.Call[[]string](d_, objc.Sel("standaloneQuarterSymbols"))
	return rv
}

// The standalone quarter symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1411487-standalonequartersymbols?language=objc
func (d_ DateFormatter) SetStandaloneQuarterSymbols(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setStandaloneQuarterSymbols:"), value)
}

// The month symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1412049-monthsymbols?language=objc
func (d_ DateFormatter) MonthSymbols() []string {
	rv := objc.Call[[]string](d_, objc.Sel("monthSymbols"))
	return rv
}

// The month symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1412049-monthsymbols?language=objc
func (d_ DateFormatter) SetMonthSymbols(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setMonthSymbols:"), value)
}

// The formatter behavior for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1409720-formatterbehavior?language=objc
func (d_ DateFormatter) FormatterBehavior() DateFormatterBehavior {
	rv := objc.Call[DateFormatterBehavior](d_, objc.Sel("formatterBehavior"))
	return rv
}

// The formatter behavior for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1409720-formatterbehavior?language=objc
func (d_ DateFormatter) SetFormatterBehavior(value DateFormatterBehavior) {
	objc.Call[objc.Void](d_, objc.Sel("setFormatterBehavior:"), value)
}

// The date format string used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1413514-dateformat?language=objc
func (d_ DateFormatter) DateFormat() string {
	rv := objc.Call[string](d_, objc.Sel("dateFormat"))
	return rv
}

// The date format string used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1413514-dateformat?language=objc
func (d_ DateFormatter) SetDateFormat(value string) {
	objc.Call[objc.Void](d_, objc.Sel("setDateFormat:"), value)
}

// The array of short weekday symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1416121-shortweekdaysymbols?language=objc
func (d_ DateFormatter) ShortWeekdaySymbols() []string {
	rv := objc.Call[[]string](d_, objc.Sel("shortWeekdaySymbols"))
	return rv
}

// The array of short weekday symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1416121-shortweekdaysymbols?language=objc
func (d_ DateFormatter) SetShortWeekdaySymbols(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setShortWeekdaySymbols:"), value)
}

// The locale for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1411973-locale?language=objc
func (d_ DateFormatter) Locale() Locale {
	rv := objc.Call[Locale](d_, objc.Sel("locale"))
	return rv
}

// The locale for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1411973-locale?language=objc
func (d_ DateFormatter) SetLocale(value ILocale) {
	objc.Call[objc.Void](d_, objc.Sel("setLocale:"), objc.Ptr(value))
}

// The era symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1418282-erasymbols?language=objc
func (d_ DateFormatter) EraSymbols() []string {
	rv := objc.Call[[]string](d_, objc.Sel("eraSymbols"))
	return rv
}

// The era symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1418282-erasymbols?language=objc
func (d_ DateFormatter) SetEraSymbols(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setEraSymbols:"), value)
}

// The array of standalone weekday symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1413618-standaloneweekdaysymbols?language=objc
func (d_ DateFormatter) StandaloneWeekdaySymbols() []string {
	rv := objc.Call[[]string](d_, objc.Sel("standaloneWeekdaySymbols"))
	return rv
}

// The array of standalone weekday symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1413618-standaloneweekdaysymbols?language=objc
func (d_ DateFormatter) SetStandaloneWeekdaySymbols(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setStandaloneWeekdaySymbols:"), value)
}

// The default date for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1408312-defaultdate?language=objc
func (d_ DateFormatter) DefaultDate() Date {
	rv := objc.Call[Date](d_, objc.Sel("defaultDate"))
	return rv
}

// The default date for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1408312-defaultdate?language=objc
func (d_ DateFormatter) SetDefaultDate(value IDate) {
	objc.Call[objc.Void](d_, objc.Sel("setDefaultDate:"), objc.Ptr(value))
}

// The time zone for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1411406-timezone?language=objc
func (d_ DateFormatter) TimeZone() TimeZone {
	rv := objc.Call[TimeZone](d_, objc.Sel("timeZone"))
	return rv
}

// The time zone for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1411406-timezone?language=objc
func (d_ DateFormatter) SetTimeZone(value ITimeZone) {
	objc.Call[objc.Void](d_, objc.Sel("setTimeZone:"), objc.Ptr(value))
}

// The time style of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1413467-timestyle?language=objc
func (d_ DateFormatter) TimeStyle() DateFormatterStyle {
	rv := objc.Call[DateFormatterStyle](d_, objc.Sel("timeStyle"))
	return rv
}

// The time style of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1413467-timestyle?language=objc
func (d_ DateFormatter) SetTimeStyle(value DateFormatterStyle) {
	objc.Call[objc.Void](d_, objc.Sel("setTimeStyle:"), value)
}

// The very short month symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1413632-veryshortmonthsymbols?language=objc
func (d_ DateFormatter) VeryShortMonthSymbols() []string {
	rv := objc.Call[[]string](d_, objc.Sel("veryShortMonthSymbols"))
	return rv
}

// The very short month symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1413632-veryshortmonthsymbols?language=objc
func (d_ DateFormatter) SetVeryShortMonthSymbols(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setVeryShortMonthSymbols:"), value)
}

// The quarter symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1417587-quartersymbols?language=objc
func (d_ DateFormatter) QuarterSymbols() []string {
	rv := objc.Call[[]string](d_, objc.Sel("quarterSymbols"))
	return rv
}

// The quarter symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1417587-quartersymbols?language=objc
func (d_ DateFormatter) SetQuarterSymbols(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setQuarterSymbols:"), value)
}

// The short standalone quarter symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1416421-shortstandalonequartersymbols?language=objc
func (d_ DateFormatter) ShortStandaloneQuarterSymbols() []string {
	rv := objc.Call[[]string](d_, objc.Sel("shortStandaloneQuarterSymbols"))
	return rv
}

// The short standalone quarter symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1416421-shortstandalonequartersymbols?language=objc
func (d_ DateFormatter) SetShortStandaloneQuarterSymbols(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setShortStandaloneQuarterSymbols:"), value)
}

// The earliest date that can be denoted by a two-digit year specifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1417203-twodigitstartdate?language=objc
func (d_ DateFormatter) TwoDigitStartDate() Date {
	rv := objc.Call[Date](d_, objc.Sel("twoDigitStartDate"))
	return rv
}

// The earliest date that can be denoted by a two-digit year specifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1417203-twodigitstartdate?language=objc
func (d_ DateFormatter) SetTwoDigitStartDate(value IDate) {
	objc.Call[objc.Void](d_, objc.Sel("setTwoDigitStartDate:"), objc.Ptr(value))
}

// Returns the default formatting behavior for instances of the class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1409266-defaultformatterbehavior?language=objc
func (dc _DateFormatterClass) DefaultFormatterBehavior() DateFormatterBehavior {
	rv := objc.Call[DateFormatterBehavior](dc, objc.Sel("defaultFormatterBehavior"))
	return rv
}

// Returns the default formatting behavior for instances of the class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1409266-defaultformatterbehavior?language=objc
func DateFormatter_DefaultFormatterBehavior() DateFormatterBehavior {
	return DateFormatterClass.DefaultFormatterBehavior()
}

// Returns the default formatting behavior for instances of the class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1409266-defaultformatterbehavior?language=objc
func (dc _DateFormatterClass) SetDefaultFormatterBehavior(value DateFormatterBehavior) {
	objc.Call[objc.Void](dc, objc.Sel("setDefaultFormatterBehavior:"), value)
}

// Returns the default formatting behavior for instances of the class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1409266-defaultformatterbehavior?language=objc
func DateFormatter_SetDefaultFormatterBehavior(value DateFormatterBehavior) {
	DateFormatterClass.SetDefaultFormatterBehavior(value)
}

// The capitalization formatting context used when formatting a date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1408066-formattingcontext?language=objc
func (d_ DateFormatter) FormattingContext() FormattingContext {
	rv := objc.Call[FormattingContext](d_, objc.Sel("formattingContext"))
	return rv
}

// The capitalization formatting context used when formatting a date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1408066-formattingcontext?language=objc
func (d_ DateFormatter) SetFormattingContext(value FormattingContext) {
	objc.Call[objc.Void](d_, objc.Sel("setFormattingContext:"), value)
}

// The AM symbol for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1409506-amsymbol?language=objc
func (d_ DateFormatter) AMSymbol() string {
	rv := objc.Call[string](d_, objc.Sel("AMSymbol"))
	return rv
}

// The AM symbol for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1409506-amsymbol?language=objc
func (d_ DateFormatter) SetAMSymbol(value string) {
	objc.Call[objc.Void](d_, objc.Sel("setAMSymbol:"), value)
}

// The date style of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1415411-datestyle?language=objc
func (d_ DateFormatter) DateStyle() DateFormatterStyle {
	rv := objc.Call[DateFormatterStyle](d_, objc.Sel("dateStyle"))
	return rv
}

// The date style of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1415411-datestyle?language=objc
func (d_ DateFormatter) SetDateStyle(value DateFormatterStyle) {
	objc.Call[objc.Void](d_, objc.Sel("setDateStyle:"), value)
}

// The array of very short standalone weekday symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1418238-veryshortstandaloneweekdaysymbol?language=objc
func (d_ DateFormatter) VeryShortStandaloneWeekdaySymbols() []string {
	rv := objc.Call[[]string](d_, objc.Sel("veryShortStandaloneWeekdaySymbols"))
	return rv
}

// The array of very short standalone weekday symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1418238-veryshortstandaloneweekdaysymbol?language=objc
func (d_ DateFormatter) SetVeryShortStandaloneWeekdaySymbols(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setVeryShortStandaloneWeekdaySymbols:"), value)
}

// The short standalone month symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1414771-shortstandalonemonthsymbols?language=objc
func (d_ DateFormatter) ShortStandaloneMonthSymbols() []string {
	rv := objc.Call[[]string](d_, objc.Sel("shortStandaloneMonthSymbols"))
	return rv
}

// The short standalone month symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1414771-shortstandalonemonthsymbols?language=objc
func (d_ DateFormatter) SetShortStandaloneMonthSymbols(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setShortStandaloneMonthSymbols:"), value)
}

// The start date of the Gregorian calendar for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1416389-gregorianstartdate?language=objc
func (d_ DateFormatter) GregorianStartDate() Date {
	rv := objc.Call[Date](d_, objc.Sel("gregorianStartDate"))
	return rv
}

// The start date of the Gregorian calendar for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1416389-gregorianstartdate?language=objc
func (d_ DateFormatter) SetGregorianStartDate(value IDate) {
	objc.Call[objc.Void](d_, objc.Sel("setGregorianStartDate:"), objc.Ptr(value))
}

// The calendar for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1413675-calendar?language=objc
func (d_ DateFormatter) Calendar() Calendar {
	rv := objc.Call[Calendar](d_, objc.Sel("calendar"))
	return rv
}

// The calendar for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1413675-calendar?language=objc
func (d_ DateFormatter) SetCalendar(value ICalendar) {
	objc.Call[objc.Void](d_, objc.Sel("setCalendar:"), objc.Ptr(value))
}

// The long era symbols for the receiver [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1418081-longerasymbols?language=objc
func (d_ DateFormatter) LongEraSymbols() []string {
	rv := objc.Call[[]string](d_, objc.Sel("longEraSymbols"))
	return rv
}

// The long era symbols for the receiver [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1418081-longerasymbols?language=objc
func (d_ DateFormatter) SetLongEraSymbols(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setLongEraSymbols:"), value)
}

// The array of short month symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1409209-shortmonthsymbols?language=objc
func (d_ DateFormatter) ShortMonthSymbols() []string {
	rv := objc.Call[[]string](d_, objc.Sel("shortMonthSymbols"))
	return rv
}

// The array of short month symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1409209-shortmonthsymbols?language=objc
func (d_ DateFormatter) SetShortMonthSymbols(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setShortMonthSymbols:"), value)
}

// The array of very short weekday symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1415109-veryshortweekdaysymbols?language=objc
func (d_ DateFormatter) VeryShortWeekdaySymbols() []string {
	rv := objc.Call[[]string](d_, objc.Sel("veryShortWeekdaySymbols"))
	return rv
}

// The array of very short weekday symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1415109-veryshortweekdaysymbols?language=objc
func (d_ DateFormatter) SetVeryShortWeekdaySymbols(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setVeryShortWeekdaySymbols:"), value)
}

// The short quarter symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1409851-shortquartersymbols?language=objc
func (d_ DateFormatter) ShortQuarterSymbols() []string {
	rv := objc.Call[[]string](d_, objc.Sel("shortQuarterSymbols"))
	return rv
}

// The short quarter symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1409851-shortquartersymbols?language=objc
func (d_ DateFormatter) SetShortQuarterSymbols(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setShortQuarterSymbols:"), value)
}

// The standalone month symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1416227-standalonemonthsymbols?language=objc
func (d_ DateFormatter) StandaloneMonthSymbols() []string {
	rv := objc.Call[[]string](d_, objc.Sel("standaloneMonthSymbols"))
	return rv
}

// The standalone month symbols for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1416227-standalonemonthsymbols?language=objc
func (d_ DateFormatter) SetStandaloneMonthSymbols(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setStandaloneMonthSymbols:"), value)
}

// Indicates whether the formatter generates the deprecated calendar date type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1411107-generatescalendardates?language=objc
func (d_ DateFormatter) GeneratesCalendarDates() bool {
	rv := objc.Call[bool](d_, objc.Sel("generatesCalendarDates"))
	return rv
}

// Indicates whether the formatter generates the deprecated calendar date type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateformatter/1411107-generatescalendardates?language=objc
func (d_ DateFormatter) SetGeneratesCalendarDates(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setGeneratesCalendarDates:"), value)
}
