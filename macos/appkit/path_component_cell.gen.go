// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PathComponentCell] class.
var PathComponentCellClass = _PathComponentCellClass{objc.GetClass("NSPathComponentCell")}

type _PathComponentCellClass struct {
	objc.Class
}

// An interface definition for the [PathComponentCell] class.
type IPathComponentCell interface {
	ITextFieldCell
	URL() foundation.URL
	SetURL(value foundation.IURL)
}

// A component of a path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcomponentcell?language=objc
type PathComponentCell struct {
	TextFieldCell
}

func PathComponentCellFrom(ptr unsafe.Pointer) PathComponentCell {
	return PathComponentCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}

func (pc _PathComponentCellClass) Alloc() PathComponentCell {
	rv := objc.Call[PathComponentCell](pc, objc.Sel("alloc"))
	return rv
}

func PathComponentCell_Alloc() PathComponentCell {
	return PathComponentCellClass.Alloc()
}

func (pc _PathComponentCellClass) New() PathComponentCell {
	rv := objc.Call[PathComponentCell](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPathComponentCell() PathComponentCell {
	return PathComponentCellClass.New()
}

func (p_ PathComponentCell) Init() PathComponentCell {
	rv := objc.Call[PathComponentCell](p_, objc.Sel("init"))
	return rv
}

func (p_ PathComponentCell) InitTextCell(string_ string) PathComponentCell {
	rv := objc.Call[PathComponentCell](p_, objc.Sel("initTextCell:"), string_)
	return rv
}

// Initializes a text field cell that displays the specified string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1642278-inittextcell?language=objc
func PathComponentCell_InitTextCell(string_ string) PathComponentCell {
	return PathComponentCellClass.Alloc().InitTextCell(string_)
}

func (p_ PathComponentCell) InitImageCell(image IImage) PathComponentCell {
	rv := objc.Call[PathComponentCell](p_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

// Returns an NSCell object initialized with the specified image and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533898-initimagecell?language=objc
func PathComponentCell_InitImageCell(image IImage) PathComponentCell {
	return PathComponentCellClass.Alloc().InitImageCell(image)
}

// The portion of the path from the root through the component represented by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcomponentcell/1534779-url?language=objc
func (p_ PathComponentCell) URL() foundation.URL {
	rv := objc.Call[foundation.URL](p_, objc.Sel("URL"))
	return rv
}

// The portion of the path from the root through the component represented by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcomponentcell/1534779-url?language=objc
func (p_ PathComponentCell) SetURL(value foundation.IURL) {
	objc.Call[objc.Void](p_, objc.Sel("setURL:"), objc.Ptr(value))
}
