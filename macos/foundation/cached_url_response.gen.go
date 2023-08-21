// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CachedURLResponse] class.
var CachedURLResponseClass = _CachedURLResponseClass{objc.GetClass("NSCachedURLResponse")}

type _CachedURLResponseClass struct {
	objc.Class
}

// An interface definition for the [CachedURLResponse] class.
type ICachedURLResponse interface {
	objc.IObject
	UserInfo() Dictionary
	Data() []byte
	Response() URLResponse
	StoragePolicy() URLCacheStoragePolicy
}

// A cached response to a URL request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscachedurlresponse?language=objc
type CachedURLResponse struct {
	objc.Object
}

func CachedURLResponseFrom(ptr unsafe.Pointer) CachedURLResponse {
	return CachedURLResponse{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ CachedURLResponse) InitWithResponseData(response IURLResponse, data []byte) CachedURLResponse {
	rv := objc.Call[CachedURLResponse](c_, objc.Sel("initWithResponse:data:"), objc.Ptr(response), data)
	return rv
}

// Creates a cached URL response instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscachedurlresponse/1413035-initwithresponse?language=objc
func NewCachedURLResponseWithResponseData(response IURLResponse, data []byte) CachedURLResponse {
	instance := CachedURLResponseClass.Alloc().InitWithResponseData(response, data)
	instance.Autorelease()
	return instance
}

func (cc _CachedURLResponseClass) Alloc() CachedURLResponse {
	rv := objc.Call[CachedURLResponse](cc, objc.Sel("alloc"))
	return rv
}

func CachedURLResponse_Alloc() CachedURLResponse {
	return CachedURLResponseClass.Alloc()
}

func (cc _CachedURLResponseClass) New() CachedURLResponse {
	rv := objc.Call[CachedURLResponse](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCachedURLResponse() CachedURLResponse {
	return CachedURLResponseClass.New()
}

func (c_ CachedURLResponse) Init() CachedURLResponse {
	rv := objc.Call[CachedURLResponse](c_, objc.Sel("init"))
	return rv
}

// The cached response’s user info dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscachedurlresponse/1411900-userinfo?language=objc
func (c_ CachedURLResponse) UserInfo() Dictionary {
	rv := objc.Call[Dictionary](c_, objc.Sel("userInfo"))
	return rv
}

// The cached response’s data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscachedurlresponse/1414011-data?language=objc
func (c_ CachedURLResponse) Data() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("data"))
	return rv
}

// The URL response object associated with the instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscachedurlresponse/1411077-response?language=objc
func (c_ CachedURLResponse) Response() URLResponse {
	rv := objc.Call[URLResponse](c_, objc.Sel("response"))
	return rv
}

// The cached response’s storage policy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscachedurlresponse/1412269-storagepolicy?language=objc
func (c_ CachedURLResponse) StoragePolicy() URLCacheStoragePolicy {
	rv := objc.Call[URLCacheStoragePolicy](c_, objc.Sel("storagePolicy"))
	return rv
}
