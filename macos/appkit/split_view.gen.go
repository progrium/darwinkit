// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SplitView] class.
var SplitViewClass = _SplitViewClass{objc.GetClass("NSSplitView")}

type _SplitViewClass struct {
	objc.Class
}

// An interface definition for the [SplitView] class.
type ISplitView interface {
	IView
	AdjustSubviews()
	IsSubviewCollapsed(subview IView) bool
	SetPositionOfDividerAtIndex(position float64, dividerIndex int)
	MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64
	MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64
	AddArrangedSubview(view IView)
	InsertArrangedSubviewAtIndex(view IView, index int)
	HoldingPriorityForSubviewAtIndex(subviewIndex int) LayoutPriority
	DrawDividerInRect(rect foundation.Rect)
	RemoveArrangedSubview(view IView)
	SetHoldingPriorityForSubviewAtIndex(priority LayoutPriority, subviewIndex int)
	DividerColor() Color
	DividerThickness() float64
	DividerStyle() SplitViewDividerStyle
	SetDividerStyle(value SplitViewDividerStyle)
	Delegate() SplitViewDelegateWrapper
	SetDelegate(value PSplitViewDelegate)
	SetDelegateObject(valueObject objc.IObject)
	ArrangedSubviews() []View
	AutosaveName() SplitViewAutosaveName
	SetAutosaveName(value SplitViewAutosaveName)
	IsVertical() bool
	SetVertical(value bool)
	ArrangesAllSubviews() bool
	SetArrangesAllSubviews(value bool)
}

// A view that arranges two or more views in a linear stack running horizontally or vertically. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview?language=objc
type SplitView struct {
	View
}

func SplitViewFrom(ptr unsafe.Pointer) SplitView {
	return SplitView{
		View: ViewFrom(ptr),
	}
}

func (sc _SplitViewClass) Alloc() SplitView {
	rv := objc.Call[SplitView](sc, objc.Sel("alloc"))
	return rv
}

func SplitView_Alloc() SplitView {
	return SplitViewClass.Alloc()
}

func (sc _SplitViewClass) New() SplitView {
	rv := objc.Call[SplitView](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSplitView() SplitView {
	return SplitViewClass.New()
}

func (s_ SplitView) Init() SplitView {
	rv := objc.Call[SplitView](s_, objc.Sel("init"))
	return rv
}

func (s_ SplitView) InitWithFrame(frameRect foundation.Rect) SplitView {
	rv := objc.Call[SplitView](s_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func SplitView_InitWithFrame(frameRect foundation.Rect) SplitView {
	return SplitViewClass.Alloc().InitWithFrame(frameRect)
}

// Adjusts the sizes of the split view’s subviews so they (plus the dividers) fill the split view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455285-adjustsubviews?language=objc
func (s_ SplitView) AdjustSubviews() {
	objc.Call[objc.Void](s_, objc.Sel("adjustSubviews"))
}

// Returns whether the specified view is in a collapsed state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455311-issubviewcollapsed?language=objc
func (s_ SplitView) IsSubviewCollapsed(subview IView) bool {
	rv := objc.Call[bool](s_, objc.Sel("isSubviewCollapsed:"), objc.Ptr(subview))
	return rv
}

// Updates the location of a divider you specify by index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455316-setposition?language=objc
func (s_ SplitView) SetPositionOfDividerAtIndex(position float64, dividerIndex int) {
	objc.Call[objc.Void](s_, objc.Sel("setPosition:ofDividerAtIndex:"), position, dividerIndex)
}

// Returns the maximum possible position of the divider at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455261-maxpossiblepositionofdivideratin?language=objc
func (s_ SplitView) MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := objc.Call[float64](s_, objc.Sel("maxPossiblePositionOfDividerAtIndex:"), dividerIndex)
	return rv
}

// Returns the minimum possible position of the divider at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455278-minpossiblepositionofdivideratin?language=objc
func (s_ SplitView) MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := objc.Call[float64](s_, objc.Sel("minPossiblePositionOfDividerAtIndex:"), dividerIndex)
	return rv
}

// Adds a view as an arranged split pane. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455295-addarrangedsubview?language=objc
func (s_ SplitView) AddArrangedSubview(view IView) {
	objc.Call[objc.Void](s_, objc.Sel("addArrangedSubview:"), objc.Ptr(view))
}

// Adds a view as an arranged split pane at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455321-insertarrangedsubview?language=objc
func (s_ SplitView) InsertArrangedSubviewAtIndex(view IView, index int) {
	objc.Call[objc.Void](s_, objc.Sel("insertArrangedSubview:atIndex:"), objc.Ptr(view), index)
}

// Returns the priority of the subview’s width or height when resizing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455287-holdingpriorityforsubviewatindex?language=objc
func (s_ SplitView) HoldingPriorityForSubviewAtIndex(subviewIndex int) LayoutPriority {
	rv := objc.Call[LayoutPriority](s_, objc.Sel("holdingPriorityForSubviewAtIndex:"), subviewIndex)
	return rv
}

// Draws a divider between two of the split view’s subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455297-drawdividerinrect?language=objc
func (s_ SplitView) DrawDividerInRect(rect foundation.Rect) {
	objc.Call[objc.Void](s_, objc.Sel("drawDividerInRect:"), rect)
}

// Removes a view as an arranged split pane. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455271-removearrangedsubview?language=objc
func (s_ SplitView) RemoveArrangedSubview(view IView) {
	objc.Call[objc.Void](s_, objc.Sel("removeArrangedSubview:"), objc.Ptr(view))
}

// Sets the priority for split view subviews to maintain their width or height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455320-setholdingpriority?language=objc
func (s_ SplitView) SetHoldingPriorityForSubviewAtIndex(priority LayoutPriority, subviewIndex int) {
	objc.Call[objc.Void](s_, objc.Sel("setHoldingPriority:forSubviewAtIndex:"), priority, subviewIndex)
}

// The color of the dividers that the split view draws between subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455267-dividercolor?language=objc
func (s_ SplitView) DividerColor() Color {
	rv := objc.Call[Color](s_, objc.Sel("dividerColor"))
	return rv
}

// The thickness of the dividers for the split view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455257-dividerthickness?language=objc
func (s_ SplitView) DividerThickness() float64 {
	rv := objc.Call[float64](s_, objc.Sel("dividerThickness"))
	return rv
}

// The style of divider between views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455291-dividerstyle?language=objc
func (s_ SplitView) DividerStyle() SplitViewDividerStyle {
	rv := objc.Call[SplitViewDividerStyle](s_, objc.Sel("dividerStyle"))
	return rv
}

// The style of divider between views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455291-dividerstyle?language=objc
func (s_ SplitView) SetDividerStyle(value SplitViewDividerStyle) {
	objc.Call[objc.Void](s_, objc.Sel("setDividerStyle:"), value)
}

// The split view’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455306-delegate?language=objc
func (s_ SplitView) Delegate() SplitViewDelegateWrapper {
	rv := objc.Call[SplitViewDelegateWrapper](s_, objc.Sel("delegate"))
	return rv
}

// The split view’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455306-delegate?language=objc
func (s_ SplitView) SetDelegate(value PSplitViewDelegate) {
	po0 := objc.WrapAsProtocol("NSSplitViewDelegate", value)
	objc.SetAssociatedObject(s_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), po0)
}

// The split view’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455306-delegate?language=objc
func (s_ SplitView) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The array of views that the split view arranges as its split panes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455253-arrangedsubviews?language=objc
func (s_ SplitView) ArrangedSubviews() []View {
	rv := objc.Call[[]View](s_, objc.Sel("arrangedSubviews"))
	return rv
}

// The name to use when the system automatically saves the split view’s divider configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455319-autosavename?language=objc
func (s_ SplitView) AutosaveName() SplitViewAutosaveName {
	rv := objc.Call[SplitViewAutosaveName](s_, objc.Sel("autosaveName"))
	return rv
}

// The name to use when the system automatically saves the split view’s divider configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455319-autosavename?language=objc
func (s_ SplitView) SetAutosaveName(value SplitViewAutosaveName) {
	objc.Call[objc.Void](s_, objc.Sel("setAutosaveName:"), value)
}

// A Boolean value that determines the geometric orientation of the split view's dividers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455318-vertical?language=objc
func (s_ SplitView) IsVertical() bool {
	rv := objc.Call[bool](s_, objc.Sel("isVertical"))
	return rv
}

// A Boolean value that determines the geometric orientation of the split view's dividers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455318-vertical?language=objc
func (s_ SplitView) SetVertical(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setVertical:"), value)
}

// A Boolean value that determines whether the split view arranges all of its subviews as split panes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455307-arrangesallsubviews?language=objc
func (s_ SplitView) ArrangesAllSubviews() bool {
	rv := objc.Call[bool](s_, objc.Sel("arrangesAllSubviews"))
	return rv
}

// A Boolean value that determines whether the split view arranges all of its subviews as split panes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/1455307-arrangesallsubviews?language=objc
func (s_ SplitView) SetArrangesAllSubviews(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setArrangesAllSubviews:"), value)
}
