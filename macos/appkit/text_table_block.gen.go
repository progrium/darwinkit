// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextTableBlock] class.
var TextTableBlockClass = _TextTableBlockClass{objc.GetClass("NSTextTableBlock")}

type _TextTableBlockClass struct {
	objc.Class
}

// An interface definition for the [TextTableBlock] class.
type ITextTableBlock interface {
	ITextBlock
	RowSpan() int
	StartingRow() int
	Table() TextTable
	StartingColumn() int
	ColumnSpan() int
}

// A text block that appears as a cell in a text table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttableblock?language=objc
type TextTableBlock struct {
	TextBlock
}

func TextTableBlockFrom(ptr unsafe.Pointer) TextTableBlock {
	return TextTableBlock{
		TextBlock: TextBlockFrom(ptr),
	}
}

func (t_ TextTableBlock) InitWithTableStartingRowRowSpanStartingColumnColumnSpan(table ITextTable, row int, rowSpan int, col int, colSpan int) TextTableBlock {
	rv := objc.Call[TextTableBlock](t_, objc.Sel("initWithTable:startingRow:rowSpan:startingColumn:columnSpan:"), objc.Ptr(table), row, rowSpan, col, colSpan)
	return rv
}

// Returns an initialized text table block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttableblock/1532894-initwithtable?language=objc
func TextTableBlock_InitWithTableStartingRowRowSpanStartingColumnColumnSpan(table ITextTable, row int, rowSpan int, col int, colSpan int) TextTableBlock {
	return TextTableBlockClass.Alloc().InitWithTableStartingRowRowSpanStartingColumnColumnSpan(table, row, rowSpan, col, colSpan)
}

func (tc _TextTableBlockClass) Alloc() TextTableBlock {
	rv := objc.Call[TextTableBlock](tc, objc.Sel("alloc"))
	return rv
}

func TextTableBlock_Alloc() TextTableBlock {
	return TextTableBlockClass.Alloc()
}

func (tc _TextTableBlockClass) New() TextTableBlock {
	rv := objc.Call[TextTableBlock](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextTableBlock() TextTableBlock {
	return TextTableBlockClass.New()
}

func (t_ TextTableBlock) Init() TextTableBlock {
	rv := objc.Call[TextTableBlock](t_, objc.Sel("init"))
	return rv
}

// Returns the number of table rows spanned by this text table block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttableblock/1528586-rowspan?language=objc
func (t_ TextTableBlock) RowSpan() int {
	rv := objc.Call[int](t_, objc.Sel("rowSpan"))
	return rv
}

// Returns the table row at which this text table block starts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttableblock/1525803-startingrow?language=objc
func (t_ TextTableBlock) StartingRow() int {
	rv := objc.Call[int](t_, objc.Sel("startingRow"))
	return rv
}

// Returns the table containing this text table block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttableblock/1525141-table?language=objc
func (t_ TextTableBlock) Table() TextTable {
	rv := objc.Call[TextTable](t_, objc.Sel("table"))
	return rv
}

// Returns the table column at which this text table block starts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttableblock/1525383-startingcolumn?language=objc
func (t_ TextTableBlock) StartingColumn() int {
	rv := objc.Call[int](t_, objc.Sel("startingColumn"))
	return rv
}

// Returns the number of table columns spanned by this text table block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttableblock/1528568-columnspan?language=objc
func (t_ TextTableBlock) ColumnSpan() int {
	rv := objc.Call[int](t_, objc.Sel("columnSpan"))
	return rv
}
