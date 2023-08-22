// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangeShuffleModeCommandEvent] class.
var ChangeShuffleModeCommandEventClass = _ChangeShuffleModeCommandEventClass{objc.GetClass("MPChangeShuffleModeCommandEvent")}

type _ChangeShuffleModeCommandEventClass struct {
	objc.Class
}

// An interface definition for the [ChangeShuffleModeCommandEvent] class.
type IChangeShuffleModeCommandEvent interface {
	IRemoteCommandEvent
	PreservesShuffleMode() bool
	ShuffleType() ShuffleType
}

// An event requesting a change in the shuffle mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeshufflemodecommandevent?language=objc
type ChangeShuffleModeCommandEvent struct {
	RemoteCommandEvent
}

func ChangeShuffleModeCommandEventFrom(ptr unsafe.Pointer) ChangeShuffleModeCommandEvent {
	return ChangeShuffleModeCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}

func (cc _ChangeShuffleModeCommandEventClass) Alloc() ChangeShuffleModeCommandEvent {
	rv := objc.Call[ChangeShuffleModeCommandEvent](cc, objc.Sel("alloc"))
	return rv
}

func ChangeShuffleModeCommandEvent_Alloc() ChangeShuffleModeCommandEvent {
	return ChangeShuffleModeCommandEventClass.Alloc()
}

func (cc _ChangeShuffleModeCommandEventClass) New() ChangeShuffleModeCommandEvent {
	rv := objc.Call[ChangeShuffleModeCommandEvent](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangeShuffleModeCommandEvent() ChangeShuffleModeCommandEvent {
	return ChangeShuffleModeCommandEventClass.New()
}

func (c_ ChangeShuffleModeCommandEvent) Init() ChangeShuffleModeCommandEvent {
	rv := objc.Call[ChangeShuffleModeCommandEvent](c_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the shuffle mode is preserved between playback sessions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeshufflemodecommandevent/2097552-preservesshufflemode?language=objc
func (c_ ChangeShuffleModeCommandEvent) PreservesShuffleMode() bool {
	rv := objc.Call[bool](c_, objc.Sel("preservesShuffleMode"))
	return rv
}

// The shuffle type used when fulfilling the event request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeshufflemodecommandevent/1649696-shuffletype?language=objc
func (c_ ChangeShuffleModeCommandEvent) ShuffleType() ShuffleType {
	rv := objc.Call[ShuffleType](c_, objc.Sel("shuffleType"))
	return rv
}
