// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// An object you use to synchronize access to Metal resources across multiple CPUs, GPUs, and processes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedevent?language=objc
type PSharedEvent interface {
	// optional
	NotifyListenerAtValueBlock(listener SharedEventListener, value uint64, block SharedEventNotificationBlock)
	HasNotifyListenerAtValueBlock() bool

	// optional
	NewSharedEventHandle() ISharedEventHandle
	HasNewSharedEventHandle() bool

	// optional
	SetSignaledValue(value uint64)
	HasSetSignaledValue() bool

	// optional
	SignaledValue() uint64
	HasSignaledValue() bool
}

// A concrete type wrapper for the [PSharedEvent] protocol.
type SharedEventWrapper struct {
	objc.Object
}

func (s_ SharedEventWrapper) HasNotifyListenerAtValueBlock() bool {
	return s_.RespondsToSelector(objc.Sel("notifyListener:atValue:block:"))
}

// Schedules a notification handler to be called after the shareable event’s signal value equals or exceeds a given value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedevent/2966574-notifylistener?language=objc
func (s_ SharedEventWrapper) NotifyListenerAtValueBlock(listener ISharedEventListener, value uint64, block SharedEventNotificationBlock) {
	objc.Call[objc.Void](s_, objc.Sel("notifyListener:atValue:block:"), objc.Ptr(listener), value, block)
}

func (s_ SharedEventWrapper) HasNewSharedEventHandle() bool {
	return s_.RespondsToSelector(objc.Sel("newSharedEventHandle"))
}

// Creates a new shareable event handle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedevent/2981025-newsharedeventhandle?language=objc
func (s_ SharedEventWrapper) NewSharedEventHandle() SharedEventHandle {
	rv := objc.Call[SharedEventHandle](s_, objc.Sel("newSharedEventHandle"))
	rv.Autorelease()
	return rv
}

func (s_ SharedEventWrapper) HasSetSignaledValue() bool {
	return s_.RespondsToSelector(objc.Sel("setSignaledValue:"))
}

// The current signal value for the shareable event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedevent/2966575-signaledvalue?language=objc
func (s_ SharedEventWrapper) SetSignaledValue(value uint64) {
	objc.Call[objc.Void](s_, objc.Sel("setSignaledValue:"), value)
}

func (s_ SharedEventWrapper) HasSignaledValue() bool {
	return s_.RespondsToSelector(objc.Sel("signaledValue"))
}

// The current signal value for the shareable event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedevent/2966575-signaledvalue?language=objc
func (s_ SharedEventWrapper) SignaledValue() uint64 {
	rv := objc.Call[uint64](s_, objc.Sel("signaledValue"))
	return rv
}
