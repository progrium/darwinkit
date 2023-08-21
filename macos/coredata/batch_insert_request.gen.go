// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BatchInsertRequest] class.
var BatchInsertRequestClass = _BatchInsertRequestClass{objc.GetClass("NSBatchInsertRequest")}

type _BatchInsertRequestClass struct {
	objc.Class
}

// An interface definition for the [BatchInsertRequest] class.
type IBatchInsertRequest interface {
	IPersistentStoreRequest
	EntityName() string
	Entity() EntityDescription
	ObjectsToInsert() []map[string]objc.Object
	SetObjectsToInsert(value []map[string]objc.IObject)
	DictionaryHandler() func(obj foundation.MutableDictionary) bool
	SetDictionaryHandler(value func(obj foundation.MutableDictionary) bool)
	ManagedObjectHandler() func(obj ManagedObject) bool
	SetManagedObjectHandler(value func(obj ManagedObject) bool)
	ResultType() BatchInsertRequestResultType
	SetResultType(value BatchInsertRequestResultType)
}

// A request to insert a batch of data in a persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest?language=objc
type BatchInsertRequest struct {
	PersistentStoreRequest
}

func BatchInsertRequestFrom(ptr unsafe.Pointer) BatchInsertRequest {
	return BatchInsertRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}

func (b_ BatchInsertRequest) InitWithEntityNameObjects(entityName string, dictionaries []map[string]objc.IObject) BatchInsertRequest {
	rv := objc.Call[BatchInsertRequest](b_, objc.Sel("initWithEntityName:objects:"), entityName, dictionaries)
	return rv
}

// Creates a batch-insertion request for a named managed entity, and provides an array of data dictionaries for insertion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/3333237-initwithentityname?language=objc
func NewBatchInsertRequestWithEntityNameObjects(entityName string, dictionaries []map[string]objc.IObject) BatchInsertRequest {
	instance := BatchInsertRequestClass.Alloc().InitWithEntityNameObjects(entityName, dictionaries)
	instance.Autorelease()
	return instance
}

func (bc _BatchInsertRequestClass) BatchInsertRequestWithEntityNameObjects(entityName string, dictionaries []map[string]objc.IObject) BatchInsertRequest {
	rv := objc.Call[BatchInsertRequest](bc, objc.Sel("batchInsertRequestWithEntityName:objects:"), entityName, dictionaries)
	return rv
}

// Creates a batch-insertion request for a named managed entity, and provides an array of data dictionaries for insertion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/3333230-batchinsertrequestwithentityname?language=objc
func BatchInsertRequest_BatchInsertRequestWithEntityNameObjects(entityName string, dictionaries []map[string]objc.IObject) BatchInsertRequest {
	return BatchInsertRequestClass.BatchInsertRequestWithEntityNameObjects(entityName, dictionaries)
}

func (b_ BatchInsertRequest) InitWithEntityManagedObjectHandler(entity IEntityDescription, handler func(obj ManagedObject) bool) BatchInsertRequest {
	rv := objc.Call[BatchInsertRequest](b_, objc.Sel("initWithEntity:managedObjectHandler:"), objc.Ptr(entity), handler)
	return rv
}

// Creates a batch-insertion request for a managed entity, and specifies a closure that inserts data into the entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/3618791-initwithentity?language=objc
func NewBatchInsertRequestWithEntityManagedObjectHandler(entity IEntityDescription, handler func(obj ManagedObject) bool) BatchInsertRequest {
	instance := BatchInsertRequestClass.Alloc().InitWithEntityManagedObjectHandler(entity, handler)
	instance.Autorelease()
	return instance
}

func (bc _BatchInsertRequestClass) Alloc() BatchInsertRequest {
	rv := objc.Call[BatchInsertRequest](bc, objc.Sel("alloc"))
	return rv
}

func BatchInsertRequest_Alloc() BatchInsertRequest {
	return BatchInsertRequestClass.Alloc()
}

func (bc _BatchInsertRequestClass) New() BatchInsertRequest {
	rv := objc.Call[BatchInsertRequest](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBatchInsertRequest() BatchInsertRequest {
	return BatchInsertRequestClass.New()
}

func (b_ BatchInsertRequest) Init() BatchInsertRequest {
	rv := objc.Call[BatchInsertRequest](b_, objc.Sel("init"))
	return rv
}

// The name of the managed entity to insert data into. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/3333232-entityname?language=objc
func (b_ BatchInsertRequest) EntityName() string {
	rv := objc.Call[string](b_, objc.Sel("entityName"))
	return rv
}

// The managed entity to insert data into. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/3333231-entity?language=objc
func (b_ BatchInsertRequest) Entity() EntityDescription {
	rv := objc.Call[EntityDescription](b_, objc.Sel("entity"))
	return rv
}

// An array of dictionaries that represents the objects to insert with the keys as attribute names and their assigned values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/3333238-objectstoinsert?language=objc
func (b_ BatchInsertRequest) ObjectsToInsert() []map[string]objc.Object {
	rv := objc.Call[[]map[string]objc.Object](b_, objc.Sel("objectsToInsert"))
	return rv
}

// An array of dictionaries that represents the objects to insert with the keys as attribute names and their assigned values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/3333238-objectstoinsert?language=objc
func (b_ BatchInsertRequest) SetObjectsToInsert(value []map[string]objc.IObject) {
	objc.Call[objc.Void](b_, objc.Sel("setObjectsToInsert:"), value)
}

// A closure that provides a dictionary for your app to insert data into. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/3618789-dictionaryhandler?language=objc
func (b_ BatchInsertRequest) DictionaryHandler() func(obj foundation.MutableDictionary) bool {
	rv := objc.Call[func(obj foundation.MutableDictionary) bool](b_, objc.Sel("dictionaryHandler"))
	return rv
}

// A closure that provides a dictionary for your app to insert data into. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/3618789-dictionaryhandler?language=objc
func (b_ BatchInsertRequest) SetDictionaryHandler(value func(obj foundation.MutableDictionary) bool) {
	objc.Call[objc.Void](b_, objc.Sel("setDictionaryHandler:"), value)
}

// A closure that provides a managed object for your app to insert data into. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/3618792-managedobjecthandler?language=objc
func (b_ BatchInsertRequest) ManagedObjectHandler() func(obj ManagedObject) bool {
	rv := objc.Call[func(obj ManagedObject) bool](b_, objc.Sel("managedObjectHandler"))
	return rv
}

// A closure that provides a managed object for your app to insert data into. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/3618792-managedobjecthandler?language=objc
func (b_ BatchInsertRequest) SetManagedObjectHandler(value func(obj ManagedObject) bool) {
	objc.Call[objc.Void](b_, objc.Sel("setManagedObjectHandler:"), value)
}

// The type of result that Core Data returns from this request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/3333239-resulttype?language=objc
func (b_ BatchInsertRequest) ResultType() BatchInsertRequestResultType {
	rv := objc.Call[BatchInsertRequestResultType](b_, objc.Sel("resultType"))
	return rv
}

// The type of result that Core Data returns from this request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/3333239-resulttype?language=objc
func (b_ BatchInsertRequest) SetResultType(value BatchInsertRequestResultType) {
	objc.Call[objc.Void](b_, objc.Sel("setResultType:"), value)
}
