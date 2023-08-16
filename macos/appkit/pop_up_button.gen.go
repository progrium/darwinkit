// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PopUpButton] class.
var PopUpButtonClass = _PopUpButtonClass{objc.GetClass("NSPopUpButton")}

type _PopUpButtonClass struct {
	objc.Class
}

// An interface definition for the [PopUpButton] class.
type IPopUpButton interface {
	IButton
	ItemAtIndex(index int) MenuItem
	IndexOfItemWithTag(tag int) int
	InsertItemWithTitleAtIndex(title string, index int)
	RemoveAllItems()
	RemoveItemAtIndex(index int)
	RemoveItemWithTitle(title string)
	SelectItemWithTag(tag int) bool
	AddItemWithTitle(title string)
	SelectItemWithTitle(title string)
	ItemWithTitle(title string) MenuItem
	IndexOfItemWithRepresentedObject(obj objc.IObject) int
	ItemTitleAtIndex(index int) string
	SelectItem(item IMenuItem)
	IndexOfItem(item IMenuItem) int
	AddItemsWithTitles(itemTitles []string)
	SelectItemAtIndex(index int)
	SynchronizeTitleAndSelectedItem()
	IndexOfItemWithTargetAndAction(target objc.IObject, actionSelector objc.Selector) int
	IndexOfItemWithTitle(title string) int
	IndexOfSelectedItem() int
	SelectedItem() MenuItem
	ItemArray() []MenuItem
	AutoenablesItems() bool
	SetAutoenablesItems(value bool)
	PullsDown() bool
	SetPullsDown(value bool)
	LastItem() MenuItem
	NumberOfItems() int
	SelectedTag() int
	TitleOfSelectedItem() string
	ItemTitles() []string
	PreferredEdge() foundation.RectEdge
	SetPreferredEdge(value foundation.RectEdge)
}

// A control for selecting an item from a list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton?language=objc
type PopUpButton struct {
	Button
}

func PopUpButtonFrom(ptr unsafe.Pointer) PopUpButton {
	return PopUpButton{
		Button: ButtonFrom(ptr),
	}
}

func (p_ PopUpButton) InitWithFramePullsDown(buttonFrame foundation.Rect, flag bool) PopUpButton {
	rv := objc.Call[PopUpButton](p_, objc.Sel("initWithFrame:pullsDown:"), buttonFrame, flag)
	return rv
}

// Returns an NSPopUpButton object initialized to the specified dimensions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1524562-initwithframe?language=objc
func PopUpButton_InitWithFramePullsDown(buttonFrame foundation.Rect, flag bool) PopUpButton {
	return PopUpButtonClass.Alloc().InitWithFramePullsDown(buttonFrame, flag)
}

func (pc _PopUpButtonClass) Alloc() PopUpButton {
	rv := objc.Call[PopUpButton](pc, objc.Sel("alloc"))
	return rv
}

func PopUpButton_Alloc() PopUpButton {
	return PopUpButtonClass.Alloc()
}

func (pc _PopUpButtonClass) New() PopUpButton {
	rv := objc.Call[PopUpButton](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPopUpButton() PopUpButton {
	return PopUpButtonClass.New()
}

func (p_ PopUpButton) Init() PopUpButton {
	rv := objc.Call[PopUpButton](p_, objc.Sel("init"))
	return rv
}

func (pc _PopUpButtonClass) CheckboxWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) PopUpButton {
	rv := objc.Call[PopUpButton](pc, objc.Sel("checkboxWithTitle:target:action:"), title, target, action)
	return rv
}

// Creates a standard checkbox with the title you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1644525-checkboxwithtitle?language=objc
func PopUpButton_CheckboxWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) PopUpButton {
	return PopUpButtonClass.CheckboxWithTitleTargetAction(title, target, action)
}

func (pc _PopUpButtonClass) ButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) PopUpButton {
	rv := objc.Call[PopUpButton](pc, objc.Sel("buttonWithTitle:target:action:"), title, target, action)
	return rv
}

// Creates a standard push button with the title you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1644256-buttonwithtitle?language=objc
func PopUpButton_ButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) PopUpButton {
	return PopUpButtonClass.ButtonWithTitleTargetAction(title, target, action)
}

func (pc _PopUpButtonClass) ButtonWithImageTargetAction(image IImage, target objc.IObject, action objc.Selector) PopUpButton {
	rv := objc.Call[PopUpButton](pc, objc.Sel("buttonWithImage:target:action:"), objc.Ptr(image), target, action)
	return rv
}

// Creates a standard push button with the image you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1644659-buttonwithimage?language=objc
func PopUpButton_ButtonWithImageTargetAction(image IImage, target objc.IObject, action objc.Selector) PopUpButton {
	return PopUpButtonClass.ButtonWithImageTargetAction(image, target, action)
}

func (pc _PopUpButtonClass) RadioButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) PopUpButton {
	rv := objc.Call[PopUpButton](pc, objc.Sel("radioButtonWithTitle:target:action:"), title, target, action)
	return rv
}

// Creates a standard radio button with the title you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1644340-radiobuttonwithtitle?language=objc
func PopUpButton_RadioButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) PopUpButton {
	return PopUpButtonClass.RadioButtonWithTitleTargetAction(title, target, action)
}

func (p_ PopUpButton) InitWithFrame(frameRect foundation.Rect) PopUpButton {
	rv := objc.Call[PopUpButton](p_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func PopUpButton_InitWithFrame(frameRect foundation.Rect) PopUpButton {
	return PopUpButtonClass.Alloc().InitWithFrame(frameRect)
}

// Returns the menu item at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1535860-itematindex?language=objc
func (p_ PopUpButton) ItemAtIndex(index int) MenuItem {
	rv := objc.Call[MenuItem](p_, objc.Sel("itemAtIndex:"), index)
	return rv
}

// Returns the index of the menu item with the specified tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1534140-indexofitemwithtag?language=objc
func (p_ PopUpButton) IndexOfItemWithTag(tag int) int {
	rv := objc.Call[int](p_, objc.Sel("indexOfItemWithTag:"), tag)
	return rv
}

// Inserts an item at the specified position in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1533750-insertitemwithtitle?language=objc
func (p_ PopUpButton) InsertItemWithTitleAtIndex(title string, index int) {
	objc.Call[objc.Void](p_, objc.Sel("insertItemWithTitle:atIndex:"), title, index)
}

// Removes all items in the receiver’s item menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1534430-removeallitems?language=objc
func (p_ PopUpButton) RemoveAllItems() {
	objc.Call[objc.Void](p_, objc.Sel("removeAllItems"))
}

// Removes the item at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1525241-removeitematindex?language=objc
func (p_ PopUpButton) RemoveItemAtIndex(index int) {
	objc.Call[objc.Void](p_, objc.Sel("removeItemAtIndex:"), index)
}

// Removes the item with the specified title from the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1530675-removeitemwithtitle?language=objc
func (p_ PopUpButton) RemoveItemWithTitle(title string) {
	objc.Call[objc.Void](p_, objc.Sel("removeItemWithTitle:"), title)
}

// Selects the menu item with the specified tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1528181-selectitemwithtag?language=objc
func (p_ PopUpButton) SelectItemWithTag(tag int) bool {
	rv := objc.Call[bool](p_, objc.Sel("selectItemWithTag:"), tag)
	return rv
}

// Adds an item with the specified title to the end of the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1529500-additemwithtitle?language=objc
func (p_ PopUpButton) AddItemWithTitle(title string) {
	objc.Call[objc.Void](p_, objc.Sel("addItemWithTitle:"), title)
}

// Selects the item with the specified title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1525666-selectitemwithtitle?language=objc
func (p_ PopUpButton) SelectItemWithTitle(title string) {
	objc.Call[objc.Void](p_, objc.Sel("selectItemWithTitle:"), title)
}

// Returns the menu item with the specified title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1524680-itemwithtitle?language=objc
func (p_ PopUpButton) ItemWithTitle(title string) MenuItem {
	rv := objc.Call[MenuItem](p_, objc.Sel("itemWithTitle:"), title)
	return rv
}

// Returns the index of the menu item that holds the specified represented object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1533993-indexofitemwithrepresentedobject?language=objc
func (p_ PopUpButton) IndexOfItemWithRepresentedObject(obj objc.IObject) int {
	rv := objc.Call[int](p_, objc.Sel("indexOfItemWithRepresentedObject:"), obj)
	return rv
}

// Returns the title of the item at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1524908-itemtitleatindex?language=objc
func (p_ PopUpButton) ItemTitleAtIndex(index int) string {
	rv := objc.Call[string](p_, objc.Sel("itemTitleAtIndex:"), index)
	return rv
}

// Selects the specified menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1527162-selectitem?language=objc
func (p_ PopUpButton) SelectItem(item IMenuItem) {
	objc.Call[objc.Void](p_, objc.Sel("selectItem:"), objc.Ptr(item))
}

// Returns the index of the specified menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1529708-indexofitem?language=objc
func (p_ PopUpButton) IndexOfItem(item IMenuItem) int {
	rv := objc.Call[int](p_, objc.Sel("indexOfItem:"), objc.Ptr(item))
	return rv
}

// Adds multiple items to the end of the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1532518-additemswithtitles?language=objc
func (p_ PopUpButton) AddItemsWithTitles(itemTitles []string) {
	objc.Call[objc.Void](p_, objc.Sel("addItemsWithTitles:"), itemTitles)
}

// Selects the item in the menu at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1534906-selectitematindex?language=objc
func (p_ PopUpButton) SelectItemAtIndex(index int) {
	objc.Call[objc.Void](p_, objc.Sel("selectItemAtIndex:"), index)
}

// Ensures that the item being displayed by the receiver agrees with the selected item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1527124-synchronizetitleandselecteditem?language=objc
func (p_ PopUpButton) SynchronizeTitleAndSelectedItem() {
	objc.Call[objc.Void](p_, objc.Sel("synchronizeTitleAndSelectedItem"))
}

// Returns the index of the menu item with the specified target and action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1535747-indexofitemwithtarget?language=objc
func (p_ PopUpButton) IndexOfItemWithTargetAndAction(target objc.IObject, actionSelector objc.Selector) int {
	rv := objc.Call[int](p_, objc.Sel("indexOfItemWithTarget:andAction:"), target, actionSelector)
	return rv
}

// Returns the index of the item with the specified title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1535406-indexofitemwithtitle?language=objc
func (p_ PopUpButton) IndexOfItemWithTitle(title string) int {
	rv := objc.Call[int](p_, objc.Sel("indexOfItemWithTitle:"), title)
	return rv
}

// The index of the item that was last selected by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1534134-indexofselecteditem?language=objc
func (p_ PopUpButton) IndexOfSelectedItem() int {
	rv := objc.Call[int](p_, objc.Sel("indexOfSelectedItem"))
	return rv
}

// The menu item that was last selected by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1526197-selecteditem?language=objc
func (p_ PopUpButton) SelectedItem() MenuItem {
	rv := objc.Call[MenuItem](p_, objc.Sel("selectedItem"))
	return rv
}

// The array of menu item objects associated with the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1535361-itemarray?language=objc
func (p_ PopUpButton) ItemArray() []MenuItem {
	rv := objc.Call[[]MenuItem](p_, objc.Sel("itemArray"))
	return rv
}

// A Boolean value indicating whether the button enables and disables its items every time a user event occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1530543-autoenablesitems?language=objc
func (p_ PopUpButton) AutoenablesItems() bool {
	rv := objc.Call[bool](p_, objc.Sel("autoenablesItems"))
	return rv
}

// A Boolean value indicating whether the button enables and disables its items every time a user event occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1530543-autoenablesitems?language=objc
func (p_ PopUpButton) SetAutoenablesItems(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setAutoenablesItems:"), value)
}

// A Boolean value indicating whether the button displays a pull-down or pop-up menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1532070-pullsdown?language=objc
func (p_ PopUpButton) PullsDown() bool {
	rv := objc.Call[bool](p_, objc.Sel("pullsDown"))
	return rv
}

// A Boolean value indicating whether the button displays a pull-down or pop-up menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1532070-pullsdown?language=objc
func (p_ PopUpButton) SetPullsDown(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setPullsDown:"), value)
}

// The last item in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1535371-lastitem?language=objc
func (p_ PopUpButton) LastItem() MenuItem {
	rv := objc.Call[MenuItem](p_, objc.Sel("lastItem"))
	return rv
}

// The number of items in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1534959-numberofitems?language=objc
func (p_ PopUpButton) NumberOfItems() int {
	rv := objc.Call[int](p_, objc.Sel("numberOfItems"))
	return rv
}

// The tag of the menu item that was last selected by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1577134-selectedtag?language=objc
func (p_ PopUpButton) SelectedTag() int {
	rv := objc.Call[int](p_, objc.Sel("selectedTag"))
	return rv
}

// The title of the item that was last selected by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1534038-titleofselecteditem?language=objc
func (p_ PopUpButton) TitleOfSelectedItem() string {
	rv := objc.Call[string](p_, objc.Sel("titleOfSelectedItem"))
	return rv
}

// An array of strings corresponding to the titles of the items in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1529271-itemtitles?language=objc
func (p_ PopUpButton) ItemTitles() []string {
	rv := objc.Call[[]string](p_, objc.Sel("itemTitles"))
	return rv
}

// The edge of the button on which to display the menu when screen space is constrained. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1535345-preferrededge?language=objc
func (p_ PopUpButton) PreferredEdge() foundation.RectEdge {
	rv := objc.Call[foundation.RectEdge](p_, objc.Sel("preferredEdge"))
	return rv
}

// The edge of the button on which to display the menu when screen space is constrained. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopupbutton/1535345-preferrededge?language=objc
func (p_ PopUpButton) SetPreferredEdge(value foundation.RectEdge) {
	objc.Call[objc.Void](p_, objc.Sel("setPreferredEdge:"), value)
}
