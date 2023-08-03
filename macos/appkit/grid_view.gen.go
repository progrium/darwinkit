// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var GridViewClass = _GridViewClass{objc.GetClass("NSGridView")}

type _GridViewClass struct {
	objc.Class
}

type IGridView interface {
	IView
	IndexOfColumn(column IGridColumn) int
	RowAtIndex(index int) GridRow
	ColumnAtIndex(index int) GridColumn
	IndexOfRow(row IGridRow) int
	AddRowWithViews(views []IView) GridRow
	InsertRowAtIndexWithViews(index int, views []IView) GridRow
	RemoveRowAtIndex(index int)
	MoveRowAtIndexToIndex(fromIndex int, toIndex int)
	AddColumnWithViews(views []IView) GridColumn
	InsertColumnAtIndexWithViews(index int, views []IView) GridColumn
	RemoveColumnAtIndex(index int)
	MoveColumnAtIndexToIndex(fromIndex int, toIndex int)
	CellAtColumnIndexRowIndex(columnIndex int, rowIndex int) GridCell
	CellForView(view IView) GridCell
	MergeCellsInHorizontalRangeVerticalRange(hRange foundation.Range, vRange foundation.Range)
	NumberOfRows() int
	NumberOfColumns() int
	ColumnSpacing() float64
	SetColumnSpacing(value float64)
	RowSpacing() float64
	SetRowSpacing(value float64)
	RowAlignment() GridRowAlignment
	SetRowAlignment(value GridRowAlignment)
	XPlacement() GridCellPlacement
	SetXPlacement(value GridCellPlacement)
	YPlacement() GridCellPlacement
	SetYPlacement(value GridCellPlacement)
}

type GridView struct {
	View
}

func MakeGridView(ptr unsafe.Pointer) GridView {
	return GridView{
		View: MakeView(ptr),
	}
}

func (gc _GridViewClass) GridViewWithNumberOfColumnsRows(columnCount int, rowCount int) GridView {
	rv := objc.CallMethod[GridView](gc, objc.GetSelector("gridViewWithNumberOfColumns:rows:"), columnCount, rowCount)
	return rv
}

func GridView_GridViewWithNumberOfColumnsRows(columnCount int, rowCount int) GridView {
	return GridViewClass.GridViewWithNumberOfColumnsRows(columnCount, rowCount)
}

func (gc _GridViewClass) GridViewWithViews(rows [][]IView) GridView {
	rv := objc.CallMethod[GridView](gc, objc.GetSelector("gridViewWithViews:"), rows)
	return rv
}

func GridView_GridViewWithViews(rows [][]IView) GridView {
	return GridViewClass.GridViewWithViews(rows)
}

func (g_ GridView) InitWithFrame(frameRect foundation.Rect) GridView {
	rv := objc.CallMethod[GridView](g_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func GridView_InitWithFrame(frameRect foundation.Rect) GridView {
	return GridViewClass.Alloc().InitWithFrame(frameRect)
}

func (g_ GridView) Init() GridView {
	rv := objc.CallMethod[GridView](g_, objc.GetSelector("init"))
	return rv
}

func GridView_Init() GridView {
	return GridViewClass.Alloc().Init()
}

func (gc _GridViewClass) Alloc() GridView {
	rv := objc.CallMethod[GridView](gc, objc.GetSelector("alloc"))
	return rv
}

func GridView_Alloc() GridView {
	return GridViewClass.Alloc()
}

func (gc _GridViewClass) New() GridView {
	rv := objc.CallMethod[GridView](gc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewGridView() GridView {
	return GridViewClass.New()
}

func GridView_New() GridView {
	return GridViewClass.New()
}

func (g_ GridView) IndexOfColumn(column IGridColumn) int {
	rv := objc.CallMethod[int](g_, objc.GetSelector("indexOfColumn:"), objc.ExtractPtr(column))
	return rv
}

func (g_ GridView) RowAtIndex(index int) GridRow {
	rv := objc.CallMethod[GridRow](g_, objc.GetSelector("rowAtIndex:"), index)
	return rv
}

func (g_ GridView) ColumnAtIndex(index int) GridColumn {
	rv := objc.CallMethod[GridColumn](g_, objc.GetSelector("columnAtIndex:"), index)
	return rv
}

func (g_ GridView) IndexOfRow(row IGridRow) int {
	rv := objc.CallMethod[int](g_, objc.GetSelector("indexOfRow:"), objc.ExtractPtr(row))
	return rv
}

func (g_ GridView) AddRowWithViews(views []IView) GridRow {
	rv := objc.CallMethod[GridRow](g_, objc.GetSelector("addRowWithViews:"), views)
	return rv
}

func (g_ GridView) InsertRowAtIndexWithViews(index int, views []IView) GridRow {
	rv := objc.CallMethod[GridRow](g_, objc.GetSelector("insertRowAtIndex:withViews:"), index, views)
	return rv
}

func (g_ GridView) RemoveRowAtIndex(index int) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("removeRowAtIndex:"), index)
}

func (g_ GridView) MoveRowAtIndexToIndex(fromIndex int, toIndex int) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("moveRowAtIndex:toIndex:"), fromIndex, toIndex)
}

func (g_ GridView) AddColumnWithViews(views []IView) GridColumn {
	rv := objc.CallMethod[GridColumn](g_, objc.GetSelector("addColumnWithViews:"), views)
	return rv
}

func (g_ GridView) InsertColumnAtIndexWithViews(index int, views []IView) GridColumn {
	rv := objc.CallMethod[GridColumn](g_, objc.GetSelector("insertColumnAtIndex:withViews:"), index, views)
	return rv
}

func (g_ GridView) RemoveColumnAtIndex(index int) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("removeColumnAtIndex:"), index)
}

func (g_ GridView) MoveColumnAtIndexToIndex(fromIndex int, toIndex int) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("moveColumnAtIndex:toIndex:"), fromIndex, toIndex)
}

func (g_ GridView) CellAtColumnIndexRowIndex(columnIndex int, rowIndex int) GridCell {
	rv := objc.CallMethod[GridCell](g_, objc.GetSelector("cellAtColumnIndex:rowIndex:"), columnIndex, rowIndex)
	return rv
}

func (g_ GridView) CellForView(view IView) GridCell {
	rv := objc.CallMethod[GridCell](g_, objc.GetSelector("cellForView:"), objc.ExtractPtr(view))
	return rv
}

func (g_ GridView) MergeCellsInHorizontalRangeVerticalRange(hRange foundation.Range, vRange foundation.Range) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("mergeCellsInHorizontalRange:verticalRange:"), hRange, vRange)
}

func (g_ GridView) NumberOfRows() int {
	rv := objc.CallMethod[int](g_, objc.GetSelector("numberOfRows"))
	return rv
}

func (g_ GridView) NumberOfColumns() int {
	rv := objc.CallMethod[int](g_, objc.GetSelector("numberOfColumns"))
	return rv
}

func (g_ GridView) ColumnSpacing() float64 {
	rv := objc.CallMethod[float64](g_, objc.GetSelector("columnSpacing"))
	return rv
}

func (g_ GridView) SetColumnSpacing(value float64) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setColumnSpacing:"), value)
}

func (g_ GridView) RowSpacing() float64 {
	rv := objc.CallMethod[float64](g_, objc.GetSelector("rowSpacing"))
	return rv
}

func (g_ GridView) SetRowSpacing(value float64) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setRowSpacing:"), value)
}

func (g_ GridView) RowAlignment() GridRowAlignment {
	rv := objc.CallMethod[GridRowAlignment](g_, objc.GetSelector("rowAlignment"))
	return rv
}

func (g_ GridView) SetRowAlignment(value GridRowAlignment) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setRowAlignment:"), value)
}

func (g_ GridView) XPlacement() GridCellPlacement {
	rv := objc.CallMethod[GridCellPlacement](g_, objc.GetSelector("xPlacement"))
	return rv
}

func (g_ GridView) SetXPlacement(value GridCellPlacement) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setXPlacement:"), value)
}

func (g_ GridView) YPlacement() GridCellPlacement {
	rv := objc.CallMethod[GridCellPlacement](g_, objc.GetSelector("yPlacement"))
	return rv
}

func (g_ GridView) SetYPlacement(value GridCellPlacement) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setYPlacement:"), value)
}
