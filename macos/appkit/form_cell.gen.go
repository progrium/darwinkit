// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FormCell] class.
var FormCellClass = _FormCellClass{objc.GetClass("NSFormCell")}

type _FormCellClass struct {
	objc.Class
}

// An interface definition for the [FormCell] class.
type IFormCell interface {
	IActionCell
	TitleAlignment() TextAlignment
	SetTitleAlignment(value TextAlignment)
	TitleWidth() float64
	SetTitleWidth(value float64)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.IAttributedString)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	PreferredTextFieldWidth() float64
	SetPreferredTextFieldWidth(value float64)
	TitleBaseWritingDirection() WritingDirection
	SetTitleBaseWritingDirection(value WritingDirection)
	PlaceholderString() string
	SetPlaceholderString(value string)
	TitleFont() Font
	SetTitleFont(value IFont)
}

// The NSFormCell class is used to implement text entry fields in a form. The left part of an NSFormCell object contains a title. The right part contains an editable text entry field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell?language=objc
type FormCell struct {
	ActionCell
}

func FormCellFrom(ptr unsafe.Pointer) FormCell {
	return FormCell{
		ActionCell: ActionCellFrom(ptr),
	}
}

func (f_ FormCell) InitTextCell(string_ string) FormCell {
	rv := objc.Call[FormCell](f_, objc.Sel("initTextCell:"), string_)
	return rv
}

// Returns an NSFormCell object initialized with the specified title string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/1526669-inittextcell?language=objc
func FormCell_InitTextCell(string_ string) FormCell {
	return FormCellClass.Alloc().InitTextCell(string_)
}

func (fc _FormCellClass) Alloc() FormCell {
	rv := objc.Call[FormCell](fc, objc.Sel("alloc"))
	return rv
}

func FormCell_Alloc() FormCell {
	return FormCellClass.Alloc()
}

func (fc _FormCellClass) New() FormCell {
	rv := objc.Call[FormCell](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFormCell() FormCell {
	return FormCellClass.New()
}

func (f_ FormCell) Init() FormCell {
	rv := objc.Call[FormCell](f_, objc.Sel("init"))
	return rv
}

func (f_ FormCell) InitImageCell(image IImage) FormCell {
	rv := objc.Call[FormCell](f_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

// Returns an NSCell object initialized with the specified image and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533898-initimagecell?language=objc
func FormCell_InitImageCell(image IImage) FormCell {
	return FormCellClass.Alloc().InitImageCell(image)
}

// The alignment of the title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/1525716-titlealignment?language=objc
func (f_ FormCell) TitleAlignment() TextAlignment {
	rv := objc.Call[TextAlignment](f_, objc.Sel("titleAlignment"))
	return rv
}

// The alignment of the title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/1525716-titlealignment?language=objc
func (f_ FormCell) SetTitleAlignment(value TextAlignment) {
	objc.Call[objc.Void](f_, objc.Sel("setTitleAlignment:"), value)
}

// The width of the title field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/1535464-titlewidth?language=objc
func (f_ FormCell) TitleWidth() float64 {
	rv := objc.Call[float64](f_, objc.Sel("titleWidth"))
	return rv
}

// The width of the title field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/1535464-titlewidth?language=objc
func (f_ FormCell) SetTitleWidth(value float64) {
	objc.Call[objc.Void](f_, objc.Sel("setTitleWidth:"), value)
}

// The cell’s attributed placeholder string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/1535914-placeholderattributedstring?language=objc
func (f_ FormCell) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](f_, objc.Sel("placeholderAttributedString"))
	return rv
}

// The cell’s attributed placeholder string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/1535914-placeholderattributedstring?language=objc
func (f_ FormCell) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	objc.Call[objc.Void](f_, objc.Sel("setPlaceholderAttributedString:"), objc.Ptr(value))
}

// The title of the cell as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/1527410-attributedtitle?language=objc
func (f_ FormCell) AttributedTitle() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](f_, objc.Sel("attributedTitle"))
	return rv
}

// The title of the cell as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/1527410-attributedtitle?language=objc
func (f_ FormCell) SetAttributedTitle(value foundation.IAttributedString) {
	objc.Call[objc.Void](f_, objc.Sel("setAttributedTitle:"), objc.Ptr(value))
}

// The preferred text field width. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/1527483-preferredtextfieldwidth?language=objc
func (f_ FormCell) PreferredTextFieldWidth() float64 {
	rv := objc.Call[float64](f_, objc.Sel("preferredTextFieldWidth"))
	return rv
}

// The preferred text field width. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/1527483-preferredtextfieldwidth?language=objc
func (f_ FormCell) SetPreferredTextFieldWidth(value float64) {
	objc.Call[objc.Void](f_, objc.Sel("setPreferredTextFieldWidth:"), value)
}

// The default writing direction used to render the form cell’s title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/1526419-titlebasewritingdirection?language=objc
func (f_ FormCell) TitleBaseWritingDirection() WritingDirection {
	rv := objc.Call[WritingDirection](f_, objc.Sel("titleBaseWritingDirection"))
	return rv
}

// The default writing direction used to render the form cell’s title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/1526419-titlebasewritingdirection?language=objc
func (f_ FormCell) SetTitleBaseWritingDirection(value WritingDirection) {
	objc.Call[objc.Void](f_, objc.Sel("setTitleBaseWritingDirection:"), value)
}

// The cell’s plain text placeholder string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/1534043-placeholderstring?language=objc
func (f_ FormCell) PlaceholderString() string {
	rv := objc.Call[string](f_, objc.Sel("placeholderString"))
	return rv
}

// The cell’s plain text placeholder string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/1534043-placeholderstring?language=objc
func (f_ FormCell) SetPlaceholderString(value string) {
	objc.Call[objc.Void](f_, objc.Sel("setPlaceholderString:"), value)
}

// The font used to draw cell’s title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/1525255-titlefont?language=objc
func (f_ FormCell) TitleFont() Font {
	rv := objc.Call[Font](f_, objc.Sel("titleFont"))
	return rv
}

// The font used to draw cell’s title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsformcell/1525255-titlefont?language=objc
func (f_ FormCell) SetTitleFont(value IFont) {
	objc.Call[objc.Void](f_, objc.Sel("setTitleFont:"), objc.Ptr(value))
}
