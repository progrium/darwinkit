package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/pdfkit"
	"github.com/progrium/darwinkit/objc"
)

func main() {
	// Initialize application
	objc.WithAutoreleasePool(func() {
		app := appkit.SharedApplication()
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)

		// Create window
		rect := foundation.MakeRect(0, 0, 800, 600)
		window := appkit.NewWindow(rect, appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable|appkit.WindowStyleMaskResizable, appkit.BackingStoreBuffered, false)
		window.SetTitle("PDFKit Example")
		window.SetReleasedWhenClosed(false)
		window.Center()

		// Create PDF view
		pdfView := pdfkit.NewPDFView()
		
		// Get sample PDF path (replace with your own PDF path)
		homedir, _ := os.UserHomeDir()
		pdfPath := filepath.Join(homedir, "Documents", "sample.pdf")
		
		// Check if the file exists, show warning if not
		if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
			fmt.Printf("Warning: PDF file not found at %s\nApplication will open with an empty PDF view.\n", pdfPath)
			fmt.Println("Please place a sample.pdf file in your Documents folder to view content.")
		} else {
			// Load PDF document
			url := foundation.URLWithFilePath(pdfPath)
			doc := pdfkit.DocumentWithURL(url)
			pdfView.SetDocument(doc)
			
			// Print document information
			fmt.Printf("Loaded PDF with %d pages\n", doc.PageCount())
			
			// Display outline if available
			outlines := doc.Outlines()
			if len(outlines) > 0 {
				fmt.Println("Document outline:")
				for i, outline := range outlines {
					fmt.Printf("  %d. %s\n", i+1, outline.Label())
				}
			}
		}
		
		// Set up auto-sizing for the PDF view
		pdfView.SetFrame(window.ContentView().Frame())
		pdfView.SetAutoresizingMask(appkit.ViewWidthSizable | appkit.ViewHeightSizable)
		
		// Add PDF view to window
		window.ContentView().AddSubview(pdfView)
		
		// Show window and run application
		window.MakeKeyAndOrderFront(nil)
		app.Run()
	})
}