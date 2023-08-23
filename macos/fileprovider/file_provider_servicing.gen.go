// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// Support for providing a custom service source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderservicing?language=objc
type PFileProviderServicing interface {
	// optional
	SupportedServiceSourcesForItemIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier, completionHandler func(arg0 []FileProviderServiceSourceWrapper, arg1 foundation.Error)) foundation.IProgress
	HasSupportedServiceSourcesForItemIdentifierCompletionHandler() bool
}

// A concrete type wrapper for the [PFileProviderServicing] protocol.
type FileProviderServicingWrapper struct {
	objc.Object
}

func (f_ FileProviderServicingWrapper) HasSupportedServiceSourcesForItemIdentifierCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("supportedServiceSourcesForItemIdentifier:completionHandler:"))
}

// Asks the File Provider extension for an array of custom communication channels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderservicing/3553311-supportedservicesourcesforitemid?language=objc
func (f_ FileProviderServicingWrapper) SupportedServiceSourcesForItemIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier, completionHandler func(arg0 []FileProviderServiceSourceWrapper, arg1 foundation.Error)) foundation.Progress {
	rv := objc.Call[foundation.Progress](f_, objc.Sel("supportedServiceSourcesForItemIdentifier:completionHandler:"), itemIdentifier, completionHandler)
	return rv
}
