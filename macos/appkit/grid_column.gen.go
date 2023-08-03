// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var GridColumnClass = _GridColumnClass{objc.GetClass("NSGridColumn")}

type _GridColumnClass struct {
	objc.Class
}

type IGridColumn interface {
	objc.IObject
	CellAtIndex(index int) GridCell
	MergeCellsInRange(range_ foundation.Range)
	GridView() GridView
	IsHidden() bool
	SetHidden(value bool)
	LeadingPadding() float64
	SetLeadingPadding(value float64)
	NumberOfCells() int
	TrailingPadding() float64
	SetTrailingPadding(value float64)
	Width() float64
	SetWidth(value float64)
	XPlacement() GridCellPlacement
	SetXPlacement(value GridCellPlacement)
}

type GridColumn struct {
	objc.Object
}

func MakeGridColumn(ptr unsafe.Pointer) GridColumn {
	return GridColumn{
		Object: objc.MakeObject(ptr),
	}
}

func (gc _GridColumnClass) Alloc() GridColumn {
	rv := objc.CallMethod[GridColumn](gc, objc.GetSelector("alloc"))
	return rv
}

func GridColumn_Alloc() GridColumn {
	return GridColumnClass.Alloc()
}

func (gc _GridColumnClass) New() GridColumn {
	rv := objc.CallMethod[GridColumn](gc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewGridColumn() GridColumn {
	return GridColumnClass.New()
}

func GridColumn_New() GridColumn {
	return GridColumnClass.New()
}

func (g_ GridColumn) Init() GridColumn {
	rv := objc.CallMethod[GridColumn](g_, objc.GetSelector("init"))
	return rv
}

func GridColumn_Init() GridColumn {
	return GridColumnClass.Alloc().Init()
}

func (g_ GridColumn) CellAtIndex(index int) GridCell {
	rv := objc.CallMethod[GridCell](g_, objc.GetSelector("cellAtIndex:"), index)
	return rv
}

func (g_ GridColumn) MergeCellsInRange(range_ foundation.Range) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("mergeCellsInRange:"), range_)
}

func (g_ GridColumn) GridView() GridView {
	rv := objc.CallMethod[GridView](g_, objc.GetSelector("gridView"))
	return rv
}

func (g_ GridColumn) IsHidden() bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("isHidden"))
	return rv
}

func (g_ GridColumn) SetHidden(value bool) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setHidden:"), value)
}

func (g_ GridColumn) LeadingPadding() float64 {
	rv := objc.CallMethod[float64](g_, objc.GetSelector("leadingPadding"))
	return rv
}

func (g_ GridColumn) SetLeadingPadding(value float64) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setLeadingPadding:"), value)
}

func (g_ GridColumn) NumberOfCells() int {
	rv := objc.CallMethod[int](g_, objc.GetSelector("numberOfCells"))
	return rv
}

func (g_ GridColumn) TrailingPadding() float64 {
	rv := objc.CallMethod[float64](g_, objc.GetSelector("trailingPadding"))
	return rv
}

func (g_ GridColumn) SetTrailingPadding(value float64) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setTrailingPadding:"), value)
}

func (g_ GridColumn) Width() float64 {
	rv := objc.CallMethod[float64](g_, objc.GetSelector("width"))
	return rv
}

func (g_ GridColumn) SetWidth(value float64) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setWidth:"), value)
}

func (g_ GridColumn) XPlacement() GridCellPlacement {
	rv := objc.CallMethod[GridCellPlacement](g_, objc.GetSelector("xPlacement"))
	return rv
}

func (g_ GridColumn) SetXPlacement(value GridCellPlacement) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setXPlacement:"), value)
}
