// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

// The types of attribute that Core Data supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributetype?language=objc
type AttributeType uint

const (
	BinaryDataAttributeType    AttributeType = 1000
	BooleanAttributeType       AttributeType = 800
	DateAttributeType          AttributeType = 900
	DecimalAttributeType       AttributeType = 400
	DoubleAttributeType        AttributeType = 500
	FloatAttributeType         AttributeType = 600
	Integer16AttributeType     AttributeType = 100
	Integer32AttributeType     AttributeType = 200
	Integer64AttributeType     AttributeType = 300
	ObjectIDAttributeType      AttributeType = 2000
	StringAttributeType        AttributeType = 700
	TransformableAttributeType AttributeType = 1800
	URIAttributeType           AttributeType = 1200
	UUIDAttributeType          AttributeType = 1100
	UndefinedAttributeType     AttributeType = 0
)

// The types of result a batch delete request can provide when it executes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchdeleterequestresulttype?language=objc
type BatchDeleteRequestResultType uint

const (
	BatchDeleteResultTypeCount      BatchDeleteRequestResultType = 2
	BatchDeleteResultTypeObjectIDs  BatchDeleteRequestResultType = 1
	BatchDeleteResultTypeStatusOnly BatchDeleteRequestResultType = 0
)

// Result types for a batch-insertion request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequestresulttype?language=objc
type BatchInsertRequestResultType uint

const (
	BatchInsertRequestResultTypeCount      BatchInsertRequestResultType = 2
	BatchInsertRequestResultTypeObjectIDs  BatchInsertRequestResultType = 1
	BatchInsertRequestResultTypeStatusOnly BatchInsertRequestResultType = 0
)

// Result types for a batch-update request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchupdaterequestresulttype?language=objc
type BatchUpdateRequestResultType uint

const (
	StatusOnlyResultType          BatchUpdateRequestResultType = 0
	UpdatedObjectIDsResultType    BatchUpdateRequestResultType = 1
	UpdatedObjectsCountResultType BatchUpdateRequestResultType = 2
)

// Constants that determine what happens when you delete a relationship’s owning managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsdeleterule?language=objc
type DeleteRule uint

const (
	CascadeDeleteRule  DeleteRule = 2
	DenyDeleteRule     DeleteRule = 3
	NoActionDeleteRule DeleteRule = 0
	NullifyDeleteRule  DeleteRule = 1
)

// The types for mapping an entity between a source model and a destination model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymappingtype?language=objc
type EntityMappingType uint

const (
	AddEntityMappingType       EntityMappingType = 2
	CopyEntityMappingType      EntityMappingType = 4
	CustomEntityMappingType    EntityMappingType = 1
	RemoveEntityMappingType    EntityMappingType = 3
	TransformEntityMappingType EntityMappingType = 5
	UndefinedEntityMappingType EntityMappingType = 0
)

// Defines the possible types of index elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexelementtype?language=objc
type FetchIndexElementType uint

const (
	FetchIndexElementTypeBinary FetchIndexElementType = 0
	FetchIndexElementTypeRTree  FetchIndexElementType = 1
)

// Constants that specify the possible result types a fetch request can return. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequestresulttype?language=objc
type FetchRequestResultType uint

const (
	CountResultType           FetchRequestResultType = 4
	DictionaryResultType      FetchRequestResultType = 2
	ManagedObjectIDResultType FetchRequestResultType = 1
	ManagedObjectResultType   FetchRequestResultType = 0
)

// Constants that specify the possible types of changes that are reported. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultschangetype?language=objc
type FetchedResultsChangeType uint

const (
	FetchedResultsChangeDelete FetchedResultsChangeType = 2
	FetchedResultsChangeInsert FetchedResultsChangeType = 1
	FetchedResultsChangeMove   FetchedResultsChangeType = 3
	FetchedResultsChangeUpdate FetchedResultsChangeType = 4
)

// The concurrency types you can use with a managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontextconcurrencytype?language=objc
type ManagedObjectContextConcurrencyType uint

const (
	ConfinementConcurrencyType  ManagedObjectContextConcurrencyType = 0
	MainQueueConcurrencyType    ManagedObjectContextConcurrencyType = 2
	PrivateQueueConcurrencyType ManagedObjectContextConcurrencyType = 1
)

// Constants that define merge policy types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergepolicytype?language=objc
type MergePolicyType uint

const (
	ErrorMergePolicyType                      MergePolicyType = 0
	MergeByPropertyObjectTrumpMergePolicyType MergePolicyType = 2
	MergeByPropertyStoreTrumpMergePolicyType  MergePolicyType = 1
	OverwriteMergePolicyType                  MergePolicyType = 3
	RollbackMergePolicyType                   MergePolicyType = 4
)

// The types of results from a persistent CloudKit container event fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainereventresulttype?language=objc
type PersistentCloudKitContainerEventResultType int

const (
	PersistentCloudKitContainerEventResultTypeCountEvents PersistentCloudKitContainerEventResultType = 1
	PersistentCloudKitContainerEventResultTypeEvents      PersistentCloudKitContainerEventResultType = 0
)

// The type of event in a persistent CloudKit container, either setup, import, or export. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainereventtype?language=objc
type PersistentCloudKitContainerEventType int

const (
	PersistentCloudKitContainerEventTypeExport PersistentCloudKitContainerEventType = 2
	PersistentCloudKitContainerEventTypeImport PersistentCloudKitContainerEventType = 1
	PersistentCloudKitContainerEventTypeSetup  PersistentCloudKitContainerEventType = 0
)

// Options that control the behavior when promoting the container’s schema to CloudKit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainerschemainitializationoptions?language=objc
type PersistentCloudKitContainerSchemaInitializationOptions uint

const (
	PersistentCloudKitContainerSchemaInitializationOptionsDryRun      PersistentCloudKitContainerSchemaInitializationOptions = 2
	PersistentCloudKitContainerSchemaInitializationOptionsNone        PersistentCloudKitContainerSchemaInitializationOptions = 0
	PersistentCloudKitContainerSchemaInitializationOptionsPrintSchema PersistentCloudKitContainerSchemaInitializationOptions = 4
)

// The types of changes to managed objects reflected in persistent history. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychangetype?language=objc
type PersistentHistoryChangeType int

const (
	PersistentHistoryChangeTypeDelete PersistentHistoryChangeType = 2
	PersistentHistoryChangeTypeInsert PersistentHistoryChangeType = 0
	PersistentHistoryChangeTypeUpdate PersistentHistoryChangeType = 1
)

// The types of results from a persistent history change request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistoryresulttype?language=objc
type PersistentHistoryResultType int

const (
	PersistentHistoryResultTypeChangesOnly            PersistentHistoryResultType = 4
	PersistentHistoryResultTypeCount                  PersistentHistoryResultType = 2
	PersistentHistoryResultTypeObjectIDs              PersistentHistoryResultType = 1
	PersistentHistoryResultTypeStatusOnly             PersistentHistoryResultType = 0
	PersistentHistoryResultTypeTransactionsAndChanges PersistentHistoryResultType = 5
	PersistentHistoryResultTypeTransactionsOnly       PersistentHistoryResultType = 3
)

// Constants that specify the types of fetch requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorerequesttype?language=objc
type PersistentStoreRequestType uint

const (
	BatchDeleteRequestType PersistentStoreRequestType = 7
	BatchInsertRequestType PersistentStoreRequestType = 5
	BatchUpdateRequestType PersistentStoreRequestType = 6
	FetchRequestType       PersistentStoreRequestType = 1
	SaveRequestType        PersistentStoreRequestType = 2
)

// These constants are used as the value corresponding to the NSPersistentStoreUbiquitousTransitionTypeKey in the user info dictionary of NSPersistentStoreCoordinatorStoresWillChangeNotification and NSPersistentStoreCoordinatorStoresDidChangeNotification notifications to identify the type of event leading to a change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoreubiquitoustransitiontype?language=objc
type PersistentStoreUbiquitousTransitionType uint

const (
	PersistentStoreUbiquitousTransitionTypeAccountAdded           PersistentStoreUbiquitousTransitionType = 1
	PersistentStoreUbiquitousTransitionTypeAccountRemoved         PersistentStoreUbiquitousTransitionType = 2
	PersistentStoreUbiquitousTransitionTypeContentRemoved         PersistentStoreUbiquitousTransitionType = 3
	PersistentStoreUbiquitousTransitionTypeInitialImportCompleted PersistentStoreUbiquitousTransitionType = 4
)

// Constants that specify the reason the managed object may need to reinitialize its values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nssnapshoteventtype?language=objc
type SnapshotEventType uint

const (
	SnapshotEventMergePolicy   SnapshotEventType = 64
	SnapshotEventRefresh       SnapshotEventType = 32
	SnapshotEventRollback      SnapshotEventType = 16
	SnapshotEventUndoDeletion  SnapshotEventType = 4
	SnapshotEventUndoInsertion SnapshotEventType = 2
	SnapshotEventUndoUpdate    SnapshotEventType = 8
)
