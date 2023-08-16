// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DiscoverAllUserIdentitiesOperation] class.
var DiscoverAllUserIdentitiesOperationClass = _DiscoverAllUserIdentitiesOperationClass{objc.GetClass("CKDiscoverAllUserIdentitiesOperation")}

type _DiscoverAllUserIdentitiesOperationClass struct {
	objc.Class
}

// An interface definition for the [DiscoverAllUserIdentitiesOperation] class.
type IDiscoverAllUserIdentitiesOperation interface {
	IOperation
}

// An operation that uses the device’s contacts to search for discoverable iCloud users. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdiscoveralluseridentitiesoperation?language=objc
type DiscoverAllUserIdentitiesOperation struct {
	Operation
}

func DiscoverAllUserIdentitiesOperationFrom(ptr unsafe.Pointer) DiscoverAllUserIdentitiesOperation {
	return DiscoverAllUserIdentitiesOperation{
		Operation: OperationFrom(ptr),
	}
}

func (dc _DiscoverAllUserIdentitiesOperationClass) Alloc() DiscoverAllUserIdentitiesOperation {
	rv := objc.Call[DiscoverAllUserIdentitiesOperation](dc, objc.Sel("alloc"))
	return rv
}

func DiscoverAllUserIdentitiesOperation_Alloc() DiscoverAllUserIdentitiesOperation {
	return DiscoverAllUserIdentitiesOperationClass.Alloc()
}

func (dc _DiscoverAllUserIdentitiesOperationClass) New() DiscoverAllUserIdentitiesOperation {
	rv := objc.Call[DiscoverAllUserIdentitiesOperation](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDiscoverAllUserIdentitiesOperation() DiscoverAllUserIdentitiesOperation {
	return DiscoverAllUserIdentitiesOperationClass.New()
}

func (d_ DiscoverAllUserIdentitiesOperation) Init() DiscoverAllUserIdentitiesOperation {
	rv := objc.Call[DiscoverAllUserIdentitiesOperation](d_, objc.Sel("init"))
	return rv
}
