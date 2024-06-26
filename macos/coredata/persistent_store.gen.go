// Code generated by DarwinKit. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [PersistentStore] class.
var PersistentStoreClass = _PersistentStoreClass{objc.GetClass("NSPersistentStore")}

type _PersistentStoreClass struct {
	objc.Class
}

// An interface definition for the [PersistentStore] class.
type IPersistentStore interface {
	objc.IObject
	DidAddToPersistentStoreCoordinator(coordinator IPersistentStoreCoordinator)
	LoadMetadata(error unsafe.Pointer) bool
	WillRemoveFromPersistentStoreCoordinator(coordinator IPersistentStoreCoordinator)
	Metadata() map[string]objc.Object
	SetMetadata(value map[string]objc.IObject)
	ConfigurationName() string
	URL() foundation.URL
	SetURL(value foundation.IURL)
	CoreSpotlightExporter() CoreDataCoreSpotlightDelegate
	Options() foundation.Dictionary
	PersistentStoreCoordinator() PersistentStoreCoordinator
	Type() string
	Identifier() string
	SetIdentifier(value string)
	IsReadOnly() bool
	SetReadOnly(value bool)
}

// The abstract base class for all Core Data persistent stores. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore?language=objc
type PersistentStore struct {
	objc.Object
}

func PersistentStoreFrom(ptr unsafe.Pointer) PersistentStore {
	return PersistentStore{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PersistentStore) InitWithPersistentStoreCoordinatorConfigurationNameURLOptions(root IPersistentStoreCoordinator, name string, url foundation.IURL, options foundation.Dictionary) PersistentStore {
	rv := objc.Call[PersistentStore](p_, objc.Sel("initWithPersistentStoreCoordinator:configurationName:URL:options:"), root, name, url, options)
	return rv
}

// Returns a store initialized with the given arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506232-initwithpersistentstorecoordinat?language=objc
func NewPersistentStoreWithPersistentStoreCoordinatorConfigurationNameURLOptions(root IPersistentStoreCoordinator, name string, url foundation.IURL, options foundation.Dictionary) PersistentStore {
	instance := PersistentStoreClass.Alloc().InitWithPersistentStoreCoordinatorConfigurationNameURLOptions(root, name, url, options)
	instance.Autorelease()
	return instance
}

func (pc _PersistentStoreClass) Alloc() PersistentStore {
	rv := objc.Call[PersistentStore](pc, objc.Sel("alloc"))
	return rv
}

func (pc _PersistentStoreClass) New() PersistentStore {
	rv := objc.Call[PersistentStore](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentStore() PersistentStore {
	return PersistentStoreClass.New()
}

func (p_ PersistentStore) Init() PersistentStore {
	rv := objc.Call[PersistentStore](p_, objc.Sel("init"))
	return rv
}

// Invoked after the persistent store has been added to the persistent store coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506873-didaddtopersistentstorecoordinat?language=objc
func (p_ PersistentStore) DidAddToPersistentStoreCoordinator(coordinator IPersistentStoreCoordinator) {
	objc.Call[objc.Void](p_, objc.Sel("didAddToPersistentStoreCoordinator:"), coordinator)
}

// Returns the metadata from the persistent store at the given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506741-metadataforpersistentstorewithur?language=objc
func (pc _PersistentStoreClass) MetadataForPersistentStoreWithURLError(url foundation.IURL, error unsafe.Pointer) map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](pc, objc.Sel("metadataForPersistentStoreWithURL:error:"), url, error)
	return rv
}

// Returns the metadata from the persistent store at the given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506741-metadataforpersistentstorewithur?language=objc
func PersistentStore_MetadataForPersistentStoreWithURLError(url foundation.IURL, error unsafe.Pointer) map[string]objc.Object {
	return PersistentStoreClass.MetadataForPersistentStoreWithURLError(url, error)
}

// Instructs the persistent store to load its metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506273-loadmetadata?language=objc
func (p_ PersistentStore) LoadMetadata(error unsafe.Pointer) bool {
	rv := objc.Call[bool](p_, objc.Sel("loadMetadata:"), error)
	return rv
}

// Returns the migration manager class for this store class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506361-migrationmanagerclass?language=objc
func (pc _PersistentStoreClass) MigrationManagerClass() objc.Class {
	rv := objc.Call[objc.Class](pc, objc.Sel("migrationManagerClass"))
	return rv
}

// Returns the migration manager class for this store class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506361-migrationmanagerclass?language=objc
func PersistentStore_MigrationManagerClass() objc.Class {
	return PersistentStoreClass.MigrationManagerClass()
}

// Sets the metadata for the store at a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506824-setmetadata?language=objc
func (pc _PersistentStoreClass) SetMetadataForPersistentStoreWithURLError(metadata map[string]objc.IObject, url foundation.IURL, error unsafe.Pointer) bool {
	rv := objc.Call[bool](pc, objc.Sel("setMetadata:forPersistentStoreWithURL:error:"), metadata, url, error)
	return rv
}

// Sets the metadata for the store at a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506824-setmetadata?language=objc
func PersistentStore_SetMetadataForPersistentStoreWithURLError(metadata map[string]objc.IObject, url foundation.IURL, error unsafe.Pointer) bool {
	return PersistentStoreClass.SetMetadataForPersistentStoreWithURLError(metadata, url, error)
}

// Invoked before the persistent store is removed from the persistent store coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506731-willremovefrompersistentstorecoo?language=objc
func (p_ PersistentStore) WillRemoveFromPersistentStoreCoordinator(coordinator IPersistentStoreCoordinator) {
	objc.Call[objc.Void](p_, objc.Sel("willRemoveFromPersistentStoreCoordinator:"), coordinator)
}

// The metadata for the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506564-metadata?language=objc
func (p_ PersistentStore) Metadata() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](p_, objc.Sel("metadata"))
	return rv
}

// The metadata for the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506564-metadata?language=objc
func (p_ PersistentStore) SetMetadata(value map[string]objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("setMetadata:"), value)
}

// The name of the managed object model configuration that creates the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506620-configurationname?language=objc
func (p_ PersistentStore) ConfigurationName() string {
	rv := objc.Call[string](p_, objc.Sel("configurationName"))
	return rv
}

// The URL for the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506700-url?language=objc
func (p_ PersistentStore) URL() foundation.URL {
	rv := objc.Call[foundation.URL](p_, objc.Sel("URL"))
	return rv
}

// The URL for the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506700-url?language=objc
func (p_ PersistentStore) SetURL(value foundation.IURL) {
	objc.Call[objc.Void](p_, objc.Sel("setURL:"), value)
}

// The spotlight exporter associated with this persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/2897191-corespotlightexporter?language=objc
func (p_ PersistentStore) CoreSpotlightExporter() CoreDataCoreSpotlightDelegate {
	rv := objc.Call[CoreDataCoreSpotlightDelegate](p_, objc.Sel("coreSpotlightExporter"))
	return rv
}

// The options that Core Data uses to create the store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506821-options?language=objc
func (p_ PersistentStore) Options() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](p_, objc.Sel("options"))
	return rv
}

// The persistent store coordinator that loads the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506226-persistentstorecoordinator?language=objc
func (p_ PersistentStore) PersistentStoreCoordinator() PersistentStoreCoordinator {
	rv := objc.Call[PersistentStoreCoordinator](p_, objc.Sel("persistentStoreCoordinator"))
	return rv
}

// The type string of the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506250-type?language=objc
func (p_ PersistentStore) Type() string {
	rv := objc.Call[string](p_, objc.Sel("type"))
	return rv
}

// The unique identifier for the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506215-identifier?language=objc
func (p_ PersistentStore) Identifier() string {
	rv := objc.Call[string](p_, objc.Sel("identifier"))
	return rv
}

// The unique identifier for the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506215-identifier?language=objc
func (p_ PersistentStore) SetIdentifier(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setIdentifier:"), value)
}

// A Boolean value that indicates whether the persistent store is read-only. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506183-readonly?language=objc
func (p_ PersistentStore) IsReadOnly() bool {
	rv := objc.Call[bool](p_, objc.Sel("isReadOnly"))
	return rv
}

// A Boolean value that indicates whether the persistent store is read-only. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506183-readonly?language=objc
func (p_ PersistentStore) SetReadOnly(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setReadOnly:"), value)
}
