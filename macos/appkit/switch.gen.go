// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Switch] class.
var SwitchClass = _SwitchClass{objc.GetClass("NSSwitch")}

type _SwitchClass struct {
	objc.Class
}

// An interface definition for the [Switch] class.
type ISwitch interface {
	IControl
	State() ControlStateValue
	SetState(value ControlStateValue)
}

// A control that offers a binary choice. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsswitch?language=objc
type Switch struct {
	Control
}

func SwitchFrom(ptr unsafe.Pointer) Switch {
	return Switch{
		Control: ControlFrom(ptr),
	}
}

func (sc _SwitchClass) Alloc() Switch {
	rv := objc.Call[Switch](sc, objc.Sel("alloc"))
	return rv
}

func Switch_Alloc() Switch {
	return SwitchClass.Alloc()
}

func (sc _SwitchClass) New() Switch {
	rv := objc.Call[Switch](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSwitch() Switch {
	return SwitchClass.New()
}

func (s_ Switch) Init() Switch {
	rv := objc.Call[Switch](s_, objc.Sel("init"))
	return rv
}

func (s_ Switch) InitWithFrame(frameRect foundation.Rect) Switch {
	rv := objc.Call[Switch](s_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func NewSwitchWithFrame(frameRect foundation.Rect) Switch {
	instance := SwitchClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// The current position of the switch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsswitch/3172699-state?language=objc
func (s_ Switch) State() ControlStateValue {
	rv := objc.Call[ControlStateValue](s_, objc.Sel("state"))
	return rv
}

// The current position of the switch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsswitch/3172699-state?language=objc
func (s_ Switch) SetState(value ControlStateValue) {
	objc.Call[objc.Void](s_, objc.Sel("setState:"), value)
}
