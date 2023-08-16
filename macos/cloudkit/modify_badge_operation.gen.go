// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ModifyBadgeOperation] class.
var ModifyBadgeOperationClass = _ModifyBadgeOperationClass{objc.GetClass("CKModifyBadgeOperation")}

type _ModifyBadgeOperationClass struct {
	objc.Class
}

// An interface definition for the [ModifyBadgeOperation] class.
type IModifyBadgeOperation interface {
	IOperation
}

// An operation that sets the value of the app icon’s badge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifybadgeoperation?language=objc
type ModifyBadgeOperation struct {
	Operation
}

func ModifyBadgeOperationFrom(ptr unsafe.Pointer) ModifyBadgeOperation {
	return ModifyBadgeOperation{
		Operation: OperationFrom(ptr),
	}
}

func (mc _ModifyBadgeOperationClass) Alloc() ModifyBadgeOperation {
	rv := objc.Call[ModifyBadgeOperation](mc, objc.Sel("alloc"))
	return rv
}

func ModifyBadgeOperation_Alloc() ModifyBadgeOperation {
	return ModifyBadgeOperationClass.Alloc()
}

func (mc _ModifyBadgeOperationClass) New() ModifyBadgeOperation {
	rv := objc.Call[ModifyBadgeOperation](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewModifyBadgeOperation() ModifyBadgeOperation {
	return ModifyBadgeOperationClass.New()
}

func (m_ ModifyBadgeOperation) Init() ModifyBadgeOperation {
	rv := objc.Call[ModifyBadgeOperation](m_, objc.Sel("init"))
	return rv
}
