// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitConcentrationMass] class.
var UnitConcentrationMassClass = _UnitConcentrationMassClass{objc.GetClass("NSUnitConcentrationMass")}

type _UnitConcentrationMassClass struct {
	objc.Class
}

// An interface definition for the [UnitConcentrationMass] class.
type IUnitConcentrationMass interface {
	IDimension
}

// A unit of measure for concentration of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitconcentrationmass?language=objc
type UnitConcentrationMass struct {
	Dimension
}

func UnitConcentrationMassFrom(ptr unsafe.Pointer) UnitConcentrationMass {
	return UnitConcentrationMass{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitConcentrationMassClass) Alloc() UnitConcentrationMass {
	rv := objc.Call[UnitConcentrationMass](uc, objc.Sel("alloc"))
	return rv
}

func UnitConcentrationMass_Alloc() UnitConcentrationMass {
	return UnitConcentrationMassClass.Alloc()
}

func (uc _UnitConcentrationMassClass) New() UnitConcentrationMass {
	rv := objc.Call[UnitConcentrationMass](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitConcentrationMass() UnitConcentrationMass {
	return UnitConcentrationMassClass.New()
}

func (u_ UnitConcentrationMass) Init() UnitConcentrationMass {
	rv := objc.Call[UnitConcentrationMass](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitConcentrationMassClass) BaseUnit() UnitConcentrationMass {
	rv := objc.Call[UnitConcentrationMass](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitConcentrationMass_BaseUnit() UnitConcentrationMass {
	return UnitConcentrationMassClass.BaseUnit()
}

func (u_ UnitConcentrationMass) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitConcentrationMass {
	rv := objc.Call[UnitConcentrationMass](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func UnitConcentrationMass_InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitConcentrationMass {
	return UnitConcentrationMassClass.Alloc().InitWithSymbolConverter(symbol, converter)
}

func (u_ UnitConcentrationMass) InitWithSymbol(symbol string) UnitConcentrationMass {
	rv := objc.Call[UnitConcentrationMass](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func UnitConcentrationMass_InitWithSymbol(symbol string) UnitConcentrationMass {
	return UnitConcentrationMassClass.Alloc().InitWithSymbol(symbol)
}

// Returns the millimoles per liter unit with the specified number of grams per mole. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitconcentrationmass/1855799-millimolesperliterwithgramspermo?language=objc
func (uc _UnitConcentrationMassClass) MillimolesPerLiterWithGramsPerMole(gramsPerMole float64) UnitConcentrationMass {
	rv := objc.Call[UnitConcentrationMass](uc, objc.Sel("millimolesPerLiterWithGramsPerMole:"), gramsPerMole)
	return rv
}

// Returns the millimoles per liter unit with the specified number of grams per mole. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitconcentrationmass/1855799-millimolesperliterwithgramspermo?language=objc
func UnitConcentrationMass_MillimolesPerLiterWithGramsPerMole(gramsPerMole float64) UnitConcentrationMass {
	return UnitConcentrationMassClass.MillimolesPerLiterWithGramsPerMole(gramsPerMole)
}

// The grams per liter unit of concentration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitconcentrationmass/1856019-gramsperliter?language=objc
func (uc _UnitConcentrationMassClass) GramsPerLiter() UnitConcentrationMass {
	rv := objc.Call[UnitConcentrationMass](uc, objc.Sel("gramsPerLiter"))
	return rv
}

// The grams per liter unit of concentration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitconcentrationmass/1856019-gramsperliter?language=objc
func UnitConcentrationMass_GramsPerLiter() UnitConcentrationMass {
	return UnitConcentrationMassClass.GramsPerLiter()
}

// The milligrams per deciliter unit of concentration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitconcentrationmass/1856024-milligramsperdeciliter?language=objc
func (uc _UnitConcentrationMassClass) MilligramsPerDeciliter() UnitConcentrationMass {
	rv := objc.Call[UnitConcentrationMass](uc, objc.Sel("milligramsPerDeciliter"))
	return rv
}

// The milligrams per deciliter unit of concentration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitconcentrationmass/1856024-milligramsperdeciliter?language=objc
func UnitConcentrationMass_MilligramsPerDeciliter() UnitConcentrationMass {
	return UnitConcentrationMassClass.MilligramsPerDeciliter()
}
