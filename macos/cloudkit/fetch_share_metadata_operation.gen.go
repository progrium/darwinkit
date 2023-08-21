// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchShareMetadataOperation] class.
var FetchShareMetadataOperationClass = _FetchShareMetadataOperationClass{objc.GetClass("CKFetchShareMetadataOperation")}

type _FetchShareMetadataOperationClass struct {
	objc.Class
}

// An interface definition for the [FetchShareMetadataOperation] class.
type IFetchShareMetadataOperation interface {
	IOperation
	RootRecordDesiredKeys() []RecordFieldKey
	SetRootRecordDesiredKeys(value []RecordFieldKey)
	FetchShareMetadataCompletionBlock() func(operationError foundation.Error)
	SetFetchShareMetadataCompletionBlock(value func(operationError foundation.Error))
	ShareURLs() []foundation.URL
	SetShareURLs(value []foundation.IURL)
	PerShareMetadataBlock() func(shareURL foundation.URL, shareMetadata ShareMetadata, error foundation.Error)
	SetPerShareMetadataBlock(value func(shareURL foundation.URL, shareMetadata ShareMetadata, error foundation.Error))
	ShouldFetchRootRecord() bool
	SetShouldFetchRootRecord(value bool)
}

// An operation that fetches metadata for one or more shares. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation?language=objc
type FetchShareMetadataOperation struct {
	Operation
}

func FetchShareMetadataOperationFrom(ptr unsafe.Pointer) FetchShareMetadataOperation {
	return FetchShareMetadataOperation{
		Operation: OperationFrom(ptr),
	}
}

func (f_ FetchShareMetadataOperation) InitWithShareURLs(shareURLs []foundation.IURL) FetchShareMetadataOperation {
	rv := objc.Call[FetchShareMetadataOperation](f_, objc.Sel("initWithShareURLs:"), shareURLs)
	return rv
}

// Creates an operation for fetching the metadata for the specified shares. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/1640495-initwithshareurls?language=objc
func NewFetchShareMetadataOperationWithShareURLs(shareURLs []foundation.IURL) FetchShareMetadataOperation {
	instance := FetchShareMetadataOperationClass.Alloc().InitWithShareURLs(shareURLs)
	instance.Autorelease()
	return instance
}

func (f_ FetchShareMetadataOperation) Init() FetchShareMetadataOperation {
	rv := objc.Call[FetchShareMetadataOperation](f_, objc.Sel("init"))
	return rv
}

func (fc _FetchShareMetadataOperationClass) Alloc() FetchShareMetadataOperation {
	rv := objc.Call[FetchShareMetadataOperation](fc, objc.Sel("alloc"))
	return rv
}

func FetchShareMetadataOperation_Alloc() FetchShareMetadataOperation {
	return FetchShareMetadataOperationClass.Alloc()
}

func (fc _FetchShareMetadataOperationClass) New() FetchShareMetadataOperation {
	rv := objc.Call[FetchShareMetadataOperation](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchShareMetadataOperation() FetchShareMetadataOperation {
	return FetchShareMetadataOperationClass.New()
}

// The fields to return when fetching the root record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/1640375-rootrecorddesiredkeys?language=objc
func (f_ FetchShareMetadataOperation) RootRecordDesiredKeys() []RecordFieldKey {
	rv := objc.Call[[]RecordFieldKey](f_, objc.Sel("rootRecordDesiredKeys"))
	return rv
}

// The fields to return when fetching the root record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/1640375-rootrecorddesiredkeys?language=objc
func (f_ FetchShareMetadataOperation) SetRootRecordDesiredKeys(value []RecordFieldKey) {
	objc.Call[objc.Void](f_, objc.Sel("setRootRecordDesiredKeys:"), value)
}

// The block to execute when the operation finishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/1640457-fetchsharemetadatacompletionbloc?language=objc
func (f_ FetchShareMetadataOperation) FetchShareMetadataCompletionBlock() func(operationError foundation.Error) {
	rv := objc.Call[func(operationError foundation.Error)](f_, objc.Sel("fetchShareMetadataCompletionBlock"))
	return rv
}

// The block to execute when the operation finishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/1640457-fetchsharemetadatacompletionbloc?language=objc
func (f_ FetchShareMetadataOperation) SetFetchShareMetadataCompletionBlock(value func(operationError foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("setFetchShareMetadataCompletionBlock:"), value)
}

// The URLs of the shares to fetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/1640508-shareurls?language=objc
func (f_ FetchShareMetadataOperation) ShareURLs() []foundation.URL {
	rv := objc.Call[[]foundation.URL](f_, objc.Sel("shareURLs"))
	return rv
}

// The URLs of the shares to fetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/1640508-shareurls?language=objc
func (f_ FetchShareMetadataOperation) SetShareURLs(value []foundation.IURL) {
	objc.Call[objc.Void](f_, objc.Sel("setShareURLs:"), value)
}

// The block to execute as the operation fetches individual shares. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/1640447-persharemetadatablock?language=objc
func (f_ FetchShareMetadataOperation) PerShareMetadataBlock() func(shareURL foundation.URL, shareMetadata ShareMetadata, error foundation.Error) {
	rv := objc.Call[func(shareURL foundation.URL, shareMetadata ShareMetadata, error foundation.Error)](f_, objc.Sel("perShareMetadataBlock"))
	return rv
}

// The block to execute as the operation fetches individual shares. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/1640447-persharemetadatablock?language=objc
func (f_ FetchShareMetadataOperation) SetPerShareMetadataBlock(value func(shareURL foundation.URL, shareMetadata ShareMetadata, error foundation.Error)) {
	objc.Call[objc.Void](f_, objc.Sel("setPerShareMetadataBlock:"), value)
}

// A Boolean value that indicates whether to retrieve the root record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/1640519-shouldfetchrootrecord?language=objc
func (f_ FetchShareMetadataOperation) ShouldFetchRootRecord() bool {
	rv := objc.Call[bool](f_, objc.Sel("shouldFetchRootRecord"))
	return rv
}

// A Boolean value that indicates whether to retrieve the root record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/1640519-shouldfetchrootrecord?language=objc
func (f_ FetchShareMetadataOperation) SetShouldFetchRootRecord(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setShouldFetchRootRecord:"), value)
}
