package main

import (
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

type TableViewDataSourceDelegate struct {
	_TableViewSetObjectValueForTableColumnRow              func(tableView appkit.TableView, object objc.Object, tableColumn appkit.TableColumn, row int)
	_NumberOfRowsInTableView                               func(tableView appkit.TableView) int
	_TableViewSortDescriptorsDidChange                     func(tableView appkit.TableView, oldDescriptors []foundation.SortDescriptor)
	_TableViewDraggingSessionEndedAtPointOperation         func(tableView appkit.TableView, session appkit.DraggingSession, screenPoint foundation.Point, operation appkit.DragOperation)
	_TableViewDraggingSessionWillBeginAtPointForRowIndexes func(tableView appkit.TableView, session appkit.DraggingSession, screenPoint foundation.Point, rowIndexes foundation.IndexSet)
	_TableViewAcceptDropRowDropOperation                   func(tableView appkit.TableView, info appkit.DraggingInfoObject, row int, dropOperation appkit.TableViewDropOperation) bool
	_TableViewObjectValueForTableColumnRow                 func(tableView appkit.TableView, tableColumn appkit.TableColumn, row int) objc.Object
	_TableViewPasteboardWriterForRow                       func(tableView appkit.TableView, row int) appkit.PasteboardWritingObject
	_TableViewUpdateDraggingItemsForDrag                   func(tableView appkit.TableView, draggingInfo appkit.DraggingInfoObject)
	_TableViewValidateDropProposedRowProposedDropOperation func(tableView appkit.TableView, info appkit.DraggingInfoObject, row int, dropOperation appkit.TableViewDropOperation) appkit.DragOperation
}

func (t *TableViewDataSourceDelegate) TableViewSetObjectValueForTableColumnRow(tableView appkit.TableView, object objc.Object, tableColumn appkit.TableColumn, row int) {
	t._TableViewObjectValueForTableColumnRow(tableView, tableColumn, row)
}

func (t *TableViewDataSourceDelegate) HasTableViewSetObjectValueForTableColumnRow() bool {
	if t._TableViewSetObjectValueForTableColumnRow != nil {
		return true
	}
	return false
}

func (t *TableViewDataSourceDelegate) SetTableViewSetObjectValueForTableColumnRow(f func(tableView appkit.TableView, object objc.Object, tableColumn appkit.TableColumn, row int)) {
	t._TableViewSetObjectValueForTableColumnRow = f
}

func (t *TableViewDataSourceDelegate) NumberOfRowsInTableView(tableView appkit.TableView) int {
	return t._NumberOfRowsInTableView(tableView)
}

func (t *TableViewDataSourceDelegate) HasNumberOfRowsInTableView() bool {
	if t._NumberOfRowsInTableView != nil {
		return true
	}
	return false
}

func (t *TableViewDataSourceDelegate) SetNumberOfRowsInTableView(f func(tableView appkit.TableView) int) {
	t._NumberOfRowsInTableView = f
}

func (t *TableViewDataSourceDelegate) TableViewSortDescriptorsDidChange(tableView appkit.TableView, oldDescriptors []foundation.SortDescriptor) {
	t._TableViewSortDescriptorsDidChange(tableView, oldDescriptors)
}

func (t *TableViewDataSourceDelegate) HasTableViewSortDescriptorsDidChange() bool {
	if t._TableViewSortDescriptorsDidChange != nil {
		return true
	}
	return false
}

func (t *TableViewDataSourceDelegate) SetTableViewSortDescriptorsDidChange(f func(tableView appkit.TableView, oldDescriptors []foundation.SortDescriptor)) {
	t._TableViewSortDescriptorsDidChange = f
}

func (t *TableViewDataSourceDelegate) TableViewDraggingSessionEndedAtPointOperation(tableView appkit.TableView, session appkit.DraggingSession, screenPoint foundation.Point, operation appkit.DragOperation) {
	t._TableViewDraggingSessionEndedAtPointOperation(tableView, session, screenPoint, operation)
}

func (t *TableViewDataSourceDelegate) HasTableViewDraggingSessionEndedAtPointOperation() bool {
	if t._TableViewDraggingSessionEndedAtPointOperation != nil {
		return true
	}
	return false
}

func (t *TableViewDataSourceDelegate) SetTableViewDraggingSessionEndedAtPointOperation(f func(tableView appkit.TableView, session appkit.DraggingSession, screenPoint foundation.Point, operation appkit.DragOperation)) {
	t._TableViewDraggingSessionEndedAtPointOperation = f
}

func (t *TableViewDataSourceDelegate) TableViewDraggingSessionWillBeginAtPointForRowIndexes(tableView appkit.TableView, session appkit.DraggingSession, screenPoint foundation.Point, rowIndexes foundation.IndexSet) {
	t._TableViewDraggingSessionWillBeginAtPointForRowIndexes(tableView, session, screenPoint, rowIndexes)
}

func (t *TableViewDataSourceDelegate) HasTableViewDraggingSessionWillBeginAtPointForRowIndexes() bool {
	if t._TableViewDraggingSessionWillBeginAtPointForRowIndexes != nil {
		return true
	}
	return false
}

func (t *TableViewDataSourceDelegate) SetTableViewDraggingSessionWillBeginAtPointForRowIndexes(f func(tableView appkit.TableView, session appkit.DraggingSession, screenPoint foundation.Point, rowIndexes foundation.IndexSet)) {
	t._TableViewDraggingSessionWillBeginAtPointForRowIndexes = f
}

func (t *TableViewDataSourceDelegate) TableViewAcceptDropRowDropOperation(tableView appkit.TableView, info appkit.DraggingInfoObject, row int, dropOperation appkit.TableViewDropOperation) bool {
	return t._TableViewAcceptDropRowDropOperation(tableView, info, row, dropOperation)
}

func (t *TableViewDataSourceDelegate) HasTableViewAcceptDropRowDropOperation() bool {
	if t._TableViewAcceptDropRowDropOperation != nil {
		return true
	}
	return false
}

func (t *TableViewDataSourceDelegate) SetTableViewAcceptDropRowDropOperation(f func(tableView appkit.TableView, info appkit.DraggingInfoObject, row int, dropOperation appkit.TableViewDropOperation) bool) {
	t._TableViewAcceptDropRowDropOperation = f
}

func (t *TableViewDataSourceDelegate) TableViewObjectValueForTableColumnRow(tableView appkit.TableView, tableColumn appkit.TableColumn, row int) objc.Object {
	return t._TableViewObjectValueForTableColumnRow(tableView, tableColumn, row)
}

func (t *TableViewDataSourceDelegate) HasTableViewObjectValueForTableColumnRow() bool {
	if t._TableViewObjectValueForTableColumnRow != nil {
		return true
	}
	return false
}

func (t *TableViewDataSourceDelegate) SetTableViewObjectValueForTableColumnRow(f func(tableView appkit.TableView, tableColumn appkit.TableColumn, row int) objc.Object) {
	t._TableViewObjectValueForTableColumnRow = f
}

func (t *TableViewDataSourceDelegate) TableViewPasteboardWriterForRow(tableView appkit.TableView, row int) appkit.PasteboardWritingObject {
	return t._TableViewPasteboardWriterForRow(tableView, row)
}

func (t *TableViewDataSourceDelegate) HasTableViewPasteboardWriterForRow() bool {
	if t._TableViewPasteboardWriterForRow != nil {
		return true
	}
	return false
}

func (t *TableViewDataSourceDelegate) SetTableViewPasteboardWriterForRow(f func(tableView appkit.TableView, row int) appkit.PasteboardWritingObject) {
	t._TableViewPasteboardWriterForRow = f
}

func (t *TableViewDataSourceDelegate) TableViewUpdateDraggingItemsForDrag(tableView appkit.TableView, draggingInfo appkit.DraggingInfoObject) {
	t._TableViewUpdateDraggingItemsForDrag(tableView, draggingInfo)
}

func (t *TableViewDataSourceDelegate) HasTableViewUpdateDraggingItemsForDrag() bool {
	if t._TableViewUpdateDraggingItemsForDrag != nil {
		return true
	}
	return false
}

func (t *TableViewDataSourceDelegate) SetTableViewUpdateDraggingItemsForDrag(f func(tableView appkit.TableView, draggingInfo appkit.DraggingInfoObject)) {
	t._TableViewUpdateDraggingItemsForDrag = f
}

func (t *TableViewDataSourceDelegate) TableViewValidateDropProposedRowProposedDropOperation(tableView appkit.TableView, info appkit.DraggingInfoObject, row int, dropOperation appkit.TableViewDropOperation) appkit.DragOperation {
	return t._TableViewValidateDropProposedRowProposedDropOperation(tableView, info, row, dropOperation)
}

func (t *TableViewDataSourceDelegate) HasTableViewValidateDropProposedRowProposedDropOperation() bool {
	if t._TableViewValidateDropProposedRowProposedDropOperation != nil {
		return true
	}
	return false
}

func (t *TableViewDataSourceDelegate) SetTableViewValidateDropProposedRowProposedDropOperation(f func(tableView appkit.TableView, info appkit.DraggingInfoObject, row int, dropOperation appkit.TableViewDropOperation) appkit.DragOperation) {
	t._TableViewValidateDropProposedRowProposedDropOperation = f
}
