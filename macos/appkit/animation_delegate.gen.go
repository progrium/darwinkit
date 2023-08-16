// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods implemented by delegates of NSAnimation objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate?language=objc
type PAnimationDelegate interface {
	// optional
	AnimationDidEnd(animation Animation)
	HasAnimationDidEnd() bool

	// optional
	AnimationDidReachProgressMark(animation Animation, progress AnimationProgress)
	HasAnimationDidReachProgressMark() bool

	// optional
	AnimationShouldStart(animation Animation) bool
	HasAnimationShouldStart() bool

	// optional
	AnimationDidStop(animation Animation)
	HasAnimationDidStop() bool
}

// A delegate implementation builder for the [PAnimationDelegate] protocol.
type AnimationDelegate struct {
	_AnimationDidEnd               func(animation Animation)
	_AnimationDidReachProgressMark func(animation Animation, progress AnimationProgress)
	_AnimationShouldStart          func(animation Animation) bool
	_AnimationDidStop              func(animation Animation)
}

func (di *AnimationDelegate) HasAnimationDidEnd() bool {
	return di._AnimationDidEnd != nil
}

// Sent to the delegate when the specified animation completes its run. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1535871-animationdidend?language=objc
func (di *AnimationDelegate) SetAnimationDidEnd(f func(animation Animation)) {
	di._AnimationDidEnd = f
}

// Sent to the delegate when the specified animation completes its run. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1535871-animationdidend?language=objc
func (di *AnimationDelegate) AnimationDidEnd(animation Animation) {
	di._AnimationDidEnd(animation)
}
func (di *AnimationDelegate) HasAnimationDidReachProgressMark() bool {
	return di._AnimationDidReachProgressMark != nil
}

// Sent to the delegate when an animation reaches a specific progress mark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1535100-animation?language=objc
func (di *AnimationDelegate) SetAnimationDidReachProgressMark(f func(animation Animation, progress AnimationProgress)) {
	di._AnimationDidReachProgressMark = f
}

// Sent to the delegate when an animation reaches a specific progress mark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1535100-animation?language=objc
func (di *AnimationDelegate) AnimationDidReachProgressMark(animation Animation, progress AnimationProgress) {
	di._AnimationDidReachProgressMark(animation, progress)
}
func (di *AnimationDelegate) HasAnimationShouldStart() bool {
	return di._AnimationShouldStart != nil
}

// Sent to the delegate just after an animation is started. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1533279-animationshouldstart?language=objc
func (di *AnimationDelegate) SetAnimationShouldStart(f func(animation Animation) bool) {
	di._AnimationShouldStart = f
}

// Sent to the delegate just after an animation is started. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1533279-animationshouldstart?language=objc
func (di *AnimationDelegate) AnimationShouldStart(animation Animation) bool {
	return di._AnimationShouldStart(animation)
}
func (di *AnimationDelegate) HasAnimationDidStop() bool {
	return di._AnimationDidStop != nil
}

// Sent to the delegate when the specified animation is stopped before it completes its run. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1534155-animationdidstop?language=objc
func (di *AnimationDelegate) SetAnimationDidStop(f func(animation Animation)) {
	di._AnimationDidStop = f
}

// Sent to the delegate when the specified animation is stopped before it completes its run. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1534155-animationdidstop?language=objc
func (di *AnimationDelegate) AnimationDidStop(animation Animation) {
	di._AnimationDidStop(animation)
}

// A concrete type wrapper for the [PAnimationDelegate] protocol.
type AnimationDelegateWrapper struct {
	objc.Object
}

func (a_ AnimationDelegateWrapper) HasAnimationDidEnd() bool {
	return a_.RespondsToSelector(objc.Sel("animationDidEnd:"))
}

// Sent to the delegate when the specified animation completes its run. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1535871-animationdidend?language=objc
func (a_ AnimationDelegateWrapper) AnimationDidEnd(animation IAnimation) {
	objc.Call[objc.Void](a_, objc.Sel("animationDidEnd:"), objc.Ptr(animation))
}

func (a_ AnimationDelegateWrapper) HasAnimationDidReachProgressMark() bool {
	return a_.RespondsToSelector(objc.Sel("animation:didReachProgressMark:"))
}

// Sent to the delegate when an animation reaches a specific progress mark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1535100-animation?language=objc
func (a_ AnimationDelegateWrapper) AnimationDidReachProgressMark(animation IAnimation, progress AnimationProgress) {
	objc.Call[objc.Void](a_, objc.Sel("animation:didReachProgressMark:"), objc.Ptr(animation), progress)
}

func (a_ AnimationDelegateWrapper) HasAnimationShouldStart() bool {
	return a_.RespondsToSelector(objc.Sel("animationShouldStart:"))
}

// Sent to the delegate just after an animation is started. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1533279-animationshouldstart?language=objc
func (a_ AnimationDelegateWrapper) AnimationShouldStart(animation IAnimation) bool {
	rv := objc.Call[bool](a_, objc.Sel("animationShouldStart:"), objc.Ptr(animation))
	return rv
}

func (a_ AnimationDelegateWrapper) HasAnimationDidStop() bool {
	return a_.RespondsToSelector(objc.Sel("animationDidStop:"))
}

// Sent to the delegate when the specified animation is stopped before it completes its run. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationdelegate/1534155-animationdidstop?language=objc
func (a_ AnimationDelegateWrapper) AnimationDidStop(animation IAnimation) {
	objc.Call[objc.Void](a_, objc.Sel("animationDidStop:"), objc.Ptr(animation))
}
