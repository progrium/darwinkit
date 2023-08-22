// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// Methods for monitoring or controlling the output of a media file capture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputdelegate?language=objc
type PCaptureFileOutputDelegate interface {
	// optional
	CaptureOutputShouldProvideSampleAccurateRecordingStart(output CaptureOutput) bool
	HasCaptureOutputShouldProvideSampleAccurateRecordingStart() bool

	// optional
	CaptureOutputDidOutputSampleBufferFromConnection(output CaptureFileOutput, sampleBuffer coremedia.SampleBufferRef, connection CaptureConnection)
	HasCaptureOutputDidOutputSampleBufferFromConnection() bool
}

// A delegate implementation builder for the [PCaptureFileOutputDelegate] protocol.
type CaptureFileOutputDelegate struct {
	_CaptureOutputShouldProvideSampleAccurateRecordingStart func(output CaptureOutput) bool
	_CaptureOutputDidOutputSampleBufferFromConnection       func(output CaptureFileOutput, sampleBuffer coremedia.SampleBufferRef, connection CaptureConnection)
}

func (di *CaptureFileOutputDelegate) HasCaptureOutputShouldProvideSampleAccurateRecordingStart() bool {
	return di._CaptureOutputShouldProvideSampleAccurateRecordingStart != nil
}

// Allows a client to opt in to frame accurate recording in captureOutput:didOutputSampleBuffer:fromConnection:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputdelegate/1388760-captureoutputshouldprovidesample?language=objc
func (di *CaptureFileOutputDelegate) SetCaptureOutputShouldProvideSampleAccurateRecordingStart(f func(output CaptureOutput) bool) {
	di._CaptureOutputShouldProvideSampleAccurateRecordingStart = f
}

// Allows a client to opt in to frame accurate recording in captureOutput:didOutputSampleBuffer:fromConnection:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputdelegate/1388760-captureoutputshouldprovidesample?language=objc
func (di *CaptureFileOutputDelegate) CaptureOutputShouldProvideSampleAccurateRecordingStart(output CaptureOutput) bool {
	return di._CaptureOutputShouldProvideSampleAccurateRecordingStart(output)
}
func (di *CaptureFileOutputDelegate) HasCaptureOutputDidOutputSampleBufferFromConnection() bool {
	return di._CaptureOutputDidOutputSampleBufferFromConnection != nil
}

// Gives the delegate the opportunity to inspect samples as they are received by the output and start and stop recording at exact times. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputdelegate/1390096-captureoutput?language=objc
func (di *CaptureFileOutputDelegate) SetCaptureOutputDidOutputSampleBufferFromConnection(f func(output CaptureFileOutput, sampleBuffer coremedia.SampleBufferRef, connection CaptureConnection)) {
	di._CaptureOutputDidOutputSampleBufferFromConnection = f
}

// Gives the delegate the opportunity to inspect samples as they are received by the output and start and stop recording at exact times. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputdelegate/1390096-captureoutput?language=objc
func (di *CaptureFileOutputDelegate) CaptureOutputDidOutputSampleBufferFromConnection(output CaptureFileOutput, sampleBuffer coremedia.SampleBufferRef, connection CaptureConnection) {
	di._CaptureOutputDidOutputSampleBufferFromConnection(output, sampleBuffer, connection)
}

// A concrete type wrapper for the [PCaptureFileOutputDelegate] protocol.
type CaptureFileOutputDelegateWrapper struct {
	objc.Object
}

func (c_ CaptureFileOutputDelegateWrapper) HasCaptureOutputShouldProvideSampleAccurateRecordingStart() bool {
	return c_.RespondsToSelector(objc.Sel("captureOutputShouldProvideSampleAccurateRecordingStart:"))
}

// Allows a client to opt in to frame accurate recording in captureOutput:didOutputSampleBuffer:fromConnection:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputdelegate/1388760-captureoutputshouldprovidesample?language=objc
func (c_ CaptureFileOutputDelegateWrapper) CaptureOutputShouldProvideSampleAccurateRecordingStart(output ICaptureOutput) bool {
	rv := objc.Call[bool](c_, objc.Sel("captureOutputShouldProvideSampleAccurateRecordingStart:"), objc.Ptr(output))
	return rv
}

func (c_ CaptureFileOutputDelegateWrapper) HasCaptureOutputDidOutputSampleBufferFromConnection() bool {
	return c_.RespondsToSelector(objc.Sel("captureOutput:didOutputSampleBuffer:fromConnection:"))
}

// Gives the delegate the opportunity to inspect samples as they are received by the output and start and stop recording at exact times. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputdelegate/1390096-captureoutput?language=objc
func (c_ CaptureFileOutputDelegateWrapper) CaptureOutputDidOutputSampleBufferFromConnection(output ICaptureFileOutput, sampleBuffer coremedia.SampleBufferRef, connection ICaptureConnection) {
	objc.Call[objc.Void](c_, objc.Sel("captureOutput:didOutputSampleBuffer:fromConnection:"), objc.Ptr(output), sampleBuffer, objc.Ptr(connection))
}
