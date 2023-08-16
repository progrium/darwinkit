// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitElectricResistance] class.
var UnitElectricResistanceClass = _UnitElectricResistanceClass{objc.GetClass("NSUnitElectricResistance")}

type _UnitElectricResistanceClass struct {
	objc.Class
}

// An interface definition for the [UnitElectricResistance] class.
type IUnitElectricResistance interface {
	IDimension
}

// A unit of measure for electric resistance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricresistance?language=objc
type UnitElectricResistance struct {
	Dimension
}

func UnitElectricResistanceFrom(ptr unsafe.Pointer) UnitElectricResistance {
	return UnitElectricResistance{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitElectricResistanceClass) Alloc() UnitElectricResistance {
	rv := objc.Call[UnitElectricResistance](uc, objc.Sel("alloc"))
	return rv
}

func UnitElectricResistance_Alloc() UnitElectricResistance {
	return UnitElectricResistanceClass.Alloc()
}

func (uc _UnitElectricResistanceClass) New() UnitElectricResistance {
	rv := objc.Call[UnitElectricResistance](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitElectricResistance() UnitElectricResistance {
	return UnitElectricResistanceClass.New()
}

func (u_ UnitElectricResistance) Init() UnitElectricResistance {
	rv := objc.Call[UnitElectricResistance](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitElectricResistanceClass) BaseUnit() UnitElectricResistance {
	rv := objc.Call[UnitElectricResistance](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitElectricResistance_BaseUnit() UnitElectricResistance {
	return UnitElectricResistanceClass.BaseUnit()
}

func (u_ UnitElectricResistance) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitElectricResistance {
	rv := objc.Call[UnitElectricResistance](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func UnitElectricResistance_InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitElectricResistance {
	return UnitElectricResistanceClass.Alloc().InitWithSymbolConverter(symbol, converter)
}

func (u_ UnitElectricResistance) InitWithSymbol(symbol string) UnitElectricResistance {
	rv := objc.Call[UnitElectricResistance](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func UnitElectricResistance_InitWithSymbol(symbol string) UnitElectricResistance {
	return UnitElectricResistanceClass.Alloc().InitWithSymbol(symbol)
}

// The ohms unit of electric resistance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricresistance/1856110-ohms?language=objc
func (uc _UnitElectricResistanceClass) Ohms() UnitElectricResistance {
	rv := objc.Call[UnitElectricResistance](uc, objc.Sel("ohms"))
	return rv
}

// The ohms unit of electric resistance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricresistance/1856110-ohms?language=objc
func UnitElectricResistance_Ohms() UnitElectricResistance {
	return UnitElectricResistanceClass.Ohms()
}

// The milliohms unit of electric resistance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricresistance/1856049-milliohms?language=objc
func (uc _UnitElectricResistanceClass) Milliohms() UnitElectricResistance {
	rv := objc.Call[UnitElectricResistance](uc, objc.Sel("milliohms"))
	return rv
}

// The milliohms unit of electric resistance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricresistance/1856049-milliohms?language=objc
func UnitElectricResistance_Milliohms() UnitElectricResistance {
	return UnitElectricResistanceClass.Milliohms()
}

// The microohms unit of electric resistance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricresistance/1856031-microohms?language=objc
func (uc _UnitElectricResistanceClass) Microohms() UnitElectricResistance {
	rv := objc.Call[UnitElectricResistance](uc, objc.Sel("microohms"))
	return rv
}

// The microohms unit of electric resistance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricresistance/1856031-microohms?language=objc
func UnitElectricResistance_Microohms() UnitElectricResistance {
	return UnitElectricResistanceClass.Microohms()
}

// The kiloohms unit of electric resistance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricresistance/1855981-kiloohms?language=objc
func (uc _UnitElectricResistanceClass) Kiloohms() UnitElectricResistance {
	rv := objc.Call[UnitElectricResistance](uc, objc.Sel("kiloohms"))
	return rv
}

// The kiloohms unit of electric resistance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricresistance/1855981-kiloohms?language=objc
func UnitElectricResistance_Kiloohms() UnitElectricResistance {
	return UnitElectricResistanceClass.Kiloohms()
}

// The megaohms unit of electric resistance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricresistance/1856009-megaohms?language=objc
func (uc _UnitElectricResistanceClass) Megaohms() UnitElectricResistance {
	rv := objc.Call[UnitElectricResistance](uc, objc.Sel("megaohms"))
	return rv
}

// The megaohms unit of electric resistance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectricresistance/1856009-megaohms?language=objc
func UnitElectricResistance_Megaohms() UnitElectricResistance {
	return UnitElectricResistanceClass.Megaohms()
}
