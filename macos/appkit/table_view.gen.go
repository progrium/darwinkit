// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TableViewClass = _TableViewClass{objc.GetClass("NSTableView")}

type _TableViewClass struct {
	objc.Class
}

type ITableView interface {
	IControl
	ReloadData()
	ReloadDataForRowIndexesColumnIndexes(rowIndexes foundation.IIndexSet, columnIndexes foundation.IIndexSet)
	MakeViewWithIdentifierOwner(identifier UserInterfaceItemIdentifier, owner objc.IObject) View
	RowViewAtRowMakeIfNecessary(row int, makeIfNecessary bool) TableRowView
	ViewAtColumnRowMakeIfNecessary(column int, row int, makeIfNecessary bool) View
	BeginUpdates()
	EndUpdates()
	MoveRowAtIndexToIndex(oldIndex int, newIndex int)
	InsertRowsAtIndexesWithAnimation(indexes foundation.IIndexSet, animationOptions TableViewAnimationOptions)
	RemoveRowsAtIndexesWithAnimation(indexes foundation.IIndexSet, animationOptions TableViewAnimationOptions)
	RowForView(view IView) int
	ColumnForView(view IView) int
	RegisterNibForIdentifier(nib INib, identifier UserInterfaceItemIdentifier)
	IndicatorImageInTableColumn(tableColumn ITableColumn) Image
	SetIndicatorImageInTableColumn(image IImage, tableColumn ITableColumn)
	AddTableColumn(tableColumn ITableColumn)
	RemoveTableColumn(tableColumn ITableColumn)
	MoveColumnToColumn(oldIndex int, newIndex int)
	ColumnWithIdentifier(identifier UserInterfaceItemIdentifier) int
	TableColumnWithIdentifier(identifier UserInterfaceItemIdentifier) TableColumn
	SelectColumnIndexesByExtendingSelection(indexes foundation.IIndexSet, extend bool)
	DeselectColumn(column int)
	IsColumnSelected(column int) bool
	SelectRowIndexesByExtendingSelection(indexes foundation.IIndexSet, extend bool)
	DeselectRow(row int)
	IsRowSelected(row int) bool
	SelectAll(sender objc.IObject)
	DeselectAll(sender objc.IObject)
	EnumerateAvailableRowViewsUsingBlock(handler func(rowView TableRowView, row int))
	EditColumnRowWithEventSelect(column int, row int, event IEvent, select_ bool)
	DidAddRowViewForRow(rowView ITableRowView, row int)
	DidRemoveRowViewForRow(rowView ITableRowView, row int)
	RectOfColumn(column int) foundation.Rect
	RectOfRow(row int) foundation.Rect
	RowsInRect(rect foundation.Rect) foundation.Range
	ColumnIndexesInRect(rect foundation.Rect) foundation.IndexSet
	ColumnAtPoint(point foundation.Point) int
	RowAtPoint(point foundation.Point) int
	FrameOfCellAtColumnRow(column int, row int) foundation.Rect
	SizeLastColumnToFit()
	NoteNumberOfRowsChanged()
	Tile()
	NoteHeightOfRowsWithIndexesChanged(indexSet foundation.IIndexSet)
	DrawRowClipRect(row int, clipRect foundation.Rect)
	DrawGridInClipRect(clipRect foundation.Rect)
	HighlightSelectionInClipRect(clipRect foundation.Rect)
	DrawBackgroundInClipRect(clipRect foundation.Rect)
	ScrollRowToVisible(row int)
	ScrollColumnToVisible(column int)
	DragImageForRowsWithIndexesTableColumnsEventOffset(dragRows foundation.IIndexSet, tableColumns []ITableColumn, dragEvent IEvent, dragImageOffset *foundation.Point) Image
	CanDragRowsWithIndexesAtPoint(rowIndexes foundation.IIndexSet, mouseDownPoint foundation.Point) bool
	SetDraggingSourceOperationMaskForLocal(mask DragOperation, isLocal bool)
	SetDropRowDropOperation(row int, dropOperation TableViewDropOperation)
	HideRowsAtIndexesWithAnimation(indexes foundation.IIndexSet, rowAnimation TableViewAnimationOptions)
	UnhideRowsAtIndexesWithAnimation(indexes foundation.IIndexSet, rowAnimation TableViewAnimationOptions)
	DataSource() TableViewDataSourceWrapper
	SetDataSource(value ITableViewDataSource)
	SetDataSource0(value objc.IObject)
	UsesStaticContents() bool
	SetUsesStaticContents(value bool)
	RegisteredNibsByIdentifier() map[UserInterfaceItemIdentifier]Nib
	DoubleAction() objc.Selector
	SetDoubleAction(value objc.Selector)
	ClickedColumn() int
	ClickedRow() int
	AllowsColumnReordering() bool
	SetAllowsColumnReordering(value bool)
	AllowsColumnResizing() bool
	SetAllowsColumnResizing(value bool)
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	AllowsColumnSelection() bool
	SetAllowsColumnSelection(value bool)
	UsesAutomaticRowHeights() bool
	SetUsesAutomaticRowHeights(value bool)
	IntercellSpacing() foundation.Size
	SetIntercellSpacing(value foundation.Size)
	RowHeight() float64
	SetRowHeight(value float64)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	UsesAlternatingRowBackgroundColors() bool
	SetUsesAlternatingRowBackgroundColors(value bool)
	Style() TableViewStyle
	SetStyle(value TableViewStyle)
	EffectiveStyle() TableViewStyle
	SelectionHighlightStyle() TableViewSelectionHighlightStyle
	SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle)
	GridColor() Color
	SetGridColor(value IColor)
	GridStyleMask() TableViewGridLineStyle
	SetGridStyleMask(value TableViewGridLineStyle)
	EffectiveRowSizeStyle() TableViewRowSizeStyle
	RowSizeStyle() TableViewRowSizeStyle
	SetRowSizeStyle(value TableViewRowSizeStyle)
	TableColumns() []TableColumn
	SelectedColumn() int
	SelectedColumnIndexes() foundation.IndexSet
	NumberOfSelectedColumns() int
	SelectedRow() int
	SelectedRowIndexes() foundation.IndexSet
	NumberOfSelectedRows() int
	AllowsTypeSelect() bool
	SetAllowsTypeSelect(value bool)
	NumberOfColumns() int
	NumberOfRows() int
	FloatsGroupRows() bool
	SetFloatsGroupRows(value bool)
	EditedColumn() int
	EditedRow() int
	HeaderView() TableHeaderView
	SetHeaderView(value ITableHeaderView)
	CornerView() View
	SetCornerView(value IView)
	ColumnAutoresizingStyle() TableViewColumnAutoresizingStyle
	SetColumnAutoresizingStyle(value TableViewColumnAutoresizingStyle)
	AutosaveTableColumns() bool
	SetAutosaveTableColumns(value bool)
	AutosaveName() TableViewAutosaveName
	SetAutosaveName(value TableViewAutosaveName)
	Delegate() TableViewDelegateWrapper
	SetDelegate(value ITableViewDelegate)
	SetDelegate0(value objc.IObject)
	HighlightedTableColumn() TableColumn
	SetHighlightedTableColumn(value ITableColumn)
	VerticalMotionCanBeginDrag() bool
	SetVerticalMotionCanBeginDrag(value bool)
	DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle
	SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle)
	SortDescriptors() []foundation.SortDescriptor
	SetSortDescriptors(value []foundation.ISortDescriptor)
	RowActionsVisible() bool
	SetRowActionsVisible(value bool)
	HiddenRowIndexes() foundation.IndexSet
}

type TableView struct {
	Control
}

func MakeTableView(ptr unsafe.Pointer) TableView {
	return TableView{
		Control: MakeControl(ptr),
	}
}

func (t_ TableView) InitWithFrame(frameRect foundation.Rect) TableView {
	rv := objc.CallMethod[TableView](t_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func TableView_InitWithFrame(frameRect foundation.Rect) TableView {
	return TableViewClass.Alloc().InitWithFrame(frameRect)
}

func (t_ TableView) Init() TableView {
	rv := objc.CallMethod[TableView](t_, objc.GetSelector("init"))
	return rv
}

func TableView_Init() TableView {
	return TableViewClass.Alloc().Init()
}

func (tc _TableViewClass) Alloc() TableView {
	rv := objc.CallMethod[TableView](tc, objc.GetSelector("alloc"))
	return rv
}

func TableView_Alloc() TableView {
	return TableViewClass.Alloc()
}

func (tc _TableViewClass) New() TableView {
	rv := objc.CallMethod[TableView](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTableView() TableView {
	return TableViewClass.New()
}

func TableView_New() TableView {
	return TableViewClass.New()
}

func (t_ TableView) ReloadData() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("reloadData"))
}

func (t_ TableView) ReloadDataForRowIndexesColumnIndexes(rowIndexes foundation.IIndexSet, columnIndexes foundation.IIndexSet) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("reloadDataForRowIndexes:columnIndexes:"), objc.ExtractPtr(rowIndexes), objc.ExtractPtr(columnIndexes))
}

func (t_ TableView) MakeViewWithIdentifierOwner(identifier UserInterfaceItemIdentifier, owner objc.IObject) View {
	rv := objc.CallMethod[View](t_, objc.GetSelector("makeViewWithIdentifier:owner:"), identifier, objc.ExtractPtr(owner))
	return rv
}

func (t_ TableView) RowViewAtRowMakeIfNecessary(row int, makeIfNecessary bool) TableRowView {
	rv := objc.CallMethod[TableRowView](t_, objc.GetSelector("rowViewAtRow:makeIfNecessary:"), row, makeIfNecessary)
	return rv
}

func (t_ TableView) ViewAtColumnRowMakeIfNecessary(column int, row int, makeIfNecessary bool) View {
	rv := objc.CallMethod[View](t_, objc.GetSelector("viewAtColumn:row:makeIfNecessary:"), column, row, makeIfNecessary)
	return rv
}

func (t_ TableView) BeginUpdates() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("beginUpdates"))
}

func (t_ TableView) EndUpdates() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("endUpdates"))
}

func (t_ TableView) MoveRowAtIndexToIndex(oldIndex int, newIndex int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("moveRowAtIndex:toIndex:"), oldIndex, newIndex)
}

func (t_ TableView) InsertRowsAtIndexesWithAnimation(indexes foundation.IIndexSet, animationOptions TableViewAnimationOptions) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("insertRowsAtIndexes:withAnimation:"), objc.ExtractPtr(indexes), animationOptions)
}

func (t_ TableView) RemoveRowsAtIndexesWithAnimation(indexes foundation.IIndexSet, animationOptions TableViewAnimationOptions) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("removeRowsAtIndexes:withAnimation:"), objc.ExtractPtr(indexes), animationOptions)
}

func (t_ TableView) RowForView(view IView) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("rowForView:"), objc.ExtractPtr(view))
	return rv
}

func (t_ TableView) ColumnForView(view IView) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("columnForView:"), objc.ExtractPtr(view))
	return rv
}

func (t_ TableView) RegisterNibForIdentifier(nib INib, identifier UserInterfaceItemIdentifier) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("registerNib:forIdentifier:"), objc.ExtractPtr(nib), identifier)
}

func (t_ TableView) IndicatorImageInTableColumn(tableColumn ITableColumn) Image {
	rv := objc.CallMethod[Image](t_, objc.GetSelector("indicatorImageInTableColumn:"), objc.ExtractPtr(tableColumn))
	return rv
}

func (t_ TableView) SetIndicatorImageInTableColumn(image IImage, tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setIndicatorImage:inTableColumn:"), objc.ExtractPtr(image), objc.ExtractPtr(tableColumn))
}

func (t_ TableView) AddTableColumn(tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("addTableColumn:"), objc.ExtractPtr(tableColumn))
}

func (t_ TableView) RemoveTableColumn(tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("removeTableColumn:"), objc.ExtractPtr(tableColumn))
}

func (t_ TableView) MoveColumnToColumn(oldIndex int, newIndex int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("moveColumn:toColumn:"), oldIndex, newIndex)
}

func (t_ TableView) ColumnWithIdentifier(identifier UserInterfaceItemIdentifier) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("columnWithIdentifier:"), identifier)
	return rv
}

func (t_ TableView) TableColumnWithIdentifier(identifier UserInterfaceItemIdentifier) TableColumn {
	rv := objc.CallMethod[TableColumn](t_, objc.GetSelector("tableColumnWithIdentifier:"), identifier)
	return rv
}

func (t_ TableView) SelectColumnIndexesByExtendingSelection(indexes foundation.IIndexSet, extend bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("selectColumnIndexes:byExtendingSelection:"), objc.ExtractPtr(indexes), extend)
}

func (t_ TableView) DeselectColumn(column int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("deselectColumn:"), column)
}

func (t_ TableView) IsColumnSelected(column int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isColumnSelected:"), column)
	return rv
}

func (t_ TableView) SelectRowIndexesByExtendingSelection(indexes foundation.IIndexSet, extend bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("selectRowIndexes:byExtendingSelection:"), objc.ExtractPtr(indexes), extend)
}

func (t_ TableView) DeselectRow(row int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("deselectRow:"), row)
}

func (t_ TableView) IsRowSelected(row int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isRowSelected:"), row)
	return rv
}

func (t_ TableView) SelectAll(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("selectAll:"), objc.ExtractPtr(sender))
}

func (t_ TableView) DeselectAll(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("deselectAll:"), objc.ExtractPtr(sender))
}

func (t_ TableView) EnumerateAvailableRowViewsUsingBlock(handler func(rowView TableRowView, row int)) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("enumerateAvailableRowViewsUsingBlock:"), handler)
}

func (t_ TableView) EditColumnRowWithEventSelect(column int, row int, event IEvent, select_ bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("editColumn:row:withEvent:select:"), column, row, objc.ExtractPtr(event), select_)
}

func (t_ TableView) DidAddRowViewForRow(rowView ITableRowView, row int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("didAddRowView:forRow:"), objc.ExtractPtr(rowView), row)
}

func (t_ TableView) DidRemoveRowViewForRow(rowView ITableRowView, row int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("didRemoveRowView:forRow:"), objc.ExtractPtr(rowView), row)
}

func (t_ TableView) RectOfColumn(column int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.GetSelector("rectOfColumn:"), column)
	return rv
}

func (t_ TableView) RectOfRow(row int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.GetSelector("rectOfRow:"), row)
	return rv
}

func (t_ TableView) RowsInRect(rect foundation.Rect) foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("rowsInRect:"), rect)
	return rv
}

func (t_ TableView) ColumnIndexesInRect(rect foundation.Rect) foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](t_, objc.GetSelector("columnIndexesInRect:"), rect)
	return rv
}

func (t_ TableView) ColumnAtPoint(point foundation.Point) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("columnAtPoint:"), point)
	return rv
}

func (t_ TableView) RowAtPoint(point foundation.Point) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("rowAtPoint:"), point)
	return rv
}

func (t_ TableView) FrameOfCellAtColumnRow(column int, row int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.GetSelector("frameOfCellAtColumn:row:"), column, row)
	return rv
}

func (t_ TableView) SizeLastColumnToFit() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("sizeLastColumnToFit"))
}

func (t_ TableView) NoteNumberOfRowsChanged() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("noteNumberOfRowsChanged"))
}

func (t_ TableView) Tile() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tile"))
}

func (t_ TableView) NoteHeightOfRowsWithIndexesChanged(indexSet foundation.IIndexSet) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("noteHeightOfRowsWithIndexesChanged:"), objc.ExtractPtr(indexSet))
}

func (t_ TableView) DrawRowClipRect(row int, clipRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("drawRow:clipRect:"), row, clipRect)
}

func (t_ TableView) DrawGridInClipRect(clipRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("drawGridInClipRect:"), clipRect)
}

func (t_ TableView) HighlightSelectionInClipRect(clipRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("highlightSelectionInClipRect:"), clipRect)
}

func (t_ TableView) DrawBackgroundInClipRect(clipRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("drawBackgroundInClipRect:"), clipRect)
}

func (t_ TableView) ScrollRowToVisible(row int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("scrollRowToVisible:"), row)
}

func (t_ TableView) ScrollColumnToVisible(column int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("scrollColumnToVisible:"), column)
}

func (t_ TableView) DragImageForRowsWithIndexesTableColumnsEventOffset(dragRows foundation.IIndexSet, tableColumns []ITableColumn, dragEvent IEvent, dragImageOffset *foundation.Point) Image {
	rv := objc.CallMethod[Image](t_, objc.GetSelector("dragImageForRowsWithIndexes:tableColumns:event:offset:"), objc.ExtractPtr(dragRows), tableColumns, objc.ExtractPtr(dragEvent), dragImageOffset)
	return rv
}

func (t_ TableView) CanDragRowsWithIndexesAtPoint(rowIndexes foundation.IIndexSet, mouseDownPoint foundation.Point) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("canDragRowsWithIndexes:atPoint:"), objc.ExtractPtr(rowIndexes), mouseDownPoint)
	return rv
}

func (t_ TableView) SetDraggingSourceOperationMaskForLocal(mask DragOperation, isLocal bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDraggingSourceOperationMask:forLocal:"), mask, isLocal)
}

func (t_ TableView) SetDropRowDropOperation(row int, dropOperation TableViewDropOperation) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDropRow:dropOperation:"), row, dropOperation)
}

func (t_ TableView) HideRowsAtIndexesWithAnimation(indexes foundation.IIndexSet, rowAnimation TableViewAnimationOptions) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("hideRowsAtIndexes:withAnimation:"), objc.ExtractPtr(indexes), rowAnimation)
}

func (t_ TableView) UnhideRowsAtIndexesWithAnimation(indexes foundation.IIndexSet, rowAnimation TableViewAnimationOptions) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("unhideRowsAtIndexes:withAnimation:"), objc.ExtractPtr(indexes), rowAnimation)
}

func (t_ TableView) DataSource() TableViewDataSourceWrapper {
	rv := objc.CallMethod[TableViewDataSourceWrapper](t_, objc.GetSelector("dataSource"))
	return rv
}

func (t_ TableView) SetDataSource(value ITableViewDataSource) {
	po := objc.WrapAsProtocol("NSTableViewDataSource", value)
	objc.SetAssociatedObject(t_, objc.AssociationKey("setDataSource"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDataSource:"), po)
}

func (t_ TableView) SetDataSource0(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDataSource:"), objc.ExtractPtr(value))
}

func (t_ TableView) UsesStaticContents() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("usesStaticContents"))
	return rv
}

func (t_ TableView) SetUsesStaticContents(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setUsesStaticContents:"), value)
}

func (t_ TableView) RegisteredNibsByIdentifier() map[UserInterfaceItemIdentifier]Nib {
	rv := objc.CallMethod[map[UserInterfaceItemIdentifier]Nib](t_, objc.GetSelector("registeredNibsByIdentifier"))
	return rv
}

func (t_ TableView) DoubleAction() objc.Selector {
	rv := objc.CallMethod[objc.Selector](t_, objc.GetSelector("doubleAction"))
	return rv
}

func (t_ TableView) SetDoubleAction(value objc.Selector) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDoubleAction:"), value)
}

func (t_ TableView) ClickedColumn() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("clickedColumn"))
	return rv
}

func (t_ TableView) ClickedRow() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("clickedRow"))
	return rv
}

func (t_ TableView) AllowsColumnReordering() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsColumnReordering"))
	return rv
}

func (t_ TableView) SetAllowsColumnReordering(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsColumnReordering:"), value)
}

func (t_ TableView) AllowsColumnResizing() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsColumnResizing"))
	return rv
}

func (t_ TableView) SetAllowsColumnResizing(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsColumnResizing:"), value)
}

func (t_ TableView) AllowsMultipleSelection() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsMultipleSelection"))
	return rv
}

func (t_ TableView) SetAllowsMultipleSelection(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsMultipleSelection:"), value)
}

func (t_ TableView) AllowsEmptySelection() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsEmptySelection"))
	return rv
}

func (t_ TableView) SetAllowsEmptySelection(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsEmptySelection:"), value)
}

func (t_ TableView) AllowsColumnSelection() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsColumnSelection"))
	return rv
}

func (t_ TableView) SetAllowsColumnSelection(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsColumnSelection:"), value)
}

func (t_ TableView) UsesAutomaticRowHeights() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("usesAutomaticRowHeights"))
	return rv
}

func (t_ TableView) SetUsesAutomaticRowHeights(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setUsesAutomaticRowHeights:"), value)
}

func (t_ TableView) IntercellSpacing() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.GetSelector("intercellSpacing"))
	return rv
}

func (t_ TableView) SetIntercellSpacing(value foundation.Size) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setIntercellSpacing:"), value)
}

func (t_ TableView) RowHeight() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("rowHeight"))
	return rv
}

func (t_ TableView) SetRowHeight(value float64) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setRowHeight:"), value)
}

func (t_ TableView) BackgroundColor() Color {
	rv := objc.CallMethod[Color](t_, objc.GetSelector("backgroundColor"))
	return rv
}

func (t_ TableView) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (t_ TableView) UsesAlternatingRowBackgroundColors() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("usesAlternatingRowBackgroundColors"))
	return rv
}

func (t_ TableView) SetUsesAlternatingRowBackgroundColors(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setUsesAlternatingRowBackgroundColors:"), value)
}

func (t_ TableView) Style() TableViewStyle {
	rv := objc.CallMethod[TableViewStyle](t_, objc.GetSelector("style"))
	return rv
}

func (t_ TableView) SetStyle(value TableViewStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setStyle:"), value)
}

func (t_ TableView) EffectiveStyle() TableViewStyle {
	rv := objc.CallMethod[TableViewStyle](t_, objc.GetSelector("effectiveStyle"))
	return rv
}

func (t_ TableView) SelectionHighlightStyle() TableViewSelectionHighlightStyle {
	rv := objc.CallMethod[TableViewSelectionHighlightStyle](t_, objc.GetSelector("selectionHighlightStyle"))
	return rv
}

func (t_ TableView) SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSelectionHighlightStyle:"), value)
}

func (t_ TableView) GridColor() Color {
	rv := objc.CallMethod[Color](t_, objc.GetSelector("gridColor"))
	return rv
}

func (t_ TableView) SetGridColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setGridColor:"), objc.ExtractPtr(value))
}

func (t_ TableView) GridStyleMask() TableViewGridLineStyle {
	rv := objc.CallMethod[TableViewGridLineStyle](t_, objc.GetSelector("gridStyleMask"))
	return rv
}

func (t_ TableView) SetGridStyleMask(value TableViewGridLineStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setGridStyleMask:"), value)
}

func (t_ TableView) EffectiveRowSizeStyle() TableViewRowSizeStyle {
	rv := objc.CallMethod[TableViewRowSizeStyle](t_, objc.GetSelector("effectiveRowSizeStyle"))
	return rv
}

func (t_ TableView) RowSizeStyle() TableViewRowSizeStyle {
	rv := objc.CallMethod[TableViewRowSizeStyle](t_, objc.GetSelector("rowSizeStyle"))
	return rv
}

func (t_ TableView) SetRowSizeStyle(value TableViewRowSizeStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setRowSizeStyle:"), value)
}

func (t_ TableView) TableColumns() []TableColumn {
	rv := objc.CallMethod[[]TableColumn](t_, objc.GetSelector("tableColumns"))
	return rv
}

func (t_ TableView) SelectedColumn() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("selectedColumn"))
	return rv
}

func (t_ TableView) SelectedColumnIndexes() foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](t_, objc.GetSelector("selectedColumnIndexes"))
	return rv
}

func (t_ TableView) NumberOfSelectedColumns() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("numberOfSelectedColumns"))
	return rv
}

func (t_ TableView) SelectedRow() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("selectedRow"))
	return rv
}

func (t_ TableView) SelectedRowIndexes() foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](t_, objc.GetSelector("selectedRowIndexes"))
	return rv
}

func (t_ TableView) NumberOfSelectedRows() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("numberOfSelectedRows"))
	return rv
}

func (t_ TableView) AllowsTypeSelect() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsTypeSelect"))
	return rv
}

func (t_ TableView) SetAllowsTypeSelect(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsTypeSelect:"), value)
}

func (t_ TableView) NumberOfColumns() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("numberOfColumns"))
	return rv
}

func (t_ TableView) NumberOfRows() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("numberOfRows"))
	return rv
}

func (t_ TableView) FloatsGroupRows() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("floatsGroupRows"))
	return rv
}

func (t_ TableView) SetFloatsGroupRows(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setFloatsGroupRows:"), value)
}

func (t_ TableView) EditedColumn() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("editedColumn"))
	return rv
}

func (t_ TableView) EditedRow() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("editedRow"))
	return rv
}

func (t_ TableView) HeaderView() TableHeaderView {
	rv := objc.CallMethod[TableHeaderView](t_, objc.GetSelector("headerView"))
	return rv
}

func (t_ TableView) SetHeaderView(value ITableHeaderView) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setHeaderView:"), objc.ExtractPtr(value))
}

func (t_ TableView) CornerView() View {
	rv := objc.CallMethod[View](t_, objc.GetSelector("cornerView"))
	return rv
}

func (t_ TableView) SetCornerView(value IView) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setCornerView:"), objc.ExtractPtr(value))
}

func (t_ TableView) ColumnAutoresizingStyle() TableViewColumnAutoresizingStyle {
	rv := objc.CallMethod[TableViewColumnAutoresizingStyle](t_, objc.GetSelector("columnAutoresizingStyle"))
	return rv
}

func (t_ TableView) SetColumnAutoresizingStyle(value TableViewColumnAutoresizingStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setColumnAutoresizingStyle:"), value)
}

func (t_ TableView) AutosaveTableColumns() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("autosaveTableColumns"))
	return rv
}

func (t_ TableView) SetAutosaveTableColumns(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAutosaveTableColumns:"), value)
}

func (t_ TableView) AutosaveName() TableViewAutosaveName {
	rv := objc.CallMethod[TableViewAutosaveName](t_, objc.GetSelector("autosaveName"))
	return rv
}

func (t_ TableView) SetAutosaveName(value TableViewAutosaveName) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAutosaveName:"), value)
}

func (t_ TableView) Delegate() TableViewDelegateWrapper {
	rv := objc.CallMethod[TableViewDelegateWrapper](t_, objc.GetSelector("delegate"))
	return rv
}

func (t_ TableView) SetDelegate(value ITableViewDelegate) {
	po := objc.WrapAsProtocol("NSTableViewDelegate", value)
	objc.SetAssociatedObject(t_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDelegate:"), po)
}

func (t_ TableView) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (t_ TableView) HighlightedTableColumn() TableColumn {
	rv := objc.CallMethod[TableColumn](t_, objc.GetSelector("highlightedTableColumn"))
	return rv
}

func (t_ TableView) SetHighlightedTableColumn(value ITableColumn) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setHighlightedTableColumn:"), objc.ExtractPtr(value))
}

func (t_ TableView) VerticalMotionCanBeginDrag() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("verticalMotionCanBeginDrag"))
	return rv
}

func (t_ TableView) SetVerticalMotionCanBeginDrag(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setVerticalMotionCanBeginDrag:"), value)
}

func (t_ TableView) DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle {
	rv := objc.CallMethod[TableViewDraggingDestinationFeedbackStyle](t_, objc.GetSelector("draggingDestinationFeedbackStyle"))
	return rv
}

func (t_ TableView) SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDraggingDestinationFeedbackStyle:"), value)
}

func (t_ TableView) SortDescriptors() []foundation.SortDescriptor {
	rv := objc.CallMethod[[]foundation.SortDescriptor](t_, objc.GetSelector("sortDescriptors"))
	return rv
}

func (t_ TableView) SetSortDescriptors(value []foundation.ISortDescriptor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSortDescriptors:"), value)
}

func (t_ TableView) RowActionsVisible() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("rowActionsVisible"))
	return rv
}

func (t_ TableView) SetRowActionsVisible(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setRowActionsVisible:"), value)
}

func (t_ TableView) HiddenRowIndexes() foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](t_, objc.GetSelector("hiddenRowIndexes"))
	return rv
}
