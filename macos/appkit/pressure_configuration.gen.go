// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PressureConfiguration] class.
var PressureConfigurationClass = _PressureConfigurationClass{objc.GetClass("NSPressureConfiguration")}

type _PressureConfigurationClass struct {
	objc.Class
}

// An interface definition for the [PressureConfiguration] class.
type IPressureConfiguration interface {
	objc.IObject
	Set()
	PressureBehavior() PressureBehavior
}

// An encapsulation of the behavior and progression of a Force Touch trackpad as it responds to specific events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressureconfiguration?language=objc
type PressureConfiguration struct {
	objc.Object
}

func PressureConfigurationFrom(ptr unsafe.Pointer) PressureConfiguration {
	return PressureConfiguration{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PressureConfiguration) InitWithPressureBehavior(pressureBehavior PressureBehavior) PressureConfiguration {
	rv := objc.Call[PressureConfiguration](p_, objc.Sel("initWithPressureBehavior:"), pressureBehavior)
	return rv
}

// Initializes a pressure configuration object with a specified pressure behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressureconfiguration/1426883-initwithpressurebehavior?language=objc
func NewPressureConfigurationWithPressureBehavior(pressureBehavior PressureBehavior) PressureConfiguration {
	instance := PressureConfigurationClass.Alloc().InitWithPressureBehavior(pressureBehavior)
	instance.Autorelease()
	return instance
}

func (pc _PressureConfigurationClass) Alloc() PressureConfiguration {
	rv := objc.Call[PressureConfiguration](pc, objc.Sel("alloc"))
	return rv
}

func PressureConfiguration_Alloc() PressureConfiguration {
	return PressureConfigurationClass.Alloc()
}

func (pc _PressureConfigurationClass) New() PressureConfiguration {
	rv := objc.Call[PressureConfiguration](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPressureConfiguration() PressureConfiguration {
	return PressureConfigurationClass.New()
}

func (p_ PressureConfiguration) Init() PressureConfiguration {
	rv := objc.Call[PressureConfiguration](p_, objc.Sel("init"))
	return rv
}

// Changes the pressure configuration of the trackpad to the initialized pressure configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressureconfiguration/1426887-set?language=objc
func (p_ PressureConfiguration) Set() {
	objc.Call[objc.Void](p_, objc.Sel("set"))
}

// The pressure behavior of the pressure configuration object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressureconfiguration/1426889-pressurebehavior?language=objc
func (p_ PressureConfiguration) PressureBehavior() PressureBehavior {
	rv := objc.Call[PressureBehavior](p_, objc.Sel("pressureBehavior"))
	return rv
}
