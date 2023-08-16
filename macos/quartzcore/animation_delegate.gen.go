// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"github.com/progrium/macdriver/objc"
)

// Methods your app can implement to respond when animations start and stop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimationdelegate?language=objc
type PAnimationDelegate interface {
	// optional
	AnimationDidStart(anim Animation)
	HasAnimationDidStart() bool

	// optional
	AnimationDidStopFinished(anim Animation, flag bool)
	HasAnimationDidStopFinished() bool
}

// A delegate implementation builder for the [PAnimationDelegate] protocol.
type AnimationDelegate struct {
	_AnimationDidStart        func(anim Animation)
	_AnimationDidStopFinished func(anim Animation, flag bool)
}

func (di *AnimationDelegate) HasAnimationDidStart() bool {
	return di._AnimationDidStart != nil
}

// Tells the delegate the animation has started. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimationdelegate/2097265-animationdidstart?language=objc
func (di *AnimationDelegate) SetAnimationDidStart(f func(anim Animation)) {
	di._AnimationDidStart = f
}

// Tells the delegate the animation has started. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimationdelegate/2097265-animationdidstart?language=objc
func (di *AnimationDelegate) AnimationDidStart(anim Animation) {
	di._AnimationDidStart(anim)
}
func (di *AnimationDelegate) HasAnimationDidStopFinished() bool {
	return di._AnimationDidStopFinished != nil
}

// Tells the delegate the animation has ended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimationdelegate/2097259-animationdidstop?language=objc
func (di *AnimationDelegate) SetAnimationDidStopFinished(f func(anim Animation, flag bool)) {
	di._AnimationDidStopFinished = f
}

// Tells the delegate the animation has ended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimationdelegate/2097259-animationdidstop?language=objc
func (di *AnimationDelegate) AnimationDidStopFinished(anim Animation, flag bool) {
	di._AnimationDidStopFinished(anim, flag)
}

// A concrete type wrapper for the [PAnimationDelegate] protocol.
type AnimationDelegateWrapper struct {
	objc.Object
}

func (a_ AnimationDelegateWrapper) HasAnimationDidStart() bool {
	return a_.RespondsToSelector(objc.Sel("animationDidStart:"))
}

// Tells the delegate the animation has started. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimationdelegate/2097265-animationdidstart?language=objc
func (a_ AnimationDelegateWrapper) AnimationDidStart(anim IAnimation) {
	objc.Call[objc.Void](a_, objc.Sel("animationDidStart:"), objc.Ptr(anim))
}

func (a_ AnimationDelegateWrapper) HasAnimationDidStopFinished() bool {
	return a_.RespondsToSelector(objc.Sel("animationDidStop:finished:"))
}

// Tells the delegate the animation has ended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimationdelegate/2097259-animationdidstop?language=objc
func (a_ AnimationDelegateWrapper) AnimationDidStopFinished(anim IAnimation, flag bool) {
	objc.Call[objc.Void](a_, objc.Sel("animationDidStop:finished:"), objc.Ptr(anim), flag)
}
