// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DatePickerCell] class.
var DatePickerCellClass = _DatePickerCellClass{objc.GetClass("NSDatePickerCell")}

type _DatePickerCellClass struct {
	objc.Class
}

// An interface definition for the [DatePickerCell] class.
type IDatePickerCell interface {
	IActionCell
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
	TextColor() Color
	SetTextColor(value IColor)
	Calendar() foundation.Calendar
	SetCalendar(value foundation.ICalendar)
	DatePickerMode() DatePickerMode
	SetDatePickerMode(value DatePickerMode)
	MaxDate() foundation.Date
	SetMaxDate(value foundation.IDate)
	DatePickerStyle() DatePickerStyle
	SetDatePickerStyle(value DatePickerStyle)
}

// An object that controls the behavior of a date picker, or of a single date picker cell in a matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell?language=objc
type DatePickerCell struct {
	ActionCell
}

func DatePickerCellFrom(ptr unsafe.Pointer) DatePickerCell {
	return DatePickerCell{
		ActionCell: ActionCellFrom(ptr),
	}
}

func (d_ DatePickerCell) InitTextCell(string_ string) DatePickerCell {
	rv := objc.Call[DatePickerCell](d_, objc.Sel("initTextCell:"), string_)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1643566-inittextcell?language=objc
func DatePickerCell_InitTextCell(string_ string) DatePickerCell {
	return DatePickerCellClass.Alloc().InitTextCell(string_)
}

func (dc _DatePickerCellClass) Alloc() DatePickerCell {
	rv := objc.Call[DatePickerCell](dc, objc.Sel("alloc"))
	return rv
}

func DatePickerCell_Alloc() DatePickerCell {
	return DatePickerCellClass.Alloc()
}

func (dc _DatePickerCellClass) New() DatePickerCell {
	rv := objc.Call[DatePickerCell](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDatePickerCell() DatePickerCell {
	return DatePickerCellClass.New()
}

func (d_ DatePickerCell) Init() DatePickerCell {
	rv := objc.Call[DatePickerCell](d_, objc.Sel("init"))
	return rv
}

func (d_ DatePickerCell) InitImageCell(image IImage) DatePickerCell {
	rv := objc.Call[DatePickerCell](d_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

// Returns an NSCell object initialized with the specified image and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533898-initimagecell?language=objc
func DatePickerCell_InitImageCell(image IImage) DatePickerCell {
	return DatePickerCellClass.Alloc().InitImageCell(image)
}

// A bitmask that indicates which visual elements are shown by the date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459639-datepickerelements?language=objc
func (d_ DatePickerCell) DatePickerElements() DatePickerElementFlags {
	rv := objc.Call[DatePickerElementFlags](d_, objc.Sel("datePickerElements"))
	return rv
}

// A bitmask that indicates which visual elements are shown by the date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459639-datepickerelements?language=objc
func (d_ DatePickerCell) SetDatePickerElements(value DatePickerElementFlags) {
	objc.Call[objc.Void](d_, objc.Sel("setDatePickerElements:"), value)
}

// The time interval that represents the date range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459589-timeinterval?language=objc
func (d_ DatePickerCell) TimeInterval() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](d_, objc.Sel("timeInterval"))
	return rv
}

// The time interval that represents the date range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459589-timeinterval?language=objc
func (d_ DatePickerCell) SetTimeInterval(value foundation.TimeInterval) {
	objc.Call[objc.Void](d_, objc.Sel("setTimeInterval:"), value)
}

// The locale used to display dates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459625-locale?language=objc
func (d_ DatePickerCell) Locale() foundation.Locale {
	rv := objc.Call[foundation.Locale](d_, objc.Sel("locale"))
	return rv
}

// The locale used to display dates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459625-locale?language=objc
func (d_ DatePickerCell) SetLocale(value foundation.ILocale) {
	objc.Call[objc.Void](d_, objc.Sel("setLocale:"), objc.Ptr(value))
}

// The delegate associated with the date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459615-delegate?language=objc
func (d_ DatePickerCell) Delegate() DatePickerCellDelegateWrapper {
	rv := objc.Call[DatePickerCellDelegateWrapper](d_, objc.Sel("delegate"))
	return rv
}

// The delegate associated with the date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459615-delegate?language=objc
func (d_ DatePickerCell) SetDelegate(value PDatePickerCellDelegate) {
	po0 := objc.WrapAsProtocol("NSDatePickerCellDelegate", value)
	objc.SetAssociatedObject(d_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](d_, objc.Sel("setDelegate:"), po0)
}

// The delegate associated with the date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459615-delegate?language=objc
func (d_ DatePickerCell) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](d_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The cell’s background color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459629-backgroundcolor?language=objc
func (d_ DatePickerCell) BackgroundColor() Color {
	rv := objc.Call[Color](d_, objc.Sel("backgroundColor"))
	return rv
}

// The cell’s background color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459629-backgroundcolor?language=objc
func (d_ DatePickerCell) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](d_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// The minimum date that the picker allows as input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459619-mindate?language=objc
func (d_ DatePickerCell) MinDate() foundation.Date {
	rv := objc.Call[foundation.Date](d_, objc.Sel("minDate"))
	return rv
}

// The minimum date that the picker allows as input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459619-mindate?language=objc
func (d_ DatePickerCell) SetMinDate(value foundation.IDate) {
	objc.Call[objc.Void](d_, objc.Sel("setMinDate:"), objc.Ptr(value))
}

// A Boolean value indicating whether the cell draws its background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459591-drawsbackground?language=objc
func (d_ DatePickerCell) DrawsBackground() bool {
	rv := objc.Call[bool](d_, objc.Sel("drawsBackground"))
	return rv
}

// A Boolean value indicating whether the cell draws its background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459591-drawsbackground?language=objc
func (d_ DatePickerCell) SetDrawsBackground(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setDrawsBackground:"), value)
}

// The time zone used to display time-related values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459633-timezone?language=objc
func (d_ DatePickerCell) TimeZone() foundation.TimeZone {
	rv := objc.Call[foundation.TimeZone](d_, objc.Sel("timeZone"))
	return rv
}

// The time zone used to display time-related values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459633-timezone?language=objc
func (d_ DatePickerCell) SetTimeZone(value foundation.ITimeZone) {
	objc.Call[objc.Void](d_, objc.Sel("setTimeZone:"), objc.Ptr(value))
}

// The date currently specified in the picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459594-datevalue?language=objc
func (d_ DatePickerCell) DateValue() foundation.Date {
	rv := objc.Call[foundation.Date](d_, objc.Sel("dateValue"))
	return rv
}

// The date currently specified in the picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459594-datevalue?language=objc
func (d_ DatePickerCell) SetDateValue(value foundation.IDate) {
	objc.Call[objc.Void](d_, objc.Sel("setDateValue:"), objc.Ptr(value))
}

// The cell’s text color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459644-textcolor?language=objc
func (d_ DatePickerCell) TextColor() Color {
	rv := objc.Call[Color](d_, objc.Sel("textColor"))
	return rv
}

// The cell’s text color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459644-textcolor?language=objc
func (d_ DatePickerCell) SetTextColor(value IColor) {
	objc.Call[objc.Void](d_, objc.Sel("setTextColor:"), objc.Ptr(value))
}

// The calendar used by the date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459609-calendar?language=objc
func (d_ DatePickerCell) Calendar() foundation.Calendar {
	rv := objc.Call[foundation.Calendar](d_, objc.Sel("calendar"))
	return rv
}

// The calendar used by the date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459609-calendar?language=objc
func (d_ DatePickerCell) SetCalendar(value foundation.ICalendar) {
	objc.Call[objc.Void](d_, objc.Sel("setCalendar:"), objc.Ptr(value))
}

// The mode in use by the date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459602-datepickermode?language=objc
func (d_ DatePickerCell) DatePickerMode() DatePickerMode {
	rv := objc.Call[DatePickerMode](d_, objc.Sel("datePickerMode"))
	return rv
}

// The mode in use by the date picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459602-datepickermode?language=objc
func (d_ DatePickerCell) SetDatePickerMode(value DatePickerMode) {
	objc.Call[objc.Void](d_, objc.Sel("setDatePickerMode:"), value)
}

// The maximum date that the picker allows as input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459640-maxdate?language=objc
func (d_ DatePickerCell) MaxDate() foundation.Date {
	rv := objc.Call[foundation.Date](d_, objc.Sel("maxDate"))
	return rv
}

// The maximum date that the picker allows as input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459640-maxdate?language=objc
func (d_ DatePickerCell) SetMaxDate(value foundation.IDate) {
	objc.Call[objc.Void](d_, objc.Sel("setMaxDate:"), objc.Ptr(value))
}

// The date picker style to use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459635-datepickerstyle?language=objc
func (d_ DatePickerCell) DatePickerStyle() DatePickerStyle {
	rv := objc.Call[DatePickerStyle](d_, objc.Sel("datePickerStyle"))
	return rv
}

// The date picker style to use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/1459635-datepickerstyle?language=objc
func (d_ DatePickerCell) SetDatePickerStyle(value DatePickerStyle) {
	objc.Call[objc.Void](d_, objc.Sel("setDatePickerStyle:"), value)
}
