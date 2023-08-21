// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitIlluminance] class.
var UnitIlluminanceClass = _UnitIlluminanceClass{objc.GetClass("NSUnitIlluminance")}

type _UnitIlluminanceClass struct {
	objc.Class
}

// An interface definition for the [UnitIlluminance] class.
type IUnitIlluminance interface {
	IDimension
}

// A unit of measure for illuminance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitilluminance?language=objc
type UnitIlluminance struct {
	Dimension
}

func UnitIlluminanceFrom(ptr unsafe.Pointer) UnitIlluminance {
	return UnitIlluminance{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitIlluminanceClass) Alloc() UnitIlluminance {
	rv := objc.Call[UnitIlluminance](uc, objc.Sel("alloc"))
	return rv
}

func UnitIlluminance_Alloc() UnitIlluminance {
	return UnitIlluminanceClass.Alloc()
}

func (uc _UnitIlluminanceClass) New() UnitIlluminance {
	rv := objc.Call[UnitIlluminance](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitIlluminance() UnitIlluminance {
	return UnitIlluminanceClass.New()
}

func (u_ UnitIlluminance) Init() UnitIlluminance {
	rv := objc.Call[UnitIlluminance](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitIlluminanceClass) BaseUnit() UnitIlluminance {
	rv := objc.Call[UnitIlluminance](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitIlluminance_BaseUnit() UnitIlluminance {
	return UnitIlluminanceClass.BaseUnit()
}

func (u_ UnitIlluminance) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitIlluminance {
	rv := objc.Call[UnitIlluminance](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func NewUnitIlluminanceWithSymbolConverter(symbol string, converter IUnitConverter) UnitIlluminance {
	instance := UnitIlluminanceClass.Alloc().InitWithSymbolConverter(symbol, converter)
	instance.Autorelease()
	return instance
}

func (u_ UnitIlluminance) InitWithSymbol(symbol string) UnitIlluminance {
	rv := objc.Call[UnitIlluminance](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func NewUnitIlluminanceWithSymbol(symbol string) UnitIlluminance {
	instance := UnitIlluminanceClass.Alloc().InitWithSymbol(symbol)
	instance.Autorelease()
	return instance
}

// The lux unit of illuminance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitilluminance/1823716-lux?language=objc
func (uc _UnitIlluminanceClass) Lux() UnitIlluminance {
	rv := objc.Call[UnitIlluminance](uc, objc.Sel("lux"))
	return rv
}

// The lux unit of illuminance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitilluminance/1823716-lux?language=objc
func UnitIlluminance_Lux() UnitIlluminance {
	return UnitIlluminanceClass.Lux()
}
