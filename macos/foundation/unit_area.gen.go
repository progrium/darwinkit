// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitArea] class.
var UnitAreaClass = _UnitAreaClass{objc.GetClass("NSUnitArea")}

type _UnitAreaClass struct {
	objc.Class
}

// An interface definition for the [UnitArea] class.
type IUnitArea interface {
	IDimension
}

// A unit of measure for area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea?language=objc
type UnitArea struct {
	Dimension
}

func UnitAreaFrom(ptr unsafe.Pointer) UnitArea {
	return UnitArea{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitAreaClass) Alloc() UnitArea {
	rv := objc.Call[UnitArea](uc, objc.Sel("alloc"))
	return rv
}

func UnitArea_Alloc() UnitArea {
	return UnitAreaClass.Alloc()
}

func (uc _UnitAreaClass) New() UnitArea {
	rv := objc.Call[UnitArea](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitArea() UnitArea {
	return UnitAreaClass.New()
}

func (u_ UnitArea) Init() UnitArea {
	rv := objc.Call[UnitArea](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitAreaClass) BaseUnit() UnitArea {
	rv := objc.Call[UnitArea](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitArea_BaseUnit() UnitArea {
	return UnitAreaClass.BaseUnit()
}

func (u_ UnitArea) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitArea {
	rv := objc.Call[UnitArea](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func UnitArea_InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitArea {
	return UnitAreaClass.Alloc().InitWithSymbolConverter(symbol, converter)
}

func (u_ UnitArea) InitWithSymbol(symbol string) UnitArea {
	rv := objc.Call[UnitArea](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func UnitArea_InitWithSymbol(symbol string) UnitArea {
	return UnitAreaClass.Alloc().InitWithSymbol(symbol)
}

// The ares unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856064-ares?language=objc
func (uc _UnitAreaClass) Ares() UnitArea {
	rv := objc.Call[UnitArea](uc, objc.Sel("ares"))
	return rv
}

// The ares unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856064-ares?language=objc
func UnitArea_Ares() UnitArea {
	return UnitAreaClass.Ares()
}

// The square inches unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856067-squareinches?language=objc
func (uc _UnitAreaClass) SquareInches() UnitArea {
	rv := objc.Call[UnitArea](uc, objc.Sel("squareInches"))
	return rv
}

// The square inches unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856067-squareinches?language=objc
func UnitArea_SquareInches() UnitArea {
	return UnitAreaClass.SquareInches()
}

// The square feet unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856014-squarefeet?language=objc
func (uc _UnitAreaClass) SquareFeet() UnitArea {
	rv := objc.Call[UnitArea](uc, objc.Sel("squareFeet"))
	return rv
}

// The square feet unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856014-squarefeet?language=objc
func UnitArea_SquareFeet() UnitArea {
	return UnitAreaClass.SquareFeet()
}

// The square millimeters unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856039-squaremillimeters?language=objc
func (uc _UnitAreaClass) SquareMillimeters() UnitArea {
	rv := objc.Call[UnitArea](uc, objc.Sel("squareMillimeters"))
	return rv
}

// The square millimeters unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856039-squaremillimeters?language=objc
func UnitArea_SquareMillimeters() UnitArea {
	return UnitAreaClass.SquareMillimeters()
}

// The square kilometers unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856053-squarekilometers?language=objc
func (uc _UnitAreaClass) SquareKilometers() UnitArea {
	rv := objc.Call[UnitArea](uc, objc.Sel("squareKilometers"))
	return rv
}

// The square kilometers unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856053-squarekilometers?language=objc
func UnitArea_SquareKilometers() UnitArea {
	return UnitAreaClass.SquareKilometers()
}

// The hectares unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856066-hectares?language=objc
func (uc _UnitAreaClass) Hectares() UnitArea {
	rv := objc.Call[UnitArea](uc, objc.Sel("hectares"))
	return rv
}

// The hectares unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856066-hectares?language=objc
func UnitArea_Hectares() UnitArea {
	return UnitAreaClass.Hectares()
}

// The square miles unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856097-squaremiles?language=objc
func (uc _UnitAreaClass) SquareMiles() UnitArea {
	rv := objc.Call[UnitArea](uc, objc.Sel("squareMiles"))
	return rv
}

// The square miles unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856097-squaremiles?language=objc
func UnitArea_SquareMiles() UnitArea {
	return UnitAreaClass.SquareMiles()
}

// The square yards unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856070-squareyards?language=objc
func (uc _UnitAreaClass) SquareYards() UnitArea {
	rv := objc.Call[UnitArea](uc, objc.Sel("squareYards"))
	return rv
}

// The square yards unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856070-squareyards?language=objc
func UnitArea_SquareYards() UnitArea {
	return UnitAreaClass.SquareYards()
}

// The square nanometers unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856072-squarenanometers?language=objc
func (uc _UnitAreaClass) SquareNanometers() UnitArea {
	rv := objc.Call[UnitArea](uc, objc.Sel("squareNanometers"))
	return rv
}

// The square nanometers unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856072-squarenanometers?language=objc
func UnitArea_SquareNanometers() UnitArea {
	return UnitAreaClass.SquareNanometers()
}

// The square megameters unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856008-squaremegameters?language=objc
func (uc _UnitAreaClass) SquareMegameters() UnitArea {
	rv := objc.Call[UnitArea](uc, objc.Sel("squareMegameters"))
	return rv
}

// The square megameters unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856008-squaremegameters?language=objc
func UnitArea_SquareMegameters() UnitArea {
	return UnitAreaClass.SquareMegameters()
}

// The square meters unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1855985-squaremeters?language=objc
func (uc _UnitAreaClass) SquareMeters() UnitArea {
	rv := objc.Call[UnitArea](uc, objc.Sel("squareMeters"))
	return rv
}

// The square meters unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1855985-squaremeters?language=objc
func UnitArea_SquareMeters() UnitArea {
	return UnitAreaClass.SquareMeters()
}

// The square centimeters unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856030-squarecentimeters?language=objc
func (uc _UnitAreaClass) SquareCentimeters() UnitArea {
	rv := objc.Call[UnitArea](uc, objc.Sel("squareCentimeters"))
	return rv
}

// The square centimeters unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856030-squarecentimeters?language=objc
func UnitArea_SquareCentimeters() UnitArea {
	return UnitAreaClass.SquareCentimeters()
}

// The square micrometers unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856041-squaremicrometers?language=objc
func (uc _UnitAreaClass) SquareMicrometers() UnitArea {
	rv := objc.Call[UnitArea](uc, objc.Sel("squareMicrometers"))
	return rv
}

// The square micrometers unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856041-squaremicrometers?language=objc
func UnitArea_SquareMicrometers() UnitArea {
	return UnitAreaClass.SquareMicrometers()
}

// The acres unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856099-acres?language=objc
func (uc _UnitAreaClass) Acres() UnitArea {
	rv := objc.Call[UnitArea](uc, objc.Sel("acres"))
	return rv
}

// The acres unit of area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitarea/1856099-acres?language=objc
func UnitArea_Acres() UnitArea {
	return UnitAreaClass.Acres()
}
