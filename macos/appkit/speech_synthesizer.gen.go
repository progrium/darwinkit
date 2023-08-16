// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SpeechSynthesizer] class.
var SpeechSynthesizerClass = _SpeechSynthesizerClass{objc.GetClass("NSSpeechSynthesizer")}

type _SpeechSynthesizerClass struct {
	objc.Class
}

// An interface definition for the [SpeechSynthesizer] class.
type ISpeechSynthesizer interface {
	objc.IObject
}

// The Cocoa interface to speech synthesis in macOS. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechsynthesizer?language=objc
type SpeechSynthesizer struct {
	objc.Object
}

func SpeechSynthesizerFrom(ptr unsafe.Pointer) SpeechSynthesizer {
	return SpeechSynthesizer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SpeechSynthesizerClass) Alloc() SpeechSynthesizer {
	rv := objc.Call[SpeechSynthesizer](sc, objc.Sel("alloc"))
	return rv
}

func SpeechSynthesizer_Alloc() SpeechSynthesizer {
	return SpeechSynthesizerClass.Alloc()
}

func (sc _SpeechSynthesizerClass) New() SpeechSynthesizer {
	rv := objc.Call[SpeechSynthesizer](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSpeechSynthesizer() SpeechSynthesizer {
	return SpeechSynthesizerClass.New()
}

func (s_ SpeechSynthesizer) Init() SpeechSynthesizer {
	rv := objc.Call[SpeechSynthesizer](s_, objc.Sel("init"))
	return rv
}
