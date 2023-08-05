// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var SliderClass = _SliderClass{objc.GetClass("NSSlider")}

type _SliderClass struct {
	objc.Class
}

type ISlider interface {
	IControl
	ClosestTickMarkValueToValue(value float64) float64
	IndexOfTickMarkAtPoint(point foundation.Point) int
	RectOfTickMarkAtIndex(index int) foundation.Rect
	TickMarkValueAtIndex(index int) float64
	SliderType() SliderType
	SetSliderType(value SliderType)
	AltIncrementValue() float64
	SetAltIncrementValue(value float64)
	KnobThickness() float64
	IsVertical() bool
	SetVertical(value bool)
	TrackFillColor() Color
	SetTrackFillColor(value IColor)
	MaxValue() float64
	SetMaxValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
	AllowsTickMarkValuesOnly() bool
	SetAllowsTickMarkValuesOnly(value bool)
	NumberOfTickMarks() int
	SetNumberOfTickMarks(value int)
	TickMarkPosition() TickMarkPosition
	SetTickMarkPosition(value TickMarkPosition)
}

type Slider struct {
	Control
}

func MakeSlider(ptr unsafe.Pointer) Slider {
	return Slider{
		Control: MakeControl(ptr),
	}
}

func (sc _SliderClass) SliderWithTargetAction(target objc.IObject, action objc.Selector) Slider {
	rv := objc.CallMethod[Slider](sc, objc.GetSelector("sliderWithTarget:action:"), objc.ExtractPtr(target), action)
	return rv
}

func Slider_SliderWithTargetAction(target objc.IObject, action objc.Selector) Slider {
	return SliderClass.SliderWithTargetAction(target, action)
}

func (sc _SliderClass) SliderWithValueMinValueMaxValueTargetAction(value float64, minValue float64, maxValue float64, target objc.IObject, action objc.Selector) Slider {
	rv := objc.CallMethod[Slider](sc, objc.GetSelector("sliderWithValue:minValue:maxValue:target:action:"), value, minValue, maxValue, objc.ExtractPtr(target), action)
	return rv
}

func Slider_SliderWithValueMinValueMaxValueTargetAction(value float64, minValue float64, maxValue float64, target objc.IObject, action objc.Selector) Slider {
	return SliderClass.SliderWithValueMinValueMaxValueTargetAction(value, minValue, maxValue, target, action)
}

func (s_ Slider) InitWithFrame(frameRect foundation.Rect) Slider {
	rv := objc.CallMethod[Slider](s_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func Slider_InitWithFrame(frameRect foundation.Rect) Slider {
	return SliderClass.Alloc().InitWithFrame(frameRect)
}

func (s_ Slider) Init() Slider {
	rv := objc.CallMethod[Slider](s_, objc.GetSelector("init"))
	return rv
}

func Slider_Init() Slider {
	return SliderClass.Alloc().Init()
}

func (sc _SliderClass) Alloc() Slider {
	rv := objc.CallMethod[Slider](sc, objc.GetSelector("alloc"))
	return rv
}

func Slider_Alloc() Slider {
	return SliderClass.Alloc()
}

func (sc _SliderClass) New() Slider {
	rv := objc.CallMethod[Slider](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewSlider() Slider {
	return SliderClass.New()
}

func Slider_New() Slider {
	return SliderClass.New()
}

func (s_ Slider) ClosestTickMarkValueToValue(value float64) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("closestTickMarkValueToValue:"), value)
	return rv
}

func (s_ Slider) IndexOfTickMarkAtPoint(point foundation.Point) int {
	rv := objc.CallMethod[int](s_, objc.GetSelector("indexOfTickMarkAtPoint:"), point)
	return rv
}

func (s_ Slider) RectOfTickMarkAtIndex(index int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("rectOfTickMarkAtIndex:"), index)
	return rv
}

func (s_ Slider) TickMarkValueAtIndex(index int) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("tickMarkValueAtIndex:"), index)
	return rv
}

func (s_ Slider) SliderType() SliderType {
	rv := objc.CallMethod[SliderType](s_, objc.GetSelector("sliderType"))
	return rv
}

func (s_ Slider) SetSliderType(value SliderType) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSliderType:"), value)
}

func (s_ Slider) AltIncrementValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("altIncrementValue"))
	return rv
}

func (s_ Slider) SetAltIncrementValue(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAltIncrementValue:"), value)
}

func (s_ Slider) KnobThickness() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("knobThickness"))
	return rv
}

func (s_ Slider) IsVertical() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("isVertical"))
	return rv
}

func (s_ Slider) SetVertical(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setVertical:"), value)
}

func (s_ Slider) TrackFillColor() Color {
	rv := objc.CallMethod[Color](s_, objc.GetSelector("trackFillColor"))
	return rv
}

func (s_ Slider) SetTrackFillColor(value IColor) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setTrackFillColor:"), objc.ExtractPtr(value))
}

func (s_ Slider) MaxValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("maxValue"))
	return rv
}

func (s_ Slider) SetMaxValue(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMaxValue:"), value)
}

func (s_ Slider) MinValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("minValue"))
	return rv
}

func (s_ Slider) SetMinValue(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMinValue:"), value)
}

func (s_ Slider) AllowsTickMarkValuesOnly() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("allowsTickMarkValuesOnly"))
	return rv
}

func (s_ Slider) SetAllowsTickMarkValuesOnly(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAllowsTickMarkValuesOnly:"), value)
}

func (s_ Slider) NumberOfTickMarks() int {
	rv := objc.CallMethod[int](s_, objc.GetSelector("numberOfTickMarks"))
	return rv
}

func (s_ Slider) SetNumberOfTickMarks(value int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setNumberOfTickMarks:"), value)
}

func (s_ Slider) TickMarkPosition() TickMarkPosition {
	rv := objc.CallMethod[TickMarkPosition](s_, objc.GetSelector("tickMarkPosition"))
	return rv
}

func (s_ Slider) SetTickMarkPosition(value TickMarkPosition) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setTickMarkPosition:"), value)
}
