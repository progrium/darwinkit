// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ViewAnimationClass = _ViewAnimationClass{objc.GetClass("NSViewAnimation")}

type _ViewAnimationClass struct {
	objc.Class
}

type IViewAnimation interface {
	IAnimation
	ViewAnimations() []map[ViewAnimationKey]objc.Object
	SetViewAnimations(value []map[ViewAnimationKey]objc.IObject)
}

type ViewAnimation struct {
	Animation
}

func MakeViewAnimation(ptr unsafe.Pointer) ViewAnimation {
	return ViewAnimation{
		Animation: MakeAnimation(ptr),
	}
}

func (v_ ViewAnimation) InitWithViewAnimations(viewAnimations []map[ViewAnimationKey]objc.IObject) ViewAnimation {
	rv := objc.CallMethod[ViewAnimation](v_, objc.GetSelector("initWithViewAnimations:"), viewAnimations)
	return rv
}

func ViewAnimation_InitWithViewAnimations(viewAnimations []map[ViewAnimationKey]objc.IObject) ViewAnimation {
	return ViewAnimationClass.Alloc().InitWithViewAnimations(viewAnimations)
}

func (v_ ViewAnimation) InitWithDurationAnimationCurve(duration foundation.TimeInterval, animationCurve AnimationCurve) ViewAnimation {
	rv := objc.CallMethod[ViewAnimation](v_, objc.GetSelector("initWithDuration:animationCurve:"), duration, animationCurve)
	return rv
}

func ViewAnimation_InitWithDurationAnimationCurve(duration foundation.TimeInterval, animationCurve AnimationCurve) ViewAnimation {
	return ViewAnimationClass.Alloc().InitWithDurationAnimationCurve(duration, animationCurve)
}

func (vc _ViewAnimationClass) Alloc() ViewAnimation {
	rv := objc.CallMethod[ViewAnimation](vc, objc.GetSelector("alloc"))
	return rv
}

func ViewAnimation_Alloc() ViewAnimation {
	return ViewAnimationClass.Alloc()
}

func (vc _ViewAnimationClass) New() ViewAnimation {
	rv := objc.CallMethod[ViewAnimation](vc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewViewAnimation() ViewAnimation {
	return ViewAnimationClass.New()
}

func ViewAnimation_New() ViewAnimation {
	return ViewAnimationClass.New()
}

func (v_ ViewAnimation) Init() ViewAnimation {
	rv := objc.CallMethod[ViewAnimation](v_, objc.GetSelector("init"))
	return rv
}

func ViewAnimation_Init() ViewAnimation {
	return ViewAnimationClass.Alloc().Init()
}

func (v_ ViewAnimation) ViewAnimations() []map[ViewAnimationKey]objc.Object {
	rv := objc.CallMethod[[]map[ViewAnimationKey]objc.Object](v_, objc.GetSelector("viewAnimations"))
	// TODO: convert slice items...
	return rv
}

func (v_ ViewAnimation) SetViewAnimations(value []map[ViewAnimationKey]objc.IObject) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("setViewAnimations:"), value)
}
