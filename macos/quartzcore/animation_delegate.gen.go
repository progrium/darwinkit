// auto-generated code, do not modify
package quartzcore

import (
	"github.com/progrium/macdriver/objc"
)

type IAnimationDelegate interface {
	ImplementsAnimationDidStart() bool
	// optional
	AnimationDidStart(anim Animation)
	ImplementsAnimationDidStopFinished() bool
	// optional
	AnimationDidStopFinished(anim Animation, flag bool)
}

type AnimationDelegate struct {
	_AnimationDidStart        func(anim Animation)
	_AnimationDidStopFinished func(anim Animation, flag bool)
}

func (di *AnimationDelegate) ImplementsAnimationDidStart() bool {
	return di._AnimationDidStart != nil
}

func (di *AnimationDelegate) SetAnimationDidStart(f func(anim Animation)) {
	di._AnimationDidStart = f
}

func (di *AnimationDelegate) AnimationDidStart(anim Animation) {
	di._AnimationDidStart(anim)
}
func (di *AnimationDelegate) ImplementsAnimationDidStopFinished() bool {
	return di._AnimationDidStopFinished != nil
}

func (di *AnimationDelegate) SetAnimationDidStopFinished(f func(anim Animation, flag bool)) {
	di._AnimationDidStopFinished = f
}

func (di *AnimationDelegate) AnimationDidStopFinished(anim Animation, flag bool) {
	di._AnimationDidStopFinished(anim, flag)
}

type AnimationDelegateWrapper struct {
	objc.Object
}

func (a_ AnimationDelegateWrapper) ImplementsAnimationDidStart() bool {
	return a_.RespondsToSelector(objc.GetSelector("animationDidStart:"))
}

func (a_ AnimationDelegateWrapper) AnimationDidStart(anim IAnimation) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("animationDidStart:"), objc.ExtractPtr(anim))
}

func (a_ AnimationDelegateWrapper) ImplementsAnimationDidStopFinished() bool {
	return a_.RespondsToSelector(objc.GetSelector("animationDidStop:finished:"))
}

func (a_ AnimationDelegateWrapper) AnimationDidStopFinished(anim IAnimation, flag bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("animationDidStop:finished:"), objc.ExtractPtr(anim), flag)
}
