// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TableHeaderCell] class.
var TableHeaderCellClass = _TableHeaderCellClass{objc.GetClass("NSTableHeaderCell")}

type _TableHeaderCellClass struct {
	objc.Class
}

// An interface definition for the [TableHeaderCell] class.
type ITableHeaderCell interface {
	ITextFieldCell
	DrawSortIndicatorWithFrameInViewAscendingPriority(cellFrame foundation.Rect, controlView IView, ascending bool, priority int)
	SortIndicatorRectForBounds(rect foundation.Rect) foundation.Rect
}

// An object that a table header view uses to draw the content of the column headers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableheadercell?language=objc
type TableHeaderCell struct {
	TextFieldCell
}

func TableHeaderCellFrom(ptr unsafe.Pointer) TableHeaderCell {
	return TableHeaderCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}

func (tc _TableHeaderCellClass) Alloc() TableHeaderCell {
	rv := objc.Call[TableHeaderCell](tc, objc.Sel("alloc"))
	return rv
}

func TableHeaderCell_Alloc() TableHeaderCell {
	return TableHeaderCellClass.Alloc()
}

func (tc _TableHeaderCellClass) New() TableHeaderCell {
	rv := objc.Call[TableHeaderCell](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTableHeaderCell() TableHeaderCell {
	return TableHeaderCellClass.New()
}

func (t_ TableHeaderCell) Init() TableHeaderCell {
	rv := objc.Call[TableHeaderCell](t_, objc.Sel("init"))
	return rv
}

func (t_ TableHeaderCell) InitTextCell(string_ string) TableHeaderCell {
	rv := objc.Call[TableHeaderCell](t_, objc.Sel("initTextCell:"), string_)
	return rv
}

// Initializes a text field cell that displays the specified string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1642278-inittextcell?language=objc
func TableHeaderCell_InitTextCell(string_ string) TableHeaderCell {
	return TableHeaderCellClass.Alloc().InitTextCell(string_)
}

func (t_ TableHeaderCell) InitImageCell(image IImage) TableHeaderCell {
	rv := objc.Call[TableHeaderCell](t_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

// Returns an NSCell object initialized with the specified image and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533898-initimagecell?language=objc
func TableHeaderCell_InitImageCell(image IImage) TableHeaderCell {
	return TableHeaderCellClass.Alloc().InitImageCell(image)
}

// Draws a sorting indicator given a cell frame contained inside a view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableheadercell/1526553-drawsortindicatorwithframe?language=objc
func (t_ TableHeaderCell) DrawSortIndicatorWithFrameInViewAscendingPriority(cellFrame foundation.Rect, controlView IView, ascending bool, priority int) {
	objc.Call[objc.Void](t_, objc.Sel("drawSortIndicatorWithFrame:inView:ascending:priority:"), cellFrame, objc.Ptr(controlView), ascending, priority)
}

// Returns the location to display the sorting indicator given theRect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableheadercell/1525964-sortindicatorrectforbounds?language=objc
func (t_ TableHeaderCell) SortIndicatorRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](t_, objc.Sel("sortIndicatorRectForBounds:"), rect)
	return rv
}
