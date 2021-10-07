package core

type NSThread struct {
	gen_NSThread
}

func NSThread_IsMainThread() bool {
	return NSThread_isMainThread()
}
