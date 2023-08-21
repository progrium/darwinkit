// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SliderCell] class.
var SliderCellClass = _SliderCellClass{objc.GetClass("NSSliderCell")}

type _SliderCellClass struct {
	objc.Class
}

// An interface definition for the [SliderCell] class.
type ISliderCell interface {
	IActionCell
	RectOfTickMarkAtIndex(index int) foundation.Rect
	BarRectFlipped(flipped bool) foundation.Rect
	DrawTickMarks()
	DrawBarInsideFlipped(rect foundation.Rect, flipped bool)
	DrawKnob(knobRect foundation.Rect)
	ClosestTickMarkValueToValue(value float64) float64
	TickMarkValueAtIndex(index int) float64
	KnobRectFlipped(flipped bool) foundation.Rect
	IndexOfTickMarkAtPoint(point foundation.Point) int
	AltIncrementValue() float64
	SetAltIncrementValue(value float64)
	TickMarkPosition() TickMarkPosition
	SetTickMarkPosition(value TickMarkPosition)
	KnobThickness() float64
	TrackRect() foundation.Rect
	SliderType() SliderType
	SetSliderType(value SliderType)
	MinValue() float64
	SetMinValue(value float64)
	MaxValue() float64
	SetMaxValue(value float64)
	IsVertical() bool
	SetVertical(value bool)
	AllowsTickMarkValuesOnly() bool
	SetAllowsTickMarkValuesOnly(value bool)
	NumberOfTickMarks() int
	SetNumberOfTickMarks(value int)
}

// The appearance and behavior of an NSSlider object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell?language=objc
type SliderCell struct {
	ActionCell
}

func SliderCellFrom(ptr unsafe.Pointer) SliderCell {
	return SliderCell{
		ActionCell: ActionCellFrom(ptr),
	}
}

func (sc _SliderCellClass) Alloc() SliderCell {
	rv := objc.Call[SliderCell](sc, objc.Sel("alloc"))
	return rv
}

func SliderCell_Alloc() SliderCell {
	return SliderCellClass.Alloc()
}

func (sc _SliderCellClass) New() SliderCell {
	rv := objc.Call[SliderCell](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSliderCell() SliderCell {
	return SliderCellClass.New()
}

func (s_ SliderCell) Init() SliderCell {
	rv := objc.Call[SliderCell](s_, objc.Sel("init"))
	return rv
}

func (s_ SliderCell) InitImageCell(image IImage) SliderCell {
	rv := objc.Call[SliderCell](s_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

// Returns an NSCell object initialized with the specified image and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533898-initimagecell?language=objc
func NewSliderCellImageCell(image IImage) SliderCell {
	instance := SliderCellClass.Alloc().InitImageCell(image)
	instance.Autorelease()
	return instance
}

func (s_ SliderCell) InitTextCell(string_ string) SliderCell {
	rv := objc.Call[SliderCell](s_, objc.Sel("initTextCell:"), string_)
	return rv
}

// Returns an NSCell object initialized with the specified string and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530851-inittextcell?language=objc
func NewSliderCellTextCell(string_ string) SliderCell {
	instance := SliderCellClass.Alloc().InitTextCell(string_)
	instance.Autorelease()
	return instance
}

// Returns the bounding rectangle of the tick mark at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444645-rectoftickmarkatindex?language=objc
func (s_ SliderCell) RectOfTickMarkAtIndex(index int) foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("rectOfTickMarkAtIndex:"), index)
	return rv
}

// Returns the rectangle in which the bar is drawn. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444629-barrectflipped?language=objc
func (s_ SliderCell) BarRectFlipped(flipped bool) foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("barRectFlipped:"), flipped)
	return rv
}

// Draws the slider’s tick marks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444633-drawtickmarks?language=objc
func (s_ SliderCell) DrawTickMarks() {
	objc.Call[objc.Void](s_, objc.Sel("drawTickMarks"))
}

// Draws the slider’s bar—but not its bezel or knob—inside the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444587-drawbarinside?language=objc
func (s_ SliderCell) DrawBarInsideFlipped(rect foundation.Rect, flipped bool) {
	objc.Call[objc.Void](s_, objc.Sel("drawBarInside:flipped:"), rect, flipped)
}

// Draws the slider knob in the given rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444600-drawknob?language=objc
func (s_ SliderCell) DrawKnob(knobRect foundation.Rect) {
	objc.Call[objc.Void](s_, objc.Sel("drawKnob:"), knobRect)
}

// Returns the value of the tick mark closest to the specified value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444627-closesttickmarkvaluetovalue?language=objc
func (s_ SliderCell) ClosestTickMarkValueToValue(value float64) float64 {
	rv := objc.Call[float64](s_, objc.Sel("closestTickMarkValueToValue:"), value)
	return rv
}

// Returns the receiver’s value represented by the tick mark at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444635-tickmarkvalueatindex?language=objc
func (s_ SliderCell) TickMarkValueAtIndex(index int) float64 {
	rv := objc.Call[float64](s_, objc.Sel("tickMarkValueAtIndex:"), index)
	return rv
}

// Returns the rectangle in which the slider knob is drawn. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444647-knobrectflipped?language=objc
func (s_ SliderCell) KnobRectFlipped(flipped bool) foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("knobRectFlipped:"), flipped)
	return rv
}

// Returns the index of the tick mark closest to the location of the slider represented by the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444614-indexoftickmarkatpoint?language=objc
func (s_ SliderCell) IndexOfTickMarkAtPoint(point foundation.Point) int {
	rv := objc.Call[int](s_, objc.Sel("indexOfTickMarkAtPoint:"), point)
	return rv
}

// The amount by which the slider changes its value when the user Option-drags the knob. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444596-altincrementvalue?language=objc
func (s_ SliderCell) AltIncrementValue() float64 {
	rv := objc.Call[float64](s_, objc.Sel("altIncrementValue"))
	return rv
}

// The amount by which the slider changes its value when the user Option-drags the knob. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444596-altincrementvalue?language=objc
func (s_ SliderCell) SetAltIncrementValue(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setAltIncrementValue:"), value)
}

// The position of the tick marks relative to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444616-tickmarkposition?language=objc
func (s_ SliderCell) TickMarkPosition() TickMarkPosition {
	rv := objc.Call[TickMarkPosition](s_, objc.Sel("tickMarkPosition"))
	return rv
}

// The position of the tick marks relative to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444616-tickmarkposition?language=objc
func (s_ SliderCell) SetTickMarkPosition(value TickMarkPosition) {
	objc.Call[objc.Void](s_, objc.Sel("setTickMarkPosition:"), value)
}

// The thickness of the slider knob, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444593-knobthickness?language=objc
func (s_ SliderCell) KnobThickness() float64 {
	rv := objc.Call[float64](s_, objc.Sel("knobThickness"))
	return rv
}

// The rectangle within which the cell tracks the pointer while the mouse button is down. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444583-trackrect?language=objc
func (s_ SliderCell) TrackRect() foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("trackRect"))
	return rv
}

// The slider type, either linear or circular. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444598-slidertype?language=objc
func (s_ SliderCell) SliderType() SliderType {
	rv := objc.Call[SliderType](s_, objc.Sel("sliderType"))
	return rv
}

// The slider type, either linear or circular. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444598-slidertype?language=objc
func (s_ SliderCell) SetSliderType(value SliderType) {
	objc.Call[objc.Void](s_, objc.Sel("setSliderType:"), value)
}

// The minimum value the slider can send to its target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444641-minvalue?language=objc
func (s_ SliderCell) MinValue() float64 {
	rv := objc.Call[float64](s_, objc.Sel("minValue"))
	return rv
}

// The minimum value the slider can send to its target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444641-minvalue?language=objc
func (s_ SliderCell) SetMinValue(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMinValue:"), value)
}

// The maximum value the slider can send to its target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444589-maxvalue?language=objc
func (s_ SliderCell) MaxValue() float64 {
	rv := objc.Call[float64](s_, objc.Sel("maxValue"))
	return rv
}

// The maximum value the slider can send to its target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444589-maxvalue?language=objc
func (s_ SliderCell) SetMaxValue(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMaxValue:"), value)
}

// An integer indicating the orientation (vertical or horizontal) of the slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444602-vertical?language=objc
func (s_ SliderCell) IsVertical() bool {
	rv := objc.Call[bool](s_, objc.Sel("isVertical"))
	return rv
}

// An integer indicating the orientation (vertical or horizontal) of the slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444602-vertical?language=objc
func (s_ SliderCell) SetVertical(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setVertical:"), value)
}

// A Boolean value indicating whether the receiver fixes its values to those values represented by its tick marks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444604-allowstickmarkvaluesonly?language=objc
func (s_ SliderCell) AllowsTickMarkValuesOnly() bool {
	rv := objc.Call[bool](s_, objc.Sel("allowsTickMarkValuesOnly"))
	return rv
}

// A Boolean value indicating whether the receiver fixes its values to those values represented by its tick marks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444604-allowstickmarkvaluesonly?language=objc
func (s_ SliderCell) SetAllowsTickMarkValuesOnly(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setAllowsTickMarkValuesOnly:"), value)
}

// The number of tick marks associated with the slider, including the tick marks assigned to the minimum and maximum values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444621-numberoftickmarks?language=objc
func (s_ SliderCell) NumberOfTickMarks() int {
	rv := objc.Call[int](s_, objc.Sel("numberOfTickMarks"))
	return rv
}

// The number of tick marks associated with the slider, including the tick marks assigned to the minimum and maximum values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/1444621-numberoftickmarks?language=objc
func (s_ SliderCell) SetNumberOfTickMarks(value int) {
	objc.Call[objc.Void](s_, objc.Sel("setNumberOfTickMarks:"), value)
}
