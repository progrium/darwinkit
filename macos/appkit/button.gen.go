// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ButtonClass = _ButtonClass{objc.GetClass("NSButton")}

type _ButtonClass struct {
	objc.Class
}

type IButton interface {
	IControl
	SetButtonType(type_ ButtonType)
	GetPeriodicDelayInterval(delay *float32, interval *float32)
	SetPeriodicDelayInterval(delay float32, interval float32)
	CompressWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions)
	MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) foundation.Size
	SetNextState()
	Highlight(flag bool)
	ContentTintColor() Color
	SetContentTintColor(value IColor)
	HasDestructiveAction() bool
	SetHasDestructiveAction(value bool)
	AlternateTitle() string
	SetAlternateTitle(value string)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	AttributedAlternateTitle() foundation.AttributedString
	SetAttributedAlternateTitle(value foundation.IAttributedString)
	Title() string
	SetTitle(value string)
	SymbolConfiguration() ImageSymbolConfiguration
	SetSymbolConfiguration(value IImageSymbolConfiguration)
	Sound() Sound
	SetSound(value ISound)
	IsSpringLoaded() bool
	SetSpringLoaded(value bool)
	MaxAcceleratorLevel() int
	SetMaxAcceleratorLevel(value int)
	Image() Image
	SetImage(value IImage)
	AlternateImage() Image
	SetAlternateImage(value IImage)
	ImagePosition() CellImagePosition
	SetImagePosition(value CellImagePosition)
	IsBordered() bool
	SetBordered(value bool)
	IsTransparent() bool
	SetTransparent(value bool)
	BezelStyle() BezelStyle
	SetBezelStyle(value BezelStyle)
	BezelColor() Color
	SetBezelColor(value IColor)
	ShowsBorderOnlyWhileMouseInside() bool
	SetShowsBorderOnlyWhileMouseInside(value bool)
	ImageHugsTitle() bool
	SetImageHugsTitle(value bool)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
	ActiveCompressionOptions() UserInterfaceCompressionOptions
	AllowsMixedState() bool
	SetAllowsMixedState(value bool)
	State() ControlStateValue
	SetState(value ControlStateValue)
	KeyEquivalent() string
	SetKeyEquivalent(value string)
	KeyEquivalentModifierMask() EventModifierFlags
	SetKeyEquivalentModifierMask(value EventModifierFlags)
}

type Button struct {
	Control
}

func MakeButton(ptr unsafe.Pointer) Button {
	return Button{
		Control: MakeControl(ptr),
	}
}

func (bc _ButtonClass) CheckboxWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) Button {
	rv := objc.CallMethod[Button](bc, objc.GetSelector("checkboxWithTitle:target:action:"), title, objc.ExtractPtr(target), action)
	return rv
}

func Button_CheckboxWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) Button {
	return ButtonClass.CheckboxWithTitleTargetAction(title, target, action)
}

func (bc _ButtonClass) ButtonWithImageTargetAction(image IImage, target objc.IObject, action objc.Selector) Button {
	rv := objc.CallMethod[Button](bc, objc.GetSelector("buttonWithImage:target:action:"), objc.ExtractPtr(image), objc.ExtractPtr(target), action)
	return rv
}

func Button_ButtonWithImageTargetAction(image IImage, target objc.IObject, action objc.Selector) Button {
	return ButtonClass.ButtonWithImageTargetAction(image, target, action)
}

func (bc _ButtonClass) RadioButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) Button {
	rv := objc.CallMethod[Button](bc, objc.GetSelector("radioButtonWithTitle:target:action:"), title, objc.ExtractPtr(target), action)
	return rv
}

func Button_RadioButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) Button {
	return ButtonClass.RadioButtonWithTitleTargetAction(title, target, action)
}

func (bc _ButtonClass) ButtonWithTitleImageTargetAction(title string, image IImage, target objc.IObject, action objc.Selector) Button {
	rv := objc.CallMethod[Button](bc, objc.GetSelector("buttonWithTitle:image:target:action:"), title, objc.ExtractPtr(image), objc.ExtractPtr(target), action)
	return rv
}

func Button_ButtonWithTitleImageTargetAction(title string, image IImage, target objc.IObject, action objc.Selector) Button {
	return ButtonClass.ButtonWithTitleImageTargetAction(title, image, target, action)
}

func (bc _ButtonClass) ButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) Button {
	rv := objc.CallMethod[Button](bc, objc.GetSelector("buttonWithTitle:target:action:"), title, objc.ExtractPtr(target), action)
	return rv
}

func Button_ButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) Button {
	return ButtonClass.ButtonWithTitleTargetAction(title, target, action)
}

func (b_ Button) InitWithFrame(frameRect foundation.Rect) Button {
	rv := objc.CallMethod[Button](b_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func Button_InitWithFrame(frameRect foundation.Rect) Button {
	return ButtonClass.Alloc().InitWithFrame(frameRect)
}

func (b_ Button) Init() Button {
	rv := objc.CallMethod[Button](b_, objc.GetSelector("init"))
	return rv
}

func Button_Init() Button {
	return ButtonClass.Alloc().Init()
}

func (bc _ButtonClass) Alloc() Button {
	rv := objc.CallMethod[Button](bc, objc.GetSelector("alloc"))
	return rv
}

func Button_Alloc() Button {
	return ButtonClass.Alloc()
}

func (bc _ButtonClass) New() Button {
	rv := objc.CallMethod[Button](bc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewButton() Button {
	return ButtonClass.New()
}

func Button_New() Button {
	return ButtonClass.New()
}

func (b_ Button) SetButtonType(type_ ButtonType) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setButtonType:"), type_)
}

func (b_ Button) GetPeriodicDelayInterval(delay *float32, interval *float32) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("getPeriodicDelay:interval:"), delay, interval)
}

func (b_ Button) SetPeriodicDelayInterval(delay float32, interval float32) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setPeriodicDelay:interval:"), delay, interval)
}

func (b_ Button) CompressWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("compressWithPrioritizedCompressionOptions:"), prioritizedOptions)
}

func (b_ Button) MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) foundation.Size {
	rv := objc.CallMethod[foundation.Size](b_, objc.GetSelector("minimumSizeWithPrioritizedCompressionOptions:"), prioritizedOptions)
	return rv
}

func (b_ Button) SetNextState() {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setNextState"))
}

func (b_ Button) Highlight(flag bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("highlight:"), flag)
}

func (b_ Button) ContentTintColor() Color {
	rv := objc.CallMethod[Color](b_, objc.GetSelector("contentTintColor"))
	return rv
}

func (b_ Button) SetContentTintColor(value IColor) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setContentTintColor:"), objc.ExtractPtr(value))
}

func (b_ Button) HasDestructiveAction() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("hasDestructiveAction"))
	return rv
}

func (b_ Button) SetHasDestructiveAction(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setHasDestructiveAction:"), value)
}

func (b_ Button) AlternateTitle() string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("alternateTitle"))
	return rv
}

func (b_ Button) SetAlternateTitle(value string) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setAlternateTitle:"), value)
}

func (b_ Button) AttributedTitle() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](b_, objc.GetSelector("attributedTitle"))
	return rv
}

func (b_ Button) SetAttributedTitle(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setAttributedTitle:"), objc.ExtractPtr(value))
}

func (b_ Button) AttributedAlternateTitle() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](b_, objc.GetSelector("attributedAlternateTitle"))
	return rv
}

func (b_ Button) SetAttributedAlternateTitle(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setAttributedAlternateTitle:"), objc.ExtractPtr(value))
}

func (b_ Button) Title() string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("title"))
	return rv
}

func (b_ Button) SetTitle(value string) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setTitle:"), value)
}

func (b_ Button) SymbolConfiguration() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](b_, objc.GetSelector("symbolConfiguration"))
	return rv
}

func (b_ Button) SetSymbolConfiguration(value IImageSymbolConfiguration) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setSymbolConfiguration:"), objc.ExtractPtr(value))
}

func (b_ Button) Sound() Sound {
	rv := objc.CallMethod[Sound](b_, objc.GetSelector("sound"))
	return rv
}

func (b_ Button) SetSound(value ISound) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setSound:"), objc.ExtractPtr(value))
}

func (b_ Button) IsSpringLoaded() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("isSpringLoaded"))
	return rv
}

func (b_ Button) SetSpringLoaded(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setSpringLoaded:"), value)
}

func (b_ Button) MaxAcceleratorLevel() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("maxAcceleratorLevel"))
	return rv
}

func (b_ Button) SetMaxAcceleratorLevel(value int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setMaxAcceleratorLevel:"), value)
}

func (b_ Button) Image() Image {
	rv := objc.CallMethod[Image](b_, objc.GetSelector("image"))
	return rv
}

func (b_ Button) SetImage(value IImage) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setImage:"), objc.ExtractPtr(value))
}

func (b_ Button) AlternateImage() Image {
	rv := objc.CallMethod[Image](b_, objc.GetSelector("alternateImage"))
	return rv
}

func (b_ Button) SetAlternateImage(value IImage) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setAlternateImage:"), objc.ExtractPtr(value))
}

func (b_ Button) ImagePosition() CellImagePosition {
	rv := objc.CallMethod[CellImagePosition](b_, objc.GetSelector("imagePosition"))
	return rv
}

func (b_ Button) SetImagePosition(value CellImagePosition) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setImagePosition:"), value)
}

func (b_ Button) IsBordered() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("isBordered"))
	return rv
}

func (b_ Button) SetBordered(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setBordered:"), value)
}

func (b_ Button) IsTransparent() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("isTransparent"))
	return rv
}

func (b_ Button) SetTransparent(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setTransparent:"), value)
}

func (b_ Button) BezelStyle() BezelStyle {
	rv := objc.CallMethod[BezelStyle](b_, objc.GetSelector("bezelStyle"))
	return rv
}

func (b_ Button) SetBezelStyle(value BezelStyle) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setBezelStyle:"), value)
}

func (b_ Button) BezelColor() Color {
	rv := objc.CallMethod[Color](b_, objc.GetSelector("bezelColor"))
	return rv
}

func (b_ Button) SetBezelColor(value IColor) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setBezelColor:"), objc.ExtractPtr(value))
}

func (b_ Button) ShowsBorderOnlyWhileMouseInside() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("showsBorderOnlyWhileMouseInside"))
	return rv
}

func (b_ Button) SetShowsBorderOnlyWhileMouseInside(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setShowsBorderOnlyWhileMouseInside:"), value)
}

func (b_ Button) ImageHugsTitle() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("imageHugsTitle"))
	return rv
}

func (b_ Button) SetImageHugsTitle(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setImageHugsTitle:"), value)
}

func (b_ Button) ImageScaling() ImageScaling {
	rv := objc.CallMethod[ImageScaling](b_, objc.GetSelector("imageScaling"))
	return rv
}

func (b_ Button) SetImageScaling(value ImageScaling) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setImageScaling:"), value)
}

func (b_ Button) ActiveCompressionOptions() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](b_, objc.GetSelector("activeCompressionOptions"))
	return rv
}

func (b_ Button) AllowsMixedState() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("allowsMixedState"))
	return rv
}

func (b_ Button) SetAllowsMixedState(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setAllowsMixedState:"), value)
}

func (b_ Button) State() ControlStateValue {
	rv := objc.CallMethod[ControlStateValue](b_, objc.GetSelector("state"))
	return rv
}

func (b_ Button) SetState(value ControlStateValue) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setState:"), value)
}

func (b_ Button) KeyEquivalent() string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("keyEquivalent"))
	return rv
}

func (b_ Button) SetKeyEquivalent(value string) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setKeyEquivalent:"), value)
}

func (b_ Button) KeyEquivalentModifierMask() EventModifierFlags {
	rv := objc.CallMethod[EventModifierFlags](b_, objc.GetSelector("keyEquivalentModifierMask"))
	return rv
}

func (b_ Button) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setKeyEquivalentModifierMask:"), value)
}
