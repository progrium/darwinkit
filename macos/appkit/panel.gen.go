// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Panel] class.
var PanelClass = _PanelClass{objc.GetClass("NSPanel")}

type _PanelClass struct {
	objc.Class
}

// An interface definition for the [Panel] class.
type IPanel interface {
	IWindow
	BecomesKeyOnlyIfNeeded() bool
	SetBecomesKeyOnlyIfNeeded(value bool)
	SetWorksWhenModal(value bool)
	SetFloatingPanel(value bool)
}

// A special kind of window that typically performs a function that is auxiliary to the main window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspanel?language=objc
type Panel struct {
	Window
}

func PanelFrom(ptr unsafe.Pointer) Panel {
	return Panel{
		Window: WindowFrom(ptr),
	}
}

func (pc _PanelClass) Alloc() Panel {
	rv := objc.Call[Panel](pc, objc.Sel("alloc"))
	return rv
}

func Panel_Alloc() Panel {
	return PanelClass.Alloc()
}

func (pc _PanelClass) New() Panel {
	rv := objc.Call[Panel](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPanel() Panel {
	return PanelClass.New()
}

func (p_ Panel) Init() Panel {
	rv := objc.Call[Panel](p_, objc.Sel("init"))
	return rv
}

func (pc _PanelClass) WindowWithContentViewController(contentViewController IViewController) Panel {
	rv := objc.Call[Panel](pc, objc.Sel("windowWithContentViewController:"), objc.Ptr(contentViewController))
	return rv
}

// Creates a titled window that contains the specified content view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419551-windowwithcontentviewcontroller?language=objc
func Panel_WindowWithContentViewController(contentViewController IViewController) Panel {
	return PanelClass.WindowWithContentViewController(contentViewController)
}

func (p_ Panel) InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) Panel {
	rv := objc.Call[Panel](p_, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.Ptr(screen))
	return rv
}

// Initializes an allocated window with the specified values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419755-initwithcontentrect?language=objc
func NewPanelWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) Panel {
	instance := PanelClass.Alloc().InitWithContentRectStyleMaskBackingDeferScreen(contentRect, style, backingStoreType, flag, screen)
	instance.Autorelease()
	return instance
}

// A Boolean value that indicates whether the receiver becomes the key window only when needed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspanel/1528836-becomeskeyonlyifneeded?language=objc
func (p_ Panel) BecomesKeyOnlyIfNeeded() bool {
	rv := objc.Call[bool](p_, objc.Sel("becomesKeyOnlyIfNeeded"))
	return rv
}

// A Boolean value that indicates whether the receiver becomes the key window only when needed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspanel/1528836-becomeskeyonlyifneeded?language=objc
func (p_ Panel) SetBecomesKeyOnlyIfNeeded(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setBecomesKeyOnlyIfNeeded:"), value)
}

// A Boolean value that indicates whether the panel receives keyboard and mouse events even when some other window is being run modally. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspanel/1525549-workswhenmodal?language=objc
func (p_ Panel) SetWorksWhenModal(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setWorksWhenModal:"), value)
}

// A Boolean value that indicates whether the receiver is a floating panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspanel/1531901-floatingpanel?language=objc
func (p_ Panel) SetFloatingPanel(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setFloatingPanel:"), value)
}
