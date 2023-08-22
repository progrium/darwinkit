// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetResourceLoader] class.
var AssetResourceLoaderClass = _AssetResourceLoaderClass{objc.GetClass("AVAssetResourceLoader")}

type _AssetResourceLoaderClass struct {
	objc.Class
}

// An interface definition for the [AssetResourceLoader] class.
type IAssetResourceLoader interface {
	objc.IObject
	SetDelegateQueue(delegate PAssetResourceLoaderDelegate, delegateQueue dispatch.Queue)
	SetDelegateObjectQueue(delegateObject objc.IObject, delegateQueue dispatch.Queue)
	PreloadsEligibleContentKeys() bool
	SetPreloadsEligibleContentKeys(value bool)
	Delegate() AssetResourceLoaderDelegateWrapper
	DelegateQueue() dispatch.Queue
}

// An object that mediates resource requests from a URL asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloader?language=objc
type AssetResourceLoader struct {
	objc.Object
}

func AssetResourceLoaderFrom(ptr unsafe.Pointer) AssetResourceLoader {
	return AssetResourceLoader{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetResourceLoaderClass) Alloc() AssetResourceLoader {
	rv := objc.Call[AssetResourceLoader](ac, objc.Sel("alloc"))
	return rv
}

func AssetResourceLoader_Alloc() AssetResourceLoader {
	return AssetResourceLoaderClass.Alloc()
}

func (ac _AssetResourceLoaderClass) New() AssetResourceLoader {
	rv := objc.Call[AssetResourceLoader](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetResourceLoader() AssetResourceLoader {
	return AssetResourceLoaderClass.New()
}

func (a_ AssetResourceLoader) Init() AssetResourceLoader {
	rv := objc.Call[AssetResourceLoader](a_, objc.Sel("init"))
	return rv
}

// Sets the delegate and dispatch queue to use with the resource loader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloader/1388314-setdelegate?language=objc
func (a_ AssetResourceLoader) SetDelegateQueue(delegate PAssetResourceLoaderDelegate, delegateQueue dispatch.Queue) {
	po0 := objc.WrapAsProtocol("AVAssetResourceLoaderDelegate", delegate)
	objc.Call[objc.Void](a_, objc.Sel("setDelegate:queue:"), po0, delegateQueue)
}

// Sets the delegate and dispatch queue to use with the resource loader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloader/1388314-setdelegate?language=objc
func (a_ AssetResourceLoader) SetDelegateObjectQueue(delegateObject objc.IObject, delegateQueue dispatch.Queue) {
	objc.Call[objc.Void](a_, objc.Sel("setDelegate:queue:"), objc.Ptr(delegateObject), delegateQueue)
}

// A Boolean value that indicates whether content keys will be loaded as quickly as possible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloader/1386939-preloadseligiblecontentkeys?language=objc
func (a_ AssetResourceLoader) PreloadsEligibleContentKeys() bool {
	rv := objc.Call[bool](a_, objc.Sel("preloadsEligibleContentKeys"))
	return rv
}

// A Boolean value that indicates whether content keys will be loaded as quickly as possible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloader/1386939-preloadseligiblecontentkeys?language=objc
func (a_ AssetResourceLoader) SetPreloadsEligibleContentKeys(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setPreloadsEligibleContentKeys:"), value)
}

// The delegate object to use when handling resource requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloader/1387913-delegate?language=objc
func (a_ AssetResourceLoader) Delegate() AssetResourceLoaderDelegateWrapper {
	rv := objc.Call[AssetResourceLoaderDelegateWrapper](a_, objc.Sel("delegate"))
	return rv
}

// The dispatch queue to use when handling resource requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloader/1387678-delegatequeue?language=objc
func (a_ AssetResourceLoader) DelegateQueue() dispatch.Queue {
	rv := objc.Call[dispatch.Queue](a_, objc.Sel("delegateQueue"))
	return rv
}
