// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MarkNotificationsReadOperation] class.
var MarkNotificationsReadOperationClass = _MarkNotificationsReadOperationClass{objc.GetClass("CKMarkNotificationsReadOperation")}

type _MarkNotificationsReadOperationClass struct {
	objc.Class
}

// An interface definition for the [MarkNotificationsReadOperation] class.
type IMarkNotificationsReadOperation interface {
	IOperation
}

// An operation that marks push notifications as read by your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmarknotificationsreadoperation?language=objc
type MarkNotificationsReadOperation struct {
	Operation
}

func MarkNotificationsReadOperationFrom(ptr unsafe.Pointer) MarkNotificationsReadOperation {
	return MarkNotificationsReadOperation{
		Operation: OperationFrom(ptr),
	}
}

func (mc _MarkNotificationsReadOperationClass) Alloc() MarkNotificationsReadOperation {
	rv := objc.Call[MarkNotificationsReadOperation](mc, objc.Sel("alloc"))
	return rv
}

func MarkNotificationsReadOperation_Alloc() MarkNotificationsReadOperation {
	return MarkNotificationsReadOperationClass.Alloc()
}

func (mc _MarkNotificationsReadOperationClass) New() MarkNotificationsReadOperation {
	rv := objc.Call[MarkNotificationsReadOperation](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMarkNotificationsReadOperation() MarkNotificationsReadOperation {
	return MarkNotificationsReadOperationClass.New()
}

func (m_ MarkNotificationsReadOperation) Init() MarkNotificationsReadOperation {
	rv := objc.Call[MarkNotificationsReadOperation](m_, objc.Sel("init"))
	return rv
}
