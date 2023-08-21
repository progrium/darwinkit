// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchRecordZoneChangesOperation] class.
var FetchRecordZoneChangesOperationClass = _FetchRecordZoneChangesOperationClass{objc.GetClass("CKFetchRecordZoneChangesOperation")}

type _FetchRecordZoneChangesOperationClass struct {
	objc.Class
}

// An interface definition for the [FetchRecordZoneChangesOperation] class.
type IFetchRecordZoneChangesOperation interface {
	IDatabaseOperation
	RecordZoneIDs() []RecordZoneID
	SetRecordZoneIDs(value []IRecordZoneID)
	RecordWithIDWasDeletedBlock() func(recordID RecordID, recordType RecordType)
	SetRecordWithIDWasDeletedBlock(value func(recordID RecordID, recordType RecordType))
	RecordWasChangedBlock() func(recordID RecordID, record Record, error foundation.Error)
	SetRecordWasChangedBlock(value func(recordID RecordID, record Record, error foundation.Error))
	FetchRecordZoneChangesCompletionBlock() func(operationError foundation.Error)
	SetFetchRecordZoneChangesCompletionBlock(value func(operationError foundation.Error))
	ConfigurationsByRecordZoneID() foundation.Dictionary
	SetConfigurationsByRecordZoneID(value foundation.Dictionary)
	FetchAllChanges() bool
	SetFetchAllChanges(value bool)
	RecordZoneFetchCompletionBlock() func(recordZoneID RecordZoneID, serverChangeToken ServerChangeToken, clientChangeTokenData []byte, moreComing bool, recordZoneError foundation.Error)
	SetRecordZoneFetchCompletionBlock(value func(recordZoneID RecordZoneID, serverChangeToken ServerChangeToken, clientChangeTokenData []byte, moreComing bool, recordZoneError foundation.Error))
	RecordZoneChangeTokensUpdatedBlock() func(recordZoneID RecordZoneID, serverChangeToken ServerChangeToken, clientChangeTokenData []byte)
	SetRecordZoneChangeTokensUpdatedBlock(value func(recordZoneID RecordZoneID, serverChangeToken ServerChangeToken, clientChangeTokenData []byte))
}

// An operation that fetches record zone changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation?language=objc
type FetchRecordZoneChangesOperation struct {
	DatabaseOperation
}

func FetchRecordZoneChangesOperationFrom(ptr unsafe.Pointer) FetchRecordZoneChangesOperation {
	return FetchRecordZoneChangesOperation{
		DatabaseOperation: DatabaseOperationFrom(ptr),
	}
}

func (f_ FetchRecordZoneChangesOperation) InitWithRecordZoneIDsConfigurationsByRecordZoneID(recordZoneIDs []IRecordZoneID, configurationsByRecordZoneID foundation.Dictionary) FetchRecordZoneChangesOperation {
	rv := objc.Call[FetchRecordZoneChangesOperation](f_, objc.Sel("initWithRecordZoneIDs:configurationsByRecordZoneID:"), recordZoneIDs, configurationsByRecordZoneID)
	return rv
}

// Creates an operation for fetching record zone changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/2980665-initwithrecordzoneids?language=objc
func NewFetchRecordZoneChangesOperationWithRecordZoneIDsConfigurationsByRecordZoneID(recordZoneIDs []IRecordZoneID, configurationsByRecordZoneID foundation.Dictionary) FetchRecordZoneChangesOperation {
	instance := FetchRecordZoneChangesOperationClass.Alloc().InitWithRecordZoneIDsConfigurationsByRecordZoneID(recordZoneIDs, configurationsByRecordZoneID)
	instance.Autorelease()
	return instance
}

func (f_ FetchRecordZoneChangesOperation) Init() FetchRecordZoneChangesOperation {
	rv := objc.Call[FetchRecordZoneChangesOperation](f_, objc.Sel("init"))
	return rv
}

func (fc _FetchRecordZoneChangesOperationClass) Alloc() FetchRecordZoneChangesOperation {
	rv := objc.Call[FetchRecordZoneChangesOperation](fc, objc.Sel("alloc"))
	return rv
}

func FetchRecordZoneChangesOperation_Alloc() FetchRecordZoneChangesOperation {
	return FetchRecordZoneChangesOperationClass.Alloc()
}

func (fc _FetchRecordZoneChangesOperationClass) New() FetchRecordZoneChangesOperation {
	rv := objc.Call[FetchRecordZoneChangesOperation](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchRecordZoneChangesOperation() FetchRecordZoneChangesOperation {
	return FetchRecordZoneChangesOperationClass.New()
}

// The IDs of the record zones that contain the records to fetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/1640463-recordzoneids?language=objc
func (f_ FetchRecordZoneChangesOperation) RecordZoneIDs() []RecordZoneID {
	rv := objc.Call[[]RecordZoneID](f_, objc.Sel("recordZoneIDs"))
	return rv
}

// The IDs of the record zones that contain the records to fetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/1640463-recordzoneids?language=objc
func (f_ FetchRecordZoneChangesOperation) SetRecordZoneIDs(value []IRecordZoneID) {
	objc.Call[objc.Void](f_, objc.Sel("setRecordZoneIDs:"), value)
}

// The block to execute when a record no longer exists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/1640470-recordwithidwasdeletedblock?language=objc
func (f_ FetchRecordZoneChangesOperation) RecordWithIDWasDeletedBlock() func(recordID RecordID, recordType RecordType) {
	rv := objc.Call[func(recordID RecordID, recordType RecordType)](f_, objc.Sel("recordWithIDWasDeletedBlock"))
	return rv
}

// The block to execute when a record no longer exists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/1640470-recordwithidwasdeletedblock?language=objc
func (f_ FetchRecordZoneChangesOperation) SetRecordWithIDWasDeletedBlock(value func(recordID RecordID, recordType RecordType)) {
	objc.Call[objc.Void](f_, objc.Sel("setRecordWithIDWasDeletedBlock:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/3793698-recordwaschangedblock?language=objc
func (f_ FetchRecordZoneChangesOperation) RecordWasChangedBlock() func(recordID RecordID, record Record, error foundation.Error) {
	rv := objc.Call[func(recordID RecordID, record Record, error foundation.Error)](f_, objc.Sel("recordWasChangedBlock"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/3793698-recordwaschangedblock?language=objc
func (f_ FetchRecordZoneChangesOperation) SetRecordWasChangedBlock(value func(recordID RecordID, record Record, error foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("setRecordWasChangedBlock:"), value)
}

// The block to execute when the operation finishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/1640409-fetchrecordzonechangescompletion?language=objc
func (f_ FetchRecordZoneChangesOperation) FetchRecordZoneChangesCompletionBlock() func(operationError foundation.Error) {
	rv := objc.Call[func(operationError foundation.Error)](f_, objc.Sel("fetchRecordZoneChangesCompletionBlock"))
	return rv
}

// The block to execute when the operation finishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/1640409-fetchrecordzonechangescompletion?language=objc
func (f_ FetchRecordZoneChangesOperation) SetFetchRecordZoneChangesCompletionBlock(value func(operationError foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("setFetchRecordZoneChangesCompletionBlock:"), value)
}

// A dictionary of configurations for fetching change operations by zone identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/2980664-configurationsbyrecordzoneid?language=objc
func (f_ FetchRecordZoneChangesOperation) ConfigurationsByRecordZoneID() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](f_, objc.Sel("configurationsByRecordZoneID"))
	return rv
}

// A dictionary of configurations for fetching change operations by zone identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/2980664-configurationsbyrecordzoneid?language=objc
func (f_ FetchRecordZoneChangesOperation) SetConfigurationsByRecordZoneID(value foundation.Dictionary) {
	objc.Call[objc.Void](f_, objc.Sel("setConfigurationsByRecordZoneID:"), value)
}

// A Boolean value that indicates whether to send repeated requests to the server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/1640386-fetchallchanges?language=objc
func (f_ FetchRecordZoneChangesOperation) FetchAllChanges() bool {
	rv := objc.Call[bool](f_, objc.Sel("fetchAllChanges"))
	return rv
}

// A Boolean value that indicates whether to send repeated requests to the server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/1640386-fetchallchanges?language=objc
func (f_ FetchRecordZoneChangesOperation) SetFetchAllChanges(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setFetchAllChanges:"), value)
}

// The block to execute when a record zone’s fetch finishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/1640411-recordzonefetchcompletionblock?language=objc
func (f_ FetchRecordZoneChangesOperation) RecordZoneFetchCompletionBlock() func(recordZoneID RecordZoneID, serverChangeToken ServerChangeToken, clientChangeTokenData []byte, moreComing bool, recordZoneError foundation.Error) {
	rv := objc.Call[func(recordZoneID RecordZoneID, serverChangeToken ServerChangeToken, clientChangeTokenData []byte, moreComing bool, recordZoneError foundation.Error)](f_, objc.Sel("recordZoneFetchCompletionBlock"))
	return rv
}

// The block to execute when a record zone’s fetch finishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/1640411-recordzonefetchcompletionblock?language=objc
func (f_ FetchRecordZoneChangesOperation) SetRecordZoneFetchCompletionBlock(value func(recordZoneID RecordZoneID, serverChangeToken ServerChangeToken, clientChangeTokenData []byte, moreComing bool, recordZoneError foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("setRecordZoneFetchCompletionBlock:"), value)
}

// The block to execute when the change token updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/1640422-recordzonechangetokensupdatedblo?language=objc
func (f_ FetchRecordZoneChangesOperation) RecordZoneChangeTokensUpdatedBlock() func(recordZoneID RecordZoneID, serverChangeToken ServerChangeToken, clientChangeTokenData []byte) {
	rv := objc.Call[func(recordZoneID RecordZoneID, serverChangeToken ServerChangeToken, clientChangeTokenData []byte)](f_, objc.Sel("recordZoneChangeTokensUpdatedBlock"))
	return rv
}

// The block to execute when the change token updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/1640422-recordzonechangetokensupdatedblo?language=objc
func (f_ FetchRecordZoneChangesOperation) SetRecordZoneChangeTokensUpdatedBlock(value func(recordZoneID RecordZoneID, serverChangeToken ServerChangeToken, clientChangeTokenData []byte)) {
	objc.Call[objc.Void](f_, objc.Sel("setRecordZoneChangeTokensUpdatedBlock:"), value)
}
