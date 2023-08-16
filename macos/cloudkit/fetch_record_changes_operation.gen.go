// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchRecordChangesOperation] class.
var FetchRecordChangesOperationClass = _FetchRecordChangesOperationClass{objc.GetClass("CKFetchRecordChangesOperation")}

type _FetchRecordChangesOperationClass struct {
	objc.Class
}

// An interface definition for the [FetchRecordChangesOperation] class.
type IFetchRecordChangesOperation interface {
	IDatabaseOperation
}

// An operation that reports on the changed and deleted records in the specified record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordchangesoperation?language=objc
type FetchRecordChangesOperation struct {
	DatabaseOperation
}

func FetchRecordChangesOperationFrom(ptr unsafe.Pointer) FetchRecordChangesOperation {
	return FetchRecordChangesOperation{
		DatabaseOperation: DatabaseOperationFrom(ptr),
	}
}

func (fc _FetchRecordChangesOperationClass) Alloc() FetchRecordChangesOperation {
	rv := objc.Call[FetchRecordChangesOperation](fc, objc.Sel("alloc"))
	return rv
}

func FetchRecordChangesOperation_Alloc() FetchRecordChangesOperation {
	return FetchRecordChangesOperationClass.Alloc()
}

func (fc _FetchRecordChangesOperationClass) New() FetchRecordChangesOperation {
	rv := objc.Call[FetchRecordChangesOperation](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchRecordChangesOperation() FetchRecordChangesOperation {
	return FetchRecordChangesOperationClass.New()
}

func (f_ FetchRecordChangesOperation) Init() FetchRecordChangesOperation {
	rv := objc.Call[FetchRecordChangesOperation](f_, objc.Sel("init"))
	return rv
}
