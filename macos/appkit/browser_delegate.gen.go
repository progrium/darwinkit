// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IBrowserDelegate interface {
	ImplementsBrowserIsColumnValid() bool
	// optional
	BrowserIsColumnValid(sender Browser, column int) bool
	ImplementsBrowserNumberOfRowsInColumn() bool
	// optional
	BrowserNumberOfRowsInColumn(sender Browser, column int) int
	ImplementsBrowserNumberOfChildrenOfItem() bool
	// optional
	BrowserNumberOfChildrenOfItem(browser Browser, item objc.Object) int
	ImplementsBrowserTitleOfColumn() bool
	// optional
	BrowserTitleOfColumn(sender Browser, column int) string
	ImplementsBrowserShouldTypeSelectForEventWithCurrentSearchString() bool
	// optional
	BrowserShouldTypeSelectForEventWithCurrentSearchString(browser Browser, event Event, searchString string) bool
	ImplementsBrowserTypeSelectStringForRowInColumn() bool
	// optional
	BrowserTypeSelectStringForRowInColumn(browser Browser, row int, column int) string
	ImplementsBrowserNextTypeSelectMatchFromRowToRowInColumnForString() bool
	// optional
	BrowserNextTypeSelectMatchFromRowToRowInColumnForString(browser Browser, startRow int, endRow int, column int, searchString string) int
	ImplementsBrowserSelectCellWithStringInColumn() bool
	// optional
	BrowserSelectCellWithStringInColumn(sender Browser, title string, column int) bool
	ImplementsBrowserSelectRowInColumn() bool
	// optional
	BrowserSelectRowInColumn(sender Browser, row int, column int) bool
	ImplementsBrowserSelectionIndexesForProposedSelectionInColumn() bool
	// optional
	BrowserSelectionIndexesForProposedSelectionInColumn(browser Browser, proposedSelectionIndexes foundation.IndexSet, column int) foundation.IIndexSet
	ImplementsBrowserChildOfItem() bool
	// optional
	BrowserChildOfItem(browser Browser, index int, item objc.Object) objc.IObject
	ImplementsBrowserIsLeafItem() bool
	// optional
	BrowserIsLeafItem(browser Browser, item objc.Object) bool
	ImplementsBrowserShouldEditItem() bool
	// optional
	BrowserShouldEditItem(browser Browser, item objc.Object) bool
	ImplementsBrowserObjectValueForItem() bool
	// optional
	BrowserObjectValueForItem(browser Browser, item objc.Object) objc.IObject
	ImplementsBrowserSetObjectValueForItem() bool
	// optional
	BrowserSetObjectValueForItem(browser Browser, object objc.Object, item objc.Object)
	ImplementsRootItemForBrowser() bool
	// optional
	RootItemForBrowser(browser Browser) objc.IObject
	ImplementsBrowserPreviewViewControllerForLeafItem() bool
	// optional
	BrowserPreviewViewControllerForLeafItem(browser Browser, item objc.Object) IViewController
	ImplementsBrowserHeaderViewControllerForItem() bool
	// optional
	BrowserHeaderViewControllerForItem(browser Browser, item objc.Object) IViewController
	ImplementsBrowserCreateRowsForColumnInMatrix() bool
	// optional
	BrowserCreateRowsForColumnInMatrix(sender Browser, column int, matrix Matrix)
	ImplementsBrowserWillDisplayCellAtRowColumn() bool
	// optional
	BrowserWillDisplayCellAtRowColumn(sender Browser, cell objc.Object, row int, column int)
	ImplementsBrowserDidChangeLastColumnToColumn() bool
	// optional
	BrowserDidChangeLastColumnToColumn(browser Browser, oldLastColumn int, column int)
	ImplementsBrowserWillScroll() bool
	// optional
	BrowserWillScroll(sender Browser)
	ImplementsBrowserDidScroll() bool
	// optional
	BrowserDidScroll(sender Browser)
	ImplementsBrowserCanDragRowsWithIndexesInColumnWithEvent() bool
	// optional
	BrowserCanDragRowsWithIndexesInColumnWithEvent(browser Browser, rowIndexes foundation.IndexSet, column int, event Event) bool
	ImplementsBrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset() bool
	// optional
	BrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset(browser Browser, rowIndexes foundation.IndexSet, column int, event Event, dragImageOffset *foundation.Point) IImage
	ImplementsBrowserValidateDropProposedRowColumnDropOperation() bool
	// optional
	BrowserValidateDropProposedRowColumnDropOperation(browser Browser, info DraggingInfoWrapper, row *int, column *int, dropOperation *BrowserDropOperation) DragOperation
	ImplementsBrowserAcceptDropAtRowColumnDropOperation() bool
	// optional
	BrowserAcceptDropAtRowColumnDropOperation(browser Browser, info DraggingInfoWrapper, row int, column int, dropOperation BrowserDropOperation) bool
	ImplementsBrowserWriteRowsWithIndexesInColumnToPasteboard() bool
	// optional
	BrowserWriteRowsWithIndexesInColumnToPasteboard(browser Browser, rowIndexes foundation.IndexSet, column int, pasteboard Pasteboard) bool
	ImplementsBrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn() bool
	// optional
	// deprecated
	BrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn(browser Browser, dropDestination foundation.URL, rowIndexes foundation.IndexSet, column int) []string
	ImplementsBrowserShouldSizeColumnForUserResizeToWidth() bool
	// optional
	BrowserShouldSizeColumnForUserResizeToWidth(browser Browser, columnIndex int, forUserResize bool, suggestedWidth float64) float64
	ImplementsBrowserSizeToFitWidthOfColumn() bool
	// optional
	BrowserSizeToFitWidthOfColumn(browser Browser, columnIndex int) float64
	ImplementsBrowserColumnConfigurationDidChange() bool
	// optional
	BrowserColumnConfigurationDidChange(notification foundation.Notification)
	ImplementsBrowserHeightOfRowInColumn() bool
	// optional
	BrowserHeightOfRowInColumn(browser Browser, row int, columnIndex int) float64
	ImplementsBrowserShouldShowCellExpansionForRowColumn() bool
	// optional
	BrowserShouldShowCellExpansionForRowColumn(browser Browser, row int, column int) bool
}

type BrowserDelegate struct {
	_BrowserIsColumnValid                                                             func(sender Browser, column int) bool
	_BrowserNumberOfRowsInColumn                                                      func(sender Browser, column int) int
	_BrowserNumberOfChildrenOfItem                                                    func(browser Browser, item objc.Object) int
	_BrowserTitleOfColumn                                                             func(sender Browser, column int) string
	_BrowserShouldTypeSelectForEventWithCurrentSearchString                           func(browser Browser, event Event, searchString string) bool
	_BrowserTypeSelectStringForRowInColumn                                            func(browser Browser, row int, column int) string
	_BrowserNextTypeSelectMatchFromRowToRowInColumnForString                          func(browser Browser, startRow int, endRow int, column int, searchString string) int
	_BrowserSelectCellWithStringInColumn                                              func(sender Browser, title string, column int) bool
	_BrowserSelectRowInColumn                                                         func(sender Browser, row int, column int) bool
	_BrowserSelectionIndexesForProposedSelectionInColumn                              func(browser Browser, proposedSelectionIndexes foundation.IndexSet, column int) foundation.IIndexSet
	_BrowserChildOfItem                                                               func(browser Browser, index int, item objc.Object) objc.IObject
	_BrowserIsLeafItem                                                                func(browser Browser, item objc.Object) bool
	_BrowserShouldEditItem                                                            func(browser Browser, item objc.Object) bool
	_BrowserObjectValueForItem                                                        func(browser Browser, item objc.Object) objc.IObject
	_BrowserSetObjectValueForItem                                                     func(browser Browser, object objc.Object, item objc.Object)
	_RootItemForBrowser                                                               func(browser Browser) objc.IObject
	_BrowserPreviewViewControllerForLeafItem                                          func(browser Browser, item objc.Object) IViewController
	_BrowserHeaderViewControllerForItem                                               func(browser Browser, item objc.Object) IViewController
	_BrowserCreateRowsForColumnInMatrix                                               func(sender Browser, column int, matrix Matrix)
	_BrowserWillDisplayCellAtRowColumn                                                func(sender Browser, cell objc.Object, row int, column int)
	_BrowserDidChangeLastColumnToColumn                                               func(browser Browser, oldLastColumn int, column int)
	_BrowserWillScroll                                                                func(sender Browser)
	_BrowserDidScroll                                                                 func(sender Browser)
	_BrowserCanDragRowsWithIndexesInColumnWithEvent                                   func(browser Browser, rowIndexes foundation.IndexSet, column int, event Event) bool
	_BrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset                    func(browser Browser, rowIndexes foundation.IndexSet, column int, event Event, dragImageOffset *foundation.Point) IImage
	_BrowserValidateDropProposedRowColumnDropOperation                                func(browser Browser, info DraggingInfoWrapper, row *int, column *int, dropOperation *BrowserDropOperation) DragOperation
	_BrowserAcceptDropAtRowColumnDropOperation                                        func(browser Browser, info DraggingInfoWrapper, row int, column int, dropOperation BrowserDropOperation) bool
	_BrowserWriteRowsWithIndexesInColumnToPasteboard                                  func(browser Browser, rowIndexes foundation.IndexSet, column int, pasteboard Pasteboard) bool
	_BrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn func(browser Browser, dropDestination foundation.URL, rowIndexes foundation.IndexSet, column int) []string
	_BrowserShouldSizeColumnForUserResizeToWidth                                      func(browser Browser, columnIndex int, forUserResize bool, suggestedWidth float64) float64
	_BrowserSizeToFitWidthOfColumn                                                    func(browser Browser, columnIndex int) float64
	_BrowserColumnConfigurationDidChange                                              func(notification foundation.Notification)
	_BrowserHeightOfRowInColumn                                                       func(browser Browser, row int, columnIndex int) float64
	_BrowserShouldShowCellExpansionForRowColumn                                       func(browser Browser, row int, column int) bool
}

func (di *BrowserDelegate) ImplementsBrowserIsColumnValid() bool {
	return di._BrowserIsColumnValid != nil
}

func (di *BrowserDelegate) SetBrowserIsColumnValid(f func(sender Browser, column int) bool) {
	di._BrowserIsColumnValid = f
}

func (di *BrowserDelegate) BrowserIsColumnValid(sender Browser, column int) bool {
	return di._BrowserIsColumnValid(sender, column)
}
func (di *BrowserDelegate) ImplementsBrowserNumberOfRowsInColumn() bool {
	return di._BrowserNumberOfRowsInColumn != nil
}

func (di *BrowserDelegate) SetBrowserNumberOfRowsInColumn(f func(sender Browser, column int) int) {
	di._BrowserNumberOfRowsInColumn = f
}

func (di *BrowserDelegate) BrowserNumberOfRowsInColumn(sender Browser, column int) int {
	return di._BrowserNumberOfRowsInColumn(sender, column)
}
func (di *BrowserDelegate) ImplementsBrowserNumberOfChildrenOfItem() bool {
	return di._BrowserNumberOfChildrenOfItem != nil
}

func (di *BrowserDelegate) SetBrowserNumberOfChildrenOfItem(f func(browser Browser, item objc.Object) int) {
	di._BrowserNumberOfChildrenOfItem = f
}

func (di *BrowserDelegate) BrowserNumberOfChildrenOfItem(browser Browser, item objc.Object) int {
	return di._BrowserNumberOfChildrenOfItem(browser, item)
}
func (di *BrowserDelegate) ImplementsBrowserTitleOfColumn() bool {
	return di._BrowserTitleOfColumn != nil
}

func (di *BrowserDelegate) SetBrowserTitleOfColumn(f func(sender Browser, column int) string) {
	di._BrowserTitleOfColumn = f
}

func (di *BrowserDelegate) BrowserTitleOfColumn(sender Browser, column int) string {
	return di._BrowserTitleOfColumn(sender, column)
}
func (di *BrowserDelegate) ImplementsBrowserShouldTypeSelectForEventWithCurrentSearchString() bool {
	return di._BrowserShouldTypeSelectForEventWithCurrentSearchString != nil
}

func (di *BrowserDelegate) SetBrowserShouldTypeSelectForEventWithCurrentSearchString(f func(browser Browser, event Event, searchString string) bool) {
	di._BrowserShouldTypeSelectForEventWithCurrentSearchString = f
}

func (di *BrowserDelegate) BrowserShouldTypeSelectForEventWithCurrentSearchString(browser Browser, event Event, searchString string) bool {
	return di._BrowserShouldTypeSelectForEventWithCurrentSearchString(browser, event, searchString)
}
func (di *BrowserDelegate) ImplementsBrowserTypeSelectStringForRowInColumn() bool {
	return di._BrowserTypeSelectStringForRowInColumn != nil
}

func (di *BrowserDelegate) SetBrowserTypeSelectStringForRowInColumn(f func(browser Browser, row int, column int) string) {
	di._BrowserTypeSelectStringForRowInColumn = f
}

func (di *BrowserDelegate) BrowserTypeSelectStringForRowInColumn(browser Browser, row int, column int) string {
	return di._BrowserTypeSelectStringForRowInColumn(browser, row, column)
}
func (di *BrowserDelegate) ImplementsBrowserNextTypeSelectMatchFromRowToRowInColumnForString() bool {
	return di._BrowserNextTypeSelectMatchFromRowToRowInColumnForString != nil
}

func (di *BrowserDelegate) SetBrowserNextTypeSelectMatchFromRowToRowInColumnForString(f func(browser Browser, startRow int, endRow int, column int, searchString string) int) {
	di._BrowserNextTypeSelectMatchFromRowToRowInColumnForString = f
}

func (di *BrowserDelegate) BrowserNextTypeSelectMatchFromRowToRowInColumnForString(browser Browser, startRow int, endRow int, column int, searchString string) int {
	return di._BrowserNextTypeSelectMatchFromRowToRowInColumnForString(browser, startRow, endRow, column, searchString)
}
func (di *BrowserDelegate) ImplementsBrowserSelectCellWithStringInColumn() bool {
	return di._BrowserSelectCellWithStringInColumn != nil
}

func (di *BrowserDelegate) SetBrowserSelectCellWithStringInColumn(f func(sender Browser, title string, column int) bool) {
	di._BrowserSelectCellWithStringInColumn = f
}

func (di *BrowserDelegate) BrowserSelectCellWithStringInColumn(sender Browser, title string, column int) bool {
	return di._BrowserSelectCellWithStringInColumn(sender, title, column)
}
func (di *BrowserDelegate) ImplementsBrowserSelectRowInColumn() bool {
	return di._BrowserSelectRowInColumn != nil
}

func (di *BrowserDelegate) SetBrowserSelectRowInColumn(f func(sender Browser, row int, column int) bool) {
	di._BrowserSelectRowInColumn = f
}

func (di *BrowserDelegate) BrowserSelectRowInColumn(sender Browser, row int, column int) bool {
	return di._BrowserSelectRowInColumn(sender, row, column)
}
func (di *BrowserDelegate) ImplementsBrowserSelectionIndexesForProposedSelectionInColumn() bool {
	return di._BrowserSelectionIndexesForProposedSelectionInColumn != nil
}

func (di *BrowserDelegate) SetBrowserSelectionIndexesForProposedSelectionInColumn(f func(browser Browser, proposedSelectionIndexes foundation.IndexSet, column int) foundation.IIndexSet) {
	di._BrowserSelectionIndexesForProposedSelectionInColumn = f
}

func (di *BrowserDelegate) BrowserSelectionIndexesForProposedSelectionInColumn(browser Browser, proposedSelectionIndexes foundation.IndexSet, column int) foundation.IIndexSet {
	return di._BrowserSelectionIndexesForProposedSelectionInColumn(browser, proposedSelectionIndexes, column)
}
func (di *BrowserDelegate) ImplementsBrowserChildOfItem() bool {
	return di._BrowserChildOfItem != nil
}

func (di *BrowserDelegate) SetBrowserChildOfItem(f func(browser Browser, index int, item objc.Object) objc.IObject) {
	di._BrowserChildOfItem = f
}

func (di *BrowserDelegate) BrowserChildOfItem(browser Browser, index int, item objc.Object) objc.IObject {
	return di._BrowserChildOfItem(browser, index, item)
}
func (di *BrowserDelegate) ImplementsBrowserIsLeafItem() bool {
	return di._BrowserIsLeafItem != nil
}

func (di *BrowserDelegate) SetBrowserIsLeafItem(f func(browser Browser, item objc.Object) bool) {
	di._BrowserIsLeafItem = f
}

func (di *BrowserDelegate) BrowserIsLeafItem(browser Browser, item objc.Object) bool {
	return di._BrowserIsLeafItem(browser, item)
}
func (di *BrowserDelegate) ImplementsBrowserShouldEditItem() bool {
	return di._BrowserShouldEditItem != nil
}

func (di *BrowserDelegate) SetBrowserShouldEditItem(f func(browser Browser, item objc.Object) bool) {
	di._BrowserShouldEditItem = f
}

func (di *BrowserDelegate) BrowserShouldEditItem(browser Browser, item objc.Object) bool {
	return di._BrowserShouldEditItem(browser, item)
}
func (di *BrowserDelegate) ImplementsBrowserObjectValueForItem() bool {
	return di._BrowserObjectValueForItem != nil
}

func (di *BrowserDelegate) SetBrowserObjectValueForItem(f func(browser Browser, item objc.Object) objc.IObject) {
	di._BrowserObjectValueForItem = f
}

func (di *BrowserDelegate) BrowserObjectValueForItem(browser Browser, item objc.Object) objc.IObject {
	return di._BrowserObjectValueForItem(browser, item)
}
func (di *BrowserDelegate) ImplementsBrowserSetObjectValueForItem() bool {
	return di._BrowserSetObjectValueForItem != nil
}

func (di *BrowserDelegate) SetBrowserSetObjectValueForItem(f func(browser Browser, object objc.Object, item objc.Object)) {
	di._BrowserSetObjectValueForItem = f
}

func (di *BrowserDelegate) BrowserSetObjectValueForItem(browser Browser, object objc.Object, item objc.Object) {
	di._BrowserSetObjectValueForItem(browser, object, item)
}
func (di *BrowserDelegate) ImplementsRootItemForBrowser() bool {
	return di._RootItemForBrowser != nil
}

func (di *BrowserDelegate) SetRootItemForBrowser(f func(browser Browser) objc.IObject) {
	di._RootItemForBrowser = f
}

func (di *BrowserDelegate) RootItemForBrowser(browser Browser) objc.IObject {
	return di._RootItemForBrowser(browser)
}
func (di *BrowserDelegate) ImplementsBrowserPreviewViewControllerForLeafItem() bool {
	return di._BrowserPreviewViewControllerForLeafItem != nil
}

func (di *BrowserDelegate) SetBrowserPreviewViewControllerForLeafItem(f func(browser Browser, item objc.Object) IViewController) {
	di._BrowserPreviewViewControllerForLeafItem = f
}

func (di *BrowserDelegate) BrowserPreviewViewControllerForLeafItem(browser Browser, item objc.Object) IViewController {
	return di._BrowserPreviewViewControllerForLeafItem(browser, item)
}
func (di *BrowserDelegate) ImplementsBrowserHeaderViewControllerForItem() bool {
	return di._BrowserHeaderViewControllerForItem != nil
}

func (di *BrowserDelegate) SetBrowserHeaderViewControllerForItem(f func(browser Browser, item objc.Object) IViewController) {
	di._BrowserHeaderViewControllerForItem = f
}

func (di *BrowserDelegate) BrowserHeaderViewControllerForItem(browser Browser, item objc.Object) IViewController {
	return di._BrowserHeaderViewControllerForItem(browser, item)
}
func (di *BrowserDelegate) ImplementsBrowserCreateRowsForColumnInMatrix() bool {
	return di._BrowserCreateRowsForColumnInMatrix != nil
}

func (di *BrowserDelegate) SetBrowserCreateRowsForColumnInMatrix(f func(sender Browser, column int, matrix Matrix)) {
	di._BrowserCreateRowsForColumnInMatrix = f
}

func (di *BrowserDelegate) BrowserCreateRowsForColumnInMatrix(sender Browser, column int, matrix Matrix) {
	di._BrowserCreateRowsForColumnInMatrix(sender, column, matrix)
}
func (di *BrowserDelegate) ImplementsBrowserWillDisplayCellAtRowColumn() bool {
	return di._BrowserWillDisplayCellAtRowColumn != nil
}

func (di *BrowserDelegate) SetBrowserWillDisplayCellAtRowColumn(f func(sender Browser, cell objc.Object, row int, column int)) {
	di._BrowserWillDisplayCellAtRowColumn = f
}

func (di *BrowserDelegate) BrowserWillDisplayCellAtRowColumn(sender Browser, cell objc.Object, row int, column int) {
	di._BrowserWillDisplayCellAtRowColumn(sender, cell, row, column)
}
func (di *BrowserDelegate) ImplementsBrowserDidChangeLastColumnToColumn() bool {
	return di._BrowserDidChangeLastColumnToColumn != nil
}

func (di *BrowserDelegate) SetBrowserDidChangeLastColumnToColumn(f func(browser Browser, oldLastColumn int, column int)) {
	di._BrowserDidChangeLastColumnToColumn = f
}

func (di *BrowserDelegate) BrowserDidChangeLastColumnToColumn(browser Browser, oldLastColumn int, column int) {
	di._BrowserDidChangeLastColumnToColumn(browser, oldLastColumn, column)
}
func (di *BrowserDelegate) ImplementsBrowserWillScroll() bool {
	return di._BrowserWillScroll != nil
}

func (di *BrowserDelegate) SetBrowserWillScroll(f func(sender Browser)) {
	di._BrowserWillScroll = f
}

func (di *BrowserDelegate) BrowserWillScroll(sender Browser) {
	di._BrowserWillScroll(sender)
}
func (di *BrowserDelegate) ImplementsBrowserDidScroll() bool {
	return di._BrowserDidScroll != nil
}

func (di *BrowserDelegate) SetBrowserDidScroll(f func(sender Browser)) {
	di._BrowserDidScroll = f
}

func (di *BrowserDelegate) BrowserDidScroll(sender Browser) {
	di._BrowserDidScroll(sender)
}
func (di *BrowserDelegate) ImplementsBrowserCanDragRowsWithIndexesInColumnWithEvent() bool {
	return di._BrowserCanDragRowsWithIndexesInColumnWithEvent != nil
}

func (di *BrowserDelegate) SetBrowserCanDragRowsWithIndexesInColumnWithEvent(f func(browser Browser, rowIndexes foundation.IndexSet, column int, event Event) bool) {
	di._BrowserCanDragRowsWithIndexesInColumnWithEvent = f
}

func (di *BrowserDelegate) BrowserCanDragRowsWithIndexesInColumnWithEvent(browser Browser, rowIndexes foundation.IndexSet, column int, event Event) bool {
	return di._BrowserCanDragRowsWithIndexesInColumnWithEvent(browser, rowIndexes, column, event)
}
func (di *BrowserDelegate) ImplementsBrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset() bool {
	return di._BrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset != nil
}

func (di *BrowserDelegate) SetBrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset(f func(browser Browser, rowIndexes foundation.IndexSet, column int, event Event, dragImageOffset *foundation.Point) IImage) {
	di._BrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset = f
}

func (di *BrowserDelegate) BrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset(browser Browser, rowIndexes foundation.IndexSet, column int, event Event, dragImageOffset *foundation.Point) IImage {
	return di._BrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset(browser, rowIndexes, column, event, dragImageOffset)
}
func (di *BrowserDelegate) ImplementsBrowserValidateDropProposedRowColumnDropOperation() bool {
	return di._BrowserValidateDropProposedRowColumnDropOperation != nil
}

func (di *BrowserDelegate) SetBrowserValidateDropProposedRowColumnDropOperation(f func(browser Browser, info DraggingInfoWrapper, row *int, column *int, dropOperation *BrowserDropOperation) DragOperation) {
	di._BrowserValidateDropProposedRowColumnDropOperation = f
}

func (di *BrowserDelegate) BrowserValidateDropProposedRowColumnDropOperation(browser Browser, info DraggingInfoWrapper, row *int, column *int, dropOperation *BrowserDropOperation) DragOperation {
	return di._BrowserValidateDropProposedRowColumnDropOperation(browser, info, row, column, dropOperation)
}
func (di *BrowserDelegate) ImplementsBrowserAcceptDropAtRowColumnDropOperation() bool {
	return di._BrowserAcceptDropAtRowColumnDropOperation != nil
}

func (di *BrowserDelegate) SetBrowserAcceptDropAtRowColumnDropOperation(f func(browser Browser, info DraggingInfoWrapper, row int, column int, dropOperation BrowserDropOperation) bool) {
	di._BrowserAcceptDropAtRowColumnDropOperation = f
}

func (di *BrowserDelegate) BrowserAcceptDropAtRowColumnDropOperation(browser Browser, info DraggingInfoWrapper, row int, column int, dropOperation BrowserDropOperation) bool {
	return di._BrowserAcceptDropAtRowColumnDropOperation(browser, info, row, column, dropOperation)
}
func (di *BrowserDelegate) ImplementsBrowserWriteRowsWithIndexesInColumnToPasteboard() bool {
	return di._BrowserWriteRowsWithIndexesInColumnToPasteboard != nil
}

func (di *BrowserDelegate) SetBrowserWriteRowsWithIndexesInColumnToPasteboard(f func(browser Browser, rowIndexes foundation.IndexSet, column int, pasteboard Pasteboard) bool) {
	di._BrowserWriteRowsWithIndexesInColumnToPasteboard = f
}

func (di *BrowserDelegate) BrowserWriteRowsWithIndexesInColumnToPasteboard(browser Browser, rowIndexes foundation.IndexSet, column int, pasteboard Pasteboard) bool {
	return di._BrowserWriteRowsWithIndexesInColumnToPasteboard(browser, rowIndexes, column, pasteboard)
}
func (di *BrowserDelegate) ImplementsBrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn() bool {
	return di._BrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn != nil
}

// deprecated
func (di *BrowserDelegate) SetBrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn(f func(browser Browser, dropDestination foundation.URL, rowIndexes foundation.IndexSet, column int) []string) {
	di._BrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn = f
}

// deprecated
func (di *BrowserDelegate) BrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn(browser Browser, dropDestination foundation.URL, rowIndexes foundation.IndexSet, column int) []string {
	return di._BrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn(browser, dropDestination, rowIndexes, column)
}
func (di *BrowserDelegate) ImplementsBrowserShouldSizeColumnForUserResizeToWidth() bool {
	return di._BrowserShouldSizeColumnForUserResizeToWidth != nil
}

func (di *BrowserDelegate) SetBrowserShouldSizeColumnForUserResizeToWidth(f func(browser Browser, columnIndex int, forUserResize bool, suggestedWidth float64) float64) {
	di._BrowserShouldSizeColumnForUserResizeToWidth = f
}

func (di *BrowserDelegate) BrowserShouldSizeColumnForUserResizeToWidth(browser Browser, columnIndex int, forUserResize bool, suggestedWidth float64) float64 {
	return di._BrowserShouldSizeColumnForUserResizeToWidth(browser, columnIndex, forUserResize, suggestedWidth)
}
func (di *BrowserDelegate) ImplementsBrowserSizeToFitWidthOfColumn() bool {
	return di._BrowserSizeToFitWidthOfColumn != nil
}

func (di *BrowserDelegate) SetBrowserSizeToFitWidthOfColumn(f func(browser Browser, columnIndex int) float64) {
	di._BrowserSizeToFitWidthOfColumn = f
}

func (di *BrowserDelegate) BrowserSizeToFitWidthOfColumn(browser Browser, columnIndex int) float64 {
	return di._BrowserSizeToFitWidthOfColumn(browser, columnIndex)
}
func (di *BrowserDelegate) ImplementsBrowserColumnConfigurationDidChange() bool {
	return di._BrowserColumnConfigurationDidChange != nil
}

func (di *BrowserDelegate) SetBrowserColumnConfigurationDidChange(f func(notification foundation.Notification)) {
	di._BrowserColumnConfigurationDidChange = f
}

func (di *BrowserDelegate) BrowserColumnConfigurationDidChange(notification foundation.Notification) {
	di._BrowserColumnConfigurationDidChange(notification)
}
func (di *BrowserDelegate) ImplementsBrowserHeightOfRowInColumn() bool {
	return di._BrowserHeightOfRowInColumn != nil
}

func (di *BrowserDelegate) SetBrowserHeightOfRowInColumn(f func(browser Browser, row int, columnIndex int) float64) {
	di._BrowserHeightOfRowInColumn = f
}

func (di *BrowserDelegate) BrowserHeightOfRowInColumn(browser Browser, row int, columnIndex int) float64 {
	return di._BrowserHeightOfRowInColumn(browser, row, columnIndex)
}
func (di *BrowserDelegate) ImplementsBrowserShouldShowCellExpansionForRowColumn() bool {
	return di._BrowserShouldShowCellExpansionForRowColumn != nil
}

func (di *BrowserDelegate) SetBrowserShouldShowCellExpansionForRowColumn(f func(browser Browser, row int, column int) bool) {
	di._BrowserShouldShowCellExpansionForRowColumn = f
}

func (di *BrowserDelegate) BrowserShouldShowCellExpansionForRowColumn(browser Browser, row int, column int) bool {
	return di._BrowserShouldShowCellExpansionForRowColumn(browser, row, column)
}

type BrowserDelegateWrapper struct {
	objc.Object
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserIsColumnValid() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:isColumnValid:"))
}

func (b_ BrowserDelegateWrapper) BrowserIsColumnValid(sender IBrowser, column int) bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("browser:isColumnValid:"), objc.ExtractPtr(sender), column)
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserNumberOfRowsInColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:numberOfRowsInColumn:"))
}

func (b_ BrowserDelegateWrapper) BrowserNumberOfRowsInColumn(sender IBrowser, column int) int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("browser:numberOfRowsInColumn:"), objc.ExtractPtr(sender), column)
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserNumberOfChildrenOfItem() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:numberOfChildrenOfItem:"))
}

func (b_ BrowserDelegateWrapper) BrowserNumberOfChildrenOfItem(browser IBrowser, item objc.IObject) int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("browser:numberOfChildrenOfItem:"), objc.ExtractPtr(browser), objc.ExtractPtr(item))
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserTitleOfColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:titleOfColumn:"))
}

func (b_ BrowserDelegateWrapper) BrowserTitleOfColumn(sender IBrowser, column int) string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("browser:titleOfColumn:"), objc.ExtractPtr(sender), column)
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserShouldTypeSelectForEventWithCurrentSearchString() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:shouldTypeSelectForEvent:withCurrentSearchString:"))
}

func (b_ BrowserDelegateWrapper) BrowserShouldTypeSelectForEventWithCurrentSearchString(browser IBrowser, event IEvent, searchString string) bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("browser:shouldTypeSelectForEvent:withCurrentSearchString:"), objc.ExtractPtr(browser), objc.ExtractPtr(event), searchString)
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserTypeSelectStringForRowInColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:typeSelectStringForRow:inColumn:"))
}

func (b_ BrowserDelegateWrapper) BrowserTypeSelectStringForRowInColumn(browser IBrowser, row int, column int) string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("browser:typeSelectStringForRow:inColumn:"), objc.ExtractPtr(browser), row, column)
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserNextTypeSelectMatchFromRowToRowInColumnForString() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:nextTypeSelectMatchFromRow:toRow:inColumn:forString:"))
}

func (b_ BrowserDelegateWrapper) BrowserNextTypeSelectMatchFromRowToRowInColumnForString(browser IBrowser, startRow int, endRow int, column int, searchString string) int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("browser:nextTypeSelectMatchFromRow:toRow:inColumn:forString:"), objc.ExtractPtr(browser), startRow, endRow, column, searchString)
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserSelectCellWithStringInColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:selectCellWithString:inColumn:"))
}

func (b_ BrowserDelegateWrapper) BrowserSelectCellWithStringInColumn(sender IBrowser, title string, column int) bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("browser:selectCellWithString:inColumn:"), objc.ExtractPtr(sender), title, column)
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserSelectRowInColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:selectRow:inColumn:"))
}

func (b_ BrowserDelegateWrapper) BrowserSelectRowInColumn(sender IBrowser, row int, column int) bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("browser:selectRow:inColumn:"), objc.ExtractPtr(sender), row, column)
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserSelectionIndexesForProposedSelectionInColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:selectionIndexesForProposedSelection:inColumn:"))
}

func (b_ BrowserDelegateWrapper) BrowserSelectionIndexesForProposedSelectionInColumn(browser IBrowser, proposedSelectionIndexes foundation.IIndexSet, column int) foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](b_, objc.GetSelector("browser:selectionIndexesForProposedSelection:inColumn:"), objc.ExtractPtr(browser), objc.ExtractPtr(proposedSelectionIndexes), column)
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserChildOfItem() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:child:ofItem:"))
}

func (b_ BrowserDelegateWrapper) BrowserChildOfItem(browser IBrowser, index int, item objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("browser:child:ofItem:"), objc.ExtractPtr(browser), index, objc.ExtractPtr(item))
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserIsLeafItem() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:isLeafItem:"))
}

func (b_ BrowserDelegateWrapper) BrowserIsLeafItem(browser IBrowser, item objc.IObject) bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("browser:isLeafItem:"), objc.ExtractPtr(browser), objc.ExtractPtr(item))
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserShouldEditItem() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:shouldEditItem:"))
}

func (b_ BrowserDelegateWrapper) BrowserShouldEditItem(browser IBrowser, item objc.IObject) bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("browser:shouldEditItem:"), objc.ExtractPtr(browser), objc.ExtractPtr(item))
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserObjectValueForItem() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:objectValueForItem:"))
}

func (b_ BrowserDelegateWrapper) BrowserObjectValueForItem(browser IBrowser, item objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("browser:objectValueForItem:"), objc.ExtractPtr(browser), objc.ExtractPtr(item))
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserSetObjectValueForItem() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:setObjectValue:forItem:"))
}

func (b_ BrowserDelegateWrapper) BrowserSetObjectValueForItem(browser IBrowser, object objc.IObject, item objc.IObject) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("browser:setObjectValue:forItem:"), objc.ExtractPtr(browser), objc.ExtractPtr(object), objc.ExtractPtr(item))
}

func (b_ BrowserDelegateWrapper) ImplementsRootItemForBrowser() bool {
	return b_.RespondsToSelector(objc.GetSelector("rootItemForBrowser:"))
}

func (b_ BrowserDelegateWrapper) RootItemForBrowser(browser IBrowser) objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("rootItemForBrowser:"), objc.ExtractPtr(browser))
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserPreviewViewControllerForLeafItem() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:previewViewControllerForLeafItem:"))
}

func (b_ BrowserDelegateWrapper) BrowserPreviewViewControllerForLeafItem(browser IBrowser, item objc.IObject) ViewController {
	rv := objc.CallMethod[ViewController](b_, objc.GetSelector("browser:previewViewControllerForLeafItem:"), objc.ExtractPtr(browser), objc.ExtractPtr(item))
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserHeaderViewControllerForItem() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:headerViewControllerForItem:"))
}

func (b_ BrowserDelegateWrapper) BrowserHeaderViewControllerForItem(browser IBrowser, item objc.IObject) ViewController {
	rv := objc.CallMethod[ViewController](b_, objc.GetSelector("browser:headerViewControllerForItem:"), objc.ExtractPtr(browser), objc.ExtractPtr(item))
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserCreateRowsForColumnInMatrix() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:createRowsForColumn:inMatrix:"))
}

func (b_ BrowserDelegateWrapper) BrowserCreateRowsForColumnInMatrix(sender IBrowser, column int, matrix IMatrix) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("browser:createRowsForColumn:inMatrix:"), objc.ExtractPtr(sender), column, objc.ExtractPtr(matrix))
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserWillDisplayCellAtRowColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:willDisplayCell:atRow:column:"))
}

func (b_ BrowserDelegateWrapper) BrowserWillDisplayCellAtRowColumn(sender IBrowser, cell objc.IObject, row int, column int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("browser:willDisplayCell:atRow:column:"), objc.ExtractPtr(sender), objc.ExtractPtr(cell), row, column)
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserDidChangeLastColumnToColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:didChangeLastColumn:toColumn:"))
}

func (b_ BrowserDelegateWrapper) BrowserDidChangeLastColumnToColumn(browser IBrowser, oldLastColumn int, column int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("browser:didChangeLastColumn:toColumn:"), objc.ExtractPtr(browser), oldLastColumn, column)
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserWillScroll() bool {
	return b_.RespondsToSelector(objc.GetSelector("browserWillScroll:"))
}

func (b_ BrowserDelegateWrapper) BrowserWillScroll(sender IBrowser) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("browserWillScroll:"), objc.ExtractPtr(sender))
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserDidScroll() bool {
	return b_.RespondsToSelector(objc.GetSelector("browserDidScroll:"))
}

func (b_ BrowserDelegateWrapper) BrowserDidScroll(sender IBrowser) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("browserDidScroll:"), objc.ExtractPtr(sender))
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserCanDragRowsWithIndexesInColumnWithEvent() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:canDragRowsWithIndexes:inColumn:withEvent:"))
}

func (b_ BrowserDelegateWrapper) BrowserCanDragRowsWithIndexesInColumnWithEvent(browser IBrowser, rowIndexes foundation.IIndexSet, column int, event IEvent) bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("browser:canDragRowsWithIndexes:inColumn:withEvent:"), objc.ExtractPtr(browser), objc.ExtractPtr(rowIndexes), column, objc.ExtractPtr(event))
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:draggingImageForRowsWithIndexes:inColumn:withEvent:offset:"))
}

func (b_ BrowserDelegateWrapper) BrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset(browser IBrowser, rowIndexes foundation.IIndexSet, column int, event IEvent, dragImageOffset *foundation.Point) Image {
	rv := objc.CallMethod[Image](b_, objc.GetSelector("browser:draggingImageForRowsWithIndexes:inColumn:withEvent:offset:"), objc.ExtractPtr(browser), objc.ExtractPtr(rowIndexes), column, objc.ExtractPtr(event), dragImageOffset)
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserValidateDropProposedRowColumnDropOperation() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:validateDrop:proposedRow:column:dropOperation:"))
}

func (b_ BrowserDelegateWrapper) BrowserValidateDropProposedRowColumnDropOperation(browser IBrowser, info IDraggingInfo, row *int, column *int, dropOperation *BrowserDropOperation) DragOperation {
	po := objc.WrapAsProtocol("NSDraggingInfo", info)
	rv := objc.CallMethod[DragOperation](b_, objc.GetSelector("browser:validateDrop:proposedRow:column:dropOperation:"), objc.ExtractPtr(browser), po, row, column, dropOperation)
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserAcceptDropAtRowColumnDropOperation() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:acceptDrop:atRow:column:dropOperation:"))
}

func (b_ BrowserDelegateWrapper) BrowserAcceptDropAtRowColumnDropOperation(browser IBrowser, info IDraggingInfo, row int, column int, dropOperation BrowserDropOperation) bool {
	po := objc.WrapAsProtocol("NSDraggingInfo", info)
	rv := objc.CallMethod[bool](b_, objc.GetSelector("browser:acceptDrop:atRow:column:dropOperation:"), objc.ExtractPtr(browser), po, row, column, dropOperation)
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserWriteRowsWithIndexesInColumnToPasteboard() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:writeRowsWithIndexes:inColumn:toPasteboard:"))
}

func (b_ BrowserDelegateWrapper) BrowserWriteRowsWithIndexesInColumnToPasteboard(browser IBrowser, rowIndexes foundation.IIndexSet, column int, pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("browser:writeRowsWithIndexes:inColumn:toPasteboard:"), objc.ExtractPtr(browser), objc.ExtractPtr(rowIndexes), column, objc.ExtractPtr(pasteboard))
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:namesOfPromisedFilesDroppedAtDestination:forDraggedRowsWithIndexes:inColumn:"))
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserShouldSizeColumnForUserResizeToWidth() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:shouldSizeColumn:forUserResize:toWidth:"))
}

func (b_ BrowserDelegateWrapper) BrowserShouldSizeColumnForUserResizeToWidth(browser IBrowser, columnIndex int, forUserResize bool, suggestedWidth float64) float64 {
	rv := objc.CallMethod[float64](b_, objc.GetSelector("browser:shouldSizeColumn:forUserResize:toWidth:"), objc.ExtractPtr(browser), columnIndex, forUserResize, suggestedWidth)
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserSizeToFitWidthOfColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:sizeToFitWidthOfColumn:"))
}

func (b_ BrowserDelegateWrapper) BrowserSizeToFitWidthOfColumn(browser IBrowser, columnIndex int) float64 {
	rv := objc.CallMethod[float64](b_, objc.GetSelector("browser:sizeToFitWidthOfColumn:"), objc.ExtractPtr(browser), columnIndex)
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserColumnConfigurationDidChange() bool {
	return b_.RespondsToSelector(objc.GetSelector("browserColumnConfigurationDidChange:"))
}

func (b_ BrowserDelegateWrapper) BrowserColumnConfigurationDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("browserColumnConfigurationDidChange:"), objc.ExtractPtr(notification))
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserHeightOfRowInColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:heightOfRow:inColumn:"))
}

func (b_ BrowserDelegateWrapper) BrowserHeightOfRowInColumn(browser IBrowser, row int, columnIndex int) float64 {
	rv := objc.CallMethod[float64](b_, objc.GetSelector("browser:heightOfRow:inColumn:"), objc.ExtractPtr(browser), row, columnIndex)
	return rv
}

func (b_ BrowserDelegateWrapper) ImplementsBrowserShouldShowCellExpansionForRowColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:shouldShowCellExpansionForRow:column:"))
}

func (b_ BrowserDelegateWrapper) BrowserShouldShowCellExpansionForRowColumn(browser IBrowser, row int, column int) bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("browser:shouldShowCellExpansionForRow:column:"), objc.ExtractPtr(browser), row, column)
	return rv
}
