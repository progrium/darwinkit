// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods implemented by delegates of NSSound objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssounddelegate?language=objc
type PSoundDelegate interface {
	// optional
	SoundDidFinishPlaying(sound Sound, flag bool)
	HasSoundDidFinishPlaying() bool
}

// A delegate implementation builder for the [PSoundDelegate] protocol.
type SoundDelegate struct {
	_SoundDidFinishPlaying func(sound Sound, flag bool)
}

func (di *SoundDelegate) HasSoundDidFinishPlaying() bool {
	return di._SoundDidFinishPlaying != nil
}

// This delegate method is called when an NSSound instance has completed playback of its sound data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssounddelegate/1477298-sound?language=objc
func (di *SoundDelegate) SetSoundDidFinishPlaying(f func(sound Sound, flag bool)) {
	di._SoundDidFinishPlaying = f
}

// This delegate method is called when an NSSound instance has completed playback of its sound data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssounddelegate/1477298-sound?language=objc
func (di *SoundDelegate) SoundDidFinishPlaying(sound Sound, flag bool) {
	di._SoundDidFinishPlaying(sound, flag)
}

// A concrete type wrapper for the [PSoundDelegate] protocol.
type SoundDelegateWrapper struct {
	objc.Object
}

func (s_ SoundDelegateWrapper) HasSoundDidFinishPlaying() bool {
	return s_.RespondsToSelector(objc.Sel("sound:didFinishPlaying:"))
}

// This delegate method is called when an NSSound instance has completed playback of its sound data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssounddelegate/1477298-sound?language=objc
func (s_ SoundDelegateWrapper) SoundDidFinishPlaying(sound ISound, flag bool) {
	objc.Call[objc.Void](s_, objc.Sel("sound:didFinishPlaying:"), objc.Ptr(sound), flag)
}
