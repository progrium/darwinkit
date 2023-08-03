// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var StepperClass = _StepperClass{objc.GetClass("NSStepper")}

type _StepperClass struct {
	objc.Class
}

type IStepper interface {
	IControl
	MaxValue() float64
	SetMaxValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
	Increment() float64
	SetIncrement(value float64)
	Autorepeat() bool
	SetAutorepeat(value bool)
	ValueWraps() bool
	SetValueWraps(value bool)
}

type Stepper struct {
	Control
}

func MakeStepper(ptr unsafe.Pointer) Stepper {
	return Stepper{
		Control: MakeControl(ptr),
	}
}

func (s_ Stepper) InitWithFrame(frameRect foundation.Rect) Stepper {
	rv := objc.CallMethod[Stepper](s_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func Stepper_InitWithFrame(frameRect foundation.Rect) Stepper {
	return StepperClass.Alloc().InitWithFrame(frameRect)
}

func (s_ Stepper) Init() Stepper {
	rv := objc.CallMethod[Stepper](s_, objc.GetSelector("init"))
	return rv
}

func Stepper_Init() Stepper {
	return StepperClass.Alloc().Init()
}

func (sc _StepperClass) Alloc() Stepper {
	rv := objc.CallMethod[Stepper](sc, objc.GetSelector("alloc"))
	return rv
}

func Stepper_Alloc() Stepper {
	return StepperClass.Alloc()
}

func (sc _StepperClass) New() Stepper {
	rv := objc.CallMethod[Stepper](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewStepper() Stepper {
	return StepperClass.New()
}

func Stepper_New() Stepper {
	return StepperClass.New()
}

func (s_ Stepper) MaxValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("maxValue"))
	return rv
}

func (s_ Stepper) SetMaxValue(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMaxValue:"), value)
}

func (s_ Stepper) MinValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("minValue"))
	return rv
}

func (s_ Stepper) SetMinValue(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMinValue:"), value)
}

func (s_ Stepper) Increment() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("increment"))
	return rv
}

func (s_ Stepper) SetIncrement(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setIncrement:"), value)
}

func (s_ Stepper) Autorepeat() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("autorepeat"))
	return rv
}

func (s_ Stepper) SetAutorepeat(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAutorepeat:"), value)
}

func (s_ Stepper) ValueWraps() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("valueWraps"))
	return rv
}

func (s_ Stepper) SetValueWraps(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setValueWraps:"), value)
}
