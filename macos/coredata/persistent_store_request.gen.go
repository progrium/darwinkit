// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PersistentStoreRequest] class.
var PersistentStoreRequestClass = _PersistentStoreRequestClass{objc.GetClass("NSPersistentStoreRequest")}

type _PersistentStoreRequestClass struct {
	objc.Class
}

// An interface definition for the [PersistentStoreRequest] class.
type IPersistentStoreRequest interface {
	objc.IObject
	RequestType() PersistentStoreRequestType
	AffectedStores() []PersistentStore
	SetAffectedStores(value []IPersistentStore)
}

// Criteria used to retrieve data from or save data to a persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorerequest?language=objc
type PersistentStoreRequest struct {
	objc.Object
}

func PersistentStoreRequestFrom(ptr unsafe.Pointer) PersistentStoreRequest {
	return PersistentStoreRequest{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PersistentStoreRequestClass) Alloc() PersistentStoreRequest {
	rv := objc.Call[PersistentStoreRequest](pc, objc.Sel("alloc"))
	return rv
}

func PersistentStoreRequest_Alloc() PersistentStoreRequest {
	return PersistentStoreRequestClass.Alloc()
}

func (pc _PersistentStoreRequestClass) New() PersistentStoreRequest {
	rv := objc.Call[PersistentStoreRequest](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentStoreRequest() PersistentStoreRequest {
	return PersistentStoreRequestClass.New()
}

func (p_ PersistentStoreRequest) Init() PersistentStoreRequest {
	rv := objc.Call[PersistentStoreRequest](p_, objc.Sel("init"))
	return rv
}

// The type of the fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorerequest/1506892-requesttype?language=objc
func (p_ PersistentStoreRequest) RequestType() PersistentStoreRequestType {
	rv := objc.Call[PersistentStoreRequestType](p_, objc.Sel("requestType"))
	return rv
}

// The stores the request should be sent to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorerequest/1506844-affectedstores?language=objc
func (p_ PersistentStoreRequest) AffectedStores() []PersistentStore {
	rv := objc.Call[[]PersistentStore](p_, objc.Sel("affectedStores"))
	return rv
}

// The stores the request should be sent to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorerequest/1506844-affectedstores?language=objc
func (p_ PersistentStoreRequest) SetAffectedStores(value []IPersistentStore) {
	objc.Call[objc.Void](p_, objc.Sel("setAffectedStores:"), value)
}
