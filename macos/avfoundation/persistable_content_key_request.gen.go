// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PersistableContentKeyRequest] class.
var PersistableContentKeyRequestClass = _PersistableContentKeyRequestClass{objc.GetClass("AVPersistableContentKeyRequest")}

type _PersistableContentKeyRequestClass struct {
	objc.Class
}

// An interface definition for the [PersistableContentKeyRequest] class.
type IPersistableContentKeyRequest interface {
	IContentKeyRequest
	PersistableContentKeyFromKeyVendorResponseOptionsError(keyVendorResponse []byte, options map[string]objc.IObject, outError foundation.IError) []byte
}

// An object that encapsulates information about a persistable content decryption key request issued from a content key session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avpersistablecontentkeyrequest?language=objc
type PersistableContentKeyRequest struct {
	ContentKeyRequest
}

func PersistableContentKeyRequestFrom(ptr unsafe.Pointer) PersistableContentKeyRequest {
	return PersistableContentKeyRequest{
		ContentKeyRequest: ContentKeyRequestFrom(ptr),
	}
}

func (pc _PersistableContentKeyRequestClass) Alloc() PersistableContentKeyRequest {
	rv := objc.Call[PersistableContentKeyRequest](pc, objc.Sel("alloc"))
	return rv
}

func PersistableContentKeyRequest_Alloc() PersistableContentKeyRequest {
	return PersistableContentKeyRequestClass.Alloc()
}

func (pc _PersistableContentKeyRequestClass) New() PersistableContentKeyRequest {
	rv := objc.Call[PersistableContentKeyRequest](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistableContentKeyRequest() PersistableContentKeyRequest {
	return PersistableContentKeyRequestClass.New()
}

func (p_ PersistableContentKeyRequest) Init() PersistableContentKeyRequest {
	rv := objc.Call[PersistableContentKeyRequest](p_, objc.Sel("init"))
	return rv
}

// Creates a persistable content key from the content key context data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avpersistablecontentkeyrequest/2799188-persistablecontentkeyfromkeyvend?language=objc
func (p_ PersistableContentKeyRequest) PersistableContentKeyFromKeyVendorResponseOptionsError(keyVendorResponse []byte, options map[string]objc.IObject, outError foundation.IError) []byte {
	rv := objc.Call[[]byte](p_, objc.Sel("persistableContentKeyFromKeyVendorResponse:options:error:"), keyVendorResponse, options, objc.Ptr(outError))
	return rv
}
