// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TextFieldClass = _TextFieldClass{objc.GetClass("NSTextField")}

type _TextFieldClass struct {
	objc.Class
}

type ITextField interface {
	IControl
	SelectText(sender objc.IObject)
	TextShouldBeginEditing(textObject IText) bool
	TextDidBeginEditing(notification foundation.INotification)
	TextDidChange(notification foundation.INotification)
	TextShouldEndEditing(textObject IText) bool
	TextDidEndEditing(notification foundation.INotification)
	IsSelectable() bool
	SetSelectable(value bool)
	IsEditable() bool
	SetEditable(value bool)
	AllowsEditingTextAttributes() bool
	SetAllowsEditingTextAttributes(value bool)
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
	PlaceholderString() string
	SetPlaceholderString(value string)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.IAttributedString)
	LineBreakStrategy() LineBreakStrategy
	SetLineBreakStrategy(value LineBreakStrategy)
	AllowsDefaultTighteningForTruncation() bool
	SetAllowsDefaultTighteningForTruncation(value bool)
	MaximumNumberOfLines() int
	SetMaximumNumberOfLines(value int)
	PreferredMaxLayoutWidth() float64
	SetPreferredMaxLayoutWidth(value float64)
	TextColor() Color
	SetTextColor(value IColor)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	IsBezeled() bool
	SetBezeled(value bool)
	BezelStyle() TextFieldBezelStyle
	SetBezelStyle(value TextFieldBezelStyle)
	IsBordered() bool
	SetBordered(value bool)
	AllowsCharacterPickerTouchBarItem() bool
	SetAllowsCharacterPickerTouchBarItem(value bool)
	IsAutomaticTextCompletionEnabled() bool
	SetAutomaticTextCompletionEnabled(value bool)
	Delegate() TextFieldDelegateWrapper
	SetDelegate(value ITextFieldDelegate)
	SetDelegate0(value objc.IObject)
}

type TextField struct {
	Control
}

func MakeTextField(ptr unsafe.Pointer) TextField {
	return TextField{
		Control: MakeControl(ptr),
	}
}

func (tc _TextFieldClass) LabelWithAttributedString(attributedStringValue foundation.IAttributedString) TextField {
	rv := objc.CallMethod[TextField](tc, objc.GetSelector("labelWithAttributedString:"), objc.ExtractPtr(attributedStringValue))
	return rv
}

func TextField_LabelWithAttributedString(attributedStringValue foundation.IAttributedString) TextField {
	return TextFieldClass.LabelWithAttributedString(attributedStringValue)
}

func (tc _TextFieldClass) LabelWithString(stringValue string) TextField {
	rv := objc.CallMethod[TextField](tc, objc.GetSelector("labelWithString:"), stringValue)
	return rv
}

func TextField_LabelWithString(stringValue string) TextField {
	return TextFieldClass.LabelWithString(stringValue)
}

func (tc _TextFieldClass) TextFieldWithString(stringValue string) TextField {
	rv := objc.CallMethod[TextField](tc, objc.GetSelector("textFieldWithString:"), stringValue)
	return rv
}

func TextField_TextFieldWithString(stringValue string) TextField {
	return TextFieldClass.TextFieldWithString(stringValue)
}

func (tc _TextFieldClass) WrappingLabelWithString(stringValue string) TextField {
	rv := objc.CallMethod[TextField](tc, objc.GetSelector("wrappingLabelWithString:"), stringValue)
	return rv
}

func TextField_WrappingLabelWithString(stringValue string) TextField {
	return TextFieldClass.WrappingLabelWithString(stringValue)
}

func (t_ TextField) InitWithFrame(frameRect foundation.Rect) TextField {
	rv := objc.CallMethod[TextField](t_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func TextField_InitWithFrame(frameRect foundation.Rect) TextField {
	return TextFieldClass.Alloc().InitWithFrame(frameRect)
}

func (t_ TextField) Init() TextField {
	rv := objc.CallMethod[TextField](t_, objc.GetSelector("init"))
	return rv
}

func TextField_Init() TextField {
	return TextFieldClass.Alloc().Init()
}

func (tc _TextFieldClass) Alloc() TextField {
	rv := objc.CallMethod[TextField](tc, objc.GetSelector("alloc"))
	return rv
}

func TextField_Alloc() TextField {
	return TextFieldClass.Alloc()
}

func (tc _TextFieldClass) New() TextField {
	rv := objc.CallMethod[TextField](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextField() TextField {
	return TextFieldClass.New()
}

func TextField_New() TextField {
	return TextFieldClass.New()
}

func (t_ TextField) SelectText(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("selectText:"), objc.ExtractPtr(sender))
}

func (t_ TextField) TextShouldBeginEditing(textObject IText) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textShouldBeginEditing:"), objc.ExtractPtr(textObject))
	return rv
}

func (t_ TextField) TextDidBeginEditing(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textDidBeginEditing:"), objc.ExtractPtr(notification))
}

func (t_ TextField) TextDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textDidChange:"), objc.ExtractPtr(notification))
}

func (t_ TextField) TextShouldEndEditing(textObject IText) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textShouldEndEditing:"), objc.ExtractPtr(textObject))
	return rv
}

func (t_ TextField) TextDidEndEditing(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textDidEndEditing:"), objc.ExtractPtr(notification))
}

func (t_ TextField) IsSelectable() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isSelectable"))
	return rv
}

func (t_ TextField) SetSelectable(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSelectable:"), value)
}

func (t_ TextField) IsEditable() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isEditable"))
	return rv
}

func (t_ TextField) SetEditable(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setEditable:"), value)
}

func (t_ TextField) AllowsEditingTextAttributes() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsEditingTextAttributes"))
	return rv
}

func (t_ TextField) SetAllowsEditingTextAttributes(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsEditingTextAttributes:"), value)
}

func (t_ TextField) ImportsGraphics() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("importsGraphics"))
	return rv
}

func (t_ TextField) SetImportsGraphics(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setImportsGraphics:"), value)
}

func (t_ TextField) PlaceholderString() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("placeholderString"))
	return rv
}

func (t_ TextField) SetPlaceholderString(value string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setPlaceholderString:"), value)
}

func (t_ TextField) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](t_, objc.GetSelector("placeholderAttributedString"))
	return rv
}

func (t_ TextField) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setPlaceholderAttributedString:"), objc.ExtractPtr(value))
}

func (t_ TextField) LineBreakStrategy() LineBreakStrategy {
	rv := objc.CallMethod[LineBreakStrategy](t_, objc.GetSelector("lineBreakStrategy"))
	return rv
}

func (t_ TextField) SetLineBreakStrategy(value LineBreakStrategy) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLineBreakStrategy:"), value)
}

func (t_ TextField) AllowsDefaultTighteningForTruncation() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsDefaultTighteningForTruncation"))
	return rv
}

func (t_ TextField) SetAllowsDefaultTighteningForTruncation(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsDefaultTighteningForTruncation:"), value)
}

func (t_ TextField) MaximumNumberOfLines() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("maximumNumberOfLines"))
	return rv
}

func (t_ TextField) SetMaximumNumberOfLines(value int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setMaximumNumberOfLines:"), value)
}

func (t_ TextField) PreferredMaxLayoutWidth() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("preferredMaxLayoutWidth"))
	return rv
}

func (t_ TextField) SetPreferredMaxLayoutWidth(value float64) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setPreferredMaxLayoutWidth:"), value)
}

func (t_ TextField) TextColor() Color {
	rv := objc.CallMethod[Color](t_, objc.GetSelector("textColor"))
	return rv
}

func (t_ TextField) SetTextColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTextColor:"), objc.ExtractPtr(value))
}

func (t_ TextField) BackgroundColor() Color {
	rv := objc.CallMethod[Color](t_, objc.GetSelector("backgroundColor"))
	return rv
}

func (t_ TextField) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (t_ TextField) DrawsBackground() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("drawsBackground"))
	return rv
}

func (t_ TextField) SetDrawsBackground(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDrawsBackground:"), value)
}

func (t_ TextField) IsBezeled() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isBezeled"))
	return rv
}

func (t_ TextField) SetBezeled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBezeled:"), value)
}

func (t_ TextField) BezelStyle() TextFieldBezelStyle {
	rv := objc.CallMethod[TextFieldBezelStyle](t_, objc.GetSelector("bezelStyle"))
	return rv
}

func (t_ TextField) SetBezelStyle(value TextFieldBezelStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBezelStyle:"), value)
}

func (t_ TextField) IsBordered() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isBordered"))
	return rv
}

func (t_ TextField) SetBordered(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBordered:"), value)
}

func (t_ TextField) AllowsCharacterPickerTouchBarItem() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsCharacterPickerTouchBarItem"))
	return rv
}

func (t_ TextField) SetAllowsCharacterPickerTouchBarItem(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsCharacterPickerTouchBarItem:"), value)
}

func (t_ TextField) IsAutomaticTextCompletionEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isAutomaticTextCompletionEnabled"))
	return rv
}

func (t_ TextField) SetAutomaticTextCompletionEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAutomaticTextCompletionEnabled:"), value)
}

func (t_ TextField) Delegate() TextFieldDelegateWrapper {
	rv := objc.CallMethod[TextFieldDelegateWrapper](t_, objc.GetSelector("delegate"))
	return rv
}

func (t_ TextField) SetDelegate(value ITextFieldDelegate) {
	po := objc.WrapAsProtocol("NSTextFieldDelegate", value)
	objc.SetAssociatedObject(t_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDelegate:"), po)
}

func (t_ TextField) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}
