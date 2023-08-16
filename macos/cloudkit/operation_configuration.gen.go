// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [OperationConfiguration] class.
var OperationConfigurationClass = _OperationConfigurationClass{objc.GetClass("CKOperationConfiguration")}

type _OperationConfigurationClass struct {
	objc.Class
}

// An interface definition for the [OperationConfiguration] class.
type IOperationConfiguration interface {
	objc.IObject
	QualityOfService() foundation.QualityOfService
	SetQualityOfService(value foundation.QualityOfService)
	AllowsCellularAccess() bool
	SetAllowsCellularAccess(value bool)
	Container() Container
	SetContainer(value IContainer)
	TimeoutIntervalForRequest() foundation.TimeInterval
	SetTimeoutIntervalForRequest(value foundation.TimeInterval)
	IsLongLived() bool
	SetLongLived(value bool)
	TimeoutIntervalForResource() foundation.TimeInterval
	SetTimeoutIntervalForResource(value foundation.TimeInterval)
}

// An object that describes how a CloudKit operation behaves. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationconfiguration?language=objc
type OperationConfiguration struct {
	objc.Object
}

func OperationConfigurationFrom(ptr unsafe.Pointer) OperationConfiguration {
	return OperationConfiguration{
		Object: objc.ObjectFrom(ptr),
	}
}

func (oc _OperationConfigurationClass) Alloc() OperationConfiguration {
	rv := objc.Call[OperationConfiguration](oc, objc.Sel("alloc"))
	return rv
}

func OperationConfiguration_Alloc() OperationConfiguration {
	return OperationConfigurationClass.Alloc()
}

func (oc _OperationConfigurationClass) New() OperationConfiguration {
	rv := objc.Call[OperationConfiguration](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOperationConfiguration() OperationConfiguration {
	return OperationConfigurationClass.New()
}

func (o_ OperationConfiguration) Init() OperationConfiguration {
	rv := objc.Call[OperationConfiguration](o_, objc.Sel("init"))
	return rv
}

// The priority that the system uses when it allocates resources to the operations that use this configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationconfiguration/2866201-qualityofservice?language=objc
func (o_ OperationConfiguration) QualityOfService() foundation.QualityOfService {
	rv := objc.Call[foundation.QualityOfService](o_, objc.Sel("qualityOfService"))
	return rv
}

// The priority that the system uses when it allocates resources to the operations that use this configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationconfiguration/2866201-qualityofservice?language=objc
func (o_ OperationConfiguration) SetQualityOfService(value foundation.QualityOfService) {
	objc.Call[objc.Void](o_, objc.Sel("setQualityOfService:"), value)
}

// A Boolean value that indicates whether operations that use this configuration can send data over the cellular network. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationconfiguration/2866217-allowscellularaccess?language=objc
func (o_ OperationConfiguration) AllowsCellularAccess() bool {
	rv := objc.Call[bool](o_, objc.Sel("allowsCellularAccess"))
	return rv
}

// A Boolean value that indicates whether operations that use this configuration can send data over the cellular network. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationconfiguration/2866217-allowscellularaccess?language=objc
func (o_ OperationConfiguration) SetAllowsCellularAccess(value bool) {
	objc.Call[objc.Void](o_, objc.Sel("setAllowsCellularAccess:"), value)
}

// The configuration’s container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationconfiguration/2866232-container?language=objc
func (o_ OperationConfiguration) Container() Container {
	rv := objc.Call[Container](o_, objc.Sel("container"))
	return rv
}

// The configuration’s container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationconfiguration/2866232-container?language=objc
func (o_ OperationConfiguration) SetContainer(value IContainer) {
	objc.Call[objc.Void](o_, objc.Sel("setContainer:"), objc.Ptr(value))
}

// The maximum amount of time that a request can take. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationconfiguration/2866210-timeoutintervalforrequest?language=objc
func (o_ OperationConfiguration) TimeoutIntervalForRequest() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](o_, objc.Sel("timeoutIntervalForRequest"))
	return rv
}

// The maximum amount of time that a request can take. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationconfiguration/2866210-timeoutintervalforrequest?language=objc
func (o_ OperationConfiguration) SetTimeoutIntervalForRequest(value foundation.TimeInterval) {
	objc.Call[objc.Void](o_, objc.Sel("setTimeoutIntervalForRequest:"), value)
}

// A Boolean value that indicates whether the operations that use this configuration are long-lived. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationconfiguration/2866225-longlived?language=objc
func (o_ OperationConfiguration) IsLongLived() bool {
	rv := objc.Call[bool](o_, objc.Sel("isLongLived"))
	return rv
}

// A Boolean value that indicates whether the operations that use this configuration are long-lived. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationconfiguration/2866225-longlived?language=objc
func (o_ OperationConfiguration) SetLongLived(value bool) {
	objc.Call[objc.Void](o_, objc.Sel("setLongLived:"), value)
}

// The maximum amount of time that a resource request can take. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationconfiguration/2866221-timeoutintervalforresource?language=objc
func (o_ OperationConfiguration) TimeoutIntervalForResource() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](o_, objc.Sel("timeoutIntervalForResource"))
	return rv
}

// The maximum amount of time that a resource request can take. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationconfiguration/2866221-timeoutintervalforresource?language=objc
func (o_ OperationConfiguration) SetTimeoutIntervalForResource(value foundation.TimeInterval) {
	objc.Call[objc.Void](o_, objc.Sel("setTimeoutIntervalForResource:"), value)
}
