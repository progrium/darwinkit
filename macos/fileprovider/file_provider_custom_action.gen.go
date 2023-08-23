// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// Support for custom actions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidercustomaction?language=objc
type PFileProviderCustomAction interface {
	// optional
	PerformActionWithIdentifierOnItemsWithIdentifiersCompletionHandler(actionIdentifier FileProviderExtensionActionIdentifier, itemIdentifiers []FileProviderItemIdentifier, completionHandler func(error foundation.Error)) foundation.IProgress
	HasPerformActionWithIdentifierOnItemsWithIdentifiersCompletionHandler() bool
}

// A concrete type wrapper for the [PFileProviderCustomAction] protocol.
type FileProviderCustomActionWrapper struct {
	objc.Object
}

func (f_ FileProviderCustomActionWrapper) HasPerformActionWithIdentifierOnItemsWithIdentifiersCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("performActionWithIdentifier:onItemsWithIdentifiers:completionHandler:"))
}

// Tells the File Provider extension to perform a custom action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidercustomaction/3553293-performactionwithidentifier?language=objc
func (f_ FileProviderCustomActionWrapper) PerformActionWithIdentifierOnItemsWithIdentifiersCompletionHandler(actionIdentifier FileProviderExtensionActionIdentifier, itemIdentifiers []FileProviderItemIdentifier, completionHandler func(error foundation.Error)) foundation.Progress {
	rv := objc.Call[foundation.Progress](f_, objc.Sel("performActionWithIdentifier:onItemsWithIdentifiers:completionHandler:"), actionIdentifier, itemIdentifiers, completionHandler)
	return rv
}
