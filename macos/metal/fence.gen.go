// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// An object that can capture, track, and manage resource dependencies across command encoders. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfence?language=objc
type PFence interface {
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

// A concrete type wrapper for the [PFence] protocol.
type FenceWrapper struct {
	objc.Object
}

func (f_ FenceWrapper) HasDevice() bool {
	return f_.RespondsToSelector(objc.Sel("device"))
}

// The device object that created the fence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfence/1648532-device?language=objc
func (f_ FenceWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](f_, objc.Sel("device"))
	return rv
}

func (f_ FenceWrapper) HasSetLabel() bool {
	return f_.RespondsToSelector(objc.Sel("setLabel:"))
}

// A string that identifies the fence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfence/1648531-label?language=objc
func (f_ FenceWrapper) SetLabel(value string) {
	objc.Call[objc.Void](f_, objc.Sel("setLabel:"), value)
}

func (f_ FenceWrapper) HasLabel() bool {
	return f_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the fence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfence/1648531-label?language=objc
func (f_ FenceWrapper) Label() string {
	rv := objc.Call[string](f_, objc.Sel("label"))
	return rv
}
