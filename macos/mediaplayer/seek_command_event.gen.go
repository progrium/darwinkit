// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SeekCommandEvent] class.
var SeekCommandEventClass = _SeekCommandEventClass{objc.GetClass("MPSeekCommandEvent")}

type _SeekCommandEventClass struct {
	objc.Class
}

// An interface definition for the [SeekCommandEvent] class.
type ISeekCommandEvent interface {
	IRemoteCommandEvent
	Type() SeekCommandEventType
}

// An event requesting that the player seek to a new position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpseekcommandevent?language=objc
type SeekCommandEvent struct {
	RemoteCommandEvent
}

func SeekCommandEventFrom(ptr unsafe.Pointer) SeekCommandEvent {
	return SeekCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}

func (sc _SeekCommandEventClass) Alloc() SeekCommandEvent {
	rv := objc.Call[SeekCommandEvent](sc, objc.Sel("alloc"))
	return rv
}

func SeekCommandEvent_Alloc() SeekCommandEvent {
	return SeekCommandEventClass.Alloc()
}

func (sc _SeekCommandEventClass) New() SeekCommandEvent {
	rv := objc.Call[SeekCommandEvent](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSeekCommandEvent() SeekCommandEvent {
	return SeekCommandEventClass.New()
}

func (s_ SeekCommandEvent) Init() SeekCommandEvent {
	rv := objc.Call[SeekCommandEvent](s_, objc.Sel("init"))
	return rv
}

// The type of seek command event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpseekcommandevent/1616783-type?language=objc
func (s_ SeekCommandEvent) Type() SeekCommandEventType {
	rv := objc.Call[SeekCommandEventType](s_, objc.Sel("type"))
	return rv
}
