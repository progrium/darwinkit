// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// Methods for creating new proxy objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcproxycreating?language=objc
type PXPCProxyCreating interface {
	// optional
	SynchronousRemoteObjectProxyWithErrorHandler(handler func(error Error)) objc.IObject
	HasSynchronousRemoteObjectProxyWithErrorHandler() bool

	// optional
	RemoteObjectProxy() objc.IObject
	HasRemoteObjectProxy() bool

	// optional
	RemoteObjectProxyWithErrorHandler(handler func(error Error)) objc.IObject
	HasRemoteObjectProxyWithErrorHandler() bool
}

// A concrete type wrapper for the [PXPCProxyCreating] protocol.
type XPCProxyCreatingWrapper struct {
	objc.Object
}

func (x_ XPCProxyCreatingWrapper) HasSynchronousRemoteObjectProxyWithErrorHandler() bool {
	return x_.RespondsToSelector(objc.Sel("synchronousRemoteObjectProxyWithErrorHandler:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcproxycreating/2879411-synchronousremoteobjectproxywith?language=objc
func (x_ XPCProxyCreatingWrapper) SynchronousRemoteObjectProxyWithErrorHandler(handler func(error Error)) objc.Object {
	rv := objc.Call[objc.Object](x_, objc.Sel("synchronousRemoteObjectProxyWithErrorHandler:"), handler)
	return rv
}

func (x_ XPCProxyCreatingWrapper) HasRemoteObjectProxy() bool {
	return x_.RespondsToSelector(objc.Sel("remoteObjectProxy"))
}

// Returns a proxy object with no error handling block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcproxycreating/1418403-remoteobjectproxy?language=objc
func (x_ XPCProxyCreatingWrapper) RemoteObjectProxy() objc.Object {
	rv := objc.Call[objc.Object](x_, objc.Sel("remoteObjectProxy"))
	return rv
}

func (x_ XPCProxyCreatingWrapper) HasRemoteObjectProxyWithErrorHandler() bool {
	return x_.RespondsToSelector(objc.Sel("remoteObjectProxyWithErrorHandler:"))
}

// Returns a proxy object that invokes the error handling block if an error occurs on the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcproxycreating/1415611-remoteobjectproxywitherrorhandle?language=objc
func (x_ XPCProxyCreatingWrapper) RemoteObjectProxyWithErrorHandler(handler func(error Error)) objc.Object {
	rv := objc.Call[objc.Object](x_, objc.Sel("remoteObjectProxyWithErrorHandler:"), handler)
	return rv
}
