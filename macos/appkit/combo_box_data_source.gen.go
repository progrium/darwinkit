// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type IComboBoxDataSource interface {
	ImplementsComboBoxCompletedString() bool
	// optional
	ComboBoxCompletedString(comboBox ComboBox, string_ string) string
	ImplementsComboBoxIndexOfItemWithStringValue() bool
	// optional
	ComboBoxIndexOfItemWithStringValue(comboBox ComboBox, string_ string) uint
	ImplementsComboBoxObjectValueForItemAtIndex() bool
	// optional
	ComboBoxObjectValueForItemAtIndex(comboBox ComboBox, index int) objc.IObject
	ImplementsNumberOfItemsInComboBox() bool
	// optional
	NumberOfItemsInComboBox(comboBox ComboBox) int
}

type ComboBoxDataSourceWrapper struct {
	objc.Object
}

func (c_ ComboBoxDataSourceWrapper) ImplementsComboBoxCompletedString() bool {
	return c_.RespondsToSelector(objc.GetSelector("comboBox:completedString:"))
}

func (c_ ComboBoxDataSourceWrapper) ComboBoxCompletedString(comboBox IComboBox, string_ string) string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("comboBox:completedString:"), objc.ExtractPtr(comboBox), string_)
	return rv
}

func (c_ ComboBoxDataSourceWrapper) ImplementsComboBoxIndexOfItemWithStringValue() bool {
	return c_.RespondsToSelector(objc.GetSelector("comboBox:indexOfItemWithStringValue:"))
}

func (c_ ComboBoxDataSourceWrapper) ComboBoxIndexOfItemWithStringValue(comboBox IComboBox, string_ string) uint {
	rv := objc.CallMethod[uint](c_, objc.GetSelector("comboBox:indexOfItemWithStringValue:"), objc.ExtractPtr(comboBox), string_)
	return rv
}

func (c_ ComboBoxDataSourceWrapper) ImplementsComboBoxObjectValueForItemAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("comboBox:objectValueForItemAtIndex:"))
}

func (c_ ComboBoxDataSourceWrapper) ComboBoxObjectValueForItemAtIndex(comboBox IComboBox, index int) objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.GetSelector("comboBox:objectValueForItemAtIndex:"), objc.ExtractPtr(comboBox), index)
	return rv
}

func (c_ ComboBoxDataSourceWrapper) ImplementsNumberOfItemsInComboBox() bool {
	return c_.RespondsToSelector(objc.GetSelector("numberOfItemsInComboBox:"))
}

func (c_ ComboBoxDataSourceWrapper) NumberOfItemsInComboBox(comboBox IComboBox) int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("numberOfItemsInComboBox:"), objc.ExtractPtr(comboBox))
	return rv
}
