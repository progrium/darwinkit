// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLProtocol] class.
var URLProtocolClass = _URLProtocolClass{objc.GetClass("NSURLProtocol")}

type _URLProtocolClass struct {
	objc.Class
}

// An interface definition for the [URLProtocol] class.
type IURLProtocol interface {
	objc.IObject
	StopLoading()
	StartLoading()
	Request() URLRequest
	CachedResponse() CachedURLResponse
	Task() URLSessionTask
	Client() URLProtocolClientWrapper
}

// An abstract class that handles the loading of protocol-specific URL data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol?language=objc
type URLProtocol struct {
	objc.Object
}

func URLProtocolFrom(ptr unsafe.Pointer) URLProtocol {
	return URLProtocol{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ URLProtocol) InitWithTaskCachedResponseClient(task IURLSessionTask, cachedResponse ICachedURLResponse, client PURLProtocolClient) URLProtocol {
	po2 := objc.WrapAsProtocol("NSURLProtocolClient", client)
	rv := objc.Call[URLProtocol](u_, objc.Sel("initWithTask:cachedResponse:client:"), objc.Ptr(task), objc.Ptr(cachedResponse), po2)
	return rv
}

// Creates a URL protocol instance to handle the task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1417672-initwithtask?language=objc
func URLProtocol_InitWithTaskCachedResponseClient(task IURLSessionTask, cachedResponse ICachedURLResponse, client PURLProtocolClient) URLProtocol {
	return URLProtocolClass.Alloc().InitWithTaskCachedResponseClient(task, cachedResponse, client)
}

func (u_ URLProtocol) InitWithRequestCachedResponseClient(request IURLRequest, cachedResponse ICachedURLResponse, client PURLProtocolClient) URLProtocol {
	po2 := objc.WrapAsProtocol("NSURLProtocolClient", client)
	rv := objc.Call[URLProtocol](u_, objc.Sel("initWithRequest:cachedResponse:client:"), objc.Ptr(request), objc.Ptr(cachedResponse), po2)
	return rv
}

// Creates a URL protocol instance to handle the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1414366-initwithrequest?language=objc
func URLProtocol_InitWithRequestCachedResponseClient(request IURLRequest, cachedResponse ICachedURLResponse, client PURLProtocolClient) URLProtocol {
	return URLProtocolClass.Alloc().InitWithRequestCachedResponseClient(request, cachedResponse, client)
}

func (uc _URLProtocolClass) Alloc() URLProtocol {
	rv := objc.Call[URLProtocol](uc, objc.Sel("alloc"))
	return rv
}

func URLProtocol_Alloc() URLProtocol {
	return URLProtocolClass.Alloc()
}

func (uc _URLProtocolClass) New() URLProtocol {
	rv := objc.Call[URLProtocol](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLProtocol() URLProtocol {
	return URLProtocolClass.New()
}

func (u_ URLProtocol) Init() URLProtocol {
	rv := objc.Call[URLProtocol](u_, objc.Sel("init"))
	return rv
}

// Stops protocol-specific loading of the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1408666-stoploading?language=objc
func (u_ URLProtocol) StopLoading() {
	objc.Call[objc.Void](u_, objc.Sel("stopLoading"))
}

// Fetches the property associated with the specified key in the specified request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1416521-propertyforkey?language=objc
func (uc _URLProtocolClass) PropertyForKeyInRequest(key string, request IURLRequest) objc.Object {
	rv := objc.Call[objc.Object](uc, objc.Sel("propertyForKey:inRequest:"), key, objc.Ptr(request))
	return rv
}

// Fetches the property associated with the specified key in the specified request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1416521-propertyforkey?language=objc
func URLProtocol_PropertyForKeyInRequest(key string, request IURLRequest) objc.Object {
	return URLProtocolClass.PropertyForKeyInRequest(key, request)
}

// Returns a canonical version of the specified request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1408650-canonicalrequestforrequest?language=objc
func (uc _URLProtocolClass) CanonicalRequestForRequest(request IURLRequest) URLRequest {
	rv := objc.Call[URLRequest](uc, objc.Sel("canonicalRequestForRequest:"), objc.Ptr(request))
	return rv
}

// Returns a canonical version of the specified request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1408650-canonicalrequestforrequest?language=objc
func URLProtocol_CanonicalRequestForRequest(request IURLRequest) URLRequest {
	return URLProtocolClass.CanonicalRequestForRequest(request)
}

// Attempts to register a subclass of NSURLProtocol, making it visible to the URL loading system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1407208-registerclass?language=objc
func (uc _URLProtocolClass) RegisterClass(protocolClass objc.IClass) bool {
	rv := objc.Call[bool](uc, objc.Sel("registerClass:"), objc.Ptr(protocolClass))
	return rv
}

// Attempts to register a subclass of NSURLProtocol, making it visible to the URL loading system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1407208-registerclass?language=objc
func URLProtocol_RegisterClass(protocolClass objc.IClass) bool {
	return URLProtocolClass.RegisterClass(protocolClass)
}

// Determines whether the protocol subclass can handle the specified task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1416997-caninitwithtask?language=objc
func (uc _URLProtocolClass) CanInitWithTask(task IURLSessionTask) bool {
	rv := objc.Call[bool](uc, objc.Sel("canInitWithTask:"), objc.Ptr(task))
	return rv
}

// Determines whether the protocol subclass can handle the specified task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1416997-caninitwithtask?language=objc
func URLProtocol_CanInitWithTask(task IURLSessionTask) bool {
	return URLProtocolClass.CanInitWithTask(task)
}

// Starts protocol-specific loading of the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1408989-startloading?language=objc
func (u_ URLProtocol) StartLoading() {
	objc.Call[objc.Void](u_, objc.Sel("startLoading"))
}

// Unregisters the specified subclass of NSURLProtocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1409131-unregisterclass?language=objc
func (uc _URLProtocolClass) UnregisterClass(protocolClass objc.IClass) {
	objc.Call[objc.Void](uc, objc.Sel("unregisterClass:"), objc.Ptr(protocolClass))
}

// Unregisters the specified subclass of NSURLProtocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1409131-unregisterclass?language=objc
func URLProtocol_UnregisterClass(protocolClass objc.IClass) {
	URLProtocolClass.UnregisterClass(protocolClass)
}

// Determines whether the protocol subclass can handle the specified request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1411389-caninitwithrequest?language=objc
func (uc _URLProtocolClass) CanInitWithRequest(request IURLRequest) bool {
	rv := objc.Call[bool](uc, objc.Sel("canInitWithRequest:"), objc.Ptr(request))
	return rv
}

// Determines whether the protocol subclass can handle the specified request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1411389-caninitwithrequest?language=objc
func URLProtocol_CanInitWithRequest(request IURLRequest) bool {
	return URLProtocolClass.CanInitWithRequest(request)
}

// Sets the property associated with the specified key in the specified request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1407897-setproperty?language=objc
func (uc _URLProtocolClass) SetPropertyForKeyInRequest(value objc.IObject, key string, request IMutableURLRequest) {
	objc.Call[objc.Void](uc, objc.Sel("setProperty:forKey:inRequest:"), value, key, objc.Ptr(request))
}

// Sets the property associated with the specified key in the specified request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1407897-setproperty?language=objc
func URLProtocol_SetPropertyForKeyInRequest(value objc.IObject, key string, request IMutableURLRequest) {
	URLProtocolClass.SetPropertyForKeyInRequest(value, key, request)
}

// A Boolean value indicating whether two requests are equivalent for cache purposes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1410008-requestiscacheequivalent?language=objc
func (uc _URLProtocolClass) RequestIsCacheEquivalentToRequest(a IURLRequest, b IURLRequest) bool {
	rv := objc.Call[bool](uc, objc.Sel("requestIsCacheEquivalent:toRequest:"), objc.Ptr(a), objc.Ptr(b))
	return rv
}

// A Boolean value indicating whether two requests are equivalent for cache purposes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1410008-requestiscacheequivalent?language=objc
func URLProtocol_RequestIsCacheEquivalentToRequest(a IURLRequest, b IURLRequest) bool {
	return URLProtocolClass.RequestIsCacheEquivalentToRequest(a, b)
}

// Removes the property associated with the specified key in the specified request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1407740-removepropertyforkey?language=objc
func (uc _URLProtocolClass) RemovePropertyForKeyInRequest(key string, request IMutableURLRequest) {
	objc.Call[objc.Void](uc, objc.Sel("removePropertyForKey:inRequest:"), key, objc.Ptr(request))
}

// Removes the property associated with the specified key in the specified request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1407740-removepropertyforkey?language=objc
func URLProtocol_RemovePropertyForKeyInRequest(key string, request IMutableURLRequest) {
	URLProtocolClass.RemovePropertyForKeyInRequest(key, request)
}

// The protocol’s request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1412383-request?language=objc
func (u_ URLProtocol) Request() URLRequest {
	rv := objc.Call[URLRequest](u_, objc.Sel("request"))
	return rv
}

// The protocol’s cached response. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1418409-cachedresponse?language=objc
func (u_ URLProtocol) CachedResponse() CachedURLResponse {
	rv := objc.Call[CachedURLResponse](u_, objc.Sel("cachedResponse"))
	return rv
}

// The protocol’s task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1407649-task?language=objc
func (u_ URLProtocol) Task() URLSessionTask {
	rv := objc.Call[URLSessionTask](u_, objc.Sel("task"))
	return rv
}

// The object the protocol uses to communicate with the URL loading system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocol/1413722-client?language=objc
func (u_ URLProtocol) Client() URLProtocolClientWrapper {
	rv := objc.Call[URLProtocolClientWrapper](u_, objc.Sel("client"))
	return rv
}
