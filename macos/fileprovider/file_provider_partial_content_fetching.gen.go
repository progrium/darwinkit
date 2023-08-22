// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// Support for fetching part of a file’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderpartialcontentfetching?language=objc
type PFileProviderPartialContentFetching interface {
	// optional
	FetchPartialContentsForItemWithIdentifierVersionRequestMinimalRangeAligningToOptionsCompletionHandler(itemIdentifier FileProviderItemIdentifier, requestedVersion FileProviderItemVersion, request FileProviderRequest, requestedRange foundation.Range, alignment uint, options FileProviderFetchContentsOptions, completionHandler func(fileContents foundation.URL, item objc.Object, retrievedRange foundation.Range, flags FileProviderMaterializationFlags, error foundation.Error)) foundation.IProgress
	HasFetchPartialContentsForItemWithIdentifierVersionRequestMinimalRangeAligningToOptionsCompletionHandler() bool
}

// A concrete type wrapper for the [PFileProviderPartialContentFetching] protocol.
type FileProviderPartialContentFetchingWrapper struct {
	objc.Object
}

func (f_ FileProviderPartialContentFetchingWrapper) HasFetchPartialContentsForItemWithIdentifierVersionRequestMinimalRangeAligningToOptionsCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("fetchPartialContentsForItemWithIdentifier:version:request:minimalRange:aligningTo:options:completionHandler:"))
}

// Tells the file provider to download the requested item from remote storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderpartialcontentfetching/3923718-fetchpartialcontentsforitemwithi?language=objc
func (f_ FileProviderPartialContentFetchingWrapper) FetchPartialContentsForItemWithIdentifierVersionRequestMinimalRangeAligningToOptionsCompletionHandler(itemIdentifier FileProviderItemIdentifier, requestedVersion IFileProviderItemVersion, request IFileProviderRequest, requestedRange foundation.Range, alignment uint, options FileProviderFetchContentsOptions, completionHandler func(fileContents foundation.URL, item objc.Object, retrievedRange foundation.Range, flags FileProviderMaterializationFlags, error foundation.Error)) foundation.Progress {
	rv := objc.Call[foundation.Progress](f_, objc.Sel("fetchPartialContentsForItemWithIdentifier:version:request:minimalRange:aligningTo:options:completionHandler:"), itemIdentifier, objc.Ptr(requestedVersion), objc.Ptr(request), requestedRange, alignment, options, completionHandler)
	return rv
}
