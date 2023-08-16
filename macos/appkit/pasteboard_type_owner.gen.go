// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// An object that serves as a data provider for data types that use lazy data fulfillment from a pasteboard request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardtypeowner?language=objc
type PPasteboardTypeOwner interface {
	// optional
	PasteboardChangedOwner(sender Pasteboard)
	HasPasteboardChangedOwner() bool

	// optional
	PasteboardProvideDataForType(sender Pasteboard, type_ PasteboardType)
	HasPasteboardProvideDataForType() bool
}

// A concrete type wrapper for the [PPasteboardTypeOwner] protocol.
type PasteboardTypeOwnerWrapper struct {
	objc.Object
}

func (p_ PasteboardTypeOwnerWrapper) HasPasteboardChangedOwner() bool {
	return p_.RespondsToSelector(objc.Sel("pasteboardChangedOwner:"))
}

// Notifies the object that the pasteboard’s owner changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardtypeowner/3005194-pasteboardchangedowner?language=objc
func (p_ PasteboardTypeOwnerWrapper) PasteboardChangedOwner(sender IPasteboard) {
	objc.Call[objc.Void](p_, objc.Sel("pasteboardChangedOwner:"), objc.Ptr(sender))
}

func (p_ PasteboardTypeOwnerWrapper) HasPasteboardProvideDataForType() bool {
	return p_.RespondsToSelector(objc.Sel("pasteboard:provideDataForType:"))
}

// Requests that the object provide data for the data type to the pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardtypeowner/3005193-pasteboard?language=objc
func (p_ PasteboardTypeOwnerWrapper) PasteboardProvideDataForType(sender IPasteboard, type_ PasteboardType) {
	objc.Call[objc.Void](p_, objc.Sel("pasteboard:provideDataForType:"), objc.Ptr(sender), type_)
}
