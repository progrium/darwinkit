// AUTO-GENERATED CODE, DO NOT MODIFY

package coremediaio

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A protocol for objects that act as device sources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevicesource?language=objc
type PExtensionDeviceSource interface {
	// optional
	SetDevicePropertiesError(deviceProperties ExtensionDeviceProperties, outError foundation.Error) bool
	HasSetDevicePropertiesError() bool

	// optional
	DevicePropertiesForPropertiesError(properties foundation.Set, outError foundation.Error) IExtensionDeviceProperties
	HasDevicePropertiesForPropertiesError() bool

	// optional
	AvailableProperties() foundation.ISet
	HasAvailableProperties() bool
}

// A concrete type wrapper for the [PExtensionDeviceSource] protocol.
type ExtensionDeviceSourceWrapper struct {
	objc.Object
}

func (e_ ExtensionDeviceSourceWrapper) HasSetDevicePropertiesError() bool {
	return e_.RespondsToSelector(objc.Sel("setDeviceProperties:error:"))
}

// Sets the state of device properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevicesource/3915850-setdeviceproperties?language=objc
func (e_ ExtensionDeviceSourceWrapper) SetDevicePropertiesError(deviceProperties IExtensionDeviceProperties, outError foundation.IError) bool {
	rv := objc.Call[bool](e_, objc.Sel("setDeviceProperties:error:"), objc.Ptr(deviceProperties), objc.Ptr(outError))
	return rv
}

func (e_ ExtensionDeviceSourceWrapper) HasDevicePropertiesForPropertiesError() bool {
	return e_.RespondsToSelector(objc.Sel("devicePropertiesForProperties:error:"))
}

// Retrieves the state of device properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevicesource/3915849-devicepropertiesforproperties?language=objc
func (e_ ExtensionDeviceSourceWrapper) DevicePropertiesForPropertiesError(properties foundation.ISet, outError foundation.IError) ExtensionDeviceProperties {
	rv := objc.Call[ExtensionDeviceProperties](e_, objc.Sel("devicePropertiesForProperties:error:"), objc.Ptr(properties), objc.Ptr(outError))
	return rv
}

func (e_ ExtensionDeviceSourceWrapper) HasAvailableProperties() bool {
	return e_.RespondsToSelector(objc.Sel("availableProperties"))
}

// A set of available properties that a device provides. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevicesource/3915848-availableproperties?language=objc
func (e_ ExtensionDeviceSourceWrapper) AvailableProperties() foundation.Set {
	rv := objc.Call[foundation.Set](e_, objc.Sel("availableProperties"))
	return rv
}
