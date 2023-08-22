// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RemoteCommandEvent] class.
var RemoteCommandEventClass = _RemoteCommandEventClass{objc.GetClass("MPRemoteCommandEvent")}

type _RemoteCommandEventClass struct {
	objc.Class
}

// An interface definition for the [RemoteCommandEvent] class.
type IRemoteCommandEvent interface {
	objc.IObject
	Command() RemoteCommand
	Timestamp() foundation.TimeInterval
}

// A description of a command sent by an external media player. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandevent?language=objc
type RemoteCommandEvent struct {
	objc.Object
}

func RemoteCommandEventFrom(ptr unsafe.Pointer) RemoteCommandEvent {
	return RemoteCommandEvent{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RemoteCommandEventClass) Alloc() RemoteCommandEvent {
	rv := objc.Call[RemoteCommandEvent](rc, objc.Sel("alloc"))
	return rv
}

func RemoteCommandEvent_Alloc() RemoteCommandEvent {
	return RemoteCommandEventClass.Alloc()
}

func (rc _RemoteCommandEventClass) New() RemoteCommandEvent {
	rv := objc.Call[RemoteCommandEvent](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRemoteCommandEvent() RemoteCommandEvent {
	return RemoteCommandEventClass.New()
}

func (r_ RemoteCommandEvent) Init() RemoteCommandEvent {
	rv := objc.Call[RemoteCommandEvent](r_, objc.Sel("init"))
	return rv
}

// The command that sent the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandevent/1616776-command?language=objc
func (r_ RemoteCommandEvent) Command() RemoteCommand {
	rv := objc.Call[RemoteCommand](r_, objc.Sel("command"))
	return rv
}

// The time the event occurred. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandevent/1616784-timestamp?language=objc
func (r_ RemoteCommandEvent) Timestamp() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](r_, objc.Sel("timestamp"))
	return rv
}
