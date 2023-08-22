// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FeedbackCommand] class.
var FeedbackCommandClass = _FeedbackCommandClass{objc.GetClass("MPFeedbackCommand")}

type _FeedbackCommandClass struct {
	objc.Class
}

// An interface definition for the [FeedbackCommand] class.
type IFeedbackCommand interface {
	IRemoteCommand
	IsActive() bool
	SetActive(value bool)
	LocalizedTitle() string
	SetLocalizedTitle(value string)
	LocalizedShortTitle() string
	SetLocalizedShortTitle(value string)
}

// An object that reflects the feedback state for the playing item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpfeedbackcommand?language=objc
type FeedbackCommand struct {
	RemoteCommand
}

func FeedbackCommandFrom(ptr unsafe.Pointer) FeedbackCommand {
	return FeedbackCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}

func (fc _FeedbackCommandClass) Alloc() FeedbackCommand {
	rv := objc.Call[FeedbackCommand](fc, objc.Sel("alloc"))
	return rv
}

func FeedbackCommand_Alloc() FeedbackCommand {
	return FeedbackCommandClass.Alloc()
}

func (fc _FeedbackCommandClass) New() FeedbackCommand {
	rv := objc.Call[FeedbackCommand](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFeedbackCommand() FeedbackCommand {
	return FeedbackCommandClass.New()
}

func (f_ FeedbackCommand) Init() FeedbackCommand {
	rv := objc.Call[FeedbackCommand](f_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the feedback’s action is on or off. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpfeedbackcommand/1622900-active?language=objc
func (f_ FeedbackCommand) IsActive() bool {
	rv := objc.Call[bool](f_, objc.Sel("isActive"))
	return rv
}

// A Boolean value that indicates whether the feedback’s action is on or off. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpfeedbackcommand/1622900-active?language=objc
func (f_ FeedbackCommand) SetActive(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setActive:"), value)
}

// A localized string used to describe the context of a command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpfeedbackcommand/1622905-localizedtitle?language=objc
func (f_ FeedbackCommand) LocalizedTitle() string {
	rv := objc.Call[string](f_, objc.Sel("localizedTitle"))
	return rv
}

// A localized string used to describe the context of a command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpfeedbackcommand/1622905-localizedtitle?language=objc
func (f_ FeedbackCommand) SetLocalizedTitle(value string) {
	objc.Call[objc.Void](f_, objc.Sel("setLocalizedTitle:"), value)
}

// A shortened version of the string used to describe the context of a command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpfeedbackcommand/1622916-localizedshorttitle?language=objc
func (f_ FeedbackCommand) LocalizedShortTitle() string {
	rv := objc.Call[string](f_, objc.Sel("localizedShortTitle"))
	return rv
}

// A shortened version of the string used to describe the context of a command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpfeedbackcommand/1622916-localizedshorttitle?language=objc
func (f_ FeedbackCommand) SetLocalizedShortTitle(value string) {
	objc.Call[objc.Void](f_, objc.Sel("setLocalizedShortTitle:"), value)
}
