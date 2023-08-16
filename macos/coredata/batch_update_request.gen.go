// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BatchUpdateRequest] class.
var BatchUpdateRequestClass = _BatchUpdateRequestClass{objc.GetClass("NSBatchUpdateRequest")}

type _BatchUpdateRequestClass struct {
	objc.Class
}

// An interface definition for the [BatchUpdateRequest] class.
type IBatchUpdateRequest interface {
	IPersistentStoreRequest
	EntityName() string
	IncludesSubentities() bool
	SetIncludesSubentities(value bool)
	Entity() EntityDescription
	Predicate() foundation.Predicate
	SetPredicate(value foundation.IPredicate)
	PropertiesToUpdate() foundation.Dictionary
	SetPropertiesToUpdate(value foundation.Dictionary)
	ResultType() BatchUpdateRequestResultType
	SetResultType(value BatchUpdateRequestResultType)
}

// A request to Core Data to do a batch update of data in a persistent store without loading any data into memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchupdaterequest?language=objc
type BatchUpdateRequest struct {
	PersistentStoreRequest
}

func BatchUpdateRequestFrom(ptr unsafe.Pointer) BatchUpdateRequest {
	return BatchUpdateRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}

func (bc _BatchUpdateRequestClass) BatchUpdateRequestWithEntityName(entityName string) BatchUpdateRequest {
	rv := objc.Call[BatchUpdateRequest](bc, objc.Sel("batchUpdateRequestWithEntityName:"), entityName)
	return rv
}

// Creates a batch-update request for a named managed entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchupdaterequest/1526273-batchupdaterequestwithentityname?language=objc
func BatchUpdateRequest_BatchUpdateRequestWithEntityName(entityName string) BatchUpdateRequest {
	return BatchUpdateRequestClass.BatchUpdateRequestWithEntityName(entityName)
}

func (b_ BatchUpdateRequest) InitWithEntityName(entityName string) BatchUpdateRequest {
	rv := objc.Call[BatchUpdateRequest](b_, objc.Sel("initWithEntityName:"), entityName)
	return rv
}

// Creates a batch-update request for a named managed entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchupdaterequest/1506702-initwithentityname?language=objc
func BatchUpdateRequest_InitWithEntityName(entityName string) BatchUpdateRequest {
	return BatchUpdateRequestClass.Alloc().InitWithEntityName(entityName)
}

func (b_ BatchUpdateRequest) InitWithEntity(entity IEntityDescription) BatchUpdateRequest {
	rv := objc.Call[BatchUpdateRequest](b_, objc.Sel("initWithEntity:"), objc.Ptr(entity))
	return rv
}

// Creates a batch-update request for a managed entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchupdaterequest/1506374-initwithentity?language=objc
func BatchUpdateRequest_InitWithEntity(entity IEntityDescription) BatchUpdateRequest {
	return BatchUpdateRequestClass.Alloc().InitWithEntity(entity)
}

func (bc _BatchUpdateRequestClass) Alloc() BatchUpdateRequest {
	rv := objc.Call[BatchUpdateRequest](bc, objc.Sel("alloc"))
	return rv
}

func BatchUpdateRequest_Alloc() BatchUpdateRequest {
	return BatchUpdateRequestClass.Alloc()
}

func (bc _BatchUpdateRequestClass) New() BatchUpdateRequest {
	rv := objc.Call[BatchUpdateRequest](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBatchUpdateRequest() BatchUpdateRequest {
	return BatchUpdateRequestClass.New()
}

func (b_ BatchUpdateRequest) Init() BatchUpdateRequest {
	rv := objc.Call[BatchUpdateRequest](b_, objc.Sel("init"))
	return rv
}

// The name of the managed entity to update data for. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchupdaterequest/1506796-entityname?language=objc
func (b_ BatchUpdateRequest) EntityName() string {
	rv := objc.Call[string](b_, objc.Sel("entityName"))
	return rv
}

// A Boolean value that indicates whether to update subentities. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchupdaterequest/1506477-includessubentities?language=objc
func (b_ BatchUpdateRequest) IncludesSubentities() bool {
	rv := objc.Call[bool](b_, objc.Sel("includesSubentities"))
	return rv
}

// A Boolean value that indicates whether to update subentities. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchupdaterequest/1506477-includessubentities?language=objc
func (b_ BatchUpdateRequest) SetIncludesSubentities(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setIncludesSubentities:"), value)
}

// The managed entity to update data for. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchupdaterequest/1506664-entity?language=objc
func (b_ BatchUpdateRequest) Entity() EntityDescription {
	rv := objc.Call[EntityDescription](b_, objc.Sel("entity"))
	return rv
}

// A predicate that identifies the objects to update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchupdaterequest/1506659-predicate?language=objc
func (b_ BatchUpdateRequest) Predicate() foundation.Predicate {
	rv := objc.Call[foundation.Predicate](b_, objc.Sel("predicate"))
	return rv
}

// A predicate that identifies the objects to update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchupdaterequest/1506659-predicate?language=objc
func (b_ BatchUpdateRequest) SetPredicate(value foundation.IPredicate) {
	objc.Call[objc.Void](b_, objc.Sel("setPredicate:"), objc.Ptr(value))
}

// A dictionary of property description pairs that describe the updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchupdaterequest/1506582-propertiestoupdate?language=objc
func (b_ BatchUpdateRequest) PropertiesToUpdate() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](b_, objc.Sel("propertiesToUpdate"))
	return rv
}

// A dictionary of property description pairs that describe the updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchupdaterequest/1506582-propertiestoupdate?language=objc
func (b_ BatchUpdateRequest) SetPropertiesToUpdate(value foundation.Dictionary) {
	objc.Call[objc.Void](b_, objc.Sel("setPropertiesToUpdate:"), value)
}

// The type of result that Core Data returns from the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchupdaterequest/1506350-resulttype?language=objc
func (b_ BatchUpdateRequest) ResultType() BatchUpdateRequestResultType {
	rv := objc.Call[BatchUpdateRequestResultType](b_, objc.Sel("resultType"))
	return rv
}

// The type of result that Core Data returns from the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchupdaterequest/1506350-resulttype?language=objc
func (b_ BatchUpdateRequest) SetResultType(value BatchUpdateRequestResultType) {
	objc.Call[objc.Void](b_, objc.Sel("setResultType:"), value)
}
