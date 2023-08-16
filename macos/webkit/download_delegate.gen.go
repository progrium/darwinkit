// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A protocol you implement to track download progress and handle redirects, authentication challenges, and failures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate?language=objc
type PDownloadDelegate interface {
	// optional
	DownloadDidFailWithErrorResumeData(download Download, error foundation.Error, resumeData []byte)
	HasDownloadDidFailWithErrorResumeData() bool

	// optional
	DownloadDidFinish(download Download)
	HasDownloadDidFinish() bool
}

// A delegate implementation builder for the [PDownloadDelegate] protocol.
type DownloadDelegate struct {
	_DownloadDidFailWithErrorResumeData func(download Download, error foundation.Error, resumeData []byte)
	_DownloadDidFinish                  func(download Download)
}

func (di *DownloadDelegate) HasDownloadDidFailWithErrorResumeData() bool {
	return di._DownloadDidFailWithErrorResumeData != nil
}

// Tells the delegate that the download failed, with error information and data you can use to restart the download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727345-download?language=objc
func (di *DownloadDelegate) SetDownloadDidFailWithErrorResumeData(f func(download Download, error foundation.Error, resumeData []byte)) {
	di._DownloadDidFailWithErrorResumeData = f
}

// Tells the delegate that the download failed, with error information and data you can use to restart the download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727345-download?language=objc
func (di *DownloadDelegate) DownloadDidFailWithErrorResumeData(download Download, error foundation.Error, resumeData []byte) {
	di._DownloadDidFailWithErrorResumeData(download, error, resumeData)
}
func (di *DownloadDelegate) HasDownloadDidFinish() bool {
	return di._DownloadDidFinish != nil
}

// Tells the delegate that the download finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727348-downloaddidfinish?language=objc
func (di *DownloadDelegate) SetDownloadDidFinish(f func(download Download)) {
	di._DownloadDidFinish = f
}

// Tells the delegate that the download finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727348-downloaddidfinish?language=objc
func (di *DownloadDelegate) DownloadDidFinish(download Download) {
	di._DownloadDidFinish(download)
}

// A concrete type wrapper for the [PDownloadDelegate] protocol.
type DownloadDelegateWrapper struct {
	objc.Object
}

func (d_ DownloadDelegateWrapper) HasDownloadDidFailWithErrorResumeData() bool {
	return d_.RespondsToSelector(objc.Sel("download:didFailWithError:resumeData:"))
}

// Tells the delegate that the download failed, with error information and data you can use to restart the download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727345-download?language=objc
func (d_ DownloadDelegateWrapper) DownloadDidFailWithErrorResumeData(download IDownload, error foundation.IError, resumeData []byte) {
	objc.Call[objc.Void](d_, objc.Sel("download:didFailWithError:resumeData:"), objc.Ptr(download), objc.Ptr(error), resumeData)
}

func (d_ DownloadDelegateWrapper) HasDownloadDidFinish() bool {
	return d_.RespondsToSelector(objc.Sel("downloadDidFinish:"))
}

// Tells the delegate that the download finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727348-downloaddidfinish?language=objc
func (d_ DownloadDelegateWrapper) DownloadDidFinish(download IDownload) {
	objc.Call[objc.Void](d_, objc.Sel("downloadDidFinish:"), objc.Ptr(download))
}
