// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines custom boundaries for a GPU frame capture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturescope?language=objc
type PCaptureScope interface {
	// optional
	BeginScope()
	HasBeginScope() bool

	// optional
	EndScope()
	HasEndScope() bool

	// optional
	CommandQueue() PCommandQueue
	HasCommandQueue() bool

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

// A concrete type wrapper for the [PCaptureScope] protocol.
type CaptureScopeWrapper struct {
	objc.Object
}

func (c_ CaptureScopeWrapper) HasBeginScope() bool {
	return c_.RespondsToSelector(objc.Sel("beginScope"))
}

// Tells Metal to begin recording command information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturescope/2869571-beginscope?language=objc
func (c_ CaptureScopeWrapper) BeginScope() {
	objc.Call[objc.Void](c_, objc.Sel("beginScope"))
}

func (c_ CaptureScopeWrapper) HasEndScope() bool {
	return c_.RespondsToSelector(objc.Sel("endScope"))
}

// Tells Metal to stop recording command information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturescope/2869574-endscope?language=objc
func (c_ CaptureScopeWrapper) EndScope() {
	objc.Call[objc.Void](c_, objc.Sel("endScope"))
}

func (c_ CaptureScopeWrapper) HasCommandQueue() bool {
	return c_.RespondsToSelector(objc.Sel("commandQueue"))
}

// The command queue that this capture scope uses to limit which commands are recorded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturescope/2869573-commandqueue?language=objc
func (c_ CaptureScopeWrapper) CommandQueue() CommandQueueWrapper {
	rv := objc.Call[CommandQueueWrapper](c_, objc.Sel("commandQueue"))
	return rv
}

func (c_ CaptureScopeWrapper) HasDevice() bool {
	return c_.RespondsToSelector(objc.Sel("device"))
}

// The device object from which you created the capture scope. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturescope/2869575-device?language=objc
func (c_ CaptureScopeWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](c_, objc.Sel("device"))
	return rv
}

func (c_ CaptureScopeWrapper) HasSetLabel() bool {
	return c_.RespondsToSelector(objc.Sel("setLabel:"))
}

// A string that identifies the capture scope. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturescope/2869572-label?language=objc
func (c_ CaptureScopeWrapper) SetLabel(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setLabel:"), value)
}

func (c_ CaptureScopeWrapper) HasLabel() bool {
	return c_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the capture scope. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturescope/2869572-label?language=objc
func (c_ CaptureScopeWrapper) Label() string {
	rv := objc.Call[string](c_, objc.Sel("label"))
	return rv
}
