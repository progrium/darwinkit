// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetResourceLoadingRequestor] class.
var AssetResourceLoadingRequestorClass = _AssetResourceLoadingRequestorClass{objc.GetClass("AVAssetResourceLoadingRequestor")}

type _AssetResourceLoadingRequestorClass struct {
	objc.Class
}

// An interface definition for the [AssetResourceLoadingRequestor] class.
type IAssetResourceLoadingRequestor interface {
	objc.IObject
	ProvidesExpiredSessionReports() bool
}

// An object that contains information about the originator of a resource-loading request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequestor?language=objc
type AssetResourceLoadingRequestor struct {
	objc.Object
}

func AssetResourceLoadingRequestorFrom(ptr unsafe.Pointer) AssetResourceLoadingRequestor {
	return AssetResourceLoadingRequestor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetResourceLoadingRequestorClass) Alloc() AssetResourceLoadingRequestor {
	rv := objc.Call[AssetResourceLoadingRequestor](ac, objc.Sel("alloc"))
	return rv
}

func AssetResourceLoadingRequestor_Alloc() AssetResourceLoadingRequestor {
	return AssetResourceLoadingRequestorClass.Alloc()
}

func (ac _AssetResourceLoadingRequestorClass) New() AssetResourceLoadingRequestor {
	rv := objc.Call[AssetResourceLoadingRequestor](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetResourceLoadingRequestor() AssetResourceLoadingRequestor {
	return AssetResourceLoadingRequestorClass.New()
}

func (a_ AssetResourceLoadingRequestor) Init() AssetResourceLoadingRequestor {
	rv := objc.Call[AssetResourceLoadingRequestor](a_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the requestor provides expired session reports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequestor/2966511-providesexpiredsessionreports?language=objc
func (a_ AssetResourceLoadingRequestor) ProvidesExpiredSessionReports() bool {
	rv := objc.Call[bool](a_, objc.Sel("providesExpiredSessionReports"))
	return rv
}
