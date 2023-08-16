// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitSpeed] class.
var UnitSpeedClass = _UnitSpeedClass{objc.GetClass("NSUnitSpeed")}

type _UnitSpeedClass struct {
	objc.Class
}

// An interface definition for the [UnitSpeed] class.
type IUnitSpeed interface {
	IDimension
}

// A unit of measure for speed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitspeed?language=objc
type UnitSpeed struct {
	Dimension
}

func UnitSpeedFrom(ptr unsafe.Pointer) UnitSpeed {
	return UnitSpeed{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitSpeedClass) Alloc() UnitSpeed {
	rv := objc.Call[UnitSpeed](uc, objc.Sel("alloc"))
	return rv
}

func UnitSpeed_Alloc() UnitSpeed {
	return UnitSpeedClass.Alloc()
}

func (uc _UnitSpeedClass) New() UnitSpeed {
	rv := objc.Call[UnitSpeed](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitSpeed() UnitSpeed {
	return UnitSpeedClass.New()
}

func (u_ UnitSpeed) Init() UnitSpeed {
	rv := objc.Call[UnitSpeed](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitSpeedClass) BaseUnit() UnitSpeed {
	rv := objc.Call[UnitSpeed](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitSpeed_BaseUnit() UnitSpeed {
	return UnitSpeedClass.BaseUnit()
}

func (u_ UnitSpeed) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitSpeed {
	rv := objc.Call[UnitSpeed](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func UnitSpeed_InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitSpeed {
	return UnitSpeedClass.Alloc().InitWithSymbolConverter(symbol, converter)
}

func (u_ UnitSpeed) InitWithSymbol(symbol string) UnitSpeed {
	rv := objc.Call[UnitSpeed](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func UnitSpeed_InitWithSymbol(symbol string) UnitSpeed {
	return UnitSpeedClass.Alloc().InitWithSymbol(symbol)
}

// The knots unit of speed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitspeed/1856050-knots?language=objc
func (uc _UnitSpeedClass) Knots() UnitSpeed {
	rv := objc.Call[UnitSpeed](uc, objc.Sel("knots"))
	return rv
}

// The knots unit of speed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitspeed/1856050-knots?language=objc
func UnitSpeed_Knots() UnitSpeed {
	return UnitSpeedClass.Knots()
}

// The meter per second unit of speed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitspeed/1856079-meterspersecond?language=objc
func (uc _UnitSpeedClass) MetersPerSecond() UnitSpeed {
	rv := objc.Call[UnitSpeed](uc, objc.Sel("metersPerSecond"))
	return rv
}

// The meter per second unit of speed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitspeed/1856079-meterspersecond?language=objc
func UnitSpeed_MetersPerSecond() UnitSpeed {
	return UnitSpeedClass.MetersPerSecond()
}

// The kilometers per hour unit of speed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitspeed/1856044-kilometersperhour?language=objc
func (uc _UnitSpeedClass) KilometersPerHour() UnitSpeed {
	rv := objc.Call[UnitSpeed](uc, objc.Sel("kilometersPerHour"))
	return rv
}

// The kilometers per hour unit of speed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitspeed/1856044-kilometersperhour?language=objc
func UnitSpeed_KilometersPerHour() UnitSpeed {
	return UnitSpeedClass.KilometersPerHour()
}

// The miles per hour unit of speed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitspeed/1856098-milesperhour?language=objc
func (uc _UnitSpeedClass) MilesPerHour() UnitSpeed {
	rv := objc.Call[UnitSpeed](uc, objc.Sel("milesPerHour"))
	return rv
}

// The miles per hour unit of speed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitspeed/1856098-milesperhour?language=objc
func UnitSpeed_MilesPerHour() UnitSpeed {
	return UnitSpeedClass.MilesPerHour()
}
