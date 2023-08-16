// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WebView] class.
var WebViewClass = _WebViewClass{objc.GetClass("WKWebView")}

type _WebViewClass struct {
	objc.Class
}

// An interface definition for the [WebView] class.
type IWebView interface {
	appkit.IView
	SetMagnificationCenteredAtPoint(magnification float64, point coregraphics.Point)
	GoBack(sender objc.IObject) objc.Object
	PrintOperationWithPrintInfo(printInfo appkit.IPrintInfo) appkit.PrintOperation
	StopLoading(sender objc.IObject) objc.Object
	PauseAllMediaPlaybackWithCompletionHandler(completionHandler func())
	GoForward() Navigation
	GoToBackForwardListItem(item IBackForwardListItem) Navigation
	RequestMediaPlaybackStateWithCompletionHandler(completionHandler func(arg0 MediaPlaybackState))
	Reload(sender objc.IObject) objc.Object
	TakeSnapshotWithConfigurationCompletionHandler(snapshotConfiguration ISnapshotConfiguration, completionHandler func(snapshotImage appkit.Image, error foundation.Error))
	SetMicrophoneCaptureStateCompletionHandler(state MediaCaptureState, completionHandler func())
	LoadDataMIMETypeCharacterEncodingNameBaseURL(data []byte, MIMEType string, characterEncodingName string, baseURL foundation.IURL) Navigation
	CloseAllMediaPresentationsWithCompletionHandler(completionHandler func())
	ReloadFromOrigin(sender objc.IObject) objc.Object
	SetAllMediaPlaybackSuspendedCompletionHandler(suspended bool, completionHandler func())
	LoadSimulatedRequestResponseHTMLString(request foundation.IURLRequest, string_ string) Navigation
	LoadFileRequestAllowingReadAccessToURL(request foundation.IURLRequest, readAccessURL foundation.IURL) Navigation
	LoadRequest(request foundation.IURLRequest) Navigation
	StartDownloadUsingRequestCompletionHandler(request foundation.IURLRequest, completionHandler func(arg0 Download))
	EvaluateJavaScriptCompletionHandler(javaScriptString string, completionHandler func(arg0 objc.Object, error foundation.Error))
	ResumeDownloadFromResumeDataCompletionHandler(resumeData []byte, completionHandler func(arg0 Download))
	LoadFileURLAllowingReadAccessToURL(URL foundation.IURL, readAccessURL foundation.IURL) Navigation
	LoadHTMLStringBaseURL(string_ string, baseURL foundation.IURL) Navigation
	SetCameraCaptureStateCompletionHandler(state MediaCaptureState, completionHandler func())
	CreatePDFWithConfigurationCompletionHandler(pdfConfiguration IPDFConfiguration, completionHandler func(pdfDocumentData []byte, error foundation.Error))
	CreateWebArchiveDataWithCompletionHandler(completionHandler func(arg0 []byte, arg1 foundation.Error))
	FindStringWithConfigurationCompletionHandler(string_ string, configuration IFindConfiguration, completionHandler func(result FindResult))
	CallAsyncJavaScriptArgumentsInFrameInContentWorldCompletionHandler(functionBody string, arguments map[string]objc.IObject, frame IFrameInfo, contentWorld IContentWorld, completionHandler func(arg0 objc.Object, error foundation.Error))
	AllowsMagnification() bool
	SetAllowsMagnification(value bool)
	EstimatedProgress() float64
	IsLoading() bool
	AllowsBackForwardNavigationGestures() bool
	SetAllowsBackForwardNavigationGestures(value bool)
	HasOnlySecureContent() bool
	CanGoBack() bool
	CustomUserAgent() string
	SetCustomUserAgent(value string)
	MediaType() string
	SetMediaType(value string)
	ThemeColor() appkit.Color
	Magnification() float64
	SetMagnification(value float64)
	CanGoForward() bool
	UIDelegate() UIDelegateWrapper
	SetUIDelegate(value PUIDelegate)
	SetUIDelegateObject(valueObject objc.IObject)
	PageZoom() float64
	SetPageZoom(value float64)
	URL() foundation.URL
	Configuration() WebViewConfiguration
	BackForwardList() BackForwardList
	MicrophoneCaptureState() MediaCaptureState
	NavigationDelegate() NavigationDelegateWrapper
	SetNavigationDelegate(value PNavigationDelegate)
	SetNavigationDelegateObject(valueObject objc.IObject)
	Title() string
	CameraCaptureState() MediaCaptureState
	UnderPageBackgroundColor() appkit.Color
	SetUnderPageBackgroundColor(value appkit.IColor)
	AllowsLinkPreview() bool
	SetAllowsLinkPreview(value bool)
	InteractionState() objc.Object
	SetInteractionState(value objc.IObject)
}

// An object that displays interactive web content, such as for an in-app browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview?language=objc
type WebView struct {
	appkit.View
}

func WebViewFrom(ptr unsafe.Pointer) WebView {
	return WebView{
		View: appkit.ViewFrom(ptr),
	}
}

func (w_ WebView) InitWithFrameConfiguration(frame coregraphics.Rect, configuration IWebViewConfiguration) WebView {
	rv := objc.Call[WebView](w_, objc.Sel("initWithFrame:configuration:"), frame, objc.Ptr(configuration))
	return rv
}

// Creates a web view and initializes it with the specified frame and configuration data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414998-initwithframe?language=objc
func WebView_InitWithFrameConfiguration(frame coregraphics.Rect, configuration IWebViewConfiguration) WebView {
	return WebViewClass.Alloc().InitWithFrameConfiguration(frame, configuration)
}

func (wc _WebViewClass) Alloc() WebView {
	rv := objc.Call[WebView](wc, objc.Sel("alloc"))
	return rv
}

func WebView_Alloc() WebView {
	return WebViewClass.Alloc()
}

func (wc _WebViewClass) New() WebView {
	rv := objc.Call[WebView](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWebView() WebView {
	return WebViewClass.New()
}

func (w_ WebView) Init() WebView {
	rv := objc.Call[WebView](w_, objc.Sel("init"))
	return rv
}

func (w_ WebView) InitWithFrame(frameRect foundation.Rect) WebView {
	rv := objc.Call[WebView](w_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func WebView_InitWithFrame(frameRect foundation.Rect) WebView {
	return WebViewClass.Alloc().InitWithFrame(frameRect)
}

// Scales the page content and centers the result on the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414996-setmagnification?language=objc
func (w_ WebView) SetMagnificationCenteredAtPoint(magnification float64, point coregraphics.Point) {
	objc.Call[objc.Void](w_, objc.Sel("setMagnification:centeredAtPoint:"), magnification, point)
}

// Navigates to the back item in the back-forward list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414975-goback?language=objc
func (w_ WebView) GoBack(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("goBack:"), sender)
	return rv
}

// Returns the print operation object to use when printing the contents of the web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3516861-printoperationwithprintinfo?language=objc
func (w_ WebView) PrintOperationWithPrintInfo(printInfo appkit.IPrintInfo) appkit.PrintOperation {
	rv := objc.Call[appkit.PrintOperation](w_, objc.Sel("printOperationWithPrintInfo:"), objc.Ptr(printInfo))
	return rv
}

// Stops loading all resources on the current page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1415013-stoploading?language=objc
func (w_ WebView) StopLoading(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("stopLoading:"), sender)
	return rv
}

// Pauses playback of all media in the web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3752240-pauseallmediaplaybackwithcomplet?language=objc
func (w_ WebView) PauseAllMediaPlaybackWithCompletionHandler(completionHandler func()) {
	objc.Call[objc.Void](w_, objc.Sel("pauseAllMediaPlaybackWithCompletionHandler:"), completionHandler)
}

// Navigates to the forward item in the back-forward list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414993-goforward?language=objc
func (w_ WebView) GoForward() Navigation {
	rv := objc.Call[Navigation](w_, objc.Sel("goForward"))
	return rv
}

// Navigates to an item from the back-forward list and sets it as the current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414991-gotobackforwardlistitem?language=objc
func (w_ WebView) GoToBackForwardListItem(item IBackForwardListItem) Navigation {
	rv := objc.Call[Navigation](w_, objc.Sel("goToBackForwardListItem:"), objc.Ptr(item))
	return rv
}

// Requests the playback status of media in the web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3752241-requestmediaplaybackstatewithcom?language=objc
func (w_ WebView) RequestMediaPlaybackStateWithCompletionHandler(completionHandler func(arg0 MediaPlaybackState)) {
	objc.Call[objc.Void](w_, objc.Sel("requestMediaPlaybackStateWithCompletionHandler:"), completionHandler)
}

// Returns a Boolean value that indicates whether WebKit natively supports resources with the specified URL scheme. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/2875370-handlesurlscheme?language=objc
func (wc _WebViewClass) HandlesURLScheme(urlScheme string) bool {
	rv := objc.Call[bool](wc, objc.Sel("handlesURLScheme:"), urlScheme)
	return rv
}

// Returns a Boolean value that indicates whether WebKit natively supports resources with the specified URL scheme. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/2875370-handlesurlscheme?language=objc
func WebView_HandlesURLScheme(urlScheme string) bool {
	return WebViewClass.HandlesURLScheme(urlScheme)
}

// Reloads the current webpage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414987-reload?language=objc
func (w_ WebView) Reload(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("reload:"), sender)
	return rv
}

// Generates a platform-native image from the web view’s contents asynchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/2873260-takesnapshotwithconfiguration?language=objc
func (w_ WebView) TakeSnapshotWithConfigurationCompletionHandler(snapshotConfiguration ISnapshotConfiguration, completionHandler func(snapshotImage appkit.Image, error foundation.Error)) {
	objc.Call[objc.Void](w_, objc.Sel("takeSnapshotWithConfiguration:completionHandler:"), objc.Ptr(snapshotConfiguration), completionHandler)
}

// Changes whether the webpage is using the microphone to capture audio. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3763098-setmicrophonecapturestate?language=objc
func (w_ WebView) SetMicrophoneCaptureStateCompletionHandler(state MediaCaptureState, completionHandler func()) {
	objc.Call[objc.Void](w_, objc.Sel("setMicrophoneCaptureState:completionHandler:"), state, completionHandler)
}

// Loads the content of the specified data object and navigates to it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1415011-loaddata?language=objc
func (w_ WebView) LoadDataMIMETypeCharacterEncodingNameBaseURL(data []byte, MIMEType string, characterEncodingName string, baseURL foundation.IURL) Navigation {
	rv := objc.Call[Navigation](w_, objc.Sel("loadData:MIMEType:characterEncodingName:baseURL:"), data, MIMEType, characterEncodingName, objc.Ptr(baseURL))
	return rv
}

// Closes all media the web view is presenting, including picture-in-picture video and fullscreen video. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3752235-closeallmediapresentationswithco?language=objc
func (w_ WebView) CloseAllMediaPresentationsWithCompletionHandler(completionHandler func()) {
	objc.Call[objc.Void](w_, objc.Sel("closeAllMediaPresentationsWithCompletionHandler:"), completionHandler)
}

// Reloads the current webpage, and performs end-to-end revalidation of the content using cache-validating conditionals, if possible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414989-reloadfromorigin?language=objc
func (w_ WebView) ReloadFromOrigin(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("reloadFromOrigin:"), sender)
	return rv
}

// Changes whether the webpage is suspending playback of all media in the page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3752242-setallmediaplaybacksuspended?language=objc
func (w_ WebView) SetAllMediaPlaybackSuspendedCompletionHandler(suspended bool, completionHandler func()) {
	objc.Call[objc.Void](w_, objc.Sel("setAllMediaPlaybackSuspended:completionHandler:"), suspended, completionHandler)
}

// Loads the web content from the HTML you provide as if the HTML were the response to the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3763095-loadsimulatedrequest?language=objc
func (w_ WebView) LoadSimulatedRequestResponseHTMLString(request foundation.IURLRequest, string_ string) Navigation {
	rv := objc.Call[Navigation](w_, objc.Sel("loadSimulatedRequest:responseHTMLString:"), objc.Ptr(request), string_)
	return rv
}

// Loads the web content from the file the URL request object specifies and navigates to that content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3752237-loadfilerequest?language=objc
func (w_ WebView) LoadFileRequestAllowingReadAccessToURL(request foundation.IURLRequest, readAccessURL foundation.IURL) Navigation {
	rv := objc.Call[Navigation](w_, objc.Sel("loadFileRequest:allowingReadAccessToURL:"), objc.Ptr(request), objc.Ptr(readAccessURL))
	return rv
}

// Loads the web content that the specified URL request object references and navigates to that content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414954-loadrequest?language=objc
func (w_ WebView) LoadRequest(request foundation.IURLRequest) Navigation {
	rv := objc.Call[Navigation](w_, objc.Sel("loadRequest:"), objc.Ptr(request))
	return rv
}

// Starts to download the resource at the URL in the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3727368-startdownloadusingrequest?language=objc
func (w_ WebView) StartDownloadUsingRequestCompletionHandler(request foundation.IURLRequest, completionHandler func(arg0 Download)) {
	objc.Call[objc.Void](w_, objc.Sel("startDownloadUsingRequest:completionHandler:"), objc.Ptr(request), completionHandler)
}

// Evaluates the specified JavaScript string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1415017-evaluatejavascript?language=objc
func (w_ WebView) EvaluateJavaScriptCompletionHandler(javaScriptString string, completionHandler func(arg0 objc.Object, error foundation.Error)) {
	objc.Call[objc.Void](w_, objc.Sel("evaluateJavaScript:completionHandler:"), javaScriptString, completionHandler)
}

// Resumes a failed or canceled download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3727367-resumedownloadfromresumedata?language=objc
func (w_ WebView) ResumeDownloadFromResumeDataCompletionHandler(resumeData []byte, completionHandler func(arg0 Download)) {
	objc.Call[objc.Void](w_, objc.Sel("resumeDownloadFromResumeData:completionHandler:"), resumeData, completionHandler)
}

// Loads the web content from the specified file and navigates to it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414973-loadfileurl?language=objc
func (w_ WebView) LoadFileURLAllowingReadAccessToURL(URL foundation.IURL, readAccessURL foundation.IURL) Navigation {
	rv := objc.Call[Navigation](w_, objc.Sel("loadFileURL:allowingReadAccessToURL:"), objc.Ptr(URL), objc.Ptr(readAccessURL))
	return rv
}

// Loads the contents of the specified HTML string and navigates to it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1415004-loadhtmlstring?language=objc
func (w_ WebView) LoadHTMLStringBaseURL(string_ string, baseURL foundation.IURL) Navigation {
	rv := objc.Call[Navigation](w_, objc.Sel("loadHTMLString:baseURL:"), string_, objc.Ptr(baseURL))
	return rv
}

// Changes whether the webpage is using the camera to capture images or video. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3763097-setcameracapturestate?language=objc
func (w_ WebView) SetCameraCaptureStateCompletionHandler(state MediaCaptureState, completionHandler func()) {
	objc.Call[objc.Void](w_, objc.Sel("setCameraCaptureState:completionHandler:"), state, completionHandler)
}

// Generates PDF data from the web view’s contents asynchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3516407-createpdfwithconfiguration?language=objc
func (w_ WebView) CreatePDFWithConfigurationCompletionHandler(pdfConfiguration IPDFConfiguration, completionHandler func(pdfDocumentData []byte, error foundation.Error)) {
	objc.Call[objc.Void](w_, objc.Sel("createPDFWithConfiguration:completionHandler:"), objc.Ptr(pdfConfiguration), completionHandler)
}

// Creates a web archive of the web view’s contents asynchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3516408-createwebarchivedatawithcompleti?language=objc
func (w_ WebView) CreateWebArchiveDataWithCompletionHandler(completionHandler func(arg0 []byte, arg1 foundation.Error)) {
	objc.Call[objc.Void](w_, objc.Sel("createWebArchiveDataWithCompletionHandler:"), completionHandler)
}

// Searches for the specified string in the web view’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3516409-findstring?language=objc
func (w_ WebView) FindStringWithConfigurationCompletionHandler(string_ string, configuration IFindConfiguration, completionHandler func(result FindResult)) {
	objc.Call[objc.Void](w_, objc.Sel("findString:withConfiguration:completionHandler:"), string_, objc.Ptr(configuration), completionHandler)
}

// Executes the specified string as an asynchronous JavaScript function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3656355-callasyncjavascript?language=objc
func (w_ WebView) CallAsyncJavaScriptArgumentsInFrameInContentWorldCompletionHandler(functionBody string, arguments map[string]objc.IObject, frame IFrameInfo, contentWorld IContentWorld, completionHandler func(arg0 objc.Object, error foundation.Error)) {
	objc.Call[objc.Void](w_, objc.Sel("callAsyncJavaScript:arguments:inFrame:inContentWorld:completionHandler:"), functionBody, arguments, objc.Ptr(frame), objc.Ptr(contentWorld), completionHandler)
}

// A Boolean value that indicates whether magnify gestures change the web view’s magnification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414983-allowsmagnification?language=objc
func (w_ WebView) AllowsMagnification() bool {
	rv := objc.Call[bool](w_, objc.Sel("allowsMagnification"))
	return rv
}

// A Boolean value that indicates whether magnify gestures change the web view’s magnification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414983-allowsmagnification?language=objc
func (w_ WebView) SetAllowsMagnification(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setAllowsMagnification:"), value)
}

// An estimate of what fraction of the current navigation has been loaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1415007-estimatedprogress?language=objc
func (w_ WebView) EstimatedProgress() float64 {
	rv := objc.Call[float64](w_, objc.Sel("estimatedProgress"))
	return rv
}

// A Boolean value that indicates whether the view is currently loading content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414964-loading?language=objc
func (w_ WebView) IsLoading() bool {
	rv := objc.Call[bool](w_, objc.Sel("isLoading"))
	return rv
}

// A Boolean value that indicates whether horizontal swipe gestures trigger backward and forward page navigation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414995-allowsbackforwardnavigationgestu?language=objc
func (w_ WebView) AllowsBackForwardNavigationGestures() bool {
	rv := objc.Call[bool](w_, objc.Sel("allowsBackForwardNavigationGestures"))
	return rv
}

// A Boolean value that indicates whether horizontal swipe gestures trigger backward and forward page navigation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414995-allowsbackforwardnavigationgestu?language=objc
func (w_ WebView) SetAllowsBackForwardNavigationGestures(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setAllowsBackForwardNavigationGestures:"), value)
}

// A Boolean value that indicates whether the web view loaded all resources on the page through securely encrypted connections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1415002-hasonlysecurecontent?language=objc
func (w_ WebView) HasOnlySecureContent() bool {
	rv := objc.Call[bool](w_, objc.Sel("hasOnlySecureContent"))
	return rv
}

// A Boolean value that indicates whether there is a valid back item in the back-forward list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414966-cangoback?language=objc
func (w_ WebView) CanGoBack() bool {
	rv := objc.Call[bool](w_, objc.Sel("canGoBack"))
	return rv
}

// The custom user agent string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414950-customuseragent?language=objc
func (w_ WebView) CustomUserAgent() string {
	rv := objc.Call[string](w_, objc.Sel("customUserAgent"))
	return rv
}

// The custom user agent string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414950-customuseragent?language=objc
func (w_ WebView) SetCustomUserAgent(value string) {
	objc.Call[objc.Void](w_, objc.Sel("setCustomUserAgent:"), value)
}

// The media type for the contents of the web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3516410-mediatype?language=objc
func (w_ WebView) MediaType() string {
	rv := objc.Call[string](w_, objc.Sel("mediaType"))
	return rv
}

// The media type for the contents of the web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3516410-mediatype?language=objc
func (w_ WebView) SetMediaType(value string) {
	objc.Call[objc.Void](w_, objc.Sel("setMediaType:"), value)
}

// The theme color that the system gets from the first valid meta tag in the webpage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3794258-themecolor?language=objc
func (w_ WebView) ThemeColor() appkit.Color {
	rv := objc.Call[appkit.Color](w_, objc.Sel("themeColor"))
	return rv
}

// The factor by which the page content is currently scaled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414985-magnification?language=objc
func (w_ WebView) Magnification() float64 {
	rv := objc.Call[float64](w_, objc.Sel("magnification"))
	return rv
}

// The factor by which the page content is currently scaled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414985-magnification?language=objc
func (w_ WebView) SetMagnification(value float64) {
	objc.Call[objc.Void](w_, objc.Sel("setMagnification:"), value)
}

// A Boolean value that indicates whether there is a valid forward item in the back-forward list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414962-cangoforward?language=objc
func (w_ WebView) CanGoForward() bool {
	rv := objc.Call[bool](w_, objc.Sel("canGoForward"))
	return rv
}

// The object you use to integrate custom user interface elements, such as contextual menus or panels, into web view interactions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1415009-uidelegate?language=objc
func (w_ WebView) UIDelegate() UIDelegateWrapper {
	rv := objc.Call[UIDelegateWrapper](w_, objc.Sel("UIDelegate"))
	return rv
}

// The object you use to integrate custom user interface elements, such as contextual menus or panels, into web view interactions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1415009-uidelegate?language=objc
func (w_ WebView) SetUIDelegate(value PUIDelegate) {
	po0 := objc.WrapAsProtocol("WKUIDelegate", value)
	objc.SetAssociatedObject(w_, objc.AssociationKey("setUIDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](w_, objc.Sel("setUIDelegate:"), po0)
}

// The object you use to integrate custom user interface elements, such as contextual menus or panels, into web view interactions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1415009-uidelegate?language=objc
func (w_ WebView) SetUIDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("setUIDelegate:"), objc.Ptr(valueObject))
}

// The scale factor by which the web view scales content relative to its bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3516411-pagezoom?language=objc
func (w_ WebView) PageZoom() float64 {
	rv := objc.Call[float64](w_, objc.Sel("pageZoom"))
	return rv
}

// The scale factor by which the web view scales content relative to its bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3516411-pagezoom?language=objc
func (w_ WebView) SetPageZoom(value float64) {
	objc.Call[objc.Void](w_, objc.Sel("setPageZoom:"), value)
}

// The URL for the current webpage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1415005-url?language=objc
func (w_ WebView) URL() foundation.URL {
	rv := objc.Call[foundation.URL](w_, objc.Sel("URL"))
	return rv
}

// The object that contains the configuration details for the web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414979-configuration?language=objc
func (w_ WebView) Configuration() WebViewConfiguration {
	rv := objc.Call[WebViewConfiguration](w_, objc.Sel("configuration"))
	return rv
}

// The web view's back-forward list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414977-backforwardlist?language=objc
func (w_ WebView) BackForwardList() BackForwardList {
	rv := objc.Call[BackForwardList](w_, objc.Sel("backForwardList"))
	return rv
}

// An enumeration case that indicates whether the webpage is using the microphone to capture audio. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3763096-microphonecapturestate?language=objc
func (w_ WebView) MicrophoneCaptureState() MediaCaptureState {
	rv := objc.Call[MediaCaptureState](w_, objc.Sel("microphoneCaptureState"))
	return rv
}

// The object you use to manage navigation behavior for the web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414971-navigationdelegate?language=objc
func (w_ WebView) NavigationDelegate() NavigationDelegateWrapper {
	rv := objc.Call[NavigationDelegateWrapper](w_, objc.Sel("navigationDelegate"))
	return rv
}

// The object you use to manage navigation behavior for the web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414971-navigationdelegate?language=objc
func (w_ WebView) SetNavigationDelegate(value PNavigationDelegate) {
	po0 := objc.WrapAsProtocol("WKNavigationDelegate", value)
	objc.SetAssociatedObject(w_, objc.AssociationKey("setNavigationDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](w_, objc.Sel("setNavigationDelegate:"), po0)
}

// The object you use to manage navigation behavior for the web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1414971-navigationdelegate?language=objc
func (w_ WebView) SetNavigationDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("setNavigationDelegate:"), objc.Ptr(valueObject))
}

// The page title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1415015-title?language=objc
func (w_ WebView) Title() string {
	rv := objc.Call[string](w_, objc.Sel("title"))
	return rv
}

// An enumeration case that indicates whether the webpage is using the camera to capture images or video. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3763093-cameracapturestate?language=objc
func (w_ WebView) CameraCaptureState() MediaCaptureState {
	rv := objc.Call[MediaCaptureState](w_, objc.Sel("cameraCaptureState"))
	return rv
}

// The color the web view displays behind the active page, visible when the user scrolls beyond the bounds of the page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3850574-underpagebackgroundcolor?language=objc
func (w_ WebView) UnderPageBackgroundColor() appkit.Color {
	rv := objc.Call[appkit.Color](w_, objc.Sel("underPageBackgroundColor"))
	return rv
}

// The color the web view displays behind the active page, visible when the user scrolls beyond the bounds of the page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3850574-underpagebackgroundcolor?language=objc
func (w_ WebView) SetUnderPageBackgroundColor(value appkit.IColor) {
	objc.Call[objc.Void](w_, objc.Sel("setUnderPageBackgroundColor:"), objc.Ptr(value))
}

// A Boolean value that determines whether pressing a link displays a preview of the destination for the link. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1415000-allowslinkpreview?language=objc
func (w_ WebView) AllowsLinkPreview() bool {
	rv := objc.Call[bool](w_, objc.Sel("allowsLinkPreview"))
	return rv
}

// A Boolean value that determines whether pressing a link displays a preview of the destination for the link. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/1415000-allowslinkpreview?language=objc
func (w_ WebView) SetAllowsLinkPreview(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setAllowsLinkPreview:"), value)
}

// An object you use to capture the current state of interaction in a web view so that you can restore that state later to another web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3752236-interactionstate?language=objc
func (w_ WebView) InteractionState() objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("interactionState"))
	return rv
}

// An object you use to capture the current state of interaction in a web view so that you can restore that state later to another web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/3752236-interactionstate?language=objc
func (w_ WebView) SetInteractionState(value objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("setInteractionState:"), value)
}
