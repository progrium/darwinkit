// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// An object that defines how a texture should be sampled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerstate?language=objc
type PSamplerState interface {
	// optional
	Device() PDevice
	HasDevice() bool

	// optional
	Label() string
	HasLabel() bool
}

// A concrete type wrapper for the [PSamplerState] protocol.
type SamplerStateWrapper struct {
	objc.Object
}

func (s_ SamplerStateWrapper) HasDevice() bool {
	return s_.RespondsToSelector(objc.Sel("device"))
}

// The device object that created the sampler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerstate/1515871-device?language=objc
func (s_ SamplerStateWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](s_, objc.Sel("device"))
	return rv
}

func (s_ SamplerStateWrapper) HasLabel() bool {
	return s_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the sampler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerstate/1516329-label?language=objc
func (s_ SamplerStateWrapper) Label() string {
	rv := objc.Call[string](s_, objc.Sel("label"))
	return rv
}
