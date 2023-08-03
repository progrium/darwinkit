package main

import (
	"fmt"

	"github.com/progrium/macdriver/macos/appkit"
)

func main() {
	ws := appkit.Workspace_SharedWorkspace()
	apps := ws.RunningApplications()
	fmt.Println(apps)

	frontmost := ws.FrontmostApplication()
	fmt.Println(frontmost)
}
