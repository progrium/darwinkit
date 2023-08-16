// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SpeechRecognizer] class.
var SpeechRecognizerClass = _SpeechRecognizerClass{objc.GetClass("NSSpeechRecognizer")}

type _SpeechRecognizerClass struct {
	objc.Class
}

// An interface definition for the [SpeechRecognizer] class.
type ISpeechRecognizer interface {
	objc.IObject
	StopListening()
	StartListening()
	ListensInForegroundOnly() bool
	SetListensInForegroundOnly(value bool)
	DisplayedCommandsTitle() string
	SetDisplayedCommandsTitle(value string)
	Delegate() SpeechRecognizerDelegateWrapper
	SetDelegate(value PSpeechRecognizerDelegate)
	SetDelegateObject(valueObject objc.IObject)
	Commands() []string
	SetCommands(value []string)
	BlocksOtherRecognizers() bool
	SetBlocksOtherRecognizers(value bool)
}

// The Cocoa interface to speech recognition in macOS. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer?language=objc
type SpeechRecognizer struct {
	objc.Object
}

func SpeechRecognizerFrom(ptr unsafe.Pointer) SpeechRecognizer {
	return SpeechRecognizer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ SpeechRecognizer) Init() SpeechRecognizer {
	rv := objc.Call[SpeechRecognizer](s_, objc.Sel("init"))
	return rv
}

func (sc _SpeechRecognizerClass) Alloc() SpeechRecognizer {
	rv := objc.Call[SpeechRecognizer](sc, objc.Sel("alloc"))
	return rv
}

func SpeechRecognizer_Alloc() SpeechRecognizer {
	return SpeechRecognizerClass.Alloc()
}

func (sc _SpeechRecognizerClass) New() SpeechRecognizer {
	rv := objc.Call[SpeechRecognizer](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSpeechRecognizer() SpeechRecognizer {
	return SpeechRecognizerClass.New()
}

// Tells the speech recognition engine to suspend listening for commands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/1529583-stoplistening?language=objc
func (s_ SpeechRecognizer) StopListening() {
	objc.Call[objc.Void](s_, objc.Sel("stopListening"))
}

// Tells the speech recognition engine to begin listening for commands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/1528154-startlistening?language=objc
func (s_ SpeechRecognizer) StartListening() {
	objc.Call[objc.Void](s_, objc.Sel("startListening"))
}

// A Boolean value that indicates whether the speech recognizer object should only enable its commands when its application is the frontmost one. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/1533405-listensinforegroundonly?language=objc
func (s_ SpeechRecognizer) ListensInForegroundOnly() bool {
	rv := objc.Call[bool](s_, objc.Sel("listensInForegroundOnly"))
	return rv
}

// A Boolean value that indicates whether the speech recognizer object should only enable its commands when its application is the frontmost one. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/1533405-listensinforegroundonly?language=objc
func (s_ SpeechRecognizer) SetListensInForegroundOnly(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setListensInForegroundOnly:"), value)
}

// The title of the commands section in the Speech Commands window or nil if there is no title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/1530622-displayedcommandstitle?language=objc
func (s_ SpeechRecognizer) DisplayedCommandsTitle() string {
	rv := objc.Call[string](s_, objc.Sel("displayedCommandsTitle"))
	return rv
}

// The title of the commands section in the Speech Commands window or nil if there is no title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/1530622-displayedcommandstitle?language=objc
func (s_ SpeechRecognizer) SetDisplayedCommandsTitle(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setDisplayedCommandsTitle:"), value)
}

// The delegate for the speech recognizer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/1528171-delegate?language=objc
func (s_ SpeechRecognizer) Delegate() SpeechRecognizerDelegateWrapper {
	rv := objc.Call[SpeechRecognizerDelegateWrapper](s_, objc.Sel("delegate"))
	return rv
}

// The delegate for the speech recognizer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/1528171-delegate?language=objc
func (s_ SpeechRecognizer) SetDelegate(value PSpeechRecognizerDelegate) {
	po0 := objc.WrapAsProtocol("NSSpeechRecognizerDelegate", value)
	objc.SetAssociatedObject(s_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), po0)
}

// The delegate for the speech recognizer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/1528171-delegate?language=objc
func (s_ SpeechRecognizer) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// An array of strings defining the commands for which the speech recognizer object should listen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/1529757-commands?language=objc
func (s_ SpeechRecognizer) Commands() []string {
	rv := objc.Call[[]string](s_, objc.Sel("commands"))
	return rv
}

// An array of strings defining the commands for which the speech recognizer object should listen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/1529757-commands?language=objc
func (s_ SpeechRecognizer) SetCommands(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setCommands:"), value)
}

// A Boolean value that indicates whether the speech recognizer object should block all other recognizers (that is, other applications attempting to understand spoken commands) when listening. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/1533022-blocksotherrecognizers?language=objc
func (s_ SpeechRecognizer) BlocksOtherRecognizers() bool {
	rv := objc.Call[bool](s_, objc.Sel("blocksOtherRecognizers"))
	return rv
}

// A Boolean value that indicates whether the speech recognizer object should block all other recognizers (that is, other applications attempting to understand spoken commands) when listening. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/1533022-blocksotherrecognizers?language=objc
func (s_ SpeechRecognizer) SetBlocksOtherRecognizers(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setBlocksOtherRecognizers:"), value)
}
