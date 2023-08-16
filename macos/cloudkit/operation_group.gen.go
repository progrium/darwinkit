// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [OperationGroup] class.
var OperationGroupClass = _OperationGroupClass{objc.GetClass("CKOperationGroup")}

type _OperationGroupClass struct {
	objc.Class
}

// An interface definition for the [OperationGroup] class.
type IOperationGroup interface {
	objc.IObject
	OperationGroupID() string
	Name() string
	SetName(value string)
	DefaultConfiguration() OperationConfiguration
	SetDefaultConfiguration(value IOperationConfiguration)
	ExpectedSendSize() OperationGroupTransferSize
	SetExpectedSendSize(value OperationGroupTransferSize)
	Quantity() uint
	SetQuantity(value uint)
	ExpectedReceiveSize() OperationGroupTransferSize
	SetExpectedReceiveSize(value OperationGroupTransferSize)
}

// An explicit association between two or more operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgroup?language=objc
type OperationGroup struct {
	objc.Object
}

func OperationGroupFrom(ptr unsafe.Pointer) OperationGroup {
	return OperationGroup{
		Object: objc.ObjectFrom(ptr),
	}
}

func (o_ OperationGroup) Init() OperationGroup {
	rv := objc.Call[OperationGroup](o_, objc.Sel("init"))
	return rv
}

func (oc _OperationGroupClass) Alloc() OperationGroup {
	rv := objc.Call[OperationGroup](oc, objc.Sel("alloc"))
	return rv
}

func OperationGroup_Alloc() OperationGroup {
	return OperationGroupClass.Alloc()
}

func (oc _OperationGroupClass) New() OperationGroup {
	rv := objc.Call[OperationGroup](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOperationGroup() OperationGroup {
	return OperationGroupClass.New()
}

// The operation group’s unique identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgroup/2866238-operationgroupid?language=objc
func (o_ OperationGroup) OperationGroupID() string {
	rv := objc.Call[string](o_, objc.Sel("operationGroupID"))
	return rv
}

// The operation group’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgroup/2866233-name?language=objc
func (o_ OperationGroup) Name() string {
	rv := objc.Call[string](o_, objc.Sel("name"))
	return rv
}

// The operation group’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgroup/2866233-name?language=objc
func (o_ OperationGroup) SetName(value string) {
	objc.Call[objc.Void](o_, objc.Sel("setName:"), value)
}

// The default configuration for operations in the group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgroup/2866240-defaultconfiguration?language=objc
func (o_ OperationGroup) DefaultConfiguration() OperationConfiguration {
	rv := objc.Call[OperationConfiguration](o_, objc.Sel("defaultConfiguration"))
	return rv
}

// The default configuration for operations in the group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgroup/2866240-defaultconfiguration?language=objc
func (o_ OperationGroup) SetDefaultConfiguration(value IOperationConfiguration) {
	objc.Call[objc.Void](o_, objc.Sel("setDefaultConfiguration:"), objc.Ptr(value))
}

// The estimated size of traffic to upload to CloudKit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgroup/2866235-expectedsendsize?language=objc
func (o_ OperationGroup) ExpectedSendSize() OperationGroupTransferSize {
	rv := objc.Call[OperationGroupTransferSize](o_, objc.Sel("expectedSendSize"))
	return rv
}

// The estimated size of traffic to upload to CloudKit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgroup/2866235-expectedsendsize?language=objc
func (o_ OperationGroup) SetExpectedSendSize(value OperationGroupTransferSize) {
	objc.Call[objc.Void](o_, objc.Sel("setExpectedSendSize:"), value)
}

// The number of operations in the operation group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgroup/2866220-quantity?language=objc
func (o_ OperationGroup) Quantity() uint {
	rv := objc.Call[uint](o_, objc.Sel("quantity"))
	return rv
}

// The number of operations in the operation group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgroup/2866220-quantity?language=objc
func (o_ OperationGroup) SetQuantity(value uint) {
	objc.Call[objc.Void](o_, objc.Sel("setQuantity:"), value)
}

// The estimated size of traffic to download from CloudKit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgroup/2866229-expectedreceivesize?language=objc
func (o_ OperationGroup) ExpectedReceiveSize() OperationGroupTransferSize {
	rv := objc.Call[OperationGroupTransferSize](o_, objc.Sel("expectedReceiveSize"))
	return rv
}

// The estimated size of traffic to download from CloudKit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgroup/2866229-expectedreceivesize?language=objc
func (o_ OperationGroup) SetExpectedReceiveSize(value OperationGroupTransferSize) {
	objc.Call[objc.Void](o_, objc.Sel("setExpectedReceiveSize:"), value)
}
