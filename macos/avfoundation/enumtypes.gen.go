// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

// Constants that define eviction priorities for a storage management policy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadedassetevictionpriority?language=objc
type AssetDownloadedAssetEvictionPriority string

const (
	AssetDownloadedAssetEvictionPriorityDefault   AssetDownloadedAssetEvictionPriority = "default"
	AssetDownloadedAssetEvictionPriorityImportant AssetDownloadedAssetEvictionPriority = "important"
)

// Values that indicate the state of an export session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsessionstatus?language=objc
type AssetExportSessionStatus int

const (
	AssetExportSessionStatusCancelled AssetExportSessionStatus = 5
	AssetExportSessionStatusCompleted AssetExportSessionStatus = 3
	AssetExportSessionStatusExporting AssetExportSessionStatus = 2
	AssetExportSessionStatusFailed    AssetExportSessionStatus = 4
	AssetExportSessionStatusUnknown   AssetExportSessionStatus = 0
	AssetExportSessionStatusWaiting   AssetExportSessionStatus = 1
)

// Constants that define aperture modes to use when generating images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegeneratoraperturemode?language=objc
type AssetImageGeneratorApertureMode string

const (
	AssetImageGeneratorApertureModeCleanAperture      AssetImageGeneratorApertureMode = "CleanAperture"
	AssetImageGeneratorApertureModeEncodedPixels      AssetImageGeneratorApertureMode = "EncodedPixels"
	AssetImageGeneratorApertureModeProductionAperture AssetImageGeneratorApertureMode = "ProductionAperture"
)

// Constants that indicate the result of an image generation request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegeneratorresult?language=objc
type AssetImageGeneratorResult int

const (
	AssetImageGeneratorCancelled AssetImageGeneratorResult = 2
	AssetImageGeneratorFailed    AssetImageGeneratorResult = 1
	AssetImageGeneratorSucceeded AssetImageGeneratorResult = 0
)

// Values that represent the possible states of an asset reader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderstatus?language=objc
type AssetReaderStatus int

const (
	AssetReaderStatusCancelled AssetReaderStatus = 4
	AssetReaderStatusCompleted AssetReaderStatus = 2
	AssetReaderStatusFailed    AssetReaderStatus = 3
	AssetReaderStatusReading   AssetReaderStatus = 1
	AssetReaderStatusUnknown   AssetReaderStatus = 0
)

// Restrictions to use when resolving references to external media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreferencerestrictions?language=objc
type AssetReferenceRestrictions uint

const (
	AssetReferenceRestrictionDefaultPolicy                AssetReferenceRestrictions = 2
	AssetReferenceRestrictionForbidAll                    AssetReferenceRestrictions = 65535
	AssetReferenceRestrictionForbidCrossSiteReference     AssetReferenceRestrictions = 4
	AssetReferenceRestrictionForbidLocalReferenceToLocal  AssetReferenceRestrictions = 8
	AssetReferenceRestrictionForbidLocalReferenceToRemote AssetReferenceRestrictions = 2
	AssetReferenceRestrictionForbidNone                   AssetReferenceRestrictions = 0
	AssetReferenceRestrictionForbidRemoteReferenceToLocal AssetReferenceRestrictions = 1
)

// Constants that define the type of a segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmenttype?language=objc
type AssetSegmentType int

const (
	AssetSegmentTypeInitialization AssetSegmentType = 1
	AssetSegmentTypeSeparable      AssetSegmentType = 2
)

// A structure that indicates how to lay out and interleave media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputmediadatalocation?language=objc
type AssetWriterInputMediaDataLocation string

const (
	AssetWriterInputMediaDataLocationBeforeMainMediaDataNotInterleaved AssetWriterInputMediaDataLocation = "AVAssetWriterInputMediaDataLocationBeforeMainMediaDataNotInterleaved"
	AssetWriterInputMediaDataLocationInterleavedWithMainMediaData      AssetWriterInputMediaDataLocation = "AVAssetWriterInputMediaDataLocationInterleavedWithMainMediaData"
)

// Values that indicate the state of an asset writer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterstatus?language=objc
type AssetWriterStatus int

const (
	AssetWriterStatusCancelled AssetWriterStatus = 4
	AssetWriterStatusCompleted AssetWriterStatus = 2
	AssetWriterStatusFailed    AssetWriterStatus = 3
	AssetWriterStatusUnknown   AssetWriterStatus = 0
	AssetWriterStatusWriting   AssetWriterStatus = 1
)

// A structure that defines the spatialization formats that a player item supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaudiospatializationformats?language=objc
type AudioSpatializationFormats uint

const (
	AudioSpatializationFormatMonoAndStereo             AudioSpatializationFormats = 3
	AudioSpatializationFormatMonoStereoAndMultichannel AudioSpatializationFormats = 7
	AudioSpatializationFormatMultichannel              AudioSpatializationFormats = 4
	AudioSpatializationFormatNone                      AudioSpatializationFormats = 0
)

// An algorithm used to set the audio pitch as the rate changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaudiotimepitchalgorithm?language=objc
type AudioTimePitchAlgorithm string

const (
	AudioTimePitchAlgorithmSpectral   AudioTimePitchAlgorithm = "Spectral"
	AudioTimePitchAlgorithmTimeDomain AudioTimePitchAlgorithm = "TimeDomain"
	AudioTimePitchAlgorithmVarispeed  AudioTimePitchAlgorithm = "Varispeed"
)

// Constants that indicate the status of an app’s authorization to capture media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avauthorizationstatus?language=objc
type AuthorizationStatus int

const (
	AuthorizationStatusAuthorized    AuthorizationStatus = 3
	AuthorizationStatusDenied        AuthorizationStatus = 2
	AuthorizationStatusNotDetermined AuthorizationStatus = 0
	AuthorizationStatusRestricted    AuthorizationStatus = 1
)

// Animation options for a caption. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionanimation?language=objc
type CaptionAnimation int

const (
	CaptionAnimationCharacterReveal CaptionAnimation = 1
	CaptionAnimationNone            CaptionAnimation = 0
)

// Constants that indicate an adjustment type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionadjustmenttype?language=objc
type CaptionConversionAdjustmentType string

const (
	CaptionConversionAdjustmentTypeTimeRange CaptionConversionAdjustmentType = "AVCaptionConversionAdjustmentTypeTimeRange"
)

// Constants that indicate the status of a validator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionvalidatorstatus?language=objc
type CaptionConversionValidatorStatus int

const (
	CaptionConversionValidatorStatusCompleted  CaptionConversionValidatorStatus = 2
	CaptionConversionValidatorStatusStopped    CaptionConversionValidatorStatus = 3
	CaptionConversionValidatorStatusUnknown    CaptionConversionValidatorStatus = 0
	CaptionConversionValidatorStatusValidating CaptionConversionValidatorStatus = 1
)

// The type of a caption conversion warning. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionwarningtype?language=objc
type CaptionConversionWarningType string

const (
	CaptionConversionWarningTypeExcessMediaData CaptionConversionWarningType = "AVCaptionConversionWarningTypeExcessMediaData"
)

// Text decorations for caption text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptiondecoration?language=objc
type CaptionDecoration uint

const (
	CaptionDecorationLineThrough CaptionDecoration = 2
	CaptionDecorationNone        CaptionDecoration = 0
	CaptionDecorationOverline    CaptionDecoration = 4
	CaptionDecorationUnderline   CaptionDecoration = 1
)

// Font styles for caption text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionfontstyle?language=objc
type CaptionFontStyle int

const (
	CaptionFontStyleItalic  CaptionFontStyle = 2
	CaptionFontStyleNormal  CaptionFontStyle = 1
	CaptionFontStyleUnknown CaptionFontStyle = 0
)

// Font weights for a caption. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionfontweight?language=objc
type CaptionFontWeight int

const (
	CaptionFontWeightBold    CaptionFontWeight = 2
	CaptionFontWeightNormal  CaptionFontWeight = 1
	CaptionFontWeightUnknown CaptionFontWeight = 0
)

// Constants that indicate the alignment of lines in a region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregiondisplayalignment?language=objc
type CaptionRegionDisplayAlignment int

const (
	CaptionRegionDisplayAlignmentAfter  CaptionRegionDisplayAlignment = 2
	CaptionRegionDisplayAlignmentBefore CaptionRegionDisplayAlignment = 0
	CaptionRegionDisplayAlignmentCenter CaptionRegionDisplayAlignment = 1
)

// Constants that indicate the scrolling effects the system applies to a region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregionscroll?language=objc
type CaptionRegionScroll int

const (
	CaptionRegionScrollNone   CaptionRegionScroll = 0
	CaptionRegionScrollRollUp CaptionRegionScroll = 1
)

// Constants that indicate the writing mode for a region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionregionwritingmode?language=objc
type CaptionRegionWritingMode int

const (
	CaptionRegionWritingModeLeftToRightAndTopToBottom CaptionRegionWritingMode = 0
	CaptionRegionWritingModeTopToBottomAndRightToLeft CaptionRegionWritingMode = 2
)

// Constants that indicate ruby text alignments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionrubyalignment?language=objc
type CaptionRubyAlignment int

const (
	CaptionRubyAlignmentCenter                 CaptionRubyAlignment = 1
	CaptionRubyAlignmentDistributeSpaceAround  CaptionRubyAlignment = 3
	CaptionRubyAlignmentDistributeSpaceBetween CaptionRubyAlignment = 2
	CaptionRubyAlignmentStart                  CaptionRubyAlignment = 0
)

// Constants that indicate ruby text positions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionrubyposition?language=objc
type CaptionRubyPosition int

const (
	CaptionRubyPositionAfter  CaptionRubyPosition = 1
	CaptionRubyPositionBefore CaptionRubyPosition = 0
)

// A structure that defines dictionary keys to configure the caption converter and validator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionsettingskey?language=objc
type CaptionSettingsKey string

const (
	CaptionMediaSubTypeKey          CaptionSettingsKey = "AVCaptionMediaSubTypeKey"
	CaptionMediaTypeKey             CaptionSettingsKey = "AVCaptionMediaTypeKey"
	CaptionTimeCodeFrameDurationKey CaptionSettingsKey = "AVCaptionTimeCodeFrameDurationKey"
	CaptionUseDropFrameTimeCodeKey  CaptionSettingsKey = "AVCaptionUseDropFrameTimeCodeKey"
)

// Text alignment options for a caption. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptiontextalignment?language=objc
type CaptionTextAlignment int

const (
	CaptionTextAlignmentCenter CaptionTextAlignment = 2
	CaptionTextAlignmentEnd    CaptionTextAlignment = 1
	CaptionTextAlignmentLeft   CaptionTextAlignment = 3
	CaptionTextAlignmentRight  CaptionTextAlignment = 4
	CaptionTextAlignmentStart  CaptionTextAlignment = 0
)

// The caption’s supported rendering policy options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptiontextcombine?language=objc
type CaptionTextCombine int

const (
	CaptionTextCombineAll         CaptionTextCombine = -1
	CaptionTextCombineFourDigits  CaptionTextCombine = 4
	CaptionTextCombineNone        CaptionTextCombine = 0
	CaptionTextCombineOneDigit    CaptionTextCombine = 1
	CaptionTextCombineThreeDigits CaptionTextCombine = 3
	CaptionTextCombineTwoDigits   CaptionTextCombine = 2
)

// A structure that defines a units for caption formats. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionunitstype?language=objc
type CaptionUnitsType int

const (
	CaptionUnitsTypeCells       CaptionUnitsType = 1
	CaptionUnitsTypePercent     CaptionUnitsType = 2
	CaptionUnitsTypeUnspecified CaptionUnitsType = 0
)

// An enumeration of auto focus systems. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureautofocussystem?language=objc
type CaptureAutoFocusSystem int

const (
	CaptureAutoFocusSystemContrastDetection CaptureAutoFocusSystem = 1
	CaptureAutoFocusSystemNone              CaptureAutoFocusSystem = 0
	CaptureAutoFocusSystemPhaseDetection    CaptureAutoFocusSystem = 2
)

// Constants that indicate the current Center Stage control mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturecenterstagecontrolmode?language=objc
type CaptureCenterStageControlMode int

// An enumeration of color spaces a device can support. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturecolorspace?language=objc
type CaptureColorSpace int

const (
	CaptureColorSpace_P3_D65 CaptureColorSpace = 1
	CaptureColorSpace_sRGB   CaptureColorSpace = 0
)

// Constants that indicate the physical position of a capture device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceposition?language=objc
type CaptureDevicePosition int

const (
	CaptureDevicePositionBack        CaptureDevicePosition = 1
	CaptureDevicePositionFront       CaptureDevicePosition = 2
	CaptureDevicePositionUnspecified CaptureDevicePosition = 0
)

// Constants that indicate the transport control’s current mode of playback, if it has one. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevicetransportcontrolsplaybackmode?language=objc
type CaptureDeviceTransportControlsPlaybackMode int

const (
	CaptureDeviceTransportControlsNotPlayingMode CaptureDeviceTransportControlsPlaybackMode = 0
	CaptureDeviceTransportControlsPlayingMode    CaptureDeviceTransportControlsPlaybackMode = 1
)

// A constant that specifies speed of transport controls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevicetransportcontrolsspeed?language=objc
type CaptureDeviceTransportControlsSpeed float64

// A structure that defines the device types the framework supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevicetype?language=objc
type CaptureDeviceType string

const (
	CaptureDeviceTypeBuiltInMicrophone      CaptureDeviceType = "AVCaptureDeviceTypeBuiltInMicrophone"
	CaptureDeviceTypeBuiltInWideAngleCamera CaptureDeviceType = "AVCaptureDeviceTypeBuiltInWideAngleCamera"
	CaptureDeviceTypeExternalUnknown        CaptureDeviceType = "AVCaptureDeviceTypeExternalUnknown"
)

// Constants that specify the exposure mode of a capture device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureexposuremode?language=objc
type CaptureExposureMode int

const (
	CaptureExposureModeAutoExpose             CaptureExposureMode = 1
	CaptureExposureModeContinuousAutoExposure CaptureExposureMode = 2
	CaptureExposureModeCustom                 CaptureExposureMode = 3
	CaptureExposureModeLocked                 CaptureExposureMode = 0
)

// Constants that specify the flash modes of a capture device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureflashmode?language=objc
type CaptureFlashMode int

const (
	CaptureFlashModeAuto CaptureFlashMode = 2
	CaptureFlashModeOff  CaptureFlashMode = 0
	CaptureFlashModeOn   CaptureFlashMode = 1
)

// Constants to specify the focus mode of a capture device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefocusmode?language=objc
type CaptureFocusMode int

const (
	CaptureFocusModeAutoFocus           CaptureFocusMode = 1
	CaptureFocusModeContinuousAutoFocus CaptureFocusMode = 2
	CaptureFocusModeLocked              CaptureFocusMode = 0
)

// Constants that define the available microphone modes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemicrophonemode?language=objc
type CaptureMicrophoneMode int

const (
	CaptureMicrophoneModeStandard       CaptureMicrophoneMode = 0
	CaptureMicrophoneModeVoiceIsolation CaptureMicrophoneMode = 2
	CaptureMicrophoneModeWideSpectrum   CaptureMicrophoneMode = 1
)

// Constants that define reasons for why the system dropped a frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureoutputdatadroppedreason?language=objc
type CaptureOutputDataDroppedReason int

const (
	CaptureOutputDataDroppedReasonDiscontinuity CaptureOutputDataDroppedReason = 3
	CaptureOutputDataDroppedReasonLateData      CaptureOutputDataDroppedReason = 1
	CaptureOutputDataDroppedReasonNone          CaptureOutputDataDroppedReason = 0
	CaptureOutputDataDroppedReasonOutOfBuffers  CaptureOutputDataDroppedReason = 2
)

// A structure that defines the conditions in which to restrict camera switching. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureprimaryconstituentdevicerestrictedswitchingbehaviorconditions?language=objc
type CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions uint

const (
	CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionExposureModeChanged CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions = 4
	CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionFocusModeChanged    CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions = 2
	CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionNone                CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions = 0
	CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionVideoZoomChanged    CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions = 1
)

// Constants that control when to allow a virtual device to switch its active primary constituent device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureprimaryconstituentdeviceswitchingbehavior?language=objc
type CapturePrimaryConstituentDeviceSwitchingBehavior int

const (
	CapturePrimaryConstituentDeviceSwitchingBehaviorAuto        CapturePrimaryConstituentDeviceSwitchingBehavior = 1
	CapturePrimaryConstituentDeviceSwitchingBehaviorLocked      CapturePrimaryConstituentDeviceSwitchingBehavior = 3
	CapturePrimaryConstituentDeviceSwitchingBehaviorRestricted  CapturePrimaryConstituentDeviceSwitchingBehavior = 2
	CapturePrimaryConstituentDeviceSwitchingBehaviorUnsupported CapturePrimaryConstituentDeviceSwitchingBehavior = 0
)

// Presets that define standard configurations for a capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesessionpreset?language=objc
type CaptureSessionPreset string

const (
	CaptureSessionPreset1280x720       CaptureSessionPreset = "AVCaptureSessionPreset1280x720"
	CaptureSessionPreset1920x1080      CaptureSessionPreset = "AVCaptureSessionPreset1920x1080"
	CaptureSessionPreset320x240        CaptureSessionPreset = "AVCaptureSessionPreset320x240"
	CaptureSessionPreset352x288        CaptureSessionPreset = "AVCaptureSessionPreset352x288"
	CaptureSessionPreset3840x2160      CaptureSessionPreset = "AVCaptureSessionPreset3840x2160"
	CaptureSessionPreset640x480        CaptureSessionPreset = "AVCaptureSessionPreset640x480"
	CaptureSessionPreset960x540        CaptureSessionPreset = "AVCaptureSessionPreset960x540"
	CaptureSessionPresetHigh           CaptureSessionPreset = "AVCaptureSessionPresetHigh"
	CaptureSessionPresetLow            CaptureSessionPreset = "AVCaptureSessionPresetLow"
	CaptureSessionPresetMedium         CaptureSessionPreset = "AVCaptureSessionPresetMedium"
	CaptureSessionPresetPhoto          CaptureSessionPreset = "AVCaptureSessionPresetPhoto"
	CaptureSessionPresetiFrame1280x720 CaptureSessionPreset = "AVCaptureSessionPresetiFrame1280x720"
	CaptureSessionPresetiFrame960x540  CaptureSessionPreset = "AVCaptureSessionPresetiFrame960x540"
)

// Constants that describe the capture device configuration user interfaces. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesystemuserinterface?language=objc
type CaptureSystemUserInterface int

const (
	CaptureSystemUserInterfaceMicrophoneModes CaptureSystemUserInterface = 2
	CaptureSystemUserInterfaceVideoEffects    CaptureSystemUserInterface = 1
)

// Constants to specify the capture device’s torch mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturetorchmode?language=objc
type CaptureTorchMode int

const (
	CaptureTorchModeAuto CaptureTorchMode = 2
	CaptureTorchModeOff  CaptureTorchMode = 0
	CaptureTorchModeOn   CaptureTorchMode = 1
)

// Constants indicating video orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideoorientation?language=objc
type CaptureVideoOrientation int

const (
	CaptureVideoOrientationLandscapeLeft      CaptureVideoOrientation = 4
	CaptureVideoOrientationLandscapeRight     CaptureVideoOrientation = 3
	CaptureVideoOrientationPortrait           CaptureVideoOrientation = 1
	CaptureVideoOrientationPortraitUpsideDown CaptureVideoOrientation = 2
)

// Constants to specify the white balance mode of a capture device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturewhitebalancemode?language=objc
type CaptureWhiteBalanceMode int

const (
	CaptureWhiteBalanceModeAutoWhiteBalance           CaptureWhiteBalanceMode = 1
	CaptureWhiteBalanceModeContinuousAutoWhiteBalance CaptureWhiteBalanceMode = 2
	CaptureWhiteBalanceModeLocked                     CaptureWhiteBalanceMode = 0
)

// A value representing the status of a content authorization request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentauthorizationstatus?language=objc
type ContentAuthorizationStatus int

const (
	ContentAuthorizationBusy         ContentAuthorizationStatus = 4
	ContentAuthorizationCancelled    ContentAuthorizationStatus = 2
	ContentAuthorizationCompleted    ContentAuthorizationStatus = 1
	ContentAuthorizationNotAvailable ContentAuthorizationStatus = 5
	ContentAuthorizationNotPossible  ContentAuthorizationStatus = 6
	ContentAuthorizationTimedOut     ContentAuthorizationStatus = 3
	ContentAuthorizationUnknown      ContentAuthorizationStatus = 0
)

// The reason for asking the client to retry a content key request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequestretryreason?language=objc
type ContentKeyRequestRetryReason string

const (
	ContentKeyRequestRetryReasonReceivedObsoleteContentKey       ContentKeyRequestRetryReason = "ReceivedObsoleteKeyResponse"
	ContentKeyRequestRetryReasonReceivedResponseWithExpiredLease ContentKeyRequestRetryReason = "ReceivedKeyResponseWithExpiredLease"
	ContentKeyRequestRetryReasonTimedOut                         ContentKeyRequestRetryReason = "ReceivedKeyResponseAfterSPCTimedOut"
)

// The status for a content key request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequeststatus?language=objc
type ContentKeyRequestStatus int

const (
	ContentKeyRequestStatusCancelled          ContentKeyRequestStatus = 4
	ContentKeyRequestStatusFailed             ContentKeyRequestStatus = 5
	ContentKeyRequestStatusReceivedResponse   ContentKeyRequestStatus = 1
	ContentKeyRequestStatusRenewed            ContentKeyRequestStatus = 2
	ContentKeyRequestStatusRequestingResponse ContentKeyRequestStatus = 0
	ContentKeyRequestStatusRetried            ContentKeyRequestStatus = 3
)

// Options for specifying additional information for generating server playback context (SPC). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessionserverplaybackcontextoption?language=objc
type ContentKeySessionServerPlaybackContextOption string

const (
	ContentKeySessionServerPlaybackContextOptionProtocolVersions ContentKeySessionServerPlaybackContextOption = "ProtocolVersionsKey"
	ContentKeySessionServerPlaybackContextOptionServerChallenge  ContentKeySessionServerPlaybackContextOption = "ServerChallenge"
)

// A key-delivery method for a content key session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysystem?language=objc
type ContentKeySystem string

const (
	ContentKeySystemAuthorizationToken ContentKeySystem = "AuthorizationTokenSystem"
	ContentKeySystemClearKey           ContentKeySystem = "ClearKeySystem"
	ContentKeySystemFairPlayStreaming  ContentKeySystem = "FairPlayStreaming"
)

// Constants that identify playback suspension reasons. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoordinatedplaybacksuspensionreason?language=objc
type CoordinatedPlaybackSuspensionReason string

const (
	CoordinatedPlaybackSuspensionReasonAudioSessionInterrupted        CoordinatedPlaybackSuspensionReason = "AVCoordinatedPlaybackSuspensionReasonAudioSessionInterrupted"
	CoordinatedPlaybackSuspensionReasonCoordinatedPlaybackNotPossible CoordinatedPlaybackSuspensionReason = "AVCoordinatedPlaybackSuspensionReasonCoordinatedPlaybackNotPossible"
	CoordinatedPlaybackSuspensionReasonPlayingInterstitial            CoordinatedPlaybackSuspensionReason = "AVCoordinatedPlaybackSuspensionReasonPlayingInterstitial"
	CoordinatedPlaybackSuspensionReasonStallRecovery                  CoordinatedPlaybackSuspensionReason = "AVCoordinatedPlaybackSuspensionReasonStallRecovery"
	CoordinatedPlaybackSuspensionReasonUserActionRequired             CoordinatedPlaybackSuspensionReason = "AVCoordinatedPlaybackSuspensionReasonUserActionRequired"
	CoordinatedPlaybackSuspensionReasonUserIsChangingCurrentTime      CoordinatedPlaybackSuspensionReason = "AVCoordinatedPlaybackSuspensionReasonUserIsChangingCurrentTime"
)

// Constants that define rate change options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorratechangeoptions?language=objc
type DelegatingPlaybackCoordinatorRateChangeOptions uint

const (
	DelegatingPlaybackCoordinatorRateChangeOptionPlayImmediately DelegatingPlaybackCoordinatorRateChangeOptions = 1
)

// Constants that define seek options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorseekoptions?language=objc
type DelegatingPlaybackCoordinatorSeekOptions uint

const (
	DelegatingPlaybackCoordinatorSeekOptionResumeImmediately DelegatingPlaybackCoordinatorSeekOptions = 1
)

// Values indicating the general accuracy of a depth data map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdepthdataaccuracy?language=objc
type DepthDataAccuracy int

const (
	DepthDataAccuracyAbsolute DepthDataAccuracy = 1
	DepthDataAccuracyRelative DepthDataAccuracy = 0
)

// Values indicating the overall quality of a depth data map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdepthdataquality?language=objc
type DepthDataQuality int

const (
	DepthDataQualityHigh DepthDataQuality = 1
	DepthDataQualityLow  DepthDataQuality = 0
)

// An enumeration that defines the errors that framework operations can generate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/averror?language=objc
type Error int

const (
	ErrorAirPlayControllerRequiresInternet             Error = -11856
	ErrorAirPlayReceiverRequiresInternet               Error = -11857
	ErrorApplicationIsNotAuthorized                    Error = -11836
	ErrorApplicationIsNotAuthorizedToUseDevice         Error = -11852
	ErrorCompositionTrackSegmentsNotContiguous         Error = -11824
	ErrorContentIsNotAuthorized                        Error = -11835
	ErrorContentIsProtected                            Error = -11831
	ErrorContentIsUnavailable                          Error = -11863
	ErrorContentKeyRequestCancelled                    Error = -11879
	ErrorContentNotUpdated                             Error = -11866
	ErrorCreateContentKeyRequestFailed                 Error = -11860
	ErrorDecodeFailed                                  Error = -11821
	ErrorDecoderNotFound                               Error = -11833
	ErrorDecoderTemporarilyUnavailable                 Error = -11839
	ErrorDeviceAlreadyUsedByAnotherSession             Error = -11804
	ErrorDeviceInUseByAnotherApplication               Error = -11815
	ErrorDeviceLockedForConfigurationByAnotherProcess  Error = -11817
	ErrorDeviceNotConnected                            Error = -11814
	ErrorDeviceWasDisconnected                         Error = -11808
	ErrorDiskFull                                      Error = -11807
	ErrorDisplayWasDisabled                            Error = -11845
	ErrorEncoderNotFound                               Error = -11834
	ErrorEncoderTemporarilyUnavailable                 Error = -11840
	ErrorExportFailed                                  Error = -11820
	ErrorExternalPlaybackNotSupportedForAsset          Error = -11870
	ErrorFailedToLoadMediaData                         Error = -11849
	ErrorFailedToParse                                 Error = -11853
	ErrorFileAlreadyExists                             Error = -11823
	ErrorFileFailedToParse                             Error = -11829
	ErrorFileFormatNotRecognized                       Error = -11828
	ErrorFileTypeDoesNotSupportSampleReferences        Error = -11854
	ErrorFormatUnsupported                             Error = -11864
	ErrorIncompatibleAsset                             Error = -11848
	ErrorIncorrectlyConfigured                         Error = -11875
	ErrorInvalidCompositionTrackSegmentDuration        Error = -11825
	ErrorInvalidCompositionTrackSegmentSourceDuration  Error = -11827
	ErrorInvalidCompositionTrackSegmentSourceStartTime Error = -11826
	ErrorInvalidOutputURLPathExtension                 Error = -11843
	ErrorInvalidSourceMedia                            Error = -11822
	ErrorInvalidVideoComposition                       Error = -11841
	ErrorMalformedDepth                                Error = -11865
	ErrorMaximumDurationReached                        Error = -11810
	ErrorMaximumFileSizeReached                        Error = -11811
	ErrorMaximumNumberOfSamplesForFileFormatReached    Error = -11813
	ErrorMaximumStillImageCaptureRequestsExceeded      Error = -11830
	ErrorMediaChanged                                  Error = -11809
	ErrorMediaDiscontinuity                            Error = -11812
	ErrorNoCompatibleAlternatesForExternalDisplay      Error = -11868
	ErrorNoDataCaptured                                Error = -11805
	ErrorNoImageAtTime                                 Error = -11832
	ErrorNoLongerPlayable                              Error = -11867
	ErrorNoSourceTrack                                 Error = -11869
	ErrorOperationCancelled                            Error = -11878
	ErrorOperationNotAllowed                           Error = -11862
	ErrorOperationNotSupportedForAsset                 Error = -11838
	ErrorOperationNotSupportedForPreset                Error = -11871
	ErrorOutOfMemory                                   Error = -11801
	ErrorReferenceForbiddenByReferencePolicy           Error = -11842
	ErrorRosettaNotInstalled                           Error = -11877
	ErrorScreenCaptureFailed                           Error = -11844
	ErrorSegmentStartedWithNonSyncSample               Error = -11876
	ErrorServerIncorrectlyConfigured                   Error = -11850
	ErrorSessionConfigurationChanged                   Error = -11806
	ErrorSessionNotRunning                             Error = -11803
	ErrorTorchLevelUnavailable                         Error = -11846
	ErrorUndecodableMediaData                          Error = -11855
	ErrorUnknown                                       Error = -11800
	ErrorUnsupportedOutputSettings                     Error = -11861
	ErrorVideoCompositorFailed                         Error = -11858
)

// The uniform type identifiers for various file formats. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfiletype?language=objc
type FileType string

const (
	FileType3GPP            FileType = "public.3gpp"
	FileType3GPP2           FileType = "public.3gpp2"
	FileTypeAC3             FileType = "public.ac3-audio"
	FileTypeAIFC            FileType = "public.aifc-audio"
	FileTypeAIFF            FileType = "public.aiff-audio"
	FileTypeAMR             FileType = "org.3gpp.adaptive-multi-rate-audio"
	FileTypeAVCI            FileType = "public.avci"
	FileTypeAppleM4A        FileType = "com.apple.m4a-audio"
	FileTypeAppleM4V        FileType = "com.apple.m4v-video"
	FileTypeAppleiTT        FileType = "com.apple.itunes-timed-text"
	FileTypeCoreAudioFormat FileType = "com.apple.coreaudio-format"
	FileTypeDNG             FileType = "com.adobe.raw-image"
	FileTypeEnhancedAC3     FileType = "public.enhanced-ac3-audio"
	FileTypeHEIC            FileType = "public.heic"
	FileTypeHEIF            FileType = "public.heif"
	FileTypeJPEG            FileType = "public.jpeg"
	FileTypeMPEG4           FileType = "public.mpeg-4"
	FileTypeMPEGLayer3      FileType = "public.mp3"
	FileTypeQuickTimeMovie  FileType = "com.apple.quicktime-movie"
	FileTypeSCC             FileType = "com.scenarist.closed-caption"
	FileTypeSunAU           FileType = "public.au-audio"
	FileTypeTIFF            FileType = "public.tiff"
	FileTypeWAVE            FileType = "com.microsoft.waveform-audio"
)

// File type profiles for streaming formats. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfiletypeprofile?language=objc
type FileTypeProfile string

const (
	FileTypeProfileMPEG4AppleHLS      FileTypeProfile = "MPEG4AppleHLS"
	FileTypeProfileMPEG4CMAFCompliant FileTypeProfile = "MPEG4CMAFCompliant"
)

// Values that indicate the loaded status of a property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avkeyvaluestatus?language=objc
type KeyValueStatus int

const (
	KeyValueStatusCancelled KeyValueStatus = 4
	KeyValueStatusFailed    KeyValueStatus = 3
	KeyValueStatusLoaded    KeyValueStatus = 2
	KeyValueStatusLoading   KeyValueStatus = 1
	KeyValueStatusUnknown   KeyValueStatus = 0
)

// A structure that defines how a layer displays a player’s visual content within the layer’s bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avlayervideogravity?language=objc
type LayerVideoGravity string

const (
	LayerVideoGravityResize           LayerVideoGravity = "AVLayerVideoGravityResize"
	LayerVideoGravityResizeAspect     LayerVideoGravity = "AVLayerVideoGravityResizeAspect"
	LayerVideoGravityResizeAspectFill LayerVideoGravity = "AVLayerVideoGravityResizeAspectFill"
)

// A structure that defines media data characteristics. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediacharacteristic?language=objc
type MediaCharacteristic string

const (
	MediaCharacteristicAudible                                 MediaCharacteristic = "AVMediaCharacteristicAudible"
	MediaCharacteristicContainsAlphaChannel                    MediaCharacteristic = "public.contains-alpha-channel"
	MediaCharacteristicContainsHDRVideo                        MediaCharacteristic = "public.contains-hdr-video"
	MediaCharacteristicContainsOnlyForcedSubtitles             MediaCharacteristic = "public.subtitles.forced-only"
	MediaCharacteristicDescribesMusicAndSoundForAccessibility  MediaCharacteristic = "public.accessibility.describes-music-and-sound"
	MediaCharacteristicDescribesVideoForAccessibility          MediaCharacteristic = "public.accessibility.describes-video"
	MediaCharacteristicDubbedTranslation                       MediaCharacteristic = "public.translation.dubbed"
	MediaCharacteristicEasyToRead                              MediaCharacteristic = "public.easy-to-read"
	MediaCharacteristicFrameBased                              MediaCharacteristic = "AVMediaCharacteristicFrameBased"
	MediaCharacteristicIsAuxiliaryContent                      MediaCharacteristic = "public.auxiliary-content"
	MediaCharacteristicIsMainProgramContent                    MediaCharacteristic = "public.main-program-content"
	MediaCharacteristicIsOriginalContent                       MediaCharacteristic = "public.original-content"
	MediaCharacteristicLanguageTranslation                     MediaCharacteristic = "public.translation"
	MediaCharacteristicLegible                                 MediaCharacteristic = "AVMediaCharacteristicLegible"
	MediaCharacteristicTranscribesSpokenDialogForAccessibility MediaCharacteristic = "public.accessibility.transcribes-spoken-dialog"
	MediaCharacteristicUsesWideGamutColorSpace                 MediaCharacteristic = "public.uses-wide-gamut-color-space"
	MediaCharacteristicVisual                                  MediaCharacteristic = "AVMediaCharacteristicVisual"
	MediaCharacteristicVoiceOverTranslation                    MediaCharacteristic = "public.translation.voice-over"
)

// An identifier for various media types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediatype?language=objc
type MediaType string

const (
	MediaTypeAudio         MediaType = "soun"
	MediaTypeClosedCaption MediaType = "clcp"
	MediaTypeDepthData     MediaType = "dpth"
	MediaTypeMetadata      MediaType = "meta"
	MediaTypeMuxed         MediaType = "muxx"
	MediaTypeSubtitle      MediaType = "sbtl"
	MediaTypeText          MediaType = "text"
	MediaTypeTimecode      MediaType = "tmcd"
	MediaTypeVideo         MediaType = "vide"
)

// A structure that defines keys for extra metadata attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataextraattributekey?language=objc
type MetadataExtraAttributeKey string

const (
	MetadataExtraAttributeBaseURIKey  MetadataExtraAttributeKey = "baseURL"
	MetadataExtraAttributeInfoKey     MetadataExtraAttributeKey = "info"
	MetadataExtraAttributeValueURIKey MetadataExtraAttributeKey = "URL"
)

// A structure that defines metadata formats. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataformat?language=objc
type MetadataFormat string

const (
	MetadataFormatHLSMetadata       MetadataFormat = "com.apple.quicktime.HLS"
	MetadataFormatID3Metadata       MetadataFormat = "org.id3"
	MetadataFormatISOUserData       MetadataFormat = "org.mp4ra"
	MetadataFormatQuickTimeMetadata MetadataFormat = "com.apple.quicktime.mdta"
	MetadataFormatQuickTimeUserData MetadataFormat = "com.apple.quicktime.udta"
	MetadataFormatUnknown           MetadataFormat = "public.unknown"
	MetadataFormatiTunesMetadata    MetadataFormat = "com.apple.itunes"
)

// A structure that defines identifiers for metadata formats. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataidentifier?language=objc
type MetadataIdentifier string

const (
	MetadataCommonIdentifierAccessibilityDescription                           MetadataIdentifier = "common/accessibilityDescription"
	MetadataCommonIdentifierAlbumName                                          MetadataIdentifier = "common/albumName"
	MetadataCommonIdentifierArtist                                             MetadataIdentifier = "common/artist"
	MetadataCommonIdentifierArtwork                                            MetadataIdentifier = "common/artwork"
	MetadataCommonIdentifierAssetIdentifier                                    MetadataIdentifier = "common/identifier"
	MetadataCommonIdentifierAuthor                                             MetadataIdentifier = "common/author"
	MetadataCommonIdentifierContributor                                        MetadataIdentifier = "common/contributor"
	MetadataCommonIdentifierCopyrights                                         MetadataIdentifier = "common/copyrights"
	MetadataCommonIdentifierCreationDate                                       MetadataIdentifier = "common/creationDate"
	MetadataCommonIdentifierCreator                                            MetadataIdentifier = "common/creator"
	MetadataCommonIdentifierDescription                                        MetadataIdentifier = "common/description"
	MetadataCommonIdentifierFormat                                             MetadataIdentifier = "common/format"
	MetadataCommonIdentifierLanguage                                           MetadataIdentifier = "common/language"
	MetadataCommonIdentifierLastModifiedDate                                   MetadataIdentifier = "common/lastModifiedDate"
	MetadataCommonIdentifierLocation                                           MetadataIdentifier = "common/location"
	MetadataCommonIdentifierMake                                               MetadataIdentifier = "common/make"
	MetadataCommonIdentifierModel                                              MetadataIdentifier = "common/model"
	MetadataCommonIdentifierPublisher                                          MetadataIdentifier = "common/publisher"
	MetadataCommonIdentifierRelation                                           MetadataIdentifier = "common/relation"
	MetadataCommonIdentifierSoftware                                           MetadataIdentifier = "common/software"
	MetadataCommonIdentifierSource                                             MetadataIdentifier = "common/source"
	MetadataCommonIdentifierSubject                                            MetadataIdentifier = "common/subject"
	MetadataCommonIdentifierTitle                                              MetadataIdentifier = "common/title"
	MetadataCommonIdentifierType                                               MetadataIdentifier = "common/type"
	MetadataIdentifier3GPUserDataAlbumAndTrack                                 MetadataIdentifier = "uiso/albm"
	MetadataIdentifier3GPUserDataAuthor                                        MetadataIdentifier = "uiso/auth"
	MetadataIdentifier3GPUserDataCollection                                    MetadataIdentifier = "uiso/coll"
	MetadataIdentifier3GPUserDataCopyright                                     MetadataIdentifier = "uiso/cprt"
	MetadataIdentifier3GPUserDataDescription                                   MetadataIdentifier = "uiso/dscp"
	MetadataIdentifier3GPUserDataGenre                                         MetadataIdentifier = "uiso/gnre"
	MetadataIdentifier3GPUserDataKeywordList                                   MetadataIdentifier = "uiso/kywd"
	MetadataIdentifier3GPUserDataLocation                                      MetadataIdentifier = "uiso/loci"
	MetadataIdentifier3GPUserDataMediaClassification                           MetadataIdentifier = "uiso/clsf"
	MetadataIdentifier3GPUserDataMediaRating                                   MetadataIdentifier = "uiso/rtng"
	MetadataIdentifier3GPUserDataPerformer                                     MetadataIdentifier = "uiso/perf"
	MetadataIdentifier3GPUserDataRecordingYear                                 MetadataIdentifier = "uiso/yrrc"
	MetadataIdentifier3GPUserDataThumbnail                                     MetadataIdentifier = "uiso/thmb"
	MetadataIdentifier3GPUserDataTitle                                         MetadataIdentifier = "uiso/titl"
	MetadataIdentifier3GPUserDataUserRating                                    MetadataIdentifier = "uiso/urat"
	MetadataIdentifierID3MetadataAlbumSortOrder                                MetadataIdentifier = "id3/TSOA"
	MetadataIdentifierID3MetadataAlbumTitle                                    MetadataIdentifier = "id3/TALB"
	MetadataIdentifierID3MetadataAttachedPicture                               MetadataIdentifier = "id3/APIC"
	MetadataIdentifierID3MetadataAudioEncryption                               MetadataIdentifier = "id3/AENC"
	MetadataIdentifierID3MetadataAudioSeekPointIndex                           MetadataIdentifier = "id3/ASPI"
	MetadataIdentifierID3MetadataBand                                          MetadataIdentifier = "id3/TPE2"
	MetadataIdentifierID3MetadataBeatsPerMinute                                MetadataIdentifier = "id3/TBPM"
	MetadataIdentifierID3MetadataComments                                      MetadataIdentifier = "id3/COMM"
	MetadataIdentifierID3MetadataCommercial                                    MetadataIdentifier = "id3/COMR"
	MetadataIdentifierID3MetadataCommercialInformation                         MetadataIdentifier = "id3/WCOM"
	MetadataIdentifierID3MetadataCommerical                                    MetadataIdentifier = "id3/COMR"
	MetadataIdentifierID3MetadataComposer                                      MetadataIdentifier = "id3/TCOM"
	MetadataIdentifierID3MetadataConductor                                     MetadataIdentifier = "id3/TPE3"
	MetadataIdentifierID3MetadataContentGroupDescription                       MetadataIdentifier = "id3/TIT1"
	MetadataIdentifierID3MetadataContentType                                   MetadataIdentifier = "id3/TCON"
	MetadataIdentifierID3MetadataCopyright                                     MetadataIdentifier = "id3/TCOP"
	MetadataIdentifierID3MetadataCopyrightInformation                          MetadataIdentifier = "id3/WCOP"
	MetadataIdentifierID3MetadataDate                                          MetadataIdentifier = "id3/TDAT"
	MetadataIdentifierID3MetadataEncodedBy                                     MetadataIdentifier = "id3/TENC"
	MetadataIdentifierID3MetadataEncodedWith                                   MetadataIdentifier = "id3/TSSE"
	MetadataIdentifierID3MetadataEncodingTime                                  MetadataIdentifier = "id3/TDEN"
	MetadataIdentifierID3MetadataEncryption                                    MetadataIdentifier = "id3/ENCR"
	MetadataIdentifierID3MetadataEqualization                                  MetadataIdentifier = "id3/EQUA"
	MetadataIdentifierID3MetadataEqualization2                                 MetadataIdentifier = "id3/EQU2"
	MetadataIdentifierID3MetadataEventTimingCodes                              MetadataIdentifier = "id3/ETCO"
	MetadataIdentifierID3MetadataFileOwner                                     MetadataIdentifier = "id3/TOWN"
	MetadataIdentifierID3MetadataFileType                                      MetadataIdentifier = "id3/TFLT"
	MetadataIdentifierID3MetadataGeneralEncapsulatedObject                     MetadataIdentifier = "id3/GEOB"
	MetadataIdentifierID3MetadataGroupIdentifier                               MetadataIdentifier = "id3/GRID"
	MetadataIdentifierID3MetadataInitialKey                                    MetadataIdentifier = "id3/TKEY"
	MetadataIdentifierID3MetadataInternationalStandardRecordingCode            MetadataIdentifier = "id3/TSRC"
	MetadataIdentifierID3MetadataInternetRadioStationName                      MetadataIdentifier = "id3/TRSN"
	MetadataIdentifierID3MetadataInternetRadioStationOwner                     MetadataIdentifier = "id3/TRSO"
	MetadataIdentifierID3MetadataInvolvedPeopleList_v23                        MetadataIdentifier = "id3/IPLS"
	MetadataIdentifierID3MetadataInvolvedPeopleList_v24                        MetadataIdentifier = "id3/TIPL"
	MetadataIdentifierID3MetadataLanguage                                      MetadataIdentifier = "id3/TLAN"
	MetadataIdentifierID3MetadataLeadPerformer                                 MetadataIdentifier = "id3/TPE1"
	MetadataIdentifierID3MetadataLength                                        MetadataIdentifier = "id3/TLEN"
	MetadataIdentifierID3MetadataLink                                          MetadataIdentifier = "id3/LINK"
	MetadataIdentifierID3MetadataLyricist                                      MetadataIdentifier = "id3/TEXT"
	MetadataIdentifierID3MetadataMPEGLocationLookupTable                       MetadataIdentifier = "id3/MLLT"
	MetadataIdentifierID3MetadataMediaType                                     MetadataIdentifier = "id3/TMED"
	MetadataIdentifierID3MetadataModifiedBy                                    MetadataIdentifier = "id3/TPE4"
	MetadataIdentifierID3MetadataMood                                          MetadataIdentifier = "id3/TMOO"
	MetadataIdentifierID3MetadataMusicCDIdentifier                             MetadataIdentifier = "id3/MCDI"
	MetadataIdentifierID3MetadataMusicianCreditsList                           MetadataIdentifier = "id3/TMCL"
	MetadataIdentifierID3MetadataOfficialArtistWebpage                         MetadataIdentifier = "id3/WOAR"
	MetadataIdentifierID3MetadataOfficialAudioFileWebpage                      MetadataIdentifier = "id3/WOAF"
	MetadataIdentifierID3MetadataOfficialAudioSourceWebpage                    MetadataIdentifier = "id3/WOAS"
	MetadataIdentifierID3MetadataOfficialInternetRadioStationHomepage          MetadataIdentifier = "id3/WORS"
	MetadataIdentifierID3MetadataOfficialPublisherWebpage                      MetadataIdentifier = "id3/WPUB"
	MetadataIdentifierID3MetadataOriginalAlbumTitle                            MetadataIdentifier = "id3/TOAL"
	MetadataIdentifierID3MetadataOriginalArtist                                MetadataIdentifier = "id3/TOPE"
	MetadataIdentifierID3MetadataOriginalFilename                              MetadataIdentifier = "id3/TOFN"
	MetadataIdentifierID3MetadataOriginalLyricist                              MetadataIdentifier = "id3/TOLY"
	MetadataIdentifierID3MetadataOriginalReleaseTime                           MetadataIdentifier = "id3/TDOR"
	MetadataIdentifierID3MetadataOriginalReleaseYear                           MetadataIdentifier = "id3/TORY"
	MetadataIdentifierID3MetadataOwnership                                     MetadataIdentifier = "id3/OWNE"
	MetadataIdentifierID3MetadataPartOfASet                                    MetadataIdentifier = "id3/TPOS"
	MetadataIdentifierID3MetadataPayment                                       MetadataIdentifier = "id3/WPAY"
	MetadataIdentifierID3MetadataPerformerSortOrder                            MetadataIdentifier = "id3/TSOP"
	MetadataIdentifierID3MetadataPlayCounter                                   MetadataIdentifier = "id3/PCNT"
	MetadataIdentifierID3MetadataPlaylistDelay                                 MetadataIdentifier = "id3/TDLY"
	MetadataIdentifierID3MetadataPopularimeter                                 MetadataIdentifier = "id3/POPM"
	MetadataIdentifierID3MetadataPositionSynchronization                       MetadataIdentifier = "id3/POSS"
	MetadataIdentifierID3MetadataPrivate                                       MetadataIdentifier = "id3/PRIV"
	MetadataIdentifierID3MetadataProducedNotice                                MetadataIdentifier = "id3/TPRO"
	MetadataIdentifierID3MetadataPublisher                                     MetadataIdentifier = "id3/TPUB"
	MetadataIdentifierID3MetadataRecommendedBufferSize                         MetadataIdentifier = "id3/RBUF"
	MetadataIdentifierID3MetadataRecordingDates                                MetadataIdentifier = "id3/TRDA"
	MetadataIdentifierID3MetadataRecordingTime                                 MetadataIdentifier = "id3/TDRC"
	MetadataIdentifierID3MetadataRelativeVolumeAdjustment                      MetadataIdentifier = "id3/RVAD"
	MetadataIdentifierID3MetadataRelativeVolumeAdjustment2                     MetadataIdentifier = "id3/RVA2"
	MetadataIdentifierID3MetadataReleaseTime                                   MetadataIdentifier = "id3/TDRL"
	MetadataIdentifierID3MetadataReverb                                        MetadataIdentifier = "id3/RVRB"
	MetadataIdentifierID3MetadataSeek                                          MetadataIdentifier = "id3/SEEK"
	MetadataIdentifierID3MetadataSetSubtitle                                   MetadataIdentifier = "id3/TSST"
	MetadataIdentifierID3MetadataSignature                                     MetadataIdentifier = "id3/SIGN"
	MetadataIdentifierID3MetadataSize                                          MetadataIdentifier = "id3/TSIZ"
	MetadataIdentifierID3MetadataSubTitle                                      MetadataIdentifier = "id3/TIT3"
	MetadataIdentifierID3MetadataSynchronizedLyric                             MetadataIdentifier = "id3/SYLT"
	MetadataIdentifierID3MetadataSynchronizedTempoCodes                        MetadataIdentifier = "id3/SYTC"
	MetadataIdentifierID3MetadataTaggingTime                                   MetadataIdentifier = "id3/TDTG"
	MetadataIdentifierID3MetadataTermsOfUse                                    MetadataIdentifier = "id3/USER"
	MetadataIdentifierID3MetadataTime                                          MetadataIdentifier = "id3/TIME"
	MetadataIdentifierID3MetadataTitleDescription                              MetadataIdentifier = "id3/TIT2"
	MetadataIdentifierID3MetadataTitleSortOrder                                MetadataIdentifier = "id3/TSOT"
	MetadataIdentifierID3MetadataTrackNumber                                   MetadataIdentifier = "id3/TRCK"
	MetadataIdentifierID3MetadataUniqueFileIdentifier                          MetadataIdentifier = "id3/UFID"
	MetadataIdentifierID3MetadataUnsynchronizedLyric                           MetadataIdentifier = "id3/USLT"
	MetadataIdentifierID3MetadataUserText                                      MetadataIdentifier = "id3/TXXX"
	MetadataIdentifierID3MetadataUserURL                                       MetadataIdentifier = "id3/WXXX"
	MetadataIdentifierID3MetadataYear                                          MetadataIdentifier = "id3/TYER"
	MetadataIdentifierISOUserDataAccessibilityDescription                      MetadataIdentifier = "uiso/ades"
	MetadataIdentifierISOUserDataCopyright                                     MetadataIdentifier = "uiso/cprt"
	MetadataIdentifierISOUserDataDate                                          MetadataIdentifier = "uiso/date"
	MetadataIdentifierISOUserDataTaggedCharacteristic                          MetadataIdentifier = "uiso/tagc"
	MetadataIdentifierIcyMetadataStreamTitle                                   MetadataIdentifier = "icy/StreamTitle"
	MetadataIdentifierIcyMetadataStreamURL                                     MetadataIdentifier = "icy/StreamUrl"
	MetadataIdentifierQuickTimeMetadataAccessibilityDescription                MetadataIdentifier = "mdta/com.apple.quicktime.accessibility.description"
	MetadataIdentifierQuickTimeMetadataAlbum                                   MetadataIdentifier = "mdta/com.apple.quicktime.album"
	MetadataIdentifierQuickTimeMetadataArranger                                MetadataIdentifier = "mdta/com.apple.quicktime.arranger"
	MetadataIdentifierQuickTimeMetadataArtist                                  MetadataIdentifier = "mdta/com.apple.quicktime.artist"
	MetadataIdentifierQuickTimeMetadataArtwork                                 MetadataIdentifier = "mdta/com.apple.quicktime.artwork"
	MetadataIdentifierQuickTimeMetadataAuthor                                  MetadataIdentifier = "mdta/com.apple.quicktime.author"
	MetadataIdentifierQuickTimeMetadataAutoLivePhoto                           MetadataIdentifier = "mdta/com.apple.quicktime.live-photo.auto"
	MetadataIdentifierQuickTimeMetadataCameraFrameReadoutTime                  MetadataIdentifier = "mdta/com.apple.quicktime.camera.framereadouttimeinmicroseconds"
	MetadataIdentifierQuickTimeMetadataCameraIdentifier                        MetadataIdentifier = "mdta/com.apple.quicktime.camera.identifier"
	MetadataIdentifierQuickTimeMetadataCollectionUser                          MetadataIdentifier = "mdta/com.apple.quicktime.collection.user"
	MetadataIdentifierQuickTimeMetadataComment                                 MetadataIdentifier = "mdta/com.apple.quicktime.comment"
	MetadataIdentifierQuickTimeMetadataComposer                                MetadataIdentifier = "mdta/com.apple.quicktime.composer"
	MetadataIdentifierQuickTimeMetadataContentIdentifier                       MetadataIdentifier = "mdta/com.apple.quicktime.content.identifier"
	MetadataIdentifierQuickTimeMetadataCopyright                               MetadataIdentifier = "mdta/com.apple.quicktime.copyright"
	MetadataIdentifierQuickTimeMetadataCreationDate                            MetadataIdentifier = "mdta/com.apple.quicktime.creationdate"
	MetadataIdentifierQuickTimeMetadataCredits                                 MetadataIdentifier = "mdta/com.apple.quicktime.credits"
	MetadataIdentifierQuickTimeMetadataDescription                             MetadataIdentifier = "mdta/com.apple.quicktime.description"
	MetadataIdentifierQuickTimeMetadataDetectedCatBody                         MetadataIdentifier = "mdta/com.apple.quicktime.detected-cat-body"
	MetadataIdentifierQuickTimeMetadataDetectedDogBody                         MetadataIdentifier = "mdta/com.apple.quicktime.detected-dog-body"
	MetadataIdentifierQuickTimeMetadataDetectedFace                            MetadataIdentifier = "mdta/com.apple.quicktime.detected-face"
	MetadataIdentifierQuickTimeMetadataDetectedHumanBody                       MetadataIdentifier = "mdta/com.apple.quicktime.detected-human-body"
	MetadataIdentifierQuickTimeMetadataDetectedSalientObject                   MetadataIdentifier = "mdta/com.apple.quicktime.detected-salient-object"
	MetadataIdentifierQuickTimeMetadataDirectionFacing                         MetadataIdentifier = "mdta/com.apple.quicktime.direction.facing"
	MetadataIdentifierQuickTimeMetadataDirectionMotion                         MetadataIdentifier = "mdta/com.apple.quicktime.direction.motion"
	MetadataIdentifierQuickTimeMetadataDirector                                MetadataIdentifier = "mdta/com.apple.quicktime.director"
	MetadataIdentifierQuickTimeMetadataDisplayName                             MetadataIdentifier = "mdta/com.apple.quicktime.displayname"
	MetadataIdentifierQuickTimeMetadataEncodedBy                               MetadataIdentifier = "mdta/com.apple.quicktime.encodedby"
	MetadataIdentifierQuickTimeMetadataGenre                                   MetadataIdentifier = "mdta/com.apple.quicktime.genre"
	MetadataIdentifierQuickTimeMetadataInformation                             MetadataIdentifier = "mdta/com.apple.quicktime.information"
	MetadataIdentifierQuickTimeMetadataIsMontage                               MetadataIdentifier = "mdta/com.apple.quicktime.is-montage"
	MetadataIdentifierQuickTimeMetadataKeywords                                MetadataIdentifier = "mdta/com.apple.quicktime.keywords"
	MetadataIdentifierQuickTimeMetadataLivePhotoVitalityScore                  MetadataIdentifier = "mdta/com.apple.quicktime.live-photo.vitality-score"
	MetadataIdentifierQuickTimeMetadataLivePhotoVitalityScoringVersion         MetadataIdentifier = "mdta/com.apple.quicktime.live-photo.vitality-scoring-version"
	MetadataIdentifierQuickTimeMetadataLocationBody                            MetadataIdentifier = "mdta/com.apple.quicktime.location.body"
	MetadataIdentifierQuickTimeMetadataLocationDate                            MetadataIdentifier = "mdta/com.apple.quicktime.location.date"
	MetadataIdentifierQuickTimeMetadataLocationHorizontalAccuracyInMeters      MetadataIdentifier = "mdta/com.apple.quicktime.location.accuracy.horizontal"
	MetadataIdentifierQuickTimeMetadataLocationISO6709                         MetadataIdentifier = "mdta/com.apple.quicktime.location.ISO6709"
	MetadataIdentifierQuickTimeMetadataLocationName                            MetadataIdentifier = "mdta/com.apple.quicktime.location.name"
	MetadataIdentifierQuickTimeMetadataLocationNote                            MetadataIdentifier = "mdta/com.apple.quicktime.location.note"
	MetadataIdentifierQuickTimeMetadataLocationRole                            MetadataIdentifier = "mdta/com.apple.quicktime.location.role"
	MetadataIdentifierQuickTimeMetadataMake                                    MetadataIdentifier = "mdta/com.apple.quicktime.make"
	MetadataIdentifierQuickTimeMetadataModel                                   MetadataIdentifier = "mdta/com.apple.quicktime.model"
	MetadataIdentifierQuickTimeMetadataOriginalArtist                          MetadataIdentifier = "mdta/com.apple.quicktime.originalartist"
	MetadataIdentifierQuickTimeMetadataPerformer                               MetadataIdentifier = "mdta/com.apple.quicktime.performer"
	MetadataIdentifierQuickTimeMetadataPhonogramRights                         MetadataIdentifier = "mdta/com.apple.quicktime.phonogramrights"
	MetadataIdentifierQuickTimeMetadataPreferredAffineTransform                MetadataIdentifier = "mdta/com.apple.quicktime.preferred-affine-transform"
	MetadataIdentifierQuickTimeMetadataProducer                                MetadataIdentifier = "mdta/com.apple.quicktime.producer"
	MetadataIdentifierQuickTimeMetadataPublisher                               MetadataIdentifier = "mdta/com.apple.quicktime.publisher"
	MetadataIdentifierQuickTimeMetadataRatingUser                              MetadataIdentifier = "mdta/com.apple.quicktime.rating.user"
	MetadataIdentifierQuickTimeMetadataSoftware                                MetadataIdentifier = "mdta/com.apple.quicktime.software"
	MetadataIdentifierQuickTimeMetadataSpatialOverCaptureQualityScore          MetadataIdentifier = "mdta/com.apple.quicktime.spatial-overcapture.quality-score"
	MetadataIdentifierQuickTimeMetadataSpatialOverCaptureQualityScoringVersion MetadataIdentifier = "mdta/com.apple.quicktime.spatial-overcapture.quality-scoring-version"
	MetadataIdentifierQuickTimeMetadataTitle                                   MetadataIdentifier = "mdta/com.apple.quicktime.title"
	MetadataIdentifierQuickTimeMetadataVideoOrientation                        MetadataIdentifier = "mdta/com.apple.quicktime.video-orientation"
	MetadataIdentifierQuickTimeMetadataYear                                    MetadataIdentifier = "mdta/com.apple.quicktime.year"
	MetadataIdentifierQuickTimeMetadataiXML                                    MetadataIdentifier = "mdta/info.ixml.xml"
	MetadataIdentifierQuickTimeUserDataAccessibilityDescription                MetadataIdentifier = "udta/%A9ade"
	MetadataIdentifierQuickTimeUserDataAlbum                                   MetadataIdentifier = "udta/%A9alb"
	MetadataIdentifierQuickTimeUserDataArranger                                MetadataIdentifier = "udta/%A9arg"
	MetadataIdentifierQuickTimeUserDataArtist                                  MetadataIdentifier = "udta/%A9ART"
	MetadataIdentifierQuickTimeUserDataAuthor                                  MetadataIdentifier = "udta/%A9aut"
	MetadataIdentifierQuickTimeUserDataChapter                                 MetadataIdentifier = "udta/%A9chp"
	MetadataIdentifierQuickTimeUserDataComment                                 MetadataIdentifier = "udta/%A9cmt"
	MetadataIdentifierQuickTimeUserDataComposer                                MetadataIdentifier = "udta/%A9com"
	MetadataIdentifierQuickTimeUserDataCopyright                               MetadataIdentifier = "udta/%A9cpy"
	MetadataIdentifierQuickTimeUserDataCreationDate                            MetadataIdentifier = "udta/%A9day"
	MetadataIdentifierQuickTimeUserDataCredits                                 MetadataIdentifier = "udta/%A9src"
	MetadataIdentifierQuickTimeUserDataDescription                             MetadataIdentifier = "udta/%A9des"
	MetadataIdentifierQuickTimeUserDataDirector                                MetadataIdentifier = "udta/%A9dir"
	MetadataIdentifierQuickTimeUserDataDisclaimer                              MetadataIdentifier = "udta/%A9dis"
	MetadataIdentifierQuickTimeUserDataEncodedBy                               MetadataIdentifier = "udta/%A9enc"
	MetadataIdentifierQuickTimeUserDataFullName                                MetadataIdentifier = "udta/%A9nam"
	MetadataIdentifierQuickTimeUserDataGenre                                   MetadataIdentifier = "udta/%A9gen"
	MetadataIdentifierQuickTimeUserDataHostComputer                            MetadataIdentifier = "udta/%A9hst"
	MetadataIdentifierQuickTimeUserDataInformation                             MetadataIdentifier = "udta/%A9inf"
	MetadataIdentifierQuickTimeUserDataKeywords                                MetadataIdentifier = "udta/%A9key"
	MetadataIdentifierQuickTimeUserDataLocationISO6709                         MetadataIdentifier = "udta/%A9xyz"
	MetadataIdentifierQuickTimeUserDataMake                                    MetadataIdentifier = "udta/%A9mak"
	MetadataIdentifierQuickTimeUserDataModel                                   MetadataIdentifier = "udta/%A9mod"
	MetadataIdentifierQuickTimeUserDataOriginalArtist                          MetadataIdentifier = "udta/%A9ope"
	MetadataIdentifierQuickTimeUserDataOriginalFormat                          MetadataIdentifier = "udta/%A9fmt"
	MetadataIdentifierQuickTimeUserDataOriginalSource                          MetadataIdentifier = "udta/%A9src"
	MetadataIdentifierQuickTimeUserDataPerformers                              MetadataIdentifier = "udta/%A9prf"
	MetadataIdentifierQuickTimeUserDataPhonogramRights                         MetadataIdentifier = "udta/%A9phg"
	MetadataIdentifierQuickTimeUserDataProducer                                MetadataIdentifier = "udta/%A9prd"
	MetadataIdentifierQuickTimeUserDataProduct                                 MetadataIdentifier = "udta/%A9PRD"
	MetadataIdentifierQuickTimeUserDataPublisher                               MetadataIdentifier = "udta/%A9pub"
	MetadataIdentifierQuickTimeUserDataSoftware                                MetadataIdentifier = "udta/%A9swr"
	MetadataIdentifierQuickTimeUserDataSpecialPlaybackRequirements             MetadataIdentifier = "udta/%A9req"
	MetadataIdentifierQuickTimeUserDataTaggedCharacteristic                    MetadataIdentifier = "udta/tagc"
	MetadataIdentifierQuickTimeUserDataTrack                                   MetadataIdentifier = "udta/%A9trk"
	MetadataIdentifierQuickTimeUserDataTrackName                               MetadataIdentifier = "udta/tnam"
	MetadataIdentifierQuickTimeUserDataURLLink                                 MetadataIdentifier = "udta/%A9url"
	MetadataIdentifierQuickTimeUserDataWarning                                 MetadataIdentifier = "udta/%A9wrn"
	MetadataIdentifierQuickTimeUserDataWriter                                  MetadataIdentifier = "udta/%A9wrt"
	MetadataIdentifieriTunesMetadataAccountKind                                MetadataIdentifier = "itsk/akID"
	MetadataIdentifieriTunesMetadataAcknowledgement                            MetadataIdentifier = "itsk/%A9cak"
	MetadataIdentifieriTunesMetadataAlbum                                      MetadataIdentifier = "itsk/%A9alb"
	MetadataIdentifieriTunesMetadataAlbumArtist                                MetadataIdentifier = "itsk/aART"
	MetadataIdentifieriTunesMetadataAppleID                                    MetadataIdentifier = "itsk/apID"
	MetadataIdentifieriTunesMetadataArranger                                   MetadataIdentifier = "itsk/%A9arg"
	MetadataIdentifieriTunesMetadataArtDirector                                MetadataIdentifier = "itsk/%A9ard"
	MetadataIdentifieriTunesMetadataArtist                                     MetadataIdentifier = "itsk/%A9ART"
	MetadataIdentifieriTunesMetadataArtistID                                   MetadataIdentifier = "itsk/atID"
	MetadataIdentifieriTunesMetadataAuthor                                     MetadataIdentifier = "itsk/%A9aut"
	MetadataIdentifieriTunesMetadataBeatsPerMin                                MetadataIdentifier = "itsk/tmpo"
	MetadataIdentifieriTunesMetadataComposer                                   MetadataIdentifier = "itsk/%A9wrt"
	MetadataIdentifieriTunesMetadataConductor                                  MetadataIdentifier = "itsk/%A9con"
	MetadataIdentifieriTunesMetadataContentRating                              MetadataIdentifier = "itsk/rtng"
	MetadataIdentifieriTunesMetadataCopyright                                  MetadataIdentifier = "itsk/cprt"
	MetadataIdentifieriTunesMetadataCoverArt                                   MetadataIdentifier = "itsk/covr"
	MetadataIdentifieriTunesMetadataCredits                                    MetadataIdentifier = "itsk/%A9src"
	MetadataIdentifieriTunesMetadataDescription                                MetadataIdentifier = "itsk/%A9des"
	MetadataIdentifieriTunesMetadataDirector                                   MetadataIdentifier = "itsk/%A9dir"
	MetadataIdentifieriTunesMetadataDiscCompilation                            MetadataIdentifier = "itsk/cpil"
	MetadataIdentifieriTunesMetadataDiscNumber                                 MetadataIdentifier = "itsk/disk"
	MetadataIdentifieriTunesMetadataEQ                                         MetadataIdentifier = "itsk/%A9equ"
	MetadataIdentifieriTunesMetadataEncodedBy                                  MetadataIdentifier = "itsk/%A9enc"
	MetadataIdentifieriTunesMetadataEncodingTool                               MetadataIdentifier = "itsk/%A9too"
	MetadataIdentifieriTunesMetadataExecProducer                               MetadataIdentifier = "itsk/%A9xpd"
	MetadataIdentifieriTunesMetadataGenreID                                    MetadataIdentifier = "itsk/geID"
	MetadataIdentifieriTunesMetadataGrouping                                   MetadataIdentifier = "itsk/grup"
	MetadataIdentifieriTunesMetadataLinerNotes                                 MetadataIdentifier = "itsk/%A9lnt"
	MetadataIdentifieriTunesMetadataLyrics                                     MetadataIdentifier = "itsk/%A9lyr"
	MetadataIdentifieriTunesMetadataOnlineExtras                               MetadataIdentifier = "itsk/%A9url"
	MetadataIdentifieriTunesMetadataOriginalArtist                             MetadataIdentifier = "itsk/%A9ope"
	MetadataIdentifieriTunesMetadataPerformer                                  MetadataIdentifier = "itsk/%A9prf"
	MetadataIdentifieriTunesMetadataPhonogramRights                            MetadataIdentifier = "itsk/%A9phg"
	MetadataIdentifieriTunesMetadataPlaylistID                                 MetadataIdentifier = "itsk/plID"
	MetadataIdentifieriTunesMetadataPredefinedGenre                            MetadataIdentifier = "itsk/gnre"
	MetadataIdentifieriTunesMetadataProducer                                   MetadataIdentifier = "itsk/%A9prd"
	MetadataIdentifieriTunesMetadataPublisher                                  MetadataIdentifier = "itsk/%A9pub"
	MetadataIdentifieriTunesMetadataRecordCompany                              MetadataIdentifier = "itsk/%A9mak"
	MetadataIdentifieriTunesMetadataReleaseDate                                MetadataIdentifier = "itsk/%A9day"
	MetadataIdentifieriTunesMetadataSoloist                                    MetadataIdentifier = "itsk/%A9sol"
	MetadataIdentifieriTunesMetadataSongID                                     MetadataIdentifier = "itsk/cnID"
	MetadataIdentifieriTunesMetadataSongName                                   MetadataIdentifier = "itsk/%A9nam"
	MetadataIdentifieriTunesMetadataSoundEngineer                              MetadataIdentifier = "itsk/%A9sne"
	MetadataIdentifieriTunesMetadataThanks                                     MetadataIdentifier = "itsk/%A9thx"
	MetadataIdentifieriTunesMetadataTrackNumber                                MetadataIdentifier = "itsk/trkn"
	MetadataIdentifieriTunesMetadataTrackSubTitle                              MetadataIdentifier = "itsk/%A9st3"
	MetadataIdentifieriTunesMetadataUserComment                                MetadataIdentifier = "itsk/%A9cmt"
	MetadataIdentifieriTunesMetadataUserGenre                                  MetadataIdentifier = "itsk/%A9gen"
)

// A structure that defines a metadata key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatakey?language=objc
type MetadataKey string

const (
	Metadata3GPUserDataKeyAlbumAndTrack                        MetadataKey = "albm"
	Metadata3GPUserDataKeyAuthor                               MetadataKey = "auth"
	Metadata3GPUserDataKeyCollection                           MetadataKey = "coll"
	Metadata3GPUserDataKeyCopyright                            MetadataKey = "cprt"
	Metadata3GPUserDataKeyDescription                          MetadataKey = "dscp"
	Metadata3GPUserDataKeyGenre                                MetadataKey = "gnre"
	Metadata3GPUserDataKeyKeywordList                          MetadataKey = "kywd"
	Metadata3GPUserDataKeyLocation                             MetadataKey = "loci"
	Metadata3GPUserDataKeyMediaClassification                  MetadataKey = "clsf"
	Metadata3GPUserDataKeyMediaRating                          MetadataKey = "rtng"
	Metadata3GPUserDataKeyPerformer                            MetadataKey = "perf"
	Metadata3GPUserDataKeyRecordingYear                        MetadataKey = "yrrc"
	Metadata3GPUserDataKeyThumbnail                            MetadataKey = "thmb"
	Metadata3GPUserDataKeyTitle                                MetadataKey = "titl"
	Metadata3GPUserDataKeyUserRating                           MetadataKey = "urat"
	MetadataCommonKeyAccessibilityDescription                  MetadataKey = "accessibilityDescription"
	MetadataCommonKeyAlbumName                                 MetadataKey = "albumName"
	MetadataCommonKeyArtist                                    MetadataKey = "artist"
	MetadataCommonKeyArtwork                                   MetadataKey = "artwork"
	MetadataCommonKeyAuthor                                    MetadataKey = "author"
	MetadataCommonKeyContributor                               MetadataKey = "contributor"
	MetadataCommonKeyCopyrights                                MetadataKey = "copyrights"
	MetadataCommonKeyCreationDate                              MetadataKey = "creationDate"
	MetadataCommonKeyCreator                                   MetadataKey = "creator"
	MetadataCommonKeyDescription                               MetadataKey = "description"
	MetadataCommonKeyFormat                                    MetadataKey = "format"
	MetadataCommonKeyIdentifier                                MetadataKey = "identifier"
	MetadataCommonKeyLanguage                                  MetadataKey = "language"
	MetadataCommonKeyLastModifiedDate                          MetadataKey = "lastModifiedDate"
	MetadataCommonKeyLocation                                  MetadataKey = "location"
	MetadataCommonKeyMake                                      MetadataKey = "make"
	MetadataCommonKeyModel                                     MetadataKey = "model"
	MetadataCommonKeyPublisher                                 MetadataKey = "publisher"
	MetadataCommonKeyRelation                                  MetadataKey = "relation"
	MetadataCommonKeySoftware                                  MetadataKey = "software"
	MetadataCommonKeySource                                    MetadataKey = "source"
	MetadataCommonKeySubject                                   MetadataKey = "subject"
	MetadataCommonKeyTitle                                     MetadataKey = "title"
	MetadataCommonKeyType                                      MetadataKey = "type"
	MetadataID3MetadataKeyAlbumSortOrder                       MetadataKey = "TSOA"
	MetadataID3MetadataKeyAlbumTitle                           MetadataKey = "TALB"
	MetadataID3MetadataKeyAttachedPicture                      MetadataKey = "APIC"
	MetadataID3MetadataKeyAudioEncryption                      MetadataKey = "AENC"
	MetadataID3MetadataKeyAudioSeekPointIndex                  MetadataKey = "ASPI"
	MetadataID3MetadataKeyBand                                 MetadataKey = "TPE2"
	MetadataID3MetadataKeyBeatsPerMinute                       MetadataKey = "TBPM"
	MetadataID3MetadataKeyComments                             MetadataKey = "COMM"
	MetadataID3MetadataKeyCommercial                           MetadataKey = "COMR"
	MetadataID3MetadataKeyCommercialInformation                MetadataKey = "WCOM"
	MetadataID3MetadataKeyCommerical                           MetadataKey = "COMR"
	MetadataID3MetadataKeyComposer                             MetadataKey = "TCOM"
	MetadataID3MetadataKeyConductor                            MetadataKey = "TPE3"
	MetadataID3MetadataKeyContentGroupDescription              MetadataKey = "TIT1"
	MetadataID3MetadataKeyContentType                          MetadataKey = "TCON"
	MetadataID3MetadataKeyCopyright                            MetadataKey = "TCOP"
	MetadataID3MetadataKeyCopyrightInformation                 MetadataKey = "WCOP"
	MetadataID3MetadataKeyDate                                 MetadataKey = "TDAT"
	MetadataID3MetadataKeyEncodedBy                            MetadataKey = "TENC"
	MetadataID3MetadataKeyEncodedWith                          MetadataKey = "TSSE"
	MetadataID3MetadataKeyEncodingTime                         MetadataKey = "TDEN"
	MetadataID3MetadataKeyEncryption                           MetadataKey = "ENCR"
	MetadataID3MetadataKeyEqualization                         MetadataKey = "EQUA"
	MetadataID3MetadataKeyEqualization2                        MetadataKey = "EQU2"
	MetadataID3MetadataKeyEventTimingCodes                     MetadataKey = "ETCO"
	MetadataID3MetadataKeyFileOwner                            MetadataKey = "TOWN"
	MetadataID3MetadataKeyFileType                             MetadataKey = "TFLT"
	MetadataID3MetadataKeyGeneralEncapsulatedObject            MetadataKey = "GEOB"
	MetadataID3MetadataKeyGroupIdentifier                      MetadataKey = "GRID"
	MetadataID3MetadataKeyInitialKey                           MetadataKey = "TKEY"
	MetadataID3MetadataKeyInternationalStandardRecordingCode   MetadataKey = "TSRC"
	MetadataID3MetadataKeyInternetRadioStationName             MetadataKey = "TRSN"
	MetadataID3MetadataKeyInternetRadioStationOwner            MetadataKey = "TRSO"
	MetadataID3MetadataKeyInvolvedPeopleList_v23               MetadataKey = "IPLS"
	MetadataID3MetadataKeyInvolvedPeopleList_v24               MetadataKey = "TIPL"
	MetadataID3MetadataKeyLanguage                             MetadataKey = "TLAN"
	MetadataID3MetadataKeyLeadPerformer                        MetadataKey = "TPE1"
	MetadataID3MetadataKeyLength                               MetadataKey = "TLEN"
	MetadataID3MetadataKeyLink                                 MetadataKey = "LINK"
	MetadataID3MetadataKeyLyricist                             MetadataKey = "TEXT"
	MetadataID3MetadataKeyMPEGLocationLookupTable              MetadataKey = "MLLT"
	MetadataID3MetadataKeyMediaType                            MetadataKey = "TMED"
	MetadataID3MetadataKeyModifiedBy                           MetadataKey = "TPE4"
	MetadataID3MetadataKeyMood                                 MetadataKey = "TMOO"
	MetadataID3MetadataKeyMusicCDIdentifier                    MetadataKey = "MCDI"
	MetadataID3MetadataKeyMusicianCreditsList                  MetadataKey = "TMCL"
	MetadataID3MetadataKeyOfficialArtistWebpage                MetadataKey = "WOAR"
	MetadataID3MetadataKeyOfficialAudioFileWebpage             MetadataKey = "WOAF"
	MetadataID3MetadataKeyOfficialAudioSourceWebpage           MetadataKey = "WOAS"
	MetadataID3MetadataKeyOfficialInternetRadioStationHomepage MetadataKey = "WORS"
	MetadataID3MetadataKeyOfficialPublisherWebpage             MetadataKey = "WPUB"
	MetadataID3MetadataKeyOriginalAlbumTitle                   MetadataKey = "TOAL"
	MetadataID3MetadataKeyOriginalArtist                       MetadataKey = "TOPE"
	MetadataID3MetadataKeyOriginalFilename                     MetadataKey = "TOFN"
	MetadataID3MetadataKeyOriginalLyricist                     MetadataKey = "TOLY"
	MetadataID3MetadataKeyOriginalReleaseTime                  MetadataKey = "TDOR"
	MetadataID3MetadataKeyOriginalReleaseYear                  MetadataKey = "TORY"
	MetadataID3MetadataKeyOwnership                            MetadataKey = "OWNE"
	MetadataID3MetadataKeyPartOfASet                           MetadataKey = "TPOS"
	MetadataID3MetadataKeyPayment                              MetadataKey = "WPAY"
	MetadataID3MetadataKeyPerformerSortOrder                   MetadataKey = "TSOP"
	MetadataID3MetadataKeyPlayCounter                          MetadataKey = "PCNT"
	MetadataID3MetadataKeyPlaylistDelay                        MetadataKey = "TDLY"
	MetadataID3MetadataKeyPopularimeter                        MetadataKey = "POPM"
	MetadataID3MetadataKeyPositionSynchronization              MetadataKey = "POSS"
	MetadataID3MetadataKeyPrivate                              MetadataKey = "PRIV"
	MetadataID3MetadataKeyProducedNotice                       MetadataKey = "TPRO"
	MetadataID3MetadataKeyPublisher                            MetadataKey = "TPUB"
	MetadataID3MetadataKeyRecommendedBufferSize                MetadataKey = "RBUF"
	MetadataID3MetadataKeyRecordingDates                       MetadataKey = "TRDA"
	MetadataID3MetadataKeyRecordingTime                        MetadataKey = "TDRC"
	MetadataID3MetadataKeyRelativeVolumeAdjustment             MetadataKey = "RVAD"
	MetadataID3MetadataKeyRelativeVolumeAdjustment2            MetadataKey = "RVA2"
	MetadataID3MetadataKeyReleaseTime                          MetadataKey = "TDRL"
	MetadataID3MetadataKeyReverb                               MetadataKey = "RVRB"
	MetadataID3MetadataKeySeek                                 MetadataKey = "SEEK"
	MetadataID3MetadataKeySetSubtitle                          MetadataKey = "TSST"
	MetadataID3MetadataKeySignature                            MetadataKey = "SIGN"
	MetadataID3MetadataKeySize                                 MetadataKey = "TSIZ"
	MetadataID3MetadataKeySubTitle                             MetadataKey = "TIT3"
	MetadataID3MetadataKeySynchronizedLyric                    MetadataKey = "SYLT"
	MetadataID3MetadataKeySynchronizedTempoCodes               MetadataKey = "SYTC"
	MetadataID3MetadataKeyTaggingTime                          MetadataKey = "TDTG"
	MetadataID3MetadataKeyTermsOfUse                           MetadataKey = "USER"
	MetadataID3MetadataKeyTime                                 MetadataKey = "TIME"
	MetadataID3MetadataKeyTitleDescription                     MetadataKey = "TIT2"
	MetadataID3MetadataKeyTitleSortOrder                       MetadataKey = "TSOT"
	MetadataID3MetadataKeyTrackNumber                          MetadataKey = "TRCK"
	MetadataID3MetadataKeyUniqueFileIdentifier                 MetadataKey = "UFID"
	MetadataID3MetadataKeyUnsynchronizedLyric                  MetadataKey = "USLT"
	MetadataID3MetadataKeyUserText                             MetadataKey = "TXXX"
	MetadataID3MetadataKeyUserURL                              MetadataKey = "WXXX"
	MetadataID3MetadataKeyYear                                 MetadataKey = "TYER"
	MetadataISOUserDataKeyAccessibilityDescription             MetadataKey = "ades"
	MetadataISOUserDataKeyCopyright                            MetadataKey = "cprt"
	MetadataISOUserDataKeyDate                                 MetadataKey = "date"
	MetadataISOUserDataKeyTaggedCharacteristic                 MetadataKey = "tagc"
	MetadataIcyMetadataKeyStreamTitle                          MetadataKey = "StreamTitle"
	MetadataIcyMetadataKeyStreamURL                            MetadataKey = "StreamUrl"
	MetadataQuickTimeMetadataKeyAccessibilityDescription       MetadataKey = "com.apple.quicktime.accessibility.description"
	MetadataQuickTimeMetadataKeyAlbum                          MetadataKey = "com.apple.quicktime.album"
	MetadataQuickTimeMetadataKeyArranger                       MetadataKey = "com.apple.quicktime.arranger"
	MetadataQuickTimeMetadataKeyArtist                         MetadataKey = "com.apple.quicktime.artist"
	MetadataQuickTimeMetadataKeyArtwork                        MetadataKey = "com.apple.quicktime.artwork"
	MetadataQuickTimeMetadataKeyAuthor                         MetadataKey = "com.apple.quicktime.author"
	MetadataQuickTimeMetadataKeyCameraFrameReadoutTime         MetadataKey = "com.apple.quicktime.camera.framereadouttimeinmicroseconds"
	MetadataQuickTimeMetadataKeyCameraIdentifier               MetadataKey = "com.apple.quicktime.camera.identifier"
	MetadataQuickTimeMetadataKeyCollectionUser                 MetadataKey = "com.apple.quicktime.collection.user"
	MetadataQuickTimeMetadataKeyComment                        MetadataKey = "com.apple.quicktime.comment"
	MetadataQuickTimeMetadataKeyComposer                       MetadataKey = "com.apple.quicktime.composer"
	MetadataQuickTimeMetadataKeyContentIdentifier              MetadataKey = "com.apple.quicktime.content.identifier"
	MetadataQuickTimeMetadataKeyCopyright                      MetadataKey = "com.apple.quicktime.copyright"
	MetadataQuickTimeMetadataKeyCreationDate                   MetadataKey = "com.apple.quicktime.creationdate"
	MetadataQuickTimeMetadataKeyCredits                        MetadataKey = "com.apple.quicktime.credits"
	MetadataQuickTimeMetadataKeyDescription                    MetadataKey = "com.apple.quicktime.description"
	MetadataQuickTimeMetadataKeyDirectionFacing                MetadataKey = "com.apple.quicktime.direction.facing"
	MetadataQuickTimeMetadataKeyDirectionMotion                MetadataKey = "com.apple.quicktime.direction.motion"
	MetadataQuickTimeMetadataKeyDirector                       MetadataKey = "com.apple.quicktime.director"
	MetadataQuickTimeMetadataKeyDisplayName                    MetadataKey = "com.apple.quicktime.displayname"
	MetadataQuickTimeMetadataKeyEncodedBy                      MetadataKey = "com.apple.quicktime.encodedby"
	MetadataQuickTimeMetadataKeyGenre                          MetadataKey = "com.apple.quicktime.genre"
	MetadataQuickTimeMetadataKeyInformation                    MetadataKey = "com.apple.quicktime.information"
	MetadataQuickTimeMetadataKeyIsMontage                      MetadataKey = "com.apple.quicktime.is-montage"
	MetadataQuickTimeMetadataKeyKeywords                       MetadataKey = "com.apple.quicktime.keywords"
	MetadataQuickTimeMetadataKeyLocationBody                   MetadataKey = "com.apple.quicktime.location.body"
	MetadataQuickTimeMetadataKeyLocationDate                   MetadataKey = "com.apple.quicktime.location.date"
	MetadataQuickTimeMetadataKeyLocationISO6709                MetadataKey = "com.apple.quicktime.location.ISO6709"
	MetadataQuickTimeMetadataKeyLocationName                   MetadataKey = "com.apple.quicktime.location.name"
	MetadataQuickTimeMetadataKeyLocationNote                   MetadataKey = "com.apple.quicktime.location.note"
	MetadataQuickTimeMetadataKeyLocationRole                   MetadataKey = "com.apple.quicktime.location.role"
	MetadataQuickTimeMetadataKeyMake                           MetadataKey = "com.apple.quicktime.make"
	MetadataQuickTimeMetadataKeyModel                          MetadataKey = "com.apple.quicktime.model"
	MetadataQuickTimeMetadataKeyOriginalArtist                 MetadataKey = "com.apple.quicktime.originalartist"
	MetadataQuickTimeMetadataKeyPerformer                      MetadataKey = "com.apple.quicktime.performer"
	MetadataQuickTimeMetadataKeyPhonogramRights                MetadataKey = "com.apple.quicktime.phonogramrights"
	MetadataQuickTimeMetadataKeyProducer                       MetadataKey = "com.apple.quicktime.producer"
	MetadataQuickTimeMetadataKeyPublisher                      MetadataKey = "com.apple.quicktime.publisher"
	MetadataQuickTimeMetadataKeyRatingUser                     MetadataKey = "com.apple.quicktime.rating.user"
	MetadataQuickTimeMetadataKeySoftware                       MetadataKey = "com.apple.quicktime.software"
	MetadataQuickTimeMetadataKeyTitle                          MetadataKey = "com.apple.quicktime.title"
	MetadataQuickTimeMetadataKeyYear                           MetadataKey = "com.apple.quicktime.year"
	MetadataQuickTimeMetadataKeyiXML                           MetadataKey = "info.ixml.xml"
	MetadataQuickTimeUserDataKeyAccessibilityDescription       MetadataKey = "@ade"
	MetadataQuickTimeUserDataKeyAlbum                          MetadataKey = "@alb"
	MetadataQuickTimeUserDataKeyArranger                       MetadataKey = "@arg"
	MetadataQuickTimeUserDataKeyArtist                         MetadataKey = "@ART"
	MetadataQuickTimeUserDataKeyAuthor                         MetadataKey = "@aut"
	MetadataQuickTimeUserDataKeyChapter                        MetadataKey = "@chp"
	MetadataQuickTimeUserDataKeyComment                        MetadataKey = "@cmt"
	MetadataQuickTimeUserDataKeyComposer                       MetadataKey = "@com"
	MetadataQuickTimeUserDataKeyCopyright                      MetadataKey = "@cpy"
	MetadataQuickTimeUserDataKeyCreationDate                   MetadataKey = "@day"
	MetadataQuickTimeUserDataKeyCredits                        MetadataKey = "@src"
	MetadataQuickTimeUserDataKeyDescription                    MetadataKey = "@des"
	MetadataQuickTimeUserDataKeyDirector                       MetadataKey = "@dir"
	MetadataQuickTimeUserDataKeyDisclaimer                     MetadataKey = "@dis"
	MetadataQuickTimeUserDataKeyEncodedBy                      MetadataKey = "@enc"
	MetadataQuickTimeUserDataKeyFullName                       MetadataKey = "@nam"
	MetadataQuickTimeUserDataKeyGenre                          MetadataKey = "@gen"
	MetadataQuickTimeUserDataKeyHostComputer                   MetadataKey = "@hst"
	MetadataQuickTimeUserDataKeyInformation                    MetadataKey = "@inf"
	MetadataQuickTimeUserDataKeyKeywords                       MetadataKey = "@key"
	MetadataQuickTimeUserDataKeyLocationISO6709                MetadataKey = "@xyz"
	MetadataQuickTimeUserDataKeyMake                           MetadataKey = "@mak"
	MetadataQuickTimeUserDataKeyModel                          MetadataKey = "@mod"
	MetadataQuickTimeUserDataKeyOriginalArtist                 MetadataKey = "@ope"
	MetadataQuickTimeUserDataKeyOriginalFormat                 MetadataKey = "@fmt"
	MetadataQuickTimeUserDataKeyOriginalSource                 MetadataKey = "@src"
	MetadataQuickTimeUserDataKeyPerformers                     MetadataKey = "@prf"
	MetadataQuickTimeUserDataKeyPhonogramRights                MetadataKey = "@phg"
	MetadataQuickTimeUserDataKeyProducer                       MetadataKey = "@prd"
	MetadataQuickTimeUserDataKeyProduct                        MetadataKey = "@PRD"
	MetadataQuickTimeUserDataKeyPublisher                      MetadataKey = "@pub"
	MetadataQuickTimeUserDataKeySoftware                       MetadataKey = "@swr"
	MetadataQuickTimeUserDataKeySpecialPlaybackRequirements    MetadataKey = "@req"
	MetadataQuickTimeUserDataKeyTaggedCharacteristic           MetadataKey = "tagc"
	MetadataQuickTimeUserDataKeyTrack                          MetadataKey = "@trk"
	MetadataQuickTimeUserDataKeyTrackName                      MetadataKey = "tnam"
	MetadataQuickTimeUserDataKeyURLLink                        MetadataKey = "@url"
	MetadataQuickTimeUserDataKeyWarning                        MetadataKey = "@wrn"
	MetadataQuickTimeUserDataKeyWriter                         MetadataKey = "@wrt"
	MetadataiTunesMetadataKeyAccountKind                       MetadataKey = "akID"
	MetadataiTunesMetadataKeyAcknowledgement                   MetadataKey = "@cak"
	MetadataiTunesMetadataKeyAlbum                             MetadataKey = "@alb"
	MetadataiTunesMetadataKeyAlbumArtist                       MetadataKey = "aART"
	MetadataiTunesMetadataKeyAppleID                           MetadataKey = "apID"
	MetadataiTunesMetadataKeyArranger                          MetadataKey = "@arg"
	MetadataiTunesMetadataKeyArtDirector                       MetadataKey = "@ard"
	MetadataiTunesMetadataKeyArtist                            MetadataKey = "@ART"
	MetadataiTunesMetadataKeyArtistID                          MetadataKey = "atID"
	MetadataiTunesMetadataKeyAuthor                            MetadataKey = "@aut"
	MetadataiTunesMetadataKeyBeatsPerMin                       MetadataKey = "tmpo"
	MetadataiTunesMetadataKeyComposer                          MetadataKey = "@wrt"
	MetadataiTunesMetadataKeyConductor                         MetadataKey = "@con"
	MetadataiTunesMetadataKeyContentRating                     MetadataKey = "rtng"
	MetadataiTunesMetadataKeyCopyright                         MetadataKey = "cprt"
	MetadataiTunesMetadataKeyCoverArt                          MetadataKey = "covr"
	MetadataiTunesMetadataKeyCredits                           MetadataKey = "@src"
	MetadataiTunesMetadataKeyDescription                       MetadataKey = "@des"
	MetadataiTunesMetadataKeyDirector                          MetadataKey = "@dir"
	MetadataiTunesMetadataKeyDiscCompilation                   MetadataKey = "cpil"
	MetadataiTunesMetadataKeyDiscNumber                        MetadataKey = "disk"
	MetadataiTunesMetadataKeyEQ                                MetadataKey = "@equ"
	MetadataiTunesMetadataKeyEncodedBy                         MetadataKey = "@enc"
	MetadataiTunesMetadataKeyEncodingTool                      MetadataKey = "@too"
	MetadataiTunesMetadataKeyExecProducer                      MetadataKey = "@xpd"
	MetadataiTunesMetadataKeyGenreID                           MetadataKey = "geID"
	MetadataiTunesMetadataKeyGrouping                          MetadataKey = "grup"
	MetadataiTunesMetadataKeyLinerNotes                        MetadataKey = "@lnt"
	MetadataiTunesMetadataKeyLyrics                            MetadataKey = "@lyr"
	MetadataiTunesMetadataKeyOnlineExtras                      MetadataKey = "@url"
	MetadataiTunesMetadataKeyOriginalArtist                    MetadataKey = "@ope"
	MetadataiTunesMetadataKeyPerformer                         MetadataKey = "@prf"
	MetadataiTunesMetadataKeyPhonogramRights                   MetadataKey = "@phg"
	MetadataiTunesMetadataKeyPlaylistID                        MetadataKey = "plID"
	MetadataiTunesMetadataKeyPredefinedGenre                   MetadataKey = "gnre"
	MetadataiTunesMetadataKeyProducer                          MetadataKey = "@prd"
	MetadataiTunesMetadataKeyPublisher                         MetadataKey = "@pub"
	MetadataiTunesMetadataKeyRecordCompany                     MetadataKey = "@mak"
	MetadataiTunesMetadataKeyReleaseDate                       MetadataKey = "@day"
	MetadataiTunesMetadataKeySoloist                           MetadataKey = "@sol"
	MetadataiTunesMetadataKeySongID                            MetadataKey = "cnID"
	MetadataiTunesMetadataKeySongName                          MetadataKey = "@nam"
	MetadataiTunesMetadataKeySoundEngineer                     MetadataKey = "@sne"
	MetadataiTunesMetadataKeyThanks                            MetadataKey = "@thx"
	MetadataiTunesMetadataKeyTrackNumber                       MetadataKey = "trkn"
	MetadataiTunesMetadataKeyTrackSubTitle                     MetadataKey = "@st3"
	MetadataiTunesMetadataKeyUserComment                       MetadataKey = "@cmt"
	MetadataiTunesMetadataKeyUserGenre                         MetadataKey = "@gen"
)

// A structure that defines a metadata key space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatakeyspace?language=objc
type MetadataKeySpace string

const (
	MetadataKeySpaceAudioFile         MetadataKeySpace = "caaf"
	MetadataKeySpaceCommon            MetadataKeySpace = "comn"
	MetadataKeySpaceHLSDateRange      MetadataKeySpace = "lsdr"
	MetadataKeySpaceID3               MetadataKeySpace = "org.id3"
	MetadataKeySpaceISOUserData       MetadataKeySpace = "uiso"
	MetadataKeySpaceIcy               MetadataKeySpace = "icy"
	MetadataKeySpaceQuickTimeMetadata MetadataKeySpace = "mdta"
	MetadataKeySpaceQuickTimeUserData MetadataKeySpace = "udta"
	MetadataKeySpaceiTunes            MetadataKeySpace = "itsk"
)

// Constants that identify metadata object types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataobjecttype?language=objc
type MetadataObjectType string

const (
	MetadataObjectTypeAztecCode           MetadataObjectType = "org.iso.Aztec"
	MetadataObjectTypeCatBody             MetadataObjectType = "catBody"
	MetadataObjectTypeCode128Code         MetadataObjectType = "org.iso.Code128"
	MetadataObjectTypeCode39Code          MetadataObjectType = "org.iso.Code39"
	MetadataObjectTypeCode39Mod43Code     MetadataObjectType = "org.iso.Code39Mod43"
	MetadataObjectTypeCode93Code          MetadataObjectType = "com.intermec.Code93"
	MetadataObjectTypeDataMatrixCode      MetadataObjectType = "org.iso.DataMatrix"
	MetadataObjectTypeDogBody             MetadataObjectType = "dogBody"
	MetadataObjectTypeEAN13Code           MetadataObjectType = "org.gs1.EAN-13"
	MetadataObjectTypeEAN8Code            MetadataObjectType = "org.gs1.EAN-8"
	MetadataObjectTypeFace                MetadataObjectType = "face"
	MetadataObjectTypeHumanBody           MetadataObjectType = "humanBody"
	MetadataObjectTypeITF14Code           MetadataObjectType = "org.gs1.ITF14"
	MetadataObjectTypeInterleaved2of5Code MetadataObjectType = "org.ansi.Interleaved2of5"
	MetadataObjectTypePDF417Code          MetadataObjectType = "org.iso.PDF417"
	MetadataObjectTypeQRCode              MetadataObjectType = "org.iso.QRCode"
	MetadataObjectTypeSalientObject       MetadataObjectType = "salientObject"
	MetadataObjectTypeUPCECode            MetadataObjectType = "org.gs1.UPC-E"
)

// A structure that defines options to control the writing of a movie header to a destination URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmoviewritingoptions?language=objc
type MovieWritingOptions uint

const (
	MovieWritingAddMovieHeaderToDestination          MovieWritingOptions = 0
	MovieWritingTruncateDestinationToMovieHeaderOnly MovieWritingOptions = 1
)

// A structure that defines preset configurations for an output settings assistant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingspreset?language=objc
type OutputSettingsPreset string

const (
	OutputSettingsPreset1280x720               OutputSettingsPreset = "AVOutputSettingsPreset1280x720"
	OutputSettingsPreset1920x1080              OutputSettingsPreset = "AVOutputSettingsPreset1920x1080"
	OutputSettingsPreset3840x2160              OutputSettingsPreset = "AVOutputSettingsPreset3840x2160"
	OutputSettingsPreset640x480                OutputSettingsPreset = "AVOutputSettingsPreset640x480"
	OutputSettingsPreset960x540                OutputSettingsPreset = "AVOutputSettingsPreset960x540"
	OutputSettingsPresetHEVC1920x1080          OutputSettingsPreset = "AVOutputSettingsPresetHEVC1920x1080"
	OutputSettingsPresetHEVC1920x1080WithAlpha OutputSettingsPreset = "AVOutputSettingsPresetHEVC1920x1080WithAlpha"
	OutputSettingsPresetHEVC3840x2160          OutputSettingsPreset = "AVOutputSettingsPresetHEVC3840x2160"
	OutputSettingsPresetHEVC3840x2160WithAlpha OutputSettingsPreset = "AVOutputSettingsPresetHEVC3840x2160WithAlpha"
)

// The actions a player can take when it finishes playing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeractionatitemend?language=objc
type PlayerActionAtItemEnd int

const (
	PlayerActionAtItemEndAdvance PlayerActionAtItemEnd = 0
	PlayerActionAtItemEndNone    PlayerActionAtItemEnd = 2
	PlayerActionAtItemEndPause   PlayerActionAtItemEnd = 1
)

// Policies that describe playback behavior when an app transitions to the background while playing video. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeraudiovisualbackgroundplaybackpolicy?language=objc
type PlayerAudiovisualBackgroundPlaybackPolicy int

const (
	PlayerAudiovisualBackgroundPlaybackPolicyAutomatic           PlayerAudiovisualBackgroundPlaybackPolicy = 1
	PlayerAudiovisualBackgroundPlaybackPolicyContinuesIfPossible PlayerAudiovisualBackgroundPlaybackPolicy = 3
	PlayerAudiovisualBackgroundPlaybackPolicyPauses              PlayerAudiovisualBackgroundPlaybackPolicy = 2
)

// Constants that define restrictions on the playback of interstitial content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventrestrictions?language=objc
type PlayerInterstitialEventRestrictions uint

const (
	PlayerInterstitialEventRestrictionConstrainsSeekingForwardInPrimaryContent      PlayerInterstitialEventRestrictions = 1
	PlayerInterstitialEventRestrictionDefaultPolicy                                 PlayerInterstitialEventRestrictions = 0
	PlayerInterstitialEventRestrictionNone                                          PlayerInterstitialEventRestrictions = 0
	PlayerInterstitialEventRestrictionRequiresPlaybackAtPreferredRateForAdvancement PlayerInterstitialEventRestrictions = 4
)

// A text styling resolution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemlegibleoutputtextstylingresolution?language=objc
type PlayerItemLegibleOutputTextStylingResolution string

const (
	PlayerItemLegibleOutputTextStylingResolutionDefault            PlayerItemLegibleOutputTextStylingResolution = "AVPlayerItemLegibleOutputTextStylingResolutionDefault"
	PlayerItemLegibleOutputTextStylingResolutionSourceAndRulesOnly PlayerItemLegibleOutputTextStylingResolution = "AVPlayerItemLegibleOutputTextStylingResolutionSourceAndRulesOnly"
)

// The statuses for a player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemstatus?language=objc
type PlayerItemStatus int

const (
	PlayerItemStatusFailed      PlayerItemStatus = 2
	PlayerItemStatusReadyToPlay PlayerItemStatus = 1
	PlayerItemStatusUnknown     PlayerItemStatus = 0
)

// Status constants that indicate whether a looper can successfully perform looping playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlooperstatus?language=objc
type PlayerLooperStatus int

const (
	PlayerLooperStatusCancelled PlayerLooperStatus = 3
	PlayerLooperStatusFailed    PlayerLooperStatus = 2
	PlayerLooperStatusReady     PlayerLooperStatus = 1
	PlayerLooperStatusUnknown   PlayerLooperStatus = 0
)

// A structure that represents a rate change reason. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerratedidchangereason?language=objc
type PlayerRateDidChangeReason string

const (
	PlayerRateDidChangeReasonAppBackgrounded         PlayerRateDidChangeReason = "AVPlayerRateDidChangeReasonAppBackgrounded"
	PlayerRateDidChangeReasonAudioSessionInterrupted PlayerRateDidChangeReason = "AVPlayerRateDidChangeReasonAudioSessionInterrupted"
	PlayerRateDidChangeReasonSetRateCalled           PlayerRateDidChangeReason = "AVPlayerRateDidChangeReasonSetRateCalled"
	PlayerRateDidChangeReasonSetRateFailed           PlayerRateDidChangeReason = "AVPlayerRateDidChangeReasonSetRateFailed"
)

// Status values that indicate whether a player can successfully play media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerstatus?language=objc
type PlayerStatus int

const (
	PlayerStatusFailed      PlayerStatus = 2
	PlayerStatusReadyToPlay PlayerStatus = 1
	PlayerStatusUnknown     PlayerStatus = 0
)

// Constants that indicate the state of playback control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayertimecontrolstatus?language=objc
type PlayerTimeControlStatus int

const (
	PlayerTimeControlStatusPaused                       PlayerTimeControlStatus = 0
	PlayerTimeControlStatusPlaying                      PlayerTimeControlStatus = 2
	PlayerTimeControlStatusWaitingToPlayAtSpecifiedRate PlayerTimeControlStatus = 1
)

// The reasons a player is waiting to begin or resume playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerwaitingreason?language=objc
type PlayerWaitingReason string

const (
	PlayerWaitingDuringInterstitialEventReason      PlayerWaitingReason = "AVPlayerWaitingDuringInterstitialEventReason"
	PlayerWaitingForCoordinatedPlaybackReason       PlayerWaitingReason = "AVPlayerWaitingForCoordinatedPlaybackReason"
	PlayerWaitingToMinimizeStallsReason             PlayerWaitingReason = "AVPlayerWaitingToMinimizeStallsReason"
	PlayerWaitingWhileEvaluatingBufferingRateReason PlayerWaitingReason = "AVPlayerWaitingWhileEvaluatingBufferingRateReason"
	PlayerWaitingWithNoItemToPlayReason             PlayerWaitingReason = "AVPlayerWaitingWithNoItemToPlayReason"
)

// The statuses for sample buffer rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avqueuedsamplebufferrenderingstatus?language=objc
type QueuedSampleBufferRenderingStatus int

const (
	QueuedSampleBufferRenderingStatusFailed    QueuedSampleBufferRenderingStatus = 2
	QueuedSampleBufferRenderingStatusRendering QueuedSampleBufferRenderingStatus = 1
	QueuedSampleBufferRenderingStatusUnknown   QueuedSampleBufferRenderingStatus = 0
)

// The modes that describe the buffer request direction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrequestdirection?language=objc
type SampleBufferRequestDirection int

const (
	SampleBufferRequestDirectionForward SampleBufferRequestDirection = 1
	SampleBufferRequestDirectionNone    SampleBufferRequestDirection = 0
	SampleBufferRequestDirectionReverse SampleBufferRequestDirection = -1
)

// The modes in which a sample buffer generator processes a request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrequestmode?language=objc
type SampleBufferRequestMode int

const (
	SampleBufferRequestModeImmediate     SampleBufferRequestMode = 0
	SampleBufferRequestModeOpportunistic SampleBufferRequestMode = 2
	SampleBufferRequestModeScheduled     SampleBufferRequestMode = 1
)

// String constants defining the types of segmentation matte images that you can capture along with the primary image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsemanticsegmentationmattetype?language=objc
type SemanticSegmentationMatteType string

const (
	SemanticSegmentationMatteTypeGlasses SemanticSegmentationMatteType = "AVSemanticSegmentationMatteTypeGlasses"
	SemanticSegmentationMatteTypeHair    SemanticSegmentationMatteType = "AVSemanticSegmentationMatteTypeHair"
	SemanticSegmentationMatteTypeSkin    SemanticSegmentationMatteType = "AVSemanticSegmentationMatteTypeSkin"
	SemanticSegmentationMatteTypeTeeth   SemanticSegmentationMatteType = "AVSemanticSegmentationMatteTypeTeeth"
)

// Constants that define track association types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtrackassociationtype?language=objc
type TrackAssociationType string

const (
	TrackAssociationTypeAudioFallback       TrackAssociationType = "fall"
	TrackAssociationTypeChapterList         TrackAssociationType = "chap"
	TrackAssociationTypeForcedSubtitlesOnly TrackAssociationType = "forc"
	TrackAssociationTypeMetadataReferent    TrackAssociationType = "cdsc"
	TrackAssociationTypeSelectionFollower   TrackAssociationType = "folw"
	TrackAssociationTypeTimecode            TrackAssociationType = "tmcd"
)

// Defines the preferences the player item uses when selecting variant playlists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvariantpreferences?language=objc
type VariantPreferences uint

const (
	VariantPreferenceNone                       VariantPreferences = 0
	VariantPreferenceScalabilityToLosslessAudio VariantPreferences = 1
)

// A value that describes how a video is scaled or cropped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideoaperturemode?language=objc
type VideoApertureMode string

const (
	VideoApertureModeCleanAperture      VideoApertureMode = "CleanAperture"
	VideoApertureModeEncodedPixels      VideoApertureMode = "EncodedPixels"
	VideoApertureModeProductionAperture VideoApertureMode = "ProductionAperture"
)

// A set of constants describing the codecs that the system supports for video capture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocodectype?language=objc
type VideoCodecType string

const (
	VideoCodecTypeAppleProRes422      VideoCodecType = "apcn"
	VideoCodecTypeAppleProRes422HQ    VideoCodecType = "apch"
	VideoCodecTypeAppleProRes422LT    VideoCodecType = "apcs"
	VideoCodecTypeAppleProRes422Proxy VideoCodecType = "apco"
	VideoCodecTypeAppleProRes4444     VideoCodecType = "ap4h"
	VideoCodecTypeH264                VideoCodecType = "avc1"
	VideoCodecTypeHEVC                VideoCodecType = "hvc1"
	VideoCodecTypeHEVCWithAlpha       VideoCodecType = "muxa"
	VideoCodecTypeJPEG                VideoCodecType = "jpeg"
)

// Constants that indicate which interlacing modes the connection applies to video flowing through it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideofieldmode?language=objc
type VideoFieldMode int

const (
	VideoFieldModeBoth        VideoFieldMode = 0
	VideoFieldModeBottomOnly  VideoFieldMode = 2
	VideoFieldModeDeinterlace VideoFieldMode = 3
	VideoFieldModeTopOnly     VideoFieldMode = 1
)

// Constants that describe a video variant’s dynamic range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideorange?language=objc
type VideoRange string

const (
	VideoRangeHLG VideoRange = "AVVideoRangeHLG"
	VideoRangePQ  VideoRange = "AVVideoRangePQ"
	VideoRangeSDR VideoRange = "AVVideoRangeSDR"
)
