// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SecureTextFieldCell] class.
var SecureTextFieldCellClass = _SecureTextFieldCellClass{objc.GetClass("NSSecureTextFieldCell")}

type _SecureTextFieldCellClass struct {
	objc.Class
}

// An interface definition for the [SecureTextFieldCell] class.
type ISecureTextFieldCell interface {
	ITextFieldCell
	EchosBullets() bool
	SetEchosBullets(value bool)
}

// A text field whose value is hidden from the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssecuretextfieldcell?language=objc
type SecureTextFieldCell struct {
	TextFieldCell
}

func SecureTextFieldCellFrom(ptr unsafe.Pointer) SecureTextFieldCell {
	return SecureTextFieldCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}

func (sc _SecureTextFieldCellClass) Alloc() SecureTextFieldCell {
	rv := objc.Call[SecureTextFieldCell](sc, objc.Sel("alloc"))
	return rv
}

func SecureTextFieldCell_Alloc() SecureTextFieldCell {
	return SecureTextFieldCellClass.Alloc()
}

func (sc _SecureTextFieldCellClass) New() SecureTextFieldCell {
	rv := objc.Call[SecureTextFieldCell](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSecureTextFieldCell() SecureTextFieldCell {
	return SecureTextFieldCellClass.New()
}

func (s_ SecureTextFieldCell) Init() SecureTextFieldCell {
	rv := objc.Call[SecureTextFieldCell](s_, objc.Sel("init"))
	return rv
}

func (s_ SecureTextFieldCell) InitTextCell(string_ string) SecureTextFieldCell {
	rv := objc.Call[SecureTextFieldCell](s_, objc.Sel("initTextCell:"), string_)
	return rv
}

// Initializes a text field cell that displays the specified string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1642278-inittextcell?language=objc
func SecureTextFieldCell_InitTextCell(string_ string) SecureTextFieldCell {
	return SecureTextFieldCellClass.Alloc().InitTextCell(string_)
}

func (s_ SecureTextFieldCell) InitImageCell(image IImage) SecureTextFieldCell {
	rv := objc.Call[SecureTextFieldCell](s_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

// Returns an NSCell object initialized with the specified image and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533898-initimagecell?language=objc
func SecureTextFieldCell_InitImageCell(image IImage) SecureTextFieldCell {
	return SecureTextFieldCellClass.Alloc().InitImageCell(image)
}

// A Boolean that indicates whether the receiver echoes a bullet character rather than each character typed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssecuretextfieldcell/1395984-echosbullets?language=objc
func (s_ SecureTextFieldCell) EchosBullets() bool {
	rv := objc.Call[bool](s_, objc.Sel("echosBullets"))
	return rv
}

// A Boolean that indicates whether the receiver echoes a bullet character rather than each character typed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssecuretextfieldcell/1395984-echosbullets?language=objc
func (s_ SecureTextFieldCell) SetEchosBullets(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setEchosBullets:"), value)
}
