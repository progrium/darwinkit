// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableAssetDownloadStorageManagementPolicy] class.
var MutableAssetDownloadStorageManagementPolicyClass = _MutableAssetDownloadStorageManagementPolicyClass{objc.GetClass("AVMutableAssetDownloadStorageManagementPolicy")}

type _MutableAssetDownloadStorageManagementPolicyClass struct {
	objc.Class
}

// An interface definition for the [MutableAssetDownloadStorageManagementPolicy] class.
type IMutableAssetDownloadStorageManagementPolicy interface {
	IAssetDownloadStorageManagementPolicy
	SetPriority(value AssetDownloadedAssetEvictionPriority)
	SetExpirationDate(value foundation.IDate)
}

// A mutable object that you use to create a new storage management policy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableassetdownloadstoragemanagementpolicy?language=objc
type MutableAssetDownloadStorageManagementPolicy struct {
	AssetDownloadStorageManagementPolicy
}

func MutableAssetDownloadStorageManagementPolicyFrom(ptr unsafe.Pointer) MutableAssetDownloadStorageManagementPolicy {
	return MutableAssetDownloadStorageManagementPolicy{
		AssetDownloadStorageManagementPolicy: AssetDownloadStorageManagementPolicyFrom(ptr),
	}
}

func (mc _MutableAssetDownloadStorageManagementPolicyClass) Alloc() MutableAssetDownloadStorageManagementPolicy {
	rv := objc.Call[MutableAssetDownloadStorageManagementPolicy](mc, objc.Sel("alloc"))
	return rv
}

func MutableAssetDownloadStorageManagementPolicy_Alloc() MutableAssetDownloadStorageManagementPolicy {
	return MutableAssetDownloadStorageManagementPolicyClass.Alloc()
}

func (mc _MutableAssetDownloadStorageManagementPolicyClass) New() MutableAssetDownloadStorageManagementPolicy {
	rv := objc.Call[MutableAssetDownloadStorageManagementPolicy](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableAssetDownloadStorageManagementPolicy() MutableAssetDownloadStorageManagementPolicy {
	return MutableAssetDownloadStorageManagementPolicyClass.New()
}

func (m_ MutableAssetDownloadStorageManagementPolicy) Init() MutableAssetDownloadStorageManagementPolicy {
	rv := objc.Call[MutableAssetDownloadStorageManagementPolicy](m_, objc.Sel("init"))
	return rv
}

// The eviction priority for a downloaded asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableassetdownloadstoragemanagementpolicy/2865558-priority?language=objc
func (m_ MutableAssetDownloadStorageManagementPolicy) SetPriority(value AssetDownloadedAssetEvictionPriority) {
	objc.Call[objc.Void](m_, objc.Sel("setPriority:"), value)
}

// The expiration date for an asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableassetdownloadstoragemanagementpolicy/2865564-expirationdate?language=objc
func (m_ MutableAssetDownloadStorageManagementPolicy) SetExpirationDate(value foundation.IDate) {
	objc.Call[objc.Void](m_, objc.Sel("setExpirationDate:"), objc.Ptr(value))
}
