// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SkipIntervalCommandEvent] class.
var SkipIntervalCommandEventClass = _SkipIntervalCommandEventClass{objc.GetClass("MPSkipIntervalCommandEvent")}

type _SkipIntervalCommandEventClass struct {
	objc.Class
}

// An interface definition for the [SkipIntervalCommandEvent] class.
type ISkipIntervalCommandEvent interface {
	IRemoteCommandEvent
	Interval() foundation.TimeInterval
}

// An event requesting a change in the current skip interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpskipintervalcommandevent?language=objc
type SkipIntervalCommandEvent struct {
	RemoteCommandEvent
}

func SkipIntervalCommandEventFrom(ptr unsafe.Pointer) SkipIntervalCommandEvent {
	return SkipIntervalCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}

func (sc _SkipIntervalCommandEventClass) Alloc() SkipIntervalCommandEvent {
	rv := objc.Call[SkipIntervalCommandEvent](sc, objc.Sel("alloc"))
	return rv
}

func SkipIntervalCommandEvent_Alloc() SkipIntervalCommandEvent {
	return SkipIntervalCommandEventClass.Alloc()
}

func (sc _SkipIntervalCommandEventClass) New() SkipIntervalCommandEvent {
	rv := objc.Call[SkipIntervalCommandEvent](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSkipIntervalCommandEvent() SkipIntervalCommandEvent {
	return SkipIntervalCommandEventClass.New()
}

func (s_ SkipIntervalCommandEvent) Init() SkipIntervalCommandEvent {
	rv := objc.Call[SkipIntervalCommandEvent](s_, objc.Sel("init"))
	return rv
}

// The chosen interval, in seconds, for the skip command event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpskipintervalcommandevent/1616767-interval?language=objc
func (s_ SkipIntervalCommandEvent) Interval() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](s_, objc.Sel("interval"))
	return rv
}
