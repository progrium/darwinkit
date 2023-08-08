// AUTO-GENERATED CODE, DO NOT MODIFY
package webkit

import "math"

type AudiovisualMediaTypes uint

const AudiovisualMediaTypeAll AudiovisualMediaTypes = math.MaxUint
const AudiovisualMediaTypeAudio AudiovisualMediaTypes = 1
const AudiovisualMediaTypeNone AudiovisualMediaTypes = 0
const AudiovisualMediaTypeVideo AudiovisualMediaTypes = 2

type ContentMode int

const ContentModeDesktop ContentMode = 2
const ContentModeMobile ContentMode = 1
const ContentModeRecommended ContentMode = 0

type DownloadRedirectPolicy int

const DownloadRedirectPolicyAllow DownloadRedirectPolicy = 1
const DownloadRedirectPolicyCancel DownloadRedirectPolicy = 0

type ErrorCode int

const ErrorAttributedStringContentFailedToLoad ErrorCode = 10
const ErrorAttributedStringContentLoadTimedOut ErrorCode = 11
const ErrorContentRuleListStoreCompileFailed ErrorCode = 6
const ErrorContentRuleListStoreLookUpFailed ErrorCode = 7
const ErrorContentRuleListStoreRemoveFailed ErrorCode = 8
const ErrorContentRuleListStoreVersionMismatch ErrorCode = 9
const ErrorJavaScriptAppBoundDomain ErrorCode = 14
const ErrorJavaScriptExceptionOccurred ErrorCode = 4
const ErrorJavaScriptInvalidFrameTarget ErrorCode = 12
const ErrorJavaScriptResultTypeIsUnsupported ErrorCode = 5
const ErrorNavigationAppBoundDomain ErrorCode = 13
const ErrorUnknown ErrorCode = 1
const ErrorWebContentProcessTerminated ErrorCode = 2
const ErrorWebViewInvalidated ErrorCode = 3

type MediaCaptureState int

const MediaCaptureStateActive MediaCaptureState = 1
const MediaCaptureStateMuted MediaCaptureState = 2
const MediaCaptureStateNone MediaCaptureState = 0

type MediaCaptureType int

const MediaCaptureTypeCamera MediaCaptureType = 0
const MediaCaptureTypeCameraAndMicrophone MediaCaptureType = 2
const MediaCaptureTypeMicrophone MediaCaptureType = 1

type MediaPlaybackState int

const MediaPlaybackStateNone MediaPlaybackState = 0
const MediaPlaybackStatePaused MediaPlaybackState = 2
const MediaPlaybackStatePlaying MediaPlaybackState = 1
const MediaPlaybackStateSuspended MediaPlaybackState = 3

type NavigationActionPolicy int

const NavigationActionPolicyAllow NavigationActionPolicy = 1
const NavigationActionPolicyCancel NavigationActionPolicy = 0
const NavigationActionPolicyDownload NavigationActionPolicy = 2

type NavigationResponsePolicy int

const NavigationResponsePolicyAllow NavigationResponsePolicy = 1
const NavigationResponsePolicyCancel NavigationResponsePolicy = 0
const NavigationResponsePolicyDownload NavigationResponsePolicy = 2

type NavigationType int

const NavigationTypeBackForward NavigationType = 2
const NavigationTypeFormResubmitted NavigationType = 4
const NavigationTypeFormSubmitted NavigationType = 1
const NavigationTypeLinkActivated NavigationType = 0
const NavigationTypeOther NavigationType = -1
const NavigationTypeReload NavigationType = 3

type PermissionDecision int

const PermissionDecisionDeny PermissionDecision = 2
const PermissionDecisionGrant PermissionDecision = 1
const PermissionDecisionPrompt PermissionDecision = 0

type UserInterfaceDirectionPolicy int

const UserInterfaceDirectionPolicyContent UserInterfaceDirectionPolicy = 0
const UserInterfaceDirectionPolicySystem UserInterfaceDirectionPolicy = 1

type UserScriptInjectionTime int

const UserScriptInjectionTimeAtDocumentEnd UserScriptInjectionTime = 1
const UserScriptInjectionTimeAtDocumentStart UserScriptInjectionTime = 0

type WebCacheModel uint

const WebCacheModelDocumentBrowser WebCacheModel = 1
const WebCacheModelDocumentViewer WebCacheModel = 0
const WebCacheModelPrimaryWebBrowser WebCacheModel = 2

type WebDragDestinationAction uint

const WebDragDestinationActionAny WebDragDestinationAction = 4294967295
const WebDragDestinationActionDHTML WebDragDestinationAction = 1
const WebDragDestinationActionEdit WebDragDestinationAction = 2
const WebDragDestinationActionLoad WebDragDestinationAction = 4
const WebDragDestinationActionNone WebDragDestinationAction = 0

type WebDragSourceAction uint

const WebDragSourceActionAny WebDragSourceAction = 4294967295
const WebDragSourceActionDHTML WebDragSourceAction = 1
const WebDragSourceActionImage WebDragSourceAction = 2
const WebDragSourceActionLink WebDragSourceAction = 4
const WebDragSourceActionNone WebDragSourceAction = 0
const WebDragSourceActionSelection WebDragSourceAction = 8

type WebNavigationType int

const WebNavigationTypeBackForward WebNavigationType = 2
const WebNavigationTypeFormResubmitted WebNavigationType = 4
const WebNavigationTypeFormSubmitted WebNavigationType = 1
const WebNavigationTypeLinkClicked WebNavigationType = 0
const WebNavigationTypeOther WebNavigationType = 5
const WebNavigationTypeReload WebNavigationType = 3

type WebViewInsertAction int

const WebViewInsertActionDropped WebViewInsertAction = 2
const WebViewInsertActionPasted WebViewInsertAction = 1
const WebViewInsertActionTyped WebViewInsertAction = 0
