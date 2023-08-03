// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var SwitchClass = _SwitchClass{objc.GetClass("NSSwitch")}

type _SwitchClass struct {
	objc.Class
}

type ISwitch interface {
	IControl
	State() ControlStateValue
	SetState(value ControlStateValue)
}

type Switch struct {
	Control
}

func MakeSwitch(ptr unsafe.Pointer) Switch {
	return Switch{
		Control: MakeControl(ptr),
	}
}

func (s_ Switch) InitWithFrame(frameRect foundation.Rect) Switch {
	rv := objc.CallMethod[Switch](s_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func Switch_InitWithFrame(frameRect foundation.Rect) Switch {
	return SwitchClass.Alloc().InitWithFrame(frameRect)
}

func (s_ Switch) Init() Switch {
	rv := objc.CallMethod[Switch](s_, objc.GetSelector("init"))
	return rv
}

func Switch_Init() Switch {
	return SwitchClass.Alloc().Init()
}

func (sc _SwitchClass) Alloc() Switch {
	rv := objc.CallMethod[Switch](sc, objc.GetSelector("alloc"))
	return rv
}

func Switch_Alloc() Switch {
	return SwitchClass.Alloc()
}

func (sc _SwitchClass) New() Switch {
	rv := objc.CallMethod[Switch](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewSwitch() Switch {
	return SwitchClass.New()
}

func Switch_New() Switch {
	return SwitchClass.New()
}

func (s_ Switch) State() ControlStateValue {
	rv := objc.CallMethod[ControlStateValue](s_, objc.GetSelector("state"))
	return rv
}

func (s_ Switch) SetState(value ControlStateValue) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setState:"), value)
}
