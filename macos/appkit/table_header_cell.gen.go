// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TableHeaderCellClass = _TableHeaderCellClass{objc.GetClass("NSTableHeaderCell")}

type _TableHeaderCellClass struct {
	objc.Class
}

type ITableHeaderCell interface {
	ITextFieldCell
	DrawSortIndicatorWithFrameInViewAscendingPriority(cellFrame foundation.Rect, controlView IView, ascending bool, priority int)
	SortIndicatorRectForBounds(rect foundation.Rect) foundation.Rect
}

type TableHeaderCell struct {
	TextFieldCell
}

func MakeTableHeaderCell(ptr unsafe.Pointer) TableHeaderCell {
	return TableHeaderCell{
		TextFieldCell: MakeTextFieldCell(ptr),
	}
}

func (t_ TableHeaderCell) InitTextCell(string_ string) TableHeaderCell {
	rv := objc.CallMethod[TableHeaderCell](t_, objc.GetSelector("initTextCell:"), string_)
	return rv
}

func TableHeaderCell_InitTextCell(string_ string) TableHeaderCell {
	return TableHeaderCellClass.Alloc().InitTextCell(string_)
}

func (t_ TableHeaderCell) InitImageCell(image IImage) TableHeaderCell {
	rv := objc.CallMethod[TableHeaderCell](t_, objc.GetSelector("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func TableHeaderCell_InitImageCell(image IImage) TableHeaderCell {
	return TableHeaderCellClass.Alloc().InitImageCell(image)
}

func (t_ TableHeaderCell) Init() TableHeaderCell {
	rv := objc.CallMethod[TableHeaderCell](t_, objc.GetSelector("init"))
	return rv
}

func TableHeaderCell_Init() TableHeaderCell {
	return TableHeaderCellClass.Alloc().Init()
}

func (tc _TableHeaderCellClass) Alloc() TableHeaderCell {
	rv := objc.CallMethod[TableHeaderCell](tc, objc.GetSelector("alloc"))
	return rv
}

func TableHeaderCell_Alloc() TableHeaderCell {
	return TableHeaderCellClass.Alloc()
}

func (tc _TableHeaderCellClass) New() TableHeaderCell {
	rv := objc.CallMethod[TableHeaderCell](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTableHeaderCell() TableHeaderCell {
	return TableHeaderCellClass.New()
}

func TableHeaderCell_New() TableHeaderCell {
	return TableHeaderCellClass.New()
}

func (t_ TableHeaderCell) DrawSortIndicatorWithFrameInViewAscendingPriority(cellFrame foundation.Rect, controlView IView, ascending bool, priority int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("drawSortIndicatorWithFrame:inView:ascending:priority:"), cellFrame, objc.ExtractPtr(controlView), ascending, priority)
}

func (t_ TableHeaderCell) SortIndicatorRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.GetSelector("sortIndicatorRectForBounds:"), rect)
	return rv
}
