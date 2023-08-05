// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ActionCellClass = _ActionCellClass{objc.GetClass("NSActionCell")}

type _ActionCellClass struct {
	objc.Class
}

type IActionCell interface {
	ICell
}

type ActionCell struct {
	Cell
}

func MakeActionCell(ptr unsafe.Pointer) ActionCell {
	return ActionCell{
		Cell: MakeCell(ptr),
	}
}

func (a_ ActionCell) InitImageCell(image IImage) ActionCell {
	rv := objc.CallMethod[ActionCell](a_, objc.GetSelector("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func ActionCell_InitImageCell(image IImage) ActionCell {
	return ActionCellClass.Alloc().InitImageCell(image)
}

func (a_ ActionCell) InitTextCell(string_ string) ActionCell {
	rv := objc.CallMethod[ActionCell](a_, objc.GetSelector("initTextCell:"), string_)
	return rv
}

func ActionCell_InitTextCell(string_ string) ActionCell {
	return ActionCellClass.Alloc().InitTextCell(string_)
}

func (a_ ActionCell) Init() ActionCell {
	rv := objc.CallMethod[ActionCell](a_, objc.GetSelector("init"))
	return rv
}

func ActionCell_Init() ActionCell {
	return ActionCellClass.Alloc().Init()
}

func (ac _ActionCellClass) Alloc() ActionCell {
	rv := objc.CallMethod[ActionCell](ac, objc.GetSelector("alloc"))
	return rv
}

func ActionCell_Alloc() ActionCell {
	return ActionCellClass.Alloc()
}

func (ac _ActionCellClass) New() ActionCell {
	rv := objc.CallMethod[ActionCell](ac, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewActionCell() ActionCell {
	return ActionCellClass.New()
}

func ActionCell_New() ActionCell {
	return ActionCellClass.New()
}
