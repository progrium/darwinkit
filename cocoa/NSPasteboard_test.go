package cocoa_test

import (
	"fmt"
	"github.com/progrium/macdriver/cocoa"
)

func ExampleNSPasteboard(){
	// get general pasteboard
	gp := cocoa.NSPasteboard_GeneralPasteboard()


	// clear pasteboard and write to pasteboard
	gp.ClearContents()
	gp.SetStringForType("test-content", cocoa.NSPasteboardTypeString)

	// read data from pasteboard
	fmt.Println(gp.StringForType(cocoa.NSPasteboardTypeString))
	// Output: "test-content"
}
