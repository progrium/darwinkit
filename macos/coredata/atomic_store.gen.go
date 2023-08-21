// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AtomicStore] class.
var AtomicStoreClass = _AtomicStoreClass{objc.GetClass("NSAtomicStore")}

type _AtomicStoreClass struct {
	objc.Class
}

// An interface definition for the [AtomicStore] class.
type IAtomicStore interface {
	IPersistentStore
	ReferenceObjectForObjectID(objectID IManagedObjectID) objc.Object
	NewReferenceObjectForManagedObject(managedObject IManagedObject) objc.Object
	Save(error foundation.IError) bool
	CacheNodeForObjectID(objectID IManagedObjectID) AtomicStoreCacheNode
	WillRemoveCacheNodes(cacheNodes foundation.ISet)
	ObjectIDForEntityReferenceObject(entity IEntityDescription, data objc.IObject) ManagedObjectID
	CacheNodes() foundation.Set
	NewCacheNodeForManagedObject(managedObject IManagedObject) AtomicStoreCacheNode
	UpdateCacheNodeFromManagedObject(node IAtomicStoreCacheNode, managedObject IManagedObject)
	AddCacheNodes(cacheNodes foundation.ISet)
	Load(error foundation.IError) bool
}

// An abstract superclass that you subclass to create a Core Data atomic store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstore?language=objc
type AtomicStore struct {
	PersistentStore
}

func AtomicStoreFrom(ptr unsafe.Pointer) AtomicStore {
	return AtomicStore{
		PersistentStore: PersistentStoreFrom(ptr),
	}
}

func (a_ AtomicStore) InitWithPersistentStoreCoordinatorConfigurationNameURLOptions(coordinator IPersistentStoreCoordinator, configurationName string, url foundation.IURL, options foundation.Dictionary) AtomicStore {
	rv := objc.Call[AtomicStore](a_, objc.Sel("initWithPersistentStoreCoordinator:configurationName:URL:options:"), objc.Ptr(coordinator), configurationName, objc.Ptr(url), options)
	return rv
}

// Creates an atomic store at the specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstore/1388054-initwithpersistentstorecoordinat?language=objc
func NewAtomicStoreWithPersistentStoreCoordinatorConfigurationNameURLOptions(coordinator IPersistentStoreCoordinator, configurationName string, url foundation.IURL, options foundation.Dictionary) AtomicStore {
	instance := AtomicStoreClass.Alloc().InitWithPersistentStoreCoordinatorConfigurationNameURLOptions(coordinator, configurationName, url, options)
	instance.Autorelease()
	return instance
}

func (ac _AtomicStoreClass) Alloc() AtomicStore {
	rv := objc.Call[AtomicStore](ac, objc.Sel("alloc"))
	return rv
}

func AtomicStore_Alloc() AtomicStore {
	return AtomicStoreClass.Alloc()
}

func (ac _AtomicStoreClass) New() AtomicStore {
	rv := objc.Call[AtomicStore](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAtomicStore() AtomicStore {
	return AtomicStoreClass.New()
}

func (a_ AtomicStore) Init() AtomicStore {
	rv := objc.Call[AtomicStore](a_, objc.Sel("init"))
	return rv
}

// Returns the reference object for a given managed object ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstore/1388046-referenceobjectforobjectid?language=objc
func (a_ AtomicStore) ReferenceObjectForObjectID(objectID IManagedObjectID) objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("referenceObjectForObjectID:"), objc.Ptr(objectID))
	return rv
}

// Returns a new reference object for a given managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstore/1388050-newreferenceobjectformanagedobje?language=objc
func (a_ AtomicStore) NewReferenceObjectForManagedObject(managedObject IManagedObject) objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("newReferenceObjectForManagedObject:"), objc.Ptr(managedObject))
	rv.Autorelease()
	return rv
}

// Saves the cache nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstore/1388056-save?language=objc
func (a_ AtomicStore) Save(error foundation.IError) bool {
	rv := objc.Call[bool](a_, objc.Sel("save:"), objc.Ptr(error))
	return rv
}

// Returns the cache node for a given managed object ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstore/1388040-cachenodeforobjectid?language=objc
func (a_ AtomicStore) CacheNodeForObjectID(objectID IManagedObjectID) AtomicStoreCacheNode {
	rv := objc.Call[AtomicStoreCacheNode](a_, objc.Sel("cacheNodeForObjectID:"), objc.Ptr(objectID))
	return rv
}

// Method invoked before the store removes the given collection of cache nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstore/1388064-willremovecachenodes?language=objc
func (a_ AtomicStore) WillRemoveCacheNodes(cacheNodes foundation.ISet) {
	objc.Call[objc.Void](a_, objc.Sel("willRemoveCacheNodes:"), objc.Ptr(cacheNodes))
}

// Returns a managed object ID from the reference data for a specified entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstore/1388058-objectidforentity?language=objc
func (a_ AtomicStore) ObjectIDForEntityReferenceObject(entity IEntityDescription, data objc.IObject) ManagedObjectID {
	rv := objc.Call[ManagedObjectID](a_, objc.Sel("objectIDForEntity:referenceObject:"), objc.Ptr(entity), data)
	return rv
}

// Returns the set of cache nodes registered with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstore/1388042-cachenodes?language=objc
func (a_ AtomicStore) CacheNodes() foundation.Set {
	rv := objc.Call[foundation.Set](a_, objc.Sel("cacheNodes"))
	return rv
}

// Returns a new cache node for a given managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstore/1388052-newcachenodeformanagedobject?language=objc
func (a_ AtomicStore) NewCacheNodeForManagedObject(managedObject IManagedObject) AtomicStoreCacheNode {
	rv := objc.Call[AtomicStoreCacheNode](a_, objc.Sel("newCacheNodeForManagedObject:"), objc.Ptr(managedObject))
	rv.Autorelease()
	return rv
}

// Updates the given cache node using the values in a given managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstore/1388044-updatecachenode?language=objc
func (a_ AtomicStore) UpdateCacheNodeFromManagedObject(node IAtomicStoreCacheNode, managedObject IManagedObject) {
	objc.Call[objc.Void](a_, objc.Sel("updateCacheNode:fromManagedObject:"), objc.Ptr(node), objc.Ptr(managedObject))
}

// Registers a set of cache nodes with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstore/1388062-addcachenodes?language=objc
func (a_ AtomicStore) AddCacheNodes(cacheNodes foundation.ISet) {
	objc.Call[objc.Void](a_, objc.Sel("addCacheNodes:"), objc.Ptr(cacheNodes))
}

// Loads the cache nodes for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstore/1388060-load?language=objc
func (a_ AtomicStore) Load(error foundation.IError) bool {
	rv := objc.Call[bool](a_, objc.Sel("load:"), objc.Ptr(error))
	return rv
}
