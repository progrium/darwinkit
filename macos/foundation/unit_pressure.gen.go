// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [UnitPressure] class.
var UnitPressureClass = _UnitPressureClass{objc.GetClass("NSUnitPressure")}

type _UnitPressureClass struct {
	objc.Class
}

// An interface definition for the [UnitPressure] class.
type IUnitPressure interface {
	IDimension
}

// A unit of measure for pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure?language=objc
type UnitPressure struct {
	Dimension
}

func UnitPressureFrom(ptr unsafe.Pointer) UnitPressure {
	return UnitPressure{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitPressureClass) Alloc() UnitPressure {
	rv := objc.Call[UnitPressure](uc, objc.Sel("alloc"))
	return rv
}

func (uc _UnitPressureClass) New() UnitPressure {
	rv := objc.Call[UnitPressure](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitPressure() UnitPressure {
	return UnitPressureClass.New()
}

func (u_ UnitPressure) Init() UnitPressure {
	rv := objc.Call[UnitPressure](u_, objc.Sel("init"))
	return rv
}

func (u_ UnitPressure) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitPressure {
	rv := objc.Call[UnitPressure](u_, objc.Sel("initWithSymbol:converter:"), symbol, converter)
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func NewUnitPressureWithSymbolConverter(symbol string, converter IUnitConverter) UnitPressure {
	instance := UnitPressureClass.Alloc().InitWithSymbolConverter(symbol, converter)
	instance.Autorelease()
	return instance
}

func (uc _UnitPressureClass) BaseUnit() UnitPressure {
	rv := objc.Call[UnitPressure](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitPressure_BaseUnit() UnitPressure {
	return UnitPressureClass.BaseUnit()
}

func (u_ UnitPressure) InitWithSymbol(symbol string) UnitPressure {
	rv := objc.Call[UnitPressure](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func NewUnitPressureWithSymbol(symbol string) UnitPressure {
	instance := UnitPressureClass.Alloc().InitWithSymbol(symbol)
	instance.Autorelease()
	return instance
}

// The gigapascals unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1855983-gigapascals?language=objc
func (uc _UnitPressureClass) Gigapascals() UnitPressure {
	rv := objc.Call[UnitPressure](uc, objc.Sel("gigapascals"))
	return rv
}

// The gigapascals unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1855983-gigapascals?language=objc
func UnitPressure_Gigapascals() UnitPressure {
	return UnitPressureClass.Gigapascals()
}

// The bars unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1856109-bars?language=objc
func (uc _UnitPressureClass) Bars() UnitPressure {
	rv := objc.Call[UnitPressure](uc, objc.Sel("bars"))
	return rv
}

// The bars unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1856109-bars?language=objc
func UnitPressure_Bars() UnitPressure {
	return UnitPressureClass.Bars()
}

// The millibars unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1856093-millibars?language=objc
func (uc _UnitPressureClass) Millibars() UnitPressure {
	rv := objc.Call[UnitPressure](uc, objc.Sel("millibars"))
	return rv
}

// The millibars unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1856093-millibars?language=objc
func UnitPressure_Millibars() UnitPressure {
	return UnitPressureClass.Millibars()
}

// The millimeters of mercury unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1856052-millimetersofmercury?language=objc
func (uc _UnitPressureClass) MillimetersOfMercury() UnitPressure {
	rv := objc.Call[UnitPressure](uc, objc.Sel("millimetersOfMercury"))
	return rv
}

// The millimeters of mercury unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1856052-millimetersofmercury?language=objc
func UnitPressure_MillimetersOfMercury() UnitPressure {
	return UnitPressureClass.MillimetersOfMercury()
}

// The hectopascals unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1856111-hectopascals?language=objc
func (uc _UnitPressureClass) Hectopascals() UnitPressure {
	rv := objc.Call[UnitPressure](uc, objc.Sel("hectopascals"))
	return rv
}

// The hectopascals unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1856111-hectopascals?language=objc
func UnitPressure_Hectopascals() UnitPressure {
	return UnitPressureClass.Hectopascals()
}

// The kilopascals unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1856057-kilopascals?language=objc
func (uc _UnitPressureClass) Kilopascals() UnitPressure {
	rv := objc.Call[UnitPressure](uc, objc.Sel("kilopascals"))
	return rv
}

// The kilopascals unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1856057-kilopascals?language=objc
func UnitPressure_Kilopascals() UnitPressure {
	return UnitPressureClass.Kilopascals()
}

// The pounds per square inch unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1856077-poundsforcepersquareinch?language=objc
func (uc _UnitPressureClass) PoundsForcePerSquareInch() UnitPressure {
	rv := objc.Call[UnitPressure](uc, objc.Sel("poundsForcePerSquareInch"))
	return rv
}

// The pounds per square inch unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1856077-poundsforcepersquareinch?language=objc
func UnitPressure_PoundsForcePerSquareInch() UnitPressure {
	return UnitPressureClass.PoundsForcePerSquareInch()
}

// The newtons per square meter unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1856096-newtonspermeterssquared?language=objc
func (uc _UnitPressureClass) NewtonsPerMetersSquared() UnitPressure {
	rv := objc.Call[UnitPressure](uc, objc.Sel("newtonsPerMetersSquared"))
	return rv
}

// The newtons per square meter unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1856096-newtonspermeterssquared?language=objc
func UnitPressure_NewtonsPerMetersSquared() UnitPressure {
	return UnitPressureClass.NewtonsPerMetersSquared()
}

// The inches of mercury unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1856074-inchesofmercury?language=objc
func (uc _UnitPressureClass) InchesOfMercury() UnitPressure {
	rv := objc.Call[UnitPressure](uc, objc.Sel("inchesOfMercury"))
	return rv
}

// The inches of mercury unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1856074-inchesofmercury?language=objc
func UnitPressure_InchesOfMercury() UnitPressure {
	return UnitPressureClass.InchesOfMercury()
}

// The megapascals unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1856115-megapascals?language=objc
func (uc _UnitPressureClass) Megapascals() UnitPressure {
	rv := objc.Call[UnitPressure](uc, objc.Sel("megapascals"))
	return rv
}

// The megapascals unit of pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitpressure/1856115-megapascals?language=objc
func UnitPressure_Megapascals() UnitPressure {
	return UnitPressureClass.Megapascals()
}
