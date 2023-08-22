// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines the methods to implement to participate in playback coordination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerplaybackcoordinatordelegate?language=objc
type PPlayerPlaybackCoordinatorDelegate interface {
	// optional
	PlaybackCoordinatorIdentifierForPlayerItem(coordinator PlayerPlaybackCoordinator, playerItem PlayerItem) string
	HasPlaybackCoordinatorIdentifierForPlayerItem() bool
}

// A delegate implementation builder for the [PPlayerPlaybackCoordinatorDelegate] protocol.
type PlayerPlaybackCoordinatorDelegate struct {
	_PlaybackCoordinatorIdentifierForPlayerItem func(coordinator PlayerPlaybackCoordinator, playerItem PlayerItem) string
}

func (di *PlayerPlaybackCoordinatorDelegate) HasPlaybackCoordinatorIdentifierForPlayerItem() bool {
	return di._PlaybackCoordinatorIdentifierForPlayerItem != nil
}

// Returns an identifier for a player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerplaybackcoordinatordelegate/3750304-playbackcoordinator?language=objc
func (di *PlayerPlaybackCoordinatorDelegate) SetPlaybackCoordinatorIdentifierForPlayerItem(f func(coordinator PlayerPlaybackCoordinator, playerItem PlayerItem) string) {
	di._PlaybackCoordinatorIdentifierForPlayerItem = f
}

// Returns an identifier for a player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerplaybackcoordinatordelegate/3750304-playbackcoordinator?language=objc
func (di *PlayerPlaybackCoordinatorDelegate) PlaybackCoordinatorIdentifierForPlayerItem(coordinator PlayerPlaybackCoordinator, playerItem PlayerItem) string {
	return di._PlaybackCoordinatorIdentifierForPlayerItem(coordinator, playerItem)
}

// A concrete type wrapper for the [PPlayerPlaybackCoordinatorDelegate] protocol.
type PlayerPlaybackCoordinatorDelegateWrapper struct {
	objc.Object
}

func (p_ PlayerPlaybackCoordinatorDelegateWrapper) HasPlaybackCoordinatorIdentifierForPlayerItem() bool {
	return p_.RespondsToSelector(objc.Sel("playbackCoordinator:identifierForPlayerItem:"))
}

// Returns an identifier for a player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerplaybackcoordinatordelegate/3750304-playbackcoordinator?language=objc
func (p_ PlayerPlaybackCoordinatorDelegateWrapper) PlaybackCoordinatorIdentifierForPlayerItem(coordinator IPlayerPlaybackCoordinator, playerItem IPlayerItem) string {
	rv := objc.Call[string](p_, objc.Sel("playbackCoordinator:identifierForPlayerItem:"), objc.Ptr(coordinator), objc.Ptr(playerItem))
	return rv
}
