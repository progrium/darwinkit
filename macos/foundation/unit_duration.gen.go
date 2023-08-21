// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitDuration] class.
var UnitDurationClass = _UnitDurationClass{objc.GetClass("NSUnitDuration")}

type _UnitDurationClass struct {
	objc.Class
}

// An interface definition for the [UnitDuration] class.
type IUnitDuration interface {
	IDimension
}

// A unit of measure for a duration of time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitduration?language=objc
type UnitDuration struct {
	Dimension
}

func UnitDurationFrom(ptr unsafe.Pointer) UnitDuration {
	return UnitDuration{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitDurationClass) Alloc() UnitDuration {
	rv := objc.Call[UnitDuration](uc, objc.Sel("alloc"))
	return rv
}

func UnitDuration_Alloc() UnitDuration {
	return UnitDurationClass.Alloc()
}

func (uc _UnitDurationClass) New() UnitDuration {
	rv := objc.Call[UnitDuration](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitDuration() UnitDuration {
	return UnitDurationClass.New()
}

func (u_ UnitDuration) Init() UnitDuration {
	rv := objc.Call[UnitDuration](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitDurationClass) BaseUnit() UnitDuration {
	rv := objc.Call[UnitDuration](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitDuration_BaseUnit() UnitDuration {
	return UnitDurationClass.BaseUnit()
}

func (u_ UnitDuration) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitDuration {
	rv := objc.Call[UnitDuration](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func NewUnitDurationWithSymbolConverter(symbol string, converter IUnitConverter) UnitDuration {
	instance := UnitDurationClass.Alloc().InitWithSymbolConverter(symbol, converter)
	instance.Autorelease()
	return instance
}

func (u_ UnitDuration) InitWithSymbol(symbol string) UnitDuration {
	rv := objc.Call[UnitDuration](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func NewUnitDurationWithSymbol(symbol string) UnitDuration {
	instance := UnitDurationClass.Alloc().InitWithSymbol(symbol)
	instance.Autorelease()
	return instance
}

// The second unit of duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitduration/1856005-seconds?language=objc
func (uc _UnitDurationClass) Seconds() UnitDuration {
	rv := objc.Call[UnitDuration](uc, objc.Sel("seconds"))
	return rv
}

// The second unit of duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitduration/1856005-seconds?language=objc
func UnitDuration_Seconds() UnitDuration {
	return UnitDurationClass.Seconds()
}

// The microsecond unit of duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitduration/3172538-microseconds?language=objc
func (uc _UnitDurationClass) Microseconds() UnitDuration {
	rv := objc.Call[UnitDuration](uc, objc.Sel("microseconds"))
	return rv
}

// The microsecond unit of duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitduration/3172538-microseconds?language=objc
func UnitDuration_Microseconds() UnitDuration {
	return UnitDurationClass.Microseconds()
}

// The picosecond unit of duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitduration/3172541-picoseconds?language=objc
func (uc _UnitDurationClass) Picoseconds() UnitDuration {
	rv := objc.Call[UnitDuration](uc, objc.Sel("picoseconds"))
	return rv
}

// The picosecond unit of duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitduration/3172541-picoseconds?language=objc
func UnitDuration_Picoseconds() UnitDuration {
	return UnitDurationClass.Picoseconds()
}

// The hour unit of duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitduration/1855994-hours?language=objc
func (uc _UnitDurationClass) Hours() UnitDuration {
	rv := objc.Call[UnitDuration](uc, objc.Sel("hours"))
	return rv
}

// The hour unit of duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitduration/1855994-hours?language=objc
func UnitDuration_Hours() UnitDuration {
	return UnitDurationClass.Hours()
}

// The millisecond unit of duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitduration/3172539-milliseconds?language=objc
func (uc _UnitDurationClass) Milliseconds() UnitDuration {
	rv := objc.Call[UnitDuration](uc, objc.Sel("milliseconds"))
	return rv
}

// The millisecond unit of duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitduration/3172539-milliseconds?language=objc
func UnitDuration_Milliseconds() UnitDuration {
	return UnitDurationClass.Milliseconds()
}

// The nanosecond unit of duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitduration/3172540-nanoseconds?language=objc
func (uc _UnitDurationClass) Nanoseconds() UnitDuration {
	rv := objc.Call[UnitDuration](uc, objc.Sel("nanoseconds"))
	return rv
}

// The nanosecond unit of duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitduration/3172540-nanoseconds?language=objc
func UnitDuration_Nanoseconds() UnitDuration {
	return UnitDurationClass.Nanoseconds()
}

// The minute unit of duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitduration/1856088-minutes?language=objc
func (uc _UnitDurationClass) Minutes() UnitDuration {
	rv := objc.Call[UnitDuration](uc, objc.Sel("minutes"))
	return rv
}

// The minute unit of duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitduration/1856088-minutes?language=objc
func UnitDuration_Minutes() UnitDuration {
	return UnitDurationClass.Minutes()
}
