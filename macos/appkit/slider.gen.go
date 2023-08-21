// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Slider] class.
var SliderClass = _SliderClass{objc.GetClass("NSSlider")}

type _SliderClass struct {
	objc.Class
}

// An interface definition for the [Slider] class.
type ISlider interface {
	IControl
	RectOfTickMarkAtIndex(index int) foundation.Rect
	ClosestTickMarkValueToValue(value float64) float64
	TickMarkValueAtIndex(index int) float64
	IndexOfTickMarkAtPoint(point foundation.Point) int
	AltIncrementValue() float64
	SetAltIncrementValue(value float64)
	TickMarkPosition() TickMarkPosition
	SetTickMarkPosition(value TickMarkPosition)
	KnobThickness() float64
	SliderType() SliderType
	SetSliderType(value SliderType)
	MinValue() float64
	SetMinValue(value float64)
	MaxValue() float64
	SetMaxValue(value float64)
	TrackFillColor() Color
	SetTrackFillColor(value IColor)
	IsVertical() bool
	SetVertical(value bool)
	AllowsTickMarkValuesOnly() bool
	SetAllowsTickMarkValuesOnly(value bool)
	NumberOfTickMarks() int
	SetNumberOfTickMarks(value int)
}

// A display of a bar representing a continuous range of numerical values and a knob representing the currently selected value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider?language=objc
type Slider struct {
	Control
}

func SliderFrom(ptr unsafe.Pointer) Slider {
	return Slider{
		Control: ControlFrom(ptr),
	}
}

func (sc _SliderClass) SliderWithValueMinValueMaxValueTargetAction(value float64, minValue float64, maxValue float64, target objc.IObject, action objc.Selector) Slider {
	rv := objc.Call[Slider](sc, objc.Sel("sliderWithValue:minValue:maxValue:target:action:"), value, minValue, maxValue, target, action)
	return rv
}

// Creates a continuous horizontal slider that represents values over the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1644495-sliderwithvalue?language=objc
func Slider_SliderWithValueMinValueMaxValueTargetAction(value float64, minValue float64, maxValue float64, target objc.IObject, action objc.Selector) Slider {
	return SliderClass.SliderWithValueMinValueMaxValueTargetAction(value, minValue, maxValue, target, action)
}

func (sc _SliderClass) SliderWithTargetAction(target objc.IObject, action objc.Selector) Slider {
	rv := objc.Call[Slider](sc, objc.Sel("sliderWithTarget:action:"), target, action)
	return rv
}

// Creates a continuous horizontal slider whose values range from 0.0 to 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1644494-sliderwithtarget?language=objc
func Slider_SliderWithTargetAction(target objc.IObject, action objc.Selector) Slider {
	return SliderClass.SliderWithTargetAction(target, action)
}

func (sc _SliderClass) Alloc() Slider {
	rv := objc.Call[Slider](sc, objc.Sel("alloc"))
	return rv
}

func Slider_Alloc() Slider {
	return SliderClass.Alloc()
}

func (sc _SliderClass) New() Slider {
	rv := objc.Call[Slider](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSlider() Slider {
	return SliderClass.New()
}

func (s_ Slider) Init() Slider {
	rv := objc.Call[Slider](s_, objc.Sel("init"))
	return rv
}

func (s_ Slider) InitWithFrame(frameRect foundation.Rect) Slider {
	rv := objc.Call[Slider](s_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func NewSliderWithFrame(frameRect foundation.Rect) Slider {
	instance := SliderClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Returns the bounding rectangle of the tick mark at the given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1532922-rectoftickmarkatindex?language=objc
func (s_ Slider) RectOfTickMarkAtIndex(index int) foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("rectOfTickMarkAtIndex:"), index)
	return rv
}

// Returns the value of the tick mark closest to the specified value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1530237-closesttickmarkvaluetovalue?language=objc
func (s_ Slider) ClosestTickMarkValueToValue(value float64) float64 {
	rv := objc.Call[float64](s_, objc.Sel("closestTickMarkValueToValue:"), value)
	return rv
}

// Returns the slider’s value represented by the tick mark at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1527230-tickmarkvalueatindex?language=objc
func (s_ Slider) TickMarkValueAtIndex(index int) float64 {
	rv := objc.Call[float64](s_, objc.Sel("tickMarkValueAtIndex:"), index)
	return rv
}

// Returns the index of the tick mark closest to the location of the slider represented by the given point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1526763-indexoftickmarkatpoint?language=objc
func (s_ Slider) IndexOfTickMarkAtPoint(point foundation.Point) int {
	rv := objc.Call[int](s_, objc.Sel("indexOfTickMarkAtPoint:"), point)
	return rv
}

// The amount by which the slider changes its value when the user Option-drags the slider knob. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1532901-altincrementvalue?language=objc
func (s_ Slider) AltIncrementValue() float64 {
	rv := objc.Call[float64](s_, objc.Sel("altIncrementValue"))
	return rv
}

// The amount by which the slider changes its value when the user Option-drags the slider knob. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1532901-altincrementvalue?language=objc
func (s_ Slider) SetAltIncrementValue(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setAltIncrementValue:"), value)
}

// Determines where the slider’s tick marks are displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1529657-tickmarkposition?language=objc
func (s_ Slider) TickMarkPosition() TickMarkPosition {
	rv := objc.Call[TickMarkPosition](s_, objc.Sel("tickMarkPosition"))
	return rv
}

// Determines where the slider’s tick marks are displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1529657-tickmarkposition?language=objc
func (s_ Slider) SetTickMarkPosition(value TickMarkPosition) {
	objc.Call[objc.Void](s_, objc.Sel("setTickMarkPosition:"), value)
}

// The knob’s thickness, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1532909-knobthickness?language=objc
func (s_ Slider) KnobThickness() float64 {
	rv := objc.Call[float64](s_, objc.Sel("knobThickness"))
	return rv
}

// The type of the slider, such as vertical or circular. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1532924-slidertype?language=objc
func (s_ Slider) SliderType() SliderType {
	rv := objc.Call[SliderType](s_, objc.Sel("sliderType"))
	return rv
}

// The type of the slider, such as vertical or circular. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1532924-slidertype?language=objc
func (s_ Slider) SetSliderType(value SliderType) {
	objc.Call[objc.Void](s_, objc.Sel("setSliderType:"), value)
}

// The minimum value the slider can send to its target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1524665-minvalue?language=objc
func (s_ Slider) MinValue() float64 {
	rv := objc.Call[float64](s_, objc.Sel("minValue"))
	return rv
}

// The minimum value the slider can send to its target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1524665-minvalue?language=objc
func (s_ Slider) SetMinValue(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMinValue:"), value)
}

// The maximum value the slider can send to its target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1532919-maxvalue?language=objc
func (s_ Slider) MaxValue() float64 {
	rv := objc.Call[float64](s_, objc.Sel("maxValue"))
	return rv
}

// The maximum value the slider can send to its target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1532919-maxvalue?language=objc
func (s_ Slider) SetMaxValue(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMaxValue:"), value)
}

// The color of the filled portion of the slider track, in appearances that support it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/2560999-trackfillcolor?language=objc
func (s_ Slider) TrackFillColor() Color {
	rv := objc.Call[Color](s_, objc.Sel("trackFillColor"))
	return rv
}

// The color of the filled portion of the slider track, in appearances that support it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/2560999-trackfillcolor?language=objc
func (s_ Slider) SetTrackFillColor(value IColor) {
	objc.Call[objc.Void](s_, objc.Sel("setTrackFillColor:"), objc.Ptr(value))
}

// An integer indicating the orientation (horizontal or vertical) of the slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1527901-vertical?language=objc
func (s_ Slider) IsVertical() bool {
	rv := objc.Call[bool](s_, objc.Sel("isVertical"))
	return rv
}

// An integer indicating the orientation (horizontal or vertical) of the slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1527901-vertical?language=objc
func (s_ Slider) SetVertical(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setVertical:"), value)
}

// A Boolean value that indicates whether the slider fixes its values to those values represented by its tick marks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1526898-allowstickmarkvaluesonly?language=objc
func (s_ Slider) AllowsTickMarkValuesOnly() bool {
	rv := objc.Call[bool](s_, objc.Sel("allowsTickMarkValuesOnly"))
	return rv
}

// A Boolean value that indicates whether the slider fixes its values to those values represented by its tick marks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1526898-allowstickmarkvaluesonly?language=objc
func (s_ Slider) SetAllowsTickMarkValuesOnly(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setAllowsTickMarkValuesOnly:"), value)
}

// The number of tick marks associated with the slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1524268-numberoftickmarks?language=objc
func (s_ Slider) NumberOfTickMarks() int {
	rv := objc.Call[int](s_, objc.Sel("numberOfTickMarks"))
	return rv
}

// The number of tick marks associated with the slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/1524268-numberoftickmarks?language=objc
func (s_ Slider) SetNumberOfTickMarks(value int) {
	objc.Call[objc.Void](s_, objc.Sel("setNumberOfTickMarks:"), value)
}
