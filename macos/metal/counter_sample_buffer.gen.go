// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A specialized memory buffer that stores a GPU’s counter set data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcountersamplebuffer?language=objc
type PCounterSampleBuffer interface {
	// optional
	ResolveCounterRange(range_ foundation.Range) []byte
	HasResolveCounterRange() bool

	// optional
	Device() PDevice
	HasDevice() bool

	// optional
	SampleCount() uint
	HasSampleCount() bool

	// optional
	Label() string
	HasLabel() bool
}

// A concrete type wrapper for the [PCounterSampleBuffer] protocol.
type CounterSampleBufferWrapper struct {
	objc.Object
}

func (c_ CounterSampleBufferWrapper) HasResolveCounterRange() bool {
	return c_.RespondsToSelector(objc.Sel("resolveCounterRange:"))
}

// Transforms samples of a GPU’s counter set from the driver’s internal format to a standard Metal data structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcountersamplebuffer/3194377-resolvecounterrange?language=objc
func (c_ CounterSampleBufferWrapper) ResolveCounterRange(range_ foundation.Range) []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("resolveCounterRange:"), range_)
	return rv
}

func (c_ CounterSampleBufferWrapper) HasDevice() bool {
	return c_.RespondsToSelector(objc.Sel("device"))
}

// The GPU device instance that owns the counter sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcountersamplebuffer/3081726-device?language=objc
func (c_ CounterSampleBufferWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](c_, objc.Sel("device"))
	return rv
}

func (c_ CounterSampleBufferWrapper) HasSampleCount() bool {
	return c_.RespondsToSelector(objc.Sel("sampleCount"))
}

// The number of samples in the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcountersamplebuffer/3081728-samplecount?language=objc
func (c_ CounterSampleBufferWrapper) SampleCount() uint {
	rv := objc.Call[uint](c_, objc.Sel("sampleCount"))
	return rv
}

func (c_ CounterSampleBufferWrapper) HasLabel() bool {
	return c_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the counter sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcountersamplebuffer/3081727-label?language=objc
func (c_ CounterSampleBufferWrapper) Label() string {
	rv := objc.Call[string](c_, objc.Sel("label"))
	return rv
}
