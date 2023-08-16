// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitVolume] class.
var UnitVolumeClass = _UnitVolumeClass{objc.GetClass("NSUnitVolume")}

type _UnitVolumeClass struct {
	objc.Class
}

// An interface definition for the [UnitVolume] class.
type IUnitVolume interface {
	IDimension
}

// A unit of measure for volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume?language=objc
type UnitVolume struct {
	Dimension
}

func UnitVolumeFrom(ptr unsafe.Pointer) UnitVolume {
	return UnitVolume{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitVolumeClass) Alloc() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("alloc"))
	return rv
}

func UnitVolume_Alloc() UnitVolume {
	return UnitVolumeClass.Alloc()
}

func (uc _UnitVolumeClass) New() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitVolume() UnitVolume {
	return UnitVolumeClass.New()
}

func (u_ UnitVolume) Init() UnitVolume {
	rv := objc.Call[UnitVolume](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitVolumeClass) BaseUnit() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitVolume_BaseUnit() UnitVolume {
	return UnitVolumeClass.BaseUnit()
}

func (u_ UnitVolume) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitVolume {
	rv := objc.Call[UnitVolume](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func UnitVolume_InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitVolume {
	return UnitVolumeClass.Alloc().InitWithSymbolConverter(symbol, converter)
}

func (u_ UnitVolume) InitWithSymbol(symbol string) UnitVolume {
	rv := objc.Call[UnitVolume](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func UnitVolume_InitWithSymbol(symbol string) UnitVolume {
	return UnitVolumeClass.Alloc().InitWithSymbol(symbol)
}

// The megaliters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856080-megaliters?language=objc
func (uc _UnitVolumeClass) Megaliters() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("megaliters"))
	return rv
}

// The megaliters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856080-megaliters?language=objc
func UnitVolume_Megaliters() UnitVolume {
	return UnitVolumeClass.Megaliters()
}

// The gallons unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856105-gallons?language=objc
func (uc _UnitVolumeClass) Gallons() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("gallons"))
	return rv
}

// The gallons unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856105-gallons?language=objc
func UnitVolume_Gallons() UnitVolume {
	return UnitVolumeClass.Gallons()
}

// The imperial gallons unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1855997-imperialgallons?language=objc
func (uc _UnitVolumeClass) ImperialGallons() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("imperialGallons"))
	return rv
}

// The imperial gallons unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1855997-imperialgallons?language=objc
func UnitVolume_ImperialGallons() UnitVolume {
	return UnitVolumeClass.ImperialGallons()
}

// The cubic kilometers unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856034-cubickilometers?language=objc
func (uc _UnitVolumeClass) CubicKilometers() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("cubicKilometers"))
	return rv
}

// The cubic kilometers unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856034-cubickilometers?language=objc
func UnitVolume_CubicKilometers() UnitVolume {
	return UnitVolumeClass.CubicKilometers()
}

// The teaspoons unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1855977-teaspoons?language=objc
func (uc _UnitVolumeClass) Teaspoons() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("teaspoons"))
	return rv
}

// The teaspoons unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1855977-teaspoons?language=objc
func UnitVolume_Teaspoons() UnitVolume {
	return UnitVolumeClass.Teaspoons()
}

// The cubic centimeters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856112-cubiccentimeters?language=objc
func (uc _UnitVolumeClass) CubicCentimeters() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("cubicCentimeters"))
	return rv
}

// The cubic centimeters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856112-cubiccentimeters?language=objc
func UnitVolume_CubicCentimeters() UnitVolume {
	return UnitVolumeClass.CubicCentimeters()
}

// The cups unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1855982-cups?language=objc
func (uc _UnitVolumeClass) Cups() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("cups"))
	return rv
}

// The cups unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1855982-cups?language=objc
func UnitVolume_Cups() UnitVolume {
	return UnitVolumeClass.Cups()
}

// The cubic decimeters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856048-cubicdecimeters?language=objc
func (uc _UnitVolumeClass) CubicDecimeters() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("cubicDecimeters"))
	return rv
}

// The cubic decimeters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856048-cubicdecimeters?language=objc
func UnitVolume_CubicDecimeters() UnitVolume {
	return UnitVolumeClass.CubicDecimeters()
}

// The imperial quarts unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1855971-imperialquarts?language=objc
func (uc _UnitVolumeClass) ImperialQuarts() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("imperialQuarts"))
	return rv
}

// The imperial quarts unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1855971-imperialquarts?language=objc
func UnitVolume_ImperialQuarts() UnitVolume {
	return UnitVolumeClass.ImperialQuarts()
}

// The quarts unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856055-quarts?language=objc
func (uc _UnitVolumeClass) Quarts() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("quarts"))
	return rv
}

// The quarts unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856055-quarts?language=objc
func UnitVolume_Quarts() UnitVolume {
	return UnitVolumeClass.Quarts()
}

// The cubic inches unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856026-cubicinches?language=objc
func (uc _UnitVolumeClass) CubicInches() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("cubicInches"))
	return rv
}

// The cubic inches unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856026-cubicinches?language=objc
func UnitVolume_CubicInches() UnitVolume {
	return UnitVolumeClass.CubicInches()
}

// The cubic feet unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856017-cubicfeet?language=objc
func (uc _UnitVolumeClass) CubicFeet() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("cubicFeet"))
	return rv
}

// The cubic feet unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856017-cubicfeet?language=objc
func UnitVolume_CubicFeet() UnitVolume {
	return UnitVolumeClass.CubicFeet()
}

// The milliliters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856071-milliliters?language=objc
func (uc _UnitVolumeClass) Milliliters() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("milliliters"))
	return rv
}

// The milliliters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856071-milliliters?language=objc
func UnitVolume_Milliliters() UnitVolume {
	return UnitVolumeClass.Milliliters()
}

// The imperial pints unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1855979-imperialpints?language=objc
func (uc _UnitVolumeClass) ImperialPints() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("imperialPints"))
	return rv
}

// The imperial pints unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1855979-imperialpints?language=objc
func UnitVolume_ImperialPints() UnitVolume {
	return UnitVolumeClass.ImperialPints()
}

// The cubic miles unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856002-cubicmiles?language=objc
func (uc _UnitVolumeClass) CubicMiles() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("cubicMiles"))
	return rv
}

// The cubic miles unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856002-cubicmiles?language=objc
func UnitVolume_CubicMiles() UnitVolume {
	return UnitVolumeClass.CubicMiles()
}

// The cubic yards unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856107-cubicyards?language=objc
func (uc _UnitVolumeClass) CubicYards() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("cubicYards"))
	return rv
}

// The cubic yards unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856107-cubicyards?language=objc
func UnitVolume_CubicYards() UnitVolume {
	return UnitVolumeClass.CubicYards()
}

// The liters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856011-liters?language=objc
func (uc _UnitVolumeClass) Liters() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("liters"))
	return rv
}

// The liters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856011-liters?language=objc
func UnitVolume_Liters() UnitVolume {
	return UnitVolumeClass.Liters()
}

// The fluid ounces unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856038-fluidounces?language=objc
func (uc _UnitVolumeClass) FluidOunces() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("fluidOunces"))
	return rv
}

// The fluid ounces unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856038-fluidounces?language=objc
func UnitVolume_FluidOunces() UnitVolume {
	return UnitVolumeClass.FluidOunces()
}

// The kiloliters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856068-kiloliters?language=objc
func (uc _UnitVolumeClass) Kiloliters() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("kiloliters"))
	return rv
}

// The kiloliters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856068-kiloliters?language=objc
func UnitVolume_Kiloliters() UnitVolume {
	return UnitVolumeClass.Kiloliters()
}

// The imperial tablespoons unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1855989-imperialtablespoons?language=objc
func (uc _UnitVolumeClass) ImperialTablespoons() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("imperialTablespoons"))
	return rv
}

// The imperial tablespoons unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1855989-imperialtablespoons?language=objc
func UnitVolume_ImperialTablespoons() UnitVolume {
	return UnitVolumeClass.ImperialTablespoons()
}

// The acre feet unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1855978-acrefeet?language=objc
func (uc _UnitVolumeClass) AcreFeet() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("acreFeet"))
	return rv
}

// The acre feet unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1855978-acrefeet?language=objc
func UnitVolume_AcreFeet() UnitVolume {
	return UnitVolumeClass.AcreFeet()
}

// The tablespoons unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1855991-tablespoons?language=objc
func (uc _UnitVolumeClass) Tablespoons() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("tablespoons"))
	return rv
}

// The tablespoons unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1855991-tablespoons?language=objc
func UnitVolume_Tablespoons() UnitVolume {
	return UnitVolumeClass.Tablespoons()
}

// The imperial teaspoons unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856103-imperialteaspoons?language=objc
func (uc _UnitVolumeClass) ImperialTeaspoons() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("imperialTeaspoons"))
	return rv
}

// The imperial teaspoons unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856103-imperialteaspoons?language=objc
func UnitVolume_ImperialTeaspoons() UnitVolume {
	return UnitVolumeClass.ImperialTeaspoons()
}

// The metric cups unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856000-metriccups?language=objc
func (uc _UnitVolumeClass) MetricCups() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("metricCups"))
	return rv
}

// The metric cups unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856000-metriccups?language=objc
func UnitVolume_MetricCups() UnitVolume {
	return UnitVolumeClass.MetricCups()
}

// The cubic millimeters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856013-cubicmillimeters?language=objc
func (uc _UnitVolumeClass) CubicMillimeters() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("cubicMillimeters"))
	return rv
}

// The cubic millimeters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856013-cubicmillimeters?language=objc
func UnitVolume_CubicMillimeters() UnitVolume {
	return UnitVolumeClass.CubicMillimeters()
}

// The cubic meters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856029-cubicmeters?language=objc
func (uc _UnitVolumeClass) CubicMeters() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("cubicMeters"))
	return rv
}

// The cubic meters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856029-cubicmeters?language=objc
func UnitVolume_CubicMeters() UnitVolume {
	return UnitVolumeClass.CubicMeters()
}

// The bushels unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856025-bushels?language=objc
func (uc _UnitVolumeClass) Bushels() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("bushels"))
	return rv
}

// The bushels unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856025-bushels?language=objc
func UnitVolume_Bushels() UnitVolume {
	return UnitVolumeClass.Bushels()
}

// The centiliters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856040-centiliters?language=objc
func (uc _UnitVolumeClass) Centiliters() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("centiliters"))
	return rv
}

// The centiliters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856040-centiliters?language=objc
func UnitVolume_Centiliters() UnitVolume {
	return UnitVolumeClass.Centiliters()
}

// The pints unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856059-pints?language=objc
func (uc _UnitVolumeClass) Pints() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("pints"))
	return rv
}

// The pints unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856059-pints?language=objc
func UnitVolume_Pints() UnitVolume {
	return UnitVolumeClass.Pints()
}

// The deciliters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856091-deciliters?language=objc
func (uc _UnitVolumeClass) Deciliters() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("deciliters"))
	return rv
}

// The deciliters unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856091-deciliters?language=objc
func UnitVolume_Deciliters() UnitVolume {
	return UnitVolumeClass.Deciliters()
}

// The imperial fluid ounces unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856101-imperialfluidounces?language=objc
func (uc _UnitVolumeClass) ImperialFluidOunces() UnitVolume {
	rv := objc.Call[UnitVolume](uc, objc.Sel("imperialFluidOunces"))
	return rv
}

// The imperial fluid ounces unit of volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitvolume/1856101-imperialfluidounces?language=objc
func UnitVolume_ImperialFluidOunces() UnitVolume {
	return UnitVolumeClass.ImperialFluidOunces()
}
