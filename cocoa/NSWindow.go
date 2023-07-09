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
	return NSWindow_Alloc().Init()
}

func NSWindow_WithContentViewController(controller NSViewControllerRef) NSWindow {
	return NSWindow_WindowWithContentViewController(controller)
}

func NSWindow_Init(rect core.NSRect, windowStyle core.NSUInteger, bufferingType NSBackingStoreType, deferCreation bool) NSWindow {
	return NSWindow_Alloc().InitWithContentRectStyleMaskBackingDefer(
		rect, core.NSUInteger(windowStyle), core.NSUInteger(bufferingType), deferCreation,
	)
}

func (w NSWindow) SetLevel(level int) {
	w.gen_NSWindow.SetLevel(core.NSInteger(level))
}

func (w NSWindow) Level() int64 {
	return int64(w.gen_NSWindow.Level())
}

func (w NSWindow) SetStyleMask(mask uint) {
	w.gen_NSWindow.SetStyleMask(core.NSUInteger(mask))
}

func (w NSWindow) StyleMask() uint {
	return uint(w.gen_NSWindow.StyleMask())
}

func (w NSWindow) SetTitle(title string) {
	w.gen_NSWindow.SetTitle(core.String(title))
}

func (w NSWindow) Title() string {
	return w.gen_NSWindow.Title().String()
}

func (w NSWindow) SetTitleVisibility(v int) {
	w.gen_NSWindow.SetTitleVisibility(core.NSInteger(v))
}

func (w NSWindow) TitleVisibility() int64 {
	return int64(w.gen_NSWindow.TitleVisibility())
}

func (w NSWindow) Opaque() bool {
	return w.IsOpaque()
}

func (w NSWindow) MovableByWindowBackground() bool {
	return w.IsMovableByWindowBackground()
}

func (w NSWindow) CollectionBehavior() uint {
	return uint(w.gen_NSWindow.CollectionBehavior())
}

func (w NSWindow) SetCollectionBehavior(collectionBehavior uint) {
	w.gen_NSWindow.SetCollectionBehavior(core.NSUInteger(collectionBehavior))
}
