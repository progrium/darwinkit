// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SliderAccessoryBehavior] class.
var SliderAccessoryBehaviorClass = _SliderAccessoryBehaviorClass{objc.GetClass("NSSliderAccessoryBehavior")}

type _SliderAccessoryBehaviorClass struct {
	objc.Class
}

// An interface definition for the [SliderAccessoryBehavior] class.
type ISliderAccessoryBehavior interface {
	objc.IObject
	HandleAction(sender ISliderAccessory)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessorybehavior?language=objc
type SliderAccessoryBehavior struct {
	objc.Object
}

func SliderAccessoryBehaviorFrom(ptr unsafe.Pointer) SliderAccessoryBehavior {
	return SliderAccessoryBehavior{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SliderAccessoryBehaviorClass) Alloc() SliderAccessoryBehavior {
	rv := objc.Call[SliderAccessoryBehavior](sc, objc.Sel("alloc"))
	return rv
}

func SliderAccessoryBehavior_Alloc() SliderAccessoryBehavior {
	return SliderAccessoryBehaviorClass.Alloc()
}

func (sc _SliderAccessoryBehaviorClass) New() SliderAccessoryBehavior {
	rv := objc.Call[SliderAccessoryBehavior](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSliderAccessoryBehavior() SliderAccessoryBehavior {
	return SliderAccessoryBehaviorClass.New()
}

func (s_ SliderAccessoryBehavior) Init() SliderAccessoryBehavior {
	rv := objc.Call[SliderAccessoryBehavior](s_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessorybehavior/2544778-behaviorwithhandler?language=objc
func (sc _SliderAccessoryBehaviorClass) BehaviorWithHandler(handler func(arg0 SliderAccessory)) SliderAccessoryBehavior {
	rv := objc.Call[SliderAccessoryBehavior](sc, objc.Sel("behaviorWithHandler:"), handler)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessorybehavior/2544778-behaviorwithhandler?language=objc
func SliderAccessoryBehavior_BehaviorWithHandler(handler func(arg0 SliderAccessory)) SliderAccessoryBehavior {
	return SliderAccessoryBehaviorClass.BehaviorWithHandler(handler)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessorybehavior/2544777-behaviorwithtarget?language=objc
func (sc _SliderAccessoryBehaviorClass) BehaviorWithTargetAction(target objc.IObject, action objc.Selector) SliderAccessoryBehavior {
	rv := objc.Call[SliderAccessoryBehavior](sc, objc.Sel("behaviorWithTarget:action:"), target, action)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessorybehavior/2544777-behaviorwithtarget?language=objc
func SliderAccessoryBehavior_BehaviorWithTargetAction(target objc.IObject, action objc.Selector) SliderAccessoryBehavior {
	return SliderAccessoryBehaviorClass.BehaviorWithTargetAction(target, action)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessorybehavior/2544759-handleaction?language=objc
func (s_ SliderAccessoryBehavior) HandleAction(sender ISliderAccessory) {
	objc.Call[objc.Void](s_, objc.Sel("handleAction:"), objc.Ptr(sender))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessorybehavior/2544771-valuestepbehavior?language=objc
func (sc _SliderAccessoryBehaviorClass) ValueStepBehavior() SliderAccessoryBehavior {
	rv := objc.Call[SliderAccessoryBehavior](sc, objc.Sel("valueStepBehavior"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessorybehavior/2544771-valuestepbehavior?language=objc
func SliderAccessoryBehavior_ValueStepBehavior() SliderAccessoryBehavior {
	return SliderAccessoryBehaviorClass.ValueStepBehavior()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessorybehavior/2544780-automaticbehavior?language=objc
func (sc _SliderAccessoryBehaviorClass) AutomaticBehavior() SliderAccessoryBehavior {
	rv := objc.Call[SliderAccessoryBehavior](sc, objc.Sel("automaticBehavior"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessorybehavior/2544780-automaticbehavior?language=objc
func SliderAccessoryBehavior_AutomaticBehavior() SliderAccessoryBehavior {
	return SliderAccessoryBehaviorClass.AutomaticBehavior()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessorybehavior/2544659-valueresetbehavior?language=objc
func (sc _SliderAccessoryBehaviorClass) ValueResetBehavior() SliderAccessoryBehavior {
	rv := objc.Call[SliderAccessoryBehavior](sc, objc.Sel("valueResetBehavior"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessorybehavior/2544659-valueresetbehavior?language=objc
func SliderAccessoryBehavior_ValueResetBehavior() SliderAccessoryBehavior {
	return SliderAccessoryBehaviorClass.ValueResetBehavior()
}
