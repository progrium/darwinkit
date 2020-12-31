package cocoa

import "github.com/progrium/macdriver/core"

type NSStatusBarButton struct {
	NSView
}

func (b NSStatusBarButton) SetImagePosition(state int) {
	b.Set("imagePosition:", state)
}

func (b NSStatusBarButton) ImagePosition() int64 {
	return b.Get("imagePosition").Int()
}

func (b NSStatusBarButton) Title() string {
	return b.Get("title").String()
}

func (b NSStatusBarButton) SetTitle(s string) {
	b.Set("title:", core.String(s))
}

func (b NSStatusBarButton) Image() NSImage {
	return NSImage{b.Get("image")}
}

func (b NSStatusBarButton) SetImage(obj NSImage) {
	b.Set("image:", obj)
}

func (b NSStatusBarButton) ToolTip() string {
	return b.Get("toolTip").String()
}

func (b NSStatusBarButton) SetToolTip(s string) {
	b.Set("toolTip:", core.String(s))
}
