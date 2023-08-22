// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureAudioDataOutput] class.
var CaptureAudioDataOutputClass = _CaptureAudioDataOutputClass{objc.GetClass("AVCaptureAudioDataOutput")}

type _CaptureAudioDataOutputClass struct {
	objc.Class
}

// An interface definition for the [CaptureAudioDataOutput] class.
type ICaptureAudioDataOutput interface {
	ICaptureOutput
	RecommendedAudioSettingsForAssetWriterWithOutputFileType(outputFileType FileType) map[string]objc.Object
	SetSampleBufferDelegateQueue(sampleBufferDelegate PCaptureAudioDataOutputSampleBufferDelegate, sampleBufferCallbackQueue dispatch.Queue)
	SetSampleBufferDelegateObjectQueue(sampleBufferDelegateObject objc.IObject, sampleBufferCallbackQueue dispatch.Queue)
	AudioSettings() map[string]objc.Object
	SetAudioSettings(value map[string]objc.IObject)
	SampleBufferCallbackQueue() dispatch.Queue
	SampleBufferDelegate() CaptureAudioDataOutputSampleBufferDelegateWrapper
}

// A capture output that records audio and provides access to audio sample buffers as they are recorded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutput?language=objc
type CaptureAudioDataOutput struct {
	CaptureOutput
}

func CaptureAudioDataOutputFrom(ptr unsafe.Pointer) CaptureAudioDataOutput {
	return CaptureAudioDataOutput{
		CaptureOutput: CaptureOutputFrom(ptr),
	}
}

func (cc _CaptureAudioDataOutputClass) New() CaptureAudioDataOutput {
	rv := objc.Call[CaptureAudioDataOutput](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureAudioDataOutput() CaptureAudioDataOutput {
	return CaptureAudioDataOutputClass.New()
}

func (c_ CaptureAudioDataOutput) Init() CaptureAudioDataOutput {
	rv := objc.Call[CaptureAudioDataOutput](c_, objc.Sel("init"))
	return rv
}

func (cc _CaptureAudioDataOutputClass) Alloc() CaptureAudioDataOutput {
	rv := objc.Call[CaptureAudioDataOutput](cc, objc.Sel("alloc"))
	return rv
}

func CaptureAudioDataOutput_Alloc() CaptureAudioDataOutput {
	return CaptureAudioDataOutputClass.Alloc()
}

// Specifies the recommended settings for use with an AVAssetWriterInput. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutput/1616308-recommendedaudiosettingsforasset?language=objc
func (c_ CaptureAudioDataOutput) RecommendedAudioSettingsForAssetWriterWithOutputFileType(outputFileType FileType) map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](c_, objc.Sel("recommendedAudioSettingsForAssetWriterWithOutputFileType:"), outputFileType)
	return rv
}

// Sets the delegate that will accept captured buffers and the dispatch queue on which the delegate will be called. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutput/1390651-setsamplebufferdelegate?language=objc
func (c_ CaptureAudioDataOutput) SetSampleBufferDelegateQueue(sampleBufferDelegate PCaptureAudioDataOutputSampleBufferDelegate, sampleBufferCallbackQueue dispatch.Queue) {
	po0 := objc.WrapAsProtocol("AVCaptureAudioDataOutputSampleBufferDelegate", sampleBufferDelegate)
	objc.Call[objc.Void](c_, objc.Sel("setSampleBufferDelegate:queue:"), po0, sampleBufferCallbackQueue)
}

// Sets the delegate that will accept captured buffers and the dispatch queue on which the delegate will be called. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutput/1390651-setsamplebufferdelegate?language=objc
func (c_ CaptureAudioDataOutput) SetSampleBufferDelegateObjectQueue(sampleBufferDelegateObject objc.IObject, sampleBufferCallbackQueue dispatch.Queue) {
	objc.Call[objc.Void](c_, objc.Sel("setSampleBufferDelegate:queue:"), objc.Ptr(sampleBufferDelegateObject), sampleBufferCallbackQueue)
}

// The settings used to decode or re-encode audio before it’s output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutput/1388527-audiosettings?language=objc
func (c_ CaptureAudioDataOutput) AudioSettings() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](c_, objc.Sel("audioSettings"))
	return rv
}

// The settings used to decode or re-encode audio before it’s output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutput/1388527-audiosettings?language=objc
func (c_ CaptureAudioDataOutput) SetAudioSettings(value map[string]objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setAudioSettings:"), value)
}

// The queue on which delegate callbacks are invoked [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutput/1389355-samplebuffercallbackqueue?language=objc
func (c_ CaptureAudioDataOutput) SampleBufferCallbackQueue() dispatch.Queue {
	rv := objc.Call[dispatch.Queue](c_, objc.Sel("sampleBufferCallbackQueue"))
	return rv
}

// The capture object’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutput/1386344-samplebufferdelegate?language=objc
func (c_ CaptureAudioDataOutput) SampleBufferDelegate() CaptureAudioDataOutputSampleBufferDelegateWrapper {
	rv := objc.Call[CaptureAudioDataOutputSampleBufferDelegateWrapper](c_, objc.Sel("sampleBufferDelegate"))
	return rv
}
