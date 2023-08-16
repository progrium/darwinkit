// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcelldatasource?language=objc
type PComboBoxCellDataSource interface {
	// optional
	NumberOfItemsInComboBoxCell(comboBoxCell ComboBoxCell) int
	HasNumberOfItemsInComboBoxCell() bool

	// optional
	ComboBoxCellObjectValueForItemAtIndex(comboBoxCell ComboBoxCell, index int) objc.IObject
	HasComboBoxCellObjectValueForItemAtIndex() bool
}

// A concrete type wrapper for the [PComboBoxCellDataSource] protocol.
type ComboBoxCellDataSourceWrapper struct {
	objc.Object
}

func (c_ ComboBoxCellDataSourceWrapper) HasNumberOfItemsInComboBoxCell() bool {
	return c_.RespondsToSelector(objc.Sel("numberOfItemsInComboBoxCell:"))
}

// Returns the number of items managed for the combo box cell by your data source object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcelldatasource/1410302-numberofitemsincomboboxcell?language=objc
func (c_ ComboBoxCellDataSourceWrapper) NumberOfItemsInComboBoxCell(comboBoxCell IComboBoxCell) int {
	rv := objc.Call[int](c_, objc.Sel("numberOfItemsInComboBoxCell:"), objc.Ptr(comboBoxCell))
	return rv
}

func (c_ ComboBoxCellDataSourceWrapper) HasComboBoxCellObjectValueForItemAtIndex() bool {
	return c_.RespondsToSelector(objc.Sel("comboBoxCell:objectValueForItemAtIndex:"))
}

// Returns the object that corresponds to the item at the given index in the combo box cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcelldatasource/1410258-comboboxcell?language=objc
func (c_ ComboBoxCellDataSourceWrapper) ComboBoxCellObjectValueForItemAtIndex(comboBoxCell IComboBoxCell, index int) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("comboBoxCell:objectValueForItemAtIndex:"), objc.Ptr(comboBoxCell), index)
	return rv
}
