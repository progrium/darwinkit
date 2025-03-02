//go:generate go run ../../generate/tools/genmod.go
package screencapturekit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework ScreenCaptureKit
import "C"

import (
	"fmt"
	"strings"
	"time"

	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/objc"
)

const (
	// Error domain for ScreenCaptureKit errors
	SCStreamErrorDomain = "com.apple.ScreenCaptureKit.SCStreamErrorDomain"
	// Common error codes
	SCStreamErrorUserDeclinedPermission = -302
	SCStreamErrorUserDeclinedRecord     = -303
	SCStreamErrorNoMatchingContent      = -304
	SCStreamTCCAuthNotGranted           = -305
	// Additional TCC related error codes
	SCStreamErrorCaptureUnavailable     = -306
	SCStreamErrorPermissionRevoked      = -307
)

var _SCScreenshotManagerClass = objc.GetClass("SCScreenshotManager")

// SCStream represents a stream that captures content from a display or window.
type SCStream struct {
	objc.Object
}

// SCShareableContent represents content that can be captured.
type SCShareableContent struct {
	objc.Object
}

// SCError represents an error that occurred during screen capture.
type SCError struct {
	objc.Object
}

func (e SCError) Error() string {
	if e.IsNil() {
		return ""
	}
	return objc.Call[string](e.Object, objc.Sel("localizedDescription"))
}

// IsTCCError returns true if the error is related to TCC (Transparency, Consent, and Control) permissions.
func (e SCError) IsTCCError() bool {
	if e.IsNil() {
		return false
	}

	// Get the error domain and code
	domain := objc.Call[string](e.Object, objc.Sel("domain"))
	code := objc.Call[int](e.Object, objc.Sel("code"))

	// Check if it's a TCC-related error
	if domain == SCStreamErrorDomain && 
		(code == SCStreamErrorUserDeclinedPermission || 
		code == SCStreamErrorUserDeclinedRecord || 
		code == SCStreamTCCAuthNotGranted ||
		code == SCStreamErrorCaptureUnavailable ||
		code == SCStreamErrorPermissionRevoked) {
		return true
	}

	// Also check the description for permission-related strings
	desc := e.Error()
	return strings.Contains(strings.ToLower(desc), "permission") || 
		   strings.Contains(strings.ToLower(desc), "authoriz") || 
		   strings.Contains(strings.ToLower(desc), "tcc") ||
		   strings.Contains(strings.ToLower(desc), "access") ||
		   strings.Contains(strings.ToLower(desc), "denied")
}

// GetPermissionInstructions returns instructions for how to grant screen recording permissions
func (e SCError) GetPermissionInstructions() string {
	if \!e.IsTCCError() {
		return ""
	}

	return `To grant screen recording permission:
1. Open System Preferences > Security & Privacy > Privacy
2. Select "Screen Recording" from the left panel
3. Check the box next to your application
4. You may need to quit and relaunch your application completely
5. If problems persist, try logging out and back in to your Mac account`
}

// SCContentSharingPicker represents a system UI for selecting shareable content.
type SCContentSharingPicker struct {
	objc.Object
}

// SCScreenshotManager manages screenshot capture.
type SCScreenshotManager struct {
	objc.Object
}

// SCRecordingOutput represents a destination for captured content.
type SCRecordingOutput struct {
	objc.Object
}

// SCStreamConfiguration represents configuration settings for a screen capture stream.
type SCStreamConfiguration struct {
	objc.Object
}

// SCContentFilter represents a filter for determining which content to capture.
type SCContentFilter struct {
	objc.Object
}

// SCDisplay represents a display that can be captured.
type SCDisplay struct {
	objc.Object
}

// SCWindow represents a window that can be captured.
type SCWindow struct {
	objc.Object
}

// SCStreamConfiguration_New returns a new stream configuration with default settings.
func SCStreamConfiguration_New() SCStreamConfiguration {
	alloc := objc.Call[objc.Object](_SCScreenshotManagerClass, objc.Sel("alloc"))
	return SCStreamConfiguration{
		Object: objc.Call[objc.Object](alloc, objc.Sel("init")),
	}
}

// SCScreenshotManager_SharedManager returns the shared screenshot manager instance.
func SCScreenshotManager_SharedManager() SCScreenshotManager {
	return SCScreenshotManager{
		Object: objc.Call[objc.Object](_SCScreenshotManagerClass, objc.Sel("sharedManager")),
	}
}

// CaptureScreenshotWithCompletion takes a screenshot of the entire screen.
// The completion handler is called with the captured image or an error.
func (m SCScreenshotManager) CaptureScreenshotWithCompletion(handler func(appkit.Image, SCError)) {
	block := objc.CreateMallocBlock(func(image objc.Object, err objc.Object) {
		handler(
			appkit.Image{Object: image},
			SCError{Object: err},
		)
	})
	objc.Call[objc.Void](m.Object, objc.Sel("captureScreenshotWithCompletion:"), block)
}

// CaptureScreenshotWithRetry attempts to take a screenshot with retry logic for TCC issues
// It will retry up to maxRetries times with the specified delay between attempts
func (m SCScreenshotManager) CaptureScreenshotWithRetry(maxRetries int, retryDelay time.Duration, handler func(appkit.Image, SCError, bool)) {
	var attempt int
	var lastError SCError

	// Create the retry handler
	retryHandler := func(image appkit.Image, err SCError) {
		attempt++
		
		// If no error or not a TCC error, we're done
		if err.IsNil() || \!err.IsTCCError() {
			handler(image, err, true) // Final result
			return
		}
		
		// Store the error for potential final report
		lastError = err
		
		// Check if we should retry
		if attempt < maxRetries {
			fmt.Printf("TCC permission issue on attempt %d/%d, retrying in %v...\n", 
				attempt, maxRetries, retryDelay)
			
			// Wait and retry
			time.Sleep(retryDelay)
			m.CaptureScreenshotWithCompletion(retryHandler)
		} else {
			// We've run out of retries
			fmt.Printf("TCC permission issues persisted after %d attempts\n", maxRetries)
			handler(image, lastError, false) // Failed after all retries
		}
	}

	// Start the first attempt
	m.CaptureScreenshotWithCompletion(retryHandler)
}

// TryEnableScreenCapturePermission attempts to prompt the user for screen capture permission.
// Returns true if permission was granted or already exists, false otherwise.
func TryEnableScreenCapturePermission() bool {
	// Note: In a real implementation, this would use the appropriate macOS APIs
	// to request screen capture permission. For now, this is a placeholder.
	fmt.Println("Requesting screen capture permission...")
	fmt.Println("Please grant permission in the system dialog that appears.")
	fmt.Println("If no dialog appears, please enable it manually in System Preferences > Security & Privacy > Privacy > Screen Recording")
	
	// Would return the result of permission request
	return true
}

// OpenSystemPreferencesTCC opens the System Preferences app to the Screen Recording section
func OpenSystemPreferencesTCC() {
	// This requires AppKit and a URL to open
	url := objc.Call[objc.Object](
		objc.GetClass("NSURL"), 
		objc.Sel("URLWithString:"), 
		"x-apple.systempreferences:com.apple.preference.security?Privacy_ScreenCapture",
	)
	workspace := objc.Call[objc.Object](
		objc.GetClass("NSWorkspace"), 
		objc.Sel("sharedWorkspace"),
	)
	objc.Call[bool](
		workspace,
		objc.Sel("openURL:"),
		url,
	)
}
