// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TableView] class.
var TableViewClass = _TableViewClass{objc.GetClass("NSTableView")}

type _TableViewClass struct {
	objc.Class
}

// An interface definition for the [TableView] class.
type ITableView interface {
	IControl
	DrawBackgroundInClipRect(clipRect foundation.Rect)
	ScrollRowToVisible(row int)
	ReloadDataForRowIndexesColumnIndexes(rowIndexes foundation.IIndexSet, columnIndexes foundation.IIndexSet)
	RegisterNibForIdentifier(nib INib, identifier UserInterfaceItemIdentifier)
	BeginUpdates()
	DrawRowClipRect(row int, clipRect foundation.Rect)
	ViewAtColumnRowMakeIfNecessary(column int, row int, makeIfNecessary bool) View
	TableColumnWithIdentifier(identifier UserInterfaceItemIdentifier) TableColumn
	RowForView(view IView) int
	SetDropRowDropOperation(row int, dropOperation TableViewDropOperation)
	EndUpdates()
	HideRowsAtIndexesWithAnimation(indexes foundation.IIndexSet, rowAnimation TableViewAnimationOptions)
	DeselectAll(sender objc.IObject)
	DeselectColumn(column int)
	RectOfRow(row int) foundation.Rect
	ColumnAtPoint(point foundation.Point) int
	UnhideRowsAtIndexesWithAnimation(indexes foundation.IIndexSet, rowAnimation TableViewAnimationOptions)
	HighlightSelectionInClipRect(clipRect foundation.Rect)
	ScrollColumnToVisible(column int)
	EnumerateAvailableRowViewsUsingBlock(handler func(rowView TableRowView, row int))
	DidRemoveRowViewForRow(rowView ITableRowView, row int)
	EditColumnRowWithEventSelect(column int, row int, event IEvent, select_ bool)
	DrawGridInClipRect(clipRect foundation.Rect)
	SelectAll(sender objc.IObject)
	IndicatorImageInTableColumn(tableColumn ITableColumn) Image
	IsColumnSelected(column int) bool
	DeselectRow(row int)
	RowsInRect(rect foundation.Rect) foundation.Range
	SetIndicatorImageInTableColumn(image IImage, tableColumn ITableColumn)
	ColumnIndexesInRect(rect foundation.Rect) foundation.IndexSet
	AddTableColumn(tableColumn ITableColumn)
	RemoveRowsAtIndexesWithAnimation(indexes foundation.IIndexSet, animationOptions TableViewAnimationOptions)
	MakeViewWithIdentifierOwner(identifier UserInterfaceItemIdentifier, owner objc.IObject) View
	InsertRowsAtIndexesWithAnimation(indexes foundation.IIndexSet, animationOptions TableViewAnimationOptions)
	RowAtPoint(point foundation.Point) int
	SelectColumnIndexesByExtendingSelection(indexes foundation.IIndexSet, extend bool)
	DragImageForRowsWithIndexesTableColumnsEventOffset(dragRows foundation.IIndexSet, tableColumns []ITableColumn, dragEvent IEvent, dragImageOffset foundation.PointPointer) Image
	RectOfColumn(column int) foundation.Rect
	SizeLastColumnToFit()
	ColumnWithIdentifier(identifier UserInterfaceItemIdentifier) int
	SelectRowIndexesByExtendingSelection(indexes foundation.IIndexSet, extend bool)
	NoteHeightOfRowsWithIndexesChanged(indexSet foundation.IIndexSet)
	CanDragRowsWithIndexesAtPoint(rowIndexes foundation.IIndexSet, mouseDownPoint foundation.Point) bool
	SetDraggingSourceOperationMaskForLocal(mask DragOperation, isLocal bool)
	FrameOfCellAtColumnRow(column int, row int) foundation.Rect
	RowViewAtRowMakeIfNecessary(row int, makeIfNecessary bool) TableRowView
	Tile()
	MoveRowAtIndexToIndex(oldIndex int, newIndex int)
	ColumnForView(view IView) int
	NoteNumberOfRowsChanged()
	DidAddRowViewForRow(rowView ITableRowView, row int)
	MoveColumnToColumn(oldIndex int, newIndex int)
	ReloadData()
	IsRowSelected(row int) bool
	RemoveTableColumn(tableColumn ITableColumn)
	DataSource() TableViewDataSourceWrapper
	SetDataSource(value PTableViewDataSource)
	SetDataSourceObject(valueObject objc.IObject)
	VerticalMotionCanBeginDrag() bool
	SetVerticalMotionCanBeginDrag(value bool)
	HighlightedTableColumn() TableColumn
	SetHighlightedTableColumn(value ITableColumn)
	DoubleAction() objc.Selector
	SetDoubleAction(value objc.Selector)
	EditedRow() int
	HiddenRowIndexes() foundation.IndexSet
	UsesStaticContents() bool
	SetUsesStaticContents(value bool)
	AutosaveTableColumns() bool
	SetAutosaveTableColumns(value bool)
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	AllowsColumnResizing() bool
	SetAllowsColumnResizing(value bool)
	RowHeight() float64
	SetRowHeight(value float64)
	EditedColumn() int
	AllowsColumnReordering() bool
	SetAllowsColumnReordering(value bool)
	ClickedColumn() int
	RowSizeStyle() TableViewRowSizeStyle
	SetRowSizeStyle(value TableViewRowSizeStyle)
	AllowsTypeSelect() bool
	SetAllowsTypeSelect(value bool)
	SortDescriptors() []foundation.SortDescriptor
	SetSortDescriptors(value []foundation.ISortDescriptor)
	Style() TableViewStyle
	SetStyle(value TableViewStyle)
	CornerView() View
	SetCornerView(value IView)
	EffectiveRowSizeStyle() TableViewRowSizeStyle
	Delegate() TableViewDelegateWrapper
	SetDelegate(value PTableViewDelegate)
	SetDelegateObject(valueObject objc.IObject)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	AutosaveName() TableViewAutosaveName
	SetAutosaveName(value TableViewAutosaveName)
	EffectiveStyle() TableViewStyle
	UsesAutomaticRowHeights() bool
	SetUsesAutomaticRowHeights(value bool)
	UsesAlternatingRowBackgroundColors() bool
	SetUsesAlternatingRowBackgroundColors(value bool)
	NumberOfSelectedRows() int
	ClickedRow() int
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	SelectedColumn() int
	GridStyleMask() TableViewGridLineStyle
	SetGridStyleMask(value TableViewGridLineStyle)
	SelectedRow() int
	SelectionHighlightStyle() TableViewSelectionHighlightStyle
	SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle)
	RowActionsVisible() bool
	SetRowActionsVisible(value bool)
	RegisteredNibsByIdentifier() map[UserInterfaceItemIdentifier]Nib
	SelectedColumnIndexes() foundation.IndexSet
	AllowsColumnSelection() bool
	SetAllowsColumnSelection(value bool)
	HeaderView() TableHeaderView
	SetHeaderView(value ITableHeaderView)
	NumberOfRows() int
	FloatsGroupRows() bool
	SetFloatsGroupRows(value bool)
	SelectedRowIndexes() foundation.IndexSet
	IntercellSpacing() foundation.Size
	SetIntercellSpacing(value foundation.Size)
	NumberOfColumns() int
	ColumnAutoresizingStyle() TableViewColumnAutoresizingStyle
	SetColumnAutoresizingStyle(value TableViewColumnAutoresizingStyle)
	GridColor() Color
	SetGridColor(value IColor)
	TableColumns() []TableColumn
	DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle
	SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle)
	NumberOfSelectedColumns() int
}

// A set of related records, displayed in rows that represent individual records and columns that represent the attributes of those records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview?language=objc
type TableView struct {
	Control
}

func TableViewFrom(ptr unsafe.Pointer) TableView {
	return TableView{
		Control: ControlFrom(ptr),
	}
}

func (t_ TableView) InitWithFrame(frameRect foundation.Rect) TableView {
	rv := objc.Call[TableView](t_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1525511-initwithframe?language=objc
func NewTableViewWithFrame(frameRect foundation.Rect) TableView {
	instance := TableViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

func (tc _TableViewClass) Alloc() TableView {
	rv := objc.Call[TableView](tc, objc.Sel("alloc"))
	return rv
}

func TableView_Alloc() TableView {
	return TableViewClass.Alloc()
}

func (tc _TableViewClass) New() TableView {
	rv := objc.Call[TableView](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTableView() TableView {
	return TableViewClass.New()
}

func (t_ TableView) Init() TableView {
	rv := objc.Call[TableView](t_, objc.Sel("init"))
	return rv
}

// Draws the background of the table view in the clip rect specified by the rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1528866-drawbackgroundincliprect?language=objc
func (t_ TableView) DrawBackgroundInClipRect(clipRect foundation.Rect) {
	objc.Call[objc.Void](t_, objc.Sel("drawBackgroundInClipRect:"), clipRect)
}

// Scrolls the view so the specified row is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1529902-scrollrowtovisible?language=objc
func (t_ TableView) ScrollRowToVisible(row int) {
	objc.Call[objc.Void](t_, objc.Sel("scrollRowToVisible:"), row)
}

// Reloads the data for only the specified rows and columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1527621-reloaddataforrowindexes?language=objc
func (t_ TableView) ReloadDataForRowIndexesColumnIndexes(rowIndexes foundation.IIndexSet, columnIndexes foundation.IIndexSet) {
	objc.Call[objc.Void](t_, objc.Sel("reloadDataForRowIndexes:columnIndexes:"), objc.Ptr(rowIndexes), objc.Ptr(columnIndexes))
}

// Registers a NIB for the specified identifier, so that view-based table views can use it to instantiate views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1524297-registernib?language=objc
func (t_ TableView) RegisterNibForIdentifier(nib INib, identifier UserInterfaceItemIdentifier) {
	objc.Call[objc.Void](t_, objc.Sel("registerNib:forIdentifier:"), objc.Ptr(nib), identifier)
}

// Begins a group of updates for the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1527288-beginupdates?language=objc
func (t_ TableView) BeginUpdates() {
	objc.Call[objc.Void](t_, objc.Sel("beginUpdates"))
}

// Draws the cells for the row at rowIndex in the columns that intersect clipRect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1533711-drawrow?language=objc
func (t_ TableView) DrawRowClipRect(row int, clipRect foundation.Rect) {
	objc.Call[objc.Void](t_, objc.Sel("drawRow:clipRect:"), row, clipRect)
}

// Returns a view at the specified row and column indexes, creating one if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1528831-viewatcolumn?language=objc
func (t_ TableView) ViewAtColumnRowMakeIfNecessary(column int, row int, makeIfNecessary bool) View {
	rv := objc.Call[View](t_, objc.Sel("viewAtColumn:row:makeIfNecessary:"), column, row, makeIfNecessary)
	return rv
}

// Returns the NSTableColumn object for the first column whose identifier is equal to the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1531134-tablecolumnwithidentifier?language=objc
func (t_ TableView) TableColumnWithIdentifier(identifier UserInterfaceItemIdentifier) TableColumn {
	rv := objc.Call[TableColumn](t_, objc.Sel("tableColumnWithIdentifier:"), identifier)
	return rv
}

// Returns the index of the row for the specified view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1526732-rowforview?language=objc
func (t_ TableView) RowForView(view IView) int {
	rv := objc.Call[int](t_, objc.Sel("rowForView:"), objc.Ptr(view))
	return rv
}

// Retargets the proposed drop operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1535123-setdroprow?language=objc
func (t_ TableView) SetDropRowDropOperation(row int, dropOperation TableViewDropOperation) {
	objc.Call[objc.Void](t_, objc.Sel("setDropRow:dropOperation:"), row, dropOperation)
}

// Ends the group of updates for the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1526267-endupdates?language=objc
func (t_ TableView) EndUpdates() {
	objc.Call[objc.Void](t_, objc.Sel("endUpdates"))
}

// Hides the specified table rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534141-hiderowsatindexes?language=objc
func (t_ TableView) HideRowsAtIndexesWithAnimation(indexes foundation.IIndexSet, rowAnimation TableViewAnimationOptions) {
	objc.Call[objc.Void](t_, objc.Sel("hideRowsAtIndexes:withAnimation:"), objc.Ptr(indexes), rowAnimation)
}

// Deselects all selected rows or columns if empty selection is allowed; otherwise does nothing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1533302-deselectall?language=objc
func (t_ TableView) DeselectAll(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("deselectAll:"), sender)
}

// Deselects the column at the specified index if it’s selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1525334-deselectcolumn?language=objc
func (t_ TableView) DeselectColumn(column int) {
	objc.Call[objc.Void](t_, objc.Sel("deselectColumn:"), column)
}

// Returns the rectangle containing the row at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1532948-rectofrow?language=objc
func (t_ TableView) RectOfRow(row int) foundation.Rect {
	rv := objc.Call[foundation.Rect](t_, objc.Sel("rectOfRow:"), row)
	return rv
}

// Returns the index of the column the specified point lies in. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1527504-columnatpoint?language=objc
func (t_ TableView) ColumnAtPoint(point foundation.Point) int {
	rv := objc.Call[int](t_, objc.Sel("columnAtPoint:"), point)
	return rv
}

// Unhides the specified table rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1527447-unhiderowsatindexes?language=objc
func (t_ TableView) UnhideRowsAtIndexesWithAnimation(indexes foundation.IIndexSet, rowAnimation TableViewAnimationOptions) {
	objc.Call[objc.Void](t_, objc.Sel("unhideRowsAtIndexes:withAnimation:"), objc.Ptr(indexes), rowAnimation)
}

// Highlights the region of the table view in the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1530758-highlightselectionincliprect?language=objc
func (t_ TableView) HighlightSelectionInClipRect(clipRect foundation.Rect) {
	objc.Call[objc.Void](t_, objc.Sel("highlightSelectionInClipRect:"), clipRect)
}

// Scrolls the view so the specified column is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1535570-scrollcolumntovisible?language=objc
func (t_ TableView) ScrollColumnToVisible(column int) {
	objc.Call[objc.Void](t_, objc.Sel("scrollColumnToVisible:"), column)
}

// Allows the enumeration of all the table rows that are known to the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1532750-enumerateavailablerowviewsusingb?language=objc
func (t_ TableView) EnumerateAvailableRowViewsUsingBlock(handler func(rowView TableRowView, row int)) {
	objc.Call[objc.Void](t_, objc.Sel("enumerateAvailableRowViewsUsingBlock:"), handler)
}

// Invoked when a row view is removed from the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1532903-didremoverowview?language=objc
func (t_ TableView) DidRemoveRowViewForRow(rowView ITableRowView, row int) {
	objc.Call[objc.Void](t_, objc.Sel("didRemoveRowView:forRow:"), objc.Ptr(rowView), row)
}

// Edits the cell at the specified column and row using the specified event and selection behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1526295-editcolumn?language=objc
func (t_ TableView) EditColumnRowWithEventSelect(column int, row int, event IEvent, select_ bool) {
	objc.Call[objc.Void](t_, objc.Sel("editColumn:row:withEvent:select:"), column, row, objc.Ptr(event), select_)
}

// Draws the grid lines within the supplied rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1527006-drawgridincliprect?language=objc
func (t_ TableView) DrawGridInClipRect(clipRect foundation.Rect) {
	objc.Call[objc.Void](t_, objc.Sel("drawGridInClipRect:"), clipRect)
}

// Selects all rows or all columns, according to whether rows or columns were most recently selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534002-selectall?language=objc
func (t_ TableView) SelectAll(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("selectAll:"), sender)
}

// Returns the indicator image of the specified table column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1524846-indicatorimageintablecolumn?language=objc
func (t_ TableView) IndicatorImageInTableColumn(tableColumn ITableColumn) Image {
	rv := objc.Call[Image](t_, objc.Sel("indicatorImageInTableColumn:"), objc.Ptr(tableColumn))
	return rv
}

// Returns a Boolean value that indicates whether the column at the specified index is selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1524306-iscolumnselected?language=objc
func (t_ TableView) IsColumnSelected(column int) bool {
	rv := objc.Call[bool](t_, objc.Sel("isColumnSelected:"), column)
	return rv
}

// Deselects the row at the specified index if it’s selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1532722-deselectrow?language=objc
func (t_ TableView) DeselectRow(row int) {
	objc.Call[objc.Void](t_, objc.Sel("deselectRow:"), row)
}

// Returns a range of indexes for the rows that lie wholly or partially within the vertical boundaries of the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1525344-rowsinrect?language=objc
func (t_ TableView) RowsInRect(rect foundation.Rect) foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("rowsInRect:"), rect)
	return rv
}

// Sets the indicator image of the specified column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534381-setindicatorimage?language=objc
func (t_ TableView) SetIndicatorImageInTableColumn(image IImage, tableColumn ITableColumn) {
	objc.Call[objc.Void](t_, objc.Sel("setIndicatorImage:inTableColumn:"), objc.Ptr(image), objc.Ptr(tableColumn))
}

// Returns the indexes of the table view’s columns that intersect the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1526013-columnindexesinrect?language=objc
func (t_ TableView) ColumnIndexesInRect(rect foundation.Rect) foundation.IndexSet {
	rv := objc.Call[foundation.IndexSet](t_, objc.Sel("columnIndexesInRect:"), rect)
	return rv
}

// Adds the specified column as the last column of the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534098-addtablecolumn?language=objc
func (t_ TableView) AddTableColumn(tableColumn ITableColumn) {
	objc.Call[objc.Void](t_, objc.Sel("addTableColumn:"), objc.Ptr(tableColumn))
}

// Removes the rows using the specified animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1524655-removerowsatindexes?language=objc
func (t_ TableView) RemoveRowsAtIndexesWithAnimation(indexes foundation.IIndexSet, animationOptions TableViewAnimationOptions) {
	objc.Call[objc.Void](t_, objc.Sel("removeRowsAtIndexes:withAnimation:"), objc.Ptr(indexes), animationOptions)
}

// Returns a new or existing view with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1535482-makeviewwithidentifier?language=objc
func (t_ TableView) MakeViewWithIdentifierOwner(identifier UserInterfaceItemIdentifier, owner objc.IObject) View {
	rv := objc.Call[View](t_, objc.Sel("makeViewWithIdentifier:owner:"), identifier, owner)
	return rv
}

// Inserts the rows using the specified animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1532406-insertrowsatindexes?language=objc
func (t_ TableView) InsertRowsAtIndexesWithAnimation(indexes foundation.IIndexSet, animationOptions TableViewAnimationOptions) {
	objc.Call[objc.Void](t_, objc.Sel("insertRowsAtIndexes:withAnimation:"), objc.Ptr(indexes), animationOptions)
}

// Returns the index of the row the specified point lies in. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1530884-rowatpoint?language=objc
func (t_ TableView) RowAtPoint(point foundation.Point) int {
	rv := objc.Call[int](t_, objc.Sel("rowAtPoint:"), point)
	return rv
}

// Sets the column selection using indexes possibly extending the selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1524416-selectcolumnindexes?language=objc
func (t_ TableView) SelectColumnIndexesByExtendingSelection(indexes foundation.IIndexSet, extend bool) {
	objc.Call[objc.Void](t_, objc.Sel("selectColumnIndexes:byExtendingSelection:"), objc.Ptr(indexes), extend)
}

// Computes and returns an image to use for dragging. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1526541-dragimageforrowswithindexes?language=objc
func (t_ TableView) DragImageForRowsWithIndexesTableColumnsEventOffset(dragRows foundation.IIndexSet, tableColumns []ITableColumn, dragEvent IEvent, dragImageOffset foundation.PointPointer) Image {
	rv := objc.Call[Image](t_, objc.Sel("dragImageForRowsWithIndexes:tableColumns:event:offset:"), objc.Ptr(dragRows), tableColumns, objc.Ptr(dragEvent), dragImageOffset)
	return rv
}

// Returns the rectangle containing the column at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1531546-rectofcolumn?language=objc
func (t_ TableView) RectOfColumn(column int) foundation.Rect {
	rv := objc.Call[foundation.Rect](t_, objc.Sel("rectOfColumn:"), column)
	return rv
}

// Resizes the last column so the table view fits exactly within its enclosing clip view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1525859-sizelastcolumntofit?language=objc
func (t_ TableView) SizeLastColumnToFit() {
	objc.Call[objc.Void](t_, objc.Sel("sizeLastColumnToFit"))
}

// Returns the index of the first column in the table view whose identifier is equal to the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1526734-columnwithidentifier?language=objc
func (t_ TableView) ColumnWithIdentifier(identifier UserInterfaceItemIdentifier) int {
	rv := objc.Call[int](t_, objc.Sel("columnWithIdentifier:"), identifier)
	return rv
}

// Sets the row selection using indexes extending the selection if specified. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1529688-selectrowindexes?language=objc
func (t_ TableView) SelectRowIndexesByExtendingSelection(indexes foundation.IIndexSet, extend bool) {
	objc.Call[objc.Void](t_, objc.Sel("selectRowIndexes:byExtendingSelection:"), objc.Ptr(indexes), extend)
}

// Informs the table view that the rows specified in indexSet have changed height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1532126-noteheightofrowswithindexeschang?language=objc
func (t_ TableView) NoteHeightOfRowsWithIndexesChanged(indexSet foundation.IIndexSet) {
	objc.Call[objc.Void](t_, objc.Sel("noteHeightOfRowsWithIndexesChanged:"), objc.Ptr(indexSet))
}

// Returns a Boolean value indicating whether the table view allows dragging the rows with the drag initiated at the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1524273-candragrowswithindexes?language=objc
func (t_ TableView) CanDragRowsWithIndexesAtPoint(rowIndexes foundation.IIndexSet, mouseDownPoint foundation.Point) bool {
	rv := objc.Call[bool](t_, objc.Sel("canDragRowsWithIndexes:atPoint:"), objc.Ptr(rowIndexes), mouseDownPoint)
	return rv
}

// Sets the default operation mask returned by draggingSourceOperationMaskForLocal: to mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1527199-setdraggingsourceoperationmask?language=objc
func (t_ TableView) SetDraggingSourceOperationMaskForLocal(mask DragOperation, isLocal bool) {
	objc.Call[objc.Void](t_, objc.Sel("setDraggingSourceOperationMask:forLocal:"), mask, isLocal)
}

// Returns a rectangle locating the cell that lies at the intersection of the specified column and row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1524517-frameofcellatcolumn?language=objc
func (t_ TableView) FrameOfCellAtColumnRow(column int, row int) foundation.Rect {
	rv := objc.Call[foundation.Rect](t_, objc.Sel("frameOfCellAtColumn:row:"), column, row)
	return rv
}

// Returns a row view at the specified index, creating one if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1525162-rowviewatrow?language=objc
func (t_ TableView) RowViewAtRowMakeIfNecessary(row int, makeIfNecessary bool) TableRowView {
	rv := objc.Call[TableRowView](t_, objc.Sel("rowViewAtRow:makeIfNecessary:"), row, makeIfNecessary)
	return rv
}

// Properly sizes the table view and its header view and marks it as needing display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1528077-tile?language=objc
func (t_ TableView) Tile() {
	objc.Call[objc.Void](t_, objc.Sel("tile"))
}

// Moves the specified row to the new row location using animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1535835-moverowatindex?language=objc
func (t_ TableView) MoveRowAtIndexToIndex(oldIndex int, newIndex int) {
	objc.Call[objc.Void](t_, objc.Sel("moveRowAtIndex:toIndex:"), oldIndex, newIndex)
}

// Returns the column index for the specified view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1529415-columnforview?language=objc
func (t_ TableView) ColumnForView(view IView) int {
	rv := objc.Call[int](t_, objc.Sel("columnForView:"), objc.Ptr(view))
	return rv
}

// Informs the table view that the number of records in its data source has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534147-notenumberofrowschanged?language=objc
func (t_ TableView) NoteNumberOfRowsChanged() {
	objc.Call[objc.Void](t_, objc.Sel("noteNumberOfRowsChanged"))
}

// Invoked when a row view is added to the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534008-didaddrowview?language=objc
func (t_ TableView) DidAddRowViewForRow(rowView ITableRowView, row int) {
	objc.Call[objc.Void](t_, objc.Sel("didAddRowView:forRow:"), objc.Ptr(rowView), row)
}

// Moves the column and heading at the specified index to the new specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1530719-movecolumn?language=objc
func (t_ TableView) MoveColumnToColumn(oldIndex int, newIndex int) {
	objc.Call[objc.Void](t_, objc.Sel("moveColumn:toColumn:"), oldIndex, newIndex)
}

// Marks the table view as needing redisplay, so it will reload the data for visible cells and draw the new values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1528382-reloaddata?language=objc
func (t_ TableView) ReloadData() {
	objc.Call[objc.Void](t_, objc.Sel("reloadData"))
}

// Returns a Boolean value that indicates whether the row at the specified index is selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1525882-isrowselected?language=objc
func (t_ TableView) IsRowSelected(row int) bool {
	rv := objc.Call[bool](t_, objc.Sel("isRowSelected:"), row)
	return rv
}

// Removes the specified column from the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1535381-removetablecolumn?language=objc
func (t_ TableView) RemoveTableColumn(tableColumn ITableColumn) {
	objc.Call[objc.Void](t_, objc.Sel("removeTableColumn:"), objc.Ptr(tableColumn))
}

// The object that provides the data displayed by the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1531866-datasource?language=objc
func (t_ TableView) DataSource() TableViewDataSourceWrapper {
	rv := objc.Call[TableViewDataSourceWrapper](t_, objc.Sel("dataSource"))
	return rv
}

// The object that provides the data displayed by the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1531866-datasource?language=objc
func (t_ TableView) SetDataSource(value PTableViewDataSource) {
	po0 := objc.WrapAsProtocol("NSTableViewDataSource", value)
	objc.SetAssociatedObject(t_, objc.AssociationKey("setDataSource"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](t_, objc.Sel("setDataSource:"), po0)
}

// The object that provides the data displayed by the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1531866-datasource?language=objc
func (t_ TableView) SetDataSourceObject(valueObject objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setDataSource:"), objc.Ptr(valueObject))
}

// A Boolean value indicating whether vertical motion is treated as a drag or selection change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534585-verticalmotioncanbegindrag?language=objc
func (t_ TableView) VerticalMotionCanBeginDrag() bool {
	rv := objc.Call[bool](t_, objc.Sel("verticalMotionCanBeginDrag"))
	return rv
}

// A Boolean value indicating whether vertical motion is treated as a drag or selection change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534585-verticalmotioncanbegindrag?language=objc
func (t_ TableView) SetVerticalMotionCanBeginDrag(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setVerticalMotionCanBeginDrag:"), value)
}

// The column highlighted in the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1524980-highlightedtablecolumn?language=objc
func (t_ TableView) HighlightedTableColumn() TableColumn {
	rv := objc.Call[TableColumn](t_, objc.Sel("highlightedTableColumn"))
	return rv
}

// The column highlighted in the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1524980-highlightedtablecolumn?language=objc
func (t_ TableView) SetHighlightedTableColumn(value ITableColumn) {
	objc.Call[objc.Void](t_, objc.Sel("setHighlightedTableColumn:"), objc.Ptr(value))
}

// The message sent to the table view’s target when the user double-clicks a cell or column header. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1526992-doubleaction?language=objc
func (t_ TableView) DoubleAction() objc.Selector {
	rv := objc.Call[objc.Selector](t_, objc.Sel("doubleAction"))
	return rv
}

// The message sent to the table view’s target when the user double-clicks a cell or column header. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1526992-doubleaction?language=objc
func (t_ TableView) SetDoubleAction(value objc.Selector) {
	objc.Call[objc.Void](t_, objc.Sel("setDoubleAction:"), value)
}

// The index of the row being edited. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534282-editedrow?language=objc
func (t_ TableView) EditedRow() int {
	rv := objc.Call[int](t_, objc.Sel("editedRow"))
	return rv
}

// The indexes of all hidden table rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534914-hiddenrowindexes?language=objc
func (t_ TableView) HiddenRowIndexes() foundation.IndexSet {
	rv := objc.Call[foundation.IndexSet](t_, objc.Sel("hiddenRowIndexes"))
	return rv
}

// A Boolean value indicating whether the table uses static data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1533450-usesstaticcontents?language=objc
func (t_ TableView) UsesStaticContents() bool {
	rv := objc.Call[bool](t_, objc.Sel("usesStaticContents"))
	return rv
}

// A Boolean value indicating whether the table uses static data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1533450-usesstaticcontents?language=objc
func (t_ TableView) SetUsesStaticContents(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setUsesStaticContents:"), value)
}

// A Boolean value indicating whether the order and width of the table view’s columns are automatically saved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1525596-autosavetablecolumns?language=objc
func (t_ TableView) AutosaveTableColumns() bool {
	rv := objc.Call[bool](t_, objc.Sel("autosaveTableColumns"))
	return rv
}

// A Boolean value indicating whether the order and width of the table view’s columns are automatically saved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1525596-autosavetablecolumns?language=objc
func (t_ TableView) SetAutosaveTableColumns(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAutosaveTableColumns:"), value)
}

// A Boolean value indicating whether the table view allows the user to select zero columns or rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1535902-allowsemptyselection?language=objc
func (t_ TableView) AllowsEmptySelection() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsEmptySelection"))
	return rv
}

// A Boolean value indicating whether the table view allows the user to select zero columns or rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1535902-allowsemptyselection?language=objc
func (t_ TableView) SetAllowsEmptySelection(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsEmptySelection:"), value)
}

// A Boolean value indicating whether the table view allows the user to resize columns by dragging between their headers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1527826-allowscolumnresizing?language=objc
func (t_ TableView) AllowsColumnResizing() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsColumnResizing"))
	return rv
}

// A Boolean value indicating whether the table view allows the user to resize columns by dragging between their headers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1527826-allowscolumnresizing?language=objc
func (t_ TableView) SetAllowsColumnResizing(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsColumnResizing:"), value)
}

// The height of each row in the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1529148-rowheight?language=objc
func (t_ TableView) RowHeight() float64 {
	rv := objc.Call[float64](t_, objc.Sel("rowHeight"))
	return rv
}

// The height of each row in the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1529148-rowheight?language=objc
func (t_ TableView) SetRowHeight(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setRowHeight:"), value)
}

// The index of the column being edited. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1532307-editedcolumn?language=objc
func (t_ TableView) EditedColumn() int {
	rv := objc.Call[int](t_, objc.Sel("editedColumn"))
	return rv
}

// A Boolean value indicating whether the table view allows the user to rearrange columns by dragging their headers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1530207-allowscolumnreordering?language=objc
func (t_ TableView) AllowsColumnReordering() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsColumnReordering"))
	return rv
}

// A Boolean value indicating whether the table view allows the user to rearrange columns by dragging their headers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1530207-allowscolumnreordering?language=objc
func (t_ TableView) SetAllowsColumnReordering(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsColumnReordering:"), value)
}

// The index of the column the user clicked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1529205-clickedcolumn?language=objc
func (t_ TableView) ClickedColumn() int {
	rv := objc.Call[int](t_, objc.Sel("clickedColumn"))
	return rv
}

// The row size style (small, medium, large, or custom) used by the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534438-rowsizestyle?language=objc
func (t_ TableView) RowSizeStyle() TableViewRowSizeStyle {
	rv := objc.Call[TableViewRowSizeStyle](t_, objc.Sel("rowSizeStyle"))
	return rv
}

// The row size style (small, medium, large, or custom) used by the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534438-rowsizestyle?language=objc
func (t_ TableView) SetRowSizeStyle(value TableViewRowSizeStyle) {
	objc.Call[objc.Void](t_, objc.Sel("setRowSizeStyle:"), value)
}

// A Boolean value indicating whether the table view allows the user to type characters to select rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1526084-allowstypeselect?language=objc
func (t_ TableView) AllowsTypeSelect() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsTypeSelect"))
	return rv
}

// A Boolean value indicating whether the table view allows the user to type characters to select rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1526084-allowstypeselect?language=objc
func (t_ TableView) SetAllowsTypeSelect(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsTypeSelect:"), value)
}

// The table view’s sort descriptors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534198-sortdescriptors?language=objc
func (t_ TableView) SortDescriptors() []foundation.SortDescriptor {
	rv := objc.Call[[]foundation.SortDescriptor](t_, objc.Sel("sortDescriptors"))
	return rv
}

// The table view’s sort descriptors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534198-sortdescriptors?language=objc
func (t_ TableView) SetSortDescriptors(value []foundation.ISortDescriptor) {
	objc.Call[objc.Void](t_, objc.Sel("setSortDescriptors:"), value)
}

// The style that the table view uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/3622475-style?language=objc
func (t_ TableView) Style() TableViewStyle {
	rv := objc.Call[TableViewStyle](t_, objc.Sel("style"))
	return rv
}

// The style that the table view uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/3622475-style?language=objc
func (t_ TableView) SetStyle(value TableViewStyle) {
	objc.Call[objc.Void](t_, objc.Sel("setStyle:"), value)
}

// The view used to draw the area to the right of the column headers and above the vertical scroller of the enclosing scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1535831-cornerview?language=objc
func (t_ TableView) CornerView() View {
	rv := objc.Call[View](t_, objc.Sel("cornerView"))
	return rv
}

// The view used to draw the area to the right of the column headers and above the vertical scroller of the enclosing scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1535831-cornerview?language=objc
func (t_ TableView) SetCornerView(value IView) {
	objc.Call[objc.Void](t_, objc.Sel("setCornerView:"), objc.Ptr(value))
}

// The effective row size style for the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1531825-effectiverowsizestyle?language=objc
func (t_ TableView) EffectiveRowSizeStyle() TableViewRowSizeStyle {
	rv := objc.Call[TableViewRowSizeStyle](t_, objc.Sel("effectiveRowSizeStyle"))
	return rv
}

// The table view’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534325-delegate?language=objc
func (t_ TableView) Delegate() TableViewDelegateWrapper {
	rv := objc.Call[TableViewDelegateWrapper](t_, objc.Sel("delegate"))
	return rv
}

// The table view’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534325-delegate?language=objc
func (t_ TableView) SetDelegate(value PTableViewDelegate) {
	po0 := objc.WrapAsProtocol("NSTableViewDelegate", value)
	objc.SetAssociatedObject(t_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), po0)
}

// The table view’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534325-delegate?language=objc
func (t_ TableView) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The color used to draw the background of the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1527806-backgroundcolor?language=objc
func (t_ TableView) BackgroundColor() Color {
	rv := objc.Call[Color](t_, objc.Sel("backgroundColor"))
	return rv
}

// The color used to draw the background of the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1527806-backgroundcolor?language=objc
func (t_ TableView) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](t_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// The name under which table information is automatically saved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534409-autosavename?language=objc
func (t_ TableView) AutosaveName() TableViewAutosaveName {
	rv := objc.Call[TableViewAutosaveName](t_, objc.Sel("autosaveName"))
	return rv
}

// The name under which table information is automatically saved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1534409-autosavename?language=objc
func (t_ TableView) SetAutosaveName(value TableViewAutosaveName) {
	objc.Call[objc.Void](t_, objc.Sel("setAutosaveName:"), value)
}

// The effective style that the table uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/3622474-effectivestyle?language=objc
func (t_ TableView) EffectiveStyle() TableViewStyle {
	rv := objc.Call[TableViewStyle](t_, objc.Sel("effectiveStyle"))
	return rv
}

// A Boolean value that indicates whether the table view uses autolayout to calculate the height of rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/2870126-usesautomaticrowheights?language=objc
func (t_ TableView) UsesAutomaticRowHeights() bool {
	rv := objc.Call[bool](t_, objc.Sel("usesAutomaticRowHeights"))
	return rv
}

// A Boolean value that indicates whether the table view uses autolayout to calculate the height of rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/2870126-usesautomaticrowheights?language=objc
func (t_ TableView) SetUsesAutomaticRowHeights(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setUsesAutomaticRowHeights:"), value)
}

// A Boolean value indicating whether the table view uses alternating row colors for its background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1533967-usesalternatingrowbackgroundcolo?language=objc
func (t_ TableView) UsesAlternatingRowBackgroundColors() bool {
	rv := objc.Call[bool](t_, objc.Sel("usesAlternatingRowBackgroundColors"))
	return rv
}

// A Boolean value indicating whether the table view uses alternating row colors for its background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1533967-usesalternatingrowbackgroundcolo?language=objc
func (t_ TableView) SetUsesAlternatingRowBackgroundColors(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setUsesAlternatingRowBackgroundColors:"), value)
}

// The number of selected rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1527463-numberofselectedrows?language=objc
func (t_ TableView) NumberOfSelectedRows() int {
	rv := objc.Call[int](t_, objc.Sel("numberOfSelectedRows"))
	return rv
}

// The index of the row the user clicked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1527357-clickedrow?language=objc
func (t_ TableView) ClickedRow() int {
	rv := objc.Call[int](t_, objc.Sel("clickedRow"))
	return rv
}

// A Boolean value indicating whether the table view allows the user to select more than one column or row at a time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1532523-allowsmultipleselection?language=objc
func (t_ TableView) AllowsMultipleSelection() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsMultipleSelection"))
	return rv
}

// A Boolean value indicating whether the table view allows the user to select more than one column or row at a time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1532523-allowsmultipleselection?language=objc
func (t_ TableView) SetAllowsMultipleSelection(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsMultipleSelection:"), value)
}

// The index of the last selected column (or the last column added to the selection). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1532974-selectedcolumn?language=objc
func (t_ TableView) SelectedColumn() int {
	rv := objc.Call[int](t_, objc.Sel("selectedColumn"))
	return rv
}

// The grid lines drawn by the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1528689-gridstylemask?language=objc
func (t_ TableView) GridStyleMask() TableViewGridLineStyle {
	rv := objc.Call[TableViewGridLineStyle](t_, objc.Sel("gridStyleMask"))
	return rv
}

// The grid lines drawn by the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1528689-gridstylemask?language=objc
func (t_ TableView) SetGridStyleMask(value TableViewGridLineStyle) {
	objc.Call[objc.Void](t_, objc.Sel("setGridStyleMask:"), value)
}

// The index of the last selected row (or the last row added to the selection). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1535010-selectedrow?language=objc
func (t_ TableView) SelectedRow() int {
	rv := objc.Call[int](t_, objc.Sel("selectedRow"))
	return rv
}

// The selection highlight style used by the table view to indicate row and column selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1526311-selectionhighlightstyle?language=objc
func (t_ TableView) SelectionHighlightStyle() TableViewSelectionHighlightStyle {
	rv := objc.Call[TableViewSelectionHighlightStyle](t_, objc.Sel("selectionHighlightStyle"))
	return rv
}

// The selection highlight style used by the table view to indicate row and column selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1526311-selectionhighlightstyle?language=objc
func (t_ TableView) SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle) {
	objc.Call[objc.Void](t_, objc.Sel("setSelectionHighlightStyle:"), value)
}

// A Boolean value indicating whether a table row’s actions are visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1533198-rowactionsvisible?language=objc
func (t_ TableView) RowActionsVisible() bool {
	rv := objc.Call[bool](t_, objc.Sel("rowActionsVisible"))
	return rv
}

// A Boolean value indicating whether a table row’s actions are visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1533198-rowactionsvisible?language=objc
func (t_ TableView) SetRowActionsVisible(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setRowActionsVisible:"), value)
}

// The dictionary of all registered nib files for view-based table view identifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1530663-registerednibsbyidentifier?language=objc
func (t_ TableView) RegisteredNibsByIdentifier() map[UserInterfaceItemIdentifier]Nib {
	rv := objc.Call[map[UserInterfaceItemIdentifier]Nib](t_, objc.Sel("registeredNibsByIdentifier"))
	return rv
}

// An index set containing the indexes of the selected columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1524283-selectedcolumnindexes?language=objc
func (t_ TableView) SelectedColumnIndexes() foundation.IndexSet {
	rv := objc.Call[foundation.IndexSet](t_, objc.Sel("selectedColumnIndexes"))
	return rv
}

// A Boolean value indicating whether the table view allows the user to select columns by clicking their headers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1525276-allowscolumnselection?language=objc
func (t_ TableView) AllowsColumnSelection() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsColumnSelection"))
	return rv
}

// A Boolean value indicating whether the table view allows the user to select columns by clicking their headers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1525276-allowscolumnselection?language=objc
func (t_ TableView) SetAllowsColumnSelection(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsColumnSelection:"), value)
}

// The view object used to draw headers over columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1535880-headerview?language=objc
func (t_ TableView) HeaderView() TableHeaderView {
	rv := objc.Call[TableHeaderView](t_, objc.Sel("headerView"))
	return rv
}

// The view object used to draw headers over columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1535880-headerview?language=objc
func (t_ TableView) SetHeaderView(value ITableHeaderView) {
	objc.Call[objc.Void](t_, objc.Sel("setHeaderView:"), objc.Ptr(value))
}

// The number of rows in the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1527941-numberofrows?language=objc
func (t_ TableView) NumberOfRows() int {
	rv := objc.Call[int](t_, objc.Sel("numberOfRows"))
	return rv
}

// A Boolean value indicating whether the table view draws grouped rows as if they are floating. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1528624-floatsgrouprows?language=objc
func (t_ TableView) FloatsGroupRows() bool {
	rv := objc.Call[bool](t_, objc.Sel("floatsGroupRows"))
	return rv
}

// A Boolean value indicating whether the table view draws grouped rows as if they are floating. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1528624-floatsgrouprows?language=objc
func (t_ TableView) SetFloatsGroupRows(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setFloatsGroupRows:"), value)
}

// An index set containing the indexes of the selected rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1533844-selectedrowindexes?language=objc
func (t_ TableView) SelectedRowIndexes() foundation.IndexSet {
	rv := objc.Call[foundation.IndexSet](t_, objc.Sel("selectedRowIndexes"))
	return rv
}

// The horizontal and vertical spacing between cells. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1524258-intercellspacing?language=objc
func (t_ TableView) IntercellSpacing() foundation.Size {
	rv := objc.Call[foundation.Size](t_, objc.Sel("intercellSpacing"))
	return rv
}

// The horizontal and vertical spacing between cells. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1524258-intercellspacing?language=objc
func (t_ TableView) SetIntercellSpacing(value foundation.Size) {
	objc.Call[objc.Void](t_, objc.Sel("setIntercellSpacing:"), value)
}

// The number of columns in the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1528902-numberofcolumns?language=objc
func (t_ TableView) NumberOfColumns() int {
	rv := objc.Call[int](t_, objc.Sel("numberOfColumns"))
	return rv
}

// The table view’s column autoresizing style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1530784-columnautoresizingstyle?language=objc
func (t_ TableView) ColumnAutoresizingStyle() TableViewColumnAutoresizingStyle {
	rv := objc.Call[TableViewColumnAutoresizingStyle](t_, objc.Sel("columnAutoresizingStyle"))
	return rv
}

// The table view’s column autoresizing style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1530784-columnautoresizingstyle?language=objc
func (t_ TableView) SetColumnAutoresizingStyle(value TableViewColumnAutoresizingStyle) {
	objc.Call[objc.Void](t_, objc.Sel("setColumnAutoresizingStyle:"), value)
}

// The color used to draw grid lines. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1524620-gridcolor?language=objc
func (t_ TableView) GridColor() Color {
	rv := objc.Call[Color](t_, objc.Sel("gridColor"))
	return rv
}

// The color used to draw grid lines. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1524620-gridcolor?language=objc
func (t_ TableView) SetGridColor(value IColor) {
	objc.Call[objc.Void](t_, objc.Sel("setGridColor:"), objc.Ptr(value))
}

// An array containing the current table column objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1528735-tablecolumns?language=objc
func (t_ TableView) TableColumns() []TableColumn {
	rv := objc.Call[[]TableColumn](t_, objc.Sel("tableColumns"))
	return rv
}

// The feedback style displayed when the user drags over the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1527570-draggingdestinationfeedbackstyle?language=objc
func (t_ TableView) DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle {
	rv := objc.Call[TableViewDraggingDestinationFeedbackStyle](t_, objc.Sel("draggingDestinationFeedbackStyle"))
	return rv
}

// The feedback style displayed when the user drags over the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1527570-draggingdestinationfeedbackstyle?language=objc
func (t_ TableView) SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle) {
	objc.Call[objc.Void](t_, objc.Sel("setDraggingDestinationFeedbackStyle:"), value)
}

// The number of selected columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1524361-numberofselectedcolumns?language=objc
func (t_ TableView) NumberOfSelectedColumns() int {
	rv := objc.Call[int](t_, objc.Sel("numberOfSelectedColumns"))
	return rv
}
