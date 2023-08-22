// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AudioMixInputParameters] class.
var AudioMixInputParametersClass = _AudioMixInputParametersClass{objc.GetClass("AVAudioMixInputParameters")}

type _AudioMixInputParametersClass struct {
	objc.Class
}

// An interface definition for the [AudioMixInputParameters] class.
type IAudioMixInputParameters interface {
	objc.IObject
	GetVolumeRampForTimeStartVolumeEndVolumeTimeRange(time coremedia.Time, startVolume *float64, endVolume *float64, timeRange *coremedia.TimeRange) bool
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm
	TrackID() objc.Object
	AudioTapProcessor() objc.Object
}

// An object that represents the parameters that you apply when adding an audio track to a mix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaudiomixinputparameters?language=objc
type AudioMixInputParameters struct {
	objc.Object
}

func AudioMixInputParametersFrom(ptr unsafe.Pointer) AudioMixInputParameters {
	return AudioMixInputParameters{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AudioMixInputParametersClass) Alloc() AudioMixInputParameters {
	rv := objc.Call[AudioMixInputParameters](ac, objc.Sel("alloc"))
	return rv
}

func AudioMixInputParameters_Alloc() AudioMixInputParameters {
	return AudioMixInputParametersClass.Alloc()
}

func (ac _AudioMixInputParametersClass) New() AudioMixInputParameters {
	rv := objc.Call[AudioMixInputParameters](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAudioMixInputParameters() AudioMixInputParameters {
	return AudioMixInputParametersClass.New()
}

func (a_ AudioMixInputParameters) Init() AudioMixInputParameters {
	rv := objc.Call[AudioMixInputParameters](a_, objc.Sel("init"))
	return rv
}

// Retrieves the volume ramp that includes the specified time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaudiomixinputparameters/1389578-getvolumerampfortime?language=objc
func (a_ AudioMixInputParameters) GetVolumeRampForTimeStartVolumeEndVolumeTimeRange(time coremedia.Time, startVolume *float64, endVolume *float64, timeRange *coremedia.TimeRange) bool {
	rv := objc.Call[bool](a_, objc.Sel("getVolumeRampForTime:startVolume:endVolume:timeRange:"), time, startVolume, endVolume, timeRange)
	return rv
}

// The processing algorithm used to manage audio pitch for scaled audio edits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaudiomixinputparameters/1387042-audiotimepitchalgorithm?language=objc
func (a_ AudioMixInputParameters) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm {
	rv := objc.Call[AudioTimePitchAlgorithm](a_, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}

// The identifier of the audio track to which the parameters should be applied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaudiomixinputparameters/1387471-trackid?language=objc
func (a_ AudioMixInputParameters) TrackID() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("trackID"))
	return rv
}

// The audio processing tap associated with the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaudiomixinputparameters/1388578-audiotapprocessor?language=objc
func (a_ AudioMixInputParameters) AudioTapProcessor() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("audioTapProcessor"))
	return rv
}
