// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var MutableURLRequestClass = _MutableURLRequestClass{objc.GetClass("NSMutableURLRequest")}

type _MutableURLRequestClass struct {
	objc.Class
}

type IMutableURLRequest interface {
	IURLRequest
	AddValueForHTTPHeaderField(value string, field string)
	SetValueForHTTPHeaderField(value string, field string)
	SetCachePolicy(value URLRequestCachePolicy)
	SetHTTPMethod(value string)
	SetURL(value IURL)
	SetHTTPBody(value []byte)
	SetHTTPBodyStream(value IInputStream)
	SetMainDocumentURL(value IURL)
	SetAllHTTPHeaderFields(value map[string]string)
	SetTimeoutInterval(value TimeInterval)
	SetHTTPShouldHandleCookies(value bool)
	SetHTTPShouldUsePipelining(value bool)
	SetAllowsCellularAccess(value bool)
	SetAllowsConstrainedNetworkAccess(value bool)
	SetAllowsExpensiveNetworkAccess(value bool)
	SetNetworkServiceType(value URLRequestNetworkServiceType)
	SetAttribution(value URLRequestAttribution)
	SetAssumesHTTP3Capable(value bool)
	RequiresDNSSECValidation() bool
	SetRequiresDNSSECValidation(value bool)
}

type MutableURLRequest struct {
	URLRequest
}

func MakeMutableURLRequest(ptr unsafe.Pointer) MutableURLRequest {
	return MutableURLRequest{
		URLRequest: MakeURLRequest(ptr),
	}
}

func (mc _MutableURLRequestClass) RequestWithURL(URL IURL) MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](mc, objc.GetSelector("requestWithURL:"), objc.ExtractPtr(URL))
	return rv
}

func MutableURLRequest_RequestWithURL(URL IURL) MutableURLRequest {
	return MutableURLRequestClass.RequestWithURL(URL)
}

func (m_ MutableURLRequest) InitWithURL(URL IURL) MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](m_, objc.GetSelector("initWithURL:"), objc.ExtractPtr(URL))
	return rv
}

func MutableURLRequest_InitWithURL(URL IURL) MutableURLRequest {
	return MutableURLRequestClass.Alloc().InitWithURL(URL)
}

func (mc _MutableURLRequestClass) RequestWithURLCachePolicyTimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](mc, objc.GetSelector("requestWithURL:cachePolicy:timeoutInterval:"), objc.ExtractPtr(URL), cachePolicy, timeoutInterval)
	return rv
}

func MutableURLRequest_RequestWithURLCachePolicyTimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) MutableURLRequest {
	return MutableURLRequestClass.RequestWithURLCachePolicyTimeoutInterval(URL, cachePolicy, timeoutInterval)
}

func (m_ MutableURLRequest) InitWithURLCachePolicyTimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](m_, objc.GetSelector("initWithURL:cachePolicy:timeoutInterval:"), objc.ExtractPtr(URL), cachePolicy, timeoutInterval)
	return rv
}

func MutableURLRequest_InitWithURLCachePolicyTimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) MutableURLRequest {
	return MutableURLRequestClass.Alloc().InitWithURLCachePolicyTimeoutInterval(URL, cachePolicy, timeoutInterval)
}

func (mc _MutableURLRequestClass) Alloc() MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](mc, objc.GetSelector("alloc"))
	return rv
}

func MutableURLRequest_Alloc() MutableURLRequest {
	return MutableURLRequestClass.Alloc()
}

func (mc _MutableURLRequestClass) New() MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](mc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewMutableURLRequest() MutableURLRequest {
	return MutableURLRequestClass.New()
}

func MutableURLRequest_New() MutableURLRequest {
	return MutableURLRequestClass.New()
}

func (m_ MutableURLRequest) Init() MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](m_, objc.GetSelector("init"))
	return rv
}

func MutableURLRequest_Init() MutableURLRequest {
	return MutableURLRequestClass.Alloc().Init()
}

func (m_ MutableURLRequest) AddValueForHTTPHeaderField(value string, field string) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("addValue:forHTTPHeaderField:"), value, field)
}

func (m_ MutableURLRequest) SetValueForHTTPHeaderField(value string, field string) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setValue:forHTTPHeaderField:"), value, field)
}

func (m_ MutableURLRequest) SetCachePolicy(value URLRequestCachePolicy) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setCachePolicy:"), value)
}

func (m_ MutableURLRequest) SetHTTPMethod(value string) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setHTTPMethod:"), value)
}

func (m_ MutableURLRequest) SetURL(value IURL) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setURL:"), objc.ExtractPtr(value))
}

func (m_ MutableURLRequest) SetHTTPBody(value []byte) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setHTTPBody:"), value)
}

func (m_ MutableURLRequest) SetHTTPBodyStream(value IInputStream) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setHTTPBodyStream:"), objc.ExtractPtr(value))
}

func (m_ MutableURLRequest) SetMainDocumentURL(value IURL) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setMainDocumentURL:"), objc.ExtractPtr(value))
}

func (m_ MutableURLRequest) SetAllHTTPHeaderFields(value map[string]string) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAllHTTPHeaderFields:"), value)
}

func (m_ MutableURLRequest) SetTimeoutInterval(value TimeInterval) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setTimeoutInterval:"), value)
}

func (m_ MutableURLRequest) SetHTTPShouldHandleCookies(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setHTTPShouldHandleCookies:"), value)
}

func (m_ MutableURLRequest) SetHTTPShouldUsePipelining(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setHTTPShouldUsePipelining:"), value)
}

func (m_ MutableURLRequest) SetAllowsCellularAccess(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAllowsCellularAccess:"), value)
}

func (m_ MutableURLRequest) SetAllowsConstrainedNetworkAccess(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAllowsConstrainedNetworkAccess:"), value)
}

func (m_ MutableURLRequest) SetAllowsExpensiveNetworkAccess(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAllowsExpensiveNetworkAccess:"), value)
}

func (m_ MutableURLRequest) SetNetworkServiceType(value URLRequestNetworkServiceType) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setNetworkServiceType:"), value)
}

func (m_ MutableURLRequest) SetAttribution(value URLRequestAttribution) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAttribution:"), value)
}

func (m_ MutableURLRequest) SetAssumesHTTP3Capable(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAssumesHTTP3Capable:"), value)
}

func (m_ MutableURLRequest) RequiresDNSSECValidation() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("requiresDNSSECValidation"))
	return rv
}

func (m_ MutableURLRequest) SetRequiresDNSSECValidation(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setRequiresDNSSECValidation:"), value)
}
