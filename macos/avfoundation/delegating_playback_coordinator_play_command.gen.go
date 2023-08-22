// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DelegatingPlaybackCoordinatorPlayCommand] class.
var DelegatingPlaybackCoordinatorPlayCommandClass = _DelegatingPlaybackCoordinatorPlayCommandClass{objc.GetClass("AVDelegatingPlaybackCoordinatorPlayCommand")}

type _DelegatingPlaybackCoordinatorPlayCommandClass struct {
	objc.Class
}

// An interface definition for the [DelegatingPlaybackCoordinatorPlayCommand] class.
type IDelegatingPlaybackCoordinatorPlayCommand interface {
	IDelegatingPlaybackCoordinatorPlaybackControlCommand
	Rate() float64
	ItemTime() coremedia.Time
	HostClockTime() coremedia.Time
}

// A command that indicates to play at a specific rate and time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorplaycommand?language=objc
type DelegatingPlaybackCoordinatorPlayCommand struct {
	DelegatingPlaybackCoordinatorPlaybackControlCommand
}

func DelegatingPlaybackCoordinatorPlayCommandFrom(ptr unsafe.Pointer) DelegatingPlaybackCoordinatorPlayCommand {
	return DelegatingPlaybackCoordinatorPlayCommand{
		DelegatingPlaybackCoordinatorPlaybackControlCommand: DelegatingPlaybackCoordinatorPlaybackControlCommandFrom(ptr),
	}
}

func (dc _DelegatingPlaybackCoordinatorPlayCommandClass) Alloc() DelegatingPlaybackCoordinatorPlayCommand {
	rv := objc.Call[DelegatingPlaybackCoordinatorPlayCommand](dc, objc.Sel("alloc"))
	return rv
}

func DelegatingPlaybackCoordinatorPlayCommand_Alloc() DelegatingPlaybackCoordinatorPlayCommand {
	return DelegatingPlaybackCoordinatorPlayCommandClass.Alloc()
}

func (dc _DelegatingPlaybackCoordinatorPlayCommandClass) New() DelegatingPlaybackCoordinatorPlayCommand {
	rv := objc.Call[DelegatingPlaybackCoordinatorPlayCommand](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDelegatingPlaybackCoordinatorPlayCommand() DelegatingPlaybackCoordinatorPlayCommand {
	return DelegatingPlaybackCoordinatorPlayCommandClass.New()
}

func (d_ DelegatingPlaybackCoordinatorPlayCommand) Init() DelegatingPlaybackCoordinatorPlayCommand {
	rv := objc.Call[DelegatingPlaybackCoordinatorPlayCommand](d_, objc.Sel("init"))
	return rv
}

// A rate to use when starting playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorplaycommand/3750268-rate?language=objc
func (d_ DelegatingPlaybackCoordinatorPlayCommand) Rate() float64 {
	rv := objc.Call[float64](d_, objc.Sel("rate"))
	return rv
}

// A time in the item timeline to use to begin playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorplaycommand/3750267-itemtime?language=objc
func (d_ DelegatingPlaybackCoordinatorPlayCommand) ItemTime() coremedia.Time {
	rv := objc.Call[coremedia.Time](d_, objc.Sel("itemTime"))
	return rv
}

// A host clock time to use to begin playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorplaycommand/3750266-hostclocktime?language=objc
func (d_ DelegatingPlaybackCoordinatorPlayCommand) HostClockTime() coremedia.Time {
	rv := objc.Call[coremedia.Time](d_, objc.Sel("hostClockTime"))
	return rv
}
