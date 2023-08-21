// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MigrationManager] class.
var MigrationManagerClass = _MigrationManagerClass{objc.GetClass("NSMigrationManager")}

type _MigrationManagerClass struct {
	objc.Class
}

// An interface definition for the [MigrationManager] class.
type IMigrationManager interface {
	objc.IObject
	DestinationEntityForEntityMapping(mEntity IEntityMapping) EntityDescription
	SourceInstancesForEntityMappingNamedDestinationInstances(mappingName string, destinationInstances []IManagedObject) []ManagedObject
	AssociateSourceInstanceWithDestinationInstanceForEntityMapping(sourceInstance IManagedObject, destinationInstance IManagedObject, entityMapping IEntityMapping)
	DestinationInstancesForEntityMappingNamedSourceInstances(mappingName string, sourceInstances []IManagedObject) []ManagedObject
	SourceEntityForEntityMapping(mEntity IEntityMapping) EntityDescription
	CancelMigrationWithError(error foundation.IError)
	Reset()
	SourceModel() ManagedObjectModel
	DestinationContext() ManagedObjectContext
	SourceContext() ManagedObjectContext
	DestinationModel() ManagedObjectModel
	UsesStoreSpecificMigrationManager() bool
	SetUsesStoreSpecificMigrationManager(value bool)
	MigrationProgress() float64
	UserInfo() foundation.Dictionary
	SetUserInfo(value foundation.Dictionary)
	CurrentEntityMapping() EntityMapping
	MappingModel() MappingModel
}

// A migration manager instance that performs a migration of data from one persistent store to another using a given mapping model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager?language=objc
type MigrationManager struct {
	objc.Object
}

func MigrationManagerFrom(ptr unsafe.Pointer) MigrationManager {
	return MigrationManager{
		Object: objc.ObjectFrom(ptr),
	}
}

func (m_ MigrationManager) InitWithSourceModelDestinationModel(sourceModel IManagedObjectModel, destinationModel IManagedObjectModel) MigrationManager {
	rv := objc.Call[MigrationManager](m_, objc.Sel("initWithSourceModel:destinationModel:"), objc.Ptr(sourceModel), objc.Ptr(destinationModel))
	return rv
}

// Initializes a migration manager instance with given source and destination models. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417583-initwithsourcemodel?language=objc
func NewMigrationManagerWithSourceModelDestinationModel(sourceModel IManagedObjectModel, destinationModel IManagedObjectModel) MigrationManager {
	instance := MigrationManagerClass.Alloc().InitWithSourceModelDestinationModel(sourceModel, destinationModel)
	instance.Autorelease()
	return instance
}

func (mc _MigrationManagerClass) Alloc() MigrationManager {
	rv := objc.Call[MigrationManager](mc, objc.Sel("alloc"))
	return rv
}

func MigrationManager_Alloc() MigrationManager {
	return MigrationManagerClass.Alloc()
}

func (mc _MigrationManagerClass) New() MigrationManager {
	rv := objc.Call[MigrationManager](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMigrationManager() MigrationManager {
	return MigrationManagerClass.New()
}

func (m_ MigrationManager) Init() MigrationManager {
	rv := objc.Call[MigrationManager](m_, objc.Sel("init"))
	return rv
}

// Returns the entity description for the destination entity of a given entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417598-destinationentityforentitymappin?language=objc
func (m_ MigrationManager) DestinationEntityForEntityMapping(mEntity IEntityMapping) EntityDescription {
	rv := objc.Call[EntityDescription](m_, objc.Sel("destinationEntityForEntityMapping:"), objc.Ptr(mEntity))
	return rv
}

// Returns the managed object instances in the source store used to create the given destination instances for the passed in property mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417580-sourceinstancesforentitymappingn?language=objc
func (m_ MigrationManager) SourceInstancesForEntityMappingNamedDestinationInstances(mappingName string, destinationInstances []IManagedObject) []ManagedObject {
	rv := objc.Call[[]ManagedObject](m_, objc.Sel("sourceInstancesForEntityMappingNamed:destinationInstances:"), mappingName, destinationInstances)
	return rv
}

// Associates a given source managed object instance with an array of destination instances for a given property mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417604-associatesourceinstance?language=objc
func (m_ MigrationManager) AssociateSourceInstanceWithDestinationInstanceForEntityMapping(sourceInstance IManagedObject, destinationInstance IManagedObject, entityMapping IEntityMapping) {
	objc.Call[objc.Void](m_, objc.Sel("associateSourceInstance:withDestinationInstance:forEntityMapping:"), objc.Ptr(sourceInstance), objc.Ptr(destinationInstance), objc.Ptr(entityMapping))
}

// Returns the managed object instances created in the destination store for the named entity mapping for the given array of source instances. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417594-destinationinstancesforentitymap?language=objc
func (m_ MigrationManager) DestinationInstancesForEntityMappingNamedSourceInstances(mappingName string, sourceInstances []IManagedObject) []ManagedObject {
	rv := objc.Call[[]ManagedObject](m_, objc.Sel("destinationInstancesForEntityMappingNamed:sourceInstances:"), mappingName, sourceInstances)
	return rv
}

// Returns the entity description for the source entity of a given entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417596-sourceentityforentitymapping?language=objc
func (m_ MigrationManager) SourceEntityForEntityMapping(mEntity IEntityMapping) EntityDescription {
	rv := objc.Call[EntityDescription](m_, objc.Sel("sourceEntityForEntityMapping:"), objc.Ptr(mEntity))
	return rv
}

// Cancels the migration with a given error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417608-cancelmigrationwitherror?language=objc
func (m_ MigrationManager) CancelMigrationWithError(error foundation.IError) {
	objc.Call[objc.Void](m_, objc.Sel("cancelMigrationWithError:"), objc.Ptr(error))
}

// Resets the association tables for the migration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417590-reset?language=objc
func (m_ MigrationManager) Reset() {
	objc.Call[objc.Void](m_, objc.Sel("reset"))
}

// The source model for the migration manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417592-sourcemodel?language=objc
func (m_ MigrationManager) SourceModel() ManagedObjectModel {
	rv := objc.Call[ManagedObjectModel](m_, objc.Sel("sourceModel"))
	return rv
}

// The managed object context the migration manager uses for writing the destination persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417578-destinationcontext?language=objc
func (m_ MigrationManager) DestinationContext() ManagedObjectContext {
	rv := objc.Call[ManagedObjectContext](m_, objc.Sel("destinationContext"))
	return rv
}

// The managed object context the migration manager uses for reading the source persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417576-sourcecontext?language=objc
func (m_ MigrationManager) SourceContext() ManagedObjectContext {
	rv := objc.Call[ManagedObjectContext](m_, objc.Sel("sourceContext"))
	return rv
}

// The destination model for the migration manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417610-destinationmodel?language=objc
func (m_ MigrationManager) DestinationModel() ManagedObjectModel {
	rv := objc.Call[ManagedObjectModel](m_, objc.Sel("destinationModel"))
	return rv
}

// A Boolean value that indicates whether the migration manager tries to use a store specific migration manager to perform the  migration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417606-usesstorespecificmigrationmanage?language=objc
func (m_ MigrationManager) UsesStoreSpecificMigrationManager() bool {
	rv := objc.Call[bool](m_, objc.Sel("usesStoreSpecificMigrationManager"))
	return rv
}

// A Boolean value that indicates whether the migration manager tries to use a store specific migration manager to perform the  migration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417606-usesstorespecificmigrationmanage?language=objc
func (m_ MigrationManager) SetUsesStoreSpecificMigrationManager(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setUsesStoreSpecificMigrationManager:"), value)
}

// A number between 0 and 1 that indicates the proportion of completeness of the migration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417602-migrationprogress?language=objc
func (m_ MigrationManager) MigrationProgress() float64 {
	rv := objc.Call[float64](m_, objc.Sel("migrationProgress"))
	return rv
}

// The user info for the migration manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417588-userinfo?language=objc
func (m_ MigrationManager) UserInfo() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](m_, objc.Sel("userInfo"))
	return rv
}

// The user info for the migration manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417588-userinfo?language=objc
func (m_ MigrationManager) SetUserInfo(value foundation.Dictionary) {
	objc.Call[objc.Void](m_, objc.Sel("setUserInfo:"), value)
}

// The entity mapping currently being processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417582-currententitymapping?language=objc
func (m_ MigrationManager) CurrentEntityMapping() EntityMapping {
	rv := objc.Call[EntityMapping](m_, objc.Sel("currentEntityMapping"))
	return rv
}

// The mapping model for the migration manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmigrationmanager/1417600-mappingmodel?language=objc
func (m_ MigrationManager) MappingModel() MappingModel {
	rv := objc.Call[MappingModel](m_, objc.Sel("mappingModel"))
	return rv
}
