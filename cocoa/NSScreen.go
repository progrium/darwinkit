package cocoa

type NSScreen struct {
	gen_NSScreen
}

func NSScreen_Main() NSScreen {
	return NSScreen_mainScreen()
}

func NSScreen_Screens() []NSScreen {
	var result []NSScreen
	screens := NSScreen_screens()
	for i := uint64(0); i < screens.Count(); i++ {
		result = append(result, NSScreen_fromRef(screens.ObjectAtIndex(i)))
	}
	return result
}
