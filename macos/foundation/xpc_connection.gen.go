// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [XPCConnection] class.
var XPCConnectionClass = _XPCConnectionClass{objc.GetClass("NSXPCConnection")}

type _XPCConnectionClass struct {
	objc.Class
}

// An interface definition for the [XPCConnection] class.
type IXPCConnection interface {
	objc.IObject
	SynchronousRemoteObjectProxyWithErrorHandler(handler func(error Error)) objc.Object
	Suspend()
	ScheduleSendBarrierBlock(block func())
	RemoteObjectProxyWithErrorHandler(handler func(error Error)) objc.Object
	Resume()
	Invalidate()
	InterruptionHandler() func()
	SetInterruptionHandler(value func())
	ExportedObject() objc.Object
	SetExportedObject(value objc.IObject)
	RemoteObjectProxy() objc.Object
	Endpoint() XPCListenerEndpoint
	RemoteObjectInterface() XPCInterface
	SetRemoteObjectInterface(value IXPCInterface)
	ServiceName() string
	ExportedInterface() XPCInterface
	SetExportedInterface(value IXPCInterface)
	InvalidationHandler() func()
	SetInvalidationHandler(value func())
}

// A bidirectional communication channel between two processes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection?language=objc
type XPCConnection struct {
	objc.Object
}

func XPCConnectionFrom(ptr unsafe.Pointer) XPCConnection {
	return XPCConnection{
		Object: objc.ObjectFrom(ptr),
	}
}

func (x_ XPCConnection) InitWithListenerEndpoint(endpoint IXPCListenerEndpoint) XPCConnection {
	rv := objc.Call[XPCConnection](x_, objc.Sel("initWithListenerEndpoint:"), objc.Ptr(endpoint))
	return rv
}

// Initializes an NSXPCConnection object to connect to an NSXPCListener object in another process, identified by an NSXPCListenerEndpoint object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1416298-initwithlistenerendpoint?language=objc
func XPCConnection_InitWithListenerEndpoint(endpoint IXPCListenerEndpoint) XPCConnection {
	return XPCConnectionClass.Alloc().InitWithListenerEndpoint(endpoint)
}

func (x_ XPCConnection) InitWithMachServiceNameOptions(name string, options XPCConnectionOptions) XPCConnection {
	rv := objc.Call[XPCConnection](x_, objc.Sel("initWithMachServiceName:options:"), name, options)
	return rv
}

// Initializes an NSXPCConnection object to connect to a LaunchAgent or LaunchDaemon with a name advertised in a launchd.plist. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1418074-initwithmachservicename?language=objc
func XPCConnection_InitWithMachServiceNameOptions(name string, options XPCConnectionOptions) XPCConnection {
	return XPCConnectionClass.Alloc().InitWithMachServiceNameOptions(name, options)
}

func (x_ XPCConnection) InitWithServiceName(serviceName string) XPCConnection {
	rv := objc.Call[XPCConnection](x_, objc.Sel("initWithServiceName:"), serviceName)
	return rv
}

// Initializes an NSXPCConnection object to connect to an NSXPCListener object in an XPC service, identified by a service name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1416370-initwithservicename?language=objc
func XPCConnection_InitWithServiceName(serviceName string) XPCConnection {
	return XPCConnectionClass.Alloc().InitWithServiceName(serviceName)
}

func (xc _XPCConnectionClass) Alloc() XPCConnection {
	rv := objc.Call[XPCConnection](xc, objc.Sel("alloc"))
	return rv
}

func XPCConnection_Alloc() XPCConnection {
	return XPCConnectionClass.Alloc()
}

func (xc _XPCConnectionClass) New() XPCConnection {
	rv := objc.Call[XPCConnection](xc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewXPCConnection() XPCConnection {
	return XPCConnectionClass.New()
}

func (x_ XPCConnection) Init() XPCConnection {
	rv := objc.Call[XPCConnection](x_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/2879410-synchronousremoteobjectproxywith?language=objc
func (x_ XPCConnection) SynchronousRemoteObjectProxyWithErrorHandler(handler func(error Error)) objc.Object {
	rv := objc.Call[objc.Object](x_, objc.Sel("synchronousRemoteObjectProxyWithErrorHandler:"), handler)
	return rv
}

// Suspends the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1410778-suspend?language=objc
func (x_ XPCConnection) Suspend() {
	objc.Call[objc.Void](x_, objc.Sel("suspend"))
}

// Add a barrier block to execute on the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/3172583-schedulesendbarrierblock?language=objc
func (x_ XPCConnection) ScheduleSendBarrierBlock(block func()) {
	objc.Call[objc.Void](x_, objc.Sel("scheduleSendBarrierBlock:"), block)
}

// Returns the current connection, in the context of a call to a method on your exported object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/3172582-currentconnection?language=objc
func (xc _XPCConnectionClass) CurrentConnection() XPCConnection {
	rv := objc.Call[XPCConnection](xc, objc.Sel("currentConnection"))
	return rv
}

// Returns the current connection, in the context of a call to a method on your exported object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/3172582-currentconnection?language=objc
func XPCConnection_CurrentConnection() XPCConnection {
	return XPCConnectionClass.CurrentConnection()
}

// Returns a proxy for the remote object (that is, the object exported from the other side of this connection) with the specified error handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1407905-remoteobjectproxywitherrorhandle?language=objc
func (x_ XPCConnection) RemoteObjectProxyWithErrorHandler(handler func(error Error)) objc.Object {
	rv := objc.Call[objc.Object](x_, objc.Sel("remoteObjectProxyWithErrorHandler:"), handler)
	return rv
}

// Starts or resumes handling of messages on a connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1408677-resume?language=objc
func (x_ XPCConnection) Resume() {
	objc.Call[objc.Void](x_, objc.Sel("resume"))
}

// Invalidates the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1412618-invalidate?language=objc
func (x_ XPCConnection) Invalidate() {
	objc.Call[objc.Void](x_, objc.Sel("invalidate"))
}

// An interruption handler that is called if the remote process exits or crashes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1407915-interruptionhandler?language=objc
func (x_ XPCConnection) InterruptionHandler() func() {
	rv := objc.Call[func()](x_, objc.Sel("interruptionHandler"))
	return rv
}

// An interruption handler that is called if the remote process exits or crashes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1407915-interruptionhandler?language=objc
func (x_ XPCConnection) SetInterruptionHandler(value func()) {
	objc.Call[objc.Void](x_, objc.Sel("setInterruptionHandler:"), value)
}

// An exported object for the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1412016-exportedobject?language=objc
func (x_ XPCConnection) ExportedObject() objc.Object {
	rv := objc.Call[objc.Object](x_, objc.Sel("exportedObject"))
	return rv
}

// An exported object for the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1412016-exportedobject?language=objc
func (x_ XPCConnection) SetExportedObject(value objc.IObject) {
	objc.Call[objc.Void](x_, objc.Sel("setExportedObject:"), value)
}

// Returns a proxy for the remote object (that is, the exportedObject from the other side of this connection). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1411031-remoteobjectproxy?language=objc
func (x_ XPCConnection) RemoteObjectProxy() objc.Object {
	rv := objc.Call[objc.Object](x_, objc.Sel("remoteObjectProxy"))
	return rv
}

// If the connection was created with an NSXPCListenerEndpoint object, returns the endpoint object used. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1411757-endpoint?language=objc
func (x_ XPCConnection) Endpoint() XPCListenerEndpoint {
	rv := objc.Call[XPCListenerEndpoint](x_, objc.Sel("endpoint"))
	return rv
}

// Defines the NSXPCInterface object that describes the protocol for the object represented by the remoteObjectProxy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1411472-remoteobjectinterface?language=objc
func (x_ XPCConnection) RemoteObjectInterface() XPCInterface {
	rv := objc.Call[XPCInterface](x_, objc.Sel("remoteObjectInterface"))
	return rv
}

// Defines the NSXPCInterface object that describes the protocol for the object represented by the remoteObjectProxy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1411472-remoteobjectinterface?language=objc
func (x_ XPCConnection) SetRemoteObjectInterface(value IXPCInterface) {
	objc.Call[objc.Void](x_, objc.Sel("setRemoteObjectInterface:"), objc.Ptr(value))
}

// The name of the XPC service that this connection was configured to connect to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1413751-servicename?language=objc
func (x_ XPCConnection) ServiceName() string {
	rv := objc.Call[string](x_, objc.Sel("serviceName"))
	return rv
}

// The NSXPCInterface object that describes the protocol for the exported object on this connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1408106-exportedinterface?language=objc
func (x_ XPCConnection) ExportedInterface() XPCInterface {
	rv := objc.Call[XPCInterface](x_, objc.Sel("exportedInterface"))
	return rv
}

// The NSXPCInterface object that describes the protocol for the exported object on this connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1408106-exportedinterface?language=objc
func (x_ XPCConnection) SetExportedInterface(value IXPCInterface) {
	objc.Call[objc.Void](x_, objc.Sel("setExportedInterface:"), objc.Ptr(value))
}

// An invalidation handler that is called if the connection can not be formed or the connection has terminated and may not be re-established. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1414358-invalidationhandler?language=objc
func (x_ XPCConnection) InvalidationHandler() func() {
	rv := objc.Call[func()](x_, objc.Sel("invalidationHandler"))
	return rv
}

// An invalidation handler that is called if the connection can not be formed or the connection has terminated and may not be re-established. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnection/1414358-invalidationhandler?language=objc
func (x_ XPCConnection) SetInvalidationHandler(value func()) {
	objc.Call[objc.Void](x_, objc.Sel("setInvalidationHandler:"), value)
}
