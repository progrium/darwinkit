// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SkipIntervalCommand] class.
var SkipIntervalCommandClass = _SkipIntervalCommandClass{objc.GetClass("MPSkipIntervalCommand")}

type _SkipIntervalCommandClass struct {
	objc.Class
}

// An interface definition for the [SkipIntervalCommand] class.
type ISkipIntervalCommand interface {
	IRemoteCommand
	PreferredIntervals() []foundation.Number
	SetPreferredIntervals(value []foundation.INumber)
}

// An object that defines the skip intervals for the player. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpskipintervalcommand?language=objc
type SkipIntervalCommand struct {
	RemoteCommand
}

func SkipIntervalCommandFrom(ptr unsafe.Pointer) SkipIntervalCommand {
	return SkipIntervalCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}

func (sc _SkipIntervalCommandClass) Alloc() SkipIntervalCommand {
	rv := objc.Call[SkipIntervalCommand](sc, objc.Sel("alloc"))
	return rv
}

func SkipIntervalCommand_Alloc() SkipIntervalCommand {
	return SkipIntervalCommandClass.Alloc()
}

func (sc _SkipIntervalCommandClass) New() SkipIntervalCommand {
	rv := objc.Call[SkipIntervalCommand](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSkipIntervalCommand() SkipIntervalCommand {
	return SkipIntervalCommandClass.New()
}

func (s_ SkipIntervalCommand) Init() SkipIntervalCommand {
	rv := objc.Call[SkipIntervalCommand](s_, objc.Sel("init"))
	return rv
}

// The available skip intervals, in seconds, for a media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpskipintervalcommand/1622899-preferredintervals?language=objc
func (s_ SkipIntervalCommand) PreferredIntervals() []foundation.Number {
	rv := objc.Call[[]foundation.Number](s_, objc.Sel("preferredIntervals"))
	return rv
}

// The available skip intervals, in seconds, for a media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpskipintervalcommand/1622899-preferredintervals?language=objc
func (s_ SkipIntervalCommand) SetPreferredIntervals(value []foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setPreferredIntervals:"), value)
}
