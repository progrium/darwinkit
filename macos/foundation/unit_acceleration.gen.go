// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitAcceleration] class.
var UnitAccelerationClass = _UnitAccelerationClass{objc.GetClass("NSUnitAcceleration")}

type _UnitAccelerationClass struct {
	objc.Class
}

// An interface definition for the [UnitAcceleration] class.
type IUnitAcceleration interface {
	IDimension
}

// A unit of measure for acceleration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitacceleration?language=objc
type UnitAcceleration struct {
	Dimension
}

func UnitAccelerationFrom(ptr unsafe.Pointer) UnitAcceleration {
	return UnitAcceleration{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitAccelerationClass) Alloc() UnitAcceleration {
	rv := objc.Call[UnitAcceleration](uc, objc.Sel("alloc"))
	return rv
}

func UnitAcceleration_Alloc() UnitAcceleration {
	return UnitAccelerationClass.Alloc()
}

func (uc _UnitAccelerationClass) New() UnitAcceleration {
	rv := objc.Call[UnitAcceleration](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitAcceleration() UnitAcceleration {
	return UnitAccelerationClass.New()
}

func (u_ UnitAcceleration) Init() UnitAcceleration {
	rv := objc.Call[UnitAcceleration](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitAccelerationClass) BaseUnit() UnitAcceleration {
	rv := objc.Call[UnitAcceleration](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitAcceleration_BaseUnit() UnitAcceleration {
	return UnitAccelerationClass.BaseUnit()
}

func (u_ UnitAcceleration) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitAcceleration {
	rv := objc.Call[UnitAcceleration](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func UnitAcceleration_InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitAcceleration {
	return UnitAccelerationClass.Alloc().InitWithSymbolConverter(symbol, converter)
}

func (u_ UnitAcceleration) InitWithSymbol(symbol string) UnitAcceleration {
	rv := objc.Call[UnitAcceleration](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func UnitAcceleration_InitWithSymbol(symbol string) UnitAcceleration {
	return UnitAccelerationClass.Alloc().InitWithSymbol(symbol)
}

// Returns the meter per second squared unit of acceleration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitacceleration/1856015-meterspersecondsquared?language=objc
func (uc _UnitAccelerationClass) MetersPerSecondSquared() UnitAcceleration {
	rv := objc.Call[UnitAcceleration](uc, objc.Sel("metersPerSecondSquared"))
	return rv
}

// Returns the meter per second squared unit of acceleration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitacceleration/1856015-meterspersecondsquared?language=objc
func UnitAcceleration_MetersPerSecondSquared() UnitAcceleration {
	return UnitAccelerationClass.MetersPerSecondSquared()
}

// Returns the gravity unit of acceleration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitacceleration/1690681-gravity?language=objc
func (uc _UnitAccelerationClass) Gravity() UnitAcceleration {
	rv := objc.Call[UnitAcceleration](uc, objc.Sel("gravity"))
	return rv
}

// Returns the gravity unit of acceleration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitacceleration/1690681-gravity?language=objc
func UnitAcceleration_Gravity() UnitAcceleration {
	return UnitAccelerationClass.Gravity()
}
