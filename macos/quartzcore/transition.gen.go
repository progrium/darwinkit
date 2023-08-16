// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Transition] class.
var TransitionClass = _TransitionClass{objc.GetClass("CATransition")}

type _TransitionClass struct {
	objc.Class
}

// An interface definition for the [Transition] class.
type ITransition interface {
	IAnimation
	Filter() objc.Object
	SetFilter(value objc.IObject)
	EndProgress() float64
	SetEndProgress(value float64)
	Subtype() TransitionSubtype
	SetSubtype(value TransitionSubtype)
	Type() TransitionType
	SetType(value TransitionType)
	StartProgress() float64
	SetStartProgress(value float64)
}

// An object that provides an animated transition between a layer's states. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransition?language=objc
type Transition struct {
	Animation
}

func TransitionFrom(ptr unsafe.Pointer) Transition {
	return Transition{
		Animation: AnimationFrom(ptr),
	}
}

func (tc _TransitionClass) Alloc() Transition {
	rv := objc.Call[Transition](tc, objc.Sel("alloc"))
	return rv
}

func Transition_Alloc() Transition {
	return TransitionClass.Alloc()
}

func (tc _TransitionClass) New() Transition {
	rv := objc.Call[Transition](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTransition() Transition {
	return TransitionClass.New()
}

func (t_ Transition) Init() Transition {
	rv := objc.Call[Transition](t_, objc.Sel("init"))
	return rv
}

func (tc _TransitionClass) Animation() Transition {
	rv := objc.Call[Transition](tc, objc.Sel("animation"))
	return rv
}

// Creates and returns a new CAAnimation instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1412479-animation?language=objc
func Transition_Animation() Transition {
	return TransitionClass.Animation()
}

// An optional Core Image filter object that provides the transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransition/1412506-filter?language=objc
func (t_ Transition) Filter() objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("filter"))
	return rv
}

// An optional Core Image filter object that provides the transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransition/1412506-filter?language=objc
func (t_ Transition) SetFilter(value objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setFilter:"), value)
}

// Indicates the end point of the receiver as a fraction of the entire transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransition/1412509-endprogress?language=objc
func (t_ Transition) EndProgress() float64 {
	rv := objc.Call[float64](t_, objc.Sel("endProgress"))
	return rv
}

// Indicates the end point of the receiver as a fraction of the entire transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransition/1412509-endprogress?language=objc
func (t_ Transition) SetEndProgress(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setEndProgress:"), value)
}

// Specifies an optional subtype that indicates the direction for the predefined motion-based transitions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransition/1412467-subtype?language=objc
func (t_ Transition) Subtype() TransitionSubtype {
	rv := objc.Call[TransitionSubtype](t_, objc.Sel("subtype"))
	return rv
}

// Specifies an optional subtype that indicates the direction for the predefined motion-based transitions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransition/1412467-subtype?language=objc
func (t_ Transition) SetSubtype(value TransitionSubtype) {
	objc.Call[objc.Void](t_, objc.Sel("setSubtype:"), value)
}

// Specifies the predefined transition type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransition/1412502-type?language=objc
func (t_ Transition) Type() TransitionType {
	rv := objc.Call[TransitionType](t_, objc.Sel("type"))
	return rv
}

// Specifies the predefined transition type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransition/1412502-type?language=objc
func (t_ Transition) SetType(value TransitionType) {
	objc.Call[objc.Void](t_, objc.Sel("setType:"), value)
}

// Indicates the start point of the receiver as a fraction of the entire transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransition/1412511-startprogress?language=objc
func (t_ Transition) StartProgress() float64 {
	rv := objc.Call[float64](t_, objc.Sel("startProgress"))
	return rv
}

// Indicates the start point of the receiver as a fraction of the entire transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catransition/1412511-startprogress?language=objc
func (t_ Transition) SetStartProgress(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setStartProgress:"), value)
}
