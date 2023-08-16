// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLRequest] class.
var URLRequestClass = _URLRequestClass{objc.GetClass("NSURLRequest")}

type _URLRequestClass struct {
	objc.Class
}

// An interface definition for the [URLRequest] class.
type IURLRequest interface {
	objc.IObject
	ValueForHTTPHeaderField(field string) string
	AllHTTPHeaderFields() map[string]string
	MainDocumentURL() URL
	HTTPShouldHandleCookies() bool
	HTTPMethod() string
	AllowsCellularAccess() bool
	URL() URL
	HTTPBodyStream() InputStream
	AssumesHTTP3Capable() bool
	AllowsExpensiveNetworkAccess() bool
	AllowsConstrainedNetworkAccess() bool
	TimeoutInterval() TimeInterval
	Attribution() URLRequestAttribution
	HTTPBody() []byte
	CachePolicy() URLRequestCachePolicy
	NetworkServiceType() URLRequestNetworkServiceType
	HTTPShouldUsePipelining() bool
}

// A URL load request that is independent of protocol or URL scheme. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest?language=objc
type URLRequest struct {
	objc.Object
}

func URLRequestFrom(ptr unsafe.Pointer) URLRequest {
	return URLRequest{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ URLRequest) InitWithURL(URL IURL) URLRequest {
	rv := objc.Call[URLRequest](u_, objc.Sel("initWithURL:"), objc.Ptr(URL))
	return rv
}

// Creates a URL request for a specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1410303-initwithurl?language=objc
func URLRequest_InitWithURL(URL IURL) URLRequest {
	return URLRequestClass.Alloc().InitWithURL(URL)
}

func (uc _URLRequestClass) RequestWithURL(URL IURL) URLRequest {
	rv := objc.Call[URLRequest](uc, objc.Sel("requestWithURL:"), objc.Ptr(URL))
	return rv
}

// Creates and returns a URL request for a specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1528603-requestwithurl?language=objc
func URLRequest_RequestWithURL(URL IURL) URLRequest {
	return URLRequestClass.RequestWithURL(URL)
}

func (uc _URLRequestClass) Alloc() URLRequest {
	rv := objc.Call[URLRequest](uc, objc.Sel("alloc"))
	return rv
}

func URLRequest_Alloc() URLRequest {
	return URLRequestClass.Alloc()
}

func (uc _URLRequestClass) New() URLRequest {
	rv := objc.Call[URLRequest](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLRequest() URLRequest {
	return URLRequestClass.New()
}

func (u_ URLRequest) Init() URLRequest {
	rv := objc.Call[URLRequest](u_, objc.Sel("init"))
	return rv
}

// Returns the value of the specified HTTP header field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1409376-valueforhttpheaderfield?language=objc
func (u_ URLRequest) ValueForHTTPHeaderField(field string) string {
	rv := objc.Call[string](u_, objc.Sel("valueForHTTPHeaderField:"), field)
	return rv
}

// A dictionary containing all of the HTTP header fields for a request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1418477-allhttpheaderfields?language=objc
func (u_ URLRequest) AllHTTPHeaderFields() map[string]string {
	rv := objc.Call[map[string]string](u_, objc.Sel("allHTTPHeaderFields"))
	return rv
}

// The main document URL associated with the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1414134-maindocumenturl?language=objc
func (u_ URLRequest) MainDocumentURL() URL {
	rv := objc.Call[URL](u_, objc.Sel("mainDocumentURL"))
	return rv
}

// A Boolean value that indicates whether the default cookie handling will be used for this request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1418369-httpshouldhandlecookies?language=objc
func (u_ URLRequest) HTTPShouldHandleCookies() bool {
	rv := objc.Call[bool](u_, objc.Sel("HTTPShouldHandleCookies"))
	return rv
}

// The HTTP request method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1413030-httpmethod?language=objc
func (u_ URLRequest) HTTPMethod() string {
	rv := objc.Call[string](u_, objc.Sel("HTTPMethod"))
	return rv
}

// A Boolean value that indicates whether the request is allowed to use the cellular radio (if present). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1412032-allowscellularaccess?language=objc
func (u_ URLRequest) AllowsCellularAccess() bool {
	rv := objc.Call[bool](u_, objc.Sel("allowsCellularAccess"))
	return rv
}

// The URL being requested. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1408996-url?language=objc
func (u_ URLRequest) URL() URL {
	rv := objc.Call[URL](u_, objc.Sel("URL"))
	return rv
}

// The request body as an input stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1407341-httpbodystream?language=objc
func (u_ URLRequest) HTTPBodyStream() InputStream {
	rv := objc.Call[InputStream](u_, objc.Sel("HTTPBodyStream"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/3735880-assumeshttp3capable?language=objc
func (u_ URLRequest) AssumesHTTP3Capable() bool {
	rv := objc.Call[bool](u_, objc.Sel("assumesHTTP3Capable"))
	return rv
}

// A Boolean value that indicates whether connections may use a network interface that the system considers expensive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/3325679-allowsexpensivenetworkaccess?language=objc
func (u_ URLRequest) AllowsExpensiveNetworkAccess() bool {
	rv := objc.Call[bool](u_, objc.Sel("allowsExpensiveNetworkAccess"))
	return rv
}

// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/3325678-allowsconstrainednetworkaccess?language=objc
func (u_ URLRequest) AllowsConstrainedNetworkAccess() bool {
	rv := objc.Call[bool](u_, objc.Sel("allowsConstrainedNetworkAccess"))
	return rv
}

// The request’s timeout interval, in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1418229-timeoutinterval?language=objc
func (u_ URLRequest) TimeoutInterval() TimeInterval {
	rv := objc.Call[TimeInterval](u_, objc.Sel("timeoutInterval"))
	return rv
}

// A Boolean value indicating whether the NSURLRequest implements the NSSecureCoding protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1416510-supportssecurecoding?language=objc
func (uc _URLRequestClass) SupportsSecureCoding() bool {
	rv := objc.Call[bool](uc, objc.Sel("supportsSecureCoding"))
	return rv
}

// A Boolean value indicating whether the NSURLRequest implements the NSSecureCoding protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1416510-supportssecurecoding?language=objc
func URLRequest_SupportsSecureCoding() bool {
	return URLRequestClass.SupportsSecureCoding()
}

// The entity that initiates the network request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/3746973-attribution?language=objc
func (u_ URLRequest) Attribution() URLRequestAttribution {
	rv := objc.Call[URLRequestAttribution](u_, objc.Sel("attribution"))
	return rv
}

// The request body. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1411317-httpbody?language=objc
func (u_ URLRequest) HTTPBody() []byte {
	rv := objc.Call[[]byte](u_, objc.Sel("HTTPBody"))
	return rv
}

// The request’s cache policy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1407944-cachepolicy?language=objc
func (u_ URLRequest) CachePolicy() URLRequestCachePolicy {
	rv := objc.Call[URLRequestCachePolicy](u_, objc.Sel("cachePolicy"))
	return rv
}

// The network service type of the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1418333-networkservicetype?language=objc
func (u_ URLRequest) NetworkServiceType() URLRequestNetworkServiceType {
	rv := objc.Call[URLRequestNetworkServiceType](u_, objc.Sel("networkServiceType"))
	return rv
}

// A Boolean value that indicates whether the request should continue transmitting data before receiving a response from an earlier transmission. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/1409170-httpshouldusepipelining?language=objc
func (u_ URLRequest) HTTPShouldUsePipelining() bool {
	rv := objc.Call[bool](u_, objc.Sel("HTTPShouldUsePipelining"))
	return rv
}
