// AUTO-GENERATED CODE, DO NOT MODIFY

package avkit

import (
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// A protocol for controlling playback from a sample buffer display layer in Picture in Picture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturesamplebufferplaybackdelegate?language=objc
type PPictureInPictureSampleBufferPlaybackDelegate interface {
	// optional
	PictureInPictureControllerShouldProhibitBackgroundAudioPlayback(pictureInPictureController PictureInPictureController) bool
	HasPictureInPictureControllerShouldProhibitBackgroundAudioPlayback() bool

	// optional
	PictureInPictureControllerIsPlaybackPaused(pictureInPictureController PictureInPictureController) bool
	HasPictureInPictureControllerIsPlaybackPaused() bool

	// optional
	PictureInPictureControllerTimeRangeForPlayback(pictureInPictureController PictureInPictureController) coremedia.TimeRange
	HasPictureInPictureControllerTimeRangeForPlayback() bool

	// optional
	PictureInPictureControllerSetPlaying(pictureInPictureController PictureInPictureController, playing bool)
	HasPictureInPictureControllerSetPlaying() bool
}

// A delegate implementation builder for the [PPictureInPictureSampleBufferPlaybackDelegate] protocol.
type PictureInPictureSampleBufferPlaybackDelegate struct {
	_PictureInPictureControllerShouldProhibitBackgroundAudioPlayback func(pictureInPictureController PictureInPictureController) bool
	_PictureInPictureControllerIsPlaybackPaused                      func(pictureInPictureController PictureInPictureController) bool
	_PictureInPictureControllerTimeRangeForPlayback                  func(pictureInPictureController PictureInPictureController) coremedia.TimeRange
	_PictureInPictureControllerSetPlaying                            func(pictureInPictureController PictureInPictureController, playing bool)
}

func (di *PictureInPictureSampleBufferPlaybackDelegate) HasPictureInPictureControllerShouldProhibitBackgroundAudioPlayback() bool {
	return di._PictureInPictureControllerShouldProhibitBackgroundAudioPlayback != nil
}

// Asks the delegate whether to always prohibit background audio playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturesamplebufferplaybackdelegate/3857563-pictureinpicturecontrollershould?language=objc
func (di *PictureInPictureSampleBufferPlaybackDelegate) SetPictureInPictureControllerShouldProhibitBackgroundAudioPlayback(f func(pictureInPictureController PictureInPictureController) bool) {
	di._PictureInPictureControllerShouldProhibitBackgroundAudioPlayback = f
}

// Asks the delegate whether to always prohibit background audio playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturesamplebufferplaybackdelegate/3857563-pictureinpicturecontrollershould?language=objc
func (di *PictureInPictureSampleBufferPlaybackDelegate) PictureInPictureControllerShouldProhibitBackgroundAudioPlayback(pictureInPictureController PictureInPictureController) bool {
	return di._PictureInPictureControllerShouldProhibitBackgroundAudioPlayback(pictureInPictureController)
}
func (di *PictureInPictureSampleBufferPlaybackDelegate) HasPictureInPictureControllerIsPlaybackPaused() bool {
	return di._PictureInPictureControllerIsPlaybackPaused != nil
}

// Asks delegate to indicate whether the playback UI reflects a playing or paused state, regardless of the current playback rate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturesamplebufferplaybackdelegate/3750336-pictureinpicturecontrollerisplay?language=objc
func (di *PictureInPictureSampleBufferPlaybackDelegate) SetPictureInPictureControllerIsPlaybackPaused(f func(pictureInPictureController PictureInPictureController) bool) {
	di._PictureInPictureControllerIsPlaybackPaused = f
}

// Asks delegate to indicate whether the playback UI reflects a playing or paused state, regardless of the current playback rate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturesamplebufferplaybackdelegate/3750336-pictureinpicturecontrollerisplay?language=objc
func (di *PictureInPictureSampleBufferPlaybackDelegate) PictureInPictureControllerIsPlaybackPaused(pictureInPictureController PictureInPictureController) bool {
	return di._PictureInPictureControllerIsPlaybackPaused(pictureInPictureController)
}
func (di *PictureInPictureSampleBufferPlaybackDelegate) HasPictureInPictureControllerTimeRangeForPlayback() bool {
	return di._PictureInPictureControllerTimeRangeForPlayback != nil
}

// Asks the delegate for the current playable time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturesamplebufferplaybackdelegate/3750337-pictureinpicturecontrollertimera?language=objc
func (di *PictureInPictureSampleBufferPlaybackDelegate) SetPictureInPictureControllerTimeRangeForPlayback(f func(pictureInPictureController PictureInPictureController) coremedia.TimeRange) {
	di._PictureInPictureControllerTimeRangeForPlayback = f
}

// Asks the delegate for the current playable time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturesamplebufferplaybackdelegate/3750337-pictureinpicturecontrollertimera?language=objc
func (di *PictureInPictureSampleBufferPlaybackDelegate) PictureInPictureControllerTimeRangeForPlayback(pictureInPictureController PictureInPictureController) coremedia.TimeRange {
	return di._PictureInPictureControllerTimeRangeForPlayback(pictureInPictureController)
}
func (di *PictureInPictureSampleBufferPlaybackDelegate) HasPictureInPictureControllerSetPlaying() bool {
	return di._PictureInPictureControllerSetPlaying != nil
}

// Tells the delegate that the user requested to begin or pause playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturesamplebufferplaybackdelegate/3750334-pictureinpicturecontroller?language=objc
func (di *PictureInPictureSampleBufferPlaybackDelegate) SetPictureInPictureControllerSetPlaying(f func(pictureInPictureController PictureInPictureController, playing bool)) {
	di._PictureInPictureControllerSetPlaying = f
}

// Tells the delegate that the user requested to begin or pause playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturesamplebufferplaybackdelegate/3750334-pictureinpicturecontroller?language=objc
func (di *PictureInPictureSampleBufferPlaybackDelegate) PictureInPictureControllerSetPlaying(pictureInPictureController PictureInPictureController, playing bool) {
	di._PictureInPictureControllerSetPlaying(pictureInPictureController, playing)
}

// A concrete type wrapper for the [PPictureInPictureSampleBufferPlaybackDelegate] protocol.
type PictureInPictureSampleBufferPlaybackDelegateWrapper struct {
	objc.Object
}

func (p_ PictureInPictureSampleBufferPlaybackDelegateWrapper) HasPictureInPictureControllerShouldProhibitBackgroundAudioPlayback() bool {
	return p_.RespondsToSelector(objc.Sel("pictureInPictureControllerShouldProhibitBackgroundAudioPlayback:"))
}

// Asks the delegate whether to always prohibit background audio playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturesamplebufferplaybackdelegate/3857563-pictureinpicturecontrollershould?language=objc
func (p_ PictureInPictureSampleBufferPlaybackDelegateWrapper) PictureInPictureControllerShouldProhibitBackgroundAudioPlayback(pictureInPictureController IPictureInPictureController) bool {
	rv := objc.Call[bool](p_, objc.Sel("pictureInPictureControllerShouldProhibitBackgroundAudioPlayback:"), objc.Ptr(pictureInPictureController))
	return rv
}

func (p_ PictureInPictureSampleBufferPlaybackDelegateWrapper) HasPictureInPictureControllerIsPlaybackPaused() bool {
	return p_.RespondsToSelector(objc.Sel("pictureInPictureControllerIsPlaybackPaused:"))
}

// Asks delegate to indicate whether the playback UI reflects a playing or paused state, regardless of the current playback rate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturesamplebufferplaybackdelegate/3750336-pictureinpicturecontrollerisplay?language=objc
func (p_ PictureInPictureSampleBufferPlaybackDelegateWrapper) PictureInPictureControllerIsPlaybackPaused(pictureInPictureController IPictureInPictureController) bool {
	rv := objc.Call[bool](p_, objc.Sel("pictureInPictureControllerIsPlaybackPaused:"), objc.Ptr(pictureInPictureController))
	return rv
}

func (p_ PictureInPictureSampleBufferPlaybackDelegateWrapper) HasPictureInPictureControllerTimeRangeForPlayback() bool {
	return p_.RespondsToSelector(objc.Sel("pictureInPictureControllerTimeRangeForPlayback:"))
}

// Asks the delegate for the current playable time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturesamplebufferplaybackdelegate/3750337-pictureinpicturecontrollertimera?language=objc
func (p_ PictureInPictureSampleBufferPlaybackDelegateWrapper) PictureInPictureControllerTimeRangeForPlayback(pictureInPictureController IPictureInPictureController) coremedia.TimeRange {
	rv := objc.Call[coremedia.TimeRange](p_, objc.Sel("pictureInPictureControllerTimeRangeForPlayback:"), objc.Ptr(pictureInPictureController))
	return rv
}

func (p_ PictureInPictureSampleBufferPlaybackDelegateWrapper) HasPictureInPictureControllerSetPlaying() bool {
	return p_.RespondsToSelector(objc.Sel("pictureInPictureController:setPlaying:"))
}

// Tells the delegate that the user requested to begin or pause playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturesamplebufferplaybackdelegate/3750334-pictureinpicturecontroller?language=objc
func (p_ PictureInPictureSampleBufferPlaybackDelegateWrapper) PictureInPictureControllerSetPlaying(pictureInPictureController IPictureInPictureController, playing bool) {
	objc.Call[objc.Void](p_, objc.Sel("pictureInPictureController:setPlaying:"), objc.Ptr(pictureInPictureController), playing)
}
