// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Text] class.
var TextClass = _TextClass{objc.GetClass("NSText")}

type _TextClass struct {
	objc.Class
}

// An interface definition for the [Text] class.
type IText interface {
	IView
	Delete(sender objc.IObject)
	CopyFont(sender objc.IObject)
	Copy(sender objc.IObject)
	PasteRuler(sender objc.IObject)
	RTFFromRange(range_ foundation.Range) []byte
	RTFDFromRange(range_ foundation.Range) []byte
	ScrollRangeToVisible(range_ foundation.Range)
	SelectAll(sender objc.IObject)
	AlignRight(sender objc.IObject)
	Underline(sender objc.IObject)
	ShowGuessPanel(sender objc.IObject)
	ChangeFont(sender objc.IObject)
	Unscript(sender objc.IObject)
	SetFontRange(font IFont, range_ foundation.Range)
	AlignLeft(sender objc.IObject)
	ReplaceCharactersInRangeWithRTF(range_ foundation.Range, rtfData []byte)
	Subscript(sender objc.IObject)
	Paste(sender objc.IObject)
	WriteRTFDToFileAtomically(path string, flag bool) bool
	SizeToFit()
	SetTextColorRange(color IColor, range_ foundation.Range)
	ToggleRuler(sender objc.IObject)
	AlignCenter(sender objc.IObject)
	Superscript(sender objc.IObject)
	ReadRTFDFromFile(path string) bool
	PasteFont(sender objc.IObject)
	Cut(sender objc.IObject)
	CopyRuler(sender objc.IObject)
	CheckSpelling(sender objc.IObject)
	MaxSize() foundation.Size
	SetMaxSize(value foundation.Size)
	UsesFontPanel() bool
	SetUsesFontPanel(value bool)
	IsEditable() bool
	SetEditable(value bool)
	Alignment() TextAlignment
	SetAlignment(value TextAlignment)
	IsHorizontallyResizable() bool
	SetHorizontallyResizable(value bool)
	SelectedRange() foundation.Range
	SetSelectedRange(value foundation.Range)
	MinSize() foundation.Size
	SetMinSize(value foundation.Size)
	Delegate() TextDelegateWrapper
	SetDelegate(value PTextDelegate)
	SetDelegateObject(valueObject objc.IObject)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	IsSelectable() bool
	SetSelectable(value bool)
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
	String() string
	SetString(value string)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	Font() Font
	SetFont(value IFont)
	IsFieldEditor() bool
	SetFieldEditor(value bool)
	IsVerticallyResizable() bool
	SetVerticallyResizable(value bool)
	IsRichText() bool
	SetRichText(value bool)
	IsRulerVisible() bool
	TextColor() Color
	SetTextColor(value IColor)
	BaseWritingDirection() WritingDirection
	SetBaseWritingDirection(value WritingDirection)
}

// The most general programmatic interface for objects that manage text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext?language=objc
type Text struct {
	View
}

func TextFrom(ptr unsafe.Pointer) Text {
	return Text{
		View: ViewFrom(ptr),
	}
}

func (t_ Text) InitWithFrame(frameRect foundation.Rect) Text {
	rv := objc.Call[Text](t_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1525191-initwithframe?language=objc
func NewTextWithFrame(frameRect foundation.Rect) Text {
	instance := TextClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

func (tc _TextClass) Alloc() Text {
	rv := objc.Call[Text](tc, objc.Sel("alloc"))
	return rv
}

func Text_Alloc() Text {
	return TextClass.Alloc()
}

func (tc _TextClass) New() Text {
	rv := objc.Call[Text](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewText() Text {
	return TextClass.New()
}

func (t_ Text) Init() Text {
	rv := objc.Call[Text](t_, objc.Sel("init"))
	return rv
}

// This action method deletes the selected text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1524660-delete?language=objc
func (t_ Text) Delete(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("delete:"), sender)
}

// This action method copies the font information for the first character of the selection (or for the insertion point) onto the font pasteboard, as NSFontPboardType. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1531255-copyfont?language=objc
func (t_ Text) CopyFont(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("copyFont:"), sender)
}

// This action method copies the selected text onto the general pasteboard, in as many formats as the receiver supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1525497-copy?language=objc
func (t_ Text) Copy(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("copy:"), sender)
}

// This action method pastes paragraph style information from the ruler pasteboard onto the selected paragraphs of a rich text object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1530420-pasteruler?language=objc
func (t_ Text) PasteRuler(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("pasteRuler:"), sender)
}

// Returns an NSData object that contains an RTF stream corresponding to the characters and attributes within aRange, omitting any attachment characters and attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1527004-rtffromrange?language=objc
func (t_ Text) RTFFromRange(range_ foundation.Range) []byte {
	rv := objc.Call[[]byte](t_, objc.Sel("RTFFromRange:"), range_)
	return rv
}

// Returns an NSData object that contains an RTFD stream corresponding to the characters and attributes within aRange. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1528394-rtfdfromrange?language=objc
func (t_ Text) RTFDFromRange(range_ foundation.Range) []byte {
	rv := objc.Call[[]byte](t_, objc.Sel("RTFDFromRange:"), range_)
	return rv
}

// Scrolls the receiver in its enclosing scroll view so the first characters of aRange are visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1529884-scrollrangetovisible?language=objc
func (t_ Text) ScrollRangeToVisible(range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("scrollRangeToVisible:"), range_)
}

// This action method selects all of the receiver’s text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1527642-selectall?language=objc
func (t_ Text) SelectAll(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("selectAll:"), sender)
}

// This action method applies right alignment to selected paragraphs (or all text if the receiver is a plain text object). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1526477-alignright?language=objc
func (t_ Text) AlignRight(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("alignRight:"), sender)
}

// Adds the underline attribute to the selected text attributes if absent; removes the attribute if present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1534203-underline?language=objc
func (t_ Text) Underline(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("underline:"), sender)
}

// This action method opens the Spelling panel, allowing the user to make a correction during spell checking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1533456-showguesspanel?language=objc
func (t_ Text) ShowGuessPanel(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("showGuessPanel:"), sender)
}

// This action method changes the font of the selection for a rich text object, or of all text for a plain text object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1531459-changefont?language=objc
func (t_ Text) ChangeFont(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("changeFont:"), sender)
}

// This action method removes any superscripting or subscripting from selected text (or all text if the receiver is a plain text object). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1527542-unscript?language=objc
func (t_ Text) Unscript(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("unscript:"), sender)
}

// Sets the font of characters within aRange to aFont. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1526283-setfont?language=objc
func (t_ Text) SetFontRange(font IFont, range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("setFont:range:"), objc.Ptr(font), range_)
}

// This action method applies left alignment to selected paragraphs (or all text if the receiver is a plain text object). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1535705-alignleft?language=objc
func (t_ Text) AlignLeft(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("alignLeft:"), sender)
}

// Replaces the characters in the given range with RTF text interpreted from the given RTF data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1527249-replacecharactersinrange?language=objc
func (t_ Text) ReplaceCharactersInRangeWithRTF(range_ foundation.Range, rtfData []byte) {
	objc.Call[objc.Void](t_, objc.Sel("replaceCharactersInRange:withRTF:"), range_, rtfData)
}

// This action method applies a subscript attribute to selected text (or all text if the receiver is a plain text object), lowering its baseline offset by a predefined amount. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1535373-subscript?language=objc
func (t_ Text) Subscript(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("subscript:"), sender)
}

// This action method pastes text from the general pasteboard at the insertion point or over the selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1527209-paste?language=objc
func (t_ Text) Paste(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("paste:"), sender)
}

// Writes the receiver’s text as RTF with attachments to a file or directory at path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1527085-writertfdtofile?language=objc
func (t_ Text) WriteRTFDToFileAtomically(path string, flag bool) bool {
	rv := objc.Call[bool](t_, objc.Sel("writeRTFDToFile:atomically:"), path, flag)
	return rv
}

// Resizes the receiver to fit its text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1533991-sizetofit?language=objc
func (t_ Text) SizeToFit() {
	objc.Call[objc.Void](t_, objc.Sel("sizeToFit"))
}

// Sets the text color of characters within the specified range to the specified color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1530110-settextcolor?language=objc
func (t_ Text) SetTextColorRange(color IColor, range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("setTextColor:range:"), objc.Ptr(color), range_)
}

// This action method shows or hides the ruler, if the receiver is enclosed in a scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1535773-toggleruler?language=objc
func (t_ Text) ToggleRuler(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("toggleRuler:"), sender)
}

// This action method applies center alignment to selected paragraphs (or all text if the receiver is a plain text object). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1535569-aligncenter?language=objc
func (t_ Text) AlignCenter(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("alignCenter:"), sender)
}

// This action method applies a superscript attribute to selected text (or all text if the receiver is a plain text object), raising its baseline offset by a predefined amount. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1525743-superscript?language=objc
func (t_ Text) Superscript(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("superscript:"), sender)
}

// Attempts to read the RTFD file at path, returning YES if successful and NO if not. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1532564-readrtfdfromfile?language=objc
func (t_ Text) ReadRTFDFromFile(path string) bool {
	rv := objc.Call[bool](t_, objc.Sel("readRTFDFromFile:"), path)
	return rv
}

// This action method pastes font information from the font pasteboard onto the selected text or insertion point of a rich text object, or over all text of a plain text object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1531099-pastefont?language=objc
func (t_ Text) PasteFont(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("pasteFont:"), sender)
}

// This action method deletes the selected text and places it onto the general pasteboard, in as many formats as the receiver supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1524858-cut?language=objc
func (t_ Text) Cut(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("cut:"), sender)
}

// This action method copies the paragraph style information for first selected paragraph onto the ruler pasteboard, as NSRulerPboardType, and expands the selection to paragraph boundaries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1533303-copyruler?language=objc
func (t_ Text) CopyRuler(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("copyRuler:"), sender)
}

// This action method searches for a misspelled word in the receiver’s text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1534926-checkspelling?language=objc
func (t_ Text) CheckSpelling(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("checkSpelling:"), sender)
}

// The receiver’s maximum size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1535900-maxsize?language=objc
func (t_ Text) MaxSize() foundation.Size {
	rv := objc.Call[foundation.Size](t_, objc.Sel("maxSize"))
	return rv
}

// The receiver’s maximum size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1535900-maxsize?language=objc
func (t_ Text) SetMaxSize(value foundation.Size) {
	objc.Call[objc.Void](t_, objc.Sel("setMaxSize:"), value)
}

// A Boolean that controls whether the receiver uses the Font panel and Font menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1527431-usesfontpanel?language=objc
func (t_ Text) UsesFontPanel() bool {
	rv := objc.Call[bool](t_, objc.Sel("usesFontPanel"))
	return rv
}

// A Boolean that controls whether the receiver uses the Font panel and Font menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1527431-usesfontpanel?language=objc
func (t_ Text) SetUsesFontPanel(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setUsesFontPanel:"), value)
}

// A Boolean that controls whether the receiver allows the user to edit its text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1529876-editable?language=objc
func (t_ Text) IsEditable() bool {
	rv := objc.Call[bool](t_, objc.Sel("isEditable"))
	return rv
}

// A Boolean that controls whether the receiver allows the user to edit its text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1529876-editable?language=objc
func (t_ Text) SetEditable(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setEditable:"), value)
}

// The alignment of all the receiver’s text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1534102-alignment?language=objc
func (t_ Text) Alignment() TextAlignment {
	rv := objc.Call[TextAlignment](t_, objc.Sel("alignment"))
	return rv
}

// The alignment of all the receiver’s text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1534102-alignment?language=objc
func (t_ Text) SetAlignment(value TextAlignment) {
	objc.Call[objc.Void](t_, objc.Sel("setAlignment:"), value)
}

// A Boolean that controls whether the receiver changes its width to fit the width of its text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1527489-horizontallyresizable?language=objc
func (t_ Text) IsHorizontallyResizable() bool {
	rv := objc.Call[bool](t_, objc.Sel("isHorizontallyResizable"))
	return rv
}

// A Boolean that controls whether the receiver changes its width to fit the width of its text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1527489-horizontallyresizable?language=objc
func (t_ Text) SetHorizontallyResizable(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setHorizontallyResizable:"), value)
}

// The receiver’s characters within aRange. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1526227-selectedrange?language=objc
func (t_ Text) SelectedRange() foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("selectedRange"))
	return rv
}

// The receiver’s characters within aRange. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1526227-selectedrange?language=objc
func (t_ Text) SetSelectedRange(value foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("setSelectedRange:"), value)
}

// The receiver’s minimum size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1526222-minsize?language=objc
func (t_ Text) MinSize() foundation.Size {
	rv := objc.Call[foundation.Size](t_, objc.Sel("minSize"))
	return rv
}

// The receiver’s minimum size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1526222-minsize?language=objc
func (t_ Text) SetMinSize(value foundation.Size) {
	objc.Call[objc.Void](t_, objc.Sel("setMinSize:"), value)
}

// The receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1529480-delegate?language=objc
func (t_ Text) Delegate() TextDelegateWrapper {
	rv := objc.Call[TextDelegateWrapper](t_, objc.Sel("delegate"))
	return rv
}

// The receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1529480-delegate?language=objc
func (t_ Text) SetDelegate(value PTextDelegate) {
	po0 := objc.WrapAsProtocol("NSTextDelegate", value)
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), po0)
}

// The receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1529480-delegate?language=objc
func (t_ Text) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The receiver’s background color to a given color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1535324-backgroundcolor?language=objc
func (t_ Text) BackgroundColor() Color {
	rv := objc.Call[Color](t_, objc.Sel("backgroundColor"))
	return rv
}

// The receiver’s background color to a given color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1535324-backgroundcolor?language=objc
func (t_ Text) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](t_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// A Boolean that controls whether the receiver allows the user to select its text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1535368-selectable?language=objc
func (t_ Text) IsSelectable() bool {
	rv := objc.Call[bool](t_, objc.Sel("isSelectable"))
	return rv
}

// A Boolean that controls whether the receiver allows the user to select its text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1535368-selectable?language=objc
func (t_ Text) SetSelectable(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setSelectable:"), value)
}

// A Boolean that controls whether the receiver allows the user to import files by dragging. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1531887-importsgraphics?language=objc
func (t_ Text) ImportsGraphics() bool {
	rv := objc.Call[bool](t_, objc.Sel("importsGraphics"))
	return rv
}

// A Boolean that controls whether the receiver allows the user to import files by dragging. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1531887-importsgraphics?language=objc
func (t_ Text) SetImportsGraphics(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setImportsGraphics:"), value)
}

// The characters of the receiver’s text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1528601-string?language=objc
func (t_ Text) String() string {
	rv := objc.Call[string](t_, objc.Sel("string"))
	return rv
}

// The characters of the receiver’s text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1528601-string?language=objc
func (t_ Text) SetString(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setString:"), value)
}

// A Boolean that controls whether the receiver draws its background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1531772-drawsbackground?language=objc
func (t_ Text) DrawsBackground() bool {
	rv := objc.Call[bool](t_, objc.Sel("drawsBackground"))
	return rv
}

// A Boolean that controls whether the receiver draws its background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1531772-drawsbackground?language=objc
func (t_ Text) SetDrawsBackground(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setDrawsBackground:"), value)
}

// The font of all the receiver’s text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1534646-font?language=objc
func (t_ Text) Font() Font {
	rv := objc.Call[Font](t_, objc.Sel("font"))
	return rv
}

// The font of all the receiver’s text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1534646-font?language=objc
func (t_ Text) SetFont(value IFont) {
	objc.Call[objc.Void](t_, objc.Sel("setFont:"), objc.Ptr(value))
}

// A Boolean that controls whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1533080-fieldeditor?language=objc
func (t_ Text) IsFieldEditor() bool {
	rv := objc.Call[bool](t_, objc.Sel("isFieldEditor"))
	return rv
}

// A Boolean that controls whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1533080-fieldeditor?language=objc
func (t_ Text) SetFieldEditor(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setFieldEditor:"), value)
}

// A Boolean that controls whether the receiver changes its height to fit the height of its text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1535082-verticallyresizable?language=objc
func (t_ Text) IsVerticallyResizable() bool {
	rv := objc.Call[bool](t_, objc.Sel("isVerticallyResizable"))
	return rv
}

// A Boolean that controls whether the receiver changes its height to fit the height of its text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1535082-verticallyresizable?language=objc
func (t_ Text) SetVerticallyResizable(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setVerticallyResizable:"), value)
}

// A Boolean that controls whether the receiver allows the user to apply attributes to specific ranges of the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1531003-richtext?language=objc
func (t_ Text) IsRichText() bool {
	rv := objc.Call[bool](t_, objc.Sel("isRichText"))
	return rv
}

// A Boolean that controls whether the receiver allows the user to apply attributes to specific ranges of the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1531003-richtext?language=objc
func (t_ Text) SetRichText(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setRichText:"), value)
}

// A Boolean value that indicates whether the receiver’s enclosing scroll view shows its ruler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1533732-rulervisible?language=objc
func (t_ Text) IsRulerVisible() bool {
	rv := objc.Call[bool](t_, objc.Sel("isRulerVisible"))
	return rv
}

// The text color of all characters in the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1534875-textcolor?language=objc
func (t_ Text) TextColor() Color {
	rv := objc.Call[Color](t_, objc.Sel("textColor"))
	return rv
}

// The text color of all characters in the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1534875-textcolor?language=objc
func (t_ Text) SetTextColor(value IColor) {
	objc.Call[objc.Void](t_, objc.Sel("setTextColor:"), objc.Ptr(value))
}

// The initial writing direction used to determine the actual writing direction for text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1526206-basewritingdirection?language=objc
func (t_ Text) BaseWritingDirection() WritingDirection {
	rv := objc.Call[WritingDirection](t_, objc.Sel("baseWritingDirection"))
	return rv
}

// The initial writing direction used to determine the actual writing direction for text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1526206-basewritingdirection?language=objc
func (t_ Text) SetBaseWritingDirection(value WritingDirection) {
	objc.Call[objc.Void](t_, objc.Sel("setBaseWritingDirection:"), value)
}
