// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// Methods for responding to events that occur while recording captured media to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate?language=objc
type PCaptureFileOutputRecordingDelegate interface {
	// optional
	CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection)
	HasCaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections() bool
}

// A delegate implementation builder for the [PCaptureFileOutputRecordingDelegate] protocol.
type CaptureFileOutputRecordingDelegate struct {
	_CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections func(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection)
}

func (di *CaptureFileOutputRecordingDelegate) HasCaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections() bool {
	return di._CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections != nil
}

// Called whenever the output is recording to a file and successfully pauses the recording at the request of a client. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate/1388838-captureoutput?language=objc
func (di *CaptureFileOutputRecordingDelegate) SetCaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections(f func(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection)) {
	di._CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections = f
}

// Called whenever the output is recording to a file and successfully pauses the recording at the request of a client. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate/1388838-captureoutput?language=objc
func (di *CaptureFileOutputRecordingDelegate) CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection) {
	di._CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections(output, fileURL, connections)
}

// A concrete type wrapper for the [PCaptureFileOutputRecordingDelegate] protocol.
type CaptureFileOutputRecordingDelegateWrapper struct {
	objc.Object
}

func (c_ CaptureFileOutputRecordingDelegateWrapper) HasCaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections() bool {
	return c_.RespondsToSelector(objc.Sel("captureOutput:didPauseRecordingToOutputFileAtURL:fromConnections:"))
}

// Called whenever the output is recording to a file and successfully pauses the recording at the request of a client. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate/1388838-captureoutput?language=objc
func (c_ CaptureFileOutputRecordingDelegateWrapper) CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections(output ICaptureFileOutput, fileURL foundation.IURL, connections []ICaptureConnection) {
	objc.Call[objc.Void](c_, objc.Sel("captureOutput:didPauseRecordingToOutputFileAtURL:fromConnections:"), objc.Ptr(output), objc.Ptr(fileURL), connections)
}
