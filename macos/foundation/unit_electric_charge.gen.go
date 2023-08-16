// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitElectricCharge] class.
var UnitElectricChargeClass = _UnitElectricChargeClass{objc.GetClass("NSUnitElectricCharge")}

type _UnitElectricChargeClass struct {
	objc.Class
}

// An interface definition for the [UnitElectricCharge] class.
type IUnitElectricCharge interface {
	IDimension
}

// A unit of measure for electric charge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccharge?language=objc
type UnitElectricCharge struct {
	Dimension
}

func UnitElectricChargeFrom(ptr unsafe.Pointer) UnitElectricCharge {
	return UnitElectricCharge{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitElectricChargeClass) Alloc() UnitElectricCharge {
	rv := objc.Call[UnitElectricCharge](uc, objc.Sel("alloc"))
	return rv
}

func UnitElectricCharge_Alloc() UnitElectricCharge {
	return UnitElectricChargeClass.Alloc()
}

func (uc _UnitElectricChargeClass) New() UnitElectricCharge {
	rv := objc.Call[UnitElectricCharge](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitElectricCharge() UnitElectricCharge {
	return UnitElectricChargeClass.New()
}

func (u_ UnitElectricCharge) Init() UnitElectricCharge {
	rv := objc.Call[UnitElectricCharge](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitElectricChargeClass) BaseUnit() UnitElectricCharge {
	rv := objc.Call[UnitElectricCharge](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitElectricCharge_BaseUnit() UnitElectricCharge {
	return UnitElectricChargeClass.BaseUnit()
}

func (u_ UnitElectricCharge) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitElectricCharge {
	rv := objc.Call[UnitElectricCharge](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func UnitElectricCharge_InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitElectricCharge {
	return UnitElectricChargeClass.Alloc().InitWithSymbolConverter(symbol, converter)
}

func (u_ UnitElectricCharge) InitWithSymbol(symbol string) UnitElectricCharge {
	rv := objc.Call[UnitElectricCharge](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func UnitElectricCharge_InitWithSymbol(symbol string) UnitElectricCharge {
	return UnitElectricChargeClass.Alloc().InitWithSymbol(symbol)
}

// The kiloampere hours unit of electric charge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccharge/1856045-kiloamperehours?language=objc
func (uc _UnitElectricChargeClass) KiloampereHours() UnitElectricCharge {
	rv := objc.Call[UnitElectricCharge](uc, objc.Sel("kiloampereHours"))
	return rv
}

// The kiloampere hours unit of electric charge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccharge/1856045-kiloamperehours?language=objc
func UnitElectricCharge_KiloampereHours() UnitElectricCharge {
	return UnitElectricChargeClass.KiloampereHours()
}

// The milliampere hours unit of electric charge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccharge/1856102-milliamperehours?language=objc
func (uc _UnitElectricChargeClass) MilliampereHours() UnitElectricCharge {
	rv := objc.Call[UnitElectricCharge](uc, objc.Sel("milliampereHours"))
	return rv
}

// The milliampere hours unit of electric charge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccharge/1856102-milliamperehours?language=objc
func UnitElectricCharge_MilliampereHours() UnitElectricCharge {
	return UnitElectricChargeClass.MilliampereHours()
}

// The coulombs unit of electric charge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccharge/1856032-coulombs?language=objc
func (uc _UnitElectricChargeClass) Coulombs() UnitElectricCharge {
	rv := objc.Call[UnitElectricCharge](uc, objc.Sel("coulombs"))
	return rv
}

// The coulombs unit of electric charge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccharge/1856032-coulombs?language=objc
func UnitElectricCharge_Coulombs() UnitElectricCharge {
	return UnitElectricChargeClass.Coulombs()
}

// The ampere hours unit of electric charge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccharge/1856117-amperehours?language=objc
func (uc _UnitElectricChargeClass) AmpereHours() UnitElectricCharge {
	rv := objc.Call[UnitElectricCharge](uc, objc.Sel("ampereHours"))
	return rv
}

// The ampere hours unit of electric charge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccharge/1856117-amperehours?language=objc
func UnitElectricCharge_AmpereHours() UnitElectricCharge {
	return UnitElectricChargeClass.AmpereHours()
}

// The microampere hours unit of electric charge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccharge/1856006-microamperehours?language=objc
func (uc _UnitElectricChargeClass) MicroampereHours() UnitElectricCharge {
	rv := objc.Call[UnitElectricCharge](uc, objc.Sel("microampereHours"))
	return rv
}

// The microampere hours unit of electric charge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccharge/1856006-microamperehours?language=objc
func UnitElectricCharge_MicroampereHours() UnitElectricCharge {
	return UnitElectricChargeClass.MicroampereHours()
}

// The megaampere hours unit of electric charge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccharge/1856020-megaamperehours?language=objc
func (uc _UnitElectricChargeClass) MegaampereHours() UnitElectricCharge {
	rv := objc.Call[UnitElectricCharge](uc, objc.Sel("megaampereHours"))
	return rv
}

// The megaampere hours unit of electric charge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccharge/1856020-megaamperehours?language=objc
func UnitElectricCharge_MegaampereHours() UnitElectricCharge {
	return UnitElectricChargeClass.MegaampereHours()
}
