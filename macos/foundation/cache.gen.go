// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Cache] class.
var CacheClass = _CacheClass{objc.GetClass("NSCache")}

type _CacheClass struct {
	objc.Class
}

// An interface definition for the [Cache] class.
type ICache interface {
	objc.IObject
	ObjectForKey(key objc.IObject) objc.Object
	SetObjectForKey(obj objc.IObject, key objc.IObject)
	RemoveObjectForKey(key objc.IObject)
	RemoveAllObjects()
	EvictsObjectsWithDiscardedContent() bool
	SetEvictsObjectsWithDiscardedContent(value bool)
	CountLimit() uint
	SetCountLimit(value uint)
	Name() string
	SetName(value string)
	Delegate() CacheDelegateWrapper
	SetDelegate(value PCacheDelegate)
	SetDelegateObject(valueObject objc.IObject)
	TotalCostLimit() uint
	SetTotalCostLimit(value uint)
}

// A mutable collection you use to temporarily store transient key-value pairs that are subject to eviction when resources are low. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscache?language=objc
type Cache struct {
	objc.Object
}

func CacheFrom(ptr unsafe.Pointer) Cache {
	return Cache{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CacheClass) Alloc() Cache {
	rv := objc.Call[Cache](cc, objc.Sel("alloc"))
	return rv
}

func Cache_Alloc() Cache {
	return CacheClass.Alloc()
}

func (cc _CacheClass) New() Cache {
	rv := objc.Call[Cache](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCache() Cache {
	return CacheClass.New()
}

func (c_ Cache) Init() Cache {
	rv := objc.Call[Cache](c_, objc.Sel("init"))
	return rv
}

// Returns the value associated with a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscache/1415458-objectforkey?language=objc
func (c_ Cache) ObjectForKey(key objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("objectForKey:"), objc.Ptr(key))
	return rv
}

// Sets the value of the specified key in the cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscache/1408223-setobject?language=objc
func (c_ Cache) SetObjectForKey(obj objc.IObject, key objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setObject:forKey:"), objc.Ptr(obj), objc.Ptr(key))
}

// Removes the value of the specified key in the cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscache/1409900-removeobjectforkey?language=objc
func (c_ Cache) RemoveObjectForKey(key objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("removeObjectForKey:"), objc.Ptr(key))
}

// Empties the cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscache/1411382-removeallobjects?language=objc
func (c_ Cache) RemoveAllObjects() {
	objc.Call[objc.Void](c_, objc.Sel("removeAllObjects"))
}

// Whether the cache will automatically evict discardable-content objects whose content has been discarded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscache/1408469-evictsobjectswithdiscardedconten?language=objc
func (c_ Cache) EvictsObjectsWithDiscardedContent() bool {
	rv := objc.Call[bool](c_, objc.Sel("evictsObjectsWithDiscardedContent"))
	return rv
}

// Whether the cache will automatically evict discardable-content objects whose content has been discarded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscache/1408469-evictsobjectswithdiscardedconten?language=objc
func (c_ Cache) SetEvictsObjectsWithDiscardedContent(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setEvictsObjectsWithDiscardedContent:"), value)
}

// The maximum number of objects the cache should hold. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscache/1416355-countlimit?language=objc
func (c_ Cache) CountLimit() uint {
	rv := objc.Call[uint](c_, objc.Sel("countLimit"))
	return rv
}

// The maximum number of objects the cache should hold. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscache/1416355-countlimit?language=objc
func (c_ Cache) SetCountLimit(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setCountLimit:"), value)
}

// The name of the cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscache/1409941-name?language=objc
func (c_ Cache) Name() string {
	rv := objc.Call[string](c_, objc.Sel("name"))
	return rv
}

// The name of the cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscache/1409941-name?language=objc
func (c_ Cache) SetName(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setName:"), value)
}

// The cache’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscache/1413061-delegate?language=objc
func (c_ Cache) Delegate() CacheDelegateWrapper {
	rv := objc.Call[CacheDelegateWrapper](c_, objc.Sel("delegate"))
	return rv
}

// The cache’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscache/1413061-delegate?language=objc
func (c_ Cache) SetDelegate(value PCacheDelegate) {
	po0 := objc.WrapAsProtocol("NSCacheDelegate", value)
	objc.Call[objc.Void](c_, objc.Sel("setDelegate:"), po0)
}

// The cache’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscache/1413061-delegate?language=objc
func (c_ Cache) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The maximum total cost that the cache can hold before it starts evicting objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscache/1407672-totalcostlimit?language=objc
func (c_ Cache) TotalCostLimit() uint {
	rv := objc.Call[uint](c_, objc.Sel("totalCostLimit"))
	return rv
}

// The maximum total cost that the cache can hold before it starts evicting objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscache/1407672-totalcostlimit?language=objc
func (c_ Cache) SetTotalCostLimit(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setTotalCostLimit:"), value)
}
