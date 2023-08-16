// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corefoundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SpringAnimation] class.
var SpringAnimationClass = _SpringAnimationClass{objc.GetClass("CASpringAnimation")}

type _SpringAnimationClass struct {
	objc.Class
}

// An interface definition for the [SpringAnimation] class.
type ISpringAnimation interface {
	IBasicAnimation
	Stiffness() float64
	SetStiffness(value float64)
	InitialVelocity() float64
	SetInitialVelocity(value float64)
	Mass() float64
	SetMass(value float64)
	Damping() float64
	SetDamping(value float64)
	SettlingDuration() corefoundation.TimeInterval
}

// An animation that applies a spring-like force to a layer's properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation?language=objc
type SpringAnimation struct {
	BasicAnimation
}

func SpringAnimationFrom(ptr unsafe.Pointer) SpringAnimation {
	return SpringAnimation{
		BasicAnimation: BasicAnimationFrom(ptr),
	}
}

func (sc _SpringAnimationClass) Alloc() SpringAnimation {
	rv := objc.Call[SpringAnimation](sc, objc.Sel("alloc"))
	return rv
}

func SpringAnimation_Alloc() SpringAnimation {
	return SpringAnimationClass.Alloc()
}

func (sc _SpringAnimationClass) New() SpringAnimation {
	rv := objc.Call[SpringAnimation](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSpringAnimation() SpringAnimation {
	return SpringAnimationClass.New()
}

func (s_ SpringAnimation) Init() SpringAnimation {
	rv := objc.Call[SpringAnimation](s_, objc.Sel("init"))
	return rv
}

func (sc _SpringAnimationClass) AnimationWithKeyPath(path string) SpringAnimation {
	rv := objc.Call[SpringAnimation](sc, objc.Sel("animationWithKeyPath:"), path)
	return rv
}

// Creates and returns an CAPropertyAnimation instance for the specified key path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/1412534-animationwithkeypath?language=objc
func SpringAnimation_AnimationWithKeyPath(path string) SpringAnimation {
	return SpringAnimationClass.AnimationWithKeyPath(path)
}

func (sc _SpringAnimationClass) Animation() SpringAnimation {
	rv := objc.Call[SpringAnimation](sc, objc.Sel("animation"))
	return rv
}

// Creates and returns a new CAAnimation instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1412479-animation?language=objc
func SpringAnimation_Animation() SpringAnimation {
	return SpringAnimationClass.Animation()
}

// The spring stiffness coefficient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/1412515-stiffness?language=objc
func (s_ SpringAnimation) Stiffness() float64 {
	rv := objc.Call[float64](s_, objc.Sel("stiffness"))
	return rv
}

// The spring stiffness coefficient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/1412515-stiffness?language=objc
func (s_ SpringAnimation) SetStiffness(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setStiffness:"), value)
}

// The initial velocity of the object attached to the spring. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/1412443-initialvelocity?language=objc
func (s_ SpringAnimation) InitialVelocity() float64 {
	rv := objc.Call[float64](s_, objc.Sel("initialVelocity"))
	return rv
}

// The initial velocity of the object attached to the spring. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/1412443-initialvelocity?language=objc
func (s_ SpringAnimation) SetInitialVelocity(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setInitialVelocity:"), value)
}

// The mass of the object attached to the end of the spring. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/1412540-mass?language=objc
func (s_ SpringAnimation) Mass() float64 {
	rv := objc.Call[float64](s_, objc.Sel("mass"))
	return rv
}

// The mass of the object attached to the end of the spring. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/1412540-mass?language=objc
func (s_ SpringAnimation) SetMass(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMass:"), value)
}

// Defines how the spring’s motion should be damped due to the forces of friction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/1412532-damping?language=objc
func (s_ SpringAnimation) Damping() float64 {
	rv := objc.Call[float64](s_, objc.Sel("damping"))
	return rv
}

// Defines how the spring’s motion should be damped due to the forces of friction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/1412532-damping?language=objc
func (s_ SpringAnimation) SetDamping(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setDamping:"), value)
}

// The estimated duration required for the spring system to be considered at rest. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/1412524-settlingduration?language=objc
func (s_ SpringAnimation) SettlingDuration() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](s_, objc.Sel("settlingDuration"))
	return rv
}
