package main

import (
	"runtime"

	"github.com/progrium/macdriver/pkg/bridge"
)

func main() {
	runtime.LockOSThread()
	bridge.Run()
}
