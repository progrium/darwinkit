// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ComboBoxClass = _ComboBoxClass{objc.GetClass("NSComboBox")}

type _ComboBoxClass struct {
	objc.Class
}

type IComboBox interface {
	ITextField
	AddItemsWithObjectValues(objects []objc.IObject)
	AddItemWithObjectValue(object objc.IObject)
	InsertItemWithObjectValueAtIndex(object objc.IObject, index int)
	RemoveAllItems()
	RemoveItemAtIndex(index int)
	RemoveItemWithObjectValue(object objc.IObject)
	IndexOfItemWithObjectValue(object objc.IObject) int
	ItemObjectValueAtIndex(index int) objc.Object
	NoteNumberOfItemsChanged()
	ReloadData()
	ScrollItemAtIndexToTop(index int)
	ScrollItemAtIndexToVisible(index int)
	DeselectItemAtIndex(index int)
	SelectItemAtIndex(index int)
	SelectItemWithObjectValue(object objc.IObject)
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	IntercellSpacing() foundation.Size
	SetIntercellSpacing(value foundation.Size)
	IsButtonBordered() bool
	SetButtonBordered(value bool)
	ItemHeight() float64
	SetItemHeight(value float64)
	NumberOfVisibleItems() int
	SetNumberOfVisibleItems(value int)
	DataSource() ComboBoxDataSourceWrapper
	SetDataSource(value IComboBoxDataSource)
	SetDataSource0(value objc.IObject)
	UsesDataSource() bool
	SetUsesDataSource(value bool)
	ObjectValues() []objc.Object
	NumberOfItems() int
	IndexOfSelectedItem() int
	ObjectValueOfSelectedItem() objc.Object
	Completes() bool
	SetCompletes(value bool)
}

type ComboBox struct {
	TextField
}

func MakeComboBox(ptr unsafe.Pointer) ComboBox {
	return ComboBox{
		TextField: MakeTextField(ptr),
	}
}

func (cc _ComboBoxClass) LabelWithAttributedString(attributedStringValue foundation.IAttributedString) ComboBox {
	rv := objc.CallMethod[ComboBox](cc, objc.GetSelector("labelWithAttributedString:"), objc.ExtractPtr(attributedStringValue))
	return rv
}

func ComboBox_LabelWithAttributedString(attributedStringValue foundation.IAttributedString) ComboBox {
	return ComboBoxClass.LabelWithAttributedString(attributedStringValue)
}

func (cc _ComboBoxClass) LabelWithString(stringValue string) ComboBox {
	rv := objc.CallMethod[ComboBox](cc, objc.GetSelector("labelWithString:"), stringValue)
	return rv
}

func ComboBox_LabelWithString(stringValue string) ComboBox {
	return ComboBoxClass.LabelWithString(stringValue)
}

func (cc _ComboBoxClass) TextFieldWithString(stringValue string) ComboBox {
	rv := objc.CallMethod[ComboBox](cc, objc.GetSelector("textFieldWithString:"), stringValue)
	return rv
}

func ComboBox_TextFieldWithString(stringValue string) ComboBox {
	return ComboBoxClass.TextFieldWithString(stringValue)
}

func (cc _ComboBoxClass) WrappingLabelWithString(stringValue string) ComboBox {
	rv := objc.CallMethod[ComboBox](cc, objc.GetSelector("wrappingLabelWithString:"), stringValue)
	return rv
}

func ComboBox_WrappingLabelWithString(stringValue string) ComboBox {
	return ComboBoxClass.WrappingLabelWithString(stringValue)
}

func (c_ ComboBox) InitWithFrame(frameRect foundation.Rect) ComboBox {
	rv := objc.CallMethod[ComboBox](c_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func ComboBox_InitWithFrame(frameRect foundation.Rect) ComboBox {
	return ComboBoxClass.Alloc().InitWithFrame(frameRect)
}

func (c_ ComboBox) Init() ComboBox {
	rv := objc.CallMethod[ComboBox](c_, objc.GetSelector("init"))
	return rv
}

func ComboBox_Init() ComboBox {
	return ComboBoxClass.Alloc().Init()
}

func (cc _ComboBoxClass) Alloc() ComboBox {
	rv := objc.CallMethod[ComboBox](cc, objc.GetSelector("alloc"))
	return rv
}

func ComboBox_Alloc() ComboBox {
	return ComboBoxClass.Alloc()
}

func (cc _ComboBoxClass) New() ComboBox {
	rv := objc.CallMethod[ComboBox](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewComboBox() ComboBox {
	return ComboBoxClass.New()
}

func ComboBox_New() ComboBox {
	return ComboBoxClass.New()
}

func (c_ ComboBox) AddItemsWithObjectValues(objects []objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("addItemsWithObjectValues:"), objects)
}

func (c_ ComboBox) AddItemWithObjectValue(object objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("addItemWithObjectValue:"), objc.ExtractPtr(object))
}

func (c_ ComboBox) InsertItemWithObjectValueAtIndex(object objc.IObject, index int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("insertItemWithObjectValue:atIndex:"), objc.ExtractPtr(object), index)
}

func (c_ ComboBox) RemoveAllItems() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("removeAllItems"))
}

func (c_ ComboBox) RemoveItemAtIndex(index int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("removeItemAtIndex:"), index)
}

func (c_ ComboBox) RemoveItemWithObjectValue(object objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("removeItemWithObjectValue:"), objc.ExtractPtr(object))
}

func (c_ ComboBox) IndexOfItemWithObjectValue(object objc.IObject) int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("indexOfItemWithObjectValue:"), objc.ExtractPtr(object))
	return rv
}

func (c_ ComboBox) ItemObjectValueAtIndex(index int) objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.GetSelector("itemObjectValueAtIndex:"), index)
	return rv
}

func (c_ ComboBox) NoteNumberOfItemsChanged() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("noteNumberOfItemsChanged"))
}

func (c_ ComboBox) ReloadData() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("reloadData"))
}

func (c_ ComboBox) ScrollItemAtIndexToTop(index int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("scrollItemAtIndexToTop:"), index)
}

func (c_ ComboBox) ScrollItemAtIndexToVisible(index int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("scrollItemAtIndexToVisible:"), index)
}

func (c_ ComboBox) DeselectItemAtIndex(index int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("deselectItemAtIndex:"), index)
}

func (c_ ComboBox) SelectItemAtIndex(index int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("selectItemAtIndex:"), index)
}

func (c_ ComboBox) SelectItemWithObjectValue(object objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("selectItemWithObjectValue:"), objc.ExtractPtr(object))
}

func (c_ ComboBox) HasVerticalScroller() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("hasVerticalScroller"))
	return rv
}

func (c_ ComboBox) SetHasVerticalScroller(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setHasVerticalScroller:"), value)
}

func (c_ ComboBox) IntercellSpacing() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("intercellSpacing"))
	return rv
}

func (c_ ComboBox) SetIntercellSpacing(value foundation.Size) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setIntercellSpacing:"), value)
}

func (c_ ComboBox) IsButtonBordered() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isButtonBordered"))
	return rv
}

func (c_ ComboBox) SetButtonBordered(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setButtonBordered:"), value)
}

func (c_ ComboBox) ItemHeight() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("itemHeight"))
	return rv
}

func (c_ ComboBox) SetItemHeight(value float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setItemHeight:"), value)
}

func (c_ ComboBox) NumberOfVisibleItems() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("numberOfVisibleItems"))
	return rv
}

func (c_ ComboBox) SetNumberOfVisibleItems(value int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setNumberOfVisibleItems:"), value)
}

func (c_ ComboBox) DataSource() ComboBoxDataSourceWrapper {
	rv := objc.CallMethod[ComboBoxDataSourceWrapper](c_, objc.GetSelector("dataSource"))
	return rv
}

func (c_ ComboBox) SetDataSource(value IComboBoxDataSource) {
	po := objc.WrapAsProtocol("NSComboBoxDataSource", value)
	objc.SetAssociatedObject(c_, objc.AssociationKey("setDataSource"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setDataSource:"), po)
}

func (c_ ComboBox) SetDataSource0(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setDataSource:"), objc.ExtractPtr(value))
}

func (c_ ComboBox) UsesDataSource() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("usesDataSource"))
	return rv
}

func (c_ ComboBox) SetUsesDataSource(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setUsesDataSource:"), value)
}

func (c_ ComboBox) ObjectValues() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](c_, objc.GetSelector("objectValues"))
	// TODO: convert slice items...
	return rv
}

func (c_ ComboBox) NumberOfItems() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("numberOfItems"))
	return rv
}

func (c_ ComboBox) IndexOfSelectedItem() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("indexOfSelectedItem"))
	return rv
}

func (c_ ComboBox) ObjectValueOfSelectedItem() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.GetSelector("objectValueOfSelectedItem"))
	return rv
}

func (c_ ComboBox) Completes() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("completes"))
	return rv
}

func (c_ ComboBox) SetCompletes(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setCompletes:"), value)
}
