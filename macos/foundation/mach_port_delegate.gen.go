// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// An interface for handling incoming Mach messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmachportdelegate?language=objc
type PMachPortDelegate interface {
	// optional
	HandleMachMessage(msg unsafe.Pointer)
	HasHandleMachMessage() bool
}

// A delegate implementation builder for the [PMachPortDelegate] protocol.
type MachPortDelegate struct {
	_HandleMachMessage func(msg unsafe.Pointer)
}

func (di *MachPortDelegate) HasHandleMachMessage() bool {
	return di._HandleMachMessage != nil
}

// Process an incoming Mach message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmachportdelegate/1399509-handlemachmessage?language=objc
func (di *MachPortDelegate) SetHandleMachMessage(f func(msg unsafe.Pointer)) {
	di._HandleMachMessage = f
}

// Process an incoming Mach message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmachportdelegate/1399509-handlemachmessage?language=objc
func (di *MachPortDelegate) HandleMachMessage(msg unsafe.Pointer) {
	di._HandleMachMessage(msg)
}

// A concrete type wrapper for the [PMachPortDelegate] protocol.
type MachPortDelegateWrapper struct {
	objc.Object
}

func (m_ MachPortDelegateWrapper) HasHandleMachMessage() bool {
	return m_.RespondsToSelector(objc.Sel("handleMachMessage:"))
}

// Process an incoming Mach message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmachportdelegate/1399509-handlemachmessage?language=objc
func (m_ MachPortDelegateWrapper) HandleMachMessage(msg unsafe.Pointer) {
	objc.Call[objc.Void](m_, objc.Sel("handleMachMessage:"), msg)
}
