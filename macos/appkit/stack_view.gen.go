// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var StackViewClass = _StackViewClass{objc.GetClass("NSStackView")}

type _StackViewClass struct {
	objc.Class
}

type IStackView interface {
	IView
	AddViewInGravity(view IView, gravity StackViewGravity)
	InsertViewAtIndexInGravity(view IView, index uint, gravity StackViewGravity)
	SetViewsInGravity(views []IView, gravity StackViewGravity)
	RemoveView(view IView)
	AddArrangedSubview(view IView)
	InsertArrangedSubviewAtIndex(view IView, index int)
	RemoveArrangedSubview(view IView)
	ViewsInGravity(gravity StackViewGravity) []View
	ClippingResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	HuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	CustomSpacingAfterView(view IView) float64
	SetCustomSpacingAfterView(spacing float64, view IView)
	VisibilityPriorityForView(view IView) StackViewVisibilityPriority
	SetVisibilityPriorityForView(priority StackViewVisibilityPriority, view IView)
	SetClippingResistancePriorityForOrientation(clippingResistancePriority LayoutPriority, orientation LayoutConstraintOrientation)
	SetHuggingPriorityForOrientation(huggingPriority LayoutPriority, orientation LayoutConstraintOrientation)
	Delegate() StackViewDelegateWrapper
	SetDelegate(value IStackViewDelegate)
	SetDelegate0(value objc.IObject)
	ArrangedSubviews() []View
	Views() []View
	DetachedViews() []View
	Orientation() UserInterfaceLayoutOrientation
	SetOrientation(value UserInterfaceLayoutOrientation)
	Alignment() LayoutAttribute
	SetAlignment(value LayoutAttribute)
	Spacing() float64
	SetSpacing(value float64)
	EdgeInsets() foundation.EdgeInsets
	SetEdgeInsets(value foundation.EdgeInsets)
	Distribution() StackViewDistribution
	SetDistribution(value StackViewDistribution)
	DetachesHiddenViews() bool
	SetDetachesHiddenViews(value bool)
}

type StackView struct {
	View
}

func MakeStackView(ptr unsafe.Pointer) StackView {
	return StackView{
		View: MakeView(ptr),
	}
}

func (sc _StackViewClass) StackViewWithViews(views []IView) StackView {
	rv := objc.CallMethod[StackView](sc, objc.GetSelector("stackViewWithViews:"), views)
	return rv
}

func StackView_StackViewWithViews(views []IView) StackView {
	return StackViewClass.StackViewWithViews(views)
}

func (s_ StackView) InitWithFrame(frameRect foundation.Rect) StackView {
	rv := objc.CallMethod[StackView](s_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func StackView_InitWithFrame(frameRect foundation.Rect) StackView {
	return StackViewClass.Alloc().InitWithFrame(frameRect)
}

func (s_ StackView) Init() StackView {
	rv := objc.CallMethod[StackView](s_, objc.GetSelector("init"))
	return rv
}

func StackView_Init() StackView {
	return StackViewClass.Alloc().Init()
}

func (sc _StackViewClass) Alloc() StackView {
	rv := objc.CallMethod[StackView](sc, objc.GetSelector("alloc"))
	return rv
}

func StackView_Alloc() StackView {
	return StackViewClass.Alloc()
}

func (sc _StackViewClass) New() StackView {
	rv := objc.CallMethod[StackView](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewStackView() StackView {
	return StackViewClass.New()
}

func StackView_New() StackView {
	return StackViewClass.New()
}

func (s_ StackView) AddViewInGravity(view IView, gravity StackViewGravity) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("addView:inGravity:"), objc.ExtractPtr(view), gravity)
}

func (s_ StackView) InsertViewAtIndexInGravity(view IView, index uint, gravity StackViewGravity) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("insertView:atIndex:inGravity:"), objc.ExtractPtr(view), index, gravity)
}

func (s_ StackView) SetViewsInGravity(views []IView, gravity StackViewGravity) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setViews:inGravity:"), views, gravity)
}

func (s_ StackView) RemoveView(view IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("removeView:"), objc.ExtractPtr(view))
}

func (s_ StackView) AddArrangedSubview(view IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("addArrangedSubview:"), objc.ExtractPtr(view))
}

func (s_ StackView) InsertArrangedSubviewAtIndex(view IView, index int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("insertArrangedSubview:atIndex:"), objc.ExtractPtr(view), index)
}

func (s_ StackView) RemoveArrangedSubview(view IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("removeArrangedSubview:"), objc.ExtractPtr(view))
}

func (s_ StackView) ViewsInGravity(gravity StackViewGravity) []View {
	rv := objc.CallMethod[[]View](s_, objc.GetSelector("viewsInGravity:"), gravity)
	return rv
}

func (s_ StackView) ClippingResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := objc.CallMethod[LayoutPriority](s_, objc.GetSelector("clippingResistancePriorityForOrientation:"), orientation)
	return rv
}

func (s_ StackView) HuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := objc.CallMethod[LayoutPriority](s_, objc.GetSelector("huggingPriorityForOrientation:"), orientation)
	return rv
}

func (s_ StackView) CustomSpacingAfterView(view IView) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("customSpacingAfterView:"), objc.ExtractPtr(view))
	return rv
}

func (s_ StackView) SetCustomSpacingAfterView(spacing float64, view IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setCustomSpacing:afterView:"), spacing, objc.ExtractPtr(view))
}

func (s_ StackView) VisibilityPriorityForView(view IView) StackViewVisibilityPriority {
	rv := objc.CallMethod[StackViewVisibilityPriority](s_, objc.GetSelector("visibilityPriorityForView:"), objc.ExtractPtr(view))
	return rv
}

func (s_ StackView) SetVisibilityPriorityForView(priority StackViewVisibilityPriority, view IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setVisibilityPriority:forView:"), priority, objc.ExtractPtr(view))
}

func (s_ StackView) SetClippingResistancePriorityForOrientation(clippingResistancePriority LayoutPriority, orientation LayoutConstraintOrientation) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setClippingResistancePriority:forOrientation:"), clippingResistancePriority, orientation)
}

func (s_ StackView) SetHuggingPriorityForOrientation(huggingPriority LayoutPriority, orientation LayoutConstraintOrientation) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setHuggingPriority:forOrientation:"), huggingPriority, orientation)
}

func (s_ StackView) Delegate() StackViewDelegateWrapper {
	rv := objc.CallMethod[StackViewDelegateWrapper](s_, objc.GetSelector("delegate"))
	return rv
}

func (s_ StackView) SetDelegate(value IStackViewDelegate) {
	po := objc.WrapAsProtocol("NSStackViewDelegate", value)
	objc.SetAssociatedObject(s_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDelegate:"), po)
}

func (s_ StackView) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (s_ StackView) ArrangedSubviews() []View {
	rv := objc.CallMethod[[]View](s_, objc.GetSelector("arrangedSubviews"))
	return rv
}

func (s_ StackView) Views() []View {
	rv := objc.CallMethod[[]View](s_, objc.GetSelector("views"))
	return rv
}

func (s_ StackView) DetachedViews() []View {
	rv := objc.CallMethod[[]View](s_, objc.GetSelector("detachedViews"))
	return rv
}

func (s_ StackView) Orientation() UserInterfaceLayoutOrientation {
	rv := objc.CallMethod[UserInterfaceLayoutOrientation](s_, objc.GetSelector("orientation"))
	return rv
}

func (s_ StackView) SetOrientation(value UserInterfaceLayoutOrientation) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setOrientation:"), value)
}

func (s_ StackView) Alignment() LayoutAttribute {
	rv := objc.CallMethod[LayoutAttribute](s_, objc.GetSelector("alignment"))
	return rv
}

func (s_ StackView) SetAlignment(value LayoutAttribute) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAlignment:"), value)
}

func (s_ StackView) Spacing() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("spacing"))
	return rv
}

func (s_ StackView) SetSpacing(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSpacing:"), value)
}

func (s_ StackView) EdgeInsets() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](s_, objc.GetSelector("edgeInsets"))
	return rv
}

func (s_ StackView) SetEdgeInsets(value foundation.EdgeInsets) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setEdgeInsets:"), value)
}

func (s_ StackView) Distribution() StackViewDistribution {
	rv := objc.CallMethod[StackViewDistribution](s_, objc.GetSelector("distribution"))
	return rv
}

func (s_ StackView) SetDistribution(value StackViewDistribution) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDistribution:"), value)
}

func (s_ StackView) DetachesHiddenViews() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("detachesHiddenViews"))
	return rv
}

func (s_ StackView) SetDetachesHiddenViews(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDetachesHiddenViews:"), value)
}
