// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ScrollViewClass = _ScrollViewClass{objc.GetClass("NSScrollView")}

type _ScrollViewClass struct {
	objc.Class
}

type IScrollView interface {
	IView
	AddFloatingSubviewForAxis(view IView, axis EventGestureAxis)
	Tile()
	FlashScrollers()
	MagnifyToFitRect(rect foundation.Rect)
	SetMagnificationCenteredAtPoint(magnification float64, point foundation.Point)
	ContentSize() foundation.Size
	DocumentVisibleRect() foundation.Rect
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	BorderType() BorderType
	SetBorderType(value BorderType)
	DocumentCursor() Cursor
	SetDocumentCursor(value ICursor)
	ContentView() ClipView
	SetContentView(value IClipView)
	DocumentView() View
	SetDocumentView(value IView)
	HorizontalScroller() Scroller
	SetHorizontalScroller(value IScroller)
	HasHorizontalScroller() bool
	SetHasHorizontalScroller(value bool)
	VerticalScroller() Scroller
	SetVerticalScroller(value IScroller)
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	AutohidesScrollers() bool
	SetAutohidesScrollers(value bool)
	HasHorizontalRuler() bool
	SetHasHorizontalRuler(value bool)
	HorizontalRulerView() RulerView
	SetHorizontalRulerView(value IRulerView)
	HasVerticalRuler() bool
	SetHasVerticalRuler(value bool)
	VerticalRulerView() RulerView
	SetVerticalRulerView(value IRulerView)
	RulersVisible() bool
	SetRulersVisible(value bool)
	AutomaticallyAdjustsContentInsets() bool
	SetAutomaticallyAdjustsContentInsets(value bool)
	ContentInsets() foundation.EdgeInsets
	SetContentInsets(value foundation.EdgeInsets)
	ScrollerInsets() foundation.EdgeInsets
	SetScrollerInsets(value foundation.EdgeInsets)
	ScrollerKnobStyle() ScrollerKnobStyle
	SetScrollerKnobStyle(value ScrollerKnobStyle)
	ScrollerStyle() ScrollerStyle
	SetScrollerStyle(value ScrollerStyle)
	LineScroll() float64
	SetLineScroll(value float64)
	HorizontalLineScroll() float64
	SetHorizontalLineScroll(value float64)
	VerticalLineScroll() float64
	SetVerticalLineScroll(value float64)
	PageScroll() float64
	SetPageScroll(value float64)
	HorizontalPageScroll() float64
	SetHorizontalPageScroll(value float64)
	VerticalPageScroll() float64
	SetVerticalPageScroll(value float64)
	ScrollsDynamically() bool
	SetScrollsDynamically(value bool)
	FindBarPosition() ScrollViewFindBarPosition
	SetFindBarPosition(value ScrollViewFindBarPosition)
	UsesPredominantAxisScrolling() bool
	SetUsesPredominantAxisScrolling(value bool)
	HorizontalScrollElasticity() ScrollElasticity
	SetHorizontalScrollElasticity(value ScrollElasticity)
	VerticalScrollElasticity() ScrollElasticity
	SetVerticalScrollElasticity(value ScrollElasticity)
	AllowsMagnification() bool
	SetAllowsMagnification(value bool)
	Magnification() float64
	SetMagnification(value float64)
	MaxMagnification() float64
	SetMaxMagnification(value float64)
	MinMagnification() float64
	SetMinMagnification(value float64)
}

type ScrollView struct {
	View
}

func MakeScrollView(ptr unsafe.Pointer) ScrollView {
	return ScrollView{
		View: MakeView(ptr),
	}
}

func (s_ ScrollView) InitWithFrame(frameRect foundation.Rect) ScrollView {
	rv := objc.CallMethod[ScrollView](s_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func ScrollView_InitWithFrame(frameRect foundation.Rect) ScrollView {
	return ScrollViewClass.Alloc().InitWithFrame(frameRect)
}

func (s_ ScrollView) Init() ScrollView {
	rv := objc.CallMethod[ScrollView](s_, objc.GetSelector("init"))
	return rv
}

func ScrollView_Init() ScrollView {
	return ScrollViewClass.Alloc().Init()
}

func (sc _ScrollViewClass) Alloc() ScrollView {
	rv := objc.CallMethod[ScrollView](sc, objc.GetSelector("alloc"))
	return rv
}

func ScrollView_Alloc() ScrollView {
	return ScrollViewClass.Alloc()
}

func (sc _ScrollViewClass) New() ScrollView {
	rv := objc.CallMethod[ScrollView](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewScrollView() ScrollView {
	return ScrollViewClass.New()
}

func ScrollView_New() ScrollView {
	return ScrollViewClass.New()
}

func (sc _ScrollViewClass) FrameSizeForContentSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(cSize foundation.Size, horizontalScrollerClass objc.IClass, verticalScrollerClass objc.IClass, type_ BorderType, controlSize ControlSize, scrollerStyle ScrollerStyle) foundation.Size {
	rv := objc.CallMethod[foundation.Size](sc, objc.GetSelector("frameSizeForContentSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:"), cSize, objc.ExtractPtr(horizontalScrollerClass), objc.ExtractPtr(verticalScrollerClass), type_, controlSize, scrollerStyle)
	return rv
}

func ScrollView_FrameSizeForContentSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(cSize foundation.Size, horizontalScrollerClass objc.IClass, verticalScrollerClass objc.IClass, type_ BorderType, controlSize ControlSize, scrollerStyle ScrollerStyle) foundation.Size {
	return ScrollViewClass.FrameSizeForContentSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(cSize, horizontalScrollerClass, verticalScrollerClass, type_, controlSize, scrollerStyle)
}

func (sc _ScrollViewClass) ContentSizeForFrameSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(fSize foundation.Size, horizontalScrollerClass objc.IClass, verticalScrollerClass objc.IClass, type_ BorderType, controlSize ControlSize, scrollerStyle ScrollerStyle) foundation.Size {
	rv := objc.CallMethod[foundation.Size](sc, objc.GetSelector("contentSizeForFrameSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:"), fSize, objc.ExtractPtr(horizontalScrollerClass), objc.ExtractPtr(verticalScrollerClass), type_, controlSize, scrollerStyle)
	return rv
}

func ScrollView_ContentSizeForFrameSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(fSize foundation.Size, horizontalScrollerClass objc.IClass, verticalScrollerClass objc.IClass, type_ BorderType, controlSize ControlSize, scrollerStyle ScrollerStyle) foundation.Size {
	return ScrollViewClass.ContentSizeForFrameSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(fSize, horizontalScrollerClass, verticalScrollerClass, type_, controlSize, scrollerStyle)
}

func (s_ ScrollView) AddFloatingSubviewForAxis(view IView, axis EventGestureAxis) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("addFloatingSubview:forAxis:"), objc.ExtractPtr(view), axis)
}

func (s_ ScrollView) Tile() {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("tile"))
}

func (s_ ScrollView) FlashScrollers() {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("flashScrollers"))
}

func (s_ ScrollView) MagnifyToFitRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("magnifyToFitRect:"), rect)
}

func (s_ ScrollView) SetMagnificationCenteredAtPoint(magnification float64, point foundation.Point) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMagnification:centeredAtPoint:"), magnification, point)
}

func (s_ ScrollView) ContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](s_, objc.GetSelector("contentSize"))
	return rv
}

func (s_ ScrollView) DocumentVisibleRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("documentVisibleRect"))
	return rv
}

func (s_ ScrollView) BackgroundColor() Color {
	rv := objc.CallMethod[Color](s_, objc.GetSelector("backgroundColor"))
	return rv
}

func (s_ ScrollView) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (s_ ScrollView) DrawsBackground() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("drawsBackground"))
	return rv
}

func (s_ ScrollView) SetDrawsBackground(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDrawsBackground:"), value)
}

func (s_ ScrollView) BorderType() BorderType {
	rv := objc.CallMethod[BorderType](s_, objc.GetSelector("borderType"))
	return rv
}

func (s_ ScrollView) SetBorderType(value BorderType) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setBorderType:"), value)
}

func (s_ ScrollView) DocumentCursor() Cursor {
	rv := objc.CallMethod[Cursor](s_, objc.GetSelector("documentCursor"))
	return rv
}

func (s_ ScrollView) SetDocumentCursor(value ICursor) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDocumentCursor:"), objc.ExtractPtr(value))
}

func (s_ ScrollView) ContentView() ClipView {
	rv := objc.CallMethod[ClipView](s_, objc.GetSelector("contentView"))
	return rv
}

func (s_ ScrollView) SetContentView(value IClipView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setContentView:"), objc.ExtractPtr(value))
}

func (s_ ScrollView) DocumentView() View {
	rv := objc.CallMethod[View](s_, objc.GetSelector("documentView"))
	return rv
}

func (s_ ScrollView) SetDocumentView(value IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDocumentView:"), objc.ExtractPtr(value))
}

func (s_ ScrollView) HorizontalScroller() Scroller {
	rv := objc.CallMethod[Scroller](s_, objc.GetSelector("horizontalScroller"))
	return rv
}

func (s_ ScrollView) SetHorizontalScroller(value IScroller) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setHorizontalScroller:"), objc.ExtractPtr(value))
}

func (s_ ScrollView) HasHorizontalScroller() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("hasHorizontalScroller"))
	return rv
}

func (s_ ScrollView) SetHasHorizontalScroller(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setHasHorizontalScroller:"), value)
}

func (s_ ScrollView) VerticalScroller() Scroller {
	rv := objc.CallMethod[Scroller](s_, objc.GetSelector("verticalScroller"))
	return rv
}

func (s_ ScrollView) SetVerticalScroller(value IScroller) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setVerticalScroller:"), objc.ExtractPtr(value))
}

func (s_ ScrollView) HasVerticalScroller() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("hasVerticalScroller"))
	return rv
}

func (s_ ScrollView) SetHasVerticalScroller(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setHasVerticalScroller:"), value)
}

func (s_ ScrollView) AutohidesScrollers() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("autohidesScrollers"))
	return rv
}

func (s_ ScrollView) SetAutohidesScrollers(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAutohidesScrollers:"), value)
}

func (sc _ScrollViewClass) RulerViewClass() objc.Class {
	rv := objc.CallMethod[objc.Class](sc, objc.GetSelector("rulerViewClass"))
	return rv
}

func ScrollView_RulerViewClass() objc.Class {
	return ScrollViewClass.RulerViewClass()
}

func (sc _ScrollViewClass) SetRulerViewClass(value objc.IClass) {
	objc.CallMethod[objc.Void](sc, objc.GetSelector("setRulerViewClass:"), objc.ExtractPtr(value))
}

func ScrollView_SetRulerViewClass(value objc.IClass) {
	ScrollViewClass.SetRulerViewClass(value)
}

func (s_ ScrollView) HasHorizontalRuler() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("hasHorizontalRuler"))
	return rv
}

func (s_ ScrollView) SetHasHorizontalRuler(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setHasHorizontalRuler:"), value)
}

func (s_ ScrollView) HorizontalRulerView() RulerView {
	rv := objc.CallMethod[RulerView](s_, objc.GetSelector("horizontalRulerView"))
	return rv
}

func (s_ ScrollView) SetHorizontalRulerView(value IRulerView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setHorizontalRulerView:"), objc.ExtractPtr(value))
}

func (s_ ScrollView) HasVerticalRuler() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("hasVerticalRuler"))
	return rv
}

func (s_ ScrollView) SetHasVerticalRuler(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setHasVerticalRuler:"), value)
}

func (s_ ScrollView) VerticalRulerView() RulerView {
	rv := objc.CallMethod[RulerView](s_, objc.GetSelector("verticalRulerView"))
	return rv
}

func (s_ ScrollView) SetVerticalRulerView(value IRulerView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setVerticalRulerView:"), objc.ExtractPtr(value))
}

func (s_ ScrollView) RulersVisible() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("rulersVisible"))
	return rv
}

func (s_ ScrollView) SetRulersVisible(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setRulersVisible:"), value)
}

func (s_ ScrollView) AutomaticallyAdjustsContentInsets() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("automaticallyAdjustsContentInsets"))
	return rv
}

func (s_ ScrollView) SetAutomaticallyAdjustsContentInsets(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAutomaticallyAdjustsContentInsets:"), value)
}

func (s_ ScrollView) ContentInsets() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](s_, objc.GetSelector("contentInsets"))
	return rv
}

func (s_ ScrollView) SetContentInsets(value foundation.EdgeInsets) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setContentInsets:"), value)
}

func (s_ ScrollView) ScrollerInsets() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](s_, objc.GetSelector("scrollerInsets"))
	return rv
}

func (s_ ScrollView) SetScrollerInsets(value foundation.EdgeInsets) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setScrollerInsets:"), value)
}

func (s_ ScrollView) ScrollerKnobStyle() ScrollerKnobStyle {
	rv := objc.CallMethod[ScrollerKnobStyle](s_, objc.GetSelector("scrollerKnobStyle"))
	return rv
}

func (s_ ScrollView) SetScrollerKnobStyle(value ScrollerKnobStyle) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setScrollerKnobStyle:"), value)
}

func (s_ ScrollView) ScrollerStyle() ScrollerStyle {
	rv := objc.CallMethod[ScrollerStyle](s_, objc.GetSelector("scrollerStyle"))
	return rv
}

func (s_ ScrollView) SetScrollerStyle(value ScrollerStyle) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setScrollerStyle:"), value)
}

func (s_ ScrollView) LineScroll() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("lineScroll"))
	return rv
}

func (s_ ScrollView) SetLineScroll(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setLineScroll:"), value)
}

func (s_ ScrollView) HorizontalLineScroll() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("horizontalLineScroll"))
	return rv
}

func (s_ ScrollView) SetHorizontalLineScroll(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setHorizontalLineScroll:"), value)
}

func (s_ ScrollView) VerticalLineScroll() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("verticalLineScroll"))
	return rv
}

func (s_ ScrollView) SetVerticalLineScroll(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setVerticalLineScroll:"), value)
}

func (s_ ScrollView) PageScroll() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("pageScroll"))
	return rv
}

func (s_ ScrollView) SetPageScroll(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setPageScroll:"), value)
}

func (s_ ScrollView) HorizontalPageScroll() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("horizontalPageScroll"))
	return rv
}

func (s_ ScrollView) SetHorizontalPageScroll(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setHorizontalPageScroll:"), value)
}

func (s_ ScrollView) VerticalPageScroll() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("verticalPageScroll"))
	return rv
}

func (s_ ScrollView) SetVerticalPageScroll(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setVerticalPageScroll:"), value)
}

func (s_ ScrollView) ScrollsDynamically() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("scrollsDynamically"))
	return rv
}

func (s_ ScrollView) SetScrollsDynamically(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setScrollsDynamically:"), value)
}

func (s_ ScrollView) FindBarPosition() ScrollViewFindBarPosition {
	rv := objc.CallMethod[ScrollViewFindBarPosition](s_, objc.GetSelector("findBarPosition"))
	return rv
}

func (s_ ScrollView) SetFindBarPosition(value ScrollViewFindBarPosition) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setFindBarPosition:"), value)
}

func (s_ ScrollView) UsesPredominantAxisScrolling() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("usesPredominantAxisScrolling"))
	return rv
}

func (s_ ScrollView) SetUsesPredominantAxisScrolling(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setUsesPredominantAxisScrolling:"), value)
}

func (s_ ScrollView) HorizontalScrollElasticity() ScrollElasticity {
	rv := objc.CallMethod[ScrollElasticity](s_, objc.GetSelector("horizontalScrollElasticity"))
	return rv
}

func (s_ ScrollView) SetHorizontalScrollElasticity(value ScrollElasticity) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setHorizontalScrollElasticity:"), value)
}

func (s_ ScrollView) VerticalScrollElasticity() ScrollElasticity {
	rv := objc.CallMethod[ScrollElasticity](s_, objc.GetSelector("verticalScrollElasticity"))
	return rv
}

func (s_ ScrollView) SetVerticalScrollElasticity(value ScrollElasticity) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setVerticalScrollElasticity:"), value)
}

func (s_ ScrollView) AllowsMagnification() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("allowsMagnification"))
	return rv
}

func (s_ ScrollView) SetAllowsMagnification(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAllowsMagnification:"), value)
}

func (s_ ScrollView) Magnification() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("magnification"))
	return rv
}

func (s_ ScrollView) SetMagnification(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMagnification:"), value)
}

func (s_ ScrollView) MaxMagnification() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("maxMagnification"))
	return rv
}

func (s_ ScrollView) SetMaxMagnification(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMaxMagnification:"), value)
}

func (s_ ScrollView) MinMagnification() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("minMagnification"))
	return rv
}

func (s_ ScrollView) SetMinMagnification(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMinMagnification:"), value)
}
