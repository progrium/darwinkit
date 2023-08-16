// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AnimationGroup] class.
var AnimationGroupClass = _AnimationGroupClass{objc.GetClass("CAAnimationGroup")}

type _AnimationGroupClass struct {
	objc.Class
}

// An interface definition for the [AnimationGroup] class.
type IAnimationGroup interface {
	IAnimation
	Animations() []Animation
	SetAnimations(value []IAnimation)
}

// An object that allows multiple animations to be grouped and run concurrently. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimationgroup?language=objc
type AnimationGroup struct {
	Animation
}

func AnimationGroupFrom(ptr unsafe.Pointer) AnimationGroup {
	return AnimationGroup{
		Animation: AnimationFrom(ptr),
	}
}

func (ac _AnimationGroupClass) Alloc() AnimationGroup {
	rv := objc.Call[AnimationGroup](ac, objc.Sel("alloc"))
	return rv
}

func AnimationGroup_Alloc() AnimationGroup {
	return AnimationGroupClass.Alloc()
}

func (ac _AnimationGroupClass) New() AnimationGroup {
	rv := objc.Call[AnimationGroup](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAnimationGroup() AnimationGroup {
	return AnimationGroupClass.New()
}

func (a_ AnimationGroup) Init() AnimationGroup {
	rv := objc.Call[AnimationGroup](a_, objc.Sel("init"))
	return rv
}

func (ac _AnimationGroupClass) Animation() AnimationGroup {
	rv := objc.Call[AnimationGroup](ac, objc.Sel("animation"))
	return rv
}

// Creates and returns a new CAAnimation instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1412479-animation?language=objc
func AnimationGroup_Animation() AnimationGroup {
	return AnimationGroupClass.Animation()
}

// An array of CAAnimation objects to be evaluated in the time space of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimationgroup/1412516-animations?language=objc
func (a_ AnimationGroup) Animations() []Animation {
	rv := objc.Call[[]Animation](a_, objc.Sel("animations"))
	return rv
}

// An array of CAAnimation objects to be evaluated in the time space of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimationgroup/1412516-animations?language=objc
func (a_ AnimationGroup) SetAnimations(value []IAnimation) {
	objc.Call[objc.Void](a_, objc.Sel("setAnimations:"), value)
}
