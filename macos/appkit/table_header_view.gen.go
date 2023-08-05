// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TableHeaderViewClass = _TableHeaderViewClass{objc.GetClass("NSTableHeaderView")}

type _TableHeaderViewClass struct {
	objc.Class
}

type ITableHeaderView interface {
	IView
	ColumnAtPoint(point foundation.Point) int
	HeaderRectOfColumn(column int) foundation.Rect
	TableView() TableView
	SetTableView(value ITableView)
	DraggedColumn() int
	DraggedDistance() float64
	ResizedColumn() int
}

type TableHeaderView struct {
	View
}

func MakeTableHeaderView(ptr unsafe.Pointer) TableHeaderView {
	return TableHeaderView{
		View: MakeView(ptr),
	}
}

func (t_ TableHeaderView) InitWithFrame(frameRect foundation.Rect) TableHeaderView {
	rv := objc.CallMethod[TableHeaderView](t_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func TableHeaderView_InitWithFrame(frameRect foundation.Rect) TableHeaderView {
	return TableHeaderViewClass.Alloc().InitWithFrame(frameRect)
}

func (t_ TableHeaderView) Init() TableHeaderView {
	rv := objc.CallMethod[TableHeaderView](t_, objc.GetSelector("init"))
	return rv
}

func TableHeaderView_Init() TableHeaderView {
	return TableHeaderViewClass.Alloc().Init()
}

func (tc _TableHeaderViewClass) Alloc() TableHeaderView {
	rv := objc.CallMethod[TableHeaderView](tc, objc.GetSelector("alloc"))
	return rv
}

func TableHeaderView_Alloc() TableHeaderView {
	return TableHeaderViewClass.Alloc()
}

func (tc _TableHeaderViewClass) New() TableHeaderView {
	rv := objc.CallMethod[TableHeaderView](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTableHeaderView() TableHeaderView {
	return TableHeaderViewClass.New()
}

func TableHeaderView_New() TableHeaderView {
	return TableHeaderViewClass.New()
}

func (t_ TableHeaderView) ColumnAtPoint(point foundation.Point) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("columnAtPoint:"), point)
	return rv
}

func (t_ TableHeaderView) HeaderRectOfColumn(column int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.GetSelector("headerRectOfColumn:"), column)
	return rv
}

func (t_ TableHeaderView) TableView() TableView {
	rv := objc.CallMethod[TableView](t_, objc.GetSelector("tableView"))
	return rv
}

func (t_ TableHeaderView) SetTableView(value ITableView) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTableView:"), objc.ExtractPtr(value))
}

func (t_ TableHeaderView) DraggedColumn() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("draggedColumn"))
	return rv
}

func (t_ TableHeaderView) DraggedDistance() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("draggedDistance"))
	return rv
}

func (t_ TableHeaderView) ResizedColumn() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("resizedColumn"))
	return rv
}
