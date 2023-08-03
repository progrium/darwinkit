// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var PanelClass = _PanelClass{objc.GetClass("NSPanel")}

type _PanelClass struct {
	objc.Class
}

type IPanel interface {
	IWindow
	SetFloatingPanel(value bool)
	BecomesKeyOnlyIfNeeded() bool
	SetBecomesKeyOnlyIfNeeded(value bool)
	SetWorksWhenModal(value bool)
}

type Panel struct {
	Window
}

func MakePanel(ptr unsafe.Pointer) Panel {
	return Panel{
		Window: MakeWindow(ptr),
	}
}

func (pc _PanelClass) WindowWithContentViewController(contentViewController IViewController) Panel {
	rv := objc.CallMethod[Panel](pc, objc.GetSelector("windowWithContentViewController:"), objc.ExtractPtr(contentViewController))
	return rv
}

func Panel_WindowWithContentViewController(contentViewController IViewController) Panel {
	return PanelClass.WindowWithContentViewController(contentViewController)
}

func (p_ Panel) InitWithContentRectStyleMaskBackingDefer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) Panel {
	rv := objc.CallMethod[Panel](p_, objc.GetSelector("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return rv
}

func Panel_InitWithContentRectStyleMaskBackingDefer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) Panel {
	return PanelClass.Alloc().InitWithContentRectStyleMaskBackingDefer(contentRect, style, backingStoreType, flag)
}

func (p_ Panel) InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) Panel {
	rv := objc.CallMethod[Panel](p_, objc.GetSelector("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.ExtractPtr(screen))
	return rv
}

func Panel_InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) Panel {
	return PanelClass.Alloc().InitWithContentRectStyleMaskBackingDeferScreen(contentRect, style, backingStoreType, flag, screen)
}

func (p_ Panel) Init() Panel {
	rv := objc.CallMethod[Panel](p_, objc.GetSelector("init"))
	return rv
}

func Panel_Init() Panel {
	return PanelClass.Alloc().Init()
}

func (pc _PanelClass) Alloc() Panel {
	rv := objc.CallMethod[Panel](pc, objc.GetSelector("alloc"))
	return rv
}

func Panel_Alloc() Panel {
	return PanelClass.Alloc()
}

func (pc _PanelClass) New() Panel {
	rv := objc.CallMethod[Panel](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPanel() Panel {
	return PanelClass.New()
}

func Panel_New() Panel {
	return PanelClass.New()
}

func (p_ Panel) SetFloatingPanel(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setFloatingPanel:"), value)
}

func (p_ Panel) BecomesKeyOnlyIfNeeded() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("becomesKeyOnlyIfNeeded"))
	return rv
}

func (p_ Panel) SetBecomesKeyOnlyIfNeeded(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setBecomesKeyOnlyIfNeeded:"), value)
}

func (p_ Panel) SetWorksWhenModal(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setWorksWhenModal:"), value)
}
