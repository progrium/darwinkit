// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var SpringAnimationClass = _SpringAnimationClass{objc.GetClass("CASpringAnimation")}

type _SpringAnimationClass struct {
	objc.Class
}

type ISpringAnimation interface {
	IBasicAnimation
	Damping() float64
	SetDamping(value float64)
	InitialVelocity() float64
	SetInitialVelocity(value float64)
	Mass() float64
	SetMass(value float64)
	Stiffness() float64
	SetStiffness(value float64)
}

type SpringAnimation struct {
	BasicAnimation
}

func MakeSpringAnimation(ptr unsafe.Pointer) SpringAnimation {
	return SpringAnimation{
		BasicAnimation: MakeBasicAnimation(ptr),
	}
}

func (sc _SpringAnimationClass) AnimationWithKeyPath(path string) SpringAnimation {
	rv := objc.CallMethod[SpringAnimation](sc, objc.GetSelector("animationWithKeyPath:"), path)
	return rv
}

func SpringAnimation_AnimationWithKeyPath(path string) SpringAnimation {
	return SpringAnimationClass.AnimationWithKeyPath(path)
}

func (sc _SpringAnimationClass) Animation() SpringAnimation {
	rv := objc.CallMethod[SpringAnimation](sc, objc.GetSelector("animation"))
	return rv
}

func SpringAnimation_Animation() SpringAnimation {
	return SpringAnimationClass.Animation()
}

func (sc _SpringAnimationClass) Alloc() SpringAnimation {
	rv := objc.CallMethod[SpringAnimation](sc, objc.GetSelector("alloc"))
	return rv
}

func SpringAnimation_Alloc() SpringAnimation {
	return SpringAnimationClass.Alloc()
}

func (sc _SpringAnimationClass) New() SpringAnimation {
	rv := objc.CallMethod[SpringAnimation](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewSpringAnimation() SpringAnimation {
	return SpringAnimationClass.New()
}

func SpringAnimation_New() SpringAnimation {
	return SpringAnimationClass.New()
}

func (s_ SpringAnimation) Init() SpringAnimation {
	rv := objc.CallMethod[SpringAnimation](s_, objc.GetSelector("init"))
	return rv
}

func SpringAnimation_Init() SpringAnimation {
	return SpringAnimationClass.Alloc().Init()
}

func (s_ SpringAnimation) Damping() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("damping"))
	return rv
}

func (s_ SpringAnimation) SetDamping(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDamping:"), value)
}

func (s_ SpringAnimation) InitialVelocity() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("initialVelocity"))
	return rv
}

func (s_ SpringAnimation) SetInitialVelocity(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setInitialVelocity:"), value)
}

func (s_ SpringAnimation) Mass() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("mass"))
	return rv
}

func (s_ SpringAnimation) SetMass(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMass:"), value)
}

func (s_ SpringAnimation) Stiffness() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("stiffness"))
	return rv
}

func (s_ SpringAnimation) SetStiffness(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setStiffness:"), value)
}
