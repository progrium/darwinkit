// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DiscoverUserIdentitiesOperation] class.
var DiscoverUserIdentitiesOperationClass = _DiscoverUserIdentitiesOperationClass{objc.GetClass("CKDiscoverUserIdentitiesOperation")}

type _DiscoverUserIdentitiesOperationClass struct {
	objc.Class
}

// An interface definition for the [DiscoverUserIdentitiesOperation] class.
type IDiscoverUserIdentitiesOperation interface {
	IOperation
}

// An operation that uses the provided criteria to search for discoverable iCloud users. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdiscoveruseridentitiesoperation?language=objc
type DiscoverUserIdentitiesOperation struct {
	Operation
}

func DiscoverUserIdentitiesOperationFrom(ptr unsafe.Pointer) DiscoverUserIdentitiesOperation {
	return DiscoverUserIdentitiesOperation{
		Operation: OperationFrom(ptr),
	}
}

func (dc _DiscoverUserIdentitiesOperationClass) Alloc() DiscoverUserIdentitiesOperation {
	rv := objc.Call[DiscoverUserIdentitiesOperation](dc, objc.Sel("alloc"))
	return rv
}

func DiscoverUserIdentitiesOperation_Alloc() DiscoverUserIdentitiesOperation {
	return DiscoverUserIdentitiesOperationClass.Alloc()
}

func (dc _DiscoverUserIdentitiesOperationClass) New() DiscoverUserIdentitiesOperation {
	rv := objc.Call[DiscoverUserIdentitiesOperation](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDiscoverUserIdentitiesOperation() DiscoverUserIdentitiesOperation {
	return DiscoverUserIdentitiesOperationClass.New()
}

func (d_ DiscoverUserIdentitiesOperation) Init() DiscoverUserIdentitiesOperation {
	rv := objc.Call[DiscoverUserIdentitiesOperation](d_, objc.Sel("init"))
	return rv
}
