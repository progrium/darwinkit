package cocoa

import (
	"github.com/progrium/macdriver/core"
)

const (
	NSVariableStatusItemLength = -1.0
	NSSquareStatusItemLength   = -2.0
)

type NSStatusBar struct {
	gen_NSStatusBar
}

func NSStatusBar_System() NSStatusBar {
	return NSStatusBar_systemStatusBar()
}

func (sb NSStatusBar) StatusItemWithLength(l float64) NSStatusItem {
	return sb.StatusItemWithLength_(core.CGFloat(l))
}

func (sb NSStatusBar) RemoveStatusItem(i NSStatusItemRef) {
	sb.RemoveStatusItem_(i)
}
