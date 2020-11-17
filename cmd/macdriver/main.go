package main

import (
	"runtime"

	"github.com/progrium/macdriver"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	macdriver.Run()
}
