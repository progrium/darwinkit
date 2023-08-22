// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetDownloadStorageManagementPolicy] class.
var AssetDownloadStorageManagementPolicyClass = _AssetDownloadStorageManagementPolicyClass{objc.GetClass("AVAssetDownloadStorageManagementPolicy")}

type _AssetDownloadStorageManagementPolicyClass struct {
	objc.Class
}

// An interface definition for the [AssetDownloadStorageManagementPolicy] class.
type IAssetDownloadStorageManagementPolicy interface {
	objc.IObject
	Priority() AssetDownloadedAssetEvictionPriority
	ExpirationDate() foundation.Date
}

// An object that defines a policy to automatically manage the storage of downloaded assets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadstoragemanagementpolicy?language=objc
type AssetDownloadStorageManagementPolicy struct {
	objc.Object
}

func AssetDownloadStorageManagementPolicyFrom(ptr unsafe.Pointer) AssetDownloadStorageManagementPolicy {
	return AssetDownloadStorageManagementPolicy{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetDownloadStorageManagementPolicyClass) Alloc() AssetDownloadStorageManagementPolicy {
	rv := objc.Call[AssetDownloadStorageManagementPolicy](ac, objc.Sel("alloc"))
	return rv
}

func AssetDownloadStorageManagementPolicy_Alloc() AssetDownloadStorageManagementPolicy {
	return AssetDownloadStorageManagementPolicyClass.Alloc()
}

func (ac _AssetDownloadStorageManagementPolicyClass) New() AssetDownloadStorageManagementPolicy {
	rv := objc.Call[AssetDownloadStorageManagementPolicy](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetDownloadStorageManagementPolicy() AssetDownloadStorageManagementPolicy {
	return AssetDownloadStorageManagementPolicyClass.New()
}

func (a_ AssetDownloadStorageManagementPolicy) Init() AssetDownloadStorageManagementPolicy {
	rv := objc.Call[AssetDownloadStorageManagementPolicy](a_, objc.Sel("init"))
	return rv
}

// The eviction priority for an asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadstoragemanagementpolicy/2865565-priority?language=objc
func (a_ AssetDownloadStorageManagementPolicy) Priority() AssetDownloadedAssetEvictionPriority {
	rv := objc.Call[AssetDownloadedAssetEvictionPriority](a_, objc.Sel("priority"))
	return rv
}

// The expiration date for an asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadstoragemanagementpolicy/2865559-expirationdate?language=objc
func (a_ AssetDownloadStorageManagementPolicy) ExpirationDate() foundation.Date {
	rv := objc.Call[foundation.Date](a_, objc.Sel("expirationDate"))
	return rv
}
