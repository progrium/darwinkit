// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DelegatingPlaybackCoordinator] class.
var DelegatingPlaybackCoordinatorClass = _DelegatingPlaybackCoordinatorClass{objc.GetClass("AVDelegatingPlaybackCoordinator")}

type _DelegatingPlaybackCoordinatorClass struct {
	objc.Class
}

// An interface definition for the [DelegatingPlaybackCoordinator] class.
type IDelegatingPlaybackCoordinator interface {
	IPlaybackCoordinator
	CoordinateRateChangeToRateOptions(rate float64, options DelegatingPlaybackCoordinatorRateChangeOptions)
	TransitionToItemWithIdentifierProposingInitialTimingBasedOnTimebase(itemIdentifier string, snapshotTimebase coremedia.TimebaseRef)
	ReapplyCurrentItemStateToPlaybackControlDelegate()
	CoordinateSeekToTimeOptions(time coremedia.Time, options DelegatingPlaybackCoordinatorSeekOptions)
	PlaybackControlDelegate() PlaybackCoordinatorPlaybackControlDelegateWrapper
	CurrentItemIdentifier() string
}

// A playback coordinator subclass that coordinates the playback of custom player objects in a connected group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinator?language=objc
type DelegatingPlaybackCoordinator struct {
	PlaybackCoordinator
}

func DelegatingPlaybackCoordinatorFrom(ptr unsafe.Pointer) DelegatingPlaybackCoordinator {
	return DelegatingPlaybackCoordinator{
		PlaybackCoordinator: PlaybackCoordinatorFrom(ptr),
	}
}

func (d_ DelegatingPlaybackCoordinator) InitWithPlaybackControlDelegate(playbackControlDelegate PPlaybackCoordinatorPlaybackControlDelegate) DelegatingPlaybackCoordinator {
	po0 := objc.WrapAsProtocol("AVPlaybackCoordinatorPlaybackControlDelegate", playbackControlDelegate)
	rv := objc.Call[DelegatingPlaybackCoordinator](d_, objc.Sel("initWithPlaybackControlDelegate:"), po0)
	return rv
}

// Creates a playback coordinator for a custom playback object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinator/3750256-initwithplaybackcontroldelegate?language=objc
func NewDelegatingPlaybackCoordinatorWithPlaybackControlDelegate(playbackControlDelegate PPlaybackCoordinatorPlaybackControlDelegate) DelegatingPlaybackCoordinator {
	instance := DelegatingPlaybackCoordinatorClass.Alloc().InitWithPlaybackControlDelegate(playbackControlDelegate)
	instance.Autorelease()
	return instance
}

func (dc _DelegatingPlaybackCoordinatorClass) Alloc() DelegatingPlaybackCoordinator {
	rv := objc.Call[DelegatingPlaybackCoordinator](dc, objc.Sel("alloc"))
	return rv
}

func DelegatingPlaybackCoordinator_Alloc() DelegatingPlaybackCoordinator {
	return DelegatingPlaybackCoordinatorClass.Alloc()
}

func (dc _DelegatingPlaybackCoordinatorClass) New() DelegatingPlaybackCoordinator {
	rv := objc.Call[DelegatingPlaybackCoordinator](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDelegatingPlaybackCoordinator() DelegatingPlaybackCoordinator {
	return DelegatingPlaybackCoordinatorClass.New()
}

func (d_ DelegatingPlaybackCoordinator) Init() DelegatingPlaybackCoordinator {
	rv := objc.Call[DelegatingPlaybackCoordinator](d_, objc.Sel("init"))
	return rv
}

// Coordinates a rate change across all participants, waiting for others to become ready, if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinator/3750250-coordinateratechangetorate?language=objc
func (d_ DelegatingPlaybackCoordinator) CoordinateRateChangeToRateOptions(rate float64, options DelegatingPlaybackCoordinatorRateChangeOptions) {
	objc.Call[objc.Void](d_, objc.Sel("coordinateRateChangeToRate:options:"), rate, options)
}

// Tells the coordinator to transition to a new item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinator/3750259-transitiontoitemwithidentifier?language=objc
func (d_ DelegatingPlaybackCoordinator) TransitionToItemWithIdentifierProposingInitialTimingBasedOnTimebase(itemIdentifier string, snapshotTimebase coremedia.TimebaseRef) {
	objc.Call[objc.Void](d_, objc.Sel("transitionToItemWithIdentifier:proposingInitialTimingBasedOnTimebase:"), itemIdentifier, snapshotTimebase)
}

// Tells the coordinator to reissue current play state commands to synchronize the current item to the state of other participants. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinator/3750258-reapplycurrentitemstatetoplaybac?language=objc
func (d_ DelegatingPlaybackCoordinator) ReapplyCurrentItemStateToPlaybackControlDelegate() {
	objc.Call[objc.Void](d_, objc.Sel("reapplyCurrentItemStateToPlaybackControlDelegate"))
}

// Coordinates a seek to the specified time for all connected participants. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinator/3750251-coordinateseektotime?language=objc
func (d_ DelegatingPlaybackCoordinator) CoordinateSeekToTimeOptions(time coremedia.Time, options DelegatingPlaybackCoordinatorSeekOptions) {
	objc.Call[objc.Void](d_, objc.Sel("coordinateSeekToTime:options:"), time, options)
}

// The delegate object for the playback coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinator/3750257-playbackcontroldelegate?language=objc
func (d_ DelegatingPlaybackCoordinator) PlaybackControlDelegate() PlaybackCoordinatorPlaybackControlDelegateWrapper {
	rv := objc.Call[PlaybackCoordinatorPlaybackControlDelegateWrapper](d_, objc.Sel("playbackControlDelegate"))
	return rv
}

// An identifier of the current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinator/3750252-currentitemidentifier?language=objc
func (d_ DelegatingPlaybackCoordinator) CurrentItemIdentifier() string {
	rv := objc.Call[string](d_, objc.Sel("currentItemIdentifier"))
	return rv
}
