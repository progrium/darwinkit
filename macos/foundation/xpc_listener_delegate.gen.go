// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// The protocol that delegates to the XPC listener use to accept or reject new connections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpclistenerdelegate?language=objc
type PXPCListenerDelegate interface {
	// optional
	ListenerShouldAcceptNewConnection(listener XPCListener, newConnection XPCConnection) bool
	HasListenerShouldAcceptNewConnection() bool
}

// A delegate implementation builder for the [PXPCListenerDelegate] protocol.
type XPCListenerDelegate struct {
	_ListenerShouldAcceptNewConnection func(listener XPCListener, newConnection XPCConnection) bool
}

func (di *XPCListenerDelegate) HasListenerShouldAcceptNewConnection() bool {
	return di._ListenerShouldAcceptNewConnection != nil
}

// Accepts or rejects a new connection to the listener. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpclistenerdelegate/1410381-listener?language=objc
func (di *XPCListenerDelegate) SetListenerShouldAcceptNewConnection(f func(listener XPCListener, newConnection XPCConnection) bool) {
	di._ListenerShouldAcceptNewConnection = f
}

// Accepts or rejects a new connection to the listener. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpclistenerdelegate/1410381-listener?language=objc
func (di *XPCListenerDelegate) ListenerShouldAcceptNewConnection(listener XPCListener, newConnection XPCConnection) bool {
	return di._ListenerShouldAcceptNewConnection(listener, newConnection)
}

// A concrete type wrapper for the [PXPCListenerDelegate] protocol.
type XPCListenerDelegateWrapper struct {
	objc.Object
}

func (x_ XPCListenerDelegateWrapper) HasListenerShouldAcceptNewConnection() bool {
	return x_.RespondsToSelector(objc.Sel("listener:shouldAcceptNewConnection:"))
}

// Accepts or rejects a new connection to the listener. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpclistenerdelegate/1410381-listener?language=objc
func (x_ XPCListenerDelegateWrapper) ListenerShouldAcceptNewConnection(listener IXPCListener, newConnection IXPCConnection) bool {
	rv := objc.Call[bool](x_, objc.Sel("listener:shouldAcceptNewConnection:"), objc.Ptr(listener), objc.Ptr(newConnection))
	return rv
}
