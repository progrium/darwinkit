// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlayerPlaybackCoordinator] class.
var PlayerPlaybackCoordinatorClass = _PlayerPlaybackCoordinatorClass{objc.GetClass("AVPlayerPlaybackCoordinator")}

type _PlayerPlaybackCoordinatorClass struct {
	objc.Class
}

// An interface definition for the [PlayerPlaybackCoordinator] class.
type IPlayerPlaybackCoordinator interface {
	IPlaybackCoordinator
	Player() Player
	Delegate() PlayerPlaybackCoordinatorDelegateWrapper
	SetDelegate(value PPlayerPlaybackCoordinatorDelegate)
	SetDelegateObject(valueObject objc.IObject)
}

// A playback coordinator subclass that coordinates the playback of player objects in a connected group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerplaybackcoordinator?language=objc
type PlayerPlaybackCoordinator struct {
	PlaybackCoordinator
}

func PlayerPlaybackCoordinatorFrom(ptr unsafe.Pointer) PlayerPlaybackCoordinator {
	return PlayerPlaybackCoordinator{
		PlaybackCoordinator: PlaybackCoordinatorFrom(ptr),
	}
}

func (pc _PlayerPlaybackCoordinatorClass) Alloc() PlayerPlaybackCoordinator {
	rv := objc.Call[PlayerPlaybackCoordinator](pc, objc.Sel("alloc"))
	return rv
}

func PlayerPlaybackCoordinator_Alloc() PlayerPlaybackCoordinator {
	return PlayerPlaybackCoordinatorClass.Alloc()
}

func (pc _PlayerPlaybackCoordinatorClass) New() PlayerPlaybackCoordinator {
	rv := objc.Call[PlayerPlaybackCoordinator](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerPlaybackCoordinator() PlayerPlaybackCoordinator {
	return PlayerPlaybackCoordinatorClass.New()
}

func (p_ PlayerPlaybackCoordinator) Init() PlayerPlaybackCoordinator {
	rv := objc.Call[PlayerPlaybackCoordinator](p_, objc.Sel("init"))
	return rv
}

// A player that participates in coordinated playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerplaybackcoordinator/3750302-player?language=objc
func (p_ PlayerPlaybackCoordinator) Player() Player {
	rv := objc.Call[Player](p_, objc.Sel("player"))
	return rv
}

// A delegate object for the playback coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerplaybackcoordinator/3750301-delegate?language=objc
func (p_ PlayerPlaybackCoordinator) Delegate() PlayerPlaybackCoordinatorDelegateWrapper {
	rv := objc.Call[PlayerPlaybackCoordinatorDelegateWrapper](p_, objc.Sel("delegate"))
	return rv
}

// A delegate object for the playback coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerplaybackcoordinator/3750301-delegate?language=objc
func (p_ PlayerPlaybackCoordinator) SetDelegate(value PPlayerPlaybackCoordinatorDelegate) {
	po0 := objc.WrapAsProtocol("AVPlayerPlaybackCoordinatorDelegate", value)
	objc.SetAssociatedObject(p_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:"), po0)
}

// A delegate object for the playback coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerplaybackcoordinator/3750301-delegate?language=objc
func (p_ PlayerPlaybackCoordinator) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}
