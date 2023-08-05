// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var URLRequestClass = _URLRequestClass{objc.GetClass("NSURLRequest")}

type _URLRequestClass struct {
	objc.Class
}

type IURLRequest interface {
	objc.IObject
	ValueForHTTPHeaderField(field string) string
	CachePolicy() URLRequestCachePolicy
	HTTPMethod() string
	URL() URL
	HTTPBody() []byte
	HTTPBodyStream() InputStream
	MainDocumentURL() URL
	AllHTTPHeaderFields() map[string]string
	TimeoutInterval() TimeInterval
	HTTPShouldHandleCookies() bool
	HTTPShouldUsePipelining() bool
	AllowsCellularAccess() bool
	AllowsConstrainedNetworkAccess() bool
	AllowsExpensiveNetworkAccess() bool
	NetworkServiceType() URLRequestNetworkServiceType
	Attribution() URLRequestAttribution
	AssumesHTTP3Capable() bool
}

type URLRequest struct {
	objc.Object
}

func MakeURLRequest(ptr unsafe.Pointer) URLRequest {
	return URLRequest{
		Object: objc.MakeObject(ptr),
	}
}

func (uc _URLRequestClass) RequestWithURL(URL IURL) URLRequest {
	rv := objc.CallMethod[URLRequest](uc, objc.GetSelector("requestWithURL:"), objc.ExtractPtr(URL))
	return rv
}

func URLRequest_RequestWithURL(URL IURL) URLRequest {
	return URLRequestClass.RequestWithURL(URL)
}

func (u_ URLRequest) InitWithURL(URL IURL) URLRequest {
	rv := objc.CallMethod[URLRequest](u_, objc.GetSelector("initWithURL:"), objc.ExtractPtr(URL))
	return rv
}

func URLRequest_InitWithURL(URL IURL) URLRequest {
	return URLRequestClass.Alloc().InitWithURL(URL)
}

func (uc _URLRequestClass) RequestWithURLCachePolicyTimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) URLRequest {
	rv := objc.CallMethod[URLRequest](uc, objc.GetSelector("requestWithURL:cachePolicy:timeoutInterval:"), objc.ExtractPtr(URL), cachePolicy, timeoutInterval)
	return rv
}

func URLRequest_RequestWithURLCachePolicyTimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) URLRequest {
	return URLRequestClass.RequestWithURLCachePolicyTimeoutInterval(URL, cachePolicy, timeoutInterval)
}

func (u_ URLRequest) InitWithURLCachePolicyTimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) URLRequest {
	rv := objc.CallMethod[URLRequest](u_, objc.GetSelector("initWithURL:cachePolicy:timeoutInterval:"), objc.ExtractPtr(URL), cachePolicy, timeoutInterval)
	return rv
}

func URLRequest_InitWithURLCachePolicyTimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) URLRequest {
	return URLRequestClass.Alloc().InitWithURLCachePolicyTimeoutInterval(URL, cachePolicy, timeoutInterval)
}

func (uc _URLRequestClass) Alloc() URLRequest {
	rv := objc.CallMethod[URLRequest](uc, objc.GetSelector("alloc"))
	return rv
}

func URLRequest_Alloc() URLRequest {
	return URLRequestClass.Alloc()
}

func (uc _URLRequestClass) New() URLRequest {
	rv := objc.CallMethod[URLRequest](uc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewURLRequest() URLRequest {
	return URLRequestClass.New()
}

func URLRequest_New() URLRequest {
	return URLRequestClass.New()
}

func (u_ URLRequest) Init() URLRequest {
	rv := objc.CallMethod[URLRequest](u_, objc.GetSelector("init"))
	return rv
}

func URLRequest_Init() URLRequest {
	return URLRequestClass.Alloc().Init()
}

func (u_ URLRequest) ValueForHTTPHeaderField(field string) string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("valueForHTTPHeaderField:"), field)
	return rv
}

func (u_ URLRequest) CachePolicy() URLRequestCachePolicy {
	rv := objc.CallMethod[URLRequestCachePolicy](u_, objc.GetSelector("cachePolicy"))
	return rv
}

func (u_ URLRequest) HTTPMethod() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("HTTPMethod"))
	return rv
}

func (u_ URLRequest) URL() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("URL"))
	return rv
}

func (u_ URLRequest) HTTPBody() []byte {
	rv := objc.CallMethod[[]byte](u_, objc.GetSelector("HTTPBody"))
	return rv
}

func (u_ URLRequest) HTTPBodyStream() InputStream {
	rv := objc.CallMethod[InputStream](u_, objc.GetSelector("HTTPBodyStream"))
	return rv
}

func (u_ URLRequest) MainDocumentURL() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("mainDocumentURL"))
	return rv
}

func (u_ URLRequest) AllHTTPHeaderFields() map[string]string {
	rv := objc.CallMethod[map[string]string](u_, objc.GetSelector("allHTTPHeaderFields"))
	return rv
}

func (u_ URLRequest) TimeoutInterval() TimeInterval {
	rv := objc.CallMethod[TimeInterval](u_, objc.GetSelector("timeoutInterval"))
	return rv
}

func (u_ URLRequest) HTTPShouldHandleCookies() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("HTTPShouldHandleCookies"))
	return rv
}

func (u_ URLRequest) HTTPShouldUsePipelining() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("HTTPShouldUsePipelining"))
	return rv
}

func (u_ URLRequest) AllowsCellularAccess() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("allowsCellularAccess"))
	return rv
}

func (u_ URLRequest) AllowsConstrainedNetworkAccess() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("allowsConstrainedNetworkAccess"))
	return rv
}

func (u_ URLRequest) AllowsExpensiveNetworkAccess() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("allowsExpensiveNetworkAccess"))
	return rv
}

func (u_ URLRequest) NetworkServiceType() URLRequestNetworkServiceType {
	rv := objc.CallMethod[URLRequestNetworkServiceType](u_, objc.GetSelector("networkServiceType"))
	return rv
}

func (uc _URLRequestClass) SupportsSecureCoding() bool {
	rv := objc.CallMethod[bool](uc, objc.GetSelector("supportsSecureCoding"))
	return rv
}

func URLRequest_SupportsSecureCoding() bool {
	return URLRequestClass.SupportsSecureCoding()
}

func (u_ URLRequest) Attribution() URLRequestAttribution {
	rv := objc.CallMethod[URLRequestAttribution](u_, objc.GetSelector("attribution"))
	return rv
}

func (u_ URLRequest) AssumesHTTP3Capable() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("assumesHTTP3Capable"))
	return rv
}
