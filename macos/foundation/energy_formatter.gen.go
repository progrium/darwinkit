// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [EnergyFormatter] class.
var EnergyFormatterClass = _EnergyFormatterClass{objc.GetClass("NSEnergyFormatter")}

type _EnergyFormatterClass struct {
	objc.Class
}

// An interface definition for the [EnergyFormatter] class.
type IEnergyFormatter interface {
	IFormatter
	StringFromJoules(numberInJoules float64) string
	StringFromValueUnit(value float64, unit EnergyFormatterUnit) string
	UnitStringFromValueUnit(value float64, unit EnergyFormatterUnit) string
	UnitStringFromJoulesUsedUnit(numberInJoules float64, unitp *EnergyFormatterUnit) string
	UnitStyle() FormattingUnitStyle
	SetUnitStyle(value FormattingUnitStyle)
	NumberFormatter() NumberFormatter
	SetNumberFormatter(value INumberFormatter)
	IsForFoodEnergyUse() bool
	SetForFoodEnergyUse(value bool)
}

// A formatter that provides localized descriptions of energy values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsenergyformatter?language=objc
type EnergyFormatter struct {
	Formatter
}

func EnergyFormatterFrom(ptr unsafe.Pointer) EnergyFormatter {
	return EnergyFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

func (ec _EnergyFormatterClass) Alloc() EnergyFormatter {
	rv := objc.Call[EnergyFormatter](ec, objc.Sel("alloc"))
	return rv
}

func EnergyFormatter_Alloc() EnergyFormatter {
	return EnergyFormatterClass.Alloc()
}

func (ec _EnergyFormatterClass) New() EnergyFormatter {
	rv := objc.Call[EnergyFormatter](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewEnergyFormatter() EnergyFormatter {
	return EnergyFormatterClass.New()
}

func (e_ EnergyFormatter) Init() EnergyFormatter {
	rv := objc.Call[EnergyFormatter](e_, objc.Sel("init"))
	return rv
}

// Returns an energy string for the provided value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsenergyformatter/1409502-stringfromjoules?language=objc
func (e_ EnergyFormatter) StringFromJoules(numberInJoules float64) string {
	rv := objc.Call[string](e_, objc.Sel("stringFromJoules:"), numberInJoules)
	return rv
}

// Returns a properly formatted energy string for the given value and unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsenergyformatter/1411983-stringfromvalue?language=objc
func (e_ EnergyFormatter) StringFromValueUnit(value float64, unit EnergyFormatterUnit) string {
	rv := objc.Call[string](e_, objc.Sel("stringFromValue:unit:"), value, unit)
	return rv
}

// Returns the unit string based on the provided value and unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsenergyformatter/1411180-unitstringfromvalue?language=objc
func (e_ EnergyFormatter) UnitStringFromValueUnit(value float64, unit EnergyFormatterUnit) string {
	rv := objc.Call[string](e_, objc.Sel("unitStringFromValue:unit:"), value, unit)
	return rv
}

// Returns the unit string for the provided value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsenergyformatter/1408656-unitstringfromjoules?language=objc
func (e_ EnergyFormatter) UnitStringFromJoulesUsedUnit(numberInJoules float64, unitp *EnergyFormatterUnit) string {
	rv := objc.Call[string](e_, objc.Sel("unitStringFromJoules:usedUnit:"), numberInJoules, unitp)
	return rv
}

// The unit style used by this formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsenergyformatter/1414075-unitstyle?language=objc
func (e_ EnergyFormatter) UnitStyle() FormattingUnitStyle {
	rv := objc.Call[FormattingUnitStyle](e_, objc.Sel("unitStyle"))
	return rv
}

// The unit style used by this formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsenergyformatter/1414075-unitstyle?language=objc
func (e_ EnergyFormatter) SetUnitStyle(value FormattingUnitStyle) {
	objc.Call[objc.Void](e_, objc.Sel("setUnitStyle:"), value)
}

// The number formatter used to format the numbers in energy strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsenergyformatter/1412614-numberformatter?language=objc
func (e_ EnergyFormatter) NumberFormatter() NumberFormatter {
	rv := objc.Call[NumberFormatter](e_, objc.Sel("numberFormatter"))
	return rv
}

// The number formatter used to format the numbers in energy strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsenergyformatter/1412614-numberformatter?language=objc
func (e_ EnergyFormatter) SetNumberFormatter(value INumberFormatter) {
	objc.Call[objc.Void](e_, objc.Sel("setNumberFormatter:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the energy value is used to measure food energy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsenergyformatter/1410338-forfoodenergyuse?language=objc
func (e_ EnergyFormatter) IsForFoodEnergyUse() bool {
	rv := objc.Call[bool](e_, objc.Sel("isForFoodEnergyUse"))
	return rv
}

// A Boolean value that indicates whether the energy value is used to measure food energy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsenergyformatter/1410338-forfoodenergyuse?language=objc
func (e_ EnergyFormatter) SetForFoodEnergyUse(value bool) {
	objc.Call[objc.Void](e_, objc.Sel("setForFoodEnergyUse:"), value)
}
