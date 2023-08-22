// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureDevice] class.
var CaptureDeviceClass = _CaptureDeviceClass{objc.GetClass("AVCaptureDevice")}

type _CaptureDeviceClass struct {
	objc.Class
}

// An interface definition for the [CaptureDevice] class.
type ICaptureDevice interface {
	objc.IObject
	HasMediaType(mediaType MediaType) bool
	SetTorchModeOnWithLevelError(torchLevel float64, outError foundation.IError) bool
	IsWhiteBalanceModeSupported(whiteBalanceMode CaptureWhiteBalanceMode) bool
	SetTransportControlsPlaybackModeSpeed(mode CaptureDeviceTransportControlsPlaybackMode, speed CaptureDeviceTransportControlsSpeed)
	UnlockForConfiguration()
	IsFocusModeSupported(focusMode CaptureFocusMode) bool
	IsExposureModeSupported(exposureMode CaptureExposureMode) bool
	LockForConfiguration(outError foundation.IError) bool
	IsTorchModeSupported(torchMode CaptureTorchMode) bool
	SetPrimaryConstituentDeviceSwitchingBehaviorRestrictedSwitchingBehaviorConditions(switchingBehavior CapturePrimaryConstituentDeviceSwitchingBehavior, restrictedSwitchingBehaviorConditions CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions)
	SupportsAVCaptureSessionPreset(preset CaptureSessionPreset) bool
	Formats() []CaptureDeviceFormat
	IsInUseByAnotherApplication() bool
	ModelID() string
	TransportControlsPlaybackMode() CaptureDeviceTransportControlsPlaybackMode
	TransportType() int32
	TorchLevel() float64
	ActiveVideoMaxFrameDuration() coremedia.Time
	SetActiveVideoMaxFrameDuration(value coremedia.Time)
	HasFlash() bool
	IsExposurePointOfInterestSupported() bool
	FocusPointOfInterest() coregraphics.Point
	SetFocusPointOfInterest(value coregraphics.Point)
	IsTorchActive() bool
	PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions
	ActivePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions
	IsConnected() bool
	FlashMode() CaptureFlashMode
	SetFlashMode(value CaptureFlashMode)
	MinimumFocusDistance() int
	IsAdjustingExposure() bool
	LocalizedName() string
	WhiteBalanceMode() CaptureWhiteBalanceMode
	SetWhiteBalanceMode(value CaptureWhiteBalanceMode)
	UniqueID() string
	ActivePrimaryConstituentDeviceSwitchingBehavior() CapturePrimaryConstituentDeviceSwitchingBehavior
	HasTorch() bool
	Position() CaptureDevicePosition
	ActiveInputSource() CaptureDeviceInputSource
	SetActiveInputSource(value ICaptureDeviceInputSource)
	IsTorchAvailable() bool
	TransportControlsSpeed() CaptureDeviceTransportControlsSpeed
	SupportedFallbackPrimaryConstituentDevices() []CaptureDevice
	IsSuspended() bool
	Manufacturer() string
	TorchMode() CaptureTorchMode
	SetTorchMode(value CaptureTorchMode)
	PrimaryConstituentDeviceSwitchingBehavior() CapturePrimaryConstituentDeviceSwitchingBehavior
	ActiveVideoMinFrameDuration() coremedia.Time
	SetActiveVideoMinFrameDuration(value coremedia.Time)
	IsFlashAvailable() bool
	IsPortraitEffectActive() bool
	FocusMode() CaptureFocusMode
	SetFocusMode(value CaptureFocusMode)
	InputSources() []CaptureDeviceInputSource
	ExposureMode() CaptureExposureMode
	SetExposureMode(value CaptureExposureMode)
	IsFocusPointOfInterestSupported() bool
	TransportControlsSupported() bool
	ActivePrimaryConstituentDevice() CaptureDevice
	LinkedDevices() []CaptureDevice
	IsAdjustingWhiteBalance() bool
	DeviceType() CaptureDeviceType
	IsAdjustingFocus() bool
	ExposurePointOfInterest() coregraphics.Point
	SetExposurePointOfInterest(value coregraphics.Point)
	ActiveColorSpace() CaptureColorSpace
	SetActiveColorSpace(value CaptureColorSpace)
	FallbackPrimaryConstituentDevices() []CaptureDevice
	SetFallbackPrimaryConstituentDevices(value []ICaptureDevice)
	ActiveFormat() CaptureDeviceFormat
	SetActiveFormat(value ICaptureDeviceFormat)
	IsCenterStageActive() bool
}

// An object that represents a hardware or virtual capture device like a camera or microphone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice?language=objc
type CaptureDevice struct {
	objc.Object
}

func CaptureDeviceFrom(ptr unsafe.Pointer) CaptureDevice {
	return CaptureDevice{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptureDeviceClass) Alloc() CaptureDevice {
	rv := objc.Call[CaptureDevice](cc, objc.Sel("alloc"))
	return rv
}

func CaptureDevice_Alloc() CaptureDevice {
	return CaptureDeviceClass.Alloc()
}

func (cc _CaptureDeviceClass) New() CaptureDevice {
	rv := objc.Call[CaptureDevice](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureDevice() CaptureDevice {
	return CaptureDeviceClass.New()
}

func (c_ CaptureDevice) Init() CaptureDevice {
	rv := objc.Call[CaptureDevice](c_, objc.Sel("init"))
	return rv
}

// Returns an authorization status that indicates whether the user grants the app permission to capture media of a particular type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1624613-authorizationstatusformediatype?language=objc
func (cc _CaptureDeviceClass) AuthorizationStatusForMediaType(mediaType MediaType) AuthorizationStatus {
	rv := objc.Call[AuthorizationStatus](cc, objc.Sel("authorizationStatusForMediaType:"), mediaType)
	return rv
}

// Returns an authorization status that indicates whether the user grants the app permission to capture media of a particular type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1624613-authorizationstatusformediatype?language=objc
func CaptureDevice_AuthorizationStatusForMediaType(mediaType MediaType) AuthorizationStatus {
	return CaptureDeviceClass.AuthorizationStatusForMediaType(mediaType)
}

// Returns a Boolean value that indicates whether the device captures media of a particular type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1389487-hasmediatype?language=objc
func (c_ CaptureDevice) HasMediaType(mediaType MediaType) bool {
	rv := objc.Call[bool](c_, objc.Sel("hasMediaType:"), mediaType)
	return rv
}

// Sets the illumination level when in torch mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1624609-settorchmodeonwithlevel?language=objc
func (c_ CaptureDevice) SetTorchModeOnWithLevelError(torchLevel float64, outError foundation.IError) bool {
	rv := objc.Call[bool](c_, objc.Sel("setTorchModeOnWithLevel:error:"), torchLevel, objc.Ptr(outError))
	return rv
}

// Returns the default device for the specified device type, media type, and position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/2361508-defaultdevicewithdevicetype?language=objc
func (cc _CaptureDeviceClass) DefaultDeviceWithDeviceTypeMediaTypePosition(deviceType CaptureDeviceType, mediaType MediaType, position CaptureDevicePosition) CaptureDevice {
	rv := objc.Call[CaptureDevice](cc, objc.Sel("defaultDeviceWithDeviceType:mediaType:position:"), deviceType, mediaType, position)
	return rv
}

// Returns the default device for the specified device type, media type, and position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/2361508-defaultdevicewithdevicetype?language=objc
func CaptureDevice_DefaultDeviceWithDeviceTypeMediaTypePosition(deviceType CaptureDeviceType, mediaType MediaType, position CaptureDevicePosition) CaptureDevice {
	return CaptureDeviceClass.DefaultDeviceWithDeviceTypeMediaTypePosition(deviceType, mediaType, position)
}

// Creates an object that represents a device with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1388904-devicewithuniqueid?language=objc
func (cc _CaptureDeviceClass) DeviceWithUniqueID(deviceUniqueID string) CaptureDevice {
	rv := objc.Call[CaptureDevice](cc, objc.Sel("deviceWithUniqueID:"), deviceUniqueID)
	return rv
}

// Creates an object that represents a device with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1388904-devicewithuniqueid?language=objc
func CaptureDevice_DeviceWithUniqueID(deviceUniqueID string) CaptureDevice {
	return CaptureDeviceClass.DeviceWithUniqueID(deviceUniqueID)
}

// Returns a Boolean value that indicates whether the device supports the specified white balance mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1388587-iswhitebalancemodesupported?language=objc
func (c_ CaptureDevice) IsWhiteBalanceModeSupported(whiteBalanceMode CaptureWhiteBalanceMode) bool {
	rv := objc.Call[bool](c_, objc.Sel("isWhiteBalanceModeSupported:"), whiteBalanceMode)
	return rv
}

// Sets the transport control’s playback mode and speed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1388481-settransportcontrolsplaybackmode?language=objc
func (c_ CaptureDevice) SetTransportControlsPlaybackModeSpeed(mode CaptureDeviceTransportControlsPlaybackMode, speed CaptureDeviceTransportControlsSpeed) {
	objc.Call[objc.Void](c_, objc.Sel("setTransportControlsPlaybackMode:speed:"), mode, speed)
}

// Returns the default device that captures the specified media type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1386589-defaultdevicewithmediatype?language=objc
func (cc _CaptureDeviceClass) DefaultDeviceWithMediaType(mediaType MediaType) CaptureDevice {
	rv := objc.Call[CaptureDevice](cc, objc.Sel("defaultDeviceWithMediaType:"), mediaType)
	return rv
}

// Returns the default device that captures the specified media type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1386589-defaultdevicewithmediatype?language=objc
func CaptureDevice_DefaultDeviceWithMediaType(mediaType MediaType) CaptureDevice {
	return CaptureDeviceClass.DefaultDeviceWithMediaType(mediaType)
}

// Displays the system’s user interface to configure video effects or microphone modes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3850459-showsystemuserinterface?language=objc
func (cc _CaptureDeviceClass) ShowSystemUserInterface(systemUserInterface CaptureSystemUserInterface) {
	objc.Call[objc.Void](cc, objc.Sel("showSystemUserInterface:"), systemUserInterface)
}

// Displays the system’s user interface to configure video effects or microphone modes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3850459-showsystemuserinterface?language=objc
func CaptureDevice_ShowSystemUserInterface(systemUserInterface CaptureSystemUserInterface) {
	CaptureDeviceClass.ShowSystemUserInterface(systemUserInterface)
}

// Releases exclusive control over device hardware properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1387917-unlockforconfiguration?language=objc
func (c_ CaptureDevice) UnlockForConfiguration() {
	objc.Call[objc.Void](c_, objc.Sel("unlockForConfiguration"))
}

// Requests the user’s permission to allow the app to capture media of a particular type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1624584-requestaccessformediatype?language=objc
func (cc _CaptureDeviceClass) RequestAccessForMediaTypeCompletionHandler(mediaType MediaType, handler func(granted bool)) {
	objc.Call[objc.Void](cc, objc.Sel("requestAccessForMediaType:completionHandler:"), mediaType, handler)
}

// Requests the user’s permission to allow the app to capture media of a particular type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1624584-requestaccessformediatype?language=objc
func CaptureDevice_RequestAccessForMediaTypeCompletionHandler(mediaType MediaType, handler func(granted bool)) {
	CaptureDeviceClass.RequestAccessForMediaTypeCompletionHandler(mediaType, handler)
}

// Returns a Boolean value that indicates whether the device supports the specified focus mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1390215-isfocusmodesupported?language=objc
func (c_ CaptureDevice) IsFocusModeSupported(focusMode CaptureFocusMode) bool {
	rv := objc.Call[bool](c_, objc.Sel("isFocusModeSupported:"), focusMode)
	return rv
}

// Returns a Boolean value that indicates whether a device supports the specified exposure mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1389048-isexposuremodesupported?language=objc
func (c_ CaptureDevice) IsExposureModeSupported(exposureMode CaptureExposureMode) bool {
	rv := objc.Call[bool](c_, objc.Sel("isExposureModeSupported:"), exposureMode)
	return rv
}

// Requests exclusive access to configure device hardware properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1387810-lockforconfiguration?language=objc
func (c_ CaptureDevice) LockForConfiguration(outError foundation.IError) bool {
	rv := objc.Call[bool](c_, objc.Sel("lockForConfiguration:"), objc.Ptr(outError))
	return rv
}

// Returns a Boolean value that indicates whether the device supports the specified torch mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1388822-istorchmodesupported?language=objc
func (c_ CaptureDevice) IsTorchModeSupported(torchMode CaptureTorchMode) bool {
	rv := objc.Call[bool](c_, objc.Sel("isTorchModeSupported:"), torchMode)
	return rv
}

// Sets the switching behavior of the primary constituent device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3875310-setprimaryconstituentdeviceswitc?language=objc
func (c_ CaptureDevice) SetPrimaryConstituentDeviceSwitchingBehaviorRestrictedSwitchingBehaviorConditions(switchingBehavior CapturePrimaryConstituentDeviceSwitchingBehavior, restrictedSwitchingBehaviorConditions CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions) {
	objc.Call[objc.Void](c_, objc.Sel("setPrimaryConstituentDeviceSwitchingBehavior:restrictedSwitchingBehaviorConditions:"), switchingBehavior, restrictedSwitchingBehaviorConditions)
}

// Returns a Boolean value that indicates whether you can use the device with capture session configured with the specified preset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1386263-supportsavcapturesessionpreset?language=objc
func (c_ CaptureDevice) SupportsAVCaptureSessionPreset(preset CaptureSessionPreset) bool {
	rv := objc.Call[bool](c_, objc.Sel("supportsAVCaptureSessionPreset:"), preset)
	return rv
}

// The capture formats a device supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1388738-formats?language=objc
func (c_ CaptureDevice) Formats() []CaptureDeviceFormat {
	rv := objc.Call[[]CaptureDeviceFormat](c_, objc.Sel("formats"))
	return rv
}

// A Boolean value that indicates whether another app is using the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1389512-inusebyanotherapplication?language=objc
func (c_ CaptureDevice) IsInUseByAnotherApplication() bool {
	rv := objc.Call[bool](c_, objc.Sel("isInUseByAnotherApplication"))
	return rv
}

// A Boolean value that indicates whether a user or an app enabled Center Stage on a device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3738419-centerstageenabled?language=objc
func (cc _CaptureDeviceClass) CenterStageEnabled() bool {
	rv := objc.Call[bool](cc, objc.Sel("centerStageEnabled"))
	return rv
}

// A Boolean value that indicates whether a user or an app enabled Center Stage on a device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3738419-centerstageenabled?language=objc
func CaptureDevice_CenterStageEnabled() bool {
	return CaptureDeviceClass.CenterStageEnabled()
}

// A Boolean value that indicates whether a user or an app enabled Center Stage on a device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3738419-centerstageenabled?language=objc
func (cc _CaptureDeviceClass) SetCenterStageEnabled(value bool) {
	objc.Call[objc.Void](cc, objc.Sel("setCenterStageEnabled:"), value)
}

// A Boolean value that indicates whether a user or an app enabled Center Stage on a device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3738419-centerstageenabled?language=objc
func CaptureDevice_SetCenterStageEnabled(value bool) {
	CaptureDeviceClass.SetCenterStageEnabled(value)
}

// A model identifier for the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1389500-modelid?language=objc
func (c_ CaptureDevice) ModelID() string {
	rv := objc.Call[string](c_, objc.Sel("modelID"))
	return rv
}

// The microphone mode that the user selects in Control Center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3850458-preferredmicrophonemode?language=objc
func (cc _CaptureDeviceClass) PreferredMicrophoneMode() CaptureMicrophoneMode {
	rv := objc.Call[CaptureMicrophoneMode](cc, objc.Sel("preferredMicrophoneMode"))
	return rv
}

// The microphone mode that the user selects in Control Center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3850458-preferredmicrophonemode?language=objc
func CaptureDevice_PreferredMicrophoneMode() CaptureMicrophoneMode {
	return CaptureDeviceClass.PreferredMicrophoneMode()
}

// The current playback mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1386373-transportcontrolsplaybackmode?language=objc
func (c_ CaptureDevice) TransportControlsPlaybackMode() CaptureDeviceTransportControlsPlaybackMode {
	rv := objc.Call[CaptureDeviceTransportControlsPlaybackMode](c_, objc.Sel("transportControlsPlaybackMode"))
	return rv
}

// The transport type of the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1387804-transporttype?language=objc
func (c_ CaptureDevice) TransportType() int32 {
	rv := objc.Call[int32](c_, objc.Sel("transportType"))
	return rv
}

// The current torch brightness level. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1624605-torchlevel?language=objc
func (c_ CaptureDevice) TorchLevel() float64 {
	rv := objc.Call[float64](c_, objc.Sel("torchLevel"))
	return rv
}

// The currently active maximum frame duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1387816-activevideomaxframeduration?language=objc
func (c_ CaptureDevice) ActiveVideoMaxFrameDuration() coremedia.Time {
	rv := objc.Call[coremedia.Time](c_, objc.Sel("activeVideoMaxFrameDuration"))
	return rv
}

// The currently active maximum frame duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1387816-activevideomaxframeduration?language=objc
func (c_ CaptureDevice) SetActiveVideoMaxFrameDuration(value coremedia.Time) {
	objc.Call[objc.Void](c_, objc.Sel("setActiveVideoMaxFrameDuration:"), value)
}

// A Boolean value that indicates whether the capture device has a flash. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1388988-hasflash?language=objc
func (c_ CaptureDevice) HasFlash() bool {
	rv := objc.Call[bool](c_, objc.Sel("hasFlash"))
	return rv
}

// A Boolean value that indicates whether the device supports a point of interest for exposure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1387263-exposurepointofinterestsupported?language=objc
func (c_ CaptureDevice) IsExposurePointOfInterestSupported() bool {
	rv := objc.Call[bool](c_, objc.Sel("isExposurePointOfInterestSupported"))
	return rv
}

// The point of interest for focusing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1385853-focuspointofinterest?language=objc
func (c_ CaptureDevice) FocusPointOfInterest() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](c_, objc.Sel("focusPointOfInterest"))
	return rv
}

// The point of interest for focusing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1385853-focuspointofinterest?language=objc
func (c_ CaptureDevice) SetFocusPointOfInterest(value coregraphics.Point) {
	objc.Call[objc.Void](c_, objc.Sel("setFocusPointOfInterest:"), value)
}

// A Boolean value that indicates whether the device’s torch is currently active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1624578-torchactive?language=objc
func (c_ CaptureDevice) IsTorchActive() bool {
	rv := objc.Call[bool](c_, objc.Sel("isTorchActive"))
	return rv
}

// The conditions that restrict the primary constituent device’s switching behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3875308-primaryconstituentdevicerestrict?language=objc
func (c_ CaptureDevice) PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions {
	rv := objc.Call[CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions](c_, objc.Sel("primaryConstituentDeviceRestrictedSwitchingBehaviorConditions"))
	return rv
}

// The conditions that restrict camera switching behavior for the active primary constituent device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3875305-activeprimaryconstituentdevicere?language=objc
func (c_ CaptureDevice) ActivePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions {
	rv := objc.Call[CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions](c_, objc.Sel("activePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions"))
	return rv
}

// A Boolean value that indicates whether a device is currently connected to the system and available for use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1389343-connected?language=objc
func (c_ CaptureDevice) IsConnected() bool {
	rv := objc.Call[bool](c_, objc.Sel("isConnected"))
	return rv
}

// The device’s current flash mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1388116-flashmode?language=objc
func (c_ CaptureDevice) FlashMode() CaptureFlashMode {
	rv := objc.Call[CaptureFlashMode](c_, objc.Sel("flashMode"))
	return rv
}

// The device’s current flash mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1388116-flashmode?language=objc
func (c_ CaptureDevice) SetFlashMode(value CaptureFlashMode) {
	objc.Call[objc.Void](c_, objc.Sel("setFlashMode:"), value)
}

// The capture device’s minimum focus distance in millimeters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3751762-minimumfocusdistance?language=objc
func (c_ CaptureDevice) MinimumFocusDistance() int {
	rv := objc.Call[int](c_, objc.Sel("minimumFocusDistance"))
	return rv
}

// A Boolean value that indicates whether the device is currently adjusting its exposure setting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1386253-adjustingexposure?language=objc
func (c_ CaptureDevice) IsAdjustingExposure() bool {
	rv := objc.Call[bool](c_, objc.Sel("isAdjustingExposure"))
	return rv
}

// A localized device name for display in the user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1388222-localizedname?language=objc
func (c_ CaptureDevice) LocalizedName() string {
	rv := objc.Call[string](c_, objc.Sel("localizedName"))
	return rv
}

// The current white balance mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1386369-whitebalancemode?language=objc
func (c_ CaptureDevice) WhiteBalanceMode() CaptureWhiteBalanceMode {
	rv := objc.Call[CaptureWhiteBalanceMode](c_, objc.Sel("whiteBalanceMode"))
	return rv
}

// The current white balance mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1386369-whitebalancemode?language=objc
func (c_ CaptureDevice) SetWhiteBalanceMode(value CaptureWhiteBalanceMode) {
	objc.Call[objc.Void](c_, objc.Sel("setWhiteBalanceMode:"), value)
}

// An identifier that uniquely identifies the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1390477-uniqueid?language=objc
func (c_ CaptureDevice) UniqueID() string {
	rv := objc.Call[string](c_, objc.Sel("uniqueID"))
	return rv
}

// The switching behavior of the active constituent device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3875306-activeprimaryconstituentdevicesw?language=objc
func (c_ CaptureDevice) ActivePrimaryConstituentDeviceSwitchingBehavior() CapturePrimaryConstituentDeviceSwitchingBehavior {
	rv := objc.Call[CapturePrimaryConstituentDeviceSwitchingBehavior](c_, objc.Sel("activePrimaryConstituentDeviceSwitchingBehavior"))
	return rv
}

// A Boolean value that specifies whether the capture device has a torch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1387674-hastorch?language=objc
func (c_ CaptureDevice) HasTorch() bool {
	rv := objc.Call[bool](c_, objc.Sel("hasTorch"))
	return rv
}

// The physical position of the capture device hardware. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1386909-position?language=objc
func (c_ CaptureDevice) Position() CaptureDevicePosition {
	rv := objc.Call[CaptureDevicePosition](c_, objc.Sel("position"))
	return rv
}

// The currently active input source of the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1390141-activeinputsource?language=objc
func (c_ CaptureDevice) ActiveInputSource() CaptureDeviceInputSource {
	rv := objc.Call[CaptureDeviceInputSource](c_, objc.Sel("activeInputSource"))
	return rv
}

// The currently active input source of the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1390141-activeinputsource?language=objc
func (c_ CaptureDevice) SetActiveInputSource(value ICaptureDeviceInputSource) {
	objc.Call[objc.Void](c_, objc.Sel("setActiveInputSource:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the torch is currently available for use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1624626-torchavailable?language=objc
func (c_ CaptureDevice) IsTorchAvailable() bool {
	rv := objc.Call[bool](c_, objc.Sel("isTorchAvailable"))
	return rv
}

// The current playback speed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1386639-transportcontrolsspeed?language=objc
func (c_ CaptureDevice) TransportControlsSpeed() CaptureDeviceTransportControlsSpeed {
	rv := objc.Call[CaptureDeviceTransportControlsSpeed](c_, objc.Sel("transportControlsSpeed"))
	return rv
}

// The constituent devices available to select as a fallback for a longer focal length primary constituent device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3875311-supportedfallbackprimaryconstitu?language=objc
func (c_ CaptureDevice) SupportedFallbackPrimaryConstituentDevices() []CaptureDevice {
	rv := objc.Call[[]CaptureDevice](c_, objc.Sel("supportedFallbackPrimaryConstituentDevices"))
	return rv
}

// A Boolean value that indicates whether the device is in a suspended state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1387761-suspended?language=objc
func (c_ CaptureDevice) IsSuspended() bool {
	rv := objc.Call[bool](c_, objc.Sel("isSuspended"))
	return rv
}

// A human-readable string for the manufacturer of the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1390193-manufacturer?language=objc
func (c_ CaptureDevice) Manufacturer() string {
	rv := objc.Call[string](c_, objc.Sel("manufacturer"))
	return rv
}

// The current torch mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1386035-torchmode?language=objc
func (c_ CaptureDevice) TorchMode() CaptureTorchMode {
	rv := objc.Call[CaptureTorchMode](c_, objc.Sel("torchMode"))
	return rv
}

// The current torch mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1386035-torchmode?language=objc
func (c_ CaptureDevice) SetTorchMode(value CaptureTorchMode) {
	objc.Call[objc.Void](c_, objc.Sel("setTorchMode:"), value)
}

// The switching behavior for the primary constituent device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3875309-primaryconstituentdeviceswitchin?language=objc
func (c_ CaptureDevice) PrimaryConstituentDeviceSwitchingBehavior() CapturePrimaryConstituentDeviceSwitchingBehavior {
	rv := objc.Call[CapturePrimaryConstituentDeviceSwitchingBehavior](c_, objc.Sel("primaryConstituentDeviceSwitchingBehavior"))
	return rv
}

// The currently active minimum frame duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1389290-activevideominframeduration?language=objc
func (c_ CaptureDevice) ActiveVideoMinFrameDuration() coremedia.Time {
	rv := objc.Call[coremedia.Time](c_, objc.Sel("activeVideoMinFrameDuration"))
	return rv
}

// The currently active minimum frame duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1389290-activevideominframeduration?language=objc
func (c_ CaptureDevice) SetActiveVideoMinFrameDuration(value coremedia.Time) {
	objc.Call[objc.Void](c_, objc.Sel("setActiveVideoMinFrameDuration:"), value)
}

// A Boolean value that indicates whether the flash is currently available for use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1624627-flashavailable?language=objc
func (c_ CaptureDevice) IsFlashAvailable() bool {
	rv := objc.Call[bool](c_, objc.Sel("isFlashAvailable"))
	return rv
}

// A Boolean value that indicates whether the Portrait video effect is active on a device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3850456-portraiteffectactive?language=objc
func (c_ CaptureDevice) IsPortraitEffectActive() bool {
	rv := objc.Call[bool](c_, objc.Sel("isPortraitEffectActive"))
	return rv
}

// The capture device’s focus mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1389191-focusmode?language=objc
func (c_ CaptureDevice) FocusMode() CaptureFocusMode {
	rv := objc.Call[CaptureFocusMode](c_, objc.Sel("focusMode"))
	return rv
}

// The capture device’s focus mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1389191-focusmode?language=objc
func (c_ CaptureDevice) SetFocusMode(value CaptureFocusMode) {
	objc.Call[objc.Void](c_, objc.Sel("setFocusMode:"), value)
}

// An array of input sources that the device supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1388404-inputsources?language=objc
func (c_ CaptureDevice) InputSources() []CaptureDeviceInputSource {
	rv := objc.Call[[]CaptureDeviceInputSource](c_, objc.Sel("inputSources"))
	return rv
}

// The exposure mode for the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1388858-exposuremode?language=objc
func (c_ CaptureDevice) ExposureMode() CaptureExposureMode {
	rv := objc.Call[CaptureExposureMode](c_, objc.Sel("exposureMode"))
	return rv
}

// The exposure mode for the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1388858-exposuremode?language=objc
func (c_ CaptureDevice) SetExposureMode(value CaptureExposureMode) {
	objc.Call[objc.Void](c_, objc.Sel("setExposureMode:"), value)
}

// A Boolean value that indicates whether the user enabled the Portrait video effect in Control Center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3850457-portraiteffectenabled?language=objc
func (cc _CaptureDeviceClass) PortraitEffectEnabled() bool {
	rv := objc.Call[bool](cc, objc.Sel("portraitEffectEnabled"))
	return rv
}

// A Boolean value that indicates whether the user enabled the Portrait video effect in Control Center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3850457-portraiteffectenabled?language=objc
func CaptureDevice_PortraitEffectEnabled() bool {
	return CaptureDeviceClass.PortraitEffectEnabled()
}

// A Boolean value that indicates whether the device supports a point of interest for focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1390436-focuspointofinterestsupported?language=objc
func (c_ CaptureDevice) IsFocusPointOfInterestSupported() bool {
	rv := objc.Call[bool](c_, objc.Sel("isFocusPointOfInterestSupported"))
	return rv
}

// A Boolean value that indicates whether the device supports transport control commands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1388479-transportcontrolssupported?language=objc
func (c_ CaptureDevice) TransportControlsSupported() bool {
	rv := objc.Call[bool](c_, objc.Sel("transportControlsSupported"))
	return rv
}

// A virtual device’s active primary constituent device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3875304-activeprimaryconstituentdevice?language=objc
func (c_ CaptureDevice) ActivePrimaryConstituentDevice() CaptureDevice {
	rv := objc.Call[CaptureDevice](c_, objc.Sel("activePrimaryConstituentDevice"))
	return rv
}

// An array of capture devices that are physically linked to a device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1388720-linkeddevices?language=objc
func (c_ CaptureDevice) LinkedDevices() []CaptureDevice {
	rv := objc.Call[[]CaptureDevice](c_, objc.Sel("linkedDevices"))
	return rv
}

// A Boolean value that indicates whether the device is currently adjusting the white balance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1386544-adjustingwhitebalance?language=objc
func (c_ CaptureDevice) IsAdjustingWhiteBalance() bool {
	rv := objc.Call[bool](c_, objc.Sel("isAdjustingWhiteBalance"))
	return rv
}

// The type of device, such as a built-in microphone or wide-angle camera. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/2361119-devicetype?language=objc
func (c_ CaptureDevice) DeviceType() CaptureDeviceType {
	rv := objc.Call[CaptureDeviceType](c_, objc.Sel("deviceType"))
	return rv
}

// A Boolean value that indicates whether the device is currently adjusting its focus setting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1390577-adjustingfocus?language=objc
func (c_ CaptureDevice) IsAdjustingFocus() bool {
	rv := objc.Call[bool](c_, objc.Sel("isAdjustingFocus"))
	return rv
}

// The point of interest for exposure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1388777-exposurepointofinterest?language=objc
func (c_ CaptureDevice) ExposurePointOfInterest() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](c_, objc.Sel("exposurePointOfInterest"))
	return rv
}

// The point of interest for exposure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1388777-exposurepointofinterest?language=objc
func (c_ CaptureDevice) SetExposurePointOfInterest(value coregraphics.Point) {
	objc.Call[objc.Void](c_, objc.Sel("setExposurePointOfInterest:"), value)
}

// A value that indicates the current mode of Center Stage control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3738418-centerstagecontrolmode?language=objc
func (cc _CaptureDeviceClass) CenterStageControlMode() CaptureCenterStageControlMode {
	rv := objc.Call[CaptureCenterStageControlMode](cc, objc.Sel("centerStageControlMode"))
	return rv
}

// A value that indicates the current mode of Center Stage control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3738418-centerstagecontrolmode?language=objc
func CaptureDevice_CenterStageControlMode() CaptureCenterStageControlMode {
	return CaptureDeviceClass.CenterStageControlMode()
}

// A value that indicates the current mode of Center Stage control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3738418-centerstagecontrolmode?language=objc
func (cc _CaptureDeviceClass) SetCenterStageControlMode(value CaptureCenterStageControlMode) {
	objc.Call[objc.Void](cc, objc.Sel("setCenterStageControlMode:"), value)
}

// A value that indicates the current mode of Center Stage control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3738418-centerstagecontrolmode?language=objc
func CaptureDevice_SetCenterStageControlMode(value CaptureCenterStageControlMode) {
	CaptureDeviceClass.SetCenterStageControlMode(value)
}

// The currently active color space for capture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1648668-activecolorspace?language=objc
func (c_ CaptureDevice) ActiveColorSpace() CaptureColorSpace {
	rv := objc.Call[CaptureColorSpace](c_, objc.Sel("activeColorSpace"))
	return rv
}

// The currently active color space for capture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1648668-activecolorspace?language=objc
func (c_ CaptureDevice) SetActiveColorSpace(value CaptureColorSpace) {
	objc.Call[objc.Void](c_, objc.Sel("setActiveColorSpace:"), value)
}

// The fallback devices to use when a constituent device with a longer focal length becomes limited by its light sensitivity or minimum focus distance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3875307-fallbackprimaryconstituentdevice?language=objc
func (c_ CaptureDevice) FallbackPrimaryConstituentDevices() []CaptureDevice {
	rv := objc.Call[[]CaptureDevice](c_, objc.Sel("fallbackPrimaryConstituentDevices"))
	return rv
}

// The fallback devices to use when a constituent device with a longer focal length becomes limited by its light sensitivity or minimum focus distance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3875307-fallbackprimaryconstituentdevice?language=objc
func (c_ CaptureDevice) SetFallbackPrimaryConstituentDevices(value []ICaptureDevice) {
	objc.Call[objc.Void](c_, objc.Sel("setFallbackPrimaryConstituentDevices:"), value)
}

// The capture format in use by the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1389221-activeformat?language=objc
func (c_ CaptureDevice) ActiveFormat() CaptureDeviceFormat {
	rv := objc.Call[CaptureDeviceFormat](c_, objc.Sel("activeFormat"))
	return rv
}

// The capture format in use by the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/1389221-activeformat?language=objc
func (c_ CaptureDevice) SetActiveFormat(value ICaptureDeviceFormat) {
	objc.Call[objc.Void](c_, objc.Sel("setActiveFormat:"), objc.Ptr(value))
}

// The device’s active microphone mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3850455-activemicrophonemode?language=objc
func (cc _CaptureDeviceClass) ActiveMicrophoneMode() CaptureMicrophoneMode {
	rv := objc.Call[CaptureMicrophoneMode](cc, objc.Sel("activeMicrophoneMode"))
	return rv
}

// The device’s active microphone mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3850455-activemicrophonemode?language=objc
func CaptureDevice_ActiveMicrophoneMode() CaptureMicrophoneMode {
	return CaptureDeviceClass.ActiveMicrophoneMode()
}

// A Boolean value that indicates whether Center Stage is active on a device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/3738417-centerstageactive?language=objc
func (c_ CaptureDevice) IsCenterStageActive() bool {
	rv := objc.Call[bool](c_, objc.Sel("isCenterStageActive"))
	return rv
}