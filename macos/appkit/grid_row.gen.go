// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var GridRowClass = _GridRowClass{objc.GetClass("NSGridRow")}

type _GridRowClass struct {
	objc.Class
}

type IGridRow interface {
	objc.IObject
	CellAtIndex(index int) GridCell
	MergeCellsInRange(range_ foundation.Range)
	NumberOfCells() int
	IsHidden() bool
	SetHidden(value bool)
	TopPadding() float64
	SetTopPadding(value float64)
	BottomPadding() float64
	SetBottomPadding(value float64)
	Height() float64
	SetHeight(value float64)
	RowAlignment() GridRowAlignment
	SetRowAlignment(value GridRowAlignment)
	YPlacement() GridCellPlacement
	SetYPlacement(value GridCellPlacement)
	GridView() GridView
}

type GridRow struct {
	objc.Object
}

func MakeGridRow(ptr unsafe.Pointer) GridRow {
	return GridRow{
		Object: objc.MakeObject(ptr),
	}
}

func (gc _GridRowClass) Alloc() GridRow {
	rv := objc.CallMethod[GridRow](gc, objc.GetSelector("alloc"))
	return rv
}

func GridRow_Alloc() GridRow {
	return GridRowClass.Alloc()
}

func (gc _GridRowClass) New() GridRow {
	rv := objc.CallMethod[GridRow](gc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewGridRow() GridRow {
	return GridRowClass.New()
}

func GridRow_New() GridRow {
	return GridRowClass.New()
}

func (g_ GridRow) Init() GridRow {
	rv := objc.CallMethod[GridRow](g_, objc.GetSelector("init"))
	return rv
}

func GridRow_Init() GridRow {
	return GridRowClass.Alloc().Init()
}

func (g_ GridRow) CellAtIndex(index int) GridCell {
	rv := objc.CallMethod[GridCell](g_, objc.GetSelector("cellAtIndex:"), index)
	return rv
}

func (g_ GridRow) MergeCellsInRange(range_ foundation.Range) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("mergeCellsInRange:"), range_)
}

func (g_ GridRow) NumberOfCells() int {
	rv := objc.CallMethod[int](g_, objc.GetSelector("numberOfCells"))
	return rv
}

func (g_ GridRow) IsHidden() bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("isHidden"))
	return rv
}

func (g_ GridRow) SetHidden(value bool) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setHidden:"), value)
}

func (g_ GridRow) TopPadding() float64 {
	rv := objc.CallMethod[float64](g_, objc.GetSelector("topPadding"))
	return rv
}

func (g_ GridRow) SetTopPadding(value float64) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setTopPadding:"), value)
}

func (g_ GridRow) BottomPadding() float64 {
	rv := objc.CallMethod[float64](g_, objc.GetSelector("bottomPadding"))
	return rv
}

func (g_ GridRow) SetBottomPadding(value float64) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setBottomPadding:"), value)
}

func (g_ GridRow) Height() float64 {
	rv := objc.CallMethod[float64](g_, objc.GetSelector("height"))
	return rv
}

func (g_ GridRow) SetHeight(value float64) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setHeight:"), value)
}

func (g_ GridRow) RowAlignment() GridRowAlignment {
	rv := objc.CallMethod[GridRowAlignment](g_, objc.GetSelector("rowAlignment"))
	return rv
}

func (g_ GridRow) SetRowAlignment(value GridRowAlignment) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setRowAlignment:"), value)
}

func (g_ GridRow) YPlacement() GridCellPlacement {
	rv := objc.CallMethod[GridCellPlacement](g_, objc.GetSelector("yPlacement"))
	return rv
}

func (g_ GridRow) SetYPlacement(value GridCellPlacement) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setYPlacement:"), value)
}

func (g_ GridRow) GridView() GridView {
	rv := objc.CallMethod[GridView](g_, objc.GetSelector("gridView"))
	return rv
}
