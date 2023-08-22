// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangePlaybackRateCommandEvent] class.
var ChangePlaybackRateCommandEventClass = _ChangePlaybackRateCommandEventClass{objc.GetClass("MPChangePlaybackRateCommandEvent")}

type _ChangePlaybackRateCommandEventClass struct {
	objc.Class
}

// An interface definition for the [ChangePlaybackRateCommandEvent] class.
type IChangePlaybackRateCommandEvent interface {
	IRemoteCommandEvent
	PlaybackRate() float64
}

// An event requesting a change in the playback rate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeplaybackratecommandevent?language=objc
type ChangePlaybackRateCommandEvent struct {
	RemoteCommandEvent
}

func ChangePlaybackRateCommandEventFrom(ptr unsafe.Pointer) ChangePlaybackRateCommandEvent {
	return ChangePlaybackRateCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}

func (cc _ChangePlaybackRateCommandEventClass) Alloc() ChangePlaybackRateCommandEvent {
	rv := objc.Call[ChangePlaybackRateCommandEvent](cc, objc.Sel("alloc"))
	return rv
}

func ChangePlaybackRateCommandEvent_Alloc() ChangePlaybackRateCommandEvent {
	return ChangePlaybackRateCommandEventClass.Alloc()
}

func (cc _ChangePlaybackRateCommandEventClass) New() ChangePlaybackRateCommandEvent {
	rv := objc.Call[ChangePlaybackRateCommandEvent](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangePlaybackRateCommandEvent() ChangePlaybackRateCommandEvent {
	return ChangePlaybackRateCommandEventClass.New()
}

func (c_ ChangePlaybackRateCommandEvent) Init() ChangePlaybackRateCommandEvent {
	rv := objc.Call[ChangePlaybackRateCommandEvent](c_, objc.Sel("init"))
	return rv
}

// The chosen playback rate for the command event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeplaybackratecommandevent/1616782-playbackrate?language=objc
func (c_ ChangePlaybackRateCommandEvent) PlaybackRate() float64 {
	rv := objc.Call[float64](c_, objc.Sel("playbackRate"))
	return rv
}
