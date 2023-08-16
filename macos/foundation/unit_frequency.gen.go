// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitFrequency] class.
var UnitFrequencyClass = _UnitFrequencyClass{objc.GetClass("NSUnitFrequency")}

type _UnitFrequencyClass struct {
	objc.Class
}

// An interface definition for the [UnitFrequency] class.
type IUnitFrequency interface {
	IDimension
}

// A unit of measure for frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency?language=objc
type UnitFrequency struct {
	Dimension
}

func UnitFrequencyFrom(ptr unsafe.Pointer) UnitFrequency {
	return UnitFrequency{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitFrequencyClass) Alloc() UnitFrequency {
	rv := objc.Call[UnitFrequency](uc, objc.Sel("alloc"))
	return rv
}

func UnitFrequency_Alloc() UnitFrequency {
	return UnitFrequencyClass.Alloc()
}

func (uc _UnitFrequencyClass) New() UnitFrequency {
	rv := objc.Call[UnitFrequency](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitFrequency() UnitFrequency {
	return UnitFrequencyClass.New()
}

func (u_ UnitFrequency) Init() UnitFrequency {
	rv := objc.Call[UnitFrequency](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitFrequencyClass) BaseUnit() UnitFrequency {
	rv := objc.Call[UnitFrequency](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitFrequency_BaseUnit() UnitFrequency {
	return UnitFrequencyClass.BaseUnit()
}

func (u_ UnitFrequency) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitFrequency {
	rv := objc.Call[UnitFrequency](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func UnitFrequency_InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitFrequency {
	return UnitFrequencyClass.Alloc().InitWithSymbolConverter(symbol, converter)
}

func (u_ UnitFrequency) InitWithSymbol(symbol string) UnitFrequency {
	rv := objc.Call[UnitFrequency](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func UnitFrequency_InitWithSymbol(symbol string) UnitFrequency {
	return UnitFrequencyClass.Alloc().InitWithSymbol(symbol)
}

// The frames per second unit of frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency/3172542-framespersecond?language=objc
func (uc _UnitFrequencyClass) FramesPerSecond() UnitFrequency {
	rv := objc.Call[UnitFrequency](uc, objc.Sel("framesPerSecond"))
	return rv
}

// The frames per second unit of frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency/3172542-framespersecond?language=objc
func UnitFrequency_FramesPerSecond() UnitFrequency {
	return UnitFrequencyClass.FramesPerSecond()
}

// The terahertz unit of frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency/1690717-terahertz?language=objc
func (uc _UnitFrequencyClass) Terahertz() UnitFrequency {
	rv := objc.Call[UnitFrequency](uc, objc.Sel("terahertz"))
	return rv
}

// The terahertz unit of frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency/1690717-terahertz?language=objc
func UnitFrequency_Terahertz() UnitFrequency {
	return UnitFrequencyClass.Terahertz()
}

// The millihertz unit of frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency/1690882-millihertz?language=objc
func (uc _UnitFrequencyClass) Millihertz() UnitFrequency {
	rv := objc.Call[UnitFrequency](uc, objc.Sel("millihertz"))
	return rv
}

// The millihertz unit of frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency/1690882-millihertz?language=objc
func UnitFrequency_Millihertz() UnitFrequency {
	return UnitFrequencyClass.Millihertz()
}

// The hertz unit of frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency/1690670-hertz?language=objc
func (uc _UnitFrequencyClass) Hertz() UnitFrequency {
	rv := objc.Call[UnitFrequency](uc, objc.Sel("hertz"))
	return rv
}

// The hertz unit of frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency/1690670-hertz?language=objc
func UnitFrequency_Hertz() UnitFrequency {
	return UnitFrequencyClass.Hertz()
}

// The megahertz unit of frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency/1690707-megahertz?language=objc
func (uc _UnitFrequencyClass) Megahertz() UnitFrequency {
	rv := objc.Call[UnitFrequency](uc, objc.Sel("megahertz"))
	return rv
}

// The megahertz unit of frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency/1690707-megahertz?language=objc
func UnitFrequency_Megahertz() UnitFrequency {
	return UnitFrequencyClass.Megahertz()
}

// The nanohertz unit of frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency/1690663-nanohertz?language=objc
func (uc _UnitFrequencyClass) Nanohertz() UnitFrequency {
	rv := objc.Call[UnitFrequency](uc, objc.Sel("nanohertz"))
	return rv
}

// The nanohertz unit of frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency/1690663-nanohertz?language=objc
func UnitFrequency_Nanohertz() UnitFrequency {
	return UnitFrequencyClass.Nanohertz()
}

// The kilohertz unit of frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency/1690767-kilohertz?language=objc
func (uc _UnitFrequencyClass) Kilohertz() UnitFrequency {
	rv := objc.Call[UnitFrequency](uc, objc.Sel("kilohertz"))
	return rv
}

// The kilohertz unit of frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency/1690767-kilohertz?language=objc
func UnitFrequency_Kilohertz() UnitFrequency {
	return UnitFrequencyClass.Kilohertz()
}

// The gigahertz unit of frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency/1690667-gigahertz?language=objc
func (uc _UnitFrequencyClass) Gigahertz() UnitFrequency {
	rv := objc.Call[UnitFrequency](uc, objc.Sel("gigahertz"))
	return rv
}

// The gigahertz unit of frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency/1690667-gigahertz?language=objc
func UnitFrequency_Gigahertz() UnitFrequency {
	return UnitFrequencyClass.Gigahertz()
}

// The microhertz unit of frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency/1690838-microhertz?language=objc
func (uc _UnitFrequencyClass) Microhertz() UnitFrequency {
	rv := objc.Call[UnitFrequency](uc, objc.Sel("microhertz"))
	return rv
}

// The microhertz unit of frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitfrequency/1690838-microhertz?language=objc
func UnitFrequency_Microhertz() UnitFrequency {
	return UnitFrequencyClass.Microhertz()
}
