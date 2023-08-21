// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [StackView] class.
var StackViewClass = _StackViewClass{objc.GetClass("NSStackView")}

type _StackViewClass struct {
	objc.Class
}

// An interface definition for the [StackView] class.
type IStackView interface {
	IView
	RemoveView(view IView)
	SetClippingResistancePriorityForOrientation(clippingResistancePriority LayoutPriority, orientation LayoutConstraintOrientation)
	SetVisibilityPriorityForView(priority StackViewVisibilityPriority, view IView)
	SetHuggingPriorityForOrientation(huggingPriority LayoutPriority, orientation LayoutConstraintOrientation)
	VisibilityPriorityForView(view IView) StackViewVisibilityPriority
	CustomSpacingAfterView(view IView) float64
	AddViewInGravity(view IView, gravity StackViewGravity)
	ViewsInGravity(gravity StackViewGravity) []View
	AddArrangedSubview(view IView)
	InsertViewAtIndexInGravity(view IView, index uint, gravity StackViewGravity)
	InsertArrangedSubviewAtIndex(view IView, index int)
	HuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	ClippingResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	SetCustomSpacingAfterView(spacing float64, view IView)
	RemoveArrangedSubview(view IView)
	SetViewsInGravity(views []IView, gravity StackViewGravity)
	Distribution() StackViewDistribution
	SetDistribution(value StackViewDistribution)
	Views() []View
	DetachedViews() []View
	EdgeInsets() foundation.EdgeInsets
	SetEdgeInsets(value foundation.EdgeInsets)
	Alignment() LayoutAttribute
	SetAlignment(value LayoutAttribute)
	Delegate() StackViewDelegateWrapper
	SetDelegate(value PStackViewDelegate)
	SetDelegateObject(valueObject objc.IObject)
	ArrangedSubviews() []View
	Spacing() float64
	SetSpacing(value float64)
	DetachesHiddenViews() bool
	SetDetachesHiddenViews(value bool)
	Orientation() UserInterfaceLayoutOrientation
	SetOrientation(value UserInterfaceLayoutOrientation)
}

// A view that arranges an array of views horizontally or vertically and updates their placement and sizing when the window size changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview?language=objc
type StackView struct {
	View
}

func StackViewFrom(ptr unsafe.Pointer) StackView {
	return StackView{
		View: ViewFrom(ptr),
	}
}

func (sc _StackViewClass) StackViewWithViews(views []IView) StackView {
	rv := objc.Call[StackView](sc, objc.Sel("stackViewWithViews:"), views)
	return rv
}

// Creates and returns a stack view with a specified array of views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488929-stackviewwithviews?language=objc
func StackView_StackViewWithViews(views []IView) StackView {
	return StackViewClass.StackViewWithViews(views)
}

func (sc _StackViewClass) Alloc() StackView {
	rv := objc.Call[StackView](sc, objc.Sel("alloc"))
	return rv
}

func StackView_Alloc() StackView {
	return StackViewClass.Alloc()
}

func (sc _StackViewClass) New() StackView {
	rv := objc.Call[StackView](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStackView() StackView {
	return StackViewClass.New()
}

func (s_ StackView) Init() StackView {
	rv := objc.Call[StackView](s_, objc.Sel("init"))
	return rv
}

func (s_ StackView) InitWithFrame(frameRect foundation.Rect) StackView {
	rv := objc.Call[StackView](s_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewStackViewWithFrame(frameRect foundation.Rect) StackView {
	instance := StackViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Removes a specified view from the stack view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488916-removeview?language=objc
func (s_ StackView) RemoveView(view IView) {
	objc.Call[objc.Void](s_, objc.Sel("removeView:"), objc.Ptr(view))
}

// Sets the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488880-setclippingresistancepriority?language=objc
func (s_ StackView) SetClippingResistancePriorityForOrientation(clippingResistancePriority LayoutPriority, orientation LayoutConstraintOrientation) {
	objc.Call[objc.Void](s_, objc.Sel("setClippingResistancePriority:forOrientation:"), clippingResistancePriority, orientation)
}

// Sets the Auto Layout priority for a view to remain attached to the stack view when Auto Layout reduces the stack view’s size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488890-setvisibilitypriority?language=objc
func (s_ StackView) SetVisibilityPriorityForView(priority StackViewVisibilityPriority, view IView) {
	objc.Call[objc.Void](s_, objc.Sel("setVisibilityPriority:forView:"), priority, objc.Ptr(view))
}

// Sets the Auto Layout priority for the stack view to minimize its size, for a specified user interface axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488904-sethuggingpriority?language=objc
func (s_ StackView) SetHuggingPriorityForOrientation(huggingPriority LayoutPriority, orientation LayoutConstraintOrientation) {
	objc.Call[objc.Void](s_, objc.Sel("setHuggingPriority:forOrientation:"), huggingPriority, orientation)
}

// Returns the visibility priority for a specified view in the stack view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488934-visibilitypriorityforview?language=objc
func (s_ StackView) VisibilityPriorityForView(view IView) StackViewVisibilityPriority {
	rv := objc.Call[StackViewVisibilityPriority](s_, objc.Sel("visibilityPriorityForView:"), objc.Ptr(view))
	return rv
}

// Returns the custom spacing, in points, between a specified view in the stack view and the view that follows it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488888-customspacingafterview?language=objc
func (s_ StackView) CustomSpacingAfterView(view IView) float64 {
	rv := objc.Call[float64](s_, objc.Sel("customSpacingAfterView:"), objc.Ptr(view))
	return rv
}

// Adds a view to the end of the stack view gravity area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488897-addview?language=objc
func (s_ StackView) AddViewInGravity(view IView, gravity StackViewGravity) {
	objc.Call[objc.Void](s_, objc.Sel("addView:inGravity:"), objc.Ptr(view), gravity)
}

// Returns the array of views in the specified gravity area in the stack view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488876-viewsingravity?language=objc
func (s_ StackView) ViewsInGravity(gravity StackViewGravity) []View {
	rv := objc.Call[[]View](s_, objc.Sel("viewsInGravity:"), gravity)
	return rv
}

// Adds the specified view to the end of the arranged subviews list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488899-addarrangedsubview?language=objc
func (s_ StackView) AddArrangedSubview(view IView) {
	objc.Call[objc.Void](s_, objc.Sel("addArrangedSubview:"), objc.Ptr(view))
}

// Adds a view to a stack view gravity area at a specified index position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488933-insertview?language=objc
func (s_ StackView) InsertViewAtIndexInGravity(view IView, index uint, gravity StackViewGravity) {
	objc.Call[objc.Void](s_, objc.Sel("insertView:atIndex:inGravity:"), objc.Ptr(view), index, gravity)
}

// Adds the provided view to the array of arranged subviews at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488896-insertarrangedsubview?language=objc
func (s_ StackView) InsertArrangedSubviewAtIndex(view IView, index int) {
	objc.Call[objc.Void](s_, objc.Sel("insertArrangedSubview:atIndex:"), objc.Ptr(view), index)
}

// Returns the Auto Layout priority for the stack view to minimize its size to fit its contained views as closely as possible, for a specified user interface axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488912-huggingpriorityfororientation?language=objc
func (s_ StackView) HuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := objc.Call[LayoutPriority](s_, objc.Sel("huggingPriorityForOrientation:"), orientation)
	return rv
}

// Returns the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488936-clippingresistancepriorityforori?language=objc
func (s_ StackView) ClippingResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := objc.Call[LayoutPriority](s_, objc.Sel("clippingResistancePriorityForOrientation:"), orientation)
	return rv
}

// Specifies the custom spacing, in points, between a specified view and the view that follows it in the stack view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488874-setcustomspacing?language=objc
func (s_ StackView) SetCustomSpacingAfterView(spacing float64, view IView) {
	objc.Call[objc.Void](s_, objc.Sel("setCustomSpacing:afterView:"), spacing, objc.Ptr(view))
}

// Removes the provided view from the stack’s array of arranged subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488925-removearrangedsubview?language=objc
func (s_ StackView) RemoveArrangedSubview(view IView) {
	objc.Call[objc.Void](s_, objc.Sel("removeArrangedSubview:"), objc.Ptr(view))
}

// Specifies an array of views for a specified gravity area in the stack view, replacing any previous views in that area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488883-setviews?language=objc
func (s_ StackView) SetViewsInGravity(views []IView, gravity StackViewGravity) {
	objc.Call[objc.Void](s_, objc.Sel("setViews:inGravity:"), views, gravity)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488882-distribution?language=objc
func (s_ StackView) Distribution() StackViewDistribution {
	rv := objc.Call[StackViewDistribution](s_, objc.Sel("distribution"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488882-distribution?language=objc
func (s_ StackView) SetDistribution(value StackViewDistribution) {
	objc.Call[objc.Void](s_, objc.Sel("setDistribution:"), value)
}

// The array of views owned by the stack view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488914-views?language=objc
func (s_ StackView) Views() []View {
	rv := objc.Call[[]View](s_, objc.Sel("views"))
	return rv
}

// An array that contains the detached views from all the stack view’s gravity areas. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488952-detachedviews?language=objc
func (s_ StackView) DetachedViews() []View {
	rv := objc.Call[[]View](s_, objc.Sel("detachedViews"))
	return rv
}

// The geometric padding, in points, inside the stack view, surrounding its views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488931-edgeinsets?language=objc
func (s_ StackView) EdgeInsets() foundation.EdgeInsets {
	rv := objc.Call[foundation.EdgeInsets](s_, objc.Sel("edgeInsets"))
	return rv
}

// The geometric padding, in points, inside the stack view, surrounding its views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488931-edgeinsets?language=objc
func (s_ StackView) SetEdgeInsets(value foundation.EdgeInsets) {
	objc.Call[objc.Void](s_, objc.Sel("setEdgeInsets:"), value)
}

// The view alignment within the stack view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488906-alignment?language=objc
func (s_ StackView) Alignment() LayoutAttribute {
	rv := objc.Call[LayoutAttribute](s_, objc.Sel("alignment"))
	return rv
}

// The view alignment within the stack view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488906-alignment?language=objc
func (s_ StackView) SetAlignment(value LayoutAttribute) {
	objc.Call[objc.Void](s_, objc.Sel("setAlignment:"), value)
}

// The delegate object for the stack view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488946-delegate?language=objc
func (s_ StackView) Delegate() StackViewDelegateWrapper {
	rv := objc.Call[StackViewDelegateWrapper](s_, objc.Sel("delegate"))
	return rv
}

// The delegate object for the stack view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488946-delegate?language=objc
func (s_ StackView) SetDelegate(value PStackViewDelegate) {
	po0 := objc.WrapAsProtocol("NSStackViewDelegate", value)
	objc.SetAssociatedObject(s_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), po0)
}

// The delegate object for the stack view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488946-delegate?language=objc
func (s_ StackView) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The array of views arranged by the stack view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488894-arrangedsubviews?language=objc
func (s_ StackView) ArrangedSubviews() []View {
	rv := objc.Call[[]View](s_, objc.Sel("arrangedSubviews"))
	return rv
}

// The minimum spacing, in points, between adjacent views in the stack view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488945-spacing?language=objc
func (s_ StackView) Spacing() float64 {
	rv := objc.Call[float64](s_, objc.Sel("spacing"))
	return rv
}

// The minimum spacing, in points, between adjacent views in the stack view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488945-spacing?language=objc
func (s_ StackView) SetSpacing(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setSpacing:"), value)
}

// A Boolean value that indicates whether the stack view removes hidden views from its view hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488940-detacheshiddenviews?language=objc
func (s_ StackView) DetachesHiddenViews() bool {
	rv := objc.Call[bool](s_, objc.Sel("detachesHiddenViews"))
	return rv
}

// A Boolean value that indicates whether the stack view removes hidden views from its view hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488940-detacheshiddenviews?language=objc
func (s_ StackView) SetDetachesHiddenViews(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setDetachesHiddenViews:"), value)
}

// The horizontal or vertical layout direction of the stack view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488950-orientation?language=objc
func (s_ StackView) Orientation() UserInterfaceLayoutOrientation {
	rv := objc.Call[UserInterfaceLayoutOrientation](s_, objc.Sel("orientation"))
	return rv
}

// The horizontal or vertical layout direction of the stack view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackview/1488950-orientation?language=objc
func (s_ StackView) SetOrientation(value UserInterfaceLayoutOrientation) {
	objc.Call[objc.Void](s_, objc.Sel("setOrientation:"), value)
}
