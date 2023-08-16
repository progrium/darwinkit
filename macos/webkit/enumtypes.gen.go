// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import "math"

// The media types that require a user gesture to begin playing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkaudiovisualmediatypes?language=objc
type AudiovisualMediaTypes uint

const (
	AudiovisualMediaTypeAll   AudiovisualMediaTypes = math.MaxUint
	AudiovisualMediaTypeAudio AudiovisualMediaTypes = 1
	AudiovisualMediaTypeNone  AudiovisualMediaTypes = 0
	AudiovisualMediaTypeVideo AudiovisualMediaTypes = 2
)

// Constants that indicate how to render web view content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkcontentmode?language=objc
type ContentMode int

const (
	ContentModeDesktop     ContentMode = 2
	ContentModeMobile      ContentMode = 1
	ContentModeRecommended ContentMode = 0
)

// An enumeration with cases that indicate whether to proceed with a redirect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloadredirectpolicy?language=objc
type DownloadRedirectPolicy int

const (
	DownloadRedirectPolicyAllow  DownloadRedirectPolicy = 1
	DownloadRedirectPolicyCancel DownloadRedirectPolicy = 0
)

// Possible error values that WebKit APIs can return. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkerrorcode?language=objc
type ErrorCode int

const (
	ErrorAttributedStringContentFailedToLoad ErrorCode = 10
	ErrorAttributedStringContentLoadTimedOut ErrorCode = 11
	ErrorContentRuleListStoreCompileFailed   ErrorCode = 6
	ErrorContentRuleListStoreLookUpFailed    ErrorCode = 7
	ErrorContentRuleListStoreRemoveFailed    ErrorCode = 8
	ErrorContentRuleListStoreVersionMismatch ErrorCode = 9
	ErrorJavaScriptAppBoundDomain            ErrorCode = 14
	ErrorJavaScriptExceptionOccurred         ErrorCode = 4
	ErrorJavaScriptInvalidFrameTarget        ErrorCode = 12
	ErrorJavaScriptResultTypeIsUnsupported   ErrorCode = 5
	ErrorNavigationAppBoundDomain            ErrorCode = 13
	ErrorUnknown                             ErrorCode = 1
	ErrorWebContentProcessTerminated         ErrorCode = 2
	ErrorWebViewInvalidated                  ErrorCode = 3
)

// An enumeration that describes whether a media device, like a camera or microphone, is currently capturing audio or video. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkmediacapturestate?language=objc
type MediaCaptureState int

const (
	MediaCaptureStateActive MediaCaptureState = 1
	MediaCaptureStateMuted  MediaCaptureState = 2
	MediaCaptureStateNone   MediaCaptureState = 0
)

// An enumeration listing the types of media devices that can capture audio, video, or both. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkmediacapturetype?language=objc
type MediaCaptureType int

const (
	MediaCaptureTypeCamera              MediaCaptureType = 0
	MediaCaptureTypeCameraAndMicrophone MediaCaptureType = 2
	MediaCaptureTypeMicrophone          MediaCaptureType = 1
)

// An enumeration that describes whether an audio or video presentation is playing, paused, or suspended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkmediaplaybackstate?language=objc
type MediaPlaybackState int

const (
	MediaPlaybackStateNone      MediaPlaybackState = 0
	MediaPlaybackStatePaused    MediaPlaybackState = 2
	MediaPlaybackStatePlaying   MediaPlaybackState = 1
	MediaPlaybackStateSuspended MediaPlaybackState = 3
)

// Constants that indicate whether to allow or cancel navigation to a webpage from an action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationactionpolicy?language=objc
type NavigationActionPolicy int

const (
	NavigationActionPolicyAllow    NavigationActionPolicy = 1
	NavigationActionPolicyCancel   NavigationActionPolicy = 0
	NavigationActionPolicyDownload NavigationActionPolicy = 2
)

// Constants that indicate whether to allow or cancel navigation to a webpage from a response. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationresponsepolicy?language=objc
type NavigationResponsePolicy int

const (
	NavigationResponsePolicyAllow    NavigationResponsePolicy = 1
	NavigationResponsePolicyCancel   NavigationResponsePolicy = 0
	NavigationResponsePolicyDownload NavigationResponsePolicy = 2
)

// The type of action that triggered the navigation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationtype?language=objc
type NavigationType int

const (
	NavigationTypeBackForward     NavigationType = 2
	NavigationTypeFormResubmitted NavigationType = 4
	NavigationTypeFormSubmitted   NavigationType = 1
	NavigationTypeLinkActivated   NavigationType = 0
	NavigationTypeOther           NavigationType = -1
	NavigationTypeReload          NavigationType = 3
)

// An enumeration of possible permission decisions for device resource access. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpermissiondecision?language=objc
type PermissionDecision int

const (
	PermissionDecisionDeny   PermissionDecision = 2
	PermissionDecisionGrant  PermissionDecision = 1
	PermissionDecisionPrompt PermissionDecision = 0
)

// The policy that determines the directionality of user interface elements in a web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuserinterfacedirectionpolicy?language=objc
type UserInterfaceDirectionPolicy int

const (
	UserInterfaceDirectionPolicyContent UserInterfaceDirectionPolicy = 0
	UserInterfaceDirectionPolicySystem  UserInterfaceDirectionPolicy = 1
)

// Constants for the times at which to inject script content into a webpage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuserscriptinjectiontime?language=objc
type UserScriptInjectionTime int

const (
	UserScriptInjectionTimeAtDocumentEnd   UserScriptInjectionTime = 1
	UserScriptInjectionTimeAtDocumentStart UserScriptInjectionTime = 0
)

// Specifies the caching model for a web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webcachemodel?language=objc
type WebCacheModel uint

const (
	WebCacheModelDocumentBrowser   WebCacheModel = 1
	WebCacheModelDocumentViewer    WebCacheModel = 0
	WebCacheModelPrimaryWebBrowser WebCacheModel = 2
)

// Actions that the destination object of a drag operation can perform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webdragdestinationaction?language=objc
type WebDragDestinationAction uint

const (
	WebDragDestinationActionAny   WebDragDestinationAction = 4294967295
	WebDragDestinationActionDHTML WebDragDestinationAction = 1
	WebDragDestinationActionEdit  WebDragDestinationAction = 2
	WebDragDestinationActionLoad  WebDragDestinationAction = 4
	WebDragDestinationActionNone  WebDragDestinationAction = 0
)

// Actions that the source object of a drag operation can perform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webdragsourceaction?language=objc
type WebDragSourceAction uint

const (
	WebDragSourceActionAny       WebDragSourceAction = 4294967295
	WebDragSourceActionDHTML     WebDragSourceAction = 1
	WebDragSourceActionImage     WebDragSourceAction = 2
	WebDragSourceActionLink      WebDragSourceAction = 4
	WebDragSourceActionNone      WebDragSourceAction = 0
	WebDragSourceActionSelection WebDragSourceAction = 8
)

// Possible values for the WebActionNavigationTypeKey key that appears in an action dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webnavigationtype?language=objc
type WebNavigationType int

const (
	WebNavigationTypeBackForward     WebNavigationType = 2
	WebNavigationTypeFormResubmitted WebNavigationType = 4
	WebNavigationTypeFormSubmitted   WebNavigationType = 1
	WebNavigationTypeLinkClicked     WebNavigationType = 0
	WebNavigationTypeOther           WebNavigationType = 5
	WebNavigationTypeReload          WebNavigationType = 3
)

// The type of user action that initiated a delegate message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webviewinsertaction?language=objc
type WebViewInsertAction int

const (
	WebViewInsertActionDropped WebViewInsertAction = 2
	WebViewInsertActionPasted  WebViewInsertAction = 1
	WebViewInsertActionTyped   WebViewInsertAction = 0
)
