package main

import (
	"github.com/progrium/darwinkit/macos"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/homekit"
	"github.com/progrium/darwinkit/objc"
)

func main() {
	macos.RunApp(func(app appkit.Application, delegate *appkit.ApplicationDelegate) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)

		// Create a window
		frame := foundation.Rect{Size: foundation.Size{600, 400}}
		window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(frame,
			appkit.ClosableWindowMask|appkit.TitledWindowMask,
			appkit.BackingStoreBuffered, false)
		objc.Retain(&window)
		window.SetTitle(foundation.String_StringWithString("HomeKit Example"))
		window.Center()

		// Create a text view
		textView := appkit.NewTextView()
		scrollView := appkit.NewScrollView()
		objc.Retain(&textView)
		objc.Retain(&scrollView)

		scrollView.SetDocumentView(textView)
		window.SetContentView(scrollView)

		// Show the window
		window.MakeKeyAndOrderFront(window)

		textView.SetString(foundation.String_StringWithString("HomeKit Demo\n\nInitializing Home Manager..."))

		// Create home manager
		homeManager := homekit.NewHomeManager()
		objc.Retain(&homeManager)

		// Register a delegate for the home manager
		homesChanged := foundation.NewBlockWithVoidObject(func(sender objc.Object) {
			homes := homeManager.Homes()
			
			if homes.Count() == 0 {
				textView.SetString(foundation.String_StringWithString("No homes found.\n\nCreating a new demo home..."))
				
				// Create a new home
				createCompletion := foundation.NewBlockWithVoidObjectError(func(home objc.Object, err foundation.Error) {
					if err.Pointer() != nil {
						errorMessage := "Failed to create home: " + err.LocalizedDescription().UTF8String()
						textView.SetString(foundation.String_StringWithString(errorMessage))
					} else {
						homeName := homekit.Home{home}.Name().UTF8String()
						message := "Created new home: " + homeName
						textView.SetString(foundation.String_StringWithString(message))
					}
				})
				
				homeManager.AddHomeWithNameCompletionHandler(foundation.String_StringWithString("Demo Home"), createCompletion)
			} else {
				var homeInfo string
				for i := uint64(0); i < homes.Count(); i++ {
					home := homekit.Home{homes.ObjectAtIndex(i)}
					homeInfo += "Home: " + home.Name().UTF8String() + "\n"
					
					rooms := home.Rooms()
					homeInfo += "  Rooms: " + foundation.String_StringWithValue(rooms.Count()).UTF8String() + "\n"
					
					for j := uint64(0); j < rooms.Count(); j++ {
						room := homekit.Room{rooms.ObjectAtIndex(j)}
						homeInfo += "    - " + room.Name().UTF8String() + "\n"
					}
					
					accessories := home.Accessories()
					homeInfo += "  Accessories: " + foundation.String_StringWithValue(accessories.Count()).UTF8String() + "\n"
					
					for j := uint64(0); j < accessories.Count(); j++ {
						accessory := homekit.Accessory{accessories.ObjectAtIndex(j)}
						homeInfo += "    - " + accessory.Name().UTF8String() + "\n"
					}
				}
				
				if homeInfo == "" {
					homeInfo = "No home information available."
				}
				
				textView.SetString(foundation.String_StringWithString("Home Manager initialized successfully!\n\n" + homeInfo))
			}
		})
		
		// Create a delegate object for the home manager
		delegateClass := objc.NewClass("HMHomeManagerDelegate", "NSObject", 0)
		delegateClass.AddMethod("homeManagerDidUpdateHomes:", func(self objc.IObject, _cmd objc.SEL, homeManager objc.Object) {
			homesChanged.Invoke(homeManager)
		})
		delegateObj := delegateClass.CreateInstance()
		objc.Retain(&delegateObj)
		
		homeManager.SetDelegate(delegateObj)

		// Close app when window is closed
		delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
			return true
		})
	})
}