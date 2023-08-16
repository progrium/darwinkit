// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// The delegate of an NSCache object implements this protocol to perform specialized actions when an object is about to be evicted or removed from the cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscachedelegate?language=objc
type PCacheDelegate interface {
	// optional
	CacheWillEvictObject(cache Cache, obj objc.Object)
	HasCacheWillEvictObject() bool
}

// A delegate implementation builder for the [PCacheDelegate] protocol.
type CacheDelegate struct {
	_CacheWillEvictObject func(cache Cache, obj objc.Object)
}

func (di *CacheDelegate) HasCacheWillEvictObject() bool {
	return di._CacheWillEvictObject != nil
}

// Called when an object is about to be evicted or removed from the cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscachedelegate/1416107-cache?language=objc
func (di *CacheDelegate) SetCacheWillEvictObject(f func(cache Cache, obj objc.Object)) {
	di._CacheWillEvictObject = f
}

// Called when an object is about to be evicted or removed from the cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscachedelegate/1416107-cache?language=objc
func (di *CacheDelegate) CacheWillEvictObject(cache Cache, obj objc.Object) {
	di._CacheWillEvictObject(cache, obj)
}

// A concrete type wrapper for the [PCacheDelegate] protocol.
type CacheDelegateWrapper struct {
	objc.Object
}

func (c_ CacheDelegateWrapper) HasCacheWillEvictObject() bool {
	return c_.RespondsToSelector(objc.Sel("cache:willEvictObject:"))
}

// Called when an object is about to be evicted or removed from the cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscachedelegate/1416107-cache?language=objc
func (c_ CacheDelegateWrapper) CacheWillEvictObject(cache ICache, obj objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("cache:willEvictObject:"), objc.Ptr(cache), obj)
}
