// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type IGestureRecognizerDelegate interface {
	ImplementsGestureRecognizerShouldAttemptToRecognizeWithEvent() bool
	// optional
	GestureRecognizerShouldAttemptToRecognizeWithEvent(gestureRecognizer GestureRecognizer, event Event) bool
	ImplementsGestureRecognizerShouldBegin() bool
	// optional
	GestureRecognizerShouldBegin(gestureRecognizer GestureRecognizer) bool
	ImplementsGestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer() bool
	// optional
	GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	ImplementsGestureRecognizerShouldRequireFailureOfGestureRecognizer() bool
	// optional
	GestureRecognizerShouldRequireFailureOfGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	ImplementsGestureRecognizerShouldBeRequiredToFailByGestureRecognizer() bool
	// optional
	GestureRecognizerShouldBeRequiredToFailByGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	ImplementsGestureRecognizerShouldReceiveTouch() bool
	// optional
	GestureRecognizerShouldReceiveTouch(gestureRecognizer GestureRecognizer, touch Touch) bool
}

type GestureRecognizerDelegate struct {
	_GestureRecognizerShouldAttemptToRecognizeWithEvent                  func(gestureRecognizer GestureRecognizer, event Event) bool
	_GestureRecognizerShouldBegin                                        func(gestureRecognizer GestureRecognizer) bool
	_GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	_GestureRecognizerShouldRequireFailureOfGestureRecognizer            func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	_GestureRecognizerShouldBeRequiredToFailByGestureRecognizer          func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	_GestureRecognizerShouldReceiveTouch                                 func(gestureRecognizer GestureRecognizer, touch Touch) bool
}

func (di *GestureRecognizerDelegate) ImplementsGestureRecognizerShouldAttemptToRecognizeWithEvent() bool {
	return di._GestureRecognizerShouldAttemptToRecognizeWithEvent != nil
}

func (di *GestureRecognizerDelegate) SetGestureRecognizerShouldAttemptToRecognizeWithEvent(f func(gestureRecognizer GestureRecognizer, event Event) bool) {
	di._GestureRecognizerShouldAttemptToRecognizeWithEvent = f
}

func (di *GestureRecognizerDelegate) GestureRecognizerShouldAttemptToRecognizeWithEvent(gestureRecognizer GestureRecognizer, event Event) bool {
	return di._GestureRecognizerShouldAttemptToRecognizeWithEvent(gestureRecognizer, event)
}
func (di *GestureRecognizerDelegate) ImplementsGestureRecognizerShouldBegin() bool {
	return di._GestureRecognizerShouldBegin != nil
}

func (di *GestureRecognizerDelegate) SetGestureRecognizerShouldBegin(f func(gestureRecognizer GestureRecognizer) bool) {
	di._GestureRecognizerShouldBegin = f
}

func (di *GestureRecognizerDelegate) GestureRecognizerShouldBegin(gestureRecognizer GestureRecognizer) bool {
	return di._GestureRecognizerShouldBegin(gestureRecognizer)
}
func (di *GestureRecognizerDelegate) ImplementsGestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer() bool {
	return di._GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer != nil
}

func (di *GestureRecognizerDelegate) SetGestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer(f func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool) {
	di._GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer = f
}

func (di *GestureRecognizerDelegate) GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool {
	return di._GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer(gestureRecognizer, otherGestureRecognizer)
}
func (di *GestureRecognizerDelegate) ImplementsGestureRecognizerShouldRequireFailureOfGestureRecognizer() bool {
	return di._GestureRecognizerShouldRequireFailureOfGestureRecognizer != nil
}

func (di *GestureRecognizerDelegate) SetGestureRecognizerShouldRequireFailureOfGestureRecognizer(f func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool) {
	di._GestureRecognizerShouldRequireFailureOfGestureRecognizer = f
}

func (di *GestureRecognizerDelegate) GestureRecognizerShouldRequireFailureOfGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool {
	return di._GestureRecognizerShouldRequireFailureOfGestureRecognizer(gestureRecognizer, otherGestureRecognizer)
}
func (di *GestureRecognizerDelegate) ImplementsGestureRecognizerShouldBeRequiredToFailByGestureRecognizer() bool {
	return di._GestureRecognizerShouldBeRequiredToFailByGestureRecognizer != nil
}

func (di *GestureRecognizerDelegate) SetGestureRecognizerShouldBeRequiredToFailByGestureRecognizer(f func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool) {
	di._GestureRecognizerShouldBeRequiredToFailByGestureRecognizer = f
}

func (di *GestureRecognizerDelegate) GestureRecognizerShouldBeRequiredToFailByGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool {
	return di._GestureRecognizerShouldBeRequiredToFailByGestureRecognizer(gestureRecognizer, otherGestureRecognizer)
}
func (di *GestureRecognizerDelegate) ImplementsGestureRecognizerShouldReceiveTouch() bool {
	return di._GestureRecognizerShouldReceiveTouch != nil
}

func (di *GestureRecognizerDelegate) SetGestureRecognizerShouldReceiveTouch(f func(gestureRecognizer GestureRecognizer, touch Touch) bool) {
	di._GestureRecognizerShouldReceiveTouch = f
}

func (di *GestureRecognizerDelegate) GestureRecognizerShouldReceiveTouch(gestureRecognizer GestureRecognizer, touch Touch) bool {
	return di._GestureRecognizerShouldReceiveTouch(gestureRecognizer, touch)
}

type GestureRecognizerDelegateWrapper struct {
	objc.Object
}

func (g_ GestureRecognizerDelegateWrapper) ImplementsGestureRecognizerShouldAttemptToRecognizeWithEvent() bool {
	return g_.RespondsToSelector(objc.GetSelector("gestureRecognizer:shouldAttemptToRecognizeWithEvent:"))
}

func (g_ GestureRecognizerDelegateWrapper) GestureRecognizerShouldAttemptToRecognizeWithEvent(gestureRecognizer IGestureRecognizer, event IEvent) bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("gestureRecognizer:shouldAttemptToRecognizeWithEvent:"), objc.ExtractPtr(gestureRecognizer), objc.ExtractPtr(event))
	return rv
}

func (g_ GestureRecognizerDelegateWrapper) ImplementsGestureRecognizerShouldBegin() bool {
	return g_.RespondsToSelector(objc.GetSelector("gestureRecognizerShouldBegin:"))
}

func (g_ GestureRecognizerDelegateWrapper) GestureRecognizerShouldBegin(gestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("gestureRecognizerShouldBegin:"), objc.ExtractPtr(gestureRecognizer))
	return rv
}

func (g_ GestureRecognizerDelegateWrapper) ImplementsGestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer() bool {
	return g_.RespondsToSelector(objc.GetSelector("gestureRecognizer:shouldRecognizeSimultaneouslyWithGestureRecognizer:"))
}

func (g_ GestureRecognizerDelegateWrapper) GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer(gestureRecognizer IGestureRecognizer, otherGestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("gestureRecognizer:shouldRecognizeSimultaneouslyWithGestureRecognizer:"), objc.ExtractPtr(gestureRecognizer), objc.ExtractPtr(otherGestureRecognizer))
	return rv
}

func (g_ GestureRecognizerDelegateWrapper) ImplementsGestureRecognizerShouldRequireFailureOfGestureRecognizer() bool {
	return g_.RespondsToSelector(objc.GetSelector("gestureRecognizer:shouldRequireFailureOfGestureRecognizer:"))
}

func (g_ GestureRecognizerDelegateWrapper) GestureRecognizerShouldRequireFailureOfGestureRecognizer(gestureRecognizer IGestureRecognizer, otherGestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("gestureRecognizer:shouldRequireFailureOfGestureRecognizer:"), objc.ExtractPtr(gestureRecognizer), objc.ExtractPtr(otherGestureRecognizer))
	return rv
}

func (g_ GestureRecognizerDelegateWrapper) ImplementsGestureRecognizerShouldBeRequiredToFailByGestureRecognizer() bool {
	return g_.RespondsToSelector(objc.GetSelector("gestureRecognizer:shouldBeRequiredToFailByGestureRecognizer:"))
}

func (g_ GestureRecognizerDelegateWrapper) GestureRecognizerShouldBeRequiredToFailByGestureRecognizer(gestureRecognizer IGestureRecognizer, otherGestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("gestureRecognizer:shouldBeRequiredToFailByGestureRecognizer:"), objc.ExtractPtr(gestureRecognizer), objc.ExtractPtr(otherGestureRecognizer))
	return rv
}

func (g_ GestureRecognizerDelegateWrapper) ImplementsGestureRecognizerShouldReceiveTouch() bool {
	return g_.RespondsToSelector(objc.GetSelector("gestureRecognizer:shouldReceiveTouch:"))
}

func (g_ GestureRecognizerDelegateWrapper) GestureRecognizerShouldReceiveTouch(gestureRecognizer IGestureRecognizer, touch ITouch) bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("gestureRecognizer:shouldReceiveTouch:"), objc.ExtractPtr(gestureRecognizer), objc.ExtractPtr(touch))
	return rv
}
