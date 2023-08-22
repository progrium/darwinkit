// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetDownloadStorageManager] class.
var AssetDownloadStorageManagerClass = _AssetDownloadStorageManagerClass{objc.GetClass("AVAssetDownloadStorageManager")}

type _AssetDownloadStorageManagerClass struct {
	objc.Class
}

// An interface definition for the [AssetDownloadStorageManager] class.
type IAssetDownloadStorageManager interface {
	objc.IObject
	StorageManagementPolicyForURL(downloadStorageURL foundation.IURL) AssetDownloadStorageManagementPolicy
	SetStorageManagementPolicyForURL(storageManagementPolicy IAssetDownloadStorageManagementPolicy, downloadStorageURL foundation.IURL)
}

// An object that manages policies to automatically purge downloaded assets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadstoragemanager?language=objc
type AssetDownloadStorageManager struct {
	objc.Object
}

func AssetDownloadStorageManagerFrom(ptr unsafe.Pointer) AssetDownloadStorageManager {
	return AssetDownloadStorageManager{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetDownloadStorageManagerClass) Alloc() AssetDownloadStorageManager {
	rv := objc.Call[AssetDownloadStorageManager](ac, objc.Sel("alloc"))
	return rv
}

func AssetDownloadStorageManager_Alloc() AssetDownloadStorageManager {
	return AssetDownloadStorageManagerClass.Alloc()
}

func (ac _AssetDownloadStorageManagerClass) New() AssetDownloadStorageManager {
	rv := objc.Call[AssetDownloadStorageManager](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetDownloadStorageManager() AssetDownloadStorageManager {
	return AssetDownloadStorageManagerClass.New()
}

func (a_ AssetDownloadStorageManager) Init() AssetDownloadStorageManager {
	rv := objc.Call[AssetDownloadStorageManager](a_, objc.Sel("init"))
	return rv
}

// Returns the shared storage manager instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadstoragemanager/2865563-shareddownloadstoragemanager?language=objc
func (ac _AssetDownloadStorageManagerClass) SharedDownloadStorageManager() AssetDownloadStorageManager {
	rv := objc.Call[AssetDownloadStorageManager](ac, objc.Sel("sharedDownloadStorageManager"))
	return rv
}

// Returns the shared storage manager instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadstoragemanager/2865563-shareddownloadstoragemanager?language=objc
func AssetDownloadStorageManager_SharedDownloadStorageManager() AssetDownloadStorageManager {
	return AssetDownloadStorageManagerClass.SharedDownloadStorageManager()
}

// Returns the storage management policy for a downloaded asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadstoragemanager/2865562-storagemanagementpolicyforurl?language=objc
func (a_ AssetDownloadStorageManager) StorageManagementPolicyForURL(downloadStorageURL foundation.IURL) AssetDownloadStorageManagementPolicy {
	rv := objc.Call[AssetDownloadStorageManagementPolicy](a_, objc.Sel("storageManagementPolicyForURL:"), objc.Ptr(downloadStorageURL))
	return rv
}

// Sets a storage policy for the downloaded asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadstoragemanager/2865557-setstoragemanagementpolicy?language=objc
func (a_ AssetDownloadStorageManager) SetStorageManagementPolicyForURL(storageManagementPolicy IAssetDownloadStorageManagementPolicy, downloadStorageURL foundation.IURL) {
	objc.Call[objc.Void](a_, objc.Sel("setStorageManagementPolicy:forURL:"), objc.Ptr(storageManagementPolicy), objc.Ptr(downloadStorageURL))
}
