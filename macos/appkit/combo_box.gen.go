// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ComboBox] class.
var ComboBoxClass = _ComboBoxClass{objc.GetClass("NSComboBox")}

type _ComboBoxClass struct {
	objc.Class
}

// An interface definition for the [ComboBox] class.
type IComboBox interface {
	ITextField
	RemoveItemWithObjectValue(object objc.IObject)
	NoteNumberOfItemsChanged()
	IndexOfItemWithObjectValue(object objc.IObject) int
	InsertItemWithObjectValueAtIndex(object objc.IObject, index int)
	ScrollItemAtIndexToVisible(index int)
	RemoveAllItems()
	RemoveItemAtIndex(index int)
	AddItemWithObjectValue(object objc.IObject)
	AddItemsWithObjectValues(objects []objc.IObject)
	DeselectItemAtIndex(index int)
	ScrollItemAtIndexToTop(index int)
	SelectItemWithObjectValue(object objc.IObject)
	SelectItemAtIndex(index int)
	ItemObjectValueAtIndex(index int) objc.Object
	ReloadData()
	DataSource() ComboBoxDataSourceWrapper
	SetDataSource(value PComboBoxDataSource)
	SetDataSourceObject(valueObject objc.IObject)
	IndexOfSelectedItem() int
	ItemHeight() float64
	SetItemHeight(value float64)
	ObjectValueOfSelectedItem() objc.Object
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	Completes() bool
	SetCompletes(value bool)
	UsesDataSource() bool
	SetUsesDataSource(value bool)
	ObjectValues() []objc.Object
	NumberOfItems() int
	NumberOfVisibleItems() int
	SetNumberOfVisibleItems(value int)
	IsButtonBordered() bool
	SetButtonBordered(value bool)
	IntercellSpacing() foundation.Size
	SetIntercellSpacing(value foundation.Size)
}

// A view that displays a list of values in a pop-up menu where the user selects a value or types in a custom value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox?language=objc
type ComboBox struct {
	TextField
}

func ComboBoxFrom(ptr unsafe.Pointer) ComboBox {
	return ComboBox{
		TextField: TextFieldFrom(ptr),
	}
}

func (cc _ComboBoxClass) Alloc() ComboBox {
	rv := objc.Call[ComboBox](cc, objc.Sel("alloc"))
	return rv
}

func ComboBox_Alloc() ComboBox {
	return ComboBoxClass.Alloc()
}

func (cc _ComboBoxClass) New() ComboBox {
	rv := objc.Call[ComboBox](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewComboBox() ComboBox {
	return ComboBoxClass.New()
}

func (c_ ComboBox) Init() ComboBox {
	rv := objc.Call[ComboBox](c_, objc.Sel("init"))
	return rv
}

func (cc _ComboBoxClass) LabelWithAttributedString(attributedStringValue foundation.IAttributedString) ComboBox {
	rv := objc.Call[ComboBox](cc, objc.Sel("labelWithAttributedString:"), objc.Ptr(attributedStringValue))
	return rv
}

// Creates a text field for use as a static label that displays styled text, doesn’t wrap, and doesn’t have selectable text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644658-labelwithattributedstring?language=objc
func ComboBox_LabelWithAttributedString(attributedStringValue foundation.IAttributedString) ComboBox {
	return ComboBoxClass.LabelWithAttributedString(attributedStringValue)
}

func (cc _ComboBoxClass) LabelWithString(stringValue string) ComboBox {
	rv := objc.Call[ComboBox](cc, objc.Sel("labelWithString:"), stringValue)
	return rv
}

// Initializes a text field for use as a static label that uses the system default font, doesn’t wrap, and doesn’t have selectable text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644377-labelwithstring?language=objc
func ComboBox_LabelWithString(stringValue string) ComboBox {
	return ComboBoxClass.LabelWithString(stringValue)
}

func (cc _ComboBoxClass) WrappingLabelWithString(stringValue string) ComboBox {
	rv := objc.Call[ComboBox](cc, objc.Sel("wrappingLabelWithString:"), stringValue)
	return rv
}

// Initializes a text field for use as a multiline static label with selectable text that uses the system default font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644543-wrappinglabelwithstring?language=objc
func ComboBox_WrappingLabelWithString(stringValue string) ComboBox {
	return ComboBoxClass.WrappingLabelWithString(stringValue)
}

func (cc _ComboBoxClass) TextFieldWithString(stringValue string) ComboBox {
	rv := objc.Call[ComboBox](cc, objc.Sel("textFieldWithString:"), stringValue)
	return rv
}

// Initializes a single-line editable text field for user input using the system default font and standard visual appearance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644706-textfieldwithstring?language=objc
func ComboBox_TextFieldWithString(stringValue string) ComboBox {
	return ComboBoxClass.TextFieldWithString(stringValue)
}

func (c_ ComboBox) InitWithFrame(frameRect foundation.Rect) ComboBox {
	rv := objc.Call[ComboBox](c_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func ComboBox_InitWithFrame(frameRect foundation.Rect) ComboBox {
	return ComboBoxClass.Alloc().InitWithFrame(frameRect)
}

// Removes all occurrences of the given object from the receiver’s internal item list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436759-removeitemwithobjectvalue?language=objc
func (c_ ComboBox) RemoveItemWithObjectValue(object objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("removeItemWithObjectValue:"), object)
}

// Informs the receiver that the number of items in its data source has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436723-notenumberofitemschanged?language=objc
func (c_ ComboBox) NoteNumberOfItemsChanged() {
	objc.Call[objc.Void](c_, objc.Sel("noteNumberOfItemsChanged"))
}

// Searches the receiver’s internal item list for the specified object and returns the lowest matching index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436707-indexofitemwithobjectvalue?language=objc
func (c_ ComboBox) IndexOfItemWithObjectValue(object objc.IObject) int {
	rv := objc.Call[int](c_, objc.Sel("indexOfItemWithObjectValue:"), object)
	return rv
}

// Inserts an object at the specified location in the receiver’s internal item list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436703-insertitemwithobjectvalue?language=objc
func (c_ ComboBox) InsertItemWithObjectValueAtIndex(object objc.IObject, index int) {
	objc.Call[objc.Void](c_, objc.Sel("insertItemWithObjectValue:atIndex:"), object, index)
}

// Scrolls the receiver’s pop-up list vertically so that the item at the specified index is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436776-scrollitematindextovisible?language=objc
func (c_ ComboBox) ScrollItemAtIndexToVisible(index int) {
	objc.Call[objc.Void](c_, objc.Sel("scrollItemAtIndexToVisible:"), index)
}

// Removes all items from the receiver’s internal item list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436721-removeallitems?language=objc
func (c_ ComboBox) RemoveAllItems() {
	objc.Call[objc.Void](c_, objc.Sel("removeAllItems"))
}

// Removes the object at the specified location from the receiver’s internal item list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436761-removeitematindex?language=objc
func (c_ ComboBox) RemoveItemAtIndex(index int) {
	objc.Call[objc.Void](c_, objc.Sel("removeItemAtIndex:"), index)
}

// Adds an object to the end of the receiver’s internal item list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436739-additemwithobjectvalue?language=objc
func (c_ ComboBox) AddItemWithObjectValue(object objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("addItemWithObjectValue:"), object)
}

// Adds multiple objects to the end of the receiver’s internal item list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436751-additemswithobjectvalues?language=objc
func (c_ ComboBox) AddItemsWithObjectValues(objects []objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("addItemsWithObjectValues:"), objects)
}

// Deselects the pop-up list item at the specified index if it’s selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436699-deselectitematindex?language=objc
func (c_ ComboBox) DeselectItemAtIndex(index int) {
	objc.Call[objc.Void](c_, objc.Sel("deselectItemAtIndex:"), index)
}

// Scrolls the receiver’s pop-up list vertically so that the item at the specified index is as close to the top as possible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436778-scrollitematindextotop?language=objc
func (c_ ComboBox) ScrollItemAtIndexToTop(index int) {
	objc.Call[objc.Void](c_, objc.Sel("scrollItemAtIndexToTop:"), index)
}

// Selects the first pop-up list item that corresponds to the given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436737-selectitemwithobjectvalue?language=objc
func (c_ ComboBox) SelectItemWithObjectValue(object objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("selectItemWithObjectValue:"), object)
}

// Selects the pop-up list row at the given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436735-selectitematindex?language=objc
func (c_ ComboBox) SelectItemAtIndex(index int) {
	objc.Call[objc.Void](c_, objc.Sel("selectItemAtIndex:"), index)
}

// Returns the object located at the given index within the receiver’s internal item list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436731-itemobjectvalueatindex?language=objc
func (c_ ComboBox) ItemObjectValueAtIndex(index int) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("itemObjectValueAtIndex:"), index)
	return rv
}

// Marks the receiver as needing redisplay, so that it will reload the data for visible pop-up items and draw the new values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436717-reloaddata?language=objc
func (c_ ComboBox) ReloadData() {
	objc.Call[objc.Void](c_, objc.Sel("reloadData"))
}

// The object that provides the item data for the combo box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436729-datasource?language=objc
func (c_ ComboBox) DataSource() ComboBoxDataSourceWrapper {
	rv := objc.Call[ComboBoxDataSourceWrapper](c_, objc.Sel("dataSource"))
	return rv
}

// The object that provides the item data for the combo box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436729-datasource?language=objc
func (c_ ComboBox) SetDataSource(value PComboBoxDataSource) {
	po0 := objc.WrapAsProtocol("NSComboBoxDataSource", value)
	objc.Call[objc.Void](c_, objc.Sel("setDataSource:"), po0)
}

// The object that provides the item data for the combo box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436729-datasource?language=objc
func (c_ ComboBox) SetDataSourceObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setDataSource:"), objc.Ptr(valueObject))
}

// The index of the last item selected from the pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436701-indexofselecteditem?language=objc
func (c_ ComboBox) IndexOfSelectedItem() int {
	rv := objc.Call[int](c_, objc.Sel("indexOfSelectedItem"))
	return rv
}

// The height of each item in the pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436767-itemheight?language=objc
func (c_ ComboBox) ItemHeight() float64 {
	rv := objc.Call[float64](c_, objc.Sel("itemHeight"))
	return rv
}

// The height of each item in the pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436767-itemheight?language=objc
func (c_ ComboBox) SetItemHeight(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setItemHeight:"), value)
}

// The object corresponding to the last item selected from the pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436743-objectvalueofselecteditem?language=objc
func (c_ ComboBox) ObjectValueOfSelectedItem() objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("objectValueOfSelectedItem"))
	return rv
}

// A Boolean value indicating whether the combo box has a vertical scroller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436705-hasverticalscroller?language=objc
func (c_ ComboBox) HasVerticalScroller() bool {
	rv := objc.Call[bool](c_, objc.Sel("hasVerticalScroller"))
	return rv
}

// A Boolean value indicating whether the combo box has a vertical scroller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436705-hasverticalscroller?language=objc
func (c_ ComboBox) SetHasVerticalScroller(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setHasVerticalScroller:"), value)
}

// A Boolean value indicating whether the combo box tries to complete what the user types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436749-completes?language=objc
func (c_ ComboBox) Completes() bool {
	rv := objc.Call[bool](c_, objc.Sel("completes"))
	return rv
}

// A Boolean value indicating whether the combo box tries to complete what the user types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436749-completes?language=objc
func (c_ ComboBox) SetCompletes(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setCompletes:"), value)
}

// A Boolean value indicating whether the combo box retrieves its items from a data source object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436727-usesdatasource?language=objc
func (c_ ComboBox) UsesDataSource() bool {
	rv := objc.Call[bool](c_, objc.Sel("usesDataSource"))
	return rv
}

// A Boolean value indicating whether the combo box retrieves its items from a data source object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436727-usesdatasource?language=objc
func (c_ ComboBox) SetUsesDataSource(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setUsesDataSource:"), value)
}

// An array of the items from the combo box’s internal list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436709-objectvalues?language=objc
func (c_ ComboBox) ObjectValues() []objc.Object {
	rv := objc.Call[[]objc.Object](c_, objc.Sel("objectValues"))
	return rv
}

// The total number of items in the pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436747-numberofitems?language=objc
func (c_ ComboBox) NumberOfItems() int {
	rv := objc.Call[int](c_, objc.Sel("numberOfItems"))
	return rv
}

// The maximum number of visible items to display in the pop-up list at one time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436741-numberofvisibleitems?language=objc
func (c_ ComboBox) NumberOfVisibleItems() int {
	rv := objc.Call[int](c_, objc.Sel("numberOfVisibleItems"))
	return rv
}

// The maximum number of visible items to display in the pop-up list at one time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436741-numberofvisibleitems?language=objc
func (c_ ComboBox) SetNumberOfVisibleItems(value int) {
	objc.Call[objc.Void](c_, objc.Sel("setNumberOfVisibleItems:"), value)
}

// A Boolean value indicating whether the combo box displays a border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436711-buttonbordered?language=objc
func (c_ ComboBox) IsButtonBordered() bool {
	rv := objc.Call[bool](c_, objc.Sel("isButtonBordered"))
	return rv
}

// A Boolean value indicating whether the combo box displays a border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436711-buttonbordered?language=objc
func (c_ ComboBox) SetButtonBordered(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setButtonBordered:"), value)
}

// The horizontal and vertical spacing between cells in the pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436771-intercellspacing?language=objc
func (c_ ComboBox) IntercellSpacing() foundation.Size {
	rv := objc.Call[foundation.Size](c_, objc.Sel("intercellSpacing"))
	return rv
}

// The horizontal and vertical spacing between cells in the pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobox/1436771-intercellspacing?language=objc
func (c_ ComboBox) SetIntercellSpacing(value foundation.Size) {
	objc.Call[objc.Void](c_, objc.Sel("setIntercellSpacing:"), value)
}
