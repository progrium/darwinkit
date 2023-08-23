// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// Support for item thumbnails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderthumbnailing?language=objc
type PFileProviderThumbnailing interface {
	// optional
	FetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler(itemIdentifiers []FileProviderItemIdentifier, size coregraphics.Size, perThumbnailCompletionHandler func(identifier FileProviderItemIdentifier, imageData []byte, error foundation.Error), completionHandler func(error foundation.Error)) foundation.IProgress
	HasFetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler() bool
}

// A concrete type wrapper for the [PFileProviderThumbnailing] protocol.
type FileProviderThumbnailingWrapper struct {
	objc.Object
}

func (f_ FileProviderThumbnailingWrapper) HasFetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("fetchThumbnailsForItemIdentifiers:requestedSize:perThumbnailCompletionHandler:completionHandler:"))
}

// Asks the file provider for a thumbnail of the specified items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderthumbnailing/3553313-fetchthumbnailsforitemidentifier?language=objc
func (f_ FileProviderThumbnailingWrapper) FetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler(itemIdentifiers []FileProviderItemIdentifier, size coregraphics.Size, perThumbnailCompletionHandler func(identifier FileProviderItemIdentifier, imageData []byte, error foundation.Error), completionHandler func(error foundation.Error)) foundation.Progress {
	rv := objc.Call[foundation.Progress](f_, objc.Sel("fetchThumbnailsForItemIdentifiers:requestedSize:perThumbnailCompletionHandler:completionHandler:"), itemIdentifiers, size, perThumbnailCompletionHandler, completionHandler)
	return rv
}
