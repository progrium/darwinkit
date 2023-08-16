// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// An object you use to synchronize access to Metal resources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlevent?language=objc
type PEvent interface {
	// optional
	Device() PDevice
	HasDevice() bool

	// optional
	SetLabel(value string)
	HasSetLabel() bool

	// optional
	Label() string
	HasLabel() bool
}

// A concrete type wrapper for the [PEvent] protocol.
type EventWrapper struct {
	objc.Object
}

func (e_ EventWrapper) HasDevice() bool {
	return e_.RespondsToSelector(objc.Sel("device"))
}

// The device object that created the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlevent/2966571-device?language=objc
func (e_ EventWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](e_, objc.Sel("device"))
	return rv
}

func (e_ EventWrapper) HasSetLabel() bool {
	return e_.RespondsToSelector(objc.Sel("setLabel:"))
}

// A string that identifies the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlevent/2966572-label?language=objc
func (e_ EventWrapper) SetLabel(value string) {
	objc.Call[objc.Void](e_, objc.Sel("setLabel:"), value)
}

func (e_ EventWrapper) HasLabel() bool {
	return e_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlevent/2966572-label?language=objc
func (e_ EventWrapper) Label() string {
	rv := objc.Call[string](e_, objc.Sel("label"))
	return rv
}
