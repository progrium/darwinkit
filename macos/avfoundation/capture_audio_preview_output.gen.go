// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureAudioPreviewOutput] class.
var CaptureAudioPreviewOutputClass = _CaptureAudioPreviewOutputClass{objc.GetClass("AVCaptureAudioPreviewOutput")}

type _CaptureAudioPreviewOutputClass struct {
	objc.Class
}

// An interface definition for the [CaptureAudioPreviewOutput] class.
type ICaptureAudioPreviewOutput interface {
	ICaptureOutput
	Volume() float64
	SetVolume(value float64)
	OutputDeviceUniqueID() string
	SetOutputDeviceUniqueID(value string)
}

// A capture output that provides a preview of the captured audio. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiopreviewoutput?language=objc
type CaptureAudioPreviewOutput struct {
	CaptureOutput
}

func CaptureAudioPreviewOutputFrom(ptr unsafe.Pointer) CaptureAudioPreviewOutput {
	return CaptureAudioPreviewOutput{
		CaptureOutput: CaptureOutputFrom(ptr),
	}
}

func (cc _CaptureAudioPreviewOutputClass) New() CaptureAudioPreviewOutput {
	rv := objc.Call[CaptureAudioPreviewOutput](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureAudioPreviewOutput() CaptureAudioPreviewOutput {
	return CaptureAudioPreviewOutputClass.New()
}

func (c_ CaptureAudioPreviewOutput) Init() CaptureAudioPreviewOutput {
	rv := objc.Call[CaptureAudioPreviewOutput](c_, objc.Sel("init"))
	return rv
}

func (cc _CaptureAudioPreviewOutputClass) Alloc() CaptureAudioPreviewOutput {
	rv := objc.Call[CaptureAudioPreviewOutput](cc, objc.Sel("alloc"))
	return rv
}

func CaptureAudioPreviewOutput_Alloc() CaptureAudioPreviewOutput {
	return CaptureAudioPreviewOutputClass.Alloc()
}

// The output volume of the audio preview. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiopreviewoutput/1390510-volume?language=objc
func (c_ CaptureAudioPreviewOutput) Volume() float64 {
	rv := objc.Call[float64](c_, objc.Sel("volume"))
	return rv
}

// The output volume of the audio preview. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiopreviewoutput/1390510-volume?language=objc
func (c_ CaptureAudioPreviewOutput) SetVolume(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setVolume:"), value)
}

// The unique identifier of the Core Audio output device to use for audio preview. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiopreviewoutput/1390610-outputdeviceuniqueid?language=objc
func (c_ CaptureAudioPreviewOutput) OutputDeviceUniqueID() string {
	rv := objc.Call[string](c_, objc.Sel("outputDeviceUniqueID"))
	return rv
}

// The unique identifier of the Core Audio output device to use for audio preview. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiopreviewoutput/1390610-outputdeviceuniqueid?language=objc
func (c_ CaptureAudioPreviewOutput) SetOutputDeviceUniqueID(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setOutputDeviceUniqueID:"), value)
}
