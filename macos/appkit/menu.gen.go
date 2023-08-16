// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Menu] class.
var MenuClass = _MenuClass{objc.GetClass("NSMenu")}

type _MenuClass struct {
	objc.Class
}

// An interface definition for the [Menu] class.
type IMenu interface {
	objc.IObject
	ItemAtIndex(index int) MenuItem
	PerformActionForItemAtIndex(index int)
	IndexOfItemWithSubmenu(submenu IMenu) int
	CancelTracking()
	AddItem(newItem IMenuItem)
	IndexOfItemWithTag(tag int) int
	InsertItemWithTitleActionKeyEquivalentAtIndex(string_ string, selector objc.Selector, charCode string, index int) MenuItem
	ItemChanged(item IMenuItem)
	RemoveAllItems()
	RemoveItemAtIndex(index int)
	RemoveItem(item IMenuItem)
	AddItemWithTitleActionKeyEquivalent(string_ string, selector objc.Selector, charCode string) MenuItem
	ItemWithTitle(title string) MenuItem
	IndexOfItemWithRepresentedObject(object objc.IObject) int
	PerformKeyEquivalent(event IEvent) bool
	SetSubmenuForItem(menu IMenu, item IMenuItem)
	IndexOfItem(item IMenuItem) int
	SubmenuAction(sender objc.IObject)
	PopUpMenuPositioningItemAtLocationInView(item IMenuItem, location foundation.Point, view IView) bool
	IndexOfItemWithTargetAndAction(target objc.IObject, actionSelector objc.Selector) int
	IndexOfItemWithTitle(title string) int
	Update()
	ItemWithTag(tag int) MenuItem
	InsertItemAtIndex(newItem IMenuItem, index int)
	CancelTrackingWithoutAnimation()
	AllowsContextMenuPlugIns() bool
	SetAllowsContextMenuPlugIns(value bool)
	Supermenu() Menu
	SetSupermenu(value IMenu)
	ItemArray() []MenuItem
	SetItemArray(value []IMenuItem)
	MenuBarHeight() float64
	HighlightedItem() MenuItem
	AutoenablesItems() bool
	SetAutoenablesItems(value bool)
	Delegate() MenuDelegateWrapper
	SetDelegate(value PMenuDelegate)
	SetDelegateObject(valueObject objc.IObject)
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
	MinimumWidth() float64
	SetMinimumWidth(value float64)
	Font() Font
	SetFont(value IFont)
	PropertiesToUpdate() MenuProperties
	NumberOfItems() int
	Title() string
	SetTitle(value string)
	ShowsStateColumn() bool
	SetShowsStateColumn(value bool)
	Size() foundation.Size
}

// An object that manages an app’s menus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu?language=objc
type Menu struct {
	objc.Object
}

func MenuFrom(ptr unsafe.Pointer) Menu {
	return Menu{
		Object: objc.ObjectFrom(ptr),
	}
}

func (m_ Menu) InitWithTitle(title string) Menu {
	rv := objc.Call[Menu](m_, objc.Sel("initWithTitle:"), title)
	return rv
}

// Initializes and returns a menu having the specified title and with autoenabling of menu items turned on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518144-initwithtitle?language=objc
func Menu_InitWithTitle(title string) Menu {
	return MenuClass.Alloc().InitWithTitle(title)
}

func (mc _MenuClass) Alloc() Menu {
	rv := objc.Call[Menu](mc, objc.Sel("alloc"))
	return rv
}

func Menu_Alloc() Menu {
	return MenuClass.Alloc()
}

func (mc _MenuClass) New() Menu {
	rv := objc.Call[Menu](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMenu() Menu {
	return MenuClass.New()
}

func (m_ Menu) Init() Menu {
	rv := objc.Call[Menu](m_, objc.Sel("init"))
	return rv
}

// Returns the menu item at a specific location of the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518218-itematindex?language=objc
func (m_ Menu) ItemAtIndex(index int) MenuItem {
	rv := objc.Call[MenuItem](m_, objc.Sel("itemAtIndex:"), index)
	return rv
}

// Causes the application to send the action message of a specified menu item to its target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518210-performactionforitematindex?language=objc
func (m_ Menu) PerformActionForItemAtIndex(index int) {
	objc.Call[objc.Void](m_, objc.Sel("performActionForItemAtIndex:"), index)
}

// Returns the index of the menu item in the menu with the given submenu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518216-indexofitemwithsubmenu?language=objc
func (m_ Menu) IndexOfItemWithSubmenu(submenu IMenu) int {
	rv := objc.Call[int](m_, objc.Sel("indexOfItemWithSubmenu:"), objc.Ptr(submenu))
	return rv
}

// Dismisses the menu and ends all menu tracking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518150-canceltracking?language=objc
func (m_ Menu) CancelTracking() {
	objc.Call[objc.Void](m_, objc.Sel("cancelTracking"))
}

// Adds a menu item to the end of the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518176-additem?language=objc
func (m_ Menu) AddItem(newItem IMenuItem) {
	objc.Call[objc.Void](m_, objc.Sel("addItem:"), objc.Ptr(newItem))
}

// Returns the index of the first menu item in the menu identified by a tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518164-indexofitemwithtag?language=objc
func (m_ Menu) IndexOfItemWithTag(tag int) int {
	rv := objc.Call[int](m_, objc.Sel("indexOfItemWithTag:"), tag)
	return rv
}

// Creates and adds a menu item at a specified location in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518146-insertitemwithtitle?language=objc
func (m_ Menu) InsertItemWithTitleActionKeyEquivalentAtIndex(string_ string, selector objc.Selector, charCode string, index int) MenuItem {
	rv := objc.Call[MenuItem](m_, objc.Sel("insertItemWithTitle:action:keyEquivalent:atIndex:"), string_, selector, charCode, index)
	return rv
}

// Invoked when a menu item is modified visually (for example, its title changes). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518154-itemchanged?language=objc
func (m_ Menu) ItemChanged(item IMenuItem) {
	objc.Call[objc.Void](m_, objc.Sel("itemChanged:"), objc.Ptr(item))
}

// Removes all the menu items in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518234-removeallitems?language=objc
func (m_ Menu) RemoveAllItems() {
	objc.Call[objc.Void](m_, objc.Sel("removeAllItems"))
}

// Removes the menu item at a specified location in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518207-removeitematindex?language=objc
func (m_ Menu) RemoveItemAtIndex(index int) {
	objc.Call[objc.Void](m_, objc.Sel("removeItemAtIndex:"), index)
}

// Removes a menu item from the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518257-removeitem?language=objc
func (m_ Menu) RemoveItem(item IMenuItem) {
	objc.Call[objc.Void](m_, objc.Sel("removeItem:"), objc.Ptr(item))
}

// Sets whether the menu bar is visible and selectable by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518200-setmenubarvisible?language=objc
func (mc _MenuClass) SetMenuBarVisible(visible bool) {
	objc.Call[objc.Void](mc, objc.Sel("setMenuBarVisible:"), visible)
}

// Sets whether the menu bar is visible and selectable by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518200-setmenubarvisible?language=objc
func Menu_SetMenuBarVisible(visible bool) {
	MenuClass.SetMenuBarVisible(visible)
}

// Creates a new menu item and adds it to the end of the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518181-additemwithtitle?language=objc
func (m_ Menu) AddItemWithTitleActionKeyEquivalent(string_ string, selector objc.Selector, charCode string) MenuItem {
	rv := objc.Call[MenuItem](m_, objc.Sel("addItemWithTitle:action:keyEquivalent:"), string_, selector, charCode)
	return rv
}

// Returns the first menu item in the menu with a specified title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518248-itemwithtitle?language=objc
func (m_ Menu) ItemWithTitle(title string) MenuItem {
	rv := objc.Call[MenuItem](m_, objc.Sel("itemWithTitle:"), title)
	return rv
}

// Returns the index of the first menu item in the menu that has a given represented object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518175-indexofitemwithrepresentedobject?language=objc
func (m_ Menu) IndexOfItemWithRepresentedObject(object objc.IObject) int {
	rv := objc.Call[int](m_, objc.Sel("indexOfItemWithRepresentedObject:"), object)
	return rv
}

// Performs the action for the menu item that corresponds to the given key equivalent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518198-performkeyequivalent?language=objc
func (m_ Menu) PerformKeyEquivalent(event IEvent) bool {
	rv := objc.Call[bool](m_, objc.Sel("performKeyEquivalent:"), objc.Ptr(event))
	return rv
}

// Assigns a menu to be a submenu of the menu controlled by a given menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518194-setsubmenu?language=objc
func (m_ Menu) SetSubmenuForItem(menu IMenu, item IMenuItem) {
	objc.Call[objc.Void](m_, objc.Sel("setSubmenu:forItem:"), objc.Ptr(menu), objc.Ptr(item))
}

// Returns the index identifying the location of a specified menu item in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518178-indexofitem?language=objc
func (m_ Menu) IndexOfItem(item IMenuItem) int {
	rv := objc.Call[int](m_, objc.Sel("indexOfItem:"), objc.Ptr(item))
	return rv
}

// The action method assigned to menu items that open submenus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518179-submenuaction?language=objc
func (m_ Menu) SubmenuAction(sender objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("submenuAction:"), sender)
}

// Pops up the menu at the specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518212-popupmenupositioningitem?language=objc
func (m_ Menu) PopUpMenuPositioningItemAtLocationInView(item IMenuItem, location foundation.Point, view IView) bool {
	rv := objc.Call[bool](m_, objc.Sel("popUpMenuPositioningItem:atLocation:inView:"), objc.Ptr(item), location, objc.Ptr(view))
	return rv
}

// Returns the index of the first menu item in the menu that has a specified action and target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518153-indexofitemwithtarget?language=objc
func (m_ Menu) IndexOfItemWithTargetAndAction(target objc.IObject, actionSelector objc.Selector) int {
	rv := objc.Call[int](m_, objc.Sel("indexOfItemWithTarget:andAction:"), target, actionSelector)
	return rv
}

// Displays a contextual menu over a view for an event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518170-popupcontextmenu?language=objc
func (mc _MenuClass) PopUpContextMenuWithEventForView(menu IMenu, event IEvent, view IView) {
	objc.Call[objc.Void](mc, objc.Sel("popUpContextMenu:withEvent:forView:"), objc.Ptr(menu), objc.Ptr(event), objc.Ptr(view))
}

// Displays a contextual menu over a view for an event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518170-popupcontextmenu?language=objc
func Menu_PopUpContextMenuWithEventForView(menu IMenu, event IEvent, view IView) {
	MenuClass.PopUpContextMenuWithEventForView(menu, event, view)
}

// Returns a Boolean value that indicates whether the menu bar is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518236-menubarvisible?language=objc
func (mc _MenuClass) MenuBarVisible() bool {
	rv := objc.Call[bool](mc, objc.Sel("menuBarVisible"))
	return rv
}

// Returns a Boolean value that indicates whether the menu bar is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518236-menubarvisible?language=objc
func Menu_MenuBarVisible() bool {
	return MenuClass.MenuBarVisible()
}

// Returns the index of the first menu item in the menu that has a specified title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518237-indexofitemwithtitle?language=objc
func (m_ Menu) IndexOfItemWithTitle(title string) int {
	rv := objc.Call[int](m_, objc.Sel("indexOfItemWithTitle:"), title)
	return rv
}

// Enables or disables the menu items of the menu based on the NSMenuValidation informal protocol and sizes the menu to fit its current menu items if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518249-update?language=objc
func (m_ Menu) Update() {
	objc.Call[objc.Void](m_, objc.Sel("update"))
}

// Returns the first menu item in the menu with the specified tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518223-itemwithtag?language=objc
func (m_ Menu) ItemWithTag(tag int) MenuItem {
	rv := objc.Call[MenuItem](m_, objc.Sel("itemWithTag:"), tag)
	return rv
}

// Inserts a menu item into the menu at a specific location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518201-insertitem?language=objc
func (m_ Menu) InsertItemAtIndex(newItem IMenuItem, index int) {
	objc.Call[objc.Void](m_, objc.Sel("insertItem:atIndex:"), objc.Ptr(newItem), index)
}

// Dismisses the menu and ends all menu tracking without displaying the associated animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518244-canceltrackingwithoutanimation?language=objc
func (m_ Menu) CancelTrackingWithoutAnimation() {
	objc.Call[objc.Void](m_, objc.Sel("cancelTrackingWithoutAnimation"))
}

// Indicates whether the pop-up menu allows appending of contextual menu plug-in items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518220-allowscontextmenuplugins?language=objc
func (m_ Menu) AllowsContextMenuPlugIns() bool {
	rv := objc.Call[bool](m_, objc.Sel("allowsContextMenuPlugIns"))
	return rv
}

// Indicates whether the pop-up menu allows appending of contextual menu plug-in items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518220-allowscontextmenuplugins?language=objc
func (m_ Menu) SetAllowsContextMenuPlugIns(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAllowsContextMenuPlugIns:"), value)
}

// The parent menu that contains the menu as a submenu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518204-supermenu?language=objc
func (m_ Menu) Supermenu() Menu {
	rv := objc.Call[Menu](m_, objc.Sel("supermenu"))
	return rv
}

// The parent menu that contains the menu as a submenu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518204-supermenu?language=objc
func (m_ Menu) SetSupermenu(value IMenu) {
	objc.Call[objc.Void](m_, objc.Sel("setSupermenu:"), objc.Ptr(value))
}

// An array containing the menu items in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518186-itemarray?language=objc
func (m_ Menu) ItemArray() []MenuItem {
	rv := objc.Call[[]MenuItem](m_, objc.Sel("itemArray"))
	return rv
}

// An array containing the menu items in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518186-itemarray?language=objc
func (m_ Menu) SetItemArray(value []IMenuItem) {
	objc.Call[objc.Void](m_, objc.Sel("setItemArray:"), value)
}

// The menu bar height for the main menu in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518141-menubarheight?language=objc
func (m_ Menu) MenuBarHeight() float64 {
	rv := objc.Call[float64](m_, objc.Sel("menuBarHeight"))
	return rv
}

// Indicates the currently highlighted item in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518222-highlighteditem?language=objc
func (m_ Menu) HighlightedItem() MenuItem {
	rv := objc.Call[MenuItem](m_, objc.Sel("highlightedItem"))
	return rv
}

// Indicates whether the menu automatically enables and disables its menu items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518227-autoenablesitems?language=objc
func (m_ Menu) AutoenablesItems() bool {
	rv := objc.Call[bool](m_, objc.Sel("autoenablesItems"))
	return rv
}

// Indicates whether the menu automatically enables and disables its menu items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518227-autoenablesitems?language=objc
func (m_ Menu) SetAutoenablesItems(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAutoenablesItems:"), value)
}

// The delegate of the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518169-delegate?language=objc
func (m_ Menu) Delegate() MenuDelegateWrapper {
	rv := objc.Call[MenuDelegateWrapper](m_, objc.Sel("delegate"))
	return rv
}

// The delegate of the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518169-delegate?language=objc
func (m_ Menu) SetDelegate(value PMenuDelegate) {
	po0 := objc.WrapAsProtocol("NSMenuDelegate", value)
	objc.SetAssociatedObject(m_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](m_, objc.Sel("setDelegate:"), po0)
}

// The delegate of the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518169-delegate?language=objc
func (m_ Menu) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// Configures the layout direction of menu items in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518254-userinterfacelayoutdirection?language=objc
func (m_ Menu) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Call[UserInterfaceLayoutDirection](m_, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}

// Configures the layout direction of menu items in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518254-userinterfacelayoutdirection?language=objc
func (m_ Menu) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Call[objc.Void](m_, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}

// The minimum width of the menu in screen coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518221-minimumwidth?language=objc
func (m_ Menu) MinimumWidth() float64 {
	rv := objc.Call[float64](m_, objc.Sel("minimumWidth"))
	return rv
}

// The minimum width of the menu in screen coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518221-minimumwidth?language=objc
func (m_ Menu) SetMinimumWidth(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setMinimumWidth:"), value)
}

// The font of the menu and its submenus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518230-font?language=objc
func (m_ Menu) Font() Font {
	rv := objc.Call[Font](m_, objc.Sel("font"))
	return rv
}

// The font of the menu and its submenus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518230-font?language=objc
func (m_ Menu) SetFont(value IFont) {
	objc.Call[objc.Void](m_, objc.Sel("setFont:"), objc.Ptr(value))
}

// The available properties for the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518245-propertiestoupdate?language=objc
func (m_ Menu) PropertiesToUpdate() MenuProperties {
	rv := objc.Call[MenuProperties](m_, objc.Sel("propertiesToUpdate"))
	return rv
}

// The number of menu items in the menu, including separator items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518202-numberofitems?language=objc
func (m_ Menu) NumberOfItems() int {
	rv := objc.Call[int](m_, objc.Sel("numberOfItems"))
	return rv
}

// The title of the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518192-title?language=objc
func (m_ Menu) Title() string {
	rv := objc.Call[string](m_, objc.Sel("title"))
	return rv
}

// The title of the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518192-title?language=objc
func (m_ Menu) SetTitle(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setTitle:"), value)
}

// Indicates whether the menu displays the state column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518253-showsstatecolumn?language=objc
func (m_ Menu) ShowsStateColumn() bool {
	rv := objc.Call[bool](m_, objc.Sel("showsStateColumn"))
	return rv
}

// Indicates whether the menu displays the state column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518253-showsstatecolumn?language=objc
func (m_ Menu) SetShowsStateColumn(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setShowsStateColumn:"), value)
}

// The size of the menu in screen coordinates [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenu/1518185-size?language=objc
func (m_ Menu) Size() foundation.Size {
	rv := objc.Call[foundation.Size](m_, objc.Sel("size"))
	return rv
}
