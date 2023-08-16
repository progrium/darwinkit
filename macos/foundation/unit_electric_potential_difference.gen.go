// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitElectricPotentialDifference] class.
var UnitElectricPotentialDifferenceClass = _UnitElectricPotentialDifferenceClass{objc.GetClass("NSUnitElectricPotentialDifference")}

type _UnitElectricPotentialDifferenceClass struct {
	objc.Class
}

// An interface definition for the [UnitElectricPotentialDifference] class.
type IUnitElectricPotentialDifference interface {
	IDimension
}

// A unit of measure for electric potential difference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricpotentialdifference?language=objc
type UnitElectricPotentialDifference struct {
	Dimension
}

func UnitElectricPotentialDifferenceFrom(ptr unsafe.Pointer) UnitElectricPotentialDifference {
	return UnitElectricPotentialDifference{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitElectricPotentialDifferenceClass) Alloc() UnitElectricPotentialDifference {
	rv := objc.Call[UnitElectricPotentialDifference](uc, objc.Sel("alloc"))
	return rv
}

func UnitElectricPotentialDifference_Alloc() UnitElectricPotentialDifference {
	return UnitElectricPotentialDifferenceClass.Alloc()
}

func (uc _UnitElectricPotentialDifferenceClass) New() UnitElectricPotentialDifference {
	rv := objc.Call[UnitElectricPotentialDifference](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitElectricPotentialDifference() UnitElectricPotentialDifference {
	return UnitElectricPotentialDifferenceClass.New()
}

func (u_ UnitElectricPotentialDifference) Init() UnitElectricPotentialDifference {
	rv := objc.Call[UnitElectricPotentialDifference](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitElectricPotentialDifferenceClass) BaseUnit() UnitElectricPotentialDifference {
	rv := objc.Call[UnitElectricPotentialDifference](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitElectricPotentialDifference_BaseUnit() UnitElectricPotentialDifference {
	return UnitElectricPotentialDifferenceClass.BaseUnit()
}

func (u_ UnitElectricPotentialDifference) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitElectricPotentialDifference {
	rv := objc.Call[UnitElectricPotentialDifference](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func UnitElectricPotentialDifference_InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitElectricPotentialDifference {
	return UnitElectricPotentialDifferenceClass.Alloc().InitWithSymbolConverter(symbol, converter)
}

func (u_ UnitElectricPotentialDifference) InitWithSymbol(symbol string) UnitElectricPotentialDifference {
	rv := objc.Call[UnitElectricPotentialDifference](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func UnitElectricPotentialDifference_InitWithSymbol(symbol string) UnitElectricPotentialDifference {
	return UnitElectricPotentialDifferenceClass.Alloc().InitWithSymbol(symbol)
}

// The volts unit of electric potential difference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricpotentialdifference/1856095-volts?language=objc
func (uc _UnitElectricPotentialDifferenceClass) Volts() UnitElectricPotentialDifference {
	rv := objc.Call[UnitElectricPotentialDifference](uc, objc.Sel("volts"))
	return rv
}

// The volts unit of electric potential difference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricpotentialdifference/1856095-volts?language=objc
func UnitElectricPotentialDifference_Volts() UnitElectricPotentialDifference {
	return UnitElectricPotentialDifferenceClass.Volts()
}

// The millivolts unit of electric potential difference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricpotentialdifference/1856094-millivolts?language=objc
func (uc _UnitElectricPotentialDifferenceClass) Millivolts() UnitElectricPotentialDifference {
	rv := objc.Call[UnitElectricPotentialDifference](uc, objc.Sel("millivolts"))
	return rv
}

// The millivolts unit of electric potential difference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricpotentialdifference/1856094-millivolts?language=objc
func UnitElectricPotentialDifference_Millivolts() UnitElectricPotentialDifference {
	return UnitElectricPotentialDifferenceClass.Millivolts()
}

// The kilovolts unit of electric potential difference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricpotentialdifference/1856086-kilovolts?language=objc
func (uc _UnitElectricPotentialDifferenceClass) Kilovolts() UnitElectricPotentialDifference {
	rv := objc.Call[UnitElectricPotentialDifference](uc, objc.Sel("kilovolts"))
	return rv
}

// The kilovolts unit of electric potential difference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricpotentialdifference/1856086-kilovolts?language=objc
func UnitElectricPotentialDifference_Kilovolts() UnitElectricPotentialDifference {
	return UnitElectricPotentialDifferenceClass.Kilovolts()
}

// The megavolts unit of electric potential difference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricpotentialdifference/1855975-megavolts?language=objc
func (uc _UnitElectricPotentialDifferenceClass) Megavolts() UnitElectricPotentialDifference {
	rv := objc.Call[UnitElectricPotentialDifference](uc, objc.Sel("megavolts"))
	return rv
}

// The megavolts unit of electric potential difference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricpotentialdifference/1855975-megavolts?language=objc
func UnitElectricPotentialDifference_Megavolts() UnitElectricPotentialDifference {
	return UnitElectricPotentialDifferenceClass.Megavolts()
}

// The microvolts unit of electric potential difference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricpotentialdifference/1856022-microvolts?language=objc
func (uc _UnitElectricPotentialDifferenceClass) Microvolts() UnitElectricPotentialDifference {
	rv := objc.Call[UnitElectricPotentialDifference](uc, objc.Sel("microvolts"))
	return rv
}

// The microvolts unit of electric potential difference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricpotentialdifference/1856022-microvolts?language=objc
func UnitElectricPotentialDifference_Microvolts() UnitElectricPotentialDifference {
	return UnitElectricPotentialDifferenceClass.Microvolts()
}
