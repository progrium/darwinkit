// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TextFieldCellClass = _TextFieldCellClass{objc.GetClass("NSTextFieldCell")}

type _TextFieldCellClass struct {
	objc.Class
}

type ITextFieldCell interface {
	IActionCell
	SetWantsNotificationForMarkedText(flag bool)
	TextColor() Color
	SetTextColor(value IColor)
	BezelStyle() TextFieldBezelStyle
	SetBezelStyle(value TextFieldBezelStyle)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	PlaceholderString() string
	SetPlaceholderString(value string)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.IAttributedString)
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)
}

type TextFieldCell struct {
	ActionCell
}

func MakeTextFieldCell(ptr unsafe.Pointer) TextFieldCell {
	return TextFieldCell{
		ActionCell: MakeActionCell(ptr),
	}
}

func (t_ TextFieldCell) InitTextCell(string_ string) TextFieldCell {
	rv := objc.CallMethod[TextFieldCell](t_, objc.GetSelector("initTextCell:"), string_)
	return rv
}

func TextFieldCell_InitTextCell(string_ string) TextFieldCell {
	return TextFieldCellClass.Alloc().InitTextCell(string_)
}

func (t_ TextFieldCell) InitImageCell(image IImage) TextFieldCell {
	rv := objc.CallMethod[TextFieldCell](t_, objc.GetSelector("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func TextFieldCell_InitImageCell(image IImage) TextFieldCell {
	return TextFieldCellClass.Alloc().InitImageCell(image)
}

func (t_ TextFieldCell) Init() TextFieldCell {
	rv := objc.CallMethod[TextFieldCell](t_, objc.GetSelector("init"))
	return rv
}

func TextFieldCell_Init() TextFieldCell {
	return TextFieldCellClass.Alloc().Init()
}

func (tc _TextFieldCellClass) Alloc() TextFieldCell {
	rv := objc.CallMethod[TextFieldCell](tc, objc.GetSelector("alloc"))
	return rv
}

func TextFieldCell_Alloc() TextFieldCell {
	return TextFieldCellClass.Alloc()
}

func (tc _TextFieldCellClass) New() TextFieldCell {
	rv := objc.CallMethod[TextFieldCell](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextFieldCell() TextFieldCell {
	return TextFieldCellClass.New()
}

func TextFieldCell_New() TextFieldCell {
	return TextFieldCellClass.New()
}

func (t_ TextFieldCell) SetWantsNotificationForMarkedText(flag bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setWantsNotificationForMarkedText:"), flag)
}

func (t_ TextFieldCell) TextColor() Color {
	rv := objc.CallMethod[Color](t_, objc.GetSelector("textColor"))
	return rv
}

func (t_ TextFieldCell) SetTextColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTextColor:"), objc.ExtractPtr(value))
}

func (t_ TextFieldCell) BezelStyle() TextFieldBezelStyle {
	rv := objc.CallMethod[TextFieldBezelStyle](t_, objc.GetSelector("bezelStyle"))
	return rv
}

func (t_ TextFieldCell) SetBezelStyle(value TextFieldBezelStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBezelStyle:"), value)
}

func (t_ TextFieldCell) BackgroundColor() Color {
	rv := objc.CallMethod[Color](t_, objc.GetSelector("backgroundColor"))
	return rv
}

func (t_ TextFieldCell) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (t_ TextFieldCell) DrawsBackground() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("drawsBackground"))
	return rv
}

func (t_ TextFieldCell) SetDrawsBackground(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDrawsBackground:"), value)
}

func (t_ TextFieldCell) PlaceholderString() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("placeholderString"))
	return rv
}

func (t_ TextFieldCell) SetPlaceholderString(value string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setPlaceholderString:"), value)
}

func (t_ TextFieldCell) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](t_, objc.GetSelector("placeholderAttributedString"))
	return rv
}

func (t_ TextFieldCell) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setPlaceholderAttributedString:"), objc.ExtractPtr(value))
}

func (t_ TextFieldCell) AllowedInputSourceLocales() []string {
	rv := objc.CallMethod[[]string](t_, objc.GetSelector("allowedInputSourceLocales"))
	return rv
}

func (t_ TextFieldCell) SetAllowedInputSourceLocales(value []string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowedInputSourceLocales:"), value)
}
