// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DelegatingPlaybackCoordinatorBufferingCommand] class.
var DelegatingPlaybackCoordinatorBufferingCommandClass = _DelegatingPlaybackCoordinatorBufferingCommandClass{objc.GetClass("AVDelegatingPlaybackCoordinatorBufferingCommand")}

type _DelegatingPlaybackCoordinatorBufferingCommandClass struct {
	objc.Class
}

// An interface definition for the [DelegatingPlaybackCoordinatorBufferingCommand] class.
type IDelegatingPlaybackCoordinatorBufferingCommand interface {
	IDelegatingPlaybackCoordinatorPlaybackControlCommand
	AnticipatedPlaybackRate() float64
	CompletionDueDate() foundation.Date
}

// A command that indicates to start buffering data in preparation for playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorbufferingcommand?language=objc
type DelegatingPlaybackCoordinatorBufferingCommand struct {
	DelegatingPlaybackCoordinatorPlaybackControlCommand
}

func DelegatingPlaybackCoordinatorBufferingCommandFrom(ptr unsafe.Pointer) DelegatingPlaybackCoordinatorBufferingCommand {
	return DelegatingPlaybackCoordinatorBufferingCommand{
		DelegatingPlaybackCoordinatorPlaybackControlCommand: DelegatingPlaybackCoordinatorPlaybackControlCommandFrom(ptr),
	}
}

func (dc _DelegatingPlaybackCoordinatorBufferingCommandClass) Alloc() DelegatingPlaybackCoordinatorBufferingCommand {
	rv := objc.Call[DelegatingPlaybackCoordinatorBufferingCommand](dc, objc.Sel("alloc"))
	return rv
}

func DelegatingPlaybackCoordinatorBufferingCommand_Alloc() DelegatingPlaybackCoordinatorBufferingCommand {
	return DelegatingPlaybackCoordinatorBufferingCommandClass.Alloc()
}

func (dc _DelegatingPlaybackCoordinatorBufferingCommandClass) New() DelegatingPlaybackCoordinatorBufferingCommand {
	rv := objc.Call[DelegatingPlaybackCoordinatorBufferingCommand](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDelegatingPlaybackCoordinatorBufferingCommand() DelegatingPlaybackCoordinatorBufferingCommand {
	return DelegatingPlaybackCoordinatorBufferingCommandClass.New()
}

func (d_ DelegatingPlaybackCoordinatorBufferingCommand) Init() DelegatingPlaybackCoordinatorBufferingCommand {
	rv := objc.Call[DelegatingPlaybackCoordinatorBufferingCommand](d_, objc.Sel("init"))
	return rv
}

// The rate at which the coordinator expects the current item to play. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorbufferingcommand/3750261-anticipatedplaybackrate?language=objc
func (d_ DelegatingPlaybackCoordinatorBufferingCommand) AnticipatedPlaybackRate() float64 {
	rv := objc.Call[float64](d_, objc.Sel("anticipatedPlaybackRate"))
	return rv
}

// The deadline by which the coordinator expects the delegate to complete execution of a command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorbufferingcommand/3778530-completionduedate?language=objc
func (d_ DelegatingPlaybackCoordinatorBufferingCommand) CompletionDueDate() foundation.Date {
	rv := objc.Call[foundation.Date](d_, objc.Sel("completionDueDate"))
	return rv
}
