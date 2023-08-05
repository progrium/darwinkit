// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var CalendarClass = _CalendarClass{objc.GetClass("NSCalendar")}

type _CalendarClass struct {
	objc.Class
}

type ICalendar interface {
	objc.IObject
	InitWithCalendarIdentifier(ident CalendarIdentifier) objc.Object
	DateMatchesComponents(date IDate, components IDateComponents) bool
	ComponentFromDate(unit CalendarUnit, date IDate) int
	ComponentsFromDate(unitFlags CalendarUnit, date IDate) DateComponents
	ComponentsFromDateToDateOptions(unitFlags CalendarUnit, startingDate IDate, resultDate IDate, opts CalendarOptions) DateComponents
	ComponentsFromDateComponentsToDateComponentsOptions(unitFlags CalendarUnit, startingDateComp IDateComponents, resultDateComp IDateComponents, options CalendarOptions) DateComponents
	ComponentsInTimeZoneFromDate(timezone ITimeZone, date IDate) DateComponents
	MaximumRangeOfUnit(unit CalendarUnit) Range
	MinimumRangeOfUnit(unit CalendarUnit) Range
	OrdinalityOfUnitInUnitForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) uint
	RangeOfUnitInUnitForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) Range
	RangeOfUnitStartDateIntervalForDate(unit CalendarUnit, datep *Date, tip *TimeInterval, date IDate) bool
	StartOfDayForDate(date IDate) Date
	EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock(start IDate, comps IDateComponents, opts CalendarOptions, block func(date Date, exactMatch bool, stop *bool))
	NextDateAfterDateMatchingComponentsOptions(date IDate, comps IDateComponents, options CalendarOptions) Date
	NextDateAfterDateMatchingHourMinuteSecondOptions(date IDate, hourValue int, minuteValue int, secondValue int, options CalendarOptions) Date
	NextDateAfterDateMatchingUnitValueOptions(date IDate, unit CalendarUnit, value int, options CalendarOptions) Date
	DateFromComponents(comps IDateComponents) Date
	DateByAddingComponentsToDateOptions(comps IDateComponents, date IDate, opts CalendarOptions) Date
	DateByAddingUnitValueToDateOptions(unit CalendarUnit, value int, date IDate, options CalendarOptions) Date
	DateBySettingHourMinuteSecondOfDateOptions(h int, m int, s int, date IDate, opts CalendarOptions) Date
	DateBySettingUnitValueOfDateOptions(unit CalendarUnit, v int, date IDate, opts CalendarOptions) Date
	DateWithEraYearMonthDayHourMinuteSecondNanosecond(eraValue int, yearValue int, monthValue int, dayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date
	DateWithEraYearForWeekOfYearWeekOfYearWeekdayHourMinuteSecondNanosecond(eraValue int, yearValue int, weekValue int, weekdayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date
	CompareDateToDateToUnitGranularity(date1 IDate, date2 IDate, unit CalendarUnit) ComparisonResult
	IsDateEqualToDateToUnitGranularity(date1 IDate, date2 IDate, unit CalendarUnit) bool
	IsDateInSameDayAsDate(date1 IDate, date2 IDate) bool
	IsDateInToday(date IDate) bool
	IsDateInTomorrow(date IDate) bool
	IsDateInWeekend(date IDate) bool
	IsDateInYesterday(date IDate) bool
	CalendarIdentifier() CalendarIdentifier
	FirstWeekday() uint
	SetFirstWeekday(value uint)
	Locale() Locale
	SetLocale(value ILocale)
	TimeZone() TimeZone
	SetTimeZone(value ITimeZone)
	MinimumDaysInFirstWeek() uint
	SetMinimumDaysInFirstWeek(value uint)
	AMSymbol() string
	PMSymbol() string
	WeekdaySymbols() []string
	ShortWeekdaySymbols() []string
	VeryShortWeekdaySymbols() []string
	StandaloneWeekdaySymbols() []string
	ShortStandaloneWeekdaySymbols() []string
	VeryShortStandaloneWeekdaySymbols() []string
	MonthSymbols() []string
	ShortMonthSymbols() []string
	VeryShortMonthSymbols() []string
	StandaloneMonthSymbols() []string
	ShortStandaloneMonthSymbols() []string
	VeryShortStandaloneMonthSymbols() []string
	QuarterSymbols() []string
	ShortQuarterSymbols() []string
	StandaloneQuarterSymbols() []string
	ShortStandaloneQuarterSymbols() []string
	EraSymbols() []string
	LongEraSymbols() []string
}

type Calendar struct {
	objc.Object
}

func MakeCalendar(ptr unsafe.Pointer) Calendar {
	return Calendar{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CalendarClass) Alloc() Calendar {
	rv := objc.CallMethod[Calendar](cc, objc.GetSelector("alloc"))
	return rv
}

func Calendar_Alloc() Calendar {
	return CalendarClass.Alloc()
}

func (cc _CalendarClass) New() Calendar {
	rv := objc.CallMethod[Calendar](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCalendar() Calendar {
	return CalendarClass.New()
}

func Calendar_New() Calendar {
	return CalendarClass.New()
}

func (c_ Calendar) Init() Calendar {
	rv := objc.CallMethod[Calendar](c_, objc.GetSelector("init"))
	return rv
}

func Calendar_Init() Calendar {
	return CalendarClass.Alloc().Init()
}

func (cc _CalendarClass) CalendarWithIdentifier(calendarIdentifierConstant CalendarIdentifier) Calendar {
	rv := objc.CallMethod[Calendar](cc, objc.GetSelector("calendarWithIdentifier:"), calendarIdentifierConstant)
	return rv
}

func Calendar_CalendarWithIdentifier(calendarIdentifierConstant CalendarIdentifier) Calendar {
	return CalendarClass.CalendarWithIdentifier(calendarIdentifierConstant)
}

func (c_ Calendar) InitWithCalendarIdentifier(ident CalendarIdentifier) objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.GetSelector("initWithCalendarIdentifier:"), ident)
	return rv
}

func (c_ Calendar) DateMatchesComponents(date IDate, components IDateComponents) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("date:matchesComponents:"), objc.ExtractPtr(date), objc.ExtractPtr(components))
	return rv
}

func (c_ Calendar) ComponentFromDate(unit CalendarUnit, date IDate) int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("component:fromDate:"), unit, objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) ComponentsFromDate(unitFlags CalendarUnit, date IDate) DateComponents {
	rv := objc.CallMethod[DateComponents](c_, objc.GetSelector("components:fromDate:"), unitFlags, objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) ComponentsFromDateToDateOptions(unitFlags CalendarUnit, startingDate IDate, resultDate IDate, opts CalendarOptions) DateComponents {
	rv := objc.CallMethod[DateComponents](c_, objc.GetSelector("components:fromDate:toDate:options:"), unitFlags, objc.ExtractPtr(startingDate), objc.ExtractPtr(resultDate), opts)
	return rv
}

func (c_ Calendar) ComponentsFromDateComponentsToDateComponentsOptions(unitFlags CalendarUnit, startingDateComp IDateComponents, resultDateComp IDateComponents, options CalendarOptions) DateComponents {
	rv := objc.CallMethod[DateComponents](c_, objc.GetSelector("components:fromDateComponents:toDateComponents:options:"), unitFlags, objc.ExtractPtr(startingDateComp), objc.ExtractPtr(resultDateComp), options)
	return rv
}

func (c_ Calendar) ComponentsInTimeZoneFromDate(timezone ITimeZone, date IDate) DateComponents {
	rv := objc.CallMethod[DateComponents](c_, objc.GetSelector("componentsInTimeZone:fromDate:"), objc.ExtractPtr(timezone), objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) MaximumRangeOfUnit(unit CalendarUnit) Range {
	rv := objc.CallMethod[Range](c_, objc.GetSelector("maximumRangeOfUnit:"), unit)
	return rv
}

func (c_ Calendar) MinimumRangeOfUnit(unit CalendarUnit) Range {
	rv := objc.CallMethod[Range](c_, objc.GetSelector("minimumRangeOfUnit:"), unit)
	return rv
}

func (c_ Calendar) OrdinalityOfUnitInUnitForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) uint {
	rv := objc.CallMethod[uint](c_, objc.GetSelector("ordinalityOfUnit:inUnit:forDate:"), smaller, larger, objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) RangeOfUnitInUnitForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) Range {
	rv := objc.CallMethod[Range](c_, objc.GetSelector("rangeOfUnit:inUnit:forDate:"), smaller, larger, objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) RangeOfUnitStartDateIntervalForDate(unit CalendarUnit, datep *Date, tip *TimeInterval, date IDate) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("rangeOfUnit:startDate:interval:forDate:"), unit, unsafe.Pointer(datep), tip, objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) StartOfDayForDate(date IDate) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("startOfDayForDate:"), objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock(start IDate, comps IDateComponents, opts CalendarOptions, block func(date Date, exactMatch bool, stop *bool)) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("enumerateDatesStartingAfterDate:matchingComponents:options:usingBlock:"), objc.ExtractPtr(start), objc.ExtractPtr(comps), opts, block)
}

func (c_ Calendar) NextDateAfterDateMatchingComponentsOptions(date IDate, comps IDateComponents, options CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("nextDateAfterDate:matchingComponents:options:"), objc.ExtractPtr(date), objc.ExtractPtr(comps), options)
	return rv
}

func (c_ Calendar) NextDateAfterDateMatchingHourMinuteSecondOptions(date IDate, hourValue int, minuteValue int, secondValue int, options CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("nextDateAfterDate:matchingHour:minute:second:options:"), objc.ExtractPtr(date), hourValue, minuteValue, secondValue, options)
	return rv
}

func (c_ Calendar) NextDateAfterDateMatchingUnitValueOptions(date IDate, unit CalendarUnit, value int, options CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("nextDateAfterDate:matchingUnit:value:options:"), objc.ExtractPtr(date), unit, value, options)
	return rv
}

func (c_ Calendar) DateFromComponents(comps IDateComponents) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("dateFromComponents:"), objc.ExtractPtr(comps))
	return rv
}

func (c_ Calendar) DateByAddingComponentsToDateOptions(comps IDateComponents, date IDate, opts CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("dateByAddingComponents:toDate:options:"), objc.ExtractPtr(comps), objc.ExtractPtr(date), opts)
	return rv
}

func (c_ Calendar) DateByAddingUnitValueToDateOptions(unit CalendarUnit, value int, date IDate, options CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("dateByAddingUnit:value:toDate:options:"), unit, value, objc.ExtractPtr(date), options)
	return rv
}

func (c_ Calendar) DateBySettingHourMinuteSecondOfDateOptions(h int, m int, s int, date IDate, opts CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("dateBySettingHour:minute:second:ofDate:options:"), h, m, s, objc.ExtractPtr(date), opts)
	return rv
}

func (c_ Calendar) DateBySettingUnitValueOfDateOptions(unit CalendarUnit, v int, date IDate, opts CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("dateBySettingUnit:value:ofDate:options:"), unit, v, objc.ExtractPtr(date), opts)
	return rv
}

func (c_ Calendar) DateWithEraYearMonthDayHourMinuteSecondNanosecond(eraValue int, yearValue int, monthValue int, dayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("dateWithEra:year:month:day:hour:minute:second:nanosecond:"), eraValue, yearValue, monthValue, dayValue, hourValue, minuteValue, secondValue, nanosecondValue)
	return rv
}

func (c_ Calendar) DateWithEraYearForWeekOfYearWeekOfYearWeekdayHourMinuteSecondNanosecond(eraValue int, yearValue int, weekValue int, weekdayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("dateWithEra:yearForWeekOfYear:weekOfYear:weekday:hour:minute:second:nanosecond:"), eraValue, yearValue, weekValue, weekdayValue, hourValue, minuteValue, secondValue, nanosecondValue)
	return rv
}

func (c_ Calendar) CompareDateToDateToUnitGranularity(date1 IDate, date2 IDate, unit CalendarUnit) ComparisonResult {
	rv := objc.CallMethod[ComparisonResult](c_, objc.GetSelector("compareDate:toDate:toUnitGranularity:"), objc.ExtractPtr(date1), objc.ExtractPtr(date2), unit)
	return rv
}

func (c_ Calendar) IsDateEqualToDateToUnitGranularity(date1 IDate, date2 IDate, unit CalendarUnit) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isDate:equalToDate:toUnitGranularity:"), objc.ExtractPtr(date1), objc.ExtractPtr(date2), unit)
	return rv
}

func (c_ Calendar) IsDateInSameDayAsDate(date1 IDate, date2 IDate) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isDate:inSameDayAsDate:"), objc.ExtractPtr(date1), objc.ExtractPtr(date2))
	return rv
}

func (c_ Calendar) IsDateInToday(date IDate) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isDateInToday:"), objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) IsDateInTomorrow(date IDate) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isDateInTomorrow:"), objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) IsDateInWeekend(date IDate) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isDateInWeekend:"), objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) IsDateInYesterday(date IDate) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isDateInYesterday:"), objc.ExtractPtr(date))
	return rv
}

func (cc _CalendarClass) CurrentCalendar() Calendar {
	rv := objc.CallMethod[Calendar](cc, objc.GetSelector("currentCalendar"))
	return rv
}

func Calendar_CurrentCalendar() Calendar {
	return CalendarClass.CurrentCalendar()
}

func (cc _CalendarClass) AutoupdatingCurrentCalendar() Calendar {
	rv := objc.CallMethod[Calendar](cc, objc.GetSelector("autoupdatingCurrentCalendar"))
	return rv
}

func Calendar_AutoupdatingCurrentCalendar() Calendar {
	return CalendarClass.AutoupdatingCurrentCalendar()
}

func (c_ Calendar) CalendarIdentifier() CalendarIdentifier {
	rv := objc.CallMethod[CalendarIdentifier](c_, objc.GetSelector("calendarIdentifier"))
	return rv
}

func (c_ Calendar) FirstWeekday() uint {
	rv := objc.CallMethod[uint](c_, objc.GetSelector("firstWeekday"))
	return rv
}

func (c_ Calendar) SetFirstWeekday(value uint) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setFirstWeekday:"), value)
}

func (c_ Calendar) Locale() Locale {
	rv := objc.CallMethod[Locale](c_, objc.GetSelector("locale"))
	return rv
}

func (c_ Calendar) SetLocale(value ILocale) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setLocale:"), objc.ExtractPtr(value))
}

func (c_ Calendar) TimeZone() TimeZone {
	rv := objc.CallMethod[TimeZone](c_, objc.GetSelector("timeZone"))
	return rv
}

func (c_ Calendar) SetTimeZone(value ITimeZone) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setTimeZone:"), objc.ExtractPtr(value))
}

func (c_ Calendar) MinimumDaysInFirstWeek() uint {
	rv := objc.CallMethod[uint](c_, objc.GetSelector("minimumDaysInFirstWeek"))
	return rv
}

func (c_ Calendar) SetMinimumDaysInFirstWeek(value uint) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMinimumDaysInFirstWeek:"), value)
}

func (c_ Calendar) AMSymbol() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("AMSymbol"))
	return rv
}

func (c_ Calendar) PMSymbol() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("PMSymbol"))
	return rv
}

func (c_ Calendar) WeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("weekdaySymbols"))
	return rv
}

func (c_ Calendar) ShortWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("shortWeekdaySymbols"))
	return rv
}

func (c_ Calendar) VeryShortWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("veryShortWeekdaySymbols"))
	return rv
}

func (c_ Calendar) StandaloneWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("standaloneWeekdaySymbols"))
	return rv
}

func (c_ Calendar) ShortStandaloneWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("shortStandaloneWeekdaySymbols"))
	return rv
}

func (c_ Calendar) VeryShortStandaloneWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("veryShortStandaloneWeekdaySymbols"))
	return rv
}

func (c_ Calendar) MonthSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("monthSymbols"))
	return rv
}

func (c_ Calendar) ShortMonthSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("shortMonthSymbols"))
	return rv
}

func (c_ Calendar) VeryShortMonthSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("veryShortMonthSymbols"))
	return rv
}

func (c_ Calendar) StandaloneMonthSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("standaloneMonthSymbols"))
	return rv
}

func (c_ Calendar) ShortStandaloneMonthSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("shortStandaloneMonthSymbols"))
	return rv
}

func (c_ Calendar) VeryShortStandaloneMonthSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("veryShortStandaloneMonthSymbols"))
	return rv
}

func (c_ Calendar) QuarterSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("quarterSymbols"))
	return rv
}

func (c_ Calendar) ShortQuarterSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("shortQuarterSymbols"))
	return rv
}

func (c_ Calendar) StandaloneQuarterSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("standaloneQuarterSymbols"))
	return rv
}

func (c_ Calendar) ShortStandaloneQuarterSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("shortStandaloneQuarterSymbols"))
	return rv
}

func (c_ Calendar) EraSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("eraSymbols"))
	return rv
}

func (c_ Calendar) LongEraSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("longEraSymbols"))
	return rv
}
