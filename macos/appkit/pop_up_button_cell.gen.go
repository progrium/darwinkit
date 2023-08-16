// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PopUpButtonCell] class.
var PopUpButtonCellClass = _PopUpButtonCellClass{objc.GetClass("NSPopUpButtonCell")}

type _PopUpButtonCellClass struct {
	objc.Class
}

// An interface definition for the [PopUpButtonCell] class.
type IPopUpButtonCell interface {
	IMenuItemCell
	ItemAtIndex(index int) MenuItem
	IndexOfItemWithTag(tag int) int
	InsertItemWithTitleAtIndex(title string, index int)
	RemoveAllItems()
	RemoveItemAtIndex(index int)
	RemoveItemWithTitle(title string)
	SelectItemWithTag(tag int) bool
	AddItemWithTitle(title string)
	SelectItemWithTitle(title string)
	DismissPopUp()
	ItemWithTitle(title string) MenuItem
	IndexOfItemWithRepresentedObject(obj objc.IObject) int
	ItemTitleAtIndex(index int) string
	SelectItem(item IMenuItem)
	IndexOfItem(item IMenuItem) int
	PerformClickWithFrameInView(frame foundation.Rect, controlView IView)
	AddItemsWithTitles(itemTitles []string)
	SelectItemAtIndex(index int)
	SynchronizeTitleAndSelectedItem()
	IndexOfItemWithTargetAndAction(target objc.IObject, actionSelector objc.Selector) int
	IndexOfItemWithTitle(title string) int
	AttachPopUpWithFrameInView(cellFrame foundation.Rect, controlView IView)
	AltersStateOfSelectedItem() bool
	SetAltersStateOfSelectedItem(value bool)
	IndexOfSelectedItem() int
	SelectedItem() MenuItem
	ArrowPosition() PopUpArrowPosition
	SetArrowPosition(value PopUpArrowPosition)
	ItemArray() []MenuItem
	AutoenablesItems() bool
	SetAutoenablesItems(value bool)
	PullsDown() bool
	SetPullsDown(value bool)
	LastItem() MenuItem
	NumberOfItems() int
	TitleOfSelectedItem() string
	ItemTitles() []string
	PreferredEdge() foundation.RectEdge
	SetPreferredEdge(value foundation.RectEdge)
	UsesItemFromMenu() bool
	SetUsesItemFromMenu(value bool)
}

// The NSPopUpButtonCell class defines the visual appearance of pop-up buttons that display pop-up or pull-down menus. Pop-up menus present the user with a set of choices, much the way radio buttons do, but using much less space. Pull-down menus also provide a set of choices but present the information in a slightly different way, usually to provide a set of commands from which the user can choose. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell?language=objc
type PopUpButtonCell struct {
	MenuItemCell
}

func PopUpButtonCellFrom(ptr unsafe.Pointer) PopUpButtonCell {
	return PopUpButtonCell{
		MenuItemCell: MenuItemCellFrom(ptr),
	}
}

func (p_ PopUpButtonCell) InitTextCellPullsDown(stringValue string, pullDown bool) PopUpButtonCell {
	rv := objc.Call[PopUpButtonCell](p_, objc.Sel("initTextCell:pullsDown:"), stringValue, pullDown)
	return rv
}

// Returns an NSPopUpButtonCell object initialized with the specified title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1528591-inittextcell?language=objc
func PopUpButtonCell_InitTextCellPullsDown(stringValue string, pullDown bool) PopUpButtonCell {
	return PopUpButtonCellClass.Alloc().InitTextCellPullsDown(stringValue, pullDown)
}

func (pc _PopUpButtonCellClass) Alloc() PopUpButtonCell {
	rv := objc.Call[PopUpButtonCell](pc, objc.Sel("alloc"))
	return rv
}

func PopUpButtonCell_Alloc() PopUpButtonCell {
	return PopUpButtonCellClass.Alloc()
}

func (pc _PopUpButtonCellClass) New() PopUpButtonCell {
	rv := objc.Call[PopUpButtonCell](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPopUpButtonCell() PopUpButtonCell {
	return PopUpButtonCellClass.New()
}

func (p_ PopUpButtonCell) Init() PopUpButtonCell {
	rv := objc.Call[PopUpButtonCell](p_, objc.Sel("init"))
	return rv
}

func (p_ PopUpButtonCell) InitTextCell(string_ string) PopUpButtonCell {
	rv := objc.Call[PopUpButtonCell](p_, objc.Sel("initTextCell:"), string_)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1641970-inittextcell?language=objc
func PopUpButtonCell_InitTextCell(string_ string) PopUpButtonCell {
	return PopUpButtonCellClass.Alloc().InitTextCell(string_)
}

func (p_ PopUpButtonCell) InitImageCell(image IImage) PopUpButtonCell {
	rv := objc.Call[PopUpButtonCell](p_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1639152-initimagecell?language=objc
func PopUpButtonCell_InitImageCell(image IImage) PopUpButtonCell {
	return PopUpButtonCellClass.Alloc().InitImageCell(image)
}

// Returns the menu item at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1525784-itematindex?language=objc
func (p_ PopUpButtonCell) ItemAtIndex(index int) MenuItem {
	rv := objc.Call[MenuItem](p_, objc.Sel("itemAtIndex:"), index)
	return rv
}

// Returns the index of the menu item with the specified tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1527486-indexofitemwithtag?language=objc
func (p_ PopUpButtonCell) IndexOfItemWithTag(tag int) int {
	rv := objc.Call[int](p_, objc.Sel("indexOfItemWithTag:"), tag)
	return rv
}

// Inserts an item at the specified position in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1531000-insertitemwithtitle?language=objc
func (p_ PopUpButtonCell) InsertItemWithTitleAtIndex(title string, index int) {
	objc.Call[objc.Void](p_, objc.Sel("insertItemWithTitle:atIndex:"), title, index)
}

// Removes all items in the receiver’s item menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1534852-removeallitems?language=objc
func (p_ PopUpButtonCell) RemoveAllItems() {
	objc.Call[objc.Void](p_, objc.Sel("removeAllItems"))
}

// Removes the item at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1528201-removeitematindex?language=objc
func (p_ PopUpButtonCell) RemoveItemAtIndex(index int) {
	objc.Call[objc.Void](p_, objc.Sel("removeItemAtIndex:"), index)
}

// Removes the item with the specified title from the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1530732-removeitemwithtitle?language=objc
func (p_ PopUpButtonCell) RemoveItemWithTitle(title string) {
	objc.Call[objc.Void](p_, objc.Sel("removeItemWithTitle:"), title)
}

// Selects the menu item with the specified tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1534888-selectitemwithtag?language=objc
func (p_ PopUpButtonCell) SelectItemWithTag(tag int) bool {
	rv := objc.Call[bool](p_, objc.Sel("selectItemWithTag:"), tag)
	return rv
}

// Adds an item with the specified title to the end of the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1526296-additemwithtitle?language=objc
func (p_ PopUpButtonCell) AddItemWithTitle(title string) {
	objc.Call[objc.Void](p_, objc.Sel("addItemWithTitle:"), title)
}

// Selects the item with the specified title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1534289-selectitemwithtitle?language=objc
func (p_ PopUpButtonCell) SelectItemWithTitle(title string) {
	objc.Call[objc.Void](p_, objc.Sel("selectItemWithTitle:"), title)
}

// Dismisses the pop-up button’s menu by ordering its window out. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1535041-dismisspopup?language=objc
func (p_ PopUpButtonCell) DismissPopUp() {
	objc.Call[objc.Void](p_, objc.Sel("dismissPopUp"))
}

// Returns the menu item with the specified title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1534638-itemwithtitle?language=objc
func (p_ PopUpButtonCell) ItemWithTitle(title string) MenuItem {
	rv := objc.Call[MenuItem](p_, objc.Sel("itemWithTitle:"), title)
	return rv
}

// Returns the index of the menu item that holds the specified represented object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1525489-indexofitemwithrepresentedobject?language=objc
func (p_ PopUpButtonCell) IndexOfItemWithRepresentedObject(obj objc.IObject) int {
	rv := objc.Call[int](p_, objc.Sel("indexOfItemWithRepresentedObject:"), obj)
	return rv
}

// Returns the title of the item at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1534248-itemtitleatindex?language=objc
func (p_ PopUpButtonCell) ItemTitleAtIndex(index int) string {
	rv := objc.Call[string](p_, objc.Sel("itemTitleAtIndex:"), index)
	return rv
}

// Selects the specified menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1525225-selectitem?language=objc
func (p_ PopUpButtonCell) SelectItem(item IMenuItem) {
	objc.Call[objc.Void](p_, objc.Sel("selectItem:"), objc.Ptr(item))
}

// Returns the index of the specified menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1532397-indexofitem?language=objc
func (p_ PopUpButtonCell) IndexOfItem(item IMenuItem) int {
	rv := objc.Call[int](p_, objc.Sel("indexOfItem:"), objc.Ptr(item))
	return rv
}

// Displays the receiver’s menu and track mouse events in it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1530205-performclickwithframe?language=objc
func (p_ PopUpButtonCell) PerformClickWithFrameInView(frame foundation.Rect, controlView IView) {
	objc.Call[objc.Void](p_, objc.Sel("performClickWithFrame:inView:"), frame, objc.Ptr(controlView))
}

// Adds multiple items to the end of the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1530154-additemswithtitles?language=objc
func (p_ PopUpButtonCell) AddItemsWithTitles(itemTitles []string) {
	objc.Call[objc.Void](p_, objc.Sel("addItemsWithTitles:"), itemTitles)
}

// Selects the item in the menu at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1534855-selectitematindex?language=objc
func (p_ PopUpButtonCell) SelectItemAtIndex(index int) {
	objc.Call[objc.Void](p_, objc.Sel("selectItemAtIndex:"), index)
}

// Synchronizes the pop-up button’s displayed item with the currently selected menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1535190-synchronizetitleandselecteditem?language=objc
func (p_ PopUpButtonCell) SynchronizeTitleAndSelectedItem() {
	objc.Call[objc.Void](p_, objc.Sel("synchronizeTitleAndSelectedItem"))
}

// Returns the index of the menu item with the specified target and action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1534909-indexofitemwithtarget?language=objc
func (p_ PopUpButtonCell) IndexOfItemWithTargetAndAction(target objc.IObject, actionSelector objc.Selector) int {
	rv := objc.Call[int](p_, objc.Sel("indexOfItemWithTarget:andAction:"), target, actionSelector)
	return rv
}

// Returns the index of the item with the specified title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1532244-indexofitemwithtitle?language=objc
func (p_ PopUpButtonCell) IndexOfItemWithTitle(title string) int {
	rv := objc.Call[int](p_, objc.Sel("indexOfItemWithTitle:"), title)
	return rv
}

// Sets up the receiver to display a menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1531648-attachpopupwithframe?language=objc
func (p_ PopUpButtonCell) AttachPopUpWithFrameInView(cellFrame foundation.Rect, controlView IView) {
	objc.Call[objc.Void](p_, objc.Sel("attachPopUpWithFrame:inView:"), cellFrame, objc.Ptr(controlView))
}

// A Boolean value that indicates if the pop-up button links the state of the selected menu item to the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1528446-altersstateofselecteditem?language=objc
func (p_ PopUpButtonCell) AltersStateOfSelectedItem() bool {
	rv := objc.Call[bool](p_, objc.Sel("altersStateOfSelectedItem"))
	return rv
}

// A Boolean value that indicates if the pop-up button links the state of the selected menu item to the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1528446-altersstateofselecteditem?language=objc
func (p_ PopUpButtonCell) SetAltersStateOfSelectedItem(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setAltersStateOfSelectedItem:"), value)
}

// The index of the item last selected by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1534937-indexofselecteditem?language=objc
func (p_ PopUpButtonCell) IndexOfSelectedItem() int {
	rv := objc.Call[int](p_, objc.Sel("indexOfSelectedItem"))
	return rv
}

// The menu item last selected by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1533239-selecteditem?language=objc
func (p_ PopUpButtonCell) SelectedItem() MenuItem {
	rv := objc.Call[MenuItem](p_, objc.Sel("selectedItem"))
	return rv
}

// The position of the arrow displayed on the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1534598-arrowposition?language=objc
func (p_ PopUpButtonCell) ArrowPosition() PopUpArrowPosition {
	rv := objc.Call[PopUpArrowPosition](p_, objc.Sel("arrowPosition"))
	return rv
}

// The position of the arrow displayed on the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1534598-arrowposition?language=objc
func (p_ PopUpButtonCell) SetArrowPosition(value PopUpArrowPosition) {
	objc.Call[objc.Void](p_, objc.Sel("setArrowPosition:"), value)
}

// An array of NSMenuItem objects that represent the items in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1533757-itemarray?language=objc
func (p_ PopUpButtonCell) ItemArray() []MenuItem {
	rv := objc.Call[[]MenuItem](p_, objc.Sel("itemArray"))
	return rv
}

// A Boolean value that indicates if the button automatically enables and disables its items every time a user event occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1530889-autoenablesitems?language=objc
func (p_ PopUpButtonCell) AutoenablesItems() bool {
	rv := objc.Call[bool](p_, objc.Sel("autoenablesItems"))
	return rv
}

// A Boolean value that indicates if the button automatically enables and disables its items every time a user event occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1530889-autoenablesitems?language=objc
func (p_ PopUpButtonCell) SetAutoenablesItems(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setAutoenablesItems:"), value)
}

// A Boolean value that indicates the behavior of the button’s menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1529949-pullsdown?language=objc
func (p_ PopUpButtonCell) PullsDown() bool {
	rv := objc.Call[bool](p_, objc.Sel("pullsDown"))
	return rv
}

// A Boolean value that indicates the behavior of the button’s menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1529949-pullsdown?language=objc
func (p_ PopUpButtonCell) SetPullsDown(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setPullsDown:"), value)
}

// The last item in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1535415-lastitem?language=objc
func (p_ PopUpButtonCell) LastItem() MenuItem {
	rv := objc.Call[MenuItem](p_, objc.Sel("lastItem"))
	return rv
}

// The number of items in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1533633-numberofitems?language=objc
func (p_ PopUpButtonCell) NumberOfItems() int {
	rv := objc.Call[int](p_, objc.Sel("numberOfItems"))
	return rv
}

// The title of the item last selected by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1529952-titleofselecteditem?language=objc
func (p_ PopUpButtonCell) TitleOfSelectedItem() string {
	rv := objc.Call[string](p_, objc.Sel("titleOfSelectedItem"))
	return rv
}

// An array of NSString objects containing the titles of every item in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1532067-itemtitles?language=objc
func (p_ PopUpButtonCell) ItemTitles() []string {
	rv := objc.Call[[]string](p_, objc.Sel("itemTitles"))
	return rv
}

// The edge of the cell from which the menu should pop out when screen conditions are restrictive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1535408-preferrededge?language=objc
func (p_ PopUpButtonCell) PreferredEdge() foundation.RectEdge {
	rv := objc.Call[foundation.RectEdge](p_, objc.Sel("preferredEdge"))
	return rv
}

// The edge of the cell from which the menu should pop out when screen conditions are restrictive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1535408-preferrededge?language=objc
func (p_ PopUpButtonCell) SetPreferredEdge(value foundation.RectEdge) {
	objc.Call[objc.Void](p_, objc.Sel("setPreferredEdge:"), value)
}

// A Boolean value that indicates if the control uses an item from the menu for its own title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1534951-usesitemfrommenu?language=objc
func (p_ PopUpButtonCell) UsesItemFromMenu() bool {
	rv := objc.Call[bool](p_, objc.Sel("usesItemFromMenu"))
	return rv
}

// A Boolean value that indicates if the control uses an item from the menu for its own title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbuttoncell/1534951-usesitemfrommenu?language=objc
func (p_ PopUpButtonCell) SetUsesItemFromMenu(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setUsesItemFromMenu:"), value)
}
