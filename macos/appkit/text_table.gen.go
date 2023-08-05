// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var TextTableClass = _TextTableClass{objc.GetClass("NSTextTable")}

type _TextTableClass struct {
	objc.Class
}

type ITextTable interface {
	ITextBlock
	NumberOfColumns() uint
	SetNumberOfColumns(value uint)
	LayoutAlgorithm() TextTableLayoutAlgorithm
	SetLayoutAlgorithm(value TextTableLayoutAlgorithm)
	CollapsesBorders() bool
	SetCollapsesBorders(value bool)
	HidesEmptyCells() bool
	SetHidesEmptyCells(value bool)
}

type TextTable struct {
	TextBlock
}

func MakeTextTable(ptr unsafe.Pointer) TextTable {
	return TextTable{
		TextBlock: MakeTextBlock(ptr),
	}
}

func (t_ TextTable) Init() TextTable {
	rv := objc.CallMethod[TextTable](t_, objc.GetSelector("init"))
	return rv
}

func TextTable_Init() TextTable {
	return TextTableClass.Alloc().Init()
}

func (tc _TextTableClass) Alloc() TextTable {
	rv := objc.CallMethod[TextTable](tc, objc.GetSelector("alloc"))
	return rv
}

func TextTable_Alloc() TextTable {
	return TextTableClass.Alloc()
}

func (tc _TextTableClass) New() TextTable {
	rv := objc.CallMethod[TextTable](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextTable() TextTable {
	return TextTableClass.New()
}

func TextTable_New() TextTable {
	return TextTableClass.New()
}

func (t_ TextTable) NumberOfColumns() uint {
	rv := objc.CallMethod[uint](t_, objc.GetSelector("numberOfColumns"))
	return rv
}

func (t_ TextTable) SetNumberOfColumns(value uint) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setNumberOfColumns:"), value)
}

func (t_ TextTable) LayoutAlgorithm() TextTableLayoutAlgorithm {
	rv := objc.CallMethod[TextTableLayoutAlgorithm](t_, objc.GetSelector("layoutAlgorithm"))
	return rv
}

func (t_ TextTable) SetLayoutAlgorithm(value TextTableLayoutAlgorithm) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLayoutAlgorithm:"), value)
}

func (t_ TextTable) CollapsesBorders() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("collapsesBorders"))
	return rv
}

func (t_ TextTable) SetCollapsesBorders(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setCollapsesBorders:"), value)
}

func (t_ TextTable) HidesEmptyCells() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("hidesEmptyCells"))
	return rv
}

func (t_ TextTable) SetHidesEmptyCells(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setHidesEmptyCells:"), value)
}
