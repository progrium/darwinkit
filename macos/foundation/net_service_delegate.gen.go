// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// The interface a net service uses to inform its delegate about the state of the service it offers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate?language=objc
type PNetServiceDelegate interface {
	// optional
	NetServiceDidResolveAddress(sender NetService)
	HasNetServiceDidResolveAddress() bool

	// optional
	NetServiceWillResolve(sender NetService)
	HasNetServiceWillResolve() bool

	// optional
	NetServiceDidPublish(sender NetService)
	HasNetServiceDidPublish() bool

	// optional
	NetServiceWillPublish(sender NetService)
	HasNetServiceWillPublish() bool

	// optional
	NetServiceDidStop(sender NetService)
	HasNetServiceDidStop() bool

	// optional
	NetServiceDidUpdateTXTRecordData(sender NetService, data []byte)
	HasNetServiceDidUpdateTXTRecordData() bool
}

// A delegate implementation builder for the [PNetServiceDelegate] protocol.
type NetServiceDelegate struct {
	_NetServiceDidResolveAddress      func(sender NetService)
	_NetServiceWillResolve            func(sender NetService)
	_NetServiceDidPublish             func(sender NetService)
	_NetServiceWillPublish            func(sender NetService)
	_NetServiceDidStop                func(sender NetService)
	_NetServiceDidUpdateTXTRecordData func(sender NetService, data []byte)
}

func (di *NetServiceDelegate) HasNetServiceDidResolveAddress() bool {
	return di._NetServiceDidResolveAddress != nil
}

// Informs the delegate that the address for a given service was resolved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1408457-netservicedidresolveaddress?language=objc
func (di *NetServiceDelegate) SetNetServiceDidResolveAddress(f func(sender NetService)) {
	di._NetServiceDidResolveAddress = f
}

// Informs the delegate that the address for a given service was resolved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1408457-netservicedidresolveaddress?language=objc
func (di *NetServiceDelegate) NetServiceDidResolveAddress(sender NetService) {
	di._NetServiceDidResolveAddress(sender)
}
func (di *NetServiceDelegate) HasNetServiceWillResolve() bool {
	return di._NetServiceWillResolve != nil
}

// Notifies the delegate that the network is ready to resolve the service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1416022-netservicewillresolve?language=objc
func (di *NetServiceDelegate) SetNetServiceWillResolve(f func(sender NetService)) {
	di._NetServiceWillResolve = f
}

// Notifies the delegate that the network is ready to resolve the service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1416022-netservicewillresolve?language=objc
func (di *NetServiceDelegate) NetServiceWillResolve(sender NetService) {
	di._NetServiceWillResolve(sender)
}
func (di *NetServiceDelegate) HasNetServiceDidPublish() bool {
	return di._NetServiceDidPublish != nil
}

// Notifies the delegate that a service was successfully published. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1416802-netservicedidpublish?language=objc
func (di *NetServiceDelegate) SetNetServiceDidPublish(f func(sender NetService)) {
	di._NetServiceDidPublish = f
}

// Notifies the delegate that a service was successfully published. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1416802-netservicedidpublish?language=objc
func (di *NetServiceDelegate) NetServiceDidPublish(sender NetService) {
	di._NetServiceDidPublish(sender)
}
func (di *NetServiceDelegate) HasNetServiceWillPublish() bool {
	return di._NetServiceWillPublish != nil
}

// Notifies the delegate that the network is ready to publish the service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1414277-netservicewillpublish?language=objc
func (di *NetServiceDelegate) SetNetServiceWillPublish(f func(sender NetService)) {
	di._NetServiceWillPublish = f
}

// Notifies the delegate that the network is ready to publish the service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1414277-netservicewillpublish?language=objc
func (di *NetServiceDelegate) NetServiceWillPublish(sender NetService) {
	di._NetServiceWillPublish(sender)
}
func (di *NetServiceDelegate) HasNetServiceDidStop() bool {
	return di._NetServiceDidStop != nil
}

// Informs the delegate that a publish or resolveWithTimeout: request was stopped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1409726-netservicedidstop?language=objc
func (di *NetServiceDelegate) SetNetServiceDidStop(f func(sender NetService)) {
	di._NetServiceDidStop = f
}

// Informs the delegate that a publish or resolveWithTimeout: request was stopped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1409726-netservicedidstop?language=objc
func (di *NetServiceDelegate) NetServiceDidStop(sender NetService) {
	di._NetServiceDidStop(sender)
}
func (di *NetServiceDelegate) HasNetServiceDidUpdateTXTRecordData() bool {
	return di._NetServiceDidUpdateTXTRecordData != nil
}

// Notifies the delegate that the TXT record for a given service has been updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1413199-netservice?language=objc
func (di *NetServiceDelegate) SetNetServiceDidUpdateTXTRecordData(f func(sender NetService, data []byte)) {
	di._NetServiceDidUpdateTXTRecordData = f
}

// Notifies the delegate that the TXT record for a given service has been updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1413199-netservice?language=objc
func (di *NetServiceDelegate) NetServiceDidUpdateTXTRecordData(sender NetService, data []byte) {
	di._NetServiceDidUpdateTXTRecordData(sender, data)
}

// A concrete type wrapper for the [PNetServiceDelegate] protocol.
type NetServiceDelegateWrapper struct {
	objc.Object
}

func (n_ NetServiceDelegateWrapper) HasNetServiceDidResolveAddress() bool {
	return n_.RespondsToSelector(objc.Sel("netServiceDidResolveAddress:"))
}

// Informs the delegate that the address for a given service was resolved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1408457-netservicedidresolveaddress?language=objc
func (n_ NetServiceDelegateWrapper) NetServiceDidResolveAddress(sender INetService) {
	objc.Call[objc.Void](n_, objc.Sel("netServiceDidResolveAddress:"), objc.Ptr(sender))
}

func (n_ NetServiceDelegateWrapper) HasNetServiceWillResolve() bool {
	return n_.RespondsToSelector(objc.Sel("netServiceWillResolve:"))
}

// Notifies the delegate that the network is ready to resolve the service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1416022-netservicewillresolve?language=objc
func (n_ NetServiceDelegateWrapper) NetServiceWillResolve(sender INetService) {
	objc.Call[objc.Void](n_, objc.Sel("netServiceWillResolve:"), objc.Ptr(sender))
}

func (n_ NetServiceDelegateWrapper) HasNetServiceDidPublish() bool {
	return n_.RespondsToSelector(objc.Sel("netServiceDidPublish:"))
}

// Notifies the delegate that a service was successfully published. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1416802-netservicedidpublish?language=objc
func (n_ NetServiceDelegateWrapper) NetServiceDidPublish(sender INetService) {
	objc.Call[objc.Void](n_, objc.Sel("netServiceDidPublish:"), objc.Ptr(sender))
}

func (n_ NetServiceDelegateWrapper) HasNetServiceWillPublish() bool {
	return n_.RespondsToSelector(objc.Sel("netServiceWillPublish:"))
}

// Notifies the delegate that the network is ready to publish the service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1414277-netservicewillpublish?language=objc
func (n_ NetServiceDelegateWrapper) NetServiceWillPublish(sender INetService) {
	objc.Call[objc.Void](n_, objc.Sel("netServiceWillPublish:"), objc.Ptr(sender))
}

func (n_ NetServiceDelegateWrapper) HasNetServiceDidStop() bool {
	return n_.RespondsToSelector(objc.Sel("netServiceDidStop:"))
}

// Informs the delegate that a publish or resolveWithTimeout: request was stopped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1409726-netservicedidstop?language=objc
func (n_ NetServiceDelegateWrapper) NetServiceDidStop(sender INetService) {
	objc.Call[objc.Void](n_, objc.Sel("netServiceDidStop:"), objc.Ptr(sender))
}

func (n_ NetServiceDelegateWrapper) HasNetServiceDidUpdateTXTRecordData() bool {
	return n_.RespondsToSelector(objc.Sel("netService:didUpdateTXTRecordData:"))
}

// Notifies the delegate that the TXT record for a given service has been updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1413199-netservice?language=objc
func (n_ NetServiceDelegateWrapper) NetServiceDidUpdateTXTRecordData(sender INetService, data []byte) {
	objc.Call[objc.Void](n_, objc.Sel("netService:didUpdateTXTRecordData:"), objc.Ptr(sender), data)
}
