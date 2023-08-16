// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that delegates of a URL connection created with Newsstand Kit implement to receive data associated with a download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondownloaddelegate?language=objc
type PURLConnectionDownloadDelegate interface {
	// optional
	ConnectionDidFinishDownloadingDestinationURL(connection URLConnection, destinationURL URL)
	HasConnectionDidFinishDownloadingDestinationURL() bool

	// optional
	ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes(connection URLConnection, totalBytesWritten int64, expectedTotalBytes int64)
	HasConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes() bool

	// optional
	ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes(connection URLConnection, bytesWritten int64, totalBytesWritten int64, expectedTotalBytes int64)
	HasConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes() bool
}

// A delegate implementation builder for the [PURLConnectionDownloadDelegate] protocol.
type URLConnectionDownloadDelegate struct {
	_ConnectionDidFinishDownloadingDestinationURL                      func(connection URLConnection, destinationURL URL)
	_ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes func(connection URLConnection, totalBytesWritten int64, expectedTotalBytes int64)
	_ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes         func(connection URLConnection, bytesWritten int64, totalBytesWritten int64, expectedTotalBytes int64)
}

func (di *URLConnectionDownloadDelegate) HasConnectionDidFinishDownloadingDestinationURL() bool {
	return di._ConnectionDidFinishDownloadingDestinationURL != nil
}

// Sent to the delegate when the URL connection has successfully downloaded the URL asset to a destination file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondownloaddelegate/1412126-connectiondidfinishdownloading?language=objc
func (di *URLConnectionDownloadDelegate) SetConnectionDidFinishDownloadingDestinationURL(f func(connection URLConnection, destinationURL URL)) {
	di._ConnectionDidFinishDownloadingDestinationURL = f
}

// Sent to the delegate when the URL connection has successfully downloaded the URL asset to a destination file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondownloaddelegate/1412126-connectiondidfinishdownloading?language=objc
func (di *URLConnectionDownloadDelegate) ConnectionDidFinishDownloadingDestinationURL(connection URLConnection, destinationURL URL) {
	di._ConnectionDidFinishDownloadingDestinationURL(connection, destinationURL)
}
func (di *URLConnectionDownloadDelegate) HasConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes() bool {
	return di._ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes != nil
}

// Sent to the delegate when an URL connection resumes downloading a URL asset that was earlier suspended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondownloaddelegate/1418157-connectiondidresumedownloading?language=objc
func (di *URLConnectionDownloadDelegate) SetConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes(f func(connection URLConnection, totalBytesWritten int64, expectedTotalBytes int64)) {
	di._ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes = f
}

// Sent to the delegate when an URL connection resumes downloading a URL asset that was earlier suspended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondownloaddelegate/1418157-connectiondidresumedownloading?language=objc
func (di *URLConnectionDownloadDelegate) ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes(connection URLConnection, totalBytesWritten int64, expectedTotalBytes int64) {
	di._ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes(connection, totalBytesWritten, expectedTotalBytes)
}
func (di *URLConnectionDownloadDelegate) HasConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes() bool {
	return di._ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes != nil
}

// Sent to the delegate to deliver progress information for a download of a URL asset to a destination file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondownloaddelegate/1418304-connection?language=objc
func (di *URLConnectionDownloadDelegate) SetConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes(f func(connection URLConnection, bytesWritten int64, totalBytesWritten int64, expectedTotalBytes int64)) {
	di._ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes = f
}

// Sent to the delegate to deliver progress information for a download of a URL asset to a destination file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondownloaddelegate/1418304-connection?language=objc
func (di *URLConnectionDownloadDelegate) ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes(connection URLConnection, bytesWritten int64, totalBytesWritten int64, expectedTotalBytes int64) {
	di._ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes(connection, bytesWritten, totalBytesWritten, expectedTotalBytes)
}

// A concrete type wrapper for the [PURLConnectionDownloadDelegate] protocol.
type URLConnectionDownloadDelegateWrapper struct {
	objc.Object
}

func (u_ URLConnectionDownloadDelegateWrapper) HasConnectionDidFinishDownloadingDestinationURL() bool {
	return u_.RespondsToSelector(objc.Sel("connectionDidFinishDownloading:destinationURL:"))
}

// Sent to the delegate when the URL connection has successfully downloaded the URL asset to a destination file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondownloaddelegate/1412126-connectiondidfinishdownloading?language=objc
func (u_ URLConnectionDownloadDelegateWrapper) ConnectionDidFinishDownloadingDestinationURL(connection IURLConnection, destinationURL IURL) {
	objc.Call[objc.Void](u_, objc.Sel("connectionDidFinishDownloading:destinationURL:"), objc.Ptr(connection), objc.Ptr(destinationURL))
}

func (u_ URLConnectionDownloadDelegateWrapper) HasConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes() bool {
	return u_.RespondsToSelector(objc.Sel("connectionDidResumeDownloading:totalBytesWritten:expectedTotalBytes:"))
}

// Sent to the delegate when an URL connection resumes downloading a URL asset that was earlier suspended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondownloaddelegate/1418157-connectiondidresumedownloading?language=objc
func (u_ URLConnectionDownloadDelegateWrapper) ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes(connection IURLConnection, totalBytesWritten int64, expectedTotalBytes int64) {
	objc.Call[objc.Void](u_, objc.Sel("connectionDidResumeDownloading:totalBytesWritten:expectedTotalBytes:"), objc.Ptr(connection), totalBytesWritten, expectedTotalBytes)
}

func (u_ URLConnectionDownloadDelegateWrapper) HasConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes() bool {
	return u_.RespondsToSelector(objc.Sel("connection:didWriteData:totalBytesWritten:expectedTotalBytes:"))
}

// Sent to the delegate to deliver progress information for a download of a URL asset to a destination file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnectiondownloaddelegate/1418304-connection?language=objc
func (u_ URLConnectionDownloadDelegateWrapper) ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes(connection IURLConnection, bytesWritten int64, totalBytesWritten int64, expectedTotalBytes int64) {
	objc.Call[objc.Void](u_, objc.Sel("connection:didWriteData:totalBytesWritten:expectedTotalBytes:"), objc.Ptr(connection), bytesWritten, totalBytesWritten, expectedTotalBytes)
}
