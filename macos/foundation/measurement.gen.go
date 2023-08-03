// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var MeasurementClass = _MeasurementClass{objc.GetClass("NSMeasurement")}

type _MeasurementClass struct {
	objc.Class
}

type IMeasurement interface {
	objc.IObject
	CanBeConvertedToUnit(unit IUnit) bool
	MeasurementByConvertingToUnit(unit IUnit) Measurement
	MeasurementByAddingMeasurement(measurement IMeasurement) Measurement
	MeasurementBySubtractingMeasurement(measurement IMeasurement) Measurement
	DoubleValue() float64
}

type Measurement struct {
	objc.Object
}

func MakeMeasurement(ptr unsafe.Pointer) Measurement {
	return Measurement{
		Object: objc.MakeObject(ptr),
	}
}

func (mc _MeasurementClass) Alloc() Measurement {
	rv := objc.CallMethod[Measurement](mc, objc.GetSelector("alloc"))
	return rv
}

func Measurement_Alloc() Measurement {
	return MeasurementClass.Alloc()
}

func (mc _MeasurementClass) New() Measurement {
	rv := objc.CallMethod[Measurement](mc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewMeasurement() Measurement {
	return MeasurementClass.New()
}

func Measurement_New() Measurement {
	return MeasurementClass.New()
}

func (m_ Measurement) Init() Measurement {
	rv := objc.CallMethod[Measurement](m_, objc.GetSelector("init"))
	return rv
}

func Measurement_Init() Measurement {
	return MeasurementClass.Alloc().Init()
}

func (m_ Measurement) CanBeConvertedToUnit(unit IUnit) bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("canBeConvertedToUnit:"), objc.ExtractPtr(unit))
	return rv
}

func (m_ Measurement) MeasurementByConvertingToUnit(unit IUnit) Measurement {
	rv := objc.CallMethod[Measurement](m_, objc.GetSelector("measurementByConvertingToUnit:"), objc.ExtractPtr(unit))
	return rv
}

func (m_ Measurement) MeasurementByAddingMeasurement(measurement IMeasurement) Measurement {
	rv := objc.CallMethod[Measurement](m_, objc.GetSelector("measurementByAddingMeasurement:"), objc.ExtractPtr(measurement))
	return rv
}

func (m_ Measurement) MeasurementBySubtractingMeasurement(measurement IMeasurement) Measurement {
	rv := objc.CallMethod[Measurement](m_, objc.GetSelector("measurementBySubtractingMeasurement:"), objc.ExtractPtr(measurement))
	return rv
}

func (m_ Measurement) DoubleValue() float64 {
	rv := objc.CallMethod[float64](m_, objc.GetSelector("doubleValue"))
	return rv
}
