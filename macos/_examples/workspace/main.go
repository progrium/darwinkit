package main

import (
	"fmt"
	"runtime"

	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/objc"
)

func main() {
	runtime.LockOSThread()
	objc.WithAutoreleasePool(func() {
		ws := appkit.Workspace_SharedWorkspace()

		fmt.Println("Running:")
		for _, app := range ws.RunningApplications() {
			fmt.Println(app.LocalizedName())
		}

		fmt.Println("\nFrontmost:")
		frontmost := ws.FrontmostApplication()
		fmt.Println(frontmost.LocalizedName())
	})
}
