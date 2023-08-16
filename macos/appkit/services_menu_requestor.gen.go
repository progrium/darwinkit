// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that support interaction with items users can share through a sharing service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsservicesmenurequestor?language=objc
type PServicesMenuRequestor interface {
	// optional
	ReadSelectionFromPasteboard(pboard Pasteboard) bool
	HasReadSelectionFromPasteboard() bool

	// optional
	WriteSelectionToPasteboardTypes(pboard Pasteboard, types []PasteboardType) bool
	HasWriteSelectionToPasteboardTypes() bool
}

// A concrete type wrapper for the [PServicesMenuRequestor] protocol.
type ServicesMenuRequestorWrapper struct {
	objc.Object
}

func (s_ ServicesMenuRequestorWrapper) HasReadSelectionFromPasteboard() bool {
	return s_.RespondsToSelector(objc.Sel("readSelectionFromPasteboard:"))
}

// Reads data from the pasteboard and uses it to replace the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsservicesmenurequestor/1428481-readselectionfrompasteboard?language=objc
func (s_ ServicesMenuRequestorWrapper) ReadSelectionFromPasteboard(pboard IPasteboard) bool {
	rv := objc.Call[bool](s_, objc.Sel("readSelectionFromPasteboard:"), objc.Ptr(pboard))
	return rv
}

func (s_ ServicesMenuRequestorWrapper) HasWriteSelectionToPasteboardTypes() bool {
	return s_.RespondsToSelector(objc.Sel("writeSelectionToPasteboard:types:"))
}

// Writes the current selection to the pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsservicesmenurequestor/1428477-writeselectiontopasteboard?language=objc
func (s_ ServicesMenuRequestorWrapper) WriteSelectionToPasteboardTypes(pboard IPasteboard, types []PasteboardType) bool {
	rv := objc.Call[bool](s_, objc.Sel("writeSelectionToPasteboard:types:"), objc.Ptr(pboard), types)
	return rv
}
