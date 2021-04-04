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

func NSScreen_Screens() []NSScreen {
	var result []NSScreen
	screens := core.NSArray{objc.Get("NSScreen").Send("screens")}
	for i := uint64(0); i < screens.Count(); i++ {
		result = append(result, NSScreen{screens.ObjectAtIndex(i)})
	}
	return result
}

func (s NSScreen) Frame() (frame core.NSRect) {
	s.Send("frame", &frame)
	return frame
}
