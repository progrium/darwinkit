// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Animation] class.
var AnimationClass = _AnimationClass{objc.GetClass("NSAnimation")}

type _AnimationClass struct {
	objc.Class
}

// An interface definition for the [Animation] class.
type IAnimation interface {
	objc.IObject
	RemoveProgressMark(progressMark AnimationProgress)
	StopAnimation()
	StartAnimation()
	ClearStartAnimation()
	StartWhenAnimationReachesProgress(animation IAnimation, startProgress AnimationProgress)
	AddProgressMark(progressMark AnimationProgress)
	ClearStopAnimation()
	StopWhenAnimationReachesProgress(animation IAnimation, stopProgress AnimationProgress)
	AnimationBlockingMode() AnimationBlockingMode
	SetAnimationBlockingMode(value AnimationBlockingMode)
	CurrentProgress() AnimationProgress
	SetCurrentProgress(value AnimationProgress)
	FrameRate() float64
	SetFrameRate(value float64)
	RunLoopModesForAnimating() []foundation.RunLoopMode
	ProgressMarks() []foundation.Number
	SetProgressMarks(value []foundation.INumber)
	Delegate() AnimationDelegateWrapper
	SetDelegate(value PAnimationDelegate)
	SetDelegateObject(valueObject objc.IObject)
	CurrentValue() float64
	IsAnimating() bool
	AnimationCurve() AnimationCurve
	SetAnimationCurve(value AnimationCurve)
	Duration() foundation.TimeInterval
	SetDuration(value foundation.TimeInterval)
}

// An object that manages the timing and progress of animations in the user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation?language=objc
type Animation struct {
	objc.Object
}

func AnimationFrom(ptr unsafe.Pointer) Animation {
	return Animation{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ Animation) InitWithDurationAnimationCurve(duration foundation.TimeInterval, animationCurve AnimationCurve) Animation {
	rv := objc.Call[Animation](a_, objc.Sel("initWithDuration:animationCurve:"), duration, animationCurve)
	return rv
}

// Returns an NSAnimation object initialized with the specified duration and animation-curve values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1530069-initwithduration?language=objc
func NewAnimationWithDurationAnimationCurve(duration foundation.TimeInterval, animationCurve AnimationCurve) Animation {
	instance := AnimationClass.Alloc().InitWithDurationAnimationCurve(duration, animationCurve)
	instance.Autorelease()
	return instance
}

func (ac _AnimationClass) Alloc() Animation {
	rv := objc.Call[Animation](ac, objc.Sel("alloc"))
	return rv
}

func Animation_Alloc() Animation {
	return AnimationClass.Alloc()
}

func (ac _AnimationClass) New() Animation {
	rv := objc.Call[Animation](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAnimation() Animation {
	return AnimationClass.New()
}

func (a_ Animation) Init() Animation {
	rv := objc.Call[Animation](a_, objc.Sel("init"))
	return rv
}

// Removes progress mark from the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1528288-removeprogressmark?language=objc
func (a_ Animation) RemoveProgressMark(progressMark AnimationProgress) {
	objc.Call[objc.Void](a_, objc.Sel("removeProgressMark:"), progressMark)
}

// Stops the animation represented by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1528672-stopanimation?language=objc
func (a_ Animation) StopAnimation() {
	objc.Call[objc.Void](a_, objc.Sel("stopAnimation"))
}

// Starts the animation represented by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1529209-startanimation?language=objc
func (a_ Animation) StartAnimation() {
	objc.Call[objc.Void](a_, objc.Sel("startAnimation"))
}

// Clears linkage to another animation that causes the receiver to start. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1535932-clearstartanimation?language=objc
func (a_ Animation) ClearStartAnimation() {
	objc.Call[objc.Void](a_, objc.Sel("clearStartAnimation"))
}

// Starts running the animation represented by the receiver when another animation reaches a specific progress mark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1526774-startwhenanimation?language=objc
func (a_ Animation) StartWhenAnimationReachesProgress(animation IAnimation, startProgress AnimationProgress) {
	objc.Call[objc.Void](a_, objc.Sel("startWhenAnimation:reachesProgress:"), objc.Ptr(animation), startProgress)
}

// Adds the progress mark to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1526360-addprogressmark?language=objc
func (a_ Animation) AddProgressMark(progressMark AnimationProgress) {
	objc.Call[objc.Void](a_, objc.Sel("addProgressMark:"), progressMark)
}

// Clears linkage to another animation that causes the receiver to stop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1524448-clearstopanimation?language=objc
func (a_ Animation) ClearStopAnimation() {
	objc.Call[objc.Void](a_, objc.Sel("clearStopAnimation"))
}

// Stops running the animation represented by the receiver when another animation reaches a specific progress mark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1530363-stopwhenanimation?language=objc
func (a_ Animation) StopWhenAnimationReachesProgress(animation IAnimation, stopProgress AnimationProgress) {
	objc.Call[objc.Void](a_, objc.Sel("stopWhenAnimation:reachesProgress:"), objc.Ptr(animation), stopProgress)
}

// The blocking mode of the animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1533725-animationblockingmode?language=objc
func (a_ Animation) AnimationBlockingMode() AnimationBlockingMode {
	rv := objc.Call[AnimationBlockingMode](a_, objc.Sel("animationBlockingMode"))
	return rv
}

// The blocking mode of the animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1533725-animationblockingmode?language=objc
func (a_ Animation) SetAnimationBlockingMode(value AnimationBlockingMode) {
	objc.Call[objc.Void](a_, objc.Sel("setAnimationBlockingMode:"), value)
}

// The current progress of the animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1530843-currentprogress?language=objc
func (a_ Animation) CurrentProgress() AnimationProgress {
	rv := objc.Call[AnimationProgress](a_, objc.Sel("currentProgress"))
	return rv
}

// The current progress of the animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1530843-currentprogress?language=objc
func (a_ Animation) SetCurrentProgress(value AnimationProgress) {
	objc.Call[objc.Void](a_, objc.Sel("setCurrentProgress:"), value)
}

// The number of frame updates per second to generate for the animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1526694-framerate?language=objc
func (a_ Animation) FrameRate() float64 {
	rv := objc.Call[float64](a_, objc.Sel("frameRate"))
	return rv
}

// The number of frame updates per second to generate for the animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1526694-framerate?language=objc
func (a_ Animation) SetFrameRate(value float64) {
	objc.Call[objc.Void](a_, objc.Sel("setFrameRate:"), value)
}

// An array of strings representing the run loop modes in which the animation can run. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1526965-runloopmodesforanimating?language=objc
func (a_ Animation) RunLoopModesForAnimating() []foundation.RunLoopMode {
	rv := objc.Call[[]foundation.RunLoopMode](a_, objc.Sel("runLoopModesForAnimating"))
	return rv
}

// An array of floating-point numbers representing current progress marks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1533642-progressmarks?language=objc
func (a_ Animation) ProgressMarks() []foundation.Number {
	rv := objc.Call[[]foundation.Number](a_, objc.Sel("progressMarks"))
	return rv
}

// An array of floating-point numbers representing current progress marks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1533642-progressmarks?language=objc
func (a_ Animation) SetProgressMarks(value []foundation.INumber) {
	objc.Call[objc.Void](a_, objc.Sel("setProgressMarks:"), value)
}

// The animation delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1524439-delegate?language=objc
func (a_ Animation) Delegate() AnimationDelegateWrapper {
	rv := objc.Call[AnimationDelegateWrapper](a_, objc.Sel("delegate"))
	return rv
}

// The animation delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1524439-delegate?language=objc
func (a_ Animation) SetDelegate(value PAnimationDelegate) {
	po0 := objc.WrapAsProtocol("NSAnimationDelegate", value)
	objc.SetAssociatedObject(a_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](a_, objc.Sel("setDelegate:"), po0)
}

// The animation delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1524439-delegate?language=objc
func (a_ Animation) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The current value of the animation effect, based on the current progress [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1531043-currentvalue?language=objc
func (a_ Animation) CurrentValue() float64 {
	rv := objc.Call[float64](a_, objc.Sel("currentValue"))
	return rv
}

// A Boolean value indicating whether the animation is in progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1527492-animating?language=objc
func (a_ Animation) IsAnimating() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAnimating"))
	return rv
}

// The timing curve for the animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1535321-animationcurve?language=objc
func (a_ Animation) AnimationCurve() AnimationCurve {
	rv := objc.Call[AnimationCurve](a_, objc.Sel("animationCurve"))
	return rv
}

// The timing curve for the animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1535321-animationcurve?language=objc
func (a_ Animation) SetAnimationCurve(value AnimationCurve) {
	objc.Call[objc.Void](a_, objc.Sel("setAnimationCurve:"), value)
}

// The duration of the animation, in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1535110-duration?language=objc
func (a_ Animation) Duration() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](a_, objc.Sel("duration"))
	return rv
}

// The duration of the animation, in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1535110-duration?language=objc
func (a_ Animation) SetDuration(value foundation.TimeInterval) {
	objc.Call[objc.Void](a_, objc.Sel("setDuration:"), value)
}
