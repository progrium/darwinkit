// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods implemented by the data provider of a pasteboard item to provide the data for a particular UTI type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboarditemdataprovider?language=objc
type PPasteboardItemDataProvider interface {
	// optional
	PasteboardFinishedWithDataProvider(pasteboard Pasteboard)
	HasPasteboardFinishedWithDataProvider() bool

	// optional
	PasteboardItemProvideDataForType(pasteboard Pasteboard, item PasteboardItem, type_ PasteboardType)
	HasPasteboardItemProvideDataForType() bool
}

// A concrete type wrapper for the [PPasteboardItemDataProvider] protocol.
type PasteboardItemDataProviderWrapper struct {
	objc.Object
}

func (p_ PasteboardItemDataProviderWrapper) HasPasteboardFinishedWithDataProvider() bool {
	return p_.RespondsToSelector(objc.Sel("pasteboardFinishedWithDataProvider:"))
}

// Informs the receiver that the pasteboard no longer needs the data provider for any of its pasteboard items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboarditemdataprovider/1508506-pasteboardfinishedwithdataprovid?language=objc
func (p_ PasteboardItemDataProviderWrapper) PasteboardFinishedWithDataProvider(pasteboard IPasteboard) {
	objc.Call[objc.Void](p_, objc.Sel("pasteboardFinishedWithDataProvider:"), objc.Ptr(pasteboard))
}

func (p_ PasteboardItemDataProviderWrapper) HasPasteboardItemProvideDataForType() bool {
	return p_.RespondsToSelector(objc.Sel("pasteboard:item:provideDataForType:"))
}

// Asks the receiver to provide data for a specified type to a given pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboarditemdataprovider/1508503-pasteboard?language=objc
func (p_ PasteboardItemDataProviderWrapper) PasteboardItemProvideDataForType(pasteboard IPasteboard, item IPasteboardItem, type_ PasteboardType) {
	objc.Call[objc.Void](p_, objc.Sel("pasteboard:item:provideDataForType:"), objc.Ptr(pasteboard), objc.Ptr(item), type_)
}
