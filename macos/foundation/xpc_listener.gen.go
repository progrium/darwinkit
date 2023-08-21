// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [XPCListener] class.
var XPCListenerClass = _XPCListenerClass{objc.GetClass("NSXPCListener")}

type _XPCListenerClass struct {
	objc.Class
}

// An interface definition for the [XPCListener] class.
type IXPCListener interface {
	objc.IObject
	Suspend()
	Resume()
	Invalidate()
	Endpoint() XPCListenerEndpoint
	Delegate() XPCListenerDelegateWrapper
	SetDelegate(value PXPCListenerDelegate)
	SetDelegateObject(valueObject objc.IObject)
}

// A listener that waits for new incoming connections, configures them, and accepts or rejects them. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpclistener?language=objc
type XPCListener struct {
	objc.Object
}

func XPCListenerFrom(ptr unsafe.Pointer) XPCListener {
	return XPCListener{
		Object: objc.ObjectFrom(ptr),
	}
}

func (x_ XPCListener) InitWithMachServiceName(name string) XPCListener {
	rv := objc.Call[XPCListener](x_, objc.Sel("initWithMachServiceName:"), name)
	return rv
}

// Initializes a listener in a LaunchAgent or LaunchDaemon which has a name advertised in a launchd.plist file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpclistener/1414106-initwithmachservicename?language=objc
func NewXPCListenerWithMachServiceName(name string) XPCListener {
	instance := XPCListenerClass.Alloc().InitWithMachServiceName(name)
	instance.Autorelease()
	return instance
}

func (xc _XPCListenerClass) Alloc() XPCListener {
	rv := objc.Call[XPCListener](xc, objc.Sel("alloc"))
	return rv
}

func XPCListener_Alloc() XPCListener {
	return XPCListenerClass.Alloc()
}

func (xc _XPCListenerClass) New() XPCListener {
	rv := objc.Call[XPCListener](xc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewXPCListener() XPCListener {
	return XPCListenerClass.New()
}

func (x_ XPCListener) Init() XPCListener {
	rv := objc.Call[XPCListener](x_, objc.Sel("init"))
	return rv
}

// Returns a new anonymous listener connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpclistener/1412648-anonymouslistener?language=objc
func (xc _XPCListenerClass) AnonymousListener() XPCListener {
	rv := objc.Call[XPCListener](xc, objc.Sel("anonymousListener"))
	return rv
}

// Returns a new anonymous listener connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpclistener/1412648-anonymouslistener?language=objc
func XPCListener_AnonymousListener() XPCListener {
	return XPCListenerClass.AnonymousListener()
}

// Suspends the listener. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpclistener/1411596-suspend?language=objc
func (x_ XPCListener) Suspend() {
	objc.Call[objc.Void](x_, objc.Sel("suspend"))
}

// Returns the singleton listener used to listen for incoming connections in an XPC service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpclistener/1408414-servicelistener?language=objc
func (xc _XPCListenerClass) ServiceListener() XPCListener {
	rv := objc.Call[XPCListener](xc, objc.Sel("serviceListener"))
	return rv
}

// Returns the singleton listener used to listen for incoming connections in an XPC service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpclistener/1408414-servicelistener?language=objc
func XPCListener_ServiceListener() XPCListener {
	return XPCListenerClass.ServiceListener()
}

// Starts processing of incoming requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpclistener/1409652-resume?language=objc
func (x_ XPCListener) Resume() {
	objc.Call[objc.Void](x_, objc.Sel("resume"))
}

// Invalidates the listener. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpclistener/1418427-invalidate?language=objc
func (x_ XPCListener) Invalidate() {
	objc.Call[objc.Void](x_, objc.Sel("invalidate"))
}

// Returns an endpoint object that may be sent over an existing connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpclistener/1408519-endpoint?language=objc
func (x_ XPCListener) Endpoint() XPCListenerEndpoint {
	rv := objc.Call[XPCListenerEndpoint](x_, objc.Sel("endpoint"))
	return rv
}

// The delegate for the listener. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpclistener/1408939-delegate?language=objc
func (x_ XPCListener) Delegate() XPCListenerDelegateWrapper {
	rv := objc.Call[XPCListenerDelegateWrapper](x_, objc.Sel("delegate"))
	return rv
}

// The delegate for the listener. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpclistener/1408939-delegate?language=objc
func (x_ XPCListener) SetDelegate(value PXPCListenerDelegate) {
	po0 := objc.WrapAsProtocol("NSXPCListenerDelegate", value)
	objc.SetAssociatedObject(x_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](x_, objc.Sel("setDelegate:"), po0)
}

// The delegate for the listener. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpclistener/1408939-delegate?language=objc
func (x_ XPCListener) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](x_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}
