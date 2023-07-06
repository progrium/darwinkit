package cocoa

type NSScreen struct {
	gen_NSScreen
}

func NSScreen_Main() NSScreen {
	return NSScreen_MainScreen()
}
