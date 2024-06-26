// Code generated by DarwinKit. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [PersistentStoreDescription] class.
var PersistentStoreDescriptionClass = _PersistentStoreDescriptionClass{objc.GetClass("NSPersistentStoreDescription")}

type _PersistentStoreDescriptionClass struct {
	objc.Class
}

// An interface definition for the [PersistentStoreDescription] class.
type IPersistentStoreDescription interface {
	objc.IObject
	SetOptionForKey(option objc.IObject, key string)
	SetValueForPragmaNamed(value objc.IObject, name string)
	URL() foundation.URL
	SetURL(value foundation.IURL)
	ShouldInferMappingModelAutomatically() bool
	SetShouldInferMappingModelAutomatically(value bool)
	Configuration() string
	SetConfiguration(value string)
	ShouldAddStoreAsynchronously() bool
	SetShouldAddStoreAsynchronously(value bool)
	SqlitePragmas() map[string]objc.Object
	CloudKitContainerOptions() PersistentCloudKitContainerOptions
	SetCloudKitContainerOptions(value IPersistentCloudKitContainerOptions)
	ShouldMigrateStoreAutomatically() bool
	SetShouldMigrateStoreAutomatically(value bool)
	Type() string
	SetType(value string)
	Options() map[string]objc.Object
	Timeout() foundation.TimeInterval
	SetTimeout(value foundation.TimeInterval)
	IsReadOnly() bool
	SetReadOnly(value bool)
}

// A description object used to create and load a persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription?language=objc
type PersistentStoreDescription struct {
	objc.Object
}

func PersistentStoreDescriptionFrom(ptr unsafe.Pointer) PersistentStoreDescription {
	return PersistentStoreDescription{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PersistentStoreDescription) InitWithURL(url foundation.IURL) PersistentStoreDescription {
	rv := objc.Call[PersistentStoreDescription](p_, objc.Sel("initWithURL:"), url)
	return rv
}

// Initializes the receiver with a URL for the store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640563-initwithurl?language=objc
func NewPersistentStoreDescriptionWithURL(url foundation.IURL) PersistentStoreDescription {
	instance := PersistentStoreDescriptionClass.Alloc().InitWithURL(url)
	instance.Autorelease()
	return instance
}

func (pc _PersistentStoreDescriptionClass) PersistentStoreDescriptionWithURL(URL foundation.IURL) PersistentStoreDescription {
	rv := objc.Call[PersistentStoreDescription](pc, objc.Sel("persistentStoreDescriptionWithURL:"), URL)
	return rv
}

// Initializes and returns a persistent store description with the given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1646473-persistentstoredescriptionwithur?language=objc
func PersistentStoreDescription_PersistentStoreDescriptionWithURL(URL foundation.IURL) PersistentStoreDescription {
	return PersistentStoreDescriptionClass.PersistentStoreDescriptionWithURL(URL)
}

func (pc _PersistentStoreDescriptionClass) Alloc() PersistentStoreDescription {
	rv := objc.Call[PersistentStoreDescription](pc, objc.Sel("alloc"))
	return rv
}

func (pc _PersistentStoreDescriptionClass) New() PersistentStoreDescription {
	rv := objc.Call[PersistentStoreDescription](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentStoreDescription() PersistentStoreDescription {
	return PersistentStoreDescriptionClass.New()
}

func (p_ PersistentStoreDescription) Init() PersistentStoreDescription {
	rv := objc.Call[PersistentStoreDescription](p_, objc.Sel("init"))
	return rv
}

// Sets an option on the store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640574-setoption?language=objc
func (p_ PersistentStoreDescription) SetOptionForKey(option objc.IObject, key string) {
	objc.Call[objc.Void](p_, objc.Sel("setOption:forKey:"), option, key)
}

// Allows you to set pragmas for the SQLite store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640598-setvalue?language=objc
func (p_ PersistentStoreDescription) SetValueForPragmaNamed(value objc.IObject, name string) {
	objc.Call[objc.Void](p_, objc.Sel("setValue:forPragmaNamed:"), value, name)
}

// The URL that the store will use for its location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640616-url?language=objc
func (p_ PersistentStoreDescription) URL() foundation.URL {
	rv := objc.Call[foundation.URL](p_, objc.Sel("URL"))
	return rv
}

// The URL that the store will use for its location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640616-url?language=objc
func (p_ PersistentStoreDescription) SetURL(value foundation.IURL) {
	objc.Call[objc.Void](p_, objc.Sel("setURL:"), value)
}

// A flag indicating whether a mapping model should be created automatically. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640623-shouldinfermappingmodelautomatic?language=objc
func (p_ PersistentStoreDescription) ShouldInferMappingModelAutomatically() bool {
	rv := objc.Call[bool](p_, objc.Sel("shouldInferMappingModelAutomatically"))
	return rv
}

// A flag indicating whether a mapping model should be created automatically. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640623-shouldinfermappingmodelautomatic?language=objc
func (p_ PersistentStoreDescription) SetShouldInferMappingModelAutomatically(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setShouldInferMappingModelAutomatically:"), value)
}

// The name of the configuration used by this store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640634-configuration?language=objc
func (p_ PersistentStoreDescription) Configuration() string {
	rv := objc.Call[string](p_, objc.Sel("configuration"))
	return rv
}

// The name of the configuration used by this store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640634-configuration?language=objc
func (p_ PersistentStoreDescription) SetConfiguration(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setConfiguration:"), value)
}

// A flag that determines whether the store is added asynchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640599-shouldaddstoreasynchronously?language=objc
func (p_ PersistentStoreDescription) ShouldAddStoreAsynchronously() bool {
	rv := objc.Call[bool](p_, objc.Sel("shouldAddStoreAsynchronously"))
	return rv
}

// A flag that determines whether the store is added asynchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640599-shouldaddstoreasynchronously?language=objc
func (p_ PersistentStoreDescription) SetShouldAddStoreAsynchronously(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setShouldAddStoreAsynchronously:"), value)
}

// The SQLite pragmas set for the associated persistent store. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640614-sqlitepragmas?language=objc
func (p_ PersistentStoreDescription) SqlitePragmas() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](p_, objc.Sel("sqlitePragmas"))
	return rv
}

// Options that customize how this store description aligns with a CloudKit database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/3141672-cloudkitcontaineroptions?language=objc
func (p_ PersistentStoreDescription) CloudKitContainerOptions() PersistentCloudKitContainerOptions {
	rv := objc.Call[PersistentCloudKitContainerOptions](p_, objc.Sel("cloudKitContainerOptions"))
	return rv
}

// Options that customize how this store description aligns with a CloudKit database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/3141672-cloudkitcontaineroptions?language=objc
func (p_ PersistentStoreDescription) SetCloudKitContainerOptions(value IPersistentCloudKitContainerOptions) {
	objc.Call[objc.Void](p_, objc.Sel("setCloudKitContainerOptions:"), value)
}

// A flag indicating whether the associated persistent store should be migrated automatically. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640566-shouldmigratestoreautomatically?language=objc
func (p_ PersistentStoreDescription) ShouldMigrateStoreAutomatically() bool {
	rv := objc.Call[bool](p_, objc.Sel("shouldMigrateStoreAutomatically"))
	return rv
}

// A flag indicating whether the associated persistent store should be migrated automatically. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640566-shouldmigratestoreautomatically?language=objc
func (p_ PersistentStoreDescription) SetShouldMigrateStoreAutomatically(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setShouldMigrateStoreAutomatically:"), value)
}

// The type of store this description represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640609-type?language=objc
func (p_ PersistentStoreDescription) Type() string {
	rv := objc.Call[string](p_, objc.Sel("type"))
	return rv
}

// The type of store this description represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640609-type?language=objc
func (p_ PersistentStoreDescription) SetType(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setType:"), value)
}

// A dictionary representation of the options set on the associated persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640571-options?language=objc
func (p_ PersistentStoreDescription) Options() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](p_, objc.Sel("options"))
	return rv
}

// The connection timeout for the associated store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640587-timeout?language=objc
func (p_ PersistentStoreDescription) Timeout() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](p_, objc.Sel("timeout"))
	return rv
}

// The connection timeout for the associated store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640587-timeout?language=objc
func (p_ PersistentStoreDescription) SetTimeout(value foundation.TimeInterval) {
	objc.Call[objc.Void](p_, objc.Sel("setTimeout:"), value)
}

// A flag that indicates whether this store will be read-only. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640626-readonly?language=objc
func (p_ PersistentStoreDescription) IsReadOnly() bool {
	rv := objc.Call[bool](p_, objc.Sel("isReadOnly"))
	return rv
}

// A flag that indicates whether this store will be read-only. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/1640626-readonly?language=objc
func (p_ PersistentStoreDescription) SetReadOnly(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setReadOnly:"), value)
}
