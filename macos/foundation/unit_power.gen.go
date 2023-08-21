// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitPower] class.
var UnitPowerClass = _UnitPowerClass{objc.GetClass("NSUnitPower")}

type _UnitPowerClass struct {
	objc.Class
}

// An interface definition for the [UnitPower] class.
type IUnitPower interface {
	IDimension
}

// A unit of measure for power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower?language=objc
type UnitPower struct {
	Dimension
}

func UnitPowerFrom(ptr unsafe.Pointer) UnitPower {
	return UnitPower{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitPowerClass) Alloc() UnitPower {
	rv := objc.Call[UnitPower](uc, objc.Sel("alloc"))
	return rv
}

func UnitPower_Alloc() UnitPower {
	return UnitPowerClass.Alloc()
}

func (uc _UnitPowerClass) New() UnitPower {
	rv := objc.Call[UnitPower](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitPower() UnitPower {
	return UnitPowerClass.New()
}

func (u_ UnitPower) Init() UnitPower {
	rv := objc.Call[UnitPower](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitPowerClass) BaseUnit() UnitPower {
	rv := objc.Call[UnitPower](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitPower_BaseUnit() UnitPower {
	return UnitPowerClass.BaseUnit()
}

func (u_ UnitPower) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitPower {
	rv := objc.Call[UnitPower](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func NewUnitPowerWithSymbolConverter(symbol string, converter IUnitConverter) UnitPower {
	instance := UnitPowerClass.Alloc().InitWithSymbolConverter(symbol, converter)
	instance.Autorelease()
	return instance
}

func (u_ UnitPower) InitWithSymbol(symbol string) UnitPower {
	rv := objc.Call[UnitPower](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func NewUnitPowerWithSymbol(symbol string) UnitPower {
	instance := UnitPowerClass.Alloc().InitWithSymbol(symbol)
	instance.Autorelease()
	return instance
}

// The femtowatts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1856043-femtowatts?language=objc
func (uc _UnitPowerClass) Femtowatts() UnitPower {
	rv := objc.Call[UnitPower](uc, objc.Sel("femtowatts"))
	return rv
}

// The femtowatts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1856043-femtowatts?language=objc
func UnitPower_Femtowatts() UnitPower {
	return UnitPowerClass.Femtowatts()
}

// The picowatts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1856104-picowatts?language=objc
func (uc _UnitPowerClass) Picowatts() UnitPower {
	rv := objc.Call[UnitPower](uc, objc.Sel("picowatts"))
	return rv
}

// The picowatts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1856104-picowatts?language=objc
func UnitPower_Picowatts() UnitPower {
	return UnitPowerClass.Picowatts()
}

// The terawatts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1856100-terawatts?language=objc
func (uc _UnitPowerClass) Terawatts() UnitPower {
	rv := objc.Call[UnitPower](uc, objc.Sel("terawatts"))
	return rv
}

// The terawatts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1856100-terawatts?language=objc
func UnitPower_Terawatts() UnitPower {
	return UnitPowerClass.Terawatts()
}

// The milliwatts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1855984-milliwatts?language=objc
func (uc _UnitPowerClass) Milliwatts() UnitPower {
	rv := objc.Call[UnitPower](uc, objc.Sel("milliwatts"))
	return rv
}

// The milliwatts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1855984-milliwatts?language=objc
func UnitPower_Milliwatts() UnitPower {
	return UnitPowerClass.Milliwatts()
}

// The watts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1856075-watts?language=objc
func (uc _UnitPowerClass) Watts() UnitPower {
	rv := objc.Call[UnitPower](uc, objc.Sel("watts"))
	return rv
}

// The watts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1856075-watts?language=objc
func UnitPower_Watts() UnitPower {
	return UnitPowerClass.Watts()
}

// The nanowatts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1855999-nanowatts?language=objc
func (uc _UnitPowerClass) Nanowatts() UnitPower {
	rv := objc.Call[UnitPower](uc, objc.Sel("nanowatts"))
	return rv
}

// The nanowatts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1855999-nanowatts?language=objc
func UnitPower_Nanowatts() UnitPower {
	return UnitPowerClass.Nanowatts()
}

// The kilowatts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1856084-kilowatts?language=objc
func (uc _UnitPowerClass) Kilowatts() UnitPower {
	rv := objc.Call[UnitPower](uc, objc.Sel("kilowatts"))
	return rv
}

// The kilowatts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1856084-kilowatts?language=objc
func UnitPower_Kilowatts() UnitPower {
	return UnitPowerClass.Kilowatts()
}

// The megawatts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1856073-megawatts?language=objc
func (uc _UnitPowerClass) Megawatts() UnitPower {
	rv := objc.Call[UnitPower](uc, objc.Sel("megawatts"))
	return rv
}

// The megawatts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1856073-megawatts?language=objc
func UnitPower_Megawatts() UnitPower {
	return UnitPowerClass.Megawatts()
}

// The horsepower unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1690871-horsepower?language=objc
func (uc _UnitPowerClass) Horsepower() UnitPower {
	rv := objc.Call[UnitPower](uc, objc.Sel("horsepower"))
	return rv
}

// The horsepower unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1690871-horsepower?language=objc
func UnitPower_Horsepower() UnitPower {
	return UnitPowerClass.Horsepower()
}

// The microwatts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1856051-microwatts?language=objc
func (uc _UnitPowerClass) Microwatts() UnitPower {
	rv := objc.Call[UnitPower](uc, objc.Sel("microwatts"))
	return rv
}

// The microwatts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1856051-microwatts?language=objc
func UnitPower_Microwatts() UnitPower {
	return UnitPowerClass.Microwatts()
}

// The gigawatts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1856108-gigawatts?language=objc
func (uc _UnitPowerClass) Gigawatts() UnitPower {
	rv := objc.Call[UnitPower](uc, objc.Sel("gigawatts"))
	return rv
}

// The gigawatts unit of power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpower/1856108-gigawatts?language=objc
func UnitPower_Gigawatts() UnitPower {
	return UnitPowerClass.Gigawatts()
}
