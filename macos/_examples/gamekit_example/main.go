package main

import (
	"github.com/progrium/darwinkit/macos"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/gamekit"
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
		window.SetTitle(foundation.String_StringWithString("GameKit Example"))
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

		// Check if user is authenticated
		localPlayer := gamekit.LocalPlayer()
		objc.Retain(&localPlayer)

		if localPlayer.IsAuthenticated() {
			playerName := localPlayer.DisplayName()
			message := "You are already authenticated with Game Center!\n\nPlayer: " + playerName.UTF8String()
			textView.SetString(foundation.String_StringWithString(message))
		} else {
			textView.SetString(foundation.String_StringWithString("You are not authenticated with Game Center.\n\nTrying to authenticate..."))

			// Authenticate the player
			authCompletion := foundation.NewBlockWithVoidError(func(err foundation.Error) {
				if err.Pointer() != nil {
					errorMessage := "Authentication failed: " + err.LocalizedDescription().UTF8String()
					textView.SetString(foundation.String_StringWithString(errorMessage))
				} else {
					playerName := localPlayer.DisplayName()
					successMessage := "Successfully authenticated with Game Center!\n\nPlayer: " + playerName.UTF8String()
					textView.SetString(foundation.String_StringWithString(successMessage))

					// Create an achievement
					achievement := gamekit.NewAchievementWithIdentifier(foundation.String_StringWithString("example.achievement.login"))
					objc.Retain(&achievement)
					achievement.SetPercentComplete(100.0)
					achievement.SetShowsCompletionBanner(true)

					// Create an array with the achievement
					achievementsArray := foundation.Array_ArrayWithObject(achievement)

					// Report the achievement
					reportCompletion := foundation.NewBlockWithVoidError(func(err foundation.Error) {
						if err.Pointer() != nil {
							reportError := "Failed to report achievement: " + err.LocalizedDescription().UTF8String()
							textView.SetString(foundation.String_StringWithString(reportError))
						} else {
							textView.SetString(foundation.String_StringWithString(successMessage + "\n\nAchievement reported successfully!"))
						}
					})

					gamekit.ReportAchievements(achievementsArray, reportCompletion)
				}
			})

			localPlayer.AuthenticateWithCompletionHandler(authCompletion)
		}

		// Close app when window is closed
		delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
			return true
		})
	})
}