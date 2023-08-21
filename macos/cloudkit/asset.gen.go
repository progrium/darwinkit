// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Asset] class.
var AssetClass = _AssetClass{objc.GetClass("CKAsset")}

type _AssetClass struct {
	objc.Class
}

// An interface definition for the [Asset] class.
type IAsset interface {
	objc.IObject
	FileURL() foundation.URL
}

// An external file that belongs to a record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckasset?language=objc
type Asset struct {
	objc.Object
}

func AssetFrom(ptr unsafe.Pointer) Asset {
	return Asset{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ Asset) InitWithFileURL(fileURL foundation.IURL) Asset {
	rv := objc.Call[Asset](a_, objc.Sel("initWithFileURL:"), objc.Ptr(fileURL))
	return rv
}

// Creates an asset that references a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckasset/1514990-initwithfileurl?language=objc
func NewAssetWithFileURL(fileURL foundation.IURL) Asset {
	instance := AssetClass.Alloc().InitWithFileURL(fileURL)
	instance.Autorelease()
	return instance
}

func (ac _AssetClass) Alloc() Asset {
	rv := objc.Call[Asset](ac, objc.Sel("alloc"))
	return rv
}

func Asset_Alloc() Asset {
	return AssetClass.Alloc()
}

func (ac _AssetClass) New() Asset {
	rv := objc.Call[Asset](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAsset() Asset {
	return AssetClass.New()
}

func (a_ Asset) Init() Asset {
	rv := objc.Call[Asset](a_, objc.Sel("init"))
	return rv
}

// The URL for accessing the asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckasset/1515050-fileurl?language=objc
func (a_ Asset) FileURL() foundation.URL {
	rv := objc.Call[foundation.URL](a_, objc.Sel("fileURL"))
	return rv
}
