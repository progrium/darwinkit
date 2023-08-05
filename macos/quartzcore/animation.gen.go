// AUTO-GENERATED CODE, DO NOT MODIFY
package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var AnimationClass = _AnimationClass{objc.GetClass("CAAnimation")}

type _AnimationClass struct {
	objc.Class
}

type IAnimation interface {
	objc.IObject
	ShouldArchiveValueForKey(key string) bool
	IsRemovedOnCompletion() bool
	SetRemovedOnCompletion(value bool)
	TimingFunction() MediaTimingFunction
	SetTimingFunction(value IMediaTimingFunction)
	Delegate() AnimationDelegateWrapper
	SetDelegate(value IAnimationDelegate)
	SetDelegate0(value objc.IObject)
	PreferredFrameRateRange() FrameRateRange
	SetPreferredFrameRateRange(value FrameRateRange)
}

type Animation struct {
	objc.Object
}

func MakeAnimation(ptr unsafe.Pointer) Animation {
	return Animation{
		Object: objc.MakeObject(ptr),
	}
}

func (ac _AnimationClass) Animation() Animation {
	rv := objc.CallMethod[Animation](ac, objc.GetSelector("animation"))
	return rv
}

func Animation_Animation() Animation {
	return AnimationClass.Animation()
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

func (ac _AnimationClass) DefaultValueForKey(key string) objc.Object {
	rv := objc.CallMethod[objc.Object](ac, objc.GetSelector("defaultValueForKey:"), key)
	return rv
}

func Animation_DefaultValueForKey(key string) objc.Object {
	return AnimationClass.DefaultValueForKey(key)
}

func (a_ Animation) ShouldArchiveValueForKey(key string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("shouldArchiveValueForKey:"), key)
	return rv
}

func (a_ Animation) IsRemovedOnCompletion() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("isRemovedOnCompletion"))
	return rv
}

func (a_ Animation) SetRemovedOnCompletion(value bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setRemovedOnCompletion:"), value)
}

func (a_ Animation) TimingFunction() MediaTimingFunction {
	rv := objc.CallMethod[MediaTimingFunction](a_, objc.GetSelector("timingFunction"))
	return rv
}

func (a_ Animation) SetTimingFunction(value IMediaTimingFunction) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setTimingFunction:"), objc.ExtractPtr(value))
}

func (a_ Animation) Delegate() AnimationDelegateWrapper {
	rv := objc.CallMethod[AnimationDelegateWrapper](a_, objc.GetSelector("delegate"))
	return rv
}

func (a_ Animation) SetDelegate(value IAnimationDelegate) {
	po := objc.WrapAsProtocol("CAAnimationDelegate", value)
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setDelegate:"), po)
}

func (a_ Animation) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (a_ Animation) PreferredFrameRateRange() FrameRateRange {
	rv := objc.CallMethod[FrameRateRange](a_, objc.GetSelector("preferredFrameRateRange"))
	return rv
}

func (a_ Animation) SetPreferredFrameRateRange(value FrameRateRange) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setPreferredFrameRateRange:"), value)
}
