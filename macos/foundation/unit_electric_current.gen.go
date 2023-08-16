// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitElectricCurrent] class.
var UnitElectricCurrentClass = _UnitElectricCurrentClass{objc.GetClass("NSUnitElectricCurrent")}

type _UnitElectricCurrentClass struct {
	objc.Class
}

// An interface definition for the [UnitElectricCurrent] class.
type IUnitElectricCurrent interface {
	IDimension
}

// A unit of measure for electric current. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccurrent?language=objc
type UnitElectricCurrent struct {
	Dimension
}

func UnitElectricCurrentFrom(ptr unsafe.Pointer) UnitElectricCurrent {
	return UnitElectricCurrent{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitElectricCurrentClass) Alloc() UnitElectricCurrent {
	rv := objc.Call[UnitElectricCurrent](uc, objc.Sel("alloc"))
	return rv
}

func UnitElectricCurrent_Alloc() UnitElectricCurrent {
	return UnitElectricCurrentClass.Alloc()
}

func (uc _UnitElectricCurrentClass) New() UnitElectricCurrent {
	rv := objc.Call[UnitElectricCurrent](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitElectricCurrent() UnitElectricCurrent {
	return UnitElectricCurrentClass.New()
}

func (u_ UnitElectricCurrent) Init() UnitElectricCurrent {
	rv := objc.Call[UnitElectricCurrent](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitElectricCurrentClass) BaseUnit() UnitElectricCurrent {
	rv := objc.Call[UnitElectricCurrent](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitElectricCurrent_BaseUnit() UnitElectricCurrent {
	return UnitElectricCurrentClass.BaseUnit()
}

func (u_ UnitElectricCurrent) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitElectricCurrent {
	rv := objc.Call[UnitElectricCurrent](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func UnitElectricCurrent_InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitElectricCurrent {
	return UnitElectricCurrentClass.Alloc().InitWithSymbolConverter(symbol, converter)
}

func (u_ UnitElectricCurrent) InitWithSymbol(symbol string) UnitElectricCurrent {
	rv := objc.Call[UnitElectricCurrent](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func UnitElectricCurrent_InitWithSymbol(symbol string) UnitElectricCurrent {
	return UnitElectricCurrentClass.Alloc().InitWithSymbol(symbol)
}

// The microamperes unit of electric current. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccurrent/1856047-microamperes?language=objc
func (uc _UnitElectricCurrentClass) Microamperes() UnitElectricCurrent {
	rv := objc.Call[UnitElectricCurrent](uc, objc.Sel("microamperes"))
	return rv
}

// The microamperes unit of electric current. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccurrent/1856047-microamperes?language=objc
func UnitElectricCurrent_Microamperes() UnitElectricCurrent {
	return UnitElectricCurrentClass.Microamperes()
}

// The amperes unit of electric current. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccurrent/1855973-amperes?language=objc
func (uc _UnitElectricCurrentClass) Amperes() UnitElectricCurrent {
	rv := objc.Call[UnitElectricCurrent](uc, objc.Sel("amperes"))
	return rv
}

// The amperes unit of electric current. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccurrent/1855973-amperes?language=objc
func UnitElectricCurrent_Amperes() UnitElectricCurrent {
	return UnitElectricCurrentClass.Amperes()
}

// The kiloamperes unit of electric current. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccurrent/1856001-kiloamperes?language=objc
func (uc _UnitElectricCurrentClass) Kiloamperes() UnitElectricCurrent {
	rv := objc.Call[UnitElectricCurrent](uc, objc.Sel("kiloamperes"))
	return rv
}

// The kiloamperes unit of electric current. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccurrent/1856001-kiloamperes?language=objc
func UnitElectricCurrent_Kiloamperes() UnitElectricCurrent {
	return UnitElectricCurrentClass.Kiloamperes()
}

// The megaamperes unit of electric current. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccurrent/1855980-megaamperes?language=objc
func (uc _UnitElectricCurrentClass) Megaamperes() UnitElectricCurrent {
	rv := objc.Call[UnitElectricCurrent](uc, objc.Sel("megaamperes"))
	return rv
}

// The megaamperes unit of electric current. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccurrent/1855980-megaamperes?language=objc
func UnitElectricCurrent_Megaamperes() UnitElectricCurrent {
	return UnitElectricCurrentClass.Megaamperes()
}

// The milliamperes unit of electric current. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccurrent/1856058-milliamperes?language=objc
func (uc _UnitElectricCurrentClass) Milliamperes() UnitElectricCurrent {
	rv := objc.Call[UnitElectricCurrent](uc, objc.Sel("milliamperes"))
	return rv
}

// The milliamperes unit of electric current. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitelectriccurrent/1856058-milliamperes?language=objc
func UnitElectricCurrent_Milliamperes() UnitElectricCurrent {
	return UnitElectricCurrentClass.Milliamperes()
}
