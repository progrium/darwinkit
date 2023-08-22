// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangeShuffleModeCommand] class.
var ChangeShuffleModeCommandClass = _ChangeShuffleModeCommandClass{objc.GetClass("MPChangeShuffleModeCommand")}

type _ChangeShuffleModeCommandClass struct {
	objc.Class
}

// An interface definition for the [ChangeShuffleModeCommand] class.
type IChangeShuffleModeCommand interface {
	IRemoteCommand
	CurrentShuffleType() ShuffleType
	SetCurrentShuffleType(value ShuffleType)
}

// An object that responds to requests to change the current shuffle mode used during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeshufflemodecommand?language=objc
type ChangeShuffleModeCommand struct {
	RemoteCommand
}

func ChangeShuffleModeCommandFrom(ptr unsafe.Pointer) ChangeShuffleModeCommand {
	return ChangeShuffleModeCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}

func (cc _ChangeShuffleModeCommandClass) Alloc() ChangeShuffleModeCommand {
	rv := objc.Call[ChangeShuffleModeCommand](cc, objc.Sel("alloc"))
	return rv
}

func ChangeShuffleModeCommand_Alloc() ChangeShuffleModeCommand {
	return ChangeShuffleModeCommandClass.Alloc()
}

func (cc _ChangeShuffleModeCommandClass) New() ChangeShuffleModeCommand {
	rv := objc.Call[ChangeShuffleModeCommand](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangeShuffleModeCommand() ChangeShuffleModeCommand {
	return ChangeShuffleModeCommandClass.New()
}

func (c_ ChangeShuffleModeCommand) Init() ChangeShuffleModeCommand {
	rv := objc.Call[ChangeShuffleModeCommand](c_, objc.Sel("init"))
	return rv
}

// The current shuffle mode for a media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeshufflemodecommand/1648341-currentshuffletype?language=objc
func (c_ ChangeShuffleModeCommand) CurrentShuffleType() ShuffleType {
	rv := objc.Call[ShuffleType](c_, objc.Sel("currentShuffleType"))
	return rv
}

// The current shuffle mode for a media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeshufflemodecommand/1648341-currentshuffletype?language=objc
func (c_ ChangeShuffleModeCommand) SetCurrentShuffleType(value ShuffleType) {
	objc.Call[objc.Void](c_, objc.Sel("setCurrentShuffleType:"), value)
}
