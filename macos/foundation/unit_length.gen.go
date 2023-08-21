// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitLength] class.
var UnitLengthClass = _UnitLengthClass{objc.GetClass("NSUnitLength")}

type _UnitLengthClass struct {
	objc.Class
}

// An interface definition for the [UnitLength] class.
type IUnitLength interface {
	IDimension
}

// A unit of measure for length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength?language=objc
type UnitLength struct {
	Dimension
}

func UnitLengthFrom(ptr unsafe.Pointer) UnitLength {
	return UnitLength{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitLengthClass) Alloc() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("alloc"))
	return rv
}

func UnitLength_Alloc() UnitLength {
	return UnitLengthClass.Alloc()
}

func (uc _UnitLengthClass) New() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitLength() UnitLength {
	return UnitLengthClass.New()
}

func (u_ UnitLength) Init() UnitLength {
	rv := objc.Call[UnitLength](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitLengthClass) BaseUnit() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitLength_BaseUnit() UnitLength {
	return UnitLengthClass.BaseUnit()
}

func (u_ UnitLength) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitLength {
	rv := objc.Call[UnitLength](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func NewUnitLengthWithSymbolConverter(symbol string, converter IUnitConverter) UnitLength {
	instance := UnitLengthClass.Alloc().InitWithSymbolConverter(symbol, converter)
	instance.Autorelease()
	return instance
}

func (u_ UnitLength) InitWithSymbol(symbol string) UnitLength {
	rv := objc.Call[UnitLength](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func NewUnitLengthWithSymbol(symbol string) UnitLength {
	instance := UnitLengthClass.Alloc().InitWithSymbol(symbol)
	instance.Autorelease()
	return instance
}

// The parsecs unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856021-parsecs?language=objc
func (uc _UnitLengthClass) Parsecs() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("parsecs"))
	return rv
}

// The parsecs unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856021-parsecs?language=objc
func UnitLength_Parsecs() UnitLength {
	return UnitLengthClass.Parsecs()
}

// The astronomical units unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856087-astronomicalunits?language=objc
func (uc _UnitLengthClass) AstronomicalUnits() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("astronomicalUnits"))
	return rv
}

// The astronomical units unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856087-astronomicalunits?language=objc
func UnitLength_AstronomicalUnits() UnitLength {
	return UnitLengthClass.AstronomicalUnits()
}

// The inches unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856018-inches?language=objc
func (uc _UnitLengthClass) Inches() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("inches"))
	return rv
}

// The inches unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856018-inches?language=objc
func UnitLength_Inches() UnitLength {
	return UnitLengthClass.Inches()
}

// The nanometers unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856004-nanometers?language=objc
func (uc _UnitLengthClass) Nanometers() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("nanometers"))
	return rv
}

// The nanometers unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856004-nanometers?language=objc
func UnitLength_Nanometers() UnitLength {
	return UnitLengthClass.Nanometers()
}

// The hectometers unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1855970-hectometers?language=objc
func (uc _UnitLengthClass) Hectometers() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("hectometers"))
	return rv
}

// The hectometers unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1855970-hectometers?language=objc
func UnitLength_Hectometers() UnitLength {
	return UnitLengthClass.Hectometers()
}

// The megameters unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856036-megameters?language=objc
func (uc _UnitLengthClass) Megameters() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("megameters"))
	return rv
}

// The megameters unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856036-megameters?language=objc
func UnitLength_Megameters() UnitLength {
	return UnitLengthClass.Megameters()
}

// The millimeters unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856046-millimeters?language=objc
func (uc _UnitLengthClass) Millimeters() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("millimeters"))
	return rv
}

// The millimeters unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856046-millimeters?language=objc
func UnitLength_Millimeters() UnitLength {
	return UnitLengthClass.Millimeters()
}

// The feet unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1855972-feet?language=objc
func (uc _UnitLengthClass) Feet() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("feet"))
	return rv
}

// The feet unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1855972-feet?language=objc
func UnitLength_Feet() UnitLength {
	return UnitLengthClass.Feet()
}

// The furlongs unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856065-furlongs?language=objc
func (uc _UnitLengthClass) Furlongs() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("furlongs"))
	return rv
}

// The furlongs unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856065-furlongs?language=objc
func UnitLength_Furlongs() UnitLength {
	return UnitLengthClass.Furlongs()
}

// The fathoms unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856090-fathoms?language=objc
func (uc _UnitLengthClass) Fathoms() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("fathoms"))
	return rv
}

// The fathoms unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856090-fathoms?language=objc
func UnitLength_Fathoms() UnitLength {
	return UnitLengthClass.Fathoms()
}

// The light years unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1855974-lightyears?language=objc
func (uc _UnitLengthClass) Lightyears() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("lightyears"))
	return rv
}

// The light years unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1855974-lightyears?language=objc
func UnitLength_Lightyears() UnitLength {
	return UnitLengthClass.Lightyears()
}

// The nautical miles unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1855986-nauticalmiles?language=objc
func (uc _UnitLengthClass) NauticalMiles() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("nauticalMiles"))
	return rv
}

// The nautical miles unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1855986-nauticalmiles?language=objc
func UnitLength_NauticalMiles() UnitLength {
	return UnitLengthClass.NauticalMiles()
}

// The miles unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856016-miles?language=objc
func (uc _UnitLengthClass) Miles() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("miles"))
	return rv
}

// The miles unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856016-miles?language=objc
func UnitLength_Miles() UnitLength {
	return UnitLengthClass.Miles()
}

// The yards unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1855993-yards?language=objc
func (uc _UnitLengthClass) Yards() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("yards"))
	return rv
}

// The yards unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1855993-yards?language=objc
func UnitLength_Yards() UnitLength {
	return UnitLengthClass.Yards()
}

// The kilometers unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856106-kilometers?language=objc
func (uc _UnitLengthClass) Kilometers() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("kilometers"))
	return rv
}

// The kilometers unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856106-kilometers?language=objc
func UnitLength_Kilometers() UnitLength {
	return UnitLengthClass.Kilometers()
}

// The meters unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1855995-meters?language=objc
func (uc _UnitLengthClass) Meters() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("meters"))
	return rv
}

// The meters unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1855995-meters?language=objc
func UnitLength_Meters() UnitLength {
	return UnitLengthClass.Meters()
}

// The picometers unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856012-picometers?language=objc
func (uc _UnitLengthClass) Picometers() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("picometers"))
	return rv
}

// The picometers unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856012-picometers?language=objc
func UnitLength_Picometers() UnitLength {
	return UnitLengthClass.Picometers()
}

// The micrometers unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1855998-micrometers?language=objc
func (uc _UnitLengthClass) Micrometers() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("micrometers"))
	return rv
}

// The micrometers unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1855998-micrometers?language=objc
func UnitLength_Micrometers() UnitLength {
	return UnitLengthClass.Micrometers()
}

// The centimeters unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856082-centimeters?language=objc
func (uc _UnitLengthClass) Centimeters() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("centimeters"))
	return rv
}

// The centimeters unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856082-centimeters?language=objc
func UnitLength_Centimeters() UnitLength {
	return UnitLengthClass.Centimeters()
}

// The decimeters unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856007-decimeters?language=objc
func (uc _UnitLengthClass) Decimeters() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("decimeters"))
	return rv
}

// The decimeters unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856007-decimeters?language=objc
func UnitLength_Decimeters() UnitLength {
	return UnitLengthClass.Decimeters()
}

// The Scandinavian miles unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856061-scandinavianmiles?language=objc
func (uc _UnitLengthClass) ScandinavianMiles() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("scandinavianMiles"))
	return rv
}

// The Scandinavian miles unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856061-scandinavianmiles?language=objc
func UnitLength_ScandinavianMiles() UnitLength {
	return UnitLengthClass.ScandinavianMiles()
}

// The decameters unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856042-decameters?language=objc
func (uc _UnitLengthClass) Decameters() UnitLength {
	rv := objc.Call[UnitLength](uc, objc.Sel("decameters"))
	return rv
}

// The decameters unit of length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitlength/1856042-decameters?language=objc
func UnitLength_Decameters() UnitLength {
	return UnitLengthClass.Decameters()
}
