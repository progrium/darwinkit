// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type IAnimationDelegate interface {
	ImplementsAnimationDidEnd() bool
	// optional
	AnimationDidEnd(animation Animation)
	ImplementsAnimationDidStop() bool
	// optional
	AnimationDidStop(animation Animation)
	ImplementsAnimationShouldStart() bool
	// optional
	AnimationShouldStart(animation Animation) bool
	ImplementsAnimationValueForProgress() bool
	// optional
	AnimationValueForProgress(animation Animation, progress AnimationProgress) float32
	ImplementsAnimationDidReachProgressMark() bool
	// optional
	AnimationDidReachProgressMark(animation Animation, progress AnimationProgress)
}

type AnimationDelegate struct {
	_AnimationDidEnd               func(animation Animation)
	_AnimationDidStop              func(animation Animation)
	_AnimationShouldStart          func(animation Animation) bool
	_AnimationValueForProgress     func(animation Animation, progress AnimationProgress) float32
	_AnimationDidReachProgressMark func(animation Animation, progress AnimationProgress)
}

func (di *AnimationDelegate) ImplementsAnimationDidEnd() bool {
	return di._AnimationDidEnd != nil
}

func (di *AnimationDelegate) SetAnimationDidEnd(f func(animation Animation)) {
	di._AnimationDidEnd = f
}

func (di *AnimationDelegate) AnimationDidEnd(animation Animation) {
	di._AnimationDidEnd(animation)
}
func (di *AnimationDelegate) ImplementsAnimationDidStop() bool {
	return di._AnimationDidStop != nil
}

func (di *AnimationDelegate) SetAnimationDidStop(f func(animation Animation)) {
	di._AnimationDidStop = f
}

func (di *AnimationDelegate) AnimationDidStop(animation Animation) {
	di._AnimationDidStop(animation)
}
func (di *AnimationDelegate) ImplementsAnimationShouldStart() bool {
	return di._AnimationShouldStart != nil
}

func (di *AnimationDelegate) SetAnimationShouldStart(f func(animation Animation) bool) {
	di._AnimationShouldStart = f
}

func (di *AnimationDelegate) AnimationShouldStart(animation Animation) bool {
	return di._AnimationShouldStart(animation)
}
func (di *AnimationDelegate) ImplementsAnimationValueForProgress() bool {
	return di._AnimationValueForProgress != nil
}

func (di *AnimationDelegate) SetAnimationValueForProgress(f func(animation Animation, progress AnimationProgress) float32) {
	di._AnimationValueForProgress = f
}

func (di *AnimationDelegate) AnimationValueForProgress(animation Animation, progress AnimationProgress) float32 {
	return di._AnimationValueForProgress(animation, progress)
}
func (di *AnimationDelegate) ImplementsAnimationDidReachProgressMark() bool {
	return di._AnimationDidReachProgressMark != nil
}

func (di *AnimationDelegate) SetAnimationDidReachProgressMark(f func(animation Animation, progress AnimationProgress)) {
	di._AnimationDidReachProgressMark = f
}

func (di *AnimationDelegate) AnimationDidReachProgressMark(animation Animation, progress AnimationProgress) {
	di._AnimationDidReachProgressMark(animation, progress)
}

type AnimationDelegateWrapper struct {
	objc.Object
}

func (a_ AnimationDelegateWrapper) ImplementsAnimationDidEnd() bool {
	return a_.RespondsToSelector(objc.GetSelector("animationDidEnd:"))
}

func (a_ AnimationDelegateWrapper) AnimationDidEnd(animation IAnimation) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("animationDidEnd:"), objc.ExtractPtr(animation))
}

func (a_ AnimationDelegateWrapper) ImplementsAnimationDidStop() bool {
	return a_.RespondsToSelector(objc.GetSelector("animationDidStop:"))
}

func (a_ AnimationDelegateWrapper) AnimationDidStop(animation IAnimation) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("animationDidStop:"), objc.ExtractPtr(animation))
}

func (a_ AnimationDelegateWrapper) ImplementsAnimationShouldStart() bool {
	return a_.RespondsToSelector(objc.GetSelector("animationShouldStart:"))
}

func (a_ AnimationDelegateWrapper) AnimationShouldStart(animation IAnimation) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("animationShouldStart:"), objc.ExtractPtr(animation))
	return rv
}

func (a_ AnimationDelegateWrapper) ImplementsAnimationValueForProgress() bool {
	return a_.RespondsToSelector(objc.GetSelector("animation:valueForProgress:"))
}

func (a_ AnimationDelegateWrapper) AnimationValueForProgress(animation IAnimation, progress AnimationProgress) float32 {
	rv := objc.CallMethod[float32](a_, objc.GetSelector("animation:valueForProgress:"), objc.ExtractPtr(animation), progress)
	return rv
}

func (a_ AnimationDelegateWrapper) ImplementsAnimationDidReachProgressMark() bool {
	return a_.RespondsToSelector(objc.GetSelector("animation:didReachProgressMark:"))
}

func (a_ AnimationDelegateWrapper) AnimationDidReachProgressMark(animation IAnimation, progress AnimationProgress) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("animation:didReachProgressMark:"), objc.ExtractPtr(animation), progress)
}
