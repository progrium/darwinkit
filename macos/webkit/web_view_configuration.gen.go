// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var WebViewConfigurationClass = _WebViewConfigurationClass{objc.GetClass("WKWebViewConfiguration")}

type _WebViewConfigurationClass struct {
	objc.Class
}

type IWebViewConfiguration interface {
	objc.IObject
	SetURLSchemeHandlerForURLScheme(urlSchemeHandler IURLSchemeHandler, urlScheme string)
	SetURLSchemeHandler0ForURLScheme(urlSchemeHandler objc.IObject, urlScheme string)
	UrlSchemeHandlerForURLScheme(urlScheme string) URLSchemeHandlerWrapper
	WebsiteDataStore() WebsiteDataStore
	SetWebsiteDataStore(value IWebsiteDataStore)
	UserContentController() UserContentController
	SetUserContentController(value IUserContentController)
	ProcessPool() ProcessPool
	SetProcessPool(value IProcessPool)
	ApplicationNameForUserAgent() string
	SetApplicationNameForUserAgent(value string)
	LimitsNavigationsToAppBoundDomains() bool
	SetLimitsNavigationsToAppBoundDomains(value bool)
	Preferences() Preferences
	SetPreferences(value IPreferences)
	DefaultWebpagePreferences() WebpagePreferences
	SetDefaultWebpagePreferences(value IWebpagePreferences)
	SuppressesIncrementalRendering() bool
	SetSuppressesIncrementalRendering(value bool)
	AllowsAirPlayForMediaPlayback() bool
	SetAllowsAirPlayForMediaPlayback(value bool)
	MediaTypesRequiringUserActionForPlayback() AudiovisualMediaTypes
	SetMediaTypesRequiringUserActionForPlayback(value AudiovisualMediaTypes)
	UserInterfaceDirectionPolicy() UserInterfaceDirectionPolicy
	SetUserInterfaceDirectionPolicy(value UserInterfaceDirectionPolicy)
	UpgradeKnownHostsToHTTPS() bool
	SetUpgradeKnownHostsToHTTPS(value bool)
}

type WebViewConfiguration struct {
	objc.Object
}

func MakeWebViewConfiguration(ptr unsafe.Pointer) WebViewConfiguration {
	return WebViewConfiguration{
		Object: objc.MakeObject(ptr),
	}
}

func (wc _WebViewConfigurationClass) Alloc() WebViewConfiguration {
	rv := objc.CallMethod[WebViewConfiguration](wc, objc.GetSelector("alloc"))
	return rv
}

func WebViewConfiguration_Alloc() WebViewConfiguration {
	return WebViewConfigurationClass.Alloc()
}

func (wc _WebViewConfigurationClass) New() WebViewConfiguration {
	rv := objc.CallMethod[WebViewConfiguration](wc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewWebViewConfiguration() WebViewConfiguration {
	return WebViewConfigurationClass.New()
}

func WebViewConfiguration_New() WebViewConfiguration {
	return WebViewConfigurationClass.New()
}

func (w_ WebViewConfiguration) Init() WebViewConfiguration {
	rv := objc.CallMethod[WebViewConfiguration](w_, objc.GetSelector("init"))
	return rv
}

func WebViewConfiguration_Init() WebViewConfiguration {
	return WebViewConfigurationClass.Alloc().Init()
}

func (w_ WebViewConfiguration) SetURLSchemeHandlerForURLScheme(urlSchemeHandler IURLSchemeHandler, urlScheme string) {
	po := objc.WrapAsProtocol("WKURLSchemeHandler", urlSchemeHandler)
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setURLSchemeHandler:forURLScheme:"), po, urlScheme)
}

func (w_ WebViewConfiguration) SetURLSchemeHandler0ForURLScheme(urlSchemeHandler objc.IObject, urlScheme string) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setURLSchemeHandler:forURLScheme:"), objc.ExtractPtr(urlSchemeHandler), urlScheme)
}

func (w_ WebViewConfiguration) UrlSchemeHandlerForURLScheme(urlScheme string) URLSchemeHandlerWrapper {
	rv := objc.CallMethod[URLSchemeHandlerWrapper](w_, objc.GetSelector("urlSchemeHandlerForURLScheme:"), urlScheme)
	return rv
}

func (w_ WebViewConfiguration) WebsiteDataStore() WebsiteDataStore {
	rv := objc.CallMethod[WebsiteDataStore](w_, objc.GetSelector("websiteDataStore"))
	return rv
}

func (w_ WebViewConfiguration) SetWebsiteDataStore(value IWebsiteDataStore) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setWebsiteDataStore:"), objc.ExtractPtr(value))
}

func (w_ WebViewConfiguration) UserContentController() UserContentController {
	rv := objc.CallMethod[UserContentController](w_, objc.GetSelector("userContentController"))
	return rv
}

func (w_ WebViewConfiguration) SetUserContentController(value IUserContentController) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setUserContentController:"), objc.ExtractPtr(value))
}

func (w_ WebViewConfiguration) ProcessPool() ProcessPool {
	rv := objc.CallMethod[ProcessPool](w_, objc.GetSelector("processPool"))
	return rv
}

func (w_ WebViewConfiguration) SetProcessPool(value IProcessPool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setProcessPool:"), objc.ExtractPtr(value))
}

func (w_ WebViewConfiguration) ApplicationNameForUserAgent() string {
	rv := objc.CallMethod[string](w_, objc.GetSelector("applicationNameForUserAgent"))
	return rv
}

func (w_ WebViewConfiguration) SetApplicationNameForUserAgent(value string) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setApplicationNameForUserAgent:"), value)
}

func (w_ WebViewConfiguration) LimitsNavigationsToAppBoundDomains() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("limitsNavigationsToAppBoundDomains"))
	return rv
}

func (w_ WebViewConfiguration) SetLimitsNavigationsToAppBoundDomains(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setLimitsNavigationsToAppBoundDomains:"), value)
}

func (w_ WebViewConfiguration) Preferences() Preferences {
	rv := objc.CallMethod[Preferences](w_, objc.GetSelector("preferences"))
	return rv
}

func (w_ WebViewConfiguration) SetPreferences(value IPreferences) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setPreferences:"), objc.ExtractPtr(value))
}

func (w_ WebViewConfiguration) DefaultWebpagePreferences() WebpagePreferences {
	rv := objc.CallMethod[WebpagePreferences](w_, objc.GetSelector("defaultWebpagePreferences"))
	return rv
}

func (w_ WebViewConfiguration) SetDefaultWebpagePreferences(value IWebpagePreferences) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setDefaultWebpagePreferences:"), objc.ExtractPtr(value))
}

func (w_ WebViewConfiguration) SuppressesIncrementalRendering() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("suppressesIncrementalRendering"))
	return rv
}

func (w_ WebViewConfiguration) SetSuppressesIncrementalRendering(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setSuppressesIncrementalRendering:"), value)
}

func (w_ WebViewConfiguration) AllowsAirPlayForMediaPlayback() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("allowsAirPlayForMediaPlayback"))
	return rv
}

func (w_ WebViewConfiguration) SetAllowsAirPlayForMediaPlayback(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setAllowsAirPlayForMediaPlayback:"), value)
}

func (w_ WebViewConfiguration) MediaTypesRequiringUserActionForPlayback() AudiovisualMediaTypes {
	rv := objc.CallMethod[AudiovisualMediaTypes](w_, objc.GetSelector("mediaTypesRequiringUserActionForPlayback"))
	return rv
}

func (w_ WebViewConfiguration) SetMediaTypesRequiringUserActionForPlayback(value AudiovisualMediaTypes) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setMediaTypesRequiringUserActionForPlayback:"), value)
}

func (w_ WebViewConfiguration) UserInterfaceDirectionPolicy() UserInterfaceDirectionPolicy {
	rv := objc.CallMethod[UserInterfaceDirectionPolicy](w_, objc.GetSelector("userInterfaceDirectionPolicy"))
	return rv
}

func (w_ WebViewConfiguration) SetUserInterfaceDirectionPolicy(value UserInterfaceDirectionPolicy) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setUserInterfaceDirectionPolicy:"), value)
}

func (w_ WebViewConfiguration) UpgradeKnownHostsToHTTPS() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("upgradeKnownHostsToHTTPS"))
	return rv
}

func (w_ WebViewConfiguration) SetUpgradeKnownHostsToHTTPS(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setUpgradeKnownHostsToHTTPS:"), value)
}
