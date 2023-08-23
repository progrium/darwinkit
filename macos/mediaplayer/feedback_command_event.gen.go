// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FeedbackCommandEvent] class.
var FeedbackCommandEventClass = _FeedbackCommandEventClass{objc.GetClass("MPFeedbackCommandEvent")}

type _FeedbackCommandEventClass struct {
	objc.Class
}

// An interface definition for the [FeedbackCommandEvent] class.
type IFeedbackCommandEvent interface {
	IRemoteCommandEvent
	IsNegative() bool
}

// An event requesting a change in the feedback setting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpfeedbackcommandevent?language=objc
type FeedbackCommandEvent struct {
	RemoteCommandEvent
}

func FeedbackCommandEventFrom(ptr unsafe.Pointer) FeedbackCommandEvent {
	return FeedbackCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}

func (fc _FeedbackCommandEventClass) Alloc() FeedbackCommandEvent {
	rv := objc.Call[FeedbackCommandEvent](fc, objc.Sel("alloc"))
	return rv
}

func FeedbackCommandEvent_Alloc() FeedbackCommandEvent {
	return FeedbackCommandEventClass.Alloc()
}

func (fc _FeedbackCommandEventClass) New() FeedbackCommandEvent {
	rv := objc.Call[FeedbackCommandEvent](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFeedbackCommandEvent() FeedbackCommandEvent {
	return FeedbackCommandEventClass.New()
}

func (f_ FeedbackCommandEvent) Init() FeedbackCommandEvent {
	rv := objc.Call[FeedbackCommandEvent](f_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether an app should perform a negative command appropriate to the target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpfeedbackcommandevent/1616773-negative?language=objc
func (f_ FeedbackCommandEvent) IsNegative() bool {
	rv := objc.Call[bool](f_, objc.Sel("isNegative"))
	return rv
}
