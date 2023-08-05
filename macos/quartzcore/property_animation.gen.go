// AUTO-GENERATED CODE, DO NOT MODIFY
package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var PropertyAnimationClass = _PropertyAnimationClass{objc.GetClass("CAPropertyAnimation")}

type _PropertyAnimationClass struct {
	objc.Class
}

type IPropertyAnimation interface {
	IAnimation
	KeyPath() string
	SetKeyPath(value string)
	IsCumulative() bool
	SetCumulative(value bool)
	IsAdditive() bool
	SetAdditive(value bool)
	ValueFunction() ValueFunction
	SetValueFunction(value IValueFunction)
}

type PropertyAnimation struct {
	Animation
}

func MakePropertyAnimation(ptr unsafe.Pointer) PropertyAnimation {
	return PropertyAnimation{
		Animation: MakeAnimation(ptr),
	}
}

func (pc _PropertyAnimationClass) AnimationWithKeyPath(path string) PropertyAnimation {
	rv := objc.CallMethod[PropertyAnimation](pc, objc.GetSelector("animationWithKeyPath:"), path)
	return rv
}

func PropertyAnimation_AnimationWithKeyPath(path string) PropertyAnimation {
	return PropertyAnimationClass.AnimationWithKeyPath(path)
}

func (pc _PropertyAnimationClass) Animation() PropertyAnimation {
	rv := objc.CallMethod[PropertyAnimation](pc, objc.GetSelector("animation"))
	return rv
}

func PropertyAnimation_Animation() PropertyAnimation {
	return PropertyAnimationClass.Animation()
}

func (pc _PropertyAnimationClass) Alloc() PropertyAnimation {
	rv := objc.CallMethod[PropertyAnimation](pc, objc.GetSelector("alloc"))
	return rv
}

func PropertyAnimation_Alloc() PropertyAnimation {
	return PropertyAnimationClass.Alloc()
}

func (pc _PropertyAnimationClass) New() PropertyAnimation {
	rv := objc.CallMethod[PropertyAnimation](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPropertyAnimation() PropertyAnimation {
	return PropertyAnimationClass.New()
}

func PropertyAnimation_New() PropertyAnimation {
	return PropertyAnimationClass.New()
}

func (p_ PropertyAnimation) Init() PropertyAnimation {
	rv := objc.CallMethod[PropertyAnimation](p_, objc.GetSelector("init"))
	return rv
}

func PropertyAnimation_Init() PropertyAnimation {
	return PropertyAnimationClass.Alloc().Init()
}

func (p_ PropertyAnimation) KeyPath() string {
	rv := objc.CallMethod[string](p_, objc.GetSelector("keyPath"))
	return rv
}

func (p_ PropertyAnimation) SetKeyPath(value string) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setKeyPath:"), value)
}

func (p_ PropertyAnimation) IsCumulative() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isCumulative"))
	return rv
}

func (p_ PropertyAnimation) SetCumulative(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setCumulative:"), value)
}

func (p_ PropertyAnimation) IsAdditive() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isAdditive"))
	return rv
}

func (p_ PropertyAnimation) SetAdditive(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setAdditive:"), value)
}

func (p_ PropertyAnimation) ValueFunction() ValueFunction {
	rv := objc.CallMethod[ValueFunction](p_, objc.GetSelector("valueFunction"))
	return rv
}

func (p_ PropertyAnimation) SetValueFunction(value IValueFunction) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setValueFunction:"), objc.ExtractPtr(value))
}
