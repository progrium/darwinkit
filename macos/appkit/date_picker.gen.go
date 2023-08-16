// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DatePicker] class.
var DatePickerClass = _DatePickerClass{objc.GetClass("NSDatePicker")}

type _DatePickerClass struct {
	objc.Class
}

// An interface definition for the [DatePicker] class.
type IDatePicker interface {
	IControl
	IsBordered() bool
	SetBordered(value bool)
	DatePickerElements() DatePickerElementFlags
	SetDatePickerElements(value DatePickerElementFlags)
	TimeInterval() foundation.TimeInterval
	SetTimeInterval(value foundation.TimeInterval)
	Locale() foundation.Locale
	SetLocale(value foundation.ILocale)
	Delegate() DatePickerCellDelegateWrapper
	SetDelegate(value PDatePickerCellDelegate)
	SetDelegateObject(valueObject objc.IObject)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	MinDate() foundation.Date
	SetMinDate(value foundation.IDate)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	TimeZone() foundation.TimeZone
	SetTimeZone(value foundation.ITimeZone)
	DateValue() foundation.Date
	SetDateValue(value foundation.IDate)
	PresentsCalendarOverlay() bool
	SetPresentsCalendarOverlay(value bool)
	TextColor() Color
	SetTextColor(value IColor)
	Calendar() foundation.Calendar
	SetCalendar(value foundation.ICalendar)
	DatePickerMode() DatePickerMode
	SetDatePickerMode(value DatePickerMode)
	MaxDate() foundation.Date
	SetMaxDate(value foundation.IDate)
	IsBezeled() bool
	SetBezeled(value bool)
	DatePickerStyle() DatePickerStyle
	SetDatePickerStyle(value DatePickerStyle)
}

// A display of a calendar date with controls for editing the date value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker?language=objc
type DatePicker struct {
	Control
}

func DatePickerFrom(ptr unsafe.Pointer) DatePicker {
	return DatePicker{
		Control: ControlFrom(ptr),
	}
}

func (dc _DatePickerClass) Alloc() DatePicker {
	rv := objc.Call[DatePicker](dc, objc.Sel("alloc"))
	return rv
}

func DatePicker_Alloc() DatePicker {
	return DatePickerClass.Alloc()
}

func (dc _DatePickerClass) New() DatePicker {
	rv := objc.Call[DatePicker](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDatePicker() DatePicker {
	return DatePickerClass.New()
}

func (d_ DatePicker) Init() DatePicker {
	rv := objc.Call[DatePicker](d_, objc.Sel("init"))
	return rv
}

func (d_ DatePicker) InitWithFrame(frameRect foundation.Rect) DatePicker {
	rv := objc.Call[DatePicker](d_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func DatePicker_InitWithFrame(frameRect foundation.Rect) DatePicker {
	return DatePickerClass.Alloc().InitWithFrame(frameRect)
}

// A Boolean value that indicates whether the date picker has a plain border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1534176-bordered?language=objc
func (d_ DatePicker) IsBordered() bool {
	rv := objc.Call[bool](d_, objc.Sel("isBordered"))
	return rv
}

// A Boolean value that indicates whether the date picker has a plain border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1534176-bordered?language=objc
func (d_ DatePicker) SetBordered(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setBordered:"), value)
}

// A bitmask that indicates which visual elements of the date picker are currently shown, and which won't be usable because they are hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1533480-datepickerelements?language=objc
func (d_ DatePicker) DatePickerElements() DatePickerElementFlags {
	rv := objc.Call[DatePickerElementFlags](d_, objc.Sel("datePickerElements"))
	return rv
}

// A bitmask that indicates which visual elements of the date picker are currently shown, and which won't be usable because they are hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1533480-datepickerelements?language=objc
func (d_ DatePicker) SetDatePickerElements(value DatePickerElementFlags) {
	objc.Call[objc.Void](d_, objc.Sel("setDatePickerElements:"), value)
}

// The time interval selected by the date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1526681-timeinterval?language=objc
func (d_ DatePicker) TimeInterval() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](d_, objc.Sel("timeInterval"))
	return rv
}

// The time interval selected by the date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1526681-timeinterval?language=objc
func (d_ DatePicker) SetTimeInterval(value foundation.TimeInterval) {
	objc.Call[objc.Void](d_, objc.Sel("setTimeInterval:"), value)
}

// The date picker’s locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1525940-locale?language=objc
func (d_ DatePicker) Locale() foundation.Locale {
	rv := objc.Call[foundation.Locale](d_, objc.Sel("locale"))
	return rv
}

// The date picker’s locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1525940-locale?language=objc
func (d_ DatePicker) SetLocale(value foundation.ILocale) {
	objc.Call[objc.Void](d_, objc.Sel("setLocale:"), objc.Ptr(value))
}

// A delegate for the date picker’s cell [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1533878-delegate?language=objc
func (d_ DatePicker) Delegate() DatePickerCellDelegateWrapper {
	rv := objc.Call[DatePickerCellDelegateWrapper](d_, objc.Sel("delegate"))
	return rv
}

// A delegate for the date picker’s cell [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1533878-delegate?language=objc
func (d_ DatePicker) SetDelegate(value PDatePickerCellDelegate) {
	po0 := objc.WrapAsProtocol("NSDatePickerCellDelegate", value)
	objc.SetAssociatedObject(d_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](d_, objc.Sel("setDelegate:"), po0)
}

// A delegate for the date picker’s cell [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1533878-delegate?language=objc
func (d_ DatePicker) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](d_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The date picker’s background color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1527710-backgroundcolor?language=objc
func (d_ DatePicker) BackgroundColor() Color {
	rv := objc.Call[Color](d_, objc.Sel("backgroundColor"))
	return rv
}

// The date picker’s background color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1527710-backgroundcolor?language=objc
func (d_ DatePicker) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](d_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// The date picker’s minimum date value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1526893-mindate?language=objc
func (d_ DatePicker) MinDate() foundation.Date {
	rv := objc.Call[foundation.Date](d_, objc.Sel("minDate"))
	return rv
}

// The date picker’s minimum date value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1526893-mindate?language=objc
func (d_ DatePicker) SetMinDate(value foundation.IDate) {
	objc.Call[objc.Void](d_, objc.Sel("setMinDate:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the date picker draws the background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1528266-drawsbackground?language=objc
func (d_ DatePicker) DrawsBackground() bool {
	rv := objc.Call[bool](d_, objc.Sel("drawsBackground"))
	return rv
}

// A Boolean value that indicates whether the date picker draws the background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1528266-drawsbackground?language=objc
func (d_ DatePicker) SetDrawsBackground(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setDrawsBackground:"), value)
}

// The time zone for the date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1535451-timezone?language=objc
func (d_ DatePicker) TimeZone() foundation.TimeZone {
	rv := objc.Call[foundation.TimeZone](d_, objc.Sel("timeZone"))
	return rv
}

// The time zone for the date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1535451-timezone?language=objc
func (d_ DatePicker) SetTimeZone(value foundation.ITimeZone) {
	objc.Call[objc.Void](d_, objc.Sel("setTimeZone:"), objc.Ptr(value))
}

// The date selected by the date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1527606-datevalue?language=objc
func (d_ DatePicker) DateValue() foundation.Date {
	rv := objc.Call[foundation.Date](d_, objc.Sel("dateValue"))
	return rv
}

// The date selected by the date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1527606-datevalue?language=objc
func (d_ DatePicker) SetDateValue(value foundation.IDate) {
	objc.Call[objc.Void](d_, objc.Sel("setDateValue:"), objc.Ptr(value))
}

// A Boolean value that indicates whether to present a graphical calendar overlay when editing a calendar element within a text-field style date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/3521157-presentscalendaroverlay?language=objc
func (d_ DatePicker) PresentsCalendarOverlay() bool {
	rv := objc.Call[bool](d_, objc.Sel("presentsCalendarOverlay"))
	return rv
}

// A Boolean value that indicates whether to present a graphical calendar overlay when editing a calendar element within a text-field style date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/3521157-presentscalendaroverlay?language=objc
func (d_ DatePicker) SetPresentsCalendarOverlay(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setPresentsCalendarOverlay:"), value)
}

// The date picker’s text color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1534294-textcolor?language=objc
func (d_ DatePicker) TextColor() Color {
	rv := objc.Call[Color](d_, objc.Sel("textColor"))
	return rv
}

// The date picker’s text color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1534294-textcolor?language=objc
func (d_ DatePicker) SetTextColor(value IColor) {
	objc.Call[objc.Void](d_, objc.Sel("setTextColor:"), objc.Ptr(value))
}

// The calendar used by the date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1533591-calendar?language=objc
func (d_ DatePicker) Calendar() foundation.Calendar {
	rv := objc.Call[foundation.Calendar](d_, objc.Sel("calendar"))
	return rv
}

// The calendar used by the date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1533591-calendar?language=objc
func (d_ DatePicker) SetCalendar(value foundation.ICalendar) {
	objc.Call[objc.Void](d_, objc.Sel("setCalendar:"), objc.Ptr(value))
}

// The date picker’s mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1527214-datepickermode?language=objc
func (d_ DatePicker) DatePickerMode() DatePickerMode {
	rv := objc.Call[DatePickerMode](d_, objc.Sel("datePickerMode"))
	return rv
}

// The date picker’s mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1527214-datepickermode?language=objc
func (d_ DatePicker) SetDatePickerMode(value DatePickerMode) {
	objc.Call[objc.Void](d_, objc.Sel("setDatePickerMode:"), value)
}

// The date picker’s maximum date value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1535887-maxdate?language=objc
func (d_ DatePicker) MaxDate() foundation.Date {
	rv := objc.Call[foundation.Date](d_, objc.Sel("maxDate"))
	return rv
}

// The date picker’s maximum date value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1535887-maxdate?language=objc
func (d_ DatePicker) SetMaxDate(value foundation.IDate) {
	objc.Call[objc.Void](d_, objc.Sel("setMaxDate:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the date picker draws a bezeled border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1533534-bezeled?language=objc
func (d_ DatePicker) IsBezeled() bool {
	rv := objc.Call[bool](d_, objc.Sel("isBezeled"))
	return rv
}

// A Boolean value that indicates whether the date picker draws a bezeled border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1533534-bezeled?language=objc
func (d_ DatePicker) SetBezeled(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setBezeled:"), value)
}

// The date picker’s style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1528570-datepickerstyle?language=objc
func (d_ DatePicker) DatePickerStyle() DatePickerStyle {
	rv := objc.Call[DatePickerStyle](d_, objc.Sel("datePickerStyle"))
	return rv
}

// The date picker’s style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/1528570-datepickerstyle?language=objc
func (d_ DatePicker) SetDatePickerStyle(value DatePickerStyle) {
	objc.Call[objc.Void](d_, objc.Sel("setDatePickerStyle:"), value)
}
