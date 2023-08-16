// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitTemperature] class.
var UnitTemperatureClass = _UnitTemperatureClass{objc.GetClass("NSUnitTemperature")}

type _UnitTemperatureClass struct {
	objc.Class
}

// An interface definition for the [UnitTemperature] class.
type IUnitTemperature interface {
	IDimension
}

// A unit of measure for temperature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunittemperature?language=objc
type UnitTemperature struct {
	Dimension
}

func UnitTemperatureFrom(ptr unsafe.Pointer) UnitTemperature {
	return UnitTemperature{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitTemperatureClass) Alloc() UnitTemperature {
	rv := objc.Call[UnitTemperature](uc, objc.Sel("alloc"))
	return rv
}

func UnitTemperature_Alloc() UnitTemperature {
	return UnitTemperatureClass.Alloc()
}

func (uc _UnitTemperatureClass) New() UnitTemperature {
	rv := objc.Call[UnitTemperature](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitTemperature() UnitTemperature {
	return UnitTemperatureClass.New()
}

func (u_ UnitTemperature) Init() UnitTemperature {
	rv := objc.Call[UnitTemperature](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitTemperatureClass) BaseUnit() UnitTemperature {
	rv := objc.Call[UnitTemperature](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitTemperature_BaseUnit() UnitTemperature {
	return UnitTemperatureClass.BaseUnit()
}

func (u_ UnitTemperature) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitTemperature {
	rv := objc.Call[UnitTemperature](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func UnitTemperature_InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitTemperature {
	return UnitTemperatureClass.Alloc().InitWithSymbolConverter(symbol, converter)
}

func (u_ UnitTemperature) InitWithSymbol(symbol string) UnitTemperature {
	rv := objc.Call[UnitTemperature](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func UnitTemperature_InitWithSymbol(symbol string) UnitTemperature {
	return UnitTemperatureClass.Alloc().InitWithSymbol(symbol)
}

// The kelvin unit of temperature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunittemperature/1690766-kelvin?language=objc
func (uc _UnitTemperatureClass) Kelvin() UnitTemperature {
	rv := objc.Call[UnitTemperature](uc, objc.Sel("kelvin"))
	return rv
}

// The kelvin unit of temperature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunittemperature/1690766-kelvin?language=objc
func UnitTemperature_Kelvin() UnitTemperature {
	return UnitTemperatureClass.Kelvin()
}

// The degree Celsius unit of temperature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunittemperature/1690835-celsius?language=objc
func (uc _UnitTemperatureClass) Celsius() UnitTemperature {
	rv := objc.Call[UnitTemperature](uc, objc.Sel("celsius"))
	return rv
}

// The degree Celsius unit of temperature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunittemperature/1690835-celsius?language=objc
func UnitTemperature_Celsius() UnitTemperature {
	return UnitTemperatureClass.Celsius()
}

// The degree Fahrenheit unit of temperature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunittemperature/1690842-fahrenheit?language=objc
func (uc _UnitTemperatureClass) Fahrenheit() UnitTemperature {
	rv := objc.Call[UnitTemperature](uc, objc.Sel("fahrenheit"))
	return rv
}

// The degree Fahrenheit unit of temperature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunittemperature/1690842-fahrenheit?language=objc
func UnitTemperature_Fahrenheit() UnitTemperature {
	return UnitTemperatureClass.Fahrenheit()
}
