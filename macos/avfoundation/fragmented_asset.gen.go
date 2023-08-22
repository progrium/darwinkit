// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FragmentedAsset] class.
var FragmentedAssetClass = _FragmentedAssetClass{objc.GetClass("AVFragmentedAsset")}

type _FragmentedAssetClass struct {
	objc.Class
}

// An interface definition for the [FragmentedAsset] class.
type IFragmentedAsset interface {
	IURLAsset
}

// An asset with a duration that the system can extend without modifying its existing media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedasset?language=objc
type FragmentedAsset struct {
	URLAsset
}

func FragmentedAssetFrom(ptr unsafe.Pointer) FragmentedAsset {
	return FragmentedAsset{
		URLAsset: URLAssetFrom(ptr),
	}
}

func (fc _FragmentedAssetClass) FragmentedAssetWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) FragmentedAsset {
	rv := objc.Call[FragmentedAsset](fc, objc.Sel("fragmentedAssetWithURL:options:"), objc.Ptr(URL), options)
	return rv
}

// Creates a fragmented asset for the media at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedasset/1508700-fragmentedassetwithurl?language=objc
func FragmentedAsset_FragmentedAssetWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) FragmentedAsset {
	return FragmentedAssetClass.FragmentedAssetWithURLOptions(URL, options)
}

func (fc _FragmentedAssetClass) Alloc() FragmentedAsset {
	rv := objc.Call[FragmentedAsset](fc, objc.Sel("alloc"))
	return rv
}

func FragmentedAsset_Alloc() FragmentedAsset {
	return FragmentedAssetClass.Alloc()
}

func (fc _FragmentedAssetClass) New() FragmentedAsset {
	rv := objc.Call[FragmentedAsset](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFragmentedAsset() FragmentedAsset {
	return FragmentedAssetClass.New()
}

func (f_ FragmentedAsset) Init() FragmentedAsset {
	rv := objc.Call[FragmentedAsset](f_, objc.Sel("init"))
	return rv
}

func (fc _FragmentedAssetClass) URLAssetWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) FragmentedAsset {
	rv := objc.Call[FragmentedAsset](fc, objc.Sel("URLAssetWithURL:options:"), objc.Ptr(URL), options)
	return rv
}

// Returns an asset that models the media resource found at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/1508727-urlassetwithurl?language=objc
func FragmentedAsset_URLAssetWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) FragmentedAsset {
	return FragmentedAssetClass.URLAssetWithURLOptions(URL, options)
}

func (f_ FragmentedAsset) InitWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) FragmentedAsset {
	rv := objc.Call[FragmentedAsset](f_, objc.Sel("initWithURL:options:"), objc.Ptr(URL), options)
	return rv
}

// Creates an asset that models the media resource at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/1385698-initwithurl?language=objc
func NewFragmentedAssetWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) FragmentedAsset {
	instance := FragmentedAssetClass.Alloc().InitWithURLOptions(URL, options)
	instance.Autorelease()
	return instance
}

func (fc _FragmentedAssetClass) AssetWithURL(URL foundation.IURL) FragmentedAsset {
	rv := objc.Call[FragmentedAsset](fc, objc.Sel("assetWithURL:"), objc.Ptr(URL))
	return rv
}

// Creates an asset that models the media at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1389943-assetwithurl?language=objc
func FragmentedAsset_AssetWithURL(URL foundation.IURL) FragmentedAsset {
	return FragmentedAssetClass.AssetWithURL(URL)
}
