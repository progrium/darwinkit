package main

import (
	"runtime"

	"github.com/progrium/macdriver/bridge"
)

func main() {
	runtime.LockOSThread()
	bridge.Run()
}
