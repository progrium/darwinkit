package main

import (
	"github.com/progrium/darwinkit/macos"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/security"
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
		window.SetTitle(foundation.String_StringWithString("Security Framework Example"))
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

		// Example keychain operations
		serviceName := "DarwinKitDemo"
		accountName := "testuser"
		password := []byte("secretpassword")

		textView.SetString(foundation.String_StringWithString("Security Framework Demo\n\nTesting Keychain Operations...\n"))

		// Try to find an existing password
		var passwordLength uint32
		var retrievedPassword []byte
		var keychainItem security.KeychainItem
		
		status := security.FindGenericPassword(serviceName, accountName, &passwordLength, &retrievedPassword, &keychainItem)
		
		if status == security.ErrSecItemNotFound {
			// Password not found, add it
			textView.SetString(foundation.String_StringWithString(textView.String().UTF8String() + 
				"\nNo existing password found. Adding new password to keychain..."))
			
			status = security.AddGenericPassword(serviceName, accountName, password, &keychainItem)
			
			if status == security.ErrSecSuccess {
				textView.SetString(foundation.String_StringWithString(textView.String().UTF8String() + 
					"\nSuccessfully added password to keychain!"))
				
				// Try to find it again
				status = security.FindGenericPassword(serviceName, accountName, &passwordLength, &retrievedPassword, &keychainItem)
				
				if status == security.ErrSecSuccess {
					textView.SetString(foundation.String_StringWithString(textView.String().UTF8String() + 
						"\nSuccessfully retrieved newly added password from keychain!"))
				} else {
					textView.SetString(foundation.String_StringWithString(textView.String().UTF8String() + 
						"\nFailed to retrieve password after adding. Error code: " + string(status)))
				}
			} else if status == security.ErrSecDuplicateItem {
				textView.SetString(foundation.String_StringWithString(textView.String().UTF8String() + 
					"\nItem already exists in keychain!"))
			} else {
				textView.SetString(foundation.String_StringWithString(textView.String().UTF8String() + 
					"\nFailed to add password to keychain. Error code: " + string(status)))
			}
		} else if status == security.ErrSecSuccess {
			// Password found
			textView.SetString(foundation.String_StringWithString(textView.String().UTF8String() + 
				"\nFound existing password in keychain. Password length: " + string(passwordLength)))
			
			// Delete the password
			deleteStatus := security.DeleteKeychainItem(keychainItem)
			
			if deleteStatus == security.ErrSecSuccess {
				textView.SetString(foundation.String_StringWithString(textView.String().UTF8String() + 
					"\nSuccessfully deleted password from keychain!"))
			} else {
				textView.SetString(foundation.String_StringWithString(textView.String().UTF8String() + 
					"\nFailed to delete password from keychain. Error code: " + string(deleteStatus)))
			}
		} else {
			// Error occurred
			textView.SetString(foundation.String_StringWithString(textView.String().UTF8String() + 
				"\nError finding password in keychain. Error code: " + string(status)))
		}

		// Close app when window is closed
		delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
			return true
		})
	})
}