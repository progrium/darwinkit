// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var FontPanelClass = _FontPanelClass{objc.GetClass("NSFontPanel")}

type _FontPanelClass struct {
	objc.Class
}

type IFontPanel interface {
	IPanel
	ReloadDefaultFontFamilies()
	SetPanelFontIsMultiple(fontObj IFont, flag bool)
	PanelConvertFont(fontObj IFont) Font
	IsEnabled() bool
	SetEnabled(value bool)
	AccessoryView() View
	SetAccessoryView(value IView)
}

type FontPanel struct {
	Panel
}

func MakeFontPanel(ptr unsafe.Pointer) FontPanel {
	return FontPanel{
		Panel: MakePanel(ptr),
	}
}

func (fc _FontPanelClass) WindowWithContentViewController(contentViewController IViewController) FontPanel {
	rv := objc.CallMethod[FontPanel](fc, objc.GetSelector("windowWithContentViewController:"), objc.ExtractPtr(contentViewController))
	return rv
}

func FontPanel_WindowWithContentViewController(contentViewController IViewController) FontPanel {
	return FontPanelClass.WindowWithContentViewController(contentViewController)
}

func (f_ FontPanel) InitWithContentRectStyleMaskBackingDefer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) FontPanel {
	rv := objc.CallMethod[FontPanel](f_, objc.GetSelector("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return rv
}

func FontPanel_InitWithContentRectStyleMaskBackingDefer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) FontPanel {
	return FontPanelClass.Alloc().InitWithContentRectStyleMaskBackingDefer(contentRect, style, backingStoreType, flag)
}

func (f_ FontPanel) InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) FontPanel {
	rv := objc.CallMethod[FontPanel](f_, objc.GetSelector("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.ExtractPtr(screen))
	return rv
}

func FontPanel_InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) FontPanel {
	return FontPanelClass.Alloc().InitWithContentRectStyleMaskBackingDeferScreen(contentRect, style, backingStoreType, flag, screen)
}

func (f_ FontPanel) Init() FontPanel {
	rv := objc.CallMethod[FontPanel](f_, objc.GetSelector("init"))
	return rv
}

func FontPanel_Init() FontPanel {
	return FontPanelClass.Alloc().Init()
}

func (fc _FontPanelClass) Alloc() FontPanel {
	rv := objc.CallMethod[FontPanel](fc, objc.GetSelector("alloc"))
	return rv
}

func FontPanel_Alloc() FontPanel {
	return FontPanelClass.Alloc()
}

func (fc _FontPanelClass) New() FontPanel {
	rv := objc.CallMethod[FontPanel](fc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewFontPanel() FontPanel {
	return FontPanelClass.New()
}

func FontPanel_New() FontPanel {
	return FontPanelClass.New()
}

func (f_ FontPanel) ReloadDefaultFontFamilies() {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("reloadDefaultFontFamilies"))
}

func (f_ FontPanel) SetPanelFontIsMultiple(fontObj IFont, flag bool) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setPanelFont:isMultiple:"), objc.ExtractPtr(fontObj), flag)
}

func (f_ FontPanel) PanelConvertFont(fontObj IFont) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("panelConvertFont:"), objc.ExtractPtr(fontObj))
	return rv
}

func (fc _FontPanelClass) SharedFontPanel() FontPanel {
	rv := objc.CallMethod[FontPanel](fc, objc.GetSelector("sharedFontPanel"))
	return rv
}

func FontPanel_SharedFontPanel() FontPanel {
	return FontPanelClass.SharedFontPanel()
}

func (fc _FontPanelClass) SharedFontPanelExists() bool {
	rv := objc.CallMethod[bool](fc, objc.GetSelector("sharedFontPanelExists"))
	return rv
}

func FontPanel_SharedFontPanelExists() bool {
	return FontPanelClass.SharedFontPanelExists()
}

func (f_ FontPanel) IsEnabled() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isEnabled"))
	return rv
}

func (f_ FontPanel) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setEnabled:"), value)
}

func (f_ FontPanel) AccessoryView() View {
	rv := objc.CallMethod[View](f_, objc.GetSelector("accessoryView"))
	return rv
}

func (f_ FontPanel) SetAccessoryView(value IView) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setAccessoryView:"), objc.ExtractPtr(value))
}
