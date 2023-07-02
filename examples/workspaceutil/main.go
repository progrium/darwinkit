package main

import (
	"fmt"

	"github.com/progrium/macdriver/cocoa"
)

func main() {
	ws := cocoa.NSWorkspace_sharedWorkspace()
	apps := ws.RunningApplications()
	fmt.Println(apps)

	frontmost := ws.FrontmostApplication()
	fmt.Println(frontmost)
}
