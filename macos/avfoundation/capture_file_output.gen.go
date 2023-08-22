// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureFileOutput] class.
var CaptureFileOutputClass = _CaptureFileOutputClass{objc.GetClass("AVCaptureFileOutput")}

type _CaptureFileOutputClass struct {
	objc.Class
}

// An interface definition for the [CaptureFileOutput] class.
type ICaptureFileOutput interface {
	ICaptureOutput
	PauseRecording()
	StartRecordingToOutputFileURLRecordingDelegate(outputFileURL foundation.IURL, delegate PCaptureFileOutputRecordingDelegate)
	StartRecordingToOutputFileURLRecordingDelegateObject(outputFileURL foundation.IURL, delegateObject objc.IObject)
	StopRecording()
	ResumeRecording()
	IsRecordingPaused() bool
	MaxRecordedFileSize() int64
	SetMaxRecordedFileSize(value int64)
	MaxRecordedDuration() coremedia.Time
	SetMaxRecordedDuration(value coremedia.Time)
	IsRecording() bool
	Delegate() CaptureFileOutputDelegateWrapper
	SetDelegate(value PCaptureFileOutputDelegate)
	SetDelegateObject(valueObject objc.IObject)
	MinFreeDiskSpaceLimit() int64
	SetMinFreeDiskSpaceLimit(value int64)
	RecordedDuration() coremedia.Time
	OutputFileURL() foundation.URL
	RecordedFileSize() int64
}

// The abstract superclass for capture outputs that can record captured data to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput?language=objc
type CaptureFileOutput struct {
	CaptureOutput
}

func CaptureFileOutputFrom(ptr unsafe.Pointer) CaptureFileOutput {
	return CaptureFileOutput{
		CaptureOutput: CaptureOutputFrom(ptr),
	}
}

func (cc _CaptureFileOutputClass) Alloc() CaptureFileOutput {
	rv := objc.Call[CaptureFileOutput](cc, objc.Sel("alloc"))
	return rv
}

func CaptureFileOutput_Alloc() CaptureFileOutput {
	return CaptureFileOutputClass.Alloc()
}

func (cc _CaptureFileOutputClass) New() CaptureFileOutput {
	rv := objc.Call[CaptureFileOutput](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureFileOutput() CaptureFileOutput {
	return CaptureFileOutputClass.New()
}

func (c_ CaptureFileOutput) Init() CaptureFileOutput {
	rv := objc.Call[CaptureFileOutput](c_, objc.Sel("init"))
	return rv
}

// Pauses recording to the current output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1386806-pauserecording?language=objc
func (c_ CaptureFileOutput) PauseRecording() {
	objc.Call[objc.Void](c_, objc.Sel("pauseRecording"))
}

// Starts recording media to the specified output URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1387224-startrecordingtooutputfileurl?language=objc
func (c_ CaptureFileOutput) StartRecordingToOutputFileURLRecordingDelegate(outputFileURL foundation.IURL, delegate PCaptureFileOutputRecordingDelegate) {
	po1 := objc.WrapAsProtocol("AVCaptureFileOutputRecordingDelegate", delegate)
	objc.Call[objc.Void](c_, objc.Sel("startRecordingToOutputFileURL:recordingDelegate:"), objc.Ptr(outputFileURL), po1)
}

// Starts recording media to the specified output URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1387224-startrecordingtooutputfileurl?language=objc
func (c_ CaptureFileOutput) StartRecordingToOutputFileURLRecordingDelegateObject(outputFileURL foundation.IURL, delegateObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("startRecordingToOutputFileURL:recordingDelegate:"), objc.Ptr(outputFileURL), objc.Ptr(delegateObject))
}

// Tells the receiver to stop recording to the current file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1389485-stoprecording?language=objc
func (c_ CaptureFileOutput) StopRecording() {
	objc.Call[objc.Void](c_, objc.Sel("stopRecording"))
}

// Resumes recording to the current output file after it was previously paused using pauseRecording. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1389849-resumerecording?language=objc
func (c_ CaptureFileOutput) ResumeRecording() {
	objc.Call[objc.Void](c_, objc.Sel("resumeRecording"))
}

// Indicates whether recording to the current output file is paused. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1385716-recordingpaused?language=objc
func (c_ CaptureFileOutput) IsRecordingPaused() bool {
	rv := objc.Call[bool](c_, objc.Sel("isRecordingPaused"))
	return rv
}

// The maximum size, in bytes, of the data that should be recorded by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1387684-maxrecordedfilesize?language=objc
func (c_ CaptureFileOutput) MaxRecordedFileSize() int64 {
	rv := objc.Call[int64](c_, objc.Sel("maxRecordedFileSize"))
	return rv
}

// The maximum size, in bytes, of the data that should be recorded by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1387684-maxrecordedfilesize?language=objc
func (c_ CaptureFileOutput) SetMaxRecordedFileSize(value int64) {
	objc.Call[objc.Void](c_, objc.Sel("setMaxRecordedFileSize:"), value)
}

// The longest duration allowed for the recording. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1387390-maxrecordedduration?language=objc
func (c_ CaptureFileOutput) MaxRecordedDuration() coremedia.Time {
	rv := objc.Call[coremedia.Time](c_, objc.Sel("maxRecordedDuration"))
	return rv
}

// The longest duration allowed for the recording. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1387390-maxrecordedduration?language=objc
func (c_ CaptureFileOutput) SetMaxRecordedDuration(value coremedia.Time) {
	objc.Call[objc.Void](c_, objc.Sel("setMaxRecordedDuration:"), value)
}

// Indicates whether recording is in progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1387539-recording?language=objc
func (c_ CaptureFileOutput) IsRecording() bool {
	rv := objc.Call[bool](c_, objc.Sel("isRecording"))
	return rv
}

// The delegate object for the capture file output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1388718-delegate?language=objc
func (c_ CaptureFileOutput) Delegate() CaptureFileOutputDelegateWrapper {
	rv := objc.Call[CaptureFileOutputDelegateWrapper](c_, objc.Sel("delegate"))
	return rv
}

// The delegate object for the capture file output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1388718-delegate?language=objc
func (c_ CaptureFileOutput) SetDelegate(value PCaptureFileOutputDelegate) {
	po0 := objc.WrapAsProtocol("AVCaptureFileOutputDelegate", value)
	objc.Call[objc.Void](c_, objc.Sel("setDelegate:"), po0)
}

// The delegate object for the capture file output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1388718-delegate?language=objc
func (c_ CaptureFileOutput) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The minimum amount of free space, in bytes, required for recording to continue on a given volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1387523-minfreediskspacelimit?language=objc
func (c_ CaptureFileOutput) MinFreeDiskSpaceLimit() int64 {
	rv := objc.Call[int64](c_, objc.Sel("minFreeDiskSpaceLimit"))
	return rv
}

// The minimum amount of free space, in bytes, required for recording to continue on a given volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1387523-minfreediskspacelimit?language=objc
func (c_ CaptureFileOutput) SetMinFreeDiskSpaceLimit(value int64) {
	objc.Call[objc.Void](c_, objc.Sel("setMinFreeDiskSpaceLimit:"), value)
}

// Indicates the duration of the media recorded to the current output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1389028-recordedduration?language=objc
func (c_ CaptureFileOutput) RecordedDuration() coremedia.Time {
	rv := objc.Call[coremedia.Time](c_, objc.Sel("recordedDuration"))
	return rv
}

// The URL to which output is directed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1388576-outputfileurl?language=objc
func (c_ CaptureFileOutput) OutputFileURL() foundation.URL {
	rv := objc.Call[foundation.URL](c_, objc.Sel("outputFileURL"))
	return rv
}

// Indicates the size, in bytes, of the data recorded to the current output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutput/1386933-recordedfilesize?language=objc
func (c_ CaptureFileOutput) RecordedFileSize() int64 {
	rv := objc.Call[int64](c_, objc.Sel("recordedFileSize"))
	return rv
}
