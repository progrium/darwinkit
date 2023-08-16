// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLSessionConfiguration] class.
var URLSessionConfigurationClass = _URLSessionConfigurationClass{objc.GetClass("NSURLSessionConfiguration")}

type _URLSessionConfigurationClass struct {
	objc.Class
}

// An interface definition for the [URLSessionConfiguration] class.
type IURLSessionConfiguration interface {
	objc.IObject
	SharedContainerIdentifier() string
	SetSharedContainerIdentifier(value string)
	WaitsForConnectivity() bool
	SetWaitsForConnectivity(value bool)
	RequestCachePolicy() URLRequestCachePolicy
	SetRequestCachePolicy(value URLRequestCachePolicy)
	ShouldUseExtendedBackgroundIdleMode() bool
	SetShouldUseExtendedBackgroundIdleMode(value bool)
	HTTPAdditionalHeaders() Dictionary
	SetHTTPAdditionalHeaders(value Dictionary)
	AllowsCellularAccess() bool
	SetAllowsCellularAccess(value bool)
	HTTPCookieStorage() HTTPCookieStorage
	SetHTTPCookieStorage(value IHTTPCookieStorage)
	HTTPCookieAcceptPolicy() HTTPCookieAcceptPolicy
	SetHTTPCookieAcceptPolicy(value HTTPCookieAcceptPolicy)
	URLCache() URLCache
	SetURLCache(value IURLCache)
	TimeoutIntervalForRequest() TimeInterval
	SetTimeoutIntervalForRequest(value TimeInterval)
	IsDiscretionary() bool
	SetDiscretionary(value bool)
	AllowsExpensiveNetworkAccess() bool
	SetAllowsExpensiveNetworkAccess(value bool)
	AllowsConstrainedNetworkAccess() bool
	SetAllowsConstrainedNetworkAccess(value bool)
	ConnectionProxyDictionary() Dictionary
	SetConnectionProxyDictionary(value Dictionary)
	HTTPShouldSetCookies() bool
	SetHTTPShouldSetCookies(value bool)
	HTTPMaximumConnectionsPerHost() int
	SetHTTPMaximumConnectionsPerHost(value int)
	URLCredentialStorage() URLCredentialStorage
	SetURLCredentialStorage(value IURLCredentialStorage)
	ProtocolClasses() []objc.Class
	SetProtocolClasses(value []objc.IClass)
	NetworkServiceType() URLRequestNetworkServiceType
	SetNetworkServiceType(value URLRequestNetworkServiceType)
	HTTPShouldUsePipelining() bool
	SetHTTPShouldUsePipelining(value bool)
	SessionSendsLaunchEvents() bool
	SetSessionSendsLaunchEvents(value bool)
	Identifier() string
	TimeoutIntervalForResource() TimeInterval
	SetTimeoutIntervalForResource(value TimeInterval)
}

// A configuration object that defines behavior and policies for a URL session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration?language=objc
type URLSessionConfiguration struct {
	objc.Object
}

func URLSessionConfigurationFrom(ptr unsafe.Pointer) URLSessionConfiguration {
	return URLSessionConfiguration{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _URLSessionConfigurationClass) Alloc() URLSessionConfiguration {
	rv := objc.Call[URLSessionConfiguration](uc, objc.Sel("alloc"))
	return rv
}

func URLSessionConfiguration_Alloc() URLSessionConfiguration {
	return URLSessionConfigurationClass.Alloc()
}

func (uc _URLSessionConfigurationClass) New() URLSessionConfiguration {
	rv := objc.Call[URLSessionConfiguration](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLSessionConfiguration() URLSessionConfiguration {
	return URLSessionConfigurationClass.New()
}

func (u_ URLSessionConfiguration) Init() URLSessionConfiguration {
	rv := objc.Call[URLSessionConfiguration](u_, objc.Sel("init"))
	return rv
}

// Creates a session configuration object that allows HTTP and HTTPS uploads or downloads to be performed in the background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1407496-backgroundsessionconfigurationwi?language=objc
func (uc _URLSessionConfigurationClass) BackgroundSessionConfigurationWithIdentifier(identifier string) URLSessionConfiguration {
	rv := objc.Call[URLSessionConfiguration](uc, objc.Sel("backgroundSessionConfigurationWithIdentifier:"), identifier)
	return rv
}

// Creates a session configuration object that allows HTTP and HTTPS uploads or downloads to be performed in the background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1407496-backgroundsessionconfigurationwi?language=objc
func URLSessionConfiguration_BackgroundSessionConfigurationWithIdentifier(identifier string) URLSessionConfiguration {
	return URLSessionConfigurationClass.BackgroundSessionConfigurationWithIdentifier(identifier)
}

// The identifier for the shared container into which files in background URL sessions should be downloaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1409450-sharedcontaineridentifier?language=objc
func (u_ URLSessionConfiguration) SharedContainerIdentifier() string {
	rv := objc.Call[string](u_, objc.Sel("sharedContainerIdentifier"))
	return rv
}

// The identifier for the shared container into which files in background URL sessions should be downloaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1409450-sharedcontaineridentifier?language=objc
func (u_ URLSessionConfiguration) SetSharedContainerIdentifier(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setSharedContainerIdentifier:"), value)
}

// A Boolean value that indicates whether the session should wait for connectivity to become available, or fail immediately. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/2908812-waitsforconnectivity?language=objc
func (u_ URLSessionConfiguration) WaitsForConnectivity() bool {
	rv := objc.Call[bool](u_, objc.Sel("waitsForConnectivity"))
	return rv
}

// A Boolean value that indicates whether the session should wait for connectivity to become available, or fail immediately. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/2908812-waitsforconnectivity?language=objc
func (u_ URLSessionConfiguration) SetWaitsForConnectivity(value bool) {
	objc.Call[objc.Void](u_, objc.Sel("setWaitsForConnectivity:"), value)
}

// A predefined constant that determines when to return a response from the cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411655-requestcachepolicy?language=objc
func (u_ URLSessionConfiguration) RequestCachePolicy() URLRequestCachePolicy {
	rv := objc.Call[URLRequestCachePolicy](u_, objc.Sel("requestCachePolicy"))
	return rv
}

// A predefined constant that determines when to return a response from the cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411655-requestcachepolicy?language=objc
func (u_ URLSessionConfiguration) SetRequestCachePolicy(value URLRequestCachePolicy) {
	objc.Call[objc.Void](u_, objc.Sel("setRequestCachePolicy:"), value)
}

// A Boolean value that indicates whether TCP connections should be kept open when the app moves to the background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1409517-shoulduseextendedbackgroundidlem?language=objc
func (u_ URLSessionConfiguration) ShouldUseExtendedBackgroundIdleMode() bool {
	rv := objc.Call[bool](u_, objc.Sel("shouldUseExtendedBackgroundIdleMode"))
	return rv
}

// A Boolean value that indicates whether TCP connections should be kept open when the app moves to the background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1409517-shoulduseextendedbackgroundidlem?language=objc
func (u_ URLSessionConfiguration) SetShouldUseExtendedBackgroundIdleMode(value bool) {
	objc.Call[objc.Void](u_, objc.Sel("setShouldUseExtendedBackgroundIdleMode:"), value)
}

// A dictionary of additional headers to send with requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411532-httpadditionalheaders?language=objc
func (u_ URLSessionConfiguration) HTTPAdditionalHeaders() Dictionary {
	rv := objc.Call[Dictionary](u_, objc.Sel("HTTPAdditionalHeaders"))
	return rv
}

// A dictionary of additional headers to send with requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411532-httpadditionalheaders?language=objc
func (u_ URLSessionConfiguration) SetHTTPAdditionalHeaders(value Dictionary) {
	objc.Call[objc.Void](u_, objc.Sel("setHTTPAdditionalHeaders:"), value)
}

// A Boolean value that determines whether connections should be made over a cellular network. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1409406-allowscellularaccess?language=objc
func (u_ URLSessionConfiguration) AllowsCellularAccess() bool {
	rv := objc.Call[bool](u_, objc.Sel("allowsCellularAccess"))
	return rv
}

// A Boolean value that determines whether connections should be made over a cellular network. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1409406-allowscellularaccess?language=objc
func (u_ URLSessionConfiguration) SetAllowsCellularAccess(value bool) {
	objc.Call[objc.Void](u_, objc.Sel("setAllowsCellularAccess:"), value)
}

// A default session configuration object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411560-defaultsessionconfiguration?language=objc
func (uc _URLSessionConfigurationClass) DefaultSessionConfiguration() URLSessionConfiguration {
	rv := objc.Call[URLSessionConfiguration](uc, objc.Sel("defaultSessionConfiguration"))
	return rv
}

// A default session configuration object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411560-defaultsessionconfiguration?language=objc
func URLSessionConfiguration_DefaultSessionConfiguration() URLSessionConfiguration {
	return URLSessionConfigurationClass.DefaultSessionConfiguration()
}

// The cookie store for storing cookies within this session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411599-httpcookiestorage?language=objc
func (u_ URLSessionConfiguration) HTTPCookieStorage() HTTPCookieStorage {
	rv := objc.Call[HTTPCookieStorage](u_, objc.Sel("HTTPCookieStorage"))
	return rv
}

// The cookie store for storing cookies within this session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411599-httpcookiestorage?language=objc
func (u_ URLSessionConfiguration) SetHTTPCookieStorage(value IHTTPCookieStorage) {
	objc.Call[objc.Void](u_, objc.Sel("setHTTPCookieStorage:"), objc.Ptr(value))
}

// A policy constant that determines when cookies should be accepted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1408933-httpcookieacceptpolicy?language=objc
func (u_ URLSessionConfiguration) HTTPCookieAcceptPolicy() HTTPCookieAcceptPolicy {
	rv := objc.Call[HTTPCookieAcceptPolicy](u_, objc.Sel("HTTPCookieAcceptPolicy"))
	return rv
}

// A policy constant that determines when cookies should be accepted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1408933-httpcookieacceptpolicy?language=objc
func (u_ URLSessionConfiguration) SetHTTPCookieAcceptPolicy(value HTTPCookieAcceptPolicy) {
	objc.Call[objc.Void](u_, objc.Sel("setHTTPCookieAcceptPolicy:"), value)
}

// The URL cache for providing cached responses to requests within the session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1410148-urlcache?language=objc
func (u_ URLSessionConfiguration) URLCache() URLCache {
	rv := objc.Call[URLCache](u_, objc.Sel("URLCache"))
	return rv
}

// The URL cache for providing cached responses to requests within the session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1410148-urlcache?language=objc
func (u_ URLSessionConfiguration) SetURLCache(value IURLCache) {
	objc.Call[objc.Void](u_, objc.Sel("setURLCache:"), objc.Ptr(value))
}

// The timeout interval to use when waiting for additional data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1408259-timeoutintervalforrequest?language=objc
func (u_ URLSessionConfiguration) TimeoutIntervalForRequest() TimeInterval {
	rv := objc.Call[TimeInterval](u_, objc.Sel("timeoutIntervalForRequest"))
	return rv
}

// The timeout interval to use when waiting for additional data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1408259-timeoutintervalforrequest?language=objc
func (u_ URLSessionConfiguration) SetTimeoutIntervalForRequest(value TimeInterval) {
	objc.Call[objc.Void](u_, objc.Sel("setTimeoutIntervalForRequest:"), value)
}

// A Boolean value that determines whether background tasks can be scheduled at the discretion of the system for optimal performance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411552-discretionary?language=objc
func (u_ URLSessionConfiguration) IsDiscretionary() bool {
	rv := objc.Call[bool](u_, objc.Sel("isDiscretionary"))
	return rv
}

// A Boolean value that determines whether background tasks can be scheduled at the discretion of the system for optimal performance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411552-discretionary?language=objc
func (u_ URLSessionConfiguration) SetDiscretionary(value bool) {
	objc.Call[objc.Void](u_, objc.Sel("setDiscretionary:"), value)
}

// A session configuration that uses no persistent storage for caches, cookies, or credentials. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1410529-ephemeralsessionconfiguration?language=objc
func (uc _URLSessionConfigurationClass) EphemeralSessionConfiguration() URLSessionConfiguration {
	rv := objc.Call[URLSessionConfiguration](uc, objc.Sel("ephemeralSessionConfiguration"))
	return rv
}

// A session configuration that uses no persistent storage for caches, cookies, or credentials. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1410529-ephemeralsessionconfiguration?language=objc
func URLSessionConfiguration_EphemeralSessionConfiguration() URLSessionConfiguration {
	return URLSessionConfigurationClass.EphemeralSessionConfiguration()
}

// A Boolean value that indicates whether connections may use a network interface that the system considers expensive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/3235752-allowsexpensivenetworkaccess?language=objc
func (u_ URLSessionConfiguration) AllowsExpensiveNetworkAccess() bool {
	rv := objc.Call[bool](u_, objc.Sel("allowsExpensiveNetworkAccess"))
	return rv
}

// A Boolean value that indicates whether connections may use a network interface that the system considers expensive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/3235752-allowsexpensivenetworkaccess?language=objc
func (u_ URLSessionConfiguration) SetAllowsExpensiveNetworkAccess(value bool) {
	objc.Call[objc.Void](u_, objc.Sel("setAllowsExpensiveNetworkAccess:"), value)
}

// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/3235751-allowsconstrainednetworkaccess?language=objc
func (u_ URLSessionConfiguration) AllowsConstrainedNetworkAccess() bool {
	rv := objc.Call[bool](u_, objc.Sel("allowsConstrainedNetworkAccess"))
	return rv
}

// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/3235751-allowsconstrainednetworkaccess?language=objc
func (u_ URLSessionConfiguration) SetAllowsConstrainedNetworkAccess(value bool) {
	objc.Call[objc.Void](u_, objc.Sel("setAllowsConstrainedNetworkAccess:"), value)
}

// A dictionary containing information about the proxy to use within this session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411499-connectionproxydictionary?language=objc
func (u_ URLSessionConfiguration) ConnectionProxyDictionary() Dictionary {
	rv := objc.Call[Dictionary](u_, objc.Sel("connectionProxyDictionary"))
	return rv
}

// A dictionary containing information about the proxy to use within this session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411499-connectionproxydictionary?language=objc
func (u_ URLSessionConfiguration) SetConnectionProxyDictionary(value Dictionary) {
	objc.Call[objc.Void](u_, objc.Sel("setConnectionProxyDictionary:"), value)
}

// A Boolean value that determines whether requests should contain cookies from the cookie store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411589-httpshouldsetcookies?language=objc
func (u_ URLSessionConfiguration) HTTPShouldSetCookies() bool {
	rv := objc.Call[bool](u_, objc.Sel("HTTPShouldSetCookies"))
	return rv
}

// A Boolean value that determines whether requests should contain cookies from the cookie store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411589-httpshouldsetcookies?language=objc
func (u_ URLSessionConfiguration) SetHTTPShouldSetCookies(value bool) {
	objc.Call[objc.Void](u_, objc.Sel("setHTTPShouldSetCookies:"), value)
}

// The maximum number of simultaneous connections to make to a given host. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1407597-httpmaximumconnectionsperhost?language=objc
func (u_ URLSessionConfiguration) HTTPMaximumConnectionsPerHost() int {
	rv := objc.Call[int](u_, objc.Sel("HTTPMaximumConnectionsPerHost"))
	return rv
}

// The maximum number of simultaneous connections to make to a given host. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1407597-httpmaximumconnectionsperhost?language=objc
func (u_ URLSessionConfiguration) SetHTTPMaximumConnectionsPerHost(value int) {
	objc.Call[objc.Void](u_, objc.Sel("setHTTPMaximumConnectionsPerHost:"), value)
}

// A credential store that provides credentials for authentication. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1410947-urlcredentialstorage?language=objc
func (u_ URLSessionConfiguration) URLCredentialStorage() URLCredentialStorage {
	rv := objc.Call[URLCredentialStorage](u_, objc.Sel("URLCredentialStorage"))
	return rv
}

// A credential store that provides credentials for authentication. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1410947-urlcredentialstorage?language=objc
func (u_ URLSessionConfiguration) SetURLCredentialStorage(value IURLCredentialStorage) {
	objc.Call[objc.Void](u_, objc.Sel("setURLCredentialStorage:"), objc.Ptr(value))
}

// An array of extra protocol subclasses that handle requests in a session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411050-protocolclasses?language=objc
func (u_ URLSessionConfiguration) ProtocolClasses() []objc.Class {
	rv := objc.Call[[]objc.Class](u_, objc.Sel("protocolClasses"))
	return rv
}

// An array of extra protocol subclasses that handle requests in a session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411050-protocolclasses?language=objc
func (u_ URLSessionConfiguration) SetProtocolClasses(value []objc.IClass) {
	objc.Call[objc.Void](u_, objc.Sel("setProtocolClasses:"), value)
}

// The type of network service for all tasks within sessions based on this configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411606-networkservicetype?language=objc
func (u_ URLSessionConfiguration) NetworkServiceType() URLRequestNetworkServiceType {
	rv := objc.Call[URLRequestNetworkServiceType](u_, objc.Sel("networkServiceType"))
	return rv
}

// The type of network service for all tasks within sessions based on this configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411606-networkservicetype?language=objc
func (u_ URLSessionConfiguration) SetNetworkServiceType(value URLRequestNetworkServiceType) {
	objc.Call[objc.Void](u_, objc.Sel("setNetworkServiceType:"), value)
}

// A Boolean value that determines whether the session should use HTTP pipelining. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411657-httpshouldusepipelining?language=objc
func (u_ URLSessionConfiguration) HTTPShouldUsePipelining() bool {
	rv := objc.Call[bool](u_, objc.Sel("HTTPShouldUsePipelining"))
	return rv
}

// A Boolean value that determines whether the session should use HTTP pipelining. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1411657-httpshouldusepipelining?language=objc
func (u_ URLSessionConfiguration) SetHTTPShouldUsePipelining(value bool) {
	objc.Call[objc.Void](u_, objc.Sel("setHTTPShouldUsePipelining:"), value)
}

// A Boolean value that indicates whether the app should be resumed or launched in the background when transfers finish. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1617174-sessionsendslaunchevents?language=objc
func (u_ URLSessionConfiguration) SessionSendsLaunchEvents() bool {
	rv := objc.Call[bool](u_, objc.Sel("sessionSendsLaunchEvents"))
	return rv
}

// A Boolean value that indicates whether the app should be resumed or launched in the background when transfers finish. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1617174-sessionsendslaunchevents?language=objc
func (u_ URLSessionConfiguration) SetSessionSendsLaunchEvents(value bool) {
	objc.Call[objc.Void](u_, objc.Sel("setSessionSendsLaunchEvents:"), value)
}

// The background session identifier of the configuration object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1408987-identifier?language=objc
func (u_ URLSessionConfiguration) Identifier() string {
	rv := objc.Call[string](u_, objc.Sel("identifier"))
	return rv
}

// The maximum amount of time that a resource request should be allowed to take. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1408153-timeoutintervalforresource?language=objc
func (u_ URLSessionConfiguration) TimeoutIntervalForResource() TimeInterval {
	rv := objc.Call[TimeInterval](u_, objc.Sel("timeoutIntervalForResource"))
	return rv
}

// The maximum amount of time that a resource request should be allowed to take. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionconfiguration/1408153-timeoutintervalforresource?language=objc
func (u_ URLSessionConfiguration) SetTimeoutIntervalForResource(value TimeInterval) {
	objc.Call[objc.Void](u_, objc.Sel("setTimeoutIntervalForResource:"), value)
}
