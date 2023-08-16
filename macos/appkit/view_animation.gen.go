// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ViewAnimation] class.
var ViewAnimationClass = _ViewAnimationClass{objc.GetClass("NSViewAnimation")}

type _ViewAnimationClass struct {
	objc.Class
}

// An interface definition for the [ViewAnimation] class.
type IViewAnimation interface {
	IAnimation
	ViewAnimations() []map[ViewAnimationKey]objc.Object
	SetViewAnimations(value []map[ViewAnimationKey]objc.IObject)
}

// An animation of an app's views, limited to changes in frame location and size, and to fade-in and fade-out effects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewanimation?language=objc
type ViewAnimation struct {
	Animation
}

func ViewAnimationFrom(ptr unsafe.Pointer) ViewAnimation {
	return ViewAnimation{
		Animation: AnimationFrom(ptr),
	}
}

func (v_ ViewAnimation) InitWithViewAnimations(viewAnimations []map[ViewAnimationKey]objc.IObject) ViewAnimation {
	rv := objc.Call[ViewAnimation](v_, objc.Sel("initWithViewAnimations:"), viewAnimations)
	return rv
}

// Returns an NSViewAnimation object initialized with the supplied information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewanimation/1531141-initwithviewanimations?language=objc
func ViewAnimation_InitWithViewAnimations(viewAnimations []map[ViewAnimationKey]objc.IObject) ViewAnimation {
	return ViewAnimationClass.Alloc().InitWithViewAnimations(viewAnimations)
}

func (vc _ViewAnimationClass) Alloc() ViewAnimation {
	rv := objc.Call[ViewAnimation](vc, objc.Sel("alloc"))
	return rv
}

func ViewAnimation_Alloc() ViewAnimation {
	return ViewAnimationClass.Alloc()
}

func (vc _ViewAnimationClass) New() ViewAnimation {
	rv := objc.Call[ViewAnimation](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewViewAnimation() ViewAnimation {
	return ViewAnimationClass.New()
}

func (v_ ViewAnimation) Init() ViewAnimation {
	rv := objc.Call[ViewAnimation](v_, objc.Sel("init"))
	return rv
}

func (v_ ViewAnimation) InitWithDurationAnimationCurve(duration foundation.TimeInterval, animationCurve AnimationCurve) ViewAnimation {
	rv := objc.Call[ViewAnimation](v_, objc.Sel("initWithDuration:animationCurve:"), duration, animationCurve)
	return rv
}

// Returns an NSAnimation object initialized with the specified duration and animation-curve values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/1530069-initwithduration?language=objc
func ViewAnimation_InitWithDurationAnimationCurve(duration foundation.TimeInterval, animationCurve AnimationCurve) ViewAnimation {
	return ViewAnimationClass.Alloc().InitWithDurationAnimationCurve(duration, animationCurve)
}

// The dictionaries defining the objects to animate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewanimation/1527416-viewanimations?language=objc
func (v_ ViewAnimation) ViewAnimations() []map[ViewAnimationKey]objc.Object {
	rv := objc.Call[[]map[ViewAnimationKey]objc.Object](v_, objc.Sel("viewAnimations"))
	return rv
}

// The dictionaries defining the objects to animate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewanimation/1527416-viewanimations?language=objc
func (v_ ViewAnimation) SetViewAnimations(value []map[ViewAnimationKey]objc.IObject) {
	objc.Call[objc.Void](v_, objc.Sel("setViewAnimations:"), value)
}
