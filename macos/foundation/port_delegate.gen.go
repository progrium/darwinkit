// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// An interface for handling incoming messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsportdelegate?language=objc
type PPortDelegate interface {
	// optional
	HandlePortMessage(message PortMessage)
	HasHandlePortMessage() bool
}

// A delegate implementation builder for the [PPortDelegate] protocol.
type PortDelegate struct {
	_HandlePortMessage func(message PortMessage)
}

func (di *PortDelegate) HasHandlePortMessage() bool {
	return di._HandlePortMessage != nil
}

// Processes a given incoming message on the port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsportdelegate/1399513-handleportmessage?language=objc
func (di *PortDelegate) SetHandlePortMessage(f func(message PortMessage)) {
	di._HandlePortMessage = f
}

// Processes a given incoming message on the port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsportdelegate/1399513-handleportmessage?language=objc
func (di *PortDelegate) HandlePortMessage(message PortMessage) {
	di._HandlePortMessage(message)
}

// A concrete type wrapper for the [PPortDelegate] protocol.
type PortDelegateWrapper struct {
	objc.Object
}

func (p_ PortDelegateWrapper) HasHandlePortMessage() bool {
	return p_.RespondsToSelector(objc.Sel("handlePortMessage:"))
}

// Processes a given incoming message on the port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsportdelegate/1399513-handleportmessage?language=objc
func (p_ PortDelegateWrapper) HandlePortMessage(message IPortMessage) {
	objc.Call[objc.Void](p_, objc.Sel("handlePortMessage:"), objc.Ptr(message))
}
