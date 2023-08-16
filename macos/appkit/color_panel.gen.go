// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ColorPanel] class.
var ColorPanelClass = _ColorPanelClass{objc.GetClass("NSColorPanel")}

type _ColorPanelClass struct {
	objc.Class
}

// An interface definition for the [ColorPanel] class.
type IColorPanel interface {
	IPanel
	AttachColorList(colorList IColorList)
	DetachColorList(colorList IColorList)
	SetAction(selector objc.Selector)
	SetTarget(target objc.IObject)
	Color() Color
	SetColor(value IColor)
	IsContinuous() bool
	SetContinuous(value bool)
	AccessoryView() View
	SetAccessoryView(value IView)
	Alpha() float64
	Mode() ColorPanelMode
	SetMode(value ColorPanelMode)
	ShowsAlpha() bool
	SetShowsAlpha(value bool)
}

// A standard user interface for selecting color in an app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel?language=objc
type ColorPanel struct {
	Panel
}

func ColorPanelFrom(ptr unsafe.Pointer) ColorPanel {
	return ColorPanel{
		Panel: PanelFrom(ptr),
	}
}

func (cc _ColorPanelClass) Alloc() ColorPanel {
	rv := objc.Call[ColorPanel](cc, objc.Sel("alloc"))
	return rv
}

func ColorPanel_Alloc() ColorPanel {
	return ColorPanelClass.Alloc()
}

func (cc _ColorPanelClass) New() ColorPanel {
	rv := objc.Call[ColorPanel](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewColorPanel() ColorPanel {
	return ColorPanelClass.New()
}

func (c_ ColorPanel) Init() ColorPanel {
	rv := objc.Call[ColorPanel](c_, objc.Sel("init"))
	return rv
}

func (cc _ColorPanelClass) WindowWithContentViewController(contentViewController IViewController) ColorPanel {
	rv := objc.Call[ColorPanel](cc, objc.Sel("windowWithContentViewController:"), objc.Ptr(contentViewController))
	return rv
}

// Creates a titled window that contains the specified content view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419551-windowwithcontentviewcontroller?language=objc
func ColorPanel_WindowWithContentViewController(contentViewController IViewController) ColorPanel {
	return ColorPanelClass.WindowWithContentViewController(contentViewController)
}

func (c_ ColorPanel) InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) ColorPanel {
	rv := objc.Call[ColorPanel](c_, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.Ptr(screen))
	return rv
}

// Initializes an allocated window with the specified values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419755-initwithcontentrect?language=objc
func ColorPanel_InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) ColorPanel {
	return ColorPanelClass.Alloc().InitWithContentRectStyleMaskBackingDeferScreen(contentRect, style, backingStoreType, flag, screen)
}

// Adds the list of NSColor objects specified to all the color pickers in the receiver that display color lists by invoking attachColorList: on all color pickers in the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1531970-attachcolorlist?language=objc
func (c_ ColorPanel) AttachColorList(colorList IColorList) {
	objc.Call[objc.Void](c_, objc.Sel("attachColorList:"), objc.Ptr(colorList))
}

// Specifies the color panel’s initial picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1532235-setpickermode?language=objc
func (cc _ColorPanelClass) SetPickerMode(mode ColorPanelMode) {
	objc.Call[objc.Void](cc, objc.Sel("setPickerMode:"), mode)
}

// Specifies the color panel’s initial picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1532235-setpickermode?language=objc
func ColorPanel_SetPickerMode(mode ColorPanelMode) {
	ColorPanelClass.SetPickerMode(mode)
}

// Determines which color selection modes are available in an application’s NSColorPanel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1534004-setpickermask?language=objc
func (cc _ColorPanelClass) SetPickerMask(mask ColorPanelOptions) {
	objc.Call[objc.Void](cc, objc.Sel("setPickerMask:"), mask)
}

// Determines which color selection modes are available in an application’s NSColorPanel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1534004-setpickermask?language=objc
func ColorPanel_SetPickerMask(mask ColorPanelOptions) {
	ColorPanelClass.SetPickerMask(mask)
}

// Removes the list of colors from all the color pickers in the receiver that display color lists by invoking detachColorList: on all color pickers in the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1529357-detachcolorlist?language=objc
func (c_ ColorPanel) DetachColorList(colorList IColorList) {
	objc.Call[objc.Void](c_, objc.Sel("detachColorList:"), objc.Ptr(colorList))
}

// Drags a color into a destination view from the specified source view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1529152-dragcolor?language=objc
func (cc _ColorPanelClass) DragColorWithEventFromView(color IColor, event IEvent, sourceView IView) bool {
	rv := objc.Call[bool](cc, objc.Sel("dragColor:withEvent:fromView:"), objc.Ptr(color), objc.Ptr(event), objc.Ptr(sourceView))
	return rv
}

// Drags a color into a destination view from the specified source view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1529152-dragcolor?language=objc
func ColorPanel_DragColorWithEventFromView(color IColor, event IEvent, sourceView IView) bool {
	return ColorPanelClass.DragColorWithEventFromView(color, event, sourceView)
}

// Sets the color panel's action message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1531244-setaction?language=objc
func (c_ ColorPanel) SetAction(selector objc.Selector) {
	objc.Call[objc.Void](c_, objc.Sel("setAction:"), selector)
}

// Sets the target of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1524753-settarget?language=objc
func (c_ ColorPanel) SetTarget(target objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setTarget:"), target)
}

// The color of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1530835-color?language=objc
func (c_ ColorPanel) Color() Color {
	rv := objc.Call[Color](c_, objc.Sel("color"))
	return rv
}

// The color of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1530835-color?language=objc
func (c_ ColorPanel) SetColor(value IColor) {
	objc.Call[objc.Void](c_, objc.Sel("setColor:"), objc.Ptr(value))
}

// Returns  a Boolean value indicating whether the NSColorPanel has been created already. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1525183-sharedcolorpanelexists?language=objc
func (cc _ColorPanelClass) SharedColorPanelExists() bool {
	rv := objc.Call[bool](cc, objc.Sel("sharedColorPanelExists"))
	return rv
}

// Returns  a Boolean value indicating whether the NSColorPanel has been created already. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1525183-sharedcolorpanelexists?language=objc
func ColorPanel_SharedColorPanelExists() bool {
	return ColorPanelClass.SharedColorPanelExists()
}

// Returns the shared NSColorPanel instance, creating it if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1524482-sharedcolorpanel?language=objc
func (cc _ColorPanelClass) SharedColorPanel() ColorPanel {
	rv := objc.Call[ColorPanel](cc, objc.Sel("sharedColorPanel"))
	return rv
}

// Returns the shared NSColorPanel instance, creating it if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1524482-sharedcolorpanel?language=objc
func ColorPanel_SharedColorPanel() ColorPanel {
	return ColorPanelClass.SharedColorPanel()
}

// A Boolean value indicating whether the receiver continuously sends the action message to the target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1528265-continuous?language=objc
func (c_ ColorPanel) IsContinuous() bool {
	rv := objc.Call[bool](c_, objc.Sel("isContinuous"))
	return rv
}

// A Boolean value indicating whether the receiver continuously sends the action message to the target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1528265-continuous?language=objc
func (c_ ColorPanel) SetContinuous(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setContinuous:"), value)
}

// The accessory view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1526892-accessoryview?language=objc
func (c_ ColorPanel) AccessoryView() View {
	rv := objc.Call[View](c_, objc.Sel("accessoryView"))
	return rv
}

// The accessory view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1526892-accessoryview?language=objc
func (c_ ColorPanel) SetAccessoryView(value IView) {
	objc.Call[objc.Void](c_, objc.Sel("setAccessoryView:"), objc.Ptr(value))
}

// The receiver’s current alpha value based on its opacity slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1526246-alpha?language=objc
func (c_ ColorPanel) Alpha() float64 {
	rv := objc.Call[float64](c_, objc.Sel("alpha"))
	return rv
}

// The mode of the receiver the mode is one of the modes allowed by the color mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1525410-mode?language=objc
func (c_ ColorPanel) Mode() ColorPanelMode {
	rv := objc.Call[ColorPanelMode](c_, objc.Sel("mode"))
	return rv
}

// The mode of the receiver the mode is one of the modes allowed by the color mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1525410-mode?language=objc
func (c_ ColorPanel) SetMode(value ColorPanelMode) {
	objc.Call[objc.Void](c_, objc.Sel("setMode:"), value)
}

// A Boolean value that indicates whether the receiver shows alpha values and an opacity slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1525279-showsalpha?language=objc
func (c_ ColorPanel) ShowsAlpha() bool {
	rv := objc.Call[bool](c_, objc.Sel("showsAlpha"))
	return rv
}

// A Boolean value that indicates whether the receiver shows alpha values and an opacity slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/1525279-showsalpha?language=objc
func (c_ ColorPanel) SetShowsAlpha(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setShowsAlpha:"), value)
}
