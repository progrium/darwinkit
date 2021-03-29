package core

import "github.com/progrium/macdriver/objc"

type NSThread struct {
	objc.Object
}

var nsThread = objc.Get("NSThread")

func NSThread_IsMainThread() bool {
	return nsThread.Send("isMainThread") != nil
}
