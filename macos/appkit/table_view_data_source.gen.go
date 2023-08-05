// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type ITableViewDataSource interface {
	ImplementsNumberOfRowsInTableView() bool
	// optional
	NumberOfRowsInTableView(tableView TableView) int
	ImplementsTableViewObjectValueForTableColumnRow() bool
	// optional
	TableViewObjectValueForTableColumnRow(tableView TableView, tableColumn TableColumn, row int) objc.IObject
	ImplementsTableViewSetObjectValueForTableColumnRow() bool
	// optional
	TableViewSetObjectValueForTableColumnRow(tableView TableView, object objc.Object, tableColumn TableColumn, row int)
	ImplementsTableViewPasteboardWriterForRow() bool
	// optional
	TableViewPasteboardWriterForRow(tableView TableView, row int) IPasteboardWriting
	ImplementsTableViewAcceptDropRowDropOperation() bool
	// optional
	TableViewAcceptDropRowDropOperation(tableView TableView, info DraggingInfoWrapper, row int, dropOperation TableViewDropOperation) bool
	ImplementsTableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes() bool
	// optional
	// deprecated
	TableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes(tableView TableView, dropDestination foundation.URL, indexSet foundation.IndexSet) []string
	ImplementsTableViewValidateDropProposedRowProposedDropOperation() bool
	// optional
	TableViewValidateDropProposedRowProposedDropOperation(tableView TableView, info DraggingInfoWrapper, row int, dropOperation TableViewDropOperation) DragOperation
	ImplementsTableViewWriteRowsWithIndexesToPasteboard() bool
	// optional
	// deprecated
	TableViewWriteRowsWithIndexesToPasteboard(tableView TableView, rowIndexes foundation.IndexSet, pboard Pasteboard) bool
	ImplementsTableViewDraggingSessionWillBeginAtPointForRowIndexes() bool
	// optional
	TableViewDraggingSessionWillBeginAtPointForRowIndexes(tableView TableView, session DraggingSession, screenPoint foundation.Point, rowIndexes foundation.IndexSet)
	ImplementsTableViewUpdateDraggingItemsForDrag() bool
	// optional
	TableViewUpdateDraggingItemsForDrag(tableView TableView, draggingInfo DraggingInfoWrapper)
	ImplementsTableViewDraggingSessionEndedAtPointOperation() bool
	// optional
	TableViewDraggingSessionEndedAtPointOperation(tableView TableView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	ImplementsTableViewSortDescriptorsDidChange() bool
	// optional
	TableViewSortDescriptorsDidChange(tableView TableView, oldDescriptors []foundation.SortDescriptor)
}

type TableViewDataSourceWrapper struct {
	objc.Object
}

func (t_ TableViewDataSourceWrapper) ImplementsNumberOfRowsInTableView() bool {
	return t_.RespondsToSelector(objc.GetSelector("numberOfRowsInTableView:"))
}

func (t_ TableViewDataSourceWrapper) NumberOfRowsInTableView(tableView ITableView) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("numberOfRowsInTableView:"), objc.ExtractPtr(tableView))
	return rv
}

func (t_ TableViewDataSourceWrapper) ImplementsTableViewObjectValueForTableColumnRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:objectValueForTableColumn:row:"))
}

func (t_ TableViewDataSourceWrapper) TableViewObjectValueForTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.GetSelector("tableView:objectValueForTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn), row)
	return rv
}

func (t_ TableViewDataSourceWrapper) ImplementsTableViewSetObjectValueForTableColumnRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:setObjectValue:forTableColumn:row:"))
}

func (t_ TableViewDataSourceWrapper) TableViewSetObjectValueForTableColumnRow(tableView ITableView, object objc.IObject, tableColumn ITableColumn, row int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:setObjectValue:forTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(object), objc.ExtractPtr(tableColumn), row)
}

func (t_ TableViewDataSourceWrapper) ImplementsTableViewPasteboardWriterForRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:pasteboardWriterForRow:"))
}

func (t_ TableViewDataSourceWrapper) TableViewPasteboardWriterForRow(tableView ITableView, row int) PasteboardWritingWrapper {
	rv := objc.CallMethod[PasteboardWritingWrapper](t_, objc.GetSelector("tableView:pasteboardWriterForRow:"), objc.ExtractPtr(tableView), row)
	return rv
}

func (t_ TableViewDataSourceWrapper) ImplementsTableViewAcceptDropRowDropOperation() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:acceptDrop:row:dropOperation:"))
}

func (t_ TableViewDataSourceWrapper) TableViewAcceptDropRowDropOperation(tableView ITableView, info IDraggingInfo, row int, dropOperation TableViewDropOperation) bool {
	po := objc.WrapAsProtocol("NSDraggingInfo", info)
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:acceptDrop:row:dropOperation:"), objc.ExtractPtr(tableView), po, row, dropOperation)
	return rv
}

func (t_ TableViewDataSourceWrapper) ImplementsTableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:namesOfPromisedFilesDroppedAtDestination:forDraggedRowsWithIndexes:"))
}

func (t_ TableViewDataSourceWrapper) ImplementsTableViewValidateDropProposedRowProposedDropOperation() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:validateDrop:proposedRow:proposedDropOperation:"))
}

func (t_ TableViewDataSourceWrapper) TableViewValidateDropProposedRowProposedDropOperation(tableView ITableView, info IDraggingInfo, row int, dropOperation TableViewDropOperation) DragOperation {
	po := objc.WrapAsProtocol("NSDraggingInfo", info)
	rv := objc.CallMethod[DragOperation](t_, objc.GetSelector("tableView:validateDrop:proposedRow:proposedDropOperation:"), objc.ExtractPtr(tableView), po, row, dropOperation)
	return rv
}

func (t_ TableViewDataSourceWrapper) ImplementsTableViewWriteRowsWithIndexesToPasteboard() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:writeRowsWithIndexes:toPasteboard:"))
}

func (t_ TableViewDataSourceWrapper) ImplementsTableViewDraggingSessionWillBeginAtPointForRowIndexes() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:draggingSession:willBeginAtPoint:forRowIndexes:"))
}

func (t_ TableViewDataSourceWrapper) TableViewDraggingSessionWillBeginAtPointForRowIndexes(tableView ITableView, session IDraggingSession, screenPoint foundation.Point, rowIndexes foundation.IIndexSet) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:draggingSession:willBeginAtPoint:forRowIndexes:"), objc.ExtractPtr(tableView), objc.ExtractPtr(session), screenPoint, objc.ExtractPtr(rowIndexes))
}

func (t_ TableViewDataSourceWrapper) ImplementsTableViewUpdateDraggingItemsForDrag() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:updateDraggingItemsForDrag:"))
}

func (t_ TableViewDataSourceWrapper) TableViewUpdateDraggingItemsForDrag(tableView ITableView, draggingInfo IDraggingInfo) {
	po := objc.WrapAsProtocol("NSDraggingInfo", draggingInfo)
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:updateDraggingItemsForDrag:"), objc.ExtractPtr(tableView), po)
}

func (t_ TableViewDataSourceWrapper) ImplementsTableViewDraggingSessionEndedAtPointOperation() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:draggingSession:endedAtPoint:operation:"))
}

func (t_ TableViewDataSourceWrapper) TableViewDraggingSessionEndedAtPointOperation(tableView ITableView, session IDraggingSession, screenPoint foundation.Point, operation DragOperation) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:draggingSession:endedAtPoint:operation:"), objc.ExtractPtr(tableView), objc.ExtractPtr(session), screenPoint, operation)
}

func (t_ TableViewDataSourceWrapper) ImplementsTableViewSortDescriptorsDidChange() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:sortDescriptorsDidChange:"))
}

func (t_ TableViewDataSourceWrapper) TableViewSortDescriptorsDidChange(tableView ITableView, oldDescriptors []foundation.ISortDescriptor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:sortDescriptorsDidChange:"), objc.ExtractPtr(tableView), oldDescriptors)
}
