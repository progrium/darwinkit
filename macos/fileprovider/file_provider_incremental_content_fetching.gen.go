// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// Support for fetching changes to the item’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderincrementalcontentfetching?language=objc
type PFileProviderIncrementalContentFetching interface {
	// optional
	FetchContentsForItemWithIdentifierVersionUsingExistingContentsAtURLExistingVersionRequestCompletionHandler(itemIdentifier FileProviderItemIdentifier, requestedVersion FileProviderItemVersion, existingContents foundation.URL, existingVersion FileProviderItemVersion, request FileProviderRequest, completionHandler func(fileContents foundation.URL, item objc.Object, error foundation.Error)) foundation.IProgress
	HasFetchContentsForItemWithIdentifierVersionUsingExistingContentsAtURLExistingVersionRequestCompletionHandler() bool
}

// A concrete type wrapper for the [PFileProviderIncrementalContentFetching] protocol.
type FileProviderIncrementalContentFetchingWrapper struct {
	objc.Object
}

func (f_ FileProviderIncrementalContentFetchingWrapper) HasFetchContentsForItemWithIdentifierVersionUsingExistingContentsAtURLExistingVersionRequestCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("fetchContentsForItemWithIdentifier:version:usingExistingContentsAtURL:existingVersion:request:completionHandler:"))
}

// Asks the file provider for an update of the specified item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderincrementalcontentfetching/3553297-fetchcontentsforitemwithidentifi?language=objc
func (f_ FileProviderIncrementalContentFetchingWrapper) FetchContentsForItemWithIdentifierVersionUsingExistingContentsAtURLExistingVersionRequestCompletionHandler(itemIdentifier FileProviderItemIdentifier, requestedVersion IFileProviderItemVersion, existingContents foundation.IURL, existingVersion IFileProviderItemVersion, request IFileProviderRequest, completionHandler func(fileContents foundation.URL, item objc.Object, error foundation.Error)) foundation.Progress {
	rv := objc.Call[foundation.Progress](f_, objc.Sel("fetchContentsForItemWithIdentifier:version:usingExistingContentsAtURL:existingVersion:request:completionHandler:"), itemIdentifier, objc.Ptr(requestedVersion), objc.Ptr(existingContents), objc.Ptr(existingVersion), objc.Ptr(request), completionHandler)
	return rv
}
