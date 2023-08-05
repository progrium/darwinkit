// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var PathComponentCellClass = _PathComponentCellClass{objc.GetClass("NSPathComponentCell")}

type _PathComponentCellClass struct {
	objc.Class
}

type IPathComponentCell interface {
	ITextFieldCell
	URL() foundation.URL
	SetURL(value foundation.IURL)
}

type PathComponentCell struct {
	TextFieldCell
}

func MakePathComponentCell(ptr unsafe.Pointer) PathComponentCell {
	return PathComponentCell{
		TextFieldCell: MakeTextFieldCell(ptr),
	}
}

func (p_ PathComponentCell) InitTextCell(string_ string) PathComponentCell {
	rv := objc.CallMethod[PathComponentCell](p_, objc.GetSelector("initTextCell:"), string_)
	return rv
}

func PathComponentCell_InitTextCell(string_ string) PathComponentCell {
	return PathComponentCellClass.Alloc().InitTextCell(string_)
}

func (p_ PathComponentCell) InitImageCell(image IImage) PathComponentCell {
	rv := objc.CallMethod[PathComponentCell](p_, objc.GetSelector("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func PathComponentCell_InitImageCell(image IImage) PathComponentCell {
	return PathComponentCellClass.Alloc().InitImageCell(image)
}

func (p_ PathComponentCell) Init() PathComponentCell {
	rv := objc.CallMethod[PathComponentCell](p_, objc.GetSelector("init"))
	return rv
}

func PathComponentCell_Init() PathComponentCell {
	return PathComponentCellClass.Alloc().Init()
}

func (pc _PathComponentCellClass) Alloc() PathComponentCell {
	rv := objc.CallMethod[PathComponentCell](pc, objc.GetSelector("alloc"))
	return rv
}

func PathComponentCell_Alloc() PathComponentCell {
	return PathComponentCellClass.Alloc()
}

func (pc _PathComponentCellClass) New() PathComponentCell {
	rv := objc.CallMethod[PathComponentCell](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPathComponentCell() PathComponentCell {
	return PathComponentCellClass.New()
}

func PathComponentCell_New() PathComponentCell {
	return PathComponentCellClass.New()
}

func (p_ PathComponentCell) URL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](p_, objc.GetSelector("URL"))
	return rv
}

func (p_ PathComponentCell) SetURL(value foundation.IURL) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setURL:"), objc.ExtractPtr(value))
}
