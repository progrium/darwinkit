// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlayerItem] class.
var PlayerItemClass = _PlayerItemClass{objc.GetClass("AVPlayerItem")}

type _PlayerItemClass struct {
	objc.Class
}

// An interface definition for the [PlayerItem] class.
type IPlayerItem interface {
	objc.IObject
	SelectMediaOptionAutomaticallyInMediaSelectionGroup(mediaSelectionGroup IMediaSelectionGroup)
	ErrorLog() PlayerItemErrorLog
	CancelPendingSeeks()
	Copy() objc.Object
	RemoveMediaDataCollector(collector IPlayerItemMediaDataCollector)
	CurrentDate() foundation.Date
	AddOutput(output IPlayerItemOutput)
	AddMediaDataCollector(collector IPlayerItemMediaDataCollector)
	CancelContentAuthorizationRequest()
	StepByCount(stepCount int)
	RemoveOutput(output IPlayerItemOutput)
	RequestContentAuthorizationAsynchronouslyWithTimeoutIntervalCompletionHandler(timeoutInterval foundation.TimeInterval, handler func())
	SelectMediaOptionInMediaSelectionGroup(mediaSelectionOption IMediaSelectionOption, mediaSelectionGroup IMediaSelectionGroup)
	CopyWithZone(zone unsafe.Pointer) objc.Object
	AccessLog() PlayerItemAccessLog
	CurrentTime() coremedia.Time
	AllowedAudioSpatializationFormats() AudioSpatializationFormats
	SetAllowedAudioSpatializationFormats(value AudioSpatializationFormats)
	CanPlayFastReverse() bool
	CanStepForward() bool
	Error() foundation.Error
	Tracks() []PlayerItemTrack
	CanPlaySlowReverse() bool
	IsPlaybackBufferEmpty() bool
	VideoComposition() VideoComposition
	SetVideoComposition(value IVideoComposition)
	IsAuthorizationRequiredForPlayback() bool
	CanPlayReverse() bool
	IsPlaybackLikelyToKeepUp() bool
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm
	SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm)
	ForwardPlaybackEndTime() coremedia.Time
	SetForwardPlaybackEndTime(value coremedia.Time)
	TextStyleRules() []TextStyleRule
	SetTextStyleRules(value []ITextStyleRule)
	PreferredPeakBitRateForExpensiveNetworks() float64
	SetPreferredPeakBitRateForExpensiveNetworks(value float64)
	PreferredPeakBitRate() float64
	SetPreferredPeakBitRate(value float64)
	VideoApertureMode() VideoApertureMode
	SetVideoApertureMode(value VideoApertureMode)
	TemplatePlayerItem() PlayerItem
	ReversePlaybackEndTime() coremedia.Time
	SetReversePlaybackEndTime(value coremedia.Time)
	Outputs() []PlayerItemOutput
	Timebase() coremedia.TimebaseRef
	RecommendedTimeOffsetFromLive() coremedia.Time
	CanUseNetworkResourcesForLiveStreamingWhilePaused() bool
	SetCanUseNetworkResourcesForLiveStreamingWhilePaused(value bool)
	CanStepBackward() bool
	ContentAuthorizationRequestStatus() ContentAuthorizationStatus
	CustomVideoCompositor() VideoCompositingWrapper
	PresentationSize() coregraphics.Size
	ConfiguredTimeOffsetFromLive() coremedia.Time
	SetConfiguredTimeOffsetFromLive(value coremedia.Time)
	SeekingWaitsForVideoCompositionRendering() bool
	SetSeekingWaitsForVideoCompositionRendering(value bool)
	LoadedTimeRanges() []foundation.Value
	AppliesPerFrameHDRDisplayMetadata() bool
	SetAppliesPerFrameHDRDisplayMetadata(value bool)
	CurrentMediaSelection() MediaSelection
	AutomaticallyHandlesInterstitialEvents() bool
	SetAutomaticallyHandlesInterstitialEvents(value bool)
	AutomaticallyPreservesTimeOffsetFromLive() bool
	SetAutomaticallyPreservesTimeOffsetFromLive(value bool)
	IsApplicationAuthorizedForPlayback() bool
	PreferredMaximumResolutionForExpensiveNetworks() coregraphics.Size
	SetPreferredMaximumResolutionForExpensiveNetworks(value coregraphics.Size)
	IsPlaybackBufferFull() bool
	IsContentAuthorizedForPlayback() bool
	AudioMix() AudioMix
	SetAudioMix(value IAudioMix)
	SeekableTimeRanges() []foundation.Value
	PreferredMaximumResolution() coregraphics.Size
	SetPreferredMaximumResolution(value coregraphics.Size)
	VariantPreferences() VariantPreferences
	SetVariantPreferences(value VariantPreferences)
	CanPlaySlowForward() bool
	AutomaticallyLoadedAssetKeys() []string
	Duration() coremedia.Time
	PreferredForwardBufferDuration() foundation.TimeInterval
	SetPreferredForwardBufferDuration(value foundation.TimeInterval)
	Status() PlayerItemStatus
	MediaDataCollectors() []PlayerItemMediaDataCollector
	StartsOnFirstEligibleVariant() bool
	SetStartsOnFirstEligibleVariant(value bool)
	Asset() Asset
	CanPlayFastForward() bool
}

// An object that models the timing and presentation state of an asset during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem?language=objc
type PlayerItem struct {
	objc.Object
}

func PlayerItemFrom(ptr unsafe.Pointer) PlayerItem {
	return PlayerItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PlayerItemClass) PlayerItemWithAsset(asset IAsset) PlayerItem {
	rv := objc.Call[PlayerItem](pc, objc.Sel("playerItemWithAsset:"), objc.Ptr(asset))
	return rv
}

// Returns a new player item for a specified asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1588087-playeritemwithasset?language=objc
func PlayerItem_PlayerItemWithAsset(asset IAsset) PlayerItem {
	return PlayerItemClass.PlayerItemWithAsset(asset)
}

func (p_ PlayerItem) InitWithURL(URL foundation.IURL) PlayerItem {
	rv := objc.Call[PlayerItem](p_, objc.Sel("initWithURL:"), objc.Ptr(URL))
	return rv
}

// Creates a player item with a specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1387558-initwithurl?language=objc
func NewPlayerItemWithURL(URL foundation.IURL) PlayerItem {
	instance := PlayerItemClass.Alloc().InitWithURL(URL)
	instance.Autorelease()
	return instance
}

func (pc _PlayerItemClass) PlayerItemWithURL(URL foundation.IURL) PlayerItem {
	rv := objc.Call[PlayerItem](pc, objc.Sel("playerItemWithURL:"), objc.Ptr(URL))
	return rv
}

// Returns a new player item with a specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1588089-playeritemwithurl?language=objc
func PlayerItem_PlayerItemWithURL(URL foundation.IURL) PlayerItem {
	return PlayerItemClass.PlayerItemWithURL(URL)
}

func (p_ PlayerItem) InitWithAsset(asset IAsset) PlayerItem {
	rv := objc.Call[PlayerItem](p_, objc.Sel("initWithAsset:"), objc.Ptr(asset))
	return rv
}

// Creates a player item for a specified asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1390707-initwithasset?language=objc
func NewPlayerItemWithAsset(asset IAsset) PlayerItem {
	instance := PlayerItemClass.Alloc().InitWithAsset(asset)
	instance.Autorelease()
	return instance
}

func (pc _PlayerItemClass) Alloc() PlayerItem {
	rv := objc.Call[PlayerItem](pc, objc.Sel("alloc"))
	return rv
}

func PlayerItem_Alloc() PlayerItem {
	return PlayerItemClass.Alloc()
}

func (pc _PlayerItemClass) New() PlayerItem {
	rv := objc.Call[PlayerItem](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerItem() PlayerItem {
	return PlayerItemClass.New()
}

func (p_ PlayerItem) Init() PlayerItem {
	rv := objc.Call[PlayerItem](p_, objc.Sel("init"))
	return rv
}

// Selects the media option in the specified media selection group that best matches the receiver’s automatic selection criteria. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388268-selectmediaoptionautomaticallyin?language=objc
func (p_ PlayerItem) SelectMediaOptionAutomaticallyInMediaSelectionGroup(mediaSelectionGroup IMediaSelectionGroup) {
	objc.Call[objc.Void](p_, objc.Sel("selectMediaOptionAutomaticallyInMediaSelectionGroup:"), objc.Ptr(mediaSelectionGroup))
}

// Returns an object that represents a snapshot of the error log. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1387573-errorlog?language=objc
func (p_ PlayerItem) ErrorLog() PlayerItemErrorLog {
	rv := objc.Call[PlayerItemErrorLog](p_, objc.Sel("errorLog"))
	return rv
}

// Cancels any pending seek requests and invokes the corresponding completion handlers if present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388316-cancelpendingseeks?language=objc
func (p_ PlayerItem) CancelPendingSeeks() {
	objc.Call[objc.Void](p_, objc.Sel("cancelPendingSeeks"))
}

// Creates a copy of the object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3861797-copy?language=objc
func (p_ PlayerItem) Copy() objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("copy"))
	rv.Autorelease()
	return rv
}

// Removes the specified media data collector from the player item’s collection of media collectors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1624163-removemediadatacollector?language=objc
func (p_ PlayerItem) RemoveMediaDataCollector(collector IPlayerItemMediaDataCollector) {
	objc.Call[objc.Void](p_, objc.Sel("removeMediaDataCollector:"), objc.Ptr(collector))
}

// Returns the current time of the item as a date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1386188-currentdate?language=objc
func (p_ PlayerItem) CurrentDate() foundation.Date {
	rv := objc.Call[foundation.Date](p_, objc.Sel("currentDate"))
	return rv
}

// Adds the specified player item output object to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1389782-addoutput?language=objc
func (p_ PlayerItem) AddOutput(output IPlayerItemOutput) {
	objc.Call[objc.Void](p_, objc.Sel("addOutput:"), objc.Ptr(output))
}

// Adds the specified media data collector to the player item’s collection of media collectors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1624164-addmediadatacollector?language=objc
func (p_ PlayerItem) AddMediaDataCollector(collector IPlayerItemMediaDataCollector) {
	objc.Call[objc.Void](p_, objc.Sel("addMediaDataCollector:"), objc.Ptr(collector))
}

// Cancels the currently outstanding content authorization request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1387726-cancelcontentauthorizationreques?language=objc
func (p_ PlayerItem) CancelContentAuthorizationRequest() {
	objc.Call[objc.Void](p_, objc.Sel("cancelContentAuthorizationRequest"))
}

// Moves the player item’s current time forward or backward by a specified number of steps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1387968-stepbycount?language=objc
func (p_ PlayerItem) StepByCount(stepCount int) {
	objc.Call[objc.Void](p_, objc.Sel("stepByCount:"), stepCount)
}

// Removes the specified player item output object from the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388756-removeoutput?language=objc
func (p_ PlayerItem) RemoveOutput(output IPlayerItemOutput) {
	objc.Call[objc.Void](p_, objc.Sel("removeOutput:"), objc.Ptr(output))
}

// Presents the user the opportunity to authorize the content for playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1390600-requestcontentauthorizationasync?language=objc
func (p_ PlayerItem) RequestContentAuthorizationAsynchronouslyWithTimeoutIntervalCompletionHandler(timeoutInterval foundation.TimeInterval, handler func()) {
	objc.Call[objc.Void](p_, objc.Sel("requestContentAuthorizationAsynchronouslyWithTimeoutInterval:completionHandler:"), timeoutInterval, handler)
}

// Selects a media option in a given media selection group and deselects all other options in that group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1389610-selectmediaoption?language=objc
func (p_ PlayerItem) SelectMediaOptionInMediaSelectionGroup(mediaSelectionOption IMediaSelectionOption, mediaSelectionGroup IMediaSelectionGroup) {
	objc.Call[objc.Void](p_, objc.Sel("selectMediaOption:inMediaSelectionGroup:"), objc.Ptr(mediaSelectionOption), objc.Ptr(mediaSelectionGroup))
}

// Creates a copy of the object with the specified zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3861798-copywithzone?language=objc
func (p_ PlayerItem) CopyWithZone(zone unsafe.Pointer) objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("copyWithZone:"), zone)
	return rv
}

// Returns an object that represents a snapshot of the network access log. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388499-accesslog?language=objc
func (p_ PlayerItem) AccessLog() PlayerItemAccessLog {
	rv := objc.Call[PlayerItemAccessLog](p_, objc.Sel("accessLog"))
	return rv
}

// Returns the current time of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1387230-currenttime?language=objc
func (p_ PlayerItem) CurrentTime() coremedia.Time {
	rv := objc.Call[coremedia.Time](p_, objc.Sel("currentTime"))
	return rv
}

// The source audio channel layouts the player item supports for spatialization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3601108-allowedaudiospatializationformat?language=objc
func (p_ PlayerItem) AllowedAudioSpatializationFormats() AudioSpatializationFormats {
	rv := objc.Call[AudioSpatializationFormats](p_, objc.Sel("allowedAudioSpatializationFormats"))
	return rv
}

// The source audio channel layouts the player item supports for spatialization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3601108-allowedaudiospatializationformat?language=objc
func (p_ PlayerItem) SetAllowedAudioSpatializationFormats(value AudioSpatializationFormats) {
	objc.Call[objc.Void](p_, objc.Sel("setAllowedAudioSpatializationFormats:"), value)
}

// A Boolean value that indicates whether the item can be quickly reversed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1390493-canplayfastreverse?language=objc
func (p_ PlayerItem) CanPlayFastReverse() bool {
	rv := objc.Call[bool](p_, objc.Sel("canPlayFastReverse"))
	return rv
}

// A Boolean value that indicates whether the item supports stepping forward. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1389656-canstepforward?language=objc
func (p_ PlayerItem) CanStepForward() bool {
	rv := objc.Call[bool](p_, objc.Sel("canStepForward"))
	return rv
}

// The error that caused the player item to fail. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1389185-error?language=objc
func (p_ PlayerItem) Error() foundation.Error {
	rv := objc.Call[foundation.Error](p_, objc.Sel("error"))
	return rv
}

// An array of player item track objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1386361-tracks?language=objc
func (p_ PlayerItem) Tracks() []PlayerItemTrack {
	rv := objc.Call[[]PlayerItemTrack](p_, objc.Sel("tracks"))
	return rv
}

// A Boolean value that indicates whether the item can play slowly backward. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1390598-canplayslowreverse?language=objc
func (p_ PlayerItem) CanPlaySlowReverse() bool {
	rv := objc.Call[bool](p_, objc.Sel("canPlaySlowReverse"))
	return rv
}

// A Boolean value that indicates whether playback has consumed all buffered media and that playback will stall or end. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1386960-playbackbufferempty?language=objc
func (p_ PlayerItem) IsPlaybackBufferEmpty() bool {
	rv := objc.Call[bool](p_, objc.Sel("isPlaybackBufferEmpty"))
	return rv
}

// The video composition settings to be applied during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388818-videocomposition?language=objc
func (p_ PlayerItem) VideoComposition() VideoComposition {
	rv := objc.Call[VideoComposition](p_, objc.Sel("videoComposition"))
	return rv
}

// The video composition settings to be applied during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388818-videocomposition?language=objc
func (p_ PlayerItem) SetVideoComposition(value IVideoComposition) {
	objc.Call[objc.Void](p_, objc.Sel("setVideoComposition:"), objc.Ptr(value))
}

// A Boolean value that indicates whether authorization is required to play the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1386106-authorizationrequiredforplayback?language=objc
func (p_ PlayerItem) IsAuthorizationRequiredForPlayback() bool {
	rv := objc.Call[bool](p_, objc.Sel("isAuthorizationRequiredForPlayback"))
	return rv
}

// A Boolean value that indicates whether the item can play in reverse. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1385591-canplayreverse?language=objc
func (p_ PlayerItem) CanPlayReverse() bool {
	rv := objc.Call[bool](p_, objc.Sel("canPlayReverse"))
	return rv
}

// A Boolean value that indicates whether the item will likely play through without stalling. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1390348-playbacklikelytokeepup?language=objc
func (p_ PlayerItem) IsPlaybackLikelyToKeepUp() bool {
	rv := objc.Call[bool](p_, objc.Sel("isPlaybackLikelyToKeepUp"))
	return rv
}

// The processing algorithm used to manage audio pitch for scaled audio edits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1385855-audiotimepitchalgorithm?language=objc
func (p_ PlayerItem) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm {
	rv := objc.Call[AudioTimePitchAlgorithm](p_, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}

// The processing algorithm used to manage audio pitch for scaled audio edits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1385855-audiotimepitchalgorithm?language=objc
func (p_ PlayerItem) SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm) {
	objc.Call[objc.Void](p_, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}

// The time at which forward playback ends. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1385622-forwardplaybackendtime?language=objc
func (p_ PlayerItem) ForwardPlaybackEndTime() coremedia.Time {
	rv := objc.Call[coremedia.Time](p_, objc.Sel("forwardPlaybackEndTime"))
	return rv
}

// The time at which forward playback ends. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1385622-forwardplaybackendtime?language=objc
func (p_ PlayerItem) SetForwardPlaybackEndTime(value coremedia.Time) {
	objc.Call[objc.Void](p_, objc.Sel("setForwardPlaybackEndTime:"), value)
}

// An array of text style rules that specify the formatting and presentation of Web Video Text Tracks (WebVTT) subtitles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1389681-textstylerules?language=objc
func (p_ PlayerItem) TextStyleRules() []TextStyleRule {
	rv := objc.Call[[]TextStyleRule](p_, objc.Sel("textStyleRules"))
	return rv
}

// An array of text style rules that specify the formatting and presentation of Web Video Text Tracks (WebVTT) subtitles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1389681-textstylerules?language=objc
func (p_ PlayerItem) SetTextStyleRules(value []ITextStyleRule) {
	objc.Call[objc.Void](p_, objc.Sel("setTextStyleRules:"), value)
}

// A limit of network bandwidth consumption by the item when connecting over expensive networks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3746589-preferredpeakbitrateforexpensive?language=objc
func (p_ PlayerItem) PreferredPeakBitRateForExpensiveNetworks() float64 {
	rv := objc.Call[float64](p_, objc.Sel("preferredPeakBitRateForExpensiveNetworks"))
	return rv
}

// A limit of network bandwidth consumption by the item when connecting over expensive networks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3746589-preferredpeakbitrateforexpensive?language=objc
func (p_ PlayerItem) SetPreferredPeakBitRateForExpensiveNetworks(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setPreferredPeakBitRateForExpensiveNetworks:"), value)
}

// The desired limit, in bits per second, of network bandwidth consumption for this item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388541-preferredpeakbitrate?language=objc
func (p_ PlayerItem) PreferredPeakBitRate() float64 {
	rv := objc.Call[float64](p_, objc.Sel("preferredPeakBitRate"))
	return rv
}

// The desired limit, in bits per second, of network bandwidth consumption for this item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388541-preferredpeakbitrate?language=objc
func (p_ PlayerItem) SetPreferredPeakBitRate(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setPreferredPeakBitRate:"), value)
}

// The video aperture mode to apply during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/2868499-videoaperturemode?language=objc
func (p_ PlayerItem) VideoApertureMode() VideoApertureMode {
	rv := objc.Call[VideoApertureMode](p_, objc.Sel("videoApertureMode"))
	return rv
}

// The video aperture mode to apply during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/2868499-videoaperturemode?language=objc
func (p_ PlayerItem) SetVideoApertureMode(value VideoApertureMode) {
	objc.Call[objc.Void](p_, objc.Sel("setVideoApertureMode:"), value)
}

// The template player item that initializes this instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3726147-templateplayeritem?language=objc
func (p_ PlayerItem) TemplatePlayerItem() PlayerItem {
	rv := objc.Call[PlayerItem](p_, objc.Sel("templatePlayerItem"))
	return rv
}

// The time at which reverse playback ends. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388438-reverseplaybackendtime?language=objc
func (p_ PlayerItem) ReversePlaybackEndTime() coremedia.Time {
	rv := objc.Call[coremedia.Time](p_, objc.Sel("reversePlaybackEndTime"))
	return rv
}

// The time at which reverse playback ends. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388438-reverseplaybackendtime?language=objc
func (p_ PlayerItem) SetReversePlaybackEndTime(value coremedia.Time) {
	objc.Call[objc.Void](p_, objc.Sel("setReversePlaybackEndTime:"), value)
}

// An array of outputs associated with the player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1389090-outputs?language=objc
func (p_ PlayerItem) Outputs() []PlayerItemOutput {
	rv := objc.Call[[]PlayerItemOutput](p_, objc.Sel("outputs"))
	return rv
}

// The timebase information for the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1387605-timebase?language=objc
func (p_ PlayerItem) Timebase() coremedia.TimebaseRef {
	rv := objc.Call[coremedia.TimebaseRef](p_, objc.Sel("timebase"))
	return rv
}

// A recommended time offset from the live time based on observed network conditions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3227883-recommendedtimeoffsetfromlive?language=objc
func (p_ PlayerItem) RecommendedTimeOffsetFromLive() coremedia.Time {
	rv := objc.Call[coremedia.Time](p_, objc.Sel("recommendedTimeOffsetFromLive"))
	return rv
}

// A Boolean value that indicates whether the player item can use network resources to keep the playback state up to date while paused. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388752-canusenetworkresourcesforlivestr?language=objc
func (p_ PlayerItem) CanUseNetworkResourcesForLiveStreamingWhilePaused() bool {
	rv := objc.Call[bool](p_, objc.Sel("canUseNetworkResourcesForLiveStreamingWhilePaused"))
	return rv
}

// A Boolean value that indicates whether the player item can use network resources to keep the playback state up to date while paused. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388752-canusenetworkresourcesforlivestr?language=objc
func (p_ PlayerItem) SetCanUseNetworkResourcesForLiveStreamingWhilePaused(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setCanUseNetworkResourcesForLiveStreamingWhilePaused:"), value)
}

// A Boolean value that indicates whether the item supports stepping backward. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1386796-canstepbackward?language=objc
func (p_ PlayerItem) CanStepBackward() bool {
	rv := objc.Call[bool](p_, objc.Sel("canStepBackward"))
	return rv
}

// The status of the most recent content authorization request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1389746-contentauthorizationrequeststatu?language=objc
func (p_ PlayerItem) ContentAuthorizationRequestStatus() ContentAuthorizationStatus {
	rv := objc.Call[ContentAuthorizationStatus](p_, objc.Sel("contentAuthorizationRequestStatus"))
	return rv
}

// The custom video compositor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1390669-customvideocompositor?language=objc
func (p_ PlayerItem) CustomVideoCompositor() VideoCompositingWrapper {
	rv := objc.Call[VideoCompositingWrapper](p_, objc.Sel("customVideoCompositor"))
	return rv
}

// The size at which the visual portion of the item is presented by the player. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388962-presentationsize?language=objc
func (p_ PlayerItem) PresentationSize() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](p_, objc.Sel("presentationSize"))
	return rv
}

// A time value that indicates the offset from the live time to start playback, or resume playback after a seek to positive infinity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3227882-configuredtimeoffsetfromlive?language=objc
func (p_ PlayerItem) ConfiguredTimeOffsetFromLive() coremedia.Time {
	rv := objc.Call[coremedia.Time](p_, objc.Sel("configuredTimeOffsetFromLive"))
	return rv
}

// A time value that indicates the offset from the live time to start playback, or resume playback after a seek to positive infinity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3227882-configuredtimeoffsetfromlive?language=objc
func (p_ PlayerItem) SetConfiguredTimeOffsetFromLive(value coremedia.Time) {
	objc.Call[objc.Void](p_, objc.Sel("setConfiguredTimeOffsetFromLive:"), value)
}

// A Boolean value that indicates whether the item’s timing follows the displayed video frame when seeking with a video composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1385726-seekingwaitsforvideocompositionr?language=objc
func (p_ PlayerItem) SeekingWaitsForVideoCompositionRendering() bool {
	rv := objc.Call[bool](p_, objc.Sel("seekingWaitsForVideoCompositionRendering"))
	return rv
}

// A Boolean value that indicates whether the item’s timing follows the displayed video frame when seeking with a video composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1385726-seekingwaitsforvideocompositionr?language=objc
func (p_ PlayerItem) SetSeekingWaitsForVideoCompositionRendering(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setSeekingWaitsForVideoCompositionRendering:"), value)
}

// An array of time ranges indicating media data that is readily available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1389953-loadedtimeranges?language=objc
func (p_ PlayerItem) LoadedTimeRanges() []foundation.Value {
	rv := objc.Call[[]foundation.Value](p_, objc.Sel("loadedTimeRanges"))
	return rv
}

// A Boolean value that indicates whether the player item applies per-frame HDR display metadata during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3656127-appliesperframehdrdisplaymetadat?language=objc
func (p_ PlayerItem) AppliesPerFrameHDRDisplayMetadata() bool {
	rv := objc.Call[bool](p_, objc.Sel("appliesPerFrameHDRDisplayMetadata"))
	return rv
}

// A Boolean value that indicates whether the player item applies per-frame HDR display metadata during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3656127-appliesperframehdrdisplaymetadat?language=objc
func (p_ PlayerItem) SetAppliesPerFrameHDRDisplayMetadata(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setAppliesPerFrameHDRDisplayMetadata:"), value)
}

// The current media selections for each of the receiver's media selection groups. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1386519-currentmediaselection?language=objc
func (p_ PlayerItem) CurrentMediaSelection() MediaSelection {
	rv := objc.Call[MediaSelection](p_, objc.Sel("currentMediaSelection"))
	return rv
}

// A Boolean value that indicates whether the player item automatically plays interstitial events according to server-side directives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3726146-automaticallyhandlesinterstitial?language=objc
func (p_ PlayerItem) AutomaticallyHandlesInterstitialEvents() bool {
	rv := objc.Call[bool](p_, objc.Sel("automaticallyHandlesInterstitialEvents"))
	return rv
}

// A Boolean value that indicates whether the player item automatically plays interstitial events according to server-side directives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3726146-automaticallyhandlesinterstitial?language=objc
func (p_ PlayerItem) SetAutomaticallyHandlesInterstitialEvents(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setAutomaticallyHandlesInterstitialEvents:"), value)
}

// A Boolean value that indicates whether the player preserves its time offset from the live time after a buffering operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3229855-automaticallypreservestimeoffset?language=objc
func (p_ PlayerItem) AutomaticallyPreservesTimeOffsetFromLive() bool {
	rv := objc.Call[bool](p_, objc.Sel("automaticallyPreservesTimeOffsetFromLive"))
	return rv
}

// A Boolean value that indicates whether the player preserves its time offset from the live time after a buffering operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3229855-automaticallypreservestimeoffset?language=objc
func (p_ PlayerItem) SetAutomaticallyPreservesTimeOffsetFromLive(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setAutomaticallyPreservesTimeOffsetFromLive:"), value)
}

// A Boolean value that indicates whether the application can be used to play the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1389929-applicationauthorizedforplayback?language=objc
func (p_ PlayerItem) IsApplicationAuthorizedForPlayback() bool {
	rv := objc.Call[bool](p_, objc.Sel("isApplicationAuthorizedForPlayback"))
	return rv
}

// An upper limit on the resolution of video to download when connecting over expensive networks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3750308-preferredmaximumresolutionforexp?language=objc
func (p_ PlayerItem) PreferredMaximumResolutionForExpensiveNetworks() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](p_, objc.Sel("preferredMaximumResolutionForExpensiveNetworks"))
	return rv
}

// An upper limit on the resolution of video to download when connecting over expensive networks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3750308-preferredmaximumresolutionforexp?language=objc
func (p_ PlayerItem) SetPreferredMaximumResolutionForExpensiveNetworks(value coregraphics.Size) {
	objc.Call[objc.Void](p_, objc.Sel("setPreferredMaximumResolutionForExpensiveNetworks:"), value)
}

// A Boolean value that indicates whether the internal media buffer is full and that further I/O is suspended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388852-playbackbufferfull?language=objc
func (p_ PlayerItem) IsPlaybackBufferFull() bool {
	rv := objc.Call[bool](p_, objc.Sel("isPlaybackBufferFull"))
	return rv
}

// A Boolean value that indicates whether the content has been authorized by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388458-contentauthorizedforplayback?language=objc
func (p_ PlayerItem) IsContentAuthorizedForPlayback() bool {
	rv := objc.Call[bool](p_, objc.Sel("isContentAuthorizedForPlayback"))
	return rv
}

// The audio mix parameters to be applied during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388037-audiomix?language=objc
func (p_ PlayerItem) AudioMix() AudioMix {
	rv := objc.Call[AudioMix](p_, objc.Sel("audioMix"))
	return rv
}

// The audio mix parameters to be applied during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388037-audiomix?language=objc
func (p_ PlayerItem) SetAudioMix(value IAudioMix) {
	objc.Call[objc.Void](p_, objc.Sel("setAudioMix:"), objc.Ptr(value))
}

// An array of time ranges within which it is possible to seek. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1386155-seekabletimeranges?language=objc
func (p_ PlayerItem) SeekableTimeRanges() []foundation.Value {
	rv := objc.Call[[]foundation.Value](p_, objc.Sel("seekableTimeRanges"))
	return rv
}

// The desired maximum resolution of a video that is to be downloaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/2867324-preferredmaximumresolution?language=objc
func (p_ PlayerItem) PreferredMaximumResolution() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](p_, objc.Sel("preferredMaximumResolution"))
	return rv
}

// The desired maximum resolution of a video that is to be downloaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/2867324-preferredmaximumresolution?language=objc
func (p_ PlayerItem) SetPreferredMaximumResolution(value coregraphics.Size) {
	objc.Call[objc.Void](p_, objc.Sel("setPreferredMaximumResolution:"), value)
}

// The preferences the player item uses when selecting variant playlists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3726149-variantpreferences?language=objc
func (p_ PlayerItem) VariantPreferences() VariantPreferences {
	rv := objc.Call[VariantPreferences](p_, objc.Sel("variantPreferences"))
	return rv
}

// The preferences the player item uses when selecting variant playlists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3726149-variantpreferences?language=objc
func (p_ PlayerItem) SetVariantPreferences(value VariantPreferences) {
	objc.Call[objc.Void](p_, objc.Sel("setVariantPreferences:"), value)
}

// A Boolean value that indicates whether the item can play slower than normal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388078-canplayslowforward?language=objc
func (p_ PlayerItem) CanPlaySlowForward() bool {
	rv := objc.Call[bool](p_, objc.Sel("canPlaySlowForward"))
	return rv
}

// The array of asset keys to be automatically loaded before the player item is ready to play. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388633-automaticallyloadedassetkeys?language=objc
func (p_ PlayerItem) AutomaticallyLoadedAssetKeys() []string {
	rv := objc.Call[[]string](p_, objc.Sel("automaticallyLoadedAssetKeys"))
	return rv
}

// The duration of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1389386-duration?language=objc
func (p_ PlayerItem) Duration() coremedia.Time {
	rv := objc.Call[coremedia.Time](p_, objc.Sel("duration"))
	return rv
}

// The duration the player should buffer media from the network ahead of the playhead to guard against playback disruption. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1643630-preferredforwardbufferduration?language=objc
func (p_ PlayerItem) PreferredForwardBufferDuration() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](p_, objc.Sel("preferredForwardBufferDuration"))
	return rv
}

// The duration the player should buffer media from the network ahead of the playhead to guard against playback disruption. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1643630-preferredforwardbufferduration?language=objc
func (p_ PlayerItem) SetPreferredForwardBufferDuration(value foundation.TimeInterval) {
	objc.Call[objc.Void](p_, objc.Sel("setPreferredForwardBufferDuration:"), value)
}

// The status of the player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1389493-status?language=objc
func (p_ PlayerItem) Status() PlayerItemStatus {
	rv := objc.Call[PlayerItemStatus](p_, objc.Sel("status"))
	return rv
}

// The collection of associated media data collectors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1624161-mediadatacollectors?language=objc
func (p_ PlayerItem) MediaDataCollectors() []PlayerItemMediaDataCollector {
	rv := objc.Call[[]PlayerItemMediaDataCollector](p_, objc.Sel("mediaDataCollectors"))
	return rv
}

// A Boolean value that indicates whether playback starts with the first eligible variant that appears in the stream’s main playlist. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3579514-startsonfirsteligiblevariant?language=objc
func (p_ PlayerItem) StartsOnFirstEligibleVariant() bool {
	rv := objc.Call[bool](p_, objc.Sel("startsOnFirstEligibleVariant"))
	return rv
}

// A Boolean value that indicates whether playback starts with the first eligible variant that appears in the stream’s main playlist. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/3579514-startsonfirsteligiblevariant?language=objc
func (p_ PlayerItem) SetStartsOnFirstEligibleVariant(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setStartsOnFirstEligibleVariant:"), value)
}

// The asset provided during initialization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1388177-asset?language=objc
func (p_ PlayerItem) Asset() Asset {
	rv := objc.Call[Asset](p_, objc.Sel("asset"))
	return rv
}

// A Boolean value that indicates whether the item can be fast forwarded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritem/1389096-canplayfastforward?language=objc
func (p_ PlayerItem) CanPlayFastForward() bool {
	rv := objc.Call[bool](p_, objc.Sel("canPlayFastForward"))
	return rv
}
