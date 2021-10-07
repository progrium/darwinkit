package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSWindow struct {
	gen_NSWindow
}

var nsWindow = objc.Get("NSWindow")

func NSWindow_New() NSWindow {
	return NSWindow_alloc().Init_asNSWindow()
}

func NSWindow_WithContentViewController(controller NSViewControllerRef) NSWindow {
	return NSWindow_windowWithContentViewController_(controller)
}

func NSWindow_Init(rect core.NSRect, windowStyle core.NSUInteger, bufferingType NSBackingStoreType, deferCreation bool) NSWindow {
	return NSWindow_alloc().InitWithContentRect_styleMask_backing_defer__asNSWindow(
		rect, core.NSUInteger(windowStyle), core.NSUInteger(bufferingType), deferCreation,
	)
}

func (w NSWindow) MakeKeyAndOrderFront(sender objc.Object) {
	w.MakeKeyAndOrderFront_(sender)
}

func (w NSWindow) SetLevel(level int) {
	w.SetLevel_(core.NSInteger(level))
}

func (w NSWindow) Level() int64 {
	return int64(w.gen_NSWindow.Level())
}

func (w NSWindow) SetStyleMask(mask uint) {
	w.SetStyleMask_(core.NSUInteger(mask))
}

func (w NSWindow) StyleMask() uint {
	return uint(w.gen_NSWindow.StyleMask())
}

func (w NSWindow) SetTitle(title string) {
	w.SetTitle_(core.String(title))
}

func (w NSWindow) Title() string {
	return w.gen_NSWindow.Title().String()
}

func (w NSWindow) SetContentView(view NSViewRef) {
	w.SetContentView_(view)
}

func (w NSWindow) ToggleFullScreen(s objc.Object) {
	w.ToggleFullScreen_(s)
}

func (w NSWindow) ContentRectForFrameRect(frameRect core.NSRect) (rect core.NSRect) {
	return w.ContentRectForFrameRect_(frameRect)
}

func (w NSWindow) SetTitlebarAppearsTransparent(b bool) {
	w.SetTitlebarAppearsTransparent_(b)
}

func (w NSWindow) SetTitleVisibility(v int) {
	w.SetTitleVisibility_(core.NSInteger(v))
}

func (w NSWindow) TitleVisibility() int64 {
	return int64(w.gen_NSWindow.TitleVisibility())
}

func (w NSWindow) SetOpaque(b bool) {
	w.SetOpaque_(b)
}

func (w NSWindow) Opaque() bool {
	return w.IsOpaque()
}

func (w NSWindow) SetIgnoresMouseEvents(b bool) {
	w.SetIgnoresMouseEvents_(b)
}

func (w NSWindow) SetMovableByWindowBackground(b bool) {
	w.SetMovableByWindowBackground_(b)
}

func (w NSWindow) MovableByWindowBackground() bool {
	return w.IsMovableByWindowBackground()
}

func (w NSWindow) SetBackgroundColor(color NSColor) {
	w.SetBackgroundColor_(color)
}

func (w NSWindow) SetFrameDisplay(frame core.NSRect, display bool) {
	w.SetFrame_display_(frame, display)
}

func (w NSWindow) CollectionBehavior() uint {
	return uint(w.gen_NSWindow.CollectionBehavior())
}

func (w NSWindow) SetCollectionBehavior(collectionBehavior uint) {
	w.SetCollectionBehavior_(core.NSUInteger(collectionBehavior))
}

// SetHasShadow sets a Boolean value that indicates whether the window has a shadow.
// https://developer.apple.com/documentation/appkit/nswindow/1419234-hasshadow?language=objc
func (w NSWindow) SetHasShadow(b bool) {
	w.SetHasShadow_(b)
}

// OrderOut removes the window from the screen list, which hides the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419660-orderout?language=objc
func (w NSWindow) OrderOut(sender objc.Object) {
	w.OrderOut_(sender)
}

// OrderFront moves the window to the front of its level in the screen list, without changing either the key window or the main window.
// https://developer.apple.com/documentation/appkit/nswindow/1419495-orderfront?language=objc
func (w NSWindow) OrderFront(sender objc.Object) {
	w.OrderFront_(sender)
}

// OrderBack moves the window to the back of its level in the screen list, without changing either the key window or the main window.
// https://developer.apple.com/documentation/appkit/nswindow/1419204-orderback?language=objc
func (w NSWindow) OrderBack(sender objc.Object) {
	w.OrderBack_(sender)
}
