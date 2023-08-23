// AUTO-GENERATED CODE, DO NOT MODIFY

package coremediaio

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ExtensionDevice] class.
var ExtensionDeviceClass = _ExtensionDeviceClass{objc.GetClass("CMIOExtensionDevice")}

type _ExtensionDeviceClass struct {
	objc.Class
}

// An interface definition for the [ExtensionDevice] class.
type IExtensionDevice interface {
	objc.IObject
	NotifyPropertiesChanged(propertyStates map[ExtensionProperty]IExtensionPropertyState)
	RemoveStreamError(stream IExtensionStream, outError foundation.IError) bool
	AddStreamError(stream IExtensionStream, outError foundation.IError) bool
	DeviceID() foundation.UUID
	Streams() []ExtensionStream
	Source() ExtensionDeviceSourceWrapper
	LocalizedName() string
	LegacyDeviceID() string
}

// An object that represents a physical or virtual device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice?language=objc
type ExtensionDevice struct {
	objc.Object
}

func ExtensionDeviceFrom(ptr unsafe.Pointer) ExtensionDevice {
	return ExtensionDevice{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _ExtensionDeviceClass) Alloc() ExtensionDevice {
	rv := objc.Call[ExtensionDevice](ec, objc.Sel("alloc"))
	return rv
}

func ExtensionDevice_Alloc() ExtensionDevice {
	return ExtensionDeviceClass.Alloc()
}

func (ec _ExtensionDeviceClass) New() ExtensionDevice {
	rv := objc.Call[ExtensionDevice](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExtensionDevice() ExtensionDevice {
	return ExtensionDeviceClass.New()
}

func (e_ ExtensionDevice) Init() ExtensionDevice {
	rv := objc.Call[ExtensionDevice](e_, objc.Sel("init"))
	return rv
}

// Notifies clients of property changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice/3915834-notifypropertieschanged?language=objc
func (e_ ExtensionDevice) NotifyPropertiesChanged(propertyStates map[ExtensionProperty]IExtensionPropertyState) {
	objc.Call[objc.Void](e_, objc.Sel("notifyPropertiesChanged:"), propertyStates)
}

// Removes a stream from the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice/3915835-removestream?language=objc
func (e_ ExtensionDevice) RemoveStreamError(stream IExtensionStream, outError foundation.IError) bool {
	rv := objc.Call[bool](e_, objc.Sel("removeStream:error:"), objc.Ptr(stream), objc.Ptr(outError))
	return rv
}

// Adds a stream to a device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice/3915826-addstream?language=objc
func (e_ ExtensionDevice) AddStreamError(stream IExtensionStream, outError foundation.IError) bool {
	rv := objc.Call[bool](e_, objc.Sel("addStream:error:"), objc.Ptr(stream), objc.Ptr(outError))
	return rv
}

// A universally unique device identifier value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice/3915827-deviceid?language=objc
func (e_ ExtensionDevice) DeviceID() foundation.UUID {
	rv := objc.Call[foundation.UUID](e_, objc.Sel("deviceID"))
	return rv
}

// An array of media streams attached to this device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice/3915837-streams?language=objc
func (e_ ExtensionDevice) Streams() []ExtensionStream {
	rv := objc.Call[[]ExtensionStream](e_, objc.Sel("streams"))
	return rv
}

// A source object for a device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice/3915836-source?language=objc
func (e_ ExtensionDevice) Source() ExtensionDeviceSourceWrapper {
	rv := objc.Call[ExtensionDeviceSourceWrapper](e_, objc.Sel("source"))
	return rv
}

// A localized name for a device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice/3915833-localizedname?language=objc
func (e_ ExtensionDevice) LocalizedName() string {
	rv := objc.Call[string](e_, objc.Sel("localizedName"))
	return rv
}

// A legacy device identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice/3915832-legacydeviceid?language=objc
func (e_ ExtensionDevice) LegacyDeviceID() string {
	rv := objc.Call[string](e_, objc.Sel("legacyDeviceID"))
	return rv
}
