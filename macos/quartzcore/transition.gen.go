// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var TransitionClass = _TransitionClass{objc.GetClass("CATransition")}

type _TransitionClass struct {
	objc.Class
}

type ITransition interface {
	IAnimation
	StartProgress() float32
	SetStartProgress(value float32)
	EndProgress() float32
	SetEndProgress(value float32)
	Type() TransitionType
	SetType(value TransitionType)
	Subtype() TransitionSubtype
	SetSubtype(value TransitionSubtype)
	Filter() objc.Object
	SetFilter(value objc.IObject)
}

type Transition struct {
	Animation
}

func MakeTransition(ptr unsafe.Pointer) Transition {
	return Transition{
		Animation: MakeAnimation(ptr),
	}
}

func (tc _TransitionClass) Animation() Transition {
	rv := objc.CallMethod[Transition](tc, objc.GetSelector("animation"))
	return rv
}

func Transition_Animation() Transition {
	return TransitionClass.Animation()
}

func (tc _TransitionClass) Alloc() Transition {
	rv := objc.CallMethod[Transition](tc, objc.GetSelector("alloc"))
	return rv
}

func Transition_Alloc() Transition {
	return TransitionClass.Alloc()
}

func (tc _TransitionClass) New() Transition {
	rv := objc.CallMethod[Transition](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTransition() Transition {
	return TransitionClass.New()
}

func Transition_New() Transition {
	return TransitionClass.New()
}

func (t_ Transition) Init() Transition {
	rv := objc.CallMethod[Transition](t_, objc.GetSelector("init"))
	return rv
}

func Transition_Init() Transition {
	return TransitionClass.Alloc().Init()
}

func (t_ Transition) StartProgress() float32 {
	rv := objc.CallMethod[float32](t_, objc.GetSelector("startProgress"))
	return rv
}

func (t_ Transition) SetStartProgress(value float32) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setStartProgress:"), value)
}

func (t_ Transition) EndProgress() float32 {
	rv := objc.CallMethod[float32](t_, objc.GetSelector("endProgress"))
	return rv
}

func (t_ Transition) SetEndProgress(value float32) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setEndProgress:"), value)
}

func (t_ Transition) Type() TransitionType {
	rv := objc.CallMethod[TransitionType](t_, objc.GetSelector("type"))
	return rv
}

func (t_ Transition) SetType(value TransitionType) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setType:"), value)
}

func (t_ Transition) Subtype() TransitionSubtype {
	rv := objc.CallMethod[TransitionSubtype](t_, objc.GetSelector("subtype"))
	return rv
}

func (t_ Transition) SetSubtype(value TransitionSubtype) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSubtype:"), value)
}

func (t_ Transition) Filter() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.GetSelector("filter"))
	return rv
}

func (t_ Transition) SetFilter(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setFilter:"), objc.ExtractPtr(value))
}
