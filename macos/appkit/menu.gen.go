// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var MenuClass = _MenuClass{objc.GetClass("NSMenu")}

type _MenuClass struct {
	objc.Class
}

type IMenu interface {
	objc.IObject
	InsertItemAtIndex(newItem IMenuItem, index int)
	InsertItemWithTitleActionKeyEquivalentAtIndex(string_ string, selector objc.Selector, charCode string, index int) MenuItem
	AddItem(newItem IMenuItem)
	AddItemWithTitleActionKeyEquivalent(string_ string, selector objc.Selector, charCode string) MenuItem
	RemoveItem(item IMenuItem)
	RemoveItemAtIndex(index int)
	ItemChanged(item IMenuItem)
	RemoveAllItems()
	ItemWithTag(tag int) MenuItem
	ItemWithTitle(title string) MenuItem
	ItemAtIndex(index int) MenuItem
	IndexOfItem(item IMenuItem) int
	IndexOfItemWithTitle(title string) int
	IndexOfItemWithTag(tag int) int
	IndexOfItemWithTargetAndAction(target objc.IObject, actionSelector objc.Selector) int
	IndexOfItemWithRepresentedObject(object objc.IObject) int
	IndexOfItemWithSubmenu(submenu IMenu) int
	SetSubmenuForItem(menu IMenu, item IMenuItem)
	SubmenuAction(sender objc.IObject)
	Update()
	PerformKeyEquivalent(event IEvent) bool
	PerformActionForItemAtIndex(index int)
	PopUpMenuPositioningItemAtLocationInView(item IMenuItem, location foundation.Point, view IView) bool
	CancelTracking()
	CancelTrackingWithoutAnimation()
	MenuBarHeight() float64
	NumberOfItems() int
	ItemArray() []MenuItem
	SetItemArray(value []IMenuItem)
	Supermenu() Menu
	SetSupermenu(value IMenu)
	AutoenablesItems() bool
	SetAutoenablesItems(value bool)
	Font() Font
	SetFont(value IFont)
	Title() string
	SetTitle(value string)
	MinimumWidth() float64
	SetMinimumWidth(value float64)
	Size() foundation.Size
	PropertiesToUpdate() MenuProperties
	AllowsContextMenuPlugIns() bool
	SetAllowsContextMenuPlugIns(value bool)
	ShowsStateColumn() bool
	SetShowsStateColumn(value bool)
	HighlightedItem() MenuItem
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
	Delegate() MenuDelegateWrapper
	SetDelegate(value IMenuDelegate)
	SetDelegate0(value objc.IObject)
}

type Menu struct {
	objc.Object
}

func MakeMenu(ptr unsafe.Pointer) Menu {
	return Menu{
		Object: objc.MakeObject(ptr),
	}
}

func (m_ Menu) InitWithTitle(title string) Menu {
	rv := objc.CallMethod[Menu](m_, objc.GetSelector("initWithTitle:"), title)
	return rv
}

func Menu_InitWithTitle(title string) Menu {
	return MenuClass.Alloc().InitWithTitle(title)
}

func (mc _MenuClass) Alloc() Menu {
	rv := objc.CallMethod[Menu](mc, objc.GetSelector("alloc"))
	return rv
}

func Menu_Alloc() Menu {
	return MenuClass.Alloc()
}

func (mc _MenuClass) New() Menu {
	rv := objc.CallMethod[Menu](mc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewMenu() Menu {
	return MenuClass.New()
}

func Menu_New() Menu {
	return MenuClass.New()
}

func (m_ Menu) Init() Menu {
	rv := objc.CallMethod[Menu](m_, objc.GetSelector("init"))
	return rv
}

func Menu_Init() Menu {
	return MenuClass.Alloc().Init()
}

func (mc _MenuClass) MenuBarVisible() bool {
	rv := objc.CallMethod[bool](mc, objc.GetSelector("menuBarVisible"))
	return rv
}

func Menu_MenuBarVisible() bool {
	return MenuClass.MenuBarVisible()
}

func (mc _MenuClass) SetMenuBarVisible(visible bool) {
	objc.CallMethod[objc.Void](mc, objc.GetSelector("setMenuBarVisible:"), visible)
}

func Menu_SetMenuBarVisible(visible bool) {
	MenuClass.SetMenuBarVisible(visible)
}

func (m_ Menu) InsertItemAtIndex(newItem IMenuItem, index int) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("insertItem:atIndex:"), objc.ExtractPtr(newItem), index)
}

func (m_ Menu) InsertItemWithTitleActionKeyEquivalentAtIndex(string_ string, selector objc.Selector, charCode string, index int) MenuItem {
	rv := objc.CallMethod[MenuItem](m_, objc.GetSelector("insertItemWithTitle:action:keyEquivalent:atIndex:"), string_, selector, charCode, index)
	return rv
}

func (m_ Menu) AddItem(newItem IMenuItem) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("addItem:"), objc.ExtractPtr(newItem))
}

func (m_ Menu) AddItemWithTitleActionKeyEquivalent(string_ string, selector objc.Selector, charCode string) MenuItem {
	rv := objc.CallMethod[MenuItem](m_, objc.GetSelector("addItemWithTitle:action:keyEquivalent:"), string_, selector, charCode)
	return rv
}

func (m_ Menu) RemoveItem(item IMenuItem) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("removeItem:"), objc.ExtractPtr(item))
}

func (m_ Menu) RemoveItemAtIndex(index int) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("removeItemAtIndex:"), index)
}

func (m_ Menu) ItemChanged(item IMenuItem) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("itemChanged:"), objc.ExtractPtr(item))
}

func (m_ Menu) RemoveAllItems() {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("removeAllItems"))
}

func (m_ Menu) ItemWithTag(tag int) MenuItem {
	rv := objc.CallMethod[MenuItem](m_, objc.GetSelector("itemWithTag:"), tag)
	return rv
}

func (m_ Menu) ItemWithTitle(title string) MenuItem {
	rv := objc.CallMethod[MenuItem](m_, objc.GetSelector("itemWithTitle:"), title)
	return rv
}

func (m_ Menu) ItemAtIndex(index int) MenuItem {
	rv := objc.CallMethod[MenuItem](m_, objc.GetSelector("itemAtIndex:"), index)
	return rv
}

func (m_ Menu) IndexOfItem(item IMenuItem) int {
	rv := objc.CallMethod[int](m_, objc.GetSelector("indexOfItem:"), objc.ExtractPtr(item))
	return rv
}

func (m_ Menu) IndexOfItemWithTitle(title string) int {
	rv := objc.CallMethod[int](m_, objc.GetSelector("indexOfItemWithTitle:"), title)
	return rv
}

func (m_ Menu) IndexOfItemWithTag(tag int) int {
	rv := objc.CallMethod[int](m_, objc.GetSelector("indexOfItemWithTag:"), tag)
	return rv
}

func (m_ Menu) IndexOfItemWithTargetAndAction(target objc.IObject, actionSelector objc.Selector) int {
	rv := objc.CallMethod[int](m_, objc.GetSelector("indexOfItemWithTarget:andAction:"), objc.ExtractPtr(target), actionSelector)
	return rv
}

func (m_ Menu) IndexOfItemWithRepresentedObject(object objc.IObject) int {
	rv := objc.CallMethod[int](m_, objc.GetSelector("indexOfItemWithRepresentedObject:"), objc.ExtractPtr(object))
	return rv
}

func (m_ Menu) IndexOfItemWithSubmenu(submenu IMenu) int {
	rv := objc.CallMethod[int](m_, objc.GetSelector("indexOfItemWithSubmenu:"), objc.ExtractPtr(submenu))
	return rv
}

func (m_ Menu) SetSubmenuForItem(menu IMenu, item IMenuItem) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setSubmenu:forItem:"), objc.ExtractPtr(menu), objc.ExtractPtr(item))
}

func (m_ Menu) SubmenuAction(sender objc.IObject) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("submenuAction:"), objc.ExtractPtr(sender))
}

func (m_ Menu) Update() {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("update"))
}

func (m_ Menu) PerformKeyEquivalent(event IEvent) bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("performKeyEquivalent:"), objc.ExtractPtr(event))
	return rv
}

func (m_ Menu) PerformActionForItemAtIndex(index int) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("performActionForItemAtIndex:"), index)
}

func (mc _MenuClass) PopUpContextMenuWithEventForView(menu IMenu, event IEvent, view IView) {
	objc.CallMethod[objc.Void](mc, objc.GetSelector("popUpContextMenu:withEvent:forView:"), objc.ExtractPtr(menu), objc.ExtractPtr(event), objc.ExtractPtr(view))
}

func Menu_PopUpContextMenuWithEventForView(menu IMenu, event IEvent, view IView) {
	MenuClass.PopUpContextMenuWithEventForView(menu, event, view)
}

func (mc _MenuClass) PopUpContextMenuWithEventForViewWithFont(menu IMenu, event IEvent, view IView, font IFont) {
	objc.CallMethod[objc.Void](mc, objc.GetSelector("popUpContextMenu:withEvent:forView:withFont:"), objc.ExtractPtr(menu), objc.ExtractPtr(event), objc.ExtractPtr(view), objc.ExtractPtr(font))
}

func Menu_PopUpContextMenuWithEventForViewWithFont(menu IMenu, event IEvent, view IView, font IFont) {
	MenuClass.PopUpContextMenuWithEventForViewWithFont(menu, event, view, font)
}

func (m_ Menu) PopUpMenuPositioningItemAtLocationInView(item IMenuItem, location foundation.Point, view IView) bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("popUpMenuPositioningItem:atLocation:inView:"), objc.ExtractPtr(item), location, objc.ExtractPtr(view))
	return rv
}

func (m_ Menu) CancelTracking() {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("cancelTracking"))
}

func (m_ Menu) CancelTrackingWithoutAnimation() {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("cancelTrackingWithoutAnimation"))
}

func (m_ Menu) MenuBarHeight() float64 {
	rv := objc.CallMethod[float64](m_, objc.GetSelector("menuBarHeight"))
	return rv
}

func (m_ Menu) NumberOfItems() int {
	rv := objc.CallMethod[int](m_, objc.GetSelector("numberOfItems"))
	return rv
}

func (m_ Menu) ItemArray() []MenuItem {
	rv := objc.CallMethod[[]MenuItem](m_, objc.GetSelector("itemArray"))
	return rv
}

func (m_ Menu) SetItemArray(value []IMenuItem) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setItemArray:"), value)
}

func (m_ Menu) Supermenu() Menu {
	rv := objc.CallMethod[Menu](m_, objc.GetSelector("supermenu"))
	return rv
}

func (m_ Menu) SetSupermenu(value IMenu) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setSupermenu:"), objc.ExtractPtr(value))
}

func (m_ Menu) AutoenablesItems() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("autoenablesItems"))
	return rv
}

func (m_ Menu) SetAutoenablesItems(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAutoenablesItems:"), value)
}

func (m_ Menu) Font() Font {
	rv := objc.CallMethod[Font](m_, objc.GetSelector("font"))
	return rv
}

func (m_ Menu) SetFont(value IFont) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setFont:"), objc.ExtractPtr(value))
}

func (m_ Menu) Title() string {
	rv := objc.CallMethod[string](m_, objc.GetSelector("title"))
	return rv
}

func (m_ Menu) SetTitle(value string) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setTitle:"), value)
}

func (m_ Menu) MinimumWidth() float64 {
	rv := objc.CallMethod[float64](m_, objc.GetSelector("minimumWidth"))
	return rv
}

func (m_ Menu) SetMinimumWidth(value float64) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setMinimumWidth:"), value)
}

func (m_ Menu) Size() foundation.Size {
	rv := objc.CallMethod[foundation.Size](m_, objc.GetSelector("size"))
	return rv
}

func (m_ Menu) PropertiesToUpdate() MenuProperties {
	rv := objc.CallMethod[MenuProperties](m_, objc.GetSelector("propertiesToUpdate"))
	return rv
}

func (m_ Menu) AllowsContextMenuPlugIns() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("allowsContextMenuPlugIns"))
	return rv
}

func (m_ Menu) SetAllowsContextMenuPlugIns(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAllowsContextMenuPlugIns:"), value)
}

func (m_ Menu) ShowsStateColumn() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("showsStateColumn"))
	return rv
}

func (m_ Menu) SetShowsStateColumn(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setShowsStateColumn:"), value)
}

func (m_ Menu) HighlightedItem() MenuItem {
	rv := objc.CallMethod[MenuItem](m_, objc.GetSelector("highlightedItem"))
	return rv
}

func (m_ Menu) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.CallMethod[UserInterfaceLayoutDirection](m_, objc.GetSelector("userInterfaceLayoutDirection"))
	return rv
}

func (m_ Menu) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setUserInterfaceLayoutDirection:"), value)
}

func (m_ Menu) Delegate() MenuDelegateWrapper {
	rv := objc.CallMethod[MenuDelegateWrapper](m_, objc.GetSelector("delegate"))
	return rv
}

func (m_ Menu) SetDelegate(value IMenuDelegate) {
	po := objc.WrapAsProtocol("NSMenuDelegate", value)
	objc.SetAssociatedObject(m_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setDelegate:"), po)
}

func (m_ Menu) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}
