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


void* WKNavigation_type_Alloc() {
	return [WKNavigation
		alloc];
}
void* WKUserScript_type_Alloc() {
	return [WKUserScript
		alloc];
}
void* WKWebView_type_Alloc() {
	return [WKWebView
		alloc];
}
BOOL WKWebView_type_HandlesURLScheme(void* urlScheme) {
	return [WKWebView
		handlesURLScheme: urlScheme];
}
void* WKWebViewConfiguration_type_Alloc() {
	return [WKWebViewConfiguration
		alloc];
}
void* WKPreferences_type_Alloc() {
	return [WKPreferences
		alloc];
}


void* WKNavigation_inst_Init(void *id) {
	return [(WKNavigation*)id
		init];
}

void* WKUserScript_inst_Init(void *id) {
	return [(WKUserScript*)id
		init];
}

void* WKUserScript_inst_Source(void *id) {
	return [(WKUserScript*)id
		source];
}

BOOL WKUserScript_inst_IsForMainFrameOnly(void *id) {
	return [(WKUserScript*)id
		isForMainFrameOnly];
}

void* WKWebView_inst_GoBack(void *id) {
	return [(WKWebView*)id
		goBack];
}

void WKWebView_inst_GoBack_(void *id, void* sender) {
	[(WKWebView*)id
		goBack: sender];
}

void* WKWebView_inst_GoForward(void *id) {
	return [(WKWebView*)id
		goForward];
}

void WKWebView_inst_GoForward_(void *id, void* sender) {
	[(WKWebView*)id
		goForward: sender];
}

void* WKWebView_inst_InitWithFrameConfiguration(void *id, NSRect frame, void* configuration) {
	return [(WKWebView*)id
		initWithFrame: frame
		configuration: configuration];
}

void* WKWebView_inst_LoadDataMIMETypeCharacterEncodingNameBaseURL(void *id, void* data, void* MIMEType, void* characterEncodingName, void* baseURL) {
	return [(WKWebView*)id
		loadData: data
		MIMEType: MIMEType
		characterEncodingName: characterEncodingName
		baseURL: baseURL];
}

void* WKWebView_inst_LoadFileRequestAllowingReadAccessToURL(void *id, void* request, void* readAccessURL) {
	return [(WKWebView*)id
		loadFileRequest: request
		allowingReadAccessToURL: readAccessURL];
}

void* WKWebView_inst_LoadFileURLAllowingReadAccessToURL(void *id, void* URL, void* readAccessURL) {
	return [(WKWebView*)id
		loadFileURL: URL
		allowingReadAccessToURL: readAccessURL];
}

void* WKWebView_inst_LoadHTMLStringBaseURL(void *id, void* string, void* baseURL) {
	return [(WKWebView*)id
		loadHTMLString: string
		baseURL: baseURL];
}

void* WKWebView_inst_LoadRequest(void *id, void* request) {
	return [(WKWebView*)id
		loadRequest: request];
}

void* WKWebView_inst_LoadSimulatedRequestResponseHTMLString(void *id, void* request, void* string) {
	return [(WKWebView*)id
		loadSimulatedRequest: request
		responseHTMLString: string];
}

void* WKWebView_inst_Reload(void *id) {
	return [(WKWebView*)id
		reload];
}

void WKWebView_inst_Reload_(void *id, void* sender) {
	[(WKWebView*)id
		reload: sender];
}

void* WKWebView_inst_ReloadFromOrigin(void *id) {
	return [(WKWebView*)id
		reloadFromOrigin];
}

void WKWebView_inst_ReloadFromOrigin_(void *id, void* sender) {
	[(WKWebView*)id
		reloadFromOrigin: sender];
}

void WKWebView_inst_StopLoading(void *id) {
	[(WKWebView*)id
		stopLoading];
}

void WKWebView_inst_StopLoading_(void *id, void* sender) {
	[(WKWebView*)id
		stopLoading: sender];
}

void* WKWebView_inst_Init(void *id) {
	return [(WKWebView*)id
		init];
}

void* WKWebView_inst_Configuration(void *id) {
	return [(WKWebView*)id
		configuration];
}

void* WKWebView_inst_UIDelegate(void *id) {
	return [(WKWebView*)id
		UIDelegate];
}

void WKWebView_inst_SetUIDelegate(void *id, void* value) {
	[(WKWebView*)id
		setUIDelegate: value];
}

void* WKWebView_inst_NavigationDelegate(void *id) {
	return [(WKWebView*)id
		navigationDelegate];
}

void WKWebView_inst_SetNavigationDelegate(void *id, void* value) {
	[(WKWebView*)id
		setNavigationDelegate: value];
}

BOOL WKWebView_inst_IsLoading(void *id) {
	return [(WKWebView*)id
		isLoading];
}

void* WKWebView_inst_Title(void *id) {
	return [(WKWebView*)id
		title];
}

void* WKWebView_inst_URL(void *id) {
	return [(WKWebView*)id
		URL];
}

void* WKWebView_inst_MediaType(void *id) {
	return [(WKWebView*)id
		mediaType];
}

void WKWebView_inst_SetMediaType(void *id, void* value) {
	[(WKWebView*)id
		setMediaType: value];
}

void* WKWebView_inst_CustomUserAgent(void *id) {
	return [(WKWebView*)id
		customUserAgent];
}

void WKWebView_inst_SetCustomUserAgent(void *id, void* value) {
	[(WKWebView*)id
		setCustomUserAgent: value];
}

BOOL WKWebView_inst_HasOnlySecureContent(void *id) {
	return [(WKWebView*)id
		hasOnlySecureContent];
}

BOOL WKWebView_inst_AllowsMagnification(void *id) {
	return [(WKWebView*)id
		allowsMagnification];
}

void WKWebView_inst_SetAllowsMagnification(void *id, BOOL value) {
	[(WKWebView*)id
		setAllowsMagnification: value];
}

double WKWebView_inst_Magnification(void *id) {
	return [(WKWebView*)id
		magnification];
}

void WKWebView_inst_SetMagnification(void *id, double value) {
	[(WKWebView*)id
		setMagnification: value];
}

BOOL WKWebView_inst_AllowsBackForwardNavigationGestures(void *id) {
	return [(WKWebView*)id
		allowsBackForwardNavigationGestures];
}

void WKWebView_inst_SetAllowsBackForwardNavigationGestures(void *id, BOOL value) {
	[(WKWebView*)id
		setAllowsBackForwardNavigationGestures: value];
}

BOOL WKWebView_inst_CanGoBack(void *id) {
	return [(WKWebView*)id
		canGoBack];
}

BOOL WKWebView_inst_CanGoForward(void *id) {
	return [(WKWebView*)id
		canGoForward];
}

BOOL WKWebView_inst_AllowsLinkPreview(void *id) {
	return [(WKWebView*)id
		allowsLinkPreview];
}

void WKWebView_inst_SetAllowsLinkPreview(void *id, BOOL value) {
	[(WKWebView*)id
		setAllowsLinkPreview: value];
}

void* WKWebView_inst_InteractionState(void *id) {
	return [(WKWebView*)id
		interactionState];
}

void WKWebView_inst_SetInteractionState(void *id, void* value) {
	[(WKWebView*)id
		setInteractionState: value];
}

void WKWebViewConfiguration_inst_SetURLSchemeHandlerForURLScheme(void *id, void* urlSchemeHandler, void* urlScheme) {
	[(WKWebViewConfiguration*)id
		setURLSchemeHandler: urlSchemeHandler
		forURLScheme: urlScheme];
}

void* WKWebViewConfiguration_inst_UrlSchemeHandlerForURLScheme(void *id, void* urlScheme) {
	return [(WKWebViewConfiguration*)id
		urlSchemeHandlerForURLScheme: urlScheme];
}

void* WKWebViewConfiguration_inst_Init(void *id) {
	return [(WKWebViewConfiguration*)id
		init];
}

void* WKWebViewConfiguration_inst_ApplicationNameForUserAgent(void *id) {
	return [(WKWebViewConfiguration*)id
		applicationNameForUserAgent];
}

void WKWebViewConfiguration_inst_SetApplicationNameForUserAgent(void *id, void* value) {
	[(WKWebViewConfiguration*)id
		setApplicationNameForUserAgent: value];
}

BOOL WKWebViewConfiguration_inst_LimitsNavigationsToAppBoundDomains(void *id) {
	return [(WKWebViewConfiguration*)id
		limitsNavigationsToAppBoundDomains];
}

void WKWebViewConfiguration_inst_SetLimitsNavigationsToAppBoundDomains(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setLimitsNavigationsToAppBoundDomains: value];
}

void* WKWebViewConfiguration_inst_Preferences(void *id) {
	return [(WKWebViewConfiguration*)id
		preferences];
}

void WKWebViewConfiguration_inst_SetPreferences(void *id, void* value) {
	[(WKWebViewConfiguration*)id
		setPreferences: value];
}

BOOL WKWebViewConfiguration_inst_IgnoresViewportScaleLimits(void *id) {
	return [(WKWebViewConfiguration*)id
		ignoresViewportScaleLimits];
}

void WKWebViewConfiguration_inst_SetIgnoresViewportScaleLimits(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setIgnoresViewportScaleLimits: value];
}

BOOL WKWebViewConfiguration_inst_SuppressesIncrementalRendering(void *id) {
	return [(WKWebViewConfiguration*)id
		suppressesIncrementalRendering];
}

void WKWebViewConfiguration_inst_SetSuppressesIncrementalRendering(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setSuppressesIncrementalRendering: value];
}

BOOL WKWebViewConfiguration_inst_AllowsInlineMediaPlayback(void *id) {
	return [(WKWebViewConfiguration*)id
		allowsInlineMediaPlayback];
}

void WKWebViewConfiguration_inst_SetAllowsInlineMediaPlayback(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setAllowsInlineMediaPlayback: value];
}

BOOL WKWebViewConfiguration_inst_AllowsAirPlayForMediaPlayback(void *id) {
	return [(WKWebViewConfiguration*)id
		allowsAirPlayForMediaPlayback];
}

void WKWebViewConfiguration_inst_SetAllowsAirPlayForMediaPlayback(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setAllowsAirPlayForMediaPlayback: value];
}

BOOL WKWebViewConfiguration_inst_AllowsPictureInPictureMediaPlayback(void *id) {
	return [(WKWebViewConfiguration*)id
		allowsPictureInPictureMediaPlayback];
}

void WKWebViewConfiguration_inst_SetAllowsPictureInPictureMediaPlayback(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setAllowsPictureInPictureMediaPlayback: value];
}

BOOL WKWebViewConfiguration_inst_UpgradeKnownHostsToHTTPS(void *id) {
	return [(WKWebViewConfiguration*)id
		upgradeKnownHostsToHTTPS];
}

void WKWebViewConfiguration_inst_SetUpgradeKnownHostsToHTTPS(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setUpgradeKnownHostsToHTTPS: value];
}

void WKPreferences_inst_SetValueForKey(void *id, void* value, void* key) {
	[(WKPreferences*)id
		setValue: value
		forKey: key];
}

void* WKPreferences_inst_Init(void *id) {
	return [(WKPreferences*)id
		init];
}

double WKPreferences_inst_MinimumFontSize(void *id) {
	return [(WKPreferences*)id
		minimumFontSize];
}

void WKPreferences_inst_SetMinimumFontSize(void *id, double value) {
	[(WKPreferences*)id
		setMinimumFontSize: value];
}

BOOL WKPreferences_inst_TabFocusesLinks(void *id) {
	return [(WKPreferences*)id
		tabFocusesLinks];
}

void WKPreferences_inst_SetTabFocusesLinks(void *id, BOOL value) {
	[(WKPreferences*)id
		setTabFocusesLinks: value];
}

BOOL WKPreferences_inst_JavaScriptCanOpenWindowsAutomatically(void *id) {
	return [(WKPreferences*)id
		javaScriptCanOpenWindowsAutomatically];
}

void WKPreferences_inst_SetJavaScriptCanOpenWindowsAutomatically(void *id, BOOL value) {
	[(WKPreferences*)id
		setJavaScriptCanOpenWindowsAutomatically: value];
}

BOOL WKPreferences_inst_IsFraudulentWebsiteWarningEnabled(void *id) {
	return [(WKPreferences*)id
		isFraudulentWebsiteWarningEnabled];
}

void WKPreferences_inst_SetFraudulentWebsiteWarningEnabled(void *id, BOOL value) {
	[(WKPreferences*)id
		setFraudulentWebsiteWarningEnabled: value];
}

BOOL WKPreferences_inst_IsTextInteractionEnabled(void *id) {
	return [(WKPreferences*)id
		isTextInteractionEnabled];
}

void WKPreferences_inst_SetTextInteractionEnabled(void *id, BOOL value) {
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

// WKNavigation_Alloc is undocumented.
func WKNavigation_Alloc() WKNavigation {
	ret := C.WKNavigation_type_Alloc()

	return WKNavigation_FromPointer(ret)
}

// WKUserScript_Alloc is undocumented.
func WKUserScript_Alloc() WKUserScript {
	ret := C.WKUserScript_type_Alloc()

	return WKUserScript_FromPointer(ret)
}

// WKWebView_Alloc is undocumented.
func WKWebView_Alloc() WKWebView {
	ret := C.WKWebView_type_Alloc()

	return WKWebView_FromPointer(ret)
}

// WKWebView_HandlesURLScheme returns a Boolean value that indicates whether WebKit natively supports resources with the specified URL scheme.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/2875370-handlesurlscheme?language=objc for details.
func WKWebView_HandlesURLScheme(urlScheme core.NSStringRef) bool {
	ret := C.WKWebView_type_HandlesURLScheme(
		objc.RefPointer(urlScheme),
	)

	return convertObjCBoolToGo(ret)
}

// WKWebViewConfiguration_Alloc is undocumented.
func WKWebViewConfiguration_Alloc() WKWebViewConfiguration {
	ret := C.WKWebViewConfiguration_type_Alloc()

	return WKWebViewConfiguration_FromPointer(ret)
}

// WKPreferences_Alloc is undocumented.
func WKPreferences_Alloc() WKPreferences {
	ret := C.WKPreferences_type_Alloc()

	return WKPreferences_FromPointer(ret)
}

type WKNavigationRef interface {
	Pointer() uintptr
	Init_AsWKNavigation() WKNavigation
}

type gen_WKNavigation struct {
	objc.Object
}

func WKNavigation_FromPointer(ptr unsafe.Pointer) WKNavigation {
	return WKNavigation{gen_WKNavigation{
		objc.Object_FromPointer(ptr),
	}}
}

func WKNavigation_FromRef(ref objc.Ref) WKNavigation {
	return WKNavigation_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init initializes a new instance of the WKNavigation class.
func (x gen_WKNavigation) Init() WKNavigation {
	ret := C.WKNavigation_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return WKNavigation_FromPointer(ret)
}

// Init_AsWKNavigation is a typed version of Init.
func (x gen_WKNavigation) Init_AsWKNavigation() WKNavigation {
	ret := C.WKNavigation_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return WKNavigation_FromPointer(ret)
}

type WKUserScriptRef interface {
	Pointer() uintptr
	Init_AsWKUserScript() WKUserScript
}

type gen_WKUserScript struct {
	objc.Object
}

func WKUserScript_FromPointer(ptr unsafe.Pointer) WKUserScript {
	return WKUserScript{gen_WKUserScript{
		objc.Object_FromPointer(ptr),
	}}
}

func WKUserScript_FromRef(ref objc.Ref) WKUserScript {
	return WKUserScript_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init initializes a new instance of the WKUserScript class.
func (x gen_WKUserScript) Init() WKUserScript {
	ret := C.WKUserScript_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return WKUserScript_FromPointer(ret)
}

// Init_AsWKUserScript is a typed version of Init.
func (x gen_WKUserScript) Init_AsWKUserScript() WKUserScript {
	ret := C.WKUserScript_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return WKUserScript_FromPointer(ret)
}

// Source returns the script’s source code.
//
// See https://developer.apple.com/documentation/webkit/wkuserscript/1537787-source?language=objc for details.
func (x gen_WKUserScript) Source() core.NSString {
	ret := C.WKUserScript_inst_Source(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// IsForMainFrameOnly returns a Boolean value that indicates whether to inject the script into the main frame or all frames.
//
// See https://developer.apple.com/documentation/webkit/wkuserscript/1537856-formainframeonly?language=objc for details.
func (x gen_WKUserScript) IsForMainFrameOnly() bool {
	ret := C.WKUserScript_inst_IsForMainFrameOnly(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

type WKWebViewRef interface {
	Pointer() uintptr
	Init_AsWKWebView() WKWebView
}

type gen_WKWebView struct {
	cocoa.NSView
}

func WKWebView_FromPointer(ptr unsafe.Pointer) WKWebView {
	return WKWebView{gen_WKWebView{
		cocoa.NSView_FromPointer(ptr),
	}}
}

func WKWebView_FromRef(ref objc.Ref) WKWebView {
	return WKWebView_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// GoBack navigates to the back item in the back-forward list.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414952-goback?language=objc for details.
func (x gen_WKWebView) GoBack() WKNavigation {
	ret := C.WKWebView_inst_GoBack(
		unsafe.Pointer(x.Pointer()),
	)

	return WKNavigation_FromPointer(ret)
}

// GoBack_ navigates to the back item in the back-forward list.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414975-goback?language=objc for details.
func (x gen_WKWebView) GoBack_(
	sender objc.Ref,
) {
	C.WKWebView_inst_GoBack_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// GoForward navigates to the forward item in the back-forward list.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414993-goforward?language=objc for details.
func (x gen_WKWebView) GoForward() WKNavigation {
	ret := C.WKWebView_inst_GoForward(
		unsafe.Pointer(x.Pointer()),
	)

	return WKNavigation_FromPointer(ret)
}

// GoForward_ navigates to the forward item in the back-forward list.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414960-goforward?language=objc for details.
func (x gen_WKWebView) GoForward_(
	sender objc.Ref,
) {
	C.WKWebView_inst_GoForward_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// InitWithFrameConfiguration creates a web view and initializes it with the specified frame and configuration data.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414998-initwithframe?language=objc for details.
func (x gen_WKWebView) InitWithFrameConfiguration(
	frame core.NSRect,
	configuration WKWebViewConfigurationRef,
) WKWebView {
	ret := C.WKWebView_inst_InitWithFrameConfiguration(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frame)),
		objc.RefPointer(configuration),
	)

	return WKWebView_FromPointer(ret)
}

// LoadDataMIMETypeCharacterEncodingNameBaseURL loads the content of the specified data object and navigates to it.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415011-loaddata?language=objc for details.
func (x gen_WKWebView) LoadDataMIMETypeCharacterEncodingNameBaseURL(
	data core.NSDataRef,
	MIMEType core.NSStringRef,
	characterEncodingName core.NSStringRef,
	baseURL core.NSURLRef,
) WKNavigation {
	ret := C.WKWebView_inst_LoadDataMIMETypeCharacterEncodingNameBaseURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(MIMEType),
		objc.RefPointer(characterEncodingName),
		objc.RefPointer(baseURL),
	)

	return WKNavigation_FromPointer(ret)
}

// LoadFileRequestAllowingReadAccessToURL is undocumented.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/3752237-loadfilerequest?language=objc for details.
func (x gen_WKWebView) LoadFileRequestAllowingReadAccessToURL(
	request core.NSURLRequestRef,
	readAccessURL core.NSURLRef,
) WKNavigation {
	ret := C.WKWebView_inst_LoadFileRequestAllowingReadAccessToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(request),
		objc.RefPointer(readAccessURL),
	)

	return WKNavigation_FromPointer(ret)
}

// LoadFileURLAllowingReadAccessToURL loads the web content from the specified file and navigates to it.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414973-loadfileurl?language=objc for details.
func (x gen_WKWebView) LoadFileURLAllowingReadAccessToURL(
	URL core.NSURLRef,
	readAccessURL core.NSURLRef,
) WKNavigation {
	ret := C.WKWebView_inst_LoadFileURLAllowingReadAccessToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(URL),
		objc.RefPointer(readAccessURL),
	)

	return WKNavigation_FromPointer(ret)
}

// LoadHTMLStringBaseURL loads the contents of the specified HTML string and navigates to it.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415004-loadhtmlstring?language=objc for details.
func (x gen_WKWebView) LoadHTMLStringBaseURL(
	string core.NSStringRef,
	baseURL core.NSURLRef,
) WKNavigation {
	ret := C.WKWebView_inst_LoadHTMLStringBaseURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(string),
		objc.RefPointer(baseURL),
	)

	return WKNavigation_FromPointer(ret)
}

// LoadRequest loads the web content referenced by the specified URL request object and navigates to it.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414954-loadrequest?language=objc for details.
func (x gen_WKWebView) LoadRequest(
	request core.NSURLRequestRef,
) WKNavigation {
	ret := C.WKWebView_inst_LoadRequest(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(request),
	)

	return WKNavigation_FromPointer(ret)
}

// LoadSimulatedRequestResponseHTMLString is undocumented.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/3763095-loadsimulatedrequest?language=objc for details.
func (x gen_WKWebView) LoadSimulatedRequestResponseHTMLString(
	request core.NSURLRequestRef,
	string core.NSStringRef,
) WKNavigation {
	ret := C.WKWebView_inst_LoadSimulatedRequestResponseHTMLString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(request),
		objc.RefPointer(string),
	)

	return WKNavigation_FromPointer(ret)
}

// Reload reloads the current webpage.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414969-reload?language=objc for details.
func (x gen_WKWebView) Reload() WKNavigation {
	ret := C.WKWebView_inst_Reload(
		unsafe.Pointer(x.Pointer()),
	)

	return WKNavigation_FromPointer(ret)
}

// Reload_ reloads the current webpage.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414987-reload?language=objc for details.
func (x gen_WKWebView) Reload_(
	sender objc.Ref,
) {
	C.WKWebView_inst_Reload_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// ReloadFromOrigin reloads the current webpage, and performs end-to-end revalidation of the content using cache-validating conditionals, if possible.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414956-reloadfromorigin?language=objc for details.
func (x gen_WKWebView) ReloadFromOrigin() WKNavigation {
	ret := C.WKWebView_inst_ReloadFromOrigin(
		unsafe.Pointer(x.Pointer()),
	)

	return WKNavigation_FromPointer(ret)
}

// ReloadFromOrigin_ reloads the current webpage, and performs end-to-end revalidation of the content using cache-validating conditionals, if possible.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414989-reloadfromorigin?language=objc for details.
func (x gen_WKWebView) ReloadFromOrigin_(
	sender objc.Ref,
) {
	C.WKWebView_inst_ReloadFromOrigin_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// StopLoading stops loading all resources on the current page.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414981-stoploading?language=objc for details.
func (x gen_WKWebView) StopLoading() {
	C.WKWebView_inst_StopLoading(
		unsafe.Pointer(x.Pointer()),
	)

	return
}

// StopLoading_ stops loading all resources on the current page.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415013-stoploading?language=objc for details.
func (x gen_WKWebView) StopLoading_(
	sender objc.Ref,
) {
	C.WKWebView_inst_StopLoading_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return
}

// Init initializes a new instance of the WKWebView class.
func (x gen_WKWebView) Init() WKWebView {
	ret := C.WKWebView_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return WKWebView_FromPointer(ret)
}

// Init_AsWKWebView is a typed version of Init.
func (x gen_WKWebView) Init_AsWKWebView() WKWebView {
	ret := C.WKWebView_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return WKWebView_FromPointer(ret)
}

// Configuration returns the object that contains the configuration details for the web view.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414979-configuration?language=objc for details.
func (x gen_WKWebView) Configuration() WKWebViewConfiguration {
	ret := C.WKWebView_inst_Configuration(
		unsafe.Pointer(x.Pointer()),
	)

	return WKWebViewConfiguration_FromPointer(ret)
}

// UIDelegate returns the object you use to integrate custom user interface elements, such as contextual menus or panels, into web view interactions.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415009-uidelegate?language=objc for details.
func (x gen_WKWebView) UIDelegate() objc.Object {
	ret := C.WKWebView_inst_UIDelegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetUIDelegate returns the object you use to integrate custom user interface elements, such as contextual menus or panels, into web view interactions.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415009-uidelegate?language=objc for details.
func (x gen_WKWebView) SetUIDelegate(
	value objc.Ref,
) {
	C.WKWebView_inst_SetUIDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// NavigationDelegate returns the object you use to manage navigation behavior for the web view.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414971-navigationdelegate?language=objc for details.
func (x gen_WKWebView) NavigationDelegate() objc.Object {
	ret := C.WKWebView_inst_NavigationDelegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetNavigationDelegate returns the object you use to manage navigation behavior for the web view.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414971-navigationdelegate?language=objc for details.
func (x gen_WKWebView) SetNavigationDelegate(
	value objc.Ref,
) {
	C.WKWebView_inst_SetNavigationDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// IsLoading returns a Boolean value that indicates whether the view is currently loading content.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414964-loading?language=objc for details.
func (x gen_WKWebView) IsLoading() bool {
	ret := C.WKWebView_inst_IsLoading(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// Title returns the page title.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415015-title?language=objc for details.
func (x gen_WKWebView) Title() core.NSString {
	ret := C.WKWebView_inst_Title(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// URL returns the URL for the current webpage.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415005-url?language=objc for details.
func (x gen_WKWebView) URL() core.NSURL {
	ret := C.WKWebView_inst_URL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_FromPointer(ret)
}

// MediaType returns the media type for the contents of the web view.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/3516410-mediatype?language=objc for details.
func (x gen_WKWebView) MediaType() core.NSString {
	ret := C.WKWebView_inst_MediaType(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetMediaType returns the media type for the contents of the web view.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/3516410-mediatype?language=objc for details.
func (x gen_WKWebView) SetMediaType(
	value core.NSStringRef,
) {
	C.WKWebView_inst_SetMediaType(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// CustomUserAgent returns the custom user agent string.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414950-customuseragent?language=objc for details.
func (x gen_WKWebView) CustomUserAgent() core.NSString {
	ret := C.WKWebView_inst_CustomUserAgent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetCustomUserAgent returns the custom user agent string.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414950-customuseragent?language=objc for details.
func (x gen_WKWebView) SetCustomUserAgent(
	value core.NSStringRef,
) {
	C.WKWebView_inst_SetCustomUserAgent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// HasOnlySecureContent returns a Boolean value that indicates whether the web view loaded all resources on the page through securely encrypted connections.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415002-hasonlysecurecontent?language=objc for details.
func (x gen_WKWebView) HasOnlySecureContent() bool {
	ret := C.WKWebView_inst_HasOnlySecureContent(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// AllowsMagnification returns a Boolean value that indicates whether magnify gestures change the web view’s magnification.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414983-allowsmagnification?language=objc for details.
func (x gen_WKWebView) AllowsMagnification() bool {
	ret := C.WKWebView_inst_AllowsMagnification(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsMagnification returns a Boolean value that indicates whether magnify gestures change the web view’s magnification.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414983-allowsmagnification?language=objc for details.
func (x gen_WKWebView) SetAllowsMagnification(
	value bool,
) {
	C.WKWebView_inst_SetAllowsMagnification(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// Magnification returns the factor by which the page content is currently scaled.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414985-magnification?language=objc for details.
func (x gen_WKWebView) Magnification() core.CGFloat {
	ret := C.WKWebView_inst_Magnification(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// SetMagnification returns the factor by which the page content is currently scaled.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414985-magnification?language=objc for details.
func (x gen_WKWebView) SetMagnification(
	value core.CGFloat,
) {
	C.WKWebView_inst_SetMagnification(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return
}

// AllowsBackForwardNavigationGestures returns a Boolean value that indicates whether horizontal swipe gestures trigger backward and forward page navigation.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414995-allowsbackforwardnavigationgestu?language=objc for details.
func (x gen_WKWebView) AllowsBackForwardNavigationGestures() bool {
	ret := C.WKWebView_inst_AllowsBackForwardNavigationGestures(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsBackForwardNavigationGestures returns a Boolean value that indicates whether horizontal swipe gestures trigger backward and forward page navigation.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414995-allowsbackforwardnavigationgestu?language=objc for details.
func (x gen_WKWebView) SetAllowsBackForwardNavigationGestures(
	value bool,
) {
	C.WKWebView_inst_SetAllowsBackForwardNavigationGestures(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// CanGoBack returns a Boolean value that indicates whether there is a valid back item in the back-forward list.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414966-cangoback?language=objc for details.
func (x gen_WKWebView) CanGoBack() bool {
	ret := C.WKWebView_inst_CanGoBack(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// CanGoForward returns a Boolean value that indicates whether there is a valid forward item in the back-forward list.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414962-cangoforward?language=objc for details.
func (x gen_WKWebView) CanGoForward() bool {
	ret := C.WKWebView_inst_CanGoForward(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// AllowsLinkPreview returns a Boolean value that determines whether pressing a link displays a preview of the destination for the link.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415000-allowslinkpreview?language=objc for details.
func (x gen_WKWebView) AllowsLinkPreview() bool {
	ret := C.WKWebView_inst_AllowsLinkPreview(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsLinkPreview returns a Boolean value that determines whether pressing a link displays a preview of the destination for the link.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415000-allowslinkpreview?language=objc for details.
func (x gen_WKWebView) SetAllowsLinkPreview(
	value bool,
) {
	C.WKWebView_inst_SetAllowsLinkPreview(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// InteractionState is undocumented.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/3752236-interactionstate?language=objc for details.
func (x gen_WKWebView) InteractionState() objc.Object {
	ret := C.WKWebView_inst_InteractionState(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_FromPointer(ret)
}

// SetInteractionState is undocumented.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/3752236-interactionstate?language=objc for details.
func (x gen_WKWebView) SetInteractionState(
	value objc.Ref,
) {
	C.WKWebView_inst_SetInteractionState(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

type WKWebViewConfigurationRef interface {
	Pointer() uintptr
	Init_AsWKWebViewConfiguration() WKWebViewConfiguration
}

type gen_WKWebViewConfiguration struct {
	objc.Object
}

func WKWebViewConfiguration_FromPointer(ptr unsafe.Pointer) WKWebViewConfiguration {
	return WKWebViewConfiguration{gen_WKWebViewConfiguration{
		objc.Object_FromPointer(ptr),
	}}
}

func WKWebViewConfiguration_FromRef(ref objc.Ref) WKWebViewConfiguration {
	return WKWebViewConfiguration_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// SetURLSchemeHandlerForURLScheme registers an object to load resources associated with the specified URL scheme.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/2875766-seturlschemehandler?language=objc for details.
func (x gen_WKWebViewConfiguration) SetURLSchemeHandlerForURLScheme(
	urlSchemeHandler objc.Ref,
	urlScheme core.NSStringRef,
) {
	C.WKWebViewConfiguration_inst_SetURLSchemeHandlerForURLScheme(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(urlSchemeHandler),
		objc.RefPointer(urlScheme),
	)

	return
}

// UrlSchemeHandlerForURLScheme returns the currently registered handler object for the specified URL scheme.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/2875767-urlschemehandlerforurlscheme?language=objc for details.
func (x gen_WKWebViewConfiguration) UrlSchemeHandlerForURLScheme(
	urlScheme core.NSStringRef,
) objc.Object {
	ret := C.WKWebViewConfiguration_inst_UrlSchemeHandlerForURLScheme(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(urlScheme),
	)

	return objc.Object_FromPointer(ret)
}

// Init initializes a new instance of the WKWebViewConfiguration class.
func (x gen_WKWebViewConfiguration) Init() WKWebViewConfiguration {
	ret := C.WKWebViewConfiguration_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return WKWebViewConfiguration_FromPointer(ret)
}

// Init_AsWKWebViewConfiguration is a typed version of Init.
func (x gen_WKWebViewConfiguration) Init_AsWKWebViewConfiguration() WKWebViewConfiguration {
	ret := C.WKWebViewConfiguration_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return WKWebViewConfiguration_FromPointer(ret)
}

// ApplicationNameForUserAgent returns the app name that appears in the user agent string.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395665-applicationnameforuseragent?language=objc for details.
func (x gen_WKWebViewConfiguration) ApplicationNameForUserAgent() core.NSString {
	ret := C.WKWebViewConfiguration_inst_ApplicationNameForUserAgent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_FromPointer(ret)
}

// SetApplicationNameForUserAgent returns the app name that appears in the user agent string.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395665-applicationnameforuseragent?language=objc for details.
func (x gen_WKWebViewConfiguration) SetApplicationNameForUserAgent(
	value core.NSStringRef,
) {
	C.WKWebViewConfiguration_inst_SetApplicationNameForUserAgent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// LimitsNavigationsToAppBoundDomains returns a Boolean value that indicates whether the web view limits navigation to pages within the app’s domain.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/3585117-limitsnavigationstoappbounddomai?language=objc for details.
func (x gen_WKWebViewConfiguration) LimitsNavigationsToAppBoundDomains() bool {
	ret := C.WKWebViewConfiguration_inst_LimitsNavigationsToAppBoundDomains(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetLimitsNavigationsToAppBoundDomains returns a Boolean value that indicates whether the web view limits navigation to pages within the app’s domain.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/3585117-limitsnavigationstoappbounddomai?language=objc for details.
func (x gen_WKWebViewConfiguration) SetLimitsNavigationsToAppBoundDomains(
	value bool,
) {
	C.WKWebViewConfiguration_inst_SetLimitsNavigationsToAppBoundDomains(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// Preferences returns the object that manages the preference-related settings for the web view.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395666-preferences?language=objc for details.
func (x gen_WKWebViewConfiguration) Preferences() WKPreferences {
	ret := C.WKWebViewConfiguration_inst_Preferences(
		unsafe.Pointer(x.Pointer()),
	)

	return WKPreferences_FromPointer(ret)
}

// SetPreferences returns the object that manages the preference-related settings for the web view.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395666-preferences?language=objc for details.
func (x gen_WKWebViewConfiguration) SetPreferences(
	value WKPreferencesRef,
) {
	C.WKWebViewConfiguration_inst_SetPreferences(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return
}

// IgnoresViewportScaleLimits returns a Boolean value that determines whether a web view allows scaling of the webpage.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/2274633-ignoresviewportscalelimits?language=objc for details.
func (x gen_WKWebViewConfiguration) IgnoresViewportScaleLimits() bool {
	ret := C.WKWebViewConfiguration_inst_IgnoresViewportScaleLimits(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetIgnoresViewportScaleLimits returns a Boolean value that determines whether a web view allows scaling of the webpage.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/2274633-ignoresviewportscalelimits?language=objc for details.
func (x gen_WKWebViewConfiguration) SetIgnoresViewportScaleLimits(
	value bool,
) {
	C.WKWebViewConfiguration_inst_SetIgnoresViewportScaleLimits(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// SuppressesIncrementalRendering returns a Boolean value that indicates whether the web view suppresses content rendering until the content is fully loaded into memory.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395663-suppressesincrementalrendering?language=objc for details.
func (x gen_WKWebViewConfiguration) SuppressesIncrementalRendering() bool {
	ret := C.WKWebViewConfiguration_inst_SuppressesIncrementalRendering(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetSuppressesIncrementalRendering returns a Boolean value that indicates whether the web view suppresses content rendering until the content is fully loaded into memory.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395663-suppressesincrementalrendering?language=objc for details.
func (x gen_WKWebViewConfiguration) SetSuppressesIncrementalRendering(
	value bool,
) {
	C.WKWebViewConfiguration_inst_SetSuppressesIncrementalRendering(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AllowsInlineMediaPlayback returns a Boolean value that indicates whether HTML5 videos play inline or use the native full-screen controller.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1614793-allowsinlinemediaplayback?language=objc for details.
func (x gen_WKWebViewConfiguration) AllowsInlineMediaPlayback() bool {
	ret := C.WKWebViewConfiguration_inst_AllowsInlineMediaPlayback(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsInlineMediaPlayback returns a Boolean value that indicates whether HTML5 videos play inline or use the native full-screen controller.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1614793-allowsinlinemediaplayback?language=objc for details.
func (x gen_WKWebViewConfiguration) SetAllowsInlineMediaPlayback(
	value bool,
) {
	C.WKWebViewConfiguration_inst_SetAllowsInlineMediaPlayback(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AllowsAirPlayForMediaPlayback returns a Boolean value that indicates whether the web view allows media playback over AirPlay.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395673-allowsairplayformediaplayback?language=objc for details.
func (x gen_WKWebViewConfiguration) AllowsAirPlayForMediaPlayback() bool {
	ret := C.WKWebViewConfiguration_inst_AllowsAirPlayForMediaPlayback(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsAirPlayForMediaPlayback returns a Boolean value that indicates whether the web view allows media playback over AirPlay.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395673-allowsairplayformediaplayback?language=objc for details.
func (x gen_WKWebViewConfiguration) SetAllowsAirPlayForMediaPlayback(
	value bool,
) {
	C.WKWebViewConfiguration_inst_SetAllowsAirPlayForMediaPlayback(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// AllowsPictureInPictureMediaPlayback returns a Boolean value that indicates whether HTML5 videos can play Picture in Picture.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1614792-allowspictureinpicturemediaplayb?language=objc for details.
func (x gen_WKWebViewConfiguration) AllowsPictureInPictureMediaPlayback() bool {
	ret := C.WKWebViewConfiguration_inst_AllowsPictureInPictureMediaPlayback(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetAllowsPictureInPictureMediaPlayback returns a Boolean value that indicates whether HTML5 videos can play Picture in Picture.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1614792-allowspictureinpicturemediaplayb?language=objc for details.
func (x gen_WKWebViewConfiguration) SetAllowsPictureInPictureMediaPlayback(
	value bool,
) {
	C.WKWebViewConfiguration_inst_SetAllowsPictureInPictureMediaPlayback(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// UpgradeKnownHostsToHTTPS is undocumented.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/3752243-upgradeknownhoststohttps?language=objc for details.
func (x gen_WKWebViewConfiguration) UpgradeKnownHostsToHTTPS() bool {
	ret := C.WKWebViewConfiguration_inst_UpgradeKnownHostsToHTTPS(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetUpgradeKnownHostsToHTTPS is undocumented.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/3752243-upgradeknownhoststohttps?language=objc for details.
func (x gen_WKWebViewConfiguration) SetUpgradeKnownHostsToHTTPS(
	value bool,
) {
	C.WKWebViewConfiguration_inst_SetUpgradeKnownHostsToHTTPS(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

type WKPreferencesRef interface {
	Pointer() uintptr
	Init_AsWKPreferences() WKPreferences
}

type gen_WKPreferences struct {
	objc.Object
}

func WKPreferences_FromPointer(ptr unsafe.Pointer) WKPreferences {
	return WKPreferences{gen_WKPreferences{
		objc.Object_FromPointer(ptr),
	}}
}

func WKPreferences_FromRef(ref objc.Ref) WKPreferences {
	return WKPreferences_FromPointer(unsafe.Pointer(ref.Pointer()))
}

// SetValueForKey is undocumented.
func (x gen_WKPreferences) SetValueForKey(
	value objc.Ref,
	key core.NSStringRef,
) {
	C.WKPreferences_inst_SetValueForKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
		objc.RefPointer(key),
	)

	return
}

// Init initializes a new instance of the WKPreferences class.
func (x gen_WKPreferences) Init() WKPreferences {
	ret := C.WKPreferences_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return WKPreferences_FromPointer(ret)
}

// Init_AsWKPreferences is a typed version of Init.
func (x gen_WKPreferences) Init_AsWKPreferences() WKPreferences {
	ret := C.WKPreferences_inst_Init(
		unsafe.Pointer(x.Pointer()),
	)

	return WKPreferences_FromPointer(ret)
}

// MinimumFontSize returns the minimum font size, in points.
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/1537155-minimumfontsize?language=objc for details.
func (x gen_WKPreferences) MinimumFontSize() core.CGFloat {
	ret := C.WKPreferences_inst_MinimumFontSize(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)
}

// SetMinimumFontSize returns the minimum font size, in points.
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/1537155-minimumfontsize?language=objc for details.
func (x gen_WKPreferences) SetMinimumFontSize(
	value core.CGFloat,
) {
	C.WKPreferences_inst_SetMinimumFontSize(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return
}

// TabFocusesLinks returns a Boolean value that indicates whether pressing the tab key changes the focus to links and form controls.
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/2818595-tabfocuseslinks?language=objc for details.
func (x gen_WKPreferences) TabFocusesLinks() bool {
	ret := C.WKPreferences_inst_TabFocusesLinks(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetTabFocusesLinks returns a Boolean value that indicates whether pressing the tab key changes the focus to links and form controls.
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/2818595-tabfocuseslinks?language=objc for details.
func (x gen_WKPreferences) SetTabFocusesLinks(
	value bool,
) {
	C.WKPreferences_inst_SetTabFocusesLinks(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// JavaScriptCanOpenWindowsAutomatically returns a Boolean value that indicates whether JavaScript can open windows without user interaction.
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/1536573-javascriptcanopenwindowsautomati?language=objc for details.
func (x gen_WKPreferences) JavaScriptCanOpenWindowsAutomatically() bool {
	ret := C.WKPreferences_inst_JavaScriptCanOpenWindowsAutomatically(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetJavaScriptCanOpenWindowsAutomatically returns a Boolean value that indicates whether JavaScript can open windows without user interaction.
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/1536573-javascriptcanopenwindowsautomati?language=objc for details.
func (x gen_WKPreferences) SetJavaScriptCanOpenWindowsAutomatically(
	value bool,
) {
	C.WKPreferences_inst_SetJavaScriptCanOpenWindowsAutomatically(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsFraudulentWebsiteWarningEnabled returns a Boolean value that indicates whether the web view shows warnings for suspected fraudulent content, such as malware or phishing attemps.
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/3335219-fraudulentwebsitewarningenabled?language=objc for details.
func (x gen_WKPreferences) IsFraudulentWebsiteWarningEnabled() bool {
	ret := C.WKPreferences_inst_IsFraudulentWebsiteWarningEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetFraudulentWebsiteWarningEnabled returns a Boolean value that indicates whether the web view shows warnings for suspected fraudulent content, such as malware or phishing attemps.
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/3335219-fraudulentwebsitewarningenabled?language=objc for details.
func (x gen_WKPreferences) SetFraudulentWebsiteWarningEnabled(
	value bool,
) {
	C.WKPreferences_inst_SetFraudulentWebsiteWarningEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}

// IsTextInteractionEnabled is undocumented.
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/3727362-textinteractionenabled?language=objc for details.
func (x gen_WKPreferences) IsTextInteractionEnabled() bool {
	ret := C.WKPreferences_inst_IsTextInteractionEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)
}

// SetTextInteractionEnabled is undocumented.
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/3727362-textinteractionenabled?language=objc for details.
func (x gen_WKPreferences) SetTextInteractionEnabled(
	value bool,
) {
	C.WKPreferences_inst_SetTextInteractionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return
}
