// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods for fine-tuning a gesture recognizer’s behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizerdelegate?language=objc
type PGestureRecognizerDelegate interface {
	// optional
	GestureRecognizerShouldBegin(gestureRecognizer GestureRecognizer) bool
	HasGestureRecognizerShouldBegin() bool

	// optional
	GestureRecognizerShouldReceiveTouch(gestureRecognizer GestureRecognizer, touch Touch) bool
	HasGestureRecognizerShouldReceiveTouch() bool
}

// A delegate implementation builder for the [PGestureRecognizerDelegate] protocol.
type GestureRecognizerDelegate struct {
	_GestureRecognizerShouldBegin        func(gestureRecognizer GestureRecognizer) bool
	_GestureRecognizerShouldReceiveTouch func(gestureRecognizer GestureRecognizer, touch Touch) bool
}

func (di *GestureRecognizerDelegate) HasGestureRecognizerShouldBegin() bool {
	return di._GestureRecognizerShouldBegin != nil
}

// Asks the delegate if a gesture recognizer should transition out of the Possible (NSGestureRecognizerStatePossible) state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizerdelegate/1535463-gesturerecognizershouldbegin?language=objc
func (di *GestureRecognizerDelegate) SetGestureRecognizerShouldBegin(f func(gestureRecognizer GestureRecognizer) bool) {
	di._GestureRecognizerShouldBegin = f
}

// Asks the delegate if a gesture recognizer should transition out of the Possible (NSGestureRecognizerStatePossible) state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizerdelegate/1535463-gesturerecognizershouldbegin?language=objc
func (di *GestureRecognizerDelegate) GestureRecognizerShouldBegin(gestureRecognizer GestureRecognizer) bool {
	return di._GestureRecognizerShouldBegin(gestureRecognizer)
}
func (di *GestureRecognizerDelegate) HasGestureRecognizerShouldReceiveTouch() bool {
	return di._GestureRecognizerShouldReceiveTouch != nil
}

// Called, for a new touch, before the system calls the touchesBegan:withEvent: method on the gesture recognizer. Return NO to prevent the gesture recognizer from seeing this touch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizerdelegate/2544787-gesturerecognizer?language=objc
func (di *GestureRecognizerDelegate) SetGestureRecognizerShouldReceiveTouch(f func(gestureRecognizer GestureRecognizer, touch Touch) bool) {
	di._GestureRecognizerShouldReceiveTouch = f
}

// Called, for a new touch, before the system calls the touchesBegan:withEvent: method on the gesture recognizer. Return NO to prevent the gesture recognizer from seeing this touch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizerdelegate/2544787-gesturerecognizer?language=objc
func (di *GestureRecognizerDelegate) GestureRecognizerShouldReceiveTouch(gestureRecognizer GestureRecognizer, touch Touch) bool {
	return di._GestureRecognizerShouldReceiveTouch(gestureRecognizer, touch)
}

// A concrete type wrapper for the [PGestureRecognizerDelegate] protocol.
type GestureRecognizerDelegateWrapper struct {
	objc.Object
}

func (g_ GestureRecognizerDelegateWrapper) HasGestureRecognizerShouldBegin() bool {
	return g_.RespondsToSelector(objc.Sel("gestureRecognizerShouldBegin:"))
}

// Asks the delegate if a gesture recognizer should transition out of the Possible (NSGestureRecognizerStatePossible) state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizerdelegate/1535463-gesturerecognizershouldbegin?language=objc
func (g_ GestureRecognizerDelegateWrapper) GestureRecognizerShouldBegin(gestureRecognizer IGestureRecognizer) bool {
	rv := objc.Call[bool](g_, objc.Sel("gestureRecognizerShouldBegin:"), objc.Ptr(gestureRecognizer))
	return rv
}

func (g_ GestureRecognizerDelegateWrapper) HasGestureRecognizerShouldReceiveTouch() bool {
	return g_.RespondsToSelector(objc.Sel("gestureRecognizer:shouldReceiveTouch:"))
}

// Called, for a new touch, before the system calls the touchesBegan:withEvent: method on the gesture recognizer. Return NO to prevent the gesture recognizer from seeing this touch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizerdelegate/2544787-gesturerecognizer?language=objc
func (g_ GestureRecognizerDelegateWrapper) GestureRecognizerShouldReceiveTouch(gestureRecognizer IGestureRecognizer, touch ITouch) bool {
	rv := objc.Call[bool](g_, objc.Sel("gestureRecognizer:shouldReceiveTouch:"), objc.Ptr(gestureRecognizer), objc.Ptr(touch))
	return rv
}
