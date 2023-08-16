// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLConnection] class.
var URLConnectionClass = _URLConnectionClass{objc.GetClass("NSURLConnection")}

type _URLConnectionClass struct {
	objc.Class
}

// An interface definition for the [URLConnection] class.
type IURLConnection interface {
	objc.IObject
	ScheduleInRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode)
	Start()
	SetDelegateQueue(queue IOperationQueue)
	UnscheduleFromRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode)
	Cancel()
	CurrentRequest() URLRequest
	OriginalRequest() URLRequest
}

// An object that enables you to start and stop URL requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnection?language=objc
type URLConnection struct {
	objc.Object
}

func URLConnectionFrom(ptr unsafe.Pointer) URLConnection {
	return URLConnection{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _URLConnectionClass) Alloc() URLConnection {
	rv := objc.Call[URLConnection](uc, objc.Sel("alloc"))
	return rv
}

func URLConnection_Alloc() URLConnection {
	return URLConnectionClass.Alloc()
}

func (uc _URLConnectionClass) New() URLConnection {
	rv := objc.Call[URLConnection](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLConnection() URLConnection {
	return URLConnectionClass.New()
}

func (u_ URLConnection) Init() URLConnection {
	rv := objc.Call[URLConnection](u_, objc.Sel("init"))
	return rv
}

// Determines the run loop and mode that the connection uses to call methods on its delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnection/1417485-scheduleinrunloop?language=objc
func (u_ URLConnection) ScheduleInRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode) {
	objc.Call[objc.Void](u_, objc.Sel("scheduleInRunLoop:forMode:"), objc.Ptr(aRunLoop), mode)
}

// Returns whether a request can be handled based on a preflight evaluation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnection/1413072-canhandlerequest?language=objc
func (uc _URLConnectionClass) CanHandleRequest(request IURLRequest) bool {
	rv := objc.Call[bool](uc, objc.Sel("canHandleRequest:"), objc.Ptr(request))
	return rv
}

// Returns whether a request can be handled based on a preflight evaluation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnection/1413072-canhandlerequest?language=objc
func URLConnection_CanHandleRequest(request IURLRequest) bool {
	return URLConnectionClass.CanHandleRequest(request)
}

// Causes the connection to begin loading data, if it has not already. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnection/1417345-start?language=objc
func (u_ URLConnection) Start() {
	objc.Call[objc.Void](u_, objc.Sel("start"))
}

// Determines the operation queue that is used to call methods on the connection’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnection/1411849-setdelegatequeue?language=objc
func (u_ URLConnection) SetDelegateQueue(queue IOperationQueue) {
	objc.Call[objc.Void](u_, objc.Sel("setDelegateQueue:"), objc.Ptr(queue))
}

// Causes the connection to stop calling delegate methods in the specified run loop and mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnection/1409722-unschedulefromrunloop?language=objc
func (u_ URLConnection) UnscheduleFromRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode) {
	objc.Call[objc.Void](u_, objc.Sel("unscheduleFromRunLoop:forMode:"), objc.Ptr(aRunLoop), mode)
}

// Cancels an asynchronous load of a request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnection/1407797-cancel?language=objc
func (u_ URLConnection) Cancel() {
	objc.Call[objc.Void](u_, objc.Sel("cancel"))
}

// The current connection request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnection/1409060-currentrequest?language=objc
func (u_ URLConnection) CurrentRequest() URLRequest {
	rv := objc.Call[URLRequest](u_, objc.Sel("currentRequest"))
	return rv
}

// A deep copy of the original connection request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlconnection/1411340-originalrequest?language=objc
func (u_ URLConnection) OriginalRequest() URLRequest {
	rv := objc.Call[URLRequest](u_, objc.Sel("originalRequest"))
	return rv
}
