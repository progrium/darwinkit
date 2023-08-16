// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Measurement] class.
var MeasurementClass = _MeasurementClass{objc.GetClass("NSMeasurement")}

type _MeasurementClass struct {
	objc.Class
}

// An interface definition for the [Measurement] class.
type IMeasurement interface {
	objc.IObject
	MeasurementByConvertingToUnit(unit IUnit) Measurement
	CanBeConvertedToUnit(unit IUnit) bool
	MeasurementBySubtractingMeasurement(measurement IMeasurement) Measurement
	MeasurementByAddingMeasurement(measurement IMeasurement) Measurement
	DoubleValue() float64
	Unit() objc.Object
}

// A numeric quantity labeled with a unit of measure, with support for unit conversion and unit-aware calculations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurement?language=objc
type Measurement struct {
	objc.Object
}

func MeasurementFrom(ptr unsafe.Pointer) Measurement {
	return Measurement{
		Object: objc.ObjectFrom(ptr),
	}
}

func (m_ Measurement) InitWithDoubleValueUnit(doubleValue float64, unit objc.IObject) Measurement {
	rv := objc.Call[Measurement](m_, objc.Sel("initWithDoubleValue:unit:"), doubleValue, objc.Ptr(unit))
	return rv
}

// Initializes a new measurement with a specified double-precision floating-point value and unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurement/1643012-initwithdoublevalue?language=objc
func Measurement_InitWithDoubleValueUnit(doubleValue float64, unit objc.IObject) Measurement {
	return MeasurementClass.Alloc().InitWithDoubleValueUnit(doubleValue, unit)
}

func (mc _MeasurementClass) Alloc() Measurement {
	rv := objc.Call[Measurement](mc, objc.Sel("alloc"))
	return rv
}

func Measurement_Alloc() Measurement {
	return MeasurementClass.Alloc()
}

func (mc _MeasurementClass) New() Measurement {
	rv := objc.Call[Measurement](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMeasurement() Measurement {
	return MeasurementClass.New()
}

func (m_ Measurement) Init() Measurement {
	rv := objc.Call[Measurement](m_, objc.Sel("init"))
	return rv
}

// Returns a measurement created by converting the receiver to the specified unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurement/1642900-measurementbyconvertingtounit?language=objc
func (m_ Measurement) MeasurementByConvertingToUnit(unit IUnit) Measurement {
	rv := objc.Call[Measurement](m_, objc.Sel("measurementByConvertingToUnit:"), objc.Ptr(unit))
	return rv
}

// Indicates whether the measurement can be converted to the given unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurement/1690850-canbeconvertedtounit?language=objc
func (m_ Measurement) CanBeConvertedToUnit(unit IUnit) bool {
	rv := objc.Call[bool](m_, objc.Sel("canBeConvertedToUnit:"), objc.Ptr(unit))
	return rv
}

// Returns a new measurement by subtracting the specified measurement from the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurement/1642872-measurementbysubtractingmeasurem?language=objc
func (m_ Measurement) MeasurementBySubtractingMeasurement(measurement IMeasurement) Measurement {
	rv := objc.Call[Measurement](m_, objc.Sel("measurementBySubtractingMeasurement:"), objc.Ptr(measurement))
	return rv
}

// Returns a new measurement by adding the receiver to the specified measurement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurement/1643170-measurementbyaddingmeasurement?language=objc
func (m_ Measurement) MeasurementByAddingMeasurement(measurement IMeasurement) Measurement {
	rv := objc.Call[Measurement](m_, objc.Sel("measurementByAddingMeasurement:"), objc.Ptr(measurement))
	return rv
}

// The measurement value, represented as a double-precision floating-point number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurement/1643124-doublevalue?language=objc
func (m_ Measurement) DoubleValue() float64 {
	rv := objc.Call[float64](m_, objc.Sel("doubleValue"))
	return rv
}

// The unit of measure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurement/1642831-unit?language=objc
func (m_ Measurement) Unit() objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("unit"))
	return rv
}
