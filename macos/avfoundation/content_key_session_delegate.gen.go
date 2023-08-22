// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that handles content key requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate?language=objc
type PContentKeySessionDelegate interface {
	// optional
	ContentKeySessionDidGenerateExpiredSessionReport(session ContentKeySession)
	HasContentKeySessionDidGenerateExpiredSessionReport() bool

	// optional
	ContentKeySessionContentProtectionSessionIdentifierDidChange(session ContentKeySession)
	HasContentKeySessionContentProtectionSessionIdentifierDidChange() bool

	// optional
	ContentKeySessionDidProvideContentKeyRequest(session ContentKeySession, keyRequest ContentKeyRequest)
	HasContentKeySessionDidProvideContentKeyRequest() bool
}

// A delegate implementation builder for the [PContentKeySessionDelegate] protocol.
type ContentKeySessionDelegate struct {
	_ContentKeySessionDidGenerateExpiredSessionReport             func(session ContentKeySession)
	_ContentKeySessionContentProtectionSessionIdentifierDidChange func(session ContentKeySession)
	_ContentKeySessionDidProvideContentKeyRequest                 func(session ContentKeySession, keyRequest ContentKeyRequest)
}

func (di *ContentKeySessionDelegate) HasContentKeySessionDidGenerateExpiredSessionReport() bool {
	return di._ContentKeySessionDidGenerateExpiredSessionReport != nil
}

// Notifies the sender that an expired session report has been generated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2966513-contentkeysessiondidgenerateexpi?language=objc
func (di *ContentKeySessionDelegate) SetContentKeySessionDidGenerateExpiredSessionReport(f func(session ContentKeySession)) {
	di._ContentKeySessionDidGenerateExpiredSessionReport = f
}

// Notifies the sender that an expired session report has been generated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2966513-contentkeysessiondidgenerateexpi?language=objc
func (di *ContentKeySessionDelegate) ContentKeySessionDidGenerateExpiredSessionReport(session ContentKeySession) {
	di._ContentKeySessionDidGenerateExpiredSessionReport(session)
}
func (di *ContentKeySessionDelegate) HasContentKeySessionContentProtectionSessionIdentifierDidChange() bool {
	return di._ContentKeySessionContentProtectionSessionIdentifierDidChange != nil
}

// Tells the receiver the content protection session identifier changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799160-contentkeysessioncontentprotecti?language=objc
func (di *ContentKeySessionDelegate) SetContentKeySessionContentProtectionSessionIdentifierDidChange(f func(session ContentKeySession)) {
	di._ContentKeySessionContentProtectionSessionIdentifierDidChange = f
}

// Tells the receiver the content protection session identifier changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799160-contentkeysessioncontentprotecti?language=objc
func (di *ContentKeySessionDelegate) ContentKeySessionContentProtectionSessionIdentifierDidChange(session ContentKeySession) {
	di._ContentKeySessionContentProtectionSessionIdentifierDidChange(session)
}
func (di *ContentKeySessionDelegate) HasContentKeySessionDidProvideContentKeyRequest() bool {
	return di._ContentKeySessionDidProvideContentKeyRequest != nil
}

// Provides the receiver with a new content key request object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799204-contentkeysession?language=objc
func (di *ContentKeySessionDelegate) SetContentKeySessionDidProvideContentKeyRequest(f func(session ContentKeySession, keyRequest ContentKeyRequest)) {
	di._ContentKeySessionDidProvideContentKeyRequest = f
}

// Provides the receiver with a new content key request object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799204-contentkeysession?language=objc
func (di *ContentKeySessionDelegate) ContentKeySessionDidProvideContentKeyRequest(session ContentKeySession, keyRequest ContentKeyRequest) {
	di._ContentKeySessionDidProvideContentKeyRequest(session, keyRequest)
}

// A concrete type wrapper for the [PContentKeySessionDelegate] protocol.
type ContentKeySessionDelegateWrapper struct {
	objc.Object
}

func (c_ ContentKeySessionDelegateWrapper) HasContentKeySessionDidGenerateExpiredSessionReport() bool {
	return c_.RespondsToSelector(objc.Sel("contentKeySessionDidGenerateExpiredSessionReport:"))
}

// Notifies the sender that an expired session report has been generated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2966513-contentkeysessiondidgenerateexpi?language=objc
func (c_ ContentKeySessionDelegateWrapper) ContentKeySessionDidGenerateExpiredSessionReport(session IContentKeySession) {
	objc.Call[objc.Void](c_, objc.Sel("contentKeySessionDidGenerateExpiredSessionReport:"), objc.Ptr(session))
}

func (c_ ContentKeySessionDelegateWrapper) HasContentKeySessionContentProtectionSessionIdentifierDidChange() bool {
	return c_.RespondsToSelector(objc.Sel("contentKeySessionContentProtectionSessionIdentifierDidChange:"))
}

// Tells the receiver the content protection session identifier changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799160-contentkeysessioncontentprotecti?language=objc
func (c_ ContentKeySessionDelegateWrapper) ContentKeySessionContentProtectionSessionIdentifierDidChange(session IContentKeySession) {
	objc.Call[objc.Void](c_, objc.Sel("contentKeySessionContentProtectionSessionIdentifierDidChange:"), objc.Ptr(session))
}

func (c_ ContentKeySessionDelegateWrapper) HasContentKeySessionDidProvideContentKeyRequest() bool {
	return c_.RespondsToSelector(objc.Sel("contentKeySession:didProvideContentKeyRequest:"))
}

// Provides the receiver with a new content key request object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate/2799204-contentkeysession?language=objc
func (c_ ContentKeySessionDelegateWrapper) ContentKeySessionDidProvideContentKeyRequest(session IContentKeySession, keyRequest IContentKeyRequest) {
	objc.Call[objc.Void](c_, objc.Sel("contentKeySession:didProvideContentKeyRequest:"), objc.Ptr(session), objc.Ptr(keyRequest))
}
