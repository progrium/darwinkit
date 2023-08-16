// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TableHeaderView] class.
var TableHeaderViewClass = _TableHeaderViewClass{objc.GetClass("NSTableHeaderView")}

type _TableHeaderViewClass struct {
	objc.Class
}

// An interface definition for the [TableHeaderView] class.
type ITableHeaderView interface {
	IView
	ColumnAtPoint(point foundation.Point) int
	HeaderRectOfColumn(column int) foundation.Rect
	DraggedDistance() float64
	DraggedColumn() int
	ResizedColumn() int
	TableView() TableView
	SetTableView(value ITableView)
}

// An object that draws headers over a table view's columns and handles mouse events in those headers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableheaderview?language=objc
type TableHeaderView struct {
	View
}

func TableHeaderViewFrom(ptr unsafe.Pointer) TableHeaderView {
	return TableHeaderView{
		View: ViewFrom(ptr),
	}
}

func (tc _TableHeaderViewClass) Alloc() TableHeaderView {
	rv := objc.Call[TableHeaderView](tc, objc.Sel("alloc"))
	return rv
}

func TableHeaderView_Alloc() TableHeaderView {
	return TableHeaderViewClass.Alloc()
}

func (tc _TableHeaderViewClass) New() TableHeaderView {
	rv := objc.Call[TableHeaderView](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTableHeaderView() TableHeaderView {
	return TableHeaderViewClass.New()
}

func (t_ TableHeaderView) Init() TableHeaderView {
	rv := objc.Call[TableHeaderView](t_, objc.Sel("init"))
	return rv
}

func (t_ TableHeaderView) InitWithFrame(frameRect foundation.Rect) TableHeaderView {
	rv := objc.Call[TableHeaderView](t_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func TableHeaderView_InitWithFrame(frameRect foundation.Rect) TableHeaderView {
	return TableHeaderViewClass.Alloc().InitWithFrame(frameRect)
}

// Returns the index of the column whose header lies under aPoint in the receiver, or –1 if no such column is found. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableheaderview/1529216-columnatpoint?language=objc
func (t_ TableHeaderView) ColumnAtPoint(point foundation.Point) int {
	rv := objc.Call[int](t_, objc.Sel("columnAtPoint:"), point)
	return rv
}

// Returns the rectangle containing the header tile for the column at columnIndex. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableheaderview/1531883-headerrectofcolumn?language=objc
func (t_ TableHeaderView) HeaderRectOfColumn(column int) foundation.Rect {
	rv := objc.Call[foundation.Rect](t_, objc.Sel("headerRectOfColumn:"), column)
	return rv
}

// The horizontal distance that the user has dragged a column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableheaderview/1527836-draggeddistance?language=objc
func (t_ TableHeaderView) DraggedDistance() float64 {
	rv := objc.Call[float64](t_, objc.Sel("draggedDistance"))
	return rv
}

// The index of the column that the user is dragging. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableheaderview/1534458-draggedcolumn?language=objc
func (t_ TableHeaderView) DraggedColumn() int {
	rv := objc.Call[int](t_, objc.Sel("draggedColumn"))
	return rv
}

// The index of the column that the user is resizing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableheaderview/1528247-resizedcolumn?language=objc
func (t_ TableHeaderView) ResizedColumn() int {
	rv := objc.Call[int](t_, objc.Sel("resizedColumn"))
	return rv
}

// The NSTableView instance that this table header view belongs to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableheaderview/1535730-tableview?language=objc
func (t_ TableHeaderView) TableView() TableView {
	rv := objc.Call[TableView](t_, objc.Sel("tableView"))
	return rv
}

// The NSTableView instance that this table header view belongs to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableheaderview/1535730-tableview?language=objc
func (t_ TableHeaderView) SetTableView(value ITableView) {
	objc.Call[objc.Void](t_, objc.Sel("setTableView:"), objc.Ptr(value))
}
