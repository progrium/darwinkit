// auto-generated code, do not modify
package webkit

type AudiovisualMediaTypes uint

const AudiovisualMediaTypeNone AudiovisualMediaTypes = 0
const AudiovisualMediaTypeAudio AudiovisualMediaTypes = 1
const AudiovisualMediaTypeVideo AudiovisualMediaTypes = 2
const AudiovisualMediaTypeAll AudiovisualMediaTypes = 18446744073709551615

type ContentMode int

const ContentModeRecommended ContentMode = 0
const ContentModeDesktop ContentMode = 2
const ContentModeMobile ContentMode = 1

type FullscreenState int

const FullscreenStateEnteringFullscreen FullscreenState = 1
const FullscreenStateExitingFullscreen FullscreenState = 3
const FullscreenStateInFullscreen FullscreenState = 2
const FullscreenStateNotInFullscreen FullscreenState = 0

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
const MediaPlaybackStatePaused MediaPlaybackState = 1
const MediaPlaybackStatePlaying MediaPlaybackState = 3
const MediaPlaybackStateSuspended MediaPlaybackState = 2

type NavigationActionPolicy int

const NavigationActionPolicyCancel NavigationActionPolicy = 0
const NavigationActionPolicyAllow NavigationActionPolicy = 1
const NavigationActionPolicyDownload NavigationActionPolicy = 2

type NavigationResponsePolicy int

const NavigationResponsePolicyCancel NavigationResponsePolicy = 0
const NavigationResponsePolicyAllow NavigationResponsePolicy = 1
const NavigationResponsePolicyDownload NavigationResponsePolicy = 2

type NavigationType int

const NavigationTypeLinkActivated NavigationType = 0
const NavigationTypeFormSubmitted NavigationType = 1
const NavigationTypeBackForward NavigationType = 2
const NavigationTypeReload NavigationType = 3
const NavigationTypeFormResubmitted NavigationType = 4
const NavigationTypeOther NavigationType = -1

type PermissionDecision int

const PermissionDecisionDeny PermissionDecision = 2
const PermissionDecisionGrant PermissionDecision = 1
const PermissionDecisionPrompt PermissionDecision = 0

type UserInterfaceDirectionPolicy int

const UserInterfaceDirectionPolicyContent UserInterfaceDirectionPolicy = 0
const UserInterfaceDirectionPolicySystem UserInterfaceDirectionPolicy = 1

type UserScriptInjectionTime int

const UserScriptInjectionTimeAtDocumentStart UserScriptInjectionTime = 0
const UserScriptInjectionTimeAtDocumentEnd UserScriptInjectionTime = 1
