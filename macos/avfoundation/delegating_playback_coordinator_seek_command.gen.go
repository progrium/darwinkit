// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DelegatingPlaybackCoordinatorSeekCommand] class.
var DelegatingPlaybackCoordinatorSeekCommandClass = _DelegatingPlaybackCoordinatorSeekCommandClass{objc.GetClass("AVDelegatingPlaybackCoordinatorSeekCommand")}

type _DelegatingPlaybackCoordinatorSeekCommandClass struct {
	objc.Class
}

// An interface definition for the [DelegatingPlaybackCoordinatorSeekCommand] class.
type IDelegatingPlaybackCoordinatorSeekCommand interface {
	IDelegatingPlaybackCoordinatorPlaybackControlCommand
	ShouldBufferInAnticipationOfPlayback() bool
	ItemTime() coremedia.Time
	AnticipatedPlaybackRate() float64
	CompletionDueDate() foundation.Date
}

// A command that indicates to seek to a new time in the item timeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorseekcommand?language=objc
type DelegatingPlaybackCoordinatorSeekCommand struct {
	DelegatingPlaybackCoordinatorPlaybackControlCommand
}

func DelegatingPlaybackCoordinatorSeekCommandFrom(ptr unsafe.Pointer) DelegatingPlaybackCoordinatorSeekCommand {
	return DelegatingPlaybackCoordinatorSeekCommand{
		DelegatingPlaybackCoordinatorPlaybackControlCommand: DelegatingPlaybackCoordinatorPlaybackControlCommandFrom(ptr),
	}
}

func (dc _DelegatingPlaybackCoordinatorSeekCommandClass) Alloc() DelegatingPlaybackCoordinatorSeekCommand {
	rv := objc.Call[DelegatingPlaybackCoordinatorSeekCommand](dc, objc.Sel("alloc"))
	return rv
}

func DelegatingPlaybackCoordinatorSeekCommand_Alloc() DelegatingPlaybackCoordinatorSeekCommand {
	return DelegatingPlaybackCoordinatorSeekCommandClass.Alloc()
}

func (dc _DelegatingPlaybackCoordinatorSeekCommandClass) New() DelegatingPlaybackCoordinatorSeekCommand {
	rv := objc.Call[DelegatingPlaybackCoordinatorSeekCommand](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDelegatingPlaybackCoordinatorSeekCommand() DelegatingPlaybackCoordinatorSeekCommand {
	return DelegatingPlaybackCoordinatorSeekCommandClass.New()
}

func (d_ DelegatingPlaybackCoordinatorSeekCommand) Init() DelegatingPlaybackCoordinatorSeekCommand {
	rv := objc.Call[DelegatingPlaybackCoordinatorSeekCommand](d_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the player starts buffering in anticipation of a request to begin playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorseekcommand/3750281-shouldbufferinanticipationofplay?language=objc
func (d_ DelegatingPlaybackCoordinatorSeekCommand) ShouldBufferInAnticipationOfPlayback() bool {
	rv := objc.Call[bool](d_, objc.Sel("shouldBufferInAnticipationOfPlayback"))
	return rv
}

// The time to seek to in the item timeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorseekcommand/3750280-itemtime?language=objc
func (d_ DelegatingPlaybackCoordinatorSeekCommand) ItemTime() coremedia.Time {
	rv := objc.Call[coremedia.Time](d_, objc.Sel("itemTime"))
	return rv
}

// The rate at which the coordinator expects playback to resume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorseekcommand/3750279-anticipatedplaybackrate?language=objc
func (d_ DelegatingPlaybackCoordinatorSeekCommand) AnticipatedPlaybackRate() float64 {
	rv := objc.Call[float64](d_, objc.Sel("anticipatedPlaybackRate"))
	return rv
}

// The deadline by which the coordinator expects the delegate to handle the command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorseekcommand/3778531-completionduedate?language=objc
func (d_ DelegatingPlaybackCoordinatorSeekCommand) CompletionDueDate() foundation.Date {
	rv := objc.Call[foundation.Date](d_, objc.Sel("completionDueDate"))
	return rv
}
