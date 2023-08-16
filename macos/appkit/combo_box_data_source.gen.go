// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxdatasource?language=objc
type PComboBoxDataSource interface {
	// optional
	NumberOfItemsInComboBox(comboBox ComboBox) int
	HasNumberOfItemsInComboBox() bool

	// optional
	ComboBoxObjectValueForItemAtIndex(comboBox ComboBox, index int) objc.IObject
	HasComboBoxObjectValueForItemAtIndex() bool
}

// A concrete type wrapper for the [PComboBoxDataSource] protocol.
type ComboBoxDataSourceWrapper struct {
	objc.Object
}

func (c_ ComboBoxDataSourceWrapper) HasNumberOfItemsInComboBox() bool {
	return c_.RespondsToSelector(objc.Sel("numberOfItemsInComboBox:"))
}

// Returns the number of items that the data source manages for the combo box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxdatasource/1436780-numberofitemsincombobox?language=objc
func (c_ ComboBoxDataSourceWrapper) NumberOfItemsInComboBox(comboBox IComboBox) int {
	rv := objc.Call[int](c_, objc.Sel("numberOfItemsInComboBox:"), objc.Ptr(comboBox))
	return rv
}

func (c_ ComboBoxDataSourceWrapper) HasComboBoxObjectValueForItemAtIndex() bool {
	return c_.RespondsToSelector(objc.Sel("comboBox:objectValueForItemAtIndex:"))
}

// Returns the object that corresponds to the item at the specified index in the combo box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxdatasource/1436753-combobox?language=objc
func (c_ ComboBoxDataSourceWrapper) ComboBoxObjectValueForItemAtIndex(comboBox IComboBox, index int) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("comboBox:objectValueForItemAtIndex:"), objc.Ptr(comboBox), index)
	return rv
}
