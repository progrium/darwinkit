// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type ICollectionViewDelegate interface {
	ImplementsCollectionViewShouldSelectItemsAtIndexPaths() bool
	// optional
	CollectionViewShouldSelectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet
	ImplementsCollectionViewDidSelectItemsAtIndexPaths() bool
	// optional
	CollectionViewDidSelectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set)
	ImplementsCollectionViewShouldDeselectItemsAtIndexPaths() bool
	// optional
	CollectionViewShouldDeselectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet
	ImplementsCollectionViewDidDeselectItemsAtIndexPaths() bool
	// optional
	CollectionViewDidDeselectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set)
	ImplementsCollectionViewShouldChangeItemsAtIndexPathsToHighlightState() bool
	// optional
	CollectionViewShouldChangeItemsAtIndexPathsToHighlightState(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) foundation.ISet
	ImplementsCollectionViewDidChangeItemsAtIndexPathsToHighlightState() bool
	// optional
	CollectionViewDidChangeItemsAtIndexPathsToHighlightState(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState)
	ImplementsCollectionViewWillDisplayItemForRepresentedObjectAtIndexPath() bool
	// optional
	CollectionViewWillDisplayItemForRepresentedObjectAtIndexPath(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)
	ImplementsCollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath() bool
	// optional
	CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)
	ImplementsCollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath() bool
	// optional
	CollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)
	ImplementsCollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath() bool
	// optional
	CollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)
	ImplementsCollectionViewTransitionLayoutForOldLayoutNewLayout() bool
	// optional
	CollectionViewTransitionLayoutForOldLayoutNewLayout(collectionView CollectionView, fromLayout CollectionViewLayout, toLayout CollectionViewLayout) ICollectionViewTransitionLayout
	ImplementsCollectionViewCanDragItemsAtIndexPathsWithEvent() bool
	// optional
	CollectionViewCanDragItemsAtIndexPathsWithEvent(collectionView CollectionView, indexPaths foundation.Set, event Event) bool
	ImplementsCollectionViewPasteboardWriterForItemAtIndexPath() bool
	// optional
	CollectionViewPasteboardWriterForItemAtIndexPath(collectionView CollectionView, indexPath foundation.IndexPath) IPasteboardWriting
	ImplementsCollectionViewWriteItemsAtIndexPathsToPasteboard() bool
	// optional
	// deprecated
	CollectionViewWriteItemsAtIndexPathsToPasteboard(collectionView CollectionView, indexPaths foundation.Set, pasteboard Pasteboard) bool
	ImplementsCollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths() bool
	// optional
	// deprecated
	CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths(collectionView CollectionView, dropURL foundation.URL, indexPaths foundation.Set) []string
	ImplementsCollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset() bool
	// optional
	CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset(collectionView CollectionView, indexPaths foundation.Set, event Event, dragImageOffset *foundation.Point) IImage
	ImplementsCollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths() bool
	// optional
	CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexPaths foundation.Set)
	ImplementsCollectionViewDraggingSessionEndedAtPointDragOperation() bool
	// optional
	CollectionViewDraggingSessionEndedAtPointDragOperation(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	ImplementsCollectionViewUpdateDraggingItemsForDrag() bool
	// optional
	CollectionViewUpdateDraggingItemsForDrag(collectionView CollectionView, draggingInfo DraggingInfoWrapper)
	ImplementsCollectionViewValidateDropProposedIndexPathDropOperation() bool
	// optional
	CollectionViewValidateDropProposedIndexPathDropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation
	ImplementsCollectionViewAcceptDropIndexPathDropOperation() bool
	// optional
	CollectionViewAcceptDropIndexPathDropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool
	ImplementsCollectionViewCanDragItemsAtIndexesWithEvent() bool
	// optional
	CollectionViewCanDragItemsAtIndexesWithEvent(collectionView CollectionView, indexes foundation.IndexSet, event Event) bool
	ImplementsCollectionViewPasteboardWriterForItemAtIndex() bool
	// optional
	CollectionViewPasteboardWriterForItemAtIndex(collectionView CollectionView, index uint) IPasteboardWriting
	ImplementsCollectionViewWriteItemsAtIndexesToPasteboard() bool
	// optional
	// deprecated
	CollectionViewWriteItemsAtIndexesToPasteboard(collectionView CollectionView, indexes foundation.IndexSet, pasteboard Pasteboard) bool
	ImplementsCollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes() bool
	// optional
	// deprecated
	CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes(collectionView CollectionView, dropURL foundation.URL, indexes foundation.IndexSet) []string
	ImplementsCollectionViewDraggingImageForItemsAtIndexesWithEventOffset() bool
	// optional
	CollectionViewDraggingImageForItemsAtIndexesWithEventOffset(collectionView CollectionView, indexes foundation.IndexSet, event Event, dragImageOffset *foundation.Point) IImage
	ImplementsCollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes() bool
	// optional
	CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexes foundation.IndexSet)
	ImplementsCollectionViewValidateDropProposedIndexDropOperation() bool
	// optional
	CollectionViewValidateDropProposedIndexDropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation
	ImplementsCollectionViewAcceptDropIndexDropOperation() bool
	// optional
	CollectionViewAcceptDropIndexDropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, index int, dropOperation CollectionViewDropOperation) bool
}

type CollectionViewDelegate struct {
	_CollectionViewShouldSelectItemsAtIndexPaths                                       func(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet
	_CollectionViewDidSelectItemsAtIndexPaths                                          func(collectionView CollectionView, indexPaths foundation.Set)
	_CollectionViewShouldDeselectItemsAtIndexPaths                                     func(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet
	_CollectionViewDidDeselectItemsAtIndexPaths                                        func(collectionView CollectionView, indexPaths foundation.Set)
	_CollectionViewShouldChangeItemsAtIndexPathsToHighlightState                       func(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) foundation.ISet
	_CollectionViewDidChangeItemsAtIndexPathsToHighlightState                          func(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState)
	_CollectionViewWillDisplayItemForRepresentedObjectAtIndexPath                      func(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)
	_CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath                 func(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)
	_CollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath               func(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)
	_CollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath        func(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)
	_CollectionViewTransitionLayoutForOldLayoutNewLayout                               func(collectionView CollectionView, fromLayout CollectionViewLayout, toLayout CollectionViewLayout) ICollectionViewTransitionLayout
	_CollectionViewCanDragItemsAtIndexPathsWithEvent                                   func(collectionView CollectionView, indexPaths foundation.Set, event Event) bool
	_CollectionViewPasteboardWriterForItemAtIndexPath                                  func(collectionView CollectionView, indexPath foundation.IndexPath) IPasteboardWriting
	_CollectionViewWriteItemsAtIndexPathsToPasteboard                                  func(collectionView CollectionView, indexPaths foundation.Set, pasteboard Pasteboard) bool
	_CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths func(collectionView CollectionView, dropURL foundation.URL, indexPaths foundation.Set) []string
	_CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset                    func(collectionView CollectionView, indexPaths foundation.Set, event Event, dragImageOffset *foundation.Point) IImage
	_CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths                 func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexPaths foundation.Set)
	_CollectionViewDraggingSessionEndedAtPointDragOperation                            func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	_CollectionViewUpdateDraggingItemsForDrag                                          func(collectionView CollectionView, draggingInfo DraggingInfoWrapper)
	_CollectionViewValidateDropProposedIndexPathDropOperation                          func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation
	_CollectionViewAcceptDropIndexPathDropOperation                                    func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool
	_CollectionViewCanDragItemsAtIndexesWithEvent                                      func(collectionView CollectionView, indexes foundation.IndexSet, event Event) bool
	_CollectionViewPasteboardWriterForItemAtIndex                                      func(collectionView CollectionView, index uint) IPasteboardWriting
	_CollectionViewWriteItemsAtIndexesToPasteboard                                     func(collectionView CollectionView, indexes foundation.IndexSet, pasteboard Pasteboard) bool
	_CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes    func(collectionView CollectionView, dropURL foundation.URL, indexes foundation.IndexSet) []string
	_CollectionViewDraggingImageForItemsAtIndexesWithEventOffset                       func(collectionView CollectionView, indexes foundation.IndexSet, event Event, dragImageOffset *foundation.Point) IImage
	_CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes                    func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexes foundation.IndexSet)
	_CollectionViewValidateDropProposedIndexDropOperation                              func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation
	_CollectionViewAcceptDropIndexDropOperation                                        func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, index int, dropOperation CollectionViewDropOperation) bool
}

func (di *CollectionViewDelegate) ImplementsCollectionViewShouldSelectItemsAtIndexPaths() bool {
	return di._CollectionViewShouldSelectItemsAtIndexPaths != nil
}

func (di *CollectionViewDelegate) SetCollectionViewShouldSelectItemsAtIndexPaths(f func(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet) {
	di._CollectionViewShouldSelectItemsAtIndexPaths = f
}

func (di *CollectionViewDelegate) CollectionViewShouldSelectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet {
	return di._CollectionViewShouldSelectItemsAtIndexPaths(collectionView, indexPaths)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewDidSelectItemsAtIndexPaths() bool {
	return di._CollectionViewDidSelectItemsAtIndexPaths != nil
}

func (di *CollectionViewDelegate) SetCollectionViewDidSelectItemsAtIndexPaths(f func(collectionView CollectionView, indexPaths foundation.Set)) {
	di._CollectionViewDidSelectItemsAtIndexPaths = f
}

func (di *CollectionViewDelegate) CollectionViewDidSelectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) {
	di._CollectionViewDidSelectItemsAtIndexPaths(collectionView, indexPaths)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewShouldDeselectItemsAtIndexPaths() bool {
	return di._CollectionViewShouldDeselectItemsAtIndexPaths != nil
}

func (di *CollectionViewDelegate) SetCollectionViewShouldDeselectItemsAtIndexPaths(f func(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet) {
	di._CollectionViewShouldDeselectItemsAtIndexPaths = f
}

func (di *CollectionViewDelegate) CollectionViewShouldDeselectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet {
	return di._CollectionViewShouldDeselectItemsAtIndexPaths(collectionView, indexPaths)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewDidDeselectItemsAtIndexPaths() bool {
	return di._CollectionViewDidDeselectItemsAtIndexPaths != nil
}

func (di *CollectionViewDelegate) SetCollectionViewDidDeselectItemsAtIndexPaths(f func(collectionView CollectionView, indexPaths foundation.Set)) {
	di._CollectionViewDidDeselectItemsAtIndexPaths = f
}

func (di *CollectionViewDelegate) CollectionViewDidDeselectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) {
	di._CollectionViewDidDeselectItemsAtIndexPaths(collectionView, indexPaths)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewShouldChangeItemsAtIndexPathsToHighlightState() bool {
	return di._CollectionViewShouldChangeItemsAtIndexPathsToHighlightState != nil
}

func (di *CollectionViewDelegate) SetCollectionViewShouldChangeItemsAtIndexPathsToHighlightState(f func(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) foundation.ISet) {
	di._CollectionViewShouldChangeItemsAtIndexPathsToHighlightState = f
}

func (di *CollectionViewDelegate) CollectionViewShouldChangeItemsAtIndexPathsToHighlightState(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) foundation.ISet {
	return di._CollectionViewShouldChangeItemsAtIndexPathsToHighlightState(collectionView, indexPaths, highlightState)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewDidChangeItemsAtIndexPathsToHighlightState() bool {
	return di._CollectionViewDidChangeItemsAtIndexPathsToHighlightState != nil
}

func (di *CollectionViewDelegate) SetCollectionViewDidChangeItemsAtIndexPathsToHighlightState(f func(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState)) {
	di._CollectionViewDidChangeItemsAtIndexPathsToHighlightState = f
}

func (di *CollectionViewDelegate) CollectionViewDidChangeItemsAtIndexPathsToHighlightState(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) {
	di._CollectionViewDidChangeItemsAtIndexPathsToHighlightState(collectionView, indexPaths, highlightState)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewWillDisplayItemForRepresentedObjectAtIndexPath() bool {
	return di._CollectionViewWillDisplayItemForRepresentedObjectAtIndexPath != nil
}

func (di *CollectionViewDelegate) SetCollectionViewWillDisplayItemForRepresentedObjectAtIndexPath(f func(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)) {
	di._CollectionViewWillDisplayItemForRepresentedObjectAtIndexPath = f
}

func (di *CollectionViewDelegate) CollectionViewWillDisplayItemForRepresentedObjectAtIndexPath(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath) {
	di._CollectionViewWillDisplayItemForRepresentedObjectAtIndexPath(collectionView, item, indexPath)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath() bool {
	return di._CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath != nil
}

func (di *CollectionViewDelegate) SetCollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath(f func(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)) {
	di._CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath = f
}

func (di *CollectionViewDelegate) CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath) {
	di._CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath(collectionView, item, indexPath)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath() bool {
	return di._CollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath != nil
}

func (di *CollectionViewDelegate) SetCollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath(f func(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)) {
	di._CollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath = f
}

func (di *CollectionViewDelegate) CollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) {
	di._CollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath(collectionView, view, elementKind, indexPath)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath() bool {
	return di._CollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath != nil
}

func (di *CollectionViewDelegate) SetCollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath(f func(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)) {
	di._CollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath = f
}

func (di *CollectionViewDelegate) CollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) {
	di._CollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath(collectionView, view, elementKind, indexPath)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewTransitionLayoutForOldLayoutNewLayout() bool {
	return di._CollectionViewTransitionLayoutForOldLayoutNewLayout != nil
}

func (di *CollectionViewDelegate) SetCollectionViewTransitionLayoutForOldLayoutNewLayout(f func(collectionView CollectionView, fromLayout CollectionViewLayout, toLayout CollectionViewLayout) ICollectionViewTransitionLayout) {
	di._CollectionViewTransitionLayoutForOldLayoutNewLayout = f
}

func (di *CollectionViewDelegate) CollectionViewTransitionLayoutForOldLayoutNewLayout(collectionView CollectionView, fromLayout CollectionViewLayout, toLayout CollectionViewLayout) ICollectionViewTransitionLayout {
	return di._CollectionViewTransitionLayoutForOldLayoutNewLayout(collectionView, fromLayout, toLayout)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewCanDragItemsAtIndexPathsWithEvent() bool {
	return di._CollectionViewCanDragItemsAtIndexPathsWithEvent != nil
}

func (di *CollectionViewDelegate) SetCollectionViewCanDragItemsAtIndexPathsWithEvent(f func(collectionView CollectionView, indexPaths foundation.Set, event Event) bool) {
	di._CollectionViewCanDragItemsAtIndexPathsWithEvent = f
}

func (di *CollectionViewDelegate) CollectionViewCanDragItemsAtIndexPathsWithEvent(collectionView CollectionView, indexPaths foundation.Set, event Event) bool {
	return di._CollectionViewCanDragItemsAtIndexPathsWithEvent(collectionView, indexPaths, event)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewPasteboardWriterForItemAtIndexPath() bool {
	return di._CollectionViewPasteboardWriterForItemAtIndexPath != nil
}

func (di *CollectionViewDelegate) SetCollectionViewPasteboardWriterForItemAtIndexPath(f func(collectionView CollectionView, indexPath foundation.IndexPath) IPasteboardWriting) {
	di._CollectionViewPasteboardWriterForItemAtIndexPath = f
}

func (di *CollectionViewDelegate) CollectionViewPasteboardWriterForItemAtIndexPath(collectionView CollectionView, indexPath foundation.IndexPath) IPasteboardWriting {
	return di._CollectionViewPasteboardWriterForItemAtIndexPath(collectionView, indexPath)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewWriteItemsAtIndexPathsToPasteboard() bool {
	return di._CollectionViewWriteItemsAtIndexPathsToPasteboard != nil
}

// deprecated
func (di *CollectionViewDelegate) SetCollectionViewWriteItemsAtIndexPathsToPasteboard(f func(collectionView CollectionView, indexPaths foundation.Set, pasteboard Pasteboard) bool) {
	di._CollectionViewWriteItemsAtIndexPathsToPasteboard = f
}

// deprecated
func (di *CollectionViewDelegate) CollectionViewWriteItemsAtIndexPathsToPasteboard(collectionView CollectionView, indexPaths foundation.Set, pasteboard Pasteboard) bool {
	return di._CollectionViewWriteItemsAtIndexPathsToPasteboard(collectionView, indexPaths, pasteboard)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths() bool {
	return di._CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths != nil
}

// deprecated
func (di *CollectionViewDelegate) SetCollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths(f func(collectionView CollectionView, dropURL foundation.URL, indexPaths foundation.Set) []string) {
	di._CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths = f
}

// deprecated
func (di *CollectionViewDelegate) CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths(collectionView CollectionView, dropURL foundation.URL, indexPaths foundation.Set) []string {
	return di._CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths(collectionView, dropURL, indexPaths)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset() bool {
	return di._CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset != nil
}

func (di *CollectionViewDelegate) SetCollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset(f func(collectionView CollectionView, indexPaths foundation.Set, event Event, dragImageOffset *foundation.Point) IImage) {
	di._CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset = f
}

func (di *CollectionViewDelegate) CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset(collectionView CollectionView, indexPaths foundation.Set, event Event, dragImageOffset *foundation.Point) IImage {
	return di._CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset(collectionView, indexPaths, event, dragImageOffset)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths() bool {
	return di._CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths != nil
}

func (di *CollectionViewDelegate) SetCollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths(f func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexPaths foundation.Set)) {
	di._CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths = f
}

func (di *CollectionViewDelegate) CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexPaths foundation.Set) {
	di._CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths(collectionView, session, screenPoint, indexPaths)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewDraggingSessionEndedAtPointDragOperation() bool {
	return di._CollectionViewDraggingSessionEndedAtPointDragOperation != nil
}

func (di *CollectionViewDelegate) SetCollectionViewDraggingSessionEndedAtPointDragOperation(f func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)) {
	di._CollectionViewDraggingSessionEndedAtPointDragOperation = f
}

func (di *CollectionViewDelegate) CollectionViewDraggingSessionEndedAtPointDragOperation(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, operation DragOperation) {
	di._CollectionViewDraggingSessionEndedAtPointDragOperation(collectionView, session, screenPoint, operation)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewUpdateDraggingItemsForDrag() bool {
	return di._CollectionViewUpdateDraggingItemsForDrag != nil
}

func (di *CollectionViewDelegate) SetCollectionViewUpdateDraggingItemsForDrag(f func(collectionView CollectionView, draggingInfo DraggingInfoWrapper)) {
	di._CollectionViewUpdateDraggingItemsForDrag = f
}

func (di *CollectionViewDelegate) CollectionViewUpdateDraggingItemsForDrag(collectionView CollectionView, draggingInfo DraggingInfoWrapper) {
	di._CollectionViewUpdateDraggingItemsForDrag(collectionView, draggingInfo)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewValidateDropProposedIndexPathDropOperation() bool {
	return di._CollectionViewValidateDropProposedIndexPathDropOperation != nil
}

func (di *CollectionViewDelegate) SetCollectionViewValidateDropProposedIndexPathDropOperation(f func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation) {
	di._CollectionViewValidateDropProposedIndexPathDropOperation = f
}

func (di *CollectionViewDelegate) CollectionViewValidateDropProposedIndexPathDropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	return di._CollectionViewValidateDropProposedIndexPathDropOperation(collectionView, draggingInfo, proposedDropIndexPath, proposedDropOperation)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewAcceptDropIndexPathDropOperation() bool {
	return di._CollectionViewAcceptDropIndexPathDropOperation != nil
}

func (di *CollectionViewDelegate) SetCollectionViewAcceptDropIndexPathDropOperation(f func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool) {
	di._CollectionViewAcceptDropIndexPathDropOperation = f
}

func (di *CollectionViewDelegate) CollectionViewAcceptDropIndexPathDropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool {
	return di._CollectionViewAcceptDropIndexPathDropOperation(collectionView, draggingInfo, indexPath, dropOperation)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewCanDragItemsAtIndexesWithEvent() bool {
	return di._CollectionViewCanDragItemsAtIndexesWithEvent != nil
}

func (di *CollectionViewDelegate) SetCollectionViewCanDragItemsAtIndexesWithEvent(f func(collectionView CollectionView, indexes foundation.IndexSet, event Event) bool) {
	di._CollectionViewCanDragItemsAtIndexesWithEvent = f
}

func (di *CollectionViewDelegate) CollectionViewCanDragItemsAtIndexesWithEvent(collectionView CollectionView, indexes foundation.IndexSet, event Event) bool {
	return di._CollectionViewCanDragItemsAtIndexesWithEvent(collectionView, indexes, event)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewPasteboardWriterForItemAtIndex() bool {
	return di._CollectionViewPasteboardWriterForItemAtIndex != nil
}

func (di *CollectionViewDelegate) SetCollectionViewPasteboardWriterForItemAtIndex(f func(collectionView CollectionView, index uint) IPasteboardWriting) {
	di._CollectionViewPasteboardWriterForItemAtIndex = f
}

func (di *CollectionViewDelegate) CollectionViewPasteboardWriterForItemAtIndex(collectionView CollectionView, index uint) IPasteboardWriting {
	return di._CollectionViewPasteboardWriterForItemAtIndex(collectionView, index)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewWriteItemsAtIndexesToPasteboard() bool {
	return di._CollectionViewWriteItemsAtIndexesToPasteboard != nil
}

// deprecated
func (di *CollectionViewDelegate) SetCollectionViewWriteItemsAtIndexesToPasteboard(f func(collectionView CollectionView, indexes foundation.IndexSet, pasteboard Pasteboard) bool) {
	di._CollectionViewWriteItemsAtIndexesToPasteboard = f
}

// deprecated
func (di *CollectionViewDelegate) CollectionViewWriteItemsAtIndexesToPasteboard(collectionView CollectionView, indexes foundation.IndexSet, pasteboard Pasteboard) bool {
	return di._CollectionViewWriteItemsAtIndexesToPasteboard(collectionView, indexes, pasteboard)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes() bool {
	return di._CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes != nil
}

// deprecated
func (di *CollectionViewDelegate) SetCollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes(f func(collectionView CollectionView, dropURL foundation.URL, indexes foundation.IndexSet) []string) {
	di._CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes = f
}

// deprecated
func (di *CollectionViewDelegate) CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes(collectionView CollectionView, dropURL foundation.URL, indexes foundation.IndexSet) []string {
	return di._CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes(collectionView, dropURL, indexes)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewDraggingImageForItemsAtIndexesWithEventOffset() bool {
	return di._CollectionViewDraggingImageForItemsAtIndexesWithEventOffset != nil
}

func (di *CollectionViewDelegate) SetCollectionViewDraggingImageForItemsAtIndexesWithEventOffset(f func(collectionView CollectionView, indexes foundation.IndexSet, event Event, dragImageOffset *foundation.Point) IImage) {
	di._CollectionViewDraggingImageForItemsAtIndexesWithEventOffset = f
}

func (di *CollectionViewDelegate) CollectionViewDraggingImageForItemsAtIndexesWithEventOffset(collectionView CollectionView, indexes foundation.IndexSet, event Event, dragImageOffset *foundation.Point) IImage {
	return di._CollectionViewDraggingImageForItemsAtIndexesWithEventOffset(collectionView, indexes, event, dragImageOffset)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes() bool {
	return di._CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes != nil
}

func (di *CollectionViewDelegate) SetCollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes(f func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexes foundation.IndexSet)) {
	di._CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes = f
}

func (di *CollectionViewDelegate) CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexes foundation.IndexSet) {
	di._CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes(collectionView, session, screenPoint, indexes)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewValidateDropProposedIndexDropOperation() bool {
	return di._CollectionViewValidateDropProposedIndexDropOperation != nil
}

func (di *CollectionViewDelegate) SetCollectionViewValidateDropProposedIndexDropOperation(f func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation) {
	di._CollectionViewValidateDropProposedIndexDropOperation = f
}

func (di *CollectionViewDelegate) CollectionViewValidateDropProposedIndexDropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	return di._CollectionViewValidateDropProposedIndexDropOperation(collectionView, draggingInfo, proposedDropIndex, proposedDropOperation)
}
func (di *CollectionViewDelegate) ImplementsCollectionViewAcceptDropIndexDropOperation() bool {
	return di._CollectionViewAcceptDropIndexDropOperation != nil
}

func (di *CollectionViewDelegate) SetCollectionViewAcceptDropIndexDropOperation(f func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, index int, dropOperation CollectionViewDropOperation) bool) {
	di._CollectionViewAcceptDropIndexDropOperation = f
}

func (di *CollectionViewDelegate) CollectionViewAcceptDropIndexDropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, index int, dropOperation CollectionViewDropOperation) bool {
	return di._CollectionViewAcceptDropIndexDropOperation(collectionView, draggingInfo, index, dropOperation)
}

type CollectionViewDelegateWrapper struct {
	objc.Object
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewShouldSelectItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:shouldSelectItemsAtIndexPaths:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewShouldSelectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, objc.GetSelector("collectionView:shouldSelectItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths))
	return rv
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewDidSelectItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:didSelectItemsAtIndexPaths:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewDidSelectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:didSelectItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths))
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewShouldDeselectItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:shouldDeselectItemsAtIndexPaths:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewShouldDeselectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, objc.GetSelector("collectionView:shouldDeselectItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths))
	return rv
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewDidDeselectItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:didDeselectItemsAtIndexPaths:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewDidDeselectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:didDeselectItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths))
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewShouldChangeItemsAtIndexPathsToHighlightState() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:shouldChangeItemsAtIndexPaths:toHighlightState:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewShouldChangeItemsAtIndexPathsToHighlightState(collectionView ICollectionView, indexPaths foundation.ISet, highlightState CollectionViewItemHighlightState) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, objc.GetSelector("collectionView:shouldChangeItemsAtIndexPaths:toHighlightState:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths), highlightState)
	return rv
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewDidChangeItemsAtIndexPathsToHighlightState() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:didChangeItemsAtIndexPaths:toHighlightState:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewDidChangeItemsAtIndexPathsToHighlightState(collectionView ICollectionView, indexPaths foundation.ISet, highlightState CollectionViewItemHighlightState) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:didChangeItemsAtIndexPaths:toHighlightState:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths), highlightState)
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewWillDisplayItemForRepresentedObjectAtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:willDisplayItem:forRepresentedObjectAtIndexPath:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewWillDisplayItemForRepresentedObjectAtIndexPath(collectionView ICollectionView, item ICollectionViewItem, indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:willDisplayItem:forRepresentedObjectAtIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(item), objc.ExtractPtr(indexPath))
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:didEndDisplayingItem:forRepresentedObjectAtIndexPath:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath(collectionView ICollectionView, item ICollectionViewItem, indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:didEndDisplayingItem:forRepresentedObjectAtIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(item), objc.ExtractPtr(indexPath))
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:willDisplaySupplementaryView:forElementKind:atIndexPath:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath(collectionView ICollectionView, view IView, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:willDisplaySupplementaryView:forElementKind:atIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(view), elementKind, objc.ExtractPtr(indexPath))
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:didEndDisplayingSupplementaryView:forElementOfKind:atIndexPath:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath(collectionView ICollectionView, view IView, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:didEndDisplayingSupplementaryView:forElementOfKind:atIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(view), elementKind, objc.ExtractPtr(indexPath))
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewTransitionLayoutForOldLayoutNewLayout() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:transitionLayoutForOldLayout:newLayout:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewTransitionLayoutForOldLayoutNewLayout(collectionView ICollectionView, fromLayout ICollectionViewLayout, toLayout ICollectionViewLayout) CollectionViewTransitionLayout {
	rv := objc.CallMethod[CollectionViewTransitionLayout](c_, objc.GetSelector("collectionView:transitionLayoutForOldLayout:newLayout:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(fromLayout), objc.ExtractPtr(toLayout))
	return rv
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewCanDragItemsAtIndexPathsWithEvent() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:canDragItemsAtIndexPaths:withEvent:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewCanDragItemsAtIndexPathsWithEvent(collectionView ICollectionView, indexPaths foundation.ISet, event IEvent) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("collectionView:canDragItemsAtIndexPaths:withEvent:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths), objc.ExtractPtr(event))
	return rv
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewPasteboardWriterForItemAtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:pasteboardWriterForItemAtIndexPath:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewPasteboardWriterForItemAtIndexPath(collectionView ICollectionView, indexPath foundation.IIndexPath) PasteboardWritingWrapper {
	rv := objc.CallMethod[PasteboardWritingWrapper](c_, objc.GetSelector("collectionView:pasteboardWriterForItemAtIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPath))
	return rv
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewWriteItemsAtIndexPathsToPasteboard() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:writeItemsAtIndexPaths:toPasteboard:"))
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:namesOfPromisedFilesDroppedAtDestination:forDraggedItemsAtIndexPaths:"))
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:draggingImageForItemsAtIndexPaths:withEvent:offset:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset(collectionView ICollectionView, indexPaths foundation.ISet, event IEvent, dragImageOffset *foundation.Point) Image {
	rv := objc.CallMethod[Image](c_, objc.GetSelector("collectionView:draggingImageForItemsAtIndexPaths:withEvent:offset:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths), objc.ExtractPtr(event), dragImageOffset)
	return rv
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexPaths:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths(collectionView ICollectionView, session IDraggingSession, screenPoint foundation.Point, indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(session), screenPoint, objc.ExtractPtr(indexPaths))
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewDraggingSessionEndedAtPointDragOperation() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:draggingSession:endedAtPoint:dragOperation:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewDraggingSessionEndedAtPointDragOperation(collectionView ICollectionView, session IDraggingSession, screenPoint foundation.Point, operation DragOperation) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:draggingSession:endedAtPoint:dragOperation:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(session), screenPoint, operation)
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewUpdateDraggingItemsForDrag() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:updateDraggingItemsForDrag:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewUpdateDraggingItemsForDrag(collectionView ICollectionView, draggingInfo IDraggingInfo) {
	po := objc.WrapAsProtocol("NSDraggingInfo", draggingInfo)
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:updateDraggingItemsForDrag:"), objc.ExtractPtr(collectionView), po)
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewValidateDropProposedIndexPathDropOperation() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:validateDrop:proposedIndexPath:dropOperation:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewValidateDropProposedIndexPathDropOperation(collectionView ICollectionView, draggingInfo IDraggingInfo, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	po := objc.WrapAsProtocol("NSDraggingInfo", draggingInfo)
	rv := objc.CallMethod[DragOperation](c_, objc.GetSelector("collectionView:validateDrop:proposedIndexPath:dropOperation:"), objc.ExtractPtr(collectionView), po, unsafe.Pointer(proposedDropIndexPath), proposedDropOperation)
	return rv
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewAcceptDropIndexPathDropOperation() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:acceptDrop:indexPath:dropOperation:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewAcceptDropIndexPathDropOperation(collectionView ICollectionView, draggingInfo IDraggingInfo, indexPath foundation.IIndexPath, dropOperation CollectionViewDropOperation) bool {
	po := objc.WrapAsProtocol("NSDraggingInfo", draggingInfo)
	rv := objc.CallMethod[bool](c_, objc.GetSelector("collectionView:acceptDrop:indexPath:dropOperation:"), objc.ExtractPtr(collectionView), po, objc.ExtractPtr(indexPath), dropOperation)
	return rv
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewCanDragItemsAtIndexesWithEvent() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:canDragItemsAtIndexes:withEvent:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewCanDragItemsAtIndexesWithEvent(collectionView ICollectionView, indexes foundation.IIndexSet, event IEvent) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("collectionView:canDragItemsAtIndexes:withEvent:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexes), objc.ExtractPtr(event))
	return rv
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewPasteboardWriterForItemAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:pasteboardWriterForItemAtIndex:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewPasteboardWriterForItemAtIndex(collectionView ICollectionView, index uint) PasteboardWritingWrapper {
	rv := objc.CallMethod[PasteboardWritingWrapper](c_, objc.GetSelector("collectionView:pasteboardWriterForItemAtIndex:"), objc.ExtractPtr(collectionView), index)
	return rv
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewWriteItemsAtIndexesToPasteboard() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:writeItemsAtIndexes:toPasteboard:"))
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:namesOfPromisedFilesDroppedAtDestination:forDraggedItemsAtIndexes:"))
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewDraggingImageForItemsAtIndexesWithEventOffset() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:draggingImageForItemsAtIndexes:withEvent:offset:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewDraggingImageForItemsAtIndexesWithEventOffset(collectionView ICollectionView, indexes foundation.IIndexSet, event IEvent, dragImageOffset *foundation.Point) Image {
	rv := objc.CallMethod[Image](c_, objc.GetSelector("collectionView:draggingImageForItemsAtIndexes:withEvent:offset:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexes), objc.ExtractPtr(event), dragImageOffset)
	return rv
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexes:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes(collectionView ICollectionView, session IDraggingSession, screenPoint foundation.Point, indexes foundation.IIndexSet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexes:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(session), screenPoint, objc.ExtractPtr(indexes))
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewValidateDropProposedIndexDropOperation() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:validateDrop:proposedIndex:dropOperation:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewValidateDropProposedIndexDropOperation(collectionView ICollectionView, draggingInfo IDraggingInfo, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	po := objc.WrapAsProtocol("NSDraggingInfo", draggingInfo)
	rv := objc.CallMethod[DragOperation](c_, objc.GetSelector("collectionView:validateDrop:proposedIndex:dropOperation:"), objc.ExtractPtr(collectionView), po, proposedDropIndex, proposedDropOperation)
	return rv
}

func (c_ CollectionViewDelegateWrapper) ImplementsCollectionViewAcceptDropIndexDropOperation() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:acceptDrop:index:dropOperation:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionViewAcceptDropIndexDropOperation(collectionView ICollectionView, draggingInfo IDraggingInfo, index int, dropOperation CollectionViewDropOperation) bool {
	po := objc.WrapAsProtocol("NSDraggingInfo", draggingInfo)
	rv := objc.CallMethod[bool](c_, objc.GetSelector("collectionView:acceptDrop:index:dropOperation:"), objc.ExtractPtr(collectionView), po, index, dropOperation)
	return rv
}
