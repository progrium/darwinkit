// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableURLRequest] class.
var MutableURLRequestClass = _MutableURLRequestClass{objc.GetClass("NSMutableURLRequest")}

type _MutableURLRequestClass struct {
	objc.Class
}

// An interface definition for the [MutableURLRequest] class.
type IMutableURLRequest interface {
	IURLRequest
	AddValueForHTTPHeaderField(value string, field string)
	SetValueForHTTPHeaderField(value string, field string)
	SetAllHTTPHeaderFields(value map[string]string)
	SetMainDocumentURL(value IURL)
	SetHTTPShouldHandleCookies(value bool)
	SetHTTPMethod(value string)
	SetAllowsCellularAccess(value bool)
	SetURL(value IURL)
	SetHTTPBodyStream(value IInputStream)
	SetAssumesHTTP3Capable(value bool)
	SetAllowsExpensiveNetworkAccess(value bool)
	SetAllowsConstrainedNetworkAccess(value bool)
	SetTimeoutInterval(value TimeInterval)
	SetAttribution(value URLRequestAttribution)
	SetHTTPBody(value []byte)
	SetCachePolicy(value URLRequestCachePolicy)
	SetNetworkServiceType(value URLRequestNetworkServiceType)
	SetHTTPShouldUsePipelining(value bool)
}

// A mutable URL load request that is independent of protocol or URL scheme. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest?language=objc
type MutableURLRequest struct {
	URLRequest
}

func MutableURLRequestFrom(ptr unsafe.Pointer) MutableURLRequest {
	return MutableURLRequest{
		URLRequest: URLRequestFrom(ptr),
	}
}

func (mc _MutableURLRequestClass) Alloc() MutableURLRequest {
	rv := objc.Call[MutableURLRequest](mc, objc.Sel("alloc"))
	return rv
}

func MutableURLRequest_Alloc() MutableURLRequest {
	return MutableURLRequestClass.Alloc()
}

func (mc _MutableURLRequestClass) New() MutableURLRequest {
	rv := objc.Call[MutableURLRequest](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableURLRequest() MutableURLRequest {
	return MutableURLRequestClass.New()
}

func (m_ MutableURLRequest) Init() MutableURLRequest {
	rv := objc.Call[MutableURLRequest](m_, objc.Sel("init"))
	return rv
}

func (m_ MutableURLRequest) InitWithURL(URL IURL) MutableURLRequest {
	rv := objc.Call[MutableURLRequest](m_, objc.Sel("initWithURL:"), objc.Ptr(URL))
	return rv
}

// Creates a URL request for a specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1410303-initwithurl?language=objc
func NewMutableURLRequestWithURL(URL IURL) MutableURLRequest {
	instance := MutableURLRequestClass.Alloc().InitWithURL(URL)
	instance.Autorelease()
	return instance
}

func (mc _MutableURLRequestClass) RequestWithURL(URL IURL) MutableURLRequest {
	rv := objc.Call[MutableURLRequest](mc, objc.Sel("requestWithURL:"), objc.Ptr(URL))
	return rv
}

// Creates and returns a URL request for a specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1528603-requestwithurl?language=objc
func MutableURLRequest_RequestWithURL(URL IURL) MutableURLRequest {
	return MutableURLRequestClass.RequestWithURL(URL)
}

// Adds a value to the header field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/1407676-addvalue?language=objc
func (m_ MutableURLRequest) AddValueForHTTPHeaderField(value string, field string) {
	objc.Call[objc.Void](m_, objc.Sel("addValue:forHTTPHeaderField:"), value, field)
}

// Sets a value for the header field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/1408793-setvalue?language=objc
func (m_ MutableURLRequest) SetValueForHTTPHeaderField(value string, field string) {
	objc.Call[objc.Void](m_, objc.Sel("setValue:forHTTPHeaderField:"), value, field)
}

// A dictionary containing all of the HTTP header fields for a request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/1414617-allhttpheaderfields?language=objc
func (m_ MutableURLRequest) SetAllHTTPHeaderFields(value map[string]string) {
	objc.Call[objc.Void](m_, objc.Sel("setAllHTTPHeaderFields:"), value)
}

// The main document URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/1415630-maindocumenturl?language=objc
func (m_ MutableURLRequest) SetMainDocumentURL(value IURL) {
	objc.Call[objc.Void](m_, objc.Sel("setMainDocumentURL:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the request should use the default cookie handling for the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/1415485-httpshouldhandlecookies?language=objc
func (m_ MutableURLRequest) SetHTTPShouldHandleCookies(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setHTTPShouldHandleCookies:"), value)
}

// The HTTP request method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/1413047-httpmethod?language=objc
func (m_ MutableURLRequest) SetHTTPMethod(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setHTTPMethod:"), value)
}

// A Boolean value that indicates whether a connection can use the device’s cellular network (if present). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/1416749-allowscellularaccess?language=objc
func (m_ MutableURLRequest) SetAllowsCellularAccess(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAllowsCellularAccess:"), value)
}

// The URL being requested. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/1410342-url?language=objc
func (m_ MutableURLRequest) SetURL(value IURL) {
	objc.Call[objc.Void](m_, objc.Sel("setURL:"), objc.Ptr(value))
}

// The request body as an input stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/1409529-httpbodystream?language=objc
func (m_ MutableURLRequest) SetHTTPBodyStream(value IInputStream) {
	objc.Call[objc.Void](m_, objc.Sel("setHTTPBodyStream:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/3735879-assumeshttp3capable?language=objc
func (m_ MutableURLRequest) SetAssumesHTTP3Capable(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAssumesHTTP3Capable:"), value)
}

// A Boolean value that indicates whether connections may use a network interface that the system considers expensive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/3325677-allowsexpensivenetworkaccess?language=objc
func (m_ MutableURLRequest) SetAllowsExpensiveNetworkAccess(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAllowsExpensiveNetworkAccess:"), value)
}

// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/3325676-allowsconstrainednetworkaccess?language=objc
func (m_ MutableURLRequest) SetAllowsConstrainedNetworkAccess(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAllowsConstrainedNetworkAccess:"), value)
}

// The request’s timeout interval, in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/1414063-timeoutinterval?language=objc
func (m_ MutableURLRequest) SetTimeoutInterval(value TimeInterval) {
	objc.Call[objc.Void](m_, objc.Sel("setTimeoutInterval:"), value)
}

// The entity that initiates the network request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/3746972-attribution?language=objc
func (m_ MutableURLRequest) SetAttribution(value URLRequestAttribution) {
	objc.Call[objc.Void](m_, objc.Sel("setAttribution:"), value)
}

// The request body. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/1409064-httpbody?language=objc
func (m_ MutableURLRequest) SetHTTPBody(value []byte) {
	objc.Call[objc.Void](m_, objc.Sel("setHTTPBody:"), value)
}

// The request’s cache policy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/1414716-cachepolicy?language=objc
func (m_ MutableURLRequest) SetCachePolicy(value URLRequestCachePolicy) {
	objc.Call[objc.Void](m_, objc.Sel("setCachePolicy:"), value)
}

// The network service type of the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/1412378-networkservicetype?language=objc
func (m_ MutableURLRequest) SetNetworkServiceType(value URLRequestNetworkServiceType) {
	objc.Call[objc.Void](m_, objc.Sel("setNetworkServiceType:"), value)
}

// A Boolean value that indicates whether the request can continue transmitting data before receiving a response from an earlier transmission. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/1412705-httpshouldusepipelining?language=objc
func (m_ MutableURLRequest) SetHTTPShouldUsePipelining(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setHTTPShouldUsePipelining:"), value)
}
