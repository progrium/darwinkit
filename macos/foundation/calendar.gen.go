// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Calendar] class.
var CalendarClass = _CalendarClass{objc.GetClass("NSCalendar")}

type _CalendarClass struct {
	objc.Class
}

// An interface definition for the [Calendar] class.
type ICalendar interface {
	objc.IObject
	ComponentsInTimeZoneFromDate(timezone ITimeZone, date IDate) DateComponents
	DateByAddingUnitValueToDateOptions(unit CalendarUnit, value int, date IDate, options CalendarOptions) Date
	DateMatchesComponents(date IDate, components IDateComponents) bool
	GetHourMinuteSecondNanosecondFromDate(hourValuePointer *int, minuteValuePointer *int, secondValuePointer *int, nanosecondValuePointer *int, date IDate)
	IsDateInToday(date IDate) bool
	DateFromComponents(comps IDateComponents) Date
	MaximumRangeOfUnit(unit CalendarUnit) Range
	DateByAddingComponentsToDateOptions(comps IDateComponents, date IDate, opts CalendarOptions) Date
	OrdinalityOfUnitInUnitForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) uint
	DateBySettingUnitValueOfDateOptions(unit CalendarUnit, v int, date IDate, opts CalendarOptions) Date
	RangeOfUnitInUnitForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) Range
	NextDateAfterDateMatchingUnitValueOptions(date IDate, unit CalendarUnit, value int, options CalendarOptions) Date
	ComponentFromDate(unit CalendarUnit, date IDate) int
	RangeOfWeekendStartDateIntervalContainingDate(datep IDate, tip *TimeInterval, date IDate) bool
	IsDateInWeekend(date IDate) bool
	ComponentsFromDate(unitFlags CalendarUnit, date IDate) DateComponents
	StartOfDayForDate(date IDate) Date
	GetEraYearMonthDayFromDate(eraValuePointer *int, yearValuePointer *int, monthValuePointer *int, dayValuePointer *int, date IDate)
	EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock(start IDate, comps IDateComponents, opts CalendarOptions, block func(date Date, exactMatch bool, stop *bool))
	InitWithCalendarIdentifier(ident CalendarIdentifier) objc.Object
	NextWeekendStartDateIntervalOptionsAfterDate(datep IDate, tip *TimeInterval, options CalendarOptions, date IDate) bool
	MinimumRangeOfUnit(unit CalendarUnit) Range
	IsDateInSameDayAsDate(date1 IDate, date2 IDate) bool
	DateWithEraYearMonthDayHourMinuteSecondNanosecond(eraValue int, yearValue int, monthValue int, dayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date
	IsDateInTomorrow(date IDate) bool
	DateBySettingHourMinuteSecondOfDateOptions(h int, m int, s int, date IDate, opts CalendarOptions) Date
	IsDateInYesterday(date IDate) bool
	CompareDateToDateToUnitGranularity(date1 IDate, date2 IDate, unit CalendarUnit) ComparisonResult
	FirstWeekday() uint
	SetFirstWeekday(value uint)
	VeryShortStandaloneMonthSymbols() []string
	ShortStandaloneWeekdaySymbols() []string
	WeekdaySymbols() []string
	PMSymbol() string
	StandaloneQuarterSymbols() []string
	MonthSymbols() []string
	ShortWeekdaySymbols() []string
	MinimumDaysInFirstWeek() uint
	SetMinimumDaysInFirstWeek(value uint)
	CalendarIdentifier() CalendarIdentifier
	Locale() Locale
	SetLocale(value ILocale)
	EraSymbols() []string
	StandaloneWeekdaySymbols() []string
	TimeZone() TimeZone
	SetTimeZone(value ITimeZone)
	VeryShortMonthSymbols() []string
	QuarterSymbols() []string
	ShortStandaloneQuarterSymbols() []string
	AMSymbol() string
	VeryShortStandaloneWeekdaySymbols() []string
	ShortStandaloneMonthSymbols() []string
	LongEraSymbols() []string
	ShortMonthSymbols() []string
	VeryShortWeekdaySymbols() []string
	ShortQuarterSymbols() []string
	StandaloneMonthSymbols() []string
}

// A definition of the relationships between calendar units and absolute points in time, providing features for calculation and comparison of dates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar?language=objc
type Calendar struct {
	objc.Object
}

func CalendarFrom(ptr unsafe.Pointer) Calendar {
	return Calendar{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CalendarClass) Alloc() Calendar {
	rv := objc.Call[Calendar](cc, objc.Sel("alloc"))
	return rv
}

func Calendar_Alloc() Calendar {
	return CalendarClass.Alloc()
}

func (cc _CalendarClass) New() Calendar {
	rv := objc.Call[Calendar](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCalendar() Calendar {
	return CalendarClass.New()
}

func (c_ Calendar) Init() Calendar {
	rv := objc.Call[Calendar](c_, objc.Sel("init"))
	return rv
}

// Returns all the date components of a date, as if in a given time zone (instead of the receiving calendar’s time zone). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1413194-componentsintimezone?language=objc
func (c_ Calendar) ComponentsInTimeZoneFromDate(timezone ITimeZone, date IDate) DateComponents {
	rv := objc.Call[DateComponents](c_, objc.Sel("componentsInTimeZone:fromDate:"), objc.Ptr(timezone), objc.Ptr(date))
	return rv
}

// Returns a date representing the absolute time calculated by adding the value of a given component to a given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1407989-datebyaddingunit?language=objc
func (c_ Calendar) DateByAddingUnitValueToDateOptions(unit CalendarUnit, value int, date IDate, options CalendarOptions) Date {
	rv := objc.Call[Date](c_, objc.Sel("dateByAddingUnit:value:toDate:options:"), unit, value, objc.Ptr(date), options)
	return rv
}

// Returns whether a given date matches all of the given date components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1407954-date?language=objc
func (c_ Calendar) DateMatchesComponents(date IDate, components IDateComponents) bool {
	rv := objc.Call[bool](c_, objc.Sel("date:matchesComponents:"), objc.Ptr(date), objc.Ptr(components))
	return rv
}

// Returns by reference the hour, minute, second, and nanosecond component values for a given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1415012-gethour?language=objc
func (c_ Calendar) GetHourMinuteSecondNanosecondFromDate(hourValuePointer *int, minuteValuePointer *int, secondValuePointer *int, nanosecondValuePointer *int, date IDate) {
	objc.Call[objc.Void](c_, objc.Sel("getHour:minute:second:nanosecond:fromDate:"), hourValuePointer, minuteValuePointer, secondValuePointer, nanosecondValuePointer, objc.Ptr(date))
}

// Indicates whether the given date is in “today.” [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1417149-isdateintoday?language=objc
func (c_ Calendar) IsDateInToday(date IDate) bool {
	rv := objc.Call[bool](c_, objc.Sel("isDateInToday:"), objc.Ptr(date))
	return rv
}

// Returns a date representing the absolute time calculated from given components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1407609-datefromcomponents?language=objc
func (c_ Calendar) DateFromComponents(comps IDateComponents) Date {
	rv := objc.Call[Date](c_, objc.Sel("dateFromComponents:"), objc.Ptr(comps))
	return rv
}

// Creates a new calendar specified by a given identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1412400-calendarwithidentifier?language=objc
func (cc _CalendarClass) CalendarWithIdentifier(calendarIdentifierConstant CalendarIdentifier) Calendar {
	rv := objc.Call[Calendar](cc, objc.Sel("calendarWithIdentifier:"), calendarIdentifierConstant)
	return rv
}

// Creates a new calendar specified by a given identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1412400-calendarwithidentifier?language=objc
func Calendar_CalendarWithIdentifier(calendarIdentifierConstant CalendarIdentifier) Calendar {
	return CalendarClass.CalendarWithIdentifier(calendarIdentifierConstant)
}

// Returns the maximum range limits of the values that a given unit can take on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1414251-maximumrangeofunit?language=objc
func (c_ Calendar) MaximumRangeOfUnit(unit CalendarUnit) Range {
	rv := objc.Call[Range](c_, objc.Sel("maximumRangeOfUnit:"), unit)
	return rv
}

// Returns a date representing the absolute time calculated by adding given components to a given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1409577-datebyaddingcomponents?language=objc
func (c_ Calendar) DateByAddingComponentsToDateOptions(comps IDateComponents, date IDate, opts CalendarOptions) Date {
	rv := objc.Call[Date](c_, objc.Sel("dateByAddingComponents:toDate:options:"), objc.Ptr(comps), objc.Ptr(date), opts)
	return rv
}

// Returns, for a given absolute time, the ordinal number of a smaller calendar unit (such as a day) within a specified larger calendar unit (such as a week). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1408595-ordinalityofunit?language=objc
func (c_ Calendar) OrdinalityOfUnitInUnitForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) uint {
	rv := objc.Call[uint](c_, objc.Sel("ordinalityOfUnit:inUnit:forDate:"), smaller, larger, objc.Ptr(date))
	return rv
}

// Returns a new date representing the date calculated by setting a specific component of a given date to a given value, while trying to keep lower components the same. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1417915-datebysettingunit?language=objc
func (c_ Calendar) DateBySettingUnitValueOfDateOptions(unit CalendarUnit, v int, date IDate, opts CalendarOptions) Date {
	rv := objc.Call[Date](c_, objc.Sel("dateBySettingUnit:value:ofDate:options:"), unit, v, objc.Ptr(date), opts)
	return rv
}

// Returns the range of absolute time values that a smaller calendar unit (such as a day) can take on in a larger calendar unit (such as a month) that includes a specified absolute time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1418344-rangeofunit?language=objc
func (c_ Calendar) RangeOfUnitInUnitForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) Range {
	rv := objc.Call[Range](c_, objc.Sel("rangeOfUnit:inUnit:forDate:"), smaller, larger, objc.Ptr(date))
	return rv
}

// Returns the next date after a given date matching the given calendar unit value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1417170-nextdateafterdate?language=objc
func (c_ Calendar) NextDateAfterDateMatchingUnitValueOptions(date IDate, unit CalendarUnit, value int, options CalendarOptions) Date {
	rv := objc.Call[Date](c_, objc.Sel("nextDateAfterDate:matchingUnit:value:options:"), objc.Ptr(date), unit, value, options)
	return rv
}

// Returns the specified date component from a given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1416505-component?language=objc
func (c_ Calendar) ComponentFromDate(unit CalendarUnit, date IDate) int {
	rv := objc.Call[int](c_, objc.Sel("component:fromDate:"), unit, objc.Ptr(date))
	return rv
}

// Returns whether a given date falls within a weekend period, and if so, returns by reference the start date and time interval of the weekend range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1413286-rangeofweekendstartdate?language=objc
func (c_ Calendar) RangeOfWeekendStartDateIntervalContainingDate(datep IDate, tip *TimeInterval, date IDate) bool {
	rv := objc.Call[bool](c_, objc.Sel("rangeOfWeekendStartDate:interval:containingDate:"), objc.Ptr(datep), tip, objc.Ptr(date))
	return rv
}

// Indicates whether a given date falls within a weekend period, as defined by the calendar and the calendar's locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1412175-isdateinweekend?language=objc
func (c_ Calendar) IsDateInWeekend(date IDate) bool {
	rv := objc.Call[bool](c_, objc.Sel("isDateInWeekend:"), objc.Ptr(date))
	return rv
}

// Returns the date components representing a given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1414841-components?language=objc
func (c_ Calendar) ComponentsFromDate(unitFlags CalendarUnit, date IDate) DateComponents {
	rv := objc.Call[DateComponents](c_, objc.Sel("components:fromDate:"), unitFlags, objc.Ptr(date))
	return rv
}

// Returns the first moment of a given date as a date instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1417161-startofdayfordate?language=objc
func (c_ Calendar) StartOfDayForDate(date IDate) Date {
	rv := objc.Call[Date](c_, objc.Sel("startOfDayForDate:"), objc.Ptr(date))
	return rv
}

// Returns by reference the era, year, week of year, and weekday component values for a given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1418143-getera?language=objc
func (c_ Calendar) GetEraYearMonthDayFromDate(eraValuePointer *int, yearValuePointer *int, monthValuePointer *int, dayValuePointer *int, date IDate) {
	objc.Call[objc.Void](c_, objc.Sel("getEra:year:month:day:fromDate:"), eraValuePointer, yearValuePointer, monthValuePointer, dayValuePointer, objc.Ptr(date))
}

// Computes the dates that match (or most closely match) a given set of components, and calls the block once for each of them, until the enumeration is stopped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1413938-enumeratedatesstartingafterdate?language=objc
func (c_ Calendar) EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock(start IDate, comps IDateComponents, opts CalendarOptions, block func(date Date, exactMatch bool, stop *bool)) {
	objc.Call[objc.Void](c_, objc.Sel("enumerateDatesStartingAfterDate:matchingComponents:options:usingBlock:"), objc.Ptr(start), objc.Ptr(comps), opts, block)
}

// Initializes a calendar according to a given identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1415991-initwithcalendaridentifier?language=objc
func (c_ Calendar) InitWithCalendarIdentifier(ident CalendarIdentifier) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("initWithCalendarIdentifier:"), ident)
	return rv
}

// Returns by reference the starting date and time interval range of the next weekend period after a given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1409905-nextweekendstartdate?language=objc
func (c_ Calendar) NextWeekendStartDateIntervalOptionsAfterDate(datep IDate, tip *TimeInterval, options CalendarOptions, date IDate) bool {
	rv := objc.Call[bool](c_, objc.Sel("nextWeekendStartDate:interval:options:afterDate:"), objc.Ptr(datep), tip, options, objc.Ptr(date))
	return rv
}

// Returns the minimum range limits of the values that a given unit can take on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1413638-minimumrangeofunit?language=objc
func (c_ Calendar) MinimumRangeOfUnit(unit CalendarUnit) Range {
	rv := objc.Call[Range](c_, objc.Sel("minimumRangeOfUnit:"), unit)
	return rv
}

// Indicates whether two dates are in the same day. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1417649-isdate?language=objc
func (c_ Calendar) IsDateInSameDayAsDate(date1 IDate, date2 IDate) bool {
	rv := objc.Call[bool](c_, objc.Sel("isDate:inSameDayAsDate:"), objc.Ptr(date1), objc.Ptr(date2))
	return rv
}

// Returns a date created with the given components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1415254-datewithera?language=objc
func (c_ Calendar) DateWithEraYearMonthDayHourMinuteSecondNanosecond(eraValue int, yearValue int, monthValue int, dayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date {
	rv := objc.Call[Date](c_, objc.Sel("dateWithEra:year:month:day:hour:minute:second:nanosecond:"), eraValue, yearValue, monthValue, dayValue, hourValue, minuteValue, secondValue, nanosecondValue)
	return rv
}

// Indicates whether the given date is in “tomorrow.” [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1410279-isdateintomorrow?language=objc
func (c_ Calendar) IsDateInTomorrow(date IDate) bool {
	rv := objc.Call[bool](c_, objc.Sel("isDateInTomorrow:"), objc.Ptr(date))
	return rv
}

// Creates a new date calculated with the given time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1407363-datebysettinghour?language=objc
func (c_ Calendar) DateBySettingHourMinuteSecondOfDateOptions(h int, m int, s int, date IDate, opts CalendarOptions) Date {
	rv := objc.Call[Date](c_, objc.Sel("dateBySettingHour:minute:second:ofDate:options:"), h, m, s, objc.Ptr(date), opts)
	return rv
}

// Indicates whether the given date is in “yesterday.” [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1409356-isdateinyesterday?language=objc
func (c_ Calendar) IsDateInYesterday(date IDate) bool {
	rv := objc.Call[bool](c_, objc.Sel("isDateInYesterday:"), objc.Ptr(date))
	return rv
}

// Indicates the ordering of two given dates based on their components down to a given unit granularity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1415661-comparedate?language=objc
func (c_ Calendar) CompareDateToDateToUnitGranularity(date1 IDate, date2 IDate, unit CalendarUnit) ComparisonResult {
	rv := objc.Call[ComparisonResult](c_, objc.Sel("compareDate:toDate:toUnitGranularity:"), objc.Ptr(date1), objc.Ptr(date2), unit)
	return rv
}

// The index of the first weekday of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1408310-firstweekday?language=objc
func (c_ Calendar) FirstWeekday() uint {
	rv := objc.Call[uint](c_, objc.Sel("firstWeekday"))
	return rv
}

// The index of the first weekday of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1408310-firstweekday?language=objc
func (c_ Calendar) SetFirstWeekday(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setFirstWeekday:"), value)
}

// A list of very short month symbols for this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1408035-veryshortstandalonemonthsymbols?language=objc
func (c_ Calendar) VeryShortStandaloneMonthSymbols() []string {
	rv := objc.Call[[]string](c_, objc.Sel("veryShortStandaloneMonthSymbols"))
	return rv
}

// A list of short standalone weekday symbols for this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1413871-shortstandaloneweekdaysymbols?language=objc
func (c_ Calendar) ShortStandaloneWeekdaySymbols() []string {
	rv := objc.Call[[]string](c_, objc.Sel("shortStandaloneWeekdaySymbols"))
	return rv
}

// A list of weekdays in this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1412939-weekdaysymbols?language=objc
func (c_ Calendar) WeekdaySymbols() []string {
	rv := objc.Call[[]string](c_, objc.Sel("weekdaySymbols"))
	return rv
}

// The symbol used to represent “PM” for this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1416343-pmsymbol?language=objc
func (c_ Calendar) PMSymbol() string {
	rv := objc.Call[string](c_, objc.Sel("PMSymbol"))
	return rv
}

// A list of standalone quarter symbols for this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1407159-standalonequartersymbols?language=objc
func (c_ Calendar) StandaloneQuarterSymbols() []string {
	rv := objc.Call[[]string](c_, objc.Sel("standaloneQuarterSymbols"))
	return rv
}

// A list of month symbols for this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1414872-monthsymbols?language=objc
func (c_ Calendar) MonthSymbols() []string {
	rv := objc.Call[[]string](c_, objc.Sel("monthSymbols"))
	return rv
}

// A list of shorter-named weekdays in this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1407268-shortweekdaysymbols?language=objc
func (c_ Calendar) ShortWeekdaySymbols() []string {
	rv := objc.Call[[]string](c_, objc.Sel("shortWeekdaySymbols"))
	return rv
}

// The minimum number of days in the first week of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1410186-minimumdaysinfirstweek?language=objc
func (c_ Calendar) MinimumDaysInFirstWeek() uint {
	rv := objc.Call[uint](c_, objc.Sel("minimumDaysInFirstWeek"))
	return rv
}

// The minimum number of days in the first week of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1410186-minimumdaysinfirstweek?language=objc
func (c_ Calendar) SetMinimumDaysInFirstWeek(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setMinimumDaysInFirstWeek:"), value)
}

// The user’s current calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1408501-currentcalendar?language=objc
func (cc _CalendarClass) CurrentCalendar() Calendar {
	rv := objc.Call[Calendar](cc, objc.Sel("currentCalendar"))
	return rv
}

// The user’s current calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1408501-currentcalendar?language=objc
func Calendar_CurrentCalendar() Calendar {
	return CalendarClass.CurrentCalendar()
}

// An identifier for the calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1408268-calendaridentifier?language=objc
func (c_ Calendar) CalendarIdentifier() CalendarIdentifier {
	rv := objc.Call[CalendarIdentifier](c_, objc.Sel("calendarIdentifier"))
	return rv
}

// The locale of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1418111-locale?language=objc
func (c_ Calendar) Locale() Locale {
	rv := objc.Call[Locale](c_, objc.Sel("locale"))
	return rv
}

// The locale of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1418111-locale?language=objc
func (c_ Calendar) SetLocale(value ILocale) {
	objc.Call[objc.Void](c_, objc.Sel("setLocale:"), objc.Ptr(value))
}

// A list of era symbols for this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1415038-erasymbols?language=objc
func (c_ Calendar) EraSymbols() []string {
	rv := objc.Call[[]string](c_, objc.Sel("eraSymbols"))
	return rv
}

// A list of standalone weekday symbols for this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1411219-standaloneweekdaysymbols?language=objc
func (c_ Calendar) StandaloneWeekdaySymbols() []string {
	rv := objc.Call[[]string](c_, objc.Sel("standaloneWeekdaySymbols"))
	return rv
}

// The time zone for the calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1409969-timezone?language=objc
func (c_ Calendar) TimeZone() TimeZone {
	rv := objc.Call[TimeZone](c_, objc.Sel("timeZone"))
	return rv
}

// The time zone for the calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1409969-timezone?language=objc
func (c_ Calendar) SetTimeZone(value ITimeZone) {
	objc.Call[objc.Void](c_, objc.Sel("setTimeZone:"), objc.Ptr(value))
}

// A list of very short month symbols for this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1412779-veryshortmonthsymbols?language=objc
func (c_ Calendar) VeryShortMonthSymbols() []string {
	rv := objc.Call[[]string](c_, objc.Sel("veryShortMonthSymbols"))
	return rv
}

// A calendar that tracks changes to user’s preferred calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1413771-autoupdatingcurrentcalendar?language=objc
func (cc _CalendarClass) AutoupdatingCurrentCalendar() Calendar {
	rv := objc.Call[Calendar](cc, objc.Sel("autoupdatingCurrentCalendar"))
	return rv
}

// A calendar that tracks changes to user’s preferred calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1413771-autoupdatingcurrentcalendar?language=objc
func Calendar_AutoupdatingCurrentCalendar() Calendar {
	return CalendarClass.AutoupdatingCurrentCalendar()
}

// A list of quarter symbols for this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1411517-quartersymbols?language=objc
func (c_ Calendar) QuarterSymbols() []string {
	rv := objc.Call[[]string](c_, objc.Sel("quarterSymbols"))
	return rv
}

// A list of short standalone quarter symbols for this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1409823-shortstandalonequartersymbols?language=objc
func (c_ Calendar) ShortStandaloneQuarterSymbols() []string {
	rv := objc.Call[[]string](c_, objc.Sel("shortStandaloneQuarterSymbols"))
	return rv
}

// The symbol used to represent “AM” for this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1416226-amsymbol?language=objc
func (c_ Calendar) AMSymbol() string {
	rv := objc.Call[string](c_, objc.Sel("AMSymbol"))
	return rv
}

// A list of very short standalone weekday symbols for this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1418273-veryshortstandaloneweekdaysymbol?language=objc
func (c_ Calendar) VeryShortStandaloneWeekdaySymbols() []string {
	rv := objc.Call[[]string](c_, objc.Sel("veryShortStandaloneWeekdaySymbols"))
	return rv
}

// A list of short standalone month symbols for this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1418180-shortstandalonemonthsymbols?language=objc
func (c_ Calendar) ShortStandaloneMonthSymbols() []string {
	rv := objc.Call[[]string](c_, objc.Sel("shortStandaloneMonthSymbols"))
	return rv
}

// A list of long era symbols for this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1414285-longerasymbols?language=objc
func (c_ Calendar) LongEraSymbols() []string {
	rv := objc.Call[[]string](c_, objc.Sel("longEraSymbols"))
	return rv
}

// A list of short month symbols for this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1408952-shortmonthsymbols?language=objc
func (c_ Calendar) ShortMonthSymbols() []string {
	rv := objc.Call[[]string](c_, objc.Sel("shortMonthSymbols"))
	return rv
}

// A list of very-shortly-named weekdays in this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1417207-veryshortweekdaysymbols?language=objc
func (c_ Calendar) VeryShortWeekdaySymbols() []string {
	rv := objc.Call[[]string](c_, objc.Sel("veryShortWeekdaySymbols"))
	return rv
}

// A list of short quarter symbols for this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1414864-shortquartersymbols?language=objc
func (c_ Calendar) ShortQuarterSymbols() []string {
	rv := objc.Call[[]string](c_, objc.Sel("shortQuarterSymbols"))
	return rv
}

// A list of standalone month symbols for this calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/1409598-standalonemonthsymbols?language=objc
func (c_ Calendar) StandaloneMonthSymbols() []string {
	rv := objc.Call[[]string](c_, objc.Sel("standaloneMonthSymbols"))
	return rv
}
