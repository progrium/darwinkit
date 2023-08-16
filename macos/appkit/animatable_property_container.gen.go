// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that defines a way to add animation to an existing class with a minimum of API impact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimatablepropertycontainer?language=objc
type PAnimatablePropertyContainer interface {
	// optional
	Animator() objc.IObject
	HasAnimator() bool

	// optional
	AnimationForKey(key AnimatablePropertyKey) objc.IObject
	HasAnimationForKey() bool

	// optional
	SetAnimations(value map[AnimatablePropertyKey]objc.Object)
	HasSetAnimations() bool

	// optional
	Animations() map[AnimatablePropertyKey]objc.IObject
	HasAnimations() bool
}

// A concrete type wrapper for the [PAnimatablePropertyContainer] protocol.
type AnimatablePropertyContainerWrapper struct {
	objc.Object
}

func (a_ AnimatablePropertyContainerWrapper) HasAnimator() bool {
	return a_.RespondsToSelector(objc.Sel("animator"))
}

// Returns a proxy object for the receiver that can be used to initiate implied animation for property changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimatablepropertycontainer/1530511-animator?language=objc
func (a_ AnimatablePropertyContainerWrapper) Animator() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("animator"))
	return rv
}

func (a_ AnimatablePropertyContainerWrapper) HasAnimationForKey() bool {
	return a_.RespondsToSelector(objc.Sel("animationForKey:"))
}

// Returns the animation that should be performed for the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimatablepropertycontainer/1526209-animationforkey?language=objc
func (a_ AnimatablePropertyContainerWrapper) AnimationForKey(key AnimatablePropertyKey) objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("animationForKey:"), key)
	return rv
}

func (a_ AnimatablePropertyContainerWrapper) HasSetAnimations() bool {
	return a_.RespondsToSelector(objc.Sel("setAnimations:"))
}

// Sets the option dictionary that maps event trigger keys to animation objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimatablepropertycontainer/1534096-animations?language=objc
func (a_ AnimatablePropertyContainerWrapper) SetAnimations(value map[AnimatablePropertyKey]objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAnimations:"), value)
}

func (a_ AnimatablePropertyContainerWrapper) HasAnimations() bool {
	return a_.RespondsToSelector(objc.Sel("animations"))
}

// Sets the option dictionary that maps event trigger keys to animation objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimatablepropertycontainer/1534096-animations?language=objc
func (a_ AnimatablePropertyContainerWrapper) Animations() map[AnimatablePropertyKey]objc.Object {
	rv := objc.Call[map[AnimatablePropertyKey]objc.Object](a_, objc.Sel("animations"))
	return rv
}
