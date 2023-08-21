// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionView] class.
var CollectionViewClass = _CollectionViewClass{objc.GetClass("NSCollectionView")}

type _CollectionViewClass struct {
	objc.Class
}

// An interface definition for the [CollectionView] class.
type ICollectionView interface {
	IView
	ItemAtIndex(index uint) CollectionViewItem
	RegisterNibForItemWithIdentifier(nib INib, identifier UserInterfaceItemIdentifier)
	IndexPathsForVisibleSupplementaryElementsOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set
	InsertSections(sections foundation.IIndexSet)
	MakeItemWithIdentifierForIndexPath(identifier UserInterfaceItemIdentifier, indexPath foundation.IIndexPath) CollectionViewItem
	NumberOfItemsInSection(section int) int
	DeselectAll(sender objc.IObject) objc.Object
	ReloadItemsAtIndexPaths(indexPaths foundation.ISet)
	FrameForItemAtIndexWithNumberOfItems(index uint, numberOfItems uint) foundation.Rect
	RegisterClassForItemWithIdentifier(itemClass objc.IClass, identifier UserInterfaceItemIdentifier)
	IndexPathForItemAtPoint(point foundation.Point) foundation.IndexPath
	DraggingImageForItemsAtIndexesWithEventOffset(indexes foundation.IIndexSet, event IEvent, dragImageOffset foundation.PointPointer) Image
	SelectAll(sender objc.IObject) objc.Object
	DeleteSections(sections foundation.IIndexSet)
	SelectItemsAtIndexPathsScrollPosition(indexPaths foundation.ISet, scrollPosition CollectionViewScrollPosition)
	ScrollToItemsAtIndexPathsScrollPosition(indexPaths foundation.ISet, scrollPosition CollectionViewScrollPosition)
	PerformBatchUpdatesCompletionHandler(updates func(), completionHandler func(finished bool))
	IndexPathForItem(item ICollectionViewItem) foundation.IndexPath
	DraggingImageForItemsAtIndexPathsWithEventOffset(indexPaths foundation.ISet, event IEvent, dragImageOffset foundation.PointPointer) Image
	MakeSupplementaryViewOfKindWithIdentifierForIndexPath(elementKind CollectionViewSupplementaryElementKind, identifier UserInterfaceItemIdentifier, indexPath foundation.IIndexPath) View
	IndexPathsForVisibleItems() foundation.Set
	DeselectItemsAtIndexPaths(indexPaths foundation.ISet)
	SetDraggingSourceOperationMaskForLocal(dragOperationMask DragOperation, localDestination bool)
	SupplementaryViewForElementKindAtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) View
	InsertItemsAtIndexPaths(indexPaths foundation.ISet)
	VisibleItems() []CollectionViewItem
	ToggleSectionCollapse(sender objc.IObject) objc.Object
	MoveSectionToSection(section int, newSection int)
	LayoutAttributesForItemAtIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	LayoutAttributesForSupplementaryElementOfKindAtIndexPath(kind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	ItemAtIndexPath(indexPath foundation.IIndexPath) CollectionViewItem
	ReloadData()
	ReloadSections(sections foundation.IIndexSet)
	DeleteItemsAtIndexPaths(indexPaths foundation.ISet)
	MoveItemAtIndexPathToIndexPath(indexPath foundation.IIndexPath, newIndexPath foundation.IIndexPath)
	VisibleSupplementaryViewsOfKind(elementKind CollectionViewSupplementaryElementKind) []View
	DataSource() CollectionViewDataSourceWrapper
	SetDataSource(value PCollectionViewDataSource)
	SetDataSourceObject(valueObject objc.IObject)
	Content() []objc.Object
	SetContent(value []objc.IObject)
	BackgroundColors() []Color
	SetBackgroundColors(value []IColor)
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	BackgroundView() View
	SetBackgroundView(value IView)
	IsFirstResponder() bool
	CollectionViewLayout() CollectionViewLayout
	SetCollectionViewLayout(value ICollectionViewLayout)
	Delegate() CollectionViewDelegateWrapper
	SetDelegate(value PCollectionViewDelegate)
	SetDelegateObject(valueObject objc.IObject)
	SelectionIndexPaths() foundation.Set
	SetSelectionIndexPaths(value foundation.ISet)
	BackgroundViewScrollsWithContent() bool
	SetBackgroundViewScrollsWithContent(value bool)
	IsSelectable() bool
	SetSelectable(value bool)
	SelectionIndexes() foundation.IndexSet
	SetSelectionIndexes(value foundation.IIndexSet)
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	NumberOfSections() int
	PrefetchDataSource() CollectionViewPrefetchingWrapper
	SetPrefetchDataSource(value PCollectionViewPrefetching)
	SetPrefetchDataSourceObject(valueObject objc.IObject)
}

// An ordered collection of data items displayed in a customizable layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview?language=objc
type CollectionView struct {
	View
}

func CollectionViewFrom(ptr unsafe.Pointer) CollectionView {
	return CollectionView{
		View: ViewFrom(ptr),
	}
}

func (cc _CollectionViewClass) Alloc() CollectionView {
	rv := objc.Call[CollectionView](cc, objc.Sel("alloc"))
	return rv
}

func CollectionView_Alloc() CollectionView {
	return CollectionViewClass.Alloc()
}

func (cc _CollectionViewClass) New() CollectionView {
	rv := objc.Call[CollectionView](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionView() CollectionView {
	return CollectionViewClass.New()
}

func (c_ CollectionView) Init() CollectionView {
	rv := objc.Call[CollectionView](c_, objc.Sel("init"))
	return rv
}

func (c_ CollectionView) InitWithFrame(frameRect foundation.Rect) CollectionView {
	rv := objc.Call[CollectionView](c_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewCollectionViewWithFrame(frameRect foundation.Rect) CollectionView {
	instance := CollectionViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Returns the collection view item for the represented object at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1526399-itematindex?language=objc
func (c_ CollectionView) ItemAtIndex(index uint) CollectionViewItem {
	rv := objc.Call[CollectionViewItem](c_, objc.Sel("itemAtIndex:"), index)
	return rv
}

// Registers a nib file to use when creating items in the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528174-registernib?language=objc
func (c_ CollectionView) RegisterNibForItemWithIdentifier(nib INib, identifier UserInterfaceItemIdentifier) {
	objc.Call[objc.Void](c_, objc.Sel("registerNib:forItemWithIdentifier:"), objc.Ptr(nib), identifier)
}

// Returns the index paths of the currently active supplementary views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528287-indexpathsforvisiblesupplementar?language=objc
func (c_ CollectionView) IndexPathsForVisibleSupplementaryElementsOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set {
	rv := objc.Call[foundation.Set](c_, objc.Sel("indexPathsForVisibleSupplementaryElementsOfKind:"), elementKind)
	return rv
}

// Inserts new sections at the specified indexes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1526126-insertsections?language=objc
func (c_ CollectionView) InsertSections(sections foundation.IIndexSet) {
	objc.Call[objc.Void](c_, objc.Sel("insertSections:"), objc.Ptr(sections))
}

// Creates or returns a reusable item object of the specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528240-makeitemwithidentifier?language=objc
func (c_ CollectionView) MakeItemWithIdentifierForIndexPath(identifier UserInterfaceItemIdentifier, indexPath foundation.IIndexPath) CollectionViewItem {
	rv := objc.Call[CollectionViewItem](c_, objc.Sel("makeItemWithIdentifier:forIndexPath:"), identifier, objc.Ptr(indexPath))
	return rv
}

// Returns the number of items in the specified section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528291-numberofitemsinsection?language=objc
func (c_ CollectionView) NumberOfItemsInSection(section int) int {
	rv := objc.Call[int](c_, objc.Sel("numberOfItemsInSection:"), section)
	return rv
}

// Deselects all items in the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528251-deselectall?language=objc
func (c_ CollectionView) DeselectAll(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("deselectAll:"), sender)
	return rv
}

// Reloads only the specified items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528166-reloaditemsatindexpaths?language=objc
func (c_ CollectionView) ReloadItemsAtIndexPaths(indexPaths foundation.ISet) {
	objc.Call[objc.Void](c_, objc.Sel("reloadItemsAtIndexPaths:"), objc.Ptr(indexPaths))
}

// Returns the frame of an item based on the number of items in the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528209-frameforitematindex?language=objc
func (c_ CollectionView) FrameForItemAtIndexWithNumberOfItems(index uint, numberOfItems uint) foundation.Rect {
	rv := objc.Call[foundation.Rect](c_, objc.Sel("frameForItemAtIndex:withNumberOfItems:"), index, numberOfItems)
	return rv
}

// Registers a class to use when creating new items in the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528268-registerclass?language=objc
func (c_ CollectionView) RegisterClassForItemWithIdentifier(itemClass objc.IClass, identifier UserInterfaceItemIdentifier) {
	objc.Call[objc.Void](c_, objc.Sel("registerClass:forItemWithIdentifier:"), objc.Ptr(itemClass), identifier)
}

// Returns the index path of the item at the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1524499-indexpathforitematpoint?language=objc
func (c_ CollectionView) IndexPathForItemAtPoint(point foundation.Point) foundation.IndexPath {
	rv := objc.Call[foundation.IndexPath](c_, objc.Sel("indexPathForItemAtPoint:"), point)
	return rv
}

// This method computes and returns an image to use for dragging. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528249-draggingimageforitemsatindexes?language=objc
func (c_ CollectionView) DraggingImageForItemsAtIndexesWithEventOffset(indexes foundation.IIndexSet, event IEvent, dragImageOffset foundation.PointPointer) Image {
	rv := objc.Call[Image](c_, objc.Sel("draggingImageForItemsAtIndexes:withEvent:offset:"), objc.Ptr(indexes), objc.Ptr(event), dragImageOffset)
	return rv
}

// Selects all items in the collection view, if doing so is possible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528156-selectall?language=objc
func (c_ CollectionView) SelectAll(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("selectAll:"), sender)
	return rv
}

// Deletes the specified sections and their contained items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1527083-deletesections?language=objc
func (c_ CollectionView) DeleteSections(sections foundation.IIndexSet) {
	objc.Call[objc.Void](c_, objc.Sel("deleteSections:"), objc.Ptr(sections))
}

// Adds the specified items to the current selection and optionally scrolls the items into position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1525656-selectitemsatindexpaths?language=objc
func (c_ CollectionView) SelectItemsAtIndexPathsScrollPosition(indexPaths foundation.ISet, scrollPosition CollectionViewScrollPosition) {
	objc.Call[objc.Void](c_, objc.Sel("selectItemsAtIndexPaths:scrollPosition:"), objc.Ptr(indexPaths), scrollPosition)
}

// Scrolls the collection view contents until the specified items are visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528227-scrolltoitemsatindexpaths?language=objc
func (c_ CollectionView) ScrollToItemsAtIndexPathsScrollPosition(indexPaths foundation.ISet, scrollPosition CollectionViewScrollPosition) {
	objc.Call[objc.Void](c_, objc.Sel("scrollToItemsAtIndexPaths:scrollPosition:"), objc.Ptr(indexPaths), scrollPosition)
}

// Encapsulates multiple insert, delete, reload, and move operations into a single animated operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1525876-performbatchupdates?language=objc
func (c_ CollectionView) PerformBatchUpdatesCompletionHandler(updates func(), completionHandler func(finished bool)) {
	objc.Call[objc.Void](c_, objc.Sel("performBatchUpdates:completionHandler:"), updates, completionHandler)
}

// Returns the index path of the specified item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528277-indexpathforitem?language=objc
func (c_ CollectionView) IndexPathForItem(item ICollectionViewItem) foundation.IndexPath {
	rv := objc.Call[foundation.IndexPath](c_, objc.Sel("indexPathForItem:"), objc.Ptr(item))
	return rv
}

// Returns an image to use for dragging the specified items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528191-draggingimageforitemsatindexpath?language=objc
func (c_ CollectionView) DraggingImageForItemsAtIndexPathsWithEventOffset(indexPaths foundation.ISet, event IEvent, dragImageOffset foundation.PointPointer) Image {
	rv := objc.Call[Image](c_, objc.Sel("draggingImageForItemsAtIndexPaths:withEvent:offset:"), objc.Ptr(indexPaths), objc.Ptr(event), dragImageOffset)
	return rv
}

// Creates or returns a reusable supplementary view of the specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528302-makesupplementaryviewofkind?language=objc
func (c_ CollectionView) MakeSupplementaryViewOfKindWithIdentifierForIndexPath(elementKind CollectionViewSupplementaryElementKind, identifier UserInterfaceItemIdentifier, indexPath foundation.IIndexPath) View {
	rv := objc.Call[View](c_, objc.Sel("makeSupplementaryViewOfKind:withIdentifier:forIndexPath:"), elementKind, identifier, objc.Ptr(indexPath))
	return rv
}

// Returns the index paths of the currently active items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528223-indexpathsforvisibleitems?language=objc
func (c_ CollectionView) IndexPathsForVisibleItems() foundation.Set {
	rv := objc.Call[foundation.Set](c_, objc.Sel("indexPathsForVisibleItems"))
	return rv
}

// Removes the specified items from the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528172-deselectitemsatindexpaths?language=objc
func (c_ CollectionView) DeselectItemsAtIndexPaths(indexPaths foundation.ISet) {
	objc.Call[objc.Void](c_, objc.Sel("deselectItemsAtIndexPaths:"), objc.Ptr(indexPaths))
}

// Configures the default value returned from draggingSourceOperationMaskForLocal:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528229-setdraggingsourceoperationmask?language=objc
func (c_ CollectionView) SetDraggingSourceOperationMaskForLocal(dragOperationMask DragOperation, localDestination bool) {
	objc.Call[objc.Void](c_, objc.Sel("setDraggingSourceOperationMask:forLocal:"), dragOperationMask, localDestination)
}

// Returns the supplementary view associated with the specified index path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1524880-supplementaryviewforelementkind?language=objc
func (c_ CollectionView) SupplementaryViewForElementKindAtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) View {
	rv := objc.Call[View](c_, objc.Sel("supplementaryViewForElementKind:atIndexPath:"), elementKind, objc.Ptr(indexPath))
	return rv
}

// Inserts new items into the collection view at the specified locations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528267-insertitemsatindexpaths?language=objc
func (c_ CollectionView) InsertItemsAtIndexPaths(indexPaths foundation.ISet) {
	objc.Call[objc.Void](c_, objc.Sel("insertItemsAtIndexPaths:"), objc.Ptr(indexPaths))
}

// Returns an array of the actively managed items in the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528194-visibleitems?language=objc
func (c_ CollectionView) VisibleItems() []CollectionViewItem {
	rv := objc.Call[[]CollectionViewItem](c_, objc.Sel("visibleItems"))
	return rv
}

// Collapses the section in which the sender resides into a single horizontally scrollable row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1644691-togglesectioncollapse?language=objc
func (c_ CollectionView) ToggleSectionCollapse(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("toggleSectionCollapse:"), sender)
	return rv
}

// Moves a section from its current location to a new location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1524485-movesection?language=objc
func (c_ CollectionView) MoveSectionToSection(section int, newSection int) {
	objc.Call[objc.Void](c_, objc.Sel("moveSection:toSection:"), section, newSection)
}

// Returns the layout information for the item at the specified index path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528243-layoutattributesforitematindexpa?language=objc
func (c_ CollectionView) LayoutAttributesForItemAtIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.Call[CollectionViewLayoutAttributes](c_, objc.Sel("layoutAttributesForItemAtIndexPath:"), objc.Ptr(indexPath))
	return rv
}

// Returns the layout information for the supplementary view at the specified index path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1526721-layoutattributesforsupplementary?language=objc
func (c_ CollectionView) LayoutAttributesForSupplementaryElementOfKindAtIndexPath(kind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.Call[CollectionViewLayoutAttributes](c_, objc.Sel("layoutAttributesForSupplementaryElementOfKind:atIndexPath:"), kind, objc.Ptr(indexPath))
	return rv
}

// Returns the item associated with the specified index path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528184-itematindexpath?language=objc
func (c_ CollectionView) ItemAtIndexPath(indexPath foundation.IIndexPath) CollectionViewItem {
	rv := objc.Call[CollectionViewItem](c_, objc.Sel("itemAtIndexPath:"), objc.Ptr(indexPath))
	return rv
}

// Reloads all of the data for the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528264-reloaddata?language=objc
func (c_ CollectionView) ReloadData() {
	objc.Call[objc.Void](c_, objc.Sel("reloadData"))
}

// Reloads the data in the specified sections of the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528306-reloadsections?language=objc
func (c_ CollectionView) ReloadSections(sections foundation.IIndexSet) {
	objc.Call[objc.Void](c_, objc.Sel("reloadSections:"), objc.Ptr(sections))
}

// Deletes the items at the specified index paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528128-deleteitemsatindexpaths?language=objc
func (c_ CollectionView) DeleteItemsAtIndexPaths(indexPaths foundation.ISet) {
	objc.Call[objc.Void](c_, objc.Sel("deleteItemsAtIndexPaths:"), objc.Ptr(indexPaths))
}

// Moves an item from one location to another in the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528159-moveitematindexpath?language=objc
func (c_ CollectionView) MoveItemAtIndexPathToIndexPath(indexPath foundation.IIndexPath, newIndexPath foundation.IIndexPath) {
	objc.Call[objc.Void](c_, objc.Sel("moveItemAtIndexPath:toIndexPath:"), objc.Ptr(indexPath), objc.Ptr(newIndexPath))
}

// Returns an array of the actively managed supplementary views in the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528203-visiblesupplementaryviewsofkind?language=objc
func (c_ CollectionView) VisibleSupplementaryViewsOfKind(elementKind CollectionViewSupplementaryElementKind) []View {
	rv := objc.Call[[]View](c_, objc.Sel("visibleSupplementaryViewsOfKind:"), elementKind)
	return rv
}

// An object that provides data for the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528222-datasource?language=objc
func (c_ CollectionView) DataSource() CollectionViewDataSourceWrapper {
	rv := objc.Call[CollectionViewDataSourceWrapper](c_, objc.Sel("dataSource"))
	return rv
}

// An object that provides data for the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528222-datasource?language=objc
func (c_ CollectionView) SetDataSource(value PCollectionViewDataSource) {
	po0 := objc.WrapAsProtocol("NSCollectionViewDataSource", value)
	objc.SetAssociatedObject(c_, objc.AssociationKey("setDataSource"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](c_, objc.Sel("setDataSource:"), po0)
}

// An object that provides data for the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528222-datasource?language=objc
func (c_ CollectionView) SetDataSourceObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setDataSource:"), objc.Ptr(valueObject))
}

// An array that provides data for the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528207-content?language=objc
func (c_ CollectionView) Content() []objc.Object {
	rv := objc.Call[[]objc.Object](c_, objc.Sel("content"))
	return rv
}

// An array that provides data for the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528207-content?language=objc
func (c_ CollectionView) SetContent(value []objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setContent:"), value)
}

// An array containing the collection view’s background colors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528220-backgroundcolors?language=objc
func (c_ CollectionView) BackgroundColors() []Color {
	rv := objc.Call[[]Color](c_, objc.Sel("backgroundColors"))
	return rv
}

// An array containing the collection view’s background colors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528220-backgroundcolors?language=objc
func (c_ CollectionView) SetBackgroundColors(value []IColor) {
	objc.Call[objc.Void](c_, objc.Sel("setBackgroundColors:"), value)
}

// A Boolean value indicating whether the collection view may have no selected items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528180-allowsemptyselection?language=objc
func (c_ CollectionView) AllowsEmptySelection() bool {
	rv := objc.Call[bool](c_, objc.Sel("allowsEmptySelection"))
	return rv
}

// A Boolean value indicating whether the collection view may have no selected items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528180-allowsemptyselection?language=objc
func (c_ CollectionView) SetAllowsEmptySelection(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setAllowsEmptySelection:"), value)
}

// The background view placed behind all items and supplementary views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528250-backgroundview?language=objc
func (c_ CollectionView) BackgroundView() View {
	rv := objc.Call[View](c_, objc.Sel("backgroundView"))
	return rv
}

// The background view placed behind all items and supplementary views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528250-backgroundview?language=objc
func (c_ CollectionView) SetBackgroundView(value IView) {
	objc.Call[objc.Void](c_, objc.Sel("setBackgroundView:"), objc.Ptr(value))
}

// A Boolean value indicating whether the collection view is the first responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528199-firstresponder?language=objc
func (c_ CollectionView) IsFirstResponder() bool {
	rv := objc.Call[bool](c_, objc.Sel("isFirstResponder"))
	return rv
}

// The layout object used to organize the collection view’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528271-collectionviewlayout?language=objc
func (c_ CollectionView) CollectionViewLayout() CollectionViewLayout {
	rv := objc.Call[CollectionViewLayout](c_, objc.Sel("collectionViewLayout"))
	return rv
}

// The layout object used to organize the collection view’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528271-collectionviewlayout?language=objc
func (c_ CollectionView) SetCollectionViewLayout(value ICollectionViewLayout) {
	objc.Call[objc.Void](c_, objc.Sel("setCollectionViewLayout:"), objc.Ptr(value))
}

// The collection view’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528246-delegate?language=objc
func (c_ CollectionView) Delegate() CollectionViewDelegateWrapper {
	rv := objc.Call[CollectionViewDelegateWrapper](c_, objc.Sel("delegate"))
	return rv
}

// The collection view’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528246-delegate?language=objc
func (c_ CollectionView) SetDelegate(value PCollectionViewDelegate) {
	po0 := objc.WrapAsProtocol("NSCollectionViewDelegate", value)
	objc.SetAssociatedObject(c_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](c_, objc.Sel("setDelegate:"), po0)
}

// The collection view’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528246-delegate?language=objc
func (c_ CollectionView) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The set of index paths representing the currently selected items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528275-selectionindexpaths?language=objc
func (c_ CollectionView) SelectionIndexPaths() foundation.Set {
	rv := objc.Call[foundation.Set](c_, objc.Sel("selectionIndexPaths"))
	return rv
}

// The set of index paths representing the currently selected items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528275-selectionindexpaths?language=objc
func (c_ CollectionView) SetSelectionIndexPaths(value foundation.ISet) {
	objc.Call[objc.Void](c_, objc.Sel("setSelectionIndexPaths:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the collection view’s background view scrolls with the items and other content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1644533-backgroundviewscrollswithcontent?language=objc
func (c_ CollectionView) BackgroundViewScrollsWithContent() bool {
	rv := objc.Call[bool](c_, objc.Sel("backgroundViewScrollsWithContent"))
	return rv
}

// A Boolean value that indicates whether the collection view’s background view scrolls with the items and other content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1644533-backgroundviewscrollswithcontent?language=objc
func (c_ CollectionView) SetBackgroundViewScrollsWithContent(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setBackgroundViewScrollsWithContent:"), value)
}

// A Boolean value that indicates whether the user may select items in the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528300-selectable?language=objc
func (c_ CollectionView) IsSelectable() bool {
	rv := objc.Call[bool](c_, objc.Sel("isSelectable"))
	return rv
}

// A Boolean value that indicates whether the user may select items in the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528300-selectable?language=objc
func (c_ CollectionView) SetSelectable(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setSelectable:"), value)
}

// The indexes of the currently selected items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1525505-selectionindexes?language=objc
func (c_ CollectionView) SelectionIndexes() foundation.IndexSet {
	rv := objc.Call[foundation.IndexSet](c_, objc.Sel("selectionIndexes"))
	return rv
}

// The indexes of the currently selected items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1525505-selectionindexes?language=objc
func (c_ CollectionView) SetSelectionIndexes(value foundation.IIndexSet) {
	objc.Call[objc.Void](c_, objc.Sel("setSelectionIndexes:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the user may select more than one item in the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1524276-allowsmultipleselection?language=objc
func (c_ CollectionView) AllowsMultipleSelection() bool {
	rv := objc.Call[bool](c_, objc.Sel("allowsMultipleSelection"))
	return rv
}

// A Boolean value that indicates whether the user may select more than one item in the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1524276-allowsmultipleselection?language=objc
func (c_ CollectionView) SetAllowsMultipleSelection(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setAllowsMultipleSelection:"), value)
}

// The number of sections in the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/1528238-numberofsections?language=objc
func (c_ CollectionView) NumberOfSections() int {
	rv := objc.Call[int](c_, objc.Sel("numberOfSections"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/2879292-prefetchdatasource?language=objc
func (c_ CollectionView) PrefetchDataSource() CollectionViewPrefetchingWrapper {
	rv := objc.Call[CollectionViewPrefetchingWrapper](c_, objc.Sel("prefetchDataSource"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/2879292-prefetchdatasource?language=objc
func (c_ CollectionView) SetPrefetchDataSource(value PCollectionViewPrefetching) {
	po0 := objc.WrapAsProtocol("NSCollectionViewPrefetching", value)
	objc.SetAssociatedObject(c_, objc.AssociationKey("setPrefetchDataSource"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](c_, objc.Sel("setPrefetchDataSource:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/2879292-prefetchdatasource?language=objc
func (c_ CollectionView) SetPrefetchDataSourceObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setPrefetchDataSource:"), objc.Ptr(valueObject))
}
