// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines the method to implement to respond to playback commands from the playback coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinatorplaybackcontroldelegate?language=objc
type PPlaybackCoordinatorPlaybackControlDelegate interface {
	// optional
	PlaybackCoordinatorDidIssuePlayCommandCompletionHandler(coordinator DelegatingPlaybackCoordinator, playCommand DelegatingPlaybackCoordinatorPlayCommand, completionHandler func())
	HasPlaybackCoordinatorDidIssuePlayCommandCompletionHandler() bool
}

// A delegate implementation builder for the [PPlaybackCoordinatorPlaybackControlDelegate] protocol.
type PlaybackCoordinatorPlaybackControlDelegate struct {
	_PlaybackCoordinatorDidIssuePlayCommandCompletionHandler func(coordinator DelegatingPlaybackCoordinator, playCommand DelegatingPlaybackCoordinatorPlayCommand, completionHandler func())
}

func (di *PlaybackCoordinatorPlaybackControlDelegate) HasPlaybackCoordinatorDidIssuePlayCommandCompletionHandler() bool {
	return di._PlaybackCoordinatorDidIssuePlayCommandCompletionHandler != nil
}

// Tells the delegate to match the playback rate to that of the group when the rate is nonzero. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinatorplaybackcontroldelegate/3750296-playbackcoordinator?language=objc
func (di *PlaybackCoordinatorPlaybackControlDelegate) SetPlaybackCoordinatorDidIssuePlayCommandCompletionHandler(f func(coordinator DelegatingPlaybackCoordinator, playCommand DelegatingPlaybackCoordinatorPlayCommand, completionHandler func())) {
	di._PlaybackCoordinatorDidIssuePlayCommandCompletionHandler = f
}

// Tells the delegate to match the playback rate to that of the group when the rate is nonzero. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinatorplaybackcontroldelegate/3750296-playbackcoordinator?language=objc
func (di *PlaybackCoordinatorPlaybackControlDelegate) PlaybackCoordinatorDidIssuePlayCommandCompletionHandler(coordinator DelegatingPlaybackCoordinator, playCommand DelegatingPlaybackCoordinatorPlayCommand, completionHandler func()) {
	di._PlaybackCoordinatorDidIssuePlayCommandCompletionHandler(coordinator, playCommand, completionHandler)
}

// A concrete type wrapper for the [PPlaybackCoordinatorPlaybackControlDelegate] protocol.
type PlaybackCoordinatorPlaybackControlDelegateWrapper struct {
	objc.Object
}

func (p_ PlaybackCoordinatorPlaybackControlDelegateWrapper) HasPlaybackCoordinatorDidIssuePlayCommandCompletionHandler() bool {
	return p_.RespondsToSelector(objc.Sel("playbackCoordinator:didIssuePlayCommand:completionHandler:"))
}

// Tells the delegate to match the playback rate to that of the group when the rate is nonzero. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinatorplaybackcontroldelegate/3750296-playbackcoordinator?language=objc
func (p_ PlaybackCoordinatorPlaybackControlDelegateWrapper) PlaybackCoordinatorDidIssuePlayCommandCompletionHandler(coordinator IDelegatingPlaybackCoordinator, playCommand IDelegatingPlaybackCoordinatorPlayCommand, completionHandler func()) {
	objc.Call[objc.Void](p_, objc.Sel("playbackCoordinator:didIssuePlayCommand:completionHandler:"), objc.Ptr(coordinator), objc.Ptr(playCommand), completionHandler)
}
