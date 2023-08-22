// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangePlaybackPositionCommandEvent] class.
var ChangePlaybackPositionCommandEventClass = _ChangePlaybackPositionCommandEventClass{objc.GetClass("MPChangePlaybackPositionCommandEvent")}

type _ChangePlaybackPositionCommandEventClass struct {
	objc.Class
}

// An interface definition for the [ChangePlaybackPositionCommandEvent] class.
type IChangePlaybackPositionCommandEvent interface {
	IRemoteCommandEvent
	PositionTime() foundation.TimeInterval
}

// An event requesting a change in the playback position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeplaybackpositioncommandevent?language=objc
type ChangePlaybackPositionCommandEvent struct {
	RemoteCommandEvent
}

func ChangePlaybackPositionCommandEventFrom(ptr unsafe.Pointer) ChangePlaybackPositionCommandEvent {
	return ChangePlaybackPositionCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}

func (cc _ChangePlaybackPositionCommandEventClass) Alloc() ChangePlaybackPositionCommandEvent {
	rv := objc.Call[ChangePlaybackPositionCommandEvent](cc, objc.Sel("alloc"))
	return rv
}

func ChangePlaybackPositionCommandEvent_Alloc() ChangePlaybackPositionCommandEvent {
	return ChangePlaybackPositionCommandEventClass.Alloc()
}

func (cc _ChangePlaybackPositionCommandEventClass) New() ChangePlaybackPositionCommandEvent {
	rv := objc.Call[ChangePlaybackPositionCommandEvent](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangePlaybackPositionCommandEvent() ChangePlaybackPositionCommandEvent {
	return ChangePlaybackPositionCommandEventClass.New()
}

func (c_ ChangePlaybackPositionCommandEvent) Init() ChangePlaybackPositionCommandEvent {
	rv := objc.Call[ChangePlaybackPositionCommandEvent](c_, objc.Sel("init"))
	return rv
}

// The playback position used when setting the current time of the player. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeplaybackpositioncommandevent/1616766-positiontime?language=objc
func (c_ ChangePlaybackPositionCommandEvent) PositionTime() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](c_, objc.Sel("positionTime"))
	return rv
}
