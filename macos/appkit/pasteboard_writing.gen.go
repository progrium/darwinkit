// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that defines the interface for retrieving a representation of an object that can be written to a pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardwriting?language=objc
type PPasteboardWriting interface {
	// optional
	WritableTypesForPasteboard(pasteboard Pasteboard) []PasteboardType
	HasWritableTypesForPasteboard() bool

	// optional
	PasteboardPropertyListForType(type_ PasteboardType) objc.IObject
	HasPasteboardPropertyListForType() bool

	// optional
	WritingOptionsForTypePasteboard(type_ PasteboardType, pasteboard Pasteboard) PasteboardWritingOptions
	HasWritingOptionsForTypePasteboard() bool
}

// A concrete type wrapper for the [PPasteboardWriting] protocol.
type PasteboardWritingWrapper struct {
	objc.Object
}

func (p_ PasteboardWritingWrapper) HasWritableTypesForPasteboard() bool {
	return p_.RespondsToSelector(objc.Sel("writableTypesForPasteboard:"))
}

// Returns an array of UTI strings of data types the receiver can write to a given pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardwriting/1533578-writabletypesforpasteboard?language=objc
func (p_ PasteboardWritingWrapper) WritableTypesForPasteboard(pasteboard IPasteboard) []PasteboardType {
	rv := objc.Call[[]PasteboardType](p_, objc.Sel("writableTypesForPasteboard:"), objc.Ptr(pasteboard))
	return rv
}

func (p_ PasteboardWritingWrapper) HasPasteboardPropertyListForType() bool {
	return p_.RespondsToSelector(objc.Sel("pasteboardPropertyListForType:"))
}

// Returns a property list object to represent the receiver on a pasteboard as an object of a specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardwriting/1529482-pasteboardpropertylistfortype?language=objc
func (p_ PasteboardWritingWrapper) PasteboardPropertyListForType(type_ PasteboardType) objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("pasteboardPropertyListForType:"), type_)
	return rv
}

func (p_ PasteboardWritingWrapper) HasWritingOptionsForTypePasteboard() bool {
	return p_.RespondsToSelector(objc.Sel("writingOptionsForType:pasteboard:"))
}

// Returns options for writing data of a specified type to a given pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardwriting/1525372-writingoptionsfortype?language=objc
func (p_ PasteboardWritingWrapper) WritingOptionsForTypePasteboard(type_ PasteboardType, pasteboard IPasteboard) PasteboardWritingOptions {
	rv := objc.Call[PasteboardWritingOptions](p_, objc.Sel("writingOptionsForType:pasteboard:"), type_, objc.Ptr(pasteboard))
	return rv
}
