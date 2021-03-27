package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSWindow struct {
	objc.Object
}

var nsWindow = objc.Get("NSWindow")

func NSWindow_New() NSWindow {
	return NSWindow{nsWindow.Alloc().Init()}
}

func NSWindow_WithContentViewController(controller objc.Object) NSWindow {
	return NSWindow{nsWindow.Send("windowWithContentViewController:", controller)}
}

func NSWindow_Init(rect core.NSRect, windowStyle core.NSUInteger, bufferingType NSBackingStoreType, deferCreation bool) NSWindow {
	obj := nsWindow.Alloc().
		Send("initWithContentRect:styleMask:backing:defer:",
			rect, windowStyle, bufferingType, deferCreation)
	return NSWindow{obj}
}

func (w NSWindow) Display() {
	w.Send("display")
}

func (w NSWindow) Center() {
	w.Send("center")
}

func (w NSWindow) MakeKeyAndOrderFront(sender objc.Object) {
	w.Send("makeKeyAndOrderFront:", sender)
}

func (w NSWindow) SetLevel(level int) {
	w.Send("setLevel:", level)
}

func (w NSWindow) Level() int64 {
	return w.Send("level").Int()
}

func (w NSWindow) SetStyleMask(mask int) {
	w.Send("setStyleMask:", mask)
}

func (w NSWindow) StyleMask() int64 {
	return w.Send("styleMask").Int()
}

func (w NSWindow) SetTitle(title string) {
	w.Send("setTitle:", core.String(title))
}

func (w NSWindow) Title() string {
	return w.Send("title").String()
}

func (w NSWindow) SetContentView(view objc.Object) {
	w.Send("setContentView:", view)
}

func (w NSWindow) ContentView() NSView {
	return NSView{w.Send("contentView")}
}

func (w NSWindow) IsVisible() bool {
	return w.Send("isVisible").Bool()
}

func (w NSWindow) Frame() (frame core.NSRect) {
	w.Send("frame", &frame)
	return frame
}

func (w NSWindow) ToggleFullScreen(s objc.Object) {
	w.Send("toggleFullScreen:", s)
}

func (w NSWindow) ContentRectForFrameRect(frameRect core.NSRect) (rect core.NSRect) {
	w.Send("contentRectForFrameRect:", frameRect, &rect)
	return rect
}

func (w NSWindow) SetTitlebarAppearsTransparent(b bool) {
	w.Set("titlebarAppearsTransparent:", b)
}

func (w NSWindow) TitlebarAppearsTransparent() bool {
	return w.Get("titlebarAppearsTransparent").Bool()
}

func (w NSWindow) SetTitleVisibility(v int) {
	w.Set("titleVisibility:", v)
}

func (w NSWindow) TitleVisibility() int64 {
	return w.Get("titleVisibility").Int()
}

func (w NSWindow) SetOpaque(b bool) {
	w.Set("opaque:", b)
}

func (w NSWindow) Opaque() bool {
	return w.Get("opaque").Bool()
}

func (w NSWindow) Close() {
	w.Send("close")
}

func (w NSWindow) SetIgnoresMouseEvents(b bool) {
	w.Set("ignoresMouseEvents:", b)
}

func (w NSWindow) IgnoresMouseEvents() bool {
	return w.Get("ignoresMouseEvents").Bool()
}

func (w NSWindow) SetMovableByWindowBackground(b bool) {
	w.Set("movableByWindowBackground:", b)
}

func (w NSWindow) MovableByWindowBackground() bool {
	return w.Get("movableByWindowBackground").Bool()
}

func (w NSWindow) BackgroundColor() NSColor {
	return NSColor{w.Get("backgroundColor")}
}

func (w NSWindow) SetBackgroundColor(color NSColor) {
	w.Set("backgroundColor:", color)
}

func (w NSWindow) SetFrameDisplay(frame core.NSRect, display bool) {
	w.Send("setFrame:display:", frame, display)
}

func (w NSWindow) CollectionBehavior() int64 {
	return w.Get("collectionBehavior").Int()
}

func (w NSWindow) SetCollectionBehavior(collectionBehavior int) {
	w.Set("collectionBehavior:", collectionBehavior)
}

// SetHasShadow sets a Boolean value that indicates whether the window has a shadow.
// https://developer.apple.com/documentation/appkit/nswindow/1419234-hasshadow?language=objc
func (w NSWindow) SetHasShadow(b bool) {
	w.Set("setHasShadow:", b)
}

// HasShadow returns a Boolean value that indicates whether the window has a shadow.
// https://developer.apple.com/documentation/appkit/nswindow/1419234-hasshadow?language=objc
func (w NSWindow) HasShadow() bool {
	return w.Get("hasShadow").Bool()
}

// OrderOut removes the window from the screen list, which hides the window.
// https://developer.apple.com/documentation/appkit/nswindow/1419660-orderout?language=objc
func (w NSWindow) OrderOut(sender objc.Object) {
	w.Send("orderOut:", sender)
}

// OrderFront moves the window to the front of its level in the screen list, without changing either the key window or the main window.
// https://developer.apple.com/documentation/appkit/nswindow/1419495-orderfront?language=objc
func (w NSWindow) OrderFront(sender objc.Object) {
	w.Send("orderFront:", sender)
}

// OrderBack moves the window to the back of its level in the screen list, without changing either the key window or the main window.
// https://developer.apple.com/documentation/appkit/nswindow/1419204-orderback?language=objc
func (w NSWindow) OrderBack(sender objc.Object) {
	w.Send("orderBack:", sender)
}
