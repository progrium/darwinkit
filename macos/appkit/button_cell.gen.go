// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ButtonCell] class.
var ButtonCellClass = _ButtonCellClass{objc.GetClass("NSButtonCell")}

type _ButtonCellClass struct {
	objc.Class
}

// An interface definition for the [ButtonCell] class.
type IButtonCell interface {
	IActionCell
	MouseEntered(event IEvent)
	DrawBezelWithFrameInView(frame foundation.Rect, controlView IView)
	MouseExited(event IEvent)
	DrawTitleWithFrameInView(title foundation.IAttributedString, frame foundation.Rect, controlView IView) foundation.Rect
	DrawImageWithFrameInView(image IImage, frame foundation.Rect, controlView IView)
	SetPeriodicDelayInterval(delay float64, interval float64)
	SetButtonType(type_ ButtonType)
	AttributedAlternateTitle() foundation.AttributedString
	SetAttributedAlternateTitle(value foundation.IAttributedString)
	SetKeyEquivalent(value string)
	IsTransparent() bool
	SetTransparent(value bool)
	AlternateTitle() string
	SetAlternateTitle(value string)
	HighlightsBy() CellStyleMask
	SetHighlightsBy(value CellStyleMask)
	ImageDimsWhenDisabled() bool
	SetImageDimsWhenDisabled(value bool)
	ImagePosition() CellImagePosition
	SetImagePosition(value CellImagePosition)
	AlternateImage() Image
	SetAlternateImage(value IImage)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
	ShowsStateBy() CellStyleMask
	SetShowsStateBy(value CellStyleMask)
	Sound() Sound
	SetSound(value ISound)
	KeyEquivalentModifierMask() EventModifierFlags
	SetKeyEquivalentModifierMask(value EventModifierFlags)
	BezelStyle() BezelStyle
	SetBezelStyle(value BezelStyle)
	ShowsBorderOnlyWhileMouseInside() bool
	SetShowsBorderOnlyWhileMouseInside(value bool)
}

// An object that defines the user interface of a button or other clickable region of a view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell?language=objc
type ButtonCell struct {
	ActionCell
}

func ButtonCellFrom(ptr unsafe.Pointer) ButtonCell {
	return ButtonCell{
		ActionCell: ActionCellFrom(ptr),
	}
}

func (b_ ButtonCell) InitImageCell(image IImage) ButtonCell {
	rv := objc.Call[ButtonCell](b_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1639152-initimagecell?language=objc
func ButtonCell_InitImageCell(image IImage) ButtonCell {
	return ButtonCellClass.Alloc().InitImageCell(image)
}

func (b_ ButtonCell) InitTextCell(string_ string) ButtonCell {
	rv := objc.Call[ButtonCell](b_, objc.Sel("initTextCell:"), string_)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1639134-inittextcell?language=objc
func ButtonCell_InitTextCell(string_ string) ButtonCell {
	return ButtonCellClass.Alloc().InitTextCell(string_)
}

func (bc _ButtonCellClass) Alloc() ButtonCell {
	rv := objc.Call[ButtonCell](bc, objc.Sel("alloc"))
	return rv
}

func ButtonCell_Alloc() ButtonCell {
	return ButtonCellClass.Alloc()
}

func (bc _ButtonCellClass) New() ButtonCell {
	rv := objc.Call[ButtonCell](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewButtonCell() ButtonCell {
	return ButtonCellClass.New()
}

func (b_ ButtonCell) Init() ButtonCell {
	rv := objc.Call[ButtonCell](b_, objc.Sel("init"))
	return rv
}

// Draws the button’s border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1524997-mouseentered?language=objc
func (b_ ButtonCell) MouseEntered(event IEvent) {
	objc.Call[objc.Void](b_, objc.Sel("mouseEntered:"), objc.Ptr(event))
}

// Draws the border of the button using the current bezel style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1524939-drawbezelwithframe?language=objc
func (b_ ButtonCell) DrawBezelWithFrameInView(frame foundation.Rect, controlView IView) {
	objc.Call[objc.Void](b_, objc.Sel("drawBezelWithFrame:inView:"), frame, objc.Ptr(controlView))
}

// Erases the button’s border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1530776-mouseexited?language=objc
func (b_ ButtonCell) MouseExited(event IEvent) {
	objc.Call[objc.Void](b_, objc.Sel("mouseExited:"), objc.Ptr(event))
}

// Draws the button’s title centered vertically in a specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1528861-drawtitle?language=objc
func (b_ ButtonCell) DrawTitleWithFrameInView(title foundation.IAttributedString, frame foundation.Rect, controlView IView) foundation.Rect {
	rv := objc.Call[foundation.Rect](b_, objc.Sel("drawTitle:withFrame:inView:"), objc.Ptr(title), frame, objc.Ptr(controlView))
	return rv
}

// Draws the image associated with the button’s current state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1531792-drawimage?language=objc
func (b_ ButtonCell) DrawImageWithFrameInView(image IImage, frame foundation.Rect, controlView IView) {
	objc.Call[objc.Void](b_, objc.Sel("drawImage:withFrame:inView:"), objc.Ptr(image), frame, objc.Ptr(controlView))
}

// Sets the message delay and interval for the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1525725-setperiodicdelay?language=objc
func (b_ ButtonCell) SetPeriodicDelayInterval(delay float64, interval float64) {
	objc.Call[objc.Void](b_, objc.Sel("setPeriodicDelay:interval:"), delay, interval)
}

// Sets how the button highlights while pressed and how it shows its state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1527474-setbuttontype?language=objc
func (b_ ButtonCell) SetButtonType(type_ ButtonType) {
	objc.Call[objc.Void](b_, objc.Sel("setButtonType:"), type_)
}

// The title displayed by the button when it’s in its alternate state, as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1526922-attributedalternatetitle?language=objc
func (b_ ButtonCell) AttributedAlternateTitle() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](b_, objc.Sel("attributedAlternateTitle"))
	return rv
}

// The title displayed by the button when it’s in its alternate state, as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1526922-attributedalternatetitle?language=objc
func (b_ ButtonCell) SetAttributedAlternateTitle(value foundation.IAttributedString) {
	objc.Call[objc.Void](b_, objc.Sel("setAttributedAlternateTitle:"), objc.Ptr(value))
}

// The button’s key-equivalent character. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1529476-keyequivalent?language=objc
func (b_ ButtonCell) SetKeyEquivalent(value string) {
	objc.Call[objc.Void](b_, objc.Sel("setKeyEquivalent:"), value)
}

// A Boolean value that indicates if the button is transparent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1530887-transparent?language=objc
func (b_ ButtonCell) IsTransparent() bool {
	rv := objc.Call[bool](b_, objc.Sel("isTransparent"))
	return rv
}

// A Boolean value that indicates if the button is transparent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1530887-transparent?language=objc
func (b_ ButtonCell) SetTransparent(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setTransparent:"), value)
}

// The string displayed by the button when it’s in its alternate state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1535382-alternatetitle?language=objc
func (b_ ButtonCell) AlternateTitle() string {
	rv := objc.Call[string](b_, objc.Sel("alternateTitle"))
	return rv
}

// The string displayed by the button when it’s in its alternate state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1535382-alternatetitle?language=objc
func (b_ ButtonCell) SetAlternateTitle(value string) {
	objc.Call[objc.Void](b_, objc.Sel("setAlternateTitle:"), value)
}

// A set of flags that indicate how the button highlights when it receives a mouse-down event (that is, when the button is pressed). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1528459-highlightsby?language=objc
func (b_ ButtonCell) HighlightsBy() CellStyleMask {
	rv := objc.Call[CellStyleMask](b_, objc.Sel("highlightsBy"))
	return rv
}

// A set of flags that indicate how the button highlights when it receives a mouse-down event (that is, when the button is pressed). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1528459-highlightsby?language=objc
func (b_ ButtonCell) SetHighlightsBy(value CellStyleMask) {
	objc.Call[objc.Void](b_, objc.Sel("setHighlightsBy:"), value)
}

// A Boolean value that indicates if the button’s image and text appear “dim” when the button is disabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1534152-imagedimswhendisabled?language=objc
func (b_ ButtonCell) ImageDimsWhenDisabled() bool {
	rv := objc.Call[bool](b_, objc.Sel("imageDimsWhenDisabled"))
	return rv
}

// A Boolean value that indicates if the button’s image and text appear “dim” when the button is disabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1534152-imagedimswhendisabled?language=objc
func (b_ ButtonCell) SetImageDimsWhenDisabled(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setImageDimsWhenDisabled:"), value)
}

// The position of the button’s image relative to its title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1529593-imageposition?language=objc
func (b_ ButtonCell) ImagePosition() CellImagePosition {
	rv := objc.Call[CellImagePosition](b_, objc.Sel("imagePosition"))
	return rv
}

// The position of the button’s image relative to its title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1529593-imageposition?language=objc
func (b_ ButtonCell) SetImagePosition(value CellImagePosition) {
	objc.Call[objc.Void](b_, objc.Sel("setImagePosition:"), value)
}

// The image the button displays in its alternate state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1527064-alternateimage?language=objc
func (b_ ButtonCell) AlternateImage() Image {
	rv := objc.Call[Image](b_, objc.Sel("alternateImage"))
	return rv
}

// The image the button displays in its alternate state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1527064-alternateimage?language=objc
func (b_ ButtonCell) SetAlternateImage(value IImage) {
	objc.Call[objc.Void](b_, objc.Sel("setAlternateImage:"), objc.Ptr(value))
}

// The background color of the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1529743-backgroundcolor?language=objc
func (b_ ButtonCell) BackgroundColor() Color {
	rv := objc.Call[Color](b_, objc.Sel("backgroundColor"))
	return rv
}

// The background color of the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1529743-backgroundcolor?language=objc
func (b_ ButtonCell) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](b_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// The title displayed by the button when it’s in its normal state as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1529303-attributedtitle?language=objc
func (b_ ButtonCell) AttributedTitle() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](b_, objc.Sel("attributedTitle"))
	return rv
}

// The title displayed by the button when it’s in its normal state as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1529303-attributedtitle?language=objc
func (b_ ButtonCell) SetAttributedTitle(value foundation.IAttributedString) {
	objc.Call[objc.Void](b_, objc.Sel("setAttributedTitle:"), objc.Ptr(value))
}

// The scale factor for the button’s image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1535104-imagescaling?language=objc
func (b_ ButtonCell) ImageScaling() ImageScaling {
	rv := objc.Call[ImageScaling](b_, objc.Sel("imageScaling"))
	return rv
}

// The scale factor for the button’s image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1535104-imagescaling?language=objc
func (b_ ButtonCell) SetImageScaling(value ImageScaling) {
	objc.Call[objc.Void](b_, objc.Sel("setImageScaling:"), value)
}

// The flags that indicate how the button cell shows its alternate state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1533225-showsstateby?language=objc
func (b_ ButtonCell) ShowsStateBy() CellStyleMask {
	rv := objc.Call[CellStyleMask](b_, objc.Sel("showsStateBy"))
	return rv
}

// The flags that indicate how the button cell shows its alternate state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1533225-showsstateby?language=objc
func (b_ ButtonCell) SetShowsStateBy(value CellStyleMask) {
	objc.Call[objc.Void](b_, objc.Sel("setShowsStateBy:"), value)
}

// The sound that’s played when the user presses the button (that is during a mouse-down event). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1525955-sound?language=objc
func (b_ ButtonCell) Sound() Sound {
	rv := objc.Call[Sound](b_, objc.Sel("sound"))
	return rv
}

// The sound that’s played when the user presses the button (that is during a mouse-down event). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1525955-sound?language=objc
func (b_ ButtonCell) SetSound(value ISound) {
	objc.Call[objc.Void](b_, objc.Sel("setSound:"), objc.Ptr(value))
}

// The mask that identifies the modifier keys for the button's key equivalent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1528315-keyequivalentmodifiermask?language=objc
func (b_ ButtonCell) KeyEquivalentModifierMask() EventModifierFlags {
	rv := objc.Call[EventModifierFlags](b_, objc.Sel("keyEquivalentModifierMask"))
	return rv
}

// The mask that identifies the modifier keys for the button's key equivalent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1528315-keyequivalentmodifiermask?language=objc
func (b_ ButtonCell) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	objc.Call[objc.Void](b_, objc.Sel("setKeyEquivalentModifierMask:"), value)
}

// The appearance of the button’s border, if it has one. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1528696-bezelstyle?language=objc
func (b_ ButtonCell) BezelStyle() BezelStyle {
	rv := objc.Call[BezelStyle](b_, objc.Sel("bezelStyle"))
	return rv
}

// The appearance of the button’s border, if it has one. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1528696-bezelstyle?language=objc
func (b_ ButtonCell) SetBezelStyle(value BezelStyle) {
	objc.Call[objc.Void](b_, objc.Sel("setBezelStyle:"), value)
}

// A Boolean value that indicates if the button displays its border only when the pointer is over it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1527903-showsborderonlywhilemouseinside?language=objc
func (b_ ButtonCell) ShowsBorderOnlyWhileMouseInside() bool {
	rv := objc.Call[bool](b_, objc.Sel("showsBorderOnlyWhileMouseInside"))
	return rv
}

// A Boolean value that indicates if the button displays its border only when the pointer is over it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1527903-showsborderonlywhilemouseinside?language=objc
func (b_ ButtonCell) SetShowsBorderOnlyWhileMouseInside(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setShowsBorderOnlyWhileMouseInside:"), value)
}
