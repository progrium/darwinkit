// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureAudioChannel] class.
var CaptureAudioChannelClass = _CaptureAudioChannelClass{objc.GetClass("AVCaptureAudioChannel")}

type _CaptureAudioChannelClass struct {
	objc.Class
}

// An interface definition for the [CaptureAudioChannel] class.
type ICaptureAudioChannel interface {
	objc.IObject
	Volume() float64
	SetVolume(value float64)
	AveragePowerLevel() float64
	PeakHoldLevel() float64
	IsEnabled() bool
	SetEnabled(value bool)
}

// An object that monitors average and peak power levels for an audio channel in a capture connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiochannel?language=objc
type CaptureAudioChannel struct {
	objc.Object
}

func CaptureAudioChannelFrom(ptr unsafe.Pointer) CaptureAudioChannel {
	return CaptureAudioChannel{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptureAudioChannelClass) Alloc() CaptureAudioChannel {
	rv := objc.Call[CaptureAudioChannel](cc, objc.Sel("alloc"))
	return rv
}

func CaptureAudioChannel_Alloc() CaptureAudioChannel {
	return CaptureAudioChannelClass.Alloc()
}

func (cc _CaptureAudioChannelClass) New() CaptureAudioChannel {
	rv := objc.Call[CaptureAudioChannel](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureAudioChannel() CaptureAudioChannel {
	return CaptureAudioChannelClass.New()
}

func (c_ CaptureAudioChannel) Init() CaptureAudioChannel {
	rv := objc.Call[CaptureAudioChannel](c_, objc.Sel("init"))
	return rv
}

// The current volume (gain) of the channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiochannel/1388396-volume?language=objc
func (c_ CaptureAudioChannel) Volume() float64 {
	rv := objc.Call[float64](c_, objc.Sel("volume"))
	return rv
}

// The current volume (gain) of the channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiochannel/1388396-volume?language=objc
func (c_ CaptureAudioChannel) SetVolume(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setVolume:"), value)
}

// The instantaneous average power level in decibels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiochannel/1387368-averagepowerlevel?language=objc
func (c_ CaptureAudioChannel) AveragePowerLevel() float64 {
	rv := objc.Call[float64](c_, objc.Sel("averagePowerLevel"))
	return rv
}

// The peak hold power level in decibels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiochannel/1390872-peakholdlevel?language=objc
func (c_ CaptureAudioChannel) PeakHoldLevel() float64 {
	rv := objc.Call[float64](c_, objc.Sel("peakHoldLevel"))
	return rv
}

// A Boolean value that indicates whether the channel is in an enabled state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiochannel/1388574-enabled?language=objc
func (c_ CaptureAudioChannel) IsEnabled() bool {
	rv := objc.Call[bool](c_, objc.Sel("isEnabled"))
	return rv
}

// A Boolean value that indicates whether the channel is in an enabled state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiochannel/1388574-enabled?language=objc
func (c_ CaptureAudioChannel) SetEnabled(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setEnabled:"), value)
}
