// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TextClass = _TextClass{objc.GetClass("NSText")}

type _TextClass struct {
	objc.Class
}

type IText interface {
	IView
	ToggleRuler(sender objc.IObject)
	ReplaceCharactersInRangeWithRTF(range_ foundation.Range, rtfData []byte)
	ReplaceCharactersInRangeWithRTFD(range_ foundation.Range, rtfdData []byte)
	ReplaceCharactersInRangeWithString(range_ foundation.Range, string_ string)
	SelectAll(sender objc.IObject)
	Copy(sender objc.IObject)
	Cut(sender objc.IObject)
	Paste(sender objc.IObject)
	CopyFont(sender objc.IObject)
	PasteFont(sender objc.IObject)
	CopyRuler(sender objc.IObject)
	PasteRuler(sender objc.IObject)
	Delete(sender objc.IObject)
	ChangeFont(sender objc.IObject)
	SetFontRange(font IFont, range_ foundation.Range)
	AlignCenter(sender objc.IObject)
	AlignLeft(sender objc.IObject)
	AlignRight(sender objc.IObject)
	SetTextColorRange(color IColor, range_ foundation.Range)
	Superscript(sender objc.IObject)
	Subscript(sender objc.IObject)
	Unscript(sender objc.IObject)
	Underline(sender objc.IObject)
	ReadRTFDFromFile(path string) bool
	WriteRTFDToFileAtomically(path string, flag bool) bool
	RTFDFromRange(range_ foundation.Range) []byte
	RTFFromRange(range_ foundation.Range) []byte
	CheckSpelling(sender objc.IObject)
	ShowGuessPanel(sender objc.IObject)
	SizeToFit()
	ScrollRangeToVisible(range_ foundation.Range)
	String() string
	SetString(value string)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	IsEditable() bool
	SetEditable(value bool)
	IsSelectable() bool
	SetSelectable(value bool)
	IsFieldEditor() bool
	SetFieldEditor(value bool)
	IsRichText() bool
	SetRichText(value bool)
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
	UsesFontPanel() bool
	SetUsesFontPanel(value bool)
	IsRulerVisible() bool
	SelectedRange() foundation.Range
	SetSelectedRange(value foundation.Range)
	Font() Font
	SetFont(value IFont)
	Alignment() TextAlignment
	SetAlignment(value TextAlignment)
	TextColor() Color
	SetTextColor(value IColor)
	BaseWritingDirection() WritingDirection
	SetBaseWritingDirection(value WritingDirection)
	MaxSize() foundation.Size
	SetMaxSize(value foundation.Size)
	MinSize() foundation.Size
	SetMinSize(value foundation.Size)
	IsVerticallyResizable() bool
	SetVerticallyResizable(value bool)
	IsHorizontallyResizable() bool
	SetHorizontallyResizable(value bool)
	Delegate() TextDelegateWrapper
	SetDelegate(value ITextDelegate)
	SetDelegate0(value objc.IObject)
}

type Text struct {
	View
}

func MakeText(ptr unsafe.Pointer) Text {
	return Text{
		View: MakeView(ptr),
	}
}

func (t_ Text) InitWithFrame(frameRect foundation.Rect) Text {
	rv := objc.CallMethod[Text](t_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func Text_InitWithFrame(frameRect foundation.Rect) Text {
	return TextClass.Alloc().InitWithFrame(frameRect)
}

func (t_ Text) Init() Text {
	rv := objc.CallMethod[Text](t_, objc.GetSelector("init"))
	return rv
}

func Text_Init() Text {
	return TextClass.Alloc().Init()
}

func (tc _TextClass) Alloc() Text {
	rv := objc.CallMethod[Text](tc, objc.GetSelector("alloc"))
	return rv
}

func Text_Alloc() Text {
	return TextClass.Alloc()
}

func (tc _TextClass) New() Text {
	rv := objc.CallMethod[Text](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewText() Text {
	return TextClass.New()
}

func Text_New() Text {
	return TextClass.New()
}

func (t_ Text) ToggleRuler(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("toggleRuler:"), objc.ExtractPtr(sender))
}

func (t_ Text) ReplaceCharactersInRangeWithRTF(range_ foundation.Range, rtfData []byte) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("replaceCharactersInRange:withRTF:"), range_, rtfData)
}

func (t_ Text) ReplaceCharactersInRangeWithRTFD(range_ foundation.Range, rtfdData []byte) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("replaceCharactersInRange:withRTFD:"), range_, rtfdData)
}

func (t_ Text) ReplaceCharactersInRangeWithString(range_ foundation.Range, string_ string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("replaceCharactersInRange:withString:"), range_, string_)
}

func (t_ Text) SelectAll(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("selectAll:"), objc.ExtractPtr(sender))
}

func (t_ Text) Copy(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("copy:"), objc.ExtractPtr(sender))
}

func (t_ Text) Cut(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("cut:"), objc.ExtractPtr(sender))
}

func (t_ Text) Paste(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("paste:"), objc.ExtractPtr(sender))
}

func (t_ Text) CopyFont(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("copyFont:"), objc.ExtractPtr(sender))
}

func (t_ Text) PasteFont(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("pasteFont:"), objc.ExtractPtr(sender))
}

func (t_ Text) CopyRuler(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("copyRuler:"), objc.ExtractPtr(sender))
}

func (t_ Text) PasteRuler(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("pasteRuler:"), objc.ExtractPtr(sender))
}

func (t_ Text) Delete(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("delete:"), objc.ExtractPtr(sender))
}

func (t_ Text) ChangeFont(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("changeFont:"), objc.ExtractPtr(sender))
}

func (t_ Text) SetFontRange(font IFont, range_ foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setFont:range:"), objc.ExtractPtr(font), range_)
}

func (t_ Text) AlignCenter(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("alignCenter:"), objc.ExtractPtr(sender))
}

func (t_ Text) AlignLeft(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("alignLeft:"), objc.ExtractPtr(sender))
}

func (t_ Text) AlignRight(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("alignRight:"), objc.ExtractPtr(sender))
}

func (t_ Text) SetTextColorRange(color IColor, range_ foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTextColor:range:"), objc.ExtractPtr(color), range_)
}

func (t_ Text) Superscript(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("superscript:"), objc.ExtractPtr(sender))
}

func (t_ Text) Subscript(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("subscript:"), objc.ExtractPtr(sender))
}

func (t_ Text) Unscript(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("unscript:"), objc.ExtractPtr(sender))
}

func (t_ Text) Underline(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("underline:"), objc.ExtractPtr(sender))
}

func (t_ Text) ReadRTFDFromFile(path string) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("readRTFDFromFile:"), path)
	return rv
}

func (t_ Text) WriteRTFDToFileAtomically(path string, flag bool) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("writeRTFDToFile:atomically:"), path, flag)
	return rv
}

func (t_ Text) RTFDFromRange(range_ foundation.Range) []byte {
	rv := objc.CallMethod[[]byte](t_, objc.GetSelector("RTFDFromRange:"), range_)
	return rv
}

func (t_ Text) RTFFromRange(range_ foundation.Range) []byte {
	rv := objc.CallMethod[[]byte](t_, objc.GetSelector("RTFFromRange:"), range_)
	return rv
}

func (t_ Text) CheckSpelling(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("checkSpelling:"), objc.ExtractPtr(sender))
}

func (t_ Text) ShowGuessPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("showGuessPanel:"), objc.ExtractPtr(sender))
}

func (t_ Text) SizeToFit() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("sizeToFit"))
}

func (t_ Text) ScrollRangeToVisible(range_ foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("scrollRangeToVisible:"), range_)
}

func (t_ Text) String() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("string"))
	return rv
}

func (t_ Text) SetString(value string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setString:"), value)
}

func (t_ Text) BackgroundColor() Color {
	rv := objc.CallMethod[Color](t_, objc.GetSelector("backgroundColor"))
	return rv
}

func (t_ Text) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (t_ Text) DrawsBackground() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("drawsBackground"))
	return rv
}

func (t_ Text) SetDrawsBackground(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDrawsBackground:"), value)
}

func (t_ Text) IsEditable() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isEditable"))
	return rv
}

func (t_ Text) SetEditable(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setEditable:"), value)
}

func (t_ Text) IsSelectable() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isSelectable"))
	return rv
}

func (t_ Text) SetSelectable(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSelectable:"), value)
}

func (t_ Text) IsFieldEditor() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isFieldEditor"))
	return rv
}

func (t_ Text) SetFieldEditor(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setFieldEditor:"), value)
}

func (t_ Text) IsRichText() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isRichText"))
	return rv
}

func (t_ Text) SetRichText(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setRichText:"), value)
}

func (t_ Text) ImportsGraphics() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("importsGraphics"))
	return rv
}

func (t_ Text) SetImportsGraphics(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setImportsGraphics:"), value)
}

func (t_ Text) UsesFontPanel() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("usesFontPanel"))
	return rv
}

func (t_ Text) SetUsesFontPanel(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setUsesFontPanel:"), value)
}

func (t_ Text) IsRulerVisible() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isRulerVisible"))
	return rv
}

func (t_ Text) SelectedRange() foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("selectedRange"))
	return rv
}

func (t_ Text) SetSelectedRange(value foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSelectedRange:"), value)
}

func (t_ Text) Font() Font {
	rv := objc.CallMethod[Font](t_, objc.GetSelector("font"))
	return rv
}

func (t_ Text) SetFont(value IFont) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setFont:"), objc.ExtractPtr(value))
}

func (t_ Text) Alignment() TextAlignment {
	rv := objc.CallMethod[TextAlignment](t_, objc.GetSelector("alignment"))
	return rv
}

func (t_ Text) SetAlignment(value TextAlignment) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAlignment:"), value)
}

func (t_ Text) TextColor() Color {
	rv := objc.CallMethod[Color](t_, objc.GetSelector("textColor"))
	return rv
}

func (t_ Text) SetTextColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTextColor:"), objc.ExtractPtr(value))
}

func (t_ Text) BaseWritingDirection() WritingDirection {
	rv := objc.CallMethod[WritingDirection](t_, objc.GetSelector("baseWritingDirection"))
	return rv
}

func (t_ Text) SetBaseWritingDirection(value WritingDirection) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBaseWritingDirection:"), value)
}

func (t_ Text) MaxSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.GetSelector("maxSize"))
	return rv
}

func (t_ Text) SetMaxSize(value foundation.Size) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setMaxSize:"), value)
}

func (t_ Text) MinSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.GetSelector("minSize"))
	return rv
}

func (t_ Text) SetMinSize(value foundation.Size) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setMinSize:"), value)
}

func (t_ Text) IsVerticallyResizable() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isVerticallyResizable"))
	return rv
}

func (t_ Text) SetVerticallyResizable(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setVerticallyResizable:"), value)
}

func (t_ Text) IsHorizontallyResizable() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isHorizontallyResizable"))
	return rv
}

func (t_ Text) SetHorizontallyResizable(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setHorizontallyResizable:"), value)
}

func (t_ Text) Delegate() TextDelegateWrapper {
	rv := objc.CallMethod[TextDelegateWrapper](t_, objc.GetSelector("delegate"))
	return rv
}

func (t_ Text) SetDelegate(value ITextDelegate) {
	po := objc.WrapAsProtocol("NSTextDelegate", value)
	objc.SetAssociatedObject(t_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDelegate:"), po)
}

func (t_ Text) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}
