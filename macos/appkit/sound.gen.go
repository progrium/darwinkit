// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var SoundClass = _SoundClass{objc.GetClass("NSSound")}

type _SoundClass struct {
	objc.Class
}

type ISound interface {
	objc.IObject
	SetName(string_ SoundName) bool
	Pause() bool
	Play() bool
	Resume() bool
	Stop() bool
	WriteToPasteboard(pasteboard IPasteboard)
	Delegate() SoundDelegateWrapper
	SetDelegate(value ISoundDelegate)
	SetDelegate0(value objc.IObject)
	Name() SoundName
	Volume() float32
	SetVolume(value float32)
	CurrentTime() foundation.TimeInterval
	SetCurrentTime(value foundation.TimeInterval)
	Loops() bool
	SetLoops(value bool)
	PlaybackDeviceIdentifier() SoundPlaybackDeviceIdentifier
	SetPlaybackDeviceIdentifier(value SoundPlaybackDeviceIdentifier)
	Duration() foundation.TimeInterval
	IsPlaying() bool
}

type Sound struct {
	objc.Object
}

func MakeSound(ptr unsafe.Pointer) Sound {
	return Sound{
		Object: objc.MakeObject(ptr),
	}
}

func (s_ Sound) InitWithContentsOfFileByReference(path string, byRef bool) Sound {
	rv := objc.CallMethod[Sound](s_, objc.GetSelector("initWithContentsOfFile:byReference:"), path, byRef)
	return rv
}

func Sound_InitWithContentsOfFileByReference(path string, byRef bool) Sound {
	return SoundClass.Alloc().InitWithContentsOfFileByReference(path, byRef)
}

func (s_ Sound) InitWithContentsOfURLByReference(url foundation.IURL, byRef bool) Sound {
	rv := objc.CallMethod[Sound](s_, objc.GetSelector("initWithContentsOfURL:byReference:"), objc.ExtractPtr(url), byRef)
	return rv
}

func Sound_InitWithContentsOfURLByReference(url foundation.IURL, byRef bool) Sound {
	return SoundClass.Alloc().InitWithContentsOfURLByReference(url, byRef)
}

func (s_ Sound) InitWithData(data []byte) Sound {
	rv := objc.CallMethod[Sound](s_, objc.GetSelector("initWithData:"), data)
	return rv
}

func Sound_InitWithData(data []byte) Sound {
	return SoundClass.Alloc().InitWithData(data)
}

func (s_ Sound) InitWithPasteboard(pasteboard IPasteboard) Sound {
	rv := objc.CallMethod[Sound](s_, objc.GetSelector("initWithPasteboard:"), objc.ExtractPtr(pasteboard))
	return rv
}

func Sound_InitWithPasteboard(pasteboard IPasteboard) Sound {
	return SoundClass.Alloc().InitWithPasteboard(pasteboard)
}

func (sc _SoundClass) Alloc() Sound {
	rv := objc.CallMethod[Sound](sc, objc.GetSelector("alloc"))
	return rv
}

func Sound_Alloc() Sound {
	return SoundClass.Alloc()
}

func (sc _SoundClass) New() Sound {
	rv := objc.CallMethod[Sound](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewSound() Sound {
	return SoundClass.New()
}

func Sound_New() Sound {
	return SoundClass.New()
}

func (s_ Sound) Init() Sound {
	rv := objc.CallMethod[Sound](s_, objc.GetSelector("init"))
	return rv
}

func Sound_Init() Sound {
	return SoundClass.Alloc().Init()
}

func (sc _SoundClass) CanInitWithPasteboard(pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](sc, objc.GetSelector("canInitWithPasteboard:"), objc.ExtractPtr(pasteboard))
	return rv
}

func Sound_CanInitWithPasteboard(pasteboard IPasteboard) bool {
	return SoundClass.CanInitWithPasteboard(pasteboard)
}

func (s_ Sound) SetName(string_ SoundName) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("setName:"), string_)
	return rv
}

func (sc _SoundClass) SoundNamed(name SoundName) Sound {
	rv := objc.CallMethod[Sound](sc, objc.GetSelector("soundNamed:"), name)
	return rv
}

func Sound_SoundNamed(name SoundName) Sound {
	return SoundClass.SoundNamed(name)
}

func (s_ Sound) Pause() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("pause"))
	return rv
}

func (s_ Sound) Play() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("play"))
	return rv
}

func (s_ Sound) Resume() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("resume"))
	return rv
}

func (s_ Sound) Stop() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("stop"))
	return rv
}

func (s_ Sound) WriteToPasteboard(pasteboard IPasteboard) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("writeToPasteboard:"), objc.ExtractPtr(pasteboard))
}

func (s_ Sound) Delegate() SoundDelegateWrapper {
	rv := objc.CallMethod[SoundDelegateWrapper](s_, objc.GetSelector("delegate"))
	return rv
}

func (s_ Sound) SetDelegate(value ISoundDelegate) {
	po := objc.WrapAsProtocol("NSSoundDelegate", value)
	objc.SetAssociatedObject(s_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDelegate:"), po)
}

func (s_ Sound) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (s_ Sound) Name() SoundName {
	rv := objc.CallMethod[SoundName](s_, objc.GetSelector("name"))
	return rv
}

func (s_ Sound) Volume() float32 {
	rv := objc.CallMethod[float32](s_, objc.GetSelector("volume"))
	return rv
}

func (s_ Sound) SetVolume(value float32) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setVolume:"), value)
}

func (s_ Sound) CurrentTime() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](s_, objc.GetSelector("currentTime"))
	return rv
}

func (s_ Sound) SetCurrentTime(value foundation.TimeInterval) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setCurrentTime:"), value)
}

func (s_ Sound) Loops() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("loops"))
	return rv
}

func (s_ Sound) SetLoops(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setLoops:"), value)
}

func (s_ Sound) PlaybackDeviceIdentifier() SoundPlaybackDeviceIdentifier {
	rv := objc.CallMethod[SoundPlaybackDeviceIdentifier](s_, objc.GetSelector("playbackDeviceIdentifier"))
	return rv
}

func (s_ Sound) SetPlaybackDeviceIdentifier(value SoundPlaybackDeviceIdentifier) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setPlaybackDeviceIdentifier:"), value)
}

func (sc _SoundClass) SoundUnfilteredTypes() []string {
	rv := objc.CallMethod[[]string](sc, objc.GetSelector("soundUnfilteredTypes"))
	return rv
}

func Sound_SoundUnfilteredTypes() []string {
	return SoundClass.SoundUnfilteredTypes()
}

func (s_ Sound) Duration() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](s_, objc.GetSelector("duration"))
	return rv
}

func (s_ Sound) IsPlaying() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("isPlaying"))
	return rv
}
