package macdriver

import (
	"fmt"

	appkit "github.com/progrium/macdriver/pkg/ns/AppKit"
	foundation "github.com/progrium/macdriver/pkg/ns/Foundation"
)

func Run() {
	window := appkit.NewNSWindow(
		foundation.NSRect{
			Origin: foundation.NSPoint{10.0, 10.0},
			Size:   foundation.NSSize{500.0, 400.0},
		},
		appkit.NSTitledWindowMask|appkit.NSClosableWindowMask|appkit.NSMiniaturizableWindowMask,
		appkit.NSBackingStoreBuffered,
		false,
	)
	window.SetTitle("Hello world")
	window.MakeKeyAndOrderFront(window)

	app := appkit.NSSharedApplication()
	app.SendMsg("setActivationPolicy:", 0)
	fmt.Println("running...")
	app.Run()
}
