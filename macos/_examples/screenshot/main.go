package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/screencapturekit"
)

func main() {
	// Ensure we run on the main thread
	runtime.LockOSThread()

	fmt.Println("ScreenCaptureKit Screenshot Example")
	fmt.Println("-----------------------------------")

	// Get the screenshot manager
	manager := screencapturekit.SCScreenshotManager_SharedManager()

	// Create a completion handler to process the screenshot
	screenshotComplete := make(chan bool, 1)

	// Define the retry handler
	retryHandler := func(image appkit.Image, err screencapturekit.SCError, success bool) {
		if \!success {
			fmt.Printf("Error taking screenshot after retries: %v\n", err)
			
			// Handle TCC permission issues
			if err.IsTCCError() {
				fmt.Println("\n*** Screen Recording Permission Required ***")
				fmt.Println(err.GetPermissionInstructions())
				
				// Open System Preferences
				fmt.Println("Opening System Preferences to help you grant permission...")
				screencapturekit.OpenSystemPreferencesTCC()
				
				fmt.Println("\nPlease restart this application after granting permission.")
			}
			
			screenshotComplete <- false
			return
		}

		fmt.Println("Screenshot captured successfully\!")

		// Get desktop path for saving
		homeDir, err := os.UserHomeDir()
		if err \!= nil {
			fmt.Printf("Error getting home directory: %v\n", err)
			screenshotComplete <- false
			return
		}

		// Create filename with timestamp
		desktopPath := filepath.Join(homeDir, "Desktop")
		filename := filepath.Join(desktopPath, fmt.Sprintf("darwinkit_screenshot_%d.png", time.Now().Unix()))

		fmt.Printf("Saving screenshot to: %s\n", filename)

		// Convert to PNG data
		tiffData := image.TIFFRepresentation()
		if tiffData.IsNil() {
			fmt.Println("Failed to convert image to TIFF")
			screenshotComplete <- false
			return
		}

		imageRep := appkit.NSBitmapImageRep_ImageRepWithData(tiffData)
		if imageRep.IsNil() {
			fmt.Println("Failed to create image representation")
			screenshotComplete <- false
			return
		}

		pngData := imageRep.RepresentationUsingTypeProperties(
			appkit.NSBitmapImageFileTypePNG,
			nil,
		)
		if pngData.IsNil() {
			fmt.Println("Failed to convert to PNG")
			screenshotComplete <- false
			return
		}

		// Create NSString for the path
		filePath := foundation.NSString_StringWithString(filename)

		// Write to file
		ok := pngData.WriteToFileAtomically(filePath, true)
		if \!ok {
			fmt.Println("Failed to write image to file")
			screenshotComplete <- false
			return
		}

		fmt.Printf("Screenshot saved to %s\n", filename)
		screenshotComplete <- true
	}

	// Take screenshot with retry for TCC issues
	fmt.Println("Taking screenshot (with automatic retry for permission issues)...")
	manager.CaptureScreenshotWithRetry(3, 1*time.Second, retryHandler)

	// Wait for completion
	result := <-screenshotComplete
	if result {
		fmt.Println("\nScreenshot process completed successfully\!")
	} else {
		fmt.Println("\nScreenshot process failed.")
		os.Exit(1)
	}
}
