// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ActionCell] class.
var ActionCellClass = _ActionCellClass{objc.GetClass("NSActionCell")}

type _ActionCellClass struct {
	objc.Class
}

// An interface definition for the [ActionCell] class.
type IActionCell interface {
	ICell
}

// An active area inside a control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsactioncell?language=objc
type ActionCell struct {
	Cell
}

func ActionCellFrom(ptr unsafe.Pointer) ActionCell {
	return ActionCell{
		Cell: CellFrom(ptr),
	}
}

func (ac _ActionCellClass) Alloc() ActionCell {
	rv := objc.Call[ActionCell](ac, objc.Sel("alloc"))
	return rv
}

func ActionCell_Alloc() ActionCell {
	return ActionCellClass.Alloc()
}

func (ac _ActionCellClass) New() ActionCell {
	rv := objc.Call[ActionCell](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewActionCell() ActionCell {
	return ActionCellClass.New()
}

func (a_ ActionCell) Init() ActionCell {
	rv := objc.Call[ActionCell](a_, objc.Sel("init"))
	return rv
}

func (a_ ActionCell) InitImageCell(image IImage) ActionCell {
	rv := objc.Call[ActionCell](a_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

// Returns an NSCell object initialized with the specified image and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533898-initimagecell?language=objc
func NewActionCellImageCell(image IImage) ActionCell {
	instance := ActionCellClass.Alloc().InitImageCell(image)
	instance.Autorelease()
	return instance
}

func (a_ ActionCell) InitTextCell(string_ string) ActionCell {
	rv := objc.Call[ActionCell](a_, objc.Sel("initTextCell:"), string_)
	return rv
}

// Returns an NSCell object initialized with the specified string and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530851-inittextcell?language=objc
func NewActionCellTextCell(string_ string) ActionCell {
	instance := ActionCellClass.Alloc().InitTextCell(string_)
	instance.Autorelease()
	return instance
}
