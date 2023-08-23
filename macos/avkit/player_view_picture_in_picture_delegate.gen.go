// AUTO-GENERATED CODE, DO NOT MODIFY

package avkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines the methods to implement to respond to Picture in Picture playback events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate?language=objc
type PPlayerViewPictureInPictureDelegate interface {
	// optional
	PlayerViewWillStartPictureInPicture(playerView PlayerView)
	HasPlayerViewWillStartPictureInPicture() bool

	// optional
	PlayerViewDidStopPictureInPicture(playerView PlayerView)
	HasPlayerViewDidStopPictureInPicture() bool

	// optional
	PlayerViewDidStartPictureInPicture(playerView PlayerView)
	HasPlayerViewDidStartPictureInPicture() bool

	// optional
	PlayerViewShouldAutomaticallyDismissAtPictureInPictureStart(playerView PlayerView) bool
	HasPlayerViewShouldAutomaticallyDismissAtPictureInPictureStart() bool

	// optional
	PlayerViewWillStopPictureInPicture(playerView PlayerView)
	HasPlayerViewWillStopPictureInPicture() bool

	// optional
	PlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler(playerView PlayerView, completionHandler func(restored bool))
	HasPlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler() bool
}

// A delegate implementation builder for the [PPlayerViewPictureInPictureDelegate] protocol.
type PlayerViewPictureInPictureDelegate struct {
	_PlayerViewWillStartPictureInPicture                                        func(playerView PlayerView)
	_PlayerViewDidStopPictureInPicture                                          func(playerView PlayerView)
	_PlayerViewDidStartPictureInPicture                                         func(playerView PlayerView)
	_PlayerViewShouldAutomaticallyDismissAtPictureInPictureStart                func(playerView PlayerView) bool
	_PlayerViewWillStopPictureInPicture                                         func(playerView PlayerView)
	_PlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler func(playerView PlayerView, completionHandler func(restored bool))
}

func (di *PlayerViewPictureInPictureDelegate) HasPlayerViewWillStartPictureInPicture() bool {
	return di._PlayerViewWillStartPictureInPicture != nil
}

// Tells the delegate that Picture in Picture playback is about to start. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate/3172696-playerviewwillstartpictureinpict?language=objc
func (di *PlayerViewPictureInPictureDelegate) SetPlayerViewWillStartPictureInPicture(f func(playerView PlayerView)) {
	di._PlayerViewWillStartPictureInPicture = f
}

// Tells the delegate that Picture in Picture playback is about to start. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate/3172696-playerviewwillstartpictureinpict?language=objc
func (di *PlayerViewPictureInPictureDelegate) PlayerViewWillStartPictureInPicture(playerView PlayerView) {
	di._PlayerViewWillStartPictureInPicture(playerView)
}
func (di *PlayerViewPictureInPictureDelegate) HasPlayerViewDidStopPictureInPicture() bool {
	return di._PlayerViewDidStopPictureInPicture != nil
}

// Tells the delegate that Picture in Picture playback stopped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate/3172694-playerviewdidstoppictureinpictur?language=objc
func (di *PlayerViewPictureInPictureDelegate) SetPlayerViewDidStopPictureInPicture(f func(playerView PlayerView)) {
	di._PlayerViewDidStopPictureInPicture = f
}

// Tells the delegate that Picture in Picture playback stopped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate/3172694-playerviewdidstoppictureinpictur?language=objc
func (di *PlayerViewPictureInPictureDelegate) PlayerViewDidStopPictureInPicture(playerView PlayerView) {
	di._PlayerViewDidStopPictureInPicture(playerView)
}
func (di *PlayerViewPictureInPictureDelegate) HasPlayerViewDidStartPictureInPicture() bool {
	return di._PlayerViewDidStartPictureInPicture != nil
}

// Tells the delegate that Picture in Picture playback started. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate/3172693-playerviewdidstartpictureinpictu?language=objc
func (di *PlayerViewPictureInPictureDelegate) SetPlayerViewDidStartPictureInPicture(f func(playerView PlayerView)) {
	di._PlayerViewDidStartPictureInPicture = f
}

// Tells the delegate that Picture in Picture playback started. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate/3172693-playerviewdidstartpictureinpictu?language=objc
func (di *PlayerViewPictureInPictureDelegate) PlayerViewDidStartPictureInPicture(playerView PlayerView) {
	di._PlayerViewDidStartPictureInPicture(playerView)
}
func (di *PlayerViewPictureInPictureDelegate) HasPlayerViewShouldAutomaticallyDismissAtPictureInPictureStart() bool {
	return di._PlayerViewShouldAutomaticallyDismissAtPictureInPictureStart != nil
}

// Asks the delegate if the player view should miniaturize when Picture in Picture starts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate/3172695-playerviewshouldautomaticallydis?language=objc
func (di *PlayerViewPictureInPictureDelegate) SetPlayerViewShouldAutomaticallyDismissAtPictureInPictureStart(f func(playerView PlayerView) bool) {
	di._PlayerViewShouldAutomaticallyDismissAtPictureInPictureStart = f
}

// Asks the delegate if the player view should miniaturize when Picture in Picture starts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate/3172695-playerviewshouldautomaticallydis?language=objc
func (di *PlayerViewPictureInPictureDelegate) PlayerViewShouldAutomaticallyDismissAtPictureInPictureStart(playerView PlayerView) bool {
	return di._PlayerViewShouldAutomaticallyDismissAtPictureInPictureStart(playerView)
}
func (di *PlayerViewPictureInPictureDelegate) HasPlayerViewWillStopPictureInPicture() bool {
	return di._PlayerViewWillStopPictureInPicture != nil
}

// Tells the delegate that Picture in Picture playback is about to stop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate/3172697-playerviewwillstoppictureinpictu?language=objc
func (di *PlayerViewPictureInPictureDelegate) SetPlayerViewWillStopPictureInPicture(f func(playerView PlayerView)) {
	di._PlayerViewWillStopPictureInPicture = f
}

// Tells the delegate that Picture in Picture playback is about to stop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate/3172697-playerviewwillstoppictureinpictu?language=objc
func (di *PlayerViewPictureInPictureDelegate) PlayerViewWillStopPictureInPicture(playerView PlayerView) {
	di._PlayerViewWillStopPictureInPicture(playerView)
}
func (di *PlayerViewPictureInPictureDelegate) HasPlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler() bool {
	return di._PlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler != nil
}

// Tells the delegate to restore the user interface before Picture in Picture playback stops. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate/3172692-playerview?language=objc
func (di *PlayerViewPictureInPictureDelegate) SetPlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler(f func(playerView PlayerView, completionHandler func(restored bool))) {
	di._PlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler = f
}

// Tells the delegate to restore the user interface before Picture in Picture playback stops. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate/3172692-playerview?language=objc
func (di *PlayerViewPictureInPictureDelegate) PlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler(playerView PlayerView, completionHandler func(restored bool)) {
	di._PlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler(playerView, completionHandler)
}

// A concrete type wrapper for the [PPlayerViewPictureInPictureDelegate] protocol.
type PlayerViewPictureInPictureDelegateWrapper struct {
	objc.Object
}

func (p_ PlayerViewPictureInPictureDelegateWrapper) HasPlayerViewWillStartPictureInPicture() bool {
	return p_.RespondsToSelector(objc.Sel("playerViewWillStartPictureInPicture:"))
}

// Tells the delegate that Picture in Picture playback is about to start. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate/3172696-playerviewwillstartpictureinpict?language=objc
func (p_ PlayerViewPictureInPictureDelegateWrapper) PlayerViewWillStartPictureInPicture(playerView IPlayerView) {
	objc.Call[objc.Void](p_, objc.Sel("playerViewWillStartPictureInPicture:"), objc.Ptr(playerView))
}

func (p_ PlayerViewPictureInPictureDelegateWrapper) HasPlayerViewDidStopPictureInPicture() bool {
	return p_.RespondsToSelector(objc.Sel("playerViewDidStopPictureInPicture:"))
}

// Tells the delegate that Picture in Picture playback stopped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate/3172694-playerviewdidstoppictureinpictur?language=objc
func (p_ PlayerViewPictureInPictureDelegateWrapper) PlayerViewDidStopPictureInPicture(playerView IPlayerView) {
	objc.Call[objc.Void](p_, objc.Sel("playerViewDidStopPictureInPicture:"), objc.Ptr(playerView))
}

func (p_ PlayerViewPictureInPictureDelegateWrapper) HasPlayerViewDidStartPictureInPicture() bool {
	return p_.RespondsToSelector(objc.Sel("playerViewDidStartPictureInPicture:"))
}

// Tells the delegate that Picture in Picture playback started. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate/3172693-playerviewdidstartpictureinpictu?language=objc
func (p_ PlayerViewPictureInPictureDelegateWrapper) PlayerViewDidStartPictureInPicture(playerView IPlayerView) {
	objc.Call[objc.Void](p_, objc.Sel("playerViewDidStartPictureInPicture:"), objc.Ptr(playerView))
}

func (p_ PlayerViewPictureInPictureDelegateWrapper) HasPlayerViewShouldAutomaticallyDismissAtPictureInPictureStart() bool {
	return p_.RespondsToSelector(objc.Sel("playerViewShouldAutomaticallyDismissAtPictureInPictureStart:"))
}

// Asks the delegate if the player view should miniaturize when Picture in Picture starts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate/3172695-playerviewshouldautomaticallydis?language=objc
func (p_ PlayerViewPictureInPictureDelegateWrapper) PlayerViewShouldAutomaticallyDismissAtPictureInPictureStart(playerView IPlayerView) bool {
	rv := objc.Call[bool](p_, objc.Sel("playerViewShouldAutomaticallyDismissAtPictureInPictureStart:"), objc.Ptr(playerView))
	return rv
}

func (p_ PlayerViewPictureInPictureDelegateWrapper) HasPlayerViewWillStopPictureInPicture() bool {
	return p_.RespondsToSelector(objc.Sel("playerViewWillStopPictureInPicture:"))
}

// Tells the delegate that Picture in Picture playback is about to stop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate/3172697-playerviewwillstoppictureinpictu?language=objc
func (p_ PlayerViewPictureInPictureDelegateWrapper) PlayerViewWillStopPictureInPicture(playerView IPlayerView) {
	objc.Call[objc.Void](p_, objc.Sel("playerViewWillStopPictureInPicture:"), objc.Ptr(playerView))
}

func (p_ PlayerViewPictureInPictureDelegateWrapper) HasPlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler() bool {
	return p_.RespondsToSelector(objc.Sel("playerView:restoreUserInterfaceForPictureInPictureStopWithCompletionHandler:"))
}

// Tells the delegate to restore the user interface before Picture in Picture playback stops. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewpictureinpicturedelegate/3172692-playerview?language=objc
func (p_ PlayerViewPictureInPictureDelegateWrapper) PlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler(playerView IPlayerView, completionHandler func(restored bool)) {
	objc.Call[objc.Void](p_, objc.Sel("playerView:restoreUserInterfaceForPictureInPictureStopWithCompletionHandler:"), objc.Ptr(playerView), completionHandler)
}
