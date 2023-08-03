// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var DateComponentsFormatterClass = _DateComponentsFormatterClass{objc.GetClass("NSDateComponentsFormatter")}

type _DateComponentsFormatterClass struct {
	objc.Class
}

type IDateComponentsFormatter interface {
	IFormatter
	StringFromDateComponents(components IDateComponents) string
	StringFromDateToDate(startDate IDate, endDate IDate) string
	StringFromTimeInterval(ti TimeInterval) string
	AllowedUnits() CalendarUnit
	SetAllowedUnits(value CalendarUnit)
	AllowsFractionalUnits() bool
	SetAllowsFractionalUnits(value bool)
	Calendar() Calendar
	SetCalendar(value ICalendar)
	CollapsesLargestUnit() bool
	SetCollapsesLargestUnit(value bool)
	IncludesApproximationPhrase() bool
	SetIncludesApproximationPhrase(value bool)
	IncludesTimeRemainingPhrase() bool
	SetIncludesTimeRemainingPhrase(value bool)
	MaximumUnitCount() int
	SetMaximumUnitCount(value int)
	UnitsStyle() DateComponentsFormatterUnitsStyle
	SetUnitsStyle(value DateComponentsFormatterUnitsStyle)
	ZeroFormattingBehavior() DateComponentsFormatterZeroFormattingBehavior
	SetZeroFormattingBehavior(value DateComponentsFormatterZeroFormattingBehavior)
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	ReferenceDate() Date
	SetReferenceDate(value IDate)
}

type DateComponentsFormatter struct {
	Formatter
}

func MakeDateComponentsFormatter(ptr unsafe.Pointer) DateComponentsFormatter {
	return DateComponentsFormatter{
		Formatter: MakeFormatter(ptr),
	}
}

func (dc _DateComponentsFormatterClass) Alloc() DateComponentsFormatter {
	rv := objc.CallMethod[DateComponentsFormatter](dc, objc.GetSelector("alloc"))
	return rv
}

func DateComponentsFormatter_Alloc() DateComponentsFormatter {
	return DateComponentsFormatterClass.Alloc()
}

func (dc _DateComponentsFormatterClass) New() DateComponentsFormatter {
	rv := objc.CallMethod[DateComponentsFormatter](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDateComponentsFormatter() DateComponentsFormatter {
	return DateComponentsFormatterClass.New()
}

func DateComponentsFormatter_New() DateComponentsFormatter {
	return DateComponentsFormatterClass.New()
}

func (d_ DateComponentsFormatter) Init() DateComponentsFormatter {
	rv := objc.CallMethod[DateComponentsFormatter](d_, objc.GetSelector("init"))
	return rv
}

func DateComponentsFormatter_Init() DateComponentsFormatter {
	return DateComponentsFormatterClass.Alloc().Init()
}

func (d_ DateComponentsFormatter) StringFromDateComponents(components IDateComponents) string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("stringFromDateComponents:"), objc.ExtractPtr(components))
	return rv
}

func (d_ DateComponentsFormatter) StringFromDateToDate(startDate IDate, endDate IDate) string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("stringFromDate:toDate:"), objc.ExtractPtr(startDate), objc.ExtractPtr(endDate))
	return rv
}

func (d_ DateComponentsFormatter) StringFromTimeInterval(ti TimeInterval) string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("stringFromTimeInterval:"), ti)
	return rv
}

func (dc _DateComponentsFormatterClass) LocalizedStringFromDateComponentsUnitsStyle(components IDateComponents, unitsStyle DateComponentsFormatterUnitsStyle) string {
	rv := objc.CallMethod[string](dc, objc.GetSelector("localizedStringFromDateComponents:unitsStyle:"), objc.ExtractPtr(components), unitsStyle)
	return rv
}

func DateComponentsFormatter_LocalizedStringFromDateComponentsUnitsStyle(components IDateComponents, unitsStyle DateComponentsFormatterUnitsStyle) string {
	return DateComponentsFormatterClass.LocalizedStringFromDateComponentsUnitsStyle(components, unitsStyle)
}

func (d_ DateComponentsFormatter) AllowedUnits() CalendarUnit {
	rv := objc.CallMethod[CalendarUnit](d_, objc.GetSelector("allowedUnits"))
	return rv
}

func (d_ DateComponentsFormatter) SetAllowedUnits(value CalendarUnit) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setAllowedUnits:"), value)
}

func (d_ DateComponentsFormatter) AllowsFractionalUnits() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("allowsFractionalUnits"))
	return rv
}

func (d_ DateComponentsFormatter) SetAllowsFractionalUnits(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setAllowsFractionalUnits:"), value)
}

func (d_ DateComponentsFormatter) Calendar() Calendar {
	rv := objc.CallMethod[Calendar](d_, objc.GetSelector("calendar"))
	return rv
}

func (d_ DateComponentsFormatter) SetCalendar(value ICalendar) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setCalendar:"), objc.ExtractPtr(value))
}

func (d_ DateComponentsFormatter) CollapsesLargestUnit() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("collapsesLargestUnit"))
	return rv
}

func (d_ DateComponentsFormatter) SetCollapsesLargestUnit(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setCollapsesLargestUnit:"), value)
}

func (d_ DateComponentsFormatter) IncludesApproximationPhrase() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("includesApproximationPhrase"))
	return rv
}

func (d_ DateComponentsFormatter) SetIncludesApproximationPhrase(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setIncludesApproximationPhrase:"), value)
}

func (d_ DateComponentsFormatter) IncludesTimeRemainingPhrase() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("includesTimeRemainingPhrase"))
	return rv
}

func (d_ DateComponentsFormatter) SetIncludesTimeRemainingPhrase(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setIncludesTimeRemainingPhrase:"), value)
}

func (d_ DateComponentsFormatter) MaximumUnitCount() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("maximumUnitCount"))
	return rv
}

func (d_ DateComponentsFormatter) SetMaximumUnitCount(value int) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setMaximumUnitCount:"), value)
}

func (d_ DateComponentsFormatter) UnitsStyle() DateComponentsFormatterUnitsStyle {
	rv := objc.CallMethod[DateComponentsFormatterUnitsStyle](d_, objc.GetSelector("unitsStyle"))
	return rv
}

func (d_ DateComponentsFormatter) SetUnitsStyle(value DateComponentsFormatterUnitsStyle) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setUnitsStyle:"), value)
}

func (d_ DateComponentsFormatter) ZeroFormattingBehavior() DateComponentsFormatterZeroFormattingBehavior {
	rv := objc.CallMethod[DateComponentsFormatterZeroFormattingBehavior](d_, objc.GetSelector("zeroFormattingBehavior"))
	return rv
}

func (d_ DateComponentsFormatter) SetZeroFormattingBehavior(value DateComponentsFormatterZeroFormattingBehavior) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setZeroFormattingBehavior:"), value)
}

func (d_ DateComponentsFormatter) FormattingContext() FormattingContext {
	rv := objc.CallMethod[FormattingContext](d_, objc.GetSelector("formattingContext"))
	return rv
}

func (d_ DateComponentsFormatter) SetFormattingContext(value FormattingContext) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setFormattingContext:"), value)
}

func (d_ DateComponentsFormatter) ReferenceDate() Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("referenceDate"))
	return rv
}

func (d_ DateComponentsFormatter) SetReferenceDate(value IDate) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setReferenceDate:"), objc.ExtractPtr(value))
}
