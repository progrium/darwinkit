// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IOutlineViewDataSource interface {
	ImplementsOutlineViewAcceptDropItemChildIndex() bool
	// optional
	OutlineViewAcceptDropItemChildIndex(outlineView OutlineView, info DraggingInfoWrapper, item objc.Object, index int) bool
	ImplementsOutlineViewChildOfItem() bool
	// optional
	OutlineViewChildOfItem(outlineView OutlineView, index int, item objc.Object) objc.IObject
	ImplementsOutlineViewDraggingSessionEndedAtPointOperation() bool
	// optional
	OutlineViewDraggingSessionEndedAtPointOperation(outlineView OutlineView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	ImplementsOutlineViewDraggingSessionWillBeginAtPointForItems() bool
	// optional
	OutlineViewDraggingSessionWillBeginAtPointForItems(outlineView OutlineView, session DraggingSession, screenPoint foundation.Point, draggedItems []objc.Object)
	ImplementsOutlineViewIsItemExpandable() bool
	// optional
	OutlineViewIsItemExpandable(outlineView OutlineView, item objc.Object) bool
	ImplementsOutlineViewItemForPersistentObject() bool
	// optional
	OutlineViewItemForPersistentObject(outlineView OutlineView, object objc.Object) objc.IObject
	ImplementsOutlineViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItems() bool
	// optional
	// deprecated
	OutlineViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItems(outlineView OutlineView, dropDestination foundation.URL, items []objc.Object) []string
	ImplementsOutlineViewNumberOfChildrenOfItem() bool
	// optional
	OutlineViewNumberOfChildrenOfItem(outlineView OutlineView, item objc.Object) int
	ImplementsOutlineViewObjectValueForTableColumnByItem() bool
	// optional
	OutlineViewObjectValueForTableColumnByItem(outlineView OutlineView, tableColumn TableColumn, item objc.Object) objc.IObject
	ImplementsOutlineViewPasteboardWriterForItem() bool
	// optional
	OutlineViewPasteboardWriterForItem(outlineView OutlineView, item objc.Object) IPasteboardWriting
	ImplementsOutlineViewPersistentObjectForItem() bool
	// optional
	OutlineViewPersistentObjectForItem(outlineView OutlineView, item objc.Object) objc.IObject
	ImplementsOutlineViewSetObjectValueForTableColumnByItem() bool
	// optional
	OutlineViewSetObjectValueForTableColumnByItem(outlineView OutlineView, object objc.Object, tableColumn TableColumn, item objc.Object)
	ImplementsOutlineViewSortDescriptorsDidChange() bool
	// optional
	OutlineViewSortDescriptorsDidChange(outlineView OutlineView, oldDescriptors []foundation.SortDescriptor)
	ImplementsOutlineViewUpdateDraggingItemsForDrag() bool
	// optional
	OutlineViewUpdateDraggingItemsForDrag(outlineView OutlineView, draggingInfo DraggingInfoWrapper)
	ImplementsOutlineViewValidateDropProposedItemProposedChildIndex() bool
	// optional
	OutlineViewValidateDropProposedItemProposedChildIndex(outlineView OutlineView, info DraggingInfoWrapper, item objc.Object, index int) DragOperation
	ImplementsOutlineViewWriteItemsToPasteboard() bool
	// optional
	// deprecated
	OutlineViewWriteItemsToPasteboard(outlineView OutlineView, items []objc.Object, pasteboard Pasteboard) bool
}

type OutlineViewDataSourceWrapper struct {
	objc.Object
}

func (o_ OutlineViewDataSourceWrapper) ImplementsOutlineViewAcceptDropItemChildIndex() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:acceptDrop:item:childIndex:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineViewAcceptDropItemChildIndex(outlineView IOutlineView, info IDraggingInfo, item objc.IObject, index int) bool {
	po := objc.WrapAsProtocol("NSDraggingInfo", info)
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:acceptDrop:item:childIndex:"), objc.ExtractPtr(outlineView), po, objc.ExtractPtr(item), index)
	return rv
}

func (o_ OutlineViewDataSourceWrapper) ImplementsOutlineViewChildOfItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:child:ofItem:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineViewChildOfItem(outlineView IOutlineView, index int, item objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("outlineView:child:ofItem:"), objc.ExtractPtr(outlineView), index, objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDataSourceWrapper) ImplementsOutlineViewDraggingSessionEndedAtPointOperation() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:draggingSession:endedAtPoint:operation:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineViewDraggingSessionEndedAtPointOperation(outlineView IOutlineView, session IDraggingSession, screenPoint foundation.Point, operation DragOperation) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:draggingSession:endedAtPoint:operation:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(session), screenPoint, operation)
}

func (o_ OutlineViewDataSourceWrapper) ImplementsOutlineViewDraggingSessionWillBeginAtPointForItems() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:draggingSession:willBeginAtPoint:forItems:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineViewDraggingSessionWillBeginAtPointForItems(outlineView IOutlineView, session IDraggingSession, screenPoint foundation.Point, draggedItems []objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:draggingSession:willBeginAtPoint:forItems:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(session), screenPoint, draggedItems)
}

func (o_ OutlineViewDataSourceWrapper) ImplementsOutlineViewIsItemExpandable() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:isItemExpandable:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineViewIsItemExpandable(outlineView IOutlineView, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:isItemExpandable:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDataSourceWrapper) ImplementsOutlineViewItemForPersistentObject() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:itemForPersistentObject:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineViewItemForPersistentObject(outlineView IOutlineView, object objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("outlineView:itemForPersistentObject:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(object))
	return rv
}

func (o_ OutlineViewDataSourceWrapper) ImplementsOutlineViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItems() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:namesOfPromisedFilesDroppedAtDestination:forDraggedItems:"))
}

func (o_ OutlineViewDataSourceWrapper) ImplementsOutlineViewNumberOfChildrenOfItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:numberOfChildrenOfItem:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineViewNumberOfChildrenOfItem(outlineView IOutlineView, item objc.IObject) int {
	rv := objc.CallMethod[int](o_, objc.GetSelector("outlineView:numberOfChildrenOfItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDataSourceWrapper) ImplementsOutlineViewObjectValueForTableColumnByItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:objectValueForTableColumn:byItem:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineViewObjectValueForTableColumnByItem(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("outlineView:objectValueForTableColumn:byItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDataSourceWrapper) ImplementsOutlineViewPasteboardWriterForItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:pasteboardWriterForItem:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineViewPasteboardWriterForItem(outlineView IOutlineView, item objc.IObject) PasteboardWritingWrapper {
	rv := objc.CallMethod[PasteboardWritingWrapper](o_, objc.GetSelector("outlineView:pasteboardWriterForItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDataSourceWrapper) ImplementsOutlineViewPersistentObjectForItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:persistentObjectForItem:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineViewPersistentObjectForItem(outlineView IOutlineView, item objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("outlineView:persistentObjectForItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDataSourceWrapper) ImplementsOutlineViewSetObjectValueForTableColumnByItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:setObjectValue:forTableColumn:byItem:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineViewSetObjectValueForTableColumnByItem(outlineView IOutlineView, object objc.IObject, tableColumn ITableColumn, item objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:setObjectValue:forTableColumn:byItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(object), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
}

func (o_ OutlineViewDataSourceWrapper) ImplementsOutlineViewSortDescriptorsDidChange() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:sortDescriptorsDidChange:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineViewSortDescriptorsDidChange(outlineView IOutlineView, oldDescriptors []foundation.ISortDescriptor) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:sortDescriptorsDidChange:"), objc.ExtractPtr(outlineView), oldDescriptors)
}

func (o_ OutlineViewDataSourceWrapper) ImplementsOutlineViewUpdateDraggingItemsForDrag() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:updateDraggingItemsForDrag:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineViewUpdateDraggingItemsForDrag(outlineView IOutlineView, draggingInfo IDraggingInfo) {
	po := objc.WrapAsProtocol("NSDraggingInfo", draggingInfo)
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:updateDraggingItemsForDrag:"), objc.ExtractPtr(outlineView), po)
}

func (o_ OutlineViewDataSourceWrapper) ImplementsOutlineViewValidateDropProposedItemProposedChildIndex() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:validateDrop:proposedItem:proposedChildIndex:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineViewValidateDropProposedItemProposedChildIndex(outlineView IOutlineView, info IDraggingInfo, item objc.IObject, index int) DragOperation {
	po := objc.WrapAsProtocol("NSDraggingInfo", info)
	rv := objc.CallMethod[DragOperation](o_, objc.GetSelector("outlineView:validateDrop:proposedItem:proposedChildIndex:"), objc.ExtractPtr(outlineView), po, objc.ExtractPtr(item), index)
	return rv
}

func (o_ OutlineViewDataSourceWrapper) ImplementsOutlineViewWriteItemsToPasteboard() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:writeItems:toPasteboard:"))
}
