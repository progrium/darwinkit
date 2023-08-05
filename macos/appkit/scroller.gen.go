// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ScrollerClass = _ScrollerClass{objc.GetClass("NSScroller")}

type _ScrollerClass struct {
	objc.Class
}

type IScroller interface {
	IControl
	SetKnobProportion(proportion float64)
	RectForPart(partCode ScrollerPart) foundation.Rect
	TestPart(point foundation.Point) ScrollerPart
	CheckSpaceForParts()
	DrawKnobSlotInRectHighlight(slotRect foundation.Rect, flag bool)
	DrawKnob()
	TrackKnob(event IEvent)
	UsableParts() UsableScrollerParts
	HitPart() ScrollerPart
	ScrollerStyle() ScrollerStyle
	SetScrollerStyle(value ScrollerStyle)
	KnobStyle() ScrollerKnobStyle
	SetKnobStyle(value ScrollerKnobStyle)
	KnobProportion() float64
}

type Scroller struct {
	Control
}

func MakeScroller(ptr unsafe.Pointer) Scroller {
	return Scroller{
		Control: MakeControl(ptr),
	}
}

func (s_ Scroller) InitWithFrame(frameRect foundation.Rect) Scroller {
	rv := objc.CallMethod[Scroller](s_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func Scroller_InitWithFrame(frameRect foundation.Rect) Scroller {
	return ScrollerClass.Alloc().InitWithFrame(frameRect)
}

func (s_ Scroller) Init() Scroller {
	rv := objc.CallMethod[Scroller](s_, objc.GetSelector("init"))
	return rv
}

func Scroller_Init() Scroller {
	return ScrollerClass.Alloc().Init()
}

func (sc _ScrollerClass) Alloc() Scroller {
	rv := objc.CallMethod[Scroller](sc, objc.GetSelector("alloc"))
	return rv
}

func Scroller_Alloc() Scroller {
	return ScrollerClass.Alloc()
}

func (sc _ScrollerClass) New() Scroller {
	rv := objc.CallMethod[Scroller](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewScroller() Scroller {
	return ScrollerClass.New()
}

func Scroller_New() Scroller {
	return ScrollerClass.New()
}

func (sc _ScrollerClass) ScrollerWidthForControlSizeScrollerStyle(controlSize ControlSize, scrollerStyle ScrollerStyle) float64 {
	rv := objc.CallMethod[float64](sc, objc.GetSelector("scrollerWidthForControlSize:scrollerStyle:"), controlSize, scrollerStyle)
	return rv
}

func Scroller_ScrollerWidthForControlSizeScrollerStyle(controlSize ControlSize, scrollerStyle ScrollerStyle) float64 {
	return ScrollerClass.ScrollerWidthForControlSizeScrollerStyle(controlSize, scrollerStyle)
}

func (s_ Scroller) SetKnobProportion(proportion float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setKnobProportion:"), proportion)
}

func (s_ Scroller) RectForPart(partCode ScrollerPart) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("rectForPart:"), partCode)
	return rv
}

func (s_ Scroller) TestPart(point foundation.Point) ScrollerPart {
	rv := objc.CallMethod[ScrollerPart](s_, objc.GetSelector("testPart:"), point)
	return rv
}

func (s_ Scroller) CheckSpaceForParts() {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("checkSpaceForParts"))
}

func (s_ Scroller) DrawKnobSlotInRectHighlight(slotRect foundation.Rect, flag bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("drawKnobSlotInRect:highlight:"), slotRect, flag)
}

func (s_ Scroller) DrawKnob() {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("drawKnob"))
}

func (s_ Scroller) TrackKnob(event IEvent) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("trackKnob:"), objc.ExtractPtr(event))
}

func (s_ Scroller) UsableParts() UsableScrollerParts {
	rv := objc.CallMethod[UsableScrollerParts](s_, objc.GetSelector("usableParts"))
	return rv
}

func (s_ Scroller) HitPart() ScrollerPart {
	rv := objc.CallMethod[ScrollerPart](s_, objc.GetSelector("hitPart"))
	return rv
}

func (sc _ScrollerClass) PreferredScrollerStyle() ScrollerStyle {
	rv := objc.CallMethod[ScrollerStyle](sc, objc.GetSelector("preferredScrollerStyle"))
	return rv
}

func Scroller_PreferredScrollerStyle() ScrollerStyle {
	return ScrollerClass.PreferredScrollerStyle()
}

func (s_ Scroller) ScrollerStyle() ScrollerStyle {
	rv := objc.CallMethod[ScrollerStyle](s_, objc.GetSelector("scrollerStyle"))
	return rv
}

func (s_ Scroller) SetScrollerStyle(value ScrollerStyle) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setScrollerStyle:"), value)
}

func (s_ Scroller) KnobStyle() ScrollerKnobStyle {
	rv := objc.CallMethod[ScrollerKnobStyle](s_, objc.GetSelector("knobStyle"))
	return rv
}

func (s_ Scroller) SetKnobStyle(value ScrollerKnobStyle) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setKnobStyle:"), value)
}

func (s_ Scroller) KnobProportion() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("knobProportion"))
	return rv
}

func (sc _ScrollerClass) IsCompatibleWithOverlayScrollers() bool {
	rv := objc.CallMethod[bool](sc, objc.GetSelector("isCompatibleWithOverlayScrollers"))
	return rv
}

func Scroller_IsCompatibleWithOverlayScrollers() bool {
	return ScrollerClass.IsCompatibleWithOverlayScrollers()
}
