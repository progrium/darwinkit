// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContentKeyResponse] class.
var ContentKeyResponseClass = _ContentKeyResponseClass{objc.GetClass("AVContentKeyResponse")}

type _ContentKeyResponseClass struct {
	objc.Class
}

// An interface definition for the [ContentKeyResponse] class.
type IContentKeyResponse interface {
	objc.IObject
}

// An object that encapsulates information about a response to a content decryption key request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyresponse?language=objc
type ContentKeyResponse struct {
	objc.Object
}

func ContentKeyResponseFrom(ptr unsafe.Pointer) ContentKeyResponse {
	return ContentKeyResponse{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ContentKeyResponseClass) ContentKeyResponseWithClearKeyDataInitializationVector(keyData []byte, initializationVector []byte) ContentKeyResponse {
	rv := objc.Call[ContentKeyResponse](cc, objc.Sel("contentKeyResponseWithClearKeyData:initializationVector:"), keyData, initializationVector)
	return rv
}

// Creates a new key response object for key data and initialization vector sent in the clear. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyresponse/2867986-contentkeyresponsewithclearkeyda?language=objc
func ContentKeyResponse_ContentKeyResponseWithClearKeyDataInitializationVector(keyData []byte, initializationVector []byte) ContentKeyResponse {
	return ContentKeyResponseClass.ContentKeyResponseWithClearKeyDataInitializationVector(keyData, initializationVector)
}

func (cc _ContentKeyResponseClass) ContentKeyResponseWithAuthorizationTokenData(authorizationTokenData []byte) ContentKeyResponse {
	rv := objc.Call[ContentKeyResponse](cc, objc.Sel("contentKeyResponseWithAuthorizationTokenData:"), authorizationTokenData)
	return rv
}

// Creates a content key response with an authorization token. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyresponse/3088798-contentkeyresponsewithauthorizat?language=objc
func ContentKeyResponse_ContentKeyResponseWithAuthorizationTokenData(authorizationTokenData []byte) ContentKeyResponse {
	return ContentKeyResponseClass.ContentKeyResponseWithAuthorizationTokenData(authorizationTokenData)
}

func (cc _ContentKeyResponseClass) ContentKeyResponseWithFairPlayStreamingKeyResponseData(keyResponseData []byte) ContentKeyResponse {
	rv := objc.Call[ContentKeyResponse](cc, objc.Sel("contentKeyResponseWithFairPlayStreamingKeyResponseData:"), keyResponseData)
	return rv
}

// Creates a content key response with an encrypted key response data blob when FairPlay Streaming is the key delivery method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyresponse/2799183-contentkeyresponsewithfairplayst?language=objc
func ContentKeyResponse_ContentKeyResponseWithFairPlayStreamingKeyResponseData(keyResponseData []byte) ContentKeyResponse {
	return ContentKeyResponseClass.ContentKeyResponseWithFairPlayStreamingKeyResponseData(keyResponseData)
}

func (cc _ContentKeyResponseClass) Alloc() ContentKeyResponse {
	rv := objc.Call[ContentKeyResponse](cc, objc.Sel("alloc"))
	return rv
}

func ContentKeyResponse_Alloc() ContentKeyResponse {
	return ContentKeyResponseClass.Alloc()
}

func (cc _ContentKeyResponseClass) New() ContentKeyResponse {
	rv := objc.Call[ContentKeyResponse](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContentKeyResponse() ContentKeyResponse {
	return ContentKeyResponseClass.New()
}

func (c_ ContentKeyResponse) Init() ContentKeyResponse {
	rv := objc.Call[ContentKeyResponse](c_, objc.Sel("init"))
	return rv
}
