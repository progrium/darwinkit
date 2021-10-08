package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSScreen struct {
	gen_NSScreen
}

func NSScreen_Main() NSScreen {
	return NSScreen_mainScreen()
}

func NSScreen_Screens() []NSScreen {
	var result []NSScreen
	screens := core.NSArray{objc.Get("NSScreen").Send("screens")}
	for i := uint64(0); i < screens.Count(); i++ {
		result = append(result, NSScreen_fromRef(screens.ObjectAtIndex(i)))
	}
	return result
}
