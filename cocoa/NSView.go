package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSView struct {
	gen_NSView
}

func NSView_Init(frame core.NSRect) NSView {
	return NSView_alloc().InitWithFrame__asNSView(frame)
}

func (v NSView) SetBackgroundColor(color NSColor) {
	v.SetBackgroundColor_(color)
}

func (v NSView) SetWantsLayer(b bool) {
	v.SetWantsLayer_(b)
}

func (v NSView) AddSubviewPositionedRelativeTo(subview NSViewRef, positioned int, relativeTo NSViewRef) {
	v.AddSubview_positioned_relativeTo_(subview, core.NSUInteger(positioned), relativeTo)
}

func (v NSView) AddSubview(subview NSViewRef) {
	v.AddSubview_(subview)
}

func (v NSView) SetFrameOrigin(p core.NSPoint) {
	v.SetFrameOrigin_(p)
}

func (v NSView) SetFrameSize(s core.NSSize) {
	v.SetFrameSize_(s)
}

func (v NSView) SetFrame(r core.NSRect) {
	v.SetFrame_(r)
}

func (v NSView) SetBoundsOrigin(p core.NSPoint) {
	v.SetBoundsOrigin_(p)
}

func (v NSView) SetBoundsSize(s core.NSSize) {
	v.SetBoundsSize_(s)
}

func (v NSView) SetBounds(r core.NSRect) {
	v.SetBounds_(r)
}

func (v NSView) SetOpaque(b bool) {
	// TODO NSView declares this property as `readonly`, do we really need this
	// setter?
	v.Set("opaque:", b)
}

func (v NSView) Opaque() bool {
	return v.IsOpaque()
}

func (v NSView) SetValueForKey(value, key objc.Object) {
	v.Send("setValue:forKey:", value, key)
}
