// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitFuelEfficiency] class.
var UnitFuelEfficiencyClass = _UnitFuelEfficiencyClass{objc.GetClass("NSUnitFuelEfficiency")}

type _UnitFuelEfficiencyClass struct {
	objc.Class
}

// An interface definition for the [UnitFuelEfficiency] class.
type IUnitFuelEfficiency interface {
	IDimension
}

// A unit of measure for fuel efficiency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfuelefficiency?language=objc
type UnitFuelEfficiency struct {
	Dimension
}

func UnitFuelEfficiencyFrom(ptr unsafe.Pointer) UnitFuelEfficiency {
	return UnitFuelEfficiency{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitFuelEfficiencyClass) Alloc() UnitFuelEfficiency {
	rv := objc.Call[UnitFuelEfficiency](uc, objc.Sel("alloc"))
	return rv
}

func UnitFuelEfficiency_Alloc() UnitFuelEfficiency {
	return UnitFuelEfficiencyClass.Alloc()
}

func (uc _UnitFuelEfficiencyClass) New() UnitFuelEfficiency {
	rv := objc.Call[UnitFuelEfficiency](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitFuelEfficiency() UnitFuelEfficiency {
	return UnitFuelEfficiencyClass.New()
}

func (u_ UnitFuelEfficiency) Init() UnitFuelEfficiency {
	rv := objc.Call[UnitFuelEfficiency](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitFuelEfficiencyClass) BaseUnit() UnitFuelEfficiency {
	rv := objc.Call[UnitFuelEfficiency](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitFuelEfficiency_BaseUnit() UnitFuelEfficiency {
	return UnitFuelEfficiencyClass.BaseUnit()
}

func (u_ UnitFuelEfficiency) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitFuelEfficiency {
	rv := objc.Call[UnitFuelEfficiency](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func NewUnitFuelEfficiencyWithSymbolConverter(symbol string, converter IUnitConverter) UnitFuelEfficiency {
	instance := UnitFuelEfficiencyClass.Alloc().InitWithSymbolConverter(symbol, converter)
	instance.Autorelease()
	return instance
}

func (u_ UnitFuelEfficiency) InitWithSymbol(symbol string) UnitFuelEfficiency {
	rv := objc.Call[UnitFuelEfficiency](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func NewUnitFuelEfficiencyWithSymbol(symbol string) UnitFuelEfficiency {
	instance := UnitFuelEfficiencyClass.Alloc().InitWithSymbol(symbol)
	instance.Autorelease()
	return instance
}

// The liters per 100 kilometers unit of fuel efficiency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfuelefficiency/1856054-litersper100kilometers?language=objc
func (uc _UnitFuelEfficiencyClass) LitersPer100Kilometers() UnitFuelEfficiency {
	rv := objc.Call[UnitFuelEfficiency](uc, objc.Sel("litersPer100Kilometers"))
	return rv
}

// The liters per 100 kilometers unit of fuel efficiency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfuelefficiency/1856054-litersper100kilometers?language=objc
func UnitFuelEfficiency_LitersPer100Kilometers() UnitFuelEfficiency {
	return UnitFuelEfficiencyClass.LitersPer100Kilometers()
}

// The miles per imperial gallon unit of fuel efficiency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfuelefficiency/1856089-milesperimperialgallon?language=objc
func (uc _UnitFuelEfficiencyClass) MilesPerImperialGallon() UnitFuelEfficiency {
	rv := objc.Call[UnitFuelEfficiency](uc, objc.Sel("milesPerImperialGallon"))
	return rv
}

// The miles per imperial gallon unit of fuel efficiency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfuelefficiency/1856089-milesperimperialgallon?language=objc
func UnitFuelEfficiency_MilesPerImperialGallon() UnitFuelEfficiency {
	return UnitFuelEfficiencyClass.MilesPerImperialGallon()
}

// The miles per gallon unit of fuel efficiency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfuelefficiency/1856085-milespergallon?language=objc
func (uc _UnitFuelEfficiencyClass) MilesPerGallon() UnitFuelEfficiency {
	rv := objc.Call[UnitFuelEfficiency](uc, objc.Sel("milesPerGallon"))
	return rv
}

// The miles per gallon unit of fuel efficiency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfuelefficiency/1856085-milespergallon?language=objc
func UnitFuelEfficiency_MilesPerGallon() UnitFuelEfficiency {
	return UnitFuelEfficiencyClass.MilesPerGallon()
}
