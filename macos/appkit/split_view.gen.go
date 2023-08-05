// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var SplitViewClass = _SplitViewClass{objc.GetClass("NSSplitView")}

type _SplitViewClass struct {
	objc.Class
}

type ISplitView interface {
	IView
	AddArrangedSubview(view IView)
	InsertArrangedSubviewAtIndex(view IView, index int)
	RemoveArrangedSubview(view IView)
	AdjustSubviews()
	IsSubviewCollapsed(subview IView) bool
	HoldingPriorityForSubviewAtIndex(subviewIndex int) LayoutPriority
	SetHoldingPriorityForSubviewAtIndex(priority LayoutPriority, subviewIndex int)
	DrawDividerInRect(rect foundation.Rect)
	MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64
	MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64
	SetPositionOfDividerAtIndex(position float64, dividerIndex int)
	Delegate() SplitViewDelegateWrapper
	SetDelegate(value ISplitViewDelegate)
	SetDelegate0(value objc.IObject)
	ArrangesAllSubviews() bool
	SetArrangesAllSubviews(value bool)
	ArrangedSubviews() []View
	IsVertical() bool
	SetVertical(value bool)
	DividerStyle() SplitViewDividerStyle
	SetDividerStyle(value SplitViewDividerStyle)
	DividerColor() Color
	DividerThickness() float64
	AutosaveName() SplitViewAutosaveName
	SetAutosaveName(value SplitViewAutosaveName)
}

type SplitView struct {
	View
}

func MakeSplitView(ptr unsafe.Pointer) SplitView {
	return SplitView{
		View: MakeView(ptr),
	}
}

func (s_ SplitView) InitWithFrame(frameRect foundation.Rect) SplitView {
	rv := objc.CallMethod[SplitView](s_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func SplitView_InitWithFrame(frameRect foundation.Rect) SplitView {
	return SplitViewClass.Alloc().InitWithFrame(frameRect)
}

func (s_ SplitView) Init() SplitView {
	rv := objc.CallMethod[SplitView](s_, objc.GetSelector("init"))
	return rv
}

func SplitView_Init() SplitView {
	return SplitViewClass.Alloc().Init()
}

func (sc _SplitViewClass) Alloc() SplitView {
	rv := objc.CallMethod[SplitView](sc, objc.GetSelector("alloc"))
	return rv
}

func SplitView_Alloc() SplitView {
	return SplitViewClass.Alloc()
}

func (sc _SplitViewClass) New() SplitView {
	rv := objc.CallMethod[SplitView](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewSplitView() SplitView {
	return SplitViewClass.New()
}

func SplitView_New() SplitView {
	return SplitViewClass.New()
}

func (s_ SplitView) AddArrangedSubview(view IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("addArrangedSubview:"), objc.ExtractPtr(view))
}

func (s_ SplitView) InsertArrangedSubviewAtIndex(view IView, index int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("insertArrangedSubview:atIndex:"), objc.ExtractPtr(view), index)
}

func (s_ SplitView) RemoveArrangedSubview(view IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("removeArrangedSubview:"), objc.ExtractPtr(view))
}

func (s_ SplitView) AdjustSubviews() {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("adjustSubviews"))
}

func (s_ SplitView) IsSubviewCollapsed(subview IView) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("isSubviewCollapsed:"), objc.ExtractPtr(subview))
	return rv
}

func (s_ SplitView) HoldingPriorityForSubviewAtIndex(subviewIndex int) LayoutPriority {
	rv := objc.CallMethod[LayoutPriority](s_, objc.GetSelector("holdingPriorityForSubviewAtIndex:"), subviewIndex)
	return rv
}

func (s_ SplitView) SetHoldingPriorityForSubviewAtIndex(priority LayoutPriority, subviewIndex int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setHoldingPriority:forSubviewAtIndex:"), priority, subviewIndex)
}

func (s_ SplitView) DrawDividerInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("drawDividerInRect:"), rect)
}

func (s_ SplitView) MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("minPossiblePositionOfDividerAtIndex:"), dividerIndex)
	return rv
}

func (s_ SplitView) MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("maxPossiblePositionOfDividerAtIndex:"), dividerIndex)
	return rv
}

func (s_ SplitView) SetPositionOfDividerAtIndex(position float64, dividerIndex int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setPosition:ofDividerAtIndex:"), position, dividerIndex)
}

func (s_ SplitView) Delegate() SplitViewDelegateWrapper {
	rv := objc.CallMethod[SplitViewDelegateWrapper](s_, objc.GetSelector("delegate"))
	return rv
}

func (s_ SplitView) SetDelegate(value ISplitViewDelegate) {
	po := objc.WrapAsProtocol("NSSplitViewDelegate", value)
	objc.SetAssociatedObject(s_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDelegate:"), po)
}

func (s_ SplitView) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (s_ SplitView) ArrangesAllSubviews() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("arrangesAllSubviews"))
	return rv
}

func (s_ SplitView) SetArrangesAllSubviews(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setArrangesAllSubviews:"), value)
}

func (s_ SplitView) ArrangedSubviews() []View {
	rv := objc.CallMethod[[]View](s_, objc.GetSelector("arrangedSubviews"))
	return rv
}

func (s_ SplitView) IsVertical() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("isVertical"))
	return rv
}

func (s_ SplitView) SetVertical(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setVertical:"), value)
}

func (s_ SplitView) DividerStyle() SplitViewDividerStyle {
	rv := objc.CallMethod[SplitViewDividerStyle](s_, objc.GetSelector("dividerStyle"))
	return rv
}

func (s_ SplitView) SetDividerStyle(value SplitViewDividerStyle) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDividerStyle:"), value)
}

func (s_ SplitView) DividerColor() Color {
	rv := objc.CallMethod[Color](s_, objc.GetSelector("dividerColor"))
	return rv
}

func (s_ SplitView) DividerThickness() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("dividerThickness"))
	return rv
}

func (s_ SplitView) AutosaveName() SplitViewAutosaveName {
	rv := objc.CallMethod[SplitViewAutosaveName](s_, objc.GetSelector("autosaveName"))
	return rv
}

func (s_ SplitView) SetAutosaveName(value SplitViewAutosaveName) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAutosaveName:"), value)
}
