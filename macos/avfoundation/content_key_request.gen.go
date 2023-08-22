// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContentKeyRequest] class.
var ContentKeyRequestClass = _ContentKeyRequestClass{objc.GetClass("AVContentKeyRequest")}

type _ContentKeyRequestClass struct {
	objc.Class
}

// An interface definition for the [ContentKeyRequest] class.
type IContentKeyRequest interface {
	objc.IObject
	ProcessContentKeyResponseError(error foundation.IError)
	ProcessContentKeyResponse(keyResponse IContentKeyResponse)
	MakeStreamingContentKeyRequestDataForAppContentIdentifierOptionsCompletionHandler(appIdentifier []byte, contentIdentifier []byte, options map[string]objc.IObject, handler func(contentKeyRequestData []byte, error foundation.Error))
	RenewsExpiringResponseData() bool
	Error() foundation.Error
	Options() map[string]objc.Object
	ContentKey() ContentKey
	CanProvidePersistableContentKey() bool
	InitializationData() []byte
	ContentKeySpecifier() ContentKeySpecifier
	Status() ContentKeyRequestStatus
	Identifier() objc.Object
}

// An object that encapsulates information about a content decryption key request issued from a content key session object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequest?language=objc
type ContentKeyRequest struct {
	objc.Object
}

func ContentKeyRequestFrom(ptr unsafe.Pointer) ContentKeyRequest {
	return ContentKeyRequest{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ContentKeyRequestClass) Alloc() ContentKeyRequest {
	rv := objc.Call[ContentKeyRequest](cc, objc.Sel("alloc"))
	return rv
}

func ContentKeyRequest_Alloc() ContentKeyRequest {
	return ContentKeyRequestClass.Alloc()
}

func (cc _ContentKeyRequestClass) New() ContentKeyRequest {
	rv := objc.Call[ContentKeyRequest](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContentKeyRequest() ContentKeyRequest {
	return ContentKeyRequestClass.New()
}

func (c_ ContentKeyRequest) Init() ContentKeyRequest {
	rv := objc.Call[ContentKeyRequest](c_, objc.Sel("init"))
	return rv
}

// Tells the receiver that the app was unable to obtain a content key response. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequest/2799159-processcontentkeyresponseerror?language=objc
func (c_ ContentKeyRequest) ProcessContentKeyResponseError(error foundation.IError) {
	objc.Call[objc.Void](c_, objc.Sel("processContentKeyResponseError:"), objc.Ptr(error))
}

// Sends the specified content key response to the receiver for processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequest/2799165-processcontentkeyresponse?language=objc
func (c_ ContentKeyRequest) ProcessContentKeyResponse(keyResponse IContentKeyResponse) {
	objc.Call[objc.Void](c_, objc.Sel("processContentKeyResponse:"), objc.Ptr(keyResponse))
}

// Obtains encrypted key request data for a specific combination of app and content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequest/2799198-makestreamingcontentkeyrequestda?language=objc
func (c_ ContentKeyRequest) MakeStreamingContentKeyRequestDataForAppContentIdentifierOptionsCompletionHandler(appIdentifier []byte, contentIdentifier []byte, options map[string]objc.IObject, handler func(contentKeyRequestData []byte, error foundation.Error)) {
	objc.Call[objc.Void](c_, objc.Sel("makeStreamingContentKeyRequestDataForApp:contentIdentifier:options:completionHandler:"), appIdentifier, contentIdentifier, options, handler)
}

// A Boolean value that indicates whether the content key request renews previously provided response data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequest/2799193-renewsexpiringresponsedata?language=objc
func (c_ ContentKeyRequest) RenewsExpiringResponseData() bool {
	rv := objc.Call[bool](c_, objc.Sel("renewsExpiringResponseData"))
	return rv
}

// The error description for a failed key request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequest/2799182-error?language=objc
func (c_ ContentKeyRequest) Error() foundation.Error {
	rv := objc.Call[foundation.Error](c_, objc.Sel("error"))
	return rv
}

// A dictionary of options used to initialize key loading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequest/3112525-options?language=objc
func (c_ ContentKeyRequest) Options() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](c_, objc.Sel("options"))
	return rv
}

// The generated content key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequest/3726102-contentkey?language=objc
func (c_ ContentKeyRequest) ContentKey() ContentKey {
	rv := objc.Call[ContentKey](c_, objc.Sel("contentKey"))
	return rv
}

// The content key request used to create a persistable content key or respond to a previous request with a persistable content key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequest/2799185-canprovidepersistablecontentkey?language=objc
func (c_ ContentKeyRequest) CanProvidePersistableContentKey() bool {
	rv := objc.Call[bool](c_, objc.Sel("canProvidePersistableContentKey"))
	return rv
}

// The data used to obtain a key response. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequest/2799186-initializationdata?language=objc
func (c_ ContentKeyRequest) InitializationData() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("initializationData"))
	return rv
}

// The requested content key specifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequest/3726103-contentkeyspecifier?language=objc
func (c_ ContentKeyRequest) ContentKeySpecifier() ContentKeySpecifier {
	rv := objc.Call[ContentKeySpecifier](c_, objc.Sel("contentKeySpecifier"))
	return rv
}

// The current state of the content key request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequest/2799190-status?language=objc
func (c_ ContentKeyRequest) Status() ContentKeyRequestStatus {
	rv := objc.Call[ContentKeyRequestStatus](c_, objc.Sel("status"))
	return rv
}

// The identifier for the content key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequest/2799205-identifier?language=objc
func (c_ ContentKeyRequest) Identifier() objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("identifier"))
	return rv
}
