// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [StepperCell] class.
var StepperCellClass = _StepperCellClass{objc.GetClass("NSStepperCell")}

type _StepperCellClass struct {
	objc.Class
}

// An interface definition for the [StepperCell] class.
type IStepperCell interface {
	IActionCell
	ValueWraps() bool
	SetValueWraps(value bool)
	MinValue() float64
	SetMinValue(value float64)
	Increment() float64
	SetIncrement(value float64)
	MaxValue() float64
	SetMaxValue(value float64)
	Autorepeat() bool
	SetAutorepeat(value bool)
}

// An NSStepperCell object controls the appearance and behavior of an NSStepper object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppercell?language=objc
type StepperCell struct {
	ActionCell
}

func StepperCellFrom(ptr unsafe.Pointer) StepperCell {
	return StepperCell{
		ActionCell: ActionCellFrom(ptr),
	}
}

func (sc _StepperCellClass) Alloc() StepperCell {
	rv := objc.Call[StepperCell](sc, objc.Sel("alloc"))
	return rv
}

func StepperCell_Alloc() StepperCell {
	return StepperCellClass.Alloc()
}

func (sc _StepperCellClass) New() StepperCell {
	rv := objc.Call[StepperCell](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStepperCell() StepperCell {
	return StepperCellClass.New()
}

func (s_ StepperCell) Init() StepperCell {
	rv := objc.Call[StepperCell](s_, objc.Sel("init"))
	return rv
}

func (s_ StepperCell) InitImageCell(image IImage) StepperCell {
	rv := objc.Call[StepperCell](s_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

// Returns an NSCell object initialized with the specified image and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533898-initimagecell?language=objc
func StepperCell_InitImageCell(image IImage) StepperCell {
	return StepperCellClass.Alloc().InitImageCell(image)
}

func (s_ StepperCell) InitTextCell(string_ string) StepperCell {
	rv := objc.Call[StepperCell](s_, objc.Sel("initTextCell:"), string_)
	return rv
}

// Returns an NSCell object initialized with the specified string and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530851-inittextcell?language=objc
func StepperCell_InitTextCell(string_ string) StepperCell {
	return StepperCellClass.Alloc().InitTextCell(string_)
}

// A Boolean value indicating whether the receiver wraps around the minimum and maximum values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppercell/1392325-valuewraps?language=objc
func (s_ StepperCell) ValueWraps() bool {
	rv := objc.Call[bool](s_, objc.Sel("valueWraps"))
	return rv
}

// A Boolean value indicating whether the receiver wraps around the minimum and maximum values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppercell/1392325-valuewraps?language=objc
func (s_ StepperCell) SetValueWraps(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setValueWraps:"), value)
}

// The minimum value for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppercell/1392327-minvalue?language=objc
func (s_ StepperCell) MinValue() float64 {
	rv := objc.Call[float64](s_, objc.Sel("minValue"))
	return rv
}

// The minimum value for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppercell/1392327-minvalue?language=objc
func (s_ StepperCell) SetMinValue(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMinValue:"), value)
}

// The amount by which the receiver will change per increment or decrement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppercell/1392331-increment?language=objc
func (s_ StepperCell) Increment() float64 {
	rv := objc.Call[float64](s_, objc.Sel("increment"))
	return rv
}

// The amount by which the receiver will change per increment or decrement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppercell/1392331-increment?language=objc
func (s_ StepperCell) SetIncrement(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setIncrement:"), value)
}

// The maximum value for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppercell/1392321-maxvalue?language=objc
func (s_ StepperCell) MaxValue() float64 {
	rv := objc.Call[float64](s_, objc.Sel("maxValue"))
	return rv
}

// The maximum value for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppercell/1392321-maxvalue?language=objc
func (s_ StepperCell) SetMaxValue(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMaxValue:"), value)
}

// A Boolean value indicating how the receiver responds to mouse events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppercell/1392323-autorepeat?language=objc
func (s_ StepperCell) Autorepeat() bool {
	rv := objc.Call[bool](s_, objc.Sel("autorepeat"))
	return rv
}

// A Boolean value indicating how the receiver responds to mouse events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppercell/1392323-autorepeat?language=objc
func (s_ StepperCell) SetAutorepeat(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setAutorepeat:"), value)
}
