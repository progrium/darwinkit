// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitEnergy] class.
var UnitEnergyClass = _UnitEnergyClass{objc.GetClass("NSUnitEnergy")}

type _UnitEnergyClass struct {
	objc.Class
}

// An interface definition for the [UnitEnergy] class.
type IUnitEnergy interface {
	IDimension
}

// A unit of measure for energy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitenergy?language=objc
type UnitEnergy struct {
	Dimension
}

func UnitEnergyFrom(ptr unsafe.Pointer) UnitEnergy {
	return UnitEnergy{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitEnergyClass) Alloc() UnitEnergy {
	rv := objc.Call[UnitEnergy](uc, objc.Sel("alloc"))
	return rv
}

func UnitEnergy_Alloc() UnitEnergy {
	return UnitEnergyClass.Alloc()
}

func (uc _UnitEnergyClass) New() UnitEnergy {
	rv := objc.Call[UnitEnergy](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitEnergy() UnitEnergy {
	return UnitEnergyClass.New()
}

func (u_ UnitEnergy) Init() UnitEnergy {
	rv := objc.Call[UnitEnergy](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitEnergyClass) BaseUnit() UnitEnergy {
	rv := objc.Call[UnitEnergy](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitEnergy_BaseUnit() UnitEnergy {
	return UnitEnergyClass.BaseUnit()
}

func (u_ UnitEnergy) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitEnergy {
	rv := objc.Call[UnitEnergy](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func NewUnitEnergyWithSymbolConverter(symbol string, converter IUnitConverter) UnitEnergy {
	instance := UnitEnergyClass.Alloc().InitWithSymbolConverter(symbol, converter)
	instance.Autorelease()
	return instance
}

func (u_ UnitEnergy) InitWithSymbol(symbol string) UnitEnergy {
	rv := objc.Call[UnitEnergy](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func NewUnitEnergyWithSymbol(symbol string) UnitEnergy {
	instance := UnitEnergyClass.Alloc().InitWithSymbol(symbol)
	instance.Autorelease()
	return instance
}

// The kilowatt hours unit of energy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitenergy/1856092-kilowatthours?language=objc
func (uc _UnitEnergyClass) KilowattHours() UnitEnergy {
	rv := objc.Call[UnitEnergy](uc, objc.Sel("kilowattHours"))
	return rv
}

// The kilowatt hours unit of energy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitenergy/1856092-kilowatthours?language=objc
func UnitEnergy_KilowattHours() UnitEnergy {
	return UnitEnergyClass.KilowattHours()
}

// The calories unit of energy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitenergy/1855990-calories?language=objc
func (uc _UnitEnergyClass) Calories() UnitEnergy {
	rv := objc.Call[UnitEnergy](uc, objc.Sel("calories"))
	return rv
}

// The calories unit of energy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitenergy/1855990-calories?language=objc
func UnitEnergy_Calories() UnitEnergy {
	return UnitEnergyClass.Calories()
}

// The kilocalories unit of energy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitenergy/1856028-kilocalories?language=objc
func (uc _UnitEnergyClass) Kilocalories() UnitEnergy {
	rv := objc.Call[UnitEnergy](uc, objc.Sel("kilocalories"))
	return rv
}

// The kilocalories unit of energy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitenergy/1856028-kilocalories?language=objc
func UnitEnergy_Kilocalories() UnitEnergy {
	return UnitEnergyClass.Kilocalories()
}

// The kilojoules unit of energy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitenergy/1856113-kilojoules?language=objc
func (uc _UnitEnergyClass) Kilojoules() UnitEnergy {
	rv := objc.Call[UnitEnergy](uc, objc.Sel("kilojoules"))
	return rv
}

// The kilojoules unit of energy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitenergy/1856113-kilojoules?language=objc
func UnitEnergy_Kilojoules() UnitEnergy {
	return UnitEnergyClass.Kilojoules()
}

// The joules unit of energy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitenergy/1855987-joules?language=objc
func (uc _UnitEnergyClass) Joules() UnitEnergy {
	rv := objc.Call[UnitEnergy](uc, objc.Sel("joules"))
	return rv
}

// The joules unit of energy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitenergy/1855987-joules?language=objc
func UnitEnergy_Joules() UnitEnergy {
	return UnitEnergyClass.Joules()
}
