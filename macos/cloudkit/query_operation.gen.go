// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [QueryOperation] class.
var QueryOperationClass = _QueryOperationClass{objc.GetClass("CKQueryOperation")}

type _QueryOperationClass struct {
	objc.Class
}

// An interface definition for the [QueryOperation] class.
type IQueryOperation interface {
	IDatabaseOperation
	Cursor() QueryCursor
	SetCursor(value IQueryCursor)
	ZoneID() RecordZoneID
	SetZoneID(value IRecordZoneID)
	Query() Query
	SetQuery(value IQuery)
	RecordMatchedBlock() func(recordID RecordID, record Record, error foundation.Error)
	SetRecordMatchedBlock(value func(recordID RecordID, record Record, error foundation.Error))
	QueryCompletionBlock() func(cursor QueryCursor, operationError foundation.Error)
	SetQueryCompletionBlock(value func(cursor QueryCursor, operationError foundation.Error))
	ResultsLimit() uint
	SetResultsLimit(value uint)
	DesiredKeys() []RecordFieldKey
	SetDesiredKeys(value []RecordFieldKey)
}

// An operation for executing queries in a database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation?language=objc
type QueryOperation struct {
	DatabaseOperation
}

func QueryOperationFrom(ptr unsafe.Pointer) QueryOperation {
	return QueryOperation{
		DatabaseOperation: DatabaseOperationFrom(ptr),
	}
}

func (q_ QueryOperation) InitWithQuery(query IQuery) QueryOperation {
	rv := objc.Call[QueryOperation](q_, objc.Sel("initWithQuery:"), objc.Ptr(query))
	return rv
}

// Creates an operation that searches for records in the specified record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/1514958-initwithquery?language=objc
func NewQueryOperationWithQuery(query IQuery) QueryOperation {
	instance := QueryOperationClass.Alloc().InitWithQuery(query)
	instance.Autorelease()
	return instance
}

func (q_ QueryOperation) InitWithCursor(cursor IQueryCursor) QueryOperation {
	rv := objc.Call[QueryOperation](q_, objc.Sel("initWithCursor:"), objc.Ptr(cursor))
	return rv
}

// Creates an operation with additional results from a previous search. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/1515033-initwithcursor?language=objc
func NewQueryOperationWithCursor(cursor IQueryCursor) QueryOperation {
	instance := QueryOperationClass.Alloc().InitWithCursor(cursor)
	instance.Autorelease()
	return instance
}

func (q_ QueryOperation) Init() QueryOperation {
	rv := objc.Call[QueryOperation](q_, objc.Sel("init"))
	return rv
}

func (qc _QueryOperationClass) Alloc() QueryOperation {
	rv := objc.Call[QueryOperation](qc, objc.Sel("alloc"))
	return rv
}

func QueryOperation_Alloc() QueryOperation {
	return QueryOperationClass.Alloc()
}

func (qc _QueryOperationClass) New() QueryOperation {
	rv := objc.Call[QueryOperation](qc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewQueryOperation() QueryOperation {
	return QueryOperationClass.New()
}

// The  cursor for continuing the search. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/1514975-cursor?language=objc
func (q_ QueryOperation) Cursor() QueryCursor {
	rv := objc.Call[QueryCursor](q_, objc.Sel("cursor"))
	return rv
}

// The  cursor for continuing the search. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/1514975-cursor?language=objc
func (q_ QueryOperation) SetCursor(value IQueryCursor) {
	objc.Call[objc.Void](q_, objc.Sel("setCursor:"), objc.Ptr(value))
}

// The ID of the record zone that contains the records to search. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/1515269-zoneid?language=objc
func (q_ QueryOperation) ZoneID() RecordZoneID {
	rv := objc.Call[RecordZoneID](q_, objc.Sel("zoneID"))
	return rv
}

// The ID of the record zone that contains the records to search. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/1515269-zoneid?language=objc
func (q_ QueryOperation) SetZoneID(value IRecordZoneID) {
	objc.Call[objc.Void](q_, objc.Sel("setZoneID:"), objc.Ptr(value))
}

// The query for the search. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/1515127-query?language=objc
func (q_ QueryOperation) Query() Query {
	rv := objc.Call[Query](q_, objc.Sel("query"))
	return rv
}

// The query for the search. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/1515127-query?language=objc
func (q_ QueryOperation) SetQuery(value IQuery) {
	objc.Call[objc.Void](q_, objc.Sel("setQuery:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/3793708-recordmatchedblock?language=objc
func (q_ QueryOperation) RecordMatchedBlock() func(recordID RecordID, record Record, error foundation.Error) {
	rv := objc.Call[func(recordID RecordID, record Record, error foundation.Error)](q_, objc.Sel("recordMatchedBlock"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/3793708-recordmatchedblock?language=objc
func (q_ QueryOperation) SetRecordMatchedBlock(value func(recordID RecordID, record Record, error foundation.Error)) {
	objc.Call[objc.Void](q_, objc.Sel("setRecordMatchedBlock:"), value)
}

// The block to execute after CloudKit retrieves all of the records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/1515067-querycompletionblock?language=objc
func (q_ QueryOperation) QueryCompletionBlock() func(cursor QueryCursor, operationError foundation.Error) {
	rv := objc.Call[func(cursor QueryCursor, operationError foundation.Error)](q_, objc.Sel("queryCompletionBlock"))
	return rv
}

// The block to execute after CloudKit retrieves all of the records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/1515067-querycompletionblock?language=objc
func (q_ QueryOperation) SetQueryCompletionBlock(value func(cursor QueryCursor, operationError foundation.Error)) {
	objc.Call[objc.Void](q_, objc.Sel("setQueryCompletionBlock:"), value)
}

// The maximum number of records to return at one time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/1515078-resultslimit?language=objc
func (q_ QueryOperation) ResultsLimit() uint {
	rv := objc.Call[uint](q_, objc.Sel("resultsLimit"))
	return rv
}

// The maximum number of records to return at one time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/1515078-resultslimit?language=objc
func (q_ QueryOperation) SetResultsLimit(value uint) {
	objc.Call[objc.Void](q_, objc.Sel("setResultsLimit:"), value)
}

// The fields of the records to fetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/1515268-desiredkeys?language=objc
func (q_ QueryOperation) DesiredKeys() []RecordFieldKey {
	rv := objc.Call[[]RecordFieldKey](q_, objc.Sel("desiredKeys"))
	return rv
}

// The fields of the records to fetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/1515268-desiredkeys?language=objc
func (q_ QueryOperation) SetDesiredKeys(value []RecordFieldKey) {
	objc.Call[objc.Void](q_, objc.Sel("setDesiredKeys:"), value)
}
