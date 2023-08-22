// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangeRepeatModeCommandEvent] class.
var ChangeRepeatModeCommandEventClass = _ChangeRepeatModeCommandEventClass{objc.GetClass("MPChangeRepeatModeCommandEvent")}

type _ChangeRepeatModeCommandEventClass struct {
	objc.Class
}

// An interface definition for the [ChangeRepeatModeCommandEvent] class.
type IChangeRepeatModeCommandEvent interface {
	IRemoteCommandEvent
	RepeatType() RepeatType
	PreservesRepeatMode() bool
}

// An event requesting a change in the repeat mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangerepeatmodecommandevent?language=objc
type ChangeRepeatModeCommandEvent struct {
	RemoteCommandEvent
}

func ChangeRepeatModeCommandEventFrom(ptr unsafe.Pointer) ChangeRepeatModeCommandEvent {
	return ChangeRepeatModeCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}

func (cc _ChangeRepeatModeCommandEventClass) Alloc() ChangeRepeatModeCommandEvent {
	rv := objc.Call[ChangeRepeatModeCommandEvent](cc, objc.Sel("alloc"))
	return rv
}

func ChangeRepeatModeCommandEvent_Alloc() ChangeRepeatModeCommandEvent {
	return ChangeRepeatModeCommandEventClass.Alloc()
}

func (cc _ChangeRepeatModeCommandEventClass) New() ChangeRepeatModeCommandEvent {
	rv := objc.Call[ChangeRepeatModeCommandEvent](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangeRepeatModeCommandEvent() ChangeRepeatModeCommandEvent {
	return ChangeRepeatModeCommandEventClass.New()
}

func (c_ ChangeRepeatModeCommandEvent) Init() ChangeRepeatModeCommandEvent {
	rv := objc.Call[ChangeRepeatModeCommandEvent](c_, objc.Sel("init"))
	return rv
}

// The repeat type used when fulfilling the event request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangerepeatmodecommandevent/1649689-repeattype?language=objc
func (c_ ChangeRepeatModeCommandEvent) RepeatType() RepeatType {
	rv := objc.Call[RepeatType](c_, objc.Sel("repeatType"))
	return rv
}

// A Boolean value that indicates whether the chosen repeat mode is preserved between playback sessions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangerepeatmodecommandevent/2097553-preservesrepeatmode?language=objc
func (c_ ChangeRepeatModeCommandEvent) PreservesRepeatMode() bool {
	rv := objc.Call[bool](c_, objc.Sel("preservesRepeatMode"))
	return rv
}
