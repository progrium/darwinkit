// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MassFormatter] class.
var MassFormatterClass = _MassFormatterClass{objc.GetClass("NSMassFormatter")}

type _MassFormatterClass struct {
	objc.Class
}

// An interface definition for the [MassFormatter] class.
type IMassFormatter interface {
	IFormatter
	StringFromValueUnit(value float64, unit MassFormatterUnit) string
	StringFromKilograms(numberInKilograms float64) string
	UnitStringFromValueUnit(value float64, unit MassFormatterUnit) string
	UnitStringFromKilogramsUsedUnit(numberInKilograms float64, unitp *MassFormatterUnit) string
	UnitStyle() FormattingUnitStyle
	SetUnitStyle(value FormattingUnitStyle)
	IsForPersonMassUse() bool
	SetForPersonMassUse(value bool)
	NumberFormatter() NumberFormatter
	SetNumberFormatter(value INumberFormatter)
}

// A formatter that provides localized descriptions of mass and weight values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmassformatter?language=objc
type MassFormatter struct {
	Formatter
}

func MassFormatterFrom(ptr unsafe.Pointer) MassFormatter {
	return MassFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

func (mc _MassFormatterClass) Alloc() MassFormatter {
	rv := objc.Call[MassFormatter](mc, objc.Sel("alloc"))
	return rv
}

func MassFormatter_Alloc() MassFormatter {
	return MassFormatterClass.Alloc()
}

func (mc _MassFormatterClass) New() MassFormatter {
	rv := objc.Call[MassFormatter](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMassFormatter() MassFormatter {
	return MassFormatterClass.New()
}

func (m_ MassFormatter) Init() MassFormatter {
	rv := objc.Call[MassFormatter](m_, objc.Sel("init"))
	return rv
}

// Returns a properly formatted mass string for the given value and unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmassformatter/1409002-stringfromvalue?language=objc
func (m_ MassFormatter) StringFromValueUnit(value float64, unit MassFormatterUnit) string {
	rv := objc.Call[string](m_, objc.Sel("stringFromValue:unit:"), value, unit)
	return rv
}

// Returns a mass string for the provided value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmassformatter/1414324-stringfromkilograms?language=objc
func (m_ MassFormatter) StringFromKilograms(numberInKilograms float64) string {
	rv := objc.Call[string](m_, objc.Sel("stringFromKilograms:"), numberInKilograms)
	return rv
}

// Returns the unit string based on the provided value and unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmassformatter/1415491-unitstringfromvalue?language=objc
func (m_ MassFormatter) UnitStringFromValueUnit(value float64, unit MassFormatterUnit) string {
	rv := objc.Call[string](m_, objc.Sel("unitStringFromValue:unit:"), value, unit)
	return rv
}

// Returns the unit string for the provided value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmassformatter/1408475-unitstringfromkilograms?language=objc
func (m_ MassFormatter) UnitStringFromKilogramsUsedUnit(numberInKilograms float64, unitp *MassFormatterUnit) string {
	rv := objc.Call[string](m_, objc.Sel("unitStringFromKilograms:usedUnit:"), numberInKilograms, unitp)
	return rv
}

// The unit style used by this formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmassformatter/1411215-unitstyle?language=objc
func (m_ MassFormatter) UnitStyle() FormattingUnitStyle {
	rv := objc.Call[FormattingUnitStyle](m_, objc.Sel("unitStyle"))
	return rv
}

// The unit style used by this formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmassformatter/1411215-unitstyle?language=objc
func (m_ MassFormatter) SetUnitStyle(value FormattingUnitStyle) {
	objc.Call[objc.Void](m_, objc.Sel("setUnitStyle:"), value)
}

// A Boolean value that indicates whether the resulting string represents a person’s mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmassformatter/1407306-forpersonmassuse?language=objc
func (m_ MassFormatter) IsForPersonMassUse() bool {
	rv := objc.Call[bool](m_, objc.Sel("isForPersonMassUse"))
	return rv
}

// A Boolean value that indicates whether the resulting string represents a person’s mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmassformatter/1407306-forpersonmassuse?language=objc
func (m_ MassFormatter) SetForPersonMassUse(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setForPersonMassUse:"), value)
}

// The number formatter used to format the numbers in a mass strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmassformatter/1418462-numberformatter?language=objc
func (m_ MassFormatter) NumberFormatter() NumberFormatter {
	rv := objc.Call[NumberFormatter](m_, objc.Sel("numberFormatter"))
	return rv
}

// The number formatter used to format the numbers in a mass strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmassformatter/1418462-numberformatter?language=objc
func (m_ MassFormatter) SetNumberFormatter(value INumberFormatter) {
	objc.Call[objc.Void](m_, objc.Sel("setNumberFormatter:"), objc.Ptr(value))
}
