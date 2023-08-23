// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// Methods for receiving sample buffers from, and monitoring the status of, a video data output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutputsamplebufferdelegate?language=objc
type PCaptureVideoDataOutputSampleBufferDelegate interface {
	// optional
	CaptureOutputDidOutputSampleBufferFromConnection(output CaptureOutput, sampleBuffer coremedia.SampleBufferRef, connection CaptureConnection)
	HasCaptureOutputDidOutputSampleBufferFromConnection() bool
}

// A delegate implementation builder for the [PCaptureVideoDataOutputSampleBufferDelegate] protocol.
type CaptureVideoDataOutputSampleBufferDelegate struct {
	_CaptureOutputDidOutputSampleBufferFromConnection func(output CaptureOutput, sampleBuffer coremedia.SampleBufferRef, connection CaptureConnection)
}

func (di *CaptureVideoDataOutputSampleBufferDelegate) HasCaptureOutputDidOutputSampleBufferFromConnection() bool {
	return di._CaptureOutputDidOutputSampleBufferFromConnection != nil
}

// Notifies the delegate that a new video frame was written. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutputsamplebufferdelegate/1385775-captureoutput?language=objc
func (di *CaptureVideoDataOutputSampleBufferDelegate) SetCaptureOutputDidOutputSampleBufferFromConnection(f func(output CaptureOutput, sampleBuffer coremedia.SampleBufferRef, connection CaptureConnection)) {
	di._CaptureOutputDidOutputSampleBufferFromConnection = f
}

// Notifies the delegate that a new video frame was written. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutputsamplebufferdelegate/1385775-captureoutput?language=objc
func (di *CaptureVideoDataOutputSampleBufferDelegate) CaptureOutputDidOutputSampleBufferFromConnection(output CaptureOutput, sampleBuffer coremedia.SampleBufferRef, connection CaptureConnection) {
	di._CaptureOutputDidOutputSampleBufferFromConnection(output, sampleBuffer, connection)
}

// A concrete type wrapper for the [PCaptureVideoDataOutputSampleBufferDelegate] protocol.
type CaptureVideoDataOutputSampleBufferDelegateWrapper struct {
	objc.Object
}

func (c_ CaptureVideoDataOutputSampleBufferDelegateWrapper) HasCaptureOutputDidOutputSampleBufferFromConnection() bool {
	return c_.RespondsToSelector(objc.Sel("captureOutput:didOutputSampleBuffer:fromConnection:"))
}

// Notifies the delegate that a new video frame was written. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutputsamplebufferdelegate/1385775-captureoutput?language=objc
func (c_ CaptureVideoDataOutputSampleBufferDelegateWrapper) CaptureOutputDidOutputSampleBufferFromConnection(output ICaptureOutput, sampleBuffer coremedia.SampleBufferRef, connection ICaptureConnection) {
	objc.Call[objc.Void](c_, objc.Sel("captureOutput:didOutputSampleBuffer:fromConnection:"), objc.Ptr(output), sampleBuffer, objc.Ptr(connection))
}
