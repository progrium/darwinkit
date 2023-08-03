// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var BasicAnimationClass = _BasicAnimationClass{objc.GetClass("CABasicAnimation")}

type _BasicAnimationClass struct {
	objc.Class
}

type IBasicAnimation interface {
	IPropertyAnimation
	FromValue() objc.Object
	SetFromValue(value objc.IObject)
	ToValue() objc.Object
	SetToValue(value objc.IObject)
	ByValue() objc.Object
	SetByValue(value objc.IObject)
}

type BasicAnimation struct {
	PropertyAnimation
}

func MakeBasicAnimation(ptr unsafe.Pointer) BasicAnimation {
	return BasicAnimation{
		PropertyAnimation: MakePropertyAnimation(ptr),
	}
}

func (bc _BasicAnimationClass) AnimationWithKeyPath(path string) BasicAnimation {
	rv := objc.CallMethod[BasicAnimation](bc, objc.GetSelector("animationWithKeyPath:"), path)
	return rv
}

func BasicAnimation_AnimationWithKeyPath(path string) BasicAnimation {
	return BasicAnimationClass.AnimationWithKeyPath(path)
}

func (bc _BasicAnimationClass) Animation() BasicAnimation {
	rv := objc.CallMethod[BasicAnimation](bc, objc.GetSelector("animation"))
	return rv
}

func BasicAnimation_Animation() BasicAnimation {
	return BasicAnimationClass.Animation()
}

func (bc _BasicAnimationClass) Alloc() BasicAnimation {
	rv := objc.CallMethod[BasicAnimation](bc, objc.GetSelector("alloc"))
	return rv
}

func BasicAnimation_Alloc() BasicAnimation {
	return BasicAnimationClass.Alloc()
}

func (bc _BasicAnimationClass) New() BasicAnimation {
	rv := objc.CallMethod[BasicAnimation](bc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewBasicAnimation() BasicAnimation {
	return BasicAnimationClass.New()
}

func BasicAnimation_New() BasicAnimation {
	return BasicAnimationClass.New()
}

func (b_ BasicAnimation) Init() BasicAnimation {
	rv := objc.CallMethod[BasicAnimation](b_, objc.GetSelector("init"))
	return rv
}

func BasicAnimation_Init() BasicAnimation {
	return BasicAnimationClass.Alloc().Init()
}

func (b_ BasicAnimation) FromValue() objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("fromValue"))
	return rv
}

func (b_ BasicAnimation) SetFromValue(value objc.IObject) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setFromValue:"), objc.ExtractPtr(value))
}

func (b_ BasicAnimation) ToValue() objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("toValue"))
	return rv
}

func (b_ BasicAnimation) SetToValue(value objc.IObject) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setToValue:"), objc.ExtractPtr(value))
}

func (b_ BasicAnimation) ByValue() objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("byValue"))
	return rv
}

func (b_ BasicAnimation) SetByValue(value objc.IObject) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setByValue:"), objc.ExtractPtr(value))
}
