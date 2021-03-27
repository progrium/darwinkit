package core

import "github.com/progrium/macdriver/objc"

type NSThread struct {
	objc.Object
}

var NSThread_ = objc.Get("NSThread")

func NSThread_IsMainThread() bool {
	return NSThread_.Send("isMainThread") != nil
}
