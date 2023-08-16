// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FontPanel] class.
var FontPanelClass = _FontPanelClass{objc.GetClass("NSFontPanel")}

type _FontPanelClass struct {
	objc.Class
}

// An interface definition for the [FontPanel] class.
type IFontPanel interface {
	IPanel
	PanelConvertFont(fontObj IFont) Font
	SetPanelFontIsMultiple(fontObj IFont, flag bool)
	ReloadDefaultFontFamilies()
	AccessoryView() View
	SetAccessoryView(value IView)
	IsEnabled() bool
	SetEnabled(value bool)
}

// The Font panel—a user interface object that displays a list of available fonts, letting the user preview them and change the font used to display text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontpanel?language=objc
type FontPanel struct {
	Panel
}

func FontPanelFrom(ptr unsafe.Pointer) FontPanel {
	return FontPanel{
		Panel: PanelFrom(ptr),
	}
}

func (fc _FontPanelClass) Alloc() FontPanel {
	rv := objc.Call[FontPanel](fc, objc.Sel("alloc"))
	return rv
}

func FontPanel_Alloc() FontPanel {
	return FontPanelClass.Alloc()
}

func (fc _FontPanelClass) New() FontPanel {
	rv := objc.Call[FontPanel](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFontPanel() FontPanel {
	return FontPanelClass.New()
}

func (f_ FontPanel) Init() FontPanel {
	rv := objc.Call[FontPanel](f_, objc.Sel("init"))
	return rv
}

func (fc _FontPanelClass) WindowWithContentViewController(contentViewController IViewController) FontPanel {
	rv := objc.Call[FontPanel](fc, objc.Sel("windowWithContentViewController:"), objc.Ptr(contentViewController))
	return rv
}

// Creates a titled window that contains the specified content view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419551-windowwithcontentviewcontroller?language=objc
func FontPanel_WindowWithContentViewController(contentViewController IViewController) FontPanel {
	return FontPanelClass.WindowWithContentViewController(contentViewController)
}

func (f_ FontPanel) InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) FontPanel {
	rv := objc.Call[FontPanel](f_, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.Ptr(screen))
	return rv
}

// Initializes an allocated window with the specified values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419755-initwithcontentrect?language=objc
func FontPanel_InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) FontPanel {
	return FontPanelClass.Alloc().InitWithContentRectStyleMaskBackingDeferScreen(contentRect, style, backingStoreType, flag, screen)
}

// Converts the specified font using the settings in the receiver, with the aid of the shared NSFontManager if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontpanel/1535338-panelconvertfont?language=objc
func (f_ FontPanel) PanelConvertFont(fontObj IFont) Font {
	rv := objc.Call[Font](f_, objc.Sel("panelConvertFont:"), objc.Ptr(fontObj))
	return rv
}

// Sets the selected font in the receiver to the specified font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontpanel/1532648-setpanelfont?language=objc
func (f_ FontPanel) SetPanelFontIsMultiple(fontObj IFont, flag bool) {
	objc.Call[objc.Void](f_, objc.Sel("setPanelFont:isMultiple:"), objc.Ptr(fontObj), flag)
}

// Triggers a reload to the default state, so that the delegate is called. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontpanel/1535396-reloaddefaultfontfamilies?language=objc
func (f_ FontPanel) ReloadDefaultFontFamilies() {
	objc.Call[objc.Void](f_, objc.Sel("reloadDefaultFontFamilies"))
}

// The specified view as the receiver’s accessory view, allowing you to add custom controls to your application’s Font panel without having to create a subclass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontpanel/1535927-accessoryview?language=objc
func (f_ FontPanel) AccessoryView() View {
	rv := objc.Call[View](f_, objc.Sel("accessoryView"))
	return rv
}

// The specified view as the receiver’s accessory view, allowing you to add custom controls to your application’s Font panel without having to create a subclass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontpanel/1535927-accessoryview?language=objc
func (f_ FontPanel) SetAccessoryView(value IView) {
	objc.Call[objc.Void](f_, objc.Sel("setAccessoryView:"), objc.Ptr(value))
}

// Returns the single NSFontPanel instance for the application, creating it if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontpanel/1527046-sharedfontpanel?language=objc
func (fc _FontPanelClass) SharedFontPanel() FontPanel {
	rv := objc.Call[FontPanel](fc, objc.Sel("sharedFontPanel"))
	return rv
}

// Returns the single NSFontPanel instance for the application, creating it if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontpanel/1527046-sharedfontpanel?language=objc
func FontPanel_SharedFontPanel() FontPanel {
	return FontPanelClass.SharedFontPanel()
}

// Returns YES if the shared Font panel has been created, NO if it hasn’t. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontpanel/1524657-sharedfontpanelexists?language=objc
func (fc _FontPanelClass) SharedFontPanelExists() bool {
	rv := objc.Call[bool](fc, objc.Sel("sharedFontPanelExists"))
	return rv
}

// Returns YES if the shared Font panel has been created, NO if it hasn’t. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontpanel/1524657-sharedfontpanelexists?language=objc
func FontPanel_SharedFontPanelExists() bool {
	return FontPanelClass.SharedFontPanelExists()
}

// A Boolean that shows whether the receiver’s Set button is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontpanel/1526041-enabled?language=objc
func (f_ FontPanel) IsEnabled() bool {
	rv := objc.Call[bool](f_, objc.Sel("isEnabled"))
	return rv
}

// A Boolean that shows whether the receiver’s Set button is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontpanel/1526041-enabled?language=objc
func (f_ FontPanel) SetEnabled(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setEnabled:"), value)
}
