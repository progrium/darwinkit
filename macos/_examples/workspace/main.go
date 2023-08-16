package main

import (
	"fmt"

	"github.com/progrium/macdriver/macos/appkit"
)

func main() {
	ws := appkit.Workspace_SharedWorkspace()
	fmt.Println("Running:")
	for _, app := range ws.RunningApplications() {
		fmt.Println(app.LocalizedName())
	}

	fmt.Println("\nFrontmost:")
	frontmost := ws.FrontmostApplication()
	fmt.Println(frontmost.LocalizedName())
}
