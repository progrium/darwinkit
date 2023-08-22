// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines the methods to implement to respond to asset-download events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloaddelegate?language=objc
type PAssetDownloadDelegate interface {
	// optional
	URLSessionAggregateAssetDownloadTaskWillDownloadToURL(session foundation.URLSession, aggregateAssetDownloadTask AggregateAssetDownloadTask, location foundation.URL)
	HasURLSessionAggregateAssetDownloadTaskWillDownloadToURL() bool
}

// A delegate implementation builder for the [PAssetDownloadDelegate] protocol.
type AssetDownloadDelegate struct {
	_URLSessionAggregateAssetDownloadTaskWillDownloadToURL func(session foundation.URLSession, aggregateAssetDownloadTask AggregateAssetDownloadTask, location foundation.URL)
}

func (di *AssetDownloadDelegate) HasURLSessionAggregateAssetDownloadTaskWillDownloadToURL() bool {
	return di._URLSessionAggregateAssetDownloadTaskWillDownloadToURL != nil
}

// Tells the delegate the final location of the asset when the download completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloaddelegate/2897241-urlsession?language=objc
func (di *AssetDownloadDelegate) SetURLSessionAggregateAssetDownloadTaskWillDownloadToURL(f func(session foundation.URLSession, aggregateAssetDownloadTask AggregateAssetDownloadTask, location foundation.URL)) {
	di._URLSessionAggregateAssetDownloadTaskWillDownloadToURL = f
}

// Tells the delegate the final location of the asset when the download completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloaddelegate/2897241-urlsession?language=objc
func (di *AssetDownloadDelegate) URLSessionAggregateAssetDownloadTaskWillDownloadToURL(session foundation.URLSession, aggregateAssetDownloadTask AggregateAssetDownloadTask, location foundation.URL) {
	di._URLSessionAggregateAssetDownloadTaskWillDownloadToURL(session, aggregateAssetDownloadTask, location)
}

// A concrete type wrapper for the [PAssetDownloadDelegate] protocol.
type AssetDownloadDelegateWrapper struct {
	objc.Object
}

func (a_ AssetDownloadDelegateWrapper) HasURLSessionAggregateAssetDownloadTaskWillDownloadToURL() bool {
	return a_.RespondsToSelector(objc.Sel("URLSession:aggregateAssetDownloadTask:willDownloadToURL:"))
}

// Tells the delegate the final location of the asset when the download completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloaddelegate/2897241-urlsession?language=objc
func (a_ AssetDownloadDelegateWrapper) URLSessionAggregateAssetDownloadTaskWillDownloadToURL(session foundation.IURLSession, aggregateAssetDownloadTask IAggregateAssetDownloadTask, location foundation.IURL) {
	objc.Call[objc.Void](a_, objc.Sel("URLSession:aggregateAssetDownloadTask:willDownloadToURL:"), objc.Ptr(session), objc.Ptr(aggregateAssetDownloadTask), objc.Ptr(location))
}
