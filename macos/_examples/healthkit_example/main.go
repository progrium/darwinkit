package main

import (
	"github.com/progrium/darwinkit/macos"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/healthkit"
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
		window.SetTitle(foundation.String_StringWithString("HealthKit Example"))
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

		// Check if health data is available
		if !healthkit.IsHealthDataAvailable() {
			textView.SetString(foundation.String_StringWithString("HealthKit data is not available on this device."))
			return
		}

		textView.SetString(foundation.String_StringWithString("HealthKit is available!\n\nRequesting authorization..."))

		// Create health store
		healthStore := healthkit.NewHealthStore()
		objc.Retain(&healthStore)

		// Get types to read and share
		typesToShare := foundation.Set_Set()
		typesToRead := foundation.Set_Set()

		// Add heart rate type to read
		heartRateType := healthkit.QuantityTypeIdentifierHeartRate()
		stepCountType := healthkit.QuantityTypeIdentifierStepCount()
		
		typesToRead.AddObject(heartRateType)
		typesToRead.AddObject(stepCountType)

		// Request authorization
		completion := foundation.NewBlockWithVoidBoolError(func(success bool, err foundation.Error) {
			if !success {
				textView.SetString(foundation.String_StringWithString("HealthKit authorization denied."))
				return
			}

			textView.SetString(foundation.String_StringWithString("HealthKit authorization granted!\n\nYou can now access heart rate and step count data."))
		})

		healthStore.RequestAuthorizationToShareTypes(typesToShare, typesToRead, completion)

		// Close app when window is closed
		delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
			return true
		})
	})
}