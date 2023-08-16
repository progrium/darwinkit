// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TabViewController] class.
var TabViewControllerClass = _TabViewControllerClass{objc.GetClass("NSTabViewController")}

type _TabViewControllerClass struct {
	objc.Class
}

// An interface definition for the [TabViewController] class.
type ITabViewController interface {
	IViewController
	TabViewItemForViewController(viewController IViewController) TabViewItem
	ToolbarAllowedItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier
	RemoveTabViewItem(tabViewItem ITabViewItem)
	ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(toolbar IToolbar, itemIdentifier ToolbarItemIdentifier, flag bool) ToolbarItem
	ToolbarDefaultItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier
	ToolbarSelectableItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier
	InsertTabViewItemAtIndex(tabViewItem ITabViewItem, index int)
	AddTabViewItem(tabViewItem ITabViewItem)
	CanPropagateSelectedChildViewControllerTitle() bool
	SetCanPropagateSelectedChildViewControllerTitle(value bool)
	TabStyle() TabViewControllerTabStyle
	SetTabStyle(value TabViewControllerTabStyle)
	TabViewItems() []TabViewItem
	SetTabViewItems(value []ITabViewItem)
	TabView() TabView
	SetTabView(value ITabView)
	TransitionOptions() ViewControllerTransitionOptions
	SetTransitionOptions(value ViewControllerTransitionOptions)
	SelectedTabViewItemIndex() int
	SetSelectedTabViewItemIndex(value int)
}

// A container view controller that manages a tab view interface, which organizes multiple pages of content but displays only one page at a time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller?language=objc
type TabViewController struct {
	ViewController
}

func TabViewControllerFrom(ptr unsafe.Pointer) TabViewController {
	return TabViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}

func (tc _TabViewControllerClass) Alloc() TabViewController {
	rv := objc.Call[TabViewController](tc, objc.Sel("alloc"))
	return rv
}

func TabViewController_Alloc() TabViewController {
	return TabViewControllerClass.Alloc()
}

func (tc _TabViewControllerClass) New() TabViewController {
	rv := objc.Call[TabViewController](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTabViewController() TabViewController {
	return TabViewControllerClass.New()
}

func (t_ TabViewController) Init() TabViewController {
	rv := objc.Call[TabViewController](t_, objc.Sel("init"))
	return rv
}

func (t_ TabViewController) InitWithNibNameBundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) TabViewController {
	rv := objc.Call[TabViewController](t_, objc.Sel("initWithNibName:bundle:"), nibNameOrNil, objc.Ptr(nibBundleOrNil))
	return rv
}

// Returns a view controller object initialized to the nib file in the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434481-initwithnibname?language=objc
func TabViewController_InitWithNibNameBundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) TabViewController {
	return TabViewControllerClass.Alloc().InitWithNibNameBundle(nibNameOrNil, nibBundleOrNil)
}

// Returns the tab view item for the specified child view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428233-tabviewitemforviewcontroller?language=objc
func (t_ TabViewController) TabViewItemForViewController(viewController IViewController) TabViewItem {
	rv := objc.Call[TabViewItem](t_, objc.Sel("tabViewItemForViewController:"), objc.Ptr(viewController))
	return rv
}

// Returns the array of identifier strings for the allowed toolbar items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428229-toolbaralloweditemidentifiers?language=objc
func (t_ TabViewController) ToolbarAllowedItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier {
	rv := objc.Call[[]ToolbarItemIdentifier](t_, objc.Sel("toolbarAllowedItemIdentifiers:"), objc.Ptr(toolbar))
	return rv
}

// Removes the specified tab view item from the tab view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428235-removetabviewitem?language=objc
func (t_ TabViewController) RemoveTabViewItem(tabViewItem ITabViewItem) {
	objc.Call[objc.Void](t_, objc.Sel("removeTabViewItem:"), objc.Ptr(tabViewItem))
}

// Returns the toolbar item for the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428218-toolbar?language=objc
func (t_ TabViewController) ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(toolbar IToolbar, itemIdentifier ToolbarItemIdentifier, flag bool) ToolbarItem {
	rv := objc.Call[ToolbarItem](t_, objc.Sel("toolbar:itemForItemIdentifier:willBeInsertedIntoToolbar:"), objc.Ptr(toolbar), itemIdentifier, flag)
	return rv
}

// Returns the array of identifier strings for the default toolbar items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428251-toolbardefaultitemidentifiers?language=objc
func (t_ TabViewController) ToolbarDefaultItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier {
	rv := objc.Call[[]ToolbarItemIdentifier](t_, objc.Sel("toolbarDefaultItemIdentifiers:"), objc.Ptr(toolbar))
	return rv
}

// Returns the array of identifier strings for the selectable toolbar items [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428261-toolbarselectableitemidentifiers?language=objc
func (t_ TabViewController) ToolbarSelectableItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier {
	rv := objc.Call[[]ToolbarItemIdentifier](t_, objc.Sel("toolbarSelectableItemIdentifiers:"), objc.Ptr(toolbar))
	return rv
}

// Inserts a tab view into the tab view controller’s list of tabs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428255-inserttabviewitem?language=objc
func (t_ TabViewController) InsertTabViewItemAtIndex(tabViewItem ITabViewItem, index int) {
	objc.Call[objc.Void](t_, objc.Sel("insertTabViewItem:atIndex:"), objc.Ptr(tabViewItem), index)
}

// Adds the specified tab to the end of the tab view controller’s list of tabs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428214-addtabviewitem?language=objc
func (t_ TabViewController) AddTabViewItem(tabViewItem ITabViewItem) {
	objc.Call[objc.Void](t_, objc.Sel("addTabViewItem:"), objc.Ptr(tabViewItem))
}

// A Boolean value indicating whether the tab view controller gets its title from the selected child view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428239-canpropagateselectedchildviewcon?language=objc
func (t_ TabViewController) CanPropagateSelectedChildViewControllerTitle() bool {
	rv := objc.Call[bool](t_, objc.Sel("canPropagateSelectedChildViewControllerTitle"))
	return rv
}

// A Boolean value indicating whether the tab view controller gets its title from the selected child view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428239-canpropagateselectedchildviewcon?language=objc
func (t_ TabViewController) SetCanPropagateSelectedChildViewControllerTitle(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setCanPropagateSelectedChildViewControllerTitle:"), value)
}

// The style used to display the tabs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428259-tabstyle?language=objc
func (t_ TabViewController) TabStyle() TabViewControllerTabStyle {
	rv := objc.Call[TabViewControllerTabStyle](t_, objc.Sel("tabStyle"))
	return rv
}

// The style used to display the tabs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428259-tabstyle?language=objc
func (t_ TabViewController) SetTabStyle(value TabViewControllerTabStyle) {
	objc.Call[objc.Void](t_, objc.Sel("setTabStyle:"), value)
}

// The array of tab view items used to manage each of the child view controllers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428222-tabviewitems?language=objc
func (t_ TabViewController) TabViewItems() []TabViewItem {
	rv := objc.Call[[]TabViewItem](t_, objc.Sel("tabViewItems"))
	return rv
}

// The array of tab view items used to manage each of the child view controllers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428222-tabviewitems?language=objc
func (t_ TabViewController) SetTabViewItems(value []ITabViewItem) {
	objc.Call[objc.Void](t_, objc.Sel("setTabViewItems:"), value)
}

// The tab view that manages the views of the interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428241-tabview?language=objc
func (t_ TabViewController) TabView() TabView {
	rv := objc.Call[TabView](t_, objc.Sel("tabView"))
	return rv
}

// The tab view that manages the views of the interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428241-tabview?language=objc
func (t_ TabViewController) SetTabView(value ITabView) {
	objc.Call[objc.Void](t_, objc.Sel("setTabView:"), objc.Ptr(value))
}

// The animation options to use when switching between tabs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428257-transitionoptions?language=objc
func (t_ TabViewController) TransitionOptions() ViewControllerTransitionOptions {
	rv := objc.Call[ViewControllerTransitionOptions](t_, objc.Sel("transitionOptions"))
	return rv
}

// The animation options to use when switching between tabs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428257-transitionoptions?language=objc
func (t_ TabViewController) SetTransitionOptions(value ViewControllerTransitionOptions) {
	objc.Call[objc.Void](t_, objc.Sel("setTransitionOptions:"), value)
}

// The index of the selected tab. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428220-selectedtabviewitemindex?language=objc
func (t_ TabViewController) SelectedTabViewItemIndex() int {
	rv := objc.Call[int](t_, objc.Sel("selectedTabViewItemIndex"))
	return rv
}

// The index of the selected tab. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/1428220-selectedtabviewitemindex?language=objc
func (t_ TabViewController) SetSelectedTabViewItemIndex(value int) {
	objc.Call[objc.Void](t_, objc.Sel("setSelectedTabViewItemIndex:"), value)
}
