// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Button] class.
var ButtonClass = _ButtonClass{objc.GetClass("NSButton")}

type _ButtonClass struct {
	objc.Class
}

// An interface definition for the [Button] class.
type IButton interface {
	IControl
	SetNextState()
	MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) foundation.Size
	CompressWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions)
	SetPeriodicDelayInterval(delay float64, interval float64)
	Highlight(flag bool)
	GetPeriodicDelayInterval(delay *float64, interval *float64)
	SetButtonType(type_ ButtonType)
	ImageHugsTitle() bool
	SetImageHugsTitle(value bool)
	AllowsMixedState() bool
	SetAllowsMixedState(value bool)
	AttributedAlternateTitle() foundation.AttributedString
	SetAttributedAlternateTitle(value foundation.IAttributedString)
	KeyEquivalent() string
	SetKeyEquivalent(value string)
	ActiveCompressionOptions() UserInterfaceCompressionOptions
	IsTransparent() bool
	SetTransparent(value bool)
	AlternateTitle() string
	SetAlternateTitle(value string)
	ContentTintColor() Color
	SetContentTintColor(value IColor)
	State() ControlStateValue
	SetState(value ControlStateValue)
	IsBordered() bool
	SetBordered(value bool)
	ImagePosition() CellImagePosition
	SetImagePosition(value CellImagePosition)
	AlternateImage() Image
	SetAlternateImage(value IImage)
	MaxAcceleratorLevel() int
	SetMaxAcceleratorLevel(value int)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
	HasDestructiveAction() bool
	SetHasDestructiveAction(value bool)
	SymbolConfiguration() ImageSymbolConfiguration
	SetSymbolConfiguration(value IImageSymbolConfiguration)
	BezelColor() Color
	SetBezelColor(value IColor)
	Title() string
	SetTitle(value string)
	Sound() Sound
	SetSound(value ISound)
	KeyEquivalentModifierMask() EventModifierFlags
	SetKeyEquivalentModifierMask(value EventModifierFlags)
	BezelStyle() BezelStyle
	SetBezelStyle(value BezelStyle)
	Image() Image
	SetImage(value IImage)
	IsSpringLoaded() bool
	SetSpringLoaded(value bool)
	ShowsBorderOnlyWhileMouseInside() bool
	SetShowsBorderOnlyWhileMouseInside(value bool)
}

// A control that defines an area on the screen that a user clicks to trigger an action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton?language=objc
type Button struct {
	Control
}

func ButtonFrom(ptr unsafe.Pointer) Button {
	return Button{
		Control: ControlFrom(ptr),
	}
}

func (bc _ButtonClass) CheckboxWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) Button {
	rv := objc.Call[Button](bc, objc.Sel("checkboxWithTitle:target:action:"), title, target, action)
	return rv
}

// Creates a standard checkbox with the title you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1644525-checkboxwithtitle?language=objc
func Button_CheckboxWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) Button {
	return ButtonClass.CheckboxWithTitleTargetAction(title, target, action)
}

func (bc _ButtonClass) ButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) Button {
	rv := objc.Call[Button](bc, objc.Sel("buttonWithTitle:target:action:"), title, target, action)
	return rv
}

// Creates a standard push button with the title you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1644256-buttonwithtitle?language=objc
func Button_ButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) Button {
	return ButtonClass.ButtonWithTitleTargetAction(title, target, action)
}

func (bc _ButtonClass) ButtonWithImageTargetAction(image IImage, target objc.IObject, action objc.Selector) Button {
	rv := objc.Call[Button](bc, objc.Sel("buttonWithImage:target:action:"), objc.Ptr(image), target, action)
	return rv
}

// Creates a standard push button with the image you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1644659-buttonwithimage?language=objc
func Button_ButtonWithImageTargetAction(image IImage, target objc.IObject, action objc.Selector) Button {
	return ButtonClass.ButtonWithImageTargetAction(image, target, action)
}

func (bc _ButtonClass) RadioButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) Button {
	rv := objc.Call[Button](bc, objc.Sel("radioButtonWithTitle:target:action:"), title, target, action)
	return rv
}

// Creates a standard radio button with the title you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1644340-radiobuttonwithtitle?language=objc
func Button_RadioButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) Button {
	return ButtonClass.RadioButtonWithTitleTargetAction(title, target, action)
}

func (bc _ButtonClass) Alloc() Button {
	rv := objc.Call[Button](bc, objc.Sel("alloc"))
	return rv
}

func Button_Alloc() Button {
	return ButtonClass.Alloc()
}

func (bc _ButtonClass) New() Button {
	rv := objc.Call[Button](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewButton() Button {
	return ButtonClass.New()
}

func (b_ Button) Init() Button {
	rv := objc.Call[Button](b_, objc.Sel("init"))
	return rv
}

func (b_ Button) InitWithFrame(frameRect foundation.Rect) Button {
	rv := objc.Call[Button](b_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func NewButtonWithFrame(frameRect foundation.Rect) Button {
	instance := ButtonClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Sets the button to its next state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1530594-setnextstate?language=objc
func (b_ Button) SetNextState() {
	objc.Call[objc.Void](b_, objc.Sel("setNextState"))
}

// Returns the minimum size of the button by using the compression options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/2952059-minimumsizewithprioritizedcompre?language=objc
func (b_ Button) MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) foundation.Size {
	rv := objc.Call[foundation.Size](b_, objc.Sel("minimumSizeWithPrioritizedCompressionOptions:"), prioritizedOptions)
	return rv
}

// Sets the priority compression options for this button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/2952060-compresswithprioritizedcompressi?language=objc
func (b_ Button) CompressWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) {
	objc.Call[objc.Void](b_, objc.Sel("compressWithPrioritizedCompressionOptions:"), prioritizedOptions)
}

// Sets the message delay and interval periods for a continuous button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1526612-setperiodicdelay?language=objc
func (b_ Button) SetPeriodicDelayInterval(delay float64, interval float64) {
	objc.Call[objc.Void](b_, objc.Sel("setPeriodicDelay:interval:"), delay, interval)
}

// Highlights (or unhighlights) the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1534156-highlight?language=objc
func (b_ Button) Highlight(flag bool) {
	objc.Call[objc.Void](b_, objc.Sel("highlight:"), flag)
}

// Returns by reference the delay and interval periods for a continuous button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1524817-getperiodicdelay?language=objc
func (b_ Button) GetPeriodicDelayInterval(delay *float64, interval *float64) {
	objc.Call[objc.Void](b_, objc.Sel("getPeriodicDelay:interval:"), delay, interval)
}

// Sets the button’s type, which affects its user interface and behavior when clicked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1524983-setbuttontype?language=objc
func (b_ Button) SetButtonType(type_ ButtonType) {
	objc.Call[objc.Void](b_, objc.Sel("setButtonType:"), type_)
}

// A Boolean value that determines how the button’s image and title are positioned together within the button bezel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/2092414-imagehugstitle?language=objc
func (b_ Button) ImageHugsTitle() bool {
	rv := objc.Call[bool](b_, objc.Sel("imageHugsTitle"))
	return rv
}

// A Boolean value that determines how the button’s image and title are positioned together within the button bezel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/2092414-imagehugstitle?language=objc
func (b_ Button) SetImageHugsTitle(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setImageHugsTitle:"), value)
}

// A Boolean value that indicates whether the button allows a mixed state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1528670-allowsmixedstate?language=objc
func (b_ Button) AllowsMixedState() bool {
	rv := objc.Call[bool](b_, objc.Sel("allowsMixedState"))
	return rv
}

// A Boolean value that indicates whether the button allows a mixed state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1528670-allowsmixedstate?language=objc
func (b_ Button) SetAllowsMixedState(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setAllowsMixedState:"), value)
}

// The title that the button displays as an attributed string when the button is in an on state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1526723-attributedalternatetitle?language=objc
func (b_ Button) AttributedAlternateTitle() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](b_, objc.Sel("attributedAlternateTitle"))
	return rv
}

// The title that the button displays as an attributed string when the button is in an on state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1526723-attributedalternatetitle?language=objc
func (b_ Button) SetAttributedAlternateTitle(value foundation.IAttributedString) {
	objc.Call[objc.Void](b_, objc.Sel("setAttributedAlternateTitle:"), objc.Ptr(value))
}

// The key-equivalent character of the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1525368-keyequivalent?language=objc
func (b_ Button) KeyEquivalent() string {
	rv := objc.Call[string](b_, objc.Sel("keyEquivalent"))
	return rv
}

// The key-equivalent character of the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1525368-keyequivalent?language=objc
func (b_ Button) SetKeyEquivalent(value string) {
	objc.Call[objc.Void](b_, objc.Sel("setKeyEquivalent:"), value)
}

// The compression options active for this button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/2952061-activecompressionoptions?language=objc
func (b_ Button) ActiveCompressionOptions() UserInterfaceCompressionOptions {
	rv := objc.Call[UserInterfaceCompressionOptions](b_, objc.Sel("activeCompressionOptions"))
	return rv
}

// A Boolean value that indicates whether the button is transparent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1529659-transparent?language=objc
func (b_ Button) IsTransparent() bool {
	rv := objc.Call[bool](b_, objc.Sel("isTransparent"))
	return rv
}

// A Boolean value that indicates whether the button is transparent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1529659-transparent?language=objc
func (b_ Button) SetTransparent(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setTransparent:"), value)
}

// The title that the button displays when the button is in an on state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1529588-alternatetitle?language=objc
func (b_ Button) AlternateTitle() string {
	rv := objc.Call[string](b_, objc.Sel("alternateTitle"))
	return rv
}

// The title that the button displays when the button is in an on state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1529588-alternatetitle?language=objc
func (b_ Button) SetAlternateTitle(value string) {
	objc.Call[objc.Void](b_, objc.Sel("setAlternateTitle:"), value)
}

// A tint color to use for the template image and text content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/3000781-contenttintcolor?language=objc
func (b_ Button) ContentTintColor() Color {
	rv := objc.Call[Color](b_, objc.Sel("contentTintColor"))
	return rv
}

// A tint color to use for the template image and text content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/3000781-contenttintcolor?language=objc
func (b_ Button) SetContentTintColor(value IColor) {
	objc.Call[objc.Void](b_, objc.Sel("setContentTintColor:"), objc.Ptr(value))
}

// The button’s state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1528907-state?language=objc
func (b_ Button) State() ControlStateValue {
	rv := objc.Call[ControlStateValue](b_, objc.Sel("state"))
	return rv
}

// The button’s state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1528907-state?language=objc
func (b_ Button) SetState(value ControlStateValue) {
	objc.Call[objc.Void](b_, objc.Sel("setState:"), value)
}

// A Boolean value that determines whether the button has a border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1525565-bordered?language=objc
func (b_ Button) IsBordered() bool {
	rv := objc.Call[bool](b_, objc.Sel("isBordered"))
	return rv
}

// A Boolean value that determines whether the button has a border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1525565-bordered?language=objc
func (b_ Button) SetBordered(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setBordered:"), value)
}

// The position of the button’s image relative to its title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1526778-imageposition?language=objc
func (b_ Button) ImagePosition() CellImagePosition {
	rv := objc.Call[CellImagePosition](b_, objc.Sel("imagePosition"))
	return rv
}

// The position of the button’s image relative to its title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1526778-imageposition?language=objc
func (b_ Button) SetImagePosition(value CellImagePosition) {
	objc.Call[objc.Void](b_, objc.Sel("setImagePosition:"), value)
}

// An alternate image that appears on the button when the button is in an on state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1533935-alternateimage?language=objc
func (b_ Button) AlternateImage() Image {
	rv := objc.Call[Image](b_, objc.Sel("alternateImage"))
	return rv
}

// An alternate image that appears on the button when the button is in an on state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1533935-alternateimage?language=objc
func (b_ Button) SetAlternateImage(value IImage) {
	objc.Call[objc.Void](b_, objc.Sel("setAlternateImage:"), objc.Ptr(value))
}

// An integer value indicating the maximum pressure level for a button of type NSMultiLevelAcceleratorButton. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1534413-maxacceleratorlevel?language=objc
func (b_ Button) MaxAcceleratorLevel() int {
	rv := objc.Call[int](b_, objc.Sel("maxAcceleratorLevel"))
	return rv
}

// An integer value indicating the maximum pressure level for a button of type NSMultiLevelAcceleratorButton. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1534413-maxacceleratorlevel?language=objc
func (b_ Button) SetMaxAcceleratorLevel(value int) {
	objc.Call[objc.Void](b_, objc.Sel("setMaxAcceleratorLevel:"), value)
}

// The title that the button displays in an off state, as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1524640-attributedtitle?language=objc
func (b_ Button) AttributedTitle() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](b_, objc.Sel("attributedTitle"))
	return rv
}

// The title that the button displays in an off state, as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1524640-attributedtitle?language=objc
func (b_ Button) SetAttributedTitle(value foundation.IAttributedString) {
	objc.Call[objc.Void](b_, objc.Sel("setAttributedTitle:"), objc.Ptr(value))
}

// The scaling mode applied to make the cell’s image fit the frame of the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/2202284-imagescaling?language=objc
func (b_ Button) ImageScaling() ImageScaling {
	rv := objc.Call[ImageScaling](b_, objc.Sel("imageScaling"))
	return rv
}

// The scaling mode applied to make the cell’s image fit the frame of the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/2202284-imagescaling?language=objc
func (b_ Button) SetImageScaling(value ImageScaling) {
	objc.Call[objc.Void](b_, objc.Sel("setImageScaling:"), value)
}

// A Boolean value that defines whether a button’s action has a destructive effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/3622469-hasdestructiveaction?language=objc
func (b_ Button) HasDestructiveAction() bool {
	rv := objc.Call[bool](b_, objc.Sel("hasDestructiveAction"))
	return rv
}

// A Boolean value that defines whether a button’s action has a destructive effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/3622469-hasdestructiveaction?language=objc
func (b_ Button) SetHasDestructiveAction(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setHasDestructiveAction:"), value)
}

// The combination of point size, weight, and scale to use when sizing and displaying symbol images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/3667453-symbolconfiguration?language=objc
func (b_ Button) SymbolConfiguration() ImageSymbolConfiguration {
	rv := objc.Call[ImageSymbolConfiguration](b_, objc.Sel("symbolConfiguration"))
	return rv
}

// The combination of point size, weight, and scale to use when sizing and displaying symbol images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/3667453-symbolconfiguration?language=objc
func (b_ Button) SetSymbolConfiguration(value IImageSymbolConfiguration) {
	objc.Call[objc.Void](b_, objc.Sel("setSymbolConfiguration:"), objc.Ptr(value))
}

// The color of the button's bezel, in appearances that support it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/2561000-bezelcolor?language=objc
func (b_ Button) BezelColor() Color {
	rv := objc.Call[Color](b_, objc.Sel("bezelColor"))
	return rv
}

// The color of the button's bezel, in appearances that support it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/2561000-bezelcolor?language=objc
func (b_ Button) SetBezelColor(value IColor) {
	objc.Call[objc.Void](b_, objc.Sel("setBezelColor:"), objc.Ptr(value))
}

// The title displayed on the button when it’s in an off state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1524430-title?language=objc
func (b_ Button) Title() string {
	rv := objc.Call[string](b_, objc.Sel("title"))
	return rv
}

// The title displayed on the button when it’s in an off state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1524430-title?language=objc
func (b_ Button) SetTitle(value string) {
	objc.Call[objc.Void](b_, objc.Sel("setTitle:"), value)
}

// The sound that plays when the user clicks the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1530910-sound?language=objc
func (b_ Button) Sound() Sound {
	rv := objc.Call[Sound](b_, objc.Sel("sound"))
	return rv
}

// The sound that plays when the user clicks the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1530910-sound?language=objc
func (b_ Button) SetSound(value ISound) {
	objc.Call[objc.Void](b_, objc.Sel("setSound:"), objc.Ptr(value))
}

// The mask specifying the modifier keys for the button’s key equivalent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1532670-keyequivalentmodifiermask?language=objc
func (b_ Button) KeyEquivalentModifierMask() EventModifierFlags {
	rv := objc.Call[EventModifierFlags](b_, objc.Sel("keyEquivalentModifierMask"))
	return rv
}

// The mask specifying the modifier keys for the button’s key equivalent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1532670-keyequivalentmodifiermask?language=objc
func (b_ Button) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	objc.Call[objc.Void](b_, objc.Sel("setKeyEquivalentModifierMask:"), value)
}

// The appearance of the button’s border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1527022-bezelstyle?language=objc
func (b_ Button) BezelStyle() BezelStyle {
	rv := objc.Call[BezelStyle](b_, objc.Sel("bezelStyle"))
	return rv
}

// The appearance of the button’s border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1527022-bezelstyle?language=objc
func (b_ Button) SetBezelStyle(value BezelStyle) {
	objc.Call[objc.Void](b_, objc.Sel("setBezelStyle:"), value)
}

// The image that appears on the button when it’s in an off state, or nil if there is no such image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1534221-image?language=objc
func (b_ Button) Image() Image {
	rv := objc.Call[Image](b_, objc.Sel("image"))
	return rv
}

// The image that appears on the button when it’s in an off state, or nil if there is no such image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1534221-image?language=objc
func (b_ Button) SetImage(value IImage) {
	objc.Call[objc.Void](b_, objc.Sel("setImage:"), objc.Ptr(value))
}

// A Boolean value that indicates whether spring loading is enabled for the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1532300-springloaded?language=objc
func (b_ Button) IsSpringLoaded() bool {
	rv := objc.Call[bool](b_, objc.Sel("isSpringLoaded"))
	return rv
}

// A Boolean value that indicates whether spring loading is enabled for the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1532300-springloaded?language=objc
func (b_ Button) SetSpringLoaded(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setSpringLoaded:"), value)
}

// A Boolean value that determines whether the button displays its border only when the pointer is over it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1532248-showsborderonlywhilemouseinside?language=objc
func (b_ Button) ShowsBorderOnlyWhileMouseInside() bool {
	rv := objc.Call[bool](b_, objc.Sel("showsBorderOnlyWhileMouseInside"))
	return rv
}

// A Boolean value that determines whether the button displays its border only when the pointer is over it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1532248-showsborderonlywhilemouseinside?language=objc
func (b_ Button) SetShowsBorderOnlyWhileMouseInside(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setShowsBorderOnlyWhileMouseInside:"), value)
}
