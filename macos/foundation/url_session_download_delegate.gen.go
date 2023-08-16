// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines methods that URL session instances call on their delegates to handle task-level events specific to download tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondownloaddelegate?language=objc
type PURLSessionDownloadDelegate interface {
	// optional
	URLSessionDownloadTaskDidFinishDownloadingToURL(session URLSession, downloadTask URLSessionDownloadTask, location URL)
	HasURLSessionDownloadTaskDidFinishDownloadingToURL() bool
}

// A delegate implementation builder for the [PURLSessionDownloadDelegate] protocol.
type URLSessionDownloadDelegate struct {
	_URLSessionDownloadTaskDidFinishDownloadingToURL func(session URLSession, downloadTask URLSessionDownloadTask, location URL)
}

func (di *URLSessionDownloadDelegate) HasURLSessionDownloadTaskDidFinishDownloadingToURL() bool {
	return di._URLSessionDownloadTaskDidFinishDownloadingToURL != nil
}

// Tells the delegate that a download task has finished downloading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondownloaddelegate/1411575-urlsession?language=objc
func (di *URLSessionDownloadDelegate) SetURLSessionDownloadTaskDidFinishDownloadingToURL(f func(session URLSession, downloadTask URLSessionDownloadTask, location URL)) {
	di._URLSessionDownloadTaskDidFinishDownloadingToURL = f
}

// Tells the delegate that a download task has finished downloading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondownloaddelegate/1411575-urlsession?language=objc
func (di *URLSessionDownloadDelegate) URLSessionDownloadTaskDidFinishDownloadingToURL(session URLSession, downloadTask URLSessionDownloadTask, location URL) {
	di._URLSessionDownloadTaskDidFinishDownloadingToURL(session, downloadTask, location)
}

// A concrete type wrapper for the [PURLSessionDownloadDelegate] protocol.
type URLSessionDownloadDelegateWrapper struct {
	objc.Object
}

func (u_ URLSessionDownloadDelegateWrapper) HasURLSessionDownloadTaskDidFinishDownloadingToURL() bool {
	return u_.RespondsToSelector(objc.Sel("URLSession:downloadTask:didFinishDownloadingToURL:"))
}

// Tells the delegate that a download task has finished downloading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondownloaddelegate/1411575-urlsession?language=objc
func (u_ URLSessionDownloadDelegateWrapper) URLSessionDownloadTaskDidFinishDownloadingToURL(session IURLSession, downloadTask IURLSessionDownloadTask, location IURL) {
	objc.Call[objc.Void](u_, objc.Sel("URLSession:downloadTask:didFinishDownloadingToURL:"), objc.Ptr(session), objc.Ptr(downloadTask), objc.Ptr(location))
}
