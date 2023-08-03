// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type IPasteboardItemDataProvider interface {
	// required
	PasteboardItemProvideDataForType(pasteboard Pasteboard, item PasteboardItem, type_ PasteboardType)
	ImplementsPasteboardFinishedWithDataProvider() bool
	// optional
	PasteboardFinishedWithDataProvider(pasteboard Pasteboard)
}

type PasteboardItemDataProviderWrapper struct {
	objc.Object
}

func (p_ PasteboardItemDataProviderWrapper) PasteboardItemProvideDataForType(pasteboard IPasteboard, item IPasteboardItem, type_ PasteboardType) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("pasteboard:item:provideDataForType:"), objc.ExtractPtr(pasteboard), objc.ExtractPtr(item), type_)
}

func (p_ PasteboardItemDataProviderWrapper) ImplementsPasteboardFinishedWithDataProvider() bool {
	return p_.RespondsToSelector(objc.GetSelector("pasteboardFinishedWithDataProvider:"))
}

func (p_ PasteboardItemDataProviderWrapper) PasteboardFinishedWithDataProvider(pasteboard IPasteboard) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("pasteboardFinishedWithDataProvider:"), objc.ExtractPtr(pasteboard))
}
