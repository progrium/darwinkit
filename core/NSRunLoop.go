package core

import "github.com/progrium/macdriver/objc"

type NSRunLoop struct {
	objc.Object
}

var nsRunLoop = objc.Get("NSRunLoop")

func NSRunLoop_Current() NSRunLoop {
	return NSRunLoop{nsRunLoop.Send("currentRunLoop")}
}

func NSRunLoop_Main() NSRunLoop {
	return NSRunLoop{nsRunLoop.Send("mainRunLoop")}
}

func (rl NSRunLoop) Run() {
	rl.Send("run")
}
