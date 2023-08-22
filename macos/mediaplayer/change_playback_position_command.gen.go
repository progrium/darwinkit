// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangePlaybackPositionCommand] class.
var ChangePlaybackPositionCommandClass = _ChangePlaybackPositionCommandClass{objc.GetClass("MPChangePlaybackPositionCommand")}

type _ChangePlaybackPositionCommandClass struct {
	objc.Class
}

// An interface definition for the [ChangePlaybackPositionCommand] class.
type IChangePlaybackPositionCommand interface {
	IRemoteCommand
}

// An object that responds to requests to change the current playback position of the playing item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeplaybackpositioncommand?language=objc
type ChangePlaybackPositionCommand struct {
	RemoteCommand
}

func ChangePlaybackPositionCommandFrom(ptr unsafe.Pointer) ChangePlaybackPositionCommand {
	return ChangePlaybackPositionCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}

func (cc _ChangePlaybackPositionCommandClass) Alloc() ChangePlaybackPositionCommand {
	rv := objc.Call[ChangePlaybackPositionCommand](cc, objc.Sel("alloc"))
	return rv
}

func ChangePlaybackPositionCommand_Alloc() ChangePlaybackPositionCommand {
	return ChangePlaybackPositionCommandClass.Alloc()
}

func (cc _ChangePlaybackPositionCommandClass) New() ChangePlaybackPositionCommand {
	rv := objc.Call[ChangePlaybackPositionCommand](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangePlaybackPositionCommand() ChangePlaybackPositionCommand {
	return ChangePlaybackPositionCommandClass.New()
}

func (c_ ChangePlaybackPositionCommand) Init() ChangePlaybackPositionCommand {
	rv := objc.Call[ChangePlaybackPositionCommand](c_, objc.Sel("init"))
	return rv
}
