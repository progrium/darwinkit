package cocoa

import (
	"github.com/progrium/macdriver/objc"
)

type NSStatusItem struct {
	gen_NSStatusItem
}

func (i NSStatusItem) Button() NSStatusBarButton {
	return NSStatusBarButton_FromRef(i.Get("button"))
}

func (i NSStatusItem) Menu() NSMenu {
	return NSMenu_FromRef(i.Get("menu"))
}

func (i NSStatusItem) Target() objc.Object {
	return i.Get("target")
}

func (i NSStatusItem) SetTarget(obj objc.Object) {
	i.Set("target:", obj)
}

func (i NSStatusItem) Action() objc.Selector {
	return objc.Sel(i.Get("action").String())
}

func (i NSStatusItem) SetAction(sel objc.Selector) {
	i.Set("action:", sel)
}
