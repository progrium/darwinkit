// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ModifyRecordsOperation] class.
var ModifyRecordsOperationClass = _ModifyRecordsOperationClass{objc.GetClass("CKModifyRecordsOperation")}

type _ModifyRecordsOperationClass struct {
	objc.Class
}

// An interface definition for the [ModifyRecordsOperation] class.
type IModifyRecordsOperation interface {
	IDatabaseOperation
	RecordsToSave() []Record
	SetRecordsToSave(value []IRecord)
	ModifyRecordsCompletionBlock() func(savedRecords []Record, deletedRecordIDs []RecordID, operationError foundation.Error)
	SetModifyRecordsCompletionBlock(value func(savedRecords []Record, deletedRecordIDs []RecordID, operationError foundation.Error))
	Atomic() bool
	SetAtomic(value bool)
	PerRecordSaveBlock() func(recordID RecordID, record Record, error foundation.Error)
	SetPerRecordSaveBlock(value func(recordID RecordID, record Record, error foundation.Error))
	SavePolicy() RecordSavePolicy
	SetSavePolicy(value RecordSavePolicy)
	ClientChangeTokenData() []byte
	SetClientChangeTokenData(value []byte)
	PerRecordDeleteBlock() func(recordID RecordID, error foundation.Error)
	SetPerRecordDeleteBlock(value func(recordID RecordID, error foundation.Error))
	PerRecordProgressBlock() func(record Record, progress float64)
	SetPerRecordProgressBlock(value func(record Record, progress float64))
	RecordIDsToDelete() []RecordID
	SetRecordIDsToDelete(value []IRecordID)
}

// An operation that modifies one or more records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation?language=objc
type ModifyRecordsOperation struct {
	DatabaseOperation
}

func ModifyRecordsOperationFrom(ptr unsafe.Pointer) ModifyRecordsOperation {
	return ModifyRecordsOperation{
		DatabaseOperation: DatabaseOperationFrom(ptr),
	}
}

func (m_ ModifyRecordsOperation) InitWithRecordsToSaveRecordIDsToDelete(records []IRecord, recordIDs []IRecordID) ModifyRecordsOperation {
	rv := objc.Call[ModifyRecordsOperation](m_, objc.Sel("initWithRecordsToSave:recordIDsToDelete:"), records, recordIDs)
	return rv
}

// Creates an operation for modifying the specified records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/1447464-initwithrecordstosave?language=objc
func ModifyRecordsOperation_InitWithRecordsToSaveRecordIDsToDelete(records []IRecord, recordIDs []IRecordID) ModifyRecordsOperation {
	return ModifyRecordsOperationClass.Alloc().InitWithRecordsToSaveRecordIDsToDelete(records, recordIDs)
}

func (m_ ModifyRecordsOperation) Init() ModifyRecordsOperation {
	rv := objc.Call[ModifyRecordsOperation](m_, objc.Sel("init"))
	return rv
}

func (mc _ModifyRecordsOperationClass) Alloc() ModifyRecordsOperation {
	rv := objc.Call[ModifyRecordsOperation](mc, objc.Sel("alloc"))
	return rv
}

func ModifyRecordsOperation_Alloc() ModifyRecordsOperation {
	return ModifyRecordsOperationClass.Alloc()
}

func (mc _ModifyRecordsOperationClass) New() ModifyRecordsOperation {
	rv := objc.Call[ModifyRecordsOperation](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewModifyRecordsOperation() ModifyRecordsOperation {
	return ModifyRecordsOperationClass.New()
}

// The records to save to the database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/1447482-recordstosave?language=objc
func (m_ ModifyRecordsOperation) RecordsToSave() []Record {
	rv := objc.Call[[]Record](m_, objc.Sel("recordsToSave"))
	return rv
}

// The records to save to the database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/1447482-recordstosave?language=objc
func (m_ ModifyRecordsOperation) SetRecordsToSave(value []IRecord) {
	objc.Call[objc.Void](m_, objc.Sel("setRecordsToSave:"), value)
}

// The block to execute after CloudKit modifies all of the records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/1447486-modifyrecordscompletionblock?language=objc
func (m_ ModifyRecordsOperation) ModifyRecordsCompletionBlock() func(savedRecords []Record, deletedRecordIDs []RecordID, operationError foundation.Error) {
	rv := objc.Call[func(savedRecords []Record, deletedRecordIDs []RecordID, operationError foundation.Error)](m_, objc.Sel("modifyRecordsCompletionBlock"))
	return rv
}

// The block to execute after CloudKit modifies all of the records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/1447486-modifyrecordscompletionblock?language=objc
func (m_ ModifyRecordsOperation) SetModifyRecordsCompletionBlock(value func(savedRecords []Record, deletedRecordIDs []RecordID, operationError foundation.Error)) {
	objc.Call[objc.Void](m_, objc.Sel("setModifyRecordsCompletionBlock:"), value)
}

// A Boolean value that indicates whether the entire operation fails when CloudKit can’t update one or more records in a record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/1447484-atomic?language=objc
func (m_ ModifyRecordsOperation) Atomic() bool {
	rv := objc.Call[bool](m_, objc.Sel("atomic"))
	return rv
}

// A Boolean value that indicates whether the entire operation fails when CloudKit can’t update one or more records in a record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/1447484-atomic?language=objc
func (m_ ModifyRecordsOperation) SetAtomic(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAtomic:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/3793705-perrecordsaveblock?language=objc
func (m_ ModifyRecordsOperation) PerRecordSaveBlock() func(recordID RecordID, record Record, error foundation.Error) {
	rv := objc.Call[func(recordID RecordID, record Record, error foundation.Error)](m_, objc.Sel("perRecordSaveBlock"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/3793705-perrecordsaveblock?language=objc
func (m_ ModifyRecordsOperation) SetPerRecordSaveBlock(value func(recordID RecordID, record Record, error foundation.Error)) {
	objc.Call[objc.Void](m_, objc.Sel("setPerRecordSaveBlock:"), value)
}

// The policy to use when saving changes to records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/1447488-savepolicy?language=objc
func (m_ ModifyRecordsOperation) SavePolicy() RecordSavePolicy {
	rv := objc.Call[RecordSavePolicy](m_, objc.Sel("savePolicy"))
	return rv
}

// The policy to use when saving changes to records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/1447488-savepolicy?language=objc
func (m_ ModifyRecordsOperation) SetSavePolicy(value RecordSavePolicy) {
	objc.Call[objc.Void](m_, objc.Sel("setSavePolicy:"), value)
}

// A token that tracks local changes to records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/1447472-clientchangetokendata?language=objc
func (m_ ModifyRecordsOperation) ClientChangeTokenData() []byte {
	rv := objc.Call[[]byte](m_, objc.Sel("clientChangeTokenData"))
	return rv
}

// A token that tracks local changes to records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/1447472-clientchangetokendata?language=objc
func (m_ ModifyRecordsOperation) SetClientChangeTokenData(value []byte) {
	objc.Call[objc.Void](m_, objc.Sel("setClientChangeTokenData:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/3794356-perrecorddeleteblock?language=objc
func (m_ ModifyRecordsOperation) PerRecordDeleteBlock() func(recordID RecordID, error foundation.Error) {
	rv := objc.Call[func(recordID RecordID, error foundation.Error)](m_, objc.Sel("perRecordDeleteBlock"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/3794356-perrecorddeleteblock?language=objc
func (m_ ModifyRecordsOperation) SetPerRecordDeleteBlock(value func(recordID RecordID, error foundation.Error)) {
	objc.Call[objc.Void](m_, objc.Sel("setPerRecordDeleteBlock:"), value)
}

// The block to execute with progress information for individual records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/1447477-perrecordprogressblock?language=objc
func (m_ ModifyRecordsOperation) PerRecordProgressBlock() func(record Record, progress float64) {
	rv := objc.Call[func(record Record, progress float64)](m_, objc.Sel("perRecordProgressBlock"))
	return rv
}

// The block to execute with progress information for individual records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/1447477-perrecordprogressblock?language=objc
func (m_ ModifyRecordsOperation) SetPerRecordProgressBlock(value func(record Record, progress float64)) {
	objc.Call[objc.Void](m_, objc.Sel("setPerRecordProgressBlock:"), value)
}

// The IDs of the records to delete permanently from the database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/1447479-recordidstodelete?language=objc
func (m_ ModifyRecordsOperation) RecordIDsToDelete() []RecordID {
	rv := objc.Call[[]RecordID](m_, objc.Sel("recordIDsToDelete"))
	return rv
}

// The IDs of the records to delete permanently from the database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/1447479-recordidstodelete?language=objc
func (m_ ModifyRecordsOperation) SetRecordIDsToDelete(value []IRecordID) {
	objc.Call[objc.Void](m_, objc.Sel("setRecordIDsToDelete:"), value)
}
