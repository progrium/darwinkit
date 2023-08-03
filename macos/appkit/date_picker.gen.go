// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var DatePickerClass = _DatePickerClass{objc.GetClass("NSDatePicker")}

type _DatePickerClass struct {
	objc.Class
}

type IDatePicker interface {
	IControl
	IsBezeled() bool
	SetBezeled(value bool)
	IsBordered() bool
	SetBordered(value bool)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	TextColor() Color
	SetTextColor(value IColor)
	DatePickerStyle() DatePickerStyle
	SetDatePickerStyle(value DatePickerStyle)
	PresentsCalendarOverlay() bool
	SetPresentsCalendarOverlay(value bool)
	Delegate() DatePickerCellDelegateWrapper
	SetDelegate(value IDatePickerCellDelegate)
	SetDelegate0(value objc.IObject)
	DatePickerElements() DatePickerElementFlags
	SetDatePickerElements(value DatePickerElementFlags)
	Calendar() foundation.Calendar
	SetCalendar(value foundation.ICalendar)
	Locale() foundation.Locale
	SetLocale(value foundation.ILocale)
	DatePickerMode() DatePickerMode
	SetDatePickerMode(value DatePickerMode)
	TimeZone() foundation.TimeZone
	SetTimeZone(value foundation.ITimeZone)
	DateValue() foundation.Date
	SetDateValue(value foundation.IDate)
	TimeInterval() foundation.TimeInterval
	SetTimeInterval(value foundation.TimeInterval)
	MinDate() foundation.Date
	SetMinDate(value foundation.IDate)
	MaxDate() foundation.Date
	SetMaxDate(value foundation.IDate)
}

type DatePicker struct {
	Control
}

func MakeDatePicker(ptr unsafe.Pointer) DatePicker {
	return DatePicker{
		Control: MakeControl(ptr),
	}
}

func (d_ DatePicker) InitWithFrame(frameRect foundation.Rect) DatePicker {
	rv := objc.CallMethod[DatePicker](d_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func DatePicker_InitWithFrame(frameRect foundation.Rect) DatePicker {
	return DatePickerClass.Alloc().InitWithFrame(frameRect)
}

func (d_ DatePicker) Init() DatePicker {
	rv := objc.CallMethod[DatePicker](d_, objc.GetSelector("init"))
	return rv
}

func DatePicker_Init() DatePicker {
	return DatePickerClass.Alloc().Init()
}

func (dc _DatePickerClass) Alloc() DatePicker {
	rv := objc.CallMethod[DatePicker](dc, objc.GetSelector("alloc"))
	return rv
}

func DatePicker_Alloc() DatePicker {
	return DatePickerClass.Alloc()
}

func (dc _DatePickerClass) New() DatePicker {
	rv := objc.CallMethod[DatePicker](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDatePicker() DatePicker {
	return DatePickerClass.New()
}

func DatePicker_New() DatePicker {
	return DatePickerClass.New()
}

func (d_ DatePicker) IsBezeled() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isBezeled"))
	return rv
}

func (d_ DatePicker) SetBezeled(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setBezeled:"), value)
}

func (d_ DatePicker) IsBordered() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isBordered"))
	return rv
}

func (d_ DatePicker) SetBordered(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setBordered:"), value)
}

func (d_ DatePicker) BackgroundColor() Color {
	rv := objc.CallMethod[Color](d_, objc.GetSelector("backgroundColor"))
	return rv
}

func (d_ DatePicker) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (d_ DatePicker) DrawsBackground() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("drawsBackground"))
	return rv
}

func (d_ DatePicker) SetDrawsBackground(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDrawsBackground:"), value)
}

func (d_ DatePicker) TextColor() Color {
	rv := objc.CallMethod[Color](d_, objc.GetSelector("textColor"))
	return rv
}

func (d_ DatePicker) SetTextColor(value IColor) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setTextColor:"), objc.ExtractPtr(value))
}

func (d_ DatePicker) DatePickerStyle() DatePickerStyle {
	rv := objc.CallMethod[DatePickerStyle](d_, objc.GetSelector("datePickerStyle"))
	return rv
}

func (d_ DatePicker) SetDatePickerStyle(value DatePickerStyle) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDatePickerStyle:"), value)
}

func (d_ DatePicker) PresentsCalendarOverlay() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("presentsCalendarOverlay"))
	return rv
}

func (d_ DatePicker) SetPresentsCalendarOverlay(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setPresentsCalendarOverlay:"), value)
}

func (d_ DatePicker) Delegate() DatePickerCellDelegateWrapper {
	rv := objc.CallMethod[DatePickerCellDelegateWrapper](d_, objc.GetSelector("delegate"))
	return rv
}

func (d_ DatePicker) SetDelegate(value IDatePickerCellDelegate) {
	po := objc.WrapAsProtocol("NSDatePickerCellDelegate", value)
	objc.SetAssociatedObject(d_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDelegate:"), po)
}

func (d_ DatePicker) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (d_ DatePicker) DatePickerElements() DatePickerElementFlags {
	rv := objc.CallMethod[DatePickerElementFlags](d_, objc.GetSelector("datePickerElements"))
	return rv
}

func (d_ DatePicker) SetDatePickerElements(value DatePickerElementFlags) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDatePickerElements:"), value)
}

func (d_ DatePicker) Calendar() foundation.Calendar {
	rv := objc.CallMethod[foundation.Calendar](d_, objc.GetSelector("calendar"))
	return rv
}

func (d_ DatePicker) SetCalendar(value foundation.ICalendar) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setCalendar:"), objc.ExtractPtr(value))
}

func (d_ DatePicker) Locale() foundation.Locale {
	rv := objc.CallMethod[foundation.Locale](d_, objc.GetSelector("locale"))
	return rv
}

func (d_ DatePicker) SetLocale(value foundation.ILocale) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setLocale:"), objc.ExtractPtr(value))
}

func (d_ DatePicker) DatePickerMode() DatePickerMode {
	rv := objc.CallMethod[DatePickerMode](d_, objc.GetSelector("datePickerMode"))
	return rv
}

func (d_ DatePicker) SetDatePickerMode(value DatePickerMode) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDatePickerMode:"), value)
}

func (d_ DatePicker) TimeZone() foundation.TimeZone {
	rv := objc.CallMethod[foundation.TimeZone](d_, objc.GetSelector("timeZone"))
	return rv
}

func (d_ DatePicker) SetTimeZone(value foundation.ITimeZone) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setTimeZone:"), objc.ExtractPtr(value))
}

func (d_ DatePicker) DateValue() foundation.Date {
	rv := objc.CallMethod[foundation.Date](d_, objc.GetSelector("dateValue"))
	return rv
}

func (d_ DatePicker) SetDateValue(value foundation.IDate) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDateValue:"), objc.ExtractPtr(value))
}

func (d_ DatePicker) TimeInterval() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](d_, objc.GetSelector("timeInterval"))
	return rv
}

func (d_ DatePicker) SetTimeInterval(value foundation.TimeInterval) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setTimeInterval:"), value)
}

func (d_ DatePicker) MinDate() foundation.Date {
	rv := objc.CallMethod[foundation.Date](d_, objc.GetSelector("minDate"))
	return rv
}

func (d_ DatePicker) SetMinDate(value foundation.IDate) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setMinDate:"), objc.ExtractPtr(value))
}

func (d_ DatePicker) MaxDate() foundation.Date {
	rv := objc.CallMethod[foundation.Date](d_, objc.GetSelector("maxDate"))
	return rv
}

func (d_ DatePicker) SetMaxDate(value foundation.IDate) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setMaxDate:"), objc.ExtractPtr(value))
}
