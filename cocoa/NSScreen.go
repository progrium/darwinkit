package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSScreen struct {
	objc.Object
}

func NSScreen_Main() NSScreen {
	return NSScreen{objc.Get("NSScreen").Send("mainScreen")}
}

func (s NSScreen) Frame() (frame core.NSRect) {
	s.Send("frame", &frame)
	return frame
}
