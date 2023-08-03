package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/helper/action"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	app := appkit.Application_SharedApplication()

	ad := &appkit.ApplicationDelegate{}
	ad.SetApplicationDidFinishLaunching(func(foundation.Notification) {

		item := appkit.StatusBar_SystemStatusBar().StatusItemWithLength(appkit.VariableStatusItemLength)
		item.Retain()
		item.Button().SetTitle("▶️ Ready")

		nextClicked := make(chan bool)
		go func() {
			state := -1
			timer := 1500
			countdown := false
			for {
				select {
				case <-time.After(1 * time.Second):
					if timer > 0 && countdown {
						timer = timer - 1
					}
					if timer <= 0 && state%2 == 1 {
						state = (state + 1) % 4
					}
				case <-nextClicked:
					state = (state + 1) % 4
					timer = map[int]int{
						0: 1500,
						1: 1500,
						2: 0,
						3: 300,
					}[state]
					if state%2 == 1 {
						countdown = true
					} else {
						countdown = false
					}
				}
				labels := map[int]string{
					0: "▶️ Ready %02d:%02d",
					1: "✴️ Working %02d:%02d",
					2: "✅ Finished %02d:%02d",
					3: "⏸️ Break %02d:%02d",
				}
				// updates to the ui should happen on the main thread to avoid segfaults
				dispatch.GetMainQueue().DispatchAsync(func() {
					item.Button().SetTitle(fmt.Sprintf(labels[state], timer/60, timer%60))
				})
			}
		}()
		nextClicked <- true

		itemNext := appkit.NewMenuItem()
		itemNext.SetTitle("Next")
		action.Set(itemNext, func(sender objc.IObject) {
			nextClicked <- true
		})

		itemQuit := appkit.NewMenuItem()
		itemQuit.SetTitle("Quit")
		itemQuit.SetAction(objc.GetSelector("terminate:"))

		menu := appkit.NewMenu()
		menu.AddItem(itemNext)
		menu.AddItem(itemQuit)
		item.SetMenu(menu)

	})
	ad.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return false
	})
	app.SetDelegate(ad)

	app.Run()

}
