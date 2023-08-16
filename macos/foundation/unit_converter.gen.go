// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitConverter] class.
var UnitConverterClass = _UnitConverterClass{objc.GetClass("NSUnitConverter")}

type _UnitConverterClass struct {
	objc.Class
}

// An interface definition for the [UnitConverter] class.
type IUnitConverter interface {
	objc.IObject
	BaseUnitValueFromValue(value float64) float64
	ValueFromBaseUnitValue(baseUnitValue float64) float64
}

// An abstract class that provides a description of how to convert a unit to and from the base unit of its dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitconverter?language=objc
type UnitConverter struct {
	objc.Object
}

func UnitConverterFrom(ptr unsafe.Pointer) UnitConverter {
	return UnitConverter{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _UnitConverterClass) Alloc() UnitConverter {
	rv := objc.Call[UnitConverter](uc, objc.Sel("alloc"))
	return rv
}

func UnitConverter_Alloc() UnitConverter {
	return UnitConverterClass.Alloc()
}

func (uc _UnitConverterClass) New() UnitConverter {
	rv := objc.Call[UnitConverter](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitConverter() UnitConverter {
	return UnitConverterClass.New()
}

func (u_ UnitConverter) Init() UnitConverter {
	rv := objc.Call[UnitConverter](u_, objc.Sel("init"))
	return rv
}

// For a given unit, returns the specified value of that unit in terms of the base unit of its dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitconverter/1823668-baseunitvaluefromvalue?language=objc
func (u_ UnitConverter) BaseUnitValueFromValue(value float64) float64 {
	rv := objc.Call[float64](u_, objc.Sel("baseUnitValueFromValue:"), value)
	return rv
}

// For a given unit, returns the specified value of the base unit in terms of that unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitconverter/1823657-valuefrombaseunitvalue?language=objc
func (u_ UnitConverter) ValueFromBaseUnitValue(baseUnitValue float64) float64 {
	rv := objc.Call[float64](u_, objc.Sel("valueFromBaseUnitValue:"), baseUnitValue)
	return rv
}
