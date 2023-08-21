// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corespotlight"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CoreDataCoreSpotlightDelegate] class.
var CoreDataCoreSpotlightDelegateClass = _CoreDataCoreSpotlightDelegateClass{objc.GetClass("NSCoreDataCoreSpotlightDelegate")}

type _CoreDataCoreSpotlightDelegateClass struct {
	objc.Class
}

// An interface definition for the [CoreDataCoreSpotlightDelegate] class.
type ICoreDataCoreSpotlightDelegate interface {
	objc.IObject
	StartSpotlightIndexing()
	SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex corespotlight.ISearchableIndex, acknowledgementHandler func())
	AttributeSetForObject(object IManagedObject) corespotlight.SearchableItemAttributeSet
	DomainIdentifier() string
	StopSpotlightIndexing()
	DeleteSpotlightIndexWithCompletionHandler(completionHandler func(error foundation.Error))
	IndexName() string
	IsIndexingEnabled() bool
}

// A set of methods that enable integration with Core Spotlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nscoredatacorespotlightdelegate?language=objc
type CoreDataCoreSpotlightDelegate struct {
	objc.Object
}

func CoreDataCoreSpotlightDelegateFrom(ptr unsafe.Pointer) CoreDataCoreSpotlightDelegate {
	return CoreDataCoreSpotlightDelegate{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ CoreDataCoreSpotlightDelegate) InitForStoreWithDescriptionCoordinator(description IPersistentStoreDescription, psc IPersistentStoreCoordinator) CoreDataCoreSpotlightDelegate {
	rv := objc.Call[CoreDataCoreSpotlightDelegate](c_, objc.Sel("initForStoreWithDescription:coordinator:"), objc.Ptr(description), objc.Ptr(psc))
	return rv
}

// Creates a Core Spotlight delegate with the specified store description and coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nscoredatacorespotlightdelegate/3751984-initforstorewithdescription?language=objc
func NewCoreDataCoreSpotlightDelegateForStoreWithDescriptionCoordinator(description IPersistentStoreDescription, psc IPersistentStoreCoordinator) CoreDataCoreSpotlightDelegate {
	instance := CoreDataCoreSpotlightDelegateClass.Alloc().InitForStoreWithDescriptionCoordinator(description, psc)
	instance.Autorelease()
	return instance
}

func (cc _CoreDataCoreSpotlightDelegateClass) Alloc() CoreDataCoreSpotlightDelegate {
	rv := objc.Call[CoreDataCoreSpotlightDelegate](cc, objc.Sel("alloc"))
	return rv
}

func CoreDataCoreSpotlightDelegate_Alloc() CoreDataCoreSpotlightDelegate {
	return CoreDataCoreSpotlightDelegateClass.Alloc()
}

func (cc _CoreDataCoreSpotlightDelegateClass) New() CoreDataCoreSpotlightDelegate {
	rv := objc.Call[CoreDataCoreSpotlightDelegate](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCoreDataCoreSpotlightDelegate() CoreDataCoreSpotlightDelegate {
	return CoreDataCoreSpotlightDelegateClass.New()
}

func (c_ CoreDataCoreSpotlightDelegate) Init() CoreDataCoreSpotlightDelegate {
	rv := objc.Call[CoreDataCoreSpotlightDelegate](c_, objc.Sel("init"))
	return rv
}

// Starts the indexing of the store’s entities. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nscoredatacorespotlightdelegate/3751985-startspotlightindexing?language=objc
func (c_ CoreDataCoreSpotlightDelegate) StartSpotlightIndexing() {
	objc.Call[objc.Void](c_, objc.Sel("startSpotlightIndexing"))
}

// Reindexes all searchable items and clears any local state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nscoredatacorespotlightdelegate/2897201-searchableindex?language=objc
func (c_ CoreDataCoreSpotlightDelegate) SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex corespotlight.ISearchableIndex, acknowledgementHandler func()) {
	objc.Call[objc.Void](c_, objc.Sel("searchableIndex:reindexAllSearchableItemsWithAcknowledgementHandler:"), objc.Ptr(searchableIndex), acknowledgementHandler)
}

// Returns the searchable attributes for the specified managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nscoredatacorespotlightdelegate/2897197-attributesetforobject?language=objc
func (c_ CoreDataCoreSpotlightDelegate) AttributeSetForObject(object IManagedObject) corespotlight.SearchableItemAttributeSet {
	rv := objc.Call[corespotlight.SearchableItemAttributeSet](c_, objc.Sel("attributeSetForObject:"), objc.Ptr(object))
	return rv
}

// Returns the domain identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nscoredatacorespotlightdelegate/2897202-domainidentifier?language=objc
func (c_ CoreDataCoreSpotlightDelegate) DomainIdentifier() string {
	rv := objc.Call[string](c_, objc.Sel("domainIdentifier"))
	return rv
}

// Stops the indexing of the store’s entities. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nscoredatacorespotlightdelegate/3751986-stopspotlightindexing?language=objc
func (c_ CoreDataCoreSpotlightDelegate) StopSpotlightIndexing() {
	objc.Call[objc.Void](c_, objc.Sel("stopSpotlightIndexing"))
}

// Deletes all searchable items from the configured index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nscoredatacorespotlightdelegate/3751982-deletespotlightindexwithcompleti?language=objc
func (c_ CoreDataCoreSpotlightDelegate) DeleteSpotlightIndexWithCompletionHandler(completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](c_, objc.Sel("deleteSpotlightIndexWithCompletionHandler:"), completionHandler)
}

// Returns the index’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nscoredatacorespotlightdelegate/2897199-indexname?language=objc
func (c_ CoreDataCoreSpotlightDelegate) IndexName() string {
	rv := objc.Call[string](c_, objc.Sel("indexName"))
	return rv
}

// A Boolean value that indicates whether Core Data is currently updating the Core Spotlight index with the persistent store’s entities. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nscoredatacorespotlightdelegate/3751983-indexingenabled?language=objc
func (c_ CoreDataCoreSpotlightDelegate) IsIndexingEnabled() bool {
	rv := objc.Call[bool](c_, objc.Sel("isIndexingEnabled"))
	return rv
}
