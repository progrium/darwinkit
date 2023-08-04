// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var KeyframeAnimationClass = _KeyframeAnimationClass{objc.GetClass("CAKeyframeAnimation")}

type _KeyframeAnimationClass struct {
	objc.Class
}

type IKeyframeAnimation interface {
	IPropertyAnimation
	Values() []objc.Object
	SetValues(value []objc.IObject)
	Path() coregraphics.PathRef
	SetPath(value coregraphics.PathRef)
	KeyTimes() []foundation.Number
	SetKeyTimes(value []foundation.INumber)
	TimingFunctions() []MediaTimingFunction
	SetTimingFunctions(value []IMediaTimingFunction)
	CalculationMode() AnimationCalculationMode
	SetCalculationMode(value AnimationCalculationMode)
	RotationMode() AnimationRotationMode
	SetRotationMode(value AnimationRotationMode)
	TensionValues() []foundation.Number
	SetTensionValues(value []foundation.INumber)
	ContinuityValues() []foundation.Number
	SetContinuityValues(value []foundation.INumber)
	BiasValues() []foundation.Number
	SetBiasValues(value []foundation.INumber)
}

type KeyframeAnimation struct {
	PropertyAnimation
}

func MakeKeyframeAnimation(ptr unsafe.Pointer) KeyframeAnimation {
	return KeyframeAnimation{
		PropertyAnimation: MakePropertyAnimation(ptr),
	}
}

func (kc _KeyframeAnimationClass) AnimationWithKeyPath(path string) KeyframeAnimation {
	rv := objc.CallMethod[KeyframeAnimation](kc, objc.GetSelector("animationWithKeyPath:"), path)
	return rv
}

func KeyframeAnimation_AnimationWithKeyPath(path string) KeyframeAnimation {
	return KeyframeAnimationClass.AnimationWithKeyPath(path)
}

func (kc _KeyframeAnimationClass) Animation() KeyframeAnimation {
	rv := objc.CallMethod[KeyframeAnimation](kc, objc.GetSelector("animation"))
	return rv
}

func KeyframeAnimation_Animation() KeyframeAnimation {
	return KeyframeAnimationClass.Animation()
}

func (kc _KeyframeAnimationClass) Alloc() KeyframeAnimation {
	rv := objc.CallMethod[KeyframeAnimation](kc, objc.GetSelector("alloc"))
	return rv
}

func KeyframeAnimation_Alloc() KeyframeAnimation {
	return KeyframeAnimationClass.Alloc()
}

func (kc _KeyframeAnimationClass) New() KeyframeAnimation {
	rv := objc.CallMethod[KeyframeAnimation](kc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewKeyframeAnimation() KeyframeAnimation {
	return KeyframeAnimationClass.New()
}

func KeyframeAnimation_New() KeyframeAnimation {
	return KeyframeAnimationClass.New()
}

func (k_ KeyframeAnimation) Init() KeyframeAnimation {
	rv := objc.CallMethod[KeyframeAnimation](k_, objc.GetSelector("init"))
	return rv
}

func KeyframeAnimation_Init() KeyframeAnimation {
	return KeyframeAnimationClass.Alloc().Init()
}

func (k_ KeyframeAnimation) Values() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](k_, objc.GetSelector("values"))
	return rv
}

func (k_ KeyframeAnimation) SetValues(value []objc.IObject) {
	objc.CallMethod[objc.Void](k_, objc.GetSelector("setValues:"), value)
}

func (k_ KeyframeAnimation) Path() coregraphics.PathRef {
	rv := objc.CallMethod[coregraphics.PathRef](k_, objc.GetSelector("path"))
	return rv
}

func (k_ KeyframeAnimation) SetPath(value coregraphics.PathRef) {
	objc.CallMethod[objc.Void](k_, objc.GetSelector("setPath:"), value)
}

func (k_ KeyframeAnimation) KeyTimes() []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](k_, objc.GetSelector("keyTimes"))
	return rv
}

func (k_ KeyframeAnimation) SetKeyTimes(value []foundation.INumber) {
	objc.CallMethod[objc.Void](k_, objc.GetSelector("setKeyTimes:"), value)
}

func (k_ KeyframeAnimation) TimingFunctions() []MediaTimingFunction {
	rv := objc.CallMethod[[]MediaTimingFunction](k_, objc.GetSelector("timingFunctions"))
	return rv
}

func (k_ KeyframeAnimation) SetTimingFunctions(value []IMediaTimingFunction) {
	objc.CallMethod[objc.Void](k_, objc.GetSelector("setTimingFunctions:"), value)
}

func (k_ KeyframeAnimation) CalculationMode() AnimationCalculationMode {
	rv := objc.CallMethod[AnimationCalculationMode](k_, objc.GetSelector("calculationMode"))
	return rv
}

func (k_ KeyframeAnimation) SetCalculationMode(value AnimationCalculationMode) {
	objc.CallMethod[objc.Void](k_, objc.GetSelector("setCalculationMode:"), value)
}

func (k_ KeyframeAnimation) RotationMode() AnimationRotationMode {
	rv := objc.CallMethod[AnimationRotationMode](k_, objc.GetSelector("rotationMode"))
	return rv
}

func (k_ KeyframeAnimation) SetRotationMode(value AnimationRotationMode) {
	objc.CallMethod[objc.Void](k_, objc.GetSelector("setRotationMode:"), value)
}

func (k_ KeyframeAnimation) TensionValues() []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](k_, objc.GetSelector("tensionValues"))
	return rv
}

func (k_ KeyframeAnimation) SetTensionValues(value []foundation.INumber) {
	objc.CallMethod[objc.Void](k_, objc.GetSelector("setTensionValues:"), value)
}

func (k_ KeyframeAnimation) ContinuityValues() []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](k_, objc.GetSelector("continuityValues"))
	return rv
}

func (k_ KeyframeAnimation) SetContinuityValues(value []foundation.INumber) {
	objc.CallMethod[objc.Void](k_, objc.GetSelector("setContinuityValues:"), value)
}

func (k_ KeyframeAnimation) BiasValues() []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](k_, objc.GetSelector("biasValues"))
	return rv
}

func (k_ KeyframeAnimation) SetBiasValues(value []foundation.INumber) {
	objc.CallMethod[objc.Void](k_, objc.GetSelector("setBiasValues:"), value)
}
