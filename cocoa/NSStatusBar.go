package cocoa

import "github.com/progrium/macdriver/objc"

const (
	NSVariableStatusItemLength = -1.0
	NSSquareStatusItemLength   = -2.0
)

type NSStatusBar struct {
	objc.Object
}

func NSStatusBar_System() NSStatusBar {
	return NSStatusBar{objc.Get("NSStatusBar").Send("systemStatusBar")}
}

func (sb NSStatusBar) StatusItemWithLength(l float64) NSStatusItem {
	return NSStatusItem{sb.Send("statusItemWithLength:", l)}
}

func (sb NSStatusBar) RemoveStatusItem(i NSStatusItem) {
	sb.Send("removeStatusItem:", i)
}
