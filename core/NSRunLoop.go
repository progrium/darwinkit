package core

type NSRunLoop struct {
	gen_NSRunLoop
}

func NSRunLoop_Current() NSRunLoop {
	return NSRunLoop_currentRunLoop()
}

func NSRunLoop_Main() NSRunLoop {
	return NSRunLoop_mainRunLoop()
}
