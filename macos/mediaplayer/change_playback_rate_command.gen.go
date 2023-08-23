// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangePlaybackRateCommand] class.
var ChangePlaybackRateCommandClass = _ChangePlaybackRateCommandClass{objc.GetClass("MPChangePlaybackRateCommand")}

type _ChangePlaybackRateCommandClass struct {
	objc.Class
}

// An interface definition for the [ChangePlaybackRateCommand] class.
type IChangePlaybackRateCommand interface {
	IRemoteCommand
	SupportedPlaybackRates() []foundation.Number
	SetSupportedPlaybackRates(value []foundation.INumber)
}

// An object that responds to requests to change the playback rate of the playing item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeplaybackratecommand?language=objc
type ChangePlaybackRateCommand struct {
	RemoteCommand
}

func ChangePlaybackRateCommandFrom(ptr unsafe.Pointer) ChangePlaybackRateCommand {
	return ChangePlaybackRateCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}

func (cc _ChangePlaybackRateCommandClass) Alloc() ChangePlaybackRateCommand {
	rv := objc.Call[ChangePlaybackRateCommand](cc, objc.Sel("alloc"))
	return rv
}

func ChangePlaybackRateCommand_Alloc() ChangePlaybackRateCommand {
	return ChangePlaybackRateCommandClass.Alloc()
}

func (cc _ChangePlaybackRateCommandClass) New() ChangePlaybackRateCommand {
	rv := objc.Call[ChangePlaybackRateCommand](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangePlaybackRateCommand() ChangePlaybackRateCommand {
	return ChangePlaybackRateCommandClass.New()
}

func (c_ ChangePlaybackRateCommand) Init() ChangePlaybackRateCommand {
	rv := objc.Call[ChangePlaybackRateCommand](c_, objc.Sel("init"))
	return rv
}

// The supported playback rates for a media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeplaybackratecommand/1622915-supportedplaybackrates?language=objc
func (c_ ChangePlaybackRateCommand) SupportedPlaybackRates() []foundation.Number {
	rv := objc.Call[[]foundation.Number](c_, objc.Sel("supportedPlaybackRates"))
	return rv
}

// The supported playback rates for a media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeplaybackratecommand/1622915-supportedplaybackrates?language=objc
func (c_ ChangePlaybackRateCommand) SetSupportedPlaybackRates(value []foundation.INumber) {
	objc.Call[objc.Void](c_, objc.Sel("setSupportedPlaybackRates:"), value)
}
