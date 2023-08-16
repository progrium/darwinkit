// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GridRow] class.
var GridRowClass = _GridRowClass{objc.GetClass("NSGridRow")}

type _GridRowClass struct {
	objc.Class
}

// An interface definition for the [GridRow] class.
type IGridRow interface {
	objc.IObject
	MergeCellsInRange(range_ foundation.Range)
	CellAtIndex(index int) GridCell
	RowAlignment() GridRowAlignment
	SetRowAlignment(value GridRowAlignment)
	IsHidden() bool
	SetHidden(value bool)
	YPlacement() GridCellPlacement
	SetYPlacement(value GridCellPlacement)
	Height() float64
	SetHeight(value float64)
	GridView() GridView
	BottomPadding() float64
	SetBottomPadding(value float64)
	NumberOfCells() int
	TopPadding() float64
	SetTopPadding(value float64)
}

// A row within a grid view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow?language=objc
type GridRow struct {
	objc.Object
}

func GridRowFrom(ptr unsafe.Pointer) GridRow {
	return GridRow{
		Object: objc.ObjectFrom(ptr),
	}
}

func (gc _GridRowClass) Alloc() GridRow {
	rv := objc.Call[GridRow](gc, objc.Sel("alloc"))
	return rv
}

func GridRow_Alloc() GridRow {
	return GridRowClass.Alloc()
}

func (gc _GridRowClass) New() GridRow {
	rv := objc.Call[GridRow](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGridRow() GridRow {
	return GridRowClass.New()
}

func (g_ GridRow) Init() GridRow {
	rv := objc.Call[GridRow](g_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/1639726-mergecellsinrange?language=objc
func (g_ GridRow) MergeCellsInRange(range_ foundation.Range) {
	objc.Call[objc.Void](g_, objc.Sel("mergeCellsInRange:"), range_)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/1639712-cellatindex?language=objc
func (g_ GridRow) CellAtIndex(index int) GridCell {
	rv := objc.Call[GridCell](g_, objc.Sel("cellAtIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/1823690-rowalignment?language=objc
func (g_ GridRow) RowAlignment() GridRowAlignment {
	rv := objc.Call[GridRowAlignment](g_, objc.Sel("rowAlignment"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/1823690-rowalignment?language=objc
func (g_ GridRow) SetRowAlignment(value GridRowAlignment) {
	objc.Call[objc.Void](g_, objc.Sel("setRowAlignment:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/1639723-hidden?language=objc
func (g_ GridRow) IsHidden() bool {
	rv := objc.Call[bool](g_, objc.Sel("isHidden"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/1639723-hidden?language=objc
func (g_ GridRow) SetHidden(value bool) {
	objc.Call[objc.Void](g_, objc.Sel("setHidden:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/1639665-yplacement?language=objc
func (g_ GridRow) YPlacement() GridCellPlacement {
	rv := objc.Call[GridCellPlacement](g_, objc.Sel("yPlacement"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/1639665-yplacement?language=objc
func (g_ GridRow) SetYPlacement(value GridCellPlacement) {
	objc.Call[objc.Void](g_, objc.Sel("setYPlacement:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/1639728-height?language=objc
func (g_ GridRow) Height() float64 {
	rv := objc.Call[float64](g_, objc.Sel("height"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/1639728-height?language=objc
func (g_ GridRow) SetHeight(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setHeight:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/1639773-gridview?language=objc
func (g_ GridRow) GridView() GridView {
	rv := objc.Call[GridView](g_, objc.Sel("gridView"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/1639739-bottompadding?language=objc
func (g_ GridRow) BottomPadding() float64 {
	rv := objc.Call[float64](g_, objc.Sel("bottomPadding"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/1639739-bottompadding?language=objc
func (g_ GridRow) SetBottomPadding(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setBottomPadding:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/1639685-numberofcells?language=objc
func (g_ GridRow) NumberOfCells() int {
	rv := objc.Call[int](g_, objc.Sel("numberOfCells"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/1639785-toppadding?language=objc
func (g_ GridRow) TopPadding() float64 {
	rv := objc.Call[float64](g_, objc.Sel("topPadding"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/1639785-toppadding?language=objc
func (g_ GridRow) SetTopPadding(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setTopPadding:"), value)
}
