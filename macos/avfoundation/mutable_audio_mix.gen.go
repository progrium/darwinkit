// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableAudioMix] class.
var MutableAudioMixClass = _MutableAudioMixClass{objc.GetClass("AVMutableAudioMix")}

type _MutableAudioMixClass struct {
	objc.Class
}

// An interface definition for the [MutableAudioMix] class.
type IMutableAudioMix interface {
	IAudioMix
	SetInputParameters(value []IAudioMixInputParameters)
}

// An object that manages the input parameters for mixing audio tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableaudiomix?language=objc
type MutableAudioMix struct {
	AudioMix
}

func MutableAudioMixFrom(ptr unsafe.Pointer) MutableAudioMix {
	return MutableAudioMix{
		AudioMix: AudioMixFrom(ptr),
	}
}

func (mc _MutableAudioMixClass) AudioMix() MutableAudioMix {
	rv := objc.Call[MutableAudioMix](mc, objc.Sel("audioMix"))
	return rv
}

// Returns a new mutable audio mix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableaudiomix/1560973-audiomix?language=objc
func MutableAudioMix_AudioMix() MutableAudioMix {
	return MutableAudioMixClass.AudioMix()
}

func (mc _MutableAudioMixClass) Alloc() MutableAudioMix {
	rv := objc.Call[MutableAudioMix](mc, objc.Sel("alloc"))
	return rv
}

func MutableAudioMix_Alloc() MutableAudioMix {
	return MutableAudioMixClass.Alloc()
}

func (mc _MutableAudioMixClass) New() MutableAudioMix {
	rv := objc.Call[MutableAudioMix](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableAudioMix() MutableAudioMix {
	return MutableAudioMixClass.New()
}

func (m_ MutableAudioMix) Init() MutableAudioMix {
	rv := objc.Call[MutableAudioMix](m_, objc.Sel("init"))
	return rv
}

// An array of input parameters for the mix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableaudiomix/1388159-inputparameters?language=objc
func (m_ MutableAudioMix) SetInputParameters(value []IAudioMixInputParameters) {
	objc.Call[objc.Void](m_, objc.Sel("setInputParameters:"), value)
}
