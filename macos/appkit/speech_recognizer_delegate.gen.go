// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods implemented by delegates of NSSpeechRecognizer objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizerdelegate?language=objc
type PSpeechRecognizerDelegate interface {
	// optional
	SpeechRecognizerDidRecognizeCommand(sender SpeechRecognizer, command string)
	HasSpeechRecognizerDidRecognizeCommand() bool
}

// A delegate implementation builder for the [PSpeechRecognizerDelegate] protocol.
type SpeechRecognizerDelegate struct {
	_SpeechRecognizerDidRecognizeCommand func(sender SpeechRecognizer, command string)
}

func (di *SpeechRecognizerDelegate) HasSpeechRecognizerDidRecognizeCommand() bool {
	return di._SpeechRecognizerDidRecognizeCommand != nil
}

// Invoked when the recognition engine has recognized the application command command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizerdelegate/1534211-speechrecognizer?language=objc
func (di *SpeechRecognizerDelegate) SetSpeechRecognizerDidRecognizeCommand(f func(sender SpeechRecognizer, command string)) {
	di._SpeechRecognizerDidRecognizeCommand = f
}

// Invoked when the recognition engine has recognized the application command command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizerdelegate/1534211-speechrecognizer?language=objc
func (di *SpeechRecognizerDelegate) SpeechRecognizerDidRecognizeCommand(sender SpeechRecognizer, command string) {
	di._SpeechRecognizerDidRecognizeCommand(sender, command)
}

// A concrete type wrapper for the [PSpeechRecognizerDelegate] protocol.
type SpeechRecognizerDelegateWrapper struct {
	objc.Object
}

func (s_ SpeechRecognizerDelegateWrapper) HasSpeechRecognizerDidRecognizeCommand() bool {
	return s_.RespondsToSelector(objc.Sel("speechRecognizer:didRecognizeCommand:"))
}

// Invoked when the recognition engine has recognized the application command command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizerdelegate/1534211-speechrecognizer?language=objc
func (s_ SpeechRecognizerDelegateWrapper) SpeechRecognizerDidRecognizeCommand(sender ISpeechRecognizer, command string) {
	objc.Call[objc.Void](s_, objc.Sel("speechRecognizer:didRecognizeCommand:"), objc.Ptr(sender), command)
}
