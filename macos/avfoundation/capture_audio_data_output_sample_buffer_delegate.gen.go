// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// Methods for receiving audio sample data from an audio capture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutputsamplebufferdelegate?language=objc
type PCaptureAudioDataOutputSampleBufferDelegate interface {
	// optional
	CaptureOutputDidOutputSampleBufferFromConnection(output CaptureOutput, sampleBuffer coremedia.SampleBufferRef, connection CaptureConnection)
	HasCaptureOutputDidOutputSampleBufferFromConnection() bool
}

// A delegate implementation builder for the [PCaptureAudioDataOutputSampleBufferDelegate] protocol.
type CaptureAudioDataOutputSampleBufferDelegate struct {
	_CaptureOutputDidOutputSampleBufferFromConnection func(output CaptureOutput, sampleBuffer coremedia.SampleBufferRef, connection CaptureConnection)
}

func (di *CaptureAudioDataOutputSampleBufferDelegate) HasCaptureOutputDidOutputSampleBufferFromConnection() bool {
	return di._CaptureOutputDidOutputSampleBufferFromConnection != nil
}

// Notifies the delegate that a sample buffer was written. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutputsamplebufferdelegate/1386039-captureoutput?language=objc
func (di *CaptureAudioDataOutputSampleBufferDelegate) SetCaptureOutputDidOutputSampleBufferFromConnection(f func(output CaptureOutput, sampleBuffer coremedia.SampleBufferRef, connection CaptureConnection)) {
	di._CaptureOutputDidOutputSampleBufferFromConnection = f
}

// Notifies the delegate that a sample buffer was written. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutputsamplebufferdelegate/1386039-captureoutput?language=objc
func (di *CaptureAudioDataOutputSampleBufferDelegate) CaptureOutputDidOutputSampleBufferFromConnection(output CaptureOutput, sampleBuffer coremedia.SampleBufferRef, connection CaptureConnection) {
	di._CaptureOutputDidOutputSampleBufferFromConnection(output, sampleBuffer, connection)
}

// A concrete type wrapper for the [PCaptureAudioDataOutputSampleBufferDelegate] protocol.
type CaptureAudioDataOutputSampleBufferDelegateWrapper struct {
	objc.Object
}

func (c_ CaptureAudioDataOutputSampleBufferDelegateWrapper) HasCaptureOutputDidOutputSampleBufferFromConnection() bool {
	return c_.RespondsToSelector(objc.Sel("captureOutput:didOutputSampleBuffer:fromConnection:"))
}

// Notifies the delegate that a sample buffer was written. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutputsamplebufferdelegate/1386039-captureoutput?language=objc
func (c_ CaptureAudioDataOutputSampleBufferDelegateWrapper) CaptureOutputDidOutputSampleBufferFromConnection(output ICaptureOutput, sampleBuffer coremedia.SampleBufferRef, connection ICaptureConnection) {
	objc.Call[objc.Void](c_, objc.Sel("captureOutput:didOutputSampleBuffer:fromConnection:"), objc.Ptr(output), sampleBuffer, objc.Ptr(connection))
}
