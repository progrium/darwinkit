// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchRecordZonesOperation] class.
var FetchRecordZonesOperationClass = _FetchRecordZonesOperationClass{objc.GetClass("CKFetchRecordZonesOperation")}

type _FetchRecordZonesOperationClass struct {
	objc.Class
}

// An interface definition for the [FetchRecordZonesOperation] class.
type IFetchRecordZonesOperation interface {
	IDatabaseOperation
	FetchRecordZonesCompletionBlock() func(recordZonesByZoneID foundation.Dictionary, operationError foundation.Error)
	SetFetchRecordZonesCompletionBlock(value func(recordZonesByZoneID foundation.Dictionary, operationError foundation.Error))
	RecordZoneIDs() []RecordZoneID
	SetRecordZoneIDs(value []IRecordZoneID)
	PerRecordZoneCompletionBlock() func(zoneID RecordZoneID, recordZone RecordZone, error foundation.Error)
	SetPerRecordZoneCompletionBlock(value func(zoneID RecordZoneID, recordZone RecordZone, error foundation.Error))
}

// An operation for retrieving record zones from a database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonesoperation?language=objc
type FetchRecordZonesOperation struct {
	DatabaseOperation
}

func FetchRecordZonesOperationFrom(ptr unsafe.Pointer) FetchRecordZonesOperation {
	return FetchRecordZonesOperation{
		DatabaseOperation: DatabaseOperationFrom(ptr),
	}
}

func (f_ FetchRecordZonesOperation) InitWithRecordZoneIDs(zoneIDs []IRecordZoneID) FetchRecordZonesOperation {
	rv := objc.Call[FetchRecordZonesOperation](f_, objc.Sel("initWithRecordZoneIDs:"), zoneIDs)
	return rv
}

// Creates an operation for fetching the specified record zones. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonesoperation/1515299-initwithrecordzoneids?language=objc
func FetchRecordZonesOperation_InitWithRecordZoneIDs(zoneIDs []IRecordZoneID) FetchRecordZonesOperation {
	return FetchRecordZonesOperationClass.Alloc().InitWithRecordZoneIDs(zoneIDs)
}

func (fc _FetchRecordZonesOperationClass) FetchAllRecordZonesOperation() FetchRecordZonesOperation {
	rv := objc.Call[FetchRecordZonesOperation](fc, objc.Sel("fetchAllRecordZonesOperation"))
	return rv
}

// Returns an operation for fetching all record zones in the current database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonesoperation/1514890-fetchallrecordzonesoperation?language=objc
func FetchRecordZonesOperation_FetchAllRecordZonesOperation() FetchRecordZonesOperation {
	return FetchRecordZonesOperationClass.FetchAllRecordZonesOperation()
}

func (f_ FetchRecordZonesOperation) Init() FetchRecordZonesOperation {
	rv := objc.Call[FetchRecordZonesOperation](f_, objc.Sel("init"))
	return rv
}

func (fc _FetchRecordZonesOperationClass) Alloc() FetchRecordZonesOperation {
	rv := objc.Call[FetchRecordZonesOperation](fc, objc.Sel("alloc"))
	return rv
}

func FetchRecordZonesOperation_Alloc() FetchRecordZonesOperation {
	return FetchRecordZonesOperationClass.Alloc()
}

func (fc _FetchRecordZonesOperationClass) New() FetchRecordZonesOperation {
	rv := objc.Call[FetchRecordZonesOperation](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchRecordZonesOperation() FetchRecordZonesOperation {
	return FetchRecordZonesOperationClass.New()
}

// The block to execute after CloudKit retrieves all of the record zones. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonesoperation/1515145-fetchrecordzonescompletionblock?language=objc
func (f_ FetchRecordZonesOperation) FetchRecordZonesCompletionBlock() func(recordZonesByZoneID foundation.Dictionary, operationError foundation.Error) {
	rv := objc.Call[func(recordZonesByZoneID foundation.Dictionary, operationError foundation.Error)](f_, objc.Sel("fetchRecordZonesCompletionBlock"))
	return rv
}

// The block to execute after CloudKit retrieves all of the record zones. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonesoperation/1515145-fetchrecordzonescompletionblock?language=objc
func (f_ FetchRecordZonesOperation) SetFetchRecordZonesCompletionBlock(value func(recordZonesByZoneID foundation.Dictionary, operationError foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("setFetchRecordZonesCompletionBlock:"), value)
}

// The IDs of the record zones to retrieve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonesoperation/1515084-recordzoneids?language=objc
func (f_ FetchRecordZonesOperation) RecordZoneIDs() []RecordZoneID {
	rv := objc.Call[[]RecordZoneID](f_, objc.Sel("recordZoneIDs"))
	return rv
}

// The IDs of the record zones to retrieve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonesoperation/1515084-recordzoneids?language=objc
func (f_ FetchRecordZonesOperation) SetRecordZoneIDs(value []IRecordZoneID) {
	objc.Call[objc.Void](f_, objc.Sel("setRecordZoneIDs:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonesoperation/3793699-perrecordzonecompletionblock?language=objc
func (f_ FetchRecordZonesOperation) PerRecordZoneCompletionBlock() func(zoneID RecordZoneID, recordZone RecordZone, error foundation.Error) {
	rv := objc.Call[func(zoneID RecordZoneID, recordZone RecordZone, error foundation.Error)](f_, objc.Sel("perRecordZoneCompletionBlock"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonesoperation/3793699-perrecordzonecompletionblock?language=objc
func (f_ FetchRecordZonesOperation) SetPerRecordZoneCompletionBlock(value func(zoneID RecordZoneID, recordZone RecordZone, error foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("setPerRecordZoneCompletionBlock:"), value)
}
