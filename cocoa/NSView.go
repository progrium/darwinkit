package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSView struct {
	objc.Object
}

func NSView_Init(frame core.NSRect) NSView {
	return NSView{objc.Get("NSView").Alloc().Send("initWithFrame:", frame)}
}

func (v NSView) Frame() (frame core.NSRect) {
	v.Send("frame", &frame)
	return frame
}

func (v NSView) Bounds() (frame core.NSRect) {
	v.Send("bounds", &frame)
	return frame
}

func (v NSView) BackgroundColor() NSColor {
	return NSColor{v.Get("backgroundColor")}
}

func (v NSView) SetBackgroundColor(color NSColor) {
	v.Set("backgroundColor:", color)
}

func (v NSView) WantsLayer() bool {
	return v.Get("wantsLayer").Bool()
}

func (v NSView) SetWantsLayer(b bool) {
	v.Set("wantsLayer:", b)
}

func (v NSView) Layer() core.CALayer {
	return core.CALayer{Object: v.Get("layer")}
}

func (v NSView) AddSubviewPositionedRelativeTo(subview objc.Object, positioned int, relativeTo objc.Object) {
	v.Send("addSubview:positioned:relativeTo:", subview, positioned, relativeTo)
}

func (v NSView) SetFrameOrigin(p core.NSPoint) {
	v.Set("frameOrigin:", p)
}

func (v NSView) SetFrameSize(s core.NSSize) {
	v.Set("frameSize:", s)
}

func (v NSView) SetBoundsOrigin(p core.NSPoint) {
	v.Set("boundsOrigin:", p)
}

func (v NSView) SetBoundsSize(s core.NSSize) {
	v.Set("boundsSize:", s)
}

func (v NSView) SetOpaque(b bool) {
	v.Set("opaque:", b)
}

func (v NSView) Opaque() bool {
	return v.Get("opaque").Bool()
}

func (v NSView) SetValueForKey(value, key objc.Object) {
	v.Send("setValue:forKey:", value, key)
}
