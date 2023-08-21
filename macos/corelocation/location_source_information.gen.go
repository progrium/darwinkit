// AUTO-GENERATED CODE, DO NOT MODIFY

package corelocation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LocationSourceInformation] class.
var LocationSourceInformationClass = _LocationSourceInformationClass{objc.GetClass("CLLocationSourceInformation")}

type _LocationSourceInformationClass struct {
	objc.Class
}

// An interface definition for the [LocationSourceInformation] class.
type ILocationSourceInformation interface {
	objc.IObject
	IsProducedByAccessory() bool
	IsSimulatedBySoftware() bool
}

// Information about the source that provides a location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationsourceinformation?language=objc
type LocationSourceInformation struct {
	objc.Object
}

func LocationSourceInformationFrom(ptr unsafe.Pointer) LocationSourceInformation {
	return LocationSourceInformation{
		Object: objc.ObjectFrom(ptr),
	}
}

func (l_ LocationSourceInformation) InitWithSoftwareSimulationStateAndExternalAccessoryState(isSoftware bool, isAccessory bool) LocationSourceInformation {
	rv := objc.Call[LocationSourceInformation](l_, objc.Sel("initWithSoftwareSimulationState:andExternalAccessoryState:"), isSoftware, isAccessory)
	return rv
}

// Creates an instance of location source information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationsourceinformation/3861805-initwithsoftwaresimulationstate?language=objc
func NewLocationSourceInformationWithSoftwareSimulationStateAndExternalAccessoryState(isSoftware bool, isAccessory bool) LocationSourceInformation {
	instance := LocationSourceInformationClass.Alloc().InitWithSoftwareSimulationStateAndExternalAccessoryState(isSoftware, isAccessory)
	instance.Autorelease()
	return instance
}

func (lc _LocationSourceInformationClass) Alloc() LocationSourceInformation {
	rv := objc.Call[LocationSourceInformation](lc, objc.Sel("alloc"))
	return rv
}

func LocationSourceInformation_Alloc() LocationSourceInformation {
	return LocationSourceInformationClass.Alloc()
}

func (lc _LocationSourceInformationClass) New() LocationSourceInformation {
	rv := objc.Call[LocationSourceInformation](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLocationSourceInformation() LocationSourceInformation {
	return LocationSourceInformationClass.New()
}

func (l_ LocationSourceInformation) Init() LocationSourceInformation {
	rv := objc.Call[LocationSourceInformation](l_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the system receives the location from an external accessory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationsourceinformation/3861806-isproducedbyaccessory?language=objc
func (l_ LocationSourceInformation) IsProducedByAccessory() bool {
	rv := objc.Call[bool](l_, objc.Sel("isProducedByAccessory"))
	return rv
}

// A Boolean value that indicates whether the system generates the location using on-device software simulation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocationsourceinformation/3861807-issimulatedbysoftware?language=objc
func (l_ LocationSourceInformation) IsSimulatedBySoftware() bool {
	rv := objc.Call[bool](l_, objc.Sel("isSimulatedBySoftware"))
	return rv
}
