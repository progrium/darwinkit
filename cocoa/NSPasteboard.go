package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

// The supported pasteboard types.
// https://developer.apple.com/documentation/appkit/nspasteboardtype?language=objc
type NSPasteboardType = string

/*
Code to get the values for those constants:
```objc
//
//  main.m
//  test-objc
//
//  Created by Jonas Faber on 16.02.21.
//

#import <Foundation/Foundation.h>
#import <AppKit/NSPasteboard.h>
#import <AppKit/NSTextView.h>

int main(int argc, const char * argv[]) {
    @autoreleasepool {

        NSLog(@"NSPasteboardTypeURL:%@", NSPasteboardTypeURL);
        NSLog(@"NSPasteboardTypeColor:%@", NSPasteboardTypeColor);
        NSLog(@"NSPasteboardTypeFileURL:%@", NSPasteboardTypeFileURL);
        NSLog(@"NSPasteboardTypeFont:%@", NSPasteboardTypeFont);
        NSLog(@"NSPasteboardTypeHTML:%@", NSPasteboardTypeHTML);
        NSLog(@"NSPasteboardTypeMultipleTextSelection:%@", NSPasteboardTypeMultipleTextSelection);
        NSLog(@"NSPasteboardTypePDF:%@", NSPasteboardTypePDF);
        NSLog(@"NSPasteboardTypePNG:%@", NSPasteboardTypePNG);
        NSLog(@"NSPasteboardTypeRTF:%@", NSPasteboardTypeRTF);
        NSLog(@"NSPasteboardTypeRTFD:%@", NSPasteboardTypeRTFD);
        NSLog(@"NSPasteboardTypeRuler:%@", NSPasteboardTypeRuler);
        NSLog(@"NSPasteboardTypeSound:%@", NSPasteboardTypeSound);
        NSLog(@"NSPasteboardTypeString:%@", NSPasteboardTypeString);
        NSLog(@"NSPasteboardTypeTabularText:%@", NSPasteboardTypeTabularText);
        NSLog(@"NSPasteboardTypeTextFinderOptions:%@", NSPasteboardTypeTextFinderOptions);
        NSLog(@"NSPasteboardTypeTIFF:%@", NSPasteboardTypeTIFF);
    }
    return 0;
}
```
*/

const (
	// NSPasteboardTypeURL holds URL data for one file or resource.
	NSPasteboardTypeURL = NSPasteboardType("public.url")
	// NSPasteboardTypeColor holds color data.
	NSPasteboardTypeColor = NSPasteboardType("com.apple.cocoa.pasteboard.color")
	// NSPasteboardTypeFileURL holds a file URL.
	NSPasteboardTypeFileURL = NSPasteboardType("public.file-url")
	// NSPasteboardTypeFont holds font and character information.
	NSPasteboardTypeFont = NSPasteboardType("com.apple.cocoa.pasteboard.character-formatting")
	// NSPasteboardTypeHTML holds type for HTML content.
	NSPasteboardTypeHTML = NSPasteboardType("public.html")
	// NSPasteboardTypeMultipleTextSelection holds multiple text selection.
	NSPasteboardTypeMultipleTextSelection = NSPasteboardType("com.apple.cocoa.pasteboard.multiple-text-selection")
	// NSPasteboardTypePDF holds PDF data.
	NSPasteboardTypePDF = NSPasteboardType("com.adobe.pdf")
	// NSPasteboardTypePNG holds PNG image data.
	NSPasteboardTypePNG = NSPasteboardType("public.png")
	// NSPasteboardTypeRTF holds rich Text Format (RTF) data.
	NSPasteboardTypeRTF = NSPasteboardType("public.rtf")
	// NSPasteboardTypeRTFD holds RTFD formatted file contents.
	NSPasteboardTypeRTFD = NSPasteboardType("com.apple.flat-rtfd")
	// NSPasteboardTypeRuler holds paragraph formatting information.
	NSPasteboardTypeRuler = NSPasteboardType("com.apple.cocoa.pasteboard.paragraph-formatting")
	// NSPasteboardTypeSound holds sound data.
	NSPasteboardTypeSound = NSPasteboardType("com.apple.cocoa.pasteboard.sound")
	// NSPasteboardTypeString holds string data.
	NSPasteboardTypeString = NSPasteboardType("public.utf8-plain-text")
	// NSPasteboardTypeTabularText holds tab-separated fields of text.
	NSPasteboardTypeTabularText = NSPasteboardType("public.utf8-tab-separated-values-text")
	// NSPasteboardTypeTextFinderOptions holds type for the Find panel metadata property list.
	NSPasteboardTypeTextFinderOptions = NSPasteboardType("com.apple.cocoa.pasteboard.find-panel-search-options")
	// NSPasteboardTypeTIFF holds tag Image File Format (TIFF) data.
	NSPasteboardTypeTIFF = NSPasteboardType("public.tiff")
)

type NSPasteboard struct {
	objc.Object
}

// Wrapper for NSPasteboard
// https://developer.apple.com/documentation/appkit/nspasteboard?language=objc
var NSPasteboard_ = NSPasteboard{objc.Get("NSPasteboard")}

// NSPasteboard_GeneralPasteboard is the shared pasteboard object to use for general content.
// https://developer.apple.com/documentation/appkit/nspasteboard/1530091-generalpasteboard?language=objc
func NSPasteboard_GeneralPasteboard() NSPasteboard {
	return NSPasteboard{NSPasteboard_.Get("generalPasteboard")}
}

// ClearContents clears the existing contents of the pasteboard.
// https://developer.apple.com/documentation/appkit/nspasteboard/1533599-clearcontents?language=objc
func (pb NSPasteboard) ClearContents() {
	pb.Send("clearContents")
}

// SetStringForType sets the given string as the representation for the specified type for the first item on the receiver.
// https://developer.apple.com/documentation/appkit/nspasteboard/1528225-setstring?language=objc
func (pb NSPasteboard) SetStringForType(s string, t NSPasteboardType) {
	pb.Send("setString:forType:", core.NSString_FromString(s),  core.NSString_FromString(t))
}
// StringForType returns a concatenation of the strings for the specified type from all the items in the receiver that contain the type.
// https://developer.apple.com/documentation/appkit/nspasteboard/1533566-stringfortype?language=objc
func (pb NSPasteboard) StringForType(t NSPasteboardType) string {
	return pb.Send("stringForType:",  core.NSString_FromString(t)).String()
}