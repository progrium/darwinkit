// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Player] class.
var PlayerClass = _PlayerClass{objc.GetClass("AVPlayer")}

type _PlayerClass struct {
	objc.Class
}

// An interface definition for the [Player] class.
type IPlayer interface {
	objc.IObject
	MediaSelectionCriteriaForMediaCharacteristic(mediaCharacteristic MediaCharacteristic) PlayerMediaSelectionCriteria
	SeekToDate(date foundation.IDate)
	PlayImmediatelyAtRate(rate float64)
	RemoveTimeObserver(observer objc.IObject)
	SetRateTimeAtHostTime(rate float64, itemTime coremedia.Time, hostClockTime coremedia.Time)
	ReplaceCurrentItemWithPlayerItem(item IPlayerItem)
	CancelPendingPrerolls()
	AddBoundaryTimeObserverForTimesQueueUsingBlock(times []foundation.IValue, queue dispatch.Queue, block func()) objc.Object
	AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval coremedia.Time, queue dispatch.Queue, block func(time coremedia.Time)) objc.Object
	SeekToTimeToleranceBeforeToleranceAfter(time coremedia.Time, toleranceBefore coremedia.Time, toleranceAfter coremedia.Time)
	SetMediaSelectionCriteriaForMediaCharacteristic(criteria IPlayerMediaSelectionCriteria, mediaCharacteristic MediaCharacteristic)
	Pause()
	Play()
	PrerollAtRateCompletionHandler(rate float64, completionHandler func(finished bool))
	CurrentTime() coremedia.Time
	SourceClock() coremedia.ClockRef
	SetSourceClock(value coremedia.ClockRef)
	AudiovisualBackgroundPlaybackPolicy() PlayerAudiovisualBackgroundPlaybackPolicy
	SetAudiovisualBackgroundPlaybackPolicy(value PlayerAudiovisualBackgroundPlaybackPolicy)
	Volume() float64
	SetVolume(value float64)
	Error() foundation.Error
	OutputObscuredDueToInsufficientExternalProtection() bool
	PreferredVideoDecoderGPURegistryID() uint64
	SetPreferredVideoDecoderGPURegistryID(value uint64)
	Rate() float64
	SetRate(value float64)
	IsMuted() bool
	SetMuted(value bool)
	ReasonForWaitingToPlay() PlayerWaitingReason
	PlaybackCoordinator() PlayerPlaybackCoordinator
	CurrentItem() PlayerItem
	AudioOutputDeviceUniqueID() string
	SetAudioOutputDeviceUniqueID(value string)
	TimeControlStatus() PlayerTimeControlStatus
	IsExternalPlaybackActive() bool
	ActionAtItemEnd() PlayerActionAtItemEnd
	SetActionAtItemEnd(value PlayerActionAtItemEnd)
	AllowsExternalPlayback() bool
	SetAllowsExternalPlayback(value bool)
	AppliesMediaSelectionCriteriaAutomatically() bool
	SetAppliesMediaSelectionCriteriaAutomatically(value bool)
	PreventsDisplaySleepDuringVideoPlayback() bool
	SetPreventsDisplaySleepDuringVideoPlayback(value bool)
	Status() PlayerStatus
	AutomaticallyWaitsToMinimizeStalling() bool
	SetAutomaticallyWaitsToMinimizeStalling(value bool)
}

// An object that provides the interface to control the player’s transport behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer?language=objc
type Player struct {
	objc.Object
}

func PlayerFrom(ptr unsafe.Pointer) Player {
	return Player{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PlayerClass) PlayerWithPlayerItem(item IPlayerItem) Player {
	rv := objc.Call[Player](pc, objc.Sel("playerWithPlayerItem:"), objc.Ptr(item))
	return rv
}

// Returns a new player initialized to play the specified player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1538390-playerwithplayeritem?language=objc
func Player_PlayerWithPlayerItem(item IPlayerItem) Player {
	return PlayerClass.PlayerWithPlayerItem(item)
}

func (p_ Player) InitWithURL(URL foundation.IURL) Player {
	rv := objc.Call[Player](p_, objc.Sel("initWithURL:"), objc.Ptr(URL))
	return rv
}

// Creates a new player to play a single audiovisual resource referenced by a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1385706-initwithurl?language=objc
func NewPlayerWithURL(URL foundation.IURL) Player {
	instance := PlayerClass.Alloc().InitWithURL(URL)
	instance.Autorelease()
	return instance
}

func (p_ Player) InitWithPlayerItem(item IPlayerItem) Player {
	rv := objc.Call[Player](p_, objc.Sel("initWithPlayerItem:"), objc.Ptr(item))
	return rv
}

// Creates a new player to play the specified player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1387104-initwithplayeritem?language=objc
func NewPlayerWithPlayerItem(item IPlayerItem) Player {
	instance := PlayerClass.Alloc().InitWithPlayerItem(item)
	instance.Autorelease()
	return instance
}

func (pc _PlayerClass) PlayerWithURL(URL foundation.IURL) Player {
	rv := objc.Call[Player](pc, objc.Sel("playerWithURL:"), objc.Ptr(URL))
	return rv
}

// Returns a new player to play a single audiovisual resource referenced by a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1538409-playerwithurl?language=objc
func Player_PlayerWithURL(URL foundation.IURL) Player {
	return PlayerClass.PlayerWithURL(URL)
}

func (pc _PlayerClass) Alloc() Player {
	rv := objc.Call[Player](pc, objc.Sel("alloc"))
	return rv
}

func Player_Alloc() Player {
	return PlayerClass.Alloc()
}

func (pc _PlayerClass) New() Player {
	rv := objc.Call[Player](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayer() Player {
	return PlayerClass.New()
}

func (p_ Player) Init() Player {
	rv := objc.Call[Player](p_, objc.Sel("init"))
	return rv
}

// Returns the automatic selection criteria for media items with the specified media characteristic. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1387825-mediaselectioncriteriaformediach?language=objc
func (p_ Player) MediaSelectionCriteriaForMediaCharacteristic(mediaCharacteristic MediaCharacteristic) PlayerMediaSelectionCriteria {
	rv := objc.Call[PlayerMediaSelectionCriteria](p_, objc.Sel("mediaSelectionCriteriaForMediaCharacteristic:"), mediaCharacteristic)
	return rv
}

// Requests that the player seek to a specified date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1386114-seektodate?language=objc
func (p_ Player) SeekToDate(date foundation.IDate) {
	objc.Call[objc.Void](p_, objc.Sel("seekToDate:"), objc.Ptr(date))
}

// Plays the available media data immediately, at the specified rate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1643480-playimmediatelyatrate?language=objc
func (p_ Player) PlayImmediatelyAtRate(rate float64) {
	objc.Call[objc.Void](p_, objc.Sel("playImmediatelyAtRate:"), rate)
}

// Cancels a previously registered periodic or boundary time observer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1387552-removetimeobserver?language=objc
func (p_ Player) RemoveTimeObserver(observer objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("removeTimeObserver:"), observer)
}

// Synchronizes the playback rate and time of the current item with an external source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1386591-setrate?language=objc
func (p_ Player) SetRateTimeAtHostTime(rate float64, itemTime coremedia.Time, hostClockTime coremedia.Time) {
	objc.Call[objc.Void](p_, objc.Sel("setRate:time:atHostTime:"), rate, itemTime, hostClockTime)
}

// Replaces the current item with a new item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1390806-replacecurrentitemwithplayeritem?language=objc
func (p_ Player) ReplaceCurrentItemWithPlayerItem(item IPlayerItem) {
	objc.Call[objc.Void](p_, objc.Sel("replaceCurrentItemWithPlayerItem:"), objc.Ptr(item))
}

// Cancels any pending preroll requests and invokes the corresponding completion handlers, if present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1388260-cancelpendingprerolls?language=objc
func (p_ Player) CancelPendingPrerolls() {
	objc.Call[objc.Void](p_, objc.Sel("cancelPendingPrerolls"))
}

// Requests the invocation of a block when specified times are traversed during normal playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1388027-addboundarytimeobserverfortimes?language=objc
func (p_ Player) AddBoundaryTimeObserverForTimesQueueUsingBlock(times []foundation.IValue, queue dispatch.Queue, block func()) objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("addBoundaryTimeObserverForTimes:queue:usingBlock:"), times, queue, block)
	return rv
}

// Requests the periodic invocation of a given block during playback to report changing time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1385829-addperiodictimeobserverforinterv?language=objc
func (p_ Player) AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval coremedia.Time, queue dispatch.Queue, block func(time coremedia.Time)) objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("addPeriodicTimeObserverForInterval:queue:usingBlock:"), interval, queue, block)
	return rv
}

// Requests that the player seek to a specified time with the amount of accuracy specified by the time tolerance values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1387741-seektotime?language=objc
func (p_ Player) SeekToTimeToleranceBeforeToleranceAfter(time coremedia.Time, toleranceBefore coremedia.Time, toleranceAfter coremedia.Time) {
	objc.Call[objc.Void](p_, objc.Sel("seekToTime:toleranceBefore:toleranceAfter:"), time, toleranceBefore, toleranceAfter)
}

// Applies automatic selection criteria for media that has the specified media characteristic. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1390563-setmediaselectioncriteria?language=objc
func (p_ Player) SetMediaSelectionCriteriaForMediaCharacteristic(criteria IPlayerMediaSelectionCriteria, mediaCharacteristic MediaCharacteristic) {
	objc.Call[objc.Void](p_, objc.Sel("setMediaSelectionCriteria:forMediaCharacteristic:"), objc.Ptr(criteria), mediaCharacteristic)
}

// Pauses playback of the current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1387895-pause?language=objc
func (p_ Player) Pause() {
	objc.Call[objc.Void](p_, objc.Sel("pause"))
}

// Begins playback of the current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1386726-play?language=objc
func (p_ Player) Play() {
	objc.Call[objc.Void](p_, objc.Sel("play"))
}

// Begins loading media data to prime the media pipelines for playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1389712-prerollatrate?language=objc
func (p_ Player) PrerollAtRateCompletionHandler(rate float64, completionHandler func(finished bool)) {
	objc.Call[objc.Void](p_, objc.Sel("prerollAtRate:completionHandler:"), rate, completionHandler)
}

// Returns the current time of the current player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1390404-currenttime?language=objc
func (p_ Player) CurrentTime() coremedia.Time {
	rv := objc.Call[coremedia.Time](p_, objc.Sel("currentTime"))
	return rv
}

// A clock the player uses for item time bases. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/3746583-sourceclock?language=objc
func (p_ Player) SourceClock() coremedia.ClockRef {
	rv := objc.Call[coremedia.ClockRef](p_, objc.Sel("sourceClock"))
	return rv
}

// A clock the player uses for item time bases. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/3746583-sourceclock?language=objc
func (p_ Player) SetSourceClock(value coremedia.ClockRef) {
	objc.Call[objc.Void](p_, objc.Sel("setSourceClock:"), value)
}

// A policy that determines how playback of audiovisual media continues when the app transitions to the background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/3787548-audiovisualbackgroundplaybackpol?language=objc
func (p_ Player) AudiovisualBackgroundPlaybackPolicy() PlayerAudiovisualBackgroundPlaybackPolicy {
	rv := objc.Call[PlayerAudiovisualBackgroundPlaybackPolicy](p_, objc.Sel("audiovisualBackgroundPlaybackPolicy"))
	return rv
}

// A policy that determines how playback of audiovisual media continues when the app transitions to the background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/3787548-audiovisualbackgroundplaybackpol?language=objc
func (p_ Player) SetAudiovisualBackgroundPlaybackPolicy(value PlayerAudiovisualBackgroundPlaybackPolicy) {
	objc.Call[objc.Void](p_, objc.Sel("setAudiovisualBackgroundPlaybackPolicy:"), value)
}

// The audio playback volume for the player. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1390127-volume?language=objc
func (p_ Player) Volume() float64 {
	rv := objc.Call[float64](p_, objc.Sel("volume"))
	return rv
}

// The audio playback volume for the player. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1390127-volume?language=objc
func (p_ Player) SetVolume(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setVolume:"), value)
}

// An error that caused a failure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1387764-error?language=objc
func (p_ Player) Error() foundation.Error {
	rv := objc.Call[foundation.Error](p_, objc.Sel("error"))
	return rv
}

// A Boolean value that indicates whether output is being obscured because of insufficient external protection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1624254-outputobscuredduetoinsufficiente?language=objc
func (p_ Player) OutputObscuredDueToInsufficientExternalProtection() bool {
	rv := objc.Call[bool](p_, objc.Sel("outputObscuredDueToInsufficientExternalProtection"))
	return rv
}

// The registry identifier for the GPU used for video decoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/2942616-preferredvideodecodergpuregistry?language=objc
func (p_ Player) PreferredVideoDecoderGPURegistryID() uint64 {
	rv := objc.Call[uint64](p_, objc.Sel("preferredVideoDecoderGPURegistryID"))
	return rv
}

// The registry identifier for the GPU used for video decoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/2942616-preferredvideodecodergpuregistry?language=objc
func (p_ Player) SetPreferredVideoDecoderGPURegistryID(value uint64) {
	objc.Call[objc.Void](p_, objc.Sel("setPreferredVideoDecoderGPURegistryID:"), value)
}

// The current playback rate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1388846-rate?language=objc
func (p_ Player) Rate() float64 {
	rv := objc.Call[float64](p_, objc.Sel("rate"))
	return rv
}

// The current playback rate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1388846-rate?language=objc
func (p_ Player) SetRate(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setRate:"), value)
}

// A Boolean value that indicates whether the audio output of the player is muted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1387544-muted?language=objc
func (p_ Player) IsMuted() bool {
	rv := objc.Call[bool](p_, objc.Sel("isMuted"))
	return rv
}

// A Boolean value that indicates whether the audio output of the player is muted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1387544-muted?language=objc
func (p_ Player) SetMuted(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setMuted:"), value)
}

// The reason the player is currently waiting for playback to begin or resume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1643486-reasonforwaitingtoplay?language=objc
func (p_ Player) ReasonForWaitingToPlay() PlayerWaitingReason {
	rv := objc.Call[PlayerWaitingReason](p_, objc.Sel("reasonForWaitingToPlay"))
	return rv
}

// The playback coordinator for the player. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/3750305-playbackcoordinator?language=objc
func (p_ Player) PlaybackCoordinator() PlayerPlaybackCoordinator {
	rv := objc.Call[PlayerPlaybackCoordinator](p_, objc.Sel("playbackCoordinator"))
	return rv
}

// The item for which the player is currently controlling playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1387569-currentitem?language=objc
func (p_ Player) CurrentItem() PlayerItem {
	rv := objc.Call[PlayerItem](p_, objc.Sel("currentItem"))
	return rv
}

// Specifies the unique ID of the Core Audio output device used to play audio. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1390717-audiooutputdeviceuniqueid?language=objc
func (p_ Player) AudioOutputDeviceUniqueID() string {
	rv := objc.Call[string](p_, objc.Sel("audioOutputDeviceUniqueID"))
	return rv
}

// Specifies the unique ID of the Core Audio output device used to play audio. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1390717-audiooutputdeviceuniqueid?language=objc
func (p_ Player) SetAudioOutputDeviceUniqueID(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setAudioOutputDeviceUniqueID:"), value)
}

// A value that indicates whether playback is in progress, paused indefinitely, or waiting for network conditions to improve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1643485-timecontrolstatus?language=objc
func (p_ Player) TimeControlStatus() PlayerTimeControlStatus {
	rv := objc.Call[PlayerTimeControlStatus](p_, objc.Sel("timeControlStatus"))
	return rv
}

// A Boolean value that indicates whether the player is currently playing video in external playback mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1388982-externalplaybackactive?language=objc
func (p_ Player) IsExternalPlaybackActive() bool {
	rv := objc.Call[bool](p_, objc.Sel("isExternalPlaybackActive"))
	return rv
}

// The action to perform when the current player item has finished playing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1387376-actionatitemend?language=objc
func (p_ Player) ActionAtItemEnd() PlayerActionAtItemEnd {
	rv := objc.Call[PlayerActionAtItemEnd](p_, objc.Sel("actionAtItemEnd"))
	return rv
}

// The action to perform when the current player item has finished playing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1387376-actionatitemend?language=objc
func (p_ Player) SetActionAtItemEnd(value PlayerActionAtItemEnd) {
	objc.Call[objc.Void](p_, objc.Sel("setActionAtItemEnd:"), value)
}

// A Boolean value that indicates whether the player allows switching to external playback mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1387441-allowsexternalplayback?language=objc
func (p_ Player) AllowsExternalPlayback() bool {
	rv := objc.Call[bool](p_, objc.Sel("allowsExternalPlayback"))
	return rv
}

// A Boolean value that indicates whether the player allows switching to external playback mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1387441-allowsexternalplayback?language=objc
func (p_ Player) SetAllowsExternalPlayback(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setAllowsExternalPlayback:"), value)
}

// A Boolean value that indicates whether the current device can present content to an HDR display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/3365978-eligibleforhdrplayback?language=objc
func (pc _PlayerClass) EligibleForHDRPlayback() bool {
	rv := objc.Call[bool](pc, objc.Sel("eligibleForHDRPlayback"))
	return rv
}

// A Boolean value that indicates whether the current device can present content to an HDR display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/3365978-eligibleforhdrplayback?language=objc
func Player_EligibleForHDRPlayback() bool {
	return PlayerClass.EligibleForHDRPlayback()
}

// A Boolean value that indicates whether the receiver should apply the current selection criteria automatically to player items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1387178-appliesmediaselectioncriteriaaut?language=objc
func (p_ Player) AppliesMediaSelectionCriteriaAutomatically() bool {
	rv := objc.Call[bool](p_, objc.Sel("appliesMediaSelectionCriteriaAutomatically"))
	return rv
}

// A Boolean value that indicates whether the receiver should apply the current selection criteria automatically to player items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1387178-appliesmediaselectioncriteriaaut?language=objc
func (p_ Player) SetAppliesMediaSelectionCriteriaAutomatically(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setAppliesMediaSelectionCriteriaAutomatically:"), value)
}

// A Boolean value that indicates whether video playback prevents display and device sleep. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/2990522-preventsdisplaysleepduringvideop?language=objc
func (p_ Player) PreventsDisplaySleepDuringVideoPlayback() bool {
	rv := objc.Call[bool](p_, objc.Sel("preventsDisplaySleepDuringVideoPlayback"))
	return rv
}

// A Boolean value that indicates whether video playback prevents display and device sleep. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/2990522-preventsdisplaysleepduringvideop?language=objc
func (p_ Player) SetPreventsDisplaySleepDuringVideoPlayback(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setPreventsDisplaySleepDuringVideoPlayback:"), value)
}

// A value that indicates the readiness of a player object for playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1388096-status?language=objc
func (p_ Player) Status() PlayerStatus {
	rv := objc.Call[PlayerStatus](p_, objc.Sel("status"))
	return rv
}

// A Boolean value that indicates whether the player should automatically delay playback in order to minimize stalling. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1643482-automaticallywaitstominimizestal?language=objc
func (p_ Player) AutomaticallyWaitsToMinimizeStalling() bool {
	rv := objc.Call[bool](p_, objc.Sel("automaticallyWaitsToMinimizeStalling"))
	return rv
}

// A Boolean value that indicates whether the player should automatically delay playback in order to minimize stalling. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/1643482-automaticallywaitstominimizestal?language=objc
func (p_ Player) SetAutomaticallyWaitsToMinimizeStalling(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setAutomaticallyWaitsToMinimizeStalling:"), value)
}
