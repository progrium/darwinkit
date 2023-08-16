// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [StepperTouchBarItem] class.
var StepperTouchBarItemClass = _StepperTouchBarItemClass{objc.GetClass("NSStepperTouchBarItem")}

type _StepperTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [StepperTouchBarItem] class.
type IStepperTouchBarItem interface {
	ITouchBarItem
	Target() objc.Object
	SetTarget(value objc.IObject)
	SetCustomizationLabel(value string)
	Action() objc.Selector
	SetAction(value objc.Selector)
	Value() float64
	SetValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
	Increment() float64
	SetIncrement(value float64)
	MaxValue() float64
	SetMaxValue(value float64)
}

// A bar item that provides a stepper control for incrementing or decrementing a value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem?language=objc
type StepperTouchBarItem struct {
	TouchBarItem
}

func StepperTouchBarItemFrom(ptr unsafe.Pointer) StepperTouchBarItem {
	return StepperTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

func (sc _StepperTouchBarItemClass) StepperTouchBarItemWithIdentifierFormatter(identifier TouchBarItemIdentifier, formatter foundation.IFormatter) StepperTouchBarItem {
	rv := objc.Call[StepperTouchBarItem](sc, objc.Sel("stepperTouchBarItemWithIdentifier:formatter:"), identifier, objc.Ptr(formatter))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/3334936-steppertouchbaritemwithidentifie?language=objc
func StepperTouchBarItem_StepperTouchBarItemWithIdentifierFormatter(identifier TouchBarItemIdentifier, formatter foundation.IFormatter) StepperTouchBarItem {
	return StepperTouchBarItemClass.StepperTouchBarItemWithIdentifierFormatter(identifier, formatter)
}

func (sc _StepperTouchBarItemClass) Alloc() StepperTouchBarItem {
	rv := objc.Call[StepperTouchBarItem](sc, objc.Sel("alloc"))
	return rv
}

func StepperTouchBarItem_Alloc() StepperTouchBarItem {
	return StepperTouchBarItemClass.Alloc()
}

func (sc _StepperTouchBarItemClass) New() StepperTouchBarItem {
	rv := objc.Call[StepperTouchBarItem](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStepperTouchBarItem() StepperTouchBarItem {
	return StepperTouchBarItemClass.New()
}

func (s_ StepperTouchBarItem) Init() StepperTouchBarItem {
	rv := objc.Call[StepperTouchBarItem](s_, objc.Sel("init"))
	return rv
}

func (s_ StepperTouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) StepperTouchBarItem {
	rv := objc.Call[StepperTouchBarItem](s_, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Creates a new item with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/2544755-initwithidentifier?language=objc
func StepperTouchBarItem_InitWithIdentifier(identifier TouchBarItemIdentifier) StepperTouchBarItem {
	return StepperTouchBarItemClass.Alloc().InitWithIdentifier(identifier)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/3334937-target?language=objc
func (s_ StepperTouchBarItem) Target() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("target"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/3334937-target?language=objc
func (s_ StepperTouchBarItem) SetTarget(value objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setTarget:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/3334931-customizationlabel?language=objc
func (s_ StepperTouchBarItem) SetCustomizationLabel(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setCustomizationLabel:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/3334930-action?language=objc
func (s_ StepperTouchBarItem) Action() objc.Selector {
	rv := objc.Call[objc.Selector](s_, objc.Sel("action"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/3334930-action?language=objc
func (s_ StepperTouchBarItem) SetAction(value objc.Selector) {
	objc.Call[objc.Void](s_, objc.Sel("setAction:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/3334938-value?language=objc
func (s_ StepperTouchBarItem) Value() float64 {
	rv := objc.Call[float64](s_, objc.Sel("value"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/3334938-value?language=objc
func (s_ StepperTouchBarItem) SetValue(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/3334934-minvalue?language=objc
func (s_ StepperTouchBarItem) MinValue() float64 {
	rv := objc.Call[float64](s_, objc.Sel("minValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/3334934-minvalue?language=objc
func (s_ StepperTouchBarItem) SetMinValue(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMinValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/3334932-increment?language=objc
func (s_ StepperTouchBarItem) Increment() float64 {
	rv := objc.Call[float64](s_, objc.Sel("increment"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/3334932-increment?language=objc
func (s_ StepperTouchBarItem) SetIncrement(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setIncrement:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/3334933-maxvalue?language=objc
func (s_ StepperTouchBarItem) MaxValue() float64 {
	rv := objc.Call[float64](s_, objc.Sel("maxValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/3334933-maxvalue?language=objc
func (s_ StepperTouchBarItem) SetMaxValue(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMaxValue:"), value)
}
