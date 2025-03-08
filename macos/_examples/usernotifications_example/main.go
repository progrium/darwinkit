package main

import (
	"github.com/progrium/darwinkit/macos"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/usernotifications"
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
		window.SetTitle(foundation.String_StringWithString("UserNotifications Example"))
		window.Center()

		// Create a text view
		textView := appkit.NewTextView()
		scrollView := appkit.NewScrollView()
		objc.Retain(&textView)
		objc.Retain(&scrollView)

		// Create a button to trigger notifications
		button := appkit.NewButtonWithTitle(foundation.String_StringWithString("Send Notification"))
		objc.Retain(&button)
		button.SetFrame(foundation.Rect{
			Origin: foundation.Point{X: 200, Y: 50},
			Size:   foundation.Size{Width: 200, Height: 30},
		})

		// Container view
		containerView := appkit.NewView()
		objc.Retain(&containerView)
		containerView.AddSubview(button)
		containerView.AddSubview(scrollView)

		// Setup text view and scroll view
		scrollView.SetFrame(foundation.Rect{
			Origin: foundation.Point{X: 0, Y: 100},
			Size:   foundation.Size{Width: 600, Height: 300},
		})
		scrollView.SetDocumentView(textView)
		
		window.SetContentView(containerView)

		// Show the window
		window.MakeKeyAndOrderFront(window)

		// Get the notification center
		notificationCenter := usernotifications.CurrentNotificationCenter()
		objc.Retain(&notificationCenter)

		// Request authorization for notifications
		textView.SetString(foundation.String_StringWithString("UserNotifications Demo\n\nRequesting notification permissions..."))
		
		authOptions := usernotifications.AuthorizationOptionBadge | 
					  usernotifications.AuthorizationOptionSound | 
					  usernotifications.AuthorizationOptionAlert
		
		authCompletion := foundation.NewBlockWithVoidBoolError(func(granted bool, err foundation.Error) {
			if err.Pointer() != nil {
				textView.SetString(foundation.String_StringWithString(textView.String().UTF8String() + 
					"\n\nError requesting notification permissions: " + err.LocalizedDescription().UTF8String()))
				return
			}
			
			if granted {
				textView.SetString(foundation.String_StringWithString(textView.String().UTF8String() + 
					"\n\nNotification permissions granted. Click 'Send Notification' to test."))
			} else {
				textView.SetString(foundation.String_StringWithString(textView.String().UTF8String() + 
					"\n\nNotification permissions denied."))
			}
		})
		
		notificationCenter.RequestAuthorizationWithOptionsCompletionHandler(authOptions, authCompletion)

		// Set up notification handler
		button.SetTarget(button)
		button.SetAction(objc.Sel("buttonClicked:"))
		
		buttonClickedBlock := foundation.NewBlockWithObjectSender(func(sender objc.Object) objc.Object {
			textView.SetString(foundation.String_StringWithString(textView.String().UTF8String() + 
				"\n\nCreating notification..."))
			
			// Create notification content
			content := usernotifications.NewMutableNotificationContent()
			objc.Retain(&content)
			content.SetTitle(foundation.String_StringWithString("DarwinKit Notification"))
			content.SetSubtitle(foundation.String_StringWithString("From UserNotifications Example"))
			content.SetBody(foundation.String_StringWithString("This is a test notification from the DarwinKit UserNotifications example."))
			content.SetSound(usernotifications.DefaultSound())
			
			// Create a trigger (5 seconds from now)
			trigger := usernotifications.TriggerWithTimeIntervalRepeats(5.0, false)
			objc.Retain(&trigger)
			
			// Create a notification request
			request := usernotifications.NotificationRequestWithIdentifierContentTrigger(
				foundation.String_StringWithString("com.darwinkit.notification.test"),
				content,
				trigger,
			)
			objc.Retain(&request)
			
			// Schedule the notification
			addCompletion := foundation.NewBlockWithVoidError(func(err foundation.Error) {
				if err.Pointer() != nil {
					textView.SetString(foundation.String_StringWithString(textView.String().UTF8String() + 
						"\nError scheduling notification: " + err.LocalizedDescription().UTF8String()))
				} else {
					textView.SetString(foundation.String_StringWithString(textView.String().UTF8String() + 
						"\nNotification scheduled successfully! You will receive it in 5 seconds."))
				}
			})
			
			notificationCenter.AddNotificationRequestWithCompletionHandler(request, addCompletion)
			
			return nil
		})
		
		button.SetTarget(buttonClickedBlock)
		
		// Create a delegate for the notification center
		notificationDelegateClass := objc.NewClass("NotificationCenterDelegate", "NSObject", 0)
		notificationDelegateClass.AddProtocol("UNUserNotificationCenterDelegate")
		
		notificationDelegateClass.AddMethod("userNotificationCenter:willPresentNotification:withCompletionHandler:", 
			func(self objc.IObject, _cmd objc.SEL, center objc.Object, notification objc.Object, completionHandler objc.Object) {
				// Convert completion handler to a block
				handler := objc.BlockFrom(completionHandler)
				// Present notification while app is in foreground
				handler.Invoke(usernotifications.NotificationPresentationOptionList | 
							 usernotifications.NotificationPresentationOptionBanner | 
							 usernotifications.NotificationPresentationOptionSound)
			})
		
		notificationDelegate := notificationDelegateClass.CreateInstance()
		objc.Retain(&notificationDelegate)
		notificationCenter.SetDelegate(notificationDelegate)

		// Close app when window is closed
		delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
			return true
		})
	})
}