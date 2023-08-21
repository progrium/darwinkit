// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchRecordsOperation] class.
var FetchRecordsOperationClass = _FetchRecordsOperationClass{objc.GetClass("CKFetchRecordsOperation")}

type _FetchRecordsOperationClass struct {
	objc.Class
}

// An interface definition for the [FetchRecordsOperation] class.
type IFetchRecordsOperation interface {
	IDatabaseOperation
	FetchRecordsCompletionBlock() func(recordsByRecordID foundation.Dictionary, operationError foundation.Error)
	SetFetchRecordsCompletionBlock(value func(recordsByRecordID foundation.Dictionary, operationError foundation.Error))
	RecordIDs() []RecordID
	SetRecordIDs(value []IRecordID)
	PerRecordCompletionBlock() func(record Record, recordID RecordID, error foundation.Error)
	SetPerRecordCompletionBlock(value func(record Record, recordID RecordID, error foundation.Error))
	PerRecordProgressBlock() func(recordID RecordID, progress float64)
	SetPerRecordProgressBlock(value func(recordID RecordID, progress float64))
	DesiredKeys() []RecordFieldKey
	SetDesiredKeys(value []RecordFieldKey)
}

// An operation for retrieving records from a database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation?language=objc
type FetchRecordsOperation struct {
	DatabaseOperation
}

func FetchRecordsOperationFrom(ptr unsafe.Pointer) FetchRecordsOperation {
	return FetchRecordsOperation{
		DatabaseOperation: DatabaseOperationFrom(ptr),
	}
}

func (f_ FetchRecordsOperation) InitWithRecordIDs(recordIDs []IRecordID) FetchRecordsOperation {
	rv := objc.Call[FetchRecordsOperation](f_, objc.Sel("initWithRecordIDs:"), recordIDs)
	return rv
}

// Creates a fetch operation for retrieving the records with the specified IDs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/1476074-initwithrecordids?language=objc
func NewFetchRecordsOperationWithRecordIDs(recordIDs []IRecordID) FetchRecordsOperation {
	instance := FetchRecordsOperationClass.Alloc().InitWithRecordIDs(recordIDs)
	instance.Autorelease()
	return instance
}

func (f_ FetchRecordsOperation) Init() FetchRecordsOperation {
	rv := objc.Call[FetchRecordsOperation](f_, objc.Sel("init"))
	return rv
}

func (fc _FetchRecordsOperationClass) FetchCurrentUserRecordOperation() FetchRecordsOperation {
	rv := objc.Call[FetchRecordsOperation](fc, objc.Sel("fetchCurrentUserRecordOperation"))
	return rv
}

// Returns a fetch operation for retrieving the current user record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/1476070-fetchcurrentuserrecordoperation?language=objc
func FetchRecordsOperation_FetchCurrentUserRecordOperation() FetchRecordsOperation {
	return FetchRecordsOperationClass.FetchCurrentUserRecordOperation()
}

func (fc _FetchRecordsOperationClass) Alloc() FetchRecordsOperation {
	rv := objc.Call[FetchRecordsOperation](fc, objc.Sel("alloc"))
	return rv
}

func FetchRecordsOperation_Alloc() FetchRecordsOperation {
	return FetchRecordsOperationClass.Alloc()
}

func (fc _FetchRecordsOperationClass) New() FetchRecordsOperation {
	rv := objc.Call[FetchRecordsOperation](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchRecordsOperation() FetchRecordsOperation {
	return FetchRecordsOperationClass.New()
}

// The block to execute after CloudKit retrieves all of the records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/1476078-fetchrecordscompletionblock?language=objc
func (f_ FetchRecordsOperation) FetchRecordsCompletionBlock() func(recordsByRecordID foundation.Dictionary, operationError foundation.Error) {
	rv := objc.Call[func(recordsByRecordID foundation.Dictionary, operationError foundation.Error)](f_, objc.Sel("fetchRecordsCompletionBlock"))
	return rv
}

// The block to execute after CloudKit retrieves all of the records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/1476078-fetchrecordscompletionblock?language=objc
func (f_ FetchRecordsOperation) SetFetchRecordsCompletionBlock(value func(recordsByRecordID foundation.Dictionary, operationError foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("setFetchRecordsCompletionBlock:"), value)
}

// The record IDs of the records to fetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/1476076-recordids?language=objc
func (f_ FetchRecordsOperation) RecordIDs() []RecordID {
	rv := objc.Call[[]RecordID](f_, objc.Sel("recordIDs"))
	return rv
}

// The record IDs of the records to fetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/1476076-recordids?language=objc
func (f_ FetchRecordsOperation) SetRecordIDs(value []IRecordID) {
	objc.Call[objc.Void](f_, objc.Sel("setRecordIDs:"), value)
}

// The block to execute when a record becomes available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/1476082-perrecordcompletionblock?language=objc
func (f_ FetchRecordsOperation) PerRecordCompletionBlock() func(record Record, recordID RecordID, error foundation.Error) {
	rv := objc.Call[func(record Record, recordID RecordID, error foundation.Error)](f_, objc.Sel("perRecordCompletionBlock"))
	return rv
}

// The block to execute when a record becomes available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/1476082-perrecordcompletionblock?language=objc
func (f_ FetchRecordsOperation) SetPerRecordCompletionBlock(value func(record Record, recordID RecordID, error foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("setPerRecordCompletionBlock:"), value)
}

// The block to execute with progress information for individual records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/1476080-perrecordprogressblock?language=objc
func (f_ FetchRecordsOperation) PerRecordProgressBlock() func(recordID RecordID, progress float64) {
	rv := objc.Call[func(recordID RecordID, progress float64)](f_, objc.Sel("perRecordProgressBlock"))
	return rv
}

// The block to execute with progress information for individual records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/1476080-perrecordprogressblock?language=objc
func (f_ FetchRecordsOperation) SetPerRecordProgressBlock(value func(recordID RecordID, progress float64)) {
	objc.Call[objc.Void](f_, objc.Sel("setPerRecordProgressBlock:"), value)
}

// The fields of the records to fetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/1476088-desiredkeys?language=objc
func (f_ FetchRecordsOperation) DesiredKeys() []RecordFieldKey {
	rv := objc.Call[[]RecordFieldKey](f_, objc.Sel("desiredKeys"))
	return rv
}

// The fields of the records to fetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/1476088-desiredkeys?language=objc
func (f_ FetchRecordsOperation) SetDesiredKeys(value []RecordFieldKey) {
	objc.Call[objc.Void](f_, objc.Sel("setDesiredKeys:"), value)
}
