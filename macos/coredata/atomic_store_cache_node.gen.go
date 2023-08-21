// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AtomicStoreCacheNode] class.
var AtomicStoreCacheNodeClass = _AtomicStoreCacheNodeClass{objc.GetClass("NSAtomicStoreCacheNode")}

type _AtomicStoreCacheNodeClass struct {
	objc.Class
}

// An interface definition for the [AtomicStoreCacheNode] class.
type IAtomicStoreCacheNode interface {
	objc.IObject
	SetValueForKey(value objc.IObject, key string)
	ValueForKey(key string) objc.Object
	PropertyCache() foundation.MutableDictionary
	SetPropertyCache(value foundation.IMutableDictionary)
	ObjectID() ManagedObjectID
}

// A concrete class that you use to represent basic nodes in a Core Data atomic store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstorecachenode?language=objc
type AtomicStoreCacheNode struct {
	objc.Object
}

func AtomicStoreCacheNodeFrom(ptr unsafe.Pointer) AtomicStoreCacheNode {
	return AtomicStoreCacheNode{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ AtomicStoreCacheNode) InitWithObjectID(moid IManagedObjectID) AtomicStoreCacheNode {
	rv := objc.Call[AtomicStoreCacheNode](a_, objc.Sel("initWithObjectID:"), objc.Ptr(moid))
	return rv
}

// Returns a cache node for the given managed object ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstorecachenode/1506754-initwithobjectid?language=objc
func NewAtomicStoreCacheNodeWithObjectID(moid IManagedObjectID) AtomicStoreCacheNode {
	instance := AtomicStoreCacheNodeClass.Alloc().InitWithObjectID(moid)
	instance.Autorelease()
	return instance
}

func (ac _AtomicStoreCacheNodeClass) Alloc() AtomicStoreCacheNode {
	rv := objc.Call[AtomicStoreCacheNode](ac, objc.Sel("alloc"))
	return rv
}

func AtomicStoreCacheNode_Alloc() AtomicStoreCacheNode {
	return AtomicStoreCacheNodeClass.Alloc()
}

func (ac _AtomicStoreCacheNodeClass) New() AtomicStoreCacheNode {
	rv := objc.Call[AtomicStoreCacheNode](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAtomicStoreCacheNode() AtomicStoreCacheNode {
	return AtomicStoreCacheNodeClass.New()
}

func (a_ AtomicStoreCacheNode) Init() AtomicStoreCacheNode {
	rv := objc.Call[AtomicStoreCacheNode](a_, objc.Sel("init"))
	return rv
}

// Sets the value for the given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstorecachenode/1506456-setvalue?language=objc
func (a_ AtomicStoreCacheNode) SetValueForKey(value objc.IObject, key string) {
	objc.Call[objc.Void](a_, objc.Sel("setValue:forKey:"), value, key)
}

// Returns the value for a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstorecachenode/1506550-valueforkey?language=objc
func (a_ AtomicStoreCacheNode) ValueForKey(key string) objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("valueForKey:"), key)
	return rv
}

// The property cache dictionary of the node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstorecachenode/1506283-propertycache?language=objc
func (a_ AtomicStoreCacheNode) PropertyCache() foundation.MutableDictionary {
	rv := objc.Call[foundation.MutableDictionary](a_, objc.Sel("propertyCache"))
	return rv
}

// The property cache dictionary of the node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstorecachenode/1506283-propertycache?language=objc
func (a_ AtomicStoreCacheNode) SetPropertyCache(value foundation.IMutableDictionary) {
	objc.Call[objc.Void](a_, objc.Sel("setPropertyCache:"), objc.Ptr(value))
}

// The managed object ID of the node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstorecachenode/1506627-objectid?language=objc
func (a_ AtomicStoreCacheNode) ObjectID() ManagedObjectID {
	rv := objc.Call[ManagedObjectID](a_, objc.Sel("objectID"))
	return rv
}
