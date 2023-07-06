package cocoa

import (
	"github.com/progrium/macdriver/core"
)

type NSStatusBarButton struct {
	gen_NSStatusBarButton
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
	b.gen_NSStatusBarButton.SetTitle(core.String(s))
}

func (b NSStatusBarButton) ToolTip() string {
	return b.Get("toolTip").String()
}

func (b NSStatusBarButton) SetToolTip(s string) {
	b.Set("toolTip:", core.String(s))
}
