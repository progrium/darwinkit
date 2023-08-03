// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var GridCellClass = _GridCellClass{objc.GetClass("NSGridCell")}

type _GridCellClass struct {
	objc.Class
}

type IGridCell interface {
	objc.IObject
	Column() GridColumn
	Row() GridRow
	ContentView() View
	SetContentView(value IView)
	CustomPlacementConstraints() []LayoutConstraint
	SetCustomPlacementConstraints(value []ILayoutConstraint)
	RowAlignment() GridRowAlignment
	SetRowAlignment(value GridRowAlignment)
	XPlacement() GridCellPlacement
	SetXPlacement(value GridCellPlacement)
	YPlacement() GridCellPlacement
	SetYPlacement(value GridCellPlacement)
}

type GridCell struct {
	objc.Object
}

func MakeGridCell(ptr unsafe.Pointer) GridCell {
	return GridCell{
		Object: objc.MakeObject(ptr),
	}
}

func (gc _GridCellClass) Alloc() GridCell {
	rv := objc.CallMethod[GridCell](gc, objc.GetSelector("alloc"))
	return rv
}

func GridCell_Alloc() GridCell {
	return GridCellClass.Alloc()
}

func (gc _GridCellClass) New() GridCell {
	rv := objc.CallMethod[GridCell](gc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewGridCell() GridCell {
	return GridCellClass.New()
}

func GridCell_New() GridCell {
	return GridCellClass.New()
}

func (g_ GridCell) Init() GridCell {
	rv := objc.CallMethod[GridCell](g_, objc.GetSelector("init"))
	return rv
}

func GridCell_Init() GridCell {
	return GridCellClass.Alloc().Init()
}

func (g_ GridCell) Column() GridColumn {
	rv := objc.CallMethod[GridColumn](g_, objc.GetSelector("column"))
	return rv
}

func (g_ GridCell) Row() GridRow {
	rv := objc.CallMethod[GridRow](g_, objc.GetSelector("row"))
	return rv
}

func (g_ GridCell) ContentView() View {
	rv := objc.CallMethod[View](g_, objc.GetSelector("contentView"))
	return rv
}

func (g_ GridCell) SetContentView(value IView) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setContentView:"), objc.ExtractPtr(value))
}

func (gc _GridCellClass) EmptyContentView() View {
	rv := objc.CallMethod[View](gc, objc.GetSelector("emptyContentView"))
	return rv
}

func GridCell_EmptyContentView() View {
	return GridCellClass.EmptyContentView()
}

func (g_ GridCell) CustomPlacementConstraints() []LayoutConstraint {
	rv := objc.CallMethod[[]LayoutConstraint](g_, objc.GetSelector("customPlacementConstraints"))
	// TODO: convert slice items...
	return rv
}

func (g_ GridCell) SetCustomPlacementConstraints(value []ILayoutConstraint) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setCustomPlacementConstraints:"), value)
}

func (g_ GridCell) RowAlignment() GridRowAlignment {
	rv := objc.CallMethod[GridRowAlignment](g_, objc.GetSelector("rowAlignment"))
	return rv
}

func (g_ GridCell) SetRowAlignment(value GridRowAlignment) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setRowAlignment:"), value)
}

func (g_ GridCell) XPlacement() GridCellPlacement {
	rv := objc.CallMethod[GridCellPlacement](g_, objc.GetSelector("xPlacement"))
	return rv
}

func (g_ GridCell) SetXPlacement(value GridCellPlacement) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setXPlacement:"), value)
}

func (g_ GridCell) YPlacement() GridCellPlacement {
	rv := objc.CallMethod[GridCellPlacement](g_, objc.GetSelector("yPlacement"))
	return rv
}

func (g_ GridCell) SetYPlacement(value GridCellPlacement) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setYPlacement:"), value)
}
