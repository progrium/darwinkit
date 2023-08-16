// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextField] class.
var TextFieldClass = _TextFieldClass{objc.GetClass("NSTextField")}

type _TextFieldClass struct {
	objc.Class
}

// An interface definition for the [TextField] class.
type ITextField interface {
	IControl
	SelectText(sender objc.IObject)
	TextShouldBeginEditing(textObject IText) bool
	TextDidChange(notification foundation.INotification)
	TextDidBeginEditing(notification foundation.INotification)
	TextShouldEndEditing(textObject IText) bool
	TextDidEndEditing(notification foundation.INotification)
	MaximumNumberOfLines() int
	SetMaximumNumberOfLines(value int)
	AllowsDefaultTighteningForTruncation() bool
	SetAllowsDefaultTighteningForTruncation(value bool)
	IsBordered() bool
	SetBordered(value bool)
	IsEditable() bool
	SetEditable(value bool)
	IsAutomaticTextCompletionEnabled() bool
	SetAutomaticTextCompletionEnabled(value bool)
	LineBreakStrategy() LineBreakStrategy
	SetLineBreakStrategy(value LineBreakStrategy)
	Delegate() TextFieldDelegateWrapper
	SetDelegate(value PTextFieldDelegate)
	SetDelegateObject(valueObject objc.IObject)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.IAttributedString)
	AllowsCharacterPickerTouchBarItem() bool
	SetAllowsCharacterPickerTouchBarItem(value bool)
	AllowsEditingTextAttributes() bool
	SetAllowsEditingTextAttributes(value bool)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	IsSelectable() bool
	SetSelectable(value bool)
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	TextColor() Color
	SetTextColor(value IColor)
	PlaceholderString() string
	SetPlaceholderString(value string)
	BezelStyle() TextFieldBezelStyle
	SetBezelStyle(value TextFieldBezelStyle)
	IsBezeled() bool
	SetBezeled(value bool)
	PreferredMaxLayoutWidth() float64
	SetPreferredMaxLayoutWidth(value float64)
}

// Text the user can select or edit to send an action message to a target when the user presses the Return key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield?language=objc
type TextField struct {
	Control
}

func TextFieldFrom(ptr unsafe.Pointer) TextField {
	return TextField{
		Control: ControlFrom(ptr),
	}
}

func (tc _TextFieldClass) LabelWithAttributedString(attributedStringValue foundation.IAttributedString) TextField {
	rv := objc.Call[TextField](tc, objc.Sel("labelWithAttributedString:"), objc.Ptr(attributedStringValue))
	return rv
}

// Creates a text field for use as a static label that displays styled text, doesn’t wrap, and doesn’t have selectable text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644658-labelwithattributedstring?language=objc
func TextField_LabelWithAttributedString(attributedStringValue foundation.IAttributedString) TextField {
	return TextFieldClass.LabelWithAttributedString(attributedStringValue)
}

func (tc _TextFieldClass) LabelWithString(stringValue string) TextField {
	rv := objc.Call[TextField](tc, objc.Sel("labelWithString:"), stringValue)
	return rv
}

// Initializes a text field for use as a static label that uses the system default font, doesn’t wrap, and doesn’t have selectable text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644377-labelwithstring?language=objc
func TextField_LabelWithString(stringValue string) TextField {
	return TextFieldClass.LabelWithString(stringValue)
}

func (tc _TextFieldClass) WrappingLabelWithString(stringValue string) TextField {
	rv := objc.Call[TextField](tc, objc.Sel("wrappingLabelWithString:"), stringValue)
	return rv
}

// Initializes a text field for use as a multiline static label with selectable text that uses the system default font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644543-wrappinglabelwithstring?language=objc
func TextField_WrappingLabelWithString(stringValue string) TextField {
	return TextFieldClass.WrappingLabelWithString(stringValue)
}

func (tc _TextFieldClass) TextFieldWithString(stringValue string) TextField {
	rv := objc.Call[TextField](tc, objc.Sel("textFieldWithString:"), stringValue)
	return rv
}

// Initializes a single-line editable text field for user input using the system default font and standard visual appearance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644706-textfieldwithstring?language=objc
func TextField_TextFieldWithString(stringValue string) TextField {
	return TextFieldClass.TextFieldWithString(stringValue)
}

func (tc _TextFieldClass) Alloc() TextField {
	rv := objc.Call[TextField](tc, objc.Sel("alloc"))
	return rv
}

func TextField_Alloc() TextField {
	return TextFieldClass.Alloc()
}

func (tc _TextFieldClass) New() TextField {
	rv := objc.Call[TextField](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextField() TextField {
	return TextFieldClass.New()
}

func (t_ TextField) Init() TextField {
	rv := objc.Call[TextField](t_, objc.Sel("init"))
	return rv
}

func (t_ TextField) InitWithFrame(frameRect foundation.Rect) TextField {
	rv := objc.Call[TextField](t_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func TextField_InitWithFrame(frameRect foundation.Rect) TextField {
	return TextFieldClass.Alloc().InitWithFrame(frameRect)
}

// Ends editing in the text field and, if it’s selectable, selects the entire text content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399430-selecttext?language=objc
func (t_ TextField) SelectText(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("selectText:"), sender)
}

// Requests permission to begin editing a text object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399399-textshouldbeginediting?language=objc
func (t_ TextField) TextShouldBeginEditing(textObject IText) bool {
	rv := objc.Call[bool](t_, objc.Sel("textShouldBeginEditing:"), objc.Ptr(textObject))
	return rv
}

// Posts a notification when the text changes, and forwards the message to the text field’s cell if it responds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399397-textdidchange?language=objc
func (t_ TextField) TextDidChange(notification foundation.INotification) {
	objc.Call[objc.Void](t_, objc.Sel("textDidChange:"), objc.Ptr(notification))
}

// Posts a notification to the default notification center that the text is about to go into edit mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399426-textdidbeginediting?language=objc
func (t_ TextField) TextDidBeginEditing(notification foundation.INotification) {
	objc.Call[objc.Void](t_, objc.Sel("textDidBeginEditing:"), objc.Ptr(notification))
}

// Performs validation on the text field’s new value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399434-textshouldendediting?language=objc
func (t_ TextField) TextShouldEndEditing(textObject IText) bool {
	rv := objc.Call[bool](t_, objc.Sel("textShouldEndEditing:"), objc.Ptr(textObject))
	return rv
}

// Posts a notification when the text is no longer in edit mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399420-textdidendediting?language=objc
func (t_ TextField) TextDidEndEditing(notification foundation.INotification) {
	objc.Call[objc.Void](t_, objc.Sel("textDidEndEditing:"), objc.Ptr(notification))
}

// The maximum number of lines a wrapping text field displays before clipping or truncating the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399424-maximumnumberoflines?language=objc
func (t_ TextField) MaximumNumberOfLines() int {
	rv := objc.Call[int](t_, objc.Sel("maximumNumberOfLines"))
	return rv
}

// The maximum number of lines a wrapping text field displays before clipping or truncating the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399424-maximumnumberoflines?language=objc
func (t_ TextField) SetMaximumNumberOfLines(value int) {
	objc.Call[objc.Void](t_, objc.Sel("setMaximumNumberOfLines:"), value)
}

// A Boolean value that controls whether single-line text fields tighten intercharacter spacing before truncating the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399405-allowsdefaulttighteningfortrunca?language=objc
func (t_ TextField) AllowsDefaultTighteningForTruncation() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsDefaultTighteningForTruncation"))
	return rv
}

// A Boolean value that controls whether single-line text fields tighten intercharacter spacing before truncating the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399405-allowsdefaulttighteningfortrunca?language=objc
func (t_ TextField) SetAllowsDefaultTighteningForTruncation(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsDefaultTighteningForTruncation:"), value)
}

// A Boolean value that controls whether the text field draws a solid black border around its contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399403-bordered?language=objc
func (t_ TextField) IsBordered() bool {
	rv := objc.Call[bool](t_, objc.Sel("isBordered"))
	return rv
}

// A Boolean value that controls whether the text field draws a solid black border around its contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399403-bordered?language=objc
func (t_ TextField) SetBordered(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setBordered:"), value)
}

// A Boolean value that controls whether the user can edit the value in the text field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399407-editable?language=objc
func (t_ TextField) IsEditable() bool {
	rv := objc.Call[bool](t_, objc.Sel("isEditable"))
	return rv
}

// A Boolean value that controls whether the user can edit the value in the text field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399407-editable?language=objc
func (t_ TextField) SetEditable(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setEditable:"), value)
}

// A Boolean value that indicates whether the text field automatically completes text as the user types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/2539554-automatictextcompletionenabled?language=objc
func (t_ TextField) IsAutomaticTextCompletionEnabled() bool {
	rv := objc.Call[bool](t_, objc.Sel("isAutomaticTextCompletionEnabled"))
	return rv
}

// A Boolean value that indicates whether the text field automatically completes text as the user types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/2539554-automatictextcompletionenabled?language=objc
func (t_ TextField) SetAutomaticTextCompletionEnabled(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAutomaticTextCompletionEnabled:"), value)
}

// The strategy that the system uses to break lines when laying out multiple lines of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/3667464-linebreakstrategy?language=objc
func (t_ TextField) LineBreakStrategy() LineBreakStrategy {
	rv := objc.Call[LineBreakStrategy](t_, objc.Sel("lineBreakStrategy"))
	return rv
}

// The strategy that the system uses to break lines when laying out multiple lines of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/3667464-linebreakstrategy?language=objc
func (t_ TextField) SetLineBreakStrategy(value LineBreakStrategy) {
	objc.Call[objc.Void](t_, objc.Sel("setLineBreakStrategy:"), value)
}

// The text field’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399437-delegate?language=objc
func (t_ TextField) Delegate() TextFieldDelegateWrapper {
	rv := objc.Call[TextFieldDelegateWrapper](t_, objc.Sel("delegate"))
	return rv
}

// The text field’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399437-delegate?language=objc
func (t_ TextField) SetDelegate(value PTextFieldDelegate) {
	po0 := objc.WrapAsProtocol("NSTextFieldDelegate", value)
	objc.SetAssociatedObject(t_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), po0)
}

// The text field’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399437-delegate?language=objc
func (t_ TextField) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The attributed string the text field displays when empty to help the user understand the text field’s purpose. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399387-placeholderattributedstring?language=objc
func (t_ TextField) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](t_, objc.Sel("placeholderAttributedString"))
	return rv
}

// The attributed string the text field displays when empty to help the user understand the text field’s purpose. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399387-placeholderattributedstring?language=objc
func (t_ TextField) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	objc.Call[objc.Void](t_, objc.Sel("setPlaceholderAttributedString:"), objc.Ptr(value))
}

// A Boolean value that controls whether the Touch Bar displays the character picker item for rich text fields. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/2539553-allowscharacterpickertouchbarite?language=objc
func (t_ TextField) AllowsCharacterPickerTouchBarItem() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsCharacterPickerTouchBarItem"))
	return rv
}

// A Boolean value that controls whether the Touch Bar displays the character picker item for rich text fields. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/2539553-allowscharacterpickertouchbarite?language=objc
func (t_ TextField) SetAllowsCharacterPickerTouchBarItem(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsCharacterPickerTouchBarItem:"), value)
}

// A Boolean value that controls whether the user can change font attributes of the text field’s string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399401-allowseditingtextattributes?language=objc
func (t_ TextField) AllowsEditingTextAttributes() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsEditingTextAttributes"))
	return rv
}

// A Boolean value that controls whether the user can change font attributes of the text field’s string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399401-allowseditingtextattributes?language=objc
func (t_ TextField) SetAllowsEditingTextAttributes(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsEditingTextAttributes:"), value)
}

// The color of the background the text field’s cell draws behind the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399389-backgroundcolor?language=objc
func (t_ TextField) BackgroundColor() Color {
	rv := objc.Call[Color](t_, objc.Sel("backgroundColor"))
	return rv
}

// The color of the background the text field’s cell draws behind the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399389-backgroundcolor?language=objc
func (t_ TextField) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](t_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// A Boolean value that determines whether the user can select the content of the text field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399422-selectable?language=objc
func (t_ TextField) IsSelectable() bool {
	rv := objc.Call[bool](t_, objc.Sel("isSelectable"))
	return rv
}

// A Boolean value that determines whether the user can select the content of the text field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399422-selectable?language=objc
func (t_ TextField) SetSelectable(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setSelectable:"), value)
}

// A Boolean value that controls whether the user can drag image files into the text field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399428-importsgraphics?language=objc
func (t_ TextField) ImportsGraphics() bool {
	rv := objc.Call[bool](t_, objc.Sel("importsGraphics"))
	return rv
}

// A Boolean value that controls whether the user can drag image files into the text field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399428-importsgraphics?language=objc
func (t_ TextField) SetImportsGraphics(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setImportsGraphics:"), value)
}

// A Boolean value that controls whether the text field’s cell draws a background color behind the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399416-drawsbackground?language=objc
func (t_ TextField) DrawsBackground() bool {
	rv := objc.Call[bool](t_, objc.Sel("drawsBackground"))
	return rv
}

// A Boolean value that controls whether the text field’s cell draws a background color behind the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399416-drawsbackground?language=objc
func (t_ TextField) SetDrawsBackground(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setDrawsBackground:"), value)
}

// The color of the text field’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399409-textcolor?language=objc
func (t_ TextField) TextColor() Color {
	rv := objc.Call[Color](t_, objc.Sel("textColor"))
	return rv
}

// The color of the text field’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399409-textcolor?language=objc
func (t_ TextField) SetTextColor(value IColor) {
	objc.Call[objc.Void](t_, objc.Sel("setTextColor:"), objc.Ptr(value))
}

// The string the text field displays when empty to help the user understand the text field’s purpose. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399391-placeholderstring?language=objc
func (t_ TextField) PlaceholderString() string {
	rv := objc.Call[string](t_, objc.Sel("placeholderString"))
	return rv
}

// The string the text field displays when empty to help the user understand the text field’s purpose. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399391-placeholderstring?language=objc
func (t_ TextField) SetPlaceholderString(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setPlaceholderString:"), value)
}

// The text field’s bezel style, square or rounded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399418-bezelstyle?language=objc
func (t_ TextField) BezelStyle() TextFieldBezelStyle {
	rv := objc.Call[TextFieldBezelStyle](t_, objc.Sel("bezelStyle"))
	return rv
}

// The text field’s bezel style, square or rounded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399418-bezelstyle?language=objc
func (t_ TextField) SetBezelStyle(value TextFieldBezelStyle) {
	objc.Call[objc.Void](t_, objc.Sel("setBezelStyle:"), value)
}

// A Boolean value that controls whether the text field draws a bezeled background around its contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399435-bezeled?language=objc
func (t_ TextField) IsBezeled() bool {
	rv := objc.Call[bool](t_, objc.Sel("isBezeled"))
	return rv
}

// A Boolean value that controls whether the text field draws a bezeled background around its contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399435-bezeled?language=objc
func (t_ TextField) SetBezeled(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setBezeled:"), value)
}

// The maximum width of the text field’s intrinsic content size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399395-preferredmaxlayoutwidth?language=objc
func (t_ TextField) PreferredMaxLayoutWidth() float64 {
	rv := objc.Call[float64](t_, objc.Sel("preferredMaxLayoutWidth"))
	return rv
}

// The maximum width of the text field’s intrinsic content size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1399395-preferredmaxlayoutwidth?language=objc
func (t_ TextField) SetPreferredMaxLayoutWidth(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setPreferredMaxLayoutWidth:"), value)
}
