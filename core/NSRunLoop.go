package core

type NSRunLoop struct {
	gen_NSRunLoop
}

func NSRunLoop_Current() NSRunLoop {
	return NSRunLoop_CurrentRunLoop()
}

func NSRunLoop_Main() NSRunLoop {
	return NSRunLoop_MainRunLoop()
}
