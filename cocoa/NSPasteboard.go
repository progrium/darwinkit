package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

// The supported pasteboard types.
// https://developer.apple.com/documentation/appkit/nspasteboardtype?language=objc
type NSPasteboardType string

// check https://github.com/progrium/macdriver/wiki/How-to-pull-enum-constants for example on how to get these values from objc
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

// Wrapper for NSPasteboard
// https://developer.apple.com/documentation/appkit/nspasteboard?language=objc
type NSPasteboard struct {
	gen_NSPasteboard
}

// SetStringForType sets the given string as the representation for the specified type for the first item on the receiver.
// https://developer.apple.com/documentation/appkit/nspasteboard/1528225-setstring?language=objc
func (pb NSPasteboard) SetStringForType(s string, t NSPasteboardType) {
	pb.Send("setString:forType:", core.String(s), core.String(string(t)))
}

// StringForType returns a concatenation of the strings for the specified type from all the items in the receiver that contain the type.
// https://developer.apple.com/documentation/appkit/nspasteboard/1533566-stringfortype?language=objc
func (pb NSPasteboard) StringForType(t NSPasteboardType) string {
	return pb.Send("stringForType:", core.String(string(t))).String()
}

// DataForType returns the data for the specified type from the first item in the receiver that contains the type.
// https://developer.apple.com/documentation/appkit/nspasteboard/1531810-datafortype?language=objc
func (pb NSPasteboard) DataForType(t NSPasteboardType) core.NSData {
	return core.NSData_FromRef(pb.Send("dataForType:", core.String(string(t))))
}

// Types is an array of the receiver’s supported data types.
// https://developer.apple.com/documentation/appkit/nspasteboard/1529599-types?language=objc
func (pb NSPasteboard) Types() []NSPasteboardType {
	ss := pb.gen_NSPasteboard.Types().Strings()
	types := make([]NSPasteboardType, len(ss))
	for i, s := range ss {
		types[i] = NSPasteboardType(s)
	}
	return types
}

// AvailableTypeFromArray scans the specified types for a type that the receiver supports.
// https://developer.apple.com/documentation/appkit/nspasteboard/1526078-availabletypefromarray?language=objc
func (pb NSPasteboard) AvailableTypeFromArray(types []NSPasteboardType) NSPasteboardType {
	strs := make([]objc.Object, len(types))
	for i, t := range types {
		strs[i] = core.String(string(t))
	}
	arr := core.NSArray_WithObjects(strs...)
	o := pb.Send("availableTypeFromArray:", arr)
	pbType := NSPasteboardType(o.String())
	if pbType == "(nil)" {
		return ""
	}
	return pbType
}
