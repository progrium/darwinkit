// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// An interface that enables the setting of a Metal device for unarchived objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdeviceprovider?language=objc
type PDeviceProvider interface {
	// optional
	MpsMTLDevice() metal.PDevice
	HasMpsMTLDevice() bool
}

// A concrete type wrapper for the [PDeviceProvider] protocol.
type DeviceProviderWrapper struct {
	objc.Object
}

func (d_ DeviceProviderWrapper) HasMpsMTLDevice() bool {
	return d_.RespondsToSelector(objc.Sel("mpsMTLDevice"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdeviceprovider/2875211-mpsmtldevice?language=objc
func (d_ DeviceProviderWrapper) MpsMTLDevice() metal.DeviceWrapper {
	rv := objc.Call[metal.DeviceWrapper](d_, objc.Sel("mpsMTLDevice"))
	return rv
}
