// Code generated by DarwinKit. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [PropertyAnimation] class.
var PropertyAnimationClass = _PropertyAnimationClass{objc.GetClass("CAPropertyAnimation")}

type _PropertyAnimationClass struct {
	objc.Class
}

// An interface definition for the [PropertyAnimation] class.
type IPropertyAnimation interface {
	IAnimation
	ValueFunction() ValueFunction
	SetValueFunction(value IValueFunction)
	IsCumulative() bool
	SetCumulative(value bool)
	IsAdditive() bool
	SetAdditive(value bool)
	KeyPath() string
	SetKeyPath(value string)
}

// An abstract subclass of CAAnimation for creating animations that manipulate the value of layer properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation?language=objc
type PropertyAnimation struct {
	Animation
}

func PropertyAnimationFrom(ptr unsafe.Pointer) PropertyAnimation {
	return PropertyAnimation{
		Animation: AnimationFrom(ptr),
	}
}

func (pc _PropertyAnimationClass) AnimationWithKeyPath(path string) PropertyAnimation {
	rv := objc.Call[PropertyAnimation](pc, objc.Sel("animationWithKeyPath:"), path)
	return rv
}

// Creates and returns an CAPropertyAnimation instance for the specified key path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/1412534-animationwithkeypath?language=objc
func PropertyAnimation_AnimationWithKeyPath(path string) PropertyAnimation {
	return PropertyAnimationClass.AnimationWithKeyPath(path)
}

func (pc _PropertyAnimationClass) Alloc() PropertyAnimation {
	rv := objc.Call[PropertyAnimation](pc, objc.Sel("alloc"))
	return rv
}

func (pc _PropertyAnimationClass) New() PropertyAnimation {
	rv := objc.Call[PropertyAnimation](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPropertyAnimation() PropertyAnimation {
	return PropertyAnimationClass.New()
}

func (p_ PropertyAnimation) Init() PropertyAnimation {
	rv := objc.Call[PropertyAnimation](p_, objc.Sel("init"))
	return rv
}

func (pc _PropertyAnimationClass) Animation() PropertyAnimation {
	rv := objc.Call[PropertyAnimation](pc, objc.Sel("animation"))
	return rv
}

// Creates and returns a new CAAnimation instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1412479-animation?language=objc
func PropertyAnimation_Animation() PropertyAnimation {
	return PropertyAnimationClass.Animation()
}

// An optional value function that is applied to interpolated values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/1412447-valuefunction?language=objc
func (p_ PropertyAnimation) ValueFunction() ValueFunction {
	rv := objc.Call[ValueFunction](p_, objc.Sel("valueFunction"))
	return rv
}

// An optional value function that is applied to interpolated values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/1412447-valuefunction?language=objc
func (p_ PropertyAnimation) SetValueFunction(value IValueFunction) {
	objc.Call[objc.Void](p_, objc.Sel("setValueFunction:"), value)
}

// Determines if the value of the property is the value at the end of the previous repeat cycle, plus the value of the current repeat cycle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/1412538-cumulative?language=objc
func (p_ PropertyAnimation) IsCumulative() bool {
	rv := objc.Call[bool](p_, objc.Sel("isCumulative"))
	return rv
}

// Determines if the value of the property is the value at the end of the previous repeat cycle, plus the value of the current repeat cycle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/1412538-cumulative?language=objc
func (p_ PropertyAnimation) SetCumulative(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setCumulative:"), value)
}

// Determines if the value specified by the animation is added to the current render tree value to produce the new render tree value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/1412493-additive?language=objc
func (p_ PropertyAnimation) IsAdditive() bool {
	rv := objc.Call[bool](p_, objc.Sel("isAdditive"))
	return rv
}

// Determines if the value specified by the animation is added to the current render tree value to produce the new render tree value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/1412493-additive?language=objc
func (p_ PropertyAnimation) SetAdditive(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setAdditive:"), value)
}

// Specifies the key path the receiver animates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/1412496-keypath?language=objc
func (p_ PropertyAnimation) KeyPath() string {
	rv := objc.Call[string](p_, objc.Sel("keyPath"))
	return rv
}

// Specifies the key path the receiver animates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/1412496-keypath?language=objc
func (p_ PropertyAnimation) SetKeyPath(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setKeyPath:"), value)
}
