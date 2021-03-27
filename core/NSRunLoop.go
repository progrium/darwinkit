package core

import "github.com/progrium/macdriver/objc"

type NSRunLoop struct {
	objc.Object
}

var NSRunLoop_ = objc.Get("NSRunLoop")

func NSRunLoop_Current() NSRunLoop {
	return NSRunLoop{NSRunLoop_.Send("currentRunLoop")}
}

func NSRunLoop_Main() NSRunLoop {
	return NSRunLoop{NSRunLoop_.Send("mainRunLoop")}
}

func (rl NSRunLoop) Run() {
	rl.Send("run")
}
