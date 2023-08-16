// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that URL download delegates implement to interact with a URL download request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownloaddelegate?language=objc
type PURLDownloadDelegate interface {
	// optional
	DownloadCanAuthenticateAgainstProtectionSpace(connection URLDownload, protectionSpace URLProtectionSpace) bool
	HasDownloadCanAuthenticateAgainstProtectionSpace() bool

	// optional
	DownloadDidFinish(download URLDownload)
	HasDownloadDidFinish() bool

	// optional
	DownloadDidBegin(download URLDownload)
	HasDownloadDidBegin() bool

	// optional
	DownloadShouldUseCredentialStorage(download URLDownload) bool
	HasDownloadShouldUseCredentialStorage() bool
}

// A delegate implementation builder for the [PURLDownloadDelegate] protocol.
type URLDownloadDelegate struct {
	_DownloadCanAuthenticateAgainstProtectionSpace func(connection URLDownload, protectionSpace URLProtectionSpace) bool
	_DownloadDidFinish                             func(download URLDownload)
	_DownloadDidBegin                              func(download URLDownload)
	_DownloadShouldUseCredentialStorage            func(download URLDownload) bool
}

func (di *URLDownloadDelegate) HasDownloadCanAuthenticateAgainstProtectionSpace() bool {
	return di._DownloadCanAuthenticateAgainstProtectionSpace != nil
}

// Sent to determine whether the delegate is able to respond to a protection space’s form of authentication. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownloaddelegate/1417213-download?language=objc
func (di *URLDownloadDelegate) SetDownloadCanAuthenticateAgainstProtectionSpace(f func(connection URLDownload, protectionSpace URLProtectionSpace) bool) {
	di._DownloadCanAuthenticateAgainstProtectionSpace = f
}

// Sent to determine whether the delegate is able to respond to a protection space’s form of authentication. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownloaddelegate/1417213-download?language=objc
func (di *URLDownloadDelegate) DownloadCanAuthenticateAgainstProtectionSpace(connection URLDownload, protectionSpace URLProtectionSpace) bool {
	return di._DownloadCanAuthenticateAgainstProtectionSpace(connection, protectionSpace)
}
func (di *URLDownloadDelegate) HasDownloadDidFinish() bool {
	return di._DownloadDidFinish != nil
}

// Sent when a download object has completed downloading successfully and has written its results to disk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownloaddelegate/1408884-downloaddidfinish?language=objc
func (di *URLDownloadDelegate) SetDownloadDidFinish(f func(download URLDownload)) {
	di._DownloadDidFinish = f
}

// Sent when a download object has completed downloading successfully and has written its results to disk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownloaddelegate/1408884-downloaddidfinish?language=objc
func (di *URLDownloadDelegate) DownloadDidFinish(download URLDownload) {
	di._DownloadDidFinish(download)
}
func (di *URLDownloadDelegate) HasDownloadDidBegin() bool {
	return di._DownloadDidBegin != nil
}

// Sent immediately after a download object begins a download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownloaddelegate/1409618-downloaddidbegin?language=objc
func (di *URLDownloadDelegate) SetDownloadDidBegin(f func(download URLDownload)) {
	di._DownloadDidBegin = f
}

// Sent immediately after a download object begins a download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownloaddelegate/1409618-downloaddidbegin?language=objc
func (di *URLDownloadDelegate) DownloadDidBegin(download URLDownload) {
	di._DownloadDidBegin(download)
}
func (di *URLDownloadDelegate) HasDownloadShouldUseCredentialStorage() bool {
	return di._DownloadShouldUseCredentialStorage != nil
}

// Sent to determine whether the URL loader should consult the credential storage to authenticate the download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownloaddelegate/1416506-downloadshouldusecredentialstora?language=objc
func (di *URLDownloadDelegate) SetDownloadShouldUseCredentialStorage(f func(download URLDownload) bool) {
	di._DownloadShouldUseCredentialStorage = f
}

// Sent to determine whether the URL loader should consult the credential storage to authenticate the download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownloaddelegate/1416506-downloadshouldusecredentialstora?language=objc
func (di *URLDownloadDelegate) DownloadShouldUseCredentialStorage(download URLDownload) bool {
	return di._DownloadShouldUseCredentialStorage(download)
}

// A concrete type wrapper for the [PURLDownloadDelegate] protocol.
type URLDownloadDelegateWrapper struct {
	objc.Object
}

func (u_ URLDownloadDelegateWrapper) HasDownloadCanAuthenticateAgainstProtectionSpace() bool {
	return u_.RespondsToSelector(objc.Sel("download:canAuthenticateAgainstProtectionSpace:"))
}

// Sent to determine whether the delegate is able to respond to a protection space’s form of authentication. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownloaddelegate/1417213-download?language=objc
func (u_ URLDownloadDelegateWrapper) DownloadCanAuthenticateAgainstProtectionSpace(connection IURLDownload, protectionSpace IURLProtectionSpace) bool {
	rv := objc.Call[bool](u_, objc.Sel("download:canAuthenticateAgainstProtectionSpace:"), objc.Ptr(connection), objc.Ptr(protectionSpace))
	return rv
}

func (u_ URLDownloadDelegateWrapper) HasDownloadDidFinish() bool {
	return u_.RespondsToSelector(objc.Sel("downloadDidFinish:"))
}

// Sent when a download object has completed downloading successfully and has written its results to disk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownloaddelegate/1408884-downloaddidfinish?language=objc
func (u_ URLDownloadDelegateWrapper) DownloadDidFinish(download IURLDownload) {
	objc.Call[objc.Void](u_, objc.Sel("downloadDidFinish:"), objc.Ptr(download))
}

func (u_ URLDownloadDelegateWrapper) HasDownloadDidBegin() bool {
	return u_.RespondsToSelector(objc.Sel("downloadDidBegin:"))
}

// Sent immediately after a download object begins a download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownloaddelegate/1409618-downloaddidbegin?language=objc
func (u_ URLDownloadDelegateWrapper) DownloadDidBegin(download IURLDownload) {
	objc.Call[objc.Void](u_, objc.Sel("downloadDidBegin:"), objc.Ptr(download))
}

func (u_ URLDownloadDelegateWrapper) HasDownloadShouldUseCredentialStorage() bool {
	return u_.RespondsToSelector(objc.Sel("downloadShouldUseCredentialStorage:"))
}

// Sent to determine whether the URL loader should consult the credential storage to authenticate the download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownloaddelegate/1416506-downloadshouldusecredentialstora?language=objc
func (u_ URLDownloadDelegateWrapper) DownloadShouldUseCredentialStorage(download IURLDownload) bool {
	rv := objc.Call[bool](u_, objc.Sel("downloadShouldUseCredentialStorage:"), objc.Ptr(download))
	return rv
}
