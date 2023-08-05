// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var AnimationClass = _AnimationClass{objc.GetClass("NSAnimation")}

type _AnimationClass struct {
	objc.Class
}

type IAnimation interface {
	objc.IObject
	StartAnimation()
	StopAnimation()
	AddProgressMark(progressMark AnimationProgress)
	RemoveProgressMark(progressMark AnimationProgress)
	StartWhenAnimationReachesProgress(animation IAnimation, startProgress AnimationProgress)
	StopWhenAnimationReachesProgress(animation IAnimation, stopProgress AnimationProgress)
	ClearStartAnimation()
	ClearStopAnimation()
	AnimationBlockingMode() AnimationBlockingMode
	SetAnimationBlockingMode(value AnimationBlockingMode)
	RunLoopModesForAnimating() []foundation.RunLoopMode
	AnimationCurve() AnimationCurve
	SetAnimationCurve(value AnimationCurve)
	Duration() foundation.TimeInterval
	SetDuration(value foundation.TimeInterval)
	FrameRate() float32
	SetFrameRate(value float32)
	Delegate() AnimationDelegateWrapper
	SetDelegate(value IAnimationDelegate)
	SetDelegate0(value objc.IObject)
	IsAnimating() bool
	CurrentProgress() AnimationProgress
	SetCurrentProgress(value AnimationProgress)
	CurrentValue() float32
	ProgressMarks() []foundation.Number
	SetProgressMarks(value []foundation.INumber)
}

type Animation struct {
	objc.Object
}

func MakeAnimation(ptr unsafe.Pointer) Animation {
	return Animation{
		Object: objc.MakeObject(ptr),
	}
}

func (a_ Animation) InitWithDurationAnimationCurve(duration foundation.TimeInterval, animationCurve AnimationCurve) Animation {
	rv := objc.CallMethod[Animation](a_, objc.GetSelector("initWithDuration:animationCurve:"), duration, animationCurve)
	return rv
}

func Animation_InitWithDurationAnimationCurve(duration foundation.TimeInterval, animationCurve AnimationCurve) Animation {
	return AnimationClass.Alloc().InitWithDurationAnimationCurve(duration, animationCurve)
}

func (ac _AnimationClass) Alloc() Animation {
	rv := objc.CallMethod[Animation](ac, objc.GetSelector("alloc"))
	return rv
}

func Animation_Alloc() Animation {
	return AnimationClass.Alloc()
}

func (ac _AnimationClass) New() Animation {
	rv := objc.CallMethod[Animation](ac, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewAnimation() Animation {
	return AnimationClass.New()
}

func Animation_New() Animation {
	return AnimationClass.New()
}

func (a_ Animation) Init() Animation {
	rv := objc.CallMethod[Animation](a_, objc.GetSelector("init"))
	return rv
}

func Animation_Init() Animation {
	return AnimationClass.Alloc().Init()
}

func (a_ Animation) StartAnimation() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("startAnimation"))
}

func (a_ Animation) StopAnimation() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("stopAnimation"))
}

func (a_ Animation) AddProgressMark(progressMark AnimationProgress) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("addProgressMark:"), progressMark)
}

func (a_ Animation) RemoveProgressMark(progressMark AnimationProgress) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("removeProgressMark:"), progressMark)
}

func (a_ Animation) StartWhenAnimationReachesProgress(animation IAnimation, startProgress AnimationProgress) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("startWhenAnimation:reachesProgress:"), objc.ExtractPtr(animation), startProgress)
}

func (a_ Animation) StopWhenAnimationReachesProgress(animation IAnimation, stopProgress AnimationProgress) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("stopWhenAnimation:reachesProgress:"), objc.ExtractPtr(animation), stopProgress)
}

func (a_ Animation) ClearStartAnimation() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("clearStartAnimation"))
}

func (a_ Animation) ClearStopAnimation() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("clearStopAnimation"))
}

func (a_ Animation) AnimationBlockingMode() AnimationBlockingMode {
	rv := objc.CallMethod[AnimationBlockingMode](a_, objc.GetSelector("animationBlockingMode"))
	return rv
}

func (a_ Animation) SetAnimationBlockingMode(value AnimationBlockingMode) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setAnimationBlockingMode:"), value)
}

func (a_ Animation) RunLoopModesForAnimating() []foundation.RunLoopMode {
	rv := objc.CallMethod[[]foundation.RunLoopMode](a_, objc.GetSelector("runLoopModesForAnimating"))
	return rv
}

func (a_ Animation) AnimationCurve() AnimationCurve {
	rv := objc.CallMethod[AnimationCurve](a_, objc.GetSelector("animationCurve"))
	return rv
}

func (a_ Animation) SetAnimationCurve(value AnimationCurve) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setAnimationCurve:"), value)
}

func (a_ Animation) Duration() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](a_, objc.GetSelector("duration"))
	return rv
}

func (a_ Animation) SetDuration(value foundation.TimeInterval) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setDuration:"), value)
}

func (a_ Animation) FrameRate() float32 {
	rv := objc.CallMethod[float32](a_, objc.GetSelector("frameRate"))
	return rv
}

func (a_ Animation) SetFrameRate(value float32) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setFrameRate:"), value)
}

func (a_ Animation) Delegate() AnimationDelegateWrapper {
	rv := objc.CallMethod[AnimationDelegateWrapper](a_, objc.GetSelector("delegate"))
	return rv
}

func (a_ Animation) SetDelegate(value IAnimationDelegate) {
	po := objc.WrapAsProtocol("NSAnimationDelegate", value)
	objc.SetAssociatedObject(a_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setDelegate:"), po)
}

func (a_ Animation) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (a_ Animation) IsAnimating() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("isAnimating"))
	return rv
}

func (a_ Animation) CurrentProgress() AnimationProgress {
	rv := objc.CallMethod[AnimationProgress](a_, objc.GetSelector("currentProgress"))
	return rv
}

func (a_ Animation) SetCurrentProgress(value AnimationProgress) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setCurrentProgress:"), value)
}

func (a_ Animation) CurrentValue() float32 {
	rv := objc.CallMethod[float32](a_, objc.GetSelector("currentValue"))
	return rv
}

func (a_ Animation) ProgressMarks() []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](a_, objc.GetSelector("progressMarks"))
	return rv
}

func (a_ Animation) SetProgressMarks(value []foundation.INumber) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setProgressMarks:"), value)
}
