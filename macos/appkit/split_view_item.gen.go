// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SplitViewItem] class.
var SplitViewItemClass = _SplitViewItemClass{objc.GetClass("NSSplitViewItem")}

type _SplitViewItemClass struct {
	objc.Class
}

// An interface definition for the [SplitViewItem] class.
type ISplitViewItem interface {
	objc.IObject
	AutomaticMaximumThickness() float64
	SetAutomaticMaximumThickness(value float64)
	PreferredThicknessFraction() float64
	SetPreferredThicknessFraction(value float64)
	MaximumThickness() float64
	SetMaximumThickness(value float64)
	HoldingPriority() LayoutPriority
	SetHoldingPriority(value LayoutPriority)
	CollapseBehavior() SplitViewItemCollapseBehavior
	SetCollapseBehavior(value SplitViewItemCollapseBehavior)
	Behavior() SplitViewItemBehavior
	IsCollapsed() bool
	SetCollapsed(value bool)
	ViewController() ViewController
	SetViewController(value IViewController)
	CanCollapse() bool
	SetCanCollapse(value bool)
	MinimumThickness() float64
	SetMinimumThickness(value float64)
	TitlebarSeparatorStyle() TitlebarSeparatorStyle
	SetTitlebarSeparatorStyle(value TitlebarSeparatorStyle)
	AllowsFullHeightLayout() bool
	SetAllowsFullHeightLayout(value bool)
	IsSpringLoaded() bool
	SetSpringLoaded(value bool)
}

// An item in a split view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem?language=objc
type SplitViewItem struct {
	objc.Object
}

func SplitViewItemFrom(ptr unsafe.Pointer) SplitViewItem {
	return SplitViewItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SplitViewItemClass) SidebarWithViewController(viewController IViewController) SplitViewItem {
	rv := objc.Call[SplitViewItem](sc, objc.Sel("sidebarWithViewController:"), objc.Ptr(viewController))
	return rv
}

// Creates a split view item that represents a sidebar for the specified view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388920-sidebarwithviewcontroller?language=objc
func SplitViewItem_SidebarWithViewController(viewController IViewController) SplitViewItem {
	return SplitViewItemClass.SidebarWithViewController(viewController)
}

func (sc _SplitViewItemClass) SplitViewItemWithViewController(viewController IViewController) SplitViewItem {
	rv := objc.Call[SplitViewItem](sc, objc.Sel("splitViewItemWithViewController:"), objc.Ptr(viewController))
	return rv
}

// Creates a split view item that represents the specified view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388918-splitviewitemwithviewcontroller?language=objc
func SplitViewItem_SplitViewItemWithViewController(viewController IViewController) SplitViewItem {
	return SplitViewItemClass.SplitViewItemWithViewController(viewController)
}

func (sc _SplitViewItemClass) ContentListWithViewController(viewController IViewController) SplitViewItem {
	rv := objc.Call[SplitViewItem](sc, objc.Sel("contentListWithViewController:"), objc.Ptr(viewController))
	return rv
}

// Creates a split view item that represents a content list for the specified view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388878-contentlistwithviewcontroller?language=objc
func SplitViewItem_ContentListWithViewController(viewController IViewController) SplitViewItem {
	return SplitViewItemClass.ContentListWithViewController(viewController)
}

func (sc _SplitViewItemClass) Alloc() SplitViewItem {
	rv := objc.Call[SplitViewItem](sc, objc.Sel("alloc"))
	return rv
}

func SplitViewItem_Alloc() SplitViewItem {
	return SplitViewItemClass.Alloc()
}

func (sc _SplitViewItemClass) New() SplitViewItem {
	rv := objc.Call[SplitViewItem](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSplitViewItem() SplitViewItem {
	return SplitViewItemClass.New()
}

func (s_ SplitViewItem) Init() SplitViewItem {
	rv := objc.Call[SplitViewItem](s_, objc.Sel("init"))
	return rv
}

// The maximum thickness of the split view item when it resizes due to automatic sizing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388889-automaticmaximumthickness?language=objc
func (s_ SplitViewItem) AutomaticMaximumThickness() float64 {
	rv := objc.Call[float64](s_, objc.Sel("automaticMaximumThickness"))
	return rv
}

// The maximum thickness of the split view item when it resizes due to automatic sizing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388889-automaticmaximumthickness?language=objc
func (s_ SplitViewItem) SetAutomaticMaximumThickness(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setAutomaticMaximumThickness:"), value)
}

// The preferred thickness of the split view item relative to the split view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388885-preferredthicknessfraction?language=objc
func (s_ SplitViewItem) PreferredThicknessFraction() float64 {
	rv := objc.Call[float64](s_, objc.Sel("preferredThicknessFraction"))
	return rv
}

// The preferred thickness of the split view item relative to the split view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388885-preferredthicknessfraction?language=objc
func (s_ SplitViewItem) SetPreferredThicknessFraction(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setPreferredThicknessFraction:"), value)
}

// The maximum thickness of the split view item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388930-maximumthickness?language=objc
func (s_ SplitViewItem) MaximumThickness() float64 {
	rv := objc.Call[float64](s_, objc.Sel("maximumThickness"))
	return rv
}

// The maximum thickness of the split view item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388930-maximumthickness?language=objc
func (s_ SplitViewItem) SetMaximumThickness(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMaximumThickness:"), value)
}

// The priority for a split view item to hold its size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388887-holdingpriority?language=objc
func (s_ SplitViewItem) HoldingPriority() LayoutPriority {
	rv := objc.Call[LayoutPriority](s_, objc.Sel("holdingPriority"))
	return rv
}

// The priority for a split view item to hold its size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388887-holdingpriority?language=objc
func (s_ SplitViewItem) SetHoldingPriority(value LayoutPriority) {
	objc.Call[objc.Void](s_, objc.Sel("setHoldingPriority:"), value)
}

// The resizing behavior when the split view item toggles its collapsed state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388851-collapsebehavior?language=objc
func (s_ SplitViewItem) CollapseBehavior() SplitViewItemCollapseBehavior {
	rv := objc.Call[SplitViewItemCollapseBehavior](s_, objc.Sel("collapseBehavior"))
	return rv
}

// The resizing behavior when the split view item toggles its collapsed state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388851-collapsebehavior?language=objc
func (s_ SplitViewItem) SetCollapseBehavior(value SplitViewItemCollapseBehavior) {
	objc.Call[objc.Void](s_, objc.Sel("setCollapseBehavior:"), value)
}

// The standard behavior type of the split view item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388899-behavior?language=objc
func (s_ SplitViewItem) Behavior() SplitViewItemBehavior {
	rv := objc.Call[SplitViewItemBehavior](s_, objc.Sel("behavior"))
	return rv
}

// A Boolean value that determines whether the child view controller that corresponds to the split view item is in a collapsed state in the split view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388891-collapsed?language=objc
func (s_ SplitViewItem) IsCollapsed() bool {
	rv := objc.Call[bool](s_, objc.Sel("isCollapsed"))
	return rv
}

// A Boolean value that determines whether the child view controller that corresponds to the split view item is in a collapsed state in the split view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388891-collapsed?language=objc
func (s_ SplitViewItem) SetCollapsed(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setCollapsed:"), value)
}

// The view controller that the split view item represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388911-viewcontroller?language=objc
func (s_ SplitViewItem) ViewController() ViewController {
	rv := objc.Call[ViewController](s_, objc.Sel("viewController"))
	return rv
}

// The view controller that the split view item represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388911-viewcontroller?language=objc
func (s_ SplitViewItem) SetViewController(value IViewController) {
	objc.Call[objc.Void](s_, objc.Sel("setViewController:"), objc.Ptr(value))
}

// A Boolean value that determines whether a user interaction can collapse the child view controller that corresponds to the split view item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388926-cancollapse?language=objc
func (s_ SplitViewItem) CanCollapse() bool {
	rv := objc.Call[bool](s_, objc.Sel("canCollapse"))
	return rv
}

// A Boolean value that determines whether a user interaction can collapse the child view controller that corresponds to the split view item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388926-cancollapse?language=objc
func (s_ SplitViewItem) SetCanCollapse(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setCanCollapse:"), value)
}

// The minimum thickness of the split view item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388853-minimumthickness?language=objc
func (s_ SplitViewItem) MinimumThickness() float64 {
	rv := objc.Call[float64](s_, objc.Sel("minimumThickness"))
	return rv
}

// The minimum thickness of the split view item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388853-minimumthickness?language=objc
func (s_ SplitViewItem) SetMinimumThickness(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMinimumThickness:"), value)
}

// The type of separator that the app displays between the title bar and content of a window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/3622473-titlebarseparatorstyle?language=objc
func (s_ SplitViewItem) TitlebarSeparatorStyle() TitlebarSeparatorStyle {
	rv := objc.Call[TitlebarSeparatorStyle](s_, objc.Sel("titlebarSeparatorStyle"))
	return rv
}

// The type of separator that the app displays between the title bar and content of a window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/3622473-titlebarseparatorstyle?language=objc
func (s_ SplitViewItem) SetTitlebarSeparatorStyle(value TitlebarSeparatorStyle) {
	objc.Call[objc.Void](s_, objc.Sel("setTitlebarSeparatorStyle:"), value)
}

// A Boolean value that indicates whether full-height sidebars appear in the window after you set a style mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/3608197-allowsfullheightlayout?language=objc
func (s_ SplitViewItem) AllowsFullHeightLayout() bool {
	rv := objc.Call[bool](s_, objc.Sel("allowsFullHeightLayout"))
	return rv
}

// A Boolean value that indicates whether full-height sidebars appear in the window after you set a style mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/3608197-allowsfullheightlayout?language=objc
func (s_ SplitViewItem) SetAllowsFullHeightLayout(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setAllowsFullHeightLayout:"), value)
}

// A Boolean value that determines whether the split view item can temporarily expand during a drag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388895-springloaded?language=objc
func (s_ SplitViewItem) IsSpringLoaded() bool {
	rv := objc.Call[bool](s_, objc.Sel("isSpringLoaded"))
	return rv
}

// A Boolean value that determines whether the split view item can temporarily expand during a drag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/1388895-springloaded?language=objc
func (s_ SplitViewItem) SetSpringLoaded(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setSpringLoaded:"), value)
}
