// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AudioMix] class.
var AudioMixClass = _AudioMixClass{objc.GetClass("AVAudioMix")}

type _AudioMixClass struct {
	objc.Class
}

// An interface definition for the [AudioMix] class.
type IAudioMix interface {
	objc.IObject
	InputParameters() []AudioMixInputParameters
}

// An object that manages the input parameters for mixing audio tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaudiomix?language=objc
type AudioMix struct {
	objc.Object
}

func AudioMixFrom(ptr unsafe.Pointer) AudioMix {
	return AudioMix{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AudioMixClass) Alloc() AudioMix {
	rv := objc.Call[AudioMix](ac, objc.Sel("alloc"))
	return rv
}

func AudioMix_Alloc() AudioMix {
	return AudioMixClass.Alloc()
}

func (ac _AudioMixClass) New() AudioMix {
	rv := objc.Call[AudioMix](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAudioMix() AudioMix {
	return AudioMixClass.New()
}

func (a_ AudioMix) Init() AudioMix {
	rv := objc.Call[AudioMix](a_, objc.Sel("init"))
	return rv
}

// An array of input parameters for the mix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaudiomix/1388791-inputparameters?language=objc
func (a_ AudioMix) InputParameters() []AudioMixInputParameters {
	rv := objc.Call[[]AudioMixInputParameters](a_, objc.Sel("inputParameters"))
	return rv
}
