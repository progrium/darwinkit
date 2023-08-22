// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetResourceRenewalRequest] class.
var AssetResourceRenewalRequestClass = _AssetResourceRenewalRequestClass{objc.GetClass("AVAssetResourceRenewalRequest")}

type _AssetResourceRenewalRequestClass struct {
	objc.Class
}

// An interface definition for the [AssetResourceRenewalRequest] class.
type IAssetResourceRenewalRequest interface {
	IAssetResourceLoadingRequest
}

// An object that encapsulates information about a resource request from a resource loader to renew a previously issued request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourcerenewalrequest?language=objc
type AssetResourceRenewalRequest struct {
	AssetResourceLoadingRequest
}

func AssetResourceRenewalRequestFrom(ptr unsafe.Pointer) AssetResourceRenewalRequest {
	return AssetResourceRenewalRequest{
		AssetResourceLoadingRequest: AssetResourceLoadingRequestFrom(ptr),
	}
}

func (ac _AssetResourceRenewalRequestClass) Alloc() AssetResourceRenewalRequest {
	rv := objc.Call[AssetResourceRenewalRequest](ac, objc.Sel("alloc"))
	return rv
}

func AssetResourceRenewalRequest_Alloc() AssetResourceRenewalRequest {
	return AssetResourceRenewalRequestClass.Alloc()
}

func (ac _AssetResourceRenewalRequestClass) New() AssetResourceRenewalRequest {
	rv := objc.Call[AssetResourceRenewalRequest](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetResourceRenewalRequest() AssetResourceRenewalRequest {
	return AssetResourceRenewalRequestClass.New()
}

func (a_ AssetResourceRenewalRequest) Init() AssetResourceRenewalRequest {
	rv := objc.Call[AssetResourceRenewalRequest](a_, objc.Sel("init"))
	return rv
}
