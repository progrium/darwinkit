// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BatchDeleteRequest] class.
var BatchDeleteRequestClass = _BatchDeleteRequestClass{objc.GetClass("NSBatchDeleteRequest")}

type _BatchDeleteRequestClass struct {
	objc.Class
}

// An interface definition for the [BatchDeleteRequest] class.
type IBatchDeleteRequest interface {
	IPersistentStoreRequest
	FetchRequest() FetchRequest
	ResultType() BatchDeleteRequestResultType
	SetResultType(value BatchDeleteRequestResultType)
}

// A request that deletes objects in the SQLite persistent store without loading them into memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchdeleterequest?language=objc
type BatchDeleteRequest struct {
	PersistentStoreRequest
}

func BatchDeleteRequestFrom(ptr unsafe.Pointer) BatchDeleteRequest {
	return BatchDeleteRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}

func (b_ BatchDeleteRequest) InitWithFetchRequest(fetch IFetchRequest) BatchDeleteRequest {
	rv := objc.Call[BatchDeleteRequest](b_, objc.Sel("initWithFetchRequest:"), objc.Ptr(fetch))
	return rv
}

// Creates a request that deletes the results of the specified fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchdeleterequest/1506302-initwithfetchrequest?language=objc
func BatchDeleteRequest_InitWithFetchRequest(fetch IFetchRequest) BatchDeleteRequest {
	return BatchDeleteRequestClass.Alloc().InitWithFetchRequest(fetch)
}

func (b_ BatchDeleteRequest) InitWithObjectIDs(objects []IManagedObjectID) BatchDeleteRequest {
	rv := objc.Call[BatchDeleteRequest](b_, objc.Sel("initWithObjectIDs:"), objects)
	return rv
}

// Creates a request that deletes the managed objects with the specified identifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchdeleterequest/1506746-initwithobjectids?language=objc
func BatchDeleteRequest_InitWithObjectIDs(objects []IManagedObjectID) BatchDeleteRequest {
	return BatchDeleteRequestClass.Alloc().InitWithObjectIDs(objects)
}

func (bc _BatchDeleteRequestClass) Alloc() BatchDeleteRequest {
	rv := objc.Call[BatchDeleteRequest](bc, objc.Sel("alloc"))
	return rv
}

func BatchDeleteRequest_Alloc() BatchDeleteRequest {
	return BatchDeleteRequestClass.Alloc()
}

func (bc _BatchDeleteRequestClass) New() BatchDeleteRequest {
	rv := objc.Call[BatchDeleteRequest](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBatchDeleteRequest() BatchDeleteRequest {
	return BatchDeleteRequestClass.New()
}

func (b_ BatchDeleteRequest) Init() BatchDeleteRequest {
	rv := objc.Call[BatchDeleteRequest](b_, objc.Sel("init"))
	return rv
}

// The fetch request that identifies the managed objects to delete. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchdeleterequest/1506206-fetchrequest?language=objc
func (b_ BatchDeleteRequest) FetchRequest() FetchRequest {
	rv := objc.Call[FetchRequest](b_, objc.Sel("fetchRequest"))
	return rv
}

// The type of result the request provides when it executes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchdeleterequest/1506389-resulttype?language=objc
func (b_ BatchDeleteRequest) ResultType() BatchDeleteRequestResultType {
	rv := objc.Call[BatchDeleteRequestResultType](b_, objc.Sel("resultType"))
	return rv
}

// The type of result the request provides when it executes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchdeleterequest/1506389-resulttype?language=objc
func (b_ BatchDeleteRequest) SetResultType(value BatchDeleteRequestResultType) {
	objc.Call[objc.Void](b_, objc.Sel("setResultType:"), value)
}
