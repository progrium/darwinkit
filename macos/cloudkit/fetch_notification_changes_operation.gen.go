// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchNotificationChangesOperation] class.
var FetchNotificationChangesOperationClass = _FetchNotificationChangesOperationClass{objc.GetClass("CKFetchNotificationChangesOperation")}

type _FetchNotificationChangesOperationClass struct {
	objc.Class
}

// An interface definition for the [FetchNotificationChangesOperation] class.
type IFetchNotificationChangesOperation interface {
	IOperation
}

// An operation that retrieves unread notifications from a CloudKit container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchnotificationchangesoperation?language=objc
type FetchNotificationChangesOperation struct {
	Operation
}

func FetchNotificationChangesOperationFrom(ptr unsafe.Pointer) FetchNotificationChangesOperation {
	return FetchNotificationChangesOperation{
		Operation: OperationFrom(ptr),
	}
}

func (fc _FetchNotificationChangesOperationClass) Alloc() FetchNotificationChangesOperation {
	rv := objc.Call[FetchNotificationChangesOperation](fc, objc.Sel("alloc"))
	return rv
}

func FetchNotificationChangesOperation_Alloc() FetchNotificationChangesOperation {
	return FetchNotificationChangesOperationClass.Alloc()
}

func (fc _FetchNotificationChangesOperationClass) New() FetchNotificationChangesOperation {
	rv := objc.Call[FetchNotificationChangesOperation](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchNotificationChangesOperation() FetchNotificationChangesOperation {
	return FetchNotificationChangesOperationClass.New()
}

func (f_ FetchNotificationChangesOperation) Init() FetchNotificationChangesOperation {
	rv := objc.Call[FetchNotificationChangesOperation](f_, objc.Sel("init"))
	return rv
}
