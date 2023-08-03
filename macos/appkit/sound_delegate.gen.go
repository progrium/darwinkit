// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type ISoundDelegate interface {
	ImplementsSoundDidFinishPlaying() bool
	// optional
	SoundDidFinishPlaying(sound Sound, flag bool)
}

type SoundDelegate struct {
	_SoundDidFinishPlaying func(sound Sound, flag bool)
}

func (di *SoundDelegate) ImplementsSoundDidFinishPlaying() bool {
	return di._SoundDidFinishPlaying != nil
}

func (di *SoundDelegate) SetSoundDidFinishPlaying(f func(sound Sound, flag bool)) {
	di._SoundDidFinishPlaying = f
}

func (di *SoundDelegate) SoundDidFinishPlaying(sound Sound, flag bool) {
	di._SoundDidFinishPlaying(sound, flag)
}

type SoundDelegateWrapper struct {
	objc.Object
}

func (s_ SoundDelegateWrapper) ImplementsSoundDidFinishPlaying() bool {
	return s_.RespondsToSelector(objc.GetSelector("sound:didFinishPlaying:"))
}

func (s_ SoundDelegateWrapper) SoundDidFinishPlaying(sound ISound, flag bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("sound:didFinishPlaying:"), objc.ExtractPtr(sound), flag)
}
