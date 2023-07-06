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
	return NSStatusBar_SystemStatusBar()
}

func (sb NSStatusBar) StatusItemWithLength(l float64) NSStatusItem {
	return sb.gen_NSStatusBar.StatusItemWithLength(core.CGFloat(l))
}
