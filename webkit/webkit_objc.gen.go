package webkit

import (
	cocoa "github.com/progrium/macdriver/cocoa"
	core "github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"unsafe"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -lobjc -framework WebKit
#define __OBJC2__ 1
#include <objc/message.h>
#include <stdlib.h>

#include <WebKit/WebKit.h>

bool webkit_convertObjCBool(BOOL b) {
	if (b) { return true; }
	return false;
}


void* WKNavigation_type_alloc() {
	return [WKNavigation
		alloc];
}
void* WKUserScript_type_alloc() {
	return [WKUserScript
		alloc];
}
void* WKWebView_type_alloc() {
	return [WKWebView
		alloc];
}
BOOL WKWebView_type_handlesURLScheme_(void* urlScheme) {
	return [WKWebView
		handlesURLScheme: urlScheme];
}
void* WKWebViewConfiguration_type_alloc() {
	return [WKWebViewConfiguration
		alloc];
}
void* WKPreferences_type_alloc() {
	return [WKPreferences
		alloc];
}


void* WKNavigation_inst_init(void *id) {
	return [(WKNavigation*)id
		init];
}

void* WKUserScript_inst_init(void *id) {
	return [(WKUserScript*)id
		init];
}

void* WKUserScript_inst_source(void *id) {
	return [(WKUserScript*)id
		source];
}

BOOL WKUserScript_inst_isForMainFrameOnly(void *id) {
	return [(WKUserScript*)id
		isForMainFrameOnly];
}

void* WKWebView_inst_initWithFrame_configuration_(void *id, NSRect frame, void* configuration) {
	return [(WKWebView*)id
		initWithFrame: frame
		configuration: configuration];
}

void* WKWebView_inst_loadRequest_(void *id, void* request) {
	return [(WKWebView*)id
		loadRequest: request];
}

void* WKWebView_inst_loadHTMLString_baseURL_(void *id, void* string, void* baseURL) {
	return [(WKWebView*)id
		loadHTMLString: string
		baseURL: baseURL];
}

void* WKWebView_inst_loadFileURL_allowingReadAccessToURL_(void *id, void* URL, void* readAccessURL) {
	return [(WKWebView*)id
		loadFileURL: URL
		allowingReadAccessToURL: readAccessURL];
}

void* WKWebView_inst_loadData_MIMEType_characterEncodingName_baseURL_(void *id, void* data, void* MIMEType, void* characterEncodingName, void* baseURL) {
	return [(WKWebView*)id
		loadData: data
		MIMEType: MIMEType
		characterEncodingName: characterEncodingName
		baseURL: baseURL];
}

void* WKWebView_inst_reload(void *id) {
	return [(WKWebView*)id
		reload];
}

void WKWebView_inst_reload_(void *id, void* sender) {
	[(WKWebView*)id
		reload: sender];
}

void* WKWebView_inst_reloadFromOrigin(void *id) {
	return [(WKWebView*)id
		reloadFromOrigin];
}

void WKWebView_inst_reloadFromOrigin_(void *id, void* sender) {
	[(WKWebView*)id
		reloadFromOrigin: sender];
}

void WKWebView_inst_stopLoading(void *id) {
	[(WKWebView*)id
		stopLoading];
}

void WKWebView_inst_stopLoading_(void *id, void* sender) {
	[(WKWebView*)id
		stopLoading: sender];
}

void WKWebView_inst_goBack_(void *id, void* sender) {
	[(WKWebView*)id
		goBack: sender];
}

void* WKWebView_inst_goBack(void *id) {
	return [(WKWebView*)id
		goBack];
}

void WKWebView_inst_goForward_(void *id, void* sender) {
	[(WKWebView*)id
		goForward: sender];
}

void* WKWebView_inst_goForward(void *id) {
	return [(WKWebView*)id
		goForward];
}

void* WKWebView_inst_loadFileRequest_allowingReadAccessToURL_(void *id, void* request, void* readAccessURL) {
	return [(WKWebView*)id
		loadFileRequest: request
		allowingReadAccessToURL: readAccessURL];
}

void* WKWebView_inst_loadSimulatedRequest_responseHTMLString_(void *id, void* request, void* string) {
	return [(WKWebView*)id
		loadSimulatedRequest: request
		responseHTMLString: string];
}

void* WKWebView_inst_init(void *id) {
	return [(WKWebView*)id
		init];
}

void* WKWebView_inst_configuration(void *id) {
	return [(WKWebView*)id
		configuration];
}

void* WKWebView_inst_UIDelegate(void *id) {
	return [(WKWebView*)id
		UIDelegate];
}

void WKWebView_inst_setUIDelegate_(void *id, void* value) {
	[(WKWebView*)id
		setUIDelegate: value];
}

void* WKWebView_inst_navigationDelegate(void *id) {
	return [(WKWebView*)id
		navigationDelegate];
}

void WKWebView_inst_setNavigationDelegate_(void *id, void* value) {
	[(WKWebView*)id
		setNavigationDelegate: value];
}

BOOL WKWebView_inst_isLoading(void *id) {
	return [(WKWebView*)id
		isLoading];
}

void* WKWebView_inst_title(void *id) {
	return [(WKWebView*)id
		title];
}

void* WKWebView_inst_URL(void *id) {
	return [(WKWebView*)id
		URL];
}

void* WKWebView_inst_mediaType(void *id) {
	return [(WKWebView*)id
		mediaType];
}

void WKWebView_inst_setMediaType_(void *id, void* value) {
	[(WKWebView*)id
		setMediaType: value];
}

void* WKWebView_inst_customUserAgent(void *id) {
	return [(WKWebView*)id
		customUserAgent];
}

void WKWebView_inst_setCustomUserAgent_(void *id, void* value) {
	[(WKWebView*)id
		setCustomUserAgent: value];
}

BOOL WKWebView_inst_hasOnlySecureContent(void *id) {
	return [(WKWebView*)id
		hasOnlySecureContent];
}

BOOL WKWebView_inst_allowsMagnification(void *id) {
	return [(WKWebView*)id
		allowsMagnification];
}

void WKWebView_inst_setAllowsMagnification_(void *id, BOOL value) {
	[(WKWebView*)id
		setAllowsMagnification: value];
}

double WKWebView_inst_magnification(void *id) {
	return [(WKWebView*)id
		magnification];
}

void WKWebView_inst_setMagnification_(void *id, double value) {
	[(WKWebView*)id
		setMagnification: value];
}

BOOL WKWebView_inst_allowsBackForwardNavigationGestures(void *id) {
	return [(WKWebView*)id
		allowsBackForwardNavigationGestures];
}

void WKWebView_inst_setAllowsBackForwardNavigationGestures_(void *id, BOOL value) {
	[(WKWebView*)id
		setAllowsBackForwardNavigationGestures: value];
}

BOOL WKWebView_inst_canGoBack(void *id) {
	return [(WKWebView*)id
		canGoBack];
}

BOOL WKWebView_inst_canGoForward(void *id) {
	return [(WKWebView*)id
		canGoForward];
}

BOOL WKWebView_inst_allowsLinkPreview(void *id) {
	return [(WKWebView*)id
		allowsLinkPreview];
}

void WKWebView_inst_setAllowsLinkPreview_(void *id, BOOL value) {
	[(WKWebView*)id
		setAllowsLinkPreview: value];
}

void* WKWebView_inst_interactionState(void *id) {
	return [(WKWebView*)id
		interactionState];
}

void WKWebView_inst_setInteractionState_(void *id, void* value) {
	[(WKWebView*)id
		setInteractionState: value];
}

void WKWebViewConfiguration_inst_setURLSchemeHandler_forURLScheme_(void *id, void* urlSchemeHandler, void* urlScheme) {
	[(WKWebViewConfiguration*)id
		setURLSchemeHandler: urlSchemeHandler
		forURLScheme: urlScheme];
}

void* WKWebViewConfiguration_inst_urlSchemeHandlerForURLScheme_(void *id, void* urlScheme) {
	return [(WKWebViewConfiguration*)id
		urlSchemeHandlerForURLScheme: urlScheme];
}

void* WKWebViewConfiguration_inst_init(void *id) {
	return [(WKWebViewConfiguration*)id
		init];
}

void* WKWebViewConfiguration_inst_applicationNameForUserAgent(void *id) {
	return [(WKWebViewConfiguration*)id
		applicationNameForUserAgent];
}

void WKWebViewConfiguration_inst_setApplicationNameForUserAgent_(void *id, void* value) {
	[(WKWebViewConfiguration*)id
		setApplicationNameForUserAgent: value];
}

BOOL WKWebViewConfiguration_inst_limitsNavigationsToAppBoundDomains(void *id) {
	return [(WKWebViewConfiguration*)id
		limitsNavigationsToAppBoundDomains];
}

void WKWebViewConfiguration_inst_setLimitsNavigationsToAppBoundDomains_(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setLimitsNavigationsToAppBoundDomains: value];
}

void* WKWebViewConfiguration_inst_preferences(void *id) {
	return [(WKWebViewConfiguration*)id
		preferences];
}

void WKWebViewConfiguration_inst_setPreferences_(void *id, void* value) {
	[(WKWebViewConfiguration*)id
		setPreferences: value];
}

BOOL WKWebViewConfiguration_inst_ignoresViewportScaleLimits(void *id) {
	return [(WKWebViewConfiguration*)id
		ignoresViewportScaleLimits];
}

void WKWebViewConfiguration_inst_setIgnoresViewportScaleLimits_(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setIgnoresViewportScaleLimits: value];
}

BOOL WKWebViewConfiguration_inst_suppressesIncrementalRendering(void *id) {
	return [(WKWebViewConfiguration*)id
		suppressesIncrementalRendering];
}

void WKWebViewConfiguration_inst_setSuppressesIncrementalRendering_(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setSuppressesIncrementalRendering: value];
}

BOOL WKWebViewConfiguration_inst_allowsInlineMediaPlayback(void *id) {
	return [(WKWebViewConfiguration*)id
		allowsInlineMediaPlayback];
}

void WKWebViewConfiguration_inst_setAllowsInlineMediaPlayback_(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setAllowsInlineMediaPlayback: value];
}

BOOL WKWebViewConfiguration_inst_allowsAirPlayForMediaPlayback(void *id) {
	return [(WKWebViewConfiguration*)id
		allowsAirPlayForMediaPlayback];
}

void WKWebViewConfiguration_inst_setAllowsAirPlayForMediaPlayback_(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setAllowsAirPlayForMediaPlayback: value];
}

BOOL WKWebViewConfiguration_inst_allowsPictureInPictureMediaPlayback(void *id) {
	return [(WKWebViewConfiguration*)id
		allowsPictureInPictureMediaPlayback];
}

void WKWebViewConfiguration_inst_setAllowsPictureInPictureMediaPlayback_(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setAllowsPictureInPictureMediaPlayback: value];
}

BOOL WKWebViewConfiguration_inst_upgradeKnownHostsToHTTPS(void *id) {
	return [(WKWebViewConfiguration*)id
		upgradeKnownHostsToHTTPS];
}

void WKWebViewConfiguration_inst_setUpgradeKnownHostsToHTTPS_(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setUpgradeKnownHostsToHTTPS: value];
}

void WKPreferences_inst_setValue_forKey_(void *id, void* value, void* key) {
	[(WKPreferences*)id
		setValue: value
		forKey: key];
}

void* WKPreferences_inst_init(void *id) {
	return [(WKPreferences*)id
		init];
}

double WKPreferences_inst_minimumFontSize(void *id) {
	return [(WKPreferences*)id
		minimumFontSize];
}

void WKPreferences_inst_setMinimumFontSize_(void *id, double value) {
	[(WKPreferences*)id
		setMinimumFontSize: value];
}

BOOL WKPreferences_inst_tabFocusesLinks(void *id) {
	return [(WKPreferences*)id
		tabFocusesLinks];
}

void WKPreferences_inst_setTabFocusesLinks_(void *id, BOOL value) {
	[(WKPreferences*)id
		setTabFocusesLinks: value];
}

BOOL WKPreferences_inst_javaScriptCanOpenWindowsAutomatically(void *id) {
	return [(WKPreferences*)id
		javaScriptCanOpenWindowsAutomatically];
}

void WKPreferences_inst_setJavaScriptCanOpenWindowsAutomatically_(void *id, BOOL value) {
	[(WKPreferences*)id
		setJavaScriptCanOpenWindowsAutomatically: value];
}

BOOL WKPreferences_inst_isFraudulentWebsiteWarningEnabled(void *id) {
	return [(WKPreferences*)id
		isFraudulentWebsiteWarningEnabled];
}

void WKPreferences_inst_setFraudulentWebsiteWarningEnabled_(void *id, BOOL value) {
	[(WKPreferences*)id
		setFraudulentWebsiteWarningEnabled: value];
}

BOOL WKPreferences_inst_isTextInteractionEnabled(void *id) {
	return [(WKPreferences*)id
		isTextInteractionEnabled];
}

void WKPreferences_inst_setTextInteractionEnabled_(void *id, BOOL value) {
	[(WKPreferences*)id
		setTextInteractionEnabled: value];
}


BOOL webkit_objc_bool_true = YES;
BOOL webkit_objc_bool_false = NO;

*/
import "C"

func convertObjCBoolToGo(b C.BOOL) bool {
	// NOTE: the prefix here is used to namespace these since the linker will
	// otherwise report a "duplicate symbol" because the C functions have the
	// same name.
	return bool(C.webkit_convertObjCBool(b))
}

func convertToObjCBool(b bool) C.BOOL {
	if b {
		return C.webkit_objc_bool_true
	}
	return C.webkit_objc_bool_false
}

func WKNavigation_alloc() (
	r0 WKNavigation,
) {
	ret := C.WKNavigation_type_alloc()
	r0 = WKNavigation_fromPointer(ret)
	return
}

func WKUserScript_alloc() (
	r0 WKUserScript,
) {
	ret := C.WKUserScript_type_alloc()
	r0 = WKUserScript_fromPointer(ret)
	return
}

func WKWebView_alloc() (
	r0 WKWebView,
) {
	ret := C.WKWebView_type_alloc()
	r0 = WKWebView_fromPointer(ret)
	return
}

func WKWebView_handlesURLScheme_(
	urlScheme core.NSStringRef,
) (
	r0 bool,
) {
	ret := C.WKWebView_type_handlesURLScheme_(
		objc.RefPointer(urlScheme),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func WKWebViewConfiguration_alloc() (
	r0 WKWebViewConfiguration,
) {
	ret := C.WKWebViewConfiguration_type_alloc()
	r0 = WKWebViewConfiguration_fromPointer(ret)
	return
}

func WKPreferences_alloc() (
	r0 WKPreferences,
) {
	ret := C.WKPreferences_type_alloc()
	r0 = WKPreferences_fromPointer(ret)
	return
}

type WKNavigationRef interface {
	Pointer() uintptr
	Init_asWKNavigation() WKNavigation
}

type gen_WKNavigation struct {
	objc.Object
}

func WKNavigation_fromPointer(ptr unsafe.Pointer) WKNavigation {
	return WKNavigation{gen_WKNavigation{
		objc.Object_fromPointer(ptr),
	}}
}

func WKNavigation_fromRef(ref objc.Ref) WKNavigation {
	return WKNavigation_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_WKNavigation) Init_asWKNavigation() (
	r0 WKNavigation,
) {
	ret := C.WKNavigation_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = WKNavigation_fromPointer(ret)
	return
}

type WKUserScriptRef interface {
	Pointer() uintptr
	Init_asWKUserScript() WKUserScript
}

type gen_WKUserScript struct {
	objc.Object
}

func WKUserScript_fromPointer(ptr unsafe.Pointer) WKUserScript {
	return WKUserScript{gen_WKUserScript{
		objc.Object_fromPointer(ptr),
	}}
}

func WKUserScript_fromRef(ref objc.Ref) WKUserScript {
	return WKUserScript_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_WKUserScript) Init_asWKUserScript() (
	r0 WKUserScript,
) {
	ret := C.WKUserScript_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = WKUserScript_fromPointer(ret)
	return
}

func (x gen_WKUserScript) Source() (
	r0 core.NSString,
) {
	ret := C.WKUserScript_inst_source(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_WKUserScript) IsForMainFrameOnly() (
	r0 bool,
) {
	ret := C.WKUserScript_inst_isForMainFrameOnly(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

type WKWebViewRef interface {
	Pointer() uintptr
	Init_asWKWebView() WKWebView
}

type gen_WKWebView struct {
	cocoa.NSView
}

func WKWebView_fromPointer(ptr unsafe.Pointer) WKWebView {
	return WKWebView{gen_WKWebView{
		cocoa.NSView_fromPointer(ptr),
	}}
}

func WKWebView_fromRef(ref objc.Ref) WKWebView {
	return WKWebView_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_WKWebView) InitWithFrame_configuration__asWKWebView(
	frame core.NSRect,
	configuration WKWebViewConfigurationRef,
) (
	r0 WKWebView,
) {
	ret := C.WKWebView_inst_initWithFrame_configuration_(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frame)),
		objc.RefPointer(configuration),
	)
	r0 = WKWebView_fromPointer(ret)
	return
}

func (x gen_WKWebView) LoadRequest_(
	request core.NSURLRequestRef,
) (
	r0 WKNavigation,
) {
	ret := C.WKWebView_inst_loadRequest_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(request),
	)
	r0 = WKNavigation_fromPointer(ret)
	return
}

func (x gen_WKWebView) LoadHTMLString_baseURL_(
	string core.NSStringRef,
	baseURL core.NSURLRef,
) (
	r0 WKNavigation,
) {
	ret := C.WKWebView_inst_loadHTMLString_baseURL_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(string),
		objc.RefPointer(baseURL),
	)
	r0 = WKNavigation_fromPointer(ret)
	return
}

func (x gen_WKWebView) LoadFileURL_allowingReadAccessToURL_(
	URL core.NSURLRef,
	readAccessURL core.NSURLRef,
) (
	r0 WKNavigation,
) {
	ret := C.WKWebView_inst_loadFileURL_allowingReadAccessToURL_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(URL),
		objc.RefPointer(readAccessURL),
	)
	r0 = WKNavigation_fromPointer(ret)
	return
}

func (x gen_WKWebView) LoadData_MIMEType_characterEncodingName_baseURL_(
	data core.NSDataRef,
	MIMEType core.NSStringRef,
	characterEncodingName core.NSStringRef,
	baseURL core.NSURLRef,
) (
	r0 WKNavigation,
) {
	ret := C.WKWebView_inst_loadData_MIMEType_characterEncodingName_baseURL_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(MIMEType),
		objc.RefPointer(characterEncodingName),
		objc.RefPointer(baseURL),
	)
	r0 = WKNavigation_fromPointer(ret)
	return
}

func (x gen_WKWebView) Reload() (
	r0 WKNavigation,
) {
	ret := C.WKWebView_inst_reload(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = WKNavigation_fromPointer(ret)
	return
}

func (x gen_WKWebView) Reload_(
	sender objc.Ref,
) {
	C.WKWebView_inst_reload_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_WKWebView) ReloadFromOrigin() (
	r0 WKNavigation,
) {
	ret := C.WKWebView_inst_reloadFromOrigin(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = WKNavigation_fromPointer(ret)
	return
}

func (x gen_WKWebView) ReloadFromOrigin_(
	sender objc.Ref,
) {
	C.WKWebView_inst_reloadFromOrigin_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_WKWebView) StopLoading() {
	C.WKWebView_inst_stopLoading(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_WKWebView) StopLoading_(
	sender objc.Ref,
) {
	C.WKWebView_inst_stopLoading_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_WKWebView) GoBack_(
	sender objc.Ref,
) {
	C.WKWebView_inst_goBack_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_WKWebView) GoBack() (
	r0 WKNavigation,
) {
	ret := C.WKWebView_inst_goBack(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = WKNavigation_fromPointer(ret)
	return
}

func (x gen_WKWebView) GoForward_(
	sender objc.Ref,
) {
	C.WKWebView_inst_goForward_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)
	return
}

func (x gen_WKWebView) GoForward() (
	r0 WKNavigation,
) {
	ret := C.WKWebView_inst_goForward(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = WKNavigation_fromPointer(ret)
	return
}

func (x gen_WKWebView) LoadFileRequest_allowingReadAccessToURL_(
	request core.NSURLRequestRef,
	readAccessURL core.NSURLRef,
) (
	r0 WKNavigation,
) {
	ret := C.WKWebView_inst_loadFileRequest_allowingReadAccessToURL_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(request),
		objc.RefPointer(readAccessURL),
	)
	r0 = WKNavigation_fromPointer(ret)
	return
}

func (x gen_WKWebView) LoadSimulatedRequest_responseHTMLString_(
	request core.NSURLRequestRef,
	string core.NSStringRef,
) (
	r0 WKNavigation,
) {
	ret := C.WKWebView_inst_loadSimulatedRequest_responseHTMLString_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(request),
		objc.RefPointer(string),
	)
	r0 = WKNavigation_fromPointer(ret)
	return
}

func (x gen_WKWebView) Init_asWKWebView() (
	r0 WKWebView,
) {
	ret := C.WKWebView_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = WKWebView_fromPointer(ret)
	return
}

func (x gen_WKWebView) Configuration() (
	r0 WKWebViewConfiguration,
) {
	ret := C.WKWebView_inst_configuration(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = WKWebViewConfiguration_fromPointer(ret)
	return
}

func (x gen_WKWebView) UIDelegate() (
	r0 objc.Object,
) {
	ret := C.WKWebView_inst_UIDelegate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_WKWebView) SetUIDelegate_(
	value objc.Ref,
) {
	C.WKWebView_inst_setUIDelegate_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_WKWebView) NavigationDelegate() (
	r0 objc.Object,
) {
	ret := C.WKWebView_inst_navigationDelegate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_WKWebView) SetNavigationDelegate_(
	value objc.Ref,
) {
	C.WKWebView_inst_setNavigationDelegate_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_WKWebView) IsLoading() (
	r0 bool,
) {
	ret := C.WKWebView_inst_isLoading(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_WKWebView) Title() (
	r0 core.NSString,
) {
	ret := C.WKWebView_inst_title(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_WKWebView) URL() (
	r0 core.NSURL,
) {
	ret := C.WKWebView_inst_URL(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSURL_fromPointer(ret)
	return
}

func (x gen_WKWebView) MediaType() (
	r0 core.NSString,
) {
	ret := C.WKWebView_inst_mediaType(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_WKWebView) SetMediaType_(
	value core.NSStringRef,
) {
	C.WKWebView_inst_setMediaType_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_WKWebView) CustomUserAgent() (
	r0 core.NSString,
) {
	ret := C.WKWebView_inst_customUserAgent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_WKWebView) SetCustomUserAgent_(
	value core.NSStringRef,
) {
	C.WKWebView_inst_setCustomUserAgent_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_WKWebView) HasOnlySecureContent() (
	r0 bool,
) {
	ret := C.WKWebView_inst_hasOnlySecureContent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_WKWebView) AllowsMagnification() (
	r0 bool,
) {
	ret := C.WKWebView_inst_allowsMagnification(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_WKWebView) SetAllowsMagnification_(
	value bool,
) {
	C.WKWebView_inst_setAllowsMagnification_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_WKWebView) Magnification() (
	r0 core.CGFloat,
) {
	ret := C.WKWebView_inst_magnification(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_WKWebView) SetMagnification_(
	value core.CGFloat,
) {
	C.WKWebView_inst_setMagnification_(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)
	return
}

func (x gen_WKWebView) AllowsBackForwardNavigationGestures() (
	r0 bool,
) {
	ret := C.WKWebView_inst_allowsBackForwardNavigationGestures(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_WKWebView) SetAllowsBackForwardNavigationGestures_(
	value bool,
) {
	C.WKWebView_inst_setAllowsBackForwardNavigationGestures_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_WKWebView) CanGoBack() (
	r0 bool,
) {
	ret := C.WKWebView_inst_canGoBack(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_WKWebView) CanGoForward() (
	r0 bool,
) {
	ret := C.WKWebView_inst_canGoForward(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_WKWebView) AllowsLinkPreview() (
	r0 bool,
) {
	ret := C.WKWebView_inst_allowsLinkPreview(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_WKWebView) SetAllowsLinkPreview_(
	value bool,
) {
	C.WKWebView_inst_setAllowsLinkPreview_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_WKWebView) InteractionState() (
	r0 objc.Object,
) {
	ret := C.WKWebView_inst_interactionState(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_WKWebView) SetInteractionState_(
	value objc.Ref,
) {
	C.WKWebView_inst_setInteractionState_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

type WKWebViewConfigurationRef interface {
	Pointer() uintptr
	Init_asWKWebViewConfiguration() WKWebViewConfiguration
}

type gen_WKWebViewConfiguration struct {
	objc.Object
}

func WKWebViewConfiguration_fromPointer(ptr unsafe.Pointer) WKWebViewConfiguration {
	return WKWebViewConfiguration{gen_WKWebViewConfiguration{
		objc.Object_fromPointer(ptr),
	}}
}

func WKWebViewConfiguration_fromRef(ref objc.Ref) WKWebViewConfiguration {
	return WKWebViewConfiguration_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_WKWebViewConfiguration) SetURLSchemeHandler_forURLScheme_(
	urlSchemeHandler objc.Ref,
	urlScheme core.NSStringRef,
) {
	C.WKWebViewConfiguration_inst_setURLSchemeHandler_forURLScheme_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(urlSchemeHandler),
		objc.RefPointer(urlScheme),
	)
	return
}

func (x gen_WKWebViewConfiguration) UrlSchemeHandlerForURLScheme_(
	urlScheme core.NSStringRef,
) (
	r0 objc.Object,
) {
	ret := C.WKWebViewConfiguration_inst_urlSchemeHandlerForURLScheme_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(urlScheme),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_WKWebViewConfiguration) Init_asWKWebViewConfiguration() (
	r0 WKWebViewConfiguration,
) {
	ret := C.WKWebViewConfiguration_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = WKWebViewConfiguration_fromPointer(ret)
	return
}

func (x gen_WKWebViewConfiguration) ApplicationNameForUserAgent() (
	r0 core.NSString,
) {
	ret := C.WKWebViewConfiguration_inst_applicationNameForUserAgent(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_WKWebViewConfiguration) SetApplicationNameForUserAgent_(
	value core.NSStringRef,
) {
	C.WKWebViewConfiguration_inst_setApplicationNameForUserAgent_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_WKWebViewConfiguration) LimitsNavigationsToAppBoundDomains() (
	r0 bool,
) {
	ret := C.WKWebViewConfiguration_inst_limitsNavigationsToAppBoundDomains(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_WKWebViewConfiguration) SetLimitsNavigationsToAppBoundDomains_(
	value bool,
) {
	C.WKWebViewConfiguration_inst_setLimitsNavigationsToAppBoundDomains_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_WKWebViewConfiguration) Preferences() (
	r0 WKPreferences,
) {
	ret := C.WKWebViewConfiguration_inst_preferences(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = WKPreferences_fromPointer(ret)
	return
}

func (x gen_WKWebViewConfiguration) SetPreferences_(
	value WKPreferencesRef,
) {
	C.WKWebViewConfiguration_inst_setPreferences_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_WKWebViewConfiguration) IgnoresViewportScaleLimits() (
	r0 bool,
) {
	ret := C.WKWebViewConfiguration_inst_ignoresViewportScaleLimits(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_WKWebViewConfiguration) SetIgnoresViewportScaleLimits_(
	value bool,
) {
	C.WKWebViewConfiguration_inst_setIgnoresViewportScaleLimits_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_WKWebViewConfiguration) SuppressesIncrementalRendering() (
	r0 bool,
) {
	ret := C.WKWebViewConfiguration_inst_suppressesIncrementalRendering(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_WKWebViewConfiguration) SetSuppressesIncrementalRendering_(
	value bool,
) {
	C.WKWebViewConfiguration_inst_setSuppressesIncrementalRendering_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_WKWebViewConfiguration) AllowsInlineMediaPlayback() (
	r0 bool,
) {
	ret := C.WKWebViewConfiguration_inst_allowsInlineMediaPlayback(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_WKWebViewConfiguration) SetAllowsInlineMediaPlayback_(
	value bool,
) {
	C.WKWebViewConfiguration_inst_setAllowsInlineMediaPlayback_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_WKWebViewConfiguration) AllowsAirPlayForMediaPlayback() (
	r0 bool,
) {
	ret := C.WKWebViewConfiguration_inst_allowsAirPlayForMediaPlayback(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_WKWebViewConfiguration) SetAllowsAirPlayForMediaPlayback_(
	value bool,
) {
	C.WKWebViewConfiguration_inst_setAllowsAirPlayForMediaPlayback_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_WKWebViewConfiguration) AllowsPictureInPictureMediaPlayback() (
	r0 bool,
) {
	ret := C.WKWebViewConfiguration_inst_allowsPictureInPictureMediaPlayback(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_WKWebViewConfiguration) SetAllowsPictureInPictureMediaPlayback_(
	value bool,
) {
	C.WKWebViewConfiguration_inst_setAllowsPictureInPictureMediaPlayback_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_WKWebViewConfiguration) UpgradeKnownHostsToHTTPS() (
	r0 bool,
) {
	ret := C.WKWebViewConfiguration_inst_upgradeKnownHostsToHTTPS(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_WKWebViewConfiguration) SetUpgradeKnownHostsToHTTPS_(
	value bool,
) {
	C.WKWebViewConfiguration_inst_setUpgradeKnownHostsToHTTPS_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

type WKPreferencesRef interface {
	Pointer() uintptr
	Init_asWKPreferences() WKPreferences
}

type gen_WKPreferences struct {
	objc.Object
}

func WKPreferences_fromPointer(ptr unsafe.Pointer) WKPreferences {
	return WKPreferences{gen_WKPreferences{
		objc.Object_fromPointer(ptr),
	}}
}

func WKPreferences_fromRef(ref objc.Ref) WKPreferences {
	return WKPreferences_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_WKPreferences) SetValue_forKey_(
	value objc.Ref,
	key core.NSStringRef,
) {
	C.WKPreferences_inst_setValue_forKey_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
		objc.RefPointer(key),
	)
	return
}

func (x gen_WKPreferences) Init_asWKPreferences() (
	r0 WKPreferences,
) {
	ret := C.WKPreferences_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = WKPreferences_fromPointer(ret)
	return
}

func (x gen_WKPreferences) MinimumFontSize() (
	r0 core.CGFloat,
) {
	ret := C.WKPreferences_inst_minimumFontSize(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.CGFloat(ret)
	return
}

func (x gen_WKPreferences) SetMinimumFontSize_(
	value core.CGFloat,
) {
	C.WKPreferences_inst_setMinimumFontSize_(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)
	return
}

func (x gen_WKPreferences) TabFocusesLinks() (
	r0 bool,
) {
	ret := C.WKPreferences_inst_tabFocusesLinks(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_WKPreferences) SetTabFocusesLinks_(
	value bool,
) {
	C.WKPreferences_inst_setTabFocusesLinks_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_WKPreferences) JavaScriptCanOpenWindowsAutomatically() (
	r0 bool,
) {
	ret := C.WKPreferences_inst_javaScriptCanOpenWindowsAutomatically(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_WKPreferences) SetJavaScriptCanOpenWindowsAutomatically_(
	value bool,
) {
	C.WKPreferences_inst_setJavaScriptCanOpenWindowsAutomatically_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_WKPreferences) IsFraudulentWebsiteWarningEnabled() (
	r0 bool,
) {
	ret := C.WKPreferences_inst_isFraudulentWebsiteWarningEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_WKPreferences) SetFraudulentWebsiteWarningEnabled_(
	value bool,
) {
	C.WKPreferences_inst_setFraudulentWebsiteWarningEnabled_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}

func (x gen_WKPreferences) IsTextInteractionEnabled() (
	r0 bool,
) {
	ret := C.WKPreferences_inst_isTextInteractionEnabled(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_WKPreferences) SetTextInteractionEnabled_(
	value bool,
) {
	C.WKPreferences_inst_setTextInteractionEnabled_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)
	return
}
