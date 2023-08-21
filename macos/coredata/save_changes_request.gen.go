// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SaveChangesRequest] class.
var SaveChangesRequestClass = _SaveChangesRequestClass{objc.GetClass("NSSaveChangesRequest")}

type _SaveChangesRequestClass struct {
	objc.Class
}

// An interface definition for the [SaveChangesRequest] class.
type ISaveChangesRequest interface {
	IPersistentStoreRequest
	DeletedObjects() foundation.Set
	InsertedObjects() foundation.Set
	UpdatedObjects() foundation.Set
	LockedObjects() foundation.Set
}

// An encapsulation of a collection of changes to be made by an object store in response to a save operation on a managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nssavechangesrequest?language=objc
type SaveChangesRequest struct {
	PersistentStoreRequest
}

func SaveChangesRequestFrom(ptr unsafe.Pointer) SaveChangesRequest {
	return SaveChangesRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}

func (s_ SaveChangesRequest) InitWithInsertedObjectsUpdatedObjectsDeletedObjectsLockedObjects(insertedObjects foundation.ISet, updatedObjects foundation.ISet, deletedObjects foundation.ISet, lockedObjects foundation.ISet) SaveChangesRequest {
	rv := objc.Call[SaveChangesRequest](s_, objc.Sel("initWithInsertedObjects:updatedObjects:deletedObjects:lockedObjects:"), objc.Ptr(insertedObjects), objc.Ptr(updatedObjects), objc.Ptr(deletedObjects), objc.Ptr(lockedObjects))
	return rv
}

// Initializes a save changes request with collections of given changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nssavechangesrequest/1500418-initwithinsertedobjects?language=objc
func NewSaveChangesRequestWithInsertedObjectsUpdatedObjectsDeletedObjectsLockedObjects(insertedObjects foundation.ISet, updatedObjects foundation.ISet, deletedObjects foundation.ISet, lockedObjects foundation.ISet) SaveChangesRequest {
	instance := SaveChangesRequestClass.Alloc().InitWithInsertedObjectsUpdatedObjectsDeletedObjectsLockedObjects(insertedObjects, updatedObjects, deletedObjects, lockedObjects)
	instance.Autorelease()
	return instance
}

func (sc _SaveChangesRequestClass) Alloc() SaveChangesRequest {
	rv := objc.Call[SaveChangesRequest](sc, objc.Sel("alloc"))
	return rv
}

func SaveChangesRequest_Alloc() SaveChangesRequest {
	return SaveChangesRequestClass.Alloc()
}

func (sc _SaveChangesRequestClass) New() SaveChangesRequest {
	rv := objc.Call[SaveChangesRequest](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSaveChangesRequest() SaveChangesRequest {
	return SaveChangesRequestClass.New()
}

func (s_ SaveChangesRequest) Init() SaveChangesRequest {
	rv := objc.Call[SaveChangesRequest](s_, objc.Sel("init"))
	return rv
}

// The objects that were deleted in the calling context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nssavechangesrequest/1500420-deletedobjects?language=objc
func (s_ SaveChangesRequest) DeletedObjects() foundation.Set {
	rv := objc.Call[foundation.Set](s_, objc.Sel("deletedObjects"))
	return rv
}

// The objects that were inserted into the calling context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nssavechangesrequest/1500416-insertedobjects?language=objc
func (s_ SaveChangesRequest) InsertedObjects() foundation.Set {
	rv := objc.Call[foundation.Set](s_, objc.Sel("insertedObjects"))
	return rv
}

// The objects that were modified in the calling context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nssavechangesrequest/1500424-updatedobjects?language=objc
func (s_ SaveChangesRequest) UpdatedObjects() foundation.Set {
	rv := objc.Call[foundation.Set](s_, objc.Sel("updatedObjects"))
	return rv
}

// The objects that were flagged for optimistic locking on the calling context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nssavechangesrequest/1500426-lockedobjects?language=objc
func (s_ SaveChangesRequest) LockedObjects() foundation.Set {
	rv := objc.Call[foundation.Set](s_, objc.Sel("lockedObjects"))
	return rv
}
