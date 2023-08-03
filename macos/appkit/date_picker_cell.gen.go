// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var DatePickerCellClass = _DatePickerCellClass{objc.GetClass("NSDatePickerCell")}

type _DatePickerCellClass struct {
	objc.Class
}

type IDatePickerCell interface {
	IActionCell
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	TextColor() Color
	SetTextColor(value IColor)
	DatePickerStyle() DatePickerStyle
	SetDatePickerStyle(value DatePickerStyle)
	DatePickerElements() DatePickerElementFlags
	SetDatePickerElements(value DatePickerElementFlags)
	DatePickerMode() DatePickerMode
	SetDatePickerMode(value DatePickerMode)
	DateValue() foundation.Date
	SetDateValue(value foundation.IDate)
	TimeInterval() foundation.TimeInterval
	SetTimeInterval(value foundation.TimeInterval)
	Calendar() foundation.Calendar
	SetCalendar(value foundation.ICalendar)
	Locale() foundation.Locale
	SetLocale(value foundation.ILocale)
	TimeZone() foundation.TimeZone
	SetTimeZone(value foundation.ITimeZone)
	MinDate() foundation.Date
	SetMinDate(value foundation.IDate)
	MaxDate() foundation.Date
	SetMaxDate(value foundation.IDate)
	Delegate() DatePickerCellDelegateWrapper
	SetDelegate(value IDatePickerCellDelegate)
	SetDelegate0(value objc.IObject)
}

type DatePickerCell struct {
	ActionCell
}

func MakeDatePickerCell(ptr unsafe.Pointer) DatePickerCell {
	return DatePickerCell{
		ActionCell: MakeActionCell(ptr),
	}
}

func (d_ DatePickerCell) InitTextCell(string_ string) DatePickerCell {
	rv := objc.CallMethod[DatePickerCell](d_, objc.GetSelector("initTextCell:"), string_)
	return rv
}

func DatePickerCell_InitTextCell(string_ string) DatePickerCell {
	return DatePickerCellClass.Alloc().InitTextCell(string_)
}

func (d_ DatePickerCell) InitImageCell(image IImage) DatePickerCell {
	rv := objc.CallMethod[DatePickerCell](d_, objc.GetSelector("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func DatePickerCell_InitImageCell(image IImage) DatePickerCell {
	return DatePickerCellClass.Alloc().InitImageCell(image)
}

func (d_ DatePickerCell) Init() DatePickerCell {
	rv := objc.CallMethod[DatePickerCell](d_, objc.GetSelector("init"))
	return rv
}

func DatePickerCell_Init() DatePickerCell {
	return DatePickerCellClass.Alloc().Init()
}

func (dc _DatePickerCellClass) Alloc() DatePickerCell {
	rv := objc.CallMethod[DatePickerCell](dc, objc.GetSelector("alloc"))
	return rv
}

func DatePickerCell_Alloc() DatePickerCell {
	return DatePickerCellClass.Alloc()
}

func (dc _DatePickerCellClass) New() DatePickerCell {
	rv := objc.CallMethod[DatePickerCell](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDatePickerCell() DatePickerCell {
	return DatePickerCellClass.New()
}

func DatePickerCell_New() DatePickerCell {
	return DatePickerCellClass.New()
}

func (d_ DatePickerCell) BackgroundColor() Color {
	rv := objc.CallMethod[Color](d_, objc.GetSelector("backgroundColor"))
	return rv
}

func (d_ DatePickerCell) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (d_ DatePickerCell) DrawsBackground() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("drawsBackground"))
	return rv
}

func (d_ DatePickerCell) SetDrawsBackground(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDrawsBackground:"), value)
}

func (d_ DatePickerCell) TextColor() Color {
	rv := objc.CallMethod[Color](d_, objc.GetSelector("textColor"))
	return rv
}

func (d_ DatePickerCell) SetTextColor(value IColor) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setTextColor:"), objc.ExtractPtr(value))
}

func (d_ DatePickerCell) DatePickerStyle() DatePickerStyle {
	rv := objc.CallMethod[DatePickerStyle](d_, objc.GetSelector("datePickerStyle"))
	return rv
}

func (d_ DatePickerCell) SetDatePickerStyle(value DatePickerStyle) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDatePickerStyle:"), value)
}

func (d_ DatePickerCell) DatePickerElements() DatePickerElementFlags {
	rv := objc.CallMethod[DatePickerElementFlags](d_, objc.GetSelector("datePickerElements"))
	return rv
}

func (d_ DatePickerCell) SetDatePickerElements(value DatePickerElementFlags) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDatePickerElements:"), value)
}

func (d_ DatePickerCell) DatePickerMode() DatePickerMode {
	rv := objc.CallMethod[DatePickerMode](d_, objc.GetSelector("datePickerMode"))
	return rv
}

func (d_ DatePickerCell) SetDatePickerMode(value DatePickerMode) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDatePickerMode:"), value)
}

func (d_ DatePickerCell) DateValue() foundation.Date {
	rv := objc.CallMethod[foundation.Date](d_, objc.GetSelector("dateValue"))
	return rv
}

func (d_ DatePickerCell) SetDateValue(value foundation.IDate) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDateValue:"), objc.ExtractPtr(value))
}

func (d_ DatePickerCell) TimeInterval() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](d_, objc.GetSelector("timeInterval"))
	return rv
}

func (d_ DatePickerCell) SetTimeInterval(value foundation.TimeInterval) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setTimeInterval:"), value)
}

func (d_ DatePickerCell) Calendar() foundation.Calendar {
	rv := objc.CallMethod[foundation.Calendar](d_, objc.GetSelector("calendar"))
	return rv
}

func (d_ DatePickerCell) SetCalendar(value foundation.ICalendar) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setCalendar:"), objc.ExtractPtr(value))
}

func (d_ DatePickerCell) Locale() foundation.Locale {
	rv := objc.CallMethod[foundation.Locale](d_, objc.GetSelector("locale"))
	return rv
}

func (d_ DatePickerCell) SetLocale(value foundation.ILocale) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setLocale:"), objc.ExtractPtr(value))
}

func (d_ DatePickerCell) TimeZone() foundation.TimeZone {
	rv := objc.CallMethod[foundation.TimeZone](d_, objc.GetSelector("timeZone"))
	return rv
}

func (d_ DatePickerCell) SetTimeZone(value foundation.ITimeZone) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setTimeZone:"), objc.ExtractPtr(value))
}

func (d_ DatePickerCell) MinDate() foundation.Date {
	rv := objc.CallMethod[foundation.Date](d_, objc.GetSelector("minDate"))
	return rv
}

func (d_ DatePickerCell) SetMinDate(value foundation.IDate) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setMinDate:"), objc.ExtractPtr(value))
}

func (d_ DatePickerCell) MaxDate() foundation.Date {
	rv := objc.CallMethod[foundation.Date](d_, objc.GetSelector("maxDate"))
	return rv
}

func (d_ DatePickerCell) SetMaxDate(value foundation.IDate) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setMaxDate:"), objc.ExtractPtr(value))
}

func (d_ DatePickerCell) Delegate() DatePickerCellDelegateWrapper {
	rv := objc.CallMethod[DatePickerCellDelegateWrapper](d_, objc.GetSelector("delegate"))
	return rv
}

func (d_ DatePickerCell) SetDelegate(value IDatePickerCellDelegate) {
	po := objc.WrapAsProtocol("NSDatePickerCellDelegate", value)
	objc.SetAssociatedObject(d_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDelegate:"), po)
}

func (d_ DatePickerCell) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}
