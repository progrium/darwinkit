// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitytable?language=objc
type PAccessibilityTable interface {
	// optional
	AccessibilityColumnHeaderUIElements() []objc.IObject
	HasAccessibilityColumnHeaderUIElements() bool

	// optional
	SetAccessibilitySelectedRows(selectedRows []AccessibilityRowWrapper)
	HasSetAccessibilitySelectedRows() bool

	// optional
	AccessibilityRowHeaderUIElements() []objc.IObject
	HasAccessibilityRowHeaderUIElements() bool

	// optional
	AccessibilityVisibleRows() []PAccessibilityRow
	HasAccessibilityVisibleRows() bool

	// optional
	AccessibilitySelectedCells() []objc.IObject
	HasAccessibilitySelectedCells() bool

	// optional
	AccessibilityVisibleColumns() []objc.IObject
	HasAccessibilityVisibleColumns() bool

	// optional
	AccessibilityLabel() string
	HasAccessibilityLabel() bool

	// optional
	AccessibilityColumns() []objc.IObject
	HasAccessibilityColumns() bool

	// optional
	AccessibilityRows() []PAccessibilityRow
	HasAccessibilityRows() bool

	// optional
	AccessibilityVisibleCells() []objc.IObject
	HasAccessibilityVisibleCells() bool

	// optional
	AccessibilitySelectedColumns() []objc.IObject
	HasAccessibilitySelectedColumns() bool

	// optional
	AccessibilitySelectedRows() []PAccessibilityRow
	HasAccessibilitySelectedRows() bool
}

// A concrete type wrapper for the [PAccessibilityTable] protocol.
type AccessibilityTableWrapper struct {
	objc.Object
}

func (a_ AccessibilityTableWrapper) HasAccessibilityColumnHeaderUIElements() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityColumnHeaderUIElements"))
}

// Returns the column header accessibility elements for the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitytable/1526621-accessibilitycolumnheaderuieleme?language=objc
func (a_ AccessibilityTableWrapper) AccessibilityColumnHeaderUIElements() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityColumnHeaderUIElements"))
	return rv
}

func (a_ AccessibilityTableWrapper) HasSetAccessibilitySelectedRows() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilitySelectedRows:"))
}

// Sets the table’s currently selected rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitytable/1534612-setaccessibilityselectedrows?language=objc
func (a_ AccessibilityTableWrapper) SetAccessibilitySelectedRows(selectedRows []PAccessibilityRow) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilitySelectedRows:"), selectedRows)
}

func (a_ AccessibilityTableWrapper) HasAccessibilityRowHeaderUIElements() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityRowHeaderUIElements"))
}

// Returns the row header accessibility elements for the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitytable/1524262-accessibilityrowheaderuielements?language=objc
func (a_ AccessibilityTableWrapper) AccessibilityRowHeaderUIElements() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityRowHeaderUIElements"))
	return rv
}

func (a_ AccessibilityTableWrapper) HasAccessibilityVisibleRows() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityVisibleRows"))
}

// Returns the visible rows for the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitytable/1535128-accessibilityvisiblerows?language=objc
func (a_ AccessibilityTableWrapper) AccessibilityVisibleRows() []AccessibilityRowWrapper {
	rv := objc.Call[[]AccessibilityRowWrapper](a_, objc.Sel("accessibilityVisibleRows"))
	return rv
}

func (a_ AccessibilityTableWrapper) HasAccessibilitySelectedCells() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySelectedCells"))
}

// The currently selected cells for the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitytable/1525577-accessibilityselectedcells?language=objc
func (a_ AccessibilityTableWrapper) AccessibilitySelectedCells() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilitySelectedCells"))
	return rv
}

func (a_ AccessibilityTableWrapper) HasAccessibilityVisibleColumns() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityVisibleColumns"))
}

// Returns the visible columns for the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitytable/1530264-accessibilityvisiblecolumns?language=objc
func (a_ AccessibilityTableWrapper) AccessibilityVisibleColumns() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityVisibleColumns"))
	return rv
}

func (a_ AccessibilityTableWrapper) HasAccessibilityLabel() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityLabel"))
}

// Returns a short description of the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitytable/1526563-accessibilitylabel?language=objc
func (a_ AccessibilityTableWrapper) AccessibilityLabel() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityLabel"))
	return rv
}

func (a_ AccessibilityTableWrapper) HasAccessibilityColumns() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityColumns"))
}

// Returns the column accessibility elements for the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitytable/1524744-accessibilitycolumns?language=objc
func (a_ AccessibilityTableWrapper) AccessibilityColumns() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityColumns"))
	return rv
}

func (a_ AccessibilityTableWrapper) HasAccessibilityRows() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityRows"))
}

// Returns the row accessibility elements for the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitytable/1526672-accessibilityrows?language=objc
func (a_ AccessibilityTableWrapper) AccessibilityRows() []AccessibilityRowWrapper {
	rv := objc.Call[[]AccessibilityRowWrapper](a_, objc.Sel("accessibilityRows"))
	return rv
}

func (a_ AccessibilityTableWrapper) HasAccessibilityVisibleCells() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityVisibleCells"))
}

// Returns the visible cells for the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitytable/1526711-accessibilityvisiblecells?language=objc
func (a_ AccessibilityTableWrapper) AccessibilityVisibleCells() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityVisibleCells"))
	return rv
}

func (a_ AccessibilityTableWrapper) HasAccessibilitySelectedColumns() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySelectedColumns"))
}

// Returns the currently selected columns for the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitytable/1528430-accessibilityselectedcolumns?language=objc
func (a_ AccessibilityTableWrapper) AccessibilitySelectedColumns() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilitySelectedColumns"))
	return rv
}

func (a_ AccessibilityTableWrapper) HasAccessibilitySelectedRows() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySelectedRows"))
}

// Returns the currently selected rows for the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitytable/1529241-accessibilityselectedrows?language=objc
func (a_ AccessibilityTableWrapper) AccessibilitySelectedRows() []AccessibilityRowWrapper {
	rv := objc.Call[[]AccessibilityRowWrapper](a_, objc.Sel("accessibilitySelectedRows"))
	return rv
}
