// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLSessionTaskTransactionMetrics] class.
var URLSessionTaskTransactionMetricsClass = _URLSessionTaskTransactionMetricsClass{objc.GetClass("NSURLSessionTaskTransactionMetrics")}

type _URLSessionTaskTransactionMetricsClass struct {
	objc.Class
}

// An interface definition for the [URLSessionTaskTransactionMetrics] class.
type IURLSessionTaskTransactionMetrics interface {
	objc.IObject
	IsCellular() bool
	DomainResolutionProtocol() URLSessionTaskMetricsDomainResolutionProtocol
	Request() URLRequest
	IsExpensive() bool
	NegotiatedTLSCipherSuite() Number
	ConnectStartDate() Date
	CountOfRequestBodyBytesSent() int64
	RemoteAddress() string
	RequestStartDate() Date
	DomainLookupStartDate() Date
	CountOfRequestBodyBytesBeforeEncoding() int64
	CountOfRequestHeaderBytesSent() int64
	RemotePort() Number
	CountOfResponseBodyBytesReceived() int64
	ConnectEndDate() Date
	CountOfResponseHeaderBytesReceived() int64
	SecureConnectionStartDate() Date
	CountOfResponseBodyBytesAfterDecoding() int64
	LocalAddress() string
	DomainLookupEndDate() Date
	IsReusedConnection() bool
	ResourceFetchType() URLSessionTaskMetricsResourceFetchType
	IsMultipath() bool
	ResponseStartDate() Date
	NetworkProtocolName() string
	LocalPort() Number
	FetchStartDate() Date
	NegotiatedTLSProtocolVersion() Number
	ResponseEndDate() Date
	IsConstrained() bool
	IsProxyConnection() bool
	Response() URLResponse
	SecureConnectionEndDate() Date
	RequestEndDate() Date
}

// An object that encapsualtes the performance metrics collected by the URL Loading System during the execution of a session task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics?language=objc
type URLSessionTaskTransactionMetrics struct {
	objc.Object
}

func URLSessionTaskTransactionMetricsFrom(ptr unsafe.Pointer) URLSessionTaskTransactionMetrics {
	return URLSessionTaskTransactionMetrics{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _URLSessionTaskTransactionMetricsClass) Alloc() URLSessionTaskTransactionMetrics {
	rv := objc.Call[URLSessionTaskTransactionMetrics](uc, objc.Sel("alloc"))
	return rv
}

func URLSessionTaskTransactionMetrics_Alloc() URLSessionTaskTransactionMetrics {
	return URLSessionTaskTransactionMetricsClass.Alloc()
}

func (uc _URLSessionTaskTransactionMetricsClass) New() URLSessionTaskTransactionMetrics {
	rv := objc.Call[URLSessionTaskTransactionMetrics](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLSessionTaskTransactionMetrics() URLSessionTaskTransactionMetrics {
	return URLSessionTaskTransactionMetricsClass.New()
}

func (u_ URLSessionTaskTransactionMetrics) Init() URLSessionTaskTransactionMetrics {
	rv := objc.Call[URLSessionTaskTransactionMetrics](u_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the connection operates over a cellular interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/3240627-cellular?language=objc
func (u_ URLSessionTaskTransactionMetrics) IsCellular() bool {
	rv := objc.Call[bool](u_, objc.Sel("isCellular"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/3674827-domainresolutionprotocol?language=objc
func (u_ URLSessionTaskTransactionMetrics) DomainResolutionProtocol() URLSessionTaskMetricsDomainResolutionProtocol {
	rv := objc.Call[URLSessionTaskMetricsDomainResolutionProtocol](u_, objc.Sel("domainResolutionProtocol"))
	return rv
}

// The transaction request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/1643144-request?language=objc
func (u_ URLSessionTaskTransactionMetrics) Request() URLRequest {
	rv := objc.Call[URLRequest](u_, objc.Sel("request"))
	return rv
}

// A Boolean value that indicates whether the connection operates over an expensive interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/3240635-expensive?language=objc
func (u_ URLSessionTaskTransactionMetrics) IsExpensive() bool {
	rv := objc.Call[bool](u_, objc.Sel("isExpensive"))
	return rv
}

// The TLS cipher suite the task negotiated with the endpoint for the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/3240639-negotiatedtlsciphersuite?language=objc
func (u_ URLSessionTaskTransactionMetrics) NegotiatedTLSCipherSuite() Number {
	rv := objc.Call[Number](u_, objc.Sel("negotiatedTLSCipherSuite"))
	return rv
}

// The time immediately before the task started establishing a TCP connection to the server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/1642815-connectstartdate?language=objc
func (u_ URLSessionTaskTransactionMetrics) ConnectStartDate() Date {
	rv := objc.Call[Date](u_, objc.Sel("connectStartDate"))
	return rv
}

// The number of bytes transferred for the request body. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/3240630-countofrequestbodybytessent?language=objc
func (u_ URLSessionTaskTransactionMetrics) CountOfRequestBodyBytesSent() int64 {
	rv := objc.Call[int64](u_, objc.Sel("countOfRequestBodyBytesSent"))
	return rv
}

// The IP address string of the remote interface for the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/3240642-remoteaddress?language=objc
func (u_ URLSessionTaskTransactionMetrics) RemoteAddress() string {
	rv := objc.Call[string](u_, objc.Sel("remoteAddress"))
	return rv
}

// The time immediately before the task started requesting the resource, regardless of whether it is retrieved from the server or local resources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/1642906-requeststartdate?language=objc
func (u_ URLSessionTaskTransactionMetrics) RequestStartDate() Date {
	rv := objc.Call[Date](u_, objc.Sel("requestStartDate"))
	return rv
}

// The time immediately before the task started the name lookup for the resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/1642859-domainlookupstartdate?language=objc
func (u_ URLSessionTaskTransactionMetrics) DomainLookupStartDate() Date {
	rv := objc.Call[Date](u_, objc.Sel("domainLookupStartDate"))
	return rv
}

// The size of the upload body data, file, or stream, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/3240629-countofrequestbodybytesbeforeenc?language=objc
func (u_ URLSessionTaskTransactionMetrics) CountOfRequestBodyBytesBeforeEncoding() int64 {
	rv := objc.Call[int64](u_, objc.Sel("countOfRequestBodyBytesBeforeEncoding"))
	return rv
}

// The number of bytes transferred for the request header. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/3240631-countofrequestheaderbytessent?language=objc
func (u_ URLSessionTaskTransactionMetrics) CountOfRequestHeaderBytesSent() int64 {
	rv := objc.Call[int64](u_, objc.Sel("countOfRequestHeaderBytesSent"))
	return rv
}

// The port number of the remote interface for the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/3240643-remoteport?language=objc
func (u_ URLSessionTaskTransactionMetrics) RemotePort() Number {
	rv := objc.Call[Number](u_, objc.Sel("remotePort"))
	return rv
}

// The number of bytes transferred for the response body. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/3240633-countofresponsebodybytesreceived?language=objc
func (u_ URLSessionTaskTransactionMetrics) CountOfResponseBodyBytesReceived() int64 {
	rv := objc.Call[int64](u_, objc.Sel("countOfResponseBodyBytesReceived"))
	return rv
}

// The time immediately after the task finished establishing the connection to the server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/1643239-connectenddate?language=objc
func (u_ URLSessionTaskTransactionMetrics) ConnectEndDate() Date {
	rv := objc.Call[Date](u_, objc.Sel("connectEndDate"))
	return rv
}

// The number of bytes transferred for the response header. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/3240634-countofresponseheaderbytesreceiv?language=objc
func (u_ URLSessionTaskTransactionMetrics) CountOfResponseHeaderBytesReceived() int64 {
	rv := objc.Call[int64](u_, objc.Sel("countOfResponseHeaderBytesReceived"))
	return rv
}

// The time immediately before the task started the TLS security handshake to secure the current connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/1643219-secureconnectionstartdate?language=objc
func (u_ URLSessionTaskTransactionMetrics) SecureConnectionStartDate() Date {
	rv := objc.Call[Date](u_, objc.Sel("secureConnectionStartDate"))
	return rv
}

// The size of data delivered to your delegate or completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/3240632-countofresponsebodybytesafterdec?language=objc
func (u_ URLSessionTaskTransactionMetrics) CountOfResponseBodyBytesAfterDecoding() int64 {
	rv := objc.Call[int64](u_, objc.Sel("countOfResponseBodyBytesAfterDecoding"))
	return rv
}

// The IP address string of the local interface for the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/3240636-localaddress?language=objc
func (u_ URLSessionTaskTransactionMetrics) LocalAddress() string {
	rv := objc.Call[string](u_, objc.Sel("localAddress"))
	return rv
}

// The time after the name lookup was completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/1643173-domainlookupenddate?language=objc
func (u_ URLSessionTaskTransactionMetrics) DomainLookupEndDate() Date {
	rv := objc.Call[Date](u_, objc.Sel("domainLookupEndDate"))
	return rv
}

// A Boolean value that indicates whether the task used a persistent connection to fetch the resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/1643233-reusedconnection?language=objc
func (u_ URLSessionTaskTransactionMetrics) IsReusedConnection() bool {
	rv := objc.Call[bool](u_, objc.Sel("isReusedConnection"))
	return rv
}

// A value that indicates whether the resource was loaded, pushed, or retrieved from the local cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/1642919-resourcefetchtype?language=objc
func (u_ URLSessionTaskTransactionMetrics) ResourceFetchType() URLSessionTaskMetricsResourceFetchType {
	rv := objc.Call[URLSessionTaskMetricsResourceFetchType](u_, objc.Sel("resourceFetchType"))
	return rv
}

// A Boolean value that indicates whether the connection uses a successfully negotiated multipath protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/3240638-multipath?language=objc
func (u_ URLSessionTaskTransactionMetrics) IsMultipath() bool {
	rv := objc.Call[bool](u_, objc.Sel("isMultipath"))
	return rv
}

// The time immediately after the task received the first byte of the response from the server or from local resources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/1642966-responsestartdate?language=objc
func (u_ URLSessionTaskTransactionMetrics) ResponseStartDate() Date {
	rv := objc.Call[Date](u_, objc.Sel("responseStartDate"))
	return rv
}

// The network protocol used to fetch the resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/1643141-networkprotocolname?language=objc
func (u_ URLSessionTaskTransactionMetrics) NetworkProtocolName() string {
	rv := objc.Call[string](u_, objc.Sel("networkProtocolName"))
	return rv
}

// The port number of the local interface for the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/3240637-localport?language=objc
func (u_ URLSessionTaskTransactionMetrics) LocalPort() Number {
	rv := objc.Call[Number](u_, objc.Sel("localPort"))
	return rv
}

// The time when the task started fetching the resource, from the server or locally. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/1643009-fetchstartdate?language=objc
func (u_ URLSessionTaskTransactionMetrics) FetchStartDate() Date {
	rv := objc.Call[Date](u_, objc.Sel("fetchStartDate"))
	return rv
}

// The TLS protocol version the task negotiated with the endpoint for the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/3240640-negotiatedtlsprotocolversion?language=objc
func (u_ URLSessionTaskTransactionMetrics) NegotiatedTLSProtocolVersion() Number {
	rv := objc.Call[Number](u_, objc.Sel("negotiatedTLSProtocolVersion"))
	return rv
}

// The time immediately after the task received the last byte of the resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/1643300-responseenddate?language=objc
func (u_ URLSessionTaskTransactionMetrics) ResponseEndDate() Date {
	rv := objc.Call[Date](u_, objc.Sel("responseEndDate"))
	return rv
}

// A Boolean value that indicates whether the connection operates over an interface marked as constrained. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/3240628-constrained?language=objc
func (u_ URLSessionTaskTransactionMetrics) IsConstrained() bool {
	rv := objc.Call[bool](u_, objc.Sel("isConstrained"))
	return rv
}

// A Boolean value that indicastes whether the task used a proxy connection to fetch the resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/1642917-proxyconnection?language=objc
func (u_ URLSessionTaskTransactionMetrics) IsProxyConnection() bool {
	rv := objc.Call[bool](u_, objc.Sel("isProxyConnection"))
	return rv
}

// The transaction response. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/1643172-response?language=objc
func (u_ URLSessionTaskTransactionMetrics) Response() URLResponse {
	rv := objc.Call[URLResponse](u_, objc.Sel("response"))
	return rv
}

// The time immediately after the security handshake completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/1643255-secureconnectionenddate?language=objc
func (u_ URLSessionTaskTransactionMetrics) SecureConnectionEndDate() Date {
	rv := objc.Call[Date](u_, objc.Sel("secureConnectionEndDate"))
	return rv
}

// The time immediately after the task finished requesting the resource, regardless of whether it was retrieved from the server or local resources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontasktransactionmetrics/1643056-requestenddate?language=objc
func (u_ URLSessionTaskTransactionMetrics) RequestEndDate() Date {
	rv := objc.Call[Date](u_, objc.Sel("requestEndDate"))
	return rv
}
