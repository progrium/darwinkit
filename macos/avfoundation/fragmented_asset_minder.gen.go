// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FragmentedAssetMinder] class.
var FragmentedAssetMinderClass = _FragmentedAssetMinderClass{objc.GetClass("AVFragmentedAssetMinder")}

type _FragmentedAssetMinderClass struct {
	objc.Class
}

// An interface definition for the [FragmentedAssetMinder] class.
type IFragmentedAssetMinder interface {
	objc.IObject
	AddFragmentedAsset(asset IAsset)
	RemoveFragmentedAsset(asset IAsset)
	MindingInterval() foundation.TimeInterval
	SetMindingInterval(value foundation.TimeInterval)
	Assets() []Asset
}

// An object that periodically checks whether the system adds new fragments to a fragmented asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedassetminder?language=objc
type FragmentedAssetMinder struct {
	objc.Object
}

func FragmentedAssetMinderFrom(ptr unsafe.Pointer) FragmentedAssetMinder {
	return FragmentedAssetMinder{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FragmentedAssetMinderClass) FragmentedAssetMinderWithAssetMindingInterval(asset IAsset, mindingInterval foundation.TimeInterval) FragmentedAssetMinder {
	rv := objc.Call[FragmentedAssetMinder](fc, objc.Sel("fragmentedAssetMinderWithAsset:mindingInterval:"), objc.Ptr(asset), mindingInterval)
	return rv
}

// Creates a fragmented asset minder containing the specified asset and minding interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedassetminder/1387182-fragmentedassetminderwithasset?language=objc
func FragmentedAssetMinder_FragmentedAssetMinderWithAssetMindingInterval(asset IAsset, mindingInterval foundation.TimeInterval) FragmentedAssetMinder {
	return FragmentedAssetMinderClass.FragmentedAssetMinderWithAssetMindingInterval(asset, mindingInterval)
}

func (f_ FragmentedAssetMinder) InitWithAssetMindingInterval(asset IAsset, mindingInterval foundation.TimeInterval) FragmentedAssetMinder {
	rv := objc.Call[FragmentedAssetMinder](f_, objc.Sel("initWithAsset:mindingInterval:"), objc.Ptr(asset), mindingInterval)
	return rv
}

// Creates a fragmented asset minder that monitors the specified asset at the indicated minding interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedassetminder/2966508-initwithasset?language=objc
func NewFragmentedAssetMinderWithAssetMindingInterval(asset IAsset, mindingInterval foundation.TimeInterval) FragmentedAssetMinder {
	instance := FragmentedAssetMinderClass.Alloc().InitWithAssetMindingInterval(asset, mindingInterval)
	instance.Autorelease()
	return instance
}

func (fc _FragmentedAssetMinderClass) Alloc() FragmentedAssetMinder {
	rv := objc.Call[FragmentedAssetMinder](fc, objc.Sel("alloc"))
	return rv
}

func FragmentedAssetMinder_Alloc() FragmentedAssetMinder {
	return FragmentedAssetMinderClass.Alloc()
}

func (fc _FragmentedAssetMinderClass) New() FragmentedAssetMinder {
	rv := objc.Call[FragmentedAssetMinder](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFragmentedAssetMinder() FragmentedAssetMinder {
	return FragmentedAssetMinderClass.New()
}

func (f_ FragmentedAssetMinder) Init() FragmentedAssetMinder {
	rv := objc.Call[FragmentedAssetMinder](f_, objc.Sel("init"))
	return rv
}

// Adds a fragmented asset to the array of minded assets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedassetminder/1387483-addfragmentedasset?language=objc
func (f_ FragmentedAssetMinder) AddFragmentedAsset(asset IAsset) {
	objc.Call[objc.Void](f_, objc.Sel("addFragmentedAsset:"), objc.Ptr(asset))
}

// Removes a fragmented asset from the array of minded assets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedassetminder/1389856-removefragmentedasset?language=objc
func (f_ FragmentedAssetMinder) RemoveFragmentedAsset(asset IAsset) {
	objc.Call[objc.Void](f_, objc.Sel("removeFragmentedAsset:"), objc.Ptr(asset))
}

// An interval that specifies when to perform a check for additional fragments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedassetminder/1390760-mindinginterval?language=objc
func (f_ FragmentedAssetMinder) MindingInterval() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](f_, objc.Sel("mindingInterval"))
	return rv
}

// An interval that specifies when to perform a check for additional fragments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedassetminder/1390760-mindinginterval?language=objc
func (f_ FragmentedAssetMinder) SetMindingInterval(value foundation.TimeInterval) {
	objc.Call[objc.Void](f_, objc.Sel("setMindingInterval:"), value)
}

// The minded array of fragmented assets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedassetminder/1390319-assets?language=objc
func (f_ FragmentedAssetMinder) Assets() []Asset {
	rv := objc.Call[[]Asset](f_, objc.Sel("assets"))
	return rv
}
