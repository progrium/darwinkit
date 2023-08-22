// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangeRepeatModeCommand] class.
var ChangeRepeatModeCommandClass = _ChangeRepeatModeCommandClass{objc.GetClass("MPChangeRepeatModeCommand")}

type _ChangeRepeatModeCommandClass struct {
	objc.Class
}

// An interface definition for the [ChangeRepeatModeCommand] class.
type IChangeRepeatModeCommand interface {
	IRemoteCommand
	CurrentRepeatType() RepeatType
	SetCurrentRepeatType(value RepeatType)
}

// An object that responds to requests to change the current repeat mode used during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangerepeatmodecommand?language=objc
type ChangeRepeatModeCommand struct {
	RemoteCommand
}

func ChangeRepeatModeCommandFrom(ptr unsafe.Pointer) ChangeRepeatModeCommand {
	return ChangeRepeatModeCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}

func (cc _ChangeRepeatModeCommandClass) Alloc() ChangeRepeatModeCommand {
	rv := objc.Call[ChangeRepeatModeCommand](cc, objc.Sel("alloc"))
	return rv
}

func ChangeRepeatModeCommand_Alloc() ChangeRepeatModeCommand {
	return ChangeRepeatModeCommandClass.Alloc()
}

func (cc _ChangeRepeatModeCommandClass) New() ChangeRepeatModeCommand {
	rv := objc.Call[ChangeRepeatModeCommand](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangeRepeatModeCommand() ChangeRepeatModeCommand {
	return ChangeRepeatModeCommandClass.New()
}

func (c_ ChangeRepeatModeCommand) Init() ChangeRepeatModeCommand {
	rv := objc.Call[ChangeRepeatModeCommand](c_, objc.Sel("init"))
	return rv
}

// The current repeat option for a media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangerepeatmodecommand/1648342-currentrepeattype?language=objc
func (c_ ChangeRepeatModeCommand) CurrentRepeatType() RepeatType {
	rv := objc.Call[RepeatType](c_, objc.Sel("currentRepeatType"))
	return rv
}

// The current repeat option for a media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangerepeatmodecommand/1648342-currentrepeattype?language=objc
func (c_ ChangeRepeatModeCommand) SetCurrentRepeatType(value RepeatType) {
	objc.Call[objc.Void](c_, objc.Sel("setCurrentRepeatType:"), value)
}
