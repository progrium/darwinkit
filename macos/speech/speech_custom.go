package speech

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// SpeechRecognizer represents a speech recognizer that can process live or pre-recorded speech
type SpeechRecognizer struct {
	objc.Object
}

// SpeechRecognizerClass is the class object for SpeechRecognizer
var SpeechRecognizerClass = objc.GetClass("SFSpeechRecognizer")

// NewSpeechRecognizer creates a new SpeechRecognizer
func NewSpeechRecognizer() SpeechRecognizer {
	obj := objc.Call[objc.Object](SpeechRecognizerClass, objc.Sel("alloc"))
	return SpeechRecognizer{objc.Call[objc.Object](obj, objc.Sel("init"))}
}

// NewSpeechRecognizerWithLocale creates a new SpeechRecognizer with a specific locale
func NewSpeechRecognizerWithLocale(locale foundation.Locale) SpeechRecognizer {
	obj := objc.Call[objc.Object](SpeechRecognizerClass, objc.Sel("alloc"))
	return SpeechRecognizer{objc.Call[objc.Object](obj, objc.Sel("initWithLocale:"), locale)}
}

// IsAvailable returns whether the speech recognizer is available
func (r SpeechRecognizer) IsAvailable() bool {
	return objc.Call[bool](r, objc.Sel("isAvailable"))
}

// SupportedLocales returns the locales supported by the speech recognizer
func SupportedLocales() foundation.Set {
	return objc.Call[foundation.Set](SpeechRecognizerClass, objc.Sel("supportedLocales"))
}

// RequestAuthorization requests authorization to perform speech recognition
func RequestAuthorization(handler objc.Block) {
	objc.Call[objc.Void](SpeechRecognizerClass, objc.Sel("requestAuthorization:"), handler)
}

// AuthorizationStatus returns the current authorization status for speech recognition
func AuthorizationStatus() SFSpeechRecognizerAuthorizationStatus {
	return SFSpeechRecognizerAuthorizationStatus(objc.Call[int](SpeechRecognizerClass, objc.Sel("authorizationStatus")))
}

// SpeechRecognitionRequest represents a request to perform speech recognition
type SpeechRecognitionRequest struct {
	objc.Object
}

// SpeechRecognitionRequestClass is the class object for SpeechRecognitionRequest
var SpeechRecognitionRequestClass = objc.GetClass("SFSpeechRecognitionRequest")

// ShouldReportPartialResults returns whether partial results should be reported
func (r SpeechRecognitionRequest) ShouldReportPartialResults() bool {
	return objc.Call[bool](r, objc.Sel("shouldReportPartialResults"))
}

// SetShouldReportPartialResults sets whether partial results should be reported
func (r SpeechRecognitionRequest) SetShouldReportPartialResults(value bool) {
	objc.Call[objc.Void](r, objc.Sel("setShouldReportPartialResults:"), value)
}

// RequiresOnDeviceRecognition returns whether on-device recognition is required
func (r SpeechRecognitionRequest) RequiresOnDeviceRecognition() bool {
	return objc.Call[bool](r, objc.Sel("requiresOnDeviceRecognition"))
}

// SetRequiresOnDeviceRecognition sets whether on-device recognition is required
func (r SpeechRecognitionRequest) SetRequiresOnDeviceRecognition(value bool) {
	objc.Call[objc.Void](r, objc.Sel("setRequiresOnDeviceRecognition:"), value)
}

// SpeechRecognitionTask represents a task for performing speech recognition
type SpeechRecognitionTask struct {
	objc.Object
}

// SpeechRecognitionTaskClass is the class object for SpeechRecognitionTask
var SpeechRecognitionTaskClass = objc.GetClass("SFSpeechRecognitionTask")

// Cancel cancels the speech recognition task
func (t SpeechRecognitionTask) Cancel() {
	objc.Call[objc.Void](t, objc.Sel("cancel"))
}

// Finish finishes the speech recognition task
func (t SpeechRecognitionTask) Finish() {
	objc.Call[objc.Void](t, objc.Sel("finish"))
}

// IsCancelled returns whether the task is cancelled
func (t SpeechRecognitionTask) IsCancelled() bool {
	return objc.Call[bool](t, objc.Sel("isCancelled"))
}

// IsFinishing returns whether the task is finishing
func (t SpeechRecognitionTask) IsFinishing() bool {
	return objc.Call[bool](t, objc.Sel("isFinishing"))
}

// State returns the state of the task
func (t SpeechRecognitionTask) State() SFSpeechRecognitionTaskState {
	return SFSpeechRecognitionTaskState(objc.Call[int](t, objc.Sel("state")))
}

// SpeechURLRecognitionRequest represents a request to recognize speech from a URL
type SpeechURLRecognitionRequest struct {
	SpeechRecognitionRequest
}

// SpeechURLRecognitionRequestClass is the class object for SpeechURLRecognitionRequest
var SpeechURLRecognitionRequestClass = objc.GetClass("SFSpeechURLRecognitionRequest")

// NewSpeechURLRecognitionRequest creates a new SpeechURLRecognitionRequest
func NewSpeechURLRecognitionRequest(url foundation.URL) SpeechURLRecognitionRequest {
	obj := objc.Call[objc.Object](SpeechURLRecognitionRequestClass, objc.Sel("alloc"))
	return SpeechURLRecognitionRequest{SpeechRecognitionRequest{objc.Call[objc.Object](obj, objc.Sel("initWithURL:"), url)}}
}

// URL returns the URL of the request
func (r SpeechURLRecognitionRequest) URL() foundation.URL {
	return objc.Call[foundation.URL](r, objc.Sel("URL"))
}

// SpeechAudioBufferRecognitionRequest represents a request to recognize speech from an audio buffer
type SpeechAudioBufferRecognitionRequest struct {
	SpeechRecognitionRequest
}

// SpeechAudioBufferRecognitionRequestClass is the class object for SpeechAudioBufferRecognitionRequest
var SpeechAudioBufferRecognitionRequestClass = objc.GetClass("SFSpeechAudioBufferRecognitionRequest")

// NewSpeechAudioBufferRecognitionRequest creates a new SpeechAudioBufferRecognitionRequest
func NewSpeechAudioBufferRecognitionRequest() SpeechAudioBufferRecognitionRequest {
	obj := objc.Call[objc.Object](SpeechAudioBufferRecognitionRequestClass, objc.Sel("alloc"))
	return SpeechAudioBufferRecognitionRequest{SpeechRecognitionRequest{objc.Call[objc.Object](obj, objc.Sel("init"))}}
}

// AppendAudioSampleBuffer appends an audio sample buffer to the request
func (r SpeechAudioBufferRecognitionRequest) AppendAudioSampleBuffer(sampleBuffer objc.Object) {
	objc.Call[objc.Void](r, objc.Sel("appendAudioSampleBuffer:"), sampleBuffer)
}

// EndAudio ends the audio for the request
func (r SpeechAudioBufferRecognitionRequest) EndAudio() {
	objc.Call[objc.Void](r, objc.Sel("endAudio"))
}

// Voice represents a voice used for speech synthesis
type Voice struct {
	objc.Object
}

// VoiceClass is the class object for Voice
var VoiceClass = objc.GetClass("SFVoice")

// CurrentLanguageCode returns the current language code
func CurrentLanguageCode() foundation.String {
	return objc.Call[foundation.String](VoiceClass, objc.Sel("currentLanguageCode"))
}

// VoiceAnalyzer represents a voice analyzer
type VoiceAnalyzer struct {
	objc.Object
}

// VoiceAnalyzerClass is the class object for VoiceAnalyzer
var VoiceAnalyzerClass = objc.GetClass("SFVoiceAnalyzer")

// NewVoiceAnalyzer creates a new VoiceAnalyzer
func NewVoiceAnalyzer(audioFile objc.Object) VoiceAnalyzer {
	obj := objc.Call[objc.Object](VoiceAnalyzerClass, objc.Sel("alloc"))
	return VoiceAnalyzer{objc.Call[objc.Object](obj, objc.Sel("initWithAudioFile:"), audioFile)}
}

// AnalyzeAudioFileForAcousticFeatures analyzes an audio file for acoustic features
func (a VoiceAnalyzer) AnalyzeAudioFileForAcousticFeatures(audioFile objc.Object, features foundation.Array, completionHandler objc.Block) {
	objc.Call[objc.Void](a, objc.Sel("analyzeAudioFile:forAcousticFeatures:completionHandler:"), audioFile, features, completionHandler)
}

// SpeechRecognition represents a recognized speech result
type SpeechRecognition struct {
	objc.Object
}

// SpeechRecognitionClass is the class object for SpeechRecognition
var SpeechRecognitionClass = objc.GetClass("SFSpeechRecognition")

// BestTranscription returns the best transcription of the speech
func (r SpeechRecognition) BestTranscription() foundation.String {
	return objc.Call[foundation.String](r, objc.Sel("bestTranscription"))
}

// IsFinal returns whether the recognition result is final
func (r SpeechRecognition) IsFinal() bool {
	return objc.Call[bool](r, objc.Sel("isFinal"))
}

// Transcriptions returns all transcriptions of the speech
func (r SpeechRecognition) Transcriptions() foundation.Array {
	return objc.Call[foundation.Array](r, objc.Sel("transcriptions"))
}