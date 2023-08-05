// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ProgressIndicatorClass = _ProgressIndicatorClass{objc.GetClass("NSProgressIndicator")}

type _ProgressIndicatorClass struct {
	objc.Class
}

type IProgressIndicator interface {
	IView
	StartAnimation(sender objc.IObject)
	StopAnimation(sender objc.IObject)
	IncrementBy(delta float64)
	SizeToFit()
	UsesThreadedAnimation() bool
	SetUsesThreadedAnimation(value bool)
	DoubleValue() float64
	SetDoubleValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
	MaxValue() float64
	SetMaxValue(value float64)
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	ControlTint() ControlTint
	SetControlTint(value ControlTint)
	IsBezeled() bool
	SetBezeled(value bool)
	IsIndeterminate() bool
	SetIndeterminate(value bool)
	Style() ProgressIndicatorStyle
	SetStyle(value ProgressIndicatorStyle)
	IsDisplayedWhenStopped() bool
	SetDisplayedWhenStopped(value bool)
}

type ProgressIndicator struct {
	View
}

func MakeProgressIndicator(ptr unsafe.Pointer) ProgressIndicator {
	return ProgressIndicator{
		View: MakeView(ptr),
	}
}

func (p_ ProgressIndicator) InitWithFrame(frameRect foundation.Rect) ProgressIndicator {
	rv := objc.CallMethod[ProgressIndicator](p_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func ProgressIndicator_InitWithFrame(frameRect foundation.Rect) ProgressIndicator {
	return ProgressIndicatorClass.Alloc().InitWithFrame(frameRect)
}

func (p_ ProgressIndicator) Init() ProgressIndicator {
	rv := objc.CallMethod[ProgressIndicator](p_, objc.GetSelector("init"))
	return rv
}

func ProgressIndicator_Init() ProgressIndicator {
	return ProgressIndicatorClass.Alloc().Init()
}

func (pc _ProgressIndicatorClass) Alloc() ProgressIndicator {
	rv := objc.CallMethod[ProgressIndicator](pc, objc.GetSelector("alloc"))
	return rv
}

func ProgressIndicator_Alloc() ProgressIndicator {
	return ProgressIndicatorClass.Alloc()
}

func (pc _ProgressIndicatorClass) New() ProgressIndicator {
	rv := objc.CallMethod[ProgressIndicator](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewProgressIndicator() ProgressIndicator {
	return ProgressIndicatorClass.New()
}

func ProgressIndicator_New() ProgressIndicator {
	return ProgressIndicatorClass.New()
}

func (p_ ProgressIndicator) StartAnimation(sender objc.IObject) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("startAnimation:"), objc.ExtractPtr(sender))
}

func (p_ ProgressIndicator) StopAnimation(sender objc.IObject) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("stopAnimation:"), objc.ExtractPtr(sender))
}

func (p_ ProgressIndicator) IncrementBy(delta float64) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("incrementBy:"), delta)
}

func (p_ ProgressIndicator) SizeToFit() {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("sizeToFit"))
}

func (p_ ProgressIndicator) UsesThreadedAnimation() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("usesThreadedAnimation"))
	return rv
}

func (p_ ProgressIndicator) SetUsesThreadedAnimation(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setUsesThreadedAnimation:"), value)
}

func (p_ ProgressIndicator) DoubleValue() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("doubleValue"))
	return rv
}

func (p_ ProgressIndicator) SetDoubleValue(value float64) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setDoubleValue:"), value)
}

func (p_ ProgressIndicator) MinValue() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("minValue"))
	return rv
}

func (p_ ProgressIndicator) SetMinValue(value float64) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setMinValue:"), value)
}

func (p_ ProgressIndicator) MaxValue() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("maxValue"))
	return rv
}

func (p_ ProgressIndicator) SetMaxValue(value float64) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setMaxValue:"), value)
}

func (p_ ProgressIndicator) ControlSize() ControlSize {
	rv := objc.CallMethod[ControlSize](p_, objc.GetSelector("controlSize"))
	return rv
}

func (p_ ProgressIndicator) SetControlSize(value ControlSize) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setControlSize:"), value)
}

func (p_ ProgressIndicator) ControlTint() ControlTint {
	rv := objc.CallMethod[ControlTint](p_, objc.GetSelector("controlTint"))
	return rv
}

func (p_ ProgressIndicator) SetControlTint(value ControlTint) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setControlTint:"), value)
}

func (p_ ProgressIndicator) IsBezeled() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isBezeled"))
	return rv
}

func (p_ ProgressIndicator) SetBezeled(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setBezeled:"), value)
}

func (p_ ProgressIndicator) IsIndeterminate() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isIndeterminate"))
	return rv
}

func (p_ ProgressIndicator) SetIndeterminate(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setIndeterminate:"), value)
}

func (p_ ProgressIndicator) Style() ProgressIndicatorStyle {
	rv := objc.CallMethod[ProgressIndicatorStyle](p_, objc.GetSelector("style"))
	return rv
}

func (p_ ProgressIndicator) SetStyle(value ProgressIndicatorStyle) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setStyle:"), value)
}

func (p_ ProgressIndicator) IsDisplayedWhenStopped() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isDisplayedWhenStopped"))
	return rv
}

func (p_ ProgressIndicator) SetDisplayedWhenStopped(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setDisplayedWhenStopped:"), value)
}
