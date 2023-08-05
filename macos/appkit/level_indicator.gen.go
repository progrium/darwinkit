// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var LevelIndicatorClass = _LevelIndicatorClass{objc.GetClass("NSLevelIndicator")}

type _LevelIndicatorClass struct {
	objc.Class
}

type ILevelIndicator interface {
	IControl
	TickMarkValueAtIndex(index int) float64
	RectOfTickMarkAtIndex(index int) foundation.Rect
	MinValue() float64
	SetMinValue(value float64)
	MaxValue() float64
	SetMaxValue(value float64)
	WarningValue() float64
	SetWarningValue(value float64)
	CriticalValue() float64
	SetCriticalValue(value float64)
	TickMarkPosition() TickMarkPosition
	SetTickMarkPosition(value TickMarkPosition)
	NumberOfTickMarks() int
	SetNumberOfTickMarks(value int)
	NumberOfMajorTickMarks() int
	SetNumberOfMajorTickMarks(value int)
	LevelIndicatorStyle() LevelIndicatorStyle
	SetLevelIndicatorStyle(value LevelIndicatorStyle)
	RatingImage() Image
	SetRatingImage(value IImage)
	DrawsTieredCapacityLevels() bool
	SetDrawsTieredCapacityLevels(value bool)
	FillColor() Color
	SetFillColor(value IColor)
	WarningFillColor() Color
	SetWarningFillColor(value IColor)
	CriticalFillColor() Color
	SetCriticalFillColor(value IColor)
	RatingPlaceholderImage() Image
	SetRatingPlaceholderImage(value IImage)
	PlaceholderVisibility() LevelIndicatorPlaceholderVisibility
	SetPlaceholderVisibility(value LevelIndicatorPlaceholderVisibility)
	IsEditable() bool
	SetEditable(value bool)
}

type LevelIndicator struct {
	Control
}

func MakeLevelIndicator(ptr unsafe.Pointer) LevelIndicator {
	return LevelIndicator{
		Control: MakeControl(ptr),
	}
}

func (l_ LevelIndicator) InitWithFrame(frameRect foundation.Rect) LevelIndicator {
	rv := objc.CallMethod[LevelIndicator](l_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func LevelIndicator_InitWithFrame(frameRect foundation.Rect) LevelIndicator {
	return LevelIndicatorClass.Alloc().InitWithFrame(frameRect)
}

func (l_ LevelIndicator) Init() LevelIndicator {
	rv := objc.CallMethod[LevelIndicator](l_, objc.GetSelector("init"))
	return rv
}

func LevelIndicator_Init() LevelIndicator {
	return LevelIndicatorClass.Alloc().Init()
}

func (lc _LevelIndicatorClass) Alloc() LevelIndicator {
	rv := objc.CallMethod[LevelIndicator](lc, objc.GetSelector("alloc"))
	return rv
}

func LevelIndicator_Alloc() LevelIndicator {
	return LevelIndicatorClass.Alloc()
}

func (lc _LevelIndicatorClass) New() LevelIndicator {
	rv := objc.CallMethod[LevelIndicator](lc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewLevelIndicator() LevelIndicator {
	return LevelIndicatorClass.New()
}

func LevelIndicator_New() LevelIndicator {
	return LevelIndicatorClass.New()
}

func (l_ LevelIndicator) TickMarkValueAtIndex(index int) float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("tickMarkValueAtIndex:"), index)
	return rv
}

func (l_ LevelIndicator) RectOfTickMarkAtIndex(index int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("rectOfTickMarkAtIndex:"), index)
	return rv
}

func (l_ LevelIndicator) MinValue() float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("minValue"))
	return rv
}

func (l_ LevelIndicator) SetMinValue(value float64) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setMinValue:"), value)
}

func (l_ LevelIndicator) MaxValue() float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("maxValue"))
	return rv
}

func (l_ LevelIndicator) SetMaxValue(value float64) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setMaxValue:"), value)
}

func (l_ LevelIndicator) WarningValue() float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("warningValue"))
	return rv
}

func (l_ LevelIndicator) SetWarningValue(value float64) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setWarningValue:"), value)
}

func (l_ LevelIndicator) CriticalValue() float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("criticalValue"))
	return rv
}

func (l_ LevelIndicator) SetCriticalValue(value float64) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setCriticalValue:"), value)
}

func (l_ LevelIndicator) TickMarkPosition() TickMarkPosition {
	rv := objc.CallMethod[TickMarkPosition](l_, objc.GetSelector("tickMarkPosition"))
	return rv
}

func (l_ LevelIndicator) SetTickMarkPosition(value TickMarkPosition) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setTickMarkPosition:"), value)
}

func (l_ LevelIndicator) NumberOfTickMarks() int {
	rv := objc.CallMethod[int](l_, objc.GetSelector("numberOfTickMarks"))
	return rv
}

func (l_ LevelIndicator) SetNumberOfTickMarks(value int) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setNumberOfTickMarks:"), value)
}

func (l_ LevelIndicator) NumberOfMajorTickMarks() int {
	rv := objc.CallMethod[int](l_, objc.GetSelector("numberOfMajorTickMarks"))
	return rv
}

func (l_ LevelIndicator) SetNumberOfMajorTickMarks(value int) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setNumberOfMajorTickMarks:"), value)
}

func (l_ LevelIndicator) LevelIndicatorStyle() LevelIndicatorStyle {
	rv := objc.CallMethod[LevelIndicatorStyle](l_, objc.GetSelector("levelIndicatorStyle"))
	return rv
}

func (l_ LevelIndicator) SetLevelIndicatorStyle(value LevelIndicatorStyle) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setLevelIndicatorStyle:"), value)
}

func (l_ LevelIndicator) RatingImage() Image {
	rv := objc.CallMethod[Image](l_, objc.GetSelector("ratingImage"))
	return rv
}

func (l_ LevelIndicator) SetRatingImage(value IImage) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setRatingImage:"), objc.ExtractPtr(value))
}

func (l_ LevelIndicator) DrawsTieredCapacityLevels() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("drawsTieredCapacityLevels"))
	return rv
}

func (l_ LevelIndicator) SetDrawsTieredCapacityLevels(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setDrawsTieredCapacityLevels:"), value)
}

func (l_ LevelIndicator) FillColor() Color {
	rv := objc.CallMethod[Color](l_, objc.GetSelector("fillColor"))
	return rv
}

func (l_ LevelIndicator) SetFillColor(value IColor) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setFillColor:"), objc.ExtractPtr(value))
}

func (l_ LevelIndicator) WarningFillColor() Color {
	rv := objc.CallMethod[Color](l_, objc.GetSelector("warningFillColor"))
	return rv
}

func (l_ LevelIndicator) SetWarningFillColor(value IColor) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setWarningFillColor:"), objc.ExtractPtr(value))
}

func (l_ LevelIndicator) CriticalFillColor() Color {
	rv := objc.CallMethod[Color](l_, objc.GetSelector("criticalFillColor"))
	return rv
}

func (l_ LevelIndicator) SetCriticalFillColor(value IColor) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setCriticalFillColor:"), objc.ExtractPtr(value))
}

func (l_ LevelIndicator) RatingPlaceholderImage() Image {
	rv := objc.CallMethod[Image](l_, objc.GetSelector("ratingPlaceholderImage"))
	return rv
}

func (l_ LevelIndicator) SetRatingPlaceholderImage(value IImage) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setRatingPlaceholderImage:"), objc.ExtractPtr(value))
}

func (l_ LevelIndicator) PlaceholderVisibility() LevelIndicatorPlaceholderVisibility {
	rv := objc.CallMethod[LevelIndicatorPlaceholderVisibility](l_, objc.GetSelector("placeholderVisibility"))
	return rv
}

func (l_ LevelIndicator) SetPlaceholderVisibility(value LevelIndicatorPlaceholderVisibility) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setPlaceholderVisibility:"), value)
}

func (l_ LevelIndicator) IsEditable() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("isEditable"))
	return rv
}

func (l_ LevelIndicator) SetEditable(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setEditable:"), value)
}
