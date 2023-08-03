// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type ITableViewDelegate interface {
	IControlTextEditingDelegate
	ImplementsTableViewViewForTableColumnRow() bool
	// optional
	TableViewViewForTableColumnRow(tableView TableView, tableColumn TableColumn, row int) IView
	ImplementsTableViewRowViewForRow() bool
	// optional
	TableViewRowViewForRow(tableView TableView, row int) ITableRowView
	ImplementsTableViewDidAddRowViewForRow() bool
	// optional
	TableViewDidAddRowViewForRow(tableView TableView, rowView TableRowView, row int)
	ImplementsTableViewDidRemoveRowViewForRow() bool
	// optional
	TableViewDidRemoveRowViewForRow(tableView TableView, rowView TableRowView, row int)
	ImplementsTableViewIsGroupRow() bool
	// optional
	TableViewIsGroupRow(tableView TableView, row int) bool
	ImplementsTableViewWillDisplayCellForTableColumnRow() bool
	// optional
	TableViewWillDisplayCellForTableColumnRow(tableView TableView, cell objc.Object, tableColumn TableColumn, row int)
	ImplementsTableViewDataCellForTableColumnRow() bool
	// optional
	TableViewDataCellForTableColumnRow(tableView TableView, tableColumn TableColumn, row int) ICell
	ImplementsTableViewShouldShowCellExpansionForTableColumnRow() bool
	// optional
	TableViewShouldShowCellExpansionForTableColumnRow(tableView TableView, tableColumn TableColumn, row int) bool
	ImplementsTableViewToolTipForCellRectTableColumnRowMouseLocation() bool
	// optional
	TableViewToolTipForCellRectTableColumnRowMouseLocation(tableView TableView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, row int, mouseLocation foundation.Point) string
	ImplementsTableViewShouldEditTableColumnRow() bool
	// optional
	TableViewShouldEditTableColumnRow(tableView TableView, tableColumn TableColumn, row int) bool
	ImplementsTableViewHeightOfRow() bool
	// optional
	TableViewHeightOfRow(tableView TableView, row int) float64
	ImplementsTableViewSizeToFitWidthOfColumn() bool
	// optional
	TableViewSizeToFitWidthOfColumn(tableView TableView, column int) float64
	ImplementsSelectionShouldChangeInTableView() bool
	// optional
	SelectionShouldChangeInTableView(tableView TableView) bool
	ImplementsTableViewShouldSelectRow() bool
	// optional
	TableViewShouldSelectRow(tableView TableView, row int) bool
	ImplementsTableViewSelectionIndexesForProposedSelection() bool
	// optional
	TableViewSelectionIndexesForProposedSelection(tableView TableView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet
	ImplementsTableViewShouldSelectTableColumn() bool
	// optional
	TableViewShouldSelectTableColumn(tableView TableView, tableColumn TableColumn) bool
	ImplementsTableViewSelectionIsChanging() bool
	// optional
	TableViewSelectionIsChanging(notification foundation.Notification)
	ImplementsTableViewSelectionDidChange() bool
	// optional
	TableViewSelectionDidChange(notification foundation.Notification)
	ImplementsTableViewShouldTypeSelectForEventWithCurrentSearchString() bool
	// optional
	TableViewShouldTypeSelectForEventWithCurrentSearchString(tableView TableView, event Event, searchString string) bool
	ImplementsTableViewTypeSelectStringForTableColumnRow() bool
	// optional
	TableViewTypeSelectStringForTableColumnRow(tableView TableView, tableColumn TableColumn, row int) string
	ImplementsTableViewNextTypeSelectMatchFromRowToRowForString() bool
	// optional
	TableViewNextTypeSelectMatchFromRowToRowForString(tableView TableView, startRow int, endRow int, searchString string) int
	ImplementsTableViewShouldReorderColumnToColumn() bool
	// optional
	TableViewShouldReorderColumnToColumn(tableView TableView, columnIndex int, newColumnIndex int) bool
	ImplementsTableViewDidDragTableColumn() bool
	// optional
	TableViewDidDragTableColumn(tableView TableView, tableColumn TableColumn)
	ImplementsTableViewColumnDidMove() bool
	// optional
	TableViewColumnDidMove(notification foundation.Notification)
	ImplementsTableViewColumnDidResize() bool
	// optional
	TableViewColumnDidResize(notification foundation.Notification)
	ImplementsTableViewDidClickTableColumn() bool
	// optional
	TableViewDidClickTableColumn(tableView TableView, tableColumn TableColumn)
	ImplementsTableViewMouseDownInHeaderOfTableColumn() bool
	// optional
	TableViewMouseDownInHeaderOfTableColumn(tableView TableView, tableColumn TableColumn)
	ImplementsTableViewShouldTrackCellForTableColumnRow() bool
	// optional
	TableViewShouldTrackCellForTableColumnRow(tableView TableView, cell Cell, tableColumn TableColumn, row int) bool
	ImplementsTableViewRowActionsForRowEdge() bool
	// optional
	TableViewRowActionsForRowEdge(tableView TableView, row int, edge TableRowActionEdge) []ITableViewRowAction
}

type TableViewDelegate struct {
	ControlTextEditingDelegate
	_TableViewViewForTableColumnRow                           func(tableView TableView, tableColumn TableColumn, row int) IView
	_TableViewRowViewForRow                                   func(tableView TableView, row int) ITableRowView
	_TableViewDidAddRowViewForRow                             func(tableView TableView, rowView TableRowView, row int)
	_TableViewDidRemoveRowViewForRow                          func(tableView TableView, rowView TableRowView, row int)
	_TableViewIsGroupRow                                      func(tableView TableView, row int) bool
	_TableViewWillDisplayCellForTableColumnRow                func(tableView TableView, cell objc.Object, tableColumn TableColumn, row int)
	_TableViewDataCellForTableColumnRow                       func(tableView TableView, tableColumn TableColumn, row int) ICell
	_TableViewShouldShowCellExpansionForTableColumnRow        func(tableView TableView, tableColumn TableColumn, row int) bool
	_TableViewToolTipForCellRectTableColumnRowMouseLocation   func(tableView TableView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, row int, mouseLocation foundation.Point) string
	_TableViewShouldEditTableColumnRow                        func(tableView TableView, tableColumn TableColumn, row int) bool
	_TableViewHeightOfRow                                     func(tableView TableView, row int) float64
	_TableViewSizeToFitWidthOfColumn                          func(tableView TableView, column int) float64
	_SelectionShouldChangeInTableView                         func(tableView TableView) bool
	_TableViewShouldSelectRow                                 func(tableView TableView, row int) bool
	_TableViewSelectionIndexesForProposedSelection            func(tableView TableView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet
	_TableViewShouldSelectTableColumn                         func(tableView TableView, tableColumn TableColumn) bool
	_TableViewSelectionIsChanging                             func(notification foundation.Notification)
	_TableViewSelectionDidChange                              func(notification foundation.Notification)
	_TableViewShouldTypeSelectForEventWithCurrentSearchString func(tableView TableView, event Event, searchString string) bool
	_TableViewTypeSelectStringForTableColumnRow               func(tableView TableView, tableColumn TableColumn, row int) string
	_TableViewNextTypeSelectMatchFromRowToRowForString        func(tableView TableView, startRow int, endRow int, searchString string) int
	_TableViewShouldReorderColumnToColumn                     func(tableView TableView, columnIndex int, newColumnIndex int) bool
	_TableViewDidDragTableColumn                              func(tableView TableView, tableColumn TableColumn)
	_TableViewColumnDidMove                                   func(notification foundation.Notification)
	_TableViewColumnDidResize                                 func(notification foundation.Notification)
	_TableViewDidClickTableColumn                             func(tableView TableView, tableColumn TableColumn)
	_TableViewMouseDownInHeaderOfTableColumn                  func(tableView TableView, tableColumn TableColumn)
	_TableViewShouldTrackCellForTableColumnRow                func(tableView TableView, cell Cell, tableColumn TableColumn, row int) bool
	_TableViewRowActionsForRowEdge                            func(tableView TableView, row int, edge TableRowActionEdge) []ITableViewRowAction
}

func (di *TableViewDelegate) ImplementsTableViewViewForTableColumnRow() bool {
	return di._TableViewViewForTableColumnRow != nil
}

func (di *TableViewDelegate) SetTableViewViewForTableColumnRow(f func(tableView TableView, tableColumn TableColumn, row int) IView) {
	di._TableViewViewForTableColumnRow = f
}

func (di *TableViewDelegate) TableViewViewForTableColumnRow(tableView TableView, tableColumn TableColumn, row int) IView {
	return di._TableViewViewForTableColumnRow(tableView, tableColumn, row)
}
func (di *TableViewDelegate) ImplementsTableViewRowViewForRow() bool {
	return di._TableViewRowViewForRow != nil
}

func (di *TableViewDelegate) SetTableViewRowViewForRow(f func(tableView TableView, row int) ITableRowView) {
	di._TableViewRowViewForRow = f
}

func (di *TableViewDelegate) TableViewRowViewForRow(tableView TableView, row int) ITableRowView {
	return di._TableViewRowViewForRow(tableView, row)
}
func (di *TableViewDelegate) ImplementsTableViewDidAddRowViewForRow() bool {
	return di._TableViewDidAddRowViewForRow != nil
}

func (di *TableViewDelegate) SetTableViewDidAddRowViewForRow(f func(tableView TableView, rowView TableRowView, row int)) {
	di._TableViewDidAddRowViewForRow = f
}

func (di *TableViewDelegate) TableViewDidAddRowViewForRow(tableView TableView, rowView TableRowView, row int) {
	di._TableViewDidAddRowViewForRow(tableView, rowView, row)
}
func (di *TableViewDelegate) ImplementsTableViewDidRemoveRowViewForRow() bool {
	return di._TableViewDidRemoveRowViewForRow != nil
}

func (di *TableViewDelegate) SetTableViewDidRemoveRowViewForRow(f func(tableView TableView, rowView TableRowView, row int)) {
	di._TableViewDidRemoveRowViewForRow = f
}

func (di *TableViewDelegate) TableViewDidRemoveRowViewForRow(tableView TableView, rowView TableRowView, row int) {
	di._TableViewDidRemoveRowViewForRow(tableView, rowView, row)
}
func (di *TableViewDelegate) ImplementsTableViewIsGroupRow() bool {
	return di._TableViewIsGroupRow != nil
}

func (di *TableViewDelegate) SetTableViewIsGroupRow(f func(tableView TableView, row int) bool) {
	di._TableViewIsGroupRow = f
}

func (di *TableViewDelegate) TableViewIsGroupRow(tableView TableView, row int) bool {
	return di._TableViewIsGroupRow(tableView, row)
}
func (di *TableViewDelegate) ImplementsTableViewWillDisplayCellForTableColumnRow() bool {
	return di._TableViewWillDisplayCellForTableColumnRow != nil
}

func (di *TableViewDelegate) SetTableViewWillDisplayCellForTableColumnRow(f func(tableView TableView, cell objc.Object, tableColumn TableColumn, row int)) {
	di._TableViewWillDisplayCellForTableColumnRow = f
}

func (di *TableViewDelegate) TableViewWillDisplayCellForTableColumnRow(tableView TableView, cell objc.Object, tableColumn TableColumn, row int) {
	di._TableViewWillDisplayCellForTableColumnRow(tableView, cell, tableColumn, row)
}
func (di *TableViewDelegate) ImplementsTableViewDataCellForTableColumnRow() bool {
	return di._TableViewDataCellForTableColumnRow != nil
}

func (di *TableViewDelegate) SetTableViewDataCellForTableColumnRow(f func(tableView TableView, tableColumn TableColumn, row int) ICell) {
	di._TableViewDataCellForTableColumnRow = f
}

func (di *TableViewDelegate) TableViewDataCellForTableColumnRow(tableView TableView, tableColumn TableColumn, row int) ICell {
	return di._TableViewDataCellForTableColumnRow(tableView, tableColumn, row)
}
func (di *TableViewDelegate) ImplementsTableViewShouldShowCellExpansionForTableColumnRow() bool {
	return di._TableViewShouldShowCellExpansionForTableColumnRow != nil
}

func (di *TableViewDelegate) SetTableViewShouldShowCellExpansionForTableColumnRow(f func(tableView TableView, tableColumn TableColumn, row int) bool) {
	di._TableViewShouldShowCellExpansionForTableColumnRow = f
}

func (di *TableViewDelegate) TableViewShouldShowCellExpansionForTableColumnRow(tableView TableView, tableColumn TableColumn, row int) bool {
	return di._TableViewShouldShowCellExpansionForTableColumnRow(tableView, tableColumn, row)
}
func (di *TableViewDelegate) ImplementsTableViewToolTipForCellRectTableColumnRowMouseLocation() bool {
	return di._TableViewToolTipForCellRectTableColumnRowMouseLocation != nil
}

func (di *TableViewDelegate) SetTableViewToolTipForCellRectTableColumnRowMouseLocation(f func(tableView TableView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, row int, mouseLocation foundation.Point) string) {
	di._TableViewToolTipForCellRectTableColumnRowMouseLocation = f
}

func (di *TableViewDelegate) TableViewToolTipForCellRectTableColumnRowMouseLocation(tableView TableView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, row int, mouseLocation foundation.Point) string {
	return di._TableViewToolTipForCellRectTableColumnRowMouseLocation(tableView, cell, rect, tableColumn, row, mouseLocation)
}
func (di *TableViewDelegate) ImplementsTableViewShouldEditTableColumnRow() bool {
	return di._TableViewShouldEditTableColumnRow != nil
}

func (di *TableViewDelegate) SetTableViewShouldEditTableColumnRow(f func(tableView TableView, tableColumn TableColumn, row int) bool) {
	di._TableViewShouldEditTableColumnRow = f
}

func (di *TableViewDelegate) TableViewShouldEditTableColumnRow(tableView TableView, tableColumn TableColumn, row int) bool {
	return di._TableViewShouldEditTableColumnRow(tableView, tableColumn, row)
}
func (di *TableViewDelegate) ImplementsTableViewHeightOfRow() bool {
	return di._TableViewHeightOfRow != nil
}

func (di *TableViewDelegate) SetTableViewHeightOfRow(f func(tableView TableView, row int) float64) {
	di._TableViewHeightOfRow = f
}

func (di *TableViewDelegate) TableViewHeightOfRow(tableView TableView, row int) float64 {
	return di._TableViewHeightOfRow(tableView, row)
}
func (di *TableViewDelegate) ImplementsTableViewSizeToFitWidthOfColumn() bool {
	return di._TableViewSizeToFitWidthOfColumn != nil
}

func (di *TableViewDelegate) SetTableViewSizeToFitWidthOfColumn(f func(tableView TableView, column int) float64) {
	di._TableViewSizeToFitWidthOfColumn = f
}

func (di *TableViewDelegate) TableViewSizeToFitWidthOfColumn(tableView TableView, column int) float64 {
	return di._TableViewSizeToFitWidthOfColumn(tableView, column)
}
func (di *TableViewDelegate) ImplementsSelectionShouldChangeInTableView() bool {
	return di._SelectionShouldChangeInTableView != nil
}

func (di *TableViewDelegate) SetSelectionShouldChangeInTableView(f func(tableView TableView) bool) {
	di._SelectionShouldChangeInTableView = f
}

func (di *TableViewDelegate) SelectionShouldChangeInTableView(tableView TableView) bool {
	return di._SelectionShouldChangeInTableView(tableView)
}
func (di *TableViewDelegate) ImplementsTableViewShouldSelectRow() bool {
	return di._TableViewShouldSelectRow != nil
}

func (di *TableViewDelegate) SetTableViewShouldSelectRow(f func(tableView TableView, row int) bool) {
	di._TableViewShouldSelectRow = f
}

func (di *TableViewDelegate) TableViewShouldSelectRow(tableView TableView, row int) bool {
	return di._TableViewShouldSelectRow(tableView, row)
}
func (di *TableViewDelegate) ImplementsTableViewSelectionIndexesForProposedSelection() bool {
	return di._TableViewSelectionIndexesForProposedSelection != nil
}

func (di *TableViewDelegate) SetTableViewSelectionIndexesForProposedSelection(f func(tableView TableView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet) {
	di._TableViewSelectionIndexesForProposedSelection = f
}

func (di *TableViewDelegate) TableViewSelectionIndexesForProposedSelection(tableView TableView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet {
	return di._TableViewSelectionIndexesForProposedSelection(tableView, proposedSelectionIndexes)
}
func (di *TableViewDelegate) ImplementsTableViewShouldSelectTableColumn() bool {
	return di._TableViewShouldSelectTableColumn != nil
}

func (di *TableViewDelegate) SetTableViewShouldSelectTableColumn(f func(tableView TableView, tableColumn TableColumn) bool) {
	di._TableViewShouldSelectTableColumn = f
}

func (di *TableViewDelegate) TableViewShouldSelectTableColumn(tableView TableView, tableColumn TableColumn) bool {
	return di._TableViewShouldSelectTableColumn(tableView, tableColumn)
}
func (di *TableViewDelegate) ImplementsTableViewSelectionIsChanging() bool {
	return di._TableViewSelectionIsChanging != nil
}

func (di *TableViewDelegate) SetTableViewSelectionIsChanging(f func(notification foundation.Notification)) {
	di._TableViewSelectionIsChanging = f
}

func (di *TableViewDelegate) TableViewSelectionIsChanging(notification foundation.Notification) {
	di._TableViewSelectionIsChanging(notification)
}
func (di *TableViewDelegate) ImplementsTableViewSelectionDidChange() bool {
	return di._TableViewSelectionDidChange != nil
}

func (di *TableViewDelegate) SetTableViewSelectionDidChange(f func(notification foundation.Notification)) {
	di._TableViewSelectionDidChange = f
}

func (di *TableViewDelegate) TableViewSelectionDidChange(notification foundation.Notification) {
	di._TableViewSelectionDidChange(notification)
}
func (di *TableViewDelegate) ImplementsTableViewShouldTypeSelectForEventWithCurrentSearchString() bool {
	return di._TableViewShouldTypeSelectForEventWithCurrentSearchString != nil
}

func (di *TableViewDelegate) SetTableViewShouldTypeSelectForEventWithCurrentSearchString(f func(tableView TableView, event Event, searchString string) bool) {
	di._TableViewShouldTypeSelectForEventWithCurrentSearchString = f
}

func (di *TableViewDelegate) TableViewShouldTypeSelectForEventWithCurrentSearchString(tableView TableView, event Event, searchString string) bool {
	return di._TableViewShouldTypeSelectForEventWithCurrentSearchString(tableView, event, searchString)
}
func (di *TableViewDelegate) ImplementsTableViewTypeSelectStringForTableColumnRow() bool {
	return di._TableViewTypeSelectStringForTableColumnRow != nil
}

func (di *TableViewDelegate) SetTableViewTypeSelectStringForTableColumnRow(f func(tableView TableView, tableColumn TableColumn, row int) string) {
	di._TableViewTypeSelectStringForTableColumnRow = f
}

func (di *TableViewDelegate) TableViewTypeSelectStringForTableColumnRow(tableView TableView, tableColumn TableColumn, row int) string {
	return di._TableViewTypeSelectStringForTableColumnRow(tableView, tableColumn, row)
}
func (di *TableViewDelegate) ImplementsTableViewNextTypeSelectMatchFromRowToRowForString() bool {
	return di._TableViewNextTypeSelectMatchFromRowToRowForString != nil
}

func (di *TableViewDelegate) SetTableViewNextTypeSelectMatchFromRowToRowForString(f func(tableView TableView, startRow int, endRow int, searchString string) int) {
	di._TableViewNextTypeSelectMatchFromRowToRowForString = f
}

func (di *TableViewDelegate) TableViewNextTypeSelectMatchFromRowToRowForString(tableView TableView, startRow int, endRow int, searchString string) int {
	return di._TableViewNextTypeSelectMatchFromRowToRowForString(tableView, startRow, endRow, searchString)
}
func (di *TableViewDelegate) ImplementsTableViewShouldReorderColumnToColumn() bool {
	return di._TableViewShouldReorderColumnToColumn != nil
}

func (di *TableViewDelegate) SetTableViewShouldReorderColumnToColumn(f func(tableView TableView, columnIndex int, newColumnIndex int) bool) {
	di._TableViewShouldReorderColumnToColumn = f
}

func (di *TableViewDelegate) TableViewShouldReorderColumnToColumn(tableView TableView, columnIndex int, newColumnIndex int) bool {
	return di._TableViewShouldReorderColumnToColumn(tableView, columnIndex, newColumnIndex)
}
func (di *TableViewDelegate) ImplementsTableViewDidDragTableColumn() bool {
	return di._TableViewDidDragTableColumn != nil
}

func (di *TableViewDelegate) SetTableViewDidDragTableColumn(f func(tableView TableView, tableColumn TableColumn)) {
	di._TableViewDidDragTableColumn = f
}

func (di *TableViewDelegate) TableViewDidDragTableColumn(tableView TableView, tableColumn TableColumn) {
	di._TableViewDidDragTableColumn(tableView, tableColumn)
}
func (di *TableViewDelegate) ImplementsTableViewColumnDidMove() bool {
	return di._TableViewColumnDidMove != nil
}

func (di *TableViewDelegate) SetTableViewColumnDidMove(f func(notification foundation.Notification)) {
	di._TableViewColumnDidMove = f
}

func (di *TableViewDelegate) TableViewColumnDidMove(notification foundation.Notification) {
	di._TableViewColumnDidMove(notification)
}
func (di *TableViewDelegate) ImplementsTableViewColumnDidResize() bool {
	return di._TableViewColumnDidResize != nil
}

func (di *TableViewDelegate) SetTableViewColumnDidResize(f func(notification foundation.Notification)) {
	di._TableViewColumnDidResize = f
}

func (di *TableViewDelegate) TableViewColumnDidResize(notification foundation.Notification) {
	di._TableViewColumnDidResize(notification)
}
func (di *TableViewDelegate) ImplementsTableViewDidClickTableColumn() bool {
	return di._TableViewDidClickTableColumn != nil
}

func (di *TableViewDelegate) SetTableViewDidClickTableColumn(f func(tableView TableView, tableColumn TableColumn)) {
	di._TableViewDidClickTableColumn = f
}

func (di *TableViewDelegate) TableViewDidClickTableColumn(tableView TableView, tableColumn TableColumn) {
	di._TableViewDidClickTableColumn(tableView, tableColumn)
}
func (di *TableViewDelegate) ImplementsTableViewMouseDownInHeaderOfTableColumn() bool {
	return di._TableViewMouseDownInHeaderOfTableColumn != nil
}

func (di *TableViewDelegate) SetTableViewMouseDownInHeaderOfTableColumn(f func(tableView TableView, tableColumn TableColumn)) {
	di._TableViewMouseDownInHeaderOfTableColumn = f
}

func (di *TableViewDelegate) TableViewMouseDownInHeaderOfTableColumn(tableView TableView, tableColumn TableColumn) {
	di._TableViewMouseDownInHeaderOfTableColumn(tableView, tableColumn)
}
func (di *TableViewDelegate) ImplementsTableViewShouldTrackCellForTableColumnRow() bool {
	return di._TableViewShouldTrackCellForTableColumnRow != nil
}

func (di *TableViewDelegate) SetTableViewShouldTrackCellForTableColumnRow(f func(tableView TableView, cell Cell, tableColumn TableColumn, row int) bool) {
	di._TableViewShouldTrackCellForTableColumnRow = f
}

func (di *TableViewDelegate) TableViewShouldTrackCellForTableColumnRow(tableView TableView, cell Cell, tableColumn TableColumn, row int) bool {
	return di._TableViewShouldTrackCellForTableColumnRow(tableView, cell, tableColumn, row)
}
func (di *TableViewDelegate) ImplementsTableViewRowActionsForRowEdge() bool {
	return di._TableViewRowActionsForRowEdge != nil
}

func (di *TableViewDelegate) SetTableViewRowActionsForRowEdge(f func(tableView TableView, row int, edge TableRowActionEdge) []ITableViewRowAction) {
	di._TableViewRowActionsForRowEdge = f
}

func (di *TableViewDelegate) TableViewRowActionsForRowEdge(tableView TableView, row int, edge TableRowActionEdge) []ITableViewRowAction {
	return di._TableViewRowActionsForRowEdge(tableView, row, edge)
}

type TableViewDelegateWrapper struct {
	ControlTextEditingDelegateWrapper
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewViewForTableColumnRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:viewForTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableViewViewForTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) View {
	rv := objc.CallMethod[View](t_, objc.GetSelector("tableView:viewForTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn), row)
	return rv
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewRowViewForRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:rowViewForRow:"))
}

func (t_ TableViewDelegateWrapper) TableViewRowViewForRow(tableView ITableView, row int) TableRowView {
	rv := objc.CallMethod[TableRowView](t_, objc.GetSelector("tableView:rowViewForRow:"), objc.ExtractPtr(tableView), row)
	return rv
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewDidAddRowViewForRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:didAddRowView:forRow:"))
}

func (t_ TableViewDelegateWrapper) TableViewDidAddRowViewForRow(tableView ITableView, rowView ITableRowView, row int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:didAddRowView:forRow:"), objc.ExtractPtr(tableView), objc.ExtractPtr(rowView), row)
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewDidRemoveRowViewForRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:didRemoveRowView:forRow:"))
}

func (t_ TableViewDelegateWrapper) TableViewDidRemoveRowViewForRow(tableView ITableView, rowView ITableRowView, row int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:didRemoveRowView:forRow:"), objc.ExtractPtr(tableView), objc.ExtractPtr(rowView), row)
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewIsGroupRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:isGroupRow:"))
}

func (t_ TableViewDelegateWrapper) TableViewIsGroupRow(tableView ITableView, row int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:isGroupRow:"), objc.ExtractPtr(tableView), row)
	return rv
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewWillDisplayCellForTableColumnRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:willDisplayCell:forTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableViewWillDisplayCellForTableColumnRow(tableView ITableView, cell objc.IObject, tableColumn ITableColumn, row int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:willDisplayCell:forTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(cell), objc.ExtractPtr(tableColumn), row)
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewDataCellForTableColumnRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:dataCellForTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableViewDataCellForTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) Cell {
	rv := objc.CallMethod[Cell](t_, objc.GetSelector("tableView:dataCellForTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn), row)
	return rv
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewShouldShowCellExpansionForTableColumnRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldShowCellExpansionForTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableViewShouldShowCellExpansionForTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:shouldShowCellExpansionForTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn), row)
	return rv
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewToolTipForCellRectTableColumnRowMouseLocation() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:toolTipForCell:rect:tableColumn:row:mouseLocation:"))
}

func (t_ TableViewDelegateWrapper) TableViewToolTipForCellRectTableColumnRowMouseLocation(tableView ITableView, cell ICell, rect *foundation.Rect, tableColumn ITableColumn, row int, mouseLocation foundation.Point) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("tableView:toolTipForCell:rect:tableColumn:row:mouseLocation:"), objc.ExtractPtr(tableView), objc.ExtractPtr(cell), rect, objc.ExtractPtr(tableColumn), row, mouseLocation)
	return rv
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewShouldEditTableColumnRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldEditTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableViewShouldEditTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:shouldEditTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn), row)
	return rv
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewHeightOfRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:heightOfRow:"))
}

func (t_ TableViewDelegateWrapper) TableViewHeightOfRow(tableView ITableView, row int) float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("tableView:heightOfRow:"), objc.ExtractPtr(tableView), row)
	return rv
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewSizeToFitWidthOfColumn() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:sizeToFitWidthOfColumn:"))
}

func (t_ TableViewDelegateWrapper) TableViewSizeToFitWidthOfColumn(tableView ITableView, column int) float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("tableView:sizeToFitWidthOfColumn:"), objc.ExtractPtr(tableView), column)
	return rv
}

func (t_ TableViewDelegateWrapper) ImplementsSelectionShouldChangeInTableView() bool {
	return t_.RespondsToSelector(objc.GetSelector("selectionShouldChangeInTableView:"))
}

func (t_ TableViewDelegateWrapper) SelectionShouldChangeInTableView(tableView ITableView) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("selectionShouldChangeInTableView:"), objc.ExtractPtr(tableView))
	return rv
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewShouldSelectRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldSelectRow:"))
}

func (t_ TableViewDelegateWrapper) TableViewShouldSelectRow(tableView ITableView, row int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:shouldSelectRow:"), objc.ExtractPtr(tableView), row)
	return rv
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewSelectionIndexesForProposedSelection() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:selectionIndexesForProposedSelection:"))
}

func (t_ TableViewDelegateWrapper) TableViewSelectionIndexesForProposedSelection(tableView ITableView, proposedSelectionIndexes foundation.IIndexSet) foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](t_, objc.GetSelector("tableView:selectionIndexesForProposedSelection:"), objc.ExtractPtr(tableView), objc.ExtractPtr(proposedSelectionIndexes))
	return rv
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewShouldSelectTableColumn() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldSelectTableColumn:"))
}

func (t_ TableViewDelegateWrapper) TableViewShouldSelectTableColumn(tableView ITableView, tableColumn ITableColumn) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:shouldSelectTableColumn:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn))
	return rv
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewSelectionIsChanging() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableViewSelectionIsChanging:"))
}

func (t_ TableViewDelegateWrapper) TableViewSelectionIsChanging(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableViewSelectionIsChanging:"), objc.ExtractPtr(notification))
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewSelectionDidChange() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableViewSelectionDidChange:"))
}

func (t_ TableViewDelegateWrapper) TableViewSelectionDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableViewSelectionDidChange:"), objc.ExtractPtr(notification))
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewShouldTypeSelectForEventWithCurrentSearchString() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldTypeSelectForEvent:withCurrentSearchString:"))
}

func (t_ TableViewDelegateWrapper) TableViewShouldTypeSelectForEventWithCurrentSearchString(tableView ITableView, event IEvent, searchString string) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:shouldTypeSelectForEvent:withCurrentSearchString:"), objc.ExtractPtr(tableView), objc.ExtractPtr(event), searchString)
	return rv
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewTypeSelectStringForTableColumnRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:typeSelectStringForTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableViewTypeSelectStringForTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("tableView:typeSelectStringForTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn), row)
	return rv
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewNextTypeSelectMatchFromRowToRowForString() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:nextTypeSelectMatchFromRow:toRow:forString:"))
}

func (t_ TableViewDelegateWrapper) TableViewNextTypeSelectMatchFromRowToRowForString(tableView ITableView, startRow int, endRow int, searchString string) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("tableView:nextTypeSelectMatchFromRow:toRow:forString:"), objc.ExtractPtr(tableView), startRow, endRow, searchString)
	return rv
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewShouldReorderColumnToColumn() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldReorderColumn:toColumn:"))
}

func (t_ TableViewDelegateWrapper) TableViewShouldReorderColumnToColumn(tableView ITableView, columnIndex int, newColumnIndex int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:shouldReorderColumn:toColumn:"), objc.ExtractPtr(tableView), columnIndex, newColumnIndex)
	return rv
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewDidDragTableColumn() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:didDragTableColumn:"))
}

func (t_ TableViewDelegateWrapper) TableViewDidDragTableColumn(tableView ITableView, tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:didDragTableColumn:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn))
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewColumnDidMove() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableViewColumnDidMove:"))
}

func (t_ TableViewDelegateWrapper) TableViewColumnDidMove(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableViewColumnDidMove:"), objc.ExtractPtr(notification))
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewColumnDidResize() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableViewColumnDidResize:"))
}

func (t_ TableViewDelegateWrapper) TableViewColumnDidResize(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableViewColumnDidResize:"), objc.ExtractPtr(notification))
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewDidClickTableColumn() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:didClickTableColumn:"))
}

func (t_ TableViewDelegateWrapper) TableViewDidClickTableColumn(tableView ITableView, tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:didClickTableColumn:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn))
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewMouseDownInHeaderOfTableColumn() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:mouseDownInHeaderOfTableColumn:"))
}

func (t_ TableViewDelegateWrapper) TableViewMouseDownInHeaderOfTableColumn(tableView ITableView, tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:mouseDownInHeaderOfTableColumn:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn))
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewShouldTrackCellForTableColumnRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldTrackCell:forTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableViewShouldTrackCellForTableColumnRow(tableView ITableView, cell ICell, tableColumn ITableColumn, row int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:shouldTrackCell:forTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(cell), objc.ExtractPtr(tableColumn), row)
	return rv
}

func (t_ TableViewDelegateWrapper) ImplementsTableViewRowActionsForRowEdge() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:rowActionsForRow:edge:"))
}

func (t_ TableViewDelegateWrapper) TableViewRowActionsForRowEdge(tableView ITableView, row int, edge TableRowActionEdge) []TableViewRowAction {
	rv := objc.CallMethod[[]TableViewRowAction](t_, objc.GetSelector("tableView:rowActionsForRow:edge:"), objc.ExtractPtr(tableView), row, edge)
	// TODO: convert slice items...
	return rv
}
