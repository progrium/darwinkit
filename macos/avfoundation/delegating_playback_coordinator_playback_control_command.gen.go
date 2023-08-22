// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DelegatingPlaybackCoordinatorPlaybackControlCommand] class.
var DelegatingPlaybackCoordinatorPlaybackControlCommandClass = _DelegatingPlaybackCoordinatorPlaybackControlCommandClass{objc.GetClass("AVDelegatingPlaybackCoordinatorPlaybackControlCommand")}

type _DelegatingPlaybackCoordinatorPlaybackControlCommandClass struct {
	objc.Class
}

// An interface definition for the [DelegatingPlaybackCoordinatorPlaybackControlCommand] class.
type IDelegatingPlaybackCoordinatorPlaybackControlCommand interface {
	objc.IObject
	Originator() CoordinatedPlaybackParticipant
	ExpectedCurrentItemIdentifier() string
}

// An abstract superclass for playback commands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorplaybackcontrolcommand?language=objc
type DelegatingPlaybackCoordinatorPlaybackControlCommand struct {
	objc.Object
}

func DelegatingPlaybackCoordinatorPlaybackControlCommandFrom(ptr unsafe.Pointer) DelegatingPlaybackCoordinatorPlaybackControlCommand {
	return DelegatingPlaybackCoordinatorPlaybackControlCommand{
		Object: objc.ObjectFrom(ptr),
	}
}

func (dc _DelegatingPlaybackCoordinatorPlaybackControlCommandClass) Alloc() DelegatingPlaybackCoordinatorPlaybackControlCommand {
	rv := objc.Call[DelegatingPlaybackCoordinatorPlaybackControlCommand](dc, objc.Sel("alloc"))
	return rv
}

func DelegatingPlaybackCoordinatorPlaybackControlCommand_Alloc() DelegatingPlaybackCoordinatorPlaybackControlCommand {
	return DelegatingPlaybackCoordinatorPlaybackControlCommandClass.Alloc()
}

func (dc _DelegatingPlaybackCoordinatorPlaybackControlCommandClass) New() DelegatingPlaybackCoordinatorPlaybackControlCommand {
	rv := objc.Call[DelegatingPlaybackCoordinatorPlaybackControlCommand](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDelegatingPlaybackCoordinatorPlaybackControlCommand() DelegatingPlaybackCoordinatorPlaybackControlCommand {
	return DelegatingPlaybackCoordinatorPlaybackControlCommandClass.New()
}

func (d_ DelegatingPlaybackCoordinatorPlaybackControlCommand) Init() DelegatingPlaybackCoordinatorPlaybackControlCommand {
	rv := objc.Call[DelegatingPlaybackCoordinatorPlaybackControlCommand](d_, objc.Sel("init"))
	return rv
}

// The participant that causes the coordinator to issue the command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorplaybackcontrolcommand/3750271-originator?language=objc
func (d_ DelegatingPlaybackCoordinatorPlaybackControlCommand) Originator() CoordinatedPlaybackParticipant {
	rv := objc.Call[CoordinatedPlaybackParticipant](d_, objc.Sel("originator"))
	return rv
}

// An item identifier the coordinator issues the command for. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorplaybackcontrolcommand/3750270-expectedcurrentitemidentifier?language=objc
func (d_ DelegatingPlaybackCoordinatorPlaybackControlCommand) ExpectedCurrentItemIdentifier() string {
	rv := objc.Call[string](d_, objc.Sel("expectedCurrentItemIdentifier"))
	return rv
}
